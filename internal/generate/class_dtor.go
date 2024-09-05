package generate

import "fmt"

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
