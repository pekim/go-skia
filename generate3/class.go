package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type class struct {
	cursor clang.Cursor
	CName  string       `json:"name"`
	Ctors  []*classCtor `json:"constructors"`
	Enums  []enum       `json:"enums"`
	dtor   classDtor
	goName string
	doc    string
}

func (c *class) enrich1(cursor clang.Cursor) {
	c.cursor = cursor
	c.goName = stripSkPrefix(c.CName)
	c.doc = cursor.RawCommentText()
	c.doc = strings.Replace(c.doc, fmt.Sprintf("\\class %s", c.CName), "", 1)
}

func (c *class) enrich2(api api) {
	var ctorCursors []clang.Cursor
	c.cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_Constructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				ctorCursors = append(ctorCursors, cursor)
			}

		case clang.Cursor_Destructor:
			c.dtor = newClassDtor(c, cursor)

		case clang.Cursor_EnumDecl:
			if enum, ok := c.findEnum(cursor.Spelling()); ok {
				enum.enrich(c, cursor)
			}
		}

		return clang.ChildVisit_Continue
	})

	if len(c.Ctors) != len(ctorCursors) {
		fatalf("class %s has %d ctors, but expected %d", c.CName, len(ctorCursors), len(c.Ctors))
	}
	for i, cursor := range ctorCursors {
		ctor := c.Ctors[i]
		if ctor != nil {
			ctor.enrich(api, c, cursor)
		}
	}
}

func (c *class) findEnum(name string) (*enum, bool) {
	for i, enum := range c.Enums {
		if enum.CName == name {
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

	f.writeDocComment(c.doc)
	f.writelnf("type %s struct {", c.goName)
	f.writeln("  sk unsafe.Pointer")
	f.writeln("}")
	f.writeln()

	for _, ctor := range c.Ctors {
		if ctor != nil {
			ctor.generate(g)
		}
	}
	c.dtor.generate(g)

	for _, enum := range c.Enums {
		enum.generate(g)
	}
}
