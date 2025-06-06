package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type typedef struct {
	cppName   string
	cName     string
	goName    string
	cType     typ
	clangType clang.Type
	doc       string
	enriched1 bool
	enriched2 bool
}

func (t *typedef) enrich1(cursor clang.Cursor) {
	t.goName = goExportedName(stripSkPrefix(cursor.Spelling()))
	t.clangType = cursor.TypedefDeclUnderlyingType()
	t.doc = docComment(cursor.ParsedComment())
	t.enriched1 = true
}

func (t *typedef) enrich2(api api) {
	if !t.enriched1 {
		fatalf("typedef %s has not been phase 1 enriched", t.cppName)
	}

	t.cType = mustTypFromClangType(t.clangType, api, "")
	t.cName = t.cType.cName
	t.enriched2 = true
	t.doc = addDocCommentLinks(t.doc, api)
}

func (t typedef) generate(g generator) {
	if !t.enriched2 {
		fatalf("typedef %s has not been enriched", t.cppName)
	}

	f := g.goFile
	f.write(t.doc)
	f.writelnf("type %s C.%s", t.goName, t.cName)
	f.writeln()
}
