package generate

import "fmt"

type class struct {
	cName  string
	goName string
	ctors  []classCtor
	enums  []enum
}

func (c class) generate(g *generator) {
	g.goFile.writelnf("type %s struct {", c.goName)
	g.goFile.writeln("  skia unsafe.Pointer")
	g.goFile.writeln("}")
	g.goFile.writeln("")

	for _, ctor := range c.ctors {
		ctor.generate(g)
	}

	for _, enum := range c.enums {
		enum.generate(g)
	}
}

type classCtor struct {
	class      class
	paramCount int
}

func (c classCtor) generate(g *generator) {
	if c.paramCount > 0 {
		return
	}

	cFuncName := fmt.Sprintf("skia_new_%s", c.class.cName)

	g.goFile.writelnf("func New%s() %s {", c.class.goName, c.class.goName)
	g.goFile.writelnfTrim(`
		c := C.%s()
		return %s {
		  skia: c,
		}
	`,
		cFuncName,
		c.class.goName,
	)
	g.goFile.writeln("}")

	g.headerFile.writelnf("void *%s();", cFuncName)

	g.cppFile.writelnf(`
		void *%s()
		{
			return (void *)(42);
		}
	`, cFuncName)
	/*
		return reinterpret_cast<void*>(new SkPaint());
	*/
}
