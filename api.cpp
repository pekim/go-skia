// This is a generated file. DO NOT EDIT.

#include <include/core/SkBitmap.h>
#include <include/core/SkCanvas.h>
#include <include/core/SkColorSpace.h>
#include <include/core/SkFont.h>
#include <include/core/SkFontMgr.h>
#include <include/core/SkFontStyle.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkRect.h>
#include <include/core/SkSize.h>
#include <include/core/SkSurfaceProps.h>
#include <include/core/SkTypeface.h>
#include <include/core/SkTypes.h>
#include <include/gpu/GpuTypes.h>
#include <include/gpu/GrBackendSurface.h>
#include <include/gpu/GrContextOptions.h>
#include <include/gpu/GrDirectContext.h>
#include <include/gpu/ganesh/gl/GrGLBackendSurface.h>
#include <include/gpu/ganesh/gl/GrGLDirectContext.h>
#include <include/gpu/gl/GrGLInterface.h>
#include <include/gpu/gl/GrGLTypes.h>

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

  sk_GrBackendRenderTarget *
  misk_new_GrBackendRenderTarget ()
  {
    return reinterpret_cast<sk_GrBackendRenderTarget *> (
        new GrBackendRenderTarget ());
  }

  sk_GrBackendRenderTarget *
  misk_new_GrBackendRenderTargetCopy (sk_GrBackendRenderTarget *c_that)
  {
    return reinterpret_cast<sk_GrBackendRenderTarget *> (
        new GrBackendRenderTarget (
            *reinterpret_cast<GrBackendRenderTarget *> (c_that)));
  }

  void
  misk_delete_GrBackendRenderTarget (sk_GrBackendRenderTarget *obj)
  {
    delete reinterpret_cast<GrBackendRenderTarget *> (obj);
  }

  void
  misk_delete_GrDirectContext (sk_GrDirectContext *obj)
  {
    delete reinterpret_cast<GrDirectContext *> (obj);
  }

  sk_GrContextOptions *
  misk_new_GrContextOptions ()
  {
    return reinterpret_cast<sk_GrContextOptions *> (new GrContextOptions ());
  }

  sk_SkBitmap *
  misk_new_Bitmap ()
  {
    return reinterpret_cast<sk_SkBitmap *> (new SkBitmap ());
  }

  sk_SkBitmap *
  misk_new_BitmapCopy (sk_SkBitmap *c_src)
  {
    return reinterpret_cast<sk_SkBitmap *> (
        new SkBitmap (*reinterpret_cast<SkBitmap *> (c_src)));
  }

  void
  misk_delete_SkBitmap (sk_SkBitmap *obj)
  {
    delete reinterpret_cast<SkBitmap *> (obj);
  }

  bool
  misk_Bitmap_ComputeIsOpaque (sk_SkBitmap *c_bm)
  {
    return SkBitmap::ComputeIsOpaque (*reinterpret_cast<SkBitmap *> (c_bm));
  }

  sk_SkCanvas *
  misk_new_Canvas ()
  {
    return reinterpret_cast<sk_SkCanvas *> (new SkCanvas ());
  }

  sk_SkCanvas *
  misk_new_CanvasWithDimensions (int c_width, int c_height,
                                 sk_SkSurfaceProps *c_props)
  {
    return reinterpret_cast<sk_SkCanvas *> (new SkCanvas (
        c_width, c_height, reinterpret_cast<SkSurfaceProps *> (c_props)));
  }

  sk_SkCanvas *
  misk_new_CanvasFromBitmap (sk_SkBitmap *c_bitmap)
  {
    return reinterpret_cast<sk_SkCanvas *> (
        new SkCanvas (*reinterpret_cast<SkBitmap *> (c_bitmap)));
  }

  sk_SkCanvas *
  misk_new_CanvasFromBitmapSurfaceProps (sk_SkBitmap *c_bitmap,
                                         sk_SkSurfaceProps *c_props)
  {
    return reinterpret_cast<sk_SkCanvas *> (
        new SkCanvas (*reinterpret_cast<SkBitmap *> (c_bitmap),
                      *reinterpret_cast<SkSurfaceProps *> (c_props)));
  }

  void
  misk_delete_SkCanvas (sk_SkCanvas *obj)
  {
    delete reinterpret_cast<SkCanvas *> (obj);
  }

  sk_SkColorSpace *
  misk_ColorSpace_MakeSRGB ()
  {
    return reinterpret_cast<sk_SkColorSpace *> (
        SkColorSpace::MakeSRGB ().release ());
  }

  sk_SkColorSpace *
  misk_ColorSpace_MakeSRGBLinear ()
  {
    return reinterpret_cast<sk_SkColorSpace *> (
        SkColorSpace::MakeSRGBLinear ().release ());
  }

  bool
  misk_ColorSpace_Equals (sk_SkColorSpace *c_p0, sk_SkColorSpace *c_p1)
  {
    return SkColorSpace::Equals (reinterpret_cast<SkColorSpace *> (c_p0),
                                 reinterpret_cast<SkColorSpace *> (c_p1));
  }

  sk_SkFont *
  misk_new_Font ()
  {
    return reinterpret_cast<sk_SkFont *> (new SkFont ());
  }

  sk_SkFont *
  misk_new_FontTypefaceSize (sk_SkTypeface *c_typeface, float c_size)
  {
    return reinterpret_cast<sk_SkFont *> (new SkFont (
        sk_ref_sp (reinterpret_cast<SkTypeface *> (c_typeface)), c_size));
  }

  sk_SkFont *
  misk_new_FontTypeface (sk_SkTypeface *c_typeface)
  {
    return reinterpret_cast<sk_SkFont *> (
        new SkFont (sk_ref_sp (reinterpret_cast<SkTypeface *> (c_typeface))));
  }

  sk_SkFont *
  misk_new_FontTypefaceSizeScaleSkew (sk_SkTypeface *c_typeface, float c_size,
                                      float c_scaleX, float c_skewX)
  {
    return reinterpret_cast<sk_SkFont *> (
        new SkFont (sk_ref_sp (reinterpret_cast<SkTypeface *> (c_typeface)),
                    c_size, c_scaleX, c_skewX));
  }

  sk_SkFontStyle *
  misk_new_FontStyle2 (int c_weight, int c_width, uint c_slant)
  {
    return reinterpret_cast<sk_SkFontStyle *> (
        new SkFontStyle (c_weight, c_width, SkFontStyle::Slant (c_slant)));
  }

  sk_SkFontStyle *
  misk_new_FontStyle ()
  {
    return reinterpret_cast<sk_SkFontStyle *> (new SkFontStyle ());
  }

  sk_SkFontStyle
  misk_FontStyle_Normal ()
  {
    auto ret = (SkFontStyle::Normal ());
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_Bold ()
  {
    auto ret = (SkFontStyle::Bold ());
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_Italic ()
  {
    auto ret = (SkFontStyle::Italic ());
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_BoldItalic ()
  {
    auto ret = (SkFontStyle::BoldItalic ());
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeEmpty ()
  {
    auto ret = (SkIRect::MakeEmpty ());
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeWH (int c_w, int c_h)
  {
    auto ret = (SkIRect::MakeWH (c_w, c_h));
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeLTRB (int c_l, int c_t, int c_r, int c_b)
  {
    auto ret = (SkIRect::MakeLTRB (c_l, c_t, c_r, c_b));
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeXYWH (int c_x, int c_y, int c_w, int c_h)
  {
    auto ret = (SkIRect::MakeXYWH (c_x, c_y, c_w, c_h));
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  bool
  misk_IRect_Intersects (sk_SkIRect c_a, sk_SkIRect c_b)
  {
    return SkIRect::Intersects (*reinterpret_cast<SkIRect *> (&c_a),
                                *reinterpret_cast<SkIRect *> (&c_b));
  }

  sk_SkPaint *
  misk_new_Paint ()
  {
    return reinterpret_cast<sk_SkPaint *> (new SkPaint ());
  }

  sk_SkPaint *
  misk_new_PaintCopy (sk_SkPaint *c_paint)
  {
    return reinterpret_cast<sk_SkPaint *> (
        new SkPaint (*reinterpret_cast<SkPaint *> (c_paint)));
  }

  void
  misk_delete_SkPaint (sk_SkPaint *obj)
  {
    delete reinterpret_cast<SkPaint *> (obj);
  }

  void
  misk_Paint_reset (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->reset ();
  }

  unsigned char
  misk_Paint_getAlpha (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getAlpha ();
  }

  void
  misk_Paint_setAlpha (sk_SkPaint *c_obj, unsigned int c_a)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setAlpha (c_a);
  }

  void
  misk_Paint_setARGB (sk_SkPaint *c_obj, unsigned int c_a, unsigned int c_r,
                      unsigned int c_g, unsigned int c_b)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setARGB (c_a, c_r, c_g, c_b);
  }

  void
  misk_Paint_setAntiAlias (sk_SkPaint *c_obj, bool c_aa)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setAntiAlias (c_aa);
  }

  void
  misk_Paint_setBlendMode (sk_SkPaint *c_obj, int c_mode)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setBlendMode (
        SkBlendMode (c_mode));
  }

  void
  misk_Paint_setColor (sk_SkPaint *c_obj, unsigned int c_color)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setColor (c_color);
  }

  void
  misk_Paint_setDither (sk_SkPaint *c_obj, bool c_dither)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setDither (c_dither);
  }

  uint
  misk_Paint_getStrokeCap (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getStrokeCap ();
  }

  void
  misk_Paint_setStrokeCap (sk_SkPaint *c_obj, uint c_cap)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setStrokeCap (
        SkPaint::Cap (c_cap));
  }

  uint
  misk_Paint_getStrokeJoin (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getStrokeJoin ();
  }

  void
  misk_Paint_setStrokeJoin (sk_SkPaint *c_obj, uint c_join)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setStrokeJoin (
        SkPaint::Join (c_join));
  }

  float
  misk_Paint_getStrokeMiter (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getStrokeMiter ();
  }

  void
  misk_Paint_setStrokeMiter (sk_SkPaint *c_obj, float c_miter)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setStrokeMiter (c_miter);
  }

  float
  misk_Paint_getStrokeWidth (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getStrokeWidth ();
  }

  void
  misk_Paint_setStrokeWidth (sk_SkPaint *c_obj, float c_width)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setStrokeWidth (c_width);
  }

  uint
  misk_Paint_getStyle (sk_SkPaint *c_obj)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->getStyle ();
  }

  void
  misk_Paint_setStyle (sk_SkPaint *c_obj, uint c_style)
  {
    return reinterpret_cast<SkPaint *> (c_obj)->setStyle (
        SkPaint::Style (c_style));
  }

  sk_SkPath *
  misk_new_Path ()
  {
    return reinterpret_cast<sk_SkPath *> (new SkPath ());
  }

  sk_SkPath *
  misk_new_PathCopy (sk_SkPath *c_path)
  {
    return reinterpret_cast<sk_SkPath *> (
        new SkPath (*reinterpret_cast<SkPath *> (c_path)));
  }

  void
  misk_delete_SkPath (sk_SkPath *obj)
  {
    delete reinterpret_cast<SkPath *> (obj);
  }

  sk_SkRect
  misk_Rect_MakeEmpty ()
  {
    auto ret = (SkRect::MakeEmpty ());
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeWH (float c_w, float c_h)
  {
    auto ret = (SkRect::MakeWH (c_w, c_h));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeSize (sk_SkSize c_size)
  {
    auto ret = (SkRect::MakeSize (*reinterpret_cast<SkSize *> (&c_size)));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeLTRB (float c_l, float c_t, float c_r, float c_b)
  {
    auto ret = (SkRect::MakeLTRB (c_l, c_t, c_r, c_b));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeXYWH (float c_x, float c_y, float c_w, float c_h)
  {
    auto ret = (SkRect::MakeXYWH (c_x, c_y, c_w, c_h));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeISize (sk_SkISize c_size)
  {
    auto ret = (SkRect::Make (*reinterpret_cast<SkISize *> (&c_size)));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeIRect (sk_SkIRect c_irect)
  {
    auto ret = (SkRect::Make (*reinterpret_cast<SkIRect *> (&c_irect)));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  bool
  misk_Rect_Intersects (sk_SkRect c_a, sk_SkRect c_b)
  {
    return SkRect::Intersects (*reinterpret_cast<SkRect *> (&c_a),
                               *reinterpret_cast<SkRect *> (&c_b));
  }

  sk_SkSurfaceProps *
  misk_new_SurfaceProps ()
  {
    return reinterpret_cast<sk_SkSurfaceProps *> (new SkSurfaceProps ());
  }

  sk_SkSurfaceProps *
  misk_new_SurfacePropsPixelGeometry (unsigned int c_flags, uint c_p1)
  {
    return reinterpret_cast<sk_SkSurfaceProps *> (
        new SkSurfaceProps (c_flags, SkPixelGeometry (c_p1)));
  }

  sk_SkSurfaceProps *
  misk_new_SurfacePropsCopy (sk_SkSurfaceProps *c_p0)
  {
    return reinterpret_cast<sk_SkSurfaceProps *> (
        new SkSurfaceProps (*reinterpret_cast<SkSurfaceProps *> (c_p0)));
  }

  bool
  misk_Typeface_Equal (sk_SkTypeface *c_facea, sk_SkTypeface *c_faceb)
  {
    return SkTypeface::Equal (reinterpret_cast<SkTypeface *> (c_facea),
                              reinterpret_cast<SkTypeface *> (c_faceb));
  }

  sk_SkTypeface *
  misk_Typeface_MakeEmpty ()
  {
    return reinterpret_cast<sk_SkTypeface *> (
        SkTypeface::MakeEmpty ().release ());
  }

  sk_GrGLInterface *
  misk_GrGLMakeNativeInterface ()
  {
    return const_cast<sk_GrGLInterface *> (
        reinterpret_cast<const sk_GrGLInterface *> (
            GrGLMakeNativeInterface ().release ()));
  }

  sk_GrBackendRenderTarget
  misk_GrBackendRenderTargetsMakeGL (int c_width, int c_height,
                                     int c_sampleCnt, int c_stencilBits,
                                     sk_GrGLFramebufferInfo c_glInfo)
  {
    auto ret = (GrBackendRenderTargets::MakeGL (
        c_width, c_height, c_sampleCnt, c_stencilBits,
        *reinterpret_cast<GrGLFramebufferInfo *> (&c_glInfo)));
    return *(reinterpret_cast<sk_GrBackendRenderTarget *> (&ret));
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterfaceOptions (sk_GrGLInterface *c_p0,
                                               sk_GrContextOptions c_p1)
  {
    return reinterpret_cast<sk_GrDirectContext *> (
        GrDirectContexts::MakeGL (
            sk_ref_sp (reinterpret_cast<GrGLInterface *> (c_p0)),
            *reinterpret_cast<GrContextOptions *> (&c_p1))
            .release ());
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterface (sk_GrGLInterface *c_p0)
  {
    return reinterpret_cast<sk_GrDirectContext *> (
        GrDirectContexts::MakeGL (
            sk_ref_sp (reinterpret_cast<GrGLInterface *> (c_p0)))
            .release ());
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLOptions (sk_GrContextOptions c_p0)
  {
    return reinterpret_cast<sk_GrDirectContext *> (
        GrDirectContexts::MakeGL (
            *reinterpret_cast<GrContextOptions *> (&c_p0))
            .release ());
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGL ()
  {
    return reinterpret_cast<sk_GrDirectContext *> (
        GrDirectContexts::MakeGL ().release ());
  }

  sk_SkFontMgr *
  sk_fontmgr_ref_default (void)
  {
#if defined(SKIA_MAC)
    return reinterpret_cast<sk_SkFontMgr *> (
        SkFontMgr_New_CoreText (nullptr).release ());
#elif defined(SKIA_UNIX)
    return reinterpret_cast<sk_SkFontMgr *> (
        SkFontMgr_New_FontConfig (nullptr).release ());
#elif defined(SKIA_WINDOWS)
    return reinterpret_cast<sk_SkFontMgr *> (
        SkFontMgr_New_DirectWrite ().release ());
#else
#error "No font manager available for this platform"
#endif
  }
}
