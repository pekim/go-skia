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
