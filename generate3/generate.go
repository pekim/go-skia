package generate

var headerFiles = []string{
	"include/core/SkCanvas.h",
	"include/core/SkPaint.h",
	"include/core/SkPath.h",
}

type generator struct {
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

	api := loadApi()
	api.generate(g)
}
