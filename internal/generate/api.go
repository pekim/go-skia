package generate

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

//go:embed api.json
var apiJson []byte

type api struct {
	Records   []record   `json:"records"`
	Enums     []enum     `json:"enums"`
	Functions []function `json:"functions"`
	Typedefs  []typedef  `json:"typedefs"`
	Variables []variable
}

func loadApi() api {
	var api api
	err := json.Unmarshal(apiJson, &api)
	fatalOnError(err)

	fmt.Print("load api ")
	var group errgroup.Group
	group.SetLimit(runtime.NumCPU() / 2)
	for _, headerFile := range headerFiles {
		group.Go(func() error {
			tu := newTranslationUnit("_skia/skia/" + headerFile)
			tu.enrichApi(&api)
			fmt.Print(".")
			return nil
		})
	}
	err = group.Wait()
	fatalOnError(err)
	fmt.Println()

	// 2nd enrichment phase
	for i := range api.Enums {
		enum := &api.Enums[i]
		enum.enrich2(api)
	}
	for i := range api.Typedefs {
		record := &api.Typedefs[i]
		record.enrich2(api)
	}
	for i := range api.Functions {
		function := &api.Functions[i]
		function.enrich2(api)
	}
	for i := range api.Records {
		record := &api.Records[i]
		record.enrich2(api)
	}
	for i := range api.Variables {
		variable := &api.Variables[i]
		variable.enrich2(api)
	}

	return api
}

func (a api) findRecord(name string) (*record, bool) {
	for i, record := range a.Records {
		if record.CppName == name {
			return &a.Records[i], true
		}
	}
	return nil, false
}

func (a api) findEnum(name string) (*enum, bool) {
	for i, enum := range a.Enums {
		if enum.CppName == name {
			return &a.Enums[i], true
		}
	}
	return nil, false
}

func (a api) findFunction(name string) (*function, bool) {
	for i, function := range a.Functions {
		if function.CppName == name {
			return &a.Functions[i], true
		}
	}
	return nil, false
}

func (a api) findTypedef(name string) (*typedef, bool) {
	for i, typedef := range a.Typedefs {
		if typedef.CppName == name {
			return &a.Typedefs[i], true
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

	for _, function := range a.Functions {
		function.generate(g)
	}

	for _, typedef := range a.Typedefs {
		typedef.generate(g)
	}

	g.headerFile.writeln()
	for _, variable := range a.Variables {
		variable.generate(g)
	}
}
