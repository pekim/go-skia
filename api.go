// This is a generated file. DO NOT EDIT.

package skia

// #cgo CXXFLAGS: -I ${SRCDIR}/skia/skia
// #cgo LDFLAGS: -L ${SRCDIR}/skia/build
// #cgo LDFLAGS: -l skia
// #cgo LDFLAGS: -l skshaper
// #cgo LDFLAGS: -l svg
// #cgo pkg-config: freetype2
// #include "api.h"
import "C"

import (
	"unsafe"
)

/*
SkMatrix holds a 3x3 matrix for transforming coordinates. This allows mapping
SkPoint and vectors with translation, scaling, skewing, rotation, and
perspective.

SkMatrix elements are in row major order.
SkMatrix constexpr default constructs to identity.

SkMatrix includes a hidden variable that classifies the type of matrix to
improve performance. SkMatrix is not thread safe unless getType() is called first.

example: https://fiddle.skia.org/c/
_063
*/
type Matrix struct {
	skia unsafe.Pointer
}

/*
Creates an identity SkMatrix:

| 1 0 0 |
| 0 1 0 |
| 0 0 1 |*/
func NewMatrix() Matrix {

	return Matrix{
		skia: C.skia_new_SkMatrix(),
	}
}

/*
ScaleToFit describes how SkMatrix is constructed to map one SkRect to another.
ScaleToFit may allow SkMatrix to have unequal horizontal and vertical scaling,
or may restrict SkMatrix to square scaling. If restricted, ScaleToFit specifies
how SkMatrix maps to the side or center of the destination SkRect.
*/
type MatrixScaleToFit int64

const (
	/*
	   scales in x and y to fill destination SkRect*/
	MatrixScaleToFit_Fill MatrixScaleToFit = 0
	/*
	   scales and aligns to left and top*/
	MatrixScaleToFit_Start MatrixScaleToFit = 1
	/*
	   scales and aligns to center*/
	MatrixScaleToFit_Center MatrixScaleToFit = 2
	/*
	   scales and aligns to right and bottom*/
	MatrixScaleToFit_End MatrixScaleToFit = 3
)

/*
Enum of bit fields for mask returned by getType().
Used to identify the complexity of SkMatrix, to optimize performance.
*/
type MatrixTypeMask int64

const (
	/*
	   identity SkMatrix; all bits clear*/
	MatrixTypeMask_Identity_Mask MatrixTypeMask = 0
	/*
	   translation SkMatrix*/
	MatrixTypeMask_Translate_Mask MatrixTypeMask = 1
	/*
	   scale SkMatrix*/
	MatrixTypeMask_Scale_Mask MatrixTypeMask = 2
	/*
	   skew or rotate SkMatrix*/
	MatrixTypeMask_Affine_Mask MatrixTypeMask = 4
	/*
	   perspective SkMatrix*/
	MatrixTypeMask_Perspective_Mask MatrixTypeMask = 8
)

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type RefCntBase struct {
	skia unsafe.Pointer
}

/*
Default construct, initializing the reference count to 1.*/
func NewRefCntBase() RefCntBase {

	return RefCntBase{
		skia: C.skia_new_SkRefCntBase(),
	}
}

func (o *RefCntBase) Delete() {
	C.skia_delete_SkRefCntBase(o.skia)
}

/*
Describes pixel and encoding. SkImageInfo can be created from SkColorInfo by
providing dimensions.

It encodes how pixel bits describe alpha, transparency; color components red, blue,
and green; and SkColorSpace, the range and linearity of colors.
*/
type ColorInfo struct {
	skia unsafe.Pointer
}

/*
Creates an SkColorInfo with kUnknown_SkColorType, kUnknown_SkAlphaType,
and no SkColorSpace.*/
func NewColorInfo() ColorInfo {

	return ColorInfo{
		skia: C.skia_new_SkColorInfo(),
	}
}

func (o *ColorInfo) Delete() {
	C.skia_delete_SkColorInfo(o.skia)
}

/*
SkPixmap provides a utility to pair SkImageInfo with pixels and row bytes.
SkPixmap is a low level class which provides convenience functions to access
raster destinations. SkCanvas can not draw SkPixmap, nor does SkPixmap provide
a direct drawing destination.

Use SkBitmap to draw pixels referenced by SkPixmap; use SkSurface to draw into
pixels referenced by SkPixmap.

SkPixmap does not try to manage the lifetime of the pixel memory. Use SkPixelRef
to manage pixel memory; SkPixelRef is safe across threads.
*/
type Pixmap struct {
	skia unsafe.Pointer
}

/*
Creates an empty SkPixmap without pixels, with kUnknown_SkColorType, with
kUnknown_SkAlphaType, and with a width and height of zero. Use
reset() to associate pixels, SkColorType, SkAlphaType, width, and height
after SkPixmap has been created.*/
func NewPixmap() Pixmap {

	return Pixmap{
		skia: C.skia_new_SkPixmap(),
	}
}

/*
SkData holds an immutable data buffer. Not only is the data immutable,
but the actual ptr that is returned (by data() or bytes()) is guaranteed
to always be the same for the life of this instance.
*/
type Data struct {
	skia unsafe.Pointer
}

/*
Specifies the structure of planes for a YUV image with optional alpha. The actual planar data
is not part of this structure and depending on usage is in external textures or pixmaps.
*/
type YUVAInfo struct {
	skia unsafe.Pointer
}

func NewYUVAInfo() YUVAInfo {

	return YUVAInfo{
		skia: C.skia_new_SkYUVAInfo(),
	}
}

type YUVAInfoYUVAChannels int64

const (
	YUVAInfoYUVAChannels_Y    YUVAInfoYUVAChannels = 0
	YUVAInfoYUVAChannels_U    YUVAInfoYUVAChannels = 1
	YUVAInfoYUVAChannels_V    YUVAInfoYUVAChannels = 2
	YUVAInfoYUVAChannels_A    YUVAInfoYUVAChannels = 3
	YUVAInfoYUVAChannels_Last YUVAInfoYUVAChannels = 3
)

/*
Specifies how YUV (and optionally A) are divided among planes. Planes are separated by
underscores in the enum value names. Within each plane the pixmap/texture channels are
mapped to the YUVA channels in the order specified, e.g. for kY_UV Y is in channel 0 of plane
0, U is in channel 0 of plane 1, and V is in channel 1 of plane 1. Channel ordering
within a pixmap/texture given the channels it contains:
A:                       0:A
Luminance/Gray:          0:Gray
Luminance/Gray + Alpha:  0:Gray, 1:A
RG                       0:R,    1:G
RGB                      0:R,    1:G, 2:B
RGBA                     0:R,    1:G, 2:B, 3:A
*/
type YUVAInfoPlaneConfig int64

const (
	YUVAInfoPlaneConfig_Unknown YUVAInfoPlaneConfig = 0
	/*
	   Plane 0: Y, Plane 1: U,  Plane 2: V*/
	YUVAInfoPlaneConfig_Y_U_V YUVAInfoPlaneConfig = 1
	/*
	   Plane 0: Y, Plane 1: V,  Plane 2: U*/
	YUVAInfoPlaneConfig_Y_V_U YUVAInfoPlaneConfig = 2
	/*
	   Plane 0: Y, Plane 1: UV*/
	YUVAInfoPlaneConfig_Y_UV YUVAInfoPlaneConfig = 3
	/*
	   Plane 0: Y, Plane 1: VU*/
	YUVAInfoPlaneConfig_Y_VU YUVAInfoPlaneConfig = 4
	/*
	   Plane 0: YUV*/
	YUVAInfoPlaneConfig_YUV YUVAInfoPlaneConfig = 5
	/*
	   Plane 0: UYV*/
	YUVAInfoPlaneConfig_UYV YUVAInfoPlaneConfig = 6
	/*
	   Plane 0: Y, Plane 1: U,  Plane 2: V, Plane 3: A*/
	YUVAInfoPlaneConfig_Y_U_V_A YUVAInfoPlaneConfig = 7
	/*
	   Plane 0: Y, Plane 1: V,  Plane 2: U, Plane 3: A*/
	YUVAInfoPlaneConfig_Y_V_U_A YUVAInfoPlaneConfig = 8
	/*
	   Plane 0: Y, Plane 1: UV, Plane 2: A*/
	YUVAInfoPlaneConfig_Y_UV_A YUVAInfoPlaneConfig = 9
	/*
	   Plane 0: Y, Plane 1: VU, Plane 2: A*/
	YUVAInfoPlaneConfig_Y_VU_A YUVAInfoPlaneConfig = 10
	/*
	   Plane 0: YUVA*/
	YUVAInfoPlaneConfig_YUVA YUVAInfoPlaneConfig = 11
	/*
	   Plane 0: UYVA*/
	YUVAInfoPlaneConfig_UYVA YUVAInfoPlaneConfig = 12
	YUVAInfoPlaneConfig_Last YUVAInfoPlaneConfig = 12
)

/*
UV subsampling is also specified in the enum value names using J:a:b notation (e.g. 4:2:0 is
1/2 horizontal and 1/2 vertical resolution for U and V). If alpha is present it is not sub-
sampled. Note that Subsampling values other than k444 are only valid with PlaneConfig values
that have U and V in different planes than Y (and A, if present).
*/
type YUVAInfoSubsampling int64

const (
	YUVAInfoSubsampling_Unknown YUVAInfoSubsampling = 0
	/*
	   No subsampling. UV values for each Y.*/
	YUVAInfoSubsampling_444 YUVAInfoSubsampling = 1
	/*
	   1 set of UV values for each 2x1 block of Y values.*/
	YUVAInfoSubsampling_422 YUVAInfoSubsampling = 2
	/*
	   1 set of UV values for each 2x2 block of Y values.*/
	YUVAInfoSubsampling_420 YUVAInfoSubsampling = 3
	/*
	   1 set of UV values for each 1x2 block of Y values.*/
	YUVAInfoSubsampling_440 YUVAInfoSubsampling = 4
	/*
	   1 set of UV values for each 4x1 block of Y values.*/
	YUVAInfoSubsampling_411 YUVAInfoSubsampling = 5
	/*
	   1 set of UV values for each 4x2 block of Y values.*/
	YUVAInfoSubsampling_410  YUVAInfoSubsampling = 6
	YUVAInfoSubsampling_Last YUVAInfoSubsampling = 6
)

/*
Describes how subsampled chroma values are sited relative to luma values.

Currently only centered siting is supported but will expand to support additional sitings.
*/
type YUVAInfoSiting int64

const (
	/*
	   Subsampled chroma value is sited at the center of the block of corresponding luma values.*/
	YUVAInfoSiting_Centered YUVAInfoSiting = 0
)

/*
SkYUVAInfo combined with per-plane SkColorTypes and row bytes. Fully specifies the SkPixmaps
for a YUVA image without the actual pixel memory and data.
*/
type YUVAPixmapInfo struct {
	skia unsafe.Pointer
}

/*
Default SkYUVAPixmapInfo is invalid.*/
func NewYUVAPixmapInfo() YUVAPixmapInfo {

	return YUVAPixmapInfo{
		skia: C.skia_new_SkYUVAPixmapInfo(),
	}
}

/*
Data type for Y, U, V, and possibly A channels independent of how values are packed into
planes.
*/
type YUVAPixmapInfoDataType int64

const (
	/*
	   8 bit unsigned normalized*/
	YUVAPixmapInfoDataType_Unorm8 YUVAPixmapInfoDataType = 0
	/*
	   16 bit unsigned normalized*/
	YUVAPixmapInfoDataType_Unorm16 YUVAPixmapInfoDataType = 1
	/*
	   16 bit (half) floating point*/
	YUVAPixmapInfoDataType_Float16 YUVAPixmapInfoDataType = 2
	/*
	   10 bit unorm for Y, U, and V. 2 bit unorm for alpha (if present).*/
	YUVAPixmapInfoDataType_Unorm10_Unorm2 YUVAPixmapInfoDataType = 3
	YUVAPixmapInfoDataType_Last           YUVAPixmapInfoDataType = 3
)

/*
Helper to store SkPixmap planes as described by a SkYUVAPixmapInfo. Can be responsible for
allocating/freeing memory for pixmaps or use external memory.
*/
type YUVAPixmaps struct {
	skia unsafe.Pointer
}

/*
Default SkYUVAPixmaps is invalid.*/
func NewYUVAPixmaps() YUVAPixmaps {

	return YUVAPixmaps{
		skia: C.skia_new_SkYUVAPixmaps(),
	}
}

func (o *YUVAPixmaps) Delete() {
	C.skia_delete_SkYUVAPixmaps(o.skia)
}

type Once struct {
	skia unsafe.Pointer
}

func NewOnce() Once {

	return Once{
		skia: C.skia_new_SkOnce(),
	}
}

type OnceState int64

const (
	OnceState_NotStarted OnceState = 0
	OnceState_Claimed    OnceState = 1
	OnceState_Done       OnceState = 2
)

/*
////////////////////////////////////////////////////////////////////////////
*/
type ColorSpace struct {
	skia unsafe.Pointer
}

/*
SkNoncopyable is the base class for objects that do not want to
be copied. It hides its copy-constructor and its assignment-operator.
*/
type Noncopyable struct {
	skia unsafe.Pointer
}

func NewNoncopyable() Noncopyable {

	return Noncopyable{
		skia: C.skia_new_SkNoncopyable(),
	}
}

/*
Abstraction layer directly on top of an image codec.
*/
type Codec struct {
	skia unsafe.Pointer
}

func (o *Codec) Delete() {
	C.skia_delete_SkCodec(o.skia)
}

/*
Error codes for various SkCodec methods.
*/
type CodecResult int64

const (
	/*
	   General return value for success.*/
	CodecResult_Success CodecResult = 0
	/*
	   The input is incomplete. A partial image was generated.*/
	CodecResult_IncompleteInput CodecResult = 1
	/*
	   Like kIncompleteInput, except the input had an error.

	   If returned from an incremental decode, decoding cannot continue,
	   even with more data.*/
	CodecResult_ErrorInInput CodecResult = 2
	/*
	   The generator cannot convert to match the request, ignoring
	   dimensions.*/
	CodecResult_InvalidConversion CodecResult = 3
	/*
	   The generator cannot scale to requested size.*/
	CodecResult_InvalidScale CodecResult = 4
	/*
	   Parameters (besides info) are invalid. e.g. NULL pixels, rowBytes
	   too small, etc.*/
	CodecResult_InvalidParameters CodecResult = 5
	/*
	   The input did not contain a valid image.*/
	CodecResult_InvalidInput CodecResult = 6
	/*
	   Fulfilling this request requires rewinding the input, which is not
	   supported for this input.*/
	CodecResult_CouldNotRewind CodecResult = 7
	/*
	   An internal error, such as OOM.*/
	CodecResult_InternalError CodecResult = 8
	/*
	   This method is not implemented by this codec.
	   FIXME: Perhaps this should be kUnsupported?*/
	CodecResult_Unimplemented CodecResult = 9
)

/*
For container formats that contain both still images and image sequences,
instruct the decoder how the output should be selected. (Refer to comments
for each value for more details.)
*/
type CodecSelectionPolicy int64

const (
	/*
	   If the container format contains both still images and image sequences,
	   SkCodec should choose one of the still images. This is the default.
	   Note that kPreferStillImage may prevent use of the animation features
	   if the input is not rewindable.*/
	CodecSelectionPolicy_PreferStillImage CodecSelectionPolicy = 0
	/*
	   If the container format contains both still images and image sequences,
	   SkCodec should choose one of the image sequences for animation.*/
	CodecSelectionPolicy_PreferAnimation CodecSelectionPolicy = 1
)

/*
Whether or not the memory passed to getPixels is zero initialized.
*/
type CodecZeroInitialized int64

const (
	/*
	   The memory passed to getPixels is zero initialized. The SkCodec
	   may take advantage of this by skipping writing zeroes.*/
	CodecZeroInitialized_Yes CodecZeroInitialized = 0
	/*
	   The memory passed to getPixels has not been initialized to zero,
	   so the SkCodec must write all zeroes to memory.

	   This is the default. It will be used if no Options struct is used.*/
	CodecZeroInitialized_No CodecZeroInitialized = 1
)

/*
The order in which rows are output from the scanline decoder is not the
same for all variations of all image types.  This explains the possible
output row orderings.
*/
type CodecSkScanlineOrder int64

const (
	CodecSkScanlineOrder_TopDown  CodecSkScanlineOrder = 0
	CodecSkScanlineOrder_BottomUp CodecSkScanlineOrder = 1
)

type CodecXformTime int64

const (
	CodecXformTime_No        CodecXformTime = 0
	CodecXformTime_Palette   CodecXformTime = 1
	CodecXformTime_DecodeRow CodecXformTime = 2
)

