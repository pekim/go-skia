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
	isArray           bool
	isLValueReference bool
	isPointer         bool
	isVoid            bool
	isSmartPointer    bool
	isOptional        bool
	record            *record
	enum              *enum
	typedef           *typedef
	subTyp            *typ
}

func typFromClangType(cType clang.Type, api api, templateRef string) (typ, error) {
	cName := strings.TrimPrefix(cType.Spelling(), "const ")
	typ := typ{
		cppName: cName,
		cName:   cName,
		isConst: cType.IsConstQualifiedType(),
	}
	// Not sure how to deal with class templates properly. So treat as a special case for now.
	if typ.cName == "SkRGBA4f<kUnpremul_SkAlphaType>" {
		typ.cName = "SkRGBA4f"
	}

	switch cType.Kind() {
	case clang.Type_Pointer:
		subTyp, err := typFromClangType(cType.PointeeType(), api, "")
		if err != nil {
			return typ, err
		}
		typ.subTyp = &subTyp
		typ.isPointer = true
		typ.goName = typ.subTyp.goName
		if typ.subTyp.goName == "byte" {
			typ.goName = "string"
		}
		return typ, nil

	case clang.Type_LValueReference:
		subTyp, err := typFromClangType(cType.PointeeType(), api, "")
		if err != nil {
			return typ, err
		}
		typ.subTyp = &subTyp
		typ.isLValueReference = true
		typ.goName = typ.subTyp.goName
		return typ, nil

	case clang.Type_IncompleteArray:
		subTyp, err := typFromClangType(cType.ArrayElementType(), api, "")
		if err != nil {
			return typ, err
		}
		typ.subTyp = &subTyp
		typ.isArray = true
		typ.goName = fmt.Sprintf("[]%s", typ.subTyp.goName)
		if typ.subTyp.goName == "byte" {
			typ.goName = "string"
		}
		return typ, nil
	}

	var recordName string
	var recordEnumName string
	nameParts := strings.Split(typ.cName, "::")
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
			typ.cName = typ.record.cStructName
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

	} else if templateRef == "sk_sp" || templateRef == "unique_ptr" {
		typ_, err := typFromClangType(cType.TemplateArgumentAsType(0), api, "")
		if err != nil {
			return typ, err
		}
		typ = typ_
		typ.isSmartPointer = true

	} else if templateRef == "optional" {
		typ_, err := typFromClangType(cType.TemplateArgumentAsType(0), api, "")
		if err != nil {
			return typ, err
		}
		typ = typ_
		typ.isOptional = true

	} else {
		switch cType.Kind() {
		case clang.Type_Void:
			typ.isVoid = true

		case clang.Type_Bool:
			typ.cName = "bool"
			typ.goName = "bool"
			typ.isPrimitive = true

		case clang.Type_Char_S:
			typ.cName = "char"
			typ.goName = "byte"
			typ.isPrimitive = true

		case clang.Type_UChar:
			typ.cName = "uchar"
			typ.goName = "byte"
			typ.isPrimitive = true

		case clang.Type_Float:
			typ.cName = "float"
			typ.goName = "float32"
			typ.isPrimitive = true

		case clang.Type_Int:
			typ.cName = "int"
			typ.goName = "int32"
			typ.isPrimitive = true

		case clang.Type_UInt:
			typ.cName = "uint"
			typ.goName = "uint32"
			typ.isPrimitive = true

		case clang.Type_UShort:
			typ.cName = "ushort"
			typ.goName = "uint16"
			typ.isPrimitive = true

		case clang.Type_ULong:
			typ.cName = "ulong"
			typ.goName = "uint32"
			typ.isPrimitive = true

		case clang.Type_Elaborated:
			typ_, err := typFromClangType(cType.CanonicalType(), api, "")
			if err != nil {
				return typ, err
			}
			typ = typ_

		default:
			return typ, fmt.Errorf("unsupported type '%s', of kind %s", cType.Spelling(), cType.Kind())
		}
	}

	return typ, nil
}

func mustTypFromClangType(cType clang.Type, api api, templateRef string) typ {
	typ, err := typFromClangType(cType, api, templateRef)
	fatalOnError(err)
	return typ
}
