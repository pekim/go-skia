package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type variable struct {
	cursor     clang.Cursor
	cppName    string
	cName      string
	goName     string
	typeCName  string
	typeGoName string
	doc        string
}

func newVariable(cursor clang.Cursor) variable {
	cppName := cursor.Spelling()

	v := variable{
		cursor:  cursor,
		cppName: cppName,
		cName:   "sk_" + cppName,
		goName:  goExportedName(stripSkPrefix(stripKPrefix(cppName))),
		doc:     docComment(cursor.ParsedComment()),
	}
	return v
}

func (v *variable) enrich2(api api) {
	typ, err := typFromClangType(v.cursor.Type(), api, "")
	if err == nil {
		if typ.isPrimitive {
			v.typeCName = typ.cName
			v.typeGoName = typ.goName
		} else if typ.typedef != nil {
			v.typeCName = typ.typedef.cName
			v.typeGoName = typ.goName
		} else {
			fatalf("unsupported type for variable, %#v", typ)
		}
	}
}

func (v *variable) generate(g generator) {
	if v.typeCName == "" || v.typeGoName == "" {
		return
	}

	g.goFile.write(v.doc)
	g.goFile.writelnf("var %s = (%s)(C.%s)", v.goName, v.typeGoName, v.cName)

	g.headerFile.writelnf("extern %s %s;", v.typeCName, v.cName)
	g.cppFile.writelnf("%s %s = %s;", v.typeCName, v.cName, v.cppName)
}
