package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

func (g *generator) visit() {
	g.transUnit.TranslationUnitCursor().Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			g.visitClass(cursor)
		case clang.Cursor_EnumDecl:
			g.visitEnum(cursor)
		}

		return clang.ChildVisit_Continue
	})
}

func (g *generator) visitClass(cursor clang.Cursor) {
	name := cursor.Spelling()
	class := class{
		cName:  name,
		goName: strings.TrimPrefix(name, "Sk"),
	}
	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {

		switch cursor.Kind() {
		case clang.Cursor_EnumDecl:
			enum := g.visitClassEnum(&class, cursor)
			class.enums = append(class.enums, enum)
		}

		return clang.ChildVisit_Continue
	})
	g.classes = append(g.classes, class)
}

func (g *generator) visitClassEnum(class *class, cursor clang.Cursor) enum {
	enum := enum{
		class: class,
		name:  cursor.Spelling(),
	}
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_EnumConstantDecl {
			enum.constants = append(enum.constants, enumConstant{
				name:  cursor.Spelling(),
				value: cursor.EnumConstantDeclValue(),
			})
		}
		return clang.ChildVisit_Continue
	})

	return enum
}

func (g *generator) visitEnum(cursor clang.Cursor) {
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		return clang.ChildVisit_Continue
	})
}
