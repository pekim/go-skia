// This is a generated file. DO NOT EDIT.

package skia

// #cgo CXXFLAGS: -I ${SRCDIR}/skia/skia
// #cgo LDFLAGS: -L ${SRCDIR}/skia/build
// #cgo LDFLAGS: -l skia
// #cgo LDFLAGS: -l skshaper
// #cgo LDFLAGS: -l svg
// #cgo pkg-config: freetype2
// #include "api.h"
import "C"

import (
	"unsafe"
)

type Matrix struct {
	skia unsafe.Pointer
}

func NewMatrix() Matrix {
	return Matrix{
		skia: C.skia_new_SkMatrix(),
	}
}

type MatrixScaleToFit int64

const (
	MatrixScaleToFit_Fill   MatrixScaleToFit = 0
	MatrixScaleToFit_Start  MatrixScaleToFit = 1
	MatrixScaleToFit_Center MatrixScaleToFit = 2
	MatrixScaleToFit_End    MatrixScaleToFit = 3
)

type MatrixTypeMask int64

const (
	MatrixTypeMask_Identity_Mask    MatrixTypeMask = 0
	MatrixTypeMask_Translate_Mask   MatrixTypeMask = 1
	MatrixTypeMask_Scale_Mask       MatrixTypeMask = 2
	MatrixTypeMask_Affine_Mask      MatrixTypeMask = 4
	MatrixTypeMask_Perspective_Mask MatrixTypeMask = 8
)

type RefCntBase struct {
	skia unsafe.Pointer
}

func NewRefCntBase() RefCntBase {
	return RefCntBase{
		skia: C.skia_new_SkRefCntBase(),
	}
}
func (o *RefCntBase) Delete() {
	C.skia_delete_SkRefCntBase(o.skia)
}

type ColorInfo struct {
	skia unsafe.Pointer
}

func NewColorInfo() ColorInfo {
	return ColorInfo{
		skia: C.skia_new_SkColorInfo(),
	}
}
func (o *ColorInfo) Delete() {
	C.skia_delete_SkColorInfo(o.skia)
}

type Pixmap struct {
	skia unsafe.Pointer
}

func NewPixmap() Pixmap {
	return Pixmap{
		skia: C.skia_new_SkPixmap(),
	}
}

type Data struct {
	skia unsafe.Pointer
}

type YUVAInfo struct {
	skia unsafe.Pointer
}

func NewYUVAInfo() YUVAInfo {
	return YUVAInfo{
		skia: C.skia_new_SkYUVAInfo(),
	}
}

type YUVAInfoYUVAChannels int64

const (
	YUVAInfoYUVAChannels_Y    YUVAInfoYUVAChannels = 0
	YUVAInfoYUVAChannels_U    YUVAInfoYUVAChannels = 1
	YUVAInfoYUVAChannels_V    YUVAInfoYUVAChannels = 2
	YUVAInfoYUVAChannels_A    YUVAInfoYUVAChannels = 3
	YUVAInfoYUVAChannels_Last YUVAInfoYUVAChannels = 3
)

type YUVAInfoPlaneConfig int64

const (
	YUVAInfoPlaneConfig_Unknown YUVAInfoPlaneConfig = 0
	YUVAInfoPlaneConfig_Y_U_V   YUVAInfoPlaneConfig = 1
	YUVAInfoPlaneConfig_Y_V_U   YUVAInfoPlaneConfig = 2
	YUVAInfoPlaneConfig_Y_UV    YUVAInfoPlaneConfig = 3
	YUVAInfoPlaneConfig_Y_VU    YUVAInfoPlaneConfig = 4
	YUVAInfoPlaneConfig_YUV     YUVAInfoPlaneConfig = 5
	YUVAInfoPlaneConfig_UYV     YUVAInfoPlaneConfig = 6
	YUVAInfoPlaneConfig_Y_U_V_A YUVAInfoPlaneConfig = 7
	YUVAInfoPlaneConfig_Y_V_U_A YUVAInfoPlaneConfig = 8
	YUVAInfoPlaneConfig_Y_UV_A  YUVAInfoPlaneConfig = 9
	YUVAInfoPlaneConfig_Y_VU_A  YUVAInfoPlaneConfig = 10
	YUVAInfoPlaneConfig_YUVA    YUVAInfoPlaneConfig = 11
	YUVAInfoPlaneConfig_UYVA    YUVAInfoPlaneConfig = 12
	YUVAInfoPlaneConfig_Last    YUVAInfoPlaneConfig = 12
)

type YUVAInfoSubsampling int64

const (
	YUVAInfoSubsampling_Unknown YUVAInfoSubsampling = 0
	YUVAInfoSubsampling_444     YUVAInfoSubsampling = 1
	YUVAInfoSubsampling_422     YUVAInfoSubsampling = 2
	YUVAInfoSubsampling_420     YUVAInfoSubsampling = 3
	YUVAInfoSubsampling_440     YUVAInfoSubsampling = 4
	YUVAInfoSubsampling_411     YUVAInfoSubsampling = 5
	YUVAInfoSubsampling_410     YUVAInfoSubsampling = 6
	YUVAInfoSubsampling_Last    YUVAInfoSubsampling = 6
)

type YUVAInfoSiting int64

const (
	YUVAInfoSiting_Centered YUVAInfoSiting = 0
)

type YUVAPixmapInfo struct {
	skia unsafe.Pointer
}

func NewYUVAPixmapInfo() YUVAPixmapInfo {
	return YUVAPixmapInfo{
		skia: C.skia_new_SkYUVAPixmapInfo(),
	}
}

type YUVAPixmapInfoDataType int64

const (
	YUVAPixmapInfoDataType_Unorm8         YUVAPixmapInfoDataType = 0
	YUVAPixmapInfoDataType_Unorm16        YUVAPixmapInfoDataType = 1
	YUVAPixmapInfoDataType_Float16        YUVAPixmapInfoDataType = 2
	YUVAPixmapInfoDataType_Unorm10_Unorm2 YUVAPixmapInfoDataType = 3
	YUVAPixmapInfoDataType_Last           YUVAPixmapInfoDataType = 3
)

type YUVAPixmaps struct {
	skia unsafe.Pointer
}

func NewYUVAPixmaps() YUVAPixmaps {
	return YUVAPixmaps{
		skia: C.skia_new_SkYUVAPixmaps(),
	}
}
func (o *YUVAPixmaps) Delete() {
	C.skia_delete_SkYUVAPixmaps(o.skia)
}

type Once struct {
	skia unsafe.Pointer
}

func NewOnce() Once {
	return Once{
		skia: C.skia_new_SkOnce(),
	}
}

type OnceState int64

const (
	OnceState_NotStarted OnceState = 0
	OnceState_Claimed    OnceState = 1
	OnceState_Done       OnceState = 2
)

type ColorSpace struct {
	skia unsafe.Pointer
}

type Noncopyable struct {
	skia unsafe.Pointer
}

func NewNoncopyable() Noncopyable {
	return Noncopyable{
		skia: C.skia_new_SkNoncopyable(),
	}
}

type Codec struct {
	skia unsafe.Pointer
}

func (o *Codec) Delete() {
	C.skia_delete_SkCodec(o.skia)
}

type CodecResult int64

const (
	CodecResult_Success           CodecResult = 0
	CodecResult_IncompleteInput   CodecResult = 1
	CodecResult_ErrorInInput      CodecResult = 2
	CodecResult_InvalidConversion CodecResult = 3
	CodecResult_InvalidScale      CodecResult = 4
	CodecResult_InvalidParameters CodecResult = 5
	CodecResult_InvalidInput      CodecResult = 6
	CodecResult_CouldNotRewind    CodecResult = 7
	CodecResult_InternalError     CodecResult = 8
	CodecResult_Unimplemented     CodecResult = 9
)

type CodecSelectionPolicy int64

const (
	CodecSelectionPolicy_PreferStillImage CodecSelectionPolicy = 0
	CodecSelectionPolicy_PreferAnimation  CodecSelectionPolicy = 1
)

type CodecZeroInitialized int64

const (
	CodecZeroInitialized_Yes CodecZeroInitialized = 0
	CodecZeroInitialized_No  CodecZeroInitialized = 1
)

type CodecSkScanlineOrder int64

const (
	CodecSkScanlineOrder_TopDown  CodecSkScanlineOrder = 0
	CodecSkScanlineOrder_BottomUp CodecSkScanlineOrder = 1
)

type CodecXformTime int64

const (
	CodecXformTime_No        CodecXformTime = 0
	CodecXformTime_Palette   CodecXformTime = 1
	CodecXformTime_DecodeRow CodecXformTime = 2
)

type Flattenable struct {
	skia unsafe.Pointer
}

type FlattenableType int64

const (
	FlattenableType_SkColorFilter FlattenableType = 0
	FlattenableType_SkBlender     FlattenableType = 1
	FlattenableType_SkDrawable    FlattenableType = 2
	FlattenableType_SkDrawLooper  FlattenableType = 3
	FlattenableType_SkImageFilter FlattenableType = 4
	FlattenableType_SkMaskFilter  FlattenableType = 5
	FlattenableType_SkPathEffect  FlattenableType = 6
	FlattenableType_SkShader      FlattenableType = 7
)

type ImageFilter struct {
	skia unsafe.Pointer
}

type ImageFilterMapDirection int64

const (
	ImageFilterMapDirection_Forward ImageFilterMapDirection = 0
	ImageFilterMapDirection_Reverse ImageFilterMapDirection = 1
)

type M44 struct {
	skia unsafe.Pointer
}

func NewM44() M44 {
	return M44{
		skia: C.skia_new_SkM44(),
	}
}

type M44Uninitialized_Constructor int64

const (
	M44Uninitialized_Constructor_ M44Uninitialized_Constructor = 0
)

type M44NaN_Constructor int64

const (
	M44NaN_Constructor_ M44NaN_Constructor = 0
)

type Paint struct {
	skia unsafe.Pointer
}

func NewPaint() Paint {
	return Paint{
		skia: C.skia_new_SkPaint(),
	}
}
func (o *Paint) Delete() {
	C.skia_delete_SkPaint(o.skia)
}

type PaintStyle int64

const (
	PaintStyle_Fill          PaintStyle = 0
	PaintStyle_Stroke        PaintStyle = 1
	PaintStyle_StrokeAndFill PaintStyle = 2
)

type PaintCap int64

const (
	PaintCap_Butt    PaintCap = 0
	PaintCap_Round   PaintCap = 1
	PaintCap_Square  PaintCap = 2
	PaintCap_Last    PaintCap = 2
	PaintCap_Default PaintCap = 0
)

type PaintJoin int64

