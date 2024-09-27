package generate

import (
	"fmt"
	"sync"
)

var headerFiles = []string{
	"include/core/SkAlphaType.h",
	"include/core/SkArc.h",
	"include/core/SkBitmap.h",
	"include/core/SkCanvas.h",
	"include/core/SkClipOp.h",
	"include/core/SkColor.h",
	"include/core/SkColorType.h",
	"include/core/SkColorSpace.h",
	"include/core/SkData.h",
	"include/core/SkFont.h",
	"include/core/SkFontArguments.h",
	"include/core/SkFontMetrics.h",
	"include/core/SkFontMgr.h",
	"include/core/SkFontStyle.h",
	"include/core/SkFontTypes.h",
	"include/core/SkImage.h",
	"include/core/SkImageInfo.h",
	"include/core/SkM44.h",
	"include/core/SkMatrix.h",
	"include/core/SkPaint.h",
	"include/core/SkPath.h",
	"include/core/SkPathTypes.h",
	"include/core/SkPixmap.h",
	"include/private/base/SkPoint_impl.h",
	"include/core/SkRect.h",
	"include/core/SkRRect.h",
	"include/core/SkRegion.h",
	"include/core/SkSize.h",
	"include/core/SkSamplingOptions.h",
	"include/core/SkString.h",
	"include/core/SkStream.h",
	"include/core/SkSurface.h",
	"include/core/SkSurfaceProps.h",
	"include/core/SkTextBlob.h",
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
	"include/pathops/SkPathOps.h",
	"modules/svg/include/SkSVGDOM.h",
	"modules/svg/include/SkSVGRenderContext.h",
	"modules/svg/include/SkSVGSVG.h",
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

	api := &api{
		variablesLock: new(sync.Mutex),
	}
	fmt.Println()
	api.parseTranslationUnits()
	api.enrich1()
	api.enrich2()
	api.generate(g)
	fmt.Println()
	api.printStats()
	fmt.Println()
}
