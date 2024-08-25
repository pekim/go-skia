package generate

type fileGo struct {
	*genFile
}

func newFileGo() *fileGo {
	f := &fileGo{
		genFile: newGenFile("./skia.go"),
	}

	f.genFile.goFmt = true
	f.writelnf(`
	package skia

	// #cgo CXXFLAGS: -I ${SRCDIR}/skia/skia
	// #include "api.h"
	import "C"

	import (
		"unsafe"
		)
		`)

	return f
}

func (f fileGo) finish() {
	f.close()
}