const (
	PaintJoin_Miter   PaintJoin = 0
	PaintJoin_Round   PaintJoin = 1
	PaintJoin_Bevel   PaintJoin = 2
	PaintJoin_Last    PaintJoin = 2
	PaintJoin_Default PaintJoin = 0
)

type RasterHandleAllocator struct {
	skia unsafe.Pointer
}

func (o *RasterHandleAllocator) Delete() {
	C.skia_delete_SkRasterHandleAllocator(o.skia)
}

type String struct {
	skia unsafe.Pointer
}

func NewString() String {
	return String{
		skia: C.skia_new_SkString(),
	}
}
func (o *String) Delete() {
	C.skia_delete_SkString(o.skia)
}

type SurfaceProps struct {
	skia unsafe.Pointer
}

func NewSurfaceProps() SurfaceProps {
	return SurfaceProps{
		skia: C.skia_new_SkSurfaceProps(),
	}
}

type SurfacePropsFlags int64

const (
	SurfacePropsFlags_Default                   SurfacePropsFlags = 0
	SurfacePropsFlags_UseDeviceIndependentFonts SurfacePropsFlags = 1
	SurfacePropsFlags_DynamicMSAA               SurfacePropsFlags = 2
	SurfacePropsFlags_AlwaysDither              SurfacePropsFlags = 4
)

type Deque struct {
	skia unsafe.Pointer
}

func (o *Deque) Delete() {
	C.skia_delete_SkDeque(o.skia)
}

type ContainerAllocator struct {
	skia unsafe.Pointer
}

type Canvas struct {
	skia unsafe.Pointer
}

func NewCanvas() Canvas {
	return Canvas{
		skia: C.skia_new_SkCanvas(),
	}
}
func (o *Canvas) Delete() {
	C.skia_delete_SkCanvas(o.skia)
}

type CanvasSaveLayerFlagsSet int64

const (
	CanvasSaveLayerFlagsSet_PreserveLCDText_SaveLayer  CanvasSaveLayerFlagsSet = 2
	CanvasSaveLayerFlagsSet_InitWithPrevious_SaveLayer CanvasSaveLayerFlagsSet = 4
	CanvasSaveLayerFlagsSet_F16ColorType               CanvasSaveLayerFlagsSet = 16
)

type CanvasPointMode int64

const (
	CanvasPointMode_Points  CanvasPointMode = 0
	CanvasPointMode_Lines   CanvasPointMode = 1
	CanvasPointMode_Polygon CanvasPointMode = 2
)

type CanvasSrcRectConstraint int64

const (
	CanvasSrcRectConstraint_Strict CanvasSrcRectConstraint = 0
	CanvasSrcRectConstraint_Fast   CanvasSrcRectConstraint = 1
)

type CanvasQuadAAFlags int64

const (
	CanvasQuadAAFlags_Left_QuadAA   CanvasQuadAAFlags = 1
	CanvasQuadAAFlags_Top_QuadAA    CanvasQuadAAFlags = 2
	CanvasQuadAAFlags_Right_QuadAA  CanvasQuadAAFlags = 4
	CanvasQuadAAFlags_Bottom_QuadAA CanvasQuadAAFlags = 8
	CanvasQuadAAFlags_None          CanvasQuadAAFlags = 0
	CanvasQuadAAFlags_All           CanvasQuadAAFlags = 15
)

type CanvasSaveLayerStrategy int64

const (
	CanvasSaveLayerStrategy_FullLayer CanvasSaveLayerStrategy = 0
	CanvasSaveLayerStrategy_NoLayer   CanvasSaveLayerStrategy = 1
)

type CanvasClipEdgeStyle int64

const (
	CanvasClipEdgeStyle_Hard CanvasClipEdgeStyle = 0
	CanvasClipEdgeStyle_Soft CanvasClipEdgeStyle = 1
)

type CanvasPredrawFlags int64

const (
	CanvasPredrawFlags_None                    CanvasPredrawFlags = 0
	CanvasPredrawFlags_OpaqueShaderOverride    CanvasPredrawFlags = 1
	CanvasPredrawFlags_NonOpaqueShaderOverride CanvasPredrawFlags = 2
	CanvasPredrawFlags_CheckForOverwrite       CanvasPredrawFlags = 4
	CanvasPredrawFlags_SkipMaskFilterAutoLayer CanvasPredrawFlags = 8
)

type CanvasDeviceCompatibleWithFilter int64

const (
	CanvasDeviceCompatibleWithFilter_Unknown CanvasDeviceCompatibleWithFilter = 0
	CanvasDeviceCompatibleWithFilter_Yes     CanvasDeviceCompatibleWithFilter = -1
)

type AutoCanvasRestore struct {
	skia unsafe.Pointer
}

func (o *AutoCanvasRestore) Delete() {
	C.skia_delete_SkAutoCanvasRestore(o.skia)
}

type Document struct {
	skia unsafe.Pointer
}

type DocumentState int64

const (
	DocumentState_BetweenPages DocumentState = 0
	DocumentState_InPage       DocumentState = 1
	DocumentState_Closed       DocumentState = 2
)

type FontStyleSet struct {
	skia unsafe.Pointer
}

type FontMgr struct {
	skia unsafe.Pointer
}

type MaskFilter struct {
	skia unsafe.Pointer
}

type Stream struct {
	skia unsafe.Pointer
}

func (o *Stream) Delete() {
	C.skia_delete_SkStream(o.skia)
}

type StreamRewindable struct {
	skia unsafe.Pointer
}

type StreamSeekable struct {
	skia unsafe.Pointer
}

type StreamAsset struct {
	skia unsafe.Pointer
}

type StreamMemory struct {
	skia unsafe.Pointer
}

type WStream struct {
	skia unsafe.Pointer
}

func (o *WStream) Delete() {
	C.skia_delete_SkWStream(o.skia)
}

type NullWStream struct {
	skia unsafe.Pointer
}

type FILEStream struct {
	skia unsafe.Pointer
}

func (o *FILEStream) Delete() {
	C.skia_delete_SkFILEStream(o.skia)
}

type MemoryStream struct {
	skia unsafe.Pointer
}

func NewMemoryStream() MemoryStream {
	return MemoryStream{
		skia: C.skia_new_SkMemoryStream(),
	}
}

type FILEWStream struct {
	skia unsafe.Pointer
}

func (o *FILEWStream) Delete() {
	C.skia_delete_SkFILEWStream(o.skia)
}

type DynamicMemoryWStream struct {
	skia unsafe.Pointer
}

func NewDynamicMemoryWStream() DynamicMemoryWStream {
	return DynamicMemoryWStream{
		skia: C.skia_new_SkDynamicMemoryWStream(),
	}
}
func (o *DynamicMemoryWStream) Delete() {
	C.skia_delete_SkDynamicMemoryWStream(o.skia)
}

type Image struct {
	skia unsafe.Pointer
}

type ImageCachingHint int64

const (
	ImageCachingHint_Allow    ImageCachingHint = 0
	ImageCachingHint_Disallow ImageCachingHint = 1
)

type ImageRescaleGamma int64

const (
	ImageRescaleGamma_Src    ImageRescaleGamma = 0
	ImageRescaleGamma_Linear ImageRescaleGamma = -1
)

type ImageRescaleMode int64

const (
	ImageRescaleMode_Nearest        ImageRescaleMode = 0
	ImageRescaleMode_Linear         ImageRescaleMode = 1
	ImageRescaleMode_RepeatedLinear ImageRescaleMode = 2
	ImageRescaleMode_RepeatedCubic  ImageRescaleMode = 3
)

type ImageLegacyBitmapMode int64

const (
	ImageLegacyBitmapMode_RO ImageLegacyBitmapMode = 0
)

type Surface struct {
	skia unsafe.Pointer
}

type SurfaceContentChangeMode int64

const (
	SurfaceContentChangeMode_Discard SurfaceContentChangeMode = 0
	SurfaceContentChangeMode_Retain  SurfaceContentChangeMode = 1
)

type SurfaceBackendHandleAccess int64

const (
	SurfaceBackendHandleAccess_FlushRead    SurfaceBackendHandleAccess = 0
	SurfaceBackendHandleAccess_FlushWrite   SurfaceBackendHandleAccess = 1
	SurfaceBackendHandleAccess_DiscardWrite SurfaceBackendHandleAccess = 2
)

type FontStyle struct {
	skia unsafe.Pointer
}

func NewFontStyle() FontStyle {
	return FontStyle{
		skia: C.skia_new_SkFontStyle(),
	}
}

type FontStyleWeight int64

const (
	FontStyleWeight_Invisible  FontStyleWeight = 0
	FontStyleWeight_Thin       FontStyleWeight = 100
	FontStyleWeight_ExtraLight FontStyleWeight = 200
	FontStyleWeight_Light      FontStyleWeight = 300
	FontStyleWeight_Normal     FontStyleWeight = 400
	FontStyleWeight_Medium     FontStyleWeight = 500
	FontStyleWeight_SemiBold   FontStyleWeight = 600
	FontStyleWeight_Bold       FontStyleWeight = 700
	FontStyleWeight_ExtraBold  FontStyleWeight = 800
	FontStyleWeight_Black      FontStyleWeight = 900
	FontStyleWeight_ExtraBlack FontStyleWeight = 1000
)

type FontStyleWidth int64

const (
	FontStyleWidth_UltraCondensed FontStyleWidth = 1
	FontStyleWidth_ExtraCondensed FontStyleWidth = 2
	FontStyleWidth_Condensed      FontStyleWidth = 3
	FontStyleWidth_SemiCondensed  FontStyleWidth = 4
	FontStyleWidth_Normal         FontStyleWidth = 5
	FontStyleWidth_SemiExpanded   FontStyleWidth = 6
	FontStyleWidth_Expanded       FontStyleWidth = 7
	FontStyleWidth_ExtraExpanded  FontStyleWidth = 8
	FontStyleWidth_UltraExpanded  FontStyleWidth = 9
)

type FontStyleSlant int64

const (
	FontStyleSlant_Upright FontStyleSlant = 0
	FontStyleSlant_Italic  FontStyleSlant = 1
	FontStyleSlant_Oblique FontStyleSlant = 2
)

type WeakRefCnt struct {
	skia unsafe.Pointer
}

func NewWeakRefCnt() WeakRefCnt {
	return WeakRefCnt{
		skia: C.skia_new_SkWeakRefCnt(),
	}
}
func (o *WeakRefCnt) Delete() {
	C.skia_delete_SkWeakRefCnt(o.skia)
}

type Typeface struct {
	skia unsafe.Pointer
}

type TypefaceSerializeBehavior int64

