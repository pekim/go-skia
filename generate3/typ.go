package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type typ struct {
	cName             string
	cgoName           string
	goName            string
	isPrimitive       bool // a simple type, that can be converted from Go type to C type with a type conversion
	isLValueReference bool
	isPointer         bool
	isVoid            bool
	isSmartPointer    bool
	class             *class
	enum              *enum
	subTyp            *typ
}

func typFromClangType(cType clang.Type, api api) typ {
	cName := strings.TrimPrefix(cType.Spelling(), "const ")
	typ := typ{
		cName:   cName,
		cgoName: cName,
	}

	if class, ok := api.findClass(typ.cName); ok {
		typ.class = class
		typ.goName = typ.class.goName

	} else if enum, ok := api.findEnum(typ.cName); ok {
		typ.enum = enum
		typ.goName = typ.enum.goName

	} else if strings.HasPrefix(typ.cName, "sk_sp<") {
		// A type like "sk_sp<SkSomeClass>" is of kind clang.Type_Elaborated.
		// Its CanonicalType's kind is clang.Type_Record.
		// The template argument 'SkSomeClass' can be got, but it is not clear how to get the 'sk_sp'
		// So determine it's a smart pointer from the C name starting with "sk_sp<".
		typ = typFromClangType(cType.TemplateArgumentAsType(0), api)
		typ.isSmartPointer = true
		typ.cName = "void *"

	} else {
		switch cType.Kind() {
		case clang.Type_Void:
			typ.isVoid = true

		case clang.Type_Bool:
			typ.cgoName = "bool"
			typ.goName = "bool"
			typ.isPrimitive = true

		case clang.Type_UChar:
			typ.cgoName = "uint"
			typ.goName = "uint"
			typ.isPrimitive = true

		case clang.Type_Int:
			typ.goName = "int"
			typ.isPrimitive = true

		case clang.Type_UInt:
			typ.cgoName = "uint"
			typ.goName = "uint"
			typ.isPrimitive = true

		case clang.Type_Elaborated:
			typ = typFromClangType(cType.CanonicalType(), api)

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
			fatalf("unsupported type '%s', of kind %s", cType.Spelling(), cType.Kind())
		}
	}

	return typ
}
