// This is a generated file. DO NOT EDIT.

package skia

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

type YUVAPixmapInfoDataType int64

const (
	YUVAPixmapInfoDataType_Unorm8         YUVAPixmapInfoDataType = 0
	YUVAPixmapInfoDataType_Unorm16        YUVAPixmapInfoDataType = 1
	YUVAPixmapInfoDataType_Float16        YUVAPixmapInfoDataType = 2
	YUVAPixmapInfoDataType_Unorm10_Unorm2 YUVAPixmapInfoDataType = 3
	YUVAPixmapInfoDataType_Last           YUVAPixmapInfoDataType = 3
)

type OnceState int64

const (
	OnceState_NotStarted OnceState = 0
	OnceState_Claimed    OnceState = 1
	OnceState_Done       OnceState = 2
)

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

type ImageFilterMapDirection int64

const (
	ImageFilterMapDirection_Forward ImageFilterMapDirection = 0
	ImageFilterMapDirection_Reverse ImageFilterMapDirection = 1
)

type M44Uninitialized_Constructor int64

const (
	M44Uninitialized_Constructor_ M44Uninitialized_Constructor = 0
)

type M44NaN_Constructor int64

const (
	M44NaN_Constructor_ M44NaN_Constructor = 0
)

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

type SurfacePropsFlags int64

const (
	SurfacePropsFlags_Default                   SurfacePropsFlags = 0
	SurfacePropsFlags_UseDeviceIndependentFonts SurfacePropsFlags = 1
	SurfacePropsFlags_DynamicMSAA               SurfacePropsFlags = 2
	SurfacePropsFlags_AlwaysDither              SurfacePropsFlags = 4
)

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

type DocumentState int64

const (
	DocumentState_BetweenPages DocumentState = 0
	DocumentState_InPage       DocumentState = 1
	DocumentState_Closed       DocumentState = 2
)

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

type TypefaceSerializeBehavior int64

const (
	TypefaceSerializeBehavior_DoIncludeData      TypefaceSerializeBehavior = 0
	TypefaceSerializeBehavior_DontIncludeData    TypefaceSerializeBehavior = 1
	TypefaceSerializeBehavior_IncludeDataIfLocal TypefaceSerializeBehavior = 2
)

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

type Path1DPathEffectStyle int64

const (
	Path1DPathEffectStyle_Translate Path1DPathEffectStyle = 0
	Path1DPathEffectStyle_Rotate    Path1DPathEffectStyle = 1
	Path1DPathEffectStyle_Morph     Path1DPathEffectStyle = 2
	Path1DPathEffectStyle_LastEnum  Path1DPathEffectStyle = 2
)

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

type PathEffectDashType int64

const (
	PathEffectDashType_None PathEffectDashType = 0
	PathEffectDashType_Dash PathEffectDashType = 1
)

type GradientShaderFlags int64

const (
	GradientShaderFlags_InterpolateColorsInPremul GradientShaderFlags = 1
)

type ImageFiltersDither int64

const (
	ImageFiltersDither_No  ImageFiltersDither = 0
	ImageFiltersDither_Yes ImageFiltersDither = -1
)

type TrimPathEffectMode int64

const (
	TrimPathEffectMode_Normal   TrimPathEffectMode = 0
	TrimPathEffectMode_Inverted TrimPathEffectMode = 1
)

type ParsePathPathEncoding int64

const (
	ParsePathPathEncoding_Absolute ParsePathPathEncoding = 0
	ParsePathPathEncoding_Relative ParsePathPathEncoding = 1
)

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

type SVGIRIType int64

const (
	SVGIRIType_Local    SVGIRIType = 0
	SVGIRIType_Nonlocal SVGIRIType = 1
	SVGIRIType_DataURI  SVGIRIType = 2
)

type SVGColorType int64

const (
	SVGColorType_CurrentColor SVGColorType = 0
	SVGColorType_Color        SVGColorType = 1
	SVGColorType_ICCColor     SVGColorType = 2
)

type SVGPaintType int64