/*
SkFlattenable is the base class for objects that need to be flattened
into a data stream for either transport or as part of the key to the
font cache.
*/
type Flattenable struct {
	skia unsafe.Pointer
}

type FlattenableType int64

const (
	FlattenableType_SkColorFilter FlattenableType = 0
	FlattenableType_SkBlender     FlattenableType = 1
	FlattenableType_SkDrawable    FlattenableType = 2
	FlattenableType_SkDrawLooper  FlattenableType = 3
	FlattenableType_SkImageFilter FlattenableType = 4
	FlattenableType_SkMaskFilter  FlattenableType = 5
	FlattenableType_SkPathEffect  FlattenableType = 6
	FlattenableType_SkShader      FlattenableType = 7
)

/*
Base class for image filters. If one is installed in the paint, then all drawing occurs as
usual, but it is as if the drawing happened into an offscreen (before the xfermode is applied).
This offscreen bitmap will then be handed to the imagefilter, who in turn creates a new bitmap
which is what will finally be drawn to the device (using the original xfermode).

The local space of image filters matches the local space of the drawn geometry. For instance if
there is rotation on the canvas, the blur will be computed along those rotated axes and not in
the device space. In order to achieve this result, the actual drawing of the geometry may happen
in an unrotated coordinate system so that the filtered image can be computed more easily, and
then it will be post transformed to match what would have been produced if the geometry were
drawn with the total canvas matrix to begin with.
*/
type ImageFilter struct {
	skia unsafe.Pointer
}

type ImageFilterMapDirection int64

const (
	ImageFilterMapDirection_Forward ImageFilterMapDirection = 0
	ImageFilterMapDirection_Reverse ImageFilterMapDirection = 1
)

/*
4x4 matrix used by SkCanvas and other parts of Skia.

Skia assumes a right-handed coordinate system:
+X goes to the right
+Y goes down
+Z goes into the screen (away from the viewer)
*/
type M44 struct {
	skia unsafe.Pointer
}

func NewM44() M44 {

	return M44{
		skia: C.skia_new_SkM44(),
	}
}

func NewM442(p0 M44Uninitialized_Constructor) M44 {
	c_p0 := C.int(p0)

	return M44{
		skia: C.skia_new_SkM442(c_p0),
	}
}

func NewM443(p0 M44NaN_Constructor) M44 {
	c_p0 := C.int(p0)

	return M44{
		skia: C.skia_new_SkM443(c_p0),
	}
}

/*
The constructor parameters are in row-major order.*/
func NewM444(m0 float32, m4 float32, m8 float32, m12 float32, m1 float32, m5 float32, m9 float32, m13 float32, m2 float32, m6 float32, m10 float32, m14 float32, m3 float32, m7 float32, m11 float32, m15 float32) M44 {
	c_m0 := C.float(m0)
	c_m4 := C.float(m4)
	c_m8 := C.float(m8)
	c_m12 := C.float(m12)
	c_m1 := C.float(m1)
	c_m5 := C.float(m5)
	c_m9 := C.float(m9)
	c_m13 := C.float(m13)
	c_m2 := C.float(m2)
	c_m6 := C.float(m6)
	c_m10 := C.float(m10)
	c_m14 := C.float(m14)
	c_m3 := C.float(m3)
	c_m7 := C.float(m7)
	c_m11 := C.float(m11)
	c_m15 := C.float(m15)

	return M44{
		skia: C.skia_new_SkM444(c_m0, c_m4, c_m8, c_m12, c_m1, c_m5, c_m9, c_m13, c_m2, c_m6, c_m10, c_m14, c_m3, c_m7, c_m11, c_m15),
	}
}

type M44Uninitialized_Constructor int64

const (
	M44Uninitialized_Constructor_ M44Uninitialized_Constructor = 0
)

type M44NaN_Constructor int64

const (
	M44NaN_Constructor_ M44NaN_Constructor = 0
)

/*
SkPaint controls options applied when drawing. SkPaint collects all
options outside of the SkCanvas clip and SkCanvas matrix.

Various options apply to strokes and fills, and images.

SkPaint collects effects and filters that describe single-pass and multiple-pass
algorithms that alter the drawing geometry, color, and transparency. For instance,
SkPaint does not directly implement dashing or blur, but contains the objects that do so.
*/
type Paint struct {
	skia unsafe.Pointer
}

/*
Constructs SkPaint with default values.



example: https://fiddle.skia.org/c/
_empty_constructor*/
func NewPaint() Paint {

	return Paint{
		skia: C.skia_new_SkPaint(),
	}
}

func (o *Paint) Delete() {
	C.skia_delete_SkPaint(o.skia)
}

/*
Set Style to fill, stroke, or both fill and stroke geometry.
The stroke and fill
share all paint attributes; for instance, they are drawn with the same color.

Use kStrokeAndFill_Style to avoid hitting the same pixels twice with a stroke draw and
a fill draw.
*/
type PaintStyle int64

const (
	/*
	   set to fill geometry*/
	PaintStyle_Fill PaintStyle = 0
	/*
	   set to stroke geometry*/
	PaintStyle_Stroke PaintStyle = 1
	/*
	   sets to stroke and fill geometry*/
	PaintStyle_StrokeAndFill PaintStyle = 2
)

/*
Cap draws at the beginning and end of an open path contour.
*/
type PaintCap int64

const (
	/*
	   no stroke extension*/
	PaintCap_Butt PaintCap = 0
	/*
	   adds circle*/
	PaintCap_Round PaintCap = 1
	/*
	   adds square*/
	PaintCap_Square PaintCap = 2
	/*
	   largest Cap value*/
	PaintCap_Last PaintCap = 2
	/*
	   equivalent to kButt_Cap*/
	PaintCap_Default PaintCap = 0
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
	/*
	   extends to miter limit*/
	PaintJoin_Miter PaintJoin = 0
	/*
	   adds circle*/
	PaintJoin_Round PaintJoin = 1
	/*
	   connects outside edges*/
	PaintJoin_Bevel PaintJoin = 2
	/*
	   equivalent to the largest value for Join*/
	PaintJoin_Last PaintJoin = 2
	/*
	   equivalent to kMiter_Join*/
	PaintJoin_Default PaintJoin = 0
)

/*
If a client wants to control the allocation of raster layers in a canvas, it should subclass
SkRasterHandleAllocator. This allocator performs two tasks:
1. controls how the memory for the pixels is allocated
2. associates a "handle" to a private object that can track the matrix/clip of the SkCanvas

This example allocates a canvas, and defers to the allocator to create the base layer.

std::unique_ptr
<SkCanvas
> canvas = SkRasterHandleAllocator::MakeCanvas(
SkImageInfo::Make(...),
std::make_unique
<MySubclassRasterHandleAllocator
>(...),
nullptr);

If you have already allocated the base layer (and its handle, release-proc etc.) then you
can pass those in using the last parameter to MakeCanvas().

Regardless of how the base layer is allocated, each time canvas->saveLayer() is called,
your allocator's allocHandle() will be called.
*/
type RasterHandleAllocator struct {
	skia unsafe.Pointer
}

func (o *RasterHandleAllocator) Delete() {
	C.skia_delete_SkRasterHandleAllocator(o.skia)
}

/*
Light weight class for managing strings. Uses reference
counting to make string assignments and copies very fast
with no extra RAM cost. Assumes UTF8 encoding.
*/
type String struct {
	skia unsafe.Pointer
}

func NewString() String {

	return String{
		skia: C.skia_new_SkString(),
	}
}

func NewString2(len uint64) String {
	c_len := C.ulong(len)

	return String{
		skia: C.skia_new_SkString2(c_len),
	}
}

func (o *String) Delete() {
	C.skia_delete_SkString(o.skia)
}

/*
Describes properties and constraints of a given SkSurface. The rendering engine can parse these
during drawing, and can sometimes optimize its performance (e.g. disabling an expensive
feature).
*/
type SurfaceProps struct {
	skia unsafe.Pointer
}

/*
No flags, unknown pixel geometry, platform-default contrast/gamma.*/
func NewSurfaceProps() SurfaceProps {

	return SurfaceProps{
		skia: C.skia_new_SkSurfaceProps(),
	}
}

/*
TODO(kschmi): Remove this constructor and replace with the one below. **/
func NewSurfaceProps2(flags uint32, p1 PixelGeometry) SurfaceProps {
	c_flags := C.uint(flags)
	c_p1 := C.int(p1)

	return SurfaceProps{
		skia: C.skia_new_SkSurfaceProps2(c_flags, c_p1),
	}
}

/*
Specified pixel geometry, text contrast, and gamma **/
func NewSurfaceProps3(flags uint32, p1 PixelGeometry, textContrast float32, textGamma float32) SurfaceProps {
	c_flags := C.uint(flags)
	c_p1 := C.int(p1)
	c_textContrast := C.float(textContrast)
	c_textGamma := C.float(textGamma)

	return SurfaceProps{
		skia: C.skia_new_SkSurfaceProps3(c_flags, c_p1, c_textContrast, c_textGamma),
	}
}

type SurfacePropsFlags int64

const (
	SurfacePropsFlags_Default                   SurfacePropsFlags = 0
	SurfacePropsFlags_UseDeviceIndependentFonts SurfacePropsFlags = 1
	SurfacePropsFlags_DynamicMSAA               SurfacePropsFlags = 2
	SurfacePropsFlags_AlwaysDither              SurfacePropsFlags = 4
)

type Deque struct {
	skia unsafe.Pointer
}

/*
elemSize specifies the size of each individual element in the deque
allocCount specifies how many elements are to be allocated as a block*/
func NewDeque(elemSize uint64, allocCount int32) Deque {
	c_elemSize := C.ulong(elemSize)
	c_allocCount := C.int(allocCount)

	return Deque{
		skia: C.skia_new_SkDeque(c_elemSize, c_allocCount),
	}
}

func (o *Deque) Delete() {
	C.skia_delete_SkDeque(o.skia)
}

type ContainerAllocator struct {
	skia unsafe.Pointer
}

func NewContainerAllocator(sizeOfT uint64, maxCapacity int32) ContainerAllocator {
	c_sizeOfT := C.ulong(sizeOfT)
	c_maxCapacity := C.int(maxCapacity)

	return ContainerAllocator{
		skia: C.skia_new_SkContainerAllocator(c_sizeOfT, c_maxCapacity),
	}
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
	skia unsafe.Pointer
}

/*
Creates an empty SkCanvas with no backing device or pixels, with
a width and height of zero.



example: https://fiddle.skia.org/c/
_empty_constructor*/
func NewCanvas() Canvas {

	return Canvas{
		skia: C.skia_new_SkCanvas(),
	}
}

func (o *Canvas) Delete() {
	C.skia_delete_SkCanvas(o.skia)
}

/*
SaveLayerFlags provides options that may be used in any combination in SaveLayerRec,
defining how layer allocated by saveLayer() operates. It may be set to zero,
kPreserveLCDText_SaveLayerFlag, kInitWithPrevious_SaveLayerFlag, or both flags.
*/
type CanvasSaveLayerFlagsSet int64

const (
	CanvasSaveLayerFlagsSet_PreserveLCDText_SaveLayer CanvasSaveLayerFlagsSet = 2
	/*
	   initializes with previous contents*/
	CanvasSaveLayerFlagsSet_InitWithPrevious_SaveLayer CanvasSaveLayerFlagsSet = 4
	CanvasSaveLayerFlagsSet_F16ColorType               CanvasSaveLayerFlagsSet = 16
)

/*
Selects if an array of points are drawn as discrete points, as lines, or as
an open polygon.
*/
type CanvasPointMode int64

const (
	/*
	   draw each point separately*/
	CanvasPointMode_Points CanvasPointMode = 0
	/*
	   draw each pair of points as a line segment*/
	CanvasPointMode_Lines CanvasPointMode = 1
	/*
	   draw the array of points as a open polygon*/
	CanvasPointMode_Polygon CanvasPointMode = 2
)

/*
SrcRectConstraint controls the behavior at the edge of source SkRect,
provided to drawImageRect() when there is any filtering. If kStrict is set,
then extra code is used to ensure it never samples outside of the src-rect.
kStrict_SrcRectConstraint disables the use of mipmaps and anisotropic filtering.
*/
type CanvasSrcRectConstraint int64

const (
	/*
	   sample only inside bounds; slower*/
	CanvasSrcRectConstraint_Strict CanvasSrcRectConstraint = 0
	/*
	   sample outside bounds; faster*/
	CanvasSrcRectConstraint_Fast CanvasSrcRectConstraint = 1
)

/*
Experimental. Controls anti-aliasing of each edge of images in an image-set.
*/
type CanvasQuadAAFlags int64

const (
	CanvasQuadAAFlags_Left_QuadAA   CanvasQuadAAFlags = 1
	CanvasQuadAAFlags_Top_QuadAA    CanvasQuadAAFlags = 2
	CanvasQuadAAFlags_Right_QuadAA  CanvasQuadAAFlags = 4
	CanvasQuadAAFlags_Bottom_QuadAA CanvasQuadAAFlags = 8
	CanvasQuadAAFlags_None          CanvasQuadAAFlags = 0
	CanvasQuadAAFlags_All           CanvasQuadAAFlags = 15
)

type CanvasSaveLayerStrategy int64

const (
	CanvasSaveLayerStrategy_FullLayer CanvasSaveLayerStrategy = 0
	CanvasSaveLayerStrategy_NoLayer   CanvasSaveLayerStrategy = 1
)

type CanvasClipEdgeStyle int64

const (
	CanvasClipEdgeStyle_Hard CanvasClipEdgeStyle = 0
	CanvasClipEdgeStyle_Soft CanvasClipEdgeStyle = 1
)

type CanvasPredrawFlags int64

const (
	CanvasPredrawFlags_None                    CanvasPredrawFlags = 0
	CanvasPredrawFlags_OpaqueShaderOverride    CanvasPredrawFlags = 1
	CanvasPredrawFlags_NonOpaqueShaderOverride CanvasPredrawFlags = 2
	CanvasPredrawFlags_CheckForOverwrite       CanvasPredrawFlags = 4
	CanvasPredrawFlags_SkipMaskFilterAutoLayer CanvasPredrawFlags = 8
)

type CanvasDeviceCompatibleWithFilter int64

const (
	CanvasDeviceCompatibleWithFilter_Unknown CanvasDeviceCompatibleWithFilter = 0
	CanvasDeviceCompatibleWithFilter_Yes     CanvasDeviceCompatibleWithFilter = -1
)

/*
Stack helper class calls SkCanvas::restoreToCount when SkAutoCanvasRestore
goes out of scope. Use this to guarantee that the canvas is restored to a known
state.
*/
type AutoCanvasRestore struct {
	skia unsafe.Pointer
}

func (o *AutoCanvasRestore) Delete() {
	C.skia_delete_SkAutoCanvasRestore(o.skia)
}

/*
High-level API for creating a document-based canvas. To use..

1. Create a document, specifying a stream to store the output.
2. For each "page" of content:
a. canvas = doc->beginPage(...)
b. draw_my_content(canvas);
c. doc->endPage();
3. Close the document with doc->close().
*/
type Document struct {
	skia unsafe.Pointer
}

type DocumentState int64

