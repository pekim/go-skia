package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type classDtor struct {
	cFuncName string
	class     *class
	doc       string
}

func newClassDtor(class *class, cursor clang.Cursor) classDtor {
	return classDtor{
		class: class,
		doc:   cursor.RawCommentText(),
	}
}

func (d *classDtor) generate(g generator) {
	if d.class == nil {
		// There is no destructor for the class.
		return
	}

	d.generateGo(g)
	d.generateHeader(g)
	d.generateCpp(g)
}

func (d *classDtor) generateGo(g generator) {
	d.cFuncName = fmt.Sprintf("misk_delete_%s", d.class.CName)

	f := g.goFile
	f.writeDocComment(d.doc)
	f.writelnf("func (o %s) Delete() {", d.class.goName)
	f.writelnf("   C.%s(o.sk)", d.cFuncName)
	f.writeln("}")
	f.writeln()
}

func (d *classDtor) generateHeader(g generator) {
	g.headerFile.writelnf("void  %s(%s *obj);", d.cFuncName, d.class.cStructName)
}

func (d *classDtor) generateCpp(g generator) {
	f := g.cppFile
	f.writelnf("void  %s(%s *obj) {", d.cFuncName, d.class.cStructName)
	f.writelnf("  delete reinterpret_cast<%s*>(obj);", d.class.CName)
	f.writeln("}")
	f.writeln()
}
