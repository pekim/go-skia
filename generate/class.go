package generate

type class struct {
	cName  string
	goName string
	enums  []enum
}

func (c class) generate(fileGo *genFile) {
	fileGo.writelnf("type %s struct {", c.goName)
	fileGo.writeln("  skia unsafe.Pointer")
	fileGo.writeln("}")
	fileGo.writeln("")

	for _, enum := range c.enums {
		enum.generate(fileGo)
	}
}
