package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type recordCtor struct {
	Suffix     string `json:"suffix"` // suffix to append to Go ctor function name
	goFuncName string
	cFuncName  string
	record     *record
	doc        string
	params     []param
	enriched   bool
}

func (c *recordCtor) enrich1(record *record, cursor clang.Cursor) {
	c.record = record
	c.goFuncName = fmt.Sprintf("New%s%s", c.record.goName, c.Suffix)
	c.cFuncName = fmt.Sprintf("misk_new_%s%s", c.record.goName, c.Suffix)
	c.doc = cursor.RawCommentText()

	paramCount := int(cursor.NumArguments())
	c.params = make([]param, paramCount)
	for i := 0; i < paramCount; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(i, arg, c.params[i].ValueNil)
		c.params[i] = param
	}

	c.enriched = true
}

func (c recordCtor) enrich2(api api) {
	for i := range c.params {
		param := &c.params[i]
		param.enrich2(api)
	}
}

func (c recordCtor) generate(g generator) {
	if !c.enriched {
		fatalf("record %s ctor has not been enriched", c.record.CppName)
	}

	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}

func (c recordCtor) generateGo(g generator) {
	f := g.goFile

	params := make([]string, len(c.params))
	cVars := make([]string, len(c.params))
	cArgs := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = fmt.Sprintf("%s %s", param.goName, param.typGoName)
		cVars[i] = param.cgoVar
		cArgs[i] = param.cName
	}

	f.writeDocComment(c.doc)
	f.writelnf("func %s(%s) %s {", c.goFuncName, strings.Join(params, ", "), c.record.goName)
	f.writeln(strings.Join(cVars, "\n"))
	f.writelnf("  retC := C.%s(%s)", c.cFuncName, strings.Join(cArgs, ", "))
	if c.record.NoWrapper {
		f.writelnf("  return *(*%s)(unsafe.Pointer(&retC))", c.record.goName)

	} else {
		f.writelnf("  return %s{sk: retC}", c.record.goName)
	}
	f.writeln("}")
	f.writeln()
}

func (c recordCtor) generateHeader(g generator) {
	f := g.headerFile

	params := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = param.cParam
	}
	f.writelnf("%s * %s(%s);", c.record.cStructName, c.cFuncName, strings.Join(params, ", "))
}

func (c recordCtor) generateCpp(g generator) {
	f := g.cppFile

	params := make([]string, len(c.params))
	args := make([]string, len(c.params))
	for i, param := range c.params {
		params[i] = param.cParam
		args[i] = param.cppArg
	}

	f.writelnf("%s * %s(%s) {", c.record.cStructName, c.cFuncName, strings.Join(params, ", "))
	f.writelnf("  return reinterpret_cast<%s*>(new %s(%s));", c.record.cStructName, c.record.CppName, strings.Join(args, ", "))
	f.writeln("}")
	f.writeln()
}
