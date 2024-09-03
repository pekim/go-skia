package generate

// func (cc classes) generate(g generator) {
// 	cc.generateGo(g)
// }

// func (cc classes) generateGo(g generator) {
// 	for _, class := range cc {
// 		g.goFile.docComment(class.Doc)
// 		g.goFile.writelnf("type %s struct {", class.goName)
// 		g.goFile.writeln("  skia unsafe.Pointer")
// 		g.goFile.writeln("}")
// 		g.goFile.writeln("")

// 		for _, enum := range class.Enums {
// 			class.generateEnum(g, enum)
// 			g.goFile.writeln("")
// 		}
// 	}
// }

// func (c class) generateEnum(g generator, enum enum) {
// 	g.goFile.docComment(enum.Doc)
// 	enumName := fmt.Sprintf("%s%s", c.goName, enum.Name)

// 	g.goFile.writelnf("type %s int", enumName)
// 	g.goFile.writeln("const    (")
// 	for i, constant := range enum.Constants {
// 		g.goFile.docComment(constant.Doc)
// 		g.goFile.writelnf("%s%s %s = %d", enumName, constant.Name, enumName, i)
// 	}
// 	g.goFile.writeln(")")
// }
