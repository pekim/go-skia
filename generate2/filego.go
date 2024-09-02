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

	// #cgo CXXFLAGS: -I ${SRCDIR}/skia/skia
	// #cgo LDFLAGS: -L ${SRCDIR}/skia/build
	// #cgo LDFLAGS: -l skia
	// #cgo LDFLAGS: -l skshaper
	// #cgo LDFLAGS: -l svg
	// #cgo pkg-config: freetype2
	// #include "api.h"
	import "C"

	// import (
	// 	"unsafe"
	// )
	`)

	return f
}

func (f fileGo) finish() {
	f.close()
}

func (f fileGo) docComment(doc []string) {
	if len(doc) == 0 {
		return
	}

	if len(doc) == 1 {
		f.write("// ")
		f.writeln(doc[0])
		return
	}

	f.writeln("/*")
	for i, para := range doc {
		if len(doc) == 0 {
			continue
		}
		f.writeln(para)
		if i < len(doc)-1 {
			f.writeln("")
		}
	}
	f.writeln("*/")
}
