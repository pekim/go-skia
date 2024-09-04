package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	CName     string `json:"name"`
	cType     typ
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

func (e *enum) enrich(class *class, cursor clang.Cursor, api api) {
	e.cType = typFromClangType(cursor.EnumDeclIntegerType(), api)
	if class != nil {
		e.goName = class.goName + e.CName
	} else {
		e.goName = stripSkPrefix(e.CName)
	}
	e.class = class
	e.doc = cursor.RawCommentText()
	if class != nil {
		e.doc = strings.Replace(e.doc, fmt.Sprintf("\\enum %s::%s", class.CName, e.CName), "", 1)
	}

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			name = stripKPrefix(name)
			name = strings.TrimSuffix(name, "_"+e.CName)
			doc := cursor.RawCommentText()
			doc = strings.Replace(doc, "!<", "", 1)

			constant := enumConstant{
				goName: e.goName + name,
				// Assume that the value is signed, and will fit in an int64.
				// Could use cursor.EnumDeclIntegerType().Kind() to find the type, but it rapidly
				// gets rather involved as the type could in theory be any integer type.
				value: cursor.EnumConstantDeclValue(),
				doc:   doc,
			}
			e.constants = append(e.constants, constant)
		}

		return clang.ChildVisit_Continue
	})
}

func (e enum) generate(g generator) {
	if e.goName == "" {
		fatalf("enum %s not initialised, perhaps misspelled in json", e.CName)
	}

	f := g.goFile

	f.writeDocComment(e.doc)
	f.writelnf("type %s int64", e.goName)
	f.writeln()

	f.writeln("const (")
	for _, constant := range e.constants {
		f.writeDocComment(constant.doc)
		f.writelnf("%s %s = %d", constant.goName, e.goName, constant.value)
	}
	f.writeln(")")
	f.writeln()
}
