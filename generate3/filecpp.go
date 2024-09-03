package generate

type fileCpp struct {
	*file
}

func newFileCpp() *fileCpp {
	f := &fileCpp{
		file: newFile("./api.cpp"),
	}

	f.writelnf(`
		#include "skia.h"

		extern "C"
  	{
		#include "api.h"

	`)

	return f
}

func (f fileCpp) finish() {
	f.writelnf(`
		}
	`)

	f.close()
	f.clangFormat()
}
