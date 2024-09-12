// This is a generated file. DO NOT EDIT.

package skia

// #cgo CXXFLAGS: -I ${SRCDIR}/_skia/skia
// #cgo LDFLAGS: -L ${SRCDIR}/_skia/build
// #cgo LDFLAGS: -l skia
// #cgo LDFLAGS: -l skshaper
// #cgo LDFLAGS: -l svg
// #cgo pkg-config: freetype2
// #cgo pkg-config: gl
//
// #include "api.h"
import "C"

import (
	"unsafe"
)

type GrGLFramebufferInfo C.sk_GrGLFramebufferInfo

type GrBackendRenderTarget struct {
	sk *C.sk_GrBackendRenderTarget
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrBackendRenderTarget has not been created.
func (o GrBackendRenderTarget) IsNil() bool {
	return o.sk == nil
}

func NewGrBackendRenderTarget() GrBackendRenderTarget {

	retC := C.misk_new_GrBackendRenderTarget()
	return GrBackendRenderTarget{sk: retC}
}

func NewGrBackendRenderTargetCopy(that GrBackendRenderTarget) GrBackendRenderTarget {
	c_that := that.sk
	retC := C.misk_new_GrBackendRenderTargetCopy(c_that)
	return GrBackendRenderTarget{sk: retC}
}

func (o GrBackendRenderTarget) Delete() {
	C.misk_delete_GrBackendRenderTarget(o.sk)
}

type GrDirectContext struct {
	sk *C.sk_GrDirectContext
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrDirectContext has not been created.
func (o GrDirectContext) IsNil() bool {
	return o.sk == nil
}

func (o GrDirectContext) Delete() {
	C.misk_delete_GrDirectContext(o.sk)
}

/*
GrContext uses the following interface to make all calls into OpenGL. When a
GrContext is created it is given a GrGLInterface. The interface's function
pointers must be valid for the OpenGL context associated with the GrContext.
On some platforms, such as Windows, function pointers for OpenGL extensions
may vary between OpenGL contexts. So the caller must be careful to use a
GrGLInterface initialized for the correct context. All functions that should
be available based on the OpenGL's version and extension string must be
non-NULL or GrContext creation will fail. This can be tested with the
validate() method when the OpenGL context has been made current.
*/
type GrGLInterface struct {
	sk *C.sk_GrGLInterface
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrGLInterface has not been created.
func (o GrGLInterface) IsNil() bool {
	return o.sk == nil
}

type GrContextOptions C.sk_GrContextOptions

func NewGrContextOptions() GrContextOptions {

	retC := C.misk_new_GrContextOptions()
	return *(*GrContextOptions)(unsafe.Pointer(&retC))
}

type GrContextOptionsEnable int64

const (
	/*
	   Forces an option to be disabled.
	*/
	GrContextOptionsEnableNo GrContextOptionsEnable = 0
	/*
	   Forces an option to be enabled.
	*/
	GrContextOptionsEnableYes GrContextOptionsEnable = 1
	/*
	   Uses Skia's default behavior, which may use runtime properties (e.g. driver version).
	*/
	GrContextOptionsEnableDefault GrContextOptionsEnable = 2
)

type GrContextOptionsShaderCacheStrategy int64

const (
	GrContextOptionsShaderCacheStrategySkSL          GrContextOptionsShaderCacheStrategy = 0
	GrContextOptionsShaderCacheStrategyBackendSource GrContextOptionsShaderCacheStrategy = 1
	GrContextOptionsShaderCacheStrategyBackendBinary GrContextOptionsShaderCacheStrategy = 2
)

/*
SkBitmap describes a two-dimensional raster pixel array. SkBitmap is built on
SkImageInfo, containing integer width and height, SkColorType and SkAlphaType
describing the pixel format, and SkColorSpace describing the range of colors.
SkBitmap points to SkPixelRef, which describes the physical array of pixels.
SkImageInfo bounds may be located anywhere fully inside SkPixelRef bounds.

SkBitmap can be drawn using SkCanvas. SkBitmap can be a drawing destination for SkCanvas
draw member functions. SkBitmap flexibility as a pixel container limits some
optimizations available to the target platform.

If pixel array is primarily read-only, use SkImage for better performance.
If pixel array is primarily written to, use SkSurface for better performance.

Declaring SkBitmap const prevents altering SkImageInfo: the SkBitmap height, width,
and so on cannot change. It does not affect SkPixelRef: a caller may write its
pixels. Declaring SkBitmap const affects SkBitmap configuration, not its contents.

SkBitmap is not thread safe. Each thread must have its own copy of SkBitmap fields,
although threads may share the underlying pixel array.
*/
type Bitmap struct {
	sk *C.sk_SkBitmap
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Bitmap has not been created.
func (o Bitmap) IsNil() bool {
	return o.sk == nil
}

/*
Creates an empty SkBitmap without pixels, with kUnknown_SkColorType,
kUnknown_SkAlphaType, and with a width and height of zero. SkPixelRef origin is
set to (0, 0).

Use setInfo() to associate SkColorType, SkAlphaType, width, and height
after SkBitmap has been created.

@return  empty SkBitmap

example: https://fiddle.skia.org/c/@Bitmap_empty_constructor
*/
func NewBitmap() Bitmap {

	retC := C.misk_new_Bitmap()
	return Bitmap{sk: retC}
}

/*
Copies settings from src to returned SkBitmap. Shares pixels if src has pixels
allocated, so both bitmaps reference the same pixels.

@param src  SkBitmap to copy SkImageInfo, and share SkPixelRef
@return     copy of src

example: https://fiddle.skia.org/c/@Bitmap_copy_const_SkBitmap
*/
func NewBitmapCopy(src Bitmap) Bitmap {
	c_src := src.sk
	retC := C.misk_new_BitmapCopy(c_src)
	return Bitmap{sk: retC}
}

/*
Decrements SkPixelRef reference count, if SkPixelRef is not nullptr.
*/
func (o Bitmap) Delete() {
	C.misk_delete_SkBitmap(o.sk)
}

/*
Returns true if all pixels are opaque. SkColorType determines how pixels
are encoded, and whether pixel describes alpha. Returns true for SkColorType
without alpha in each pixel; for other SkColorType, returns true if all
pixels have alpha values equivalent to 1.0 or greater.

For SkColorType kRGB_565_SkColorType or kGray_8_SkColorType: always
returns true. For SkColorType kAlpha_8_SkColorType, kBGRA_8888_SkColorType,
kRGBA_8888_SkColorType: returns true if all pixel alpha values are 255.
For SkColorType kARGB_4444_SkColorType: returns true if all pixel alpha values are 15.
For kRGBA_F16_SkColorType: returns true if all pixel alpha values are 1.0 or
greater.

Returns false for kUnknown_SkColorType.

@param bm  SkBitmap to check
@return    true if all pixels have opaque values or SkColorType is opaque
*/
func BitmapComputeIsOpaque(bm Bitmap) bool {
	c_bm := bm.sk
	retC := C.misk_Bitmap_ComputeIsOpaque(c_bm)
	return bool(retC)
}

/*
SkCanvas provides an interface for drawing, and how the drawing is clipped and transformed.
SkCanvas contains a stack of SkMatrix and clip values.

SkCanvas and SkPaint together provide the state to draw into SkSurface or SkDevice.
Each SkCanvas draw call transforms the geometry of the object by the concatenation of all
SkMatrix values in the stack. The transformed geometry is clipped by the intersection
of all of clip values in the stack. The SkCanvas draw calls use SkPaint to supply drawing
state such as color, SkTypeface, text size, stroke width, SkShader and so on.

To draw to a pixel-based destination, create raster surface or GPU surface.
Request SkCanvas from SkSurface to obtain the interface to draw.
SkCanvas generated by raster surface draws to memory visible to the CPU.
SkCanvas generated by GPU surface uses Vulkan or OpenGL to draw to the GPU.

To draw to a document, obtain SkCanvas from SVG canvas, document PDF, or SkPictureRecorder.
SkDocument based SkCanvas and other SkCanvas subclasses reference SkDevice describing the
destination.

SkCanvas can be constructed to draw to SkBitmap without first creating raster surface.
This approach may be deprecated in the future.
*/
type Canvas struct {
	sk *C.sk_SkCanvas
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Canvas has not been created.
func (o Canvas) IsNil() bool {
	return o.sk == nil
}

/*
Creates an empty SkCanvas with no backing device or pixels, with
a width and height of zero.

@return  empty SkCanvas

example: https://fiddle.skia.org/c/@Canvas_empty_constructor
*/
func NewCanvas() Canvas {

	retC := C.misk_new_Canvas()
	return Canvas{sk: retC}
}

/*
Creates SkCanvas of the specified dimensions without a SkSurface.
Used by subclasses with custom implementations for draw member functions.

If props equals nullptr, SkSurfaceProps are created with
SkSurfaceProps::InitType settings, which choose the pixel striping
direction and order. Since a platform may dynamically change its direction when
the device is rotated, and since a platform may have multiple monitors with
different characteristics, it is best not to rely on this legacy behavior.

@param width   zero or greater
@param height  zero or greater
@param props   LCD striping orientation and setting for device independent fonts;
may be nullptr
@return        SkCanvas placeholder with dimensions

example: https://fiddle.skia.org/c/@Canvas_int_int_const_SkSurfaceProps_star
*/
func NewCanvasWithDimensions(width int, height int, props SurfaceProps) Canvas {
	c_width := C.int(width)
	c_height := C.int(height)
	c_props := props.sk
	retC := C.misk_new_CanvasWithDimensions(c_width, c_height, c_props)
	return Canvas{sk: retC}
}

/*
Constructs a canvas that draws into bitmap.
Sets kUnknown_SkPixelGeometry in constructed SkSurface.

SkBitmap is copied so that subsequently editing bitmap will not affect
constructed SkCanvas.

May be deprecated in the future.

@param bitmap  width, height, SkColorType, SkAlphaType, and pixel
storage of raster surface
@return        SkCanvas that can be used to draw into bitmap

example: https://fiddle.skia.org/c/@Canvas_copy_const_SkBitmap
*/
func NewCanvasFromBitmap(bitmap Bitmap) Canvas {
	c_bitmap := bitmap.sk
	retC := C.misk_new_CanvasFromBitmap(c_bitmap)
	return Canvas{sk: retC}
}

/*
Constructs a canvas that draws into bitmap.
Use props to match the device characteristics, like LCD striping.

bitmap is copied so that subsequently editing bitmap will not affect
constructed SkCanvas.

@param bitmap  width, height, SkColorType, SkAlphaType,
and pixel storage of raster surface
@param props   order and orientation of RGB striping; and whether to use
device independent fonts
@return        SkCanvas that can be used to draw into bitmap

example: https://fiddle.skia.org/c/@Canvas_const_SkBitmap_const_SkSurfaceProps
*/
func NewCanvasFromBitmapSurfaceProps(bitmap Bitmap, props SurfaceProps) Canvas {
	c_bitmap := bitmap.sk
	c_props := props.sk
	retC := C.misk_new_CanvasFromBitmapSurfaceProps(c_bitmap, c_props)
	return Canvas{sk: retC}
}

/*
Draws saved layers, if any.
Frees up resources used by SkCanvas.

example: https://fiddle.skia.org/c/@Canvas_destructor
*/
func (o Canvas) Delete() {
	C.misk_delete_SkCanvas(o.sk)
}

type CanvasClipEdgeStyle int64

const (
	CanvasClipEdgeStyleHard CanvasClipEdgeStyle = 0
	CanvasClipEdgeStyleSoft CanvasClipEdgeStyle = 1
)

/*
Selects if an array of points are drawn as discrete points, as lines, or as
an open polygon.
*/
type CanvasPointMode int64

const (
	// draw each point separately
	CanvasPointModePoints CanvasPointMode = 0
	// draw each pair of points as a line segment
	CanvasPointModeLines CanvasPointMode = 1
	// draw the array of points as a open polygon
	CanvasPointModePolygon CanvasPointMode = 2
)

/*
Experimental. Controls anti-aliasing of each edge of images in an image-set.
*/
type CanvasQuadAAFlags int64

const (
	CanvasQuadAAFlagsLeft_QuadAAFlag   CanvasQuadAAFlags = 1
	CanvasQuadAAFlagsTop_QuadAAFlag    CanvasQuadAAFlags = 2
	CanvasQuadAAFlagsRight_QuadAAFlag  CanvasQuadAAFlags = 4
	CanvasQuadAAFlagsBottom_QuadAAFlag CanvasQuadAAFlags = 8
	CanvasQuadAAFlagsNone              CanvasQuadAAFlags = 0
	CanvasQuadAAFlagsAll               CanvasQuadAAFlags = 15
)

/*
SaveLayerFlags provides options that may be used in any combination in SaveLayerRec,
defining how layer allocated by saveLayer() operates. It may be set to zero,
kPreserveLCDText_SaveLayerFlag, kInitWithPrevious_SaveLayerFlag, or both flags.
*/
type CanvasSaveLayerFlagsSet int64

const (
	CanvasSaveLayerFlagsSetPreserveLCDText_SaveLayerFlag CanvasSaveLayerFlagsSet = 2
	// initializes with previous contents
	CanvasSaveLayerFlagsSetInitWithPrevious_SaveLayerFlag CanvasSaveLayerFlagsSet = 4
	CanvasSaveLayerFlagsSetF16ColorType                   CanvasSaveLayerFlagsSet = 16
)

type CanvasSaveLayerStrategy int64

const (
	CanvasSaveLayerStrategyFullLayer CanvasSaveLayerStrategy = 0
	CanvasSaveLayerStrategyNoLayer   CanvasSaveLayerStrategy = 1
)

/*
SrcRectConstraint controls the behavior at the edge of source SkRect,
provided to drawImageRect() when there is any filtering. If kStrict is set,
then extra code is used to ensure it never samples outside of the src-rect.
kStrict_SrcRectConstraint disables the use of mipmaps and anisotropic filtering.
*/
type CanvasSrcRectConstraint int64

const (
	// sample only inside bounds; slower
	CanvasSrcRectConstraintStrict CanvasSrcRectConstraint = 0
	// sample outside bounds; faster
	CanvasSrcRectConstraintFast CanvasSrcRectConstraint = 1
)

type ColorSpace struct {
	sk *C.sk_SkColorSpace
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the ColorSpace has not been created.
func (o ColorSpace) IsNil() bool {
	return o.sk == nil
}

/*
Create the sRGB color space.
*/
func ColorSpaceMakeSRGB() ColorSpace {
	retC := C.misk_ColorSpace_MakeSRGB()
	return ColorSpace{sk: retC}
}

/*
Colorspace with the sRGB primaries, but a linear (1.0) gamma.
*/
func ColorSpaceMakeSRGBLinear() ColorSpace {
	retC := C.misk_ColorSpace_MakeSRGBLinear()
	return ColorSpace{sk: retC}
}

/*
If both are null, we return true. If one is null and the other is not, we return false.
If both are non-null, we do a deeper compare.
*/
func ColorSpaceEquals(p0 ColorSpace, p1 ColorSpace) bool {
	c_p0 := p0.sk
	c_p1 := p1.sk
	retC := C.misk_ColorSpace_Equals(c_p0, c_p1)
	return bool(retC)
}

/*
SkFont controls options applied when drawing and measuring text.
*/
type Font struct {
	sk *C.sk_SkFont
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Font has not been created.
func (o Font) IsNil() bool {
	return o.sk == nil
}

/*
Constructs SkFont with default values.

@return  default initialized SkFont
*/
func NewFont() Font {

	retC := C.misk_new_Font()
	return Font{sk: retC}
}

/*
Constructs SkFont with default values with SkTypeface and size in points.

@param typeface  font and style used to draw and measure text
@param size      typographic height of text
@return          initialized SkFont
*/
func NewFontTypefaceSize(typeface Typeface, size float32) Font {
	c_typeface := typeface.sk
	c_size := C.float(size)
	retC := C.misk_new_FontTypefaceSize(c_typeface, c_size)
	return Font{sk: retC}
}

/*
Constructs SkFont with default values with SkTypeface.

@param typeface  font and style used to draw and measure text
@return          initialized SkFont
*/
func NewFontTypeface(typeface Typeface) Font {
	c_typeface := typeface.sk
	retC := C.misk_new_FontTypeface(c_typeface)
	return Font{sk: retC}
}

/*
Constructs SkFont with default values with SkTypeface and size in points,
horizontal scale, and horizontal skew. Horizontal scale emulates condensed
and expanded fonts. Horizontal skew emulates oblique fonts.

@param typeface  font and style used to draw and measure text
@param size      typographic height of text
@param scaleX    text horizontal scale
@param skewX     additional shear on x-axis relative to y-axis
@return          initialized SkFont
*/
func NewFontTypefaceSizeScaleSkew(typeface Typeface, size float32, scaleX float32, skewX float32) Font {
	c_typeface := typeface.sk
	c_size := C.float(size)
	c_scaleX := C.float(scaleX)
	c_skewX := C.float(skewX)
	retC := C.misk_new_FontTypefaceSizeScaleSkew(c_typeface, c_size, c_scaleX, c_skewX)
	return Font{sk: retC}
}

type FontMgr struct {
	sk *C.sk_SkFontMgr
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontMgr has not been created.
func (o FontMgr) IsNil() bool {
	return o.sk == nil
}

type FontStyle struct {
	sk *C.sk_SkFontStyle
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontStyle has not been created.
func (o FontStyle) IsNil() bool {
	return o.sk == nil
}

func NewFontStyle2(weight int, width int, slant FontStyleSlant) FontStyle {
	c_weight := C.int(weight)
	c_width := C.int(width)
	c_slant := C.uint(slant)
	retC := C.misk_new_FontStyle2(c_weight, c_width, c_slant)
	return FontStyle{sk: retC}
}

func NewFontStyle() FontStyle {

	retC := C.misk_new_FontStyle()
	return FontStyle{sk: retC}
}

func FontStyleNormal() FontStyle {
	retC := C.misk_FontStyle_Normal()
	return FontStyle{sk: &retC}
}

func FontStyleBold() FontStyle {
	retC := C.misk_FontStyle_Bold()
	return FontStyle{sk: &retC}
}

func FontStyleItalic() FontStyle {
	retC := C.misk_FontStyle_Italic()
	return FontStyle{sk: &retC}
}

func FontStyleBoldItalic() FontStyle {
	retC := C.misk_FontStyle_BoldItalic()
	return FontStyle{sk: &retC}
}

type FontStyleSlant int64

const (
	FontStyleSlantUpright FontStyleSlant = 0
	FontStyleSlantItalic  FontStyleSlant = 1
	FontStyleSlantOblique FontStyleSlant = 2
)

/*
SkIRect holds four 32-bit integer coordinates describing the upper and
lower bounds of a rectangle. SkIRect may be created from outer bounds or
from position, width, and height. SkIRect describes an area; if its right
is less than or equal to its left, or if its bottom is less than or equal to
its top, it is considered empty.
*/
type IRect C.sk_SkIRect

/*
Returns constructed SkIRect set to (0, 0, 0, 0).
Many other rectangles are empty; if left is equal to or greater than right,
or if top is equal to or greater than bottom. Setting all members to zero
is a convenience, but does not designate a special empty rectangle.

@return  bounds (0, 0, 0, 0)
*/
func IRectMakeEmpty() IRect {
	retC := C.misk_IRect_MakeEmpty()
	return IRect(retC)
}

/*
Returns constructed SkIRect set to (0, 0, w, h). Does not validate input; w or h
may be negative.

@param w  width of constructed SkIRect
@param h  height of constructed SkIRect
@return   bounds (0, 0, w, h)
*/
func IRectMakeWH(w int, h int) IRect {
	c_w := C.int(w)
	c_h := C.int(h)
	retC := C.misk_IRect_MakeWH(c_w, c_h)
	return IRect(retC)
}

/*
Returns constructed SkIRect set to (l, t, r, b). Does not sort input; SkIRect may
result in fLeft greater than fRight, or fTop greater than fBottom.

@param l  integer stored in fLeft
@param t  integer stored in fTop
@param r  integer stored in fRight
@param b  integer stored in fBottom
@return   bounds (l, t, r, b)
*/
func IRectMakeLTRB(l int, t int, r int, b int) IRect {
	c_l := C.int(l)
	c_t := C.int(t)
	c_r := C.int(r)
	c_b := C.int(b)
	retC := C.misk_IRect_MakeLTRB(c_l, c_t, c_r, c_b)
	return IRect(retC)
}

/*
Returns constructed SkIRect set to: (x, y, x + w, y + h).
Does not validate input; w or h may be negative.

@param x  stored in fLeft
@param y  stored in fTop
@param w  added to x and stored in fRight
@param h  added to y and stored in fBottom
@return   bounds at (x, y) with width w and height h
*/
func IRectMakeXYWH(x int, y int, w int, h int) IRect {
	c_x := C.int(x)
	c_y := C.int(y)
	c_w := C.int(w)
	c_h := C.int(h)
	retC := C.misk_IRect_MakeXYWH(c_x, c_y, c_w, c_h)
	return IRect(retC)
}

/*
Returns true if a intersects b.
Returns false if either a or b is empty, or do not intersect.

@param a  SkIRect to intersect
@param b  SkIRect to intersect
@return   true if a and b have area in common
*/
func IRectIntersects(a IRect, b IRect) bool {
	c_a := *(*C.sk_SkIRect)(unsafe.Pointer(&a))
	c_b := *(*C.sk_SkIRect)(unsafe.Pointer(&b))
	retC := C.misk_IRect_Intersects(c_a, c_b)
	return bool(retC)
}

type ISize C.sk_SkISize

/*
SkPaint controls options applied when drawing. SkPaint collects all
options outside of the SkCanvas clip and SkCanvas matrix.

Various options apply to strokes and fills, and images.

SkPaint collects effects and filters that describe single-pass and multiple-pass
algorithms that alter the drawing geometry, color, and transparency. For instance,
SkPaint does not directly implement dashing or blur, but contains the objects that do so.
*/
type Paint struct {
	sk *C.sk_SkPaint
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Paint has not been created.
func (o Paint) IsNil() bool {
	return o.sk == nil
}

/*
Constructs SkPaint with default values.

@return  default initialized SkPaint

example: https://fiddle.skia.org/c/@Paint_empty_constructor
*/
func NewPaint() Paint {

	retC := C.misk_new_Paint()
	return Paint{sk: retC}
}

/*
Makes a shallow copy of SkPaint. SkPathEffect, SkShader,
SkMaskFilter, SkColorFilter, and SkImageFilter are shared
between the original paint and the copy. Objects containing SkRefCnt increment
their references by one.

The referenced objects SkPathEffect, SkShader, SkMaskFilter, SkColorFilter,
and SkImageFilter cannot be modified after they are created.
This prevents objects with SkRefCnt from being modified once SkPaint refers to them.

@param paint  original to copy
@return       shallow copy of paint

example: https://fiddle.skia.org/c/@Paint_copy_const_SkPaint
*/
func NewPaintCopy(paint Paint) Paint {
	c_paint := paint.sk
	retC := C.misk_new_PaintCopy(c_paint)
	return Paint{sk: retC}
}

/*
Decreases SkPaint SkRefCnt of owned objects: SkPathEffect, SkShader,
SkMaskFilter, SkColorFilter, and SkImageFilter. If the
objects containing SkRefCnt go to zero, they are deleted.
*/
func (o Paint) Delete() {
	C.misk_delete_SkPaint(o.sk)
}

/*
Sets all SkPaint contents to their initial values. This is equivalent to replacing
SkPaint with the result of SkPaint().

example: https://fiddle.skia.org/c/@Paint_reset
*/
func (o Paint) Reset() {
	c_obj := o.sk
	C.misk_Paint_reset(c_obj)
}

func (o Paint) GetAlpha() uint {
	c_obj := o.sk
	retC := C.misk_Paint_getAlpha(c_obj)
	return uint(retC)
}

func (o Paint) SetAlpha(a uint) {
	c_obj := o.sk
	c_a := C.uint(a)
	C.misk_Paint_setAlpha(c_obj, c_a)
}

/*
Sets color used when drawing solid fills. The color components range from 0 to 255.
The color is unpremultiplied; alpha sets the transparency independent of RGB.

@param a  amount of alpha, from fully transparent (0) to fully opaque (255)
@param r  amount of red, from no red (0) to full red (255)
@param g  amount of green, from no green (0) to full green (255)
@param b  amount of blue, from no blue (0) to full blue (255)

example: https://fiddle.skia.org/c/@Paint_setARGB
*/
func (o Paint) SetARGB(a uint, r uint, g uint, b uint) {
	c_obj := o.sk
	c_a := C.uint(a)
	c_r := C.uint(r)
	c_g := C.uint(g)
	c_b := C.uint(b)
	C.misk_Paint_setARGB(c_obj, c_a, c_r, c_g, c_b)
}

/*
Requests, but does not require, that edge pixels draw opaque or with
partial transparency.
@param aa  setting for antialiasing
*/
func (o Paint) SetAntiAlias(aa bool) {
	c_obj := o.sk
	c_aa := C.bool(aa)
	C.misk_Paint_setAntiAlias(c_obj, c_aa)
}

/*
Helper method for calling setBlender().

	This sets a blender that implements the specified blendmode enum.
*/
func (o Paint) SetBlendMode(mode BlendMode) {
	c_obj := o.sk
	c_mode := C.int(mode)
	C.misk_Paint_setBlendMode(c_obj, c_mode)
}

/*
Sets alpha and RGB used when stroking and filling. The color is a 32-bit value,
unpremultiplied, packing 8-bit components for alpha, red, blue, and green.

@param color  unpremultiplied ARGB

example: https://fiddle.skia.org/c/@Paint_setColor
*/
func (o Paint) SetColor(color Color) {
	c_obj := o.sk
	c_color := C.uint(color)
	C.misk_Paint_setColor(c_obj, c_color)
}

/*
Requests, but does not require, to distribute color error.
@param dither  setting for ditering
*/
func (o Paint) SetDither(dither bool) {
	c_obj := o.sk
	c_dither := C.bool(dither)
	C.misk_Paint_setDither(c_obj, c_dither)
}

/*
Returns the geometry drawn at the beginning and end of strokes.
*/
func (o Paint) GetStrokeCap() PaintCap {
	c_obj := o.sk
	retC := C.misk_Paint_getStrokeCap(c_obj)
	return PaintCap(retC)
}

/*
Sets the geometry drawn at the beginning and end of strokes.

example: https://fiddle.skia.org/c/@Paint_setStrokeCap_a
example: https://fiddle.skia.org/c/@Paint_setStrokeCap_b
*/
func (o Paint) SetStrokeCap(cap PaintCap) {
	c_obj := o.sk
	c_cap := C.uint(cap)
	C.misk_Paint_setStrokeCap(c_obj, c_cap)
}

/*
Returns the geometry drawn at the corners of strokes.
*/
func (o Paint) GetStrokeJoin() PaintJoin {
	c_obj := o.sk
	retC := C.misk_Paint_getStrokeJoin(c_obj)
	return PaintJoin(retC)
}

/*
Sets the geometry drawn at the corners of strokes.

example: https://fiddle.skia.org/c/@Paint_setStrokeJoin
*/
func (o Paint) SetStrokeJoin(join PaintJoin) {
	c_obj := o.sk
	c_join := C.uint(join)
	C.misk_Paint_setStrokeJoin(c_obj, c_join)
}

/*
Returns the limit at which a sharp corner is drawn beveled.

@return  zero and greater miter limit
*/
func (o Paint) GetStrokeMiter() float32 {
	c_obj := o.sk
	retC := C.misk_Paint_getStrokeMiter(c_obj)
	return float32(retC)
}

/*
Sets the limit at which a sharp corner is drawn beveled.
Valid values are zero and greater.
Has no effect if miter is less than zero.

@param miter  zero and greater miter limit

example: https://fiddle.skia.org/c/@Paint_setStrokeMiter
*/
func (o Paint) SetStrokeMiter(miter float32) {
	c_obj := o.sk
	c_miter := C.float(miter)
	C.misk_Paint_setStrokeMiter(c_obj, c_miter)
}

/*
Returns the thickness of the pen used by SkPaint to
outline the shape.

@return  zero for hairline, greater than zero for pen thickness
*/
func (o Paint) GetStrokeWidth() float32 {
	c_obj := o.sk
	retC := C.misk_Paint_getStrokeWidth(c_obj)
	return float32(retC)
}

/*
Sets the thickness of the pen used by the paint to outline the shape.
A stroke-width of zero is treated as "hairline" width. Hairlines are always exactly one
pixel wide in device space (their thickness does not change as the canvas is scaled).
Negative stroke-widths are invalid; setting a negative width will have no effect.

@param width  zero thickness for hairline; greater than zero for pen thickness

example: https://fiddle.skia.org/c/@Miter_Limit
example: https://fiddle.skia.org/c/@Paint_setStrokeWidth
*/
func (o Paint) SetStrokeWidth(width float32) {
	c_obj := o.sk
	c_width := C.float(width)
	C.misk_Paint_setStrokeWidth(c_obj, c_width)
}

/*
Returns whether the geometry is filled, stroked, or filled and stroked.
*/
func (o Paint) GetStyle() PaintStyle {
	c_obj := o.sk
	retC := C.misk_Paint_getStyle(c_obj)
	return PaintStyle(retC)
}

/*
Sets whether the geometry is filled, stroked, or filled and stroked.
Has no effect if style is not a legal SkPaint::Style value.

example: https://fiddle.skia.org/c/@Paint_setStyle
example: https://fiddle.skia.org/c/@Stroke_Width
*/
func (o Paint) SetStyle(style PaintStyle) {
	c_obj := o.sk
	c_style := C.uint(style)
	C.misk_Paint_setStyle(c_obj, c_style)
}

/*
Cap draws at the beginning and end of an open path contour.
*/
type PaintCap int64

const (
	// no stroke extension
	PaintCapButt PaintCap = 0
	// adds circle
	PaintCapRound PaintCap = 1
	// adds square
	PaintCapSquare PaintCap = 2
	// largest Cap value
	PaintCapLast PaintCap = 2
	// equivalent to kButt_Cap
	PaintCapDefault PaintCap = 0
)

/*
Join specifies how corners are drawn when a shape is stroked. Join
affects the four corners of a stroked rectangle, and the connected segments in a
stroked path.

Choose miter join to draw sharp corners. Choose round join to draw a circle with a
radius equal to the stroke width on top of the corner. Choose bevel join to minimally
connect the thick strokes.

The fill path constructed to describe the stroked path respects the join setting but may
not contain the actual join. For instance, a fill path constructed with round joins does
not necessarily include circles at each connected segment.
*/
type PaintJoin int64

const (
	// extends to miter limit
	PaintJoinMiter PaintJoin = 0
	// adds circle
	PaintJoinRound PaintJoin = 1
	// connects outside edges
	PaintJoinBevel PaintJoin = 2
	// equivalent to the largest value for Join
	PaintJoinLast PaintJoin = 2
	// equivalent to kMiter_Join
	PaintJoinDefault PaintJoin = 0
)

/*
Set Style to fill, stroke, or both fill and stroke geometry.
The stroke and fill
share all paint attributes; for instance, they are drawn with the same color.

Use kStrokeAndFill_Style to avoid hitting the same pixels twice with a stroke draw and
a fill draw.
*/
type PaintStyle int64

const (
	// set to fill geometry
	PaintStyleFill PaintStyle = 0
	// set to stroke geometry
	PaintStyleStroke PaintStyle = 1
	// sets to stroke and fill geometry
	PaintStyleStrokeAndFill PaintStyle = 2
)

/*
SkPath contain geometry. SkPath may be empty, or contain one or more verbs that
outline a figure. SkPath always starts with a move verb to a Cartesian coordinate,
and may be followed by additional verbs that add lines or curves.
Adding a close verb makes the geometry into a continuous loop, a closed contour.
SkPath may contain any number of contours, each beginning with a move verb.

SkPath contours may contain only a move verb, or may also contain lines,
quadratic beziers, conics, and cubic beziers. SkPath contours may be open or
closed.

When used to draw a filled area, SkPath describes whether the fill is inside or
outside the geometry. SkPath also describes the winding rule used to fill
overlapping contours.

Internally, SkPath lazily computes metrics likes bounds and convexity. Call
SkPath::updateBoundsCache to make SkPath thread safe.
*/
type Path struct {
	sk *C.sk_SkPath
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Path has not been created.
func (o Path) IsNil() bool {
	return o.sk == nil
}

/*
Constructs an empty SkPath. By default, SkPath has no verbs, no SkPoint, and no weights.
FillType is set to kWinding.

@return  empty SkPath

example: https://fiddle.skia.org/c/@Path_empty_constructor
*/
func NewPath() Path {

	retC := C.misk_new_Path()
	return Path{sk: retC}
}

/*
Constructs a copy of an existing path.
Copy constructor makes two paths identical by value. Internally, path and
the returned result share pointer values. The underlying verb array, SkPoint array
and weights are copied when modified.

Creating a SkPath copy is very efficient and never allocates memory.
SkPath are always copied by value from the interface; the underlying shared
pointers are not exposed.

@param path  SkPath to copy by value
@return      copy of SkPath

example: https://fiddle.skia.org/c/@Path_copy_const_SkPath
*/
func NewPathCopy(path Path) Path {
	c_path := path.sk
	retC := C.misk_new_PathCopy(c_path)
	return Path{sk: retC}
}

/*
Releases ownership of any shared data and deletes data if SkPath is sole owner.

example: https://fiddle.skia.org/c/@Path_destructor
*/
func (o Path) Delete() {
	C.misk_delete_SkPath(o.sk)
}

/*
AddPathMode chooses how addPath() appends. Adding one SkPath to another can extend
the last contour or start a new contour.
*/
type PathAddPathMode int64

const (
	/*
	   Contours are appended to the destination path as new contours.
	*/
	PathAddPathModeAppend PathAddPathMode = 0
	/*
	   Extends the last contour of the destination path with the first countour
	   of the source path, connecting them with a line.  If the last contour is
	   closed, a new empty contour starting at its start point is extended instead.
	   If the destination path is empty, the result is the source path.
	   The last path of the result is closed only if the last path of the source is.
	*/
	PathAddPathModeExtend PathAddPathMode = 1
)

/*
Four oval parts with radii (rx, ry) start at last SkPath SkPoint and ends at (x, y).
ArcSize and Direction select one of the four oval parts.
*/
type PathArcSize int64

const (
	// smaller of arc pair
	PathArcSizeSmall PathArcSize = 0
	// larger of arc pair
	PathArcSizeLarge PathArcSize = 1
)

/*
SegmentMask constants correspond to each drawing Verb type in SkPath; for
instance, if SkPath only contains lines, only the kLine_SegmentMask bit is set.
*/
type PathSegmentMask int64

const (
	PathSegmentMaskLine  PathSegmentMask = 1
	PathSegmentMaskQuad  PathSegmentMask = 2
	PathSegmentMaskConic PathSegmentMask = 4
	PathSegmentMaskCubic PathSegmentMask = 8
)

/*
Verb instructs SkPath how to interpret one or more SkPoint and optional conic weight;
manage contour, and terminate SkPath.
*/
type PathVerb int64

const (
	PathVerbMove  PathVerb = 0
	PathVerbLine  PathVerb = 1
	PathVerbQuad  PathVerb = 2
	PathVerbConic PathVerb = 3
	PathVerbCubic PathVerb = 4
	PathVerbClose PathVerb = 5
	PathVerbDone  PathVerb = 6
)

/*
SkRect holds four float coordinates describing the upper and
lower bounds of a rectangle. SkRect may be created from outer bounds or
from position, width, and height. SkRect describes an area; if its right
is less than or equal to its left, or if its bottom is less than or equal to
its top, it is considered empty.
*/
type Rect C.sk_SkRect

/*
Returns constructed SkRect set to (0, 0, 0, 0).
Many other rectangles are empty; if left is equal to or greater than right,
or if top is equal to or greater than bottom. Setting all members to zero
is a convenience, but does not designate a special empty rectangle.

@return  bounds (0, 0, 0, 0)
*/
func RectMakeEmpty() Rect {
	retC := C.misk_Rect_MakeEmpty()
	return Rect(retC)
}

/*
Returns constructed SkRect set to float values (0, 0, w, h). Does not
validate input; w or h may be negative.

Passing integer values may generate a compiler warning since SkRect cannot
represent 32-bit integers exactly. Use SkIRect for an exact integer rectangle.

@param w  float width of constructed SkRect
@param h  float height of constructed SkRect
@return   bounds (0, 0, w, h)
*/
func RectMakeWH(w float32, h float32) Rect {
	c_w := C.float(w)
	c_h := C.float(h)
	retC := C.misk_Rect_MakeWH(c_w, c_h)
	return Rect(retC)
}

/*
Returns constructed SkRect set to (0, 0, size.width(), size.height()). Does not
validate input; size.width() or size.height() may be negative.

@param size  float values for SkRect width and height
@return      bounds (0, 0, size.width(), size.height())
*/
func RectMakeSize(size Size) Rect {
	c_size := *(*C.sk_SkSize)(unsafe.Pointer(&size))
	retC := C.misk_Rect_MakeSize(c_size)
	return Rect(retC)
}

/*
Returns constructed SkRect set to (l, t, r, b). Does not sort input; SkRect may
result in fLeft greater than fRight, or fTop greater than fBottom.

@param l  float stored in fLeft
@param t  float stored in fTop
@param r  float stored in fRight
@param b  float stored in fBottom
@return   bounds (l, t, r, b)
*/
func RectMakeLTRB(l float32, t float32, r float32, b float32) Rect {
	c_l := C.float(l)
	c_t := C.float(t)
	c_r := C.float(r)
	c_b := C.float(b)
	retC := C.misk_Rect_MakeLTRB(c_l, c_t, c_r, c_b)
	return Rect(retC)
}

/*
Returns constructed SkRect set to (x, y, x + w, y + h).
Does not validate input; w or h may be negative.

@param x  stored in fLeft
@param y  stored in fTop
@param w  added to x and stored in fRight
@param h  added to y and stored in fBottom
@return   bounds at (x, y) with width w and height h
*/
func RectMakeXYWH(x float32, y float32, w float32, h float32) Rect {
	c_x := C.float(x)
	c_y := C.float(y)
	c_w := C.float(w)
	c_h := C.float(h)
	retC := C.misk_Rect_MakeXYWH(c_x, c_y, c_w, c_h)
	return Rect(retC)
}

/*
Returns constructed SkIRect set to (0, 0, size.width(), size.height()).
Does not validate input; size.width() or size.height() may be negative.

@param size  integer values for SkRect width and height
@return      bounds (0, 0, size.width(), size.height())
*/
func RectMakeISize(size ISize) Rect {
	c_size := *(*C.sk_SkISize)(unsafe.Pointer(&size))
	retC := C.misk_Rect_MakeISize(c_size)
	return Rect(retC)
}

/*
Returns constructed SkIRect set to irect, promoting integers to float.
Does not validate input; fLeft may be greater than fRight, fTop may be greater
than fBottom.

@param irect  integer unsorted bounds
@return       irect members converted to float
*/
func RectMakeIRect(irect IRect) Rect {
	c_irect := *(*C.sk_SkIRect)(unsafe.Pointer(&irect))
	retC := C.misk_Rect_MakeIRect(c_irect)
	return Rect(retC)
}

/*
Returns true if a intersects b.
Returns false if either a or b is empty, or do not intersect.

@param a  SkRect to intersect
@param b  SkRect to intersect
@return   true if a and b have area in common
*/
func RectIntersects(a Rect, b Rect) bool {
	c_a := *(*C.sk_SkRect)(unsafe.Pointer(&a))
	c_b := *(*C.sk_SkRect)(unsafe.Pointer(&b))
	retC := C.misk_Rect_Intersects(c_a, c_b)
	return bool(retC)
}

// /////////////////////////////////////////////////////////////////////////////
type Size C.sk_SkSize

/*
Describes properties and constraints of a given SkSurface. The rendering engine can parse these
during drawing, and can sometimes optimize its performance (e.g. disabling an expensive
feature).
*/
type SurfaceProps struct {
	sk *C.sk_SkSurfaceProps
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the SurfaceProps has not been created.
func (o SurfaceProps) IsNil() bool {
	return o.sk == nil
}

/*
No flags, unknown pixel geometry, platform-default contrast/gamma.
*/
func NewSurfaceProps() SurfaceProps {

	retC := C.misk_new_SurfaceProps()
	return SurfaceProps{sk: retC}
}

/*
TODO(kschmi): Remove this constructor and replace with the one below. *
*/
func NewSurfacePropsPixelGeometry(flags uint, p1 PixelGeometry) SurfaceProps {
	c_flags := C.uint(flags)
	c_p1 := C.uint(p1)
	retC := C.misk_new_SurfacePropsPixelGeometry(c_flags, c_p1)
	return SurfaceProps{sk: retC}
}

func NewSurfacePropsCopy(p0 SurfaceProps) SurfaceProps {
	c_p0 := p0.sk
	retC := C.misk_new_SurfacePropsCopy(c_p0)
	return SurfaceProps{sk: retC}
}

type SurfacePropsFlags int64

const (
	SurfacePropsFlagsDefault_Flag                   SurfacePropsFlags = 0
	SurfacePropsFlagsUseDeviceIndependentFonts_Flag SurfacePropsFlags = 1
	SurfacePropsFlagsDynamicMSAA_Flag               SurfacePropsFlags = 2
	SurfacePropsFlagsAlwaysDither_Flag              SurfacePropsFlags = 4
)

/*
The SkTypeface class specifies the typeface and intrinsic style of a font.
This is used in the paint, along with optionally algorithmic settings like
textSize, textSkewX, textScaleX, kFakeBoldText_Mask, to specify
how text appears when drawn (and measured).

Typeface objects are immutable, and so they can be shared between threads.
*/
type Typeface struct {
	sk *C.sk_SkTypeface
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Typeface has not been created.
func (o Typeface) IsNil() bool {
	return o.sk == nil
}

/*
Returns true if the two typefaces reference the same underlying font,
handling either being null (treating null as not equal to any font).
*/
func TypefaceEqual(facea Typeface, faceb Typeface) bool {
	c_facea := facea.sk
	c_faceb := faceb.sk
	retC := C.misk_Typeface_Equal(c_facea, c_faceb)
	return bool(retC)
}

/*
Returns a non-null typeface which contains no glyphs.
*/
func TypefaceMakeEmpty() Typeface {
	retC := C.misk_Typeface_MakeEmpty()
	return Typeface{sk: retC}
}

type BlendMode int64

const ()

/*
Description of how the LCD strips are arranged for each pixel. If this is unknown, or the
pixels are meant to be "portable" and/or transformed before showing (e.g. rotated, scaled)
then use kUnknown_SkPixelGeometry.
*/
type PixelGeometry int64

const (
	PixelGeometryUnknown PixelGeometry = 0
	PixelGeometryRGB_H   PixelGeometry = 1
	PixelGeometryBGR_H   PixelGeometry = 2
	PixelGeometryRGB_V   PixelGeometry = 3
	PixelGeometryBGR_V   PixelGeometry = 4
)

/*
Is the data protected on the GPU or not.
*/
type SkgpuProtected int64

const (
	SkgpuProtectedNo  SkgpuProtected = 0
	SkgpuProtectedYes SkgpuProtected = -1
)

/*
Rather than depend on platform-specific GL headers and libraries, we require
the client to provide a struct of GL function pointers. This struct can be
specified per-GrContext as a parameter to GrContext::MakeGL. If no interface is
passed to MakeGL then a default GL interface is created using GrGLMakeNativeInterface().
If this returns nullptr then GrContext::MakeGL() will fail.

The implementation of GrGLMakeNativeInterface is platform-specific. Several
implementations have been provided (for GLX, WGL, EGL, etc), along with an
implementation that simply returns nullptr. Clients should select the most
appropriate one to build.
*/
func GrGLMakeNativeInterface() GrGLInterface {
	retC := C.misk_GrGLMakeNativeInterface()
	return GrGLInterface{sk: retC}
}

func GrBackendRenderTargetsMakeGL(width int, height int, sampleCnt int, stencilBits int, glInfo GrGLFramebufferInfo) GrBackendRenderTarget {
	c_width := C.int(width)
	c_height := C.int(height)
	c_sampleCnt := C.int(sampleCnt)
	c_stencilBits := C.int(stencilBits)
	c_glInfo := *(*C.sk_GrGLFramebufferInfo)(unsafe.Pointer(&glInfo))
	retC := C.misk_GrBackendRenderTargetsMakeGL(c_width, c_height, c_sampleCnt, c_stencilBits, c_glInfo)
	return GrBackendRenderTarget{sk: &retC}
}

/*
Creates a GrDirectContext for a backend context. GrGLInterface must be non-null.
*/
func GrDirectContextsMakeGLInterfaceOptions(p0 GrGLInterface, p1 GrContextOptions) GrDirectContext {
	c_p0 := p0.sk
	c_p1 := *(*C.sk_GrContextOptions)(unsafe.Pointer(&p1))
	retC := C.misk_GrDirectContextsMakeGLInterfaceOptions(c_p0, c_p1)
	return GrDirectContext{sk: retC}
}

func GrDirectContextsMakeGLInterface(p0 GrGLInterface) GrDirectContext {
	c_p0 := p0.sk
	retC := C.misk_GrDirectContextsMakeGLInterface(c_p0)
	return GrDirectContext{sk: retC}
}

func GrDirectContextsMakeGLOptions(p0 GrContextOptions) GrDirectContext {
	c_p0 := *(*C.sk_GrContextOptions)(unsafe.Pointer(&p0))
	retC := C.misk_GrDirectContextsMakeGLOptions(c_p0)
	return GrDirectContext{sk: retC}
}

func GrDirectContextsMakeGL() GrDirectContext {
	retC := C.misk_GrDirectContextsMakeGL()
	return GrDirectContext{sk: retC}
}

/*
Returns color value from 8-bit component values. Asserts if SK_DEBUG is defined
if a, r, g, or b exceed 255. Since color is unpremultiplied, a may be smaller
than the largest of r, g, and b.

@param a  amount of alpha, from fully transparent (0) to fully opaque (255)
@param r  amount of red, from no red (0) to full red (255)
@param g  amount of green, from no green (0) to full green (255)
@param b  amount of blue, from no blue (0) to full blue (255)
@return   color and alpha, unpremultiplied
*/
func SkColorSetARGB(a uint, r uint, g uint, b uint) Color {
	c_a := C.uint(a)
	c_r := C.uint(r)
	c_g := C.uint(g)
	c_b := C.uint(b)
	retC := C.misk_SkColorSetARGB(c_a, c_r, c_g, c_b)
	return Color(retC)
}

/*
Returns unpremultiplied color with red, blue, and green set from c; and alpha set
from a. Alpha component of c is ignored and is replaced by a in result.

@param c  packed RGB, eight bits per component
@param a  alpha: transparent at zero, fully opaque at 255
@return   color with transparency
*/
func SkColorSetA(c Color, a uint) Color {
	c_c := C.uint(c)
	c_a := C.uint(a)
	retC := C.misk_SkColorSetA(c_c, c_a)
	return Color(retC)
}

/*
Returns a SkPMColor value from unpremultiplied 8-bit component values.

@param a  amount of alpha, from fully transparent (0) to fully opaque (255)
@param r  amount of red, from no red (0) to full red (255)
@param g  amount of green, from no green (0) to full green (255)
@param b  amount of blue, from no blue (0) to full blue (255)
@return   premultiplied color
*/
func SkPreMultiplyARGB(a uint, r uint, g uint, b uint) PMColor {
	c_a := C.uint(a)
	c_r := C.uint(r)
	c_g := C.uint(g)
	c_b := C.uint(b)
	retC := C.misk_SkPreMultiplyARGB(c_a, c_r, c_g, c_b)
	return PMColor(retC)
}

/*
Returns pmcolor closest to color c. Multiplies c RGB components by the c alpha,
and arranges the bytes to match the format of kN32_SkColorType.

@param c  unpremultiplied ARGB color
@return   premultiplied color
*/
func SkPreMultiplyColor(c Color) PMColor {
	c_c := C.uint(c)
	retC := C.misk_SkPreMultiplyColor(c_c)
	return PMColor(retC)
}

/*
8-bit type for an alpha value. 255 is 100% opaque, zero is 100% transparent.
*/
type Alpha C.uint

/*
32-bit ARGB color value, unpremultiplied. Color components are always in
a known order. This is different from SkPMColor, which has its bytes in a configuration
dependent order, to match the format of kBGRA_8888_SkColorType bitmaps. SkColor
is the type used to specify colors in SkPaint and in gradients.

Color that is premultiplied has the same component values as color
that is unpremultiplied if alpha is 255, fully opaque, although may have the
component values in a different order.
*/
type Color C.uint

/*
32-bit ARGB color value, premultiplied. The byte order for this value is
configuration dependent, matching the format of kBGRA_8888_SkColorType bitmaps.
This is different from SkColor, which is unpremultiplied, and is always in the
same byte order.
*/
type PMColor C.uint

func FontMgrRefDefault() FontMgr {
	return FontMgr{
		sk: C.sk_fontmgr_ref_default(),
	}
}