const (
	DocumentState_BetweenPages DocumentState = 0
	DocumentState_InPage       DocumentState = 1
	DocumentState_Closed       DocumentState = 2
)

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type FontStyleSet struct {
	skia unsafe.Pointer
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type FontMgr struct {
	skia unsafe.Pointer
}

/*
SkMaskFilter is the base class for object that perform transformations on
the mask before drawing it. An example subclass is Blur.
*/
type MaskFilter struct {
	skia unsafe.Pointer
}

/*
SkStream -- abstraction for a source of bytes. Subclasses can be backed by
memory, or a file, or something else.
*/
type Stream struct {
	skia unsafe.Pointer
}

func (o *Stream) Delete() {
	C.skia_delete_SkStream(o.skia)
}

/*
SkStreamRewindable is a SkStream for which rewind and duplicate are required.
*/
type StreamRewindable struct {
	skia unsafe.Pointer
}

/*
SkStreamSeekable is a SkStreamRewindable for which position, seek, move, and fork are required.
*/
type StreamSeekable struct {
	skia unsafe.Pointer
}

/*
SkStreamAsset is a SkStreamSeekable for which getLength is required.
*/
type StreamAsset struct {
	skia unsafe.Pointer
}

/*
SkStreamMemory is a SkStreamAsset for which getMemoryBase is required.
*/
type StreamMemory struct {
	skia unsafe.Pointer
}

type WStream struct {
	skia unsafe.Pointer
}

func (o *WStream) Delete() {
	C.skia_delete_SkWStream(o.skia)
}

type NullWStream struct {
	skia unsafe.Pointer
}

/*
A stream that wraps a C FILE* file stream.
*/
type FILEStream struct {
	skia unsafe.Pointer
}

func (o *FILEStream) Delete() {
	C.skia_delete_SkFILEStream(o.skia)
}

/*
SkStreamMemory is a SkStreamAsset for which getMemoryBase is required.
*/
type MemoryStream struct {
	skia unsafe.Pointer
}

func NewMemoryStream() MemoryStream {

	return MemoryStream{
		skia: C.skia_new_SkMemoryStream(),
	}
}

/*
We allocate (and free) the memory. Write to it via getMemoryBase()*/
func NewMemoryStream2(length uint64) MemoryStream {
	c_length := C.ulong(length)

	return MemoryStream{
		skia: C.skia_new_SkMemoryStream2(c_length),
	}
}

/*
//////////////////////////////////////////////////////////////////////////////////////////
*/
type FILEWStream struct {
	skia unsafe.Pointer
}

func (o *FILEWStream) Delete() {
	C.skia_delete_SkFILEWStream(o.skia)
}

type DynamicMemoryWStream struct {
	skia unsafe.Pointer
}

func NewDynamicMemoryWStream() DynamicMemoryWStream {

	return DynamicMemoryWStream{
		skia: C.skia_new_SkDynamicMemoryWStream(),
	}
}

func (o *DynamicMemoryWStream) Delete() {
	C.skia_delete_SkDynamicMemoryWStream(o.skia)
}

/*
SkImage describes a two dimensional array of pixels to draw. The pixels may be
decoded in a raster bitmap, encoded in a SkPicture or compressed data stream,
or located in GPU memory as a GPU texture.

SkImage cannot be modified after it is created. SkImage may allocate additional
storage as needed; for instance, an encoded SkImage may decode when drawn.

SkImage width and height are greater than zero. Creating an SkImage with zero width
or height returns SkImage equal to nullptr.

SkImage may be created from SkBitmap, SkPixmap, SkSurface, SkPicture, encoded streams,
GPU texture, YUV_ColorSpace data, or hardware buffer. Encoded streams supported
include BMP, GIF, HEIF, ICO, JPEG, PNG, WBMP, WebP. Supported encoding details
vary with platform.

See SkImages namespace for the static factory methods to make SkImages.

Clients should *not* subclass SkImage as there is a lot of internal machinery that is
not publicly accessible.
*/
type Image struct {
	skia unsafe.Pointer
}

/*
CachingHint selects whether Skia may internally cache SkBitmap generated by
decoding SkImage, or by copying SkImage from GPU to CPU. The default behavior
allows caching SkBitmap.

Choose kDisallow_CachingHint if SkImage pixels are to be used only once, or
if SkImage pixels reside in a cache outside of Skia, or to reduce memory pressure.

Choosing kAllow_CachingHint does not ensure that pixels will be cached.
SkImage pixels may not be cached if memory requirements are too large or
pixels are not accessible.
*/
type ImageCachingHint int64

const (
	/*
	   allows internally caching decoded and copied pixels*/
	ImageCachingHint_Allow ImageCachingHint = 0
	/*
	   disallows internally caching decoded and copied pixels*/
	ImageCachingHint_Disallow ImageCachingHint = 1
)

type ImageRescaleGamma int64

const (
	ImageRescaleGamma_Src    ImageRescaleGamma = 0
	ImageRescaleGamma_Linear ImageRescaleGamma = -1
)

type ImageRescaleMode int64

const (
	ImageRescaleMode_Nearest        ImageRescaleMode = 0
	ImageRescaleMode_Linear         ImageRescaleMode = 1
	ImageRescaleMode_RepeatedLinear ImageRescaleMode = 2
	ImageRescaleMode_RepeatedCubic  ImageRescaleMode = 3
)

/*
Deprecated.
*/
type ImageLegacyBitmapMode int64

const (
	/*
	   returned bitmap is read-only and immutable*/
	ImageLegacyBitmapMode_RO ImageLegacyBitmapMode = 0
)

/*
SkSurface is responsible for managing the pixels that a canvas draws into. The pixels can be
allocated either in CPU memory (a raster surface) or on the GPU (a GrRenderTarget surface).
SkSurface takes care of allocating a SkCanvas that will draw into the surface. Call
surface->getCanvas() to use that canvas (but don't delete it, it is owned by the surface).
SkSurface always has non-zero dimensions. If there is a request for a new surface, and either
of the requested dimensions are zero, then nullptr will be returned.

Clients should *not* subclass SkSurface as there is a lot of internal machinery that is
not publicly accessible.
*/
type Surface struct {
	skia unsafe.Pointer
}

/*
ContentChangeMode members are parameters to notifyContentWillChange().
*/
type SurfaceContentChangeMode int64

const (
	/*
	   discards surface on change*/
	SurfaceContentChangeMode_Discard SurfaceContentChangeMode = 0
	/*
	   preserves surface on change*/
	SurfaceContentChangeMode_Retain SurfaceContentChangeMode = 1
)

type SurfaceBackendHandleAccess int64

const (
	/*
	   back-end object is readable*/
	SurfaceBackendHandleAccess_FlushRead SurfaceBackendHandleAccess = 0
	/*
	   back-end object is writable*/
	SurfaceBackendHandleAccess_FlushWrite SurfaceBackendHandleAccess = 1
	/*
	   back-end object must be overwritten*/
	SurfaceBackendHandleAccess_DiscardWrite SurfaceBackendHandleAccess = 2
)

type FontStyle struct {
	skia unsafe.Pointer
}

func NewFontStyle(weight int32, width int32, slant FontStyleSlant) FontStyle {
	c_weight := C.int(weight)
	c_width := C.int(width)
	c_slant := C.int(slant)

	return FontStyle{
		skia: C.skia_new_SkFontStyle(c_weight, c_width, c_slant),
	}
}

func NewFontStyle2() FontStyle {

	return FontStyle{
		skia: C.skia_new_SkFontStyle2(),
	}
}

type FontStyleWeight int64

const (
	FontStyleWeight_Invisible  FontStyleWeight = 0
	FontStyleWeight_Thin       FontStyleWeight = 100
	FontStyleWeight_ExtraLight FontStyleWeight = 200
	FontStyleWeight_Light      FontStyleWeight = 300
	FontStyleWeight_Normal     FontStyleWeight = 400
	FontStyleWeight_Medium     FontStyleWeight = 500
	FontStyleWeight_SemiBold   FontStyleWeight = 600
	FontStyleWeight_Bold       FontStyleWeight = 700
	FontStyleWeight_ExtraBold  FontStyleWeight = 800
	FontStyleWeight_Black      FontStyleWeight = 900
	FontStyleWeight_ExtraBlack FontStyleWeight = 1000
)

type FontStyleWidth int64

const (
	FontStyleWidth_UltraCondensed FontStyleWidth = 1
	FontStyleWidth_ExtraCondensed FontStyleWidth = 2
	FontStyleWidth_Condensed      FontStyleWidth = 3
	FontStyleWidth_SemiCondensed  FontStyleWidth = 4
	FontStyleWidth_Normal         FontStyleWidth = 5
	FontStyleWidth_SemiExpanded   FontStyleWidth = 6
	FontStyleWidth_Expanded       FontStyleWidth = 7
	FontStyleWidth_ExtraExpanded  FontStyleWidth = 8
	FontStyleWidth_UltraExpanded  FontStyleWidth = 9
)

type FontStyleSlant int64

const (
	FontStyleSlant_Upright FontStyleSlant = 0
	FontStyleSlant_Italic  FontStyleSlant = 1
	FontStyleSlant_Oblique FontStyleSlant = 2
)

/*
SkWeakRefCnt is the base class for objects that may be shared by multiple
objects. When an existing strong owner wants to share a reference, it calls
ref(). When a strong owner wants to release its reference, it calls
unref(). When the shared object's strong reference count goes to zero as
the result of an unref() call, its (virtual) weak_dispose method is called.
It is an error for the destructor to be called explicitly (or via the
object going out of scope on the stack or calling delete) if
getRefCnt() > 1.

In addition to strong ownership, an owner may instead obtain a weak
reference by calling weak_ref(). A call to weak_ref() must be balanced by a
call to weak_unref(). To obtain a strong reference from a weak reference,
call try_ref(). If try_ref() returns true, the owner's pointer is now also
a strong reference on which unref() must be called. Note that this does not
affect the original weak reference, weak_unref() must still be called. When
the weak reference count goes to zero, the object is deleted. While the
weak reference count is positive and the strong reference count is zero the
object still exists, but will be in the disposed state. It is up to the
object to define what this means.

Note that a strong reference implicitly implies a weak reference. As a
result, it is allowable for the owner of a strong ref to call try_ref().
This will have the same effect as calling ref(), but may be more expensive.

Example:

SkWeakRefCnt myRef = strongRef.weak_ref();
... // strongRef.unref() may or may not be called
if (myRef.try_ref()) {
... // use myRef
myRef.unref();
} else {
// myRef is in the disposed state
}
myRef.weak_unref();
*/
type WeakRefCnt struct {
	skia unsafe.Pointer
}

/*
Default construct, initializing the reference counts to 1.
The strong references collectively hold one weak reference. When the
strong reference count goes to zero, the collectively held weak
reference is released.*/
func NewWeakRefCnt() WeakRefCnt {

	return WeakRefCnt{
		skia: C.skia_new_SkWeakRefCnt(),
	}
}

func (o *WeakRefCnt) Delete() {
	C.skia_delete_SkWeakRefCnt(o.skia)
}

/*
The SkTypeface class specifies the typeface and intrinsic style of a font.
This is used in the paint, along with optionally algorithmic settings like
textSize, textSkewX, textScaleX, kFakeBoldText_Mask, to specify
how text appears when drawn (and measured).

Typeface objects are immutable, and so they can be shared between threads.
*/
type Typeface struct {
	skia unsafe.Pointer
}

/*
A typeface can serialize just a descriptor (names, etc.), or it can also include the
actual font data (which can be large). This enum controls how serialize() decides what
to serialize.
*/
type TypefaceSerializeBehavior int64

const (
	TypefaceSerializeBehavior_DoIncludeData      TypefaceSerializeBehavior = 0
	TypefaceSerializeBehavior_DontIncludeData    TypefaceSerializeBehavior = 1
	TypefaceSerializeBehavior_IncludeDataIfLocal TypefaceSerializeBehavior = 2
)

/*
SkFont controls options applied when drawing and measuring text.
*/
type Font struct {
	skia unsafe.Pointer
}

/*
Constructs SkFont with default values.*/
func NewFont() Font {

	return Font{
		skia: C.skia_new_SkFont(),
	}
}

/*
Whether edge pixels draw opaque or with partial transparency.
*/
type FontEdging int64

const (
	/*
	   no transparent pixels on glyph edges*/
	FontEdging_Alias FontEdging = 0
	/*
	   may have transparent pixels on glyph edges*/
	FontEdging_AntiAlias FontEdging = 1
	/*
	   glyph positioned in pixel using transparency*/
	FontEdging_SubpixelAntiAlias FontEdging = 2
)

type FontPrivFlags int64

const (
	FontPrivFlags_ForceAutoHinting_Priv FontPrivFlags = 1
	FontPrivFlags_EmbeddedBitmaps_Priv  FontPrivFlags = 2
	FontPrivFlags_Subpixel_Priv         FontPrivFlags = 4
	FontPrivFlags_LinearMetrics_Priv    FontPrivFlags = 8
	FontPrivFlags_Embolden_Priv         FontPrivFlags = 16
	FontPrivFlags_BaselineSnap_Priv     FontPrivFlags = 32
)

/*
SkTextBlob combines multiple text runs into an immutable container. Each text
run consists of glyphs, SkPaint, and position. Only parts of SkPaint related to
fonts and text rendering are used by run.
*/
type TextBlob struct {
	skia unsafe.Pointer
}

/*
Helper class for constructing SkTextBlob.
*/
type TextBlobBuilder struct {
	skia unsafe.Pointer
}

/*
Constructs empty SkTextBlobBuilder. By default, SkTextBlobBuilder has no runs.



example: https://fiddle.skia.org/c/
_empty_constructor*/
func NewTextBlobBuilder() TextBlobBuilder {

	return TextBlobBuilder{
		skia: C.skia_new_SkTextBlobBuilder(),
	}
}

func (o *TextBlobBuilder) Delete() {
	C.skia_delete_SkTextBlobBuilder(o.skia)
}

type Path1DPathEffect struct {
	skia unsafe.Pointer
}

type Path1DPathEffectStyle int64

const (
	Path1DPathEffectStyle_Translate Path1DPathEffectStyle = 0
	Path1DPathEffectStyle_Rotate    Path1DPathEffectStyle = 1
	Path1DPathEffectStyle_Morph     Path1DPathEffectStyle = 2
	Path1DPathEffectStyle_LastEnum  Path1DPathEffectStyle = 2
)

type Line2DPathEffect struct {
	skia unsafe.Pointer
}

type Path2DPathEffect struct {
	skia unsafe.Pointer
}

/*
ColorFilters are optional objects in the drawing pipeline. When present in
a paint, they are called with the "src" colors, and return new colors, which
are then passed onto the next stage (either ImageFilter or Xfermode).

All subclasses are required to be reentrant-safe : it must be legal to share
the same instance between several threads.
*/
type ColorFilter struct {
	skia unsafe.Pointer
}

type ColorFilters struct {
	skia unsafe.Pointer
}

/*
ColorFilters are optional objects in the drawing pipeline. When present in
a paint, they are called with the "src" colors, and return new colors, which
are then passed onto the next stage (either ImageFilter or Xfermode).

All subclasses are required to be reentrant-safe : it must be legal to share
the same instance between several threads.
*/
type ColorMatrixFilter struct {
	skia unsafe.Pointer
}

/*
SkCornerPathEffect is a subclass of SkPathEffect that can turn sharp corners
into various treatments (e.g. rounded corners)
*/
type CornerPathEffect struct {
	skia unsafe.Pointer
}

type DashPathEffect struct {
	skia unsafe.Pointer
}

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
	skia unsafe.Pointer
}

/*
Constructs an empty SkPath. By default, SkPath has no verbs, no SkPoint, and no weights.
FillType is set to kWinding.



example: https://fiddle.skia.org/c/
_empty_constructor*/
func NewPath() Path {

	return Path{
		skia: C.skia_new_SkPath(),
	}
}

func (o *Path) Delete() {
	C.skia_delete_SkPath(o.skia)
}

/*
Four oval parts with radii (rx, ry) start at last SkPath SkPoint and ends at (x, y).
ArcSize and Direction select one of the four oval parts.
*/
type PathArcSize int64

const (
	/*
	   smaller of arc pair*/
	PathArcSize_Small PathArcSize = 0
	/*
	   larger of arc pair*/
	PathArcSize_Large PathArcSize = 1
)

/*
AddPathMode chooses how addPath() appends. Adding one SkPath to another can extend
the last contour or start a new contour.
*/
type PathAddPathMode int64

const (
	/*
	   Contours are appended to the destination path as new contours.*/
	PathAddPathMode_Append PathAddPathMode = 0
	/*
	   Extends the last contour of the destination path with the first countour
	   of the source path, connecting them with a line.  If the last contour is
	   closed, a new empty contour starting at its start point is extended instead.
	   If the destination path is empty, the result is the source path.
	   The last path of the result is closed only if the last path of the source is.*/
	PathAddPathMode_Extend PathAddPathMode = 1
)

/*
SegmentMask constants correspond to each drawing Verb type in SkPath; for
instance, if SkPath only contains lines, only the kLine_SegmentMask bit is set.
*/
type PathSegmentMask int64

const (
	PathSegmentMask_Line  PathSegmentMask = 1
	PathSegmentMask_Quad  PathSegmentMask = 2
	PathSegmentMask_Conic PathSegmentMask = 4
	PathSegmentMask_Cubic PathSegmentMask = 8
)

/*
Verb instructs SkPath how to interpret one or more SkPoint and optional conic weight;
manage contour, and terminate SkPath.
*/
type PathVerb int64

const (
	PathVerb_Move  PathVerb = 0
	PathVerb_Line  PathVerb = 1
	PathVerb_Quad  PathVerb = 2
	PathVerb_Conic PathVerb = 3
	PathVerb_Cubic PathVerb = 4
	PathVerb_Close PathVerb = 5
	PathVerb_Done  PathVerb = 6
)

/*
SkPathEffect is the base class for objects in the SkPaint that affect
the geometry of a drawing primitive before it is transformed by the
canvas' matrix and drawn.

Dashing is implemented as a subclass of SkPathEffect.
*/
type PathEffect struct {
	skia unsafe.Pointer
}

type PathEffectDashType int64

