package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type typ struct {
	cName             string
	goName            string
	isLValueReference bool
	class             *class
	subTyp            *typ
}

func typFromClangType(cType clang.Type, api api) typ {
	typ := typ{
		cName: strings.TrimPrefix(cType.Spelling(), "const "),
	}

	if class, ok := api.findClass(typ.cName); ok {
		typ.class = class
		typ.goName = typ.class.goName
	} else {
		switch cType.Kind() {
		case clang.Type_LValueReference:
			subTyp := typFromClangType(cType.PointeeType(), api)
			typ.subTyp = &subTyp
			typ.isLValueReference = true
			typ.goName = typ.subTyp.goName

		default:
			panic(fmt.Sprintf("unsupported type '%s'", cType.Spelling()))
		}
	}

	return typ
}
