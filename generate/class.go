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
}

func (c *classCtor) generate(g *generator) {
	for _, param := range c.params {
		if !param.supported() {
			return
		}
	}

	nameSuffix := ""
	c.class.ctorN++
	if c.class.ctorN > 1 {
		nameSuffix = strconv.Itoa(c.class.ctorN)
	}
	cFuncName := fmt.Sprintf("skia_new_%s%s", c.class.cName, nameSuffix)

	var goParams = make([]string, len(c.params))
	for p, param := range c.params {
		goParams[p] = param.goDecl()
	}
	allGoParams := strings.Join(goParams, ", ")
	var goCArgs = make([]string, len(c.params))
	for p, param := range c.params {
		goCArgs[p] = param.goCArg()
	}
	allGoCArgs := strings.Join(goCArgs, ", ")
	g.goFile.writelnfTrim(`
		func New%s%s(%s) %s {
			return %s {
		  	skia: C.%s(%s),
			}
		}
	`,
		c.class.goName, nameSuffix, allGoParams, c.class.goName,
		c.class.goName,
		cFuncName, allGoCArgs,
	)

	var cParams = make([]string, len(c.params))
	for p, param := range c.params {
		cParams[p] = param.cDecl()
	}
	allCParams := strings.Join(cParams, ", ")
	var cArgs = make([]string, len(c.params))
	for p, param := range c.params {
		cArgs[p] = param.cArg()
	}
	allCArgs := strings.Join(cArgs, ", ")
	g.headerFile.writelnf("void *%s(%s);", cFuncName, allCParams)

	g.cppFile.writelnf(`
		void *%s(%s)
		{
			return reinterpret_cast<void*>(new %s(%s));
		}
	`,
		cFuncName,
		allCParams,
		c.class.cName, allCArgs,
	)
}

type classDtor struct {
	class class
}

func (c classDtor) generate(g *generator) {
	cFuncName := fmt.Sprintf("skia_delete_%s", c.class.cName)

	g.goFile.writelnfTrim(`
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
