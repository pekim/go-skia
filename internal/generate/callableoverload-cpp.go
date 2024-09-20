package generate

import (
	"fmt"
	"strings"
)

func (o callableOverload) generateCpp(g generator) {
	f := g.cppFile

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

	f.writelnf("%s %s(%s) {", returnDecl, o.cFuncName, o.cParamsDecl)
	o.generateCppBody(g)
	f.writeln("}")
	f.writeln()
}

func (o callableOverload) generateCppBody(g generator) {
	f := g.cppFile

	var args []string
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		args = append(args, param.cppArg)
	}
	o.cppArgs = strings.Join(args, ", ")

	skSpRelease := ""
	if o.retrn.isSmartPointer {
		skSpRelease = ".release()"
	}

	if o.retrn.record != nil {
		if o.isStatic {
			if o.retrn.isPointer || o.retrn.isSmartPointer {
				f.writelnf("  auto ret = reinterpret_cast<%s *> (%s::%s(%s)%s);",
					o.retrn.record.cStructName,
					o.record.CppName,
					o.cppName,
					o.cppArgs,
					skSpRelease)
				f.writeln("  return ret;")
			} else {
				f.writelnf("  auto ret = (%s::%s(%s)%s);",
					o.record.CppName,
					o.cppName,
					o.cppArgs,
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
					o.cppArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.record.cStructName,
				)
			} else {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					o.cppArgs,
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
				f.writelnf("  auto ret = %sreinterpret_cast<%s %s *> (%s(%s)%s)%s;",
					constCastStart,
					returnConst,
					o.retrn.record.cStructName,
					o.cppName,
					o.cppArgs,
					skSpRelease,
					constCastEnd)
				f.writeln("  return ret;")
			} else {
				f.writelnf("  auto ret = (%s(%s)%s);",
					o.cppName,
					o.cppArgs,
					skSpRelease)
				f.writelnf("  return *(reinterpret_cast<%s %s *> (&ret));",
					returnConst,
					o.retrn.record.cStructName,
				)
			}
		}
	} else {
		if o.isStatic {
			f.writelnf("  return %s::%s(%s)%s;", o.record.CppName, o.cppName, o.cppArgs, skSpRelease)
		} else if o.record != nil {
			if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					o.cppArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.cName,
				)
			} else if o.retrn.isPointer || o.retrn.isSmartPointer && o.retrn.subTyp.record != nil {
				f.writelnf("  auto ret = reinterpret_cast<%s *>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					o.cppArgs,
					skSpRelease)
				f.writelnf("  return (reinterpret_cast<%s *> (ret));",
					o.retrn.subTyp.record.cStructName,
				)
			} else {
				f.writelnf("  return reinterpret_cast<%s*>(c_obj)->%s(%s)%s;",
					o.record.CppName,
					o.cppName,
					o.cppArgs,
					skSpRelease,
				)
				// f.writeln("  return ret;")
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
					f.writelnf("  auto ret = %sreinterpret_cast<%s %s *> (%s(%s)%s)%s;",
						constCastStart,
						returnConst,
						o.retrn.record.cStructName,
						o.cppName,
						o.cppArgs,
						skSpRelease,
						constCastEnd)
					f.writeln("  return ret;")
				} else {
					f.writelnf("  auto ret = (%s(%s)%s);",
						o.cppName,
						o.cppArgs,
						skSpRelease)
					f.writelnf("  return *(reinterpret_cast<%s %s *> (&ret));",
						returnConst,
						o.retrn.record.cStructName,
					)
				}
			} else {
				f.writelnf("  return %s(%s)%s;", o.cppName, o.cppArgs, skSpRelease)
			}
		}
	}
}
