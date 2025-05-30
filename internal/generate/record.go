package generate

import (
	"fmt"
	"regexp"

	"github.com/go-clang/clang-v15/clang"
)

var reSkRefCnt regexp.Regexp

func init() {
	reSkRefCnt = *regexp.MustCompile(`^Sk\w*RefCnt`)
}

// record represents a class or struct.
type record struct {
	cursor            clang.Cursor
	CppName           string        `json:"name"`
	Ctors             []*recordCtor `json:"constructors"`
	DtorCreate        bool          `json:"dtorCreate"` // create a destructor, even if there isn't one in the AST
	Enums             []enum        `json:"enums"`
	Methods           []callable    `json:"methods"`
	Records           []record      `json:"records"`
	NoWrapper         bool          `json:"noWrapper"`
	As                []string      `json:"as"`
	asRecords         []*record
	dtor              recordDtor
	fields            []field
	size              int
	isClass           bool // false for struct, true for class
	goName            string
	cStructName       string
	derivedFromRefCnt bool
	doc               string
	parent            *record
	enriched1         bool
	enriched2         bool
}

func (r *record) enrich1(cursor clang.Cursor, parent *record) {
	parentCppName := ""
	if parent != nil {
		parentCppName = parent.CppName
	}
	r.cursor = cursor
	r.goName = stripSkPrefix(parentCppName) + stripSkPrefix(r.CppName)
	r.cStructName = fmt.Sprintf("sk_%s%s", parentCppName, r.CppName)
	r.doc = docComment(cursor.ParsedComment())
	r.size = int(cursor.Type().SizeOf())
	r.parent = parent
	r.enriched2 = true
	r.isClass = cursor.Kind() == clang.Cursor_ClassDecl

	var ctorsEnriched int
	dtorCreated := false
	dtorNotPublic := false
	r.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_Constructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				if len(r.Ctors) <= ctorsEnriched {
					fatalf("record %s only has %d ctors configured but at least %d encountered", r.CppName, len(r.Ctors), ctorsEnriched+1)
				}
				ctor := r.Ctors[ctorsEnriched]
				if ctor != nil {
					r.Ctors[ctorsEnriched].enrich1(r, &cursor)
				}
				ctorsEnriched++
			}

		case clang.Cursor_Destructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				r.dtor = newRecordDtor(r, &cursor)
				dtorCreated = true
			} else {
				dtorNotPublic = true
			}

		case clang.Cursor_EnumDecl:
			if enum, ok := r.findEnum(cursor.Spelling()); ok {
				enum.enrich1(r, cursor)
			}

		case clang.Cursor_CXXBaseSpecifier:
			if reSkRefCnt.MatchString(cursor.Spelling()) {
				r.derivedFromRefCnt = true
			}

		case clang.Cursor_CXXMethod:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				if method, ok := r.findMethod(cursor.Spelling()); ok {
					method.enrich1(r, cursor)
				}
			}

		case clang.Cursor_FieldDecl:
			r.fields = append(r.fields, newField(cursor))

		case
			clang.Cursor_StructDecl,
			clang.Cursor_ClassDecl:

			if record, ok := r.findRecord(cursor.Spelling()); ok {
				record.enrich1(cursor, r)
			}
		}

		return clang.ChildVisit_Continue
	})

	if len(r.Ctors) != ctorsEnriched {
		if len(r.Ctors) == 1 && ctorsEnriched == 0 {
			// Enrich a default constructor.
			r.Ctors[0].enrich1(r, nil)
		} else {
			fatalf("record %s has %d ctors, but expected %d", r.CppName, ctorsEnriched, len(r.Ctors))
		}
	}

	// Create dtors that were not explicitly present in the AST.
	if r.DtorCreate || (!dtorCreated && !dtorNotPublic) && len(r.Ctors) > 0 && !r.NoWrapper {
		r.dtor = newRecordDtor(r, nil)
	}

	r.enriched1 = true
}

func (r *record) enrich2(api api) {
	if !r.enriched1 {
		fatalf("record %s has not been phase 1 enriched", r.CppName)
	}

	r.doc = addDocCommentLinks(r.doc, api)

	for i := range r.Enums {
		enum := &r.Enums[i]
		enum.enrich2(api)
	}

	for i := range r.Records {
		record := &r.Records[i]
		record.enrich2(api)
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
		method.enrich2(r, api)
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

func (r record) findRecord(name string) (*record, bool) {
	for i, record := range r.Records {
		if record.CppName == name {
			return &r.Records[i], true
		}
	}
	return nil, false
}

func (r record) qualifiedCppName() string {
	if r.parent == nil {
		return r.CppName
	}
	return r.parent.CppName + "::" + r.CppName
}

func (r record) generate(g generator) {
	if !r.enriched2 {
		fatalf("record %s has not been enriched", r.CppName)
	}

	r.generateGoType(g)
	r.generateFieldsMethods(g)
	r.generateNilMethod(g)
	r.generateAsMethods(g)
	r.generateUnref(g)

	for _, ctor := range r.Ctors {
		if ctor != nil {
			ctor.generate(g)
		}
	}
	r.dtor.generate(g)

	for _, method := range r.Methods {
		method.record = &r
		method.generate(g)
	}

	for _, enum := range r.Enums {
		enum.generate(g)
	}

	for _, record := range r.Records {
		record.generate(g)
	}

	r.generateHeaderFile(g)
	r.generateCppFile(g)
}

func (r record) generateGoType(g generator) {
	f := g.goFile
	f.write(r.doc)
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
	f.write(field.doc)
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
	f.writelnf("// If it is nil is may indicate that the %s has not been created, or has been deleted with [%s.Delete].",
		r.goName, r.goName)
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

func (r record) generateUnref(g generator) {
	if !r.derivedFromRefCnt {
		return
	}

	f := g.goFile
	f.writelnf("func (o *%s) Unref() {", r.goName)
	f.writelnf("  C.misk_unref_%s(o.sk)", r.CppName)
	f.writeln("  o.sk = nil")
	f.writeln("}")
}

func (r record) generateHeaderFile(g generator) {
	f := g.headerFile
	if r.derivedFromRefCnt {
		f.writelnf("void misk_unref_%s(%s *c_obj);", r.CppName, r.cStructName)
	}
	f.writeln()
}

func (r record) generateCppFile(g generator) {
	f := g.cppFile

	if !r.derivedFromRefCnt {
		return
	}

	f.writelnf("void misk_unref_%s(%s *c_obj) {", r.CppName, r.cStructName)
	f.writelnf("reinterpret_cast<%s *> (c_obj)->unref();", r.CppName)
	f.writeln("}")
	f.writeln()
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
	f.writelnf("extern int sk_sizeof_%s;", r.cStructName)
	f.writeln()
}

func (r record) generateCPPStructSize(g generator) {
	g.cppFile.writelnf("int sk_sizeof_%s = sizeof(%s);", r.cStructName, r.cStructName)
}

func (r record) generateStructSizeAssertion(g generator) {
	g.goFile.writelnf(`assertSizesMatch("%s", C.sizeof_%s, int(C.sk_sizeof_%s))`,
		r.goName, r.cStructName, r.cStructName)
}
