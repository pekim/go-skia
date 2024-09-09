package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type field struct {
	name      string
	typ       string
	size      int // bytes
	align     int // bytes
	offset    int // bits
	public    bool
	clangType clang.Type
}

func newField(cursor clang.Cursor) field {
	f := field{
		name:      cursor.DisplayName(),
		size:      int(cursor.Type().SizeOf()),
		align:     int(cursor.Type().AlignOf()),
		offset:    int(cursor.OffsetOfField()),
		public:    cursor.AccessSpecifier() == clang.AccessSpecifier_Public,
		clangType: cursor.Type(),
	}

	return f
}

func (f *field) enrich2(api api) {
	typ, err := typFromClangType(f.clangType, api)
	if err == nil && typ.isPrimitive {
		f.typ = typ.cName
	}
}
