package generate

import (
	"fmt"
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

func (e enum) generate(fileGo *genFile) {
	if len(e.constants) == 0 {
		return
	}

	// generate the enum type
	goName := e.name
	if e.class != nil {
		goName = e.class.goName + e.name
	}
	fileGo.writelnf("type %s int64", goName)
	fileGo.writeln("")

	// generate the enum's constants
	fileGo.writeln("const (")
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

		fileGo.writelnf("  %s %s = %d", name, goName, constant.value)
	}
	fileGo.writeln(")")

	fileGo.writeln("")
}
