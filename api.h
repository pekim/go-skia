// This is a generated file. DO NOT EDIT.

#include <stdbool.h>
#include <sys/types.h>
typedef signed char schar;
typedef unsigned char uchar;

#ifdef __cplusplus
extern "C"
{
#endif // __cplusplus

  void *misk_new_Bitmap ();
  void *misk_new_BitmapCopy (void *c_src);
  void misk_delete_SkBitmap (void *obj);
  void *misk_new_Canvas ();
  void *misk_new_CanvasWithDimensions (int c_width, int c_height,
                                       void *c_props);
  void misk_delete_SkCanvas (void *obj);
  void *misk_new_Paint ();
  void *misk_new_PaintCopy (void *c_paint);
  void misk_delete_SkPaint (void *obj);
  void *misk_new_Path ();
  void *misk_new_PathCopy (void *c_path);
  void misk_delete_SkPath (void *obj);
  void *misk_new_SurfaceProps ();
  void *misk_new_SurfacePropsPixelGeometry (unsigned int c_flags,
                                            unsigned int c_p1);
  void *misk_new_SurfacePropsCopy (void *c_p0);

#ifdef __cplusplus
}
#endif // __cplusplus
