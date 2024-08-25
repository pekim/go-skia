package generate

import (
	"fmt"

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
	fmt.Println("class", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("  ", cursor.Spelling(), cursor.Kind())
		switch cursor.Kind() {
		case clang.Cursor_EnumDecl:
			g.visitClassEnum(cursor)
		}

		return clang.ChildVisit_Continue
	})
}

func (g *generator) visitClassEnum(cursor clang.Cursor) {
	fmt.Println("  enum", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("    ", cursor.Spelling(), cursor.Kind(), cursor.EnumConstantDeclValue())
		return clang.ChildVisit_Continue
	})
}

func (g *generator) visitEnum(cursor clang.Cursor) {
	fmt.Println("enum", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("  ", cursor.Spelling(), cursor.Kind(), cursor.EnumConstantDeclValue())
		return clang.ChildVisit_Continue
	})
}
