package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type callableOverload struct {
	Suffix            string `json:"suffix"`
	cppName           string
	goFuncName        string
	cFuncName         string
	record            *record
	doc               string
	isStatic          bool
	isNonStaticMethod bool    // Go function is a method, and C function has a receiver first parameter
	Params            []param `json:"params"`
	cParamsDecl       string
	resultType        clang.Type
	retrn             typ
}

func (o *callableOverload) enrich1(callable *callable, record *record, cursor clang.Cursor) {
	o.cppName = callable.CppName
	o.record = record
	o.isStatic = cursor.CXXMethod_IsStatic()
	o.isNonStaticMethod = o.record != nil && !o.isStatic
	if o.isStatic {
		o.goFuncName = fmt.Sprintf("%s%s%s", record.goName, o.cppName, o.Suffix)
	} else {
		o.goFuncName = fmt.Sprintf("%s%s", goExportedName(o.cppName), o.Suffix)
	}
	if o.record != nil {
		o.cFuncName = fmt.Sprintf("misk_%s_%s%s", o.record.goName, o.cppName, o.Suffix)
	} else {
		cName := strings.ReplaceAll(o.cppName, "::", "") // remove any "::" in the name
		o.cFuncName = fmt.Sprintf("misk_%s%s", cName, o.Suffix)
	}
	o.doc = cursor.RawCommentText()
	o.resultType = cursor.ResultType()

	paramCount := int(cursor.NumArguments())
	if len(o.Params) > 0 {
		if len(o.Params) > 0 && len(o.Params) != paramCount {
			fatalf("function %s, expected %d params but have %d", o.cppName, len(o.Params), paramCount)
		}
		// Do not make o.Params, as it's already present from unmarshalled JSON.
	} else {
		o.Params = make([]param, paramCount)
	}
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg, o.Params[i].ValueNil)
		o.Params[i] = param
	}
}

func (o *callableOverload) enrich2(api api) {
	for i := range o.Params {
		param := &o.Params[i]
		param.enrich2(api)
	}
	o.setCParamsDecl()
	o.retrn = mustTypFromClangType(o.resultType, api)
}

func (o *callableOverload) setCParamsDecl() {
	var params []string
	if o.isNonStaticMethod {
		params = append(params, fmt.Sprintf("%s *c_obj", o.record.cStructName))
	}
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		params = append(params, param.cParam)
	}
	o.cParamsDecl = strings.Join(params, ", ")
}

func (o callableOverload) generate(g generator) {
	o.generateGo(g)
	o.generateHeader(g)
	o.generateCpp(g)
}

func (o callableOverload) writeGoBody(g generator) {
	f := g.goFile

	var cVars []string
	var cArgs []string
	firstCVar, firstCArg, haveExtraCArg := o.firstGoCArg()
	if haveExtraCArg {
		cVars = append(cVars, firstCVar)
		cArgs = append(cArgs, firstCArg)
	}
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		cVars = append(cVars, param.cgoVar)
		cArgs = append(cArgs, param.cName)
	}

	// build the call to the function/method
	call := fmt.Sprintf("C.%s(%s)", o.cFuncName, strings.Join(cArgs, ", "))

	// write the creation of the vars that will be passed as args to the C function
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
		} else if o.retrn.typedef != nil {
			f.writelnf("  return %s(retC)", o.retrn.goName)
		} else {
			fatalf("return type '%s' not supported", o.retrn.goName)
		}
	}
}

// firstCArg provides an (extra) first C arg if the overload is a non-static
// class function.
func (o callableOverload) firstGoCArg() (string, string, bool) {
	if !o.isNonStaticMethod {
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

	// make function's params declarations
	var params []string
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		params = append(params, fmt.Sprintf("%s %s", param.goName, param.typGoName))
	}

	// make return declaration
	var returnDecl string
	if !o.retrn.isVoid {
		returnDecl = o.retrn.goName
	}
	// make receiver if a non-static class method
	var receiver string
	if o.isNonStaticMethod {
		receiver = fmt.Sprintf("(o %s)", o.record.goName)
	}

	f.writeDocComment(o.doc)
	f.writelnf("func %s %s(%s) %s {", receiver, o.goFuncName, strings.Join(params, ", "), returnDecl)
	o.writeGoBody(g)
	f.writeln("}")
	f.writeln()
}

