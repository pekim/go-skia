package generate

import "strings"

type generator struct {
	// classes    []class
	// enums      []enum
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

	classes := loadClasses()
	classes.generate(g)
}

func stripSkPrefix(name string) string {
	return strings.TrimPrefix(name, "Sk")
}
