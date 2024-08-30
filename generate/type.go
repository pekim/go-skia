package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

// Go types
const (
	GoByte      = "byte"
	GoInt8      = "int8"
	GoUInt8     = "uint8"
	GoInt16     = "int16"
	GoUInt16    = "uint16"
	GoInt32     = "int32"
	GoUInt32    = "uint32"
	GoInt64     = "int64"
	GoUInt64    = "uint64"
	GoFloat32   = "float32"
	GoFloat64   = "float64"
	GoBool      = "bool"
	GoInterface = "interface"
	GoPointer   = "unsafe.Pointer"
)

// C types
const (
	CChar      = "char"
	CSChar     = "schar"
	CUChar     = "uchar"
	CShort     = "short"
	CUShort    = "ushort"
	CInt       = "int"
	CUInt      = "uint"
	CLongInt   = "long"
	CULongInt  = "ulong"
	CLongLong  = "longlong"
	CULongLong = "ulonglong"
	CFloat     = "float"
	CDouble    = "double"
)

// typ represents a generation type.
type typ struct {
	cName         string
	cgoName       string
	goName        string
	pointerLevel  int
	isPrimitive   bool
	isArray       bool
	arraySize     int64
	isEnumLiteral bool
}

// yypeFromClangType returns the typ from a Clang type.
func typeFromClangType(cType clang.Type) (typ, error) {
	typ_ := typ{
		cName:        cType.Spelling(),
		pointerLevel: 0,
		isPrimitive:  true,
		isArray:      false,
	}

	switch cType.Kind() {
	case clang.Type_Char_S:
		typ_.cgoName = CSChar
		typ_.goName = GoInt8

	case clang.Type_Char_U:
		typ_.cgoName = CUChar
		typ_.goName = GoUInt8

	case clang.Type_Int:
		typ_.cgoName = CInt
		typ_.goName = GoInt32

	case clang.Type_Short:
		typ_.cgoName = CShort
		typ_.goName = GoInt16

	case clang.Type_UShort:
		typ_.cgoName = CUShort
		typ_.goName = GoUInt16

	case clang.Type_UInt:
		typ_.cgoName = CUInt
		typ_.goName = GoUInt32

	case clang.Type_Long:
		typ_.cgoName = CLongInt
		typ_.goName = GoInt64

	case clang.Type_ULong:
		typ_.cgoName = CULongInt
		typ_.goName = GoUInt64

	case clang.Type_LongLong:
		typ_.cgoName = CLongLong
		typ_.goName = GoInt64

	case clang.Type_ULongLong:
		typ_.cgoName = CULongLong
		typ_.goName = GoUInt64

	case clang.Type_Float:
		typ_.cgoName = CFloat
		typ_.goName = GoFloat32

	case clang.Type_Double:
		typ_.cgoName = CDouble
		typ_.goName = GoFloat64

	case clang.Type_Bool:
		typ_.goName = GoBool

	case clang.Type_Void:
		typ_.cgoName = "void"
		typ_.goName = "void" // There isn't really any such Go type.

	case clang.Type_ConstantArray:
		subTyp, err := typeFromClangType(cType.ArrayElementType())
		if err != nil {
			return typ{}, err
		}

		typ_.cgoName = subTyp.cgoName
		typ_.goName = subTyp.goName
		typ_.pointerLevel += subTyp.pointerLevel
		typ_.isArray = true
		typ_.arraySize = cType.ArraySize()

	case clang.Type_Typedef:
		typ_.isPrimitive = false
		typ_.cgoName = cType.Declaration().Type().Spelling()
		typ_.goName = trimSkiaPrefix(cType.Declaration().Type().Spelling())

		if cType.CanonicalType().Kind() == clang.Type_Enum {
			typ_.isEnumLiteral = true
			typ_.isPrimitive = true
		}

	case clang.Type_Pointer:
		typ_.pointerLevel++

		subTyp, err := typeFromClangType(cType.PointeeType())
		if err != nil {
			return typ{}, err
		}

		typ_.cgoName = subTyp.cgoName
		typ_.goName = subTyp.goName
		typ_.pointerLevel += subTyp.pointerLevel
		typ_.isPrimitive = subTyp.isPrimitive

	case clang.Type_Record:
		typ_.cgoName = cType.Declaration().Type().Spelling()
		typ_.goName = trimSkiaPrefix(typ_.cgoName)
		typ_.isPrimitive = false

	case clang.Type_Enum:
		typ_.goName = trimSkiaPrefix(cType.Declaration().DisplayName())
		typ_.isEnumLiteral = true
		typ_.isPrimitive = true

	case clang.Type_Elaborated:
		return typeFromClangType(cType.CanonicalType())

	case clang.Type_Unexposed: // there is a bug in clang for enums the kind is set to unexposed dunno why, bug persisted since 2013: https://llvm.org/bugs/show_bug.cgi?id=15089
		subTyp, err := typeFromClangType(cType.CanonicalType())
		if err != nil {
			return typ{}, err
		}

		typ_.cgoName = subTyp.cgoName
		typ_.goName = subTyp.goName
		typ_.pointerLevel += subTyp.pointerLevel
		typ_.isPrimitive = subTyp.isPrimitive

	default:
		return typ{}, fmt.Errorf("unhandled type %q of kind %q", cType.Spelling(), cType.Kind().Spelling())
	}

	return typ_, nil
}

func trimSkiaPrefix(name string) string {
	return strings.TrimPrefix(name, "Sk")
}
