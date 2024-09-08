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
    float Left;
    float Top;
    float Right;
    float Bottom;
  } sk_SkRect;

  typedef struct
  {
    float Width;
    float Height;
  } sk_SkSize;

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

  sk_SkFont *misk_new_Font ();
  sk_SkFont *misk_new_FontTypefaceSize (sk_SkTypeface *c_typeface,
                                        float c_size);
  sk_SkFont *misk_new_FontTypeface (sk_SkTypeface *c_typeface);
  sk_SkFont *misk_new_FontTypefaceSizeScaleSkew (sk_SkTypeface *c_typeface,
                                                 float c_size, float c_scaleX,
                                                 float c_skewX);

  sk_SkFontStyle *misk_new_FontStyle2 (int c_weight, int c_width,
                                       unsigned int c_slant);
  sk_SkFontStyle *misk_new_FontStyle ();
  sk_SkFontStyle misk_FontStyle_Normal ();
  sk_SkFontStyle misk_FontStyle_Bold ();
  sk_SkFontStyle misk_FontStyle_Italic ();
  sk_SkFontStyle misk_FontStyle_BoldItalic ();

  sk_SkIRect misk_IRect_MakeEmpty ();
  sk_SkIRect misk_IRect_MakeWH (int c_w, int c_h);
  sk_SkIRect misk_IRect_MakeLTRB (int c_l, int c_t, int c_r, int c_b);
  sk_SkIRect misk_IRect_MakeXYWH (int c_x, int c_y, int c_w, int c_h);
  bool misk_IRect_Intersects (sk_SkIRect *c_a, sk_SkIRect *c_b);

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

  sk_SkRect misk_Rect_MakeEmpty ();
  sk_SkRect misk_Rect_MakeWH (float c_w, float c_h);
  sk_SkRect misk_Rect_MakeSize (sk_SkSize *c_size);
  sk_SkRect misk_Rect_MakeLTRB (float c_l, float c_t, float c_r, float c_b);
  sk_SkRect misk_Rect_MakeXYWH (float c_x, float c_y, float c_w, float c_h);
  sk_SkRect misk_Rect_MakeISize (sk_SkISize *c_size);
  sk_SkRect misk_Rect_MakeIRect (sk_SkIRect *c_irect);
  bool misk_Rect_Intersects (sk_SkRect *c_a, sk_SkRect *c_b);

  bool misk_Typeface_Equal (sk_SkTypeface *c_facea, sk_SkTypeface *c_faceb);
  sk_SkTypeface *misk_Typeface_MakeEmpty ();

  sk_SkFontMgr *sk_fontmgr_ref_default (void);

#ifdef __cplusplus
}
#endif // __cplusplus
