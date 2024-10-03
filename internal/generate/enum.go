package generate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	CppName   string `json:"name"`
	cType     typ
	clangType clang.Type
	goName    string
	record    *record
	doc       string
	constants []enumConstant
	enriched  bool
	goType    string
}

type enumConstant struct {
	goName string
	value  int64
	doc    string
}

func (e *enum) enrich1(record *record, cursor clang.Cursor) {
	e.clangType = cursor.EnumDeclIntegerType()
	if record != nil {
		e.goName = record.goName + e.CppName
	} else {
		e.goName = stripSkPrefix(e.CppName)
	}
	e.goName = goExportedName(e.goName)
	e.record = record
	e.doc = docComment(cursor.ParsedComment())

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			name = stripKPrefix(name)
			name = strings.TrimSuffix(name, "_"+e.CppName)
			doc := docComment(cursor.ParsedComment())

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

func (e *enum) enrich2(api api) {
	e.cType = mustTypFromClangType(e.clangType, api, "")
	e.goType = e.cType.goName
	e.enriched = true
	e.doc = addDocCommentLinks(e.doc, api)
}

func (e enum) generate(g generator) {
	if len(e.constants) == 0 {
		fmt.Printf("warning: enum %s has no constants\n", e.CppName)
	}
	if !e.enriched {
		fatalf("enum %s has not been enriched", e.CppName)
	}

	f := g.goFile

	f.write(e.doc)
	f.writelnf("type %s %s", e.goName, e.goType)
	f.writeln()

	f.writeln("const (")
	for _, constant := range e.constants {
		value := strconv.Itoa(int(constant.value))
		if e.goType == "bool" {
			value = "false"
			if constant.value != 0 {
				value = "true"
			}
		}

		f.write(constant.doc)
		f.writelnf("%s %s = %s", constant.goName, e.goName, value)
	}
	f.writeln(")")
	f.writeln()
}
