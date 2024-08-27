package generate

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type class struct {
	cName    string
	goName   string
	abstract bool
	ctors    []classCtor
	ctorN    int // for ctor name function uniqueness
	dtors    []classDtor
	enums    []enum
}

func (c *class) generate(g *generator) {
	g.goFile.writelnf("type %s struct {", c.goName)
	g.goFile.writeln("  skia unsafe.Pointer")
	g.goFile.writeln("}")
	g.goFile.writeln("")

	skipCtorsClasses := []string{
		// class has an implicitly deleted ctor
		"SkTableMaskFilter",
		// "undefined symbol: typeinfo for SkWStream"
		"SkNullWStream",
	}

	if !c.abstract && !slices.Contains(skipCtorsClasses, c.cName) {
		for _, ctor := range c.ctors {
			ctor.generate(g)
		}
	}

	for _, dtor := range c.dtors {
		dtor.generate(g)
	}

	for _, enum := range c.enums {
		enum.generate(g)
	}
}

type classCtor struct {
	class  *class
	params []param

	nameSuffix string
	cFuncName  string
}

func (c *classCtor) generate(g *generator) {
	for i, param := range c.params {
		if !param.supported() {
			// fmt.Println(c.class.cName, param.cName, param.cDecl_)
			return
		}

		if param.cName == "" {
			param.cName = fmt.Sprintf("p%d", i)
			param.goName = param.cName
		}
	}

	c.class.ctorN++
	if c.class.ctorN > 1 {
		c.nameSuffix = strconv.Itoa(c.class.ctorN)
	}
	c.cFuncName = fmt.Sprintf("skia_new_%s%s", c.class.cName, c.nameSuffix)

	c.generateGo(g)
	c.generateHeader(g)
	c.generateCpp(g)
}

type classDtor struct {
	class class
}

func (c classDtor) generate(g *generator) {
	cFuncName := fmt.Sprintf("skia_delete_%s", c.class.cName)

	g.goFile.writelnf(`
		func (o *%s) Delete() {
			C.%s(o.skia)
		}
	`,
		c.class.goName,
		cFuncName,
	)

	g.headerFile.writelnf("void %s(void *obj);", cFuncName)

	g.cppFile.writelnf(`
		void %s(void *obj)
		{
			delete reinterpret_cast<%s*>(obj);
		}
	`, cFuncName, c.class.cName)
}

func (c *classCtor) generateGo(g *generator) {
	goParams := makeParamsString(c.params, func(p param) string { return p.goDecl() })
	cArgNames := makeParamsString(c.params, func(p param) string { return p.goCName })

	cArgs := make([]string, len(c.params))
	for p, param := range c.params {
		cArgs[p] = param.goCArg()
	}
	allCArgs := strings.Join(cArgs, "\n")
	allCArgs += "\n"

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
	params := makeParamsString(c.params, func(p param) string { return p.cDecl() })
	g.headerFile.writelnf("void *%s(%s);", c.cFuncName, params)
}

func (c *classCtor) generateCpp(g *generator) {
	cParams := makeParamsString(c.params, func(p param) string { return p.cDecl() })
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