const (
	SVGPaintType_None  SVGPaintType = 0
	SVGPaintType_Color SVGPaintType = 1
	SVGPaintType_IRI   SVGPaintType = 2
)

type SVGFuncIRIType int64

const (
	SVGFuncIRIType_None SVGFuncIRIType = 0
	SVGFuncIRIType_IRI  SVGFuncIRIType = 1
)

type SVGLineJoinType int64

const (
	SVGLineJoinType_Miter   SVGLineJoinType = 0
	SVGLineJoinType_Round   SVGLineJoinType = 1
	SVGLineJoinType_Bevel   SVGLineJoinType = 2
	SVGLineJoinType_Inherit SVGLineJoinType = 3
)

type SVGSpreadMethodType int64

const (
	SVGSpreadMethodType_Pad     SVGSpreadMethodType = 0
	SVGSpreadMethodType_Repeat  SVGSpreadMethodType = 1
	SVGSpreadMethodType_Reflect SVGSpreadMethodType = 2
)

type SVGFillRuleType int64

const (
	SVGFillRuleType_NonZero SVGFillRuleType = 0
	SVGFillRuleType_EvenOdd SVGFillRuleType = 1
	SVGFillRuleType_Inherit SVGFillRuleType = 2
)

type SVGVisibilityType int64

const (
	SVGVisibilityType_Visible  SVGVisibilityType = 0
	SVGVisibilityType_Hidden   SVGVisibilityType = 1
	SVGVisibilityType_Collapse SVGVisibilityType = 2
	SVGVisibilityType_Inherit  SVGVisibilityType = 3
)

type SVGDashArrayType int64

const (
	SVGDashArrayType_None      SVGDashArrayType = 0
	SVGDashArrayType_DashArray SVGDashArrayType = 1
	SVGDashArrayType_Inherit   SVGDashArrayType = 2
)

type SVGStopColorType int64

const (
	SVGStopColorType_Color        SVGStopColorType = 0
	SVGStopColorType_CurrentColor SVGStopColorType = 1
	SVGStopColorType_ICCColor     SVGStopColorType = 2
	SVGStopColorType_Inherit      SVGStopColorType = 3
)

type SVGObjectBoundingBoxUnitsType int64

const (
	SVGObjectBoundingBoxUnitsType_UserSpaceOnUse    SVGObjectBoundingBoxUnitsType = 0
	SVGObjectBoundingBoxUnitsType_ObjectBoundingBox SVGObjectBoundingBoxUnitsType = 1
)

type SVGFontFamilyType int64

const (
	SVGFontFamilyType_Family  SVGFontFamilyType = 0
	SVGFontFamilyType_Inherit SVGFontFamilyType = 1
)

type SVGFontStyleType int64

const (
	SVGFontStyleType_Normal  SVGFontStyleType = 0
	SVGFontStyleType_Italic  SVGFontStyleType = 1
	SVGFontStyleType_Oblique SVGFontStyleType = 2
	SVGFontStyleType_Inherit SVGFontStyleType = 3
)

type SVGFontSizeType int64

const (
	SVGFontSizeType_Length  SVGFontSizeType = 0
	SVGFontSizeType_Inherit SVGFontSizeType = 1
)

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

type SVGTextAnchorType int64

const (
	SVGTextAnchorType_Start   SVGTextAnchorType = 0
	SVGTextAnchorType_Middle  SVGTextAnchorType = 1
	SVGTextAnchorType_End     SVGTextAnchorType = 2
	SVGTextAnchorType_Inherit SVGTextAnchorType = 3
)

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

type SVGSVGType int64

const (
	SVGSVGType_Root  SVGSVGType = 0
	SVGSVGType_Inner SVGSVGType = 1
)

type SVGLengthContextLengthType int64

const (
	SVGLengthContextLengthType_Horizontal SVGLengthContextLengthType = 0
	SVGLengthContextLengthType_Vertical   SVGLengthContextLengthType = 1
	SVGLengthContextLengthType_Other      SVGLengthContextLengthType = 2
)

type SVGRenderContextApplyFlags int64

const (
	SVGRenderContextApplyFlags_Leaf SVGRenderContextApplyFlags = 1
)

