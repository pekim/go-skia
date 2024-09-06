package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type field struct {
	name   string
	typ    string
	size   int // bytes
	align  int // bytes
	offset int // bits
	public bool
}

func newField(cursor clang.Cursor, api api) field {
	f := field{
		name:   cursor.DisplayName(),
		size:   int(cursor.Type().SizeOf()),
		align:  int(cursor.Type().AlignOf()),
		offset: int(cursor.OffsetOfField()),
		public: cursor.AccessSpecifier() == clang.AccessSpecifier_Public,
	}

	typ, err := typFromClangType(cursor.Type(), api)
	if err == nil && typ.isPrimitive {
		f.typ = typ.cName
	}

	return f
}
