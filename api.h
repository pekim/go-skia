// This is a generated file. DO NOT EDIT.

#include <stdbool.h>
#include <sys/types.h>
typedef signed char schar;
typedef unsigned char uchar;

#ifdef __cplusplus
extern "C"
{
#endif // __cplusplus

  void *sk_fontmgr_ref_default (void);
  typedef struct
  {
    uchar fPixelRef[8];
    uchar fPixmap[40];
    uchar fMips[8];
  } sk_SkBitmap;
  void *misk_new_Bitmap ();
  void *misk_new_BitmapCopy (void *c_src);
  void misk_delete_SkBitmap (void *obj);
  bool misk_Bitmap_ComputeIsOpaque (void *c_bm);

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
  void *misk_new_Canvas ();
  void *misk_new_CanvasWithDimensions (int c_width, int c_height,
                                       void *c_props);
  void *misk_new_CanvasFromBitmap (void *c_bitmap);
  void *misk_new_CanvasFromBitmapSurfaceProps (void *c_bitmap, void *c_props);
  void misk_delete_SkCanvas (void *obj);

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
  void *misk_ColorSpace_MakeSRGB ();
  void *misk_ColorSpace_MakeSRGBLinear ();
  bool misk_ColorSpace_Equals (void *c_p0, void *c_p1);

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
    uchar fWidth[4];
    uchar fMiterLimit[4];
    uchar pad_9[8];
  } sk_SkPaint;
  void *misk_new_Paint ();
  void *misk_new_PaintCopy (void *c_paint);
  void misk_delete_SkPaint (void *obj);

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
  void *misk_new_Path ();
  void *misk_new_PathCopy (void *c_path);
  void misk_delete_SkPath (void *obj);

  typedef struct
  {
    unsigned int fFlags;
    uchar fPixelGeometry[4];
    uchar fTextContrast[4];
    uchar fTextGamma[4];
  } sk_SkSurfaceProps;
  void *misk_new_SurfaceProps ();
  void *misk_new_SurfacePropsPixelGeometry (unsigned int c_flags,
                                            unsigned int c_p1);
  void *misk_new_SurfacePropsCopy (void *c_p0);

#ifdef __cplusplus
}
#endif // __cplusplus
