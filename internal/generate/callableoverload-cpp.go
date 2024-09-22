package generate

import (
	"fmt"
	"strings"
)

func (o callableOverload) generateCpp(g generator) {
	f := g.cppFile

	returnDecl := o.retrn.cppName
	// returnPtr := ""
	if o.retrn.enum != nil {
		returnDecl = o.retrn.enum.cType.cName
	} else if o.retrn.typedef != nil {
		returnDecl = o.retrn.typedef.cName
	} else if o.retrn.record != nil {
		returnDecl = o.retrn.record.cStructName
		if o.retrn.isPointer || o.retrn.isSmartPointer {
			returnDecl += "*"
		}
	} else if (o.retrn.isPointer || o.retrn.isSmartPointer) && o.retrn.subTyp.record != nil {
		returnDecl = o.retrn.subTyp.record.cStructName + "*"
	}

	f.writelnf("%s %s(%s) {", returnDecl, o.cFuncName, o.cParamsDecl)
	o.generateCppApiCallStatement(g)
	o.generateCppReturnStatement(g)
	f.writeln("}")
	f.writeln()
}

func (o callableOverload) generateCppApiCallStatement(g generator) {
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

	// If returning something, assign it to a variable.
	if !o.retrn.isVoid {
		f.write("  auto ret = ")
	}

	// Generate the function call.
	if o.isStatic {
		// Call a class static member function.
		f.writef("%s::%s(%s)", o.record.CppName, o.cppName, o.cppArgs)
	} else if o.record != nil {
		// Call a class member function.
		f.writef("reinterpret_cast<%s*>(c_obj)->%s(%s)", o.record.CppName, o.cppName, o.cppArgs)
	} else {
		// Call a plain (non-class) function.
		f.writef("%s(%s)", o.cppName, o.cppArgs)
	}

	if o.retrn.isSmartPointer {
		// Release smart pointers.
		f.write(".release()")
	}
	f.writeln(";")
}

// generateCppReturnStatement convert the return value (in 'ret' var), and returns it.
func (o *callableOverload) generateCppReturnStatement(g generator) {
	if o.retrn.isVoid {
		return
	}

	returnConst := ""
	if o.retrn.isConst {
		returnConst = "const"
	}

	var returnValue string
	var pointerTypeName string
	if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
		// pointer to a primitive
		returnValue = fmt.Sprintf("reinterpret_cast<%s %s *> (ret)", returnConst, o.retrn.subTyp.cName)
	} else if o.retrn.isPointer || o.retrn.isSmartPointer && o.retrn.subTyp != nil && o.retrn.subTyp.record != nil {
		// pointer to a record
		returnValue = fmt.Sprintf("reinterpret_cast<%s %s *> (ret)", returnConst, o.retrn.subTyp.record.cStructName)
		pointerTypeName = o.retrn.subTyp.record.cStructName
	} else if o.retrn.isPointer || o.retrn.isSmartPointer {
		returnValue = fmt.Sprintf("reinterpret_cast<%s %s *> (ret)", returnConst, o.retrn.record.cStructName)
		pointerTypeName = o.retrn.record.cStructName
	} else if o.retrn.record != nil {
		if o.record != nil {
			returnValue = fmt.Sprintf("*(reinterpret_cast<%s %s *> (&ret))", returnConst, o.retrn.record.cStructName)
		} else {
			returnValue = fmt.Sprintf("*(reinterpret_cast<%s %s *> (&ret))", returnConst, o.retrn.record.cStructName)
		}
	} else {
		returnValue = "ret"
	}

	if o.retrn.isConst {
		returnValue = fmt.Sprintf("const_cast<%s *>(%s)",
			pointerTypeName,
			returnValue)
	}

	f := g.cppFile
	f.writelnf("  return %s;", returnValue)
}
