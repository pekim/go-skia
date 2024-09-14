package generate

var headerFiles = []string{
	"include/core/SkArc.h",
	"include/core/SkBitmap.h",
	"include/core/SkCanvas.h",
	"include/core/SkClipOp.h",
	"include/core/SkColor.h",
	"include/core/SkColorSpace.h",
	"include/core/SkFont.h",
	"include/core/SkFontMgr.h",
	"include/core/SkFontStyle.h",
	"include/core/SkImage.h",
	"include/core/SkImageInfo.h",
	"include/core/SkM44.h",
	"include/core/SkMatrix.h",
	"include/core/SkPaint.h",
	"include/core/SkPath.h",
	"include/core/SkPixmap.h",
	"include/private/base/SkPoint_impl.h",
	"include/core/SkRect.h",
	"include/core/SkRRect.h",
	"include/core/SkRegion.h",
	"include/core/SkSize.h",
	"include/core/SkSamplingOptions.h",
	"include/core/SkSurface.h",
	"include/core/SkSurfaceProps.h",
	"include/core/SkTypeface.h",
	"include/core/SkTypes.h",
	"include/gpu/gl/GrGLInterface.h",
	"include/gpu/gl/GrGLTypes.h",
	"include/gpu/ganesh/gl/GrGLBackendSurface.h",
	"include/gpu/ganesh/gl/GrGLDirectContext.h",
	"include/gpu/ganesh/SkSurfaceGanesh.h",
	"include/gpu/GrTypes.h",
	"include/gpu/GpuTypes.h",
	"include/gpu/GrBackendSurface.h",
	"include/gpu/GrContextOptions.h",
	"include/gpu/GrDirectContext.h",
	"include/gpu/GrRecordingContext.h",
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