const (
	TypefaceSerializeBehavior_DoIncludeData      TypefaceSerializeBehavior = 0
	TypefaceSerializeBehavior_DontIncludeData    TypefaceSerializeBehavior = 1
	TypefaceSerializeBehavior_IncludeDataIfLocal TypefaceSerializeBehavior = 2
)

type Font struct {
	skia unsafe.Pointer
}

func NewFont() Font {
	return Font{
		skia: C.skia_new_SkFont(),
	}
}

type FontEdging int64

const (
	FontEdging_Alias             FontEdging = 0
	FontEdging_AntiAlias         FontEdging = 1
	FontEdging_SubpixelAntiAlias FontEdging = 2
)

type FontPrivFlags int64

const (
	FontPrivFlags_ForceAutoHinting_Priv FontPrivFlags = 1
	FontPrivFlags_EmbeddedBitmaps_Priv  FontPrivFlags = 2
	FontPrivFlags_Subpixel_Priv         FontPrivFlags = 4
	FontPrivFlags_LinearMetrics_Priv    FontPrivFlags = 8
	FontPrivFlags_Embolden_Priv         FontPrivFlags = 16
	FontPrivFlags_BaselineSnap_Priv     FontPrivFlags = 32
)

type TextBlob struct {
	skia unsafe.Pointer
}

type TextBlobBuilder struct {
	skia unsafe.Pointer
}

func NewTextBlobBuilder() TextBlobBuilder {
	return TextBlobBuilder{
		skia: C.skia_new_SkTextBlobBuilder(),
	}
}
func (o *TextBlobBuilder) Delete() {
	C.skia_delete_SkTextBlobBuilder(o.skia)
}

type Path1DPathEffect struct {
	skia unsafe.Pointer
}

type Path1DPathEffectStyle int64

const (
	Path1DPathEffectStyle_Translate Path1DPathEffectStyle = 0
	Path1DPathEffectStyle_Rotate    Path1DPathEffectStyle = 1
	Path1DPathEffectStyle_Morph     Path1DPathEffectStyle = 2
	Path1DPathEffectStyle_LastEnum  Path1DPathEffectStyle = 2
)

type Line2DPathEffect struct {
	skia unsafe.Pointer
}

type Path2DPathEffect struct {
	skia unsafe.Pointer
}

type ColorFilter struct {
	skia unsafe.Pointer
}

type ColorFilters struct {
	skia unsafe.Pointer
}

type ColorMatrixFilter struct {
	skia unsafe.Pointer
}

type CornerPathEffect struct {
	skia unsafe.Pointer
}

type DashPathEffect struct {
	skia unsafe.Pointer
}

type Path struct {
	skia unsafe.Pointer
}

func NewPath() Path {
	return Path{
		skia: C.skia_new_SkPath(),
	}
}
func (o *Path) Delete() {
	C.skia_delete_SkPath(o.skia)
}

type PathArcSize int64

const (
	PathArcSize_Small PathArcSize = 0
	PathArcSize_Large PathArcSize = 1
)

type PathAddPathMode int64

const (
	PathAddPathMode_Append PathAddPathMode = 0
	PathAddPathMode_Extend PathAddPathMode = 1
)

type PathSegmentMask int64

const (
	PathSegmentMask_Line  PathSegmentMask = 1
	PathSegmentMask_Quad  PathSegmentMask = 2
	PathSegmentMask_Conic PathSegmentMask = 4
	PathSegmentMask_Cubic PathSegmentMask = 8
)

type PathVerb int64

const (
	PathVerb_Move  PathVerb = 0
	PathVerb_Line  PathVerb = 1
	PathVerb_Quad  PathVerb = 2
	PathVerb_Conic PathVerb = 3
	PathVerb_Cubic PathVerb = 4
	PathVerb_Close PathVerb = 5
	PathVerb_Done  PathVerb = 6
)

type PathEffect struct {
	skia unsafe.Pointer
}

type PathEffectDashType int64

const (
	PathEffectDashType_None PathEffectDashType = 0
	PathEffectDashType_Dash PathEffectDashType = 1
)

type DiscretePathEffect struct {
	skia unsafe.Pointer
}

type Shader struct {
	skia unsafe.Pointer
}

type GradientShader struct {
	skia unsafe.Pointer
}

type GradientShaderFlags int64

const (
	GradientShaderFlags_InterpolateColorsInPremul GradientShaderFlags = 1
)

type Picture struct {
	skia unsafe.Pointer
}

func (o *Picture) Delete() {
	C.skia_delete_SkPicture(o.skia)
}

type ImageFilters struct {
	skia unsafe.Pointer
}

type ImageFiltersDither int64

const (
	ImageFiltersDither_No  ImageFiltersDither = 0
	ImageFiltersDither_Yes ImageFiltersDither = -1
)

type ShaderMaskFilter struct {
	skia unsafe.Pointer
}

type TableMaskFilter struct {
	skia unsafe.Pointer
}

type TrimPathEffect struct {
	skia unsafe.Pointer
}

type TrimPathEffectMode int64

const (
	TrimPathEffectMode_Normal   TrimPathEffectMode = 0
	TrimPathEffectMode_Inverted TrimPathEffectMode = 1
)

type DataTable struct {
	skia unsafe.Pointer
}

type Encoder struct {
	skia unsafe.Pointer
}

func (o *Encoder) Delete() {
	C.skia_delete_SkEncoder(o.skia)
}

type GrGLExtensions struct {
	skia unsafe.Pointer
}

func NewGrGLExtensions() GrGLExtensions {
	return GrGLExtensions{
		skia: C.skia_new_GrGLExtensions(),
	}
}

type GrColorFormatDesc struct {
	skia unsafe.Pointer
}

type GrBackendFormat struct {
	skia unsafe.Pointer
}

func NewGrBackendFormat() GrBackendFormat {
	return GrBackendFormat{
		skia: C.skia_new_GrBackendFormat(),
	}
}
func (o *GrBackendFormat) Delete() {
	C.skia_delete_GrBackendFormat(o.skia)
}

type GrBackendTexture struct {
	skia unsafe.Pointer
}

func NewGrBackendTexture() GrBackendTexture {
	return GrBackendTexture{
		skia: C.skia_new_GrBackendTexture(),
	}
}
func (o *GrBackendTexture) Delete() {
	C.skia_delete_GrBackendTexture(o.skia)
}

type GrBackendRenderTarget struct {
	skia unsafe.Pointer
}

func NewGrBackendRenderTarget() GrBackendRenderTarget {
	return GrBackendRenderTarget{
		skia: C.skia_new_GrBackendRenderTarget(),
	}
}
func (o *GrBackendRenderTarget) Delete() {
	C.skia_delete_GrBackendRenderTarget(o.skia)
}

type GrDriverBugWorkarounds struct {
	skia unsafe.Pointer
}

func NewGrDriverBugWorkarounds() GrDriverBugWorkarounds {
	return GrDriverBugWorkarounds{
		skia: C.skia_new_GrDriverBugWorkarounds(),
	}
}
func (o *GrDriverBugWorkarounds) Delete() {
	C.skia_delete_GrDriverBugWorkarounds(o.skia)
}

type GrContext_Base struct {
	skia unsafe.Pointer
}

func (o *GrContext_Base) Delete() {
	C.skia_delete_GrContext_Base(o.skia)
}

type Semaphore struct {
	skia unsafe.Pointer
}

func (o *Semaphore) Delete() {
	C.skia_delete_SkSemaphore(o.skia)
}

type Mutex struct {
	skia unsafe.Pointer
}

func NewMutex() Mutex {
	return Mutex{
		skia: C.skia_new_SkMutex(),
	}
}
func (o *Mutex) Delete() {
	C.skia_delete_SkMutex(o.skia)
}

type AutoMutexExclusive struct {
	skia unsafe.Pointer
}

func (o *AutoMutexExclusive) Delete() {
	C.skia_delete_SkAutoMutexExclusive(o.skia)
}

type GrImageContext struct {
	skia unsafe.Pointer
}

func (o *GrImageContext) Delete() {
	C.skia_delete_GrImageContext(o.skia)
}

type GrRecordingContext struct {
	skia unsafe.Pointer
}

func (o *GrRecordingContext) Delete() {
	C.skia_delete_GrRecordingContext(o.skia)
}

type GrDirectContext struct {
	skia unsafe.Pointer
}

func (o *GrDirectContext) Delete() {
	C.skia_delete_GrDirectContext(o.skia)
}

type TDStorage struct {
	skia unsafe.Pointer
}

func (o *TDStorage) Delete() {
	C.skia_delete_SkTDStorage(o.skia)
}

type OpBuilder struct {
	skia unsafe.Pointer
}

type ParsePath struct {
	skia unsafe.Pointer
}

type ParsePathPathEncoding int64

const (
	ParsePathPathEncoding_Absolute ParsePathPathEncoding = 0
	ParsePathPathEncoding_Relative ParsePathPathEncoding = 1
)

type PDFBitmap struct {
	skia unsafe.Pointer
}

type PDFUnion struct {
	skia unsafe.Pointer
}

func (o *PDFUnion) Delete() {
	C.skia_delete_SkPDFUnion(o.skia)
}

type PDFUnionType int64

const (
	PDFUnionType_Destroyed       PDFUnionType = 0
	PDFUnionType_Int             PDFUnionType = 1
	PDFUnionType_ColorComponent  PDFUnionType = 2
	PDFUnionType_ColorComponentF PDFUnionType = 3
	PDFUnionType_Bool            PDFUnionType = 4
	PDFUnionType_Scalar          PDFUnionType = 5
	PDFUnionType_Name            PDFUnionType = 6
	PDFUnionType_ByteString      PDFUnionType = 7
	PDFUnionType_TextString      PDFUnionType = 8
	PDFUnionType_NameSkS         PDFUnionType = 9
	PDFUnionType_ByteStringSkS   PDFUnionType = 10
	PDFUnionType_TextStringSkS   PDFUnionType = 11
	PDFUnionType_Object          PDFUnionType = 12
	PDFUnionType_Ref             PDFUnionType = 13
)

type PDFObject struct {
	skia unsafe.Pointer
}

func (o *PDFObject) Delete() {
	C.skia_delete_SkPDFObject(o.skia)
}

type PDFArray struct {
	skia unsafe.Pointer
}

func NewPDFArray() PDFArray {
	return PDFArray{
		skia: C.skia_new_SkPDFArray(),
	}
}
func (o *PDFArray) Delete() {
	C.skia_delete_SkPDFArray(o.skia)
}

type PDFDict struct {
	skia unsafe.Pointer
}

func (o *PDFDict) Delete() {
	C.skia_delete_SkPDFDict(o.skia)
}

