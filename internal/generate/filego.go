package generate

type fileGo struct {
	*file
}

func newFileGo() *fileGo {
	f := &fileGo{
		file: newFile("./api.go"),
	}

	f.file.goFmt = true
	f.writelnf(`
	package skia

	// #cgo CXXFLAGS: -I ${SRCDIR}/skia
	// #cgo LDFLAGS: -l skia
	// #cgo LDFLAGS: -l skshaper
	// #cgo LDFLAGS: -l svg
	// #cgo LDFLAGS: -l skunicode_core
	// #cgo LDFLAGS: -l skunicode_icu
	// #cgo pkg-config: freetype2
  // #cgo pkg-config: gl
	//
	// #include "api.h"
	// #include <stdlib.h>
	import "C"

	import (
		"fmt"
	  "sync"
		"unsafe"
	)

  func assertSizesMatch(name string, cSize int, cppSize int) {
    if cSize != cppSize {
      panic(fmt.Sprintf("struct size differs : sk_%%s %%d != %%s %%d",
        name, cSize,
        name, cppSize,
    ))
    }
  }

	`)

	return f
}

func (f fileGo) finish() {
	f.writeln(`
		var fontMgr = sync.OnceValue(func() FontMgr {
			return FontMgr{
				sk: C.sk_fontmgr_ref_default(),
			}
		})

		// FontMgrDefault returns a FontMgr, that can be used to get Typefaces.
		func FontMgrDefault() FontMgr {
			return fontMgr()
		}
	`)

	f.close()
}

func (f fileGo) generateStructSizeAssertions(g generator, records []record) {
	f.writeln("func assertAllStructsSizesMatch() {")

	for _, record := range records {
		record.generateStructSizeAssertion(g)
		for _, record := range record.Records {
			record.generateStructSizeAssertion(g)
		}
	}

	f.writeln("}")
}
