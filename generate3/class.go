package generate

import (
	"fmt"
	"strings"

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
	c.doc = strings.Replace(c.doc, fmt.Sprintf("* \\class %s", c.Name), "", 1)

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

func (c class) generate(g generator) {
	c.generateGo(g)
}

func (c class) generateGo(g generator) {
	f := g.goFile

	f.writeComment(c.doc)
	f.writelnf("type %s class", c.goName)
	f.writeln("")

	for _, enum := range c.Enums {
		enum.generate(g)
	}
}