func (o callableOverload) generateHeader(g generator) {
	f := g.headerFile

	returnDecl := o.retrn.cppName
	returnPtr := ""
	if o.retrn.enum != nil {
		returnDecl = o.retrn.enum.cType.cName
	} else if o.retrn.typedef != nil {
		returnDecl = o.retrn.typedef.cName
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

	f.writelnf("%s %s(%s);", returnDecl, o.cFuncName, o.cParamsDecl)
}

func (o callableOverload) generateCpp(g generator) {
	f := g.cppFile

	var args []string
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		args = append(args, param.cppArg)
	}
	allArgs := strings.Join(args, ", ")

	returnDecl := o.retrn.cppName
	returnPtr := ""
	if o.retrn.enum != nil {
		returnDecl = o.retrn.enum.cType.cName
	} else if o.retrn.typedef != nil {
		returnDecl = o.retrn.typedef.cName
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

	f.writelnf("%s %s(%s) {", returnDecl, o.cFuncName, o.cParamsDecl)
	if o.retrn.record != nil {
		if o.isStatic {
			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  return reinterpret_cast<%s *> (%s::%s(%s)%s);",
					o.retrn.record.cStructName,
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
			} else {
				f.writelnf("  auto ret = (%s::%s(%s)%s);",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					o.retrn.record.cStructName,
				)
			}
		} else if o.record != nil {
			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.record.cStructName,
				)
			} else {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s *> (&ret));",
					o.retrn.record.cStructName,
				)
			}
		} else {
			var returnConst string
			var constCastStart string
			var constCastEnd string
			if o.retrn.isConst {
				returnConst = "const"
				constCastStart = fmt.Sprintf("const_cast<%s *>(", o.retrn.record.cStructName)
				constCastEnd = ")"
			}

			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  return %sreinterpret_cast<%s %s *> (%s(%s)%s)%s;",
					constCastStart,
					returnConst,
					o.retrn.record.cStructName,
					o.cppName,
					allArgs,
					skSpRelease,
					constCastEnd)
			} else {
				f.writelnf("  auto ret = (%s(%s)%s);",
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s %s *> (&ret));",
					returnConst,
					o.retrn.record.cStructName,
				)
			}
		}
	} else {
		if o.isStatic {
			f.writelnf("  return %s::%s(%s)%s;", o.record.CppName, o.cppName, allArgs, skSpRelease)
		} else if o.record != nil {
			if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.cName,
				)
			} else if o.retrn.isPointer || o.retrn.isSmartPointer && o.retrn.subTyp.record != nil {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.record.cStructName,
				)
			} else {
				f.writelnf("  return reinterpret_cast<%s*>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					allArgs,
					skSpRelease,
				)
			}
		} else {
			var returnConst string
			var constCastStart string
			var constCastEnd string
			if o.retrn.isConst {
				returnConst = "const"
				constCastStart = fmt.Sprintf("const_cast<%s *>(", o.retrn.record.cStructName)
				constCastEnd = ")"
			}

			if o.retrn.record != nil {
				if o.retrn.isPointer || o.retrn.isSmartPointer {
					f.writelnf("  return %sreinterpret_cast<%s %s *> (%s(%s)%s)%s;",
						constCastStart,
						returnConst,
						o.retrn.record.cStructName,
						o.cppName,
						allArgs,
						skSpRelease,
						constCastEnd)
				} else {
					f.writelnf("  auto ret = (%s(%s)%s);",
						o.cppName,
						allArgs,
						skSpRelease)
					f.writelnf("  return *(reinterpret_cast<%s %s *> (&ret));",
						returnConst,
						o.retrn.record.cStructName,
					)
				}
			} else {
				f.writelnf("  return %s(%s)%s;", o.cppName, allArgs, skSpRelease)
			}
		}
	}
	f.writeln("}")
	f.writeln()
}
