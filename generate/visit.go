package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

func (g *generator) visit() {
	g.transUnit.TranslationUnitCursor().Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_ClassDecl:
			class := newClass(cursor)
			if class.isPublic {
				g.classes = append(g.classes, class)
			}

		case clang.Cursor_EnumDecl:
			if !cursor.IsAnonymous() {
				g.enums = append(g.enums, newEnum(cursor, nil))
			}
		}

		return clang.ChildVisit_Continue
	})
}
