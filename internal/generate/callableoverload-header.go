package generate

import "fmt"

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
	if o.retrn.isConst {
		returnDecl = "const " + returnDecl
	}

	f.writelnf("%s %s(%s);", returnDecl, o.cFuncName, o.cParamsDecl)
}
