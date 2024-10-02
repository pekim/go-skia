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

	// #cgo CXXFLAGS: -I ${SRCDIR}/header
	// #cgo LDFLAGS: -L ${SRCDIR}/lib
	// #cgo LDFLAGS: -l skia
	// #cgo LDFLAGS: -l skshaper
	// #cgo LDFLAGS: -l svg
	// #cgo pkg-config: freetype2
  // #cgo pkg-config: gl
	//
	// #include "api.h"
	// #include <stdlib.h>
	import "C"

	import (
	  "sync"
		"unsafe"
	)
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
