package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cName     string
	cgoName   string
	cgoVar    string
	cParam    string
	cArg      string
	goName    string
	clangType clang.Type
	typ       typ
}

func newParam(paramIndex int, cursor clang.Cursor) param {
	cName := cursor.DisplayName()
	if cName == "" {
		cName = fmt.Sprintf("p%d", paramIndex)
	}
	cgoName := "c_" + cName

	p := param{
		cName:     cName,
		cgoName:   cgoName,
		goName:    validGoName(cName),
		clangType: cursor.Type(),
	}

	return p
}

func (p *param) enrich2(api api) {
	p.typ = mustTypFromClangType(p.clangType, api)

	if p.typ.isPrimitive {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cgoName, p.typ.cgoName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.cName, p.cgoName)
		p.cArg = p.cgoName

	} else if p.typ.enum != nil {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cgoName, p.typ.enum.cType.cgoName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.enum.cType.cName, p.cgoName)
		p.cArg = fmt.Sprintf("%s(%s)", p.typ.cName, p.cgoName)

	} else if p.typ.isLValueReference && p.typ.subTyp.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cgoName, p.goName)
		p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.subTyp.cName, p.cgoName)
		p.cArg = fmt.Sprintf("*reinterpret_cast<%s*>(%s)", p.typ.subTyp.cName, p.cgoName)

	} else if p.typ.isPointer && p.typ.subTyp.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cgoName, p.goName)
		p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.subTyp.cName, p.cgoName)
		p.cArg = fmt.Sprintf("reinterpret_cast<%s*>(%s)", p.typ.subTyp.cName, p.cgoName)

	} else if p.typ.isSmartPointer && p.typ.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cgoName, p.goName)
		p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.cName, p.cgoName)
		p.cArg = fmt.Sprintf("sk_ref_sp(reinterpret_cast<%s*>(%s))", p.typ.cName, p.cgoName)

	} else {
		fatalf("unhandled cgoVar %s for param with typ %#v", p.cName, p.typ)
	}

}
