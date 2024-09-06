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
    uchar pad_0[16];
  } sk_SkFontMgr;

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
    unsigned int fFlags;
    uchar fPixelGeometry[4];
    float fTextContrast;
    float fTextGamma;
  } sk_SkSurfaceProps;

  typedef struct
  {
    int Left;
    int Top;
    int Right;
    int Bottom;
  } sk_SkIRect;

  typedef struct
  {
    float Left;
    float Top;
    float Right;
    float Bottom;
  } sk_SkRect;

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

  sk_SkColorSpace *misk_ColorSpace_MakeSRGB ();
  sk_SkColorSpace *misk_ColorSpace_MakeSRGBLinear ();
  bool misk_ColorSpace_Equals (sk_SkColorSpace *c_p0, sk_SkColorSpace *c_p1);

  sk_SkPaint *misk_new_Paint ();
  sk_SkPaint *misk_new_PaintCopy (sk_SkPaint *c_paint);
  void misk_delete_SkPaint (sk_SkPaint *obj);

  sk_SkPath *misk_new_Path ();
  sk_SkPath *misk_new_PathCopy (sk_SkPath *c_path);
  void misk_delete_SkPath (sk_SkPath *obj);

  sk_SkSurfaceProps *misk_new_SurfaceProps ();
  sk_SkSurfaceProps *misk_new_SurfacePropsPixelGeometry (unsigned int c_flags,
                                                         unsigned int c_p1);
  sk_SkSurfaceProps *misk_new_SurfacePropsCopy (sk_SkSurfaceProps *c_p0);

  sk_SkIRect misk_IRect_MakeLTRB (int c_l, int c_t, int c_r, int c_b);

  sk_SkFontMgr *sk_fontmgr_ref_default (void);

#ifdef __cplusplus
}
#endif // __cplusplus
