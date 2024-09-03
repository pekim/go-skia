package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type class struct {
	Name   string `json:"name"`
	goName string
	doc    string
	Enums  []enum `json:"enums"`
}

func (c *class) enrich(cursor clang.Cursor) {
	c.goName = stripSkPrefix(c.Name)
	c.doc = cursor.RawCommentText()

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumDecl:
			if enum, ok := c.findEnum(cursor.Spelling()); ok {
				enum.enrich(c, cursor)
			}
		}

		return clang.ChildVisit_Continue
	})
}

func (c *class) findEnum(name string) (*enum, bool) {
	for i, enum := range c.Enums {
		if enum.Name == name {
			return &c.Enums[i], true
		}
	}
	return nil, false
}

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
