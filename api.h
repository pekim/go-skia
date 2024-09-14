// This is a generated file. DO NOT EDIT.

#include <stdbool.h>
#include <sys/types.h>
typedef signed char schar;
typedef unsigned char uchar;

#ifdef __cplusplus
extern "C"
{
#endif // __cplusplus

  typedef struct
  {
    unsigned int FBOID;
    unsigned int Format;
    uchar Protected[1];
    uchar pad_3[3];
  } sk_GrGLFramebufferInfo;

  typedef struct
  {
    bool fIsValid;
    bool fFramebufferOnly;
    uchar pad_2[2];
    int fWidth;
    int fHeight;
    int fSampleCnt;
    int fStencilBits;
    uchar fBackend[4];
    uchar fRTData[184];
    uchar pad_8[16];
  } sk_GrBackendRenderTarget;

  typedef struct
  {
    uchar pad_0[120];
    uchar fDeleteCallbackHelper[8];
    uchar fDirectContextID[4];
    uchar pad_2[4];
    uchar fTaskGroup[8];
    uchar fStrikeCache[8];
    uchar fGpu[8];
    uchar fResourceCache[8];
    uchar fResourceProvider[8];
    int fInsideReleaseProcCnt;
    bool fDidTestPMConversions;
    bool fPMUPMConversionsRoundTrip;
    uchar pad_10[2];
    uchar fPersistentCache[8];
    uchar fMappedBufferManager[8];
    uchar fAtlasManager[8];
    uchar fSmallPathAtlasMgr[8];
  } sk_GrDirectContext;

  typedef struct
  {
    uchar pad_0[64];
    uchar fStats[8];
    uchar fAuditTrail[8];
    uchar fArenas[24];
    uchar fDrawingManager[8];
    uchar fProxyProvider[8];
  } sk_GrRecordingContext;

  typedef struct
  {
    uchar pad_0[12];
    bool fOOMed;
    bool fSuppressErrorLogging;
    uchar pad_2[2];
    uchar Standard[4];
    uchar pad_3[4];
    uchar Extensions[24];
    uchar Functions[8240];
  } sk_GrGLInterface;

  typedef struct
  {
    bool SuppressPrints;
    uchar pad_1[3];
    uchar SkipGLErrorChecks[4];
    int MaxTextureSizeOverride;
    int BufferMapThreshold;
    uchar MinimumStagingBufferSize[8];
    uchar Executor[8];
    bool DoManualMipmapping;
    bool DisableCoverageCountingPaths;
    bool DisableDistanceFieldPaths;
    bool AllowPathMaskCaching;
    bool DisableGpuYUVConversion;
    uchar pad_11[3];
    uchar GlyphCacheTextureMaximumBytes[8];
    float MinDistanceFieldFontSize;
    float GlyphsAsPathsFontSize;
    uchar AllowMultipleGlyphCacheTextures[4];
    bool AvoidStencilBuffers;
    uchar pad_16[3];
    uchar UseDrawInsteadOfClear[4];
    uchar ReduceOpsTaskSplitting[4];
    bool PreferExternalImagesOverES3;
    bool DisableDriverCorrectnessWorkarounds;
    uchar pad_20[2];
    int RuntimeProgramCacheSize;
    uchar PersistentCache[8];
    uchar ShaderCacheStrategy[4];
    uchar pad_23[4];
    uchar ShaderErrorHandler[8];
    int InternalMultisampleCount;
    int MaxCachedVulkanSecondaryCommandBuffers;
    bool SuppressMipmapSupport;
    bool DisableTessellationPathRenderer;
    bool EnableExperimentalHardwareTessellation;
    bool SupportBilerpFromGlyphAtlas;
    bool ReducedShaderVariations;
    bool AllowMSAAOnNewIntel;
    bool AlwaysUseTexStorageWhenAvailable;
    uchar pad_33[1];
    uchar ContextDeleteContext[8];
    uchar ContextDeleteProc[8];
    uchar DriverBugWorkarounds[17];
    uchar pad_36[7];
  } sk_GrContextOptions;

  typedef struct
  {
    uchar Oval[16];
    float StartAngle;
    float SweepAngle;
    uchar Type[1];
    uchar pad_4[3];
  } sk_SkArc;

  typedef struct
  {
    uchar fPixelRef[8];
    uchar fPixmap[40];
    uchar fMips[8];
  } sk_SkBitmap;

  typedef struct
  {
    uchar pad_0[8];
    uchar fMCRecStorage[3072];
    uchar fMCStack[56];
    uchar fMCRec[8];
    uchar fRootDevice[8];
    uchar fProps[16];
    int fSaveCount;
    uchar pad_6[4];
    uchar fAllocator[8];
    uchar fSurfaceBase[8];
    uchar fClipRestrictionRect[16];
    int fClipRestrictionSaveCount;
    uchar fQuickRejectBounds[16];
    uchar pad_11[4];
    uchar fScratchGlyphRunBuilder[8];
  } sk_SkCanvas;

  typedef struct
  {
    uchar pad_0[4];
    unsigned int fTransferFnHash;
    unsigned int fToXYZD50Hash;
    uchar fTransferFn[28];
    uchar fToXYZD50[36];
    uchar fInvTransferFn[28];
    uchar fFromXYZD50[36];
    uchar fLazyDstFieldsOnce[1];
    uchar pad_7[3];
  } sk_SkColorSpace;

  typedef struct
  {
    uchar fTypeface[8];
    float fSize;
    float fScaleX;
    float fSkewX;
    unsigned char fFlags;
    unsigned char fEdging;
    unsigned char fHinting;
    uchar pad_7[1];
  } sk_SkFont;

  typedef struct
  {
    uchar pad_0[16];
  } sk_SkFontMgr;

  typedef struct
  {
    int fValue;
  } sk_SkFontStyle;

  typedef struct
  {
    uchar pad_0[16];
    uchar fInfo[24];
    unsigned int fUniqueID;
    uchar pad_2[4];
  } sk_SkImage;

  typedef struct
  {
    uchar fColorInfo[16];
    uchar fDimensions[8];
  } sk_SkImageInfo;

  typedef struct
  {
    uchar fMat[64];
  } sk_SkM44;

  typedef struct
  {
    int X;
    int Y;
  } sk_SkIPoint;

  typedef struct
  {
    int Left;
    int Top;
    int Right;
    int Bottom;
  } sk_SkIRect;

  typedef struct
  {
    int Width;
    int Height;
  } sk_SkISize;

  typedef struct
  {
    uchar fMat[36];
    int fTypeMask;
  } sk_SkMatrix;

  typedef struct
  {
    uchar fPathEffect[8];
    uchar fShader[8];
    uchar fMaskFilter[8];
    uchar fColorFilter[8];
    uchar fImageFilter[8];
    uchar fBlender[8];
    uchar fColor4f[16];
    float fWidth;
    float fMiterLimit;
    uchar pad_9[8];
  } sk_SkPaint;

  typedef struct
  {
    uchar fPathRef[8];
    int fLastMoveToIndex;
    uchar fConvexity[1];
    uchar fFirstDirection[1];
    unsigned char fFillType;
    // TODO misalignment (perhaps there are bitfields around here?)
    unsigned char fIsVolatile;
  } sk_SkPath;

  typedef struct
  {
    float X;
    float Y;
  } sk_SkPoint;

  typedef struct
  {
    uchar fPixels[8];
    uchar fRowBytes[8];
    uchar fInfo[24];
  } sk_SkPixmap;

  typedef struct
  {
    float Left;
    float Top;
    float Right;
    float Bottom;
  } sk_SkRect;

  typedef struct
  {
    uchar fRect[16];
    uchar fRadii[32];
    int fType;
  } sk_SkRRect;

  typedef struct
  {
    uchar fBounds[16];
    uchar fRunHead[8];
  } sk_SkRegion;

  typedef struct
  {
    float R;
    // TODO misalignment (perhaps there are bitfields around here?)
    float G;
    // TODO misalignment (perhaps there are bitfields around here?)
    float B;
    // TODO misalignment (perhaps there are bitfields around here?)
    float A;
  } sk_SkRGBA4f;

  typedef struct
  {
    int maxAniso;
    bool useCubic;
    uchar pad_2[3];
    uchar cubic[8];
    uchar ilter[4];
    uchar mipmap[4];
  } sk_SkSamplingOptions;

  typedef struct
  {
    float Width;
    float Height;
  } sk_SkSize;

  typedef struct
  {
    uchar pad_0[12];
    uchar fProps[16];
    int fWidth;
    int fHeight;
    unsigned int fGenerationID;
  } sk_SkSurface;

  typedef struct
  {
    unsigned int fFlags;
    uchar fPixelGeometry[4];
    float fTextContrast;
    float fTextGamma;
  } sk_SkSurfaceProps;

  typedef struct
  {
    uchar pad_0[16];
    unsigned int fUniqueID;
    uchar fStyle[4];
    uchar fBounds[16];
    uchar fBoundsOnce[1];
    bool fIsFixedPitch;
    uchar pad_5[6];
  } sk_SkTypeface;

  sk_GrBackendRenderTarget *misk_new_GrBackendRenderTarget ();
  sk_GrBackendRenderTarget *
  misk_new_GrBackendRenderTargetCopy (sk_GrBackendRenderTarget *c_that);
  void misk_delete_GrBackendRenderTarget (sk_GrBackendRenderTarget *obj);

  void misk_delete_GrDirectContext (sk_GrDirectContext *obj);
  void misk_GrDirectContext_flushAndSubmit (sk_GrDirectContext *c_obj,
                                            int c_sync);

  void misk_delete_GrRecordingContext (sk_GrRecordingContext *obj);

  sk_GrContextOptions *misk_new_GrContextOptions ();

  sk_SkArc *misk_new_Arc ();
  sk_SkArc *misk_new_ArcCopy (sk_SkArc *c_arc);

  sk_SkBitmap *misk_new_Bitmap ();
  sk_SkBitmap *misk_new_BitmapCopy (sk_SkBitmap *c_src);
  void misk_delete_SkBitmap (sk_SkBitmap *obj);
  bool misk_Bitmap_ComputeIsOpaque (sk_SkBitmap *c_bm);

  sk_SkCanvas *misk_new_Canvas ();
  sk_SkCanvas *misk_new_CanvasWithDimensions (int c_width, int c_height,
                                              sk_SkSurfaceProps *c_props);
  sk_SkCanvas *misk_new_CanvasFromBitmap (sk_SkBitmap *c_bitmap);
  sk_SkCanvas *
  misk_new_CanvasFromBitmapSurfaceProps (sk_SkBitmap *c_bitmap,
                                         sk_SkSurfaceProps *c_props);
  void misk_delete_SkCanvas (sk_SkCanvas *obj);
  bool misk_Canvas_getProps (sk_SkCanvas *c_obj, sk_SkSurfaceProps *c_props);
  bool misk_Canvas_peekPixels (sk_SkCanvas *c_obj, sk_SkPixmap *c_pixmap);
  int misk_Canvas_save (sk_SkCanvas *c_obj);
  void misk_Canvas_restore (sk_SkCanvas *c_obj);
  int misk_Canvas_getSaveCount (sk_SkCanvas *c_obj);
  void misk_Canvas_restoreToCount (sk_SkCanvas *c_obj, int c_saveCount);
  void misk_Canvas_translate (sk_SkCanvas *c_obj, float c_dx, float c_dy);
  void misk_Canvas_scale (sk_SkCanvas *c_obj, float c_sx, float c_sy);
  void misk_Canvas_rotate (sk_SkCanvas *c_obj, float c_degrees);
  void misk_Canvas_rotateAboutPoint (sk_SkCanvas *c_obj, float c_degrees,
                                     float c_px, float c_py);
  void misk_Canvas_skew (sk_SkCanvas *c_obj, float c_sx, float c_sy);
  void misk_Canvas_setMatrixM44 (sk_SkCanvas *c_obj, sk_SkM44 *c_matrix);
  void misk_Canvas_setMatrix (sk_SkCanvas *c_obj, sk_SkMatrix *c_matrix);
  void misk_Canvas_resetMatrix (sk_SkCanvas *c_obj);
  void misk_Canvas_clipRect (sk_SkCanvas *c_obj, sk_SkRect c_rect, int c_op,
                             bool c_doAntiAlias);
  void misk_Canvas_clipRRect (sk_SkCanvas *c_obj, sk_SkRRect c_rrect, int c_op,
                              bool c_doAntiAlias);
  void misk_Canvas_clipPath (sk_SkCanvas *c_obj, sk_SkPath *c_path, int c_op,
                             bool c_doAntiAlias);
  void misk_Canvas_clipRegion (sk_SkCanvas *c_obj, sk_SkRegion *c_deviceRgn,
                               int c_op);
  bool misk_Canvas_quickRejectRect (sk_SkCanvas *c_obj, sk_SkRect c_rect);
  bool misk_Canvas_quickRejectPath (sk_SkCanvas *c_obj, sk_SkPath *c_path);
  sk_SkRect misk_Canvas_getLocalClipBoundsRect (sk_SkCanvas *c_obj);
  bool misk_Canvas_getLocalClipBoundsPath (sk_SkCanvas *c_obj,
                                           sk_SkRect c_bounds);
  sk_SkIRect misk_Canvas_getDeviceClipBounds (sk_SkCanvas *c_obj);
  bool misk_Canvas_getDeviceClipBoundsRect (sk_SkCanvas *c_obj,
                                            sk_SkIRect c_bounds);
  void misk_Canvas_drawColor (sk_SkCanvas *c_obj, uint c_color, int c_mode);
  void misk_Canvas_drawColor4f (sk_SkCanvas *c_obj, sk_SkRGBA4f c_color,
                                int c_mode);
  void misk_Canvas_clear (sk_SkCanvas *c_obj, uint c_color);
  void misk_Canvas_clear4f (sk_SkCanvas *c_obj, sk_SkRGBA4f c_color);
  void misk_Canvas_discard (sk_SkCanvas *c_obj);
  void misk_Canvas_drawPaint (sk_SkCanvas *c_obj, sk_SkPaint *c_paint);
  void misk_Canvas_drawPointScalars (sk_SkCanvas *c_obj, float c_x, float c_y,
                                     sk_SkPaint *c_paint);
  void misk_Canvas_drawPoint (sk_SkCanvas *c_obj, sk_SkPoint c_p,
                              sk_SkPaint *c_paint);
  void misk_Canvas_drawLineScalars (sk_SkCanvas *c_obj, float c_x0, float c_y0,
                                    float c_x1, float c_y1,
                                    sk_SkPaint *c_paint);
  void misk_Canvas_drawLinePoints (sk_SkCanvas *c_obj, sk_SkPoint c_p0,
                                   sk_SkPoint c_p1, sk_SkPaint *c_paint);
  void misk_Canvas_drawRect (sk_SkCanvas *c_obj, sk_SkRect c_rect,
                             sk_SkPaint *c_paint);
  void misk_Canvas_drawIRect (sk_SkCanvas *c_obj, sk_SkIRect c_rect,
                              sk_SkPaint *c_paint);
  void misk_Canvas_drawRegion (sk_SkCanvas *c_obj, sk_SkRegion *c_region,
                               sk_SkPaint *c_paint);
  void misk_Canvas_drawOval (sk_SkCanvas *c_obj, sk_SkRect c_oval,
                             sk_SkPaint *c_paint);
  void misk_Canvas_drawRRect (sk_SkCanvas *c_obj, sk_SkRRect c_rrect,
                              sk_SkPaint *c_paint);
  void misk_Canvas_drawDRRect (sk_SkCanvas *c_obj, sk_SkRRect c_outer,
                               sk_SkRRect c_inner, sk_SkPaint *c_paint);
  void misk_Canvas_drawCircleScalars (sk_SkCanvas *c_obj, float c_cx,
                                      float c_cy, float c_radius,
                                      sk_SkPaint *c_paint);
  void misk_Canvas_drawCirclePoint (sk_SkCanvas *c_obj, sk_SkPoint c_center,
                                    float c_radius, sk_SkPaint *c_paint);
  void misk_Canvas_drawArc (sk_SkCanvas *c_obj, sk_SkRect c_oval,
                            float c_startAngle, float c_sweepAngle,
                            bool c_useCenter, sk_SkPaint *c_paint);
  void misk_Canvas_drawArcArc (sk_SkCanvas *c_obj, sk_SkArc *c_arc,
                               sk_SkPaint *c_paint);
  void misk_Canvas_drawRoundRect (sk_SkCanvas *c_obj, sk_SkRect c_rect,
                                  float c_rx, float c_ry, sk_SkPaint *c_paint);
  void misk_Canvas_drawPath (sk_SkCanvas *c_obj, sk_SkPath *c_path,
                             sk_SkPaint *c_paint);
  void misk_Canvas_drawImage (sk_SkCanvas *c_obj, sk_SkImage *c_image,
                              float c_left, float c_top);
  void misk_Canvas_drawImageSamplingOptions (sk_SkCanvas *c_obj,
                                             sk_SkImage *c_p0, float c_x,
                                             float c_y,
                                             sk_SkSamplingOptions *c_p3,
                                             sk_SkPaint *c_p4);
  void misk_Canvas_drawImageRect (sk_SkCanvas *c_obj, sk_SkImage *c_p0,
                                  sk_SkRect c_src, sk_SkRect c_dst,
                                  sk_SkSamplingOptions *c_p3, sk_SkPaint *c_p4,
                                  uint c_p5);
  void misk_Canvas_drawImageNine (sk_SkCanvas *c_obj, sk_SkImage *c_image,
                                  sk_SkIRect c_center, sk_SkRect c_dst,
                                  int c_filter, sk_SkPaint *c_paint);

  sk_SkColorSpace *misk_ColorSpace_MakeSRGB ();
  sk_SkColorSpace *misk_ColorSpace_MakeSRGBLinear ();
  bool misk_ColorSpace_Equals (sk_SkColorSpace *c_p0, sk_SkColorSpace *c_p1);

  sk_SkFont *misk_new_Font ();
  sk_SkFont *misk_new_FontTypefaceSize (sk_SkTypeface *c_typeface,
                                        float c_size);
  sk_SkFont *misk_new_FontTypeface (sk_SkTypeface *c_typeface);
  sk_SkFont *misk_new_FontTypefaceSizeScaleSkew (sk_SkTypeface *c_typeface,
                                                 float c_size, float c_scaleX,
                                                 float c_skewX);

  sk_SkFontStyle *misk_new_FontStyle2 (int c_weight, int c_width,
                                       uint c_slant);
  sk_SkFontStyle *misk_new_FontStyle ();
  sk_SkFontStyle misk_FontStyle_Normal ();
  sk_SkFontStyle misk_FontStyle_Bold ();
  sk_SkFontStyle misk_FontStyle_Italic ();
  sk_SkFontStyle misk_FontStyle_BoldItalic ();

  sk_SkImageInfo *misk_new_ImageInfo ();

  sk_SkM44 *misk_new_M44Copy (sk_SkM44 *c_src);
  sk_SkM44 *misk_new_M44 ();
  sk_SkM44 *misk_new_M44AB (sk_SkM44 *c_a, sk_SkM44 *c_b);
  sk_SkM44 *misk_new_M44Scalars (float c_m0, float c_m4, float c_m8,
                                 float c_m12, float c_m1, float c_m5,
                                 float c_m9, float c_m13, float c_m2,
                                 float c_m6, float c_m10, float c_m14,
                                 float c_m3, float c_m7, float c_m11,
                                 float c_m15);

  sk_SkIRect misk_IRect_MakeEmpty ();
  sk_SkIRect misk_IRect_MakeWH (int c_w, int c_h);
  sk_SkIRect misk_IRect_MakeLTRB (int c_l, int c_t, int c_r, int c_b);
  sk_SkIRect misk_IRect_MakeXYWH (int c_x, int c_y, int c_w, int c_h);
  sk_SkIRect misk_IRect_MakeSize (sk_SkISize c_size);
  bool misk_IRect_Intersects (sk_SkIRect c_a, sk_SkIRect c_b);
  int misk_IRect_x (sk_SkIRect *c_obj);
  int misk_IRect_y (sk_SkIRect *c_obj);
  int misk_IRect_width (sk_SkIRect *c_obj);
  int misk_IRect_height (sk_SkIRect *c_obj);
  bool misk_IRect_isEmpty (sk_SkIRect *c_obj);
  void misk_IRect_setEmpty (sk_SkIRect *c_obj);
  void misk_IRect_setLTRB (sk_SkIRect *c_obj, int c_left, int c_top,
                           int c_right, int c_bottom);
  void misk_IRect_setXYWH (sk_SkIRect *c_obj, int c_x, int c_y, int c_width,
                           int c_height);
  void misk_IRect_setWH (sk_SkIRect *c_obj, int c_width, int c_height);
  void misk_IRect_offset (sk_SkIRect *c_obj, int c_dx, int c_dy);
  void misk_IRect_offsetPoint (sk_SkIRect *c_obj, sk_SkIPoint *c_delta);
  void misk_IRect_offsetTo (sk_SkIRect *c_obj, int c_newX, int c_newY);
  void misk_IRect_inset (sk_SkIRect *c_obj, int c_dx, int c_dy);
  void misk_IRect_outset (sk_SkIRect *c_obj, int c_dx, int c_dy);
  void misk_IRect_adjust (sk_SkIRect *c_obj, int c_dL, int c_dT, int c_dR,
                          int c_dB);
  bool misk_IRect_contains (sk_SkIRect *c_obj, int c_x, int c_y);
  bool misk_IRect_containsRect (sk_SkIRect *c_obj, sk_SkIRect c_r);
  bool misk_IRect_containsNoEmptyCheck (sk_SkIRect *c_obj, sk_SkIRect c_r);
  bool misk_IRect_intersect (sk_SkIRect *c_obj, sk_SkIRect c_r);
  void misk_IRect_join (sk_SkIRect *c_obj, sk_SkIRect c_r);
  void misk_IRect_sort (sk_SkIRect *c_obj);

  sk_SkMatrix *misk_new_Matrix ();

  sk_SkPaint *misk_new_Paint ();
  sk_SkPaint *misk_new_PaintCopy (sk_SkPaint *c_paint);
  void misk_delete_SkPaint (sk_SkPaint *obj);
  void misk_Paint_reset (sk_SkPaint *c_obj);
  unsigned char misk_Paint_getAlpha (sk_SkPaint *c_obj);
  void misk_Paint_setAlpha (sk_SkPaint *c_obj, unsigned int c_a);
  void misk_Paint_setARGB (sk_SkPaint *c_obj, unsigned int c_a,
                           unsigned int c_r, unsigned int c_g,
                           unsigned int c_b);
  void misk_Paint_setAntiAlias (sk_SkPaint *c_obj, bool c_aa);
  void misk_Paint_setBlendMode (sk_SkPaint *c_obj, int c_mode);
  void misk_Paint_setColor (sk_SkPaint *c_obj, uint c_color);
  void misk_Paint_setDither (sk_SkPaint *c_obj, bool c_dither);
  uint misk_Paint_getStrokeCap (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeCap (sk_SkPaint *c_obj, uint c_cap);
  uint misk_Paint_getStrokeJoin (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeJoin (sk_SkPaint *c_obj, uint c_join);
  float misk_Paint_getStrokeMiter (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeMiter (sk_SkPaint *c_obj, float c_miter);
  float misk_Paint_getStrokeWidth (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeWidth (sk_SkPaint *c_obj, float c_width);
  uint misk_Paint_getStyle (sk_SkPaint *c_obj);
  void misk_Paint_setStyle (sk_SkPaint *c_obj, uint c_style);

  sk_SkPath *misk_new_Path ();
  sk_SkPath *misk_new_PathCopy (sk_SkPath *c_path);
  void misk_delete_SkPath (sk_SkPath *obj);

  sk_SkPixmap *misk_new_Pixmap ();

  sk_SkRect misk_Rect_MakeEmpty ();
  sk_SkRect misk_Rect_MakeWH (float c_w, float c_h);
  sk_SkRect misk_Rect_MakeSize (sk_SkSize c_size);
  sk_SkRect misk_Rect_MakeLTRB (float c_l, float c_t, float c_r, float c_b);
  sk_SkRect misk_Rect_MakeXYWH (float c_x, float c_y, float c_w, float c_h);
  sk_SkRect misk_Rect_MakeISize (sk_SkISize c_size);
  sk_SkRect misk_Rect_MakeIRect (sk_SkIRect c_irect);
  bool misk_Rect_Intersects (sk_SkRect c_a, sk_SkRect c_b);
  float misk_Rect_x (sk_SkRect *c_obj);
  float misk_Rect_y (sk_SkRect *c_obj);
  float misk_Rect_width (sk_SkRect *c_obj);
  float misk_Rect_height (sk_SkRect *c_obj);
  float misk_Rect_centerX (sk_SkRect *c_obj);
  float misk_Rect_centerY (sk_SkRect *c_obj);
  bool misk_Rect_isEmpty (sk_SkRect *c_obj);
  void misk_Rect_setEmpty (sk_SkRect *c_obj);
  void misk_Rect_setLTRB (sk_SkRect *c_obj, float c_left, float c_top,
                          float c_right, float c_bottom);
  void misk_Rect_setXYWH (sk_SkRect *c_obj, float c_x, float c_y,
                          float c_width, float c_height);
  void misk_Rect_setWH (sk_SkRect *c_obj, float c_width, float c_height);
  void misk_Rect_offset (sk_SkRect *c_obj, float c_dx, float c_dy);
  void misk_Rect_offsetTo (sk_SkRect *c_obj, float c_newX, float c_newY);
  void misk_Rect_inset (sk_SkRect *c_obj, float c_dx, float c_dy);
  void misk_Rect_outset (sk_SkRect *c_obj, float c_dx, float c_dy);
  bool misk_Rect_contains (sk_SkRect *c_obj, float c_x, float c_y);
  bool misk_Rect_containsRect (sk_SkRect *c_obj, sk_SkRect c_r);
  bool misk_Rect_intersect (sk_SkRect *c_obj, sk_SkRect c_r);
  void misk_Rect_join (sk_SkRect *c_obj, sk_SkRect c_r);
  void misk_Rect_sort (sk_SkRect *c_obj);

  sk_SkRRect *misk_new_RRect ();
  sk_SkRRect *misk_new_RRectCopy (sk_SkRRect c_rrect);

  sk_SkRegion *misk_new_Region ();
  sk_SkRegion *misk_new_RegionCopy (sk_SkRegion *c_region);
  sk_SkRegion *misk_new_RegionCopyRect (sk_SkIRect c_rect);
  void misk_delete_SkRegion (sk_SkRegion *obj);

  sk_SkSamplingOptions *misk_new_SamplingOptions ();
  sk_SkSamplingOptions *
  misk_new_SamplingOptionsCopy (sk_SkSamplingOptions *c_p0);

  sk_SkCanvas *misk_Surface_getCanvas (sk_SkSurface *c_obj);

  sk_SkSurfaceProps *misk_new_SurfaceProps ();
  sk_SkSurfaceProps *misk_new_SurfacePropsPixelGeometry (unsigned int c_flags,
                                                         uint c_p1);
  sk_SkSurfaceProps *misk_new_SurfacePropsCopy (sk_SkSurfaceProps *c_p0);

  bool misk_Typeface_Equal (sk_SkTypeface *c_facea, sk_SkTypeface *c_faceb);
  sk_SkTypeface *misk_Typeface_MakeEmpty ();

  sk_GrGLInterface *misk_GrGLMakeNativeInterface ();
  sk_GrBackendRenderTarget
  misk_GrBackendRenderTargetsMakeGL (int c_width, int c_height,
                                     int c_sampleCnt, int c_stencilBits,
                                     sk_GrGLFramebufferInfo c_glInfo);
  sk_SkSurface *misk_SkSurfacesWrapBackendRenderTarget (
      sk_GrRecordingContext *c_context,
      sk_GrBackendRenderTarget *c_backendRenderTarget, int c_origin,
      int c_colorType, sk_SkColorSpace *c_colorSpace,
      sk_SkSurfaceProps *c_surfaceProps);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterfaceOptions (sk_GrGLInterface *c_p0,
                                               sk_GrContextOptions c_p1);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterface (sk_GrGLInterface *c_p0);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLOptions (sk_GrContextOptions c_p0);
  sk_GrDirectContext *misk_GrDirectContextsMakeGL ();
  uint misk_SkColorSetARGB (unsigned int c_a, unsigned int c_r,
                            unsigned int c_g, unsigned int c_b);
  uint misk_SkColorSetA (uint c_c, unsigned int c_a);
  uint misk_SkPreMultiplyARGB (unsigned int c_a, unsigned int c_r,
                               unsigned int c_g, unsigned int c_b);
  uint misk_SkPreMultiplyColor (uint c_c);
  sk_SkFontMgr *sk_fontmgr_ref_default (void);

#ifdef __cplusplus
}
#endif // __cplusplus
