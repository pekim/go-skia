package generate

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed class.json
var classJson []byte

type classes []class

type class struct {
	Name   string `json:"name"`
	goName string
	Doc    []string `json:"doc"`
	Enums  []enum   `json:"enums"`
}

type enum struct {
	Name      string         `json:"name"`
	Doc       []string       `json:"doc"`
	Constants []enumConstant `json:"constants"`
}

type enumConstant struct {
	Name string   `json:"name"`
	Doc  []string `json:"doc"`
}

func loadClasses() classes {
	var classes classes
	err := json.Unmarshal(classJson, &classes)
	if err != nil {
		panic(err)
	}

	for i, class := range classes {
		classes[i].goName = stripSkPrefix(class.Name)
	}

	return classes
}

func (cc classes) generate(g generator) {
	cc.generateGo(g)
}

func (cc classes) generateGo(g generator) {
	for _, class := range cc {
		g.goFile.docComment(class.Doc)
		g.goFile.writelnf("type %s struct {", class.goName)
		g.goFile.writeln("}")
		g.goFile.writeln("")

		for _, enum := range class.Enums {
			class.generateEnum(g, enum)
			g.goFile.writeln("")
		}
	}
}

func (c class) generateEnum(g generator, enum enum) {
	g.goFile.docComment(enum.Doc)
	enumName := fmt.Sprintf("%s%s", c.goName, enum.Name)

	g.goFile.writelnf("type %s int", enumName)
	g.goFile.writeln("const    (")
	for i, constant := range enum.Constants {
		g.goFile.docComment(constant.Doc)
		g.goFile.writelnf("%s%s %s = %d", enumName, constant.Name, enumName, i)
	}
	g.goFile.writeln(")")
}