type ArenaAlloc struct {
	skia unsafe.Pointer
}

func (o *ArenaAlloc) Delete() {
	C.skia_delete_SkArenaAlloc(o.skia)
}

type ArenaAllocWithReset struct {
	skia unsafe.Pointer
}

type PDFTagTree struct {
	skia unsafe.Pointer
}

func NewPDFTagTree() PDFTagTree {
	return PDFTagTree{
		skia: C.skia_new_SkPDFTagTree(),
	}
}
func (o *PDFTagTree) Delete() {
	C.skia_delete_SkPDFTagTree(o.skia)
}

type PDFOffsetMap struct {
	skia unsafe.Pointer
}

type PDFDocument struct {
	skia unsafe.Pointer
}

func (o *PDFDocument) Delete() {
	C.skia_delete_SkPDFDocument(o.skia)
}

type Shaper struct {
	skia unsafe.Pointer
}

func (o *Shaper) Delete() {
	C.skia_delete_SkShaper(o.skia)
}

type TextBlobBuilderRunHandler struct {
	skia unsafe.Pointer
}

type SVGDOM struct {
	skia unsafe.Pointer
}

type SVGLength struct {
	skia unsafe.Pointer
}

func NewSVGLength() SVGLength {
	return SVGLength{
		skia: C.skia_new_SkSVGLength(),
	}
}

type SVGLengthUnit int64

const (
	SVGLengthUnit_Unknown    SVGLengthUnit = 0
	SVGLengthUnit_Number     SVGLengthUnit = 1
	SVGLengthUnit_Percentage SVGLengthUnit = 2
	SVGLengthUnit_EMS        SVGLengthUnit = 3
	SVGLengthUnit_EXS        SVGLengthUnit = 4
	SVGLengthUnit_PX         SVGLengthUnit = 5
	SVGLengthUnit_CM         SVGLengthUnit = 6
	SVGLengthUnit_MM         SVGLengthUnit = 7
	SVGLengthUnit_IN         SVGLengthUnit = 8
	SVGLengthUnit_PT         SVGLengthUnit = 9
	SVGLengthUnit_PC         SVGLengthUnit = 10
)

type SVGIRI struct {
	skia unsafe.Pointer
}

func NewSVGIRI() SVGIRI {
	return SVGIRI{
		skia: C.skia_new_SkSVGIRI(),
	}
}

type SVGIRIType int64

const (
	SVGIRIType_Local    SVGIRIType = 0
	SVGIRIType_Nonlocal SVGIRIType = 1
	SVGIRIType_DataURI  SVGIRIType = 2
)

type SVGColor struct {
	skia unsafe.Pointer
}

func NewSVGColor() SVGColor {
	return SVGColor{
		skia: C.skia_new_SkSVGColor(),
	}
}

type SVGColorType int64

const (
	SVGColorType_CurrentColor SVGColorType = 0
	SVGColorType_Color        SVGColorType = 1
	SVGColorType_ICCColor     SVGColorType = 2
)

type SVGPaint struct {
	skia unsafe.Pointer
}

func NewSVGPaint() SVGPaint {
	return SVGPaint{
		skia: C.skia_new_SkSVGPaint(),
	}
}

type SVGPaintType int64

const (
	SVGPaintType_None  SVGPaintType = 0
	SVGPaintType_Color SVGPaintType = 1
	SVGPaintType_IRI   SVGPaintType = 2
)

type SVGFuncIRI struct {
	skia unsafe.Pointer
}

func NewSVGFuncIRI() SVGFuncIRI {
	return SVGFuncIRI{
		skia: C.skia_new_SkSVGFuncIRI(),
	}
}

type SVGFuncIRIType int64

const (
	SVGFuncIRIType_None SVGFuncIRIType = 0
	SVGFuncIRIType_IRI  SVGFuncIRIType = 1
)

type SVGLineJoin struct {
	skia unsafe.Pointer
}

func NewSVGLineJoin() SVGLineJoin {
	return SVGLineJoin{
		skia: C.skia_new_SkSVGLineJoin(),
	}
}

type SVGLineJoinType int64

const (
	SVGLineJoinType_Miter   SVGLineJoinType = 0
	SVGLineJoinType_Round   SVGLineJoinType = 1
	SVGLineJoinType_Bevel   SVGLineJoinType = 2
	SVGLineJoinType_Inherit SVGLineJoinType = 3
)

type SVGSpreadMethod struct {
	skia unsafe.Pointer
}

func NewSVGSpreadMethod() SVGSpreadMethod {
	return SVGSpreadMethod{
		skia: C.skia_new_SkSVGSpreadMethod(),
	}
}

type SVGSpreadMethodType int64

const (
	SVGSpreadMethodType_Pad     SVGSpreadMethodType = 0
	SVGSpreadMethodType_Repeat  SVGSpreadMethodType = 1
	SVGSpreadMethodType_Reflect SVGSpreadMethodType = 2
)

type SVGFillRule struct {
	skia unsafe.Pointer
}

func NewSVGFillRule() SVGFillRule {
	return SVGFillRule{
		skia: C.skia_new_SkSVGFillRule(),
	}
}

type SVGFillRuleType int64

const (
	SVGFillRuleType_NonZero SVGFillRuleType = 0
	SVGFillRuleType_EvenOdd SVGFillRuleType = 1
	SVGFillRuleType_Inherit SVGFillRuleType = 2
)

type SVGVisibility struct {
	skia unsafe.Pointer
}

func NewSVGVisibility() SVGVisibility {
	return SVGVisibility{
		skia: C.skia_new_SkSVGVisibility(),
	}
}

type SVGVisibilityType int64

const (
	SVGVisibilityType_Visible  SVGVisibilityType = 0
	SVGVisibilityType_Hidden   SVGVisibilityType = 1
	SVGVisibilityType_Collapse SVGVisibilityType = 2
	SVGVisibilityType_Inherit  SVGVisibilityType = 3
)

type SVGDashArray struct {
	skia unsafe.Pointer
}

func NewSVGDashArray() SVGDashArray {
	return SVGDashArray{
		skia: C.skia_new_SkSVGDashArray(),
	}
}

type SVGDashArrayType int64

const (
	SVGDashArrayType_None      SVGDashArrayType = 0
	SVGDashArrayType_DashArray SVGDashArrayType = 1
	SVGDashArrayType_Inherit   SVGDashArrayType = 2
)

type SVGStopColor struct {
	skia unsafe.Pointer
}

func NewSVGStopColor() SVGStopColor {
	return SVGStopColor{
		skia: C.skia_new_SkSVGStopColor(),
	}
}

type SVGStopColorType int64

const (
	SVGStopColorType_Color        SVGStopColorType = 0
	SVGStopColorType_CurrentColor SVGStopColorType = 1
	SVGStopColorType_ICCColor     SVGStopColorType = 2
	SVGStopColorType_Inherit      SVGStopColorType = 3
)

type SVGObjectBoundingBoxUnits struct {
	skia unsafe.Pointer
}

func NewSVGObjectBoundingBoxUnits() SVGObjectBoundingBoxUnits {
	return SVGObjectBoundingBoxUnits{
		skia: C.skia_new_SkSVGObjectBoundingBoxUnits(),
	}
}

type SVGObjectBoundingBoxUnitsType int64

const (
	SVGObjectBoundingBoxUnitsType_UserSpaceOnUse    SVGObjectBoundingBoxUnitsType = 0
	SVGObjectBoundingBoxUnitsType_ObjectBoundingBox SVGObjectBoundingBoxUnitsType = 1
)

type SVGFontFamily struct {
	skia unsafe.Pointer
}

func NewSVGFontFamily() SVGFontFamily {
	return SVGFontFamily{
		skia: C.skia_new_SkSVGFontFamily(),
	}
}

type SVGFontFamilyType int64

const (
	SVGFontFamilyType_Family  SVGFontFamilyType = 0
	SVGFontFamilyType_Inherit SVGFontFamilyType = 1
)

type SVGFontStyle struct {
	skia unsafe.Pointer
}

func NewSVGFontStyle() SVGFontStyle {
	return SVGFontStyle{
		skia: C.skia_new_SkSVGFontStyle(),
	}
}

type SVGFontStyleType int64

const (
	SVGFontStyleType_Normal  SVGFontStyleType = 0
	SVGFontStyleType_Italic  SVGFontStyleType = 1
	SVGFontStyleType_Oblique SVGFontStyleType = 2
	SVGFontStyleType_Inherit SVGFontStyleType = 3
)

type SVGFontSize struct {
	skia unsafe.Pointer
}

func NewSVGFontSize() SVGFontSize {
	return SVGFontSize{
		skia: C.skia_new_SkSVGFontSize(),
	}
}

type SVGFontSizeType int64

const (
	SVGFontSizeType_Length  SVGFontSizeType = 0
	SVGFontSizeType_Inherit SVGFontSizeType = 1
)

type SVGFontWeight struct {
	skia unsafe.Pointer
}

func NewSVGFontWeight() SVGFontWeight {
	return SVGFontWeight{
		skia: C.skia_new_SkSVGFontWeight(),
	}
}

type SVGFontWeightType int64

const (
	SVGFontWeightType_100     SVGFontWeightType = 0
	SVGFontWeightType_200     SVGFontWeightType = 1
	SVGFontWeightType_300     SVGFontWeightType = 2
	SVGFontWeightType_400     SVGFontWeightType = 3
	SVGFontWeightType_500     SVGFontWeightType = 4
	SVGFontWeightType_600     SVGFontWeightType = 5
	SVGFontWeightType_700     SVGFontWeightType = 6
	SVGFontWeightType_800     SVGFontWeightType = 7
	SVGFontWeightType_900     SVGFontWeightType = 8
	SVGFontWeightType_Normal  SVGFontWeightType = 9
	SVGFontWeightType_Bold    SVGFontWeightType = 10
	SVGFontWeightType_Bolder  SVGFontWeightType = 11
	SVGFontWeightType_Lighter SVGFontWeightType = 12
	SVGFontWeightType_Inherit SVGFontWeightType = 13
)

type SVGTextAnchor struct {
	skia unsafe.Pointer
}

func NewSVGTextAnchor() SVGTextAnchor {
	return SVGTextAnchor{
		skia: C.skia_new_SkSVGTextAnchor(),
	}
}

type SVGTextAnchorType int64

const (
	SVGTextAnchorType_Start   SVGTextAnchorType = 0
	SVGTextAnchorType_Middle  SVGTextAnchorType = 1
	SVGTextAnchorType_End     SVGTextAnchorType = 2
	SVGTextAnchorType_Inherit SVGTextAnchorType = 3
)

type SVGFeInputType struct {
	skia unsafe.Pointer
}

