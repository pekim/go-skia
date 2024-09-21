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

	// Gather the arguments in a to string.
	var args []string
	for _, param := range o.Params {
		if param.ValueNil {
			// no more params
			break
		}
		args = append(args, param.cppArg)
	}
	o.cppArgs = strings.Join(args, ", ")

	// Smart pointers are released.
	skSpRelease := ""
	if o.retrn.isSmartPointer {
		skSpRelease = ".release()"
	}

	// If returning something, assign it to a variable.
	if !o.retrn.isVoid {
		f.writeln("  auto ret =")
	}

	// Generate the function call.
	if o.isStatic {
		// Call a class static member function.
		f.writelnf("  %s::%s(%s)%s;", o.record.CppName, o.cppName, o.cppArgs, skSpRelease)
	} else if o.record != nil {
		// Call a class member function.
		f.writelnf("  reinterpret_cast<%s*>(c_obj)->%s(%s)%s;",
			o.record.CppName,
			o.cppName,
			o.cppArgs,
			skSpRelease,
		)
	} else {
		// Call a plain (non-class) function.
		f.writelnf("  %s(%s)%s;", o.cppName, o.cppArgs, skSpRelease)
	}

	// Convert the return value (in 'ret' var), and return it.
	o.generateCppReturn(g)
}

func (o *callableOverload) generateCppReturn(g generator) {
	returnConst := ""
	constCastStart := ""
	constCastEnd := ""
	if o.retrn.isConst {
		returnConst = "const"
		constCastStart = fmt.Sprintf("const_cast<%s *>(", o.retrn.record.cStructName)
		constCastEnd = ")"
	}

	var returnValue string
	if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
		returnValue = fmt.Sprintf("reinterpret_cast<%s *> (ret)", o.retrn.subTyp.cName)
	} else if o.retrn.isPointer || o.retrn.isSmartPointer && o.retrn.subTyp != nil && o.retrn.subTyp.record != nil {
		returnValue = fmt.Sprintf("reinterpret_cast<%s *> (ret)", o.retrn.subTyp.record.cStructName)
	} else if o.retrn.isPointer || o.retrn.isSmartPointer {
		if o.retrn.isConst {
			returnValue = fmt.Sprintf("%sreinterpret_cast<%s %s *> (ret)%s",
				constCastStart,
				returnConst,
				o.retrn.record.cStructName,
				constCastEnd)
		} else {
			returnValue = fmt.Sprintf("reinterpret_cast<%s *> (ret)", o.retrn.record.cStructName)
		}
	} else if o.retrn.record != nil {
		if o.record != nil {
			returnValue = fmt.Sprintf("*(reinterpret_cast<%s *> (&ret))", o.retrn.record.cStructName)
		} else {
			returnValue = fmt.Sprintf("*(reinterpret_cast<%s %s *> (&ret))",
				returnConst,
				o.retrn.record.cStructName,
			)
		}
	} else if o.retrn.isVoid {
		returnValue = ""
	} else {
		returnValue = "ret"
	}

	f := g.cppFile
	f.writelnf("  return %s;", returnValue)
}
