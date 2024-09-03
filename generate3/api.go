package generate

import (
	_ "embed"
	"encoding/json"
)

//go:embed api.json
var apiJson []byte

type api struct {
	Classes classes `json:"classes"`
}

type classes []class

type class struct {
	Name  string `json:"name"`
	Enums []enum `json:"enums"`
}

type enum struct {
	Name string `json:"name"`
	// constants []enumConstant
}

// type enumConstant struct {
// 	name string
// }

func loadApi() api {
	var api api
	err := json.Unmarshal(apiJson, &api)
	if err != nil {
		panic(err)
	}

	return api
}
