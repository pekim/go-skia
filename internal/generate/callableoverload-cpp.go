package generate

import (
	"fmt"
	"strings"
)

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
