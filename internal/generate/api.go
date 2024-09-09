package generate

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed api.json
var apiJson []byte

type api struct {
	Records []record `json:"records"`
	Enums   []enum   `json:"enums"`
}

func loadApi() api {
	var api api
	err := json.Unmarshal(apiJson, &api)
	fatalOnError(err)

	fmt.Print("load api ")
	for _, headerFile := range headerFiles {
		fmt.Print(".")
		tu := newTranslationUnit("_skia/skia/" + headerFile)
		tu.enrichApi(&api)
	}
	fmt.Println()

	// for i := range api.Functions {
	// 	function := &api.Functions[i]
	// 	function.enrich2(api)
	// }
	for i := range api.Enums {
		enum := &api.Enums[i]
		enum.enrich2(api)
	}
	for i := range api.Records {
		record := &api.Records[i]
		record.enrich2(api)
	}

	return api
}

func (a api) findRecord(name string) (*record, bool) {
	for i, record := range a.Records {
		if record.CName == name {
			return &a.Records[i], true
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

	for _, record := range a.Records {
		record.generateCStruct(g)
	}

	for _, record := range a.Records {
		record.generate(g)
	}

	for _, enum := range a.Enums {
		enum.generate(g)
	}
}
