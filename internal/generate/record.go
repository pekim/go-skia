package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

// record represents a class or struct.
type record struct {
	cursor      clang.Cursor
	CppName     string        `json:"name"`
	Ctors       []*recordCtor `json:"constructors"`
	Enums       []enum        `json:"enums"`
	Methods     []callable    `json:"methods"`
	NoWrapper   bool          `json:"noWrapper"`
	As          []string      `json:"as"`
	asRecords   []*record
	dtor        recordDtor
	fields      []field
	size        int
	goName      string
	cStructName string
	doc         string
	enriched    bool
}

func (r *record) enrich1(cursor clang.Cursor) {
	r.cursor = cursor
	r.goName = stripSkPrefix(r.CppName)
	r.cStructName = fmt.Sprintf("sk_%s", r.CppName)
	r.doc = cursor.RawCommentText()
	r.doc = strings.Replace(r.doc, fmt.Sprintf("\\class %s", r.CppName), "", 1)
	r.doc = strings.Replace(r.doc, fmt.Sprintf("\\struct %s", r.CppName), "", 1)
	r.size = int(cursor.Type().SizeOf())
	r.enriched = true

	var ctorsEnriched int
	r.cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_Constructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				ctor := r.Ctors[ctorsEnriched]
				if ctor != nil {
					r.Ctors[ctorsEnriched].enrich1(r, cursor)
				}
				ctorsEnriched++
			}

		case clang.Cursor_Destructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				r.dtor = newRecordDtor(r, cursor)
			}

		case clang.Cursor_EnumDecl:
			if enum, ok := r.findEnum(cursor.Spelling()); ok {
				enum.enrich1(r, cursor)
			}

		case clang.Cursor_CXXMethod:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				if method, ok := r.findMethod(cursor.Spelling()); ok {
					method.enrich1(r, cursor)
				}
			}

		case clang.Cursor_FieldDecl:
			r.fields = append(r.fields, newField(cursor))

			// case clang.Cursor_VarDecl:
			// 	fmt.Println(r.CppName, cursor.Spelling())
		}

		return clang.ChildVisit_Continue
	})

	if len(r.Ctors) != ctorsEnriched {
		fatalf("record %s has %d ctors, but expected %d", r.CppName, ctorsEnriched, len(r.Ctors))
	}
}

func (r *record) enrich2(api api) {
	for i := range r.Enums {
		enum := &r.Enums[i]
		enum.enrich2(api)
	}

	for i := range r.fields {
		field := &r.fields[i]
		field.enrich2(api)
	}

	for i := range r.Ctors {
		ctor := r.Ctors[i]
		if ctor != nil {
			ctor.enrich2(api)
		}
	}

	for i := range r.Methods {
		method := &r.Methods[i]
		method.enrich2(api)
	}

	for _, as := range r.As {
		asRecord, ok := api.findRecord(as)
		if !ok {
			fatalf("failed to find record %s for As method for %s record", as, r.CppName)
		}
		r.asRecords = append(r.asRecords, asRecord)
	}
}

func (r *record) findEnum(name string) (*enum, bool) {
	for i, enum := range r.Enums {
		if enum.CppName == name {
			return &r.Enums[i], true
		}
	}
	return nil, false
}

func (r *record) findMethod(name string) (*callable, bool) {
	for i, method := range r.Methods {
		if method.CppName == name {
			return &r.Methods[i], true
		}
	}
	return nil, false
}

func (r record) generate(g generator) {
	if !r.enriched {
		fatalf("record %s has not been enriched", r.CppName)
	}

	r.generateGoType(g)
	r.generateFieldsMethods(g)
	r.generateNilMethod(g)
	r.generateAsMethods(g)

	for _, ctor := range r.Ctors {
		if ctor != nil {
			ctor.generate(g)
		}
	}
	r.dtor.generate(g)

	for _, method := range r.Methods {
		method.generate(g)
	}

	for _, enum := range r.Enums {
		enum.generate(g)
	}

	g.headerFile.writeln()
}

func (r record) generateGoType(g generator) {
	f := g.goFile
	f.writeDocComment(r.doc)
	if r.NoWrapper {
		f.writelnf("type %s C.%s", r.goName, r.cStructName)
	} else {
		f.writelnf("type %s struct {", r.goName)
		f.writelnf("  sk *C.%s", r.cStructName)
		f.writeln("}")
	}
	f.writeln()
}

func (r record) generateFieldsMethods(g generator) {
	for _, field := range r.fields {
		if !field.public || field.cType == "" {
			continue
		}

		r.generateFieldGetter(g, field)
		r.generateFieldSetter(g, field)
	}
}

func (r record) generateFieldGetter(g generator, field field) {
	methodName := goExportedName(field.name)
	if field.name[0] == 'f' {
		methodName = goExportedName(field.name[1:])
	}

	cStruct := "o"
	if !r.NoWrapper {
		cStruct = "o.sk"
	}

	f := g.goFile
	f.writeDocComment(field.doc)
	f.writelnf(`func (o %s) %s() %s {
			return %s(%s.%s)
		}`,
		r.goName, methodName, field.goType,
		field.goType, cStruct, field.name,
	)
	f.writeln()
}

func (r record) generateFieldSetter(g generator, field field) {
	fieldName := field.name
	if field.name[0] == 'f' {
		fieldName = field.name[1:]
	}
	methodName := goExportedName("Set" + fieldName)

	cStruct := "o"
	if !r.NoWrapper {
		cStruct = "o.sk"
	}

	f := g.goFile
	f.writelnf(`func (o *%s) %s(value %s) {
			%s.%s = C.%s(value)
		}`,
		r.goName, methodName, field.goType,
		cStruct, field.name, field.cType,
	)
	f.writeln()
}

func (r record) generateNilMethod(g generator) {
	if r.NoWrapper {
		return
	}

	f := g.goFile
	f.writeln("// IsNil returns true if the raw skia object pointer is nil.")
	f.writelnf("// If it is nil is may indicate that the %s has not been created.", r.goName)
	f.writelnf("func (o %s) IsNil() bool {", r.goName)
	f.writeln("  return o.sk == nil")
	f.writeln("}")
	f.writeln()
}

func (r record) generateAsMethods(g generator) {
	f := g.goFile

	for _, asRecord := range r.asRecords {
		methodName := fmt.Sprintf("As%s", asRecord.goName)
		f.writelnf("// %s converts the %s to a %s.", methodName, r.goName, asRecord.goName)
		f.writelnf("func (o %s) %s() %s {", r.goName, methodName, asRecord.goName)
		f.writelnf("  return %s{sk: (*C.%s)(unsafe.Pointer(o.sk))}", asRecord.goName, asRecord.cStructName)
		f.writeln("}")
		f.writeln()
	}
}

func (r record) generateCStruct(g generator) {
	f := g.headerFile

	f.writeln("typedef struct {")

	offset := 0
	for i, field := range r.fields {
		fieldOffset := (field.offset / 8)
		if fieldOffset > offset {
			f.writelnf("  uchar pad_%d[%d];", i, fieldOffset-offset)
			offset = fieldOffset
		}
		if fieldOffset < offset {
			f.writeln("  // TODO misalignment (perhaps there are bitfields around here?)")
		}

		name := field.name
		if field.cType != "" {
			f.writelnf("  %s %s;", field.cType, name)
		} else {
			f.writelnf("  uchar %s[%d];", name, field.size)
		}

		offset += field.size
	}

	if offset < r.size {
		f.writelnf("  uchar pad_%d[%d];", len(r.fields), r.size-offset)
	}

	f.writelnf("} %s;", r.cStructName)
	f.writeln()
}
