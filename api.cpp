// This is a generated file. DO NOT EDIT.

#include <include/core/SkBitmap.h>
#include <include/core/SkCanvas.h>
#include <include/core/SkColorSpace.h>
#include <include/core/SkFontMgr.h>
#include <include/core/SkFontStyle.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkRect.h>
#include <include/core/SkSurfaceProps.h>

#if defined(SKIA_MAC)
#include "include/ports/SkFontMgr_mac_ct.h"
#endif

#if defined(SKIA_UNIX)
#include "include/ports/SkFontMgr_fontconfig.h"
#endif

#if defined(SKIA_WINDOWS)
#include "include/ports/SkTypeface_win.h"
#endif

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

  bool
  misk_Bitmap_ComputeIsOpaque (void *c_bm)
  {
    return SkBitmap::ComputeIsOpaque (*reinterpret_cast<SkBitmap *> (c_bm));
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

  void *
  misk_new_CanvasFromBitmap (void *c_bitmap)
  {
    return reinterpret_cast<void *> (
        new SkCanvas (*reinterpret_cast<SkBitmap *> (c_bitmap)));
  }

  void *
  misk_new_CanvasFromBitmapSurfaceProps (void *c_bitmap, void *c_props)
  {
    return reinterpret_cast<void *> (
        new SkCanvas (*reinterpret_cast<SkBitmap *> (c_bitmap),
                      *reinterpret_cast<SkSurfaceProps *> (c_props)));
  }

  void
  misk_delete_SkCanvas (void *obj)
  {
    delete reinterpret_cast<SkCanvas *> (obj);
  }

  void *
  misk_ColorSpace_MakeSRGB ()
  {
    return reinterpret_cast<void *> (SkColorSpace::MakeSRGB ().release ());
  }

  void *
  misk_ColorSpace_MakeSRGBLinear ()
  {
    return reinterpret_cast<void *> (
        SkColorSpace::MakeSRGBLinear ().release ());
  }

  bool
  misk_ColorSpace_Equals (void *c_p0, void *c_p1)
  {
    return SkColorSpace::Equals (reinterpret_cast<SkColorSpace *> (c_p0),
                                 reinterpret_cast<SkColorSpace *> (c_p1));
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

  void *
  sk_fontmgr_ref_default (void)
  {
#if defined(SKIA_MAC)
    return reinterpret_cast<void *> (
        SkFontMgr_New_CoreText (nullptr).release ());
#elif defined(SKIA_UNIX)
    return reinterpret_cast<void *> (
        SkFontMgr_New_FontConfig (nullptr).release ());
#elif defined(SKIA_WINDOWS)
    return reinterpret_cast<void *> (SkFontMgr_New_DirectWrite ().release ());
#else
#error "No font manager available for this platform"
#endif
  }
}
