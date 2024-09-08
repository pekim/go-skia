package generate

import (
	"strings"
)

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

	// #cgo CXXFLAGS: -I ${SRCDIR}/_skia/skia
	// #cgo LDFLAGS: -L ${SRCDIR}/_skia/build
	// #cgo LDFLAGS: -l skia
	// #cgo LDFLAGS: -l skshaper
	// #cgo LDFLAGS: -l svg
	// #cgo pkg-config: freetype2
  // #cgo pkg-config: gl
	//
	// #include "api.h"
	import "C"

	`)

	return f
}

func (f fileGo) finish() {
	f.writeln(`
		func FontMgrRefDefault() FontMgr {
			return FontMgr {
				sk: C.sk_fontmgr_ref_default(),
			}
		}
	`)

	f.close()
}

func (f fileGo) writeDocComment(comment string) {
	if len(comment) == 0 {
		return
	}

	// Tidy up comment, and avoid godoc seeing unintended indentation.
	//		- replace "/** ... */" with "/* ... */"
	// 		- trim leading (and trailing) white space from lines
	comment = strings.TrimSpace(comment)
	if strings.HasPrefix(comment, "/**") {
		comment = strings.TrimPrefix(comment, "/**")
		comment = strings.TrimSuffix(comment, "*/")
		lines := strings.Split(comment, "\n")
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
			lines[i] = strings.TrimPrefix(lines[i], "*")
		}
		comment = "/*\n" + strings.Join(lines, "\n") + "\n*/"
	}

	f.writeln(comment)

}
