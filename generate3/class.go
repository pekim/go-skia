package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type class struct {
	Name   string       `json:"name"`
	Ctors  []*classCtor `json:"constructors"`
	Enums  []enum       `json:"enums"`
	goName string
	doc    string
}

func (c *class) enrich(cursor clang.Cursor) {
	c.goName = stripSkPrefix(c.Name)
	c.doc = cursor.RawCommentText()
	c.doc = strings.Replace(c.doc, fmt.Sprintf("* \\class %s", c.Name), "", 1)

	var ctorCursors []clang.Cursor
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_Constructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				ctorCursors = append(ctorCursors, cursor)
			}

		case clang.Cursor_EnumDecl:
			if enum, ok := c.findEnum(cursor.Spelling()); ok {
				enum.enrich(c, cursor)
			}
		}

		return clang.ChildVisit_Continue
	})

	if len(c.Ctors) != len(ctorCursors) {
		panic(fmt.Sprintf("class %s has %d ctors, but expected %d", c.Name, len(ctorCursors), len(c.Ctors)))
	}
	for i, cursor := range ctorCursors {
		ctor := c.Ctors[i]
		if ctor != nil {
			ctor.enrich(c, cursor)
		}
	}
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

	for _, ctor := range c.Ctors {
		if ctor != nil {
			ctor.generate(g)
		}
	}

	for _, enum := range c.Enums {
		enum.generate(g)
	}
}
