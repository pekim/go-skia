package generate

import (
	"fmt"

	"github.com/go-clang/clang-v15/clang"
)

type classCtor struct {
	Suffix     string `json:"suffix"` // suffix to append to Go ctor function name
	goFuncName string
	cFuncName  string
	class      *class
	doc        string
	params     []param
}

func (c *classCtor) enrich(class *class, cursor clang.Cursor) {
	c.class = class
	c.goFuncName = fmt.Sprintf("New%s%s", c.class.goName, c.Suffix)
	c.cFuncName = fmt.Sprintf("misk_new_%s%s", c.class.goName, c.Suffix)
	c.doc = cursor.RawCommentText()

	paramCount := int(cursor.NumArguments())
	c.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(arg)
		c.params[i] = param
	}
}

func (c *classCtor) generate(g generator) {
	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}

func (c classCtor) generateGo(g generator) {
	f := g.goFile

	f.writeDocComment(c.doc)
	f.writelnf("func %s() %s {", c.goFuncName, c.class.goName)
	f.writelnf("  return %s(C.%s())", c.class.goName, c.cFuncName)
	f.writeln("}")
	f.writeln()
}

func (c classCtor) generateHeader(g generator) {
	f := g.headerFile

	f.writelnf("void * %s();", c.cFuncName)
}

func (c classCtor) generateCpp(g generator) {
	f := g.cppFile

	f.writelnf("void * %s() {", c.cFuncName)
	f.writelnf("  return reinterpret_cast<void*>(new %s());", c.class.Name)
	f.writeln("}")
	f.writeln()
}