const (
	/*
	   ignores the info parameter*/
	PathEffectDashType_None PathEffectDashType = 0
	/*
	   fills in all of the info parameter*/
	PathEffectDashType_Dash PathEffectDashType = 1
)

/*
This path effect chops a path into discrete segments, and randomly displaces them.
*/
type DiscretePathEffect struct {
	skia unsafe.Pointer
}

/*
Shaders specify the source color(s) for what is being drawn. If a paint
has no shader, then the paint's color is used. If the paint has a
shader, then the shader's color(s) are use instead, but they are
modulated by the paint's alpha. This makes it easy to create a shader
once (e.g. bitmap tiling or gradient) and then change its transparency
w/o having to modify the original shader... only the paint's alpha needs
to be modified.
*/
type Shader struct {
	skia unsafe.Pointer
}

/*
SkGradientShader hosts factories for creating subclasses of SkShader that
render linear and radial gradients. In general, degenerate cases should not
produce surprising results, but there are several types of degeneracies:

A linear gradient made from the same two points.
A radial gradient with a radius of zero.
A sweep gradient where the start and end angle are the same.
A two point conical gradient where the two centers and the two radii are
the same.

For any degenerate gradient with a decal tile mode, it will draw empty since the interpolating
region is zero area and the outer region is discarded by the decal mode.

For any degenerate gradient with a repeat or mirror tile mode, it will draw a solid color that
is the average gradient color, since infinitely many repetitions of the gradients will fill the
shape.

For a clamped gradient, every type is well-defined at the limit except for linear gradients. The
radial gradient with zero radius becomes the last color. The sweep gradient draws the sector
from 0 to the provided angle with the first color, with a hardstop switching to the last color.
When the provided angle is 0, this is just the solid last color again. Similarly, the two point
conical gradient becomes a circle filled with the first color, sized to the provided radius,
with a hardstop switching to the last color. When the two radii are both zero, this is just the
solid last color.

As a linear gradient approaches the degenerate case, its shader will approach the appearance of
two half planes, each filled by the first and last colors of the gradient. The planes will be
oriented perpendicular to the vector between the two defining points of the gradient. However,
once they become the same point, Skia cannot reconstruct what that expected orientation is. To
provide a stable and predictable color in this case, Skia just uses the last color as a solid
fill to be similar to many of the other degenerate gradients' behaviors in clamp mode.
*/
type GradientShader struct {
	skia unsafe.Pointer
}

type GradientShaderFlags int64

const (
	/*
	   By default gradients will interpolate their colors in unpremul space
	   and then premultiply each of the results. By setting this flag, the
	   gradients will premultiply their colors first, and then interpolate
	   between them.
	   example: https://fiddle.skia.org/c/
	   _MakeLinear*/
	GradientShaderFlags_InterpolateColorsInPremul GradientShaderFlags = 1
)

/*
SkPicture records drawing commands made to SkCanvas. The command stream may be
played in whole or in part at a later time.

SkPicture is an abstract class. SkPicture may be generated by SkPictureRecorder
or SkDrawable, or from SkPicture previously saved to SkData or SkStream.

SkPicture may contain any SkCanvas drawing command, as well as one or more
SkCanvas matrix or SkCanvas clip. SkPicture has a cull SkRect, which is used as
a bounding box hint. To limit SkPicture bounds, use SkCanvas clip when
recording or drawing SkPicture.
*/
type Picture struct {
	skia unsafe.Pointer
}

func (o *Picture) Delete() {
	C.skia_delete_SkPicture(o.skia)
}

type ImageFilters struct {
	skia unsafe.Pointer
}

type ImageFiltersDither int64

const (
	ImageFiltersDither_No  ImageFiltersDither = 0
	ImageFiltersDither_Yes ImageFiltersDither = -1
)

type ShaderMaskFilter struct {
	skia unsafe.Pointer
}

/*
Applies a table lookup on each of the alpha values in the mask.
Helper methods create some common tables (e.g. gamma, clipping)
*/
type TableMaskFilter struct {
	skia unsafe.Pointer
}

type TrimPathEffect struct {
	skia unsafe.Pointer
}

type TrimPathEffectMode int64

const (
	TrimPathEffectMode_Normal   TrimPathEffectMode = 0
	TrimPathEffectMode_Inverted TrimPathEffectMode = 1
)

/*
Like SkData, SkDataTable holds an immutable data buffer. The data buffer is
organized into a table of entries, each with a length, so the entries are
not required to all be the same size.
*/
type DataTable struct {
	skia unsafe.Pointer
}

type Encoder struct {
	skia unsafe.Pointer
}

func (o *Encoder) Delete() {
	C.skia_delete_SkEncoder(o.skia)
}

/*
This helper queries the current GL context for its extensions, remembers them, and can be
queried. It supports both glGetString- and glGetStringi-style extension string APIs and will
use the latter if it is available. It also will query for EGL extensions if a eglQueryString
implementation is provided.
*/
type GrGLExtensions struct {
	skia unsafe.Pointer
}

func NewGrGLExtensions() GrGLExtensions {

	return GrGLExtensions{
		skia: C.skia_new_GrGLExtensions(),
	}
}

/*
Describes a GrColorType by how many bits are used for each color component and how they are
encoded. Currently all the non-zero channels share a single GrColorTypeEncoding. This could be
expanded to store separate encodings and to indicate which bits belong to which components.
*/
type GrColorFormatDesc struct {
	skia unsafe.Pointer
}

type GrBackendFormat struct {
	skia unsafe.Pointer
}

func NewGrBackendFormat() GrBackendFormat {

	return GrBackendFormat{
		skia: C.skia_new_GrBackendFormat(),
	}
}

func (o *GrBackendFormat) Delete() {
	C.skia_delete_GrBackendFormat(o.skia)
}

type GrBackendTexture struct {
	skia unsafe.Pointer
}

func NewGrBackendTexture() GrBackendTexture {

	return GrBackendTexture{
		skia: C.skia_new_GrBackendTexture(),
	}
}

func (o *GrBackendTexture) Delete() {
	C.skia_delete_GrBackendTexture(o.skia)
}

type GrBackendRenderTarget struct {
	skia unsafe.Pointer
}

func NewGrBackendRenderTarget() GrBackendRenderTarget {

	return GrBackendRenderTarget{
		skia: C.skia_new_GrBackendRenderTarget(),
	}
}

func (o *GrBackendRenderTarget) Delete() {
	C.skia_delete_GrBackendRenderTarget(o.skia)
}

type GrDriverBugWorkarounds struct {
	skia unsafe.Pointer
}

func NewGrDriverBugWorkarounds() GrDriverBugWorkarounds {

	return GrDriverBugWorkarounds{
		skia: C.skia_new_GrDriverBugWorkarounds(),
	}
}

