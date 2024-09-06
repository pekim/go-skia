package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cName   string
	cgoName string
	cgoVar  string
	cParam  string
	cArg    string
	goName  string
	typ     typ
}

func newParam(paramIndex int, cursor clang.Cursor, api api) param {
	cName := cursor.DisplayName()
	if cName == "" {
		cName = fmt.Sprintf("p%d", paramIndex)
	}
	typ := mustTypFromClangType(cursor.Type(), api)
	cgoName := "c_" + cName

	p := param{
		cName:   cName,
		cgoName: cgoName,
		goName:  validGoName(cName),
		typ:     typ,
	}

	if typ.isPrimitive {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cgoName, p.typ.cgoName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.cName, p.cgoName)
		p.cArg = p.cgoName

	} else if typ.enum != nil {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cgoName, typ.enum.cType.cgoName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", typ.enum.cType.cName, p.cgoName)
		p.cArg = fmt.Sprintf("%s(%s)", typ.cName, p.cgoName)

	} else if typ.isLValueReference && typ.subTyp.class != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cgoName, p.goName)
		p.cParam = fmt.Sprintf("void *%s", p.cgoName)
		p.cArg = fmt.Sprintf("*reinterpret_cast<%s*>(%s)", p.typ.subTyp.cName, p.cgoName)

	} else if typ.isPointer && typ.subTyp.class != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cgoName, p.goName)
		p.cParam = fmt.Sprintf("void *%s", p.cgoName)
		p.cArg = fmt.Sprintf("reinterpret_cast<%s*>(%s)", p.typ.subTyp.cName, p.cgoName)

	} else {
		fatalf("unhandled cgoVar %s for param with typ %#v", cName, typ)
	}

	return p
}
