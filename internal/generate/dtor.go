package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type recordDtor struct {
	cFuncName string
	record    *record
	doc       string
}

func newRecordDtor(record *record, cursor *clang.Cursor) recordDtor {
	d := recordDtor{record: record}
	if cursor != nil {
		d.doc = docComment(cursor.ParsedComment())

	}
	return d
}

func (d *recordDtor) generate(g generator) {
	if d.record == nil {
		// There is no destructor for the record.
		return
	}

	d.generateGo(g)
	d.generateHeader(g)
	d.generateCpp(g)
}

func (d *recordDtor) generateGo(g generator) {
	d.cFuncName = fmt.Sprintf("misk_delete_%s", d.record.cppName)

	f := g.goFile
	f.write(d.doc)
	f.writelnf("func (o *%s) Delete() {", d.record.goName)
	f.writelnf("   C.%s(o.sk)", d.cFuncName)
	f.writeln("  o.sk = nil")
	f.writeln("}")
	f.writeln()
}

func (d *recordDtor) generateHeader(g generator) {
	g.headerFile.writelnf("void  %s(%s *obj);", d.cFuncName, d.record.cStructName)
}

func (d *recordDtor) generateCpp(g generator) {
	f := g.cppFile
	f.writelnf("void  %s(%s *obj) {", d.cFuncName, d.record.cStructName)
	f.writelnf("  delete reinterpret_cast<%s*>(obj);", d.record.qualifiedCppName())
	f.writeln("}")
	f.writeln()
}