func (o *GrDriverBugWorkarounds) Delete() {
	C.skia_delete_GrDriverBugWorkarounds(o.skia)
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type GrContext_Base struct {
	skia unsafe.Pointer
}

func (o *GrContext_Base) Delete() {
	C.skia_delete_GrContext_Base(o.skia)
}

type Semaphore struct {
	skia unsafe.Pointer
}

func NewSemaphore(count int32) Semaphore {
	c_count := C.int(count)

	return Semaphore{
		skia: C.skia_new_SkSemaphore(c_count),
	}
}

func (o *Semaphore) Delete() {
	C.skia_delete_SkSemaphore(o.skia)
}

type Mutex struct {
	skia unsafe.Pointer
}

func NewMutex() Mutex {

	return Mutex{
		skia: C.skia_new_SkMutex(),
	}
}

func (o *Mutex) Delete() {
	C.skia_delete_SkMutex(o.skia)
}

type AutoMutexExclusive struct {
	skia unsafe.Pointer
}

func (o *AutoMutexExclusive) Delete() {
	C.skia_delete_SkAutoMutexExclusive(o.skia)
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type GrImageContext struct {
	skia unsafe.Pointer
}

func (o *GrImageContext) Delete() {
	C.skia_delete_GrImageContext(o.skia)
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type GrRecordingContext struct {
	skia unsafe.Pointer
}

func (o *GrRecordingContext) Delete() {
	C.skia_delete_GrRecordingContext(o.skia)
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type GrDirectContext struct {
	skia unsafe.Pointer
}

func (o *GrDirectContext) Delete() {
	C.skia_delete_GrDirectContext(o.skia)
}

type TDStorage struct {
	skia unsafe.Pointer
}

func NewTDStorage(sizeOfT int32) TDStorage {
	c_sizeOfT := C.int(sizeOfT)

	return TDStorage{
		skia: C.skia_new_SkTDStorage(c_sizeOfT),
	}
}

func (o *TDStorage) Delete() {
	C.skia_delete_SkTDStorage(o.skia)
}

/*
Perform a series of path operations, optimized for unioning many paths together.
*/
type OpBuilder struct {
	skia unsafe.Pointer
}

type ParsePath struct {
	skia unsafe.Pointer
}

type ParsePathPathEncoding int64

const (
	ParsePathPathEncoding_Absolute ParsePathPathEncoding = 0
	ParsePathPathEncoding_Relative ParsePathPathEncoding = 1
)

type PDFBitmap struct {
	skia unsafe.Pointer
}

/*
A SkPDFUnion is a non-virtualized implementation of the
non-compound, non-specialized PDF Object types: Name, String,
Number, Boolean.
*/
type PDFUnion struct {
	skia unsafe.Pointer
}

func (o *PDFUnion) Delete() {
	C.skia_delete_SkPDFUnion(o.skia)
}

type PDFUnionType int64

const (
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Destroyed PDFUnionType = 0
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Int PDFUnionType = 1
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_ColorComponent PDFUnionType = 2
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_ColorComponentF PDFUnionType = 3
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Bool PDFUnionType = 4
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Scalar PDFUnionType = 5
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Name PDFUnionType = 6
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_ByteString PDFUnionType = 7
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_TextString PDFUnionType = 8
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_NameSkS PDFUnionType = 9
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_ByteStringSkS PDFUnionType = 10
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_TextStringSkS PDFUnionType = 11
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Object PDFUnionType = 12
	/*
	   It is an error to call emitObject() or addResources() on an kDestroyed object.*/
	PDFUnionType_Ref PDFUnionType = 13
)

/*
A PDF Object is the base class for primitive elements in a PDF file.  A
common subtype is used to ease the use of indirect object references,
which are common in the PDF format.
*/
type PDFObject struct {
	skia unsafe.Pointer
}

func (o *PDFObject) Delete() {
	C.skia_delete_SkPDFObject(o.skia)
}

/*
An array object in a PDF.
*/
type PDFArray struct {
	skia unsafe.Pointer
}

/*
Create a PDF array. Maximum length is 8191.*/
func NewPDFArray() PDFArray {

	return PDFArray{
		skia: C.skia_new_SkPDFArray(),
	}
}

func (o *PDFArray) Delete() {
	C.skia_delete_SkPDFArray(o.skia)
}

/*
A dictionary object in a PDF.
*/
type PDFDict struct {
	skia unsafe.Pointer
}

func (o *PDFDict) Delete() {
	C.skia_delete_SkPDFDict(o.skia)
}

type ArenaAlloc struct {
	skia unsafe.Pointer
}

func NewArenaAlloc(firstHeapAllocation uint64) ArenaAlloc {
	c_firstHeapAllocation := C.ulong(firstHeapAllocation)

	return ArenaAlloc{
		skia: C.skia_new_SkArenaAlloc(c_firstHeapAllocation),
	}
}

func (o *ArenaAlloc) Delete() {
	C.skia_delete_SkArenaAlloc(o.skia)
}

type ArenaAllocWithReset struct {
	skia unsafe.Pointer
}

func NewArenaAllocWithReset(firstHeapAllocation uint64) ArenaAllocWithReset {
	c_firstHeapAllocation := C.ulong(firstHeapAllocation)

	return ArenaAllocWithReset{
		skia: C.skia_new_SkArenaAllocWithReset(c_firstHeapAllocation),
	}
}

type PDFTagTree struct {
	skia unsafe.Pointer
}

func NewPDFTagTree() PDFTagTree {

	return PDFTagTree{
		skia: C.skia_new_SkPDFTagTree(),
	}
}

func (o *PDFTagTree) Delete() {
	C.skia_delete_SkPDFTagTree(o.skia)
}

type PDFOffsetMap struct {
	skia unsafe.Pointer
}

/*
Concrete implementation of SkDocument that creates PDF files. This
class does not produced linearized or optimized PDFs; instead it
it attempts to use a minimum amount of RAM.
*/
type PDFDocument struct {
	skia unsafe.Pointer
}

func (o *PDFDocument) Delete() {
	C.skia_delete_SkPDFDocument(o.skia)
}

type Shaper struct {
	skia unsafe.Pointer
}

func (o *Shaper) Delete() {
	C.skia_delete_SkShaper(o.skia)
}

/*
Helper for shaping text directly into a SkTextBlob.
*/
type TextBlobBuilderRunHandler struct {
	skia unsafe.Pointer
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type SVGDOM struct {
	skia unsafe.Pointer
}

type SVGLength struct {
	skia unsafe.Pointer
}

func NewSVGLength() SVGLength {

	return SVGLength{
		skia: C.skia_new_SkSVGLength(),
	}
}

func NewSVGLength2(v float32, u SVGLengthUnit) SVGLength {
	c_v := C.float(v)
	c_u := C.int(u)

	return SVGLength{
		skia: C.skia_new_SkSVGLength2(c_v, c_u),
	}
}

type SVGLengthUnit int64

const (
	SVGLengthUnit_Unknown    SVGLengthUnit = 0
	SVGLengthUnit_Number     SVGLengthUnit = 1
	SVGLengthUnit_Percentage SVGLengthUnit = 2
	SVGLengthUnit_EMS        SVGLengthUnit = 3
	SVGLengthUnit_EXS        SVGLengthUnit = 4
	SVGLengthUnit_PX         SVGLengthUnit = 5
	SVGLengthUnit_CM         SVGLengthUnit = 6
	SVGLengthUnit_MM         SVGLengthUnit = 7
	SVGLengthUnit_IN         SVGLengthUnit = 8
	SVGLengthUnit_PT         SVGLengthUnit = 9
	SVGLengthUnit_PC         SVGLengthUnit = 10
)

type SVGIRI struct {
	skia unsafe.Pointer
}

func NewSVGIRI() SVGIRI {

	return SVGIRI{
		skia: C.skia_new_SkSVGIRI(),
	}
}

type SVGIRIType int64

const (
	SVGIRIType_Local    SVGIRIType = 0
	SVGIRIType_Nonlocal SVGIRIType = 1
	SVGIRIType_DataURI  SVGIRIType = 2
)

type SVGColor struct {
	skia unsafe.Pointer
}

func NewSVGColor() SVGColor {

	return SVGColor{
		skia: C.skia_new_SkSVGColor(),
	}
}

type SVGColorType int64

const (
	SVGColorType_CurrentColor SVGColorType = 0
	SVGColorType_Color        SVGColorType = 1
	SVGColorType_ICCColor     SVGColorType = 2
)

type SVGPaint struct {
	skia unsafe.Pointer
}

func NewSVGPaint() SVGPaint {

	return SVGPaint{
		skia: C.skia_new_SkSVGPaint(),
	}
}

func NewSVGPaint2(t SVGPaintType) SVGPaint {
	c_t := C.int(t)

	return SVGPaint{
		skia: C.skia_new_SkSVGPaint2(c_t),
	}
}

type SVGPaintType int64

const (
	SVGPaintType_None  SVGPaintType = 0
	SVGPaintType_Color SVGPaintType = 1
	SVGPaintType_IRI   SVGPaintType = 2
)

type SVGFuncIRI struct {
	skia unsafe.Pointer
}

func NewSVGFuncIRI() SVGFuncIRI {

	return SVGFuncIRI{
		skia: C.skia_new_SkSVGFuncIRI(),
	}
}

func NewSVGFuncIRI2(t SVGFuncIRIType) SVGFuncIRI {
	c_t := C.int(t)

	return SVGFuncIRI{
		skia: C.skia_new_SkSVGFuncIRI2(c_t),
	}
}

type SVGFuncIRIType int64

const (
	SVGFuncIRIType_None SVGFuncIRIType = 0
	SVGFuncIRIType_IRI  SVGFuncIRIType = 1
)

type SVGLineJoin struct {
	skia unsafe.Pointer
}

func NewSVGLineJoin() SVGLineJoin {

	return SVGLineJoin{
		skia: C.skia_new_SkSVGLineJoin(),
	}
}

func NewSVGLineJoin2(t SVGLineJoinType) SVGLineJoin {
	c_t := C.int(t)

	return SVGLineJoin{
		skia: C.skia_new_SkSVGLineJoin2(c_t),
	}
}

type SVGLineJoinType int64

const (
	SVGLineJoinType_Miter   SVGLineJoinType = 0
	SVGLineJoinType_Round   SVGLineJoinType = 1
	SVGLineJoinType_Bevel   SVGLineJoinType = 2
	SVGLineJoinType_Inherit SVGLineJoinType = 3
)

type SVGSpreadMethod struct {
	skia unsafe.Pointer
}

func NewSVGSpreadMethod() SVGSpreadMethod {

	return SVGSpreadMethod{
		skia: C.skia_new_SkSVGSpreadMethod(),
	}
}

func NewSVGSpreadMethod2(t SVGSpreadMethodType) SVGSpreadMethod {
	c_t := C.int(t)

	return SVGSpreadMethod{
		skia: C.skia_new_SkSVGSpreadMethod2(c_t),
	}
}

type SVGSpreadMethodType int64

const (
	SVGSpreadMethodType_Pad     SVGSpreadMethodType = 0
	SVGSpreadMethodType_Repeat  SVGSpreadMethodType = 1
	SVGSpreadMethodType_Reflect SVGSpreadMethodType = 2
)

type SVGFillRule struct {
	skia unsafe.Pointer
}

func NewSVGFillRule() SVGFillRule {

	return SVGFillRule{
		skia: C.skia_new_SkSVGFillRule(),
	}
}

func NewSVGFillRule2(t SVGFillRuleType) SVGFillRule {
	c_t := C.int(t)

	return SVGFillRule{
		skia: C.skia_new_SkSVGFillRule2(c_t),
	}
}

type SVGFillRuleType int64

const (
	SVGFillRuleType_NonZero SVGFillRuleType = 0
	SVGFillRuleType_EvenOdd SVGFillRuleType = 1
	SVGFillRuleType_Inherit SVGFillRuleType = 2
)

type SVGVisibility struct {
	skia unsafe.Pointer
}

func NewSVGVisibility() SVGVisibility {

	return SVGVisibility{
		skia: C.skia_new_SkSVGVisibility(),
	}
}

func NewSVGVisibility2(t SVGVisibilityType) SVGVisibility {
	c_t := C.int(t)

	return SVGVisibility{
		skia: C.skia_new_SkSVGVisibility2(c_t),
	}
}

type SVGVisibilityType int64

const (
	SVGVisibilityType_Visible  SVGVisibilityType = 0
	SVGVisibilityType_Hidden   SVGVisibilityType = 1
	SVGVisibilityType_Collapse SVGVisibilityType = 2
	SVGVisibilityType_Inherit  SVGVisibilityType = 3
)

type SVGDashArray struct {
	skia unsafe.Pointer
}

func NewSVGDashArray() SVGDashArray {

	return SVGDashArray{
		skia: C.skia_new_SkSVGDashArray(),
	}
}

func NewSVGDashArray2(t SVGDashArrayType) SVGDashArray {
	c_t := C.int(t)

	return SVGDashArray{
		skia: C.skia_new_SkSVGDashArray2(c_t),
	}
}

type SVGDashArrayType int64

const (
	SVGDashArrayType_None      SVGDashArrayType = 0
	SVGDashArrayType_DashArray SVGDashArrayType = 1
	SVGDashArrayType_Inherit   SVGDashArrayType = 2
)

type SVGStopColor struct {
	skia unsafe.Pointer
}

func NewSVGStopColor() SVGStopColor {

	return SVGStopColor{
		skia: C.skia_new_SkSVGStopColor(),
	}
}

func NewSVGStopColor2(t SVGStopColorType) SVGStopColor {
	c_t := C.int(t)

	return SVGStopColor{
		skia: C.skia_new_SkSVGStopColor2(c_t),
	}
}

type SVGStopColorType int64

const (
	SVGStopColorType_Color        SVGStopColorType = 0
	SVGStopColorType_CurrentColor SVGStopColorType = 1
	SVGStopColorType_ICCColor     SVGStopColorType = 2
	SVGStopColorType_Inherit      SVGStopColorType = 3
)

type SVGObjectBoundingBoxUnits struct {
	skia unsafe.Pointer
}

func NewSVGObjectBoundingBoxUnits() SVGObjectBoundingBoxUnits {

	return SVGObjectBoundingBoxUnits{
		skia: C.skia_new_SkSVGObjectBoundingBoxUnits(),
	}
}

func NewSVGObjectBoundingBoxUnits2(t SVGObjectBoundingBoxUnitsType) SVGObjectBoundingBoxUnits {
	c_t := C.int(t)

	return SVGObjectBoundingBoxUnits{
		skia: C.skia_new_SkSVGObjectBoundingBoxUnits2(c_t),
	}
}

type SVGObjectBoundingBoxUnitsType int64

const (
	SVGObjectBoundingBoxUnitsType_UserSpaceOnUse    SVGObjectBoundingBoxUnitsType = 0
	SVGObjectBoundingBoxUnitsType_ObjectBoundingBox SVGObjectBoundingBoxUnitsType = 1
)

type SVGFontFamily struct {
	skia unsafe.Pointer
}

func NewSVGFontFamily() SVGFontFamily {

	return SVGFontFamily{
		skia: C.skia_new_SkSVGFontFamily(),
	}
}

type SVGFontFamilyType int64

const (
	SVGFontFamilyType_Family  SVGFontFamilyType = 0
	SVGFontFamilyType_Inherit SVGFontFamilyType = 1
)

type SVGFontStyle struct {
	skia unsafe.Pointer
}

func NewSVGFontStyle() SVGFontStyle {

	return SVGFontStyle{
		skia: C.skia_new_SkSVGFontStyle(),
	}
}

func NewSVGFontStyle2(t SVGFontStyleType) SVGFontStyle {
	c_t := C.int(t)

	return SVGFontStyle{
		skia: C.skia_new_SkSVGFontStyle2(c_t),
	}
}

type SVGFontStyleType int64

const (
	SVGFontStyleType_Normal  SVGFontStyleType = 0
	SVGFontStyleType_Italic  SVGFontStyleType = 1
	SVGFontStyleType_Oblique SVGFontStyleType = 2
	SVGFontStyleType_Inherit SVGFontStyleType = 3
)

type SVGFontSize struct {
	skia unsafe.Pointer
}

func NewSVGFontSize() SVGFontSize {

	return SVGFontSize{
		skia: C.skia_new_SkSVGFontSize(),
	}
}

type SVGFontSizeType int64

const (
	SVGFontSizeType_Length  SVGFontSizeType = 0
	SVGFontSizeType_Inherit SVGFontSizeType = 1
)

type SVGFontWeight struct {
	skia unsafe.Pointer
}

func NewSVGFontWeight() SVGFontWeight {

	return SVGFontWeight{
		skia: C.skia_new_SkSVGFontWeight(),
	}
}

func NewSVGFontWeight2(t SVGFontWeightType) SVGFontWeight {
	c_t := C.int(t)

	return SVGFontWeight{
		skia: C.skia_new_SkSVGFontWeight2(c_t),
	}
}

type SVGFontWeightType int64

const (
	SVGFontWeightType_100     SVGFontWeightType = 0
	SVGFontWeightType_200     SVGFontWeightType = 1
	SVGFontWeightType_300     SVGFontWeightType = 2
	SVGFontWeightType_400     SVGFontWeightType = 3
	SVGFontWeightType_500     SVGFontWeightType = 4
	SVGFontWeightType_600     SVGFontWeightType = 5
	SVGFontWeightType_700     SVGFontWeightType = 6
	SVGFontWeightType_800     SVGFontWeightType = 7
	SVGFontWeightType_900     SVGFontWeightType = 8
	SVGFontWeightType_Normal  SVGFontWeightType = 9
	SVGFontWeightType_Bold    SVGFontWeightType = 10
	SVGFontWeightType_Bolder  SVGFontWeightType = 11
	SVGFontWeightType_Lighter SVGFontWeightType = 12
	SVGFontWeightType_Inherit SVGFontWeightType = 13
)

type SVGTextAnchor struct {
	skia unsafe.Pointer
}

func NewSVGTextAnchor() SVGTextAnchor {

	return SVGTextAnchor{
		skia: C.skia_new_SkSVGTextAnchor(),
	}
}

func NewSVGTextAnchor2(t SVGTextAnchorType) SVGTextAnchor {
	c_t := C.int(t)

	return SVGTextAnchor{
		skia: C.skia_new_SkSVGTextAnchor2(c_t),
	}
}

type SVGTextAnchorType int64

const (
	SVGTextAnchorType_Start   SVGTextAnchorType = 0
	SVGTextAnchorType_Middle  SVGTextAnchorType = 1
	SVGTextAnchorType_End     SVGTextAnchorType = 2
	SVGTextAnchorType_Inherit SVGTextAnchorType = 3
)

type SVGFeInputType struct {
	skia unsafe.Pointer
}

func NewSVGFeInputType() SVGFeInputType {

	return SVGFeInputType{
		skia: C.skia_new_SkSVGFeInputType(),
	}
}

func NewSVGFeInputType2(t SVGFeInputTypeType) SVGFeInputType {
	c_t := C.int(t)

	return SVGFeInputType{
		skia: C.skia_new_SkSVGFeInputType2(c_t),
	}
}

type SVGFeInputTypeType int64

const (
	SVGFeInputTypeType_SourceGraphic            SVGFeInputTypeType = 0
	SVGFeInputTypeType_SourceAlpha              SVGFeInputTypeType = 1
	SVGFeInputTypeType_BackgroundImage          SVGFeInputTypeType = 2
	SVGFeInputTypeType_BackgroundAlpha          SVGFeInputTypeType = 3
	SVGFeInputTypeType_FillPaint                SVGFeInputTypeType = 4
	SVGFeInputTypeType_StrokePaint              SVGFeInputTypeType = 5
	SVGFeInputTypeType_FilterPrimitiveReference SVGFeInputTypeType = 6
	SVGFeInputTypeType_Unspecified              SVGFeInputTypeType = 7
)

type SVGFeTurbulenceBaseFrequency struct {
	skia unsafe.Pointer
}

func NewSVGFeTurbulenceBaseFrequency() SVGFeTurbulenceBaseFrequency {

	return SVGFeTurbulenceBaseFrequency{
		skia: C.skia_new_SkSVGFeTurbulenceBaseFrequency(),
	}
}

func NewSVGFeTurbulenceBaseFrequency2(freqX float32, freqY float32) SVGFeTurbulenceBaseFrequency {
	c_freqX := C.float(freqX)
	c_freqY := C.float(freqY)

	return SVGFeTurbulenceBaseFrequency{
		skia: C.skia_new_SkSVGFeTurbulenceBaseFrequency2(c_freqX, c_freqY),
	}
}

/*
SkNoncopyable is the base class for objects that do not want to
be copied. It hides its copy-constructor and its assignment-operator.
*/
type SVGAttributeParser struct {
	skia unsafe.Pointer
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type SVGNode struct {
	skia unsafe.Pointer
}

func (o *SVGNode) Delete() {
	C.skia_delete_SkSVGNode(o.skia)
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type SVGTransformableNode struct {
	skia unsafe.Pointer
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type SVGContainer struct {
	skia unsafe.Pointer
}

/*
SkRefCntBase is the base class for objects that may be shared by multiple
objects. When an existing owner wants to share a reference, it calls ref().
When an owner wants to release its reference, it calls unref(). When the
shared object's reference count goes to zero as the result of an unref()
call, its (virtual) destructor is called. It is an error for the
destructor to be called explicitly (or via the object going out of scope on
the stack or calling delete) if getRefCnt() > 1.
*/
type SVGSVG struct {
	skia unsafe.Pointer
}

type SVGSVGType int64

const (
	SVGSVGType_Root  SVGSVGType = 0
	SVGSVGType_Inner SVGSVGType = 1
)

type SVGLengthContext struct {
	skia unsafe.Pointer
}

type SVGLengthContextLengthType int64

const (
	SVGLengthContextLengthType_Horizontal SVGLengthContextLengthType = 0
	SVGLengthContextLengthType_Vertical   SVGLengthContextLengthType = 1
	SVGLengthContextLengthType_Other      SVGLengthContextLengthType = 2
)

type SVGRenderContext struct {
	skia unsafe.Pointer
}

func (o *SVGRenderContext) Delete() {
	C.skia_delete_SkSVGRenderContext(o.skia)
}

type SVGRenderContextApplyFlags int64

const (
	SVGRenderContextApplyFlags_Leaf SVGRenderContextApplyFlags = 1
)

/*
When we transform points through a matrix containing perspective (the bottom row is something
other than 0,0,1), the bruteforce math can produce confusing results (since we might divide
by 0, or a negative w value). By default, methods that map rects and paths will apply
perspective clipping, but this can be changed by specifying kYes to those methods.
*/
type ApplyPerspectiveClip int64

const (
	/*
	   Don't pre-clip the geometry before applying the (perspective) matrix*/
	ApplyPerspectiveClip_No ApplyPerspectiveClip = 0
	/*
	   Do pre-clip the geometry before applying the (perspective) matrix*/
	ApplyPerspectiveClip_Yes ApplyPerspectiveClip = 1
)

type EncodedOrigin int64

const (
	EncodedOrigin_TopLeft     EncodedOrigin = 1
	EncodedOrigin_TopRight    EncodedOrigin = 2
	EncodedOrigin_BottomRight EncodedOrigin = 3
	EncodedOrigin_BottomLeft  EncodedOrigin = 4
	EncodedOrigin_LeftTop     EncodedOrigin = 5
	EncodedOrigin_RightTop    EncodedOrigin = 6
	EncodedOrigin_RightBottom EncodedOrigin = 7
	EncodedOrigin_LeftBottom  EncodedOrigin = 8
	EncodedOrigin_Default     EncodedOrigin = 1
	EncodedOrigin_Last        EncodedOrigin = 8
)

/*
Describes how to interpret the alpha component of a pixel. A pixel may
be opaque, or alpha, describing multiple levels of transparency.

In simple blending, alpha weights the draw color and the destination
color to create a new color. If alpha describes a weight from zero to one:

new color = draw color * alpha + destination color * (1 - alpha)

In practice alpha is encoded in two or more bits, where 1.0 equals all bits set.

RGB may have alpha included in each component value; the stored
value is the original RGB multiplied by alpha. Premultiplied color
components improve performance.
*/
type AlphaType int64

const (
	/*
	   uninitialized*/
	AlphaType_Unknown AlphaType = 0
	/*
	   pixel is opaque*/
	AlphaType_Opaque AlphaType = 1
	/*
	   pixel components are premultiplied by alpha*/
	AlphaType_Premul AlphaType = 2
	/*
	   pixel components are independent of alpha*/
	AlphaType_Unpremul AlphaType = 3
	/*
	   last valid value*/
	AlphaType_LastEnum AlphaType = 3
)

/*
Describes how pixel bits encode color. A pixel may be an alpha mask, a grayscale, RGB, or ARGB.

kN32_SkColorType selects the native 32-bit ARGB format for the current configuration. This can
lead to inconsistent results across platforms, so use with caution.
*/
type ColorType int64

const (
	/*
	   uninitialized*/
	ColorType_Unknown ColorType = 0
	/*
	   pixel with alpha in 8-bit byte*/
	ColorType_Alpha_8 ColorType = 1
	/*
	   pixel with 5 bits red, 6 bits green, 5 bits blue, in 16-bit word*/
	ColorType_RGB_565 ColorType = 2
	/*
	   pixel with 4 bits for alpha, red, green, blue; in 16-bit word*/
	ColorType_ARGB_4444 ColorType = 3
	/*
	   pixel with 8 bits for red, green, blue, alpha; in 32-bit word*/
	ColorType_RGBA_8888 ColorType = 4
	/*
	   pixel with 8 bits each for red, green, blue; in 32-bit word*/
	ColorType_RGB_888x ColorType = 5
	/*
	   pixel with 8 bits for blue, green, red, alpha; in 32-bit word*/
	ColorType_BGRA_8888 ColorType = 6
	/*
	   10 bits for red, green, blue; 2 bits for alpha; in 32-bit word*/
	ColorType_RGBA_1010102 ColorType = 7
	/*
	   10 bits for blue, green, red; 2 bits for alpha; in 32-bit word*/
	ColorType_BGRA_1010102 ColorType = 8
	/*
	   pixel with 10 bits each for red, green, blue; in 32-bit word*/
	ColorType_RGB_101010x ColorType = 9
	/*
	   pixel with 10 bits each for blue, green, red; in 32-bit word*/
	ColorType_BGR_101010x ColorType = 10
	/*
	   pixel with 10 bits each for blue, green, red; in 32-bit word, extended range*/
	ColorType_BGR_101010x_XR ColorType = 11
	/*
	   pixel with 10 bits each for blue, green, red, alpha; in 64-bit word, extended range*/
	ColorType_BGRA_10101010_XR ColorType = 12
	/*
	   pixel with 10 used bits (most significant) followed by 6 unused*/
	ColorType_RGBA_10x6 ColorType = 13
	/*
	   pixel with grayscale level in 8-bit byte*/
	ColorType_Gray_8 ColorType = 14
	/*
	   pixel with half floats in [0,1] for red, green, blue, alpha;*/
	ColorType_RGBA_F16Norm ColorType = 15
	/*
	   pixel with half floats for red, green, blue, alpha;*/
	ColorType_RGBA_F16 ColorType = 16
	/*
	   pixel using C float for red, green, blue, alpha; in 128-bit word*/
	ColorType_RGBA_F32 ColorType = 17
	/*
	   pixel with a uint8_t for red and green*/
	ColorType_R8G8_unorm ColorType = 18
	/*
	   pixel with a half float for alpha*/
	ColorType_A16_float ColorType = 19
	/*
	   pixel with a half float for red and green*/
	ColorType_R16G16_float ColorType = 20
	/*
	   pixel with a little endian uint16_t for alpha*/
	ColorType_A16_unorm ColorType = 21
	/*
	   pixel with a little endian uint16_t for red and green*/
	ColorType_R16G16_unorm ColorType = 22
	/*
	   pixel with a little endian uint16_t for red, green, blue*/
	ColorType_R16G16B16A16_unorm ColorType = 23
	ColorType_SRGBA_8888         ColorType = 24
	ColorType_R8_unorm           ColorType = 25
	/*
	   last valid value*/
	ColorType_LastEnum ColorType = 25
	/*
	   native 32-bit RGBA encoding*/
	ColorType_N32 ColorType = 4
)

/*
Describes color range of YUV pixels. The color mapping from YUV to RGB varies
depending on the source. YUV pixels may be generated by JPEG images, standard
video streams, or high definition video streams. Each has its own mapping from
YUV to RGB.

JPEG YUV values encode the full range of 0 to 255 for all three components.
Video YUV values often range from 16 to 235 for Y and from 16 to 240 for U and V (limited).
Details of encoding and conversion to RGB are described in YCbCr color space.

The identity colorspace exists to provide a utility mapping from Y to R, U to G and V to B.
It can be used to visualize the YUV planes or to explicitly post process the YUV channels.
*/
type YUVColorSpace int64

const (
	/*
	   describes full range*/
	YUVColorSpace_JPEG_Full YUVColorSpace = 0
	/*
	   describes SDTV range*/
	YUVColorSpace_Rec601_Limited YUVColorSpace = 1
	/*
	   describes HDTV range*/
	YUVColorSpace_Rec709_Full    YUVColorSpace = 2
	YUVColorSpace_Rec709_Limited YUVColorSpace = 3
	/*
	   describes UHDTV range, non-constant-luminance*/
	YUVColorSpace_BT2020_8bit_Full     YUVColorSpace = 4
	YUVColorSpace_BT2020_8bit_Limited  YUVColorSpace = 5
	YUVColorSpace_BT2020_10bit_Full    YUVColorSpace = 6
	YUVColorSpace_BT2020_10bit_Limited YUVColorSpace = 7
	YUVColorSpace_BT2020_12bit_Full    YUVColorSpace = 8
	YUVColorSpace_BT2020_12bit_Limited YUVColorSpace = 9
	/*
	   describes FCC range*/
	YUVColorSpace_FCC_Full    YUVColorSpace = 10
	YUVColorSpace_FCC_Limited YUVColorSpace = 11
	/*
	   describes SMPTE240M range*/
	YUVColorSpace_SMPTE240_Full    YUVColorSpace = 12
	YUVColorSpace_SMPTE240_Limited YUVColorSpace = 13
	/*
	   describes YDZDX range*/
	YUVColorSpace_YDZDX_Full    YUVColorSpace = 14
	YUVColorSpace_YDZDX_Limited YUVColorSpace = 15
	/*
	   describes GBR range*/
	YUVColorSpace_GBR_Full    YUVColorSpace = 16
	YUVColorSpace_GBR_Limited YUVColorSpace = 17
	/*
	   describes YCgCo matrix*/
	YUVColorSpace_YCgCo_8bit_Full     YUVColorSpace = 18
	YUVColorSpace_YCgCo_8bit_Limited  YUVColorSpace = 19
	YUVColorSpace_YCgCo_10bit_Full    YUVColorSpace = 20
	YUVColorSpace_YCgCo_10bit_Limited YUVColorSpace = 21
	YUVColorSpace_YCgCo_12bit_Full    YUVColorSpace = 22
	YUVColorSpace_YCgCo_12bit_Limited YUVColorSpace = 23
	/*
	   maps Y->R, U->G, V->B*/
	YUVColorSpace_Identity YUVColorSpace = 24
	/*
	   last valid value*/
	YUVColorSpace_LastEnum YUVColorSpace = 24
	YUVColorSpace_JPEG     YUVColorSpace = 0
	YUVColorSpace_Rec601   YUVColorSpace = 1
	YUVColorSpace_Rec709   YUVColorSpace = 3
	YUVColorSpace_BT2020   YUVColorSpace = 5
)

/*
Describes different color channels one can manipulate
*/
type ColorChannel int64

const (
	ColorChannel_R        ColorChannel = 0
	ColorChannel_G        ColorChannel = 1
	ColorChannel_B        ColorChannel = 2
	ColorChannel_A        ColorChannel = 3
	ColorChannel_LastEnum ColorChannel = 3
)

/*
Used to represent the channels available in a color type or texture format as a mask.
*/
type ColorChannelFlag int64

const (
	ColorChannelFlag_Red                           ColorChannelFlag = 1
	ColorChannelFlag_Green                         ColorChannelFlag = 2
	ColorChannelFlag_Blue                          ColorChannelFlag = 4
	ColorChannelFlag_Alpha                         ColorChannelFlag = 8
	ColorChannelFlag_Gray                          ColorChannelFlag = 16
	ColorChannelFlag_GrayAlpha_SkColorChannelFlags ColorChannelFlag = 24
	ColorChannelFlag_RG_SkColorChannelFlags        ColorChannelFlag = 3
	ColorChannelFlag_RGB_SkColorChannelFlags       ColorChannelFlag = 7
	ColorChannelFlag_RGBA_SkColorChannelFlags      ColorChannelFlag = 15
)

type FilterMode int64

const (
	FilterMode_Nearest FilterMode = 0
	FilterMode_Linear  FilterMode = 1
	FilterMode_Last    FilterMode = 1
)

type MipmapMode int64

const (
	MipmapMode_None    MipmapMode = 0
	MipmapMode_Nearest MipmapMode = 1
	MipmapMode_Linear  MipmapMode = 2
	MipmapMode_Last    MipmapMode = 2
)

type BlurStyle int64

const (
	/*
	   fuzzy inside and outside*/
	BlurStyle_Normal BlurStyle = 0
	/*
	   solid inside, fuzzy outside*/
	BlurStyle_Solid BlurStyle = 1
	/*
	   nothing inside, fuzzy outside*/
	BlurStyle_Outer BlurStyle = 2
	/*
	   fuzzy inside, nothing outside*/
	BlurStyle_Inner    BlurStyle = 3
	BlurStyle_LastEnum BlurStyle = 3
)

/*
Blends are operators that take in two colors (source, destination) and return a new color.
Many of these operate the same on all 4 components: red, green, blue, alpha. For these,
we just document what happens to one component, rather than naming each one separately.

Different SkColorTypes have different representations for color components:
8-bit: 0..255
6-bit: 0..63
5-bit: 0..31
4-bit: 0..15
floats: 0...1

The documentation is expressed as if the component values are always 0..1 (floats).

For brevity, the documentation uses the following abbreviations
s  : source
d  : destination
sa : source alpha
da : destination alpha

Results are abbreviated
r  : if all 4 components are computed in the same manner
ra : result alpha component
rc : result "color": red, green, blue components
*/
type BlendMode int64

const (
	/*
	   r = 0*/
	BlendMode_Clear BlendMode = 0
	/*
	   r = s*/
	BlendMode_Src BlendMode = 1
	/*
	   r = d*/
	BlendMode_Dst BlendMode = 2
	/*
	   r = s + (1-sa)*d*/
	BlendMode_SrcOver BlendMode = 3
	/*
	   r = d + (1-da)*s*/
	BlendMode_DstOver BlendMode = 4
	/*
	   r = s * da*/
	BlendMode_SrcIn BlendMode = 5
	/*
	   r = d * sa*/
	BlendMode_DstIn BlendMode = 6
	/*
	   r = s * (1-da)*/
	BlendMode_SrcOut BlendMode = 7
	/*
	   r = d * (1-sa)*/
	BlendMode_DstOut BlendMode = 8
	/*
	   r = s*da + d*(1-sa)*/
	BlendMode_SrcATop BlendMode = 9
	/*
	   r = d*sa + s*(1-da)*/
	BlendMode_DstATop BlendMode = 10
	/*
	   r = s*(1-da) + d*(1-sa)*/
	BlendMode_Xor BlendMode = 11
	/*
	   r = min(s + d, 1)*/
	BlendMode_Plus BlendMode = 12
	/*
	   r = s*d*/
	BlendMode_Modulate BlendMode = 13
	/*
	   r = s + d - s*d*/
	BlendMode_Screen BlendMode = 14
	/*
	   multiply or screen, depending on destination*/
	BlendMode_Overlay BlendMode = 15
	/*
	   rc = s + d - max(s*da, d*sa), ra = kSrcOver*/
	BlendMode_Darken BlendMode = 16
	/*
	   rc = s + d - min(s*da, d*sa), ra = kSrcOver*/
	BlendMode_Lighten BlendMode = 17
	/*
	   brighten destination to reflect source*/
	BlendMode_ColorDodge BlendMode = 18
	/*
	   darken destination to reflect source*/
	BlendMode_ColorBurn BlendMode = 19
	/*
	   multiply or screen, depending on source*/
	BlendMode_HardLight BlendMode = 20
	/*
	   lighten or darken, depending on source*/
	BlendMode_SoftLight BlendMode = 21
	/*
	   rc = s + d - 2*(min(s*da, d*sa)), ra = kSrcOver*/
	BlendMode_Difference BlendMode = 22
	/*
	   rc = s + d - two(s*d), ra = kSrcOver*/
	BlendMode_Exclusion BlendMode = 23
	/*
	   r = s*(1-da) + d*(1-sa) + s*d*/
	BlendMode_Multiply BlendMode = 24
	/*
	   hue of source with saturation and luminosity of destination*/
	BlendMode_Hue BlendMode = 25
	/*
	   saturation of source with hue and luminosity of destination*/
	BlendMode_Saturation BlendMode = 26
	/*
	   hue and saturation of source with luminosity of destination*/
	BlendMode_Color BlendMode = 27
	/*
	   luminosity of source with hue and saturation of destination*/
	BlendMode_Luminosity BlendMode = 28
	/*
	   last porter duff blend mode*/
	BlendMode_LastCoeffMode BlendMode = 14
	/*
	   last blend mode operating separately on components*/
	BlendMode_LastSeparableMode BlendMode = 24
	/*
	   last valid value*/
	BlendMode_LastMode BlendMode = 28
)

/*
For Porter-Duff SkBlendModes (those
<
= kLastCoeffMode), these coefficients describe the blend
equation used. Coefficient-based blend modes specify an equation:
('dstCoeff' * dst + 'srcCoeff' * src), where the coefficient values are constants, functions of
the src or dst alpha, or functions of the src or dst color.
*/
type BlendModeCoeff int64

const (
	BlendModeCoeff_Zero BlendModeCoeff = 0
	/*
	   0*/
	BlendModeCoeff_One BlendModeCoeff = 1
	/*
	   1*/
	BlendModeCoeff_SC BlendModeCoeff = 2
	/*
	   src color*/
	BlendModeCoeff_ISC BlendModeCoeff = 3
	/*
	   inverse src color (i.e. 1 - sc)*/
	BlendModeCoeff_DC BlendModeCoeff = 4
	/*
	   dst color*/
	BlendModeCoeff_IDC BlendModeCoeff = 5
	/*
	   inverse dst color (i.e. 1 - dc)*/
	BlendModeCoeff_SA BlendModeCoeff = 6
	/*
	   src alpha*/
	BlendModeCoeff_ISA BlendModeCoeff = 7
	/*
	   inverse src alpha (i.e. 1 - sa)*/
	BlendModeCoeff_DA BlendModeCoeff = 8
	/*
	   dst alpha*/
	BlendModeCoeff_IDA BlendModeCoeff = 9
	/*
	   inverse dst alpha (i.e. 1 - da)*/
	BlendModeCoeff_CoeffCount BlendModeCoeff = 10
)

type ClipOp int64

const (
	ClipOp_Difference    ClipOp = 0
	ClipOp_Intersect     ClipOp = 1
	ClipOp_Max_EnumValue ClipOp = 1
)

type TextEncoding int64

const (
	/*
	   uses bytes to represent UTF-8 or ASCII*/
	TextEncoding_UTF8 TextEncoding = 0
	/*
	   uses two byte words to represent most of Unicode*/
	TextEncoding_UTF16 TextEncoding = 1
	/*
	   uses four byte words to represent all of Unicode*/
	TextEncoding_UTF32 TextEncoding = 2
	/*
	   uses two byte words to represent glyph indices*/
	TextEncoding_GlyphID TextEncoding = 3
)

type FontHinting int64

const (
	/*
	   glyph outlines unchanged*/
	FontHinting_None FontHinting = 0
	/*
	   minimal modification to improve constrast*/
	FontHinting_Slight FontHinting = 1
	/*
	   glyph outlines modified to improve constrast*/
	FontHinting_Normal FontHinting = 2
	/*
	   modifies glyph outlines for maximum constrast*/
	FontHinting_Full FontHinting = 3
)

/*
Description of how the LCD strips are arranged for each pixel. If this is unknown, or the
pixels are meant to be "portable" and/or transformed before showing (e.g. rotated, scaled)
then use kUnknown_SkPixelGeometry.
*/
type PixelGeometry int64

const (
	PixelGeometry_Unknown PixelGeometry = 0
	PixelGeometry_RGB_H   PixelGeometry = 1
	PixelGeometry_BGR_H   PixelGeometry = 2
	PixelGeometry_RGB_V   PixelGeometry = 3
	PixelGeometry_BGR_V   PixelGeometry = 4
)

type PathFillType int64

const (
	/*
	   Specifies that "inside" is computed by a non-zero sum of signed edge crossings*/
	PathFillType_Winding PathFillType = 0
	/*
	   Specifies that "inside" is computed by an odd number of edge crossings*/
	PathFillType_EvenOdd PathFillType = 1
	/*
	   Same as Winding, but draws outside of the path, rather than inside*/
	PathFillType_InverseWinding PathFillType = 2
	/*
	   Same as EvenOdd, but draws outside of the path, rather than inside*/
	PathFillType_InverseEvenOdd PathFillType = 3
)

type PathDirection int64

const (
	/*
	   clockwise direction for adding closed contours*/
	PathDirection_CW PathDirection = 0
	/*
	   counter-clockwise direction for adding closed contours*/
	PathDirection_CCW PathDirection = 1
)

type TileMode int64

const (
	/*
	   Replicate the edge color if the shader draws outside of its
	   original bounds.*/
	TileMode_Clamp TileMode = 0
	/*
	   Repeat the shader's image horizontally and vertically.*/
	TileMode_Repeat TileMode = 1
	/*
	   Repeat the shader's image horizontally and vertically, alternating
	   mirror images so that adjacent images always seam.*/
	TileMode_Mirror TileMode = 2
	/*
	   Only draw within the original domain, return transparent-black everywhere else.*/
	TileMode_Decal TileMode = 3
	/*
	   Only draw within the original domain, return transparent-black everywhere else.*/
	TileMode_LastTileMode TileMode = 3
)

/*
Possible 3D APIs that may be used by Ganesh.
*/
type GrBackendApi int64

const (
	GrBackendApi_OpenGL   GrBackendApi = 0
	GrBackendApi_Vulkan   GrBackendApi = 1
	GrBackendApi_Metal    GrBackendApi = 2
	GrBackendApi_Direct3D GrBackendApi = 3
	/*
	   Mock is a backend that does not draw anything. It is used for unit tests
	   and to measure CPU overhead.*/
	GrBackendApi_Mock GrBackendApi = 4
	/*
	   Ganesh doesn't support some context types (e.g. Dawn) and will return Unsupported.*/
	GrBackendApi_Unsupported GrBackendApi = 5
	/*
	   Added here to support the legacy GrBackend enum value and clients who referenced it using
	   GrBackend::kOpenGL_GrBackend.*/
	GrBackendApi_OpenGL_GrBackend GrBackendApi = 0
)

/*
GPU SkImage and SkSurfaces can be stored such that (0, 0) in texture space may correspond to
either the top-left or bottom-left content pixel.
*/
type GrSurfaceOrigin int64

const (
	GrSurfaceOrigin_TopLeft    GrSurfaceOrigin = 0
	GrSurfaceOrigin_BottomLeft GrSurfaceOrigin = 1
)

/*
A GrContext's cache of backend context state can be partially invalidated.
These enums are specific to the GL backend and we'd add a new set for an alternative backend.
*/
type GrGLBackendState int64

const (
	GrGLBackendState_RenderTarget   GrGLBackendState = 1
	GrGLBackendState_TextureBinding GrGLBackendState = 2
	GrGLBackendState_View           GrGLBackendState = 4
	GrGLBackendState_Blend          GrGLBackendState = 8
	GrGLBackendState_MSAAEnable     GrGLBackendState = 16
	GrGLBackendState_Vertex         GrGLBackendState = 32
	GrGLBackendState_Stencil        GrGLBackendState = 64
	GrGLBackendState_PixelStore     GrGLBackendState = 128
	GrGLBackendState_Program        GrGLBackendState = 256
	GrGLBackendState_FixedFunction  GrGLBackendState = 512
	GrGLBackendState_Misc           GrGLBackendState = 1024
	GrGLBackendState_ALL            GrGLBackendState = 65535
)

/*
Enum used as return value when flush with semaphores so the client knows whether the valid
semaphores will be submitted on the next GrContext::submit call.
*/
type GrSemaphoresSubmitted int64

const (
	GrSemaphoresSubmitted_No  GrSemaphoresSubmitted = 0
	GrSemaphoresSubmitted_Yes GrSemaphoresSubmitted = -1
)

type GrPurgeResourceOptions int64

const (
	GrPurgeResourceOptions_AllResources         GrPurgeResourceOptions = 0
	GrPurgeResourceOptions_ScratchResourcesOnly GrPurgeResourceOptions = 1
)

type GrSyncCpu int64

const (
	GrSyncCpu_No  GrSyncCpu = 0
	GrSyncCpu_Yes GrSyncCpu = -1
)

/*
Classifies GL contexts by which standard they implement (currently as OpenGL vs. OpenGL ES).
*/
type GrGLStandard int64

const (
	GrGLStandard_None  GrGLStandard = 0
	GrGLStandard_GL    GrGLStandard = 1
	GrGLStandard_GLES  GrGLStandard = 2
	GrGLStandard_WebGL GrGLStandard = 3
)

/*
The supported GL formats represented as an enum. Actual support by GrContext depends on GL
context version and extensions.
*/
type GrGLFormat int64

const (
	GrGLFormat_Unknown              GrGLFormat = 0
	GrGLFormat_RGBA8                GrGLFormat = 1
	GrGLFormat_R8                   GrGLFormat = 2
	GrGLFormat_ALPHA8               GrGLFormat = 3
	GrGLFormat_LUMINANCE8           GrGLFormat = 4
	GrGLFormat_LUMINANCE8_ALPHA8    GrGLFormat = 5
	GrGLFormat_BGRA8                GrGLFormat = 6
	GrGLFormat_RGB565               GrGLFormat = 7
	GrGLFormat_RGBA16F              GrGLFormat = 8
	GrGLFormat_R16F                 GrGLFormat = 9
	GrGLFormat_RGB8                 GrGLFormat = 10
	GrGLFormat_RGBX8                GrGLFormat = 11
	GrGLFormat_RG8                  GrGLFormat = 12
	GrGLFormat_RGB10_A2             GrGLFormat = 13
	GrGLFormat_RGBA4                GrGLFormat = 14
	GrGLFormat_SRGB8_ALPHA8         GrGLFormat = 15
	GrGLFormat_COMPRESSED_ETC1_RGB8 GrGLFormat = 16
	GrGLFormat_COMPRESSED_RGB8_ETC2 GrGLFormat = 17
	GrGLFormat_COMPRESSED_RGB8_BC1  GrGLFormat = 18
	GrGLFormat_COMPRESSED_RGBA8_BC1 GrGLFormat = 19
	GrGLFormat_R16                  GrGLFormat = 20
	GrGLFormat_RG16                 GrGLFormat = 21
	GrGLFormat_RGBA16               GrGLFormat = 22
	GrGLFormat_RG16F                GrGLFormat = 23
	GrGLFormat_LUMINANCE16F         GrGLFormat = 24
	GrGLFormat_LastColorFormat      GrGLFormat = 24
	GrGLFormat_STENCIL_INDEX8       GrGLFormat = 25
	GrGLFormat_STENCIL_INDEX16      GrGLFormat = 26
	GrGLFormat_DEPTH24_STENCIL8     GrGLFormat = 27
	GrGLFormat_Last                 GrGLFormat = 27
)

type TextureCompressionType int64

const (
	TextureCompressionType_None            TextureCompressionType = 0
	TextureCompressionType_ETC2_RGB8_UNORM TextureCompressionType = 1
	TextureCompressionType_BC1_RGB8_UNORM  TextureCompressionType = 2
	TextureCompressionType_BC1_RGBA8_UNORM TextureCompressionType = 3
	TextureCompressionType_Last            TextureCompressionType = 3
	TextureCompressionType_ETC1_RGB8       TextureCompressionType = 1
)

/*
Geometric primitives used for drawing.
*/
type GrPrimitiveType int64

const (
	GrPrimitiveType_Triangles     GrPrimitiveType = 0
	GrPrimitiveType_TriangleStrip GrPrimitiveType = 1
	GrPrimitiveType_Points        GrPrimitiveType = 2
	GrPrimitiveType_Lines         GrPrimitiveType = 3
	GrPrimitiveType_LineStrip     GrPrimitiveType = 4
)

type GrPrimitiveRestart int64

const (
	GrPrimitiveRestart_No  GrPrimitiveRestart = 0
	GrPrimitiveRestart_Yes GrPrimitiveRestart = -1
)

/*
Should a created surface be texturable?
*/
type GrTexturable int64

const (
	GrTexturable_No  GrTexturable = 0
	GrTexturable_Yes GrTexturable = -1
)

type GrDDLProvider int64

const (
	GrDDLProvider_No  GrDDLProvider = 0
	GrDDLProvider_Yes GrDDLProvider = -1
)

/*
Ownership rules for external GPU resources imported into Skia.
*/
type GrWrapOwnership int64

const (
	/*
	   Skia will assume the client will keep the resource alive and Skia will not free it.*/
	GrWrapOwnership_Borrow GrWrapOwnership = 0
	/*
	   Skia will assume ownership of the resource and free it.*/
	GrWrapOwnership_Adopt GrWrapOwnership = 1
)

type GrWrapCacheable int64

const (
	/*
	   The wrapped resource will be removed from the cache as soon as it becomes purgeable. It may
	   still be assigned and found by a unique key, but the presence of the key will not be used to
	   keep the resource alive when it has no references.*/
	GrWrapCacheable_No GrWrapCacheable = 0
	/*
	   The wrapped resource is allowed to remain in the GrResourceCache when it has no references
	   but has a unique key. Such resources should only be given unique keys when it is known that
	   the key will eventually be removed from the resource or invalidated via the message bus.*/
	GrWrapCacheable_Yes GrWrapCacheable = -1
)

type GrBudgetedType int64

const (
	/*
	   The resource is budgeted and is subject to purging under budget pressure.*/
	GrBudgetedType_Budgeted GrBudgetedType = 0
	/*
	   The resource is unbudgeted and is purged as soon as it has no refs regardless of whether
	   it has a unique or scratch key.*/
	GrBudgetedType_UnbudgetedUncacheable GrBudgetedType = 1
	/*
	   The resource is unbudgeted and is allowed to remain in the cache with no refs if it
	   has a unique key. Scratch keys are ignored.*/
	GrBudgetedType_UnbudgetedCacheable GrBudgetedType = 2
)

type GrScissorTest int64

const (
	GrScissorTest_Disabled GrScissorTest = 0
	GrScissorTest_Enabled  GrScissorTest = -1
)

type GrMemoryless int64

const (
	/*
	   The texture will be allocated normally and will affect memory budgets.*/
	GrMemoryless_No GrMemoryless = 0
	/*
	   The texture will be not use GPU memory and will not affect memory budgets.*/
	GrMemoryless_Yes GrMemoryless = -1
)

type GrSemaphoreWrapType int64

const (
	GrSemaphoreWrapType_WillSignal GrSemaphoreWrapType = 0
	GrSemaphoreWrapType_WillWait   GrSemaphoreWrapType = 1
)

/*
This enum is used to specify the load operation to be used when an OpsTask/GrOpsRenderPass
begins execution.
*/
type GrLoadOp int64

const (
	GrLoadOp_Load    GrLoadOp = 0
	GrLoadOp_Clear   GrLoadOp = 1
	GrLoadOp_Discard GrLoadOp = 2
)

/*
This enum is used to specify the store operation to be used when an OpsTask/GrOpsRenderPass
ends execution.
*/
type GrStoreOp int64

const (
	GrStoreOp_Store   GrStoreOp = 0
	GrStoreOp_Discard GrStoreOp = 1
)

/*
Used to control antialiasing in draw calls.
*/
type GrAA int64

const (
	GrAA_No  GrAA = 0
	GrAA_Yes GrAA = -1
)

type GrFillRule int64

const (
	GrFillRule_Nonzero GrFillRule = 0
	GrFillRule_EvenOdd GrFillRule = -1
)

/*
This enum indicates the type of antialiasing to be performed.
*/
type GrAAType int64

const (
	/*
	   No antialiasing*/
	GrAAType_None GrAAType = 0
	/*
	   Use fragment shader code to blend with a fractional pixel coverage.*/
	GrAAType_Coverage GrAAType = 1
	/*
	   Use normal MSAA.*/
	GrAAType_MSAA GrAAType = 2
	/*
	   Use normal MSAA.*/
	GrAAType_Last GrAAType = 2
)

/*
Some pixel configs are inherently clamped to [0,1], some are allowed to go outside that range,
and some are FP but manually clamped in the XP.
*/
type GrClampType int64

const (
	GrClampType_Auto   GrClampType = 0
	GrClampType_Manual GrClampType = 1
	GrClampType_None   GrClampType = 2
)

/*
A number of rectangle/quadrilateral drawing APIs can control anti-aliasing on a per edge basis.
These masks specify which edges are AA'ed. The intent for this is to support tiling with seamless
boundaries, where the inner edges are non-AA and the outer edges are AA. Regular rectangle draws
simply use kAll or kNone depending on if they want anti-aliasing or not.

In APIs that support per-edge AA, GrQuadAAFlags is the only AA-control parameter that is
provided (compared to the typical GrAA parameter). kNone is equivalent to GrAA::kNo, and any
other set of edge flags would require GrAA::kYes (with rendering output dependent on how that
maps to GrAAType for a given SurfaceDrawContext).

These values are identical to SkCanvas::QuadAAFlags.
*/
type GrQuadAAFlags int64

const (
	GrQuadAAFlags_Left   GrQuadAAFlags = 1
	GrQuadAAFlags_Top    GrQuadAAFlags = 2
	GrQuadAAFlags_Right  GrQuadAAFlags = 4
	GrQuadAAFlags_Bottom GrQuadAAFlags = 8
	GrQuadAAFlags_None   GrQuadAAFlags = 0
	GrQuadAAFlags_All    GrQuadAAFlags = 15
)

/*
The type of texture. Backends other than GL currently only use the 2D value but the type must
still be known at the API-neutral layer as it used to determine whether MIP maps, renderability,
and sampling parameters are legal for proxies that will be instantiated with wrapped textures.
*/
type GrTextureType int64

const (
	GrTextureType_None      GrTextureType = 0
	GrTextureType_2D        GrTextureType = 1
	GrTextureType_Rectangle GrTextureType = 2
	GrTextureType_External  GrTextureType = 3
)

type GrShaderType int64

const (
	GrShaderType_Vertex   GrShaderType = 0
	GrShaderType_Fragment GrShaderType = 1
	GrShaderType_Last     GrShaderType = 1
)

type GrShaderFlags int64

const (
	GrShaderFlags_None              GrShaderFlags = 0
	GrShaderFlags_Vertex_GrShader   GrShaderFlags = 1
	GrShaderFlags_Fragment_GrShader GrShaderFlags = 2
)

/*
Types used to describe format of vertices in arrays.
*/
type GrVertexAttribType int64

const (
	GrVertexAttribType_Float        GrVertexAttribType = 0
	GrVertexAttribType_Float2       GrVertexAttribType = 1
	GrVertexAttribType_Float3       GrVertexAttribType = 2
	GrVertexAttribType_Float4       GrVertexAttribType = 3
	GrVertexAttribType_Half         GrVertexAttribType = 4
	GrVertexAttribType_Half2        GrVertexAttribType = 5
	GrVertexAttribType_Half4        GrVertexAttribType = 6
	GrVertexAttribType_Int2         GrVertexAttribType = 7
	GrVertexAttribType_Int3         GrVertexAttribType = 8
	GrVertexAttribType_Int4         GrVertexAttribType = 9
	GrVertexAttribType_Byte         GrVertexAttribType = 10
	GrVertexAttribType_Byte2        GrVertexAttribType = 11
	GrVertexAttribType_Byte4        GrVertexAttribType = 12
	GrVertexAttribType_UByte        GrVertexAttribType = 13
	GrVertexAttribType_UByte2       GrVertexAttribType = 14
	GrVertexAttribType_UByte4       GrVertexAttribType = 15
	GrVertexAttribType_UByte_norm   GrVertexAttribType = 16
	GrVertexAttribType_UByte4_norm  GrVertexAttribType = 17
	GrVertexAttribType_Short2       GrVertexAttribType = 18
	GrVertexAttribType_Short4       GrVertexAttribType = 19
	GrVertexAttribType_UShort2      GrVertexAttribType = 20
	GrVertexAttribType_UShort2_norm GrVertexAttribType = 21
	GrVertexAttribType_Int          GrVertexAttribType = 22
	GrVertexAttribType_UInt         GrVertexAttribType = 23
	GrVertexAttribType_UShort_norm  GrVertexAttribType = 24
	GrVertexAttribType_UShort4_norm GrVertexAttribType = 25
	GrVertexAttribType_Last         GrVertexAttribType = 25
)

/*
We have coverage effects that clip rendering to the edge of some geometric primitive.
This enum specifies how that clipping is performed. Not all factories that take a
GrClipEdgeType will succeed with all values and it is up to the caller to verify success.
*/
type GrClipEdgeType int64

const (
	GrClipEdgeType_FillBW        GrClipEdgeType = 0
	GrClipEdgeType_FillAA        GrClipEdgeType = 1
	GrClipEdgeType_InverseFillBW GrClipEdgeType = 2
	GrClipEdgeType_InverseFillAA GrClipEdgeType = 3
	GrClipEdgeType_Last          GrClipEdgeType = 3
)

/*
Indicates the type of pending IO operations that can be recorded for gpu resources.
*/
type GrIOType int64

const (
	GrIOType_Read  GrIOType = 0
	GrIOType_Write GrIOType = 1
	GrIOType_RW    GrIOType = 2
)

/*
Indicates the type of data that a GPU buffer will be used for.
*/
type GrGpuBufferType int64

const (
	GrGpuBufferType_Vertex       GrGpuBufferType = 0
	GrGpuBufferType_Index        GrGpuBufferType = 1
	GrGpuBufferType_DrawIndirect GrGpuBufferType = 2
	GrGpuBufferType_XferCpuToGpu GrGpuBufferType = 3
	GrGpuBufferType_XferGpuToCpu GrGpuBufferType = 4
	GrGpuBufferType_Uniform      GrGpuBufferType = 5
)

/*
Provides a performance hint regarding the frequency at which a data store will be accessed.
*/
type GrAccessPattern int64

const (
	/*
	   Data store will be respecified repeatedly and used many times.*/
	GrAccessPattern_Dynamic GrAccessPattern = 0
	/*
	   Data store will be specified once and used many times. (Thus disqualified from caching.)*/
	GrAccessPattern_Static GrAccessPattern = 1
	/*
	   Data store will be specified once and used at most a few times. (Also can't be cached.)*/
	GrAccessPattern_Stream GrAccessPattern = 2
	/*
	   Data store will be specified once and used at most a few times. (Also can't be cached.)*/
	GrAccessPattern_Last GrAccessPattern = 2
)

type GrInternalSurfaceFlags int64

const (
	GrInternalSurfaceFlags_None                        GrInternalSurfaceFlags = 0
	GrInternalSurfaceFlags_ReadOnly                    GrInternalSurfaceFlags = 1
	GrInternalSurfaceFlags_GLRTFBOIDIs0                GrInternalSurfaceFlags = 2
	GrInternalSurfaceFlags_RequiresManualMSAAResolve   GrInternalSurfaceFlags = 4
	GrInternalSurfaceFlags_FramebufferOnly             GrInternalSurfaceFlags = 8
	GrInternalSurfaceFlags_VkRTSupportsInputAttachment GrInternalSurfaceFlags = 16
)

/*
Specifies if the holder owns the backend, OpenGL or Vulkan, object.
*/
type GrBackendObjectOwnership int64

const (
	/*
	   Holder does not destroy the backend object.*/
	GrBackendObjectOwnership_Borrowed GrBackendObjectOwnership = 0
	/*
	   Holder destroys the backend object.*/
	GrBackendObjectOwnership_Owned GrBackendObjectOwnership = -1
)

/*
Used to include or exclude specific GPU path renderers for testing purposes.
*/
type GpuPathRenderers int64

const (
	GpuPathRenderers_None             GpuPathRenderers = 0
	GpuPathRenderers_DashLine         GpuPathRenderers = 1
	GpuPathRenderers_Atlas            GpuPathRenderers = 2
	GpuPathRenderers_Tessellation     GpuPathRenderers = 4
	GpuPathRenderers_CoverageCounting GpuPathRenderers = 8
	GpuPathRenderers_AAHairline       GpuPathRenderers = 16
	GpuPathRenderers_AAConvex         GpuPathRenderers = 32
	GpuPathRenderers_AALinearizing    GpuPathRenderers = 64
	GpuPathRenderers_Small            GpuPathRenderers = 128
	GpuPathRenderers_Triangulating    GpuPathRenderers = 256
	GpuPathRenderers_Default          GpuPathRenderers = 511
)

/*
Used to describe the current state of Mips on a GrTexture
*/
type GrMipmapStatus int64

const (
	GrMipmapStatus_NotAllocated GrMipmapStatus = 0
	GrMipmapStatus_Dirty        GrMipmapStatus = 1
	GrMipmapStatus_Valid        GrMipmapStatus = 2
)

/*
Like SkColorType this describes a layout of pixel data in CPU memory. It specifies the channels,
their type, and width. This exists so that the GPU backend can have private types that have no
analog in the public facing SkColorType enum and omit types not implemented in the GPU backend.
It does not refer to a texture format and the mapping to texture formats may be many-to-many.
It does not specify the sRGB encoding of the stored values. The components are listed in order of
where they appear in memory. In other words the first component listed is in the low bits and
the last component in the high bits.
*/
type GrColorType int64

const (
	GrColorType_Unknown          GrColorType = 0
	GrColorType_Alpha_8          GrColorType = 1
	GrColorType_BGR_565          GrColorType = 2
	GrColorType_RGB_565          GrColorType = 3
	GrColorType_ABGR_4444        GrColorType = 4
	GrColorType_RGBA_8888        GrColorType = 5
	GrColorType_RGBA_8888_SRGB   GrColorType = 6
	GrColorType_RGB_888x         GrColorType = 7
	GrColorType_RG_88            GrColorType = 8
	GrColorType_BGRA_8888        GrColorType = 9
	GrColorType_RGBA_1010102     GrColorType = 10
	GrColorType_BGRA_1010102     GrColorType = 11
	GrColorType_RGBA_10x6        GrColorType = 12
	GrColorType_Gray_8           GrColorType = 13
	GrColorType_GrayAlpha_88     GrColorType = 14
	GrColorType_Alpha_F16        GrColorType = 15
	GrColorType_RGBA_F16         GrColorType = 16
	GrColorType_RGBA_F16_Clamped GrColorType = 17
	GrColorType_RGBA_F32         GrColorType = 18
	GrColorType_Alpha_16         GrColorType = 19
	GrColorType_RG_1616          GrColorType = 20
	GrColorType_RG_F16           GrColorType = 21
	GrColorType_RGBA_16161616    GrColorType = 22
	GrColorType_Alpha_8xxx       GrColorType = 23
	GrColorType_Alpha_F32xxx     GrColorType = 24
	GrColorType_Gray_8xxx        GrColorType = 25
	GrColorType_R_8xxx           GrColorType = 26
	GrColorType_RGB_888          GrColorType = 27
	GrColorType_R_8              GrColorType = 28
	GrColorType_R_16             GrColorType = 29
	GrColorType_R_F16            GrColorType = 30
	GrColorType_Gray_F16         GrColorType = 31
	GrColorType_BGRA_4444        GrColorType = 32
	GrColorType_ARGB_4444        GrColorType = 33
	GrColorType_Last             GrColorType = 33
)

/*
Describes the encoding of channel data in a GrColorType.
*/
type GrColorTypeEncoding int64

const (
	GrColorTypeEncoding_Unorm     GrColorTypeEncoding = 0
	GrColorTypeEncoding_SRGBUnorm GrColorTypeEncoding = 1
	GrColorTypeEncoding_Float     GrColorTypeEncoding = 2
)

type GrDstSampleFlags int64

const (
	GrDstSampleFlags_None                   GrDstSampleFlags = 0
	GrDstSampleFlags_RequiresTextureBarrier GrDstSampleFlags = 1
	GrDstSampleFlags_AsInputAttachment      GrDstSampleFlags = 2
)

type GrDriverBugWorkaroundType int64

const (
	GrDriverBugWorkaroundType_ADD_AND_TRUE_TO_LOOP_CONDITION                GrDriverBugWorkaroundType = 0
	GrDriverBugWorkaroundType_DISABLE_BLEND_EQUATION_ADVANCED               GrDriverBugWorkaroundType = 1
	GrDriverBugWorkaroundType_DISABLE_DISCARD_FRAMEBUFFER                   GrDriverBugWorkaroundType = 2
	GrDriverBugWorkaroundType_DISABLE_DUAL_SOURCE_BLENDING_SUPPORT          GrDriverBugWorkaroundType = 3
	GrDriverBugWorkaroundType_DISABLE_TEXTURE_STORAGE                       GrDriverBugWorkaroundType = 4
	GrDriverBugWorkaroundType_DISALLOW_LARGE_INSTANCED_DRAW                 GrDriverBugWorkaroundType = 5
	GrDriverBugWorkaroundType_EMULATE_ABS_INT_FUNCTION                      GrDriverBugWorkaroundType = 6
	GrDriverBugWorkaroundType_FLUSH_ON_FRAMEBUFFER_CHANGE                   GrDriverBugWorkaroundType = 7
	GrDriverBugWorkaroundType_FORCE_UPDATE_SCISSOR_STATE_WHEN_BINDING_FBO0  GrDriverBugWorkaroundType = 8
	GrDriverBugWorkaroundType_GL_CLEAR_BROKEN                               GrDriverBugWorkaroundType = 9
	GrDriverBugWorkaroundType_MAX_FRAGMENT_UNIFORM_VECTORS_32               GrDriverBugWorkaroundType = 10
	GrDriverBugWorkaroundType_MAX_MSAA_SAMPLE_COUNT_4                       GrDriverBugWorkaroundType = 11
	GrDriverBugWorkaroundType_PACK_PARAMETERS_WORKAROUND_WITH_PACK_BUFFER   GrDriverBugWorkaroundType = 12
	GrDriverBugWorkaroundType_REMOVE_POW_WITH_CONSTANT_EXPONENT             GrDriverBugWorkaroundType = 13
	GrDriverBugWorkaroundType_REWRITE_DO_WHILE_LOOPS                        GrDriverBugWorkaroundType = 14
	GrDriverBugWorkaroundType_UNBIND_ATTACHMENTS_ON_BOUND_RENDER_FBO_DELETE GrDriverBugWorkaroundType = 15
	GrDriverBugWorkaroundType_UNFOLD_SHORT_CIRCUIT_AS_TERNARY_OPERATION     GrDriverBugWorkaroundType = 16
	GrDriverBugWorkaroundType_NUMBER_OF_GPU_DRIVER_BUG_WORKAROUND_TYPES     GrDriverBugWorkaroundType = 17
)

/*
The logical operations that can be performed when combining two paths.
*/
type PathOp int64

const (
	/*
	   subtract the op path from the first path*/
	PathOp_Difference PathOp = 0
	/*
	   intersect the two paths*/
	PathOp_Intersect PathOp = 1
	/*
	   union (inclusive-or) the two paths*/
	PathOp_Union PathOp = 2
	/*
	   exclusive-or the two paths*/
	PathOp_XOR PathOp = 3
	/*
	   subtract the first path from the op path*/
	PathOp_ReverseDifference PathOp = 4
)

type PDFSteamCompressionEnabled int64

const (
	PDFSteamCompressionEnabled_No      PDFSteamCompressionEnabled = 0
	PDFSteamCompressionEnabled_Yes     PDFSteamCompressionEnabled = -1
	PDFSteamCompressionEnabled_Default PDFSteamCompressionEnabled = -1
)

type SVGPropertyState int64

const (
	SVGPropertyState_Unspecified SVGPropertyState = 0
	SVGPropertyState_Inherit     SVGPropertyState = 1
	SVGPropertyState_Value       SVGPropertyState = 2
)

type SVGLineCap int64

const (
	SVGLineCap_Butt   SVGLineCap = 0
	SVGLineCap_Round  SVGLineCap = 1
	SVGLineCap_Square SVGLineCap = 2
)

type SVGFeColorMatrixType int64

const (
	SVGFeColorMatrixType_Matrix           SVGFeColorMatrixType = 0
	SVGFeColorMatrixType_Saturate         SVGFeColorMatrixType = 1
	SVGFeColorMatrixType_HueRotate        SVGFeColorMatrixType = 2
	SVGFeColorMatrixType_LuminanceToAlpha SVGFeColorMatrixType = 3
)

type SVGFeCompositeOperator int64

const (
	SVGFeCompositeOperator_Over       SVGFeCompositeOperator = 0
	SVGFeCompositeOperator_In         SVGFeCompositeOperator = 1
	SVGFeCompositeOperator_Out        SVGFeCompositeOperator = 2
	SVGFeCompositeOperator_Atop       SVGFeCompositeOperator = 3
	SVGFeCompositeOperator_Xor        SVGFeCompositeOperator = 4
	SVGFeCompositeOperator_Arithmetic SVGFeCompositeOperator = 5
)

type SVGXmlSpace int64

const (
	SVGXmlSpace_Default  SVGXmlSpace = 0
	SVGXmlSpace_Preserve SVGXmlSpace = 1
)

type SVGColorspace int64

const (
	SVGColorspace_Auto      SVGColorspace = 0
	SVGColorspace_SRGB      SVGColorspace = 1
	SVGColorspace_LinearRGB SVGColorspace = 2
)

type SVGDisplay int64

const (
	SVGDisplay_Inline SVGDisplay = 0
	SVGDisplay_None   SVGDisplay = 1
)

type SVGFeFuncType int64

const (
	SVGFeFuncType_Identity SVGFeFuncType = 0
	SVGFeFuncType_Table    SVGFeFuncType = 1
	SVGFeFuncType_Discrete SVGFeFuncType = 2
	SVGFeFuncType_Linear   SVGFeFuncType = 3
	SVGFeFuncType_Gamma    SVGFeFuncType = 4
)

type SVGAttribute int64

const (
	SVGAttribute_ClipRule                  SVGAttribute = 0
	SVGAttribute_Color                     SVGAttribute = 1
	SVGAttribute_ColorInterpolation        SVGAttribute = 2
	SVGAttribute_ColorInterpolationFilters SVGAttribute = 3
	SVGAttribute_Cx                        SVGAttribute = 4
	SVGAttribute_Cy                        SVGAttribute = 5
	SVGAttribute_Fill                      SVGAttribute = 6
	SVGAttribute_FillOpacity               SVGAttribute = 7
	SVGAttribute_FillRule                  SVGAttribute = 8
	SVGAttribute_Filter                    SVGAttribute = 9
	SVGAttribute_FilterUnits               SVGAttribute = 10
	SVGAttribute_FontFamily                SVGAttribute = 11
	SVGAttribute_FontSize                  SVGAttribute = 12
	SVGAttribute_FontStyle                 SVGAttribute = 13
	SVGAttribute_FontWeight                SVGAttribute = 14
	SVGAttribute_Fx                        SVGAttribute = 15
	SVGAttribute_Fy                        SVGAttribute = 16
	SVGAttribute_GradientUnits             SVGAttribute = 17
	SVGAttribute_GradientTransform         SVGAttribute = 18
	SVGAttribute_Height                    SVGAttribute = 19
	SVGAttribute_Href                      SVGAttribute = 20
	SVGAttribute_Opacity                   SVGAttribute = 21
	SVGAttribute_Points                    SVGAttribute = 22
	SVGAttribute_PreserveAspectRatio       SVGAttribute = 23
	SVGAttribute_R                         SVGAttribute = 24
	SVGAttribute_Rx                        SVGAttribute = 25
	SVGAttribute_Ry                        SVGAttribute = 26
	SVGAttribute_SpreadMethod              SVGAttribute = 27
	SVGAttribute_Stroke                    SVGAttribute = 28
	SVGAttribute_StrokeDashArray           SVGAttribute = 29
	SVGAttribute_StrokeDashOffset          SVGAttribute = 30
	SVGAttribute_StrokeOpacity             SVGAttribute = 31
	SVGAttribute_StrokeLineCap             SVGAttribute = 32
	SVGAttribute_StrokeLineJoin            SVGAttribute = 33
	SVGAttribute_StrokeMiterLimit          SVGAttribute = 34
	SVGAttribute_StrokeWidth               SVGAttribute = 35
	SVGAttribute_Transform                 SVGAttribute = 36
	SVGAttribute_Text                      SVGAttribute = 37
	SVGAttribute_TextAnchor                SVGAttribute = 38
	SVGAttribute_ViewBox                   SVGAttribute = 39
	SVGAttribute_Visibility                SVGAttribute = 40
	SVGAttribute_Width                     SVGAttribute = 41
	SVGAttribute_X                         SVGAttribute = 42
	SVGAttribute_X1                        SVGAttribute = 43
	SVGAttribute_X2                        SVGAttribute = 44
	SVGAttribute_Y                         SVGAttribute = 45
	SVGAttribute_Y1                        SVGAttribute = 46
	SVGAttribute_Y2                        SVGAttribute = 47
	SVGAttribute_Unknown                   SVGAttribute = 48
)

type SVGTag int64

const (
	SVGTag_Circle              SVGTag = 0
	SVGTag_ClipPath            SVGTag = 1
	SVGTag_Defs                SVGTag = 2
	SVGTag_Ellipse             SVGTag = 3
	SVGTag_FeBlend             SVGTag = 4
	SVGTag_FeColorMatrix       SVGTag = 5
	SVGTag_FeComponentTransfer SVGTag = 6
	SVGTag_FeComposite         SVGTag = 7
	SVGTag_FeDiffuseLighting   SVGTag = 8
	SVGTag_FeDisplacementMap   SVGTag = 9
	SVGTag_FeDistantLight      SVGTag = 10
	SVGTag_FeFlood             SVGTag = 11
	SVGTag_FeFuncA             SVGTag = 12
	SVGTag_FeFuncR             SVGTag = 13
	SVGTag_FeFuncG             SVGTag = 14
	SVGTag_FeFuncB             SVGTag = 15
	SVGTag_FeGaussianBlur      SVGTag = 16
	SVGTag_FeImage             SVGTag = 17
	SVGTag_FeMerge             SVGTag = 18
	SVGTag_FeMergeNode         SVGTag = 19
	SVGTag_FeMorphology        SVGTag = 20
	SVGTag_FeOffset            SVGTag = 21
	SVGTag_FePointLight        SVGTag = 22
	SVGTag_FeSpecularLighting  SVGTag = 23
	SVGTag_FeSpotLight         SVGTag = 24
	SVGTag_FeTurbulence        SVGTag = 25
	SVGTag_Filter              SVGTag = 26
	SVGTag_G                   SVGTag = 27
	SVGTag_Image               SVGTag = 28
	SVGTag_Line                SVGTag = 29
	SVGTag_LinearGradient      SVGTag = 30
	SVGTag_Mask                SVGTag = 31
	SVGTag_Path                SVGTag = 32
	SVGTag_Pattern             SVGTag = 33
	SVGTag_Polygon             SVGTag = 34
	SVGTag_Polyline            SVGTag = 35
	SVGTag_RadialGradient      SVGTag = 36
	SVGTag_Rect                SVGTag = 37
	SVGTag_Stop                SVGTag = 38
	SVGTag_Svg                 SVGTag = 39
	SVGTag_Text                SVGTag = 40
	SVGTag_TextLiteral         SVGTag = 41
	SVGTag_TextPath            SVGTag = 42
	SVGTag_TSpan               SVGTag = 43
	SVGTag_Use                 SVGTag = 44
)
