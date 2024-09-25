package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type typedef struct {
	CppName   string `json:"name"`
	cName     string
	goName    string
	cType     typ
	clangType clang.Type
	doc       string
	enriched  bool
}

func (t *typedef) enrich1(cursor clang.Cursor) {
	t.goName = goExportedName(stripSkPrefix(cursor.Spelling()))
	t.clangType = cursor.TypedefDeclUnderlyingType()
	t.doc = cursor.RawCommentText()
}

func (t *typedef) enrich2(api api) {
	t.cType = mustTypFromClangType(t.clangType, api, "")
	t.cName = t.cType.cName
	t.enriched = true
}

func (t typedef) generate(g generator) {
	if !t.enriched {
		fatalf("typedef %s has not been enriched", t.CppName)
	}

	f := g.goFile
	f.writeDocComment(t.doc)
	f.writelnf("type %s C.%s", t.goName, t.cName)
	f.writeln()
}
