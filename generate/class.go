package generate

import (
	"slices"

	"github.com/go-clang/clang-v15/clang"
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
	isPublic bool
}

func newClass(cursor clang.Cursor) class {
	name := cursor.Spelling()
	c := class{
		cName:    name,
		goName:   trimSkiaPrefix(name),
		abstract: cursor.CXXRecord_IsAbstract(),
		comment:  parsedCommentToGoComment(cursor.ParsedComment()),
	}

	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_CXXAccessSpecifier:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				c.isPublic = true
			}

		case clang.Cursor_Constructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				c.ctors = append(c.ctors, newClassCtor(&c, cursor))
			}

		case clang.Cursor_Destructor:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				c.dtors = append(c.dtors, classDtor{class: c})
			}

		case clang.Cursor_EnumDecl:
			c.enums = append(c.enums, newEnum(cursor, &c))
		}

		return clang.ChildVisit_Continue
	})

	return c
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
