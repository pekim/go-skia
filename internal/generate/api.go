package generate

import (
	"fmt"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type api struct {
	records       []record
	enums         []enum
	functions     []callable
	typedefs      []typedef
	variables     []variable
	tus           []translationUnit
	variablesLock *sync.Mutex
}

func (api *api) parseTranslationUnits() {
	fmt.Print("parse header files ")
	start := time.Now()
	api.tus = make([]translationUnit, len(headerFiles))
	var group errgroup.Group
	group.SetLimit(runtime.NumCPU())
	for i, headerFile := range headerFiles {
		group.Go(func() error {
			api.tus[i] = newTranslationUnit("_skia/" + headerFile)
			fmt.Print(".")
			return nil
		})
	}
	err := group.Wait()
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
	for i := range api.enums {
		enum := &api.enums[i]
		enum.enrich2(api)
	}
	for i := range api.typedefs {
		record := &api.typedefs[i]
		record.enrich2(api)
	}
	for i := range api.functions {
		function := &api.functions[i]
		function.enrich2(nil, api)
	}
	for i := range api.records {
		record := &api.records[i]
		record.enrich2(api)
	}
	for i := range api.variables {
		variable := &api.variables[i]
		variable.enrich2(api)
	}
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

func (api api) findRecord(name string) (*record, bool) {
	for i, record := range api.records {
		if record.cppName == name {
			return &api.records[i], true
		}
	}
	return nil, false
}

func (api api) findEnum(name string) (*enum, bool) {
	for i, enum := range api.enums {
		if enum.cppName == name {
			return &api.enums[i], true
		}
	}
	return nil, false
}

func (api api) findFunction(name string) (*callable, bool) {
	for i, function := range api.functions {
		if function.cppName == name {
			return &api.functions[i], true
		}
	}
	return nil, false
}

func (api api) findTypedef(name string) (*typedef, bool) {
	for i, typedef := range api.typedefs {
		if typedef.cppName == name {
			return &api.typedefs[i], true
		}
	}
	return nil, false
}

func (api api) generate(g generator) {
	fmt.Print("generate")
	start := time.Now()

	g.goFile.generateStructSizeAssertions(g, api.records)

	for _, record := range api.records {
		record.generateCStruct(g)
		record.generateCPPStructSize(g)
		for _, record := range record.records {
			record.generateCStruct(g)
			record.generateCPPStructSize(g)
		}
	}
	g.cppFile.writeln()

	for _, record := range api.records {
		record.generate(g)
	}

	for _, enum := range api.enums {
		enum.generate(g)
	}

	for _, function := range api.functions {
		function.generate(g)
	}

	for _, typedef := range api.typedefs {
		typedef.generate(g)
	}

	g.headerFile.writeln()
	slices.SortFunc(api.variables, func(a, b variable) int {
		return strings.Compare(a.cppName, b.cppName)
	})
	for _, variable := range api.variables {
		variable.generate(g)
	}

	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

func (api api) printStats() {
	classCount := 0
	structCount := 0
	classEnumCount := 0
	classMethodCount := 0
	classRecordCount := 0
	structEnumCount := 0
	structMethodCount := 0
	structRecordCount := 0
	for _, record := range api.records {
		if record.isClass {
			classCount++
			classEnumCount += len(record.enums)
			classMethodCount += len(record.methods)
			classRecordCount += len(record.records)
		} else {
			structCount++
			structEnumCount += len(record.enums)
			structMethodCount += len(record.methods)
			structRecordCount += len(record.records)
		}
	}

	fmt.Printf("%4d classes (with %d enums, %d methods, %d records)\n",
		classCount, classEnumCount, classMethodCount, classRecordCount)
	fmt.Printf("%4d enums\n", len(api.enums))
	fmt.Printf("%4d functions\n", len(api.functions))
	fmt.Printf("%4d structs (with %d enums, %d methods, %d records)\n",
		structCount, structEnumCount, structMethodCount, structRecordCount)
	fmt.Printf("%4d typedefs\n", len(api.typedefs))
}

func (api api) goNameForCppName(cppName string) (string, bool) {
	if enum, ok := api.findEnum(cppName); ok {
		return enum.goName, true
	}
	if record, ok := api.findRecord(cppName); ok {
		return record.goName, true
	}
	if typedef, ok := api.findTypedef(cppName); ok {
		return typedef.goName, true
	}

	return "", false
}
