package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type callableOverload struct {
	Suffix     string `json:"suffix"`
	cppName    string
	goFuncName string
	cFuncName  string
	record     *record
	doc        string
	isStatic   bool
	params     []param
	resultType clang.Type
	retrn      typ
}

type callable struct {
	CppName       string              `json:"name"`
	Overloads     []*callableOverload `json:"overloads"`
	enrichedCount int
}

// enrich1 enriches one of the callable's methods.
func (c *callable) enrich1(record *record, cursor clang.Cursor) {
	var overload *callableOverload
	if len(c.Overloads) == 0 {
		// If no overloads were configured in JSON it implies that there is just one.
		c.Overloads = []*callableOverload{{}}
	}
	if c.enrichedCount >= len(c.Overloads) {
		if record != nil {
			fatalf("record %s, method %s, %d overloads configured, but more encountered",
				c.Overloads[0].record.CppName, c.CppName, len(c.Overloads))
		} else {
			fatalf("function %s, %d overloads configured, but more encountered", c.CppName, len(c.Overloads))
		}
	}

	// Overloads configured in JSON are expected to be in the same order as they are encountered
	// in the AST for a translation unit.
	overload = c.Overloads[c.enrichedCount]
	if overload == nil {
		// The overload was configured as 'null' in JSON. That indicates that is it not to be
		// generated. So don't enrich it, skip it.
		c.enrichedCount++
		return
	}

	overload.cppName = c.CppName
	overload.record = record
	overload.isStatic = cursor.CXXMethod_IsStatic()
	if overload.isStatic {
		overload.goFuncName = fmt.Sprintf("%s%s%s", record.goName, c.CppName, overload.Suffix)
	} else {
		overload.goFuncName = fmt.Sprintf("%s%s", goExportedName(c.CppName), overload.Suffix)
	}
	overload.cFuncName = fmt.Sprintf("misk_%s_%s%s", record.goName, c.CppName, overload.Suffix)
	overload.doc = cursor.RawCommentText()
	overload.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	overload.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg)
		overload.params[i] = param
	}

	c.enrichedCount++
}

func (c *callable) enrich2(api api) {
	for i := range c.Overloads {
		overload := c.Overloads[i]
		if overload == nil {
			continue
		}
		for i := range overload.params {
			param := &overload.params[i]
			param.enrich2(api)
		}
		overload.retrn = mustTypFromClangType(overload.resultType, api)
	}
}

func (c callable) generate(g generator) {
	if len(c.Overloads) == 0 {
		fatalf("method %s, 0 overloads configured, and none enriched", c.CppName)
	}
	if c.enrichedCount < len(c.Overloads) {
		fatalf("record %s, method %s, only %d of %d overloads enriched",
			c.Overloads[0].record.CppName, c.CppName, c.enrichedCount, len(c.Overloads))
	}

	for _, method := range c.Overloads {
		if method != nil {
			method.generate(g)
		}
	}
}

func (o callableOverload) generate(g generator) {
	o.generateGo(g)
	o.generateHeader(g)
	o.generateCpp(g)
}

func (o callableOverload) generateGo(g generator) {
	f := g.goFile

	argsCount := len(o.params)
	argsOffset := 0
	if !o.isStatic {
		argsCount++
		argsOffset = 1
	}
	params := make([]string, len(o.params))
	cVars := make([]string, argsCount)
	cArgs := make([]string, argsCount)
	if !o.isStatic {
		if o.record.NoWrapper {
			cVars[0] = fmt.Sprintf("c_obj := (*C.%s)(&o)", o.record.cStructName)
		} else {
			cVars[0] = "c_obj := o.sk"
		}
		cArgs[0] = "c_obj"
	}
	for i, param := range o.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typGoName)
		cVars[argsOffset+i] = param.cgoVar
		cArgs[argsOffset+i] = param.cName
	}

	var returnDecl string
	if !o.retrn.isVoid {
		returnDecl = o.retrn.goName
	}

	var receiver string
	if !o.isStatic {
		receiver = fmt.Sprintf("(o %s)", o.record.goName)
	}
	call := fmt.Sprintf("C.%s(%s)", o.cFuncName, strings.Join(cArgs, ", "))

	f.writeDocComment(o.doc)
	f.writelnf("func %s %s(%s) %s {", receiver, o.goFuncName, strings.Join(params, ", "), returnDecl)
	if len(cVars) > 0 {
		f.writeln(strings.Join(cVars, "\n"))
	}
	if o.retrn.isVoid {
		f.writelnf("  %s", call)
	} else {
		f.writelnf("  retC := %s", call)
		if o.retrn.isPrimitive {
			f.writelnf("  return %s(retC)", o.retrn.goName)
		} else if o.retrn.enum != nil {
			f.writelnf("  return %s(retC)", o.retrn.enum.goName)
		} else if o.retrn.record != nil {
			if o.retrn.record.NoWrapper {
				f.writelnf("  return %s(retC)", o.retrn.record.goName)
			} else if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  return %s{sk: retC}", o.retrn.record.goName)
			} else {
				f.writelnf("  return %s{sk: &retC}", o.retrn.record.goName)
			}
		} else if (o.retrn.isPointer || o.retrn.isSmartPointer) && o.retrn.subTyp.record != nil {
			f.writelnf("  return %s{sk: retC}", o.retrn.subTyp.record.goName)
		} else if o.retrn.goName == "string" {
			f.writelnf("  return C.GoString( retC)")
		} else {
			fatalf("return type '%s' not supported", o.retrn.goName)
		}
	}
	f.writeln("}")
	f.writeln()
}