func NewSVGFeInputType() SVGFeInputType {
	return SVGFeInputType{
		skia: C.skia_new_SkSVGFeInputType(),
	}
}

type SVGFeInputTypeType int64

const (
	SVGFeInputTypeType_SourceGraphic            SVGFeInputTypeType = 0
	SVGFeInputTypeType_SourceAlpha              SVGFeInputTypeType = 1
	SVGFeInputTypeType_BackgroundImage          SVGFeInputTypeType = 2
	SVGFeInputTypeType_BackgroundAlpha          SVGFeInputTypeType = 3
	SVGFeInputTypeType_FillPaint                SVGFeInputTypeType = 4
	SVGFeInputTypeType_StrokePaint              SVGFeInputTypeType = 5
	SVGFeInputTypeType_FilterPrimitiveReference SVGFeInputTypeType = 6
	SVGFeInputTypeType_Unspecified              SVGFeInputTypeType = 7
)

type SVGFeTurbulenceBaseFrequency struct {
	skia unsafe.Pointer
}

func NewSVGFeTurbulenceBaseFrequency() SVGFeTurbulenceBaseFrequency {
	return SVGFeTurbulenceBaseFrequency{
		skia: C.skia_new_SkSVGFeTurbulenceBaseFrequency(),
	}
}

type SVGAttributeParser struct {
	skia unsafe.Pointer
}

type SVGNode struct {
	skia unsafe.Pointer
}

func (o *SVGNode) Delete() {
	C.skia_delete_SkSVGNode(o.skia)
}

type SVGTransformableNode struct {
	skia unsafe.Pointer
}

type SVGContainer struct {
	skia unsafe.Pointer
}

type SVGSVG struct {
	skia unsafe.Pointer
}

type SVGSVGType int64

const (
	SVGSVGType_Root  SVGSVGType = 0
	SVGSVGType_Inner SVGSVGType = 1
)

type SVGLengthContext struct {
	skia unsafe.Pointer
}

type SVGLengthContextLengthType int64

const (
	SVGLengthContextLengthType_Horizontal SVGLengthContextLengthType = 0
	SVGLengthContextLengthType_Vertical   SVGLengthContextLengthType = 1
	SVGLengthContextLengthType_Other      SVGLengthContextLengthType = 2
)

type SVGRenderContext struct {
	skia unsafe.Pointer
}

func (o *SVGRenderContext) Delete() {
	C.skia_delete_SkSVGRenderContext(o.skia)
}

type SVGRenderContextApplyFlags int64

const (
	SVGRenderContextApplyFlags_Leaf SVGRenderContextApplyFlags = 1
)

type ApplyPerspectiveClip int64

const (
	ApplyPerspectiveClip_No  ApplyPerspectiveClip = 0
	ApplyPerspectiveClip_Yes ApplyPerspectiveClip = 1
)

type EncodedOrigin int64

const (
	EncodedOrigin_TopLeft     EncodedOrigin = 1
	EncodedOrigin_TopRight    EncodedOrigin = 2
	EncodedOrigin_BottomRight EncodedOrigin = 3
	EncodedOrigin_BottomLeft  EncodedOrigin = 4
	EncodedOrigin_LeftTop     EncodedOrigin = 5
	EncodedOrigin_RightTop    EncodedOrigin = 6
	EncodedOrigin_RightBottom EncodedOrigin = 7
	EncodedOrigin_LeftBottom  EncodedOrigin = 8
	EncodedOrigin_Default     EncodedOrigin = 1
	EncodedOrigin_Last        EncodedOrigin = 8
)

type AlphaType int64

const (
	AlphaType_Unknown  AlphaType = 0
	AlphaType_Opaque   AlphaType = 1
	AlphaType_Premul   AlphaType = 2
	AlphaType_Unpremul AlphaType = 3
	AlphaType_LastEnum AlphaType = 3
)

type ColorType int64

const (
	ColorType_Unknown            ColorType = 0
	ColorType_Alpha_8            ColorType = 1
	ColorType_RGB_565            ColorType = 2
	ColorType_ARGB_4444          ColorType = 3
	ColorType_RGBA_8888          ColorType = 4
	ColorType_RGB_888x           ColorType = 5
	ColorType_BGRA_8888          ColorType = 6
	ColorType_RGBA_1010102       ColorType = 7
	ColorType_BGRA_1010102       ColorType = 8
	ColorType_RGB_101010x        ColorType = 9
	ColorType_BGR_101010x        ColorType = 10
	ColorType_BGR_101010x_XR     ColorType = 11
	ColorType_BGRA_10101010_XR   ColorType = 12
	ColorType_RGBA_10x6          ColorType = 13
	ColorType_Gray_8             ColorType = 14
	ColorType_RGBA_F16Norm       ColorType = 15
	ColorType_RGBA_F16           ColorType = 16
	ColorType_RGBA_F32           ColorType = 17
	ColorType_R8G8_unorm         ColorType = 18
	ColorType_A16_float          ColorType = 19
	ColorType_R16G16_float       ColorType = 20
	ColorType_A16_unorm          ColorType = 21
	ColorType_R16G16_unorm       ColorType = 22
	ColorType_R16G16B16A16_unorm ColorType = 23
	ColorType_SRGBA_8888         ColorType = 24
	ColorType_R8_unorm           ColorType = 25
	ColorType_LastEnum           ColorType = 25
	ColorType_N32                ColorType = 4
)

type YUVColorSpace int64

const (
	YUVColorSpace_JPEG_Full            YUVColorSpace = 0
	YUVColorSpace_Rec601_Limited       YUVColorSpace = 1
	YUVColorSpace_Rec709_Full          YUVColorSpace = 2
	YUVColorSpace_Rec709_Limited       YUVColorSpace = 3
	YUVColorSpace_BT2020_8bit_Full     YUVColorSpace = 4
	YUVColorSpace_BT2020_8bit_Limited  YUVColorSpace = 5
	YUVColorSpace_BT2020_10bit_Full    YUVColorSpace = 6
	YUVColorSpace_BT2020_10bit_Limited YUVColorSpace = 7
	YUVColorSpace_BT2020_12bit_Full    YUVColorSpace = 8
	YUVColorSpace_BT2020_12bit_Limited YUVColorSpace = 9
	YUVColorSpace_FCC_Full             YUVColorSpace = 10
	YUVColorSpace_FCC_Limited          YUVColorSpace = 11
	YUVColorSpace_SMPTE240_Full        YUVColorSpace = 12
	YUVColorSpace_SMPTE240_Limited     YUVColorSpace = 13
	YUVColorSpace_YDZDX_Full           YUVColorSpace = 14
	YUVColorSpace_YDZDX_Limited        YUVColorSpace = 15
	YUVColorSpace_GBR_Full             YUVColorSpace = 16
	YUVColorSpace_GBR_Limited          YUVColorSpace = 17
	YUVColorSpace_YCgCo_8bit_Full      YUVColorSpace = 18
	YUVColorSpace_YCgCo_8bit_Limited   YUVColorSpace = 19
	YUVColorSpace_YCgCo_10bit_Full     YUVColorSpace = 20
	YUVColorSpace_YCgCo_10bit_Limited  YUVColorSpace = 21
	YUVColorSpace_YCgCo_12bit_Full     YUVColorSpace = 22
	YUVColorSpace_YCgCo_12bit_Limited  YUVColorSpace = 23
	YUVColorSpace_Identity             YUVColorSpace = 24
	YUVColorSpace_LastEnum             YUVColorSpace = 24
	YUVColorSpace_JPEG                 YUVColorSpace = 0
	YUVColorSpace_Rec601               YUVColorSpace = 1
	YUVColorSpace_Rec709               YUVColorSpace = 3
	YUVColorSpace_BT2020               YUVColorSpace = 5
)

type ColorChannel int64

const (
	ColorChannel_R        ColorChannel = 0
	ColorChannel_G        ColorChannel = 1
	ColorChannel_B        ColorChannel = 2
	ColorChannel_A        ColorChannel = 3
	ColorChannel_LastEnum ColorChannel = 3
)

type ColorChannelFlag int64

const (
	ColorChannelFlag_Red                           ColorChannelFlag = 1
	ColorChannelFlag_Green                         ColorChannelFlag = 2
	ColorChannelFlag_Blue                          ColorChannelFlag = 4
	ColorChannelFlag_Alpha                         ColorChannelFlag = 8
	ColorChannelFlag_Gray                          ColorChannelFlag = 16
	ColorChannelFlag_GrayAlpha_SkColorChannelFlags ColorChannelFlag = 24
	ColorChannelFlag_RG_SkColorChannelFlags        ColorChannelFlag = 3
	ColorChannelFlag_RGB_SkColorChannelFlags       ColorChannelFlag = 7
	ColorChannelFlag_RGBA_SkColorChannelFlags      ColorChannelFlag = 15
)

type FilterMode int64

const (
	FilterMode_Nearest FilterMode = 0
	FilterMode_Linear  FilterMode = 1
	FilterMode_Last    FilterMode = 1
)

type MipmapMode int64

const (
	MipmapMode_None    MipmapMode = 0
	MipmapMode_Nearest MipmapMode = 1
	MipmapMode_Linear  MipmapMode = 2
	MipmapMode_Last    MipmapMode = 2
)

type BlurStyle int64

const (
	BlurStyle_Normal   BlurStyle = 0
	BlurStyle_Solid    BlurStyle = 1
	BlurStyle_Outer    BlurStyle = 2
	BlurStyle_Inner    BlurStyle = 3
	BlurStyle_LastEnum BlurStyle = 3
)

type BlendMode int64

const (
	BlendMode_Clear             BlendMode = 0
	BlendMode_Src               BlendMode = 1
	BlendMode_Dst               BlendMode = 2
	BlendMode_SrcOver           BlendMode = 3
	BlendMode_DstOver           BlendMode = 4
	BlendMode_SrcIn             BlendMode = 5
	BlendMode_DstIn             BlendMode = 6
	BlendMode_SrcOut            BlendMode = 7
	BlendMode_DstOut            BlendMode = 8
	BlendMode_SrcATop           BlendMode = 9
	BlendMode_DstATop           BlendMode = 10
	BlendMode_Xor               BlendMode = 11
	BlendMode_Plus              BlendMode = 12
	BlendMode_Modulate          BlendMode = 13
	BlendMode_Screen            BlendMode = 14
	BlendMode_Overlay           BlendMode = 15
	BlendMode_Darken            BlendMode = 16
	BlendMode_Lighten           BlendMode = 17
	BlendMode_ColorDodge        BlendMode = 18
	BlendMode_ColorBurn         BlendMode = 19
	BlendMode_HardLight         BlendMode = 20
	BlendMode_SoftLight         BlendMode = 21
	BlendMode_Difference        BlendMode = 22
	BlendMode_Exclusion         BlendMode = 23
	BlendMode_Multiply          BlendMode = 24
	BlendMode_Hue               BlendMode = 25
	BlendMode_Saturation        BlendMode = 26
	BlendMode_Color             BlendMode = 27
	BlendMode_Luminosity        BlendMode = 28
	BlendMode_LastCoeffMode     BlendMode = 14
	BlendMode_LastSeparableMode BlendMode = 24
	BlendMode_LastMode          BlendMode = 28
)

