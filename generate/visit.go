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
		cName:    name,
		goName:   strings.TrimPrefix(name, "Sk"),
		abstract: cursor.CXXRecord_IsAbstract(),
	}

	isPublic := false
	cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_CXXAccessSpecifier:
			if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
				isPublic = true
			}

		case clang.Cursor_Constructor:
			g.visitClassCtor(&class, cursor)

		case clang.Cursor_Destructor:
			g.visitClassDtor(&class, cursor)

		case clang.Cursor_EnumDecl:
			g.visitClassEnum(&class, cursor)
		}

		return clang.ChildVisit_Continue
	})

	if isPublic {
		g.classes = append(g.classes, class)
	}
}

func (g *generator) visitClassCtor(class *class, cursor clang.Cursor) {
	ctor := classCtor{
		class: class,
	}

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_ParmDecl:
			ctor.params = append(ctor.params, g.visitParamDecl(cursor))
		}

		return clang.ChildVisit_Continue
	})

	if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
		class.ctors = append(class.ctors, ctor)
	}
}

func (g *generator) visitParamDecl(cursor clang.Cursor) param {
	cName := cursor.Spelling()
	typ := cursor.Type()
	cDecl := typ.Spelling()
	p := newParam(cName, cDecl)

	p.kind = typ.Kind()
	p.array = typ.ArrayElementType().Kind() != clang.Type_Invalid
	p.arrayKind = typ.ArrayElementType().Kind()

	return p
}

func (g *generator) visitClassDtor(class *class, cursor clang.Cursor) {
	if cursor.AccessSpecifier() == clang.AccessSpecifier_Public {
		class.dtors = append(class.dtors, classDtor{class: *class})
	}
}

func (g *generator) visitClassEnum(class *class, cursor clang.Cursor) {
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

	class.enums = append(class.enums, enum)
}

func (g *generator) visitEnum(cursor clang.Cursor) {
	if cursor.IsAnonymous() {
		return
	}

	enum := enum{
		name: cursor.Spelling(),
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

	g.enums = append(g.enums, enum)
}
