package generate

import (
	"fmt"
	"strings"

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

func (c *classCtor) enrich(api api, class *class, cursor clang.Cursor) {
	c.class = class
	c.goFuncName = fmt.Sprintf("New%s%s", c.class.goName, c.Suffix)
	c.cFuncName = fmt.Sprintf("misk_new_%s%s", c.class.goName, c.Suffix)
	c.doc = cursor.RawCommentText()

	paramCount := int(cursor.NumArguments())
	c.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(arg, api)
		c.params[i] = param
	}
}

func (c classCtor) generate(g generator) {
	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}

func (c classCtor) generateGo(g generator) {
	f := g.goFile

	params := make([]string, len(c.params))
	cVars := make([]string, len(c.params))
	cArgs := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typ.goName)
		cVars[i] = param.cgoVar
		cArgs[i] = param.cgoName
	}

	f.writeDocComment(c.doc)
	f.writelnf("func %s(%s) %s {", c.goFuncName, strings.Join(params, ", "), c.class.goName)
	f.writeln(strings.Join(cVars, "\n"))
	f.writelnf("  retC := C.%s(%s)", c.cFuncName, strings.Join(cArgs, ", "))
	f.writelnf("  return %s{sk: unsafe.Pointer(retC)}", c.class.goName)
	f.writeln("}")
	f.writeln()
}

func (c classCtor) generateHeader(g generator) {
	f := g.headerFile

	params := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = param.cParam
	}
	f.writelnf("void * %s(%s);", c.cFuncName, strings.Join(params, ", "))
}

func (c classCtor) generateCpp(g generator) {
	f := g.cppFile

	params := make([]string, len(c.params))
	args := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = param.cParam
		args[i] = param.cArg
	}

	f.writelnf("void * %s(%s) {", c.cFuncName, strings.Join(params, ", "))
	f.writelnf("  return reinterpret_cast<void*>(new %s(%s));", c.class.Name, strings.Join(args, ", "))
	f.writeln("}")
	f.writeln()
}
