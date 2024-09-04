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
	typ := typFromClangType(cursor.Type(), api)
	cgoName := "c_" + cName

	p := param{
		cName:   cName,
		cgoName: cgoName,
		goName:  validGoName(cName),
		typ:     typ,
	}

	if typ.isLValueReference && typ.subTyp.class != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", cgoName, p.goName)
		p.cParam = fmt.Sprintf("void *%s", p.cgoName)
		p.cArg = fmt.Sprintf("*reinterpret_cast<%s*>(%s)", p.typ.subTyp.cName, p.cgoName)
	} else {
		panic(fmt.Sprintf("unhandled cgoVar for param with typ %#v", typ))
	}

	return p
}
