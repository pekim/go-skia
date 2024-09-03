package generate

import (
	_ "embed"
	"encoding/json"
)

//go:embed api.json
var apiJson []byte

type api struct {
	Classes []class `json:"classes"`
}

func loadApi() api {
	var api api
	err := json.Unmarshal(apiJson, &api)
	if err != nil {
		panic(err)
	}

	for _, headerFile := range headerFiles {
		tu := newTranslationUnit("skia/skia/" + headerFile)
		tu.enrichApi(&api)
	}

	return api
}

func (a api) findClass(name string) (*class, bool) {
	for i, class := range a.Classes {
		if class.Name == name {
			return &a.Classes[i], true
		}
	}
	return nil, false
}

func (a api) generate(g generator) {
	for _, class := range a.Classes {
		class.generate(g)
	}
}
