package generate

import (
	"fmt"
	"strings"
)

func (o callableOverload) generateCpp(g generator) {
	f := g.cppFile

	returnDecl := o.retrn.cppName
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
	if o.retrn.isConst {
		returnDecl = "const " + returnDecl
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

// generateCppReturnStatement converts the return value (in 'ret' var), and returns it.
func (o *callableOverload) generateCppReturnStatement(g generator) {
	if o.retrn.isVoid {
		return
	}

	f := g.cppFile
	isPointer := o.retrn.isPointer || o.retrn.isSmartPointer

	if !isPointer && o.retrn.record == nil {
		// simple return type, with no conversion required
		f.writeln("  return ret;")
		return
	}

	var cTypeName string
	if o.retrn.isPointer && o.retrn.subTyp.isPrimitive {
		cTypeName = o.retrn.subTyp.cName
	} else if isPointer && o.retrn.subTyp != nil && o.retrn.subTyp.record != nil {
		cTypeName = o.retrn.subTyp.record.cStructName
	} else if isPointer {
		cTypeName = o.retrn.record.cStructName
	} else if o.retrn.record != nil {
		cTypeName = o.retrn.record.cStructName
	} else {
		fatal("unsupported return type")
	}

	returnConst := ""
	if o.retrn.isConst {
		returnConst = "const"
	}

	cast := fmt.Sprintf("reinterpret_cast<%s %s *>", returnConst, cTypeName)

	f.write("  return ")
	if isPointer {
		f.writef("%s(ret)", cast)
	} else {
		f.writef("*(%s(&ret))", cast)
	}
	f.writeln(";")
}
