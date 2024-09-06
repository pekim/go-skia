package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type struct_ struct {
	cursor      clang.Cursor
	CName       string   `json:"name"`
	Methods     []method `json:"methods"`
	fields      []field
	size        int
	goName      string
	cStructName string
	doc         string
	enriched    bool
}

func (s *struct_) enrich1(cursor clang.Cursor) {
	s.cursor = cursor
	s.goName = stripSkPrefix(s.CName)
	s.cStructName = fmt.Sprintf("sk_%s", s.CName)
	s.doc = cursor.RawCommentText()
	s.doc = strings.Replace(s.doc, fmt.Sprintf("\\class %s", s.CName), "", 1)
	s.size = int(cursor.Type().SizeOf())
	s.enriched = true
}

func (s *struct_) enrich2(api api) {
	s.cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {

		// case clang.Cursor_CXXMethod:
		// 	if method, ok := s.findMethod(cursor.Spelling()); ok {
		// 		method.enrich(api, s, cursor)
		// 	}

		case clang.Cursor_FieldDecl:
			s.fields = append(s.fields, newField(cursor, api))

		}

		return clang.ChildVisit_Continue
	})
	// fmt.Println(c.CName, c.fields)
	// fmt.Println(c.cursor.Type().SizeOf())
}

func (s *struct_) findMethod(name string) (*method, bool) {
	for i, method := range s.Methods {
		if method.CName == name {
			return &s.Methods[i], true
		}
	}
	return nil, false
}

func (s struct_) generate(g generator) {
	if !s.enriched {
		fatalf("struct %s has not been enriched", s.CName)
	}

	f := g.goFile
	f.writeDocComment(s.doc)
	f.writelnf("type %s struct {", s.goName)
	f.writeln("  sk unsafe.Pointer")
	f.writeln("}")
	f.writeln()

	s.generateCStruct(g)

	for _, method := range s.Methods {
		method.generate(g)
	}

	g.headerFile.writeln()
}

func (s struct_) generateCStruct(g generator) {
	f := g.headerFile

	f.writeln("typedef struct {")

	offset := 0
	for i, field := range s.fields {
		fieldOffset := (field.offset / 8)
		if fieldOffset > offset {
			f.writelnf("  uchar pad_%d[%d];", i, fieldOffset-offset)
			offset = fieldOffset
		}
		if fieldOffset < offset {
			f.writeln("  // TODO misalignment (perhaps there are bitfields around here?)")
		}

		name := field.name
		if strings.HasPrefix(field.name, "f") && field.public {
			name = name[1:]
		}

		if field.typ != "" {
			f.writelnf("  %s %s;", field.typ, name)
		} else {
			f.writelnf("  uchar %s[%d];", name, field.size)
		}

		offset += field.size
	}

	if offset < s.size {
		f.writelnf("  uchar pad_%d[%d];", len(s.fields), s.size-offset)
	}

	f.writelnf("} %s;", s.cStructName)
}
