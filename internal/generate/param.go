package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cppName   string
	cName     string
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
		cppName:   cName,
		cName:     cgoName,
		goName:    validGoName(cName),
		clangType: cursor.Type(),
	}

	return p
}

func (p *param) enrich2(api api) {
	p.typ = mustTypFromClangType(p.clangType, api)

	if p.typ.isPrimitive {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, p.typ.cName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.cppName, p.cName)
		p.cArg = p.cName

	} else if p.typ.enum != nil {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, p.typ.enum.cType.cName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.enum.cType.cName, p.cName)
		p.cArg = fmt.Sprintf("%s(%s)", p.typ.cppName, p.cName)

	} else if p.typ.typedef != nil {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, p.typ.typedef.cName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.typedef.cType.cName, p.cName)
		p.cArg = fmt.Sprintf("%s(%s)", p.typ.cppName, p.cName)

	} else if p.typ.isLValueReference && p.typ.subTyp.record != nil {
		if p.typ.subTyp.record.NoWrapper {
			p.cgoVar = fmt.Sprintf("%s := *(*C.%s)(unsafe.Pointer(&%s))", p.cName, p.typ.subTyp.record.cStructName, p.goName)
			p.cParam = fmt.Sprintf("sk_%s %s", p.typ.subTyp.cppName, p.cName)
			p.cArg = fmt.Sprintf("*reinterpret_cast<%s*>(&%s)", p.typ.subTyp.cppName, p.cName)
		} else {
			p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
			p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.subTyp.cppName, p.cName)
			p.cArg = fmt.Sprintf("*reinterpret_cast<%s*>(%s)", p.typ.subTyp.cppName, p.cName)
		}

	} else if p.typ.isPointer && p.typ.subTyp.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
		p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.subTyp.cppName, p.cName)
		p.cArg = fmt.Sprintf("reinterpret_cast<%s*>(%s)", p.typ.subTyp.cppName, p.cName)

	} else if p.typ.isSmartPointer && p.typ.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
		p.cParam = fmt.Sprintf("sk_%s *%s", p.typ.cppName, p.cName)
		p.cArg = fmt.Sprintf("sk_ref_sp(reinterpret_cast<%s*>(%s))", p.typ.cppName, p.cName)

	} else {
		fatalf("unhandled cgoVar %s for param with typ %#v", p.cppName, p.typ)
	}
}
