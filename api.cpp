// This is a generated file. DO NOT EDIT.

#include <include/core/SkCanvas.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkSurfaceProps.h>

extern "C"
{
#include "api.h"

  void *
  misk_new_Canvas ()
  {
    return reinterpret_cast<void *> (new SkCanvas ());
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
  misk_new_SurfacePropsCopy (void *c_p0)
  {
    return reinterpret_cast<void *> (
        new SkSurfaceProps (*reinterpret_cast<SkSurfaceProps *> (c_p0)));
  }
}
