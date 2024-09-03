package generate

import (
	"fmt"
)

type generator struct {
	api        api
	goFile     *fileGo
	headerFile *fileHeader
	cppFile    *fileCpp
}

func Generate() {
	g := generator{}
	g.goFile = newFileGo()
	defer g.goFile.finish()

	g.headerFile = newFileHeader()
	defer g.headerFile.finish()

	g.cppFile = newFileCpp()
	defer g.cppFile.finish()

	g.api = loadApi()
	fmt.Println(len(g.api.Classes))
	// api.generate(g)
}
