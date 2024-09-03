package generate

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

type translationUnit struct {
	headerFile string
	clang.TranslationUnit
}

var clangResourceDir = sync.OnceValue[string](func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		panic(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		panic(errors.New("no output when getting clang resource dir"))
	}
	if !strings.HasPrefix(resDir, "/") {
		panic(fmt.Errorf("expected clang resource dir to start with '/', but it %s", resDir))
	}

	return resDir
})

func newTranslationUnit(headerFile string) translationUnit {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-I", "./skia/skia/",
		"-x", "c++-header",
	}

	transUnit := translationUnit{
		headerFile: headerFile,
	}
	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(headerFile, parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies, &transUnit.TranslationUnit)
	if errCode != clang.Error_Success {
		panic(errCode)
	}
	return transUnit
}

func (tu translationUnit) enrichApi(api *api) {
	tu.TranslationUnitCursor().Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		file, _, _, _ := cursor.Location().FileLocation()
		if file.Name() != tu.headerFile {
			// Skip declarations from other files.
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			if !cursorHasChildren(cursor) {
				// Skip forward declarations.
				break
			}
			if class, ok := api.findClass(cursor.Spelling()); ok {
				class.enrich(cursor, *api)
			}

		case clang.Cursor_EnumDecl:
			// TODO

		case clang.Cursor_StructDecl:
			// TODO
		}

		return clang.ChildVisit_Continue
	})

}

func cursorHasChildren(cursor clang.Cursor) bool {
	childCount := 0
	cursor.Visit(func(_, _ clang.Cursor) (status clang.ChildVisitResult) {
		childCount++
		return clang.ChildVisit_Break
	})

	return childCount > 0
}

type speller interface {
	Spelling() string
}

const dumpIndent = "  "

func init() {
	// Prevent 'unused' warning if dumpCursor is not called.
	// Avoids having to comment out the function when it's not used.
	dumpCursor(clang.Cursor{}, "dummy")
}

func dumpCursor(cursor clang.Cursor, prefix string) {
	if prefix == "dummy" {
		return
	}

	fmt.Printf("%s%s : %s : %s\n", prefix,
		cursor.Kind().Spelling(),
		cursor.Spelling(),
		cursor.Type().Kind())

	dumpSpelling(prefix, "Type", cursor.Type())
	if cursor.AccessSpecifier() != clang.AccessSpecifier_Invalid {
		dumpSpelling(prefix, "AccessSpecifier", cursor.AccessSpecifier())
	}
	dumpCursorValue(prefix, "DisplayName", cursor.DisplayName())
	if cursor.CXXRecord_IsAbstract() {
		fmt.Println(prefix, "Abstract")
	}

	cursor.Visit(func(cursor, _ clang.Cursor) (status clang.ChildVisitResult) {
		dumpCursor(cursor, prefix+dumpIndent)
		return clang.ChildVisit_Continue
	})
}

func dumpCursorValue(prefix string, name string, value string) {
	if len(value) > 0 {
		fmt.Printf("%s%s = %s\n", prefix+dumpIndent+dumpIndent, name, value)
	}
}

func dumpSpelling(prefix string, name string, speller speller) {
	value := speller.Spelling()
	dumpCursorValue(prefix, name, value)
}
