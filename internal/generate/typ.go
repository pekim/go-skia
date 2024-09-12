package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type typ struct {
	cppName           string
	cName             string
	goName            string
	isConst           bool
	isPrimitive       bool // a simple type, that can be converted from Go type to C type with a type conversion
	isLValueReference bool
	isPointer         bool
	isVoid            bool
	isSmartPointer    bool
	record            *record
	enum              *enum
	typedef           *typedef
	subTyp            *typ
}

func typFromClangType(cType clang.Type, api api) (typ, error) {
	cName := strings.TrimPrefix(cType.Spelling(), "const ")
	typ := typ{
		cppName: cName,
		cName:   cName,
		isConst: cType.IsConstQualifiedType(),
	}

	var recordName string
	var recordEnumName string
	nameParts := strings.Split(typ.cppName, "::")
	if len(nameParts) > 0 {
		recordName = nameParts[0]
	}
	if len(nameParts) > 1 {
		recordEnumName = nameParts[1]
	}

	if record, ok := api.findRecord(recordName); ok {
		if recordEnumName == "" {
			typ.record = record
			typ.goName = typ.record.goName
		} else {
			if enum, ok := record.findEnum(recordEnumName); ok {
				typ.enum = enum
				typ.goName = typ.enum.goName
			}
		}

	} else if enum, ok := api.findEnum(typ.cppName); ok {
		typ.enum = enum
		typ.goName = typ.enum.goName

	} else if typedef, ok := api.findTypedef(typ.cppName); ok {
		typ.typedef = typedef
		typ.goName = typ.typedef.goName

	} else if strings.HasPrefix(typ.cppName, "sk_sp<") {
		// A type like "sk_sp<SkSomeClass>" is of kind clang.Type_Elaborated.
		// Its CanonicalType's kind is clang.Type_Record.
		// The template argument 'SkSomeClass' can be got, but it is not clear how to get the 'sk_sp'
		// So determine it's a smart pointer from the C name starting with "sk_sp<".
		typ_, err := typFromClangType(cType.TemplateArgumentAsType(0), api)
		if err != nil {
			return typ, err
		}
		typ = typ_
		typ.isSmartPointer = true

	} else {
		switch cType.Kind() {
		case clang.Type_Void:
			typ.isVoid = true

		case clang.Type_Bool:
			typ.cName = "bool"
			typ.goName = "bool"
			typ.isPrimitive = true

		case clang.Type_UChar:
			typ.cName = "uint"
			typ.goName = "uint"
			typ.isPrimitive = true

		case clang.Type_Float:
			typ.cName = "float"
			typ.goName = "float32"
			typ.isPrimitive = true

		case clang.Type_Int:
			typ.goName = "int"
			typ.isPrimitive = true

		case clang.Type_UInt:
			typ.cName = "uint"
			typ.goName = "uint"
			typ.isPrimitive = true

		case clang.Type_Elaborated:
			typ_, err := typFromClangType(cType.CanonicalType(), api)
			if err != nil {
				return typ, err
			}
			typ = typ_

		case clang.Type_Pointer:
			subTyp, err := typFromClangType(cType.PointeeType(), api)
			if err != nil {
				return typ, err
			}
			typ.subTyp = &subTyp
			typ.isPointer = true
			typ.goName = typ.subTyp.goName

		case clang.Type_LValueReference:
			subTyp, err := typFromClangType(cType.PointeeType(), api)
			if err != nil {
				return typ, err
			}
			typ.subTyp = &subTyp
			typ.isLValueReference = true
			typ.goName = typ.subTyp.goName

		default:
			return typ, fmt.Errorf("unsupported type '%s', of kind %s", cType.Spelling(), cType.Kind())
		}
	}

	return typ, nil
}

func mustTypFromClangType(cType clang.Type, api api) typ {
	typ, err := typFromClangType(cType, api)
	fatalOnError(err)
	return typ
}