type BlendModeCoeff int64

const (
	BlendModeCoeff_Zero       BlendModeCoeff = 0
	BlendModeCoeff_One        BlendModeCoeff = 1
	BlendModeCoeff_SC         BlendModeCoeff = 2
	BlendModeCoeff_ISC        BlendModeCoeff = 3
	BlendModeCoeff_DC         BlendModeCoeff = 4
	BlendModeCoeff_IDC        BlendModeCoeff = 5
	BlendModeCoeff_SA         BlendModeCoeff = 6
	BlendModeCoeff_ISA        BlendModeCoeff = 7
	BlendModeCoeff_DA         BlendModeCoeff = 8
	BlendModeCoeff_IDA        BlendModeCoeff = 9
	BlendModeCoeff_CoeffCount BlendModeCoeff = 10
)

type ClipOp int64

const (
	ClipOp_Difference    ClipOp = 0
	ClipOp_Intersect     ClipOp = 1
	ClipOp_Max_EnumValue ClipOp = 1
)

type TextEncoding int64

const (
	TextEncoding_UTF8    TextEncoding = 0
	TextEncoding_UTF16   TextEncoding = 1
	TextEncoding_UTF32   TextEncoding = 2
	TextEncoding_GlyphID TextEncoding = 3
)

type FontHinting int64

const (
	FontHinting_None   FontHinting = 0
	FontHinting_Slight FontHinting = 1
	FontHinting_Normal FontHinting = 2
	FontHinting_Full   FontHinting = 3
)

type PixelGeometry int64

const (
	PixelGeometry_Unknown PixelGeometry = 0
	PixelGeometry_RGB_H   PixelGeometry = 1
	PixelGeometry_BGR_H   PixelGeometry = 2
	PixelGeometry_RGB_V   PixelGeometry = 3
	PixelGeometry_BGR_V   PixelGeometry = 4
)

type PathFillType int64

const (
	PathFillType_Winding        PathFillType = 0
	PathFillType_EvenOdd        PathFillType = 1
	PathFillType_InverseWinding PathFillType = 2
	PathFillType_InverseEvenOdd PathFillType = 3
)

type PathDirection int64

const (
	PathDirection_CW  PathDirection = 0
	PathDirection_CCW PathDirection = 1
)

type TileMode int64

const (
	TileMode_Clamp        TileMode = 0
	TileMode_Repeat       TileMode = 1
	TileMode_Mirror       TileMode = 2
	TileMode_Decal        TileMode = 3
	TileMode_LastTileMode TileMode = 3
)

type GrBackendApi int64

const (
	GrBackendApi_OpenGL           GrBackendApi = 0
	GrBackendApi_Vulkan           GrBackendApi = 1
	GrBackendApi_Metal            GrBackendApi = 2
	GrBackendApi_Direct3D         GrBackendApi = 3
	GrBackendApi_Mock             GrBackendApi = 4
	GrBackendApi_Unsupported      GrBackendApi = 5
	GrBackendApi_OpenGL_GrBackend GrBackendApi = 0
)

type GrSurfaceOrigin int64

const (
	GrSurfaceOrigin_TopLeft    GrSurfaceOrigin = 0
	GrSurfaceOrigin_BottomLeft GrSurfaceOrigin = 1
)

type GrGLBackendState int64

const (
	GrGLBackendState_RenderTarget   GrGLBackendState = 1
	GrGLBackendState_TextureBinding GrGLBackendState = 2
	GrGLBackendState_View           GrGLBackendState = 4
	GrGLBackendState_Blend          GrGLBackendState = 8
	GrGLBackendState_MSAAEnable     GrGLBackendState = 16
	GrGLBackendState_Vertex         GrGLBackendState = 32
	GrGLBackendState_Stencil        GrGLBackendState = 64
	GrGLBackendState_PixelStore     GrGLBackendState = 128
	GrGLBackendState_Program        GrGLBackendState = 256
	GrGLBackendState_FixedFunction  GrGLBackendState = 512
	GrGLBackendState_Misc           GrGLBackendState = 1024
	GrGLBackendState_ALL            GrGLBackendState = 65535
)

type GrSemaphoresSubmitted int64

const (
	GrSemaphoresSubmitted_No  GrSemaphoresSubmitted = 0
	GrSemaphoresSubmitted_Yes GrSemaphoresSubmitted = -1
)

type GrPurgeResourceOptions int64

const (
	GrPurgeResourceOptions_AllResources         GrPurgeResourceOptions = 0
	GrPurgeResourceOptions_ScratchResourcesOnly GrPurgeResourceOptions = 1
)

type GrSyncCpu int64

const (
	GrSyncCpu_No  GrSyncCpu = 0
	GrSyncCpu_Yes GrSyncCpu = -1
)

type GrGLStandard int64

const (
	GrGLStandard_None  GrGLStandard = 0
	GrGLStandard_GL    GrGLStandard = 1
	GrGLStandard_GLES  GrGLStandard = 2
	GrGLStandard_WebGL GrGLStandard = 3
)

type GrGLFormat int64

const (
	GrGLFormat_Unknown              GrGLFormat = 0
	GrGLFormat_RGBA8                GrGLFormat = 1
	GrGLFormat_R8                   GrGLFormat = 2
	GrGLFormat_ALPHA8               GrGLFormat = 3
	GrGLFormat_LUMINANCE8           GrGLFormat = 4
	GrGLFormat_LUMINANCE8_ALPHA8    GrGLFormat = 5
	GrGLFormat_BGRA8                GrGLFormat = 6
	GrGLFormat_RGB565               GrGLFormat = 7
	GrGLFormat_RGBA16F              GrGLFormat = 8
	GrGLFormat_R16F                 GrGLFormat = 9
	GrGLFormat_RGB8                 GrGLFormat = 10
	GrGLFormat_RGBX8                GrGLFormat = 11
	GrGLFormat_RG8                  GrGLFormat = 12
	GrGLFormat_RGB10_A2             GrGLFormat = 13
	GrGLFormat_RGBA4                GrGLFormat = 14
	GrGLFormat_SRGB8_ALPHA8         GrGLFormat = 15
	GrGLFormat_COMPRESSED_ETC1_RGB8 GrGLFormat = 16
	GrGLFormat_COMPRESSED_RGB8_ETC2 GrGLFormat = 17
	GrGLFormat_COMPRESSED_RGB8_BC1  GrGLFormat = 18
	GrGLFormat_COMPRESSED_RGBA8_BC1 GrGLFormat = 19
	GrGLFormat_R16                  GrGLFormat = 20
	GrGLFormat_RG16                 GrGLFormat = 21
	GrGLFormat_RGBA16               GrGLFormat = 22
	GrGLFormat_RG16F                GrGLFormat = 23
	GrGLFormat_LUMINANCE16F         GrGLFormat = 24
	GrGLFormat_LastColorFormat      GrGLFormat = 24
	GrGLFormat_STENCIL_INDEX8       GrGLFormat = 25
	GrGLFormat_STENCIL_INDEX16      GrGLFormat = 26
	GrGLFormat_DEPTH24_STENCIL8     GrGLFormat = 27
	GrGLFormat_Last                 GrGLFormat = 27
)

type TextureCompressionType int64

const (
	TextureCompressionType_None            TextureCompressionType = 0
	TextureCompressionType_ETC2_RGB8_UNORM TextureCompressionType = 1
	TextureCompressionType_BC1_RGB8_UNORM  TextureCompressionType = 2
	TextureCompressionType_BC1_RGBA8_UNORM TextureCompressionType = 3
	TextureCompressionType_Last            TextureCompressionType = 3
	TextureCompressionType_ETC1_RGB8       TextureCompressionType = 1
)

type GrPrimitiveType int64

const (
	GrPrimitiveType_Triangles     GrPrimitiveType = 0
	GrPrimitiveType_TriangleStrip GrPrimitiveType = 1
	GrPrimitiveType_Points        GrPrimitiveType = 2
	GrPrimitiveType_Lines         GrPrimitiveType = 3
	GrPrimitiveType_LineStrip     GrPrimitiveType = 4
)

type GrPrimitiveRestart int64

const (
	GrPrimitiveRestart_No  GrPrimitiveRestart = 0
	GrPrimitiveRestart_Yes GrPrimitiveRestart = -1
)

type GrTexturable int64

const (
	GrTexturable_No  GrTexturable = 0
	GrTexturable_Yes GrTexturable = -1
)

type GrDDLProvider int64

const (
	GrDDLProvider_No  GrDDLProvider = 0
	GrDDLProvider_Yes GrDDLProvider = -1
)

type GrWrapOwnership int64

const (
	GrWrapOwnership_Borrow GrWrapOwnership = 0
	GrWrapOwnership_Adopt  GrWrapOwnership = 1
)

type GrWrapCacheable int64

const (
	GrWrapCacheable_No  GrWrapCacheable = 0
	GrWrapCacheable_Yes GrWrapCacheable = -1
)

type GrBudgetedType int64

const (
	GrBudgetedType_Budgeted              GrBudgetedType = 0
	GrBudgetedType_UnbudgetedUncacheable GrBudgetedType = 1
	GrBudgetedType_UnbudgetedCacheable   GrBudgetedType = 2
)

type GrScissorTest int64

const (
	GrScissorTest_Disabled GrScissorTest = 0
	GrScissorTest_Enabled  GrScissorTest = -1
)

type GrMemoryless int64

const (
	GrMemoryless_No  GrMemoryless = 0
	GrMemoryless_Yes GrMemoryless = -1
)

type GrSemaphoreWrapType int64

const (
	GrSemaphoreWrapType_WillSignal GrSemaphoreWrapType = 0
	GrSemaphoreWrapType_WillWait   GrSemaphoreWrapType = 1
)

type GrLoadOp int64

const (
	GrLoadOp_Load    GrLoadOp = 0
	GrLoadOp_Clear   GrLoadOp = 1
	GrLoadOp_Discard GrLoadOp = 2
)

type GrStoreOp int64

const (
	GrStoreOp_Store   GrStoreOp = 0
	GrStoreOp_Discard GrStoreOp = 1
)

