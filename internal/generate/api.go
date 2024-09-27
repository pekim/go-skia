package generate

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"muzzammil.xyz/jsonc"
)

//go:embed api.jsonc
var apiJson []byte

type api struct {
	Records       []record   `json:"records"`
	Enums         []enum     `json:"enums"`
	Functions     []callable `json:"functions"`
	Typedefs      []typedef  `json:"typedefs"`
	Variables     []variable
	variablesLock *sync.Mutex
}

func loadApi() api {
	var api api
	err := json.Unmarshal(jsonc.ToJSON(apiJson), &api)
	fatalOnError(err)

	fmt.Print("load api ")
	start := time.Now()
	api.variablesLock = new(sync.Mutex)
	tus := make([]translationUnit, len(headerFiles))
	var group errgroup.Group
	group.SetLimit(runtime.NumCPU())
	for i, headerFile := range headerFiles {
		group.Go(func() error {
			tus[i] = newTranslationUnit("_skia/skia/" + headerFile)
			fmt.Print(".")
			return nil
		})
	}
	err = group.Wait()
	fatalOnError(err)
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())

	fmt.Print("enrich 1")
	start = time.Now()
	for _, tu := range tus {
		tu.enrichApi(&api)
	}
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())

	// 2nd enrichment phase
	fmt.Print("enrich 2")
	start = time.Now()
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
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())

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

func (a api) findFunction(name string) (*callable, bool) {
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
	fmt.Print("generate")
	start := time.Now()

	for _, record := range a.Records {
		record.generateCStruct(g)
		for _, record := range record.Records {
			record.generateCStruct(g)
		}
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
	slices.SortFunc(a.Variables, func(a, b variable) int {
		return strings.Compare(a.cppName, b.cppName)
	})
	for _, variable := range a.Variables {
		variable.generate(g)
	}

	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}
