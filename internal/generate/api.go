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
	tus           []translationUnit
	variablesLock *sync.Mutex
}

func (api *api) parseTranslationUnits() {
	err := json.Unmarshal(jsonc.ToJSON(apiJson), &api)
	fatalOnError(err)

	fmt.Print("load api ")
	start := time.Now()
	api.tus = make([]translationUnit, len(headerFiles))
	var group errgroup.Group
	group.SetLimit(runtime.NumCPU())
	for i, headerFile := range headerFiles {
		group.Go(func() error {
			api.tus[i] = newTranslationUnit("_skia/skia/" + headerFile)
			fmt.Print(".")
			return nil
		})
	}
	err = group.Wait()
	fatalOnError(err)
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

func (api *api) enrich1() {
	fmt.Print("enrich 1")
	start := time.Now()
	for _, tu := range api.tus {
		tu.enrichApi(api)
	}
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

func (api api) enrich2() {
	fmt.Print("enrich 2")
	start := time.Now()
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
}

func (api api) findRecord(name string) (*record, bool) {
	for i, record := range api.Records {
		if record.CppName == name {
			return &api.Records[i], true
		}
	}
	return nil, false
}

func (api api) findEnum(name string) (*enum, bool) {
	for i, enum := range api.Enums {
		if enum.CppName == name {
			return &api.Enums[i], true
		}
	}
	return nil, false
}

func (api api) findFunction(name string) (*callable, bool) {
	for i, function := range api.Functions {
		if function.CppName == name {
			return &api.Functions[i], true
		}
	}
	return nil, false
}

func (api api) findTypedef(name string) (*typedef, bool) {
	for i, typedef := range api.Typedefs {
		if typedef.CppName == name {
			return &api.Typedefs[i], true
		}
	}
	return nil, false
}

func (api api) generate(g generator) {
	fmt.Print("generate")
	start := time.Now()

	for _, record := range api.Records {
		record.generateCStruct(g)
		for _, record := range record.Records {
			record.generateCStruct(g)
		}
	}

	for _, record := range api.Records {
		record.generate(g)
	}

	for _, enum := range api.Enums {
		enum.generate(g)
	}

	for _, function := range api.Functions {
		function.generate(g)
	}

	for _, typedef := range api.Typedefs {
		typedef.generate(g)
	}

	g.headerFile.writeln()
	slices.SortFunc(api.Variables, func(a, b variable) int {
		return strings.Compare(a.cppName, b.cppName)
	})
	for _, variable := range api.Variables {
		variable.generate(g)
	}

	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}