type GrAA int64

const (
	GrAA_No  GrAA = 0
	GrAA_Yes GrAA = -1
)

type GrFillRule int64

const (
	GrFillRule_Nonzero GrFillRule = 0
	GrFillRule_EvenOdd GrFillRule = -1
)

type GrAAType int64

const (
	GrAAType_None     GrAAType = 0
	GrAAType_Coverage GrAAType = 1
	GrAAType_MSAA     GrAAType = 2
	GrAAType_Last     GrAAType = 2
)

type GrClampType int64

const (
	GrClampType_Auto   GrClampType = 0
	GrClampType_Manual GrClampType = 1
	GrClampType_None   GrClampType = 2
)

type GrQuadAAFlags int64

const (
	GrQuadAAFlags_Left   GrQuadAAFlags = 1
	GrQuadAAFlags_Top    GrQuadAAFlags = 2
	GrQuadAAFlags_Right  GrQuadAAFlags = 4
	GrQuadAAFlags_Bottom GrQuadAAFlags = 8
	GrQuadAAFlags_None   GrQuadAAFlags = 0
	GrQuadAAFlags_All    GrQuadAAFlags = 15
)

type GrTextureType int64

const (
	GrTextureType_None      GrTextureType = 0
	GrTextureType_2D        GrTextureType = 1
	GrTextureType_Rectangle GrTextureType = 2
	GrTextureType_External  GrTextureType = 3
)

type GrShaderType int64

const (
	GrShaderType_Vertex   GrShaderType = 0
	GrShaderType_Fragment GrShaderType = 1
	GrShaderType_Last     GrShaderType = 1
)

type GrShaderFlags int64

const (
	GrShaderFlags_None              GrShaderFlags = 0
	GrShaderFlags_Vertex_GrShader   GrShaderFlags = 1
	GrShaderFlags_Fragment_GrShader GrShaderFlags = 2
)

type GrVertexAttribType int64

const (
	GrVertexAttribType_Float        GrVertexAttribType = 0
	GrVertexAttribType_Float2       GrVertexAttribType = 1
	GrVertexAttribType_Float3       GrVertexAttribType = 2
	GrVertexAttribType_Float4       GrVertexAttribType = 3
	GrVertexAttribType_Half         GrVertexAttribType = 4
	GrVertexAttribType_Half2        GrVertexAttribType = 5
	GrVertexAttribType_Half4        GrVertexAttribType = 6
	GrVertexAttribType_Int2         GrVertexAttribType = 7
	GrVertexAttribType_Int3         GrVertexAttribType = 8
	GrVertexAttribType_Int4         GrVertexAttribType = 9
	GrVertexAttribType_Byte         GrVertexAttribType = 10
	GrVertexAttribType_Byte2        GrVertexAttribType = 11
	GrVertexAttribType_Byte4        GrVertexAttribType = 12
	GrVertexAttribType_UByte        GrVertexAttribType = 13
	GrVertexAttribType_UByte2       GrVertexAttribType = 14
	GrVertexAttribType_UByte4       GrVertexAttribType = 15
	GrVertexAttribType_UByte_norm   GrVertexAttribType = 16
	GrVertexAttribType_UByte4_norm  GrVertexAttribType = 17
	GrVertexAttribType_Short2       GrVertexAttribType = 18
	GrVertexAttribType_Short4       GrVertexAttribType = 19
	GrVertexAttribType_UShort2      GrVertexAttribType = 20
	GrVertexAttribType_UShort2_norm GrVertexAttribType = 21
	GrVertexAttribType_Int          GrVertexAttribType = 22
	GrVertexAttribType_UInt         GrVertexAttribType = 23
	GrVertexAttribType_UShort_norm  GrVertexAttribType = 24
	GrVertexAttribType_UShort4_norm GrVertexAttribType = 25
	GrVertexAttribType_Last         GrVertexAttribType = 25
)

type GrClipEdgeType int64

const (
	GrClipEdgeType_FillBW        GrClipEdgeType = 0
	GrClipEdgeType_FillAA        GrClipEdgeType = 1
	GrClipEdgeType_InverseFillBW GrClipEdgeType = 2
	GrClipEdgeType_InverseFillAA GrClipEdgeType = 3
	GrClipEdgeType_Last          GrClipEdgeType = 3
)

type GrIOType int64

const (
	GrIOType_Read  GrIOType = 0
	GrIOType_Write GrIOType = 1
	GrIOType_RW    GrIOType = 2
)

type GrGpuBufferType int64

const (
	GrGpuBufferType_Vertex       GrGpuBufferType = 0
	GrGpuBufferType_Index        GrGpuBufferType = 1
	GrGpuBufferType_DrawIndirect GrGpuBufferType = 2
	GrGpuBufferType_XferCpuToGpu GrGpuBufferType = 3
	GrGpuBufferType_XferGpuToCpu GrGpuBufferType = 4
	GrGpuBufferType_Uniform      GrGpuBufferType = 5
)

type GrAccessPattern int64

const (
	GrAccessPattern_Dynamic GrAccessPattern = 0
	GrAccessPattern_Static  GrAccessPattern = 1
	GrAccessPattern_Stream  GrAccessPattern = 2
	GrAccessPattern_Last    GrAccessPattern = 2
)

type GrInternalSurfaceFlags int64

const (
	GrInternalSurfaceFlags_None                        GrInternalSurfaceFlags = 0
	GrInternalSurfaceFlags_ReadOnly                    GrInternalSurfaceFlags = 1
	GrInternalSurfaceFlags_GLRTFBOIDIs0                GrInternalSurfaceFlags = 2
	GrInternalSurfaceFlags_RequiresManualMSAAResolve   GrInternalSurfaceFlags = 4
	GrInternalSurfaceFlags_FramebufferOnly             GrInternalSurfaceFlags = 8
	GrInternalSurfaceFlags_VkRTSupportsInputAttachment GrInternalSurfaceFlags = 16
)

type GrBackendObjectOwnership int64

const (
	GrBackendObjectOwnership_Borrowed GrBackendObjectOwnership = 0
	GrBackendObjectOwnership_Owned    GrBackendObjectOwnership = -1
)

type GpuPathRenderers int64

const (
	GpuPathRenderers_None             GpuPathRenderers = 0
	GpuPathRenderers_DashLine         GpuPathRenderers = 1
	GpuPathRenderers_Atlas            GpuPathRenderers = 2
	GpuPathRenderers_Tessellation     GpuPathRenderers = 4
	GpuPathRenderers_CoverageCounting GpuPathRenderers = 8
	GpuPathRenderers_AAHairline       GpuPathRenderers = 16
	GpuPathRenderers_AAConvex         GpuPathRenderers = 32
	GpuPathRenderers_AALinearizing    GpuPathRenderers = 64
	GpuPathRenderers_Small            GpuPathRenderers = 128
	GpuPathRenderers_Triangulating    GpuPathRenderers = 256
	GpuPathRenderers_Default          GpuPathRenderers = 511
)

type GrMipmapStatus int64

const (
	GrMipmapStatus_NotAllocated GrMipmapStatus = 0
	GrMipmapStatus_Dirty        GrMipmapStatus = 1
	GrMipmapStatus_Valid        GrMipmapStatus = 2
)

type GrColorType int64

const (
	GrColorType_Unknown          GrColorType = 0
	GrColorType_Alpha_8          GrColorType = 1
	GrColorType_BGR_565          GrColorType = 2
	GrColorType_RGB_565          GrColorType = 3
	GrColorType_ABGR_4444        GrColorType = 4
	GrColorType_RGBA_8888        GrColorType = 5
	GrColorType_RGBA_8888_SRGB   GrColorType = 6
	GrColorType_RGB_888x         GrColorType = 7
	GrColorType_RG_88            GrColorType = 8
	GrColorType_BGRA_8888        GrColorType = 9
	GrColorType_RGBA_1010102     GrColorType = 10
	GrColorType_BGRA_1010102     GrColorType = 11
	GrColorType_RGBA_10x6        GrColorType = 12
	GrColorType_Gray_8           GrColorType = 13
	GrColorType_GrayAlpha_88     GrColorType = 14
	GrColorType_Alpha_F16        GrColorType = 15
	GrColorType_RGBA_F16         GrColorType = 16
	GrColorType_RGBA_F16_Clamped GrColorType = 17
	GrColorType_RGBA_F32         GrColorType = 18
	GrColorType_Alpha_16         GrColorType = 19
	GrColorType_RG_1616          GrColorType = 20
	GrColorType_RG_F16           GrColorType = 21
	GrColorType_RGBA_16161616    GrColorType = 22
	GrColorType_Alpha_8xxx       GrColorType = 23
	GrColorType_Alpha_F32xxx     GrColorType = 24
	GrColorType_Gray_8xxx        GrColorType = 25
	GrColorType_R_8xxx           GrColorType = 26
	GrColorType_RGB_888          GrColorType = 27
	GrColorType_R_8              GrColorType = 28
	GrColorType_R_16             GrColorType = 29
	GrColorType_R_F16            GrColorType = 30
	GrColorType_Gray_F16         GrColorType = 31
	GrColorType_BGRA_4444        GrColorType = 32
	GrColorType_ARGB_4444        GrColorType = 33
	GrColorType_Last             GrColorType = 33
)

type GrColorTypeEncoding int64

const (
	GrColorTypeEncoding_Unorm     GrColorTypeEncoding = 0
	GrColorTypeEncoding_SRGBUnorm GrColorTypeEncoding = 1
	GrColorTypeEncoding_Float     GrColorTypeEncoding = 2
)

type GrDstSampleFlags int64

const (
	GrDstSampleFlags_None                   GrDstSampleFlags = 0
	GrDstSampleFlags_RequiresTextureBarrier GrDstSampleFlags = 1
	GrDstSampleFlags_AsInputAttachment      GrDstSampleFlags = 2
)

type GrDriverBugWorkaroundType int64

