package generate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	class     *class
	name      string
	comment   string
	constants []enumConstant
}

type enumConstant struct {
	name    string
	value   int64
	comment string
}

func newEnum(cursor clang.Cursor, class *class) enum {
	e := enum{
		class:   class,
		name:    cursor.Spelling(),
		comment: parsedCommentToGoComment(cursor.ParsedComment()),
	}
	cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		if cursor.Kind() == clang.Cursor_EnumConstantDecl {
			e.constants = append(e.constants, enumConstant{
				name:    cursor.Spelling(),
				value:   cursor.EnumConstantDeclValue(),
				comment: parsedCommentToGoComment(cursor.ParsedComment()),
			})
		}
		return clang.ChildVisit_Continue
	})

	return e
}

func (e enum) generate(g *generator) {
	if len(e.constants) == 0 {
		return
	}

	// Skip global enums that would result in a name conflict with a class enum.
	if slices.Contains([]string{"SkPathSegmentMask", "SkPathVerb"}, e.name) {
		return
	}

	// generate the enum type
	if e.comment != "" {
		g.goFile.writeln(e.comment)
	}
	goName := trimSkiaPrefix(e.name)
	if e.class != nil {
		goName = e.class.goName + e.name
	}
	g.goFile.writelnf("type %s int64", goName)
	g.goFile.writeln("")

	// generate the enum's constants
	g.goFile.writeln("const (")
	for _, constant := range e.constants {
		// skip legacy names that would generate duplicate names
		if strings.HasSuffix(constant.name, "BackendHandleAccess") {
			continue
		}

		name := constant.name
		name = strings.TrimPrefix(name, "k")
		name = strings.TrimSuffix(name, e.name)
		name = strings.TrimSuffix(name, "_")
		if strings.HasSuffix(e.name, "Flags") || strings.HasSuffix(e.name, "FlagsSet") {
			name = strings.TrimSuffix(name, "Flag")
			name = strings.TrimSuffix(name, "_")
		}
		name = fmt.Sprintf("%s_%s", goName, name)

		if constant.comment != "" {
			g.goFile.writeln(constant.comment)
		}
		g.goFile.writelnf("  %s %s = %d", name, goName, constant.value)
	}
	g.goFile.writeln(")")

	g.goFile.writeln("")
}
