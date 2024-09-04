package generate

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed api.json
var apiJson []byte

type api struct {
	Classes []class `json:"classes"`
	Enums   []enum  `json:"enums"`
}

func loadApi() api {
	var api api
	err := json.Unmarshal(apiJson, &api)
	fatalOnError(err)

	fmt.Print("load api ")
	for _, headerFile := range headerFiles {
		fmt.Print(".")
		tu := newTranslationUnit("skia/skia/" + headerFile)
		tu.enrichApi(&api)
	}
	fmt.Println()

	return api
}

func (a api) findClass(name string) (*class, bool) {
	for i, class := range a.Classes {
		if class.CName == name {
			return &a.Classes[i], true
		}
	}
	return nil, false
}

func (a api) findEnum(name string) (*enum, bool) {
	for i, enum := range a.Enums {
		if enum.CName == name {
			return &a.Enums[i], true
		}
	}
	return nil, false
}

func (a api) generate(g generator) {
	fmt.Println("generate")

	for _, class := range a.Classes {
		class.generate(g)
	}

	for _, enum := range a.Enums {
		enum.generate(g)
	}
}
