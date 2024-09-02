package generate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type classCtor struct {
	class      *class
	nameSuffix string
	cFuncName  string
	comment    string
	params     []param
}

func newClassCtor(class *class, cursor clang.Cursor) classCtor {
	ctor := classCtor{
		class:   class,
		comment: parsedCommentToGoComment(cursor.ParsedComment()),
	}

	numParam := int(cursor.NumArguments())
	ctor.params = make([]param, 0, numParam)
	for i := 0; i < numParam; i++ {
		arg := cursor.Argument(uint32(i))
		param := newParam(arg, i)
		ctor.params = append(ctor.params, param)
	}

	return ctor
}

func (c *classCtor) generate(g *generator) {
	c.class.ctorN++
	if c.class.ctorN > 1 {
		c.nameSuffix = strconv.Itoa(c.class.ctorN)
	}
	c.cFuncName = fmt.Sprintf("skia_new_%s%s", c.class.cName, c.nameSuffix)

	for _, param := range c.params {
		if param.isStruct {
			g.goFile.writelnf("// function %s not supported; param %s, is struct", c.cFuncName, param.cName)
			g.goFile.writeln("")
			return
		}

		supported, reason := param.supported()
		if !supported {
			g.goFile.writelnf("// function %s not supported; param %s, %s", c.cFuncName, param.cName, reason)
			g.goFile.writeln("")
			return
		}
	}

	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}

func (c *classCtor) generateGo(g *generator) {
	goParams := makeParamsString(c.params, func(p param) string { return p.goDecl() })
	cArgNames := makeParamsString(c.params, func(p param) string { return p.cgoName })

	cArgs := make([]string, len(c.params))
	for p, param := range c.params {
		cArgs[p] = param.goCArg()
	}
	allCArgs := strings.Join(cArgs, "\n")
	allCArgs += "\n"

	if c.comment != "" {
		g.goFile.write(c.comment)
	}
	g.goFile.writelnf(`
		func New%s%s(%s) %s {
			%s
			return %s {
		  	skia: C.%s(%s),
			}
		}
	`,
		c.class.goName, c.nameSuffix, goParams, c.class.goName,
		allCArgs,
		c.class.goName,
		c.cFuncName, cArgNames,
	)
}

func (c *classCtor) generateHeader(g *generator) {
	params := makeParamsString(c.params, func(p param) string { return p.cParamDecl() })
	g.headerFile.writelnf("void *%s(%s);", c.cFuncName, params)
}

func (c *classCtor) generateCpp(g *generator) {
	cParams := makeParamsString(c.params, func(p param) string { return p.cParamDecl() })
	cArgs := makeParamsString(c.params, func(p param) string { return p.cArg() })

	g.cppFile.writelnf(`
		void *%s(%s)
		{
			return reinterpret_cast<void*>(new %s(%s));
		}
	`,
		c.cFuncName,
		cParams,
		c.class.cName, cArgs,
	)
}
