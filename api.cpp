// This is a generated file. DO NOT EDIT.

#include <include/core/SkArc.h>
#include <include/core/SkBitmap.h>
#include <include/core/SkCanvas.h>
#include <include/core/SkClipOp.h>
#include <include/core/SkColor.h>
#include <include/core/SkColorSpace.h>
#include <include/core/SkColorType.h>
#include <include/core/SkData.h>
#include <include/core/SkFont.h>
#include <include/core/SkFontMgr.h>
#include <include/core/SkFontStyle.h>
#include <include/core/SkFontTypes.h>
#include <include/core/SkImage.h>
#include <include/core/SkImageInfo.h>
#include <include/core/SkM44.h>
#include <include/core/SkMatrix.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkPixmap.h>
#include <include/core/SkRRect.h>
#include <include/core/SkRect.h>
#include <include/core/SkRegion.h>
#include <include/core/SkSamplingOptions.h>
#include <include/core/SkSize.h>
#include <include/core/SkString.h>
#include <include/core/SkSurface.h>
#include <include/core/SkSurfaceProps.h>
#include <include/core/SkTextBlob.h>
#include <include/core/SkTypeface.h>
#include <include/core/SkTypes.h>
#include <include/gpu/GpuTypes.h>
#include <include/gpu/GrBackendSurface.h>
#include <include/gpu/GrContextOptions.h>
#include <include/gpu/GrDirectContext.h>
#include <include/gpu/GrRecordingContext.h>
#include <include/gpu/GrTypes.h>
#include <include/gpu/ganesh/SkSurfaceGanesh.h>
#include <include/gpu/ganesh/gl/GrGLBackendSurface.h>
#include <include/gpu/ganesh/gl/GrGLDirectContext.h>
#include <include/gpu/gl/GrGLInterface.h>
#include <include/gpu/gl/GrGLTypes.h>
#include <include/private/base/SkPoint_impl.h>

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

  sk_SkISize
  misk_GrBackendRenderTarget_dimensions (sk_GrBackendRenderTarget *c_obj)
  {
    auto ret
        = reinterpret_cast<GrBackendRenderTarget *> (c_obj)->dimensions ();
    return *(reinterpret_cast<sk_SkISize *> (&ret));
  }

  int
  misk_GrBackendRenderTarget_width (sk_GrBackendRenderTarget *c_obj)
  {
    auto ret = reinterpret_cast<GrBackendRenderTarget *> (c_obj)->width ();
    return ret;
  }

  int
  misk_GrBackendRenderTarget_height (sk_GrBackendRenderTarget *c_obj)
  {
    auto ret = reinterpret_cast<GrBackendRenderTarget *> (c_obj)->height ();
    return ret;
  }

  int
  misk_GrBackendRenderTarget_sampleCnt (sk_GrBackendRenderTarget *c_obj)
  {
    auto ret = reinterpret_cast<GrBackendRenderTarget *> (c_obj)->sampleCnt ();
    return ret;
  }

  int
  misk_GrBackendRenderTarget_stencilBits (sk_GrBackendRenderTarget *c_obj)
  {
    auto ret
        = reinterpret_cast<GrBackendRenderTarget *> (c_obj)->stencilBits ();
    return ret;
  }

  bool
  misk_GrBackendRenderTarget_isFramebufferOnly (
      sk_GrBackendRenderTarget *c_obj)
  {
    auto ret = reinterpret_cast<GrBackendRenderTarget *> (c_obj)
                   ->isFramebufferOnly ();
    return ret;
  }

  void
  misk_delete_GrDirectContext (sk_GrDirectContext *obj)
  {
    delete reinterpret_cast<GrDirectContext *> (obj);
  }

  void
  misk_GrDirectContext_flushAndSubmit (sk_GrDirectContext *c_obj, bool c_sync)
  {
    reinterpret_cast<GrDirectContext *> (c_obj)->flushAndSubmit (
        GrSyncCpu (c_sync));
  }

  bool
  misk_GrDirectContext_flush (sk_GrDirectContext *c_obj,
                              sk_GrFlushInfo *c_info)
  {
    auto ret = reinterpret_cast<GrDirectContext *> (c_obj)->flush (
        *reinterpret_cast<GrFlushInfo *> (c_info));
    return (bool)ret;
  }

  bool
  misk_GrDirectContext_submit (sk_GrDirectContext *c_obj, bool c_sync)
  {
    auto ret = reinterpret_cast<GrDirectContext *> (c_obj)->submit (
        GrSyncCpu (c_sync));
    return ret;
  }

  void
  misk_delete_GrRecordingContext (sk_GrRecordingContext *obj)
  {
    delete reinterpret_cast<GrRecordingContext *> (obj);
  }

  sk_GrContextOptions *
  misk_new_GrContextOptions ()
  {
    return reinterpret_cast<sk_GrContextOptions *> (new GrContextOptions ());
  }

  sk_SkArc *
  misk_new_Arc ()
  {
    return reinterpret_cast<sk_SkArc *> (new SkArc ());
  }

  sk_SkArc *
  misk_new_ArcCopy (sk_SkArc *c_arc)
  {
    return reinterpret_cast<sk_SkArc *> (
        new SkArc (*reinterpret_cast<SkArc *> (c_arc)));
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
    auto ret
        = SkBitmap::ComputeIsOpaque (*reinterpret_cast<SkBitmap *> (c_bm));
    return ret;
  }

  int
  misk_Bitmap_bytesPerPixel (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->bytesPerPixel ();
    return ret;
  }

  int
  misk_Bitmap_rowBytesAsPixels (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->rowBytesAsPixels ();
    return ret;
  }

  int
  misk_Bitmap_shiftPerPixel (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->shiftPerPixel ();
    return ret;
  }

  bool
  misk_Bitmap_drawsNothing (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->drawsNothing ();
    return ret;
  }

  unsigned long
  misk_Bitmap_rowBytes (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->rowBytes ();
    return ret;
  }

  bool
  misk_Bitmap_isImmutable (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->isImmutable ();
    return ret;
  }

  void
  misk_Bitmap_setImmutable (sk_SkBitmap *c_obj)
  {
    reinterpret_cast<SkBitmap *> (c_obj)->setImmutable ();
  }

  sk_SkIRect
  misk_Bitmap_bounds (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->bounds ();
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkISize
  misk_Bitmap_dimensions (sk_SkBitmap *c_obj)
  {
    auto ret = reinterpret_cast<SkBitmap *> (c_obj)->dimensions ();
    return *(reinterpret_cast<sk_SkISize *> (&ret));
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

  bool
  misk_Canvas_getProps (sk_SkCanvas *c_obj, sk_SkSurfaceProps *c_props)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getProps (
        reinterpret_cast<SkSurfaceProps *> (c_props));
    return ret;
  }

  sk_SkSurfaceProps
  misk_Canvas_getBaseProps (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getBaseProps ();
    return *(reinterpret_cast<sk_SkSurfaceProps *> (&ret));
  }

  sk_SkSurfaceProps
  misk_Canvas_getTopProps (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getTopProps ();
    return *(reinterpret_cast<sk_SkSurfaceProps *> (&ret));
  }

  sk_SkISize
  misk_Canvas_getBaseLayerSize (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getBaseLayerSize ();
    return *(reinterpret_cast<sk_SkISize *> (&ret));
  }

  sk_SkSurface *
  misk_Canvas_makeSurface (sk_SkCanvas *c_obj, sk_SkImageInfo *c_info,
                           sk_SkSurfaceProps *c_props)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)
                   ->makeSurface (*reinterpret_cast<SkImageInfo *> (c_info),
                                  reinterpret_cast<SkSurfaceProps *> (c_props))
                   .release ();
    return reinterpret_cast<sk_SkSurface *> (ret);
  }

  sk_SkSurface *
  misk_Canvas_getSurface (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getSurface ();
    return reinterpret_cast<sk_SkSurface *> (ret);
  }

  sk_GrRecordingContext *
  misk_Canvas_recordingContext (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->recordingContext ();
    return reinterpret_cast<sk_GrRecordingContext *> (ret);
  }

  bool
  misk_Canvas_peekPixels (sk_SkCanvas *c_obj, sk_SkPixmap *c_pixmap)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->peekPixels (
        reinterpret_cast<SkPixmap *> (c_pixmap));
    return ret;
  }

  bool
  misk_Canvas_readPixelsImageInfo (sk_SkCanvas *c_obj,
                                   sk_SkImageInfo *c_dstInfo,
                                   void *c_dstPixels, ulong c_dstRowBytes,
                                   int c_srcX, int c_srcY)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->readPixels (
        *reinterpret_cast<SkImageInfo *> (c_dstInfo),
        reinterpret_cast<void *> (c_dstPixels), c_dstRowBytes, c_srcX, c_srcY);
    return ret;
  }

  bool
  misk_Canvas_readPixelsPixmap (sk_SkCanvas *c_obj, sk_SkPixmap *c_pixmap,
                                int c_srcX, int c_srcY)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->readPixels (
        *reinterpret_cast<SkPixmap *> (c_pixmap), c_srcX, c_srcY);
    return ret;
  }

  bool
  misk_Canvas_readPixelsBitmap (sk_SkCanvas *c_obj, sk_SkBitmap *c_bitmap,
                                int c_srcX, int c_srcY)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->readPixels (
        *reinterpret_cast<SkBitmap *> (c_bitmap), c_srcX, c_srcY);
    return ret;
  }

  bool
  misk_Canvas_writePixelsImageInfo (sk_SkCanvas *c_obj, sk_SkImageInfo *c_info,
                                    void *c_pixels, ulong c_rowBytes, int c_x,
                                    int c_y)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->writePixels (
        *reinterpret_cast<SkImageInfo *> (c_info),
        reinterpret_cast<void *> (c_pixels), c_rowBytes, c_x, c_y);
    return ret;
  }

  bool
  misk_Canvas_writePixelsBitmap (sk_SkCanvas *c_obj, sk_SkBitmap *c_bitmap,
                                 int c_x, int c_y)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->writePixels (
        *reinterpret_cast<SkBitmap *> (c_bitmap), c_x, c_y);
    return ret;
  }

  int
  misk_Canvas_save (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->save ();
    return ret;
  }

  int
  misk_Canvas_saveLayer (sk_SkCanvas *c_obj, sk_SkRect c_bounds,
                         sk_SkPaint *c_paint)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->saveLayer (
        reinterpret_cast<SkRect *> (&c_bounds),
        reinterpret_cast<SkPaint *> (c_paint));
    return ret;
  }

  int
  misk_Canvas_saveLayerAlpha (sk_SkCanvas *c_obj, sk_SkRect c_bounds,
                              uint c_alpha)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->saveLayerAlpha (
        reinterpret_cast<SkRect *> (&c_bounds), c_alpha);
    return ret;
  }

  void
  misk_Canvas_restore (sk_SkCanvas *c_obj)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->restore ();
  }

  int
  misk_Canvas_getSaveCount (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getSaveCount ();
    return ret;
  }

  void
  misk_Canvas_restoreToCount (sk_SkCanvas *c_obj, int c_saveCount)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->restoreToCount (c_saveCount);
  }

  void
  misk_Canvas_translate (sk_SkCanvas *c_obj, float c_dx, float c_dy)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->translate (c_dx, c_dy);
  }

  void
  misk_Canvas_scale (sk_SkCanvas *c_obj, float c_sx, float c_sy)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->scale (c_sx, c_sy);
  }

  void
  misk_Canvas_rotate (sk_SkCanvas *c_obj, float c_degrees)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->rotate (c_degrees);
  }

  void
  misk_Canvas_rotateAboutPoint (sk_SkCanvas *c_obj, float c_degrees,
                                float c_px, float c_py)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->rotate (c_degrees, c_px, c_py);
  }

  void
  misk_Canvas_skew (sk_SkCanvas *c_obj, float c_sx, float c_sy)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->skew (c_sx, c_sy);
  }

  void
  misk_Canvas_concatMatrix (sk_SkCanvas *c_obj, sk_SkMatrix *c_matrix)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->concat (
        *reinterpret_cast<SkMatrix *> (c_matrix));
  }

  void
  misk_Canvas_concatM44 (sk_SkCanvas *c_obj, sk_SkM44 *c_p0)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->concat (
        *reinterpret_cast<SkM44 *> (c_p0));
  }

  void
  misk_Canvas_setMatrixM44 (sk_SkCanvas *c_obj, sk_SkM44 *c_matrix)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->setMatrix (
        *reinterpret_cast<SkM44 *> (c_matrix));
  }

  void
  misk_Canvas_setMatrix (sk_SkCanvas *c_obj, sk_SkMatrix *c_matrix)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->setMatrix (
        *reinterpret_cast<SkMatrix *> (c_matrix));
  }

  void
  misk_Canvas_resetMatrix (sk_SkCanvas *c_obj)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->resetMatrix ();
  }

  void
  misk_Canvas_clipRect (sk_SkCanvas *c_obj, sk_SkRect c_rect, int c_op,
                        bool c_doAntiAlias)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clipRect (
        *reinterpret_cast<SkRect *> (&c_rect), SkClipOp (c_op), c_doAntiAlias);
  }

  void
  misk_Canvas_clipRRect (sk_SkCanvas *c_obj, sk_SkRRect c_rrect, int c_op,
                         bool c_doAntiAlias)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clipRRect (
        *reinterpret_cast<SkRRect *> (&c_rrect), SkClipOp (c_op),
        c_doAntiAlias);
  }

  void
  misk_Canvas_clipPath (sk_SkCanvas *c_obj, sk_SkPath *c_path, int c_op,
                        bool c_doAntiAlias)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clipPath (
        *reinterpret_cast<SkPath *> (c_path), SkClipOp (c_op), c_doAntiAlias);
  }

  void
  misk_Canvas_clipRegion (sk_SkCanvas *c_obj, sk_SkRegion *c_deviceRgn,
                          int c_op)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clipRegion (
        *reinterpret_cast<SkRegion *> (c_deviceRgn), SkClipOp (c_op));
  }

  bool
  misk_Canvas_quickRejectRect (sk_SkCanvas *c_obj, sk_SkRect c_rect)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->quickReject (
        *reinterpret_cast<SkRect *> (&c_rect));
    return ret;
  }

  bool
  misk_Canvas_quickRejectPath (sk_SkCanvas *c_obj, sk_SkPath *c_path)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->quickReject (
        *reinterpret_cast<SkPath *> (c_path));
    return ret;
  }

  sk_SkRect
  misk_Canvas_getLocalClipBoundsRect (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getLocalClipBounds ();
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  bool
  misk_Canvas_getLocalClipBoundsPath (sk_SkCanvas *c_obj, sk_SkRect c_bounds)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getLocalClipBounds (
        reinterpret_cast<SkRect *> (&c_bounds));
    return ret;
  }

  sk_SkIRect
  misk_Canvas_getDeviceClipBounds (sk_SkCanvas *c_obj)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getDeviceClipBounds ();
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  bool
  misk_Canvas_getDeviceClipBoundsRect (sk_SkCanvas *c_obj, sk_SkIRect c_bounds)
  {
    auto ret = reinterpret_cast<SkCanvas *> (c_obj)->getDeviceClipBounds (
        reinterpret_cast<SkIRect *> (&c_bounds));
    return ret;
  }

  void
  misk_Canvas_drawColor (sk_SkCanvas *c_obj, uint c_color, int c_mode)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawColor (SkColor (c_color),
                                                     SkBlendMode (c_mode));
  }

  void
  misk_Canvas_drawColor4f (sk_SkCanvas *c_obj, sk_SkRGBA4f c_color, int c_mode)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawColor (
        *reinterpret_cast<SkRGBA4f<kUnpremul_SkAlphaType> *> (&c_color),
        SkBlendMode (c_mode));
  }

  void
  misk_Canvas_clear (sk_SkCanvas *c_obj, uint c_color)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clear (SkColor (c_color));
  }

  void
  misk_Canvas_clear4f (sk_SkCanvas *c_obj, sk_SkRGBA4f c_color)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->clear (
        *reinterpret_cast<SkRGBA4f<kUnpremul_SkAlphaType> *> (&c_color));
  }

  void
  misk_Canvas_discard (sk_SkCanvas *c_obj)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->discard ();
  }

  void
  misk_Canvas_drawPaint (sk_SkCanvas *c_obj, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawPaint (
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawPointScalars (sk_SkCanvas *c_obj, float c_x, float c_y,
                                sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawPoint (
        c_x, c_y, *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawPoint (sk_SkCanvas *c_obj, sk_SkPoint c_p,
                         sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawPoint (
        *reinterpret_cast<SkPoint *> (&c_p),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawLineScalars (sk_SkCanvas *c_obj, float c_x0, float c_y0,
                               float c_x1, float c_y1, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawLine (
        c_x0, c_y0, c_x1, c_y1, *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawLinePoints (sk_SkCanvas *c_obj, sk_SkPoint c_p0,
                              sk_SkPoint c_p1, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawLine (
        *reinterpret_cast<SkPoint *> (&c_p0),
        *reinterpret_cast<SkPoint *> (&c_p1),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawRect (sk_SkCanvas *c_obj, sk_SkRect c_rect,
                        sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawRect (
        *reinterpret_cast<SkRect *> (&c_rect),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawIRect (sk_SkCanvas *c_obj, sk_SkIRect c_rect,
                         sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawIRect (
        *reinterpret_cast<SkIRect *> (&c_rect),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawRegion (sk_SkCanvas *c_obj, sk_SkRegion *c_region,
                          sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawRegion (
        *reinterpret_cast<SkRegion *> (c_region),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawOval (sk_SkCanvas *c_obj, sk_SkRect c_oval,
                        sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawOval (
        *reinterpret_cast<SkRect *> (&c_oval),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawRRect (sk_SkCanvas *c_obj, sk_SkRRect c_rrect,
                         sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawRRect (
        *reinterpret_cast<SkRRect *> (&c_rrect),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawDRRect (sk_SkCanvas *c_obj, sk_SkRRect c_outer,
                          sk_SkRRect c_inner, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawDRRect (
        *reinterpret_cast<SkRRect *> (&c_outer),
        *reinterpret_cast<SkRRect *> (&c_inner),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawCircleScalars (sk_SkCanvas *c_obj, float c_cx, float c_cy,
                                 float c_radius, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawCircle (
        c_cx, c_cy, c_radius, *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawCirclePoint (sk_SkCanvas *c_obj, sk_SkPoint c_center,
                               float c_radius, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawCircle (
        *reinterpret_cast<SkPoint *> (&c_center), c_radius,
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawArc (sk_SkCanvas *c_obj, sk_SkRect c_oval,
                       float c_startAngle, float c_sweepAngle,
                       bool c_useCenter, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawArc (
        *reinterpret_cast<SkRect *> (&c_oval), c_startAngle, c_sweepAngle,
        c_useCenter, *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawArcArc (sk_SkCanvas *c_obj, sk_SkArc *c_arc,
                          sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawArc (
        *reinterpret_cast<SkArc *> (c_arc),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawRoundRect (sk_SkCanvas *c_obj, sk_SkRect c_rect, float c_rx,
                             float c_ry, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawRoundRect (
        *reinterpret_cast<SkRect *> (&c_rect), c_rx, c_ry,
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawPath (sk_SkCanvas *c_obj, sk_SkPath *c_path,
                        sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawPath (
        *reinterpret_cast<SkPath *> (c_path),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawImage (sk_SkCanvas *c_obj, sk_SkImage *c_image, float c_left,
                         float c_top)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawImage (
        reinterpret_cast<SkImage *> (c_image), c_left, c_top);
  }

  void
  misk_Canvas_drawImageSamplingOptions (sk_SkCanvas *c_obj, sk_SkImage *c_p0,
                                        float c_x, float c_y,
                                        sk_SkSamplingOptions *c_p3,
                                        sk_SkPaint *c_p4)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawImage (
        reinterpret_cast<SkImage *> (c_p0), c_x, c_y,
        *reinterpret_cast<SkSamplingOptions *> (c_p3),
        reinterpret_cast<SkPaint *> (c_p4));
  }

  void
  misk_Canvas_drawImageRect (sk_SkCanvas *c_obj, sk_SkImage *c_p0,
                             sk_SkRect c_src, sk_SkRect c_dst,
                             sk_SkSamplingOptions *c_p3, sk_SkPaint *c_p4,
                             uint c_p5)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawImageRect (
        reinterpret_cast<SkImage *> (c_p0),
        *reinterpret_cast<SkRect *> (&c_src),
        *reinterpret_cast<SkRect *> (&c_dst),
        *reinterpret_cast<SkSamplingOptions *> (c_p3),
        reinterpret_cast<SkPaint *> (c_p4),
        SkCanvas::SrcRectConstraint (c_p5));
  }

  void
  misk_Canvas_drawImageNine (sk_SkCanvas *c_obj, sk_SkImage *c_image,
                             sk_SkIRect c_center, sk_SkRect c_dst,
                             int c_filter, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawImageNine (
        reinterpret_cast<SkImage *> (c_image),
        *reinterpret_cast<SkIRect *> (&c_center),
        *reinterpret_cast<SkRect *> (&c_dst), SkFilterMode (c_filter),
        reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawString (sk_SkCanvas *c_obj, char *c_str, float c_x,
                          float c_y, sk_SkFont *c_font, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawString (
        (char *)c_str, c_x, c_y, *reinterpret_cast<SkFont *> (c_font),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawGlyphs (sk_SkCanvas *c_obj, int c_count, ushort *c_glyphs,
                          sk_SkPoint *c_positions, uint *c_clusters,
                          int c_textByteCount, char *c_utf8text,
                          sk_SkPoint c_origin, sk_SkFont *c_font,
                          sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawGlyphs (
        c_count, (unsigned short *)c_glyphs, (SkPoint *)c_positions,
        (unsigned int *)c_clusters, c_textByteCount, (char *)c_utf8text,
        *reinterpret_cast<SkPoint *> (&c_origin),
        *reinterpret_cast<SkFont *> (c_font),
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  void
  misk_Canvas_drawTextBlob (sk_SkCanvas *c_obj, sk_SkTextBlob *c_blob,
                            float c_x, float c_y, sk_SkPaint *c_paint)
  {
    reinterpret_cast<SkCanvas *> (c_obj)->drawTextBlob (
        reinterpret_cast<SkTextBlob *> (c_blob), c_x, c_y,
        *reinterpret_cast<SkPaint *> (c_paint));
  }

  sk_SkColorSpace *
  misk_ColorSpace_MakeSRGB ()
  {
    auto ret = SkColorSpace::MakeSRGB ().release ();
    return reinterpret_cast<sk_SkColorSpace *> (ret);
  }

  sk_SkColorSpace *
  misk_ColorSpace_MakeSRGBLinear ()
  {
    auto ret = SkColorSpace::MakeSRGBLinear ().release ();
    return reinterpret_cast<sk_SkColorSpace *> (ret);
  }

  bool
  misk_ColorSpace_Equals (sk_SkColorSpace *c_p0, sk_SkColorSpace *c_p1)
  {
    auto ret = SkColorSpace::Equals (reinterpret_cast<SkColorSpace *> (c_p0),
                                     reinterpret_cast<SkColorSpace *> (c_p1));
    return ret;
  }

  unsigned long
  misk_Data_size (sk_SkData *c_obj)
  {
    auto ret = reinterpret_cast<SkData *> (c_obj)->size ();
    return ret;
  }

  bool
  misk_Data_isEmpty (sk_SkData *c_obj)
  {
    auto ret = reinterpret_cast<SkData *> (c_obj)->isEmpty ();
    return ret;
  }

  sk_SkData *
  misk_Data_MakeWithCopy (void *c_data, ulong c_length)
  {
    auto ret
        = SkData::MakeWithCopy (reinterpret_cast<void *> (c_data), c_length)
              .release ();
    return reinterpret_cast<sk_SkData *> (ret);
  }

  sk_SkData *
  misk_Data_MakeZeroInitialized (ulong c_length)
  {
    auto ret = SkData::MakeZeroInitialized (c_length).release ();
    return reinterpret_cast<sk_SkData *> (ret);
  }

  sk_SkData *
  misk_Data_MakeWithCString (char *c_cstr)
  {
    auto ret = SkData::MakeWithCString ((char *)c_cstr).release ();
    return reinterpret_cast<sk_SkData *> (ret);
  }

  sk_SkData *
  misk_Data_MakeWithoutCopy (void *c_data, ulong c_length)
  {
    auto ret
        = SkData::MakeWithoutCopy (reinterpret_cast<void *> (c_data), c_length)
              .release ();
    return reinterpret_cast<sk_SkData *> (ret);
  }

  sk_SkData *
  misk_Data_MakeFromFileName (char *c_path)
  {
    auto ret = SkData::MakeFromFileName ((char *)c_path).release ();
    return reinterpret_cast<sk_SkData *> (ret);
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

  sk_SkFontStyleSet *
  misk_FontMgr_matchFamily (sk_SkFontMgr *c_obj, char *c_familyName)
  {
    auto ret = reinterpret_cast<SkFontMgr *> (c_obj)
                   ->matchFamily ((char *)c_familyName)
                   .release ();
    return reinterpret_cast<sk_SkFontStyleSet *> (ret);
  }

  sk_SkTypeface *
  misk_FontMgr_matchFamilyStyle (sk_SkFontMgr *c_obj, char *c_familyName,
                                 sk_SkFontStyle *c_p1)
  {
    auto ret = reinterpret_cast<SkFontMgr *> (c_obj)
                   ->matchFamilyStyle ((char *)c_familyName,
                                       *reinterpret_cast<SkFontStyle *> (c_p1))
                   .release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
  }

  sk_SkTypeface *
  misk_FontMgr_makeFromData (sk_SkFontMgr *c_obj, sk_SkData *c_p0,
                             int c_ttcIndex)
  {
    auto ret
        = reinterpret_cast<SkFontMgr *> (c_obj)
              ->makeFromData (sk_ref_sp (reinterpret_cast<SkData *> (c_p0)),
                              c_ttcIndex)
              .release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
  }

  sk_SkTypeface *
  misk_FontMgr_makeFromFile (sk_SkFontMgr *c_obj, char *c_path, int c_ttcIndex)
  {
    auto ret = reinterpret_cast<SkFontMgr *> (c_obj)
                   ->makeFromFile ((char *)c_path, c_ttcIndex)
                   .release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
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
    auto ret = SkFontStyle::Normal ();
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_Bold ()
  {
    auto ret = SkFontStyle::Bold ();
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_Italic ()
  {
    auto ret = SkFontStyle::Italic ();
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  sk_SkFontStyle
  misk_FontStyle_BoldItalic ()
  {
    auto ret = SkFontStyle::BoldItalic ();
    return *(reinterpret_cast<sk_SkFontStyle *> (&ret));
  }

  int
  misk_FontStyleSet_count (sk_SkFontStyleSet *c_obj)
  {
    auto ret = reinterpret_cast<SkFontStyleSet *> (c_obj)->count ();
    return ret;
  }

  void
  misk_FontStyleSet_getStyle (sk_SkFontStyleSet *c_obj, int c_index,
                              sk_SkFontStyle *c_p1, sk_SkString *c_style)
  {
    reinterpret_cast<SkFontStyleSet *> (c_obj)->getStyle (
        c_index, reinterpret_cast<SkFontStyle *> (c_p1),
        reinterpret_cast<SkString *> (c_style));
  }

  sk_SkTypeface *
  misk_FontStyleSet_createTypeface (sk_SkFontStyleSet *c_obj, int c_index)
  {
    auto ret = reinterpret_cast<SkFontStyleSet *> (c_obj)
                   ->createTypeface (c_index)
                   .release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
  }

  sk_SkTypeface *
  misk_FontStyleSet_matchStyle (sk_SkFontStyleSet *c_obj,
                                sk_SkFontStyle *c_pattern)
  {
    auto ret = reinterpret_cast<SkFontStyleSet *> (c_obj)
                   ->matchStyle (*reinterpret_cast<SkFontStyle *> (c_pattern))
                   .release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
  }

  int
  misk_Image_width (sk_SkImage *c_obj)
  {
    auto ret = reinterpret_cast<SkImage *> (c_obj)->width ();
    return ret;
  }

  int
  misk_Image_height (sk_SkImage *c_obj)
  {
    auto ret = reinterpret_cast<SkImage *> (c_obj)->height ();
    return ret;
  }

  sk_SkISize
  misk_Image_dimensions (sk_SkImage *c_obj)
  {
    auto ret = reinterpret_cast<SkImage *> (c_obj)->dimensions ();
    return *(reinterpret_cast<sk_SkISize *> (&ret));
  }

  sk_SkIRect
  misk_Image_bounds (sk_SkImage *c_obj)
  {
    auto ret = reinterpret_cast<SkImage *> (c_obj)->bounds ();
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkImageInfo *
  misk_new_ImageInfo ()
  {
    return reinterpret_cast<sk_SkImageInfo *> (new SkImageInfo ());
  }

  sk_SkM44 *
  misk_new_M44Copy (sk_SkM44 *c_src)
  {
    return reinterpret_cast<sk_SkM44 *> (
        new SkM44 (*reinterpret_cast<SkM44 *> (c_src)));
  }

  sk_SkM44 *
  misk_new_M44 ()
  {
    return reinterpret_cast<sk_SkM44 *> (new SkM44 ());
  }

  sk_SkM44 *
  misk_new_M44AB (sk_SkM44 *c_a, sk_SkM44 *c_b)
  {
    return reinterpret_cast<sk_SkM44 *> (new SkM44 (
        *reinterpret_cast<SkM44 *> (c_a), *reinterpret_cast<SkM44 *> (c_b)));
  }

  sk_SkM44 *
  misk_new_M44Scalars (float c_m0, float c_m4, float c_m8, float c_m12,
                       float c_m1, float c_m5, float c_m9, float c_m13,
                       float c_m2, float c_m6, float c_m10, float c_m14,
                       float c_m3, float c_m7, float c_m11, float c_m15)
  {
    return reinterpret_cast<sk_SkM44 *> (
        new SkM44 (c_m0, c_m4, c_m8, c_m12, c_m1, c_m5, c_m9, c_m13, c_m2,
                   c_m6, c_m10, c_m14, c_m3, c_m7, c_m11, c_m15));
  }

  sk_SkIRect
  misk_IRect_MakeEmpty ()
  {
    auto ret = SkIRect::MakeEmpty ();
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeWH (int c_w, int c_h)
  {
    auto ret = SkIRect::MakeWH (c_w, c_h);
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeLTRB (int c_l, int c_t, int c_r, int c_b)
  {
    auto ret = SkIRect::MakeLTRB (c_l, c_t, c_r, c_b);
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeXYWH (int c_x, int c_y, int c_w, int c_h)
  {
    auto ret = SkIRect::MakeXYWH (c_x, c_y, c_w, c_h);
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  sk_SkIRect
  misk_IRect_MakeSize (sk_SkISize c_size)
  {
    auto ret = SkIRect::MakeSize (*reinterpret_cast<SkISize *> (&c_size));
    return *(reinterpret_cast<sk_SkIRect *> (&ret));
  }

  bool
  misk_IRect_Intersects (sk_SkIRect c_a, sk_SkIRect c_b)
  {
    auto ret = SkIRect::Intersects (*reinterpret_cast<SkIRect *> (&c_a),
                                    *reinterpret_cast<SkIRect *> (&c_b));
    return ret;
  }

  int
  misk_IRect_x (sk_SkIRect *c_obj)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->x ();
    return ret;
  }

  int
  misk_IRect_y (sk_SkIRect *c_obj)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->y ();
    return ret;
  }

  int
  misk_IRect_width (sk_SkIRect *c_obj)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->width ();
    return ret;
  }

  int
  misk_IRect_height (sk_SkIRect *c_obj)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->height ();
    return ret;
  }

  bool
  misk_IRect_isEmpty (sk_SkIRect *c_obj)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->isEmpty ();
    return ret;
  }

  void
  misk_IRect_setEmpty (sk_SkIRect *c_obj)
  {
    reinterpret_cast<SkIRect *> (c_obj)->setEmpty ();
  }

  void
  misk_IRect_setLTRB (sk_SkIRect *c_obj, int c_left, int c_top, int c_right,
                      int c_bottom)
  {
    reinterpret_cast<SkIRect *> (c_obj)->setLTRB (c_left, c_top, c_right,
                                                  c_bottom);
  }

  void
  misk_IRect_setXYWH (sk_SkIRect *c_obj, int c_x, int c_y, int c_width,
                      int c_height)
  {
    reinterpret_cast<SkIRect *> (c_obj)->setXYWH (c_x, c_y, c_width, c_height);
  }

  void
  misk_IRect_setWH (sk_SkIRect *c_obj, int c_width, int c_height)
  {
    reinterpret_cast<SkIRect *> (c_obj)->setWH (c_width, c_height);
  }

  void
  misk_IRect_offset (sk_SkIRect *c_obj, int c_dx, int c_dy)
  {
    reinterpret_cast<SkIRect *> (c_obj)->offset (c_dx, c_dy);
  }

  void
  misk_IRect_offsetPoint (sk_SkIRect *c_obj, sk_SkIPoint *c_delta)
  {
    reinterpret_cast<SkIRect *> (c_obj)->offset (
        *reinterpret_cast<SkIPoint *> (c_delta));
  }

  void
  misk_IRect_offsetTo (sk_SkIRect *c_obj, int c_newX, int c_newY)
  {
    reinterpret_cast<SkIRect *> (c_obj)->offsetTo (c_newX, c_newY);
  }

  void
  misk_IRect_inset (sk_SkIRect *c_obj, int c_dx, int c_dy)
  {
    reinterpret_cast<SkIRect *> (c_obj)->inset (c_dx, c_dy);
  }

  void
  misk_IRect_outset (sk_SkIRect *c_obj, int c_dx, int c_dy)
  {
    reinterpret_cast<SkIRect *> (c_obj)->outset (c_dx, c_dy);
  }

  void
  misk_IRect_adjust (sk_SkIRect *c_obj, int c_dL, int c_dT, int c_dR, int c_dB)
  {
    reinterpret_cast<SkIRect *> (c_obj)->adjust (c_dL, c_dT, c_dR, c_dB);
  }

  bool
  misk_IRect_contains (sk_SkIRect *c_obj, int c_x, int c_y)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->contains (c_x, c_y);
    return ret;
  }

  bool
  misk_IRect_containsRect (sk_SkIRect *c_obj, sk_SkIRect c_r)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->contains (
        *reinterpret_cast<SkIRect *> (&c_r));
    return ret;
  }

  bool
  misk_IRect_containsNoEmptyCheck (sk_SkIRect *c_obj, sk_SkIRect c_r)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->containsNoEmptyCheck (
        *reinterpret_cast<SkIRect *> (&c_r));
    return ret;
  }

  bool
  misk_IRect_intersect (sk_SkIRect *c_obj, sk_SkIRect c_r)
  {
    auto ret = reinterpret_cast<SkIRect *> (c_obj)->intersect (
        *reinterpret_cast<SkIRect *> (&c_r));
    return ret;
  }

  void
  misk_IRect_join (sk_SkIRect *c_obj, sk_SkIRect c_r)
  {
    reinterpret_cast<SkIRect *> (c_obj)->join (
        *reinterpret_cast<SkIRect *> (&c_r));
  }

  void
  misk_IRect_sort (sk_SkIRect *c_obj)
  {
    reinterpret_cast<SkIRect *> (c_obj)->sort ();
  }

  sk_SkMatrix *
  misk_new_Matrix ()
  {
    return reinterpret_cast<sk_SkMatrix *> (new SkMatrix ());
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
    reinterpret_cast<SkPaint *> (c_obj)->reset ();
  }

  unsigned char
  misk_Paint_getAlpha (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getAlpha ();
    return ret;
  }

  void
  misk_Paint_setAlpha (sk_SkPaint *c_obj, uint c_a)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setAlpha (c_a);
  }

  void
  misk_Paint_setARGB (sk_SkPaint *c_obj, uint c_a, uint c_r, uint c_g,
                      uint c_b)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setARGB (c_a, c_r, c_g, c_b);
  }

  void
  misk_Paint_setAntiAlias (sk_SkPaint *c_obj, bool c_aa)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setAntiAlias (c_aa);
  }

  void
  misk_Paint_setBlendMode (sk_SkPaint *c_obj, int c_mode)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setBlendMode (SkBlendMode (c_mode));
  }

  void
  misk_Paint_setColor (sk_SkPaint *c_obj, uint c_color)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setColor (SkColor (c_color));
  }

  void
  misk_Paint_setDither (sk_SkPaint *c_obj, bool c_dither)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setDither (c_dither);
  }

  uint
  misk_Paint_getStrokeCap (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getStrokeCap ();
    return (uint)ret;
  }

  void
  misk_Paint_setStrokeCap (sk_SkPaint *c_obj, uint c_cap)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setStrokeCap (SkPaint::Cap (c_cap));
  }

  uchar
  misk_Paint_getStrokeJoin (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getStrokeJoin ();
    return (uchar)ret;
  }

  void
  misk_Paint_setStrokeJoin (sk_SkPaint *c_obj, uchar c_join)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setStrokeJoin (
        SkPaint::Join (c_join));
  }

  float
  misk_Paint_getStrokeMiter (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getStrokeMiter ();
    return ret;
  }

  void
  misk_Paint_setStrokeMiter (sk_SkPaint *c_obj, float c_miter)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setStrokeMiter (c_miter);
  }

  float
  misk_Paint_getStrokeWidth (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getStrokeWidth ();
    return ret;
  }

  void
  misk_Paint_setStrokeWidth (sk_SkPaint *c_obj, float c_width)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setStrokeWidth (c_width);
  }

  uchar
  misk_Paint_getStyle (sk_SkPaint *c_obj)
  {
    auto ret = reinterpret_cast<SkPaint *> (c_obj)->getStyle ();
    return (uchar)ret;
  }

  void
  misk_Paint_setStyle (sk_SkPaint *c_obj, uchar c_style)
  {
    reinterpret_cast<SkPaint *> (c_obj)->setStyle (SkPaint::Style (c_style));
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

  sk_SkPixmap *
  misk_new_Pixmap ()
  {
    return reinterpret_cast<sk_SkPixmap *> (new SkPixmap ());
  }

  sk_SkPixmap *
  misk_new_PixmapImageInfo (sk_SkImageInfo *c_info, void *c_addr,
                            ulong c_rowBytes)
  {
    return reinterpret_cast<sk_SkPixmap *> (
        new SkPixmap (*reinterpret_cast<SkImageInfo *> (c_info),
                      reinterpret_cast<void *> (c_addr), c_rowBytes));
  }

  sk_SkRect
  misk_Rect_MakeEmpty ()
  {
    auto ret = SkRect::MakeEmpty ();
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeWH (float c_w, float c_h)
  {
    auto ret = SkRect::MakeWH (c_w, c_h);
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeSize (sk_SkSize c_size)
  {
    auto ret = SkRect::MakeSize (*reinterpret_cast<SkSize *> (&c_size));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeLTRB (float c_l, float c_t, float c_r, float c_b)
  {
    auto ret = SkRect::MakeLTRB (c_l, c_t, c_r, c_b);
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeXYWH (float c_x, float c_y, float c_w, float c_h)
  {
    auto ret = SkRect::MakeXYWH (c_x, c_y, c_w, c_h);
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeISize (sk_SkISize c_size)
  {
    auto ret = SkRect::Make (*reinterpret_cast<SkISize *> (&c_size));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  sk_SkRect
  misk_Rect_MakeIRect (sk_SkIRect c_irect)
  {
    auto ret = SkRect::Make (*reinterpret_cast<SkIRect *> (&c_irect));
    return *(reinterpret_cast<sk_SkRect *> (&ret));
  }

  bool
  misk_Rect_Intersects (sk_SkRect c_a, sk_SkRect c_b)
  {
    auto ret = SkRect::Intersects (*reinterpret_cast<SkRect *> (&c_a),
                                   *reinterpret_cast<SkRect *> (&c_b));
    return ret;
  }

  float
  misk_Rect_x (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->x ();
    return ret;
  }

  float
  misk_Rect_y (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->y ();
    return ret;
  }

  float
  misk_Rect_width (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->width ();
    return ret;
  }

  float
  misk_Rect_height (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->height ();
    return ret;
  }

  float
  misk_Rect_centerX (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->centerX ();
    return ret;
  }

  float
  misk_Rect_centerY (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->centerY ();
    return ret;
  }

  bool
  misk_Rect_isEmpty (sk_SkRect *c_obj)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->isEmpty ();
    return ret;
  }

  void
  misk_Rect_setEmpty (sk_SkRect *c_obj)
  {
    reinterpret_cast<SkRect *> (c_obj)->setEmpty ();
  }

  void
  misk_Rect_setLTRB (sk_SkRect *c_obj, float c_left, float c_top,
                     float c_right, float c_bottom)
  {
    reinterpret_cast<SkRect *> (c_obj)->setLTRB (c_left, c_top, c_right,
                                                 c_bottom);
  }

  void
  misk_Rect_setXYWH (sk_SkRect *c_obj, float c_x, float c_y, float c_width,
                     float c_height)
  {
    reinterpret_cast<SkRect *> (c_obj)->setXYWH (c_x, c_y, c_width, c_height);
  }

  void
  misk_Rect_setWH (sk_SkRect *c_obj, float c_width, float c_height)
  {
    reinterpret_cast<SkRect *> (c_obj)->setWH (c_width, c_height);
  }

  void
  misk_Rect_offset (sk_SkRect *c_obj, float c_dx, float c_dy)
  {
    reinterpret_cast<SkRect *> (c_obj)->offset (c_dx, c_dy);
  }

  void
  misk_Rect_offsetTo (sk_SkRect *c_obj, float c_newX, float c_newY)
  {
    reinterpret_cast<SkRect *> (c_obj)->offsetTo (c_newX, c_newY);
  }

  void
  misk_Rect_inset (sk_SkRect *c_obj, float c_dx, float c_dy)
  {
    reinterpret_cast<SkRect *> (c_obj)->inset (c_dx, c_dy);
  }

  void
  misk_Rect_outset (sk_SkRect *c_obj, float c_dx, float c_dy)
  {
    reinterpret_cast<SkRect *> (c_obj)->outset (c_dx, c_dy);
  }

  bool
  misk_Rect_contains (sk_SkRect *c_obj, float c_x, float c_y)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->contains (c_x, c_y);
    return ret;
  }

  bool
  misk_Rect_containsRect (sk_SkRect *c_obj, sk_SkRect c_r)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->contains (
        *reinterpret_cast<SkRect *> (&c_r));
    return ret;
  }

  bool
  misk_Rect_intersect (sk_SkRect *c_obj, sk_SkRect c_r)
  {
    auto ret = reinterpret_cast<SkRect *> (c_obj)->intersect (
        *reinterpret_cast<SkRect *> (&c_r));
    return ret;
  }

  void
  misk_Rect_join (sk_SkRect *c_obj, sk_SkRect c_r)
  {
    reinterpret_cast<SkRect *> (c_obj)->join (
        *reinterpret_cast<SkRect *> (&c_r));
  }

  void
  misk_Rect_sort (sk_SkRect *c_obj)
  {
    reinterpret_cast<SkRect *> (c_obj)->sort ();
  }

  sk_SkRRect *
  misk_new_RRect ()
  {
    return reinterpret_cast<sk_SkRRect *> (new SkRRect ());
  }

  sk_SkRRect *
  misk_new_RRectCopy (sk_SkRRect c_rrect)
  {
    return reinterpret_cast<sk_SkRRect *> (
        new SkRRect (*reinterpret_cast<SkRRect *> (&c_rrect)));
  }

  sk_SkRegion *
  misk_new_Region ()
  {
    return reinterpret_cast<sk_SkRegion *> (new SkRegion ());
  }

  sk_SkRegion *
  misk_new_RegionCopy (sk_SkRegion *c_region)
  {
    return reinterpret_cast<sk_SkRegion *> (
        new SkRegion (*reinterpret_cast<SkRegion *> (c_region)));
  }

  sk_SkRegion *
  misk_new_RegionCopyRect (sk_SkIRect c_rect)
  {
    return reinterpret_cast<sk_SkRegion *> (
        new SkRegion (*reinterpret_cast<SkIRect *> (&c_rect)));
  }

  void
  misk_delete_SkRegion (sk_SkRegion *obj)
  {
    delete reinterpret_cast<SkRegion *> (obj);
  }

  sk_SkSamplingOptions *
  misk_new_SamplingOptions ()
  {
    return reinterpret_cast<sk_SkSamplingOptions *> (new SkSamplingOptions ());
  }

  sk_SkSamplingOptions *
  misk_new_SamplingOptionsCopy (sk_SkSamplingOptions *c_p0)
  {
    return reinterpret_cast<sk_SkSamplingOptions *> (
        new SkSamplingOptions (*reinterpret_cast<SkSamplingOptions *> (c_p0)));
  }

  sk_SkString *
  misk_new_String (char *c_text)
  {
    return reinterpret_cast<sk_SkString *> (new SkString ((char *)c_text));
  }

  void
  misk_delete_SkString (sk_SkString *obj)
  {
    delete reinterpret_cast<SkString *> (obj);
  }

  char *
  misk_String_data (sk_SkString *c_obj)
  {
    auto ret = reinterpret_cast<SkString *> (c_obj)->data ();
    return reinterpret_cast<char *> (ret);
  }

  sk_SkCanvas *
  misk_Surface_getCanvas (sk_SkSurface *c_obj)
  {
    auto ret = reinterpret_cast<SkSurface *> (c_obj)->getCanvas ();
    return reinterpret_cast<sk_SkCanvas *> (ret);
  }

  sk_SkSurfaceProps *
  misk_new_SurfaceProps ()
  {
    return reinterpret_cast<sk_SkSurfaceProps *> (new SkSurfaceProps ());
  }

  sk_SkSurfaceProps *
  misk_new_SurfacePropsPixelGeometry (uint c_flags, uint c_p1)
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

  sk_SkTextBlob *
  misk_TextBlob_MakeFromString (char *c_string, sk_SkFont *c_font,
                                int c_encoding)
  {
    auto ret = SkTextBlob::MakeFromString (
                   (char *)c_string, *reinterpret_cast<SkFont *> (c_font),
                   SkTextEncoding (c_encoding))
                   .release ();
    return reinterpret_cast<sk_SkTextBlob *> (ret);
  }

  sk_SkTextBlob *
  misk_TextBlob_MakeFromPosTextH (void *c_text, ulong c_byteLength,
                                  float *c_xpos, float c_constY,
                                  sk_SkFont *c_font, int c_encoding)
  {
    auto ret
        = SkTextBlob::MakeFromPosTextH (
              reinterpret_cast<void *> (c_text), c_byteLength, (float *)c_xpos,
              c_constY, *reinterpret_cast<SkFont *> (c_font),
              SkTextEncoding (c_encoding))
              .release ();
    return reinterpret_cast<sk_SkTextBlob *> (ret);
  }

  sk_SkTextBlob *
  misk_TextBlob_MakeFromPosText (void *c_text, ulong c_byteLength,
                                 sk_SkPoint *c_pos, sk_SkFont *c_font,
                                 int c_encoding)
  {
    auto ret = SkTextBlob::MakeFromPosText (
                   reinterpret_cast<void *> (c_text), c_byteLength,
                   (SkPoint *)c_pos, *reinterpret_cast<SkFont *> (c_font),
                   SkTextEncoding (c_encoding))
                   .release ();
    return reinterpret_cast<sk_SkTextBlob *> (ret);
  }

  bool
  misk_Typeface_Equal (sk_SkTypeface *c_facea, sk_SkTypeface *c_faceb)
  {
    auto ret = SkTypeface::Equal (reinterpret_cast<SkTypeface *> (c_facea),
                                  reinterpret_cast<SkTypeface *> (c_faceb));
    return ret;
  }

  sk_SkTypeface *
  misk_Typeface_MakeEmpty ()
  {
    auto ret = SkTypeface::MakeEmpty ().release ();
    return reinterpret_cast<sk_SkTypeface *> (ret);
  }

  const sk_GrGLInterface *
  misk_GrGLMakeNativeInterface ()
  {
    auto ret = GrGLMakeNativeInterface ().release ();
    return reinterpret_cast<const sk_GrGLInterface *> (ret);
  }

  sk_GrBackendRenderTarget
  misk_GrBackendRenderTargetsMakeGL (int c_width, int c_height,
                                     int c_sampleCnt, int c_stencilBits,
                                     sk_GrGLFramebufferInfo c_glInfo)
  {
    auto ret = GrBackendRenderTargets::MakeGL (
        c_width, c_height, c_sampleCnt, c_stencilBits,
        *reinterpret_cast<GrGLFramebufferInfo *> (&c_glInfo));
    return *(reinterpret_cast<sk_GrBackendRenderTarget *> (&ret));
  }

  sk_SkSurface *
  misk_SkSurfacesWrapBackendRenderTarget (
      sk_GrRecordingContext *c_context,
      sk_GrBackendRenderTarget *c_backendRenderTarget, int c_origin,
      int c_colorType, sk_SkColorSpace *c_colorSpace,
      sk_SkSurfaceProps *c_surfaceProps)
  {
    auto ret = SkSurfaces::WrapBackendRenderTarget (
                   reinterpret_cast<GrRecordingContext *> (c_context),
                   *reinterpret_cast<GrBackendRenderTarget *> (
                       c_backendRenderTarget),
                   GrSurfaceOrigin (c_origin), SkColorType (c_colorType),
                   sk_ref_sp (reinterpret_cast<SkColorSpace *> (c_colorSpace)),
                   reinterpret_cast<SkSurfaceProps *> (c_surfaceProps))
                   .release ();
    return reinterpret_cast<sk_SkSurface *> (ret);
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterfaceOptions (sk_GrGLInterface *c_p0,
                                               sk_GrContextOptions c_p1)
  {
    auto ret = GrDirectContexts::MakeGL (
                   sk_ref_sp (reinterpret_cast<GrGLInterface *> (c_p0)),
                   *reinterpret_cast<GrContextOptions *> (&c_p1))
                   .release ();
    return reinterpret_cast<sk_GrDirectContext *> (ret);
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLInterface (sk_GrGLInterface *c_p0)
  {
    auto ret = GrDirectContexts::MakeGL (
                   sk_ref_sp (reinterpret_cast<GrGLInterface *> (c_p0)))
                   .release ();
    return reinterpret_cast<sk_GrDirectContext *> (ret);
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGLOptions (sk_GrContextOptions c_p0)
  {
    auto ret = GrDirectContexts::MakeGL (
                   *reinterpret_cast<GrContextOptions *> (&c_p0))
                   .release ();
    return reinterpret_cast<sk_GrDirectContext *> (ret);
  }

  sk_GrDirectContext *
  misk_GrDirectContextsMakeGL ()
  {
    auto ret = GrDirectContexts::MakeGL ().release ();
    return reinterpret_cast<sk_GrDirectContext *> (ret);
  }

  uint
  misk_SkColorSetARGB (uint c_a, uint c_r, uint c_g, uint c_b)
  {
    auto ret = SkColorSetARGB (c_a, c_r, c_g, c_b);
    return ret;
  }

  uint
  misk_SkColorSetA (uint c_c, uint c_a)
  {
    auto ret = SkColorSetA (SkColor (c_c), c_a);
    return ret;
  }

  uint
  misk_SkPreMultiplyARGB (uint c_a, uint c_r, uint c_g, uint c_b)
  {
    auto ret = SkPreMultiplyARGB (c_a, c_r, c_g, c_b);
    return ret;
  }

  uint
  misk_SkPreMultiplyColor (uint c_c)
  {
    auto ret = SkPreMultiplyColor (SkColor (c_c));
    return ret;
  }

  uchar sk_SK_AlphaOPAQUE = SK_AlphaOPAQUE;
  uchar sk_SK_AlphaTRANSPARENT = SK_AlphaTRANSPARENT;
  uint sk_SK_ColorBLACK = SK_ColorBLACK;
  uint sk_SK_ColorBLUE = SK_ColorBLUE;
  uint sk_SK_ColorCYAN = SK_ColorCYAN;
  uint sk_SK_ColorDKGRAY = SK_ColorDKGRAY;
  uint sk_SK_ColorGRAY = SK_ColorGRAY;
  uint sk_SK_ColorGREEN = SK_ColorGREEN;
  uint sk_SK_ColorLTGRAY = SK_ColorLTGRAY;
  uint sk_SK_ColorMAGENTA = SK_ColorMAGENTA;
  uint sk_SK_ColorRED = SK_ColorRED;
  uint sk_SK_ColorTRANSPARENT = SK_ColorTRANSPARENT;
  uint sk_SK_ColorWHITE = SK_ColorWHITE;
  uint sk_SK_ColorYELLOW = SK_ColorYELLOW;
  uint sk_SK_InvalidGenID = SK_InvalidGenID;
  uint sk_SK_InvalidUniqueID = SK_InvalidUniqueID;
  uint sk_SK_MSecMax = SK_MSecMax;
  uint sk_kAll_GrBackendState = kAll_GrBackendState;
  int sk_kGrGLStandardCnt = kGrGLStandardCnt;
  int sk_kSkColorTypeCnt = kSkColorTypeCnt;
  int sk_kSkFilterModeCount = kSkFilterModeCount;
  int sk_kSkMipmapModeCount = kSkMipmapModeCount;
  int sk_kSkStrAppendS32_MaxSize = kSkStrAppendS32_MaxSize;
  int sk_kSkStrAppendS64_MaxSize = kSkStrAppendS64_MaxSize;
  int sk_kSkStrAppendScalar_MaxSize = kSkStrAppendScalar_MaxSize;
  int sk_kSkStrAppendU32_MaxSize = kSkStrAppendU32_MaxSize;
  int sk_kSkStrAppendU64_MaxSize = kSkStrAppendU64_MaxSize;

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
