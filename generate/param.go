package generate

import "fmt"

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
		p.cDecl_ == "uint32_t"
}

func (p param) goDecl() string {
	switch p.cDecl_ {
	case "SkScalar":
		return fmt.Sprintf("%s float32", p.goName)

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

	case "uint32_t":
		return p.cName

	default:
		panic("unsupported param type")
	}
}
