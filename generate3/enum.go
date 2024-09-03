package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	Name      string `json:"name"`
	goName    string
	class     *class
	doc       string
	constants []enumConstant
}

type enumConstant struct {
	goName string
	value  int64
	doc    string
}

func (e *enum) enrich(class *class, cursor clang.Cursor) {
	e.goName = class.goName + e.Name
	e.class = class
	e.doc = cursor.RawCommentText()
	e.doc = strings.Replace(e.doc, fmt.Sprintf("* \\enum %s::%s", class.Name, e.Name), "", 1)

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			name = stripKPrefix(name)
			name = strings.TrimSuffix(name, "_"+e.Name)
			doc := cursor.RawCommentText()
			doc = strings.Replace(doc, "!<", "", 1)

			constant := enumConstant{
				goName: e.goName + name,
				value:  cursor.EnumConstantDeclValue(),
				doc:    doc,
			}
			e.constants = append(e.constants, constant)
		}

		return clang.ChildVisit_Continue
	})
}

func (e enum) generate(g generator) {
	f := g.goFile

	f.writeln(e.doc)
	f.writelnf("type %s int64", e.goName)
	f.writeln("")

	f.writeln("const (")
	for _, constant := range e.constants {
		f.writeln(constant.doc)
		f.writelnf("%s %s = %d", constant.goName, e.goName, constant.value)
	}
	f.writeln(")")
	f.writeln("")
}