type SkApplyPerspectiveClip int64

const (
	SkApplyPerspectiveClip_No  SkApplyPerspectiveClip = 0
	SkApplyPerspectiveClip_Yes SkApplyPerspectiveClip = 1
)

type SkEncodedOrigin int64

const (
	SkEncodedOrigin_TopLeft     SkEncodedOrigin = 1
	SkEncodedOrigin_TopRight    SkEncodedOrigin = 2
	SkEncodedOrigin_BottomRight SkEncodedOrigin = 3
	SkEncodedOrigin_BottomLeft  SkEncodedOrigin = 4
	SkEncodedOrigin_LeftTop     SkEncodedOrigin = 5
	SkEncodedOrigin_RightTop    SkEncodedOrigin = 6
	SkEncodedOrigin_RightBottom SkEncodedOrigin = 7
	SkEncodedOrigin_LeftBottom  SkEncodedOrigin = 8
	SkEncodedOrigin_Default     SkEncodedOrigin = 1
	SkEncodedOrigin_Last        SkEncodedOrigin = 8
)

type SkAlphaType int64

const (
	SkAlphaType_Unknown  SkAlphaType = 0
	SkAlphaType_Opaque   SkAlphaType = 1
	SkAlphaType_Premul   SkAlphaType = 2
	SkAlphaType_Unpremul SkAlphaType = 3
	SkAlphaType_LastEnum SkAlphaType = 3
)

type SkColorType int64

const (
	SkColorType_Unknown            SkColorType = 0
	SkColorType_Alpha_8            SkColorType = 1
	SkColorType_RGB_565            SkColorType = 2
	SkColorType_ARGB_4444          SkColorType = 3
	SkColorType_RGBA_8888          SkColorType = 4
	SkColorType_RGB_888x           SkColorType = 5
	SkColorType_BGRA_8888          SkColorType = 6
	SkColorType_RGBA_1010102       SkColorType = 7
	SkColorType_BGRA_1010102       SkColorType = 8
	SkColorType_RGB_101010x        SkColorType = 9
	SkColorType_BGR_101010x        SkColorType = 10
	SkColorType_BGR_101010x_XR     SkColorType = 11
	SkColorType_BGRA_10101010_XR   SkColorType = 12
	SkColorType_RGBA_10x6          SkColorType = 13
	SkColorType_Gray_8             SkColorType = 14
	SkColorType_RGBA_F16Norm       SkColorType = 15
	SkColorType_RGBA_F16           SkColorType = 16
	SkColorType_RGBA_F32           SkColorType = 17
	SkColorType_R8G8_unorm         SkColorType = 18
	SkColorType_A16_float          SkColorType = 19
	SkColorType_R16G16_float       SkColorType = 20
	SkColorType_A16_unorm          SkColorType = 21
	SkColorType_R16G16_unorm       SkColorType = 22
	SkColorType_R16G16B16A16_unorm SkColorType = 23
	SkColorType_SRGBA_8888         SkColorType = 24
	SkColorType_R8_unorm           SkColorType = 25
	SkColorType_LastEnum           SkColorType = 25
	SkColorType_N32                SkColorType = 4
)

type SkYUVColorSpace int64

