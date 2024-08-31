package generate

import (
	"slices"
)

type class struct {
	cName    string
	goName   string
	comment  string
	abstract bool
	ctors    []classCtor
	ctorN    int // for ctor name function uniqueness
	dtors    []classDtor
	enums    []enum
}

func (c *class) generate(g *generator) {
	if c.comment != "" {
		g.goFile.writeln(c.comment)
	}
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
