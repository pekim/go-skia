package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type field struct {
	name      string
	cType     string
	goType    string
	size      int // bytes
	align     int // bytes
	offset    int // bits
	public    bool
	clangType clang.Type
	doc       string
}

func newField(cursor clang.Cursor) field {
	doc := docComment(cursor.ParsedComment())

	f := field{
		name:      cursor.DisplayName(),
		size:      int(cursor.Type().SizeOf()),
		align:     int(cursor.Type().AlignOf()),
		offset:    int(cursor.OffsetOfField()),
		public:    cursor.AccessSpecifier() == clang.AccessSpecifier_Public,
		clangType: cursor.Type(),
		doc:       doc,
	}
	return f
}

func (f *field) enrich2(api api) {
	typ, err := typFromClangType(f.clangType, api, "")
	if err == nil && typ.isPrimitive {
		f.cType = typ.cName
		f.goType = typ.goName
	}
	f.doc = addDocCommentLinks(f.doc, api)
}
