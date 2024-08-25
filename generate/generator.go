package generate

import (
	"fmt"
	"path"

	"github.com/go-clang/clang-v15/clang"
)

type generator struct {
	transUnit clang.TranslationUnit
}

func (g *generator) parse() {
	resourcesDir, err := clangResourceDir()
	if err != nil {
		panic(err)
	}

	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-I", "./skia/skia/",
		"-x", "c++-header",
	}

	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2("./generate/skia.h", parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies, &g.transUnit)
	if errCode != clang.Error_Success {
		panic(errCode)
	}
}

func (g *generator) generate() {
	g.transUnit.TranslationUnitCursor().Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			g.generateClass(cursor)
		case clang.Cursor_EnumDecl:
			g.generateEnum(cursor)
		}

		return clang.ChildVisit_Continue
	})
}

func (g *generator) generateClass(cursor clang.Cursor) {
	fmt.Println("class", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("  ", cursor.Spelling(), cursor.Kind())
		switch cursor.Kind() {
		case clang.Cursor_EnumDecl:
			g.generateClassEnum(cursor)
		}

		return clang.ChildVisit_Continue
	})
}

func (g *generator) generateClassEnum(cursor clang.Cursor) {
	fmt.Println("  enum", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("    ", cursor.Spelling(), cursor.Kind(), cursor.EnumConstantDeclValue())
		return clang.ChildVisit_Continue
	})
}

func (g *generator) generateEnum(cursor clang.Cursor) {
	fmt.Println("enum", cursor.Spelling())
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		fmt.Println("  ", cursor.Spelling(), cursor.Kind(), cursor.EnumConstantDeclValue())
		return clang.ChildVisit_Continue
	})
}
