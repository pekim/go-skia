package generate

import (
	_ "embed"
	"encoding/json"
)

type classes []class

type class struct {
	Name string `json:"name"`
}

//go:embed class.json
var classJson []byte

func loadClasses() classes {
	var classes classes
	err := json.Unmarshal(classJson, &classes)
	if err != nil {
		panic(err)
	}
	return classes
}
