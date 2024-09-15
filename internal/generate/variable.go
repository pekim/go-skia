package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type variable struct {
	cursor  clang.Cursor
	cppName string
	cName   string
	goName  string
	typ     *typ
	doc     string
}

func newVariable(cursor clang.Cursor) variable {
	cppName := cursor.Spelling()

	v := variable{
		cursor:  cursor,
		cppName: cppName,
		cName:   "sk_" + cppName,
		goName:  goExportedName(stripSkPrefix(cppName)),
		doc:     cursor.RawCommentText(),
	}
	return v
}

func (v *variable) enrich2(api api) {
	typ, err := typFromClangType(v.cursor.Type(), api)
	if err == nil && typ.typedef != nil {
		v.typ = &typ
		fmt.Println("TODO if its not a typdef")
	}
}

func (v *variable) generate(g generator) {
	if v.typ == nil {
		return
	}

	g.goFile.writeDocComment(v.doc)
	g.goFile.writelnf("var %s = (%s)(C.%s)", v.goName, v.typ.goName, v.cName)

	g.headerFile.writelnf("extern %s %s;", v.typ.typedef.cName, v.cName)
	g.cppFile.writelnf("%s %s = %s;", v.typ.typedef.cName, v.cName, v.cppName)
}
