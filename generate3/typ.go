package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type typ struct {
	cName             string
	goName            string
	isPrimitive       bool // a simple type, that can be converted from Go type to C type with a type conversion
	isLValueReference bool
	isPointer         bool
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
		case clang.Type_Int:
			typ.goName = "int"
			typ.isPrimitive = true

		case clang.Type_Pointer:
			subTyp := typFromClangType(cType.PointeeType(), api)
			typ.subTyp = &subTyp
			typ.isPointer = true
			typ.goName = typ.subTyp.goName

		case clang.Type_LValueReference:
			subTyp := typFromClangType(cType.PointeeType(), api)
			typ.subTyp = &subTyp
			typ.isLValueReference = true
			typ.goName = typ.subTyp.goName

		default:
			fatalf("unsupported type '%s'", cType.Spelling())
		}
	}

	return typ
}
