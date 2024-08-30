package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cName   string
	cgoName string
	goName  string
	typ     typ
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

	typ, err := typeFromClangType(cursor.Type())
	if err != nil {
		panic(err)
	}
	p.typ = typ

	return p
}

func (p param) supported() bool {
	if p.typ.isArray {
		// TODO support array params
		return false
	}
	if p.typ.pointerLevel > 0 {
		// TODO support pointer params
		return false
	}
	if p.typ.isEnumLiteral {
		// TODO support enum params
		return false
	}

	if p.typ.isPrimitive {
		return true
	}

	return false
}

func (p param) goDecl() string {
	return fmt.Sprintf("%s %s", p.goName, p.typ.goName)
}

func (p param) goCArg() string {
	return fmt.Sprintf("%s := C.%s(%s)", p.cgoName, p.typ.cgoName, p.goName)
}

func (p param) cParamDecl() string {
	return fmt.Sprintf("%s %s", p.typ.cgoName, p.cgoName)
}

func (p param) cArg() string {
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
