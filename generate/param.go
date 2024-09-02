package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cName    string
	cgoName  string
	goName   string
	typ      typ
	isClass  bool
	isStruct bool
}

func newParam(cursor clang.Cursor, n int) param {
	p := param{
		cName: cursor.DisplayName(),
	}
	if p.cName == "" {
		p.cName = fmt.Sprintf("p%d", n)
	}

	p.cgoName = fmt.Sprintf("c_%s", p.cName)
	p.goName = validGoName(p.cName)

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_TypeRef {
			if strings.Contains(cursor.Spelling(), "struct ") {
				p.isStruct = true
			}
			if strings.Contains(cursor.Spelling(), "class ") {
				p.isClass = true
			}
		}
		return clang.ChildVisit_Continue
	})

	typ, err := typeFromClangType(cursor.Type())
	if err != nil {
		panic(err)
	}
	p.typ = typ

	return p
}

func (p param) supported() (bool, string) {
	if p.typ.isLValueReference {
		return true, ""
	}

	if p.typ.unsupported != "" {
		return false, p.typ.unsupported
	}
	if p.typ.isArray {
		// TODO support array params
		return false, "array not supported"
	}
	if p.typ.pointerLevel > 0 {
		// TODO support pointer params
		return false, "pointer not supported"
	}

	if p.typ.isPrimitive {
		return true, ""
	}

	return false, "not supported"
}

func (p param) goDecl() string {
	return fmt.Sprintf("%s %s", p.goName, p.typ.goName)
}

func (p param) goCArg() string {
	if p.typ.isLValueReference {
		if p.typ.isPrimitive {
			return fmt.Sprintf("%s := (*C.%s)(&%s)", p.cgoName, p.typ.cgoName, p.goName)
		}
		return fmt.Sprintf("%s := %s.skia", p.cgoName, p.goName)
	}
	return fmt.Sprintf("%s := C.%s(%s)", p.cgoName, p.typ.cgoName, p.goName)
}

func (p param) cParamDecl() string {
	if p.typ.isLValueReference {
		if p.typ.isPrimitive {
			return fmt.Sprintf("%s* %s", p.typ.cName, p.cgoName)
		}
		return fmt.Sprintf("void* %s", p.cgoName)
	}
	return fmt.Sprintf("%s %s", p.typ.cgoName, p.cgoName)
}

func (p param) cArg() string {
	if p.typ.isEnumLiteral {
		return fmt.Sprintf("(%s)%s", p.typ.cName, p.cgoName)
	}
	if p.typ.isLValueReference {
		if p.typ.isPrimitive {
			return fmt.Sprintf("reinterpret_cast<%s &>(*%s)", p.typ.cName, p.cgoName)
		}
		return fmt.Sprintf("reinterpret_cast<%s &>(%s)", p.typ.cName, p.cgoName)
	}
	return p.cgoName
}

func makeParamsString(params []param, makeParam func(p param) string) string {
	var paramsStrings = make([]string, len(params))
	for p, param := range params {
		paramsStrings[p] = makeParam(param)
	}
	return strings.Join(paramsStrings, ", ")
}

func validGoName(name string) string {
	if name == "type" {
		return "typ"
	}
	return name
}
