package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

func (g *generator) visit() {
	g.transUnit.TranslationUnitCursor().Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			// dumpCursor(cursor, "")
			class := newClass(cursor)
			if class.isPublic {
				g.classes = append(g.classes, class)
			}

		case clang.Cursor_EnumDecl:
			// dumpCursor(cursor, "")
			if !cursor.IsAnonymous() {
				g.enums = append(g.enums, newEnum(cursor, nil))
			}

		case clang.Cursor_StructDecl:
			// fmt.Println(cursor.DisplayName())
		}

		return clang.ChildVisit_Continue
	})
}

// type speller interface {
// 	Spelling() string
// }

// const dumpIndent = "  "

// func dumpCursor(cursor clang.Cursor, prefix string) {
// 	fmt.Printf("%s%s : %s : %s\n", prefix,
// 		cursor.Kind().Spelling(),
// 		cursor.Spelling(),
// 		cursor.Type().Kind())

// 	dumpSpelling(prefix, "Type", cursor.Type())
// 	if cursor.AccessSpecifier() != clang.AccessSpecifier_Invalid {
// 		dumpSpelling(prefix, "AccessSpecifier", cursor.AccessSpecifier())
// 	}
// 	dumpCursorValue(prefix, "DisplayName", cursor.DisplayName())
// 	if cursor.CXXRecord_IsAbstract() {
// 		fmt.Println(prefix, "Abstract")
// 	}

// 	cursor.Visit(func(cursor, _ clang.Cursor) (status clang.ChildVisitResult) {
// 		dumpCursor(cursor, prefix+dumpIndent)
// 		return clang.ChildVisit_Continue
// 	})
// }

// func dumpCursorValue(prefix string, name string, value string) {
// 	if len(value) > 0 {
// 		fmt.Printf("%s%s = %s\n", prefix+dumpIndent+dumpIndent, name, value)
// 	}
// }

// func dumpSpelling(prefix string, name string, speller speller) {
// 	value := speller.Spelling()
// 	dumpCursorValue(prefix, name, value)
// }
