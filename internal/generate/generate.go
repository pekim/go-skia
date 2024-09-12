package generate

var headerFiles = []string{
	"include/core/SkBitmap.h",
	"include/core/SkCanvas.h",
	"include/core/SkColor.h",
	"include/core/SkColorSpace.h",
	"include/core/SkFont.h",
	"include/core/SkFontMgr.h",
	"include/core/SkFontStyle.h",
	"include/core/SkPaint.h",
	"include/core/SkPath.h",
	"include/core/SkRect.h",
	"include/core/SkSize.h",
	"include/core/SkSurfaceProps.h",
	"include/core/SkTypeface.h",
	"include/core/SkTypes.h",
	"include/gpu/gl/GrGLInterface.h",
	"include/gpu/gl/GrGLTypes.h",
	"include/gpu/ganesh/gl/GrGLBackendSurface.h",
	"include/gpu/ganesh/gl/GrGLDirectContext.h",
	"include/gpu/GpuTypes.h",
	"include/gpu/GrBackendSurface.h",
	"include/gpu/GrContextOptions.h",
	"include/gpu/GrDirectContext.h",
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
