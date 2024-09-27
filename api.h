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
    uint fFBOID;
    uint fFormat;
    uchar fProtected[1];
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
    ulong fNumSemaphores;
    uchar fSignalSemaphores[8];
    uchar fFinishedProc[8];
    uchar fFinishedContext[8];
    uchar fSubmittedProc[8];
    uchar fSubmittedContext[8];
  } sk_GrFlushInfo;

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
    uchar fStandard[4];
    uchar pad_3[4];
    uchar fExtensions[24];
    uchar fFunctions[8240];
  } sk_GrGLInterface;

  typedef struct
  {
    bool fSuppressPrints;
    uchar pad_1[3];
    uchar fSkipGLErrorChecks[4];
    int fMaxTextureSizeOverride;
    int fBufferMapThreshold;
    ulong fMinimumStagingBufferSize;
    uchar fExecutor[8];
    bool fDoManualMipmapping;
    bool fDisableCoverageCountingPaths;
    bool fDisableDistanceFieldPaths;
    bool fAllowPathMaskCaching;
    bool fDisableGpuYUVConversion;
    uchar pad_11[3];
    ulong fGlyphCacheTextureMaximumBytes;
    float fMinDistanceFieldFontSize;
    float fGlyphsAsPathsFontSize;
    uchar fAllowMultipleGlyphCacheTextures[4];
    bool fAvoidStencilBuffers;
    uchar pad_16[3];
    uchar fUseDrawInsteadOfClear[4];
    uchar fReduceOpsTaskSplitting[4];
    bool fPreferExternalImagesOverES3;
    bool fDisableDriverCorrectnessWorkarounds;
    uchar pad_20[2];
    int fRuntimeProgramCacheSize;
    uchar fPersistentCache[8];
    uchar fShaderCacheStrategy[4];
    uchar pad_23[4];
    uchar fShaderErrorHandler[8];
    int fInternalMultisampleCount;
    int fMaxCachedVulkanSecondaryCommandBuffers;
    bool fSuppressMipmapSupport;
    bool fDisableTessellationPathRenderer;
    bool fEnableExperimentalHardwareTessellation;
    bool fSupportBilerpFromGlyphAtlas;
    bool fReducedShaderVariations;
    bool fAllowMSAAOnNewIntel;
    bool fAlwaysUseTexStorageWhenAvailable;
    uchar pad_33[1];
    uchar fContextDeleteContext[8];
    uchar fContextDeleteProc[8];
    uchar fDriverBugWorkarounds[17];
    uchar pad_36[7];
  } sk_GrContextOptions;

  typedef struct
  {
    uchar fOval[16];
    float fStartAngle;
    float fSweepAngle;
    uchar fType[1];
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
    uint fTransferFnHash;
    uint fToXYZD50Hash;
    uchar fTransferFn[28];
    uchar fToXYZD50[36];
    uchar fInvTransferFn[28];
    uchar fFromXYZD50[36];
    uchar fLazyDstFieldsOnce[1];
    uchar pad_7[3];
  } sk_SkColorSpace;

  typedef struct
  {
    uchar pad_0[8];
    uchar fReleaseProc[8];
    uchar fReleaseProcContext[8];
    uchar fPtr[8];
    ulong fSize;
  } sk_SkData;

  typedef struct
  {
    uchar fTypeface[8];
    float fSize;
    float fScaleX;
    float fSkewX;
    uchar fFlags;
    uchar fEdging;
    uchar fHinting;
    uchar pad_7[1];
  } sk_SkFont;

  typedef struct
  {
    int fCollectionIndex;
    uchar pad_1[4];
    uchar fVariationDesignPosition[16];
    uchar fPalette[24];
  } sk_SkFontArguments;

  typedef struct
  {
    uchar coordinates[8];
    int coordinateCount;
    uchar pad_2[4];
  } sk_SkFontArgumentsVariationPosition;

  typedef struct
  {
    int index;
    uchar pad_1[4];
    uchar overrides[8];
    int overrideCount;
    uchar pad_3[4];
  } sk_SkFontArgumentsPalette;

  typedef struct
  {
    uint fFlags;
    float fTop;
    float fAscent;
    float fDescent;
    float fBottom;
    float fLeading;
    float fAvgCharWidth;
    float fMaxCharWidth;
    float fXMin;
    float fXMax;
    float fXHeight;
    float fCapHeight;
    float fUnderlineThickness;
    float fUnderlinePosition;
    float fStrikeoutThickness;
    float fStrikeoutPosition;
  } sk_SkFontMetrics;

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
  } sk_SkFontStyleSet;

  typedef struct
  {
    uchar pad_0[16];
    uchar fInfo[24];
    uint fUniqueID;
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
    uchar pad_0[8];
    uchar fData[8];
    ulong fOffset;
  } sk_SkMemoryStream;

  typedef struct
  {
    int fX;
    int fY;
  } sk_SkIPoint;

  typedef struct
  {
    int fLeft;
    int fTop;
    int fRight;
    int fBottom;
  } sk_SkIRect;

  typedef struct
  {
    int fWidth;
    int fHeight;
  } sk_SkISize;

  typedef struct
  {
    uchar fMat[36];
    int fTypeMask;
  } sk_SkMatrix;

  typedef struct
  {
    uchar fPathRefs[16];
    uchar fOps[24];
  } sk_SkOpBuilder;

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
    uchar fFillType;
    // TODO misalignment (perhaps there are bitfields around here?)
    uchar fIsVolatile;
  } sk_SkPath;

  typedef struct
  {
    float fX;
    float fY;
  } sk_SkPoint;

  typedef struct
  {
    uchar fPixels[8];
    ulong fRowBytes;
    uchar fInfo[24];
  } sk_SkPixmap;

  typedef struct
  {
    float fLeft;
    float fTop;
    float fRight;
    float fBottom;
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
    float fR;
    // TODO misalignment (perhaps there are bitfields around here?)
    float fG;
    // TODO misalignment (perhaps there are bitfields around here?)
    float fB;
    // TODO misalignment (perhaps there are bitfields around here?)
    float fA;
  } sk_SkRGBA4f;

  typedef struct
  {
    int maxAniso;
    bool useCubic;
    uchar pad_2[3];
    uchar cubic[8];
    uchar filter[4];
    uchar mipmap[4];
  } sk_SkSamplingOptions;

  typedef struct
  {
    float fWidth;
    float fHeight;
  } sk_SkSize;

  typedef struct
  {
    uchar pad_0[8];
  } sk_SkStream;

  typedef struct
  {
    uchar fRec[8];
  } sk_SkString;

  typedef struct
  {
    uchar pad_0[12];
    uchar fProps[16];
    int fWidth;
    int fHeight;
    uint fGenerationID;
  } sk_SkSurface;

  typedef struct
  {
    uint fFlags;
    uchar fPixelGeometry[4];
    float fTextContrast;
    float fTextGamma;
  } sk_SkSurfaceProps;

  typedef struct
  {
    uchar pad_0[16];
    uchar fRoot[8];
    uchar fFontMgr[8];
    uchar fTextShapingFactory[8];
    uchar fResourceProvider[8];
    uchar fIDMapper[16];
    uchar fContainerSize[8];
  } sk_SkSVGDOM;

  typedef struct
  {
    uchar pad_0[776];
    uchar fX[8];
    uchar fY[8];
    uchar fWidth[8];
    uchar fHeight[8];
    uchar fPreserveAspectRatio[8];
    uchar fViewBox[20];
    uchar fType[4];
  } sk_SkSVGSVG;

  typedef struct
  {
    uchar fViewport[8];
    float fDPI;
  } sk_SkSVGLengthContext;

  typedef struct
  {
    uchar pad_0[4];
    uchar fBounds[16];
    uint fUniqueID;
    uchar fCacheID[4];
    uchar pad_3[4];
    uchar fPurgeDelegate[8];
    ulong fStorageSize;
  } sk_SkTextBlob;

  typedef struct
  {
    uchar fStorage[8];
    ulong fStorageSize;
    ulong fStorageUsed;
    uchar fBounds[16];
    int fRunCount;
    bool fDeferredBounds;
    uchar pad_6[3];
    ulong fLastRun;
    uchar fCurrentRunBuffer[32];
  } sk_SkTextBlobBuilder;

  typedef struct
  {
    uchar glyphs[8];
    uchar pos[8];
    uchar utf8text[8];
    uchar clusters[8];
  } sk_SkTextBlobBuilderRunBuffer;

  typedef struct
  {
    uchar pad_0[16];
    uint fUniqueID;
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
  sk_SkISize
  misk_GrBackendRenderTarget_dimensions (sk_GrBackendRenderTarget *c_obj);
  int misk_GrBackendRenderTarget_width (sk_GrBackendRenderTarget *c_obj);
  int misk_GrBackendRenderTarget_height (sk_GrBackendRenderTarget *c_obj);
  int misk_GrBackendRenderTarget_sampleCnt (sk_GrBackendRenderTarget *c_obj);
  int misk_GrBackendRenderTarget_stencilBits (sk_GrBackendRenderTarget *c_obj);
  bool misk_GrBackendRenderTarget_isFramebufferOnly (
      sk_GrBackendRenderTarget *c_obj);

  void misk_delete_GrDirectContext (sk_GrDirectContext *obj);
  void misk_GrDirectContext_flushAndSubmit (sk_GrDirectContext *c_obj,
                                            bool c_sync);
  bool misk_GrDirectContext_flush (sk_GrDirectContext *c_obj,
                                   sk_GrFlushInfo *c_info);
  bool misk_GrDirectContext_submit (sk_GrDirectContext *c_obj, bool c_sync);

  void misk_delete_GrRecordingContext (sk_GrRecordingContext *obj);

  void misk_unref_GrGLInterface (sk_GrGLInterface *c_obj);

  sk_GrContextOptions *misk_new_GrContextOptions ();

  sk_SkArc *misk_new_Arc ();
  sk_SkArc *misk_new_ArcCopy (sk_SkArc *c_arc);

  sk_SkBitmap *misk_new_Bitmap ();
  sk_SkBitmap *misk_new_BitmapCopy (sk_SkBitmap *c_src);
  void misk_delete_SkBitmap (sk_SkBitmap *obj);
  bool misk_Bitmap_ComputeIsOpaque (sk_SkBitmap *c_bm);
  int misk_Bitmap_bytesPerPixel (sk_SkBitmap *c_obj);
  int misk_Bitmap_rowBytesAsPixels (sk_SkBitmap *c_obj);
  int misk_Bitmap_shiftPerPixel (sk_SkBitmap *c_obj);
  bool misk_Bitmap_drawsNothing (sk_SkBitmap *c_obj);
  unsigned long misk_Bitmap_rowBytes (sk_SkBitmap *c_obj);
  bool misk_Bitmap_isImmutable (sk_SkBitmap *c_obj);
  void misk_Bitmap_setImmutable (sk_SkBitmap *c_obj);
  sk_SkIRect misk_Bitmap_bounds (sk_SkBitmap *c_obj);
  sk_SkISize misk_Bitmap_dimensions (sk_SkBitmap *c_obj);

  sk_SkCanvas *misk_new_Canvas ();
  sk_SkCanvas *misk_new_CanvasWithDimensions (int c_width, int c_height,
                                              sk_SkSurfaceProps *c_props);
  sk_SkCanvas *misk_new_CanvasFromBitmap (sk_SkBitmap *c_bitmap);
  sk_SkCanvas *
  misk_new_CanvasFromBitmapSurfaceProps (sk_SkBitmap *c_bitmap,
                                         sk_SkSurfaceProps *c_props);
  void misk_delete_SkCanvas (sk_SkCanvas *obj);
  bool misk_Canvas_getProps (sk_SkCanvas *c_obj, sk_SkSurfaceProps *c_props);
  sk_SkSurfaceProps misk_Canvas_getBaseProps (sk_SkCanvas *c_obj);
  sk_SkSurfaceProps misk_Canvas_getTopProps (sk_SkCanvas *c_obj);
  sk_SkISize misk_Canvas_getBaseLayerSize (sk_SkCanvas *c_obj);
  sk_SkSurface *misk_Canvas_makeSurface (sk_SkCanvas *c_obj,
                                         sk_SkImageInfo *c_info,
                                         sk_SkSurfaceProps *c_props);
  sk_SkSurface *misk_Canvas_getSurface (sk_SkCanvas *c_obj);
  sk_GrRecordingContext *misk_Canvas_recordingContext (sk_SkCanvas *c_obj);
  bool misk_Canvas_peekPixels (sk_SkCanvas *c_obj, sk_SkPixmap *c_pixmap);
  bool misk_Canvas_readPixelsImageInfo (sk_SkCanvas *c_obj,
                                        sk_SkImageInfo *c_dstInfo,
                                        void *c_dstPixels, ulong c_dstRowBytes,
                                        int c_srcX, int c_srcY);
  bool misk_Canvas_readPixelsPixmap (sk_SkCanvas *c_obj, sk_SkPixmap *c_pixmap,
                                     int c_srcX, int c_srcY);
  bool misk_Canvas_readPixelsBitmap (sk_SkCanvas *c_obj, sk_SkBitmap *c_bitmap,
                                     int c_srcX, int c_srcY);
  bool misk_Canvas_writePixelsImageInfo (sk_SkCanvas *c_obj,
                                         sk_SkImageInfo *c_info,
                                         void *c_pixels, ulong c_rowBytes,
                                         int c_x, int c_y);
  bool misk_Canvas_writePixelsBitmap (sk_SkCanvas *c_obj,
                                      sk_SkBitmap *c_bitmap, int c_x, int c_y);
  int misk_Canvas_save (sk_SkCanvas *c_obj);
  int misk_Canvas_saveLayer (sk_SkCanvas *c_obj, sk_SkRect c_bounds,
                             sk_SkPaint *c_paint);
  int misk_Canvas_saveLayerAlpha (sk_SkCanvas *c_obj, sk_SkRect c_bounds,
                                  uint c_alpha);
  void misk_Canvas_restore (sk_SkCanvas *c_obj);
  int misk_Canvas_getSaveCount (sk_SkCanvas *c_obj);
  void misk_Canvas_restoreToCount (sk_SkCanvas *c_obj, int c_saveCount);
  void misk_Canvas_translate (sk_SkCanvas *c_obj, float c_dx, float c_dy);
  void misk_Canvas_scale (sk_SkCanvas *c_obj, float c_sx, float c_sy);
  void misk_Canvas_rotate (sk_SkCanvas *c_obj, float c_degrees);
  void misk_Canvas_rotateAboutPoint (sk_SkCanvas *c_obj, float c_degrees,
                                     float c_px, float c_py);
  void misk_Canvas_skew (sk_SkCanvas *c_obj, float c_sx, float c_sy);
  void misk_Canvas_concatMatrix (sk_SkCanvas *c_obj, sk_SkMatrix *c_matrix);
  void misk_Canvas_concatM44 (sk_SkCanvas *c_obj, sk_SkM44 *c_p0);
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
                                           sk_SkRect *c_bounds);
  sk_SkIRect misk_Canvas_getDeviceClipBounds (sk_SkCanvas *c_obj);
  bool misk_Canvas_getDeviceClipBoundsRect (sk_SkCanvas *c_obj,
                                            sk_SkIRect *c_bounds);
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
  void misk_Canvas_drawString (sk_SkCanvas *c_obj, char *c_str, float c_x,
                               float c_y, sk_SkFont *c_font,
                               sk_SkPaint *c_paint);
  void misk_Canvas_drawGlyphs (sk_SkCanvas *c_obj, int c_count,
                               ushort *c_glyphs, sk_SkPoint *c_positions,
                               uint *c_clusters, int c_textByteCount,
                               char *c_utf8text, sk_SkPoint c_origin,
                               sk_SkFont *c_font, sk_SkPaint *c_paint);
  void misk_Canvas_drawTextBlob (sk_SkCanvas *c_obj, sk_SkTextBlob *c_blob,
                                 float c_x, float c_y, sk_SkPaint *c_paint);

  sk_SkColorSpace *misk_ColorSpace_MakeSRGB ();
  sk_SkColorSpace *misk_ColorSpace_MakeSRGBLinear ();
  bool misk_ColorSpace_Equals (sk_SkColorSpace *c_p0, sk_SkColorSpace *c_p1);
  void misk_unref_SkColorSpace (sk_SkColorSpace *c_obj);

  unsigned long misk_Data_size (sk_SkData *c_obj);
  bool misk_Data_isEmpty (sk_SkData *c_obj);
  sk_SkData *misk_Data_MakeWithCopy (void *c_data, ulong c_length);
  sk_SkData *misk_Data_MakeZeroInitialized (ulong c_length);
  sk_SkData *misk_Data_MakeWithCString (char *c_cstr);
  sk_SkData *misk_Data_MakeWithoutCopy (void *c_data, ulong c_length);
  sk_SkData *misk_Data_MakeFromFileName (char *c_path);
  void misk_unref_SkData (sk_SkData *c_obj);

  sk_SkFont *misk_new_Font ();
  sk_SkFont *misk_new_FontTypefaceSize (sk_SkTypeface *c_typeface,
                                        float c_size);
  sk_SkFont *misk_new_FontTypeface (sk_SkTypeface *c_typeface);
  sk_SkFont *misk_new_FontTypefaceSizeScaleSkew (sk_SkTypeface *c_typeface,
                                                 float c_size, float c_scaleX,
                                                 float c_skewX);
  float misk_Font_getMetrics (sk_SkFont *c_obj, sk_SkFontMetrics *c_metrics);
  void misk_Font_getXPos (sk_SkFont *c_obj, ushort *c_glyphs, int c_count,
                          float *c_xpos, float c_origin);
  float misk_Font_measureText (sk_SkFont *c_obj, void *c_text,
                               ulong c_byteLength, int c_encoding,
                               sk_SkRect *c_bounds);
  float misk_Font_measureTextPaint (sk_SkFont *c_obj, void *c_text,
                                    ulong c_byteLength, int c_encoding,
                                    sk_SkRect *c_bounds, sk_SkPaint *c_paint);
  void misk_Font_setForceAutoHinting (sk_SkFont *c_obj,
                                      bool c_forceAutoHinting);
  void misk_Font_setHinting (sk_SkFont *c_obj, int c_hintingLevel);
  void misk_Font_setSubpixel (sk_SkFont *c_obj, bool c_subpixel);
  int misk_Font_textToGlyphs (sk_SkFont *c_obj, void *c_text,
                              ulong c_byteLength, int c_encoding,
                              ushort *c_glyphs, int c_maxGlyphCount);
  unsigned short misk_Font_unicharToGlyph (sk_SkFont *c_obj, int c_uni);
  void misk_Font_unicharsToGlyphs (sk_SkFont *c_obj, int *c_uni, int c_count,
                                   ushort *c_glyphs);
  void misk_Font_getWidthsBounds (sk_SkFont *c_obj, ushort *c_glyphs,
                                  int c_count, float *c_widths,
                                  sk_SkRect *c_bounds);
  void misk_Font_getWidths (sk_SkFont *c_obj, ushort *c_glyphs, int c_count,
                            float *c_widths);

  sk_SkFontArguments *misk_new_FontArguments ();
  sk_SkFontArguments
  misk_FontArguments_setCollectionIndex (sk_SkFontArguments *c_obj,
                                         int c_collectionIndex);
  int misk_FontArguments_getCollectionIndex (sk_SkFontArguments *c_obj);
  sk_SkFontArgumentsVariationPosition
  misk_FontArguments_getVariationDesignPosition (sk_SkFontArguments *c_obj);
  sk_SkFontArgumentsPalette
  misk_FontArguments_getPalette (sk_SkFontArguments *c_obj);

  sk_SkFontStyleSet *misk_FontMgr_matchFamily (sk_SkFontMgr *c_obj,
                                               char *c_familyName);
  sk_SkTypeface *misk_FontMgr_matchFamilyStyle (sk_SkFontMgr *c_obj,
                                                char *c_familyName,
                                                sk_SkFontStyle *c_p1);
  sk_SkTypeface *misk_FontMgr_makeFromData (sk_SkFontMgr *c_obj,
                                            sk_SkData *c_p0, int c_ttcIndex);
  sk_SkTypeface *misk_FontMgr_makeFromFile (sk_SkFontMgr *c_obj, char *c_path,
                                            int c_ttcIndex);
  void misk_unref_SkFontMgr (sk_SkFontMgr *c_obj);

  sk_SkFontStyle *misk_new_FontStyle2 (int c_weight, int c_width,
                                       uint c_slant);
  sk_SkFontStyle *misk_new_FontStyle ();
  sk_SkFontStyle misk_FontStyle_Normal ();
  sk_SkFontStyle misk_FontStyle_Bold ();
  sk_SkFontStyle misk_FontStyle_Italic ();
  sk_SkFontStyle misk_FontStyle_BoldItalic ();

  int misk_FontStyleSet_count (sk_SkFontStyleSet *c_obj);
  void misk_FontStyleSet_getStyle (sk_SkFontStyleSet *c_obj, int c_index,
                                   sk_SkFontStyle *c_p1, sk_SkString *c_style);
  sk_SkTypeface *misk_FontStyleSet_createTypeface (sk_SkFontStyleSet *c_obj,
                                                   int c_index);
  sk_SkTypeface *misk_FontStyleSet_matchStyle (sk_SkFontStyleSet *c_obj,
                                               sk_SkFontStyle *c_pattern);
  void misk_unref_SkFontStyleSet (sk_SkFontStyleSet *c_obj);

  sk_SkImageInfo misk_Image_imageInfo (sk_SkImage *c_obj);
  int misk_Image_width (sk_SkImage *c_obj);
  int misk_Image_height (sk_SkImage *c_obj);
  sk_SkISize misk_Image_dimensions (sk_SkImage *c_obj);
  sk_SkIRect misk_Image_bounds (sk_SkImage *c_obj);
  unsigned int misk_Image_uniqueID (sk_SkImage *c_obj);
  int misk_Image_alphaType (sk_SkImage *c_obj);
  int misk_Image_colorType (sk_SkImage *c_obj);
  sk_SkColorSpace *misk_Image_colorSpace (sk_SkImage *c_obj);
  bool misk_Image_isAlphaOnly (sk_SkImage *c_obj);
  bool misk_Image_isOpaque (sk_SkImage *c_obj);
  bool misk_Image_readPixels (sk_SkImage *c_obj, sk_GrDirectContext *c_context,
                              sk_SkImageInfo *c_dstInfo, void *c_dstPixels,
                              ulong c_dstRowBytes, int c_srcX, int c_srcY,
                              uint c_cachingHint);
  sk_SkImage *misk_Image_makeSubset (sk_SkImage *c_obj,
                                     sk_GrDirectContext *c_direct,
                                     sk_SkIRect c_subset);
  void misk_unref_SkImage (sk_SkImage *c_obj);

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

  sk_SkMemoryStream *misk_MemoryStream_Make (sk_SkData *c_data);
  sk_SkMemoryStream *misk_MemoryStream_MakeDirect (void *c_data,
                                                   ulong c_length);

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

  sk_SkOpBuilder *misk_new_OpBuilder ();
  void misk_OpBuilder_add (sk_SkOpBuilder *c_obj, sk_SkPath *c_path,
                           uint c__operator);
  bool misk_OpBuilder_resolve (sk_SkOpBuilder *c_obj, sk_SkPath *c_result);

  sk_SkPaint *misk_new_Paint ();
  sk_SkPaint *misk_new_PaintCopy (sk_SkPaint *c_paint);
  void misk_delete_SkPaint (sk_SkPaint *obj);
  void misk_Paint_reset (sk_SkPaint *c_obj);
  unsigned char misk_Paint_getAlpha (sk_SkPaint *c_obj);
  void misk_Paint_setAlpha (sk_SkPaint *c_obj, uint c_a);
  void misk_Paint_setARGB (sk_SkPaint *c_obj, uint c_a, uint c_r, uint c_g,
                           uint c_b);
  void misk_Paint_setAntiAlias (sk_SkPaint *c_obj, bool c_aa);
  void misk_Paint_setBlendMode (sk_SkPaint *c_obj, int c_mode);
  uint misk_Paint_getColor (sk_SkPaint *c_obj);
  void misk_Paint_setColor (sk_SkPaint *c_obj, uint c_color);
  void misk_Paint_setDither (sk_SkPaint *c_obj, bool c_dither);
  uint misk_Paint_getStrokeCap (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeCap (sk_SkPaint *c_obj, uint c_cap);
  uchar misk_Paint_getStrokeJoin (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeJoin (sk_SkPaint *c_obj, uchar c_join);
  float misk_Paint_getStrokeMiter (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeMiter (sk_SkPaint *c_obj, float c_miter);
  float misk_Paint_getStrokeWidth (sk_SkPaint *c_obj);
  void misk_Paint_setStrokeWidth (sk_SkPaint *c_obj, float c_width);
  uchar misk_Paint_getStyle (sk_SkPaint *c_obj);
  void misk_Paint_setStyle (sk_SkPaint *c_obj, uchar c_style);

  sk_SkPath *misk_new_Path ();
  sk_SkPath *misk_new_PathCopy (sk_SkPath *c_path);
  void misk_delete_SkPath (sk_SkPath *obj);
  bool misk_Path_isInterpolatable (sk_SkPath *c_obj, sk_SkPath *c_compare);
  bool misk_Path_interpolate (sk_SkPath *c_obj, sk_SkPath *c_ending,
                              float c_weight, sk_SkPath *c_out);
  int misk_Path_getFillType (sk_SkPath *c_obj);
  void misk_Path_setFillType (sk_SkPath *c_obj, int c_ft);
  bool misk_Path_isInverseFillType (sk_SkPath *c_obj);
  void misk_Path_toggleInverseFillType (sk_SkPath *c_obj);
  bool misk_Path_isConvex (sk_SkPath *c_obj);
  bool misk_Path_isOval (sk_SkPath *c_obj, sk_SkRect c_bounds);
  bool misk_Path_isRRect (sk_SkPath *c_obj, sk_SkRRect c_rrect);
  sk_SkPath misk_Path_reset (sk_SkPath *c_obj);
  sk_SkPath misk_Path_rewind (sk_SkPath *c_obj);
  bool misk_Path_isEmpty (sk_SkPath *c_obj);
  bool misk_Path_isLastContourClosed (sk_SkPath *c_obj);
  bool misk_Path_isFinite (sk_SkPath *c_obj);
  bool misk_Path_isVolatile (sk_SkPath *c_obj);
  sk_SkPath misk_Path_setIsVolatile (sk_SkPath *c_obj, bool c_isVolatile);
  int misk_Path_countPoints (sk_SkPath *c_obj);
  sk_SkPoint misk_Path_getPoint (sk_SkPath *c_obj, int c_index);
  int misk_Path_getPoints (sk_SkPath *c_obj, sk_SkPoint *c_points, int c_max);
  int misk_Path_countVerbs (sk_SkPath *c_obj);
  int misk_Path_getVerbs (sk_SkPath *c_obj, uchar *c_verbs, int c_max);
  unsigned long misk_Path_approximateBytesUsed (sk_SkPath *c_obj);
  void misk_Path_swap (sk_SkPath *c_obj, sk_SkPath *c_other);
  sk_SkRect misk_Path_getBounds (sk_SkPath *c_obj);
  void misk_Path_updateBoundsCache (sk_SkPath *c_obj);
  sk_SkRect misk_Path_computeTightBounds (sk_SkPath *c_obj);
  bool misk_Path_conservativelyContainsRect (sk_SkPath *c_obj,
                                             sk_SkRect c_rect);
  void misk_Path_incReserve (sk_SkPath *c_obj, int c_extraPtCount,
                             int c_extraVerbCount, int c_extraConicCount);
  sk_SkPath misk_Path_moveToPoint (sk_SkPath *c_obj, float c_x, float c_y);
  sk_SkPath misk_Path_moveTo (sk_SkPath *c_obj, sk_SkPoint *c_p);
  sk_SkPath misk_Path_rMoveTo (sk_SkPath *c_obj, float c_dx, float c_dy);
  sk_SkPath misk_Path_lineTo (sk_SkPath *c_obj, float c_x, float c_y);
  sk_SkPath misk_Path_lineToPoint (sk_SkPath *c_obj, sk_SkPoint *c_p);
  sk_SkPath misk_Path_rLineTo (sk_SkPath *c_obj, float c_dx, float c_dy);
  sk_SkPath misk_Path_quadTo (sk_SkPath *c_obj, float c_x1, float c_y1,
                              float c_x2, float c_y2);
  sk_SkPath misk_Path_quadToPoint (sk_SkPath *c_obj, sk_SkPoint *c_p1,
                                   sk_SkPoint *c_p2);
  sk_SkPath misk_Path_rQuadTo (sk_SkPath *c_obj, float c_dx1, float c_dy1,
                               float c_dx2, float c_dy2);
  sk_SkPath misk_Path_conicTo (sk_SkPath *c_obj, float c_x1, float c_y1,
                               float c_x2, float c_y2, float c_w);
  sk_SkPath misk_Path_conicToPoints (sk_SkPath *c_obj, sk_SkPoint *c_p1,
                                     sk_SkPoint *c_p2, float c_w);
  sk_SkPath misk_Path_rConicTo (sk_SkPath *c_obj, float c_dx1, float c_dy1,
                                float c_dx2, float c_dy2, float c_w);
  sk_SkPath misk_Path_cubicTo (sk_SkPath *c_obj, float c_x1, float c_y1,
                               float c_x2, float c_y2, float c_x3, float c_y3);
  sk_SkPath misk_Path_cubicToPoints (sk_SkPath *c_obj, sk_SkPoint *c_p1,
                                     sk_SkPoint *c_p2, sk_SkPoint *c_p3);
  sk_SkPath misk_Path_rCubicTo (sk_SkPath *c_obj, float c_dx1, float c_dy1,
                                float c_dx2, float c_dy2, float c_dx3,
                                float c_dy3);
  sk_SkPath misk_Path_arcTo1 (sk_SkPath *c_obj, sk_SkRect c_oval,
                              float c_startAngle, float c_sweepAngle,
                              bool c_forceMoveTo);
  sk_SkPath misk_Path_arcTo2 (sk_SkPath *c_obj, float c_x1, float c_y1,
                              float c_x2, float c_y2, float c_radius);
  sk_SkPath misk_Path_arcTo3 (sk_SkPath *c_obj, sk_SkPoint c_p1,
                              sk_SkPoint c_p2, float c_radius);
  sk_SkPath misk_Path_arcTo4 (sk_SkPath *c_obj, float c_rx, float c_ry,
                              float c_xAxisRotate, uint c_largeArc,
                              int c_sweep, float c_x, float c_y);
  sk_SkPath misk_Path_arcTo5 (sk_SkPath *c_obj, sk_SkPoint c_r,
                              float c_xAxisRotate, uint c_largeArc,
                              int c_sweep, sk_SkPoint c_xy);
  sk_SkPath misk_Path_rArcTo (sk_SkPath *c_obj, float c_rx, float c_ry,
                              float c_xAxisRotate, uint c_largeArc,
                              int c_sweep, float c_dx, float c_dy);
  sk_SkPath misk_Path_close (sk_SkPath *c_obj);
  bool misk_Path_isRect (sk_SkPath *c_obj, sk_SkRect *c_rect, bool *c_isClosed,
                         int *c_direction);
  sk_SkPath misk_Path_addRect1 (sk_SkPath *c_obj, sk_SkRect c_rect, int c_dir,
                                uint c_start);
  sk_SkPath misk_Path_addRect2 (sk_SkPath *c_obj, sk_SkRect c_rect, int c_dir);
  sk_SkPath misk_Path_addRect3 (sk_SkPath *c_obj, float c_left, float c_top,
                                float c_right, float c_bottom, int c_dir);
  sk_SkPath misk_Path_addOval1 (sk_SkPath *c_obj, sk_SkRect c_oval, int c_dir);
  sk_SkPath misk_Path_addOval2 (sk_SkPath *c_obj, sk_SkRect c_oval, int c_dir,
                                uint c_start);
  sk_SkPath misk_Path_addCircle (sk_SkPath *c_obj, float c_x, float c_y,
                                 float c_radius, int c_dir);
  sk_SkPath misk_Path_addArc (sk_SkPath *c_obj, sk_SkRect c_oval,
                              float c_startAngle, float c_sweepAngle);
  sk_SkPath misk_Path_addRoundRect1 (sk_SkPath *c_obj, sk_SkRect c_rect,
                                     float c_rx, float c_ry, int c_dir);
  sk_SkPath misk_Path_addRoundRect2 (sk_SkPath *c_obj, sk_SkRect c_rect,
                                     float *c_radii, int c_dir);
  sk_SkPath misk_Path_addRRect1 (sk_SkPath *c_obj, sk_SkRRect c_rrect,
                                 int c_dir);
  sk_SkPath misk_Path_addRRect2 (sk_SkPath *c_obj, sk_SkRRect c_rrect,
                                 int c_dir, uint c_start);
  sk_SkPath misk_Path_addPoly (sk_SkPath *c_obj, sk_SkPoint *c_pts,
                               int c_count, bool c_close);
  sk_SkPath misk_Path_addPath1 (sk_SkPath *c_obj, sk_SkPath *c_src, float c_dx,
                                float c_dy, uint c_mode);
  sk_SkPath misk_Path_addPath2 (sk_SkPath *c_obj, sk_SkPath *c_src,
                                uint c_mode);
  sk_SkPath misk_Path_addPath3 (sk_SkPath *c_obj, sk_SkPath *c_src,
                                sk_SkMatrix *c_matrix, uint c_mode);
  sk_SkPath misk_Path_reverseAddPath (sk_SkPath *c_obj, sk_SkPath *c_src);
  void misk_Path_offset1 (sk_SkPath *c_obj, float c_dx, float c_dy,
                          sk_SkPath *c_dst);
  sk_SkPath misk_Path_offset2 (sk_SkPath *c_obj, float c_dx, float c_dy);
  void misk_Path_transform1 (sk_SkPath *c_obj, sk_SkMatrix *c_matrix,
                             sk_SkPath *c_dst, int c_pc);
  sk_SkPath misk_Path_transform2 (sk_SkPath *c_obj, sk_SkMatrix *c_matrix,
                                  int c_pc);
  sk_SkPath misk_Path_makeTransform (sk_SkPath *c_obj, sk_SkMatrix *c_m,
                                     int c_pc);
  sk_SkPath misk_Path_makeScale (sk_SkPath *c_obj, float c_sx, float c_sy);
  bool misk_Path_getLastPt (sk_SkPath *c_obj, sk_SkPoint *c_lastPt);
  void misk_Path_setLastPt1 (sk_SkPath *c_obj, float c_x, float c_y);
  void misk_Path_setLastPt2 (sk_SkPath *c_obj, sk_SkPoint *c_p);
  unsigned int misk_Path_getSegmentMasks (sk_SkPath *c_obj);
  bool misk_Path_contains (sk_SkPath *c_obj, float c_x, float c_y);
  sk_SkData *misk_Path_serialize (sk_SkPath *c_obj);
  bool misk_Path_isValid (sk_SkPath *c_obj);

  sk_SkPixmap *misk_new_Pixmap ();
  sk_SkPixmap *misk_new_PixmapImageInfo (sk_SkImageInfo *c_info, void *c_addr,
                                         ulong c_rowBytes);

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

  void misk_delete_SkStream (sk_SkStream *obj);

  sk_SkString *misk_new_String (char *c_text);
  void misk_delete_SkString (sk_SkString *obj);
  char *misk_String_data (sk_SkString *c_obj);

  sk_SkCanvas *misk_Surface_getCanvas (sk_SkSurface *c_obj);
  sk_SkSurface *misk_Surface_makeSurface (sk_SkSurface *c_obj, int c_width,
                                          int c_height);
  void misk_unref_SkSurface (sk_SkSurface *c_obj);

  sk_SkSurfaceProps *misk_new_SurfaceProps ();
  sk_SkSurfaceProps *misk_new_SurfacePropsPixelGeometry (uint c_flags,
                                                         uint c_p1);
  sk_SkSurfaceProps *misk_new_SurfacePropsCopy (sk_SkSurfaceProps *c_p0);

  sk_SkSVGSVG *misk_SVGDOM_getRoot (sk_SkSVGDOM *c_obj);
  void misk_SVGDOM_setContainerSize (sk_SkSVGDOM *c_obj, sk_SkSize c_p0);
  sk_SkSize misk_SVGDOM_containerSize (sk_SkSVGDOM *c_obj);
  void misk_SVGDOM_render (sk_SkSVGDOM *c_obj, sk_SkCanvas *c_p0);
  sk_SkSVGDOM *misk_SVGDOM_MakeFromStream (sk_SkStream *c_str);
  void misk_unref_SkSVGDOM (sk_SkSVGDOM *c_obj);

  sk_SkSize misk_SVGSVG_intrinsicSize (sk_SkSVGSVG *c_obj,
                                       sk_SkSVGLengthContext *c_p0);

  sk_SkSVGLengthContext *misk_new_SVGLengthContext (sk_SkSize c_viewport,
                                                    float c_dpi);
  sk_SkSize misk_SVGLengthContext_viewPort (sk_SkSVGLengthContext *c_obj);
  void misk_SVGLengthContext_setViewPort (sk_SkSVGLengthContext *c_obj,
                                          sk_SkSize c_viewport);

  sk_SkRect misk_TextBlob_bounds (sk_SkTextBlob *c_obj);
  unsigned int misk_TextBlob_uniqueID (sk_SkTextBlob *c_obj);
  sk_SkTextBlob *misk_TextBlob_MakeFromString (char *c_string,
                                               sk_SkFont *c_font,
                                               int c_encoding);
  sk_SkTextBlob *misk_TextBlob_MakeFromPosTextH (void *c_text,
                                                 ulong c_byteLength,
                                                 float *c_xpos, float c_constY,
                                                 sk_SkFont *c_font,
                                                 int c_encoding);
  sk_SkTextBlob *misk_TextBlob_MakeFromPosText (void *c_text,
                                                ulong c_byteLength,
                                                sk_SkPoint *c_pos,
                                                sk_SkFont *c_font,
                                                int c_encoding);
  void misk_unref_SkTextBlob (sk_SkTextBlob *c_obj);

  sk_SkTextBlobBuilder *misk_new_TextBlobBuilder ();
  void misk_delete_SkTextBlobBuilder (sk_SkTextBlobBuilder *obj);
  sk_SkTextBlob *misk_TextBlobBuilder_make (sk_SkTextBlobBuilder *c_obj);
  sk_SkTextBlobBuilderRunBuffer
  misk_TextBlobBuilder_allocRun (sk_SkTextBlobBuilder *c_obj,
                                 sk_SkFont *c_font, int c_count, float c_x,
                                 float c_y, sk_SkRect *c_bounds);
  sk_SkTextBlobBuilderRunBuffer
  misk_TextBlobBuilder_allocRunPosH (sk_SkTextBlobBuilder *c_obj,
                                     sk_SkFont *c_font, int c_count, float c_y,
                                     sk_SkRect *c_bounds);
  sk_SkTextBlobBuilderRunBuffer
  misk_TextBlobBuilder_allocRunPos (sk_SkTextBlobBuilder *c_obj,
                                    sk_SkFont *c_font, int c_count,
                                    sk_SkRect *c_bounds);
  sk_SkTextBlobBuilderRunBuffer misk_TextBlobBuilder_allocRunText (
      sk_SkTextBlobBuilder *c_obj, sk_SkFont *c_font, int c_count, float c_x,
      float c_y, int c_textByteCount, sk_SkRect *c_bounds);
  sk_SkTextBlobBuilderRunBuffer misk_TextBlobBuilder_allocRunTextPosH (
      sk_SkTextBlobBuilder *c_obj, sk_SkFont *c_font, int c_count, float c_y,
      int c_textByteCount, sk_SkRect *c_bounds);
  sk_SkTextBlobBuilderRunBuffer misk_TextBlobBuilder_allocRunTextPos (
      sk_SkTextBlobBuilder *c_obj, sk_SkFont *c_font, int c_count,
      int c_textByteCount, sk_SkRect *c_bounds);
  sk_SkPoint *
  misk_TextBlobBuilderRunBuffer_points (sk_SkTextBlobBuilderRunBuffer *c_obj);

  sk_SkFontStyle misk_Typeface_fontStyle (sk_SkTypeface *c_obj);
  bool misk_Typeface_isBold (sk_SkTypeface *c_obj);
  bool misk_Typeface_isItalic (sk_SkTypeface *c_obj);
  bool misk_Typeface_isFixedPitch (sk_SkTypeface *c_obj);
  unsigned int misk_Typeface_uniqueID (sk_SkTypeface *c_obj);
  sk_SkTypeface *misk_Typeface_makeClone (sk_SkTypeface *c_obj,
                                          sk_SkFontArguments *c_p0);
  void misk_Typeface_unicharsToGlyphs (sk_SkTypeface *c_obj, int *c_uni,
                                       int c_count, ushort *c_glyphs);
  int misk_Typeface_textToGlyphs (sk_SkTypeface *c_obj, void *c_text,
                                  ulong c_byteLength, int c_encoding,
                                  ushort *c_glyphs, int c_maxGlyphCount);
  unsigned short misk_Typeface_unicharToGlyph (sk_SkTypeface *c_obj,
                                               int c_unichar);
  int misk_Typeface_countGlyphs (sk_SkTypeface *c_obj);
  int misk_Typeface_countTables (sk_SkTypeface *c_obj);
  int misk_Typeface_getUnitsPerEm (sk_SkTypeface *c_obj);
  void misk_Typeface_getFamilyName (sk_SkTypeface *c_obj, sk_SkString *c_name);
  bool misk_Typeface_Equal (sk_SkTypeface *c_facea, sk_SkTypeface *c_faceb);
  sk_SkTypeface *misk_Typeface_MakeEmpty ();
  void misk_unref_SkTypeface (sk_SkTypeface *c_obj);

  sk_GrBackendRenderTarget
  misk_GrBackendRenderTargetsMakeGL (int c_width, int c_height,
                                     int c_sampleCnt, int c_stencilBits,
                                     sk_GrGLFramebufferInfo c_glInfo);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterfaceOptions (sk_GrGLInterface *c_p0,
                                               sk_GrContextOptions c_p1);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterface (sk_GrGLInterface *c_p0);
  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLOptions (sk_GrContextOptions c_p0);
  sk_GrDirectContext *misk_GrDirectContextsMakeGL ();
  const sk_GrGLInterface *misk_GrGLMakeNativeInterface ();
  uint misk_SkColorSetARGB (uint c_a, uint c_r, uint c_g, uint c_b);
  uint misk_SkColorSetA (uint c_c, uint c_a);
  sk_SkImage *misk_SkImagesDeferredFromEncodedData (sk_SkData *c_encoded,
                                                    int *c_alphaType);
  sk_SkImage *misk_SkImagesRasterFromData (sk_SkImageInfo *c_info,
                                           sk_SkData *c_pixels,
                                           ulong c_rowBytes);
  uint misk_SkPreMultiplyARGB (uint c_a, uint c_r, uint c_g, uint c_b);
  uint misk_SkPreMultiplyColor (uint c_c);
  sk_SkSurface *misk_SkSurfacesWrapBackendRenderTarget (
      sk_GrRecordingContext *c_context,
      sk_GrBackendRenderTarget *c_backendRenderTarget, int c_origin,
      int c_colorType, sk_SkColorSpace *c_colorSpace,
      sk_SkSurfaceProps *c_surfaceProps);
  bool misk_Op (sk_SkPath *c_one, sk_SkPath *c_two, uint c_op,
                sk_SkPath *c_result);
  bool misk_Simplify (sk_SkPath *c_path, sk_SkPath *c_result);
  bool misk_TightBounds (sk_SkPath *c_path, sk_SkRect c_result);
  bool misk_AsWinding (sk_SkPath *c_path, sk_SkPath *c_result);

  extern uchar sk_SK_AlphaOPAQUE;
  extern uchar sk_SK_AlphaTRANSPARENT;
  extern uint sk_SK_ColorBLACK;
  extern uint sk_SK_ColorBLUE;
  extern uint sk_SK_ColorCYAN;
  extern uint sk_SK_ColorDKGRAY;
  extern uint sk_SK_ColorGRAY;
  extern uint sk_SK_ColorGREEN;
  extern uint sk_SK_ColorLTGRAY;
  extern uint sk_SK_ColorMAGENTA;
  extern uint sk_SK_ColorRED;
  extern uint sk_SK_ColorTRANSPARENT;
  extern uint sk_SK_ColorWHITE;
  extern uint sk_SK_ColorYELLOW;
  extern uint sk_SK_InvalidGenID;
  extern uint sk_SK_InvalidUniqueID;
  extern uint sk_SK_MSecMax;
  extern uint sk_kAll_GrBackendState;
  extern int sk_kGrGLStandardCnt;
  extern int sk_kSkColorTypeCnt;
  extern int sk_kSkFilterModeCount;
  extern int sk_kSkMipmapModeCount;
  extern int sk_kSkStrAppendS32_MaxSize;
  extern int sk_kSkStrAppendS64_MaxSize;
  extern int sk_kSkStrAppendScalar_MaxSize;
  extern int sk_kSkStrAppendU32_MaxSize;
  extern int sk_kSkStrAppendU64_MaxSize;

  sk_SkFontMgr *sk_fontmgr_ref_default (void);

#ifdef __cplusplus
}
#endif // __cplusplus
