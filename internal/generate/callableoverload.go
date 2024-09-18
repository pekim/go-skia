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

func (o *callableOverload) enrich1(callable *callable, record *record, cursor clang.Cursor) {
	o.cppName = callable.CppName
	o.record = record
	o.isStatic = cursor.CXXMethod_IsStatic()
	if o.isStatic {
		o.goFuncName = fmt.Sprintf("%s%s%s", record.goName, o.cppName, o.Suffix)
	} else {
		o.goFuncName = fmt.Sprintf("%s%s", goExportedName(o.cppName), o.Suffix)
	}
	o.cFuncName = fmt.Sprintf("misk_%s_%s%s", record.goName, o.cppName, o.Suffix)
	o.doc = cursor.RawCommentText()
	o.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	o.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg)
		o.params[i] = param
	}
}

func (o *callableOverload) enrich2(api api) {
	for i := range o.params {
		param := &o.params[i]
		param.enrich2(api)
	}
	o.retrn = mustTypFromClangType(o.resultType, api)
}

func (o callableOverload) generate(g generator) {
	o.generateGo(g)
	o.generateHeader(g)
	o.generateCpp(g)
}

// firstCArg provides an (extra) first C arg if the overload is a non-static
// class function.
func (o callableOverload) firstCArg() (string, string, bool) {
	if o.record == nil || o.isStatic {
		return "", "", false
	}

	cVar := "c_obj := o.sk"
	if o.record.NoWrapper {
		cVar = fmt.Sprintf("c_obj := (*C.%s)(&o)", o.record.cStructName)
	}
	cArg := "c_obj"
	return cVar, cArg, true
}

func (o callableOverload) generateGo(g generator) {
	f := g.goFile
	var params []string
	var cVars []string
	var cArgs []string

	firstCVar, firstCArg, haveExtraCArg := o.firstCArg()
	if haveExtraCArg {
		cVars = append(cVars, firstCVar)
		cArgs = append(cArgs, firstCArg)
	}
	for _, param := range o.params {
		params = append(params, fmt.Sprintf("%s %s", param.goName, param.typGoName))
		cVars = append(cVars, param.cgoVar)
		cArgs = append(cArgs, param.cName)
	}

	// make return declaration
	var returnDecl string
	if !o.retrn.isVoid {
		returnDecl = o.retrn.goName
	}

	// build the call to the function/method
	var receiver string
	if o.record != nil && !o.isStatic {
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
