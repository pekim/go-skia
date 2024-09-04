// This is a generated file. DO NOT EDIT.

#include <include/core/SkBitmap.h>
#include <include/core/SkCanvas.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkSurfaceProps.h>

extern "C"
{
#include "api.h"

  void *
  misk_new_Bitmap ()
  {
    return reinterpret_cast<void *> (new SkBitmap ());
  }

  void *
  misk_new_BitmapCopy (void *c_src)
  {
    return reinterpret_cast<void *> (
        new SkBitmap (*reinterpret_cast<SkBitmap *> (c_src)));
  }

  void
  misk_delete_SkBitmap (void *obj)
  {
    delete reinterpret_cast<SkBitmap *> (obj);
  }

  void *
  misk_new_Canvas ()
  {
    return reinterpret_cast<void *> (new SkCanvas ());
  }

  void *
  misk_new_CanvasWithDimensions (int c_width, int c_height, void *c_props)
  {
    return reinterpret_cast<void *> (new SkCanvas (
        c_width, c_height, reinterpret_cast<SkSurfaceProps *> (c_props)));
  }

  void
  misk_delete_SkCanvas (void *obj)
  {
    delete reinterpret_cast<SkCanvas *> (obj);
  }

  void *
  misk_new_Paint ()
  {
    return reinterpret_cast<void *> (new SkPaint ());
  }

  void *
  misk_new_PaintCopy (void *c_paint)
  {
    return reinterpret_cast<void *> (
        new SkPaint (*reinterpret_cast<SkPaint *> (c_paint)));
  }

  void
  misk_delete_SkPaint (void *obj)
  {
    delete reinterpret_cast<SkPaint *> (obj);
  }

  void *
  misk_new_Path ()
  {
    return reinterpret_cast<void *> (new SkPath ());
  }

  void *
  misk_new_PathCopy (void *c_path)
  {
    return reinterpret_cast<void *> (
        new SkPath (*reinterpret_cast<SkPath *> (c_path)));
  }

  void
  misk_delete_SkPath (void *obj)
  {
    delete reinterpret_cast<SkPath *> (obj);
  }

  void *
  misk_new_SurfaceProps ()
  {
    return reinterpret_cast<void *> (new SkSurfaceProps ());
  }

  void *
  misk_new_SurfacePropsPixelGeometry (unsigned int c_flags, unsigned int c_p1)
  {
    return reinterpret_cast<void *> (
        new SkSurfaceProps (c_flags, SkPixelGeometry (c_p1)));
  }

  void *
  misk_new_SurfacePropsCopy (void *c_p0)
  {
    return reinterpret_cast<void *> (
        new SkSurfaceProps (*reinterpret_cast<SkSurfaceProps *> (c_p0)));
  }
}