const (
	SkYUVColorSpace_JPEG_Full            SkYUVColorSpace = 0
	SkYUVColorSpace_Rec601_Limited       SkYUVColorSpace = 1
	SkYUVColorSpace_Rec709_Full          SkYUVColorSpace = 2
	SkYUVColorSpace_Rec709_Limited       SkYUVColorSpace = 3
	SkYUVColorSpace_BT2020_8bit_Full     SkYUVColorSpace = 4
	SkYUVColorSpace_BT2020_8bit_Limited  SkYUVColorSpace = 5
	SkYUVColorSpace_BT2020_10bit_Full    SkYUVColorSpace = 6
	SkYUVColorSpace_BT2020_10bit_Limited SkYUVColorSpace = 7
	SkYUVColorSpace_BT2020_12bit_Full    SkYUVColorSpace = 8
	SkYUVColorSpace_BT2020_12bit_Limited SkYUVColorSpace = 9
	SkYUVColorSpace_FCC_Full             SkYUVColorSpace = 10
	SkYUVColorSpace_FCC_Limited          SkYUVColorSpace = 11
	SkYUVColorSpace_SMPTE240_Full        SkYUVColorSpace = 12
	SkYUVColorSpace_SMPTE240_Limited     SkYUVColorSpace = 13
	SkYUVColorSpace_YDZDX_Full           SkYUVColorSpace = 14
	SkYUVColorSpace_YDZDX_Limited        SkYUVColorSpace = 15
	SkYUVColorSpace_GBR_Full             SkYUVColorSpace = 16
	SkYUVColorSpace_GBR_Limited          SkYUVColorSpace = 17
	SkYUVColorSpace_YCgCo_8bit_Full      SkYUVColorSpace = 18
	SkYUVColorSpace_YCgCo_8bit_Limited   SkYUVColorSpace = 19
	SkYUVColorSpace_YCgCo_10bit_Full     SkYUVColorSpace = 20
	SkYUVColorSpace_YCgCo_10bit_Limited  SkYUVColorSpace = 21
	SkYUVColorSpace_YCgCo_12bit_Full     SkYUVColorSpace = 22
	SkYUVColorSpace_YCgCo_12bit_Limited  SkYUVColorSpace = 23
	SkYUVColorSpace_Identity             SkYUVColorSpace = 24
	SkYUVColorSpace_LastEnum             SkYUVColorSpace = 24
	SkYUVColorSpace_JPEG                 SkYUVColorSpace = 0
	SkYUVColorSpace_Rec601               SkYUVColorSpace = 1
	SkYUVColorSpace_Rec709               SkYUVColorSpace = 3
	SkYUVColorSpace_BT2020               SkYUVColorSpace = 5
)

type SkColorChannel int64

const (
	SkColorChannel_R        SkColorChannel = 0
	SkColorChannel_G        SkColorChannel = 1
	SkColorChannel_B        SkColorChannel = 2
	SkColorChannel_A        SkColorChannel = 3
	SkColorChannel_LastEnum SkColorChannel = 3
)

type SkColorChannelFlag int64

const (
	SkColorChannelFlag_Red                           SkColorChannelFlag = 1
	SkColorChannelFlag_Green                         SkColorChannelFlag = 2
	SkColorChannelFlag_Blue                          SkColorChannelFlag = 4
	SkColorChannelFlag_Alpha                         SkColorChannelFlag = 8
	SkColorChannelFlag_Gray                          SkColorChannelFlag = 16
	SkColorChannelFlag_GrayAlpha_SkColorChannelFlags SkColorChannelFlag = 24
	SkColorChannelFlag_RG_SkColorChannelFlags        SkColorChannelFlag = 3
	SkColorChannelFlag_RGB_SkColorChannelFlags       SkColorChannelFlag = 7
	SkColorChannelFlag_RGBA_SkColorChannelFlags      SkColorChannelFlag = 15
)

type SkFilterMode int64

const (
	SkFilterMode_Nearest SkFilterMode = 0
	SkFilterMode_Linear  SkFilterMode = 1
	SkFilterMode_Last    SkFilterMode = 1
)

type SkMipmapMode int64

const (
	SkMipmapMode_None    SkMipmapMode = 0
	SkMipmapMode_Nearest SkMipmapMode = 1
	SkMipmapMode_Linear  SkMipmapMode = 2
	SkMipmapMode_Last    SkMipmapMode = 2
)

type SkBlurStyle int64

const (
	SkBlurStyle_Normal   SkBlurStyle = 0
	SkBlurStyle_Solid    SkBlurStyle = 1
	SkBlurStyle_Outer    SkBlurStyle = 2
	SkBlurStyle_Inner    SkBlurStyle = 3
	SkBlurStyle_LastEnum SkBlurStyle = 3
)

type SkBlendMode int64

