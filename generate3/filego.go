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

	import (
		"unsafe"
	)

	type class unsafe.Pointer
	`)

	return f
}

func (f fileGo) finish() {
	f.close()
}

func (f fileGo) writeComment(comment string) {
	if len(comment) > 0 {
		f.writeln(comment)
	}
}
