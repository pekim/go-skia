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
	cppArg    string
	goName    string
	clangType clang.Type
	typ       typ
	typGoName string
	ValueNil  bool `json:"valueNil"`
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
	if p.ValueNil {
		return
	}

	p.typ = mustTypFromClangType(p.clangType, api)
	p.typGoName = p.typ.goName

	if p.typ.isPrimitive {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, p.typ.cName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.cName, p.cName)
		p.cppArg = p.cName

	} else if p.typ.enum != nil {
		cType := p.typ.enum.cType.cName
		if cType == "bool" {
			cType = "int"
		}
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, cType, p.goName)
		p.cParam = fmt.Sprintf("%s %s", cType, p.cName)
		p.cppArg = fmt.Sprintf("%s(%s)", p.typ.cppName, p.cName)

	} else if p.typ.typedef != nil {
		p.cgoVar = fmt.Sprintf("%s := C.%s(%s)", p.cName, p.typ.typedef.cName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.typedef.cType.cName, p.cName)
		p.cppArg = fmt.Sprintf("%s(%s)", p.typ.cppName, p.cName)

	} else if p.typ.isArray {
		p.cgoVar = fmt.Sprintf("%s := (*C.%s)(unsafe.Pointer(&%s[0]))", p.cName, p.typ.subTyp.cName, p.goName)
		if p.typ.goName == "string" {
			p.cgoVar = fmt.Sprintf(`%s := C.CString(%s)
			defer C.free(unsafe.Pointer(%s))`,
				p.cName, p.goName, p.cName)
		}
		p.cParam = fmt.Sprintf("%s *%s", p.typ.subTyp.cName, p.cName)
		p.cppArg = fmt.Sprintf("(%s*)%s", p.typ.subTyp.cppName, p.cName)

	} else if p.typ.isPointer && p.typ.subTyp.isVoid {
		p.typGoName = "[]byte"
		p.cgoVar = fmt.Sprintf("%s := unsafe.Pointer(&%s[0])", p.cName, p.goName)
		p.cParam = fmt.Sprintf("void *%s", p.cName)
		p.cppArg = fmt.Sprintf("reinterpret_cast<%s*>(%s)", p.typ.subTyp.cppName, p.cName)

	} else if p.typ.goName == "string" {
		p.cgoVar = fmt.Sprintf(`%s := C.CString(%s)
	defer C.free(unsafe.Pointer(%s))`,
			p.cName, p.goName, p.cName)
		p.cParam = fmt.Sprintf("%s *%s", p.typ.subTyp.cName, p.cName)
		p.cppArg = fmt.Sprintf("(%s*)%s", p.typ.subTyp.cppName, p.cName)

	} else if p.typ.isLValueReference && p.typ.subTyp.record != nil {
		if p.typ.subTyp.record.NoWrapper {
			p.cgoVar = fmt.Sprintf("%s := *(*C.%s)(unsafe.Pointer(&%s))", p.cName, p.typ.subTyp.record.cStructName, p.goName)
			p.cParam = fmt.Sprintf("%s %s", p.typ.subTyp.cName, p.cName)
			p.cppArg = fmt.Sprintf("*reinterpret_cast<%s*>(&%s)", p.typ.subTyp.cppName, p.cName)
		} else {
			p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
			p.cParam = fmt.Sprintf("%s *%s", p.typ.subTyp.cName, p.cName)
			p.cppArg = fmt.Sprintf("*reinterpret_cast<%s*>(%s)", p.typ.subTyp.cppName, p.cName)
		}

	} else if p.typ.isPointer && p.typ.subTyp.record != nil {
		if p.typ.subTyp.record.NoWrapper {
			p.cgoVar = fmt.Sprintf("%s := *(*C.%s)(unsafe.Pointer(&%s))", p.cName, p.typ.subTyp.record.cStructName, p.goName)
			p.cParam = fmt.Sprintf("%s %s", p.typ.subTyp.cName, p.cName)
			p.cppArg = fmt.Sprintf("reinterpret_cast<%s*>(&%s)", p.typ.subTyp.cppName, p.cName)
		} else {
			p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
			p.cParam = fmt.Sprintf("%s *%s", p.typ.subTyp.cName, p.cName)
			p.cppArg = fmt.Sprintf("reinterpret_cast<%s*>(%s)", p.typ.subTyp.cppName, p.cName)
		}

	} else if p.typ.isSmartPointer && p.typ.record != nil {
		p.cgoVar = fmt.Sprintf("%s := %s.sk", p.cName, p.goName)
		p.cParam = fmt.Sprintf("%s *%s", p.typ.cName, p.cName)
		p.cppArg = fmt.Sprintf("sk_ref_sp(reinterpret_cast<%s*>(%s))", p.typ.cppName, p.cName)

	} else if p.typ.record != nil {
		p.cgoVar = fmt.Sprintf("%s := *(*C.%s)(unsafe.Pointer(&%s))", p.cName, p.typ.record.cStructName, p.goName)
		p.cParam = fmt.Sprintf("%s %s", p.typ.cName, p.cName)
		p.cppArg = fmt.Sprintf("*reinterpret_cast<%s*>(&%s)", p.typ.cppName, p.cName)

	} else {
		fatalf("unhandled param '%s' with typ %#v", p.cppName, p.typ)
	}
}
