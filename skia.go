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

type TextBlobGlyphPositioning int64

const ()

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
