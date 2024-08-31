package generate

import (
	"fmt"
	"slices"
	"strings"
)

type enum struct {
	class     *class
	name      string
	constants []enumConstant
}

type enumConstant struct {
	name  string
	value int64
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

		g.goFile.writelnf("  %s %s = %d", name, goName, constant.value)
	}
	g.goFile.writeln(")")

	g.goFile.writeln("")
}