const (
	SkBlendMode_Clear             SkBlendMode = 0
	SkBlendMode_Src               SkBlendMode = 1
	SkBlendMode_Dst               SkBlendMode = 2
	SkBlendMode_SrcOver           SkBlendMode = 3
	SkBlendMode_DstOver           SkBlendMode = 4
	SkBlendMode_SrcIn             SkBlendMode = 5
	SkBlendMode_DstIn             SkBlendMode = 6
	SkBlendMode_SrcOut            SkBlendMode = 7
	SkBlendMode_DstOut            SkBlendMode = 8
	SkBlendMode_SrcATop           SkBlendMode = 9
	SkBlendMode_DstATop           SkBlendMode = 10
	SkBlendMode_Xor               SkBlendMode = 11
	SkBlendMode_Plus              SkBlendMode = 12
	SkBlendMode_Modulate          SkBlendMode = 13
	SkBlendMode_Screen            SkBlendMode = 14
	SkBlendMode_Overlay           SkBlendMode = 15
	SkBlendMode_Darken            SkBlendMode = 16
	SkBlendMode_Lighten           SkBlendMode = 17
	SkBlendMode_ColorDodge        SkBlendMode = 18
	SkBlendMode_ColorBurn         SkBlendMode = 19
	SkBlendMode_HardLight         SkBlendMode = 20
	SkBlendMode_SoftLight         SkBlendMode = 21
	SkBlendMode_Difference        SkBlendMode = 22
	SkBlendMode_Exclusion         SkBlendMode = 23
	SkBlendMode_Multiply          SkBlendMode = 24
	SkBlendMode_Hue               SkBlendMode = 25
	SkBlendMode_Saturation        SkBlendMode = 26
	SkBlendMode_Color             SkBlendMode = 27
	SkBlendMode_Luminosity        SkBlendMode = 28
	SkBlendMode_LastCoeffMode     SkBlendMode = 14
	SkBlendMode_LastSeparableMode SkBlendMode = 24
	SkBlendMode_LastMode          SkBlendMode = 28
)

type SkBlendModeCoeff int64

const (
	SkBlendModeCoeff_Zero       SkBlendModeCoeff = 0
	SkBlendModeCoeff_One        SkBlendModeCoeff = 1
	SkBlendModeCoeff_SC         SkBlendModeCoeff = 2
	SkBlendModeCoeff_ISC        SkBlendModeCoeff = 3
	SkBlendModeCoeff_DC         SkBlendModeCoeff = 4
	SkBlendModeCoeff_IDC        SkBlendModeCoeff = 5
	SkBlendModeCoeff_SA         SkBlendModeCoeff = 6
	SkBlendModeCoeff_ISA        SkBlendModeCoeff = 7
	SkBlendModeCoeff_DA         SkBlendModeCoeff = 8
	SkBlendModeCoeff_IDA        SkBlendModeCoeff = 9
	SkBlendModeCoeff_CoeffCount SkBlendModeCoeff = 10
)

type SkClipOp int64

const (
	SkClipOp_Difference    SkClipOp = 0
	SkClipOp_Intersect     SkClipOp = 1
	SkClipOp_Max_EnumValue SkClipOp = 1
)

type SkTextEncoding int64

const (
	SkTextEncoding_UTF8    SkTextEncoding = 0
	SkTextEncoding_UTF16   SkTextEncoding = 1
	SkTextEncoding_UTF32   SkTextEncoding = 2
	SkTextEncoding_GlyphID SkTextEncoding = 3
)

type SkFontHinting int64

const (
	SkFontHinting_None   SkFontHinting = 0
	SkFontHinting_Slight SkFontHinting = 1
	SkFontHinting_Normal SkFontHinting = 2
	SkFontHinting_Full   SkFontHinting = 3
)

type SkPixelGeometry int64

const (
	SkPixelGeometry_Unknown SkPixelGeometry = 0
	SkPixelGeometry_RGB_H   SkPixelGeometry = 1
	SkPixelGeometry_BGR_H   SkPixelGeometry = 2
	SkPixelGeometry_RGB_V   SkPixelGeometry = 3
	SkPixelGeometry_BGR_V   SkPixelGeometry = 4
)

type SkPathFillType int64