const (
	GrDriverBugWorkaroundType_ADD_AND_TRUE_TO_LOOP_CONDITION                GrDriverBugWorkaroundType = 0
	GrDriverBugWorkaroundType_DISABLE_BLEND_EQUATION_ADVANCED               GrDriverBugWorkaroundType = 1
	GrDriverBugWorkaroundType_DISABLE_DISCARD_FRAMEBUFFER                   GrDriverBugWorkaroundType = 2
	GrDriverBugWorkaroundType_DISABLE_DUAL_SOURCE_BLENDING_SUPPORT          GrDriverBugWorkaroundType = 3
	GrDriverBugWorkaroundType_DISABLE_TEXTURE_STORAGE                       GrDriverBugWorkaroundType = 4
	GrDriverBugWorkaroundType_DISALLOW_LARGE_INSTANCED_DRAW                 GrDriverBugWorkaroundType = 5
	GrDriverBugWorkaroundType_EMULATE_ABS_INT_FUNCTION                      GrDriverBugWorkaroundType = 6
	GrDriverBugWorkaroundType_FLUSH_ON_FRAMEBUFFER_CHANGE                   GrDriverBugWorkaroundType = 7
	GrDriverBugWorkaroundType_FORCE_UPDATE_SCISSOR_STATE_WHEN_BINDING_FBO0  GrDriverBugWorkaroundType = 8
	GrDriverBugWorkaroundType_GL_CLEAR_BROKEN                               GrDriverBugWorkaroundType = 9
	GrDriverBugWorkaroundType_MAX_FRAGMENT_UNIFORM_VECTORS_32               GrDriverBugWorkaroundType = 10
	GrDriverBugWorkaroundType_MAX_MSAA_SAMPLE_COUNT_4                       GrDriverBugWorkaroundType = 11
	GrDriverBugWorkaroundType_PACK_PARAMETERS_WORKAROUND_WITH_PACK_BUFFER   GrDriverBugWorkaroundType = 12
	GrDriverBugWorkaroundType_REMOVE_POW_WITH_CONSTANT_EXPONENT             GrDriverBugWorkaroundType = 13
	GrDriverBugWorkaroundType_REWRITE_DO_WHILE_LOOPS                        GrDriverBugWorkaroundType = 14
	GrDriverBugWorkaroundType_UNBIND_ATTACHMENTS_ON_BOUND_RENDER_FBO_DELETE GrDriverBugWorkaroundType = 15
	GrDriverBugWorkaroundType_UNFOLD_SHORT_CIRCUIT_AS_TERNARY_OPERATION     GrDriverBugWorkaroundType = 16
	GrDriverBugWorkaroundType_NUMBER_OF_GPU_DRIVER_BUG_WORKAROUND_TYPES     GrDriverBugWorkaroundType = 17
)

type PathOp int64

const (
	PathOp_Difference        PathOp = 0
	PathOp_Intersect         PathOp = 1
	PathOp_Union             PathOp = 2
	PathOp_XOR               PathOp = 3
	PathOp_ReverseDifference PathOp = 4
)

type PDFSteamCompressionEnabled int64

const (
	PDFSteamCompressionEnabled_No      PDFSteamCompressionEnabled = 0
	PDFSteamCompressionEnabled_Yes     PDFSteamCompressionEnabled = -1
	PDFSteamCompressionEnabled_Default PDFSteamCompressionEnabled = -1
)

type SVGPropertyState int64

const (
	SVGPropertyState_Unspecified SVGPropertyState = 0
	SVGPropertyState_Inherit     SVGPropertyState = 1
	SVGPropertyState_Value       SVGPropertyState = 2
)

type SVGLineCap int64

const (
	SVGLineCap_Butt   SVGLineCap = 0
	SVGLineCap_Round  SVGLineCap = 1
	SVGLineCap_Square SVGLineCap = 2
)

type SVGFeColorMatrixType int64

const (
	SVGFeColorMatrixType_Matrix           SVGFeColorMatrixType = 0
	SVGFeColorMatrixType_Saturate         SVGFeColorMatrixType = 1
	SVGFeColorMatrixType_HueRotate        SVGFeColorMatrixType = 2
	SVGFeColorMatrixType_LuminanceToAlpha SVGFeColorMatrixType = 3
)

type SVGFeCompositeOperator int64

const (
	SVGFeCompositeOperator_Over       SVGFeCompositeOperator = 0
	SVGFeCompositeOperator_In         SVGFeCompositeOperator = 1
	SVGFeCompositeOperator_Out        SVGFeCompositeOperator = 2
	SVGFeCompositeOperator_Atop       SVGFeCompositeOperator = 3
	SVGFeCompositeOperator_Xor        SVGFeCompositeOperator = 4
	SVGFeCompositeOperator_Arithmetic SVGFeCompositeOperator = 5
)

type SVGXmlSpace int64

const (
	SVGXmlSpace_Default  SVGXmlSpace = 0
	SVGXmlSpace_Preserve SVGXmlSpace = 1
)

type SVGColorspace int64

const (
	SVGColorspace_Auto      SVGColorspace = 0
	SVGColorspace_SRGB      SVGColorspace = 1
	SVGColorspace_LinearRGB SVGColorspace = 2
)

type SVGDisplay int64

const (
	SVGDisplay_Inline SVGDisplay = 0
	SVGDisplay_None   SVGDisplay = 1
)

type SVGFeFuncType int64

const (
	SVGFeFuncType_Identity SVGFeFuncType = 0
	SVGFeFuncType_Table    SVGFeFuncType = 1
	SVGFeFuncType_Discrete SVGFeFuncType = 2
	SVGFeFuncType_Linear   SVGFeFuncType = 3
	SVGFeFuncType_Gamma    SVGFeFuncType = 4
)

type SVGAttribute int64

const (
	SVGAttribute_ClipRule                  SVGAttribute = 0
	SVGAttribute_Color                     SVGAttribute = 1
	SVGAttribute_ColorInterpolation        SVGAttribute = 2
	SVGAttribute_ColorInterpolationFilters SVGAttribute = 3
	SVGAttribute_Cx                        SVGAttribute = 4
	SVGAttribute_Cy                        SVGAttribute = 5
	SVGAttribute_Fill                      SVGAttribute = 6
	SVGAttribute_FillOpacity               SVGAttribute = 7
	SVGAttribute_FillRule                  SVGAttribute = 8
	SVGAttribute_Filter                    SVGAttribute = 9
	SVGAttribute_FilterUnits               SVGAttribute = 10
	SVGAttribute_FontFamily                SVGAttribute = 11
	SVGAttribute_FontSize                  SVGAttribute = 12
	SVGAttribute_FontStyle                 SVGAttribute = 13
	SVGAttribute_FontWeight                SVGAttribute = 14
	SVGAttribute_Fx                        SVGAttribute = 15
	SVGAttribute_Fy                        SVGAttribute = 16
	SVGAttribute_GradientUnits             SVGAttribute = 17
	SVGAttribute_GradientTransform         SVGAttribute = 18
	SVGAttribute_Height                    SVGAttribute = 19
	SVGAttribute_Href                      SVGAttribute = 20
	SVGAttribute_Opacity                   SVGAttribute = 21
	SVGAttribute_Points                    SVGAttribute = 22
	SVGAttribute_PreserveAspectRatio       SVGAttribute = 23
	SVGAttribute_R                         SVGAttribute = 24
	SVGAttribute_Rx                        SVGAttribute = 25
	SVGAttribute_Ry                        SVGAttribute = 26
	SVGAttribute_SpreadMethod              SVGAttribute = 27
	SVGAttribute_Stroke                    SVGAttribute = 28
	SVGAttribute_StrokeDashArray           SVGAttribute = 29
	SVGAttribute_StrokeDashOffset          SVGAttribute = 30
	SVGAttribute_StrokeOpacity             SVGAttribute = 31
	SVGAttribute_StrokeLineCap             SVGAttribute = 32
	SVGAttribute_StrokeLineJoin            SVGAttribute = 33
	SVGAttribute_StrokeMiterLimit          SVGAttribute = 34
	SVGAttribute_StrokeWidth               SVGAttribute = 35
	SVGAttribute_Transform                 SVGAttribute = 36
	SVGAttribute_Text                      SVGAttribute = 37
	SVGAttribute_TextAnchor                SVGAttribute = 38
	SVGAttribute_ViewBox                   SVGAttribute = 39
	SVGAttribute_Visibility                SVGAttribute = 40
	SVGAttribute_Width                     SVGAttribute = 41
	SVGAttribute_X                         SVGAttribute = 42
	SVGAttribute_X1                        SVGAttribute = 43
	SVGAttribute_X2                        SVGAttribute = 44
	SVGAttribute_Y                         SVGAttribute = 45
	SVGAttribute_Y1                        SVGAttribute = 46
	SVGAttribute_Y2                        SVGAttribute = 47
	SVGAttribute_Unknown                   SVGAttribute = 48
)

type SVGTag int64

const (
	SVGTag_Circle              SVGTag = 0
	SVGTag_ClipPath            SVGTag = 1
	SVGTag_Defs                SVGTag = 2
	SVGTag_Ellipse             SVGTag = 3
	SVGTag_FeBlend             SVGTag = 4
	SVGTag_FeColorMatrix       SVGTag = 5
	SVGTag_FeComponentTransfer SVGTag = 6
	SVGTag_FeComposite         SVGTag = 7
	SVGTag_FeDiffuseLighting   SVGTag = 8
	SVGTag_FeDisplacementMap   SVGTag = 9
	SVGTag_FeDistantLight      SVGTag = 10
	SVGTag_FeFlood             SVGTag = 11
	SVGTag_FeFuncA             SVGTag = 12
	SVGTag_FeFuncR             SVGTag = 13
	SVGTag_FeFuncG             SVGTag = 14
	SVGTag_FeFuncB             SVGTag = 15
	SVGTag_FeGaussianBlur      SVGTag = 16
	SVGTag_FeImage             SVGTag = 17
	SVGTag_FeMerge             SVGTag = 18
	SVGTag_FeMergeNode         SVGTag = 19
	SVGTag_FeMorphology        SVGTag = 20
	SVGTag_FeOffset            SVGTag = 21
	SVGTag_FePointLight        SVGTag = 22
	SVGTag_FeSpecularLighting  SVGTag = 23
	SVGTag_FeSpotLight         SVGTag = 24
	SVGTag_FeTurbulence        SVGTag = 25
	SVGTag_Filter              SVGTag = 26
	SVGTag_G                   SVGTag = 27
	SVGTag_Image               SVGTag = 28
	SVGTag_Line                SVGTag = 29
	SVGTag_LinearGradient      SVGTag = 30
	SVGTag_Mask                SVGTag = 31
	SVGTag_Path                SVGTag = 32
	SVGTag_Pattern             SVGTag = 33
	SVGTag_Polygon             SVGTag = 34
	SVGTag_Polyline            SVGTag = 35
	SVGTag_RadialGradient      SVGTag = 36
	SVGTag_Rect                SVGTag = 37
	SVGTag_Stop                SVGTag = 38
	SVGTag_Svg                 SVGTag = 39
	SVGTag_Text                SVGTag = 40
	SVGTag_TextLiteral         SVGTag = 41
	SVGTag_TextPath            SVGTag = 42
	SVGTag_TSpan               SVGTag = 43
	SVGTag_Use                 SVGTag = 44
)
