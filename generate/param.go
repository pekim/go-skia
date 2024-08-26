package generate

import (
	"fmt"
	"strings"
)

type param struct {
	cName  string
	goName string
	cDecl_ string
}

func newParam(cName string, cDecl string) param {
	return param{
		cName:  cName,
		cDecl_: cDecl,
		goName: cName,
	}
}

func (p param) supported() bool {
	return p.cDecl_ == "SkScalar" ||
		p.cDecl_ == "uint32_t" ||
		p.cDecl_ == "size_t"
}

func (p param) goDecl() string {
	switch p.cDecl_ {
	case "SkScalar":
		return fmt.Sprintf("%s float32", p.goName)

	case "size_t":
		return fmt.Sprintf("%s uint", p.goName)

	case "uint32_t":
		return fmt.Sprintf("%s uint32", p.goName)

	default:
		panic("unsupported param type")
	}
}

func (p param) goCArg() string {
	switch p.cDecl_ {
	case "SkScalar":
		return fmt.Sprintf("C.float(%s)", p.goName)

	case "size_t":
		return fmt.Sprintf("C.size_t(%s)", p.goName)

	case "uint32_t":
		return fmt.Sprintf("C.uint32_t(%s)", p.goName)

	default:
		panic("unsupported param type")
	}
}

func (p param) cDecl() string {
	switch p.cDecl_ {
	case "SkScalar":
		return fmt.Sprintf("float %s", p.cName)

	case "size_t":
		return fmt.Sprintf("size_t %s", p.cName)

	case "uint32_t":
		return fmt.Sprintf("uint32_t %s", p.cName)

	default:
		panic("unsupported param type")
	}
}

func (p param) cArg() string {
	switch p.cDecl_ {
	case "SkScalar":
		return fmt.Sprintf("(SkScalar)%s", p.cName)

	case "size_t":
		return p.cName

	case "uint32_t":
		return p.cName

	default:
		panic("unsupported param type")
	}
}

func makeParamsString(params []param, makeParam func(p param) string) string {
	var paramsStrings = make([]string, len(params))
	for p, param := range params {
		paramsStrings[p] = makeParam(param)
	}
	return strings.Join(paramsStrings, ", ")
}
