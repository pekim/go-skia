package generate

import (
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	Name      string `json:"name"`
	goName    string
	class     *class
	constants []enumConstant
}

type enumConstant struct {
	goName string
	value  int64
}

func (e *enum) enrich(class *class, cursor clang.Cursor) {
	e.goName = class.goName + e.Name
	e.class = class

	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		switch cursor.Kind() {
		case clang.Cursor_EnumConstantDecl:
			name := cursor.Spelling()
			name = stripKPrefix(name)
			name = strings.TrimSuffix(name, "_"+e.Name)

			constant := enumConstant{
				goName: e.goName + name,
				value:  cursor.EnumConstantDeclValue(),
			}
			e.constants = append(e.constants, constant)
		}

		return clang.ChildVisit_Continue
	})
}
