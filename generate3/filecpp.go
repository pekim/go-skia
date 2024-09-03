package generate

type fileCpp struct {
	*file
}

func newFileCpp() *fileCpp {
	f := &fileCpp{
		file: newFile("./api.cpp"),
	}

	for _, headerFile := range headerFiles {
		f.writelnf("#include <%s>", headerFile)
	}
	f.writeln("")

	f.writelnf(`
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
