package generate

type fileCpp struct {
	*genFile
}

func newFileCpp() *fileCpp {
	f := &fileCpp{
		genFile: newGenFile("./api.cpp"),
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
}
