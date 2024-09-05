package generate

func (g *generator) generate() {
	g.goFile = newFileGo()
	defer g.goFile.finish()

	g.headerFile = newFileHeader()
	defer g.headerFile.finish()

	g.cppFile = newFileCpp()
	defer g.cppFile.finish()

	for _, class := range g.classes {
		class.generate(g)
	}

	for _, enum := range g.enums {
		enum.generate(g)
	}
}