func (o callableOverload) generateHeader(g generator) {
	f := g.headerFile

	paramsCount := len(o.params)
	paramOffset := 0
	if !o.isStatic {
		paramsCount++
		paramOffset = 1
	}
	params := make([]string, paramsCount)
	if !o.isStatic {
		params[0] = fmt.Sprintf("%s *c_obj", o.record.cStructName)
	}
	for i, param := range o.params {
		params[paramOffset+i] = param.cParam
	}

	returnDecl := o.retrn.cppName
	returnPtr := ""
	if o.retrn.enum != nil {
		returnDecl = o.retrn.enum.cType.cName
	} else if o.retrn.record != nil {
		if o.retrn.isPointer || o.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", o.retrn.record.cStructName, returnPtr)
	} else if (o.retrn.isPointer || o.retrn.isSmartPointer) && o.retrn.subTyp.record != nil {
		returnPtr = "*"
		returnDecl = fmt.Sprintf("%s%s", o.retrn.subTyp.record.cStructName, returnPtr)
	} else if o.retrn.goName == "string" {
		returnDecl = "char*"
	}

	f.writelnf("%s %s(%s);", returnDecl, o.cFuncName, strings.Join(params, ", "))
}

func (o callableOverload) generateCpp(g generator) {
	f := g.cppFile

	paramsCount := len(o.params)
	paramOffset := 0
	if !o.isStatic {
		paramsCount++
		paramOffset = 1
	}
	params := make([]string, paramsCount)
	args := make([]string, len(o.params))
	if !o.isStatic {
		params[0] = fmt.Sprintf("%s *c_obj", o.record.cStructName)
	}
	for i, param := range o.params {
		params[paramOffset+i] = param.cParam
		args[i] = param.cppArg
	}

	returnDecl := o.retrn.cppName
	returnPtr := ""
	if o.retrn.enum != nil {
		returnDecl = o.retrn.enum.cType.cName
	} else if o.retrn.record != nil {
		if o.retrn.isPointer || o.retrn.isSmartPointer {
			returnPtr = "*"
		}
		returnDecl = fmt.Sprintf("%s%s", o.retrn.record.cStructName, returnPtr)
	} else if (o.retrn.isPointer || o.retrn.isSmartPointer) && o.retrn.subTyp.record != nil {
		returnPtr = "*"
		returnDecl = fmt.Sprintf("%s%s", o.retrn.subTyp.record.cStructName, returnPtr)
	}

	skSpRelease := ""
	if o.retrn.isSmartPointer {
		skSpRelease = ".release()"
	}

	f.writelnf("%s %s(%s) {", returnDecl, o.cFuncName, strings.Join(params, ", "))
	if o.retrn.record != nil {
		if o.isStatic {
			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  return reinterpret_cast<%s *> (%s::%s(%s)%s);",
					o.retrn.record.cStructName,
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
			} else {
				f.writelnf("  auto ret = (%s::%s(%s)%s);",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					o.retrn.record.cStructName,
				)
			}
		} else {
			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.record.cStructName,
				)
			} else {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					o.retrn.record.cStructName,
				)
			}
		}
	} else {
		if o.isStatic {
			f.writelnf("  return %s::%s(%s)%s;", o.record.CppName, o.cppName, strings.Join(args, ", "), skSpRelease)
		} else {
			if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.cName,
				)
			} else if o.retrn.isPointer || o.retrn.isSmartPointer && o.retrn.subTyp.record != nil {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.record.cStructName,
				)
			} else {
				f.writelnf("  return reinterpret_cast<%s*>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					strings.Join(args, ", "),
					skSpRelease,
				)
			}
		}
	}
	f.writeln("}")
	f.writeln()
}