const (
	SkPathFillType_Winding        SkPathFillType = 0
	SkPathFillType_EvenOdd        SkPathFillType = 1
	SkPathFillType_InverseWinding SkPathFillType = 2
	SkPathFillType_InverseEvenOdd SkPathFillType = 3
)

type SkPathDirection int64

const (
	SkPathDirection_CW  SkPathDirection = 0
	SkPathDirection_CCW SkPathDirection = 1
)

type SkPathSegmentMask int64

const (
	SkPathSegmentMask_Line  SkPathSegmentMask = 1
	SkPathSegmentMask_Quad  SkPathSegmentMask = 2
	SkPathSegmentMask_Conic SkPathSegmentMask = 4
	SkPathSegmentMask_Cubic SkPathSegmentMask = 8
)

type SkPathVerb int64

const (
	SkPathVerb_Move  SkPathVerb = 0
	SkPathVerb_Line  SkPathVerb = 1
	SkPathVerb_Quad  SkPathVerb = 2
	SkPathVerb_Conic SkPathVerb = 3
	SkPathVerb_Cubic SkPathVerb = 4
	SkPathVerb_Close SkPathVerb = 5
)

type SkTileMode int64

const (
	SkTileMode_Clamp        SkTileMode = 0
	SkTileMode_Repeat       SkTileMode = 1
	SkTileMode_Mirror       SkTileMode = 2
	SkTileMode_Decal        SkTileMode = 3
	SkTileMode_LastTileMode SkTileMode = 3
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

type SkTextureCompressionType int64

const (
	SkTextureCompressionType_None            SkTextureCompressionType = 0
	SkTextureCompressionType_ETC2_RGB8_UNORM SkTextureCompressionType = 1
	SkTextureCompressionType_BC1_RGB8_UNORM  SkTextureCompressionType = 2
	SkTextureCompressionType_BC1_RGBA8_UNORM SkTextureCompressionType = 3
	SkTextureCompressionType_Last            SkTextureCompressionType = 3
	SkTextureCompressionType_ETC1_RGB8       SkTextureCompressionType = 1
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

type SkPathOp int64

const (
	SkPathOp_Difference        SkPathOp = 0
	SkPathOp_Intersect         SkPathOp = 1
	SkPathOp_Union             SkPathOp = 2
	SkPathOp_XOR               SkPathOp = 3
	SkPathOp_ReverseDifference SkPathOp = 4
)

type SkPDFSteamCompressionEnabled int64

const (
	SkPDFSteamCompressionEnabled_No      SkPDFSteamCompressionEnabled = 0
	SkPDFSteamCompressionEnabled_Yes     SkPDFSteamCompressionEnabled = -1
	SkPDFSteamCompressionEnabled_Default SkPDFSteamCompressionEnabled = -1
)

type SkSVGPropertyState int64

const (
	SkSVGPropertyState_Unspecified SkSVGPropertyState = 0
	SkSVGPropertyState_Inherit     SkSVGPropertyState = 1
	SkSVGPropertyState_Value       SkSVGPropertyState = 2
)

type SkSVGLineCap int64

const (
	SkSVGLineCap_Butt   SkSVGLineCap = 0
	SkSVGLineCap_Round  SkSVGLineCap = 1
	SkSVGLineCap_Square SkSVGLineCap = 2
)

type SkSVGFeColorMatrixType int64

const (
	SkSVGFeColorMatrixType_Matrix           SkSVGFeColorMatrixType = 0
	SkSVGFeColorMatrixType_Saturate         SkSVGFeColorMatrixType = 1
	SkSVGFeColorMatrixType_HueRotate        SkSVGFeColorMatrixType = 2
	SkSVGFeColorMatrixType_LuminanceToAlpha SkSVGFeColorMatrixType = 3
)

type SkSVGFeCompositeOperator int64

const (
	SkSVGFeCompositeOperator_Over       SkSVGFeCompositeOperator = 0
	SkSVGFeCompositeOperator_In         SkSVGFeCompositeOperator = 1
	SkSVGFeCompositeOperator_Out        SkSVGFeCompositeOperator = 2
	SkSVGFeCompositeOperator_Atop       SkSVGFeCompositeOperator = 3
	SkSVGFeCompositeOperator_Xor        SkSVGFeCompositeOperator = 4
	SkSVGFeCompositeOperator_Arithmetic SkSVGFeCompositeOperator = 5
)

type SkSVGXmlSpace int64

const (
	SkSVGXmlSpace_Default  SkSVGXmlSpace = 0
	SkSVGXmlSpace_Preserve SkSVGXmlSpace = 1
)

type SkSVGColorspace int64

const (
	SkSVGColorspace_Auto      SkSVGColorspace = 0
	SkSVGColorspace_SRGB      SkSVGColorspace = 1
	SkSVGColorspace_LinearRGB SkSVGColorspace = 2
)

type SkSVGDisplay int64

const (
	SkSVGDisplay_Inline SkSVGDisplay = 0
	SkSVGDisplay_None   SkSVGDisplay = 1
)

type SkSVGFeFuncType int64

const (
	SkSVGFeFuncType_Identity SkSVGFeFuncType = 0
	SkSVGFeFuncType_Table    SkSVGFeFuncType = 1
	SkSVGFeFuncType_Discrete SkSVGFeFuncType = 2
	SkSVGFeFuncType_Linear   SkSVGFeFuncType = 3
	SkSVGFeFuncType_Gamma    SkSVGFeFuncType = 4
)

type SkSVGAttribute int64

const (
	SkSVGAttribute_ClipRule                  SkSVGAttribute = 0
	SkSVGAttribute_Color                     SkSVGAttribute = 1
	SkSVGAttribute_ColorInterpolation        SkSVGAttribute = 2
	SkSVGAttribute_ColorInterpolationFilters SkSVGAttribute = 3
	SkSVGAttribute_Cx                        SkSVGAttribute = 4
	SkSVGAttribute_Cy                        SkSVGAttribute = 5
	SkSVGAttribute_Fill                      SkSVGAttribute = 6
	SkSVGAttribute_FillOpacity               SkSVGAttribute = 7
	SkSVGAttribute_FillRule                  SkSVGAttribute = 8
	SkSVGAttribute_Filter                    SkSVGAttribute = 9
	SkSVGAttribute_FilterUnits               SkSVGAttribute = 10
	SkSVGAttribute_FontFamily                SkSVGAttribute = 11
	SkSVGAttribute_FontSize                  SkSVGAttribute = 12
	SkSVGAttribute_FontStyle                 SkSVGAttribute = 13
	SkSVGAttribute_FontWeight                SkSVGAttribute = 14
	SkSVGAttribute_Fx                        SkSVGAttribute = 15
	SkSVGAttribute_Fy                        SkSVGAttribute = 16
	SkSVGAttribute_GradientUnits             SkSVGAttribute = 17
	SkSVGAttribute_GradientTransform         SkSVGAttribute = 18
	SkSVGAttribute_Height                    SkSVGAttribute = 19
	SkSVGAttribute_Href                      SkSVGAttribute = 20
	SkSVGAttribute_Opacity                   SkSVGAttribute = 21
	SkSVGAttribute_Points                    SkSVGAttribute = 22
	SkSVGAttribute_PreserveAspectRatio       SkSVGAttribute = 23
	SkSVGAttribute_R                         SkSVGAttribute = 24
	SkSVGAttribute_Rx                        SkSVGAttribute = 25
	SkSVGAttribute_Ry                        SkSVGAttribute = 26
	SkSVGAttribute_SpreadMethod              SkSVGAttribute = 27
	SkSVGAttribute_Stroke                    SkSVGAttribute = 28
	SkSVGAttribute_StrokeDashArray           SkSVGAttribute = 29
	SkSVGAttribute_StrokeDashOffset          SkSVGAttribute = 30
	SkSVGAttribute_StrokeOpacity             SkSVGAttribute = 31
	SkSVGAttribute_StrokeLineCap             SkSVGAttribute = 32
	SkSVGAttribute_StrokeLineJoin            SkSVGAttribute = 33
	SkSVGAttribute_StrokeMiterLimit          SkSVGAttribute = 34
	SkSVGAttribute_StrokeWidth               SkSVGAttribute = 35
	SkSVGAttribute_Transform                 SkSVGAttribute = 36
	SkSVGAttribute_Text                      SkSVGAttribute = 37
	SkSVGAttribute_TextAnchor                SkSVGAttribute = 38
	SkSVGAttribute_ViewBox                   SkSVGAttribute = 39
	SkSVGAttribute_Visibility                SkSVGAttribute = 40
	SkSVGAttribute_Width                     SkSVGAttribute = 41
	SkSVGAttribute_X                         SkSVGAttribute = 42
	SkSVGAttribute_X1                        SkSVGAttribute = 43
	SkSVGAttribute_X2                        SkSVGAttribute = 44
	SkSVGAttribute_Y                         SkSVGAttribute = 45
	SkSVGAttribute_Y1                        SkSVGAttribute = 46
	SkSVGAttribute_Y2                        SkSVGAttribute = 47
	SkSVGAttribute_Unknown                   SkSVGAttribute = 48
)

type SkSVGTag int64

const (
	SkSVGTag_Circle              SkSVGTag = 0
	SkSVGTag_ClipPath            SkSVGTag = 1
	SkSVGTag_Defs                SkSVGTag = 2
	SkSVGTag_Ellipse             SkSVGTag = 3
	SkSVGTag_FeBlend             SkSVGTag = 4
	SkSVGTag_FeColorMatrix       SkSVGTag = 5
	SkSVGTag_FeComponentTransfer SkSVGTag = 6
	SkSVGTag_FeComposite         SkSVGTag = 7
	SkSVGTag_FeDiffuseLighting   SkSVGTag = 8
	SkSVGTag_FeDisplacementMap   SkSVGTag = 9
	SkSVGTag_FeDistantLight      SkSVGTag = 10
	SkSVGTag_FeFlood             SkSVGTag = 11
	SkSVGTag_FeFuncA             SkSVGTag = 12
	SkSVGTag_FeFuncR             SkSVGTag = 13
	SkSVGTag_FeFuncG             SkSVGTag = 14
	SkSVGTag_FeFuncB             SkSVGTag = 15
	SkSVGTag_FeGaussianBlur      SkSVGTag = 16
	SkSVGTag_FeImage             SkSVGTag = 17
	SkSVGTag_FeMerge             SkSVGTag = 18
	SkSVGTag_FeMergeNode         SkSVGTag = 19
	SkSVGTag_FeMorphology        SkSVGTag = 20
	SkSVGTag_FeOffset            SkSVGTag = 21
	SkSVGTag_FePointLight        SkSVGTag = 22
	SkSVGTag_FeSpecularLighting  SkSVGTag = 23
	SkSVGTag_FeSpotLight         SkSVGTag = 24
	SkSVGTag_FeTurbulence        SkSVGTag = 25
	SkSVGTag_Filter              SkSVGTag = 26
	SkSVGTag_G                   SkSVGTag = 27
	SkSVGTag_Image               SkSVGTag = 28
	SkSVGTag_Line                SkSVGTag = 29
	SkSVGTag_LinearGradient      SkSVGTag = 30
	SkSVGTag_Mask                SkSVGTag = 31
	SkSVGTag_Path                SkSVGTag = 32
	SkSVGTag_Pattern             SkSVGTag = 33
	SkSVGTag_Polygon             SkSVGTag = 34
	SkSVGTag_Polyline            SkSVGTag = 35
	SkSVGTag_RadialGradient      SkSVGTag = 36
	SkSVGTag_Rect                SkSVGTag = 37
	SkSVGTag_Stop                SkSVGTag = 38
	SkSVGTag_Svg                 SkSVGTag = 39
	SkSVGTag_Text                SkSVGTag = 40
	SkSVGTag_TextLiteral         SkSVGTag = 41
	SkSVGTag_TextPath            SkSVGTag = 42
	SkSVGTag_TSpan               SkSVGTag = 43
	SkSVGTag_Use                 SkSVGTag = 44
)
