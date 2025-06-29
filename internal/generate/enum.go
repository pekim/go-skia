package generate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	cppName   string
	cType     typ
	clangType clang.Type
	goName    string
	record    *record
	doc       string
	constants []enumConstant
	enriched1 bool
	enriched2 bool
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
		e.goName = record.goName + e.cppName
	} else {
		e.goName = stripSkPrefix(e.cppName)
	}
	e.goName = goExportedName(e.goName)
	e.record = record
	e.doc = docComment(cursor.ParsedComment())

	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			name = stripKPrefix(name)
			name = strings.TrimSuffix(name, "_"+e.cppName)
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

	e.enriched1 = true
}

func (e *enum) enrich2(api api) {
	if !e.enriched1 {
		fatalf("enum %s has not been phase 1 enriched", e.cppName)
	}

	e.cType = mustTypFromClangType(e.clangType, api, "")
	e.goType = e.cType.goName
	e.enriched2 = true
	e.doc = addDocCommentLinks(e.doc, api)
}

func (e enum) generate(g generator) {
	if len(e.constants) == 0 {
		fmt.Printf("warning: enum %s has no constants\n", e.cppName)
	}
	if !e.enriched2 {
		fatalf("enum %s has not been enriched", e.cppName)
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
