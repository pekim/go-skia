// This is a generated file. DO NOT EDIT.

package skia

// #cgo CXXFLAGS: -I ${SRCDIR}/header
// #cgo LDFLAGS: -L ${SRCDIR}/lib
// #cgo LDFLAGS: -l skia
// #cgo LDFLAGS: -l skshaper
// #cgo LDFLAGS: -l svg
// #cgo pkg-config: freetype2
// #cgo pkg-config: gl
//
// #include "api.h"
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type GrGLFramebufferInfo C.sk_GrGLFramebufferInfo

func (o GrGLFramebufferInfo) FBOID() uint32 {
	return uint32(o.fFBOID)
}

func (o *GrGLFramebufferInfo) SetFBOID(value uint32) {
	o.fFBOID = C.uint(value)
}

func (o GrGLFramebufferInfo) Format() uint32 {
	return uint32(o.fFormat)
}

func (o *GrGLFramebufferInfo) SetFormat(value uint32) {
	o.fFormat = C.uint(value)
}

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

func (o *GrBackendRenderTarget) Delete() {
	C.misk_delete_GrBackendRenderTarget(o.sk)
	o.sk = nil
}

func (o GrBackendRenderTarget) Dimensions() ISize {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_dimensions(c_obj)
	return ISize(retC)
}

func (o GrBackendRenderTarget) Width() int32 {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_width(c_obj)
	return int32(retC)
}

func (o GrBackendRenderTarget) Height() int32 {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_height(c_obj)
	return int32(retC)
}

func (o GrBackendRenderTarget) SampleCnt() int32 {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_sampleCnt(c_obj)
	return int32(retC)
}

func (o GrBackendRenderTarget) StencilBits() int32 {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_stencilBits(c_obj)
	return int32(retC)
}

func (o GrBackendRenderTarget) IsFramebufferOnly() bool {
	c_obj := o.sk
	retC := C.misk_GrBackendRenderTarget_isFramebufferOnly(c_obj)
	return bool(retC)
}

type GrDirectContext struct {
	sk *C.sk_GrDirectContext
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrDirectContext has not been created.
func (o GrDirectContext) IsNil() bool {
	return o.sk == nil
}

// AsGrRecordingContext converts the GrDirectContext to a GrRecordingContext.
func (o GrDirectContext) AsGrRecordingContext() GrRecordingContext {
	return GrRecordingContext{sk: (*C.sk_GrRecordingContext)(unsafe.Pointer(o.sk))}
}

func (o *GrDirectContext) Delete() {
	C.misk_delete_GrDirectContext(o.sk)
	o.sk = nil
}

/*
Call to ensure all drawing to the context has been flushed and submitted to the underlying 3D
API. This is equivalent to calling GrContext::flush with a default GrFlushInfo followed by
GrContext::submit(sync).
*/
func (o GrDirectContext) FlushAndSubmit(sync GrSyncCpu) {
	c_obj := o.sk
	c_sync := C.bool(sync)
	C.misk_GrDirectContext_flushAndSubmit(c_obj, c_sync)
}

/*
Call to ensure all drawing to the context has been flushed to underlying 3D API specific
objects. A call to `submit` is always required to ensure work is actually sent to
the gpu. Some specific API details:

	GL: Commands are actually sent to the driver, but glFlush is never called. Thus some
	    sync objects from the flush will not be valid until a submission occurs.

	Vulkan/Metal/D3D/Dawn: Commands are recorded to the backend APIs corresponding command
	    buffer or encoder objects. However, these objects are not sent to the gpu until a
	    submission occurs.

If the return is GrSemaphoresSubmitted::kYes, only initialized GrBackendSemaphores will be
submitted to the gpu during the next submit call (it is possible Skia failed to create a
subset of the semaphores). The client should not wait on these semaphores until after submit
has been called, and must keep them alive until then. If this call returns
GrSemaphoresSubmitted::kNo, the GPU backend will not submit any semaphores to be signaled on
the GPU. Thus the client should not have the GPU wait on any of the semaphores passed in with
the GrFlushInfo. Regardless of whether semaphores were submitted to the GPU or not, the
client is still responsible for deleting any initialized semaphores.
Regardless of semaphore submission the context will still be flushed. It should be
emphasized that a return value of GrSemaphoresSubmitted::kNo does not mean the flush did not
happen. It simply means there were no semaphores submitted to the GPU. A caller should only
take this as a failure if they passed in semaphores to be submitted.
*/
func (o GrDirectContext) Flush(info GrFlushInfo) GrSemaphoresSubmitted {
	c_obj := o.sk
	c_info := info.sk
	retC := C.misk_GrDirectContext_flush(c_obj, c_info)
	return GrSemaphoresSubmitted(retC)
}

/*
Submit outstanding work to the gpu from all previously un-submitted flushes. The return
value of the submit will indicate whether or not the submission to the GPU was successful.

If the call returns true, all previously passed in semaphores in flush calls will have been
submitted to the GPU and they can safely be waited on. The caller should wait on those
semaphores or perform some other global synchronization before deleting the semaphores.

If it returns false, then those same semaphores will not have been submitted and we will not
try to submit them again. The caller is free to delete the semaphores at any time.

If sync flag is GrSyncCpu::kYes, this function will return once the gpu has finished with all
submitted work.
*/
func (o GrDirectContext) Submit(sync GrSyncCpu) bool {
	c_obj := o.sk
	c_sync := C.bool(sync)
	retC := C.misk_GrDirectContext_submit(c_obj, c_sync)
	return bool(retC)
}

/*
Struct to supply options to flush calls.

After issuing all commands, fNumSemaphore semaphores will be signaled by the gpu. The client
passes in an array of fNumSemaphores GrBackendSemaphores. In general these GrBackendSemaphore's
can be either initialized or not. If they are initialized, the backend uses the passed in
semaphore. If it is not initialized, a new semaphore is created and the GrBackendSemaphore
object is initialized with that semaphore. The semaphores are not sent to the GPU until the next
GrContext::submit call is made. See the GrContext::submit for more information.

The client will own and be responsible for deleting the underlying semaphores that are stored
and returned in initialized GrBackendSemaphore objects. The GrBackendSemaphore objects
themselves can be deleted as soon as this function returns.

If a finishedProc is provided, the finishedProc will be called when all work submitted to the gpu
from this flush call and all previous flush calls has finished on the GPU. If the flush call
fails due to an error and nothing ends up getting sent to the GPU, the finished proc is called
immediately.

If a submittedProc is provided, the submittedProc will be called when all work from this flush
call is submitted to the GPU. If the flush call fails due to an error and nothing will get sent
to the GPU, the submitted proc is called immediately. It is possibly that when work is finally
submitted, that the submission actual fails. In this case we will not reattempt to do the
submission. Skia notifies the client of these via the success bool passed into the submittedProc.
The submittedProc is useful to the client to know when semaphores that were sent with the flush
have actually been submitted to the GPU so that they can be waited on (or deleted if the submit
fails).
GrBackendSemaphores are not supported for the GL backend and will be ignored if set.
*/
type GrFlushInfo struct {
	sk *C.sk_GrFlushInfo
}

func (o GrFlushInfo) NumSemaphores() uint32 {
	return uint32(o.sk.fNumSemaphores)
}

func (o *GrFlushInfo) SetNumSemaphores(value uint32) {
	o.sk.fNumSemaphores = C.ulong(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrFlushInfo has not been created.
func (o GrFlushInfo) IsNil() bool {
	return o.sk == nil
}

type GrRecordingContext struct {
	sk *C.sk_GrRecordingContext
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the GrRecordingContext has not been created.
func (o GrRecordingContext) IsNil() bool {
	return o.sk == nil
}

func (o *GrRecordingContext) Delete() {
	C.misk_delete_GrRecordingContext(o.sk)
	o.sk = nil
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

func (o *GrGLInterface) Unref() {
	C.misk_unref_GrGLInterface(o.sk)
	o.sk = nil
}
func (o *GrGLInterface) Delete() {
	C.misk_delete_GrGLInterface(o.sk)
	o.sk = nil
}

type GrContextOptions C.sk_GrContextOptions

func (o GrContextOptions) SuppressPrints() bool {
	return bool(o.fSuppressPrints)
}

func (o *GrContextOptions) SetSuppressPrints(value bool) {
	o.fSuppressPrints = C.bool(value)
}

/*
Overrides: These options override feature detection using backend API queries. These
overrides can only reduce the feature set or limits, never increase them beyond the
detected values.
*/
func (o GrContextOptions) MaxTextureSizeOverride() int32 {
	return int32(o.fMaxTextureSizeOverride)
}

func (o *GrContextOptions) SetMaxTextureSizeOverride(value int32) {
	o.fMaxTextureSizeOverride = C.int(value)
}

/*
the threshold in bytes above which we will use a buffer mapping API to map vertex and index
buffers to CPU memory in order to update them.  A value of -1 means the GrContext should
deduce the optimal value for this platform.
*/
func (o GrContextOptions) BufferMapThreshold() int32 {
	return int32(o.fBufferMapThreshold)
}

func (o *GrContextOptions) SetBufferMapThreshold(value int32) {
	o.fBufferMapThreshold = C.int(value)
}

/*
Default minimum size to use when allocating buffers for uploading data to textures. The
larger the value the more uploads can be packed into one buffer, but at the cost of
more gpu memory allocated that may not be used. Uploads larger than the minimum will still
work by allocating a dedicated buffer.
*/
func (o GrContextOptions) MinimumStagingBufferSize() uint32 {
	return uint32(o.fMinimumStagingBufferSize)
}

func (o *GrContextOptions) SetMinimumStagingBufferSize(value uint32) {
	o.fMinimumStagingBufferSize = C.ulong(value)
}

/*
Construct mipmaps manually, via repeated downsampling draw-calls. This is used when
the driver's implementation (glGenerateMipmap) contains bugs. This requires mipmap
level control (ie desktop or ES3).
*/
func (o GrContextOptions) DoManualMipmapping() bool {
	return bool(o.fDoManualMipmapping)
}

func (o *GrContextOptions) SetDoManualMipmapping(value bool) {
	o.fDoManualMipmapping = C.bool(value)
}

/*
Disables the use of coverage counting shortcuts to render paths. Coverage counting can cause
artifacts along shared edges if care isn't taken to ensure both contours wind in the same
direction.
*/
func (o GrContextOptions) DisableCoverageCountingPaths() bool {
	return bool(o.fDisableCoverageCountingPaths)
}

func (o *GrContextOptions) SetDisableCoverageCountingPaths(value bool) {
	o.fDisableCoverageCountingPaths = C.bool(value)
}

/*
Disables distance field rendering for paths. Distance field computation can be expensive,
and yields no benefit if a path is not rendered multiple times with different transforms.
*/
func (o GrContextOptions) DisableDistanceFieldPaths() bool {
	return bool(o.fDisableDistanceFieldPaths)
}

func (o *GrContextOptions) SetDisableDistanceFieldPaths(value bool) {
	o.fDisableDistanceFieldPaths = C.bool(value)
}

/*
If true this allows path mask textures to be cached. This is only really useful if paths
are commonly rendered at the same scale and fractional translation.
*/
func (o GrContextOptions) AllowPathMaskCaching() bool {
	return bool(o.fAllowPathMaskCaching)
}

func (o *GrContextOptions) SetAllowPathMaskCaching(value bool) {
	o.fAllowPathMaskCaching = C.bool(value)
}

/*
If true, the GPU will not be used to perform YUV -> RGB conversion when generating
textures from codec-backed images.
*/
func (o GrContextOptions) DisableGpuYUVConversion() bool {
	return bool(o.fDisableGpuYUVConversion)
}

func (o *GrContextOptions) SetDisableGpuYUVConversion(value bool) {
	o.fDisableGpuYUVConversion = C.bool(value)
}

/*
The maximum size of cache textures used for Skia's Glyph cache.
*/
func (o GrContextOptions) GlyphCacheTextureMaximumBytes() uint32 {
	return uint32(o.fGlyphCacheTextureMaximumBytes)
}

func (o *GrContextOptions) SetGlyphCacheTextureMaximumBytes(value uint32) {
	o.fGlyphCacheTextureMaximumBytes = C.ulong(value)
}

/*
Below this threshold size in device space distance field fonts won't be used. Distance field
fonts don't support hinting which is more important at smaller sizes.
*/
func (o GrContextOptions) MinDistanceFieldFontSize() float32 {
	return float32(o.fMinDistanceFieldFontSize)
}

func (o *GrContextOptions) SetMinDistanceFieldFontSize(value float32) {
	o.fMinDistanceFieldFontSize = C.float(value)
}

func (o GrContextOptions) GlyphsAsPathsFontSize() float32 {
	return float32(o.fGlyphsAsPathsFontSize)
}

func (o *GrContextOptions) SetGlyphsAsPathsFontSize(value float32) {
	o.fGlyphsAsPathsFontSize = C.float(value)
}

/*
Bugs on certain drivers cause stencil buffers to leak. This flag causes Skia to avoid
allocating stencil buffers and use alternate rasterization paths, avoiding the leak.
*/
func (o GrContextOptions) AvoidStencilBuffers() bool {
	return bool(o.fAvoidStencilBuffers)
}

func (o *GrContextOptions) SetAvoidStencilBuffers(value bool) {
	o.fAvoidStencilBuffers = C.bool(value)
}

/*
Some ES3 contexts report the ES2 external image extension, but not the ES3 version.
If support for external images is critical, enabling this option will cause Ganesh to limit
shaders to the ES2 shading language in that situation.
*/
func (o GrContextOptions) PreferExternalImagesOverES3() bool {
	return bool(o.fPreferExternalImagesOverES3)
}

func (o *GrContextOptions) SetPreferExternalImagesOverES3(value bool) {
	o.fPreferExternalImagesOverES3 = C.bool(value)
}

/*
Disables correctness workarounds that are enabled for particular GPUs, OSes, or drivers.
This does not affect code path choices that are made for perfomance reasons nor does it
override other GrContextOption settings.
*/
func (o GrContextOptions) DisableDriverCorrectnessWorkarounds() bool {
	return bool(o.fDisableDriverCorrectnessWorkarounds)
}

func (o *GrContextOptions) SetDisableDriverCorrectnessWorkarounds(value bool) {
	o.fDisableDriverCorrectnessWorkarounds = C.bool(value)
}

/*
Maximum number of GPU programs or pipelines to keep active in the runtime cache.
*/
func (o GrContextOptions) RuntimeProgramCacheSize() int32 {
	return int32(o.fRuntimeProgramCacheSize)
}

func (o *GrContextOptions) SetRuntimeProgramCacheSize(value int32) {
	o.fRuntimeProgramCacheSize = C.int(value)
}

/*
Specifies the number of samples Ganesh should use when performing internal draws with MSAA
(hardware capabilities permitting).

If 0, Ganesh will disable internal code paths that use multisampling.
*/
func (o GrContextOptions) InternalMultisampleCount() int32 {
	return int32(o.fInternalMultisampleCount)
}

func (o *GrContextOptions) SetInternalMultisampleCount(value int32) {
	o.fInternalMultisampleCount = C.int(value)
}

/*
In Skia's vulkan backend a single GrContext submit equates to the submission of a single
primary command buffer to the VkQueue. This value specifies how many vulkan secondary command
buffers we will cache for reuse on a given primary command buffer. A single submit may use
more than this many secondary command buffers, but after the primary command buffer is
finished on the GPU it will only hold on to this many secondary command buffers for reuse.

A value of -1 means we will pick a limit value internally.
*/
func (o GrContextOptions) MaxCachedVulkanSecondaryCommandBuffers() int32 {
	return int32(o.fMaxCachedVulkanSecondaryCommandBuffers)
}

func (o *GrContextOptions) SetMaxCachedVulkanSecondaryCommandBuffers(value int32) {
	o.fMaxCachedVulkanSecondaryCommandBuffers = C.int(value)
}

/*
If true, the caps will never support mipmaps.
*/
func (o GrContextOptions) SuppressMipmapSupport() bool {
	return bool(o.fSuppressMipmapSupport)
}

func (o *GrContextOptions) SetSuppressMipmapSupport(value bool) {
	o.fSuppressMipmapSupport = C.bool(value)
}

/*
If true, the TessellationPathRenderer will not be used for path rendering.
If false, will fallback to any driver workarounds, if set.
*/
func (o GrContextOptions) DisableTessellationPathRenderer() bool {
	return bool(o.fDisableTessellationPathRenderer)
}

func (o *GrContextOptions) SetDisableTessellationPathRenderer(value bool) {
	o.fDisableTessellationPathRenderer = C.bool(value)
}

/*
If true, and if supported, enables hardware tessellation in the caps.
DEPRECATED: This value is ignored; experimental hardware tessellation is always disabled.
*/
func (o GrContextOptions) EnableExperimentalHardwareTessellation() bool {
	return bool(o.fEnableExperimentalHardwareTessellation)
}

func (o *GrContextOptions) SetEnableExperimentalHardwareTessellation(value bool) {
	o.fEnableExperimentalHardwareTessellation = C.bool(value)
}

/*
If true, then add 1 pixel padding to all glyph masks in the atlas to support bi-lerp
rendering of all glyphs. This must be set to true to use Slugs.
*/
func (o GrContextOptions) SupportBilerpFromGlyphAtlas() bool {
	return bool(o.fSupportBilerpFromGlyphAtlas)
}

func (o *GrContextOptions) SetSupportBilerpFromGlyphAtlas(value bool) {
	o.fSupportBilerpFromGlyphAtlas = C.bool(value)
}

/*
Uses a reduced variety of shaders. May perform less optimally in steady state but can reduce
jank due to shader compilations.
*/
func (o GrContextOptions) ReducedShaderVariations() bool {
	return bool(o.fReducedShaderVariations)
}

func (o *GrContextOptions) SetReducedShaderVariations(value bool) {
	o.fReducedShaderVariations = C.bool(value)
}

/*
If true, then allow to enable MSAA on new Intel GPUs.
*/
func (o GrContextOptions) AllowMSAAOnNewIntel() bool {
	return bool(o.fAllowMSAAOnNewIntel)
}

func (o *GrContextOptions) SetAllowMSAAOnNewIntel(value bool) {
	o.fAllowMSAAOnNewIntel = C.bool(value)
}

/*
Currently on ARM Android we disable the use of GL TexStorage because of memory regressions.
However, some clients may still want to use TexStorage. For example, TexStorage support is
required for creating protected textures.

This flag has no impact on non GL backends.
*/
func (o GrContextOptions) AlwaysUseTexStorageWhenAvailable() bool {
	return bool(o.fAlwaysUseTexStorageWhenAvailable)
}

func (o *GrContextOptions) SetAlwaysUseTexStorageWhenAvailable(value bool) {
	o.fAlwaysUseTexStorageWhenAvailable = C.bool(value)
}

func NewGrContextOptions() GrContextOptions {

	retC := C.misk_new_GrContextOptions()
	return *(*GrContextOptions)(unsafe.Pointer(&retC))
}

type GrContextOptionsEnable int32

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

type GrContextOptionsShaderCacheStrategy int32

const (
	GrContextOptionsShaderCacheStrategySkSL          GrContextOptionsShaderCacheStrategy = 0
	GrContextOptionsShaderCacheStrategyBackendSource GrContextOptionsShaderCacheStrategy = 1
	GrContextOptionsShaderCacheStrategyBackendBinary GrContextOptionsShaderCacheStrategy = 2
)

type Arc struct {
	sk *C.sk_SkArc
}

func (o Arc) StartAngle() float32 {
	return float32(o.sk.fStartAngle)
}

func (o *Arc) SetStartAngle(value float32) {
	o.sk.fStartAngle = C.float(value)
}

func (o Arc) SweepAngle() float32 {
	return float32(o.sk.fSweepAngle)
}

func (o *Arc) SetSweepAngle(value float32) {
	o.sk.fSweepAngle = C.float(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Arc has not been created.
func (o Arc) IsNil() bool {
	return o.sk == nil
}

func NewArc() Arc {

	retC := C.misk_new_Arc()
	return Arc{sk: retC}
}

func NewArcCopy(arc Arc) Arc {
	c_arc := arc.sk
	retC := C.misk_new_ArcCopy(c_arc)
	return Arc{sk: retC}
}

func (o *Arc) Delete() {
	C.misk_delete_SkArc(o.sk)
	o.sk = nil
}

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
func (o *Bitmap) Delete() {
	C.misk_delete_SkBitmap(o.sk)
	o.sk = nil
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
Returns number of bytes per pixel required by SkColorType.
Returns zero if colorType( is kUnknown_SkColorType.

@return  bytes in pixel
*/
func (o Bitmap) BytesPerPixel() int32 {
	c_obj := o.sk
	retC := C.misk_Bitmap_bytesPerPixel(c_obj)
	return int32(retC)
}

/*
Returns number of pixels that fit on row. Should be greater than or equal to
width().

@return  maximum pixels per row
*/
func (o Bitmap) RowBytesAsPixels() int32 {
	c_obj := o.sk
	retC := C.misk_Bitmap_rowBytesAsPixels(c_obj)
	return int32(retC)
}

/*
Returns bit shift converting row bytes to row pixels.
Returns zero for kUnknown_SkColorType.

@return  one of: 0, 1, 2, 3; left shift to convert pixels to bytes
*/
func (o Bitmap) ShiftPerPixel() int32 {
	c_obj := o.sk
	retC := C.misk_Bitmap_shiftPerPixel(c_obj)
	return int32(retC)
}

/*
Returns true if width() or height() are zero, or if SkPixelRef is nullptr.
If true, SkBitmap has no effect when drawn or drawn into.

@return  true if drawing has no effect
*/
func (o Bitmap) DrawsNothing() bool {
	c_obj := o.sk
	retC := C.misk_Bitmap_drawsNothing(c_obj)
	return bool(retC)
}

/*
Returns row bytes, the interval from one pixel row to the next. Row bytes
is at least as large as: width() * info().bytesPerPixel().

Returns zero if colorType() is kUnknown_SkColorType, or if row bytes supplied to
setInfo() is not large enough to hold a row of pixels.

@return  byte length of pixel row
*/
func (o Bitmap) RowBytes() uint32 {
	c_obj := o.sk
	retC := C.misk_Bitmap_rowBytes(c_obj)
	return uint32(retC)
}

/*
Returns true if pixels can not change.

Most immutable SkBitmap checks trigger an assert only on debug builds.

@return  true if pixels are immutable

example: https://fiddle.skia.org/c/@Bitmap_isImmutable
*/
func (o Bitmap) IsImmutable() bool {
	c_obj := o.sk
	retC := C.misk_Bitmap_isImmutable(c_obj)
	return bool(retC)
}

/*
Sets internal flag to mark SkBitmap as immutable. Once set, pixels can not change.
Any other bitmap sharing the same SkPixelRef are also marked as immutable.
Once SkPixelRef is marked immutable, the setting cannot be cleared.

Writing to immutable SkBitmap pixels triggers an assert on debug builds.

example: https://fiddle.skia.org/c/@Bitmap_setImmutable
*/
func (o Bitmap) SetImmutable() {
	c_obj := o.sk
	C.misk_Bitmap_setImmutable(c_obj)
}

/*
Returns SkIRect { 0, 0, width(), height() }.

@return  integral rectangle from origin to width() and height()
*/
func (o Bitmap) Bounds() IRect {
	c_obj := o.sk
	retC := C.misk_Bitmap_bounds(c_obj)
	return IRect(retC)
}

/*
Returns SkISize { width(), height() }.

@return  integral size of width() and height()
*/
func (o Bitmap) Dimensions() ISize {
	c_obj := o.sk
	retC := C.misk_Bitmap_dimensions(c_obj)
	return ISize(retC)
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
func NewCanvasWithDimensions(width int32, height int32, props SurfaceProps) Canvas {
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
func (o *Canvas) Delete() {
	C.misk_delete_SkCanvas(o.sk)
	o.sk = nil
}

/*
Copies SkSurfaceProps, if SkCanvas is associated with raster surface or
GPU surface, and returns true. Otherwise, returns false and leave props unchanged.

@param props  storage for writable SkSurfaceProps
@return       true if SkSurfaceProps was copied

DEPRECATED: Replace usage with getBaseProps() or getTopProps()

example: https://fiddle.skia.org/c/@Canvas_getProps
*/
func (o Canvas) GetProps(props SurfaceProps) bool {
	c_obj := o.sk
	c_props := props.sk
	retC := C.misk_Canvas_getProps(c_obj, c_props)
	return bool(retC)
}

/*
Returns the SkSurfaceProps associated with the canvas (i.e., at the base of the layer
stack).

@return  base SkSurfaceProps
*/
func (o Canvas) GetBaseProps() SurfaceProps {
	c_obj := o.sk
	retC := C.misk_Canvas_getBaseProps(c_obj)
	return SurfaceProps{sk: &retC}
}

/*
Returns the SkSurfaceProps associated with the canvas that are currently active (i.e., at
the top of the layer stack). This can differ from getBaseProps depending on the flags
passed to saveLayer (see SaveLayerFlagsSet).

@return  SkSurfaceProps active in the current/top layer
*/
func (o Canvas) GetTopProps() SurfaceProps {
	c_obj := o.sk
	retC := C.misk_Canvas_getTopProps(c_obj)
	return SurfaceProps{sk: &retC}
}

/*
Gets the size of the base or root layer in global canvas coordinates. The
origin of the base layer is always (0,0). The area available for drawing may be
smaller (due to clipping or saveLayer).

@return  integral width and height of base layer

example: https://fiddle.skia.org/c/@Canvas_getBaseLayerSize
*/
func (o Canvas) GetBaseLayerSize() ISize {
	c_obj := o.sk
	retC := C.misk_Canvas_getBaseLayerSize(c_obj)
	return ISize(retC)
}

/*
Creates SkSurface matching info and props, and associates it with SkCanvas.
Returns nullptr if no match found.

If props is nullptr, matches SkSurfaceProps in SkCanvas. If props is nullptr and SkCanvas
does not have SkSurfaceProps, creates SkSurface with default SkSurfaceProps.

@param info   width, height, SkColorType, SkAlphaType, and SkColorSpace
@param props  SkSurfaceProps to match; may be nullptr to match SkCanvas
@return       SkSurface matching info and props, or nullptr if no match is available

example: https://fiddle.skia.org/c/@Canvas_makeSurface
*/
func (o Canvas) MakeSurface(info ImageInfo, props SurfaceProps) Surface {
	c_obj := o.sk
	c_info := info.sk
	c_props := props.sk
	retC := C.misk_Canvas_makeSurface(c_obj, c_info, c_props)
	return Surface{sk: retC}
}

/*
Sometimes a canvas is owned by a surface. If it is, getSurface() will return a bare

	pointer to that surface, else this will return nullptr.
*/
func (o Canvas) GetSurface() Surface {
	c_obj := o.sk
	retC := C.misk_Canvas_getSurface(c_obj)
	return Surface{sk: retC}
}

/*
Returns Ganesh context of the GPU surface associated with SkCanvas.

@return  GPU context, if available; nullptr otherwise

example: https://fiddle.skia.org/c/@Canvas_recordingContext
*/
func (o Canvas) RecordingContext() GrRecordingContext {
	c_obj := o.sk
	retC := C.misk_Canvas_recordingContext(c_obj)
	return GrRecordingContext{sk: retC}
}

/*
Returns true if SkCanvas has direct access to its pixels.

Pixels are readable when SkDevice is raster. Pixels are not readable when SkCanvas
is returned from GPU surface, returned by SkDocument::beginPage, returned by
SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility class
like DebugCanvas.

pixmap is valid only while SkCanvas is in scope and unchanged. Any
SkCanvas or SkSurface call may invalidate the pixmap values.

@param pixmap  storage for pixel state if pixels are readable; otherwise, ignored
@return        true if SkCanvas has direct access to pixels

example: https://fiddle.skia.org/c/@Canvas_peekPixels
*/
func (o Canvas) PeekPixels(pixmap Pixmap) bool {
	c_obj := o.sk
	c_pixmap := pixmap.sk
	retC := C.misk_Canvas_peekPixels(c_obj, c_pixmap)
	return bool(retC)
}

/*
Copies SkRect of pixels from SkCanvas into dstPixels. SkMatrix and clip are
ignored.

Source SkRect corners are (srcX, srcY) and (imageInfo().width(), imageInfo().height()).
Destination SkRect corners are (0, 0) and (dstInfo.width(), dstInfo.height()).
Copies each readable pixel intersecting both rectangles, without scaling,
converting to dstInfo.colorType() and dstInfo.alphaType() if required.

Pixels are readable when SkDevice is raster, or backed by a GPU.
Pixels are not readable when SkCanvas is returned by SkDocument::beginPage,
returned by SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility
class like DebugCanvas.

The destination pixel storage must be allocated by the caller.

Pixel values are converted only if SkColorType and SkAlphaType
do not match. Only pixels within both source and destination rectangles
are copied. dstPixels contents outside SkRect intersection are unchanged.

Pass negative values for srcX or srcY to offset pixels across or down destination.

Does not copy, and returns false if:
- Source and destination rectangles do not intersect.
- SkCanvas pixels could not be converted to dstInfo.colorType() or dstInfo.alphaType().
- SkCanvas pixels are not readable; for instance, SkCanvas is document-based.
- dstRowBytes is too small to contain one row of pixels.

@param dstInfo      width, height, SkColorType, and SkAlphaType of dstPixels
@param dstPixels    storage for pixels; dstInfo.height() times dstRowBytes, or larger
@param dstRowBytes  size of one destination row; dstInfo.width() times pixel size, or larger
@param srcX         offset into readable pixels on x-axis; may be negative
@param srcY         offset into readable pixels on y-axis; may be negative
@return             true if pixels were copied
*/
func (o Canvas) ReadPixelsImageInfo(dstInfo ImageInfo, dstPixels []byte, dstRowBytes uint32, srcX int32, srcY int32) bool {
	c_obj := o.sk
	c_dstInfo := dstInfo.sk
	c_dstPixels := unsafe.Pointer(&dstPixels[0])
	c_dstRowBytes := C.ulong(dstRowBytes)
	c_srcX := C.int(srcX)
	c_srcY := C.int(srcY)
	retC := C.misk_Canvas_readPixelsImageInfo(c_obj, c_dstInfo, c_dstPixels, c_dstRowBytes, c_srcX, c_srcY)
	return bool(retC)
}

/*
Copies SkRect of pixels from SkCanvas into pixmap. SkMatrix and clip are
ignored.

Source SkRect corners are (srcX, srcY) and (imageInfo().width(), imageInfo().height()).
Destination SkRect corners are (0, 0) and (pixmap.width(), pixmap.height()).
Copies each readable pixel intersecting both rectangles, without scaling,
converting to pixmap.colorType() and pixmap.alphaType() if required.

Pixels are readable when SkDevice is raster, or backed by a GPU.
Pixels are not readable when SkCanvas is returned by SkDocument::beginPage,
returned by SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility
class like DebugCanvas.

Caller must allocate pixel storage in pixmap if needed.

Pixel values are converted only if SkColorType and SkAlphaType
do not match. Only pixels within both source and destination SkRect
are copied. pixmap pixels contents outside SkRect intersection are unchanged.

Pass negative values for srcX or srcY to offset pixels across or down pixmap.

Does not copy, and returns false if:
- Source and destination rectangles do not intersect.
- SkCanvas pixels could not be converted to pixmap.colorType() or pixmap.alphaType().
- SkCanvas pixels are not readable; for instance, SkCanvas is document-based.
- SkPixmap pixels could not be allocated.
- pixmap.rowBytes() is too small to contain one row of pixels.

@param pixmap  storage for pixels copied from SkCanvas
@param srcX    offset into readable pixels on x-axis; may be negative
@param srcY    offset into readable pixels on y-axis; may be negative
@return        true if pixels were copied

example: https://fiddle.skia.org/c/@Canvas_readPixels_2
*/
func (o Canvas) ReadPixelsPixmap(pixmap Pixmap, srcX int32, srcY int32) bool {
	c_obj := o.sk
	c_pixmap := pixmap.sk
	c_srcX := C.int(srcX)
	c_srcY := C.int(srcY)
	retC := C.misk_Canvas_readPixelsPixmap(c_obj, c_pixmap, c_srcX, c_srcY)
	return bool(retC)
}

/*
Copies SkRect of pixels from SkCanvas into bitmap. SkMatrix and clip are
ignored.

Source SkRect corners are (srcX, srcY) and (imageInfo().width(), imageInfo().height()).
Destination SkRect corners are (0, 0) and (bitmap.width(), bitmap.height()).
Copies each readable pixel intersecting both rectangles, without scaling,
converting to bitmap.colorType() and bitmap.alphaType() if required.

Pixels are readable when SkDevice is raster, or backed by a GPU.
Pixels are not readable when SkCanvas is returned by SkDocument::beginPage,
returned by SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility
class like DebugCanvas.

Caller must allocate pixel storage in bitmap if needed.

SkBitmap values are converted only if SkColorType and SkAlphaType
do not match. Only pixels within both source and destination rectangles
are copied. SkBitmap pixels outside SkRect intersection are unchanged.

Pass negative values for srcX or srcY to offset pixels across or down bitmap.

Does not copy, and returns false if:
- Source and destination rectangles do not intersect.
- SkCanvas pixels could not be converted to bitmap.colorType() or bitmap.alphaType().
- SkCanvas pixels are not readable; for instance, SkCanvas is document-based.
- bitmap pixels could not be allocated.
- bitmap.rowBytes() is too small to contain one row of pixels.

@param bitmap  storage for pixels copied from SkCanvas
@param srcX    offset into readable pixels on x-axis; may be negative
@param srcY    offset into readable pixels on y-axis; may be negative
@return        true if pixels were copied

example: https://fiddle.skia.org/c/@Canvas_readPixels_3
*/
func (o Canvas) ReadPixelsBitmap(bitmap Bitmap, srcX int32, srcY int32) bool {
	c_obj := o.sk
	c_bitmap := bitmap.sk
	c_srcX := C.int(srcX)
	c_srcY := C.int(srcY)
	retC := C.misk_Canvas_readPixelsBitmap(c_obj, c_bitmap, c_srcX, c_srcY)
	return bool(retC)
}

/*
Copies SkRect from pixels to SkCanvas. SkMatrix and clip are ignored.
Source SkRect corners are (0, 0) and (info.width(), info.height()).
Destination SkRect corners are (x, y) and
(imageInfo().width(), imageInfo().height()).

Copies each readable pixel intersecting both rectangles, without scaling,
converting to imageInfo().colorType() and imageInfo().alphaType() if required.

Pixels are writable when SkDevice is raster, or backed by a GPU.
Pixels are not writable when SkCanvas is returned by SkDocument::beginPage,
returned by SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility
class like DebugCanvas.

Pixel values are converted only if SkColorType and SkAlphaType
do not match. Only pixels within both source and destination rectangles
are copied. SkCanvas pixels outside SkRect intersection are unchanged.

Pass negative values for x or y to offset pixels to the left or
above SkCanvas pixels.

Does not copy, and returns false if:
- Source and destination rectangles do not intersect.
- pixels could not be converted to SkCanvas imageInfo().colorType() or
imageInfo().alphaType().
- SkCanvas pixels are not writable; for instance, SkCanvas is document-based.
- rowBytes is too small to contain one row of pixels.

@param info      width, height, SkColorType, and SkAlphaType of pixels
@param pixels    pixels to copy, of size info.height() times rowBytes, or larger
@param rowBytes  size of one row of pixels; info.width() times pixel size, or larger
@param x         offset into SkCanvas writable pixels on x-axis; may be negative
@param y         offset into SkCanvas writable pixels on y-axis; may be negative
@return          true if pixels were written to SkCanvas

example: https://fiddle.skia.org/c/@Canvas_writePixels
*/
func (o Canvas) WritePixelsImageInfo(info ImageInfo, pixels []byte, rowBytes uint32, x int32, y int32) bool {
	c_obj := o.sk
	c_info := info.sk
	c_pixels := unsafe.Pointer(&pixels[0])
	c_rowBytes := C.ulong(rowBytes)
	c_x := C.int(x)
	c_y := C.int(y)
	retC := C.misk_Canvas_writePixelsImageInfo(c_obj, c_info, c_pixels, c_rowBytes, c_x, c_y)
	return bool(retC)
}

/*
Copies SkRect from pixels to SkCanvas. SkMatrix and clip are ignored.
Source SkRect corners are (0, 0) and (bitmap.width(), bitmap.height()).

Destination SkRect corners are (x, y) and
(imageInfo().width(), imageInfo().height()).

Copies each readable pixel intersecting both rectangles, without scaling,
converting to imageInfo().colorType() and imageInfo().alphaType() if required.

Pixels are writable when SkDevice is raster, or backed by a GPU.
Pixels are not writable when SkCanvas is returned by SkDocument::beginPage,
returned by SkPictureRecorder::beginRecording, or SkCanvas is the base of a utility
class like DebugCanvas.

Pixel values are converted only if SkColorType and SkAlphaType
do not match. Only pixels within both source and destination rectangles
are copied. SkCanvas pixels outside SkRect intersection are unchanged.

Pass negative values for x or y to offset pixels to the left or
above SkCanvas pixels.

Does not copy, and returns false if:
- Source and destination rectangles do not intersect.
- bitmap does not have allocated pixels.
- bitmap pixels could not be converted to SkCanvas imageInfo().colorType() or
imageInfo().alphaType().
- SkCanvas pixels are not writable; for instance, SkCanvas is document based.
- bitmap pixels are inaccessible; for instance, bitmap wraps a texture.

@param bitmap  contains pixels copied to SkCanvas
@param x       offset into SkCanvas writable pixels on x-axis; may be negative
@param y       offset into SkCanvas writable pixels on y-axis; may be negative
@return        true if pixels were written to SkCanvas

example: https://fiddle.skia.org/c/@Canvas_writePixels_2
example: https://fiddle.skia.org/c/@State_Stack_a
example: https://fiddle.skia.org/c/@State_Stack_b
*/
func (o Canvas) WritePixelsBitmap(bitmap Bitmap, x int32, y int32) bool {
	c_obj := o.sk
	c_bitmap := bitmap.sk
	c_x := C.int(x)
	c_y := C.int(y)
	retC := C.misk_Canvas_writePixelsBitmap(c_obj, c_bitmap, c_x, c_y)
	return bool(retC)
}

/*
Saves SkMatrix and clip.
Calling restore() discards changes to SkMatrix and clip,
restoring the SkMatrix and clip to their state when save() was called.

SkMatrix may be changed by translate(), scale(), rotate(), skew(), concat(), setMatrix(),
and resetMatrix(). Clip may be changed by clipRect(), clipRRect(), clipPath(), clipRegion().

Saved SkCanvas state is put on a stack; multiple calls to save() should be balance
by an equal number of calls to restore().

Call restoreToCount() with result to restore this and subsequent saves.

@return  depth of saved stack

example: https://fiddle.skia.org/c/@Canvas_save
*/
func (o Canvas) Save() int32 {
	c_obj := o.sk
	retC := C.misk_Canvas_save(c_obj)
	return int32(retC)
}

/*
Saves SkMatrix and clip, and allocates a SkSurface for subsequent drawing.
Calling restore() discards changes to SkMatrix and clip, and draws the SkSurface.

SkMatrix may be changed by translate(), scale(), rotate(), skew(), concat(),
setMatrix(), and resetMatrix(). Clip may be changed by clipRect(), clipRRect(),
clipPath(), clipRegion().

SkRect bounds suggests but does not define the SkSurface size. To clip drawing to
a specific rectangle, use clipRect().

Optional SkPaint paint applies alpha, SkColorFilter, SkImageFilter, and
SkBlendMode when restore() is called.

Call restoreToCount() with returned value to restore this and subsequent saves.

@param bounds  hint to limit the size of the layer; may be nullptr
@param paint   graphics state for layer; may be nullptr
@return        depth of saved stack

example: https://fiddle.skia.org/c/@Canvas_saveLayer
example: https://fiddle.skia.org/c/@Canvas_saveLayer_4
*/
func (o Canvas) SaveLayer(bounds Rect, paint Paint) int32 {
	c_obj := o.sk
	c_bounds := *(*C.sk_SkRect)(unsafe.Pointer(&bounds))
	c_paint := paint.sk
	retC := C.misk_Canvas_saveLayer(c_obj, c_bounds, c_paint)
	return int32(retC)
}

func (o Canvas) SaveLayerAlpha(bounds Rect, alpha uint32) int32 {
	c_obj := o.sk
	c_bounds := *(*C.sk_SkRect)(unsafe.Pointer(&bounds))
	c_alpha := C.uint(alpha)
	retC := C.misk_Canvas_saveLayerAlpha(c_obj, c_bounds, c_alpha)
	return int32(retC)
}

/*
Removes changes to SkMatrix and clip since SkCanvas state was
last saved. The state is removed from the stack.

Does nothing if the stack is empty.

example: https://fiddle.skia.org/c/@AutoCanvasRestore_restore

example: https://fiddle.skia.org/c/@Canvas_restore
*/
func (o Canvas) Restore() {
	c_obj := o.sk
	C.misk_Canvas_restore(c_obj)
}

/*
Returns the number of saved states, each containing: SkMatrix and clip.
Equals the number of save() calls less the number of restore() calls plus one.
The save count of a new canvas is one.

@return  depth of save state stack

example: https://fiddle.skia.org/c/@Canvas_getSaveCount
*/
func (o Canvas) GetSaveCount() int32 {
	c_obj := o.sk
	retC := C.misk_Canvas_getSaveCount(c_obj)
	return int32(retC)
}

/*
Restores state to SkMatrix and clip values when save(), saveLayer(),
saveLayerPreserveLCDTextRequests(), or saveLayerAlpha() returned saveCount.

Does nothing if saveCount is greater than state stack count.
Restores state to initial values if saveCount is less than or equal to one.

@param saveCount  depth of state stack to restore

example: https://fiddle.skia.org/c/@Canvas_restoreToCount
*/
func (o Canvas) RestoreToCount(saveCount int32) {
	c_obj := o.sk
	c_saveCount := C.int(saveCount)
	C.misk_Canvas_restoreToCount(c_obj, c_saveCount)
}

/*
Translates SkMatrix by dx along the x-axis and dy along the y-axis.

Mathematically, replaces SkMatrix with a translation matrix
premultiplied with SkMatrix.

This has the effect of moving the drawing by (dx, dy) before transforming
the result with SkMatrix.

@param dx  distance to translate on x-axis
@param dy  distance to translate on y-axis

example: https://fiddle.skia.org/c/@Canvas_translate
*/
func (o Canvas) Translate(dx float32, dy float32) {
	c_obj := o.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	C.misk_Canvas_translate(c_obj, c_dx, c_dy)
}

/*
Scales SkMatrix by sx on the x-axis and sy on the y-axis.

Mathematically, replaces SkMatrix with a scale matrix
premultiplied with SkMatrix.

This has the effect of scaling the drawing by (sx, sy) before transforming
the result with SkMatrix.

@param sx  amount to scale on x-axis
@param sy  amount to scale on y-axis

example: https://fiddle.skia.org/c/@Canvas_scale
*/
func (o Canvas) Scale(sx float32, sy float32) {
	c_obj := o.sk
	c_sx := C.float(sx)
	c_sy := C.float(sy)
	C.misk_Canvas_scale(c_obj, c_sx, c_sy)
}

/*
Rotates SkMatrix by degrees. Positive degrees rotates clockwise.

Mathematically, replaces SkMatrix with a rotation matrix
premultiplied with SkMatrix.

This has the effect of rotating the drawing by degrees before transforming
the result with SkMatrix.

@param degrees  amount to rotate, in degrees

example: https://fiddle.skia.org/c/@Canvas_rotate
*/
func (o Canvas) Rotate(degrees float32) {
	c_obj := o.sk
	c_degrees := C.float(degrees)
	C.misk_Canvas_rotate(c_obj, c_degrees)
}

/*
Rotates SkMatrix by degrees about a point at (px, py). Positive degrees rotates
clockwise.

Mathematically, constructs a rotation matrix; premultiplies the rotation matrix by
a translation matrix; then replaces SkMatrix with the resulting matrix
premultiplied with SkMatrix.

This has the effect of rotating the drawing about a given point before
transforming the result with SkMatrix.

@param degrees  amount to rotate, in degrees
@param px       x-axis value of the point to rotate about
@param py       y-axis value of the point to rotate about

example: https://fiddle.skia.org/c/@Canvas_rotate_2
*/
func (o Canvas) RotateAboutPoint(degrees float32, px float32, py float32) {
	c_obj := o.sk
	c_degrees := C.float(degrees)
	c_px := C.float(px)
	c_py := C.float(py)
	C.misk_Canvas_rotateAboutPoint(c_obj, c_degrees, c_px, c_py)
}

/*
Skews SkMatrix by sx on the x-axis and sy on the y-axis. A positive value of sx
skews the drawing right as y-axis values increase; a positive value of sy skews
the drawing down as x-axis values increase.

Mathematically, replaces SkMatrix with a skew matrix premultiplied with SkMatrix.

This has the effect of skewing the drawing by (sx, sy) before transforming
the result with SkMatrix.

@param sx  amount to skew on x-axis
@param sy  amount to skew on y-axis

example: https://fiddle.skia.org/c/@Canvas_skew
*/
func (o Canvas) Skew(sx float32, sy float32) {
	c_obj := o.sk
	c_sx := C.float(sx)
	c_sy := C.float(sy)
	C.misk_Canvas_skew(c_obj, c_sx, c_sy)
}

/*
Replaces SkMatrix with matrix premultiplied with existing SkMatrix.

This has the effect of transforming the drawn geometry by matrix, before
transforming the result with existing SkMatrix.

@param matrix  matrix to premultiply with existing SkMatrix

example: https://fiddle.skia.org/c/@Canvas_concat
*/
func (o Canvas) ConcatMatrix(matrix Matrix) {
	c_obj := o.sk
	c_matrix := matrix.sk
	C.misk_Canvas_concatMatrix(c_obj, c_matrix)
}

func (o Canvas) ConcatM44(p0 M44) {
	c_obj := o.sk
	c_p0 := p0.sk
	C.misk_Canvas_concatM44(c_obj, c_p0)
}

/*
Replaces SkMatrix with matrix.
Unlike concat(), any prior matrix state is overwritten.

@param matrix  matrix to copy, replacing existing SkMatrix

example: https://fiddle.skia.org/c/@Canvas_setMatrix
*/
func (o Canvas) SetMatrixM44(matrix M44) {
	c_obj := o.sk
	c_matrix := matrix.sk
	C.misk_Canvas_setMatrixM44(c_obj, c_matrix)
}

func (o Canvas) SetMatrix(matrix Matrix) {
	c_obj := o.sk
	c_matrix := matrix.sk
	C.misk_Canvas_setMatrix(c_obj, c_matrix)
}

/*
Sets SkMatrix to the identity matrix.
Any prior matrix state is overwritten.

example: https://fiddle.skia.org/c/@Canvas_resetMatrix
*/
func (o Canvas) ResetMatrix() {
	c_obj := o.sk
	C.misk_Canvas_resetMatrix(c_obj)
}

/*
Replaces clip with the intersection or difference of clip and rect,
with an aliased or anti-aliased clip edge. rect is transformed by SkMatrix
before it is combined with clip.

@param rect         SkRect to combine with clip
@param op           SkClipOp to apply to clip
@param doAntiAlias  true if clip is to be anti-aliased

example: https://fiddle.skia.org/c/@Canvas_clipRect
*/
func (o Canvas) ClipRect(rect Rect, op ClipOp, doAntiAlias bool) {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_op := C.int(op)
	c_doAntiAlias := C.bool(doAntiAlias)
	C.misk_Canvas_clipRect(c_obj, c_rect, c_op, c_doAntiAlias)
}

/*
Replaces clip with the intersection or difference of clip and rrect,
with an aliased or anti-aliased clip edge.
rrect is transformed by SkMatrix
before it is combined with clip.

@param rrect        SkRRect to combine with clip
@param op           SkClipOp to apply to clip
@param doAntiAlias  true if clip is to be anti-aliased

example: https://fiddle.skia.org/c/@Canvas_clipRRect
*/
func (o Canvas) ClipRRect(rrect RRect, op ClipOp, doAntiAlias bool) {
	c_obj := o.sk
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	c_op := C.int(op)
	c_doAntiAlias := C.bool(doAntiAlias)
	C.misk_Canvas_clipRRect(c_obj, c_rrect, c_op, c_doAntiAlias)
}

/*
Replaces clip with the intersection or difference of clip and path,
with an aliased or anti-aliased clip edge. SkPath::FillType determines if path
describes the area inside or outside its contours; and if path contour overlaps
itself or another path contour, whether the overlaps form part of the area.
path is transformed by SkMatrix before it is combined with clip.

@param path         SkPath to combine with clip
@param op           SkClipOp to apply to clip
@param doAntiAlias  true if clip is to be anti-aliased

example: https://fiddle.skia.org/c/@Canvas_clipPath
*/
func (o Canvas) ClipPath(path Path, op ClipOp, doAntiAlias bool) {
	c_obj := o.sk
	c_path := path.sk
	c_op := C.int(op)
	c_doAntiAlias := C.bool(doAntiAlias)
	C.misk_Canvas_clipPath(c_obj, c_path, c_op, c_doAntiAlias)
}

/*
Replaces clip with the intersection or difference of clip and SkRegion deviceRgn.
Resulting clip is aliased; pixels are fully contained by the clip.
deviceRgn is unaffected by SkMatrix.

@param deviceRgn  SkRegion to combine with clip
@param op         SkClipOp to apply to clip

example: https://fiddle.skia.org/c/@Canvas_clipRegion
*/
func (o Canvas) ClipRegion(deviceRgn Region, op ClipOp) {
	c_obj := o.sk
	c_deviceRgn := deviceRgn.sk
	c_op := C.int(op)
	C.misk_Canvas_clipRegion(c_obj, c_deviceRgn, c_op)
}

/*
Returns true if SkRect rect, transformed by SkMatrix, can be quickly determined to be
outside of clip. May return false even though rect is outside of clip.

Use to check if an area to be drawn is clipped out, to skip subsequent draw calls.

@param rect  SkRect to compare with clip
@return      true if rect, transformed by SkMatrix, does not intersect clip

example: https://fiddle.skia.org/c/@Canvas_quickReject
*/
func (o Canvas) QuickRejectRect(rect Rect) bool {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	retC := C.misk_Canvas_quickRejectRect(c_obj, c_rect)
	return bool(retC)
}

/*
Returns true if path, transformed by SkMatrix, can be quickly determined to be
outside of clip. May return false even though path is outside of clip.

Use to check if an area to be drawn is clipped out, to skip subsequent draw calls.

@param path  SkPath to compare with clip
@return      true if path, transformed by SkMatrix, does not intersect clip

example: https://fiddle.skia.org/c/@Canvas_quickReject_2
*/
func (o Canvas) QuickRejectPath(path Path) bool {
	c_obj := o.sk
	c_path := path.sk
	retC := C.misk_Canvas_quickRejectPath(c_obj, c_path)
	return bool(retC)
}

/*
Returns bounds of clip, transformed by inverse of SkMatrix. If clip is empty,
return SkRect::MakeEmpty, where all SkRect sides equal zero.

SkRect returned is outset by one to account for partial pixel coverage if clip
is anti-aliased.

@return  bounds of clip in local coordinates

example: https://fiddle.skia.org/c/@Canvas_getLocalClipBounds
*/
func (o Canvas) GetLocalClipBoundsRect() Rect {
	c_obj := o.sk
	retC := C.misk_Canvas_getLocalClipBoundsRect(c_obj)
	return Rect(retC)
}

/*
Returns bounds of clip, transformed by inverse of SkMatrix. If clip is empty,
return false, and set bounds to SkRect::MakeEmpty, where all SkRect sides equal zero.

bounds is outset by one to account for partial pixel coverage if clip
is anti-aliased.

@param bounds  SkRect of clip in local coordinates
@return        true if clip bounds is not empty
*/
func (o Canvas) GetLocalClipBoundsPath(bounds *Rect) bool {
	c_obj := o.sk
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_Canvas_getLocalClipBoundsPath(c_obj, c_bounds)
	return bool(retC)
}

/*
Returns SkIRect bounds of clip, unaffected by SkMatrix. If clip is empty,
return SkRect::MakeEmpty, where all SkRect sides equal zero.

Unlike getLocalClipBounds(), returned SkIRect is not outset.

@return  bounds of clip in base device coordinates

example: https://fiddle.skia.org/c/@Canvas_getDeviceClipBounds
*/
func (o Canvas) GetDeviceClipBounds() IRect {
	c_obj := o.sk
	retC := C.misk_Canvas_getDeviceClipBounds(c_obj)
	return IRect(retC)
}

/*
Returns SkIRect bounds of clip, unaffected by SkMatrix. If clip is empty,
return false, and set bounds to SkRect::MakeEmpty, where all SkRect sides equal zero.

Unlike getLocalClipBounds(), bounds is not outset.

@param bounds  SkRect of clip in device coordinates
@return        true if clip bounds is not empty
*/
func (o Canvas) GetDeviceClipBoundsRect(bounds *IRect) bool {
	c_obj := o.sk
	c_bounds := (*C.sk_SkIRect)(unsafe.Pointer(bounds))
	retC := C.misk_Canvas_getDeviceClipBoundsRect(c_obj, c_bounds)
	return bool(retC)
}

/*
Fills clip with color color.
mode determines how ARGB is combined with destination.

@param color  unpremultiplied ARGB
@param mode   SkBlendMode used to combine source color and destination

example: https://fiddle.skia.org/c/@Canvas_drawColor
*/
func (o Canvas) DrawColor(color Color, mode BlendMode) {
	c_obj := o.sk
	c_color := C.uint(color)
	c_mode := C.int(mode)
	C.misk_Canvas_drawColor(c_obj, c_color, c_mode)
}

/*
Fills clip with color color.
mode determines how ARGB is combined with destination.

@param color  SkColor4f representing unpremultiplied color.
@param mode   SkBlendMode used to combine source color and destination
*/
func (o Canvas) DrawColor4f(color RGBA4f, mode BlendMode) {
	c_obj := o.sk
	c_color := *(*C.sk_SkRGBA4f)(unsafe.Pointer(&color))
	c_mode := C.int(mode)
	C.misk_Canvas_drawColor4f(c_obj, c_color, c_mode)
}

/*
Fills clip with color color using SkBlendMode::kSrc.
This has the effect of replacing all pixels contained by clip with color.

@param color  unpremultiplied ARGB
*/
func (o Canvas) Clear(color Color) {
	c_obj := o.sk
	c_color := C.uint(color)
	C.misk_Canvas_clear(c_obj, c_color)
}

/*
Fills clip with color color using SkBlendMode::kSrc.
This has the effect of replacing all pixels contained by clip with color.

@param color  SkColor4f representing unpremultiplied color.
*/
func (o Canvas) Clear4f(color RGBA4f) {
	c_obj := o.sk
	c_color := *(*C.sk_SkRGBA4f)(unsafe.Pointer(&color))
	C.misk_Canvas_clear4f(c_obj, c_color)
}

/*
Makes SkCanvas contents undefined. Subsequent calls that read SkCanvas pixels,
such as drawing with SkBlendMode, return undefined results. discard() does
not change clip or SkMatrix.

discard() may do nothing, depending on the implementation of SkSurface or SkDevice
that created SkCanvas.

discard() allows optimized performance on subsequent draws by removing
cached data associated with SkSurface or SkDevice.
It is not necessary to call discard() once done with SkCanvas;
any cached data is deleted when owning SkSurface or SkDevice is deleted.
*/
func (o Canvas) Discard() {
	c_obj := o.sk
	C.misk_Canvas_discard(c_obj)
}

/*
Fills clip with SkPaint paint. SkPaint components, SkShader,
SkColorFilter, SkImageFilter, and SkBlendMode affect drawing;
SkMaskFilter and SkPathEffect in paint are ignored.

@param paint  graphics state used to fill SkCanvas

example: https://fiddle.skia.org/c/@Canvas_drawPaint
*/
func (o Canvas) DrawPaint(paint Paint) {
	c_obj := o.sk
	c_paint := paint.sk
	C.misk_Canvas_drawPaint(c_obj, c_paint)
}

/*
Draws point at (x, y) using clip, SkMatrix and SkPaint paint.

The shape of point drawn depends on paint SkPaint::Cap.
If paint is set to SkPaint::kRound_Cap, draw a circle of diameter
SkPaint stroke width. If paint is set to SkPaint::kSquare_Cap or SkPaint::kButt_Cap,
draw a square of width and height SkPaint stroke width.
SkPaint::Style is ignored, as if were set to SkPaint::kStroke_Style.

@param x      left edge of circle or square
@param y      top edge of circle or square
@param paint  stroke, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawPoint
*/
func (o Canvas) DrawPointScalars(x float32, y float32, paint Paint) {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	c_paint := paint.sk
	C.misk_Canvas_drawPointScalars(c_obj, c_x, c_y, c_paint)
}

/*
Draws point p using clip, SkMatrix and SkPaint paint.

The shape of point drawn depends on paint SkPaint::Cap.
If paint is set to SkPaint::kRound_Cap, draw a circle of diameter
SkPaint stroke width. If paint is set to SkPaint::kSquare_Cap or SkPaint::kButt_Cap,
draw a square of width and height SkPaint stroke width.
SkPaint::Style is ignored, as if were set to SkPaint::kStroke_Style.

@param p      top-left edge of circle or square
@param paint  stroke, blend, color, and so on, used to draw
*/
func (o Canvas) DrawPoint(p Point, paint Paint) {
	c_obj := o.sk
	c_p := *(*C.sk_SkPoint)(unsafe.Pointer(&p))
	c_paint := paint.sk
	C.misk_Canvas_drawPoint(c_obj, c_p, c_paint)
}

/*
Draws line segment from (x0, y0) to (x1, y1) using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint stroke width describes the line thickness;
SkPaint::Cap draws the end rounded or square;
SkPaint::Style is ignored, as if were set to SkPaint::kStroke_Style.

@param x0     start of line segment on x-axis
@param y0     start of line segment on y-axis
@param x1     end of line segment on x-axis
@param y1     end of line segment on y-axis
@param paint  stroke, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawLine
*/
func (o Canvas) DrawLineScalars(x0 float32, y0 float32, x1 float32, y1 float32, paint Paint) {
	c_obj := o.sk
	c_x0 := C.float(x0)
	c_y0 := C.float(y0)
	c_x1 := C.float(x1)
	c_y1 := C.float(y1)
	c_paint := paint.sk
	C.misk_Canvas_drawLineScalars(c_obj, c_x0, c_y0, c_x1, c_y1, c_paint)
}

/*
Draws line segment from p0 to p1 using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint stroke width describes the line thickness;
SkPaint::Cap draws the end rounded or square;
SkPaint::Style is ignored, as if were set to SkPaint::kStroke_Style.

@param p0     start of line segment
@param p1     end of line segment
@param paint  stroke, blend, color, and so on, used to draw
*/
func (o Canvas) DrawLinePoints(p0 Point, p1 Point, paint Paint) {
	c_obj := o.sk
	c_p0 := *(*C.sk_SkPoint)(unsafe.Pointer(&p0))
	c_p1 := *(*C.sk_SkPoint)(unsafe.Pointer(&p1))
	c_paint := paint.sk
	C.misk_Canvas_drawLinePoints(c_obj, c_p0, c_p1, c_paint)
}

/*
Draws SkRect rect using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint::Style determines if rectangle is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness, and
SkPaint::Join draws the corners rounded or square.

@param rect   rectangle to draw
@param paint  stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawRect
*/
func (o Canvas) DrawRect(rect Rect, paint Paint) {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_paint := paint.sk
	C.misk_Canvas_drawRect(c_obj, c_rect, c_paint)
}

/*
Draws SkIRect rect using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint::Style determines if rectangle is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness, and
SkPaint::Join draws the corners rounded or square.

@param rect   rectangle to draw
@param paint  stroke or fill, blend, color, and so on, used to draw
*/
func (o Canvas) DrawIRect(rect IRect, paint Paint) {
	c_obj := o.sk
	c_rect := *(*C.sk_SkIRect)(unsafe.Pointer(&rect))
	c_paint := paint.sk
	C.misk_Canvas_drawIRect(c_obj, c_rect, c_paint)
}

/*
Draws SkRegion region using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint::Style determines if rectangle is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness, and
SkPaint::Join draws the corners rounded or square.

@param region  region to draw
@param paint   SkPaint stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawRegion
*/
func (o Canvas) DrawRegion(region Region, paint Paint) {
	c_obj := o.sk
	c_region := region.sk
	c_paint := paint.sk
	C.misk_Canvas_drawRegion(c_obj, c_region, c_paint)
}

/*
Draws oval oval using clip, SkMatrix, and SkPaint.
In paint: SkPaint::Style determines if oval is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.

@param oval   SkRect bounds of oval
@param paint  SkPaint stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawOval
*/
func (o Canvas) DrawOval(oval Rect, paint Paint) {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_paint := paint.sk
	C.misk_Canvas_drawOval(c_obj, c_oval, c_paint)
}

/*
Draws SkRRect rrect using clip, SkMatrix, and SkPaint paint.
In paint: SkPaint::Style determines if rrect is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.

rrect may represent a rectangle, circle, oval, uniformly rounded rectangle, or
may have any combination of positive non-square radii for the four corners.

@param rrect  SkRRect with up to eight corner radii to draw
@param paint  SkPaint stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawRRect
*/
func (o Canvas) DrawRRect(rrect RRect, paint Paint) {
	c_obj := o.sk
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	c_paint := paint.sk
	C.misk_Canvas_drawRRect(c_obj, c_rrect, c_paint)
}

/*
Draws SkRRect outer and inner
using clip, SkMatrix, and SkPaint paint.
outer must contain inner or the drawing is undefined.
In paint: SkPaint::Style determines if SkRRect is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.
If stroked and SkRRect corner has zero length radii, SkPaint::Join can
draw corners rounded or square.

GPU-backed platforms optimize drawing when both outer and inner are
concave and outer contains inner. These platforms may not be able to draw
SkPath built with identical data as fast.

@param outer  SkRRect outer bounds to draw
@param inner  SkRRect inner bounds to draw
@param paint  SkPaint stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawDRRect_a
example: https://fiddle.skia.org/c/@Canvas_drawDRRect_b
*/
func (o Canvas) DrawDRRect(outer RRect, inner RRect, paint Paint) {
	c_obj := o.sk
	c_outer := *(*C.sk_SkRRect)(unsafe.Pointer(&outer))
	c_inner := *(*C.sk_SkRRect)(unsafe.Pointer(&inner))
	c_paint := paint.sk
	C.misk_Canvas_drawDRRect(c_obj, c_outer, c_inner, c_paint)
}

/*
Draws circle at (cx, cy) with radius using clip, SkMatrix, and SkPaint paint.
If radius is zero or less, nothing is drawn.
In paint: SkPaint::Style determines if circle is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.

@param cx      circle center on the x-axis
@param cy      circle center on the y-axis
@param radius  half the diameter of circle
@param paint   SkPaint stroke or fill, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawCircle
*/
func (o Canvas) DrawCircleScalars(cx float32, cy float32, radius float32, paint Paint) {
	c_obj := o.sk
	c_cx := C.float(cx)
	c_cy := C.float(cy)
	c_radius := C.float(radius)
	c_paint := paint.sk
	C.misk_Canvas_drawCircleScalars(c_obj, c_cx, c_cy, c_radius, c_paint)
}

/*
Draws circle at center with radius using clip, SkMatrix, and SkPaint paint.
If radius is zero or less, nothing is drawn.
In paint: SkPaint::Style determines if circle is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.

@param center  circle center
@param radius  half the diameter of circle
@param paint   SkPaint stroke or fill, blend, color, and so on, used to draw
*/
func (o Canvas) DrawCirclePoint(center Point, radius float32, paint Paint) {
	c_obj := o.sk
	c_center := *(*C.sk_SkPoint)(unsafe.Pointer(&center))
	c_radius := C.float(radius)
	c_paint := paint.sk
	C.misk_Canvas_drawCirclePoint(c_obj, c_center, c_radius, c_paint)
}

/*
Draws arc using clip, SkMatrix, and SkPaint paint.

Arc is part of oval bounded by oval, sweeping from startAngle to startAngle plus
sweepAngle. startAngle and sweepAngle are in degrees.

startAngle of zero places start point at the right middle edge of oval.
A positive sweepAngle places arc end point clockwise from start point;
a negative sweepAngle places arc end point counterclockwise from start point.
sweepAngle may exceed 360 degrees, a full circle.
If useCenter is true, draw a wedge that includes lines from oval
center to arc end points. If useCenter is false, draw arc between end points.

If SkRect oval is empty or sweepAngle is zero, nothing is drawn.

@param oval        SkRect bounds of oval containing arc to draw
@param startAngle  angle in degrees where arc begins
@param sweepAngle  sweep angle in degrees; positive is clockwise
@param useCenter   if true, include the center of the oval
@param paint       SkPaint stroke or fill, blend, color, and so on, used to draw
*/
func (o Canvas) DrawArc(oval Rect, startAngle float32, sweepAngle float32, useCenter bool, paint Paint) {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_startAngle := C.float(startAngle)
	c_sweepAngle := C.float(sweepAngle)
	c_useCenter := C.bool(useCenter)
	c_paint := paint.sk
	C.misk_Canvas_drawArc(c_obj, c_oval, c_startAngle, c_sweepAngle, c_useCenter, c_paint)
}

/*
Draws arc using clip, SkMatrix, and SkPaint paint.

Arc is part of oval bounded by oval, sweeping from startAngle to startAngle plus
sweepAngle. startAngle and sweepAngle are in degrees.

startAngle of zero places start point at the right middle edge of oval.
A positive sweepAngle places arc end point clockwise from start point;
a negative sweepAngle places arc end point counterclockwise from start point.
sweepAngle may exceed 360 degrees, a full circle.
If useCenter is true, draw a wedge that includes lines from oval
center to arc end points. If useCenter is false, draw arc between end points.

If SkRect oval is empty or sweepAngle is zero, nothing is drawn.

@param arc    SkArc specifying oval, startAngle, sweepAngle, and arc-vs-wedge
@param paint  SkPaint stroke or fill, blend, color, and so on, used to draw
*/
func (o Canvas) DrawArcArc(arc Arc, paint Paint) {
	c_obj := o.sk
	c_arc := arc.sk
	c_paint := paint.sk
	C.misk_Canvas_drawArcArc(c_obj, c_arc, c_paint)
}

/*
Draws SkRRect bounded by SkRect rect, with corner radii (rx, ry) using clip,
SkMatrix, and SkPaint paint.

In paint: SkPaint::Style determines if SkRRect is stroked or filled;
if stroked, SkPaint stroke width describes the line thickness.
If rx or ry are less than zero, they are treated as if they are zero.
If rx plus ry exceeds rect width or rect height, radii are scaled down to fit.
If rx and ry are zero, SkRRect is drawn as SkRect and if stroked is affected by
SkPaint::Join.

@param rect   SkRect bounds of SkRRect to draw
@param rx     axis length on x-axis of oval describing rounded corners
@param ry     axis length on y-axis of oval describing rounded corners
@param paint  stroke, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawRoundRect
*/
func (o Canvas) DrawRoundRect(rect Rect, rx float32, ry float32, paint Paint) {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_rx := C.float(rx)
	c_ry := C.float(ry)
	c_paint := paint.sk
	C.misk_Canvas_drawRoundRect(c_obj, c_rect, c_rx, c_ry, c_paint)
}

/*
Draws SkPath path using clip, SkMatrix, and SkPaint paint.
SkPath contains an array of path contour, each of which may be open or closed.

In paint: SkPaint::Style determines if SkRRect is stroked or filled:
if filled, SkPath::FillType determines whether path contour describes inside or
outside of fill; if stroked, SkPaint stroke width describes the line thickness,
SkPaint::Cap describes line ends, and SkPaint::Join describes how
corners are drawn.

@param path   SkPath to draw
@param paint  stroke, blend, color, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawPath
*/
func (o Canvas) DrawPath(path Path, paint Paint) {
	c_obj := o.sk
	c_path := path.sk
	c_paint := paint.sk
	C.misk_Canvas_drawPath(c_obj, c_path, c_paint)
}

func (o Canvas) DrawImage(image Image, left float32, top float32) {
	c_obj := o.sk
	c_image := image.sk
	c_left := C.float(left)
	c_top := C.float(top)
	C.misk_Canvas_drawImage(c_obj, c_image, c_left, c_top)
}

func (o Canvas) DrawImageSamplingOptions(p0 Image, x float32, y float32, p3 SamplingOptions, p4 Paint) {
	c_obj := o.sk
	c_p0 := p0.sk
	c_x := C.float(x)
	c_y := C.float(y)
	c_p3 := p3.sk
	c_p4 := p4.sk
	C.misk_Canvas_drawImageSamplingOptions(c_obj, c_p0, c_x, c_y, c_p3, c_p4)
}

func (o Canvas) DrawImageRect(p0 Image, src Rect, dst Rect, p3 SamplingOptions, p4 Paint, p5 CanvasSrcRectConstraint) {
	c_obj := o.sk
	c_p0 := p0.sk
	c_src := *(*C.sk_SkRect)(unsafe.Pointer(&src))
	c_dst := *(*C.sk_SkRect)(unsafe.Pointer(&dst))
	c_p3 := p3.sk
	c_p4 := p4.sk
	c_p5 := C.uint(p5)
	C.misk_Canvas_drawImageRect(c_obj, c_p0, c_src, c_dst, c_p3, c_p4, c_p5)
}

func (o Canvas) DrawImageRectNoSrc(p0 Image, dst Rect, p2 SamplingOptions, p3 Paint) {
	c_obj := o.sk
	c_p0 := p0.sk
	c_dst := *(*C.sk_SkRect)(unsafe.Pointer(&dst))
	c_p2 := p2.sk
	c_p3 := p3.sk
	C.misk_Canvas_drawImageRectNoSrc(c_obj, c_p0, c_dst, c_p2, c_p3)
}

/*
Draws SkImage image stretched proportionally to fit into SkRect dst.
SkIRect center divides the image into nine sections: four sides, four corners, and
the center. Corners are unmodified or scaled down proportionately if their sides
are larger than dst; center and four sides are scaled to fit remaining space, if any.

Additionally transform draw using clip, SkMatrix, and optional SkPaint paint.

If SkPaint paint is supplied, apply SkColorFilter, alpha, SkImageFilter, and
SkBlendMode. If image is kAlpha_8_SkColorType, apply SkShader.
If paint contains SkMaskFilter, generate mask from image bounds.
Any SkMaskFilter on paint is ignored as is paint anti-aliasing state.

If generated mask extends beyond image bounds, replicate image edge colors, just
as SkShader made from SkImage::makeShader with SkShader::kClamp_TileMode set
replicates the image edge color when it samples outside of its bounds.

@param image   SkImage containing pixels, dimensions, and format
@param center  SkIRect edge of image corners and sides
@param dst     destination SkRect of image to draw to
@param filter  what technique to use when sampling the image
@param paint   SkPaint containing SkBlendMode, SkColorFilter, SkImageFilter,
and so on; or nullptr
*/
func (o Canvas) DrawImageNine(image Image, center IRect, dst Rect, filter FilterMode, paint Paint) {
	c_obj := o.sk
	c_image := image.sk
	c_center := *(*C.sk_SkIRect)(unsafe.Pointer(&center))
	c_dst := *(*C.sk_SkRect)(unsafe.Pointer(&dst))
	c_filter := C.int(filter)
	c_paint := paint.sk
	C.misk_Canvas_drawImageNine(c_obj, c_image, c_center, c_dst, c_filter, c_paint)
}

/*
Draws null terminated string, with origin at (x, y), using clip, SkMatrix,
SkFont font, and SkPaint paint.

This function uses the default character-to-glyph mapping from the
SkTypeface in font.  It does not perform typeface fallback for
characters not found in the SkTypeface.  It does not perform kerning;
glyphs are positioned based on their default advances.

String str is encoded as UTF-8.

Text size is affected by SkMatrix and font text size. Default text
size is 12 point.

All elements of paint: SkPathEffect, SkMaskFilter, SkShader,
SkColorFilter, and SkImageFilter; apply to text. By
default, draws filled black glyphs.

@param str     character code points drawn,
ending with a char value of zero
@param x       start of string on x-axis
@param y       start of string on y-axis
@param font    typeface, text size and so, used to describe the text
@param paint   blend, color, and so on, used to draw
*/
func (o Canvas) DrawString(str string, x float32, y float32, font Font, paint Paint) {
	c_obj := o.sk
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))
	c_x := C.float(x)
	c_y := C.float(y)
	c_font := font.sk
	c_paint := paint.sk
	C.misk_Canvas_drawString(c_obj, c_str, c_x, c_y, c_font, c_paint)
}

/*
Draws count glyphs, at positions relative to origin styled with font and paint with
supporting utf8 and cluster information.

This function draw glyphs at the given positions relative to the given origin.
It does not perform typeface fallback for glyphs not found in the SkTypeface in font.

The drawing obeys the current transform matrix and clipping.

All elements of paint: SkPathEffect, SkMaskFilter, SkShader,
SkColorFilter, and SkImageFilter; apply to text. By
default, draws filled black glyphs.

@param count           number of glyphs to draw
@param glyphs          the array of glyphIDs to draw
@param positions       where to draw each glyph relative to origin
@param clusters        array of size count of cluster information
@param textByteCount   size of the utf8text
@param utf8text        utf8text supporting information for the glyphs
@param origin          the origin of all the positions
@param font            typeface, text size and so, used to describe the text
@param paint           blend, color, and so on, used to draw
*/
func (o Canvas) DrawGlyphs(count int32, glyphs []uint16, positions []Point, clusters []uint32, textByteCount int32, utf8text string, origin Point, font Font, paint Paint) {
	c_obj := o.sk
	c_count := C.int(count)
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_positions := (*C.sk_SkPoint)(unsafe.Pointer(&positions[0]))
	c_clusters := (*C.uint)(unsafe.Pointer(&clusters[0]))
	c_textByteCount := C.int(textByteCount)
	c_utf8text := C.CString(utf8text)
	defer C.free(unsafe.Pointer(c_utf8text))
	c_origin := *(*C.sk_SkPoint)(unsafe.Pointer(&origin))
	c_font := font.sk
	c_paint := paint.sk
	C.misk_Canvas_drawGlyphs(c_obj, c_count, c_glyphs, c_positions, c_clusters, c_textByteCount, c_utf8text, c_origin, c_font, c_paint)
}

/*
Draws SkTextBlob blob at (x, y), using clip, SkMatrix, and SkPaint paint.

blob contains glyphs, their positions, and paint attributes specific to text:
SkTypeface, SkPaint text size, SkPaint text scale x,
SkPaint text skew x, SkPaint::Align, SkPaint::Hinting, anti-alias, SkPaint fake bold,
SkPaint font embedded bitmaps, SkPaint full hinting spacing, LCD text, SkPaint linear text,
and SkPaint subpixel text.

SkTextEncoding must be set to SkTextEncoding::kGlyphID.

Elements of paint: anti-alias, SkBlendMode, color including alpha,
SkColorFilter, SkPaint dither, SkMaskFilter, SkPathEffect, SkShader, and
SkPaint::Style; apply to blob. If SkPaint contains SkPaint::kStroke_Style:
SkPaint miter limit, SkPaint::Cap, SkPaint::Join, and SkPaint stroke width;
apply to SkPath created from blob.

@param blob   glyphs, positions, and their paints' text size, typeface, and so on
@param x      horizontal offset applied to blob
@param y      vertical offset applied to blob
@param paint  blend, color, stroking, and so on, used to draw

example: https://fiddle.skia.org/c/@Canvas_drawTextBlob
*/
func (o Canvas) DrawTextBlob(blob TextBlob, x float32, y float32, paint Paint) {
	c_obj := o.sk
	c_blob := blob.sk
	c_x := C.float(x)
	c_y := C.float(y)
	c_paint := paint.sk
	C.misk_Canvas_drawTextBlob(c_obj, c_blob, c_x, c_y, c_paint)
}

type CanvasClipEdgeStyle uint32

const (
	CanvasClipEdgeStyleHard CanvasClipEdgeStyle = 0
	CanvasClipEdgeStyleSoft CanvasClipEdgeStyle = 1
)

/*
Selects if an array of points are drawn as discrete points, as lines, or as
an open polygon.
*/
type CanvasPointMode uint32

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
type CanvasQuadAAFlags uint32

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
type CanvasSaveLayerFlagsSet uint32

const (
	CanvasSaveLayerFlagsSetPreserveLCDText_SaveLayerFlag CanvasSaveLayerFlagsSet = 2
	// initializes with previous contents
	CanvasSaveLayerFlagsSetInitWithPrevious_SaveLayerFlag CanvasSaveLayerFlagsSet = 4
	CanvasSaveLayerFlagsSetF16ColorType                   CanvasSaveLayerFlagsSet = 16
)

type CanvasSaveLayerStrategy uint32

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
type CanvasSrcRectConstraint uint32

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

func (o *ColorSpace) Unref() {
	C.misk_unref_SkColorSpace(o.sk)
	o.sk = nil
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
SkData holds an immutable data buffer. Not only is the data immutable,
but the actual ptr that is returned (by data() or bytes()) is guaranteed
to always be the same for the life of this instance.
*/
type Data struct {
	sk *C.sk_SkData
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Data has not been created.
func (o Data) IsNil() bool {
	return o.sk == nil
}

func (o *Data) Unref() {
	C.misk_unref_SkData(o.sk)
	o.sk = nil
}

/*
Returns the number of bytes stored.
*/
func (o Data) Size() uint32 {
	c_obj := o.sk
	retC := C.misk_Data_size(c_obj)
	return uint32(retC)
}

func (o Data) IsEmpty() bool {
	c_obj := o.sk
	retC := C.misk_Data_isEmpty(c_obj)
	return bool(retC)
}

/*
Create a new dataref by copying the specified data
*/
func DataMakeWithCopy(data []byte, length uint32) Data {
	c_data := unsafe.Pointer(&data[0])
	c_length := C.ulong(length)
	retC := C.misk_Data_MakeWithCopy(c_data, c_length)
	return Data{sk: retC}
}

/*
Create a new data with zero-initialized contents. The caller should call writable_data()
to write into the buffer, but this must be done before another ref() is made.
*/
func DataMakeZeroInitialized(length uint32) Data {
	c_length := C.ulong(length)
	retC := C.misk_Data_MakeZeroInitialized(c_length)
	return Data{sk: retC}
}

/*
Create a new dataref by copying the specified c-string
(a null-terminated array of bytes). The returned SkData will have size()
equal to strlen(cstr) + 1. If cstr is NULL, it will be treated the same
as "".
*/
func DataMakeWithCString(cstr string) Data {
	c_cstr := C.CString(cstr)
	defer C.free(unsafe.Pointer(c_cstr))
	retC := C.misk_Data_MakeWithCString(c_cstr)
	return Data{sk: retC}
}

/*
Call this when the data parameter is already const and will outlive the lifetime of the
SkData. Suitable for with const globals.
*/
func DataMakeWithoutCopy(data []byte, length uint32) Data {
	c_data := unsafe.Pointer(&data[0])
	c_length := C.ulong(length)
	retC := C.misk_Data_MakeWithoutCopy(c_data, c_length)
	return Data{sk: retC}
}

/*
Create a new dataref the file with the specified path.
If the file cannot be opened, this returns NULL.
*/
func DataMakeFromFileName(path string) Data {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	retC := C.misk_Data_MakeFromFileName(c_path)
	return Data{sk: retC}
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

func (o *Font) Delete() {
	C.misk_delete_SkFont(o.sk)
	o.sk = nil
}

/*
Returns SkFontMetrics associated with SkTypeface.
The return value is the recommended spacing between lines: the sum of metrics
descent, ascent, and leading.
If metrics is not nullptr, SkFontMetrics is copied to metrics.
Results are scaled by text size but does not take into account
dimensions required by text scale, text skew, fake bold,
style stroke, and SkPathEffect.

@param metrics  storage for SkFontMetrics; may be nullptr
@return         recommended spacing between lines
*/
func (o Font) GetMetrics(metrics *FontMetrics) float32 {
	c_obj := o.sk
	c_metrics := (*C.sk_SkFontMetrics)(unsafe.Pointer(metrics))
	retC := C.misk_Font_getMetrics(c_obj, c_metrics)
	return float32(retC)
}

/*
Retrieves the x-positions for each glyph, beginning at the specified origin. The caller
must allocated at least count number of elements in the xpos[] array.

@param glyphs   array of glyph indices to be positioned
@param count    number of glyphs
@param xpos     returns glyphs x-positions
@param origin   x-position of the first glyph. Defaults to 0.
*/
func (o Font) GetXPos(glyphs []uint16, count int32, xpos []float32, origin float32) {
	c_obj := o.sk
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_count := C.int(count)
	c_xpos := (*C.float)(unsafe.Pointer(&xpos[0]))
	c_origin := C.float(origin)
	C.misk_Font_getXPos(c_obj, c_glyphs, c_count, c_xpos, c_origin)
}

/*
Returns the advance width of text.
The advance is the normal distance to move before drawing additional text.
Returns the bounding box of text if bounds is not nullptr.

@param text        character storage encoded with SkTextEncoding
@param byteLength  length of character storage in bytes
@param bounds      returns bounding box relative to (0, 0) if not nullptr
@return            the sum of the default advance widths
*/
func (o Font) MeasureText(text []byte, byteLength uint32, encoding TextEncoding, bounds *Rect) float32 {
	c_obj := o.sk
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_encoding := C.int(encoding)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_Font_measureText(c_obj, c_text, c_byteLength, c_encoding, c_bounds)
	return float32(retC)
}

/*
Returns the advance width of text.
The advance is the normal distance to move before drawing additional text.
Returns the bounding box of text if bounds is not nullptr. The paint
stroke settings, mask filter, or path effect may modify the bounds.

@param text        character storage encoded with SkTextEncoding
@param byteLength  length of character storage in bytes
@param bounds      returns bounding box relative to (0, 0) if not nullptr
@param paint       optional; may be nullptr
@return            the sum of the default advance widths
*/
func (o Font) MeasureTextPaint(text []byte, byteLength uint32, encoding TextEncoding, bounds *Rect, paint Paint) float32 {
	c_obj := o.sk
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_encoding := C.int(encoding)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	c_paint := paint.sk
	retC := C.misk_Font_measureTextPaint(c_obj, c_text, c_byteLength, c_encoding, c_bounds, c_paint)
	return float32(retC)
}

/*
Sets whether to always hint glyphs.
If forceAutoHinting is set, instructs the font manager to always hint glyphs.

Only affects platforms that use FreeType as the font manager.

@param forceAutoHinting  setting to always hint glyphs
*/
func (o Font) SetForceAutoHinting(forceAutoHinting bool) {
	c_obj := o.sk
	c_forceAutoHinting := C.bool(forceAutoHinting)
	C.misk_Font_setForceAutoHinting(c_obj, c_forceAutoHinting)
}

/*
Sets level of glyph outline adjustment.
Does not check for valid values of hintingLevel.
*/
func (o Font) SetHinting(hintingLevel FontHinting) {
	c_obj := o.sk
	c_hintingLevel := C.int(hintingLevel)
	C.misk_Font_setHinting(c_obj, c_hintingLevel)
}

/*
Requests, but does not require, that glyphs respect sub-pixel positioning.

@param subpixel  setting for sub-pixel positioning
*/
func (o Font) SetSubpixel(subpixel bool) {
	c_obj := o.sk
	c_subpixel := C.bool(subpixel)
	C.misk_Font_setSubpixel(c_obj, c_subpixel)
}

/*
Converts text into glyph indices.
Returns the number of glyph indices represented by text.
SkTextEncoding specifies how text represents characters or glyphs.
glyphs may be nullptr, to compute the glyph count.

Does not check text for valid character codes or valid glyph indices.

If byteLength equals zero, returns zero.
If byteLength includes a partial character, the partial character is ignored.

If encoding is SkTextEncoding::kUTF8 and text contains an invalid UTF-8 sequence,
zero is returned.

When encoding is SkTextEncoding::kUTF8, SkTextEncoding::kUTF16, or
SkTextEncoding::kUTF32; then each Unicode codepoint is mapped to a
single glyph.  This function uses the default character-to-glyph
mapping from the SkTypeface and maps characters not found in the
SkTypeface to zero.

If maxGlyphCount is not sufficient to store all the glyphs, no glyphs are copied.
The total glyph count is returned for subsequent buffer reallocation.

@param text          character storage encoded with SkTextEncoding
@param byteLength    length of character storage in bytes
@param glyphs        storage for glyph indices; may be nullptr
@param maxGlyphCount storage capacity
@return              number of glyphs represented by text of length byteLength
*/
func (o Font) TextToGlyphs(text []byte, byteLength uint32, encoding TextEncoding, glyphs []uint16, maxGlyphCount int32) int32 {
	c_obj := o.sk
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_encoding := C.int(encoding)
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_maxGlyphCount := C.int(maxGlyphCount)
	retC := C.misk_Font_textToGlyphs(c_obj, c_text, c_byteLength, c_encoding, c_glyphs, c_maxGlyphCount)
	return int32(retC)
}

/*
Returns glyph index for Unicode character.

If the character is not supported by the SkTypeface, returns 0.

@param uni  Unicode character
@return     glyph index
*/
func (o Font) UnicharToGlyph(uni int32) uint16 {
	c_obj := o.sk
	c_uni := C.int(uni)
	retC := C.misk_Font_unicharToGlyph(c_obj, c_uni)
	return uint16(retC)
}

func (o Font) UnicharsToGlyphs(uni []int32, count int32, glyphs []uint16) {
	c_obj := o.sk
	c_uni := (*C.int)(unsafe.Pointer(&uni[0]))
	c_count := C.int(count)
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	C.misk_Font_unicharsToGlyphs(c_obj, c_uni, c_count, c_glyphs)
}

/*
Returns text size in points.

@return  typographic height of text
*/
func (o Font) GetSize() float32 {
	c_obj := o.sk
	retC := C.misk_Font_getSize(c_obj)
	return float32(retC)
}

/*
Does not alter SkTypeface SkRefCnt.

@return  non-null SkTypeface
*/
func (o Font) GetTypeface() Typeface {
	c_obj := o.sk
	retC := C.misk_Font_getTypeface(c_obj)
	return Typeface{sk: retC}
}

/*
DEPRECATED
Retrieves the advance and bounds for each glyph in glyphs.
Both widths and bounds may be nullptr.
If widths is not nullptr, widths must be an array of count entries.
if bounds is not nullptr, bounds must be an array of count entries.

@param glyphs      array of glyph indices to be measured
@param count       number of glyphs
@param widths      returns text advances for each glyph; may be nullptr
@param bounds      returns bounds for each glyph relative to (0, 0); may be nullptr
*/
func (o Font) GetWidthsBounds(glyphs []uint16, count int32, widths []float32, bounds []Rect) {
	c_obj := o.sk
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_count := C.int(count)
	c_widths := (*C.float)(unsafe.Pointer(&widths[0]))
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(&bounds[0]))
	C.misk_Font_getWidthsBounds(c_obj, c_glyphs, c_count, c_widths, c_bounds)
}

/*
Retrieves the advance and bounds for each glyph in glyphs.
Both widths and bounds may be nullptr.
If widths is not nullptr, widths must be an array of count entries.
if bounds is not nullptr, bounds must be an array of count entries.

@param glyphs      array of glyph indices to be measured
@param count       number of glyphs
@param widths      returns text advances for each glyph
*/
func (o Font) GetWidths(glyphs []uint16, count int32, widths []float32) {
	c_obj := o.sk
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_count := C.int(count)
	c_widths := (*C.float)(unsafe.Pointer(&widths[0]))
	C.misk_Font_getWidths(c_obj, c_glyphs, c_count, c_widths)
}

/*
Represents a set of actual arguments for a font.
*/
type FontArguments struct {
	sk *C.sk_SkFontArguments
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontArguments has not been created.
func (o FontArguments) IsNil() bool {
	return o.sk == nil
}

func NewFontArguments() FontArguments {

	retC := C.misk_new_FontArguments()
	return FontArguments{sk: retC}
}

func (o *FontArguments) Delete() {
	C.misk_delete_SkFontArguments(o.sk)
	o.sk = nil
}

/*
Specify the index of the desired font.

	Font formats like ttc, dfont, cff, cid, pfr, t42, t1, and fon may actually be indexed
	collections of fonts.
*/
func (o FontArguments) SetCollectionIndex(collectionIndex int32) FontArguments {
	c_obj := o.sk
	c_collectionIndex := C.int(collectionIndex)
	retC := C.misk_FontArguments_setCollectionIndex(c_obj, c_collectionIndex)
	return FontArguments{sk: &retC}
}

func (o FontArguments) GetCollectionIndex() int32 {
	c_obj := o.sk
	retC := C.misk_FontArguments_getCollectionIndex(c_obj)
	return int32(retC)
}

func (o FontArguments) GetVariationDesignPosition() FontArgumentsVariationPosition {
	c_obj := o.sk
	retC := C.misk_FontArguments_getVariationDesignPosition(c_obj)
	return FontArgumentsVariationPosition{sk: &retC}
}

func (o FontArguments) GetPalette() FontArgumentsPalette {
	c_obj := o.sk
	retC := C.misk_FontArguments_getPalette(c_obj)
	return FontArgumentsPalette{sk: &retC}
}

type FontArgumentsVariationPosition struct {
	sk *C.sk_SkFontArgumentsVariationPosition
}

func (o FontArgumentsVariationPosition) CoordinateCount() int32 {
	return int32(o.sk.coordinateCount)
}

func (o *FontArgumentsVariationPosition) SetcoordinateCount(value int32) {
	o.sk.coordinateCount = C.int(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontArgumentsVariationPosition has not been created.
func (o FontArgumentsVariationPosition) IsNil() bool {
	return o.sk == nil
}

/*
Specify a palette to use and overrides for palette entries.

	`overrides` is a list of pairs of palette entry index and color.
	The overriden palette entries will use the associated color.
	Override pairs with palette entry indices out of range will not be applied.
	Later override entries override earlier ones.
*/
type FontArgumentsPalette struct {
	sk *C.sk_SkFontArgumentsPalette
}

func (o FontArgumentsPalette) Index() int32 {
	return int32(o.sk.index)
}

func (o *FontArgumentsPalette) Setindex(value int32) {
	o.sk.index = C.int(value)
}

func (o FontArgumentsPalette) OverrideCount() int32 {
	return int32(o.sk.overrideCount)
}

func (o *FontArgumentsPalette) SetoverrideCount(value int32) {
	o.sk.overrideCount = C.int(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontArgumentsPalette has not been created.
func (o FontArgumentsPalette) IsNil() bool {
	return o.sk == nil
}

/*
The metrics of an SkFont.
The metric values are consistent with the Skia y-down coordinate system.
*/
type FontMetrics C.sk_SkFontMetrics

// !< FontMetricsFlags indicating which metrics are valid
func (o FontMetrics) Flags() uint32 {
	return uint32(o.fFlags)
}

func (o *FontMetrics) SetFlags(value uint32) {
	o.fFlags = C.uint(value)
}

// !< greatest extent above origin of any glyph bounding box, typically negative; deprecated with variable fonts
func (o FontMetrics) Top() float32 {
	return float32(o.fTop)
}

func (o *FontMetrics) SetTop(value float32) {
	o.fTop = C.float(value)
}

// !< distance to reserve above baseline, typically negative
func (o FontMetrics) Ascent() float32 {
	return float32(o.fAscent)
}

func (o *FontMetrics) SetAscent(value float32) {
	o.fAscent = C.float(value)
}

// !< distance to reserve below baseline, typically positive
func (o FontMetrics) Descent() float32 {
	return float32(o.fDescent)
}

func (o *FontMetrics) SetDescent(value float32) {
	o.fDescent = C.float(value)
}

// !< greatest extent below origin of any glyph bounding box, typically positive; deprecated with variable fonts
func (o FontMetrics) Bottom() float32 {
	return float32(o.fBottom)
}

func (o *FontMetrics) SetBottom(value float32) {
	o.fBottom = C.float(value)
}

// !< distance to add between lines, typically positive or zero
func (o FontMetrics) Leading() float32 {
	return float32(o.fLeading)
}

func (o *FontMetrics) SetLeading(value float32) {
	o.fLeading = C.float(value)
}

// !< average character width, zero if unknown
func (o FontMetrics) AvgCharWidth() float32 {
	return float32(o.fAvgCharWidth)
}

func (o *FontMetrics) SetAvgCharWidth(value float32) {
	o.fAvgCharWidth = C.float(value)
}

// !< maximum character width, zero if unknown
func (o FontMetrics) MaxCharWidth() float32 {
	return float32(o.fMaxCharWidth)
}

func (o *FontMetrics) SetMaxCharWidth(value float32) {
	o.fMaxCharWidth = C.float(value)
}

// !< greatest extent to left of origin of any glyph bounding box, typically negative; deprecated with variable fonts
func (o FontMetrics) XMin() float32 {
	return float32(o.fXMin)
}

func (o *FontMetrics) SetXMin(value float32) {
	o.fXMin = C.float(value)
}

// !< greatest extent to right of origin of any glyph bounding box, typically positive; deprecated with variable fonts
func (o FontMetrics) XMax() float32 {
	return float32(o.fXMax)
}

func (o *FontMetrics) SetXMax(value float32) {
	o.fXMax = C.float(value)
}

// !< height of lower-case 'x', zero if unknown, typically negative
func (o FontMetrics) XHeight() float32 {
	return float32(o.fXHeight)
}

func (o *FontMetrics) SetXHeight(value float32) {
	o.fXHeight = C.float(value)
}

// !< height of an upper-case letter, zero if unknown, typically negative
func (o FontMetrics) CapHeight() float32 {
	return float32(o.fCapHeight)
}

func (o *FontMetrics) SetCapHeight(value float32) {
	o.fCapHeight = C.float(value)
}

// !< underline thickness
func (o FontMetrics) UnderlineThickness() float32 {
	return float32(o.fUnderlineThickness)
}

func (o *FontMetrics) SetUnderlineThickness(value float32) {
	o.fUnderlineThickness = C.float(value)
}

// !< distance from baseline to top of stroke, typically positive
func (o FontMetrics) UnderlinePosition() float32 {
	return float32(o.fUnderlinePosition)
}

func (o *FontMetrics) SetUnderlinePosition(value float32) {
	o.fUnderlinePosition = C.float(value)
}

// !< strikeout thickness
func (o FontMetrics) StrikeoutThickness() float32 {
	return float32(o.fStrikeoutThickness)
}

func (o *FontMetrics) SetStrikeoutThickness(value float32) {
	o.fStrikeoutThickness = C.float(value)
}

// !< distance from baseline to bottom of stroke, typically negative
func (o FontMetrics) StrikeoutPosition() float32 {
	return float32(o.fStrikeoutPosition)
}

func (o *FontMetrics) SetStrikeoutPosition(value float32) {
	o.fStrikeoutPosition = C.float(value)
}

type FontMgr struct {
	sk *C.sk_SkFontMgr
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontMgr has not been created.
func (o FontMgr) IsNil() bool {
	return o.sk == nil
}

func (o *FontMgr) Unref() {
	C.misk_unref_SkFontMgr(o.sk)
	o.sk = nil
}

/*
The caller must call unref() on the returned object.
Never returns NULL; will return an empty set if the name is not found.

Passing nullptr as the parameter will return the default system family.
Note that most systems don't have a default system family, so passing nullptr will often
result in the empty set.

It is possible that this will return a style set not accessible from
createStyleSet(int) due to hidden or auto-activated fonts.
*/
func (o FontMgr) MatchFamily(familyName string) FontStyleSet {
	c_obj := o.sk
	c_familyName := C.CString(familyName)
	defer C.free(unsafe.Pointer(c_familyName))
	retC := C.misk_FontMgr_matchFamily(c_obj, c_familyName)
	return FontStyleSet{sk: retC}
}

/*
Find the closest matching typeface to the specified familyName and style
and return a ref to it. The caller must call unref() on the returned
object. Will return nullptr if no 'good' match is found.

Passing |nullptr| as the parameter for |familyName| will return the
default system font.

It is possible that this will return a style set not accessible from
createStyleSet(int) or matchFamily(const char[]) due to hidden or
auto-activated fonts.
*/
func (o FontMgr) MatchFamilyStyle(familyName string, p1 FontStyle) Typeface {
	c_obj := o.sk
	c_familyName := C.CString(familyName)
	defer C.free(unsafe.Pointer(c_familyName))
	c_p1 := p1.sk
	retC := C.misk_FontMgr_matchFamilyStyle(c_obj, c_familyName, c_p1)
	return Typeface{sk: retC}
}

/*
Create a typeface for the specified data and TTC index (pass 0 for none)
or NULL if the data is not recognized. The caller must call unref() on
the returned object if it is not null.
*/
func (o FontMgr) MakeFromData(p0 Data, ttcIndex int32) Typeface {
	c_obj := o.sk
	c_p0 := p0.sk
	c_ttcIndex := C.int(ttcIndex)
	retC := C.misk_FontMgr_makeFromData(c_obj, c_p0, c_ttcIndex)
	return Typeface{sk: retC}
}

/*
Create a typeface for the specified fileName and TTC index
(pass 0 for none) or NULL if the file is not found, or its contents are
not recognized. The caller must call unref() on the returned object
if it is not null.
*/
func (o FontMgr) MakeFromFile(path string, ttcIndex int32) Typeface {
	c_obj := o.sk
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	c_ttcIndex := C.int(ttcIndex)
	retC := C.misk_FontMgr_makeFromFile(c_obj, c_path, c_ttcIndex)
	return Typeface{sk: retC}
}

type FontStyle struct {
	sk *C.sk_SkFontStyle
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontStyle has not been created.
func (o FontStyle) IsNil() bool {
	return o.sk == nil
}

func NewFontStyle2(weight int32, width int32, slant FontStyleSlant) FontStyle {
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

func (o *FontStyle) Delete() {
	C.misk_delete_SkFontStyle(o.sk)
	o.sk = nil
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

type FontStyleSlant uint32

const (
	FontStyleSlantUpright FontStyleSlant = 0
	FontStyleSlantItalic  FontStyleSlant = 1
	FontStyleSlantOblique FontStyleSlant = 2
)

type FontStyleWeight uint32

const (
	FontStyleWeightInvisible  FontStyleWeight = 0
	FontStyleWeightThin       FontStyleWeight = 100
	FontStyleWeightExtraLight FontStyleWeight = 200
	FontStyleWeightLight      FontStyleWeight = 300
	FontStyleWeightNormal     FontStyleWeight = 400
	FontStyleWeightMedium     FontStyleWeight = 500
	FontStyleWeightSemiBold   FontStyleWeight = 600
	FontStyleWeightBold       FontStyleWeight = 700
	FontStyleWeightExtraBold  FontStyleWeight = 800
	FontStyleWeightBlack      FontStyleWeight = 900
	FontStyleWeightExtraBlack FontStyleWeight = 1000
)

type FontStyleWidth uint32

const (
	FontStyleWidthUltraCondensed FontStyleWidth = 1
	FontStyleWidthExtraCondensed FontStyleWidth = 2
	FontStyleWidthCondensed      FontStyleWidth = 3
	FontStyleWidthSemiCondensed  FontStyleWidth = 4
	FontStyleWidthNormal         FontStyleWidth = 5
	FontStyleWidthSemiExpanded   FontStyleWidth = 6
	FontStyleWidthExpanded       FontStyleWidth = 7
	FontStyleWidthExtraExpanded  FontStyleWidth = 8
	FontStyleWidthUltraExpanded  FontStyleWidth = 9
)

type FontStyleSet struct {
	sk *C.sk_SkFontStyleSet
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the FontStyleSet has not been created.
func (o FontStyleSet) IsNil() bool {
	return o.sk == nil
}

func (o *FontStyleSet) Unref() {
	C.misk_unref_SkFontStyleSet(o.sk)
	o.sk = nil
}
func (o FontStyleSet) Count() int32 {
	c_obj := o.sk
	retC := C.misk_FontStyleSet_count(c_obj)
	return int32(retC)
}

func (o FontStyleSet) GetStyle(index int32, p1 FontStyle, style String) {
	c_obj := o.sk
	c_index := C.int(index)
	c_p1 := p1.sk
	c_style := style.sk
	C.misk_FontStyleSet_getStyle(c_obj, c_index, c_p1, c_style)
}

func (o FontStyleSet) CreateTypeface(index int32) Typeface {
	c_obj := o.sk
	c_index := C.int(index)
	retC := C.misk_FontStyleSet_createTypeface(c_obj, c_index)
	return Typeface{sk: retC}
}

func (o FontStyleSet) MatchStyle(pattern FontStyle) Typeface {
	c_obj := o.sk
	c_pattern := pattern.sk
	retC := C.misk_FontStyleSet_matchStyle(c_obj, c_pattern)
	return Typeface{sk: retC}
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
	sk *C.sk_SkImage
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Image has not been created.
func (o Image) IsNil() bool {
	return o.sk == nil
}

func (o *Image) Unref() {
	C.misk_unref_SkImage(o.sk)
	o.sk = nil
}

/*
Returns a SkImageInfo describing the width, height, color type, alpha type, and color space
of the SkImage.

@return  image info of SkImage.
*/
func (o Image) ImageInfo() ImageInfo {
	c_obj := o.sk
	retC := C.misk_Image_imageInfo(c_obj)
	return ImageInfo{sk: &retC}
}

/*
Returns pixel count in each row.

@return  pixel width in SkImage
*/
func (o Image) Width() int32 {
	c_obj := o.sk
	retC := C.misk_Image_width(c_obj)
	return int32(retC)
}

/*
Returns pixel row count.

@return  pixel height in SkImage
*/
func (o Image) Height() int32 {
	c_obj := o.sk
	retC := C.misk_Image_height(c_obj)
	return int32(retC)
}

/*
Returns SkISize { width(), height() }.

@return  integral size of width() and height()
*/
func (o Image) Dimensions() ISize {
	c_obj := o.sk
	retC := C.misk_Image_dimensions(c_obj)
	return ISize(retC)
}

/*
Returns SkIRect { 0, 0, width(), height() }.

@return  integral rectangle from origin to width() and height()
*/
func (o Image) Bounds() IRect {
	c_obj := o.sk
	retC := C.misk_Image_bounds(c_obj)
	return IRect(retC)
}

/*
Returns value unique to image. SkImage contents cannot change after SkImage is
created. Any operation to create a new SkImage will receive generate a new
unique number.

@return  unique identifier
*/
func (o Image) UniqueID() uint32 {
	c_obj := o.sk
	retC := C.misk_Image_uniqueID(c_obj)
	return uint32(retC)
}

/*
Returns SkAlphaType.

SkAlphaType returned was a parameter to an SkImage constructor,
or was parsed from encoded data.

@return  SkAlphaType in SkImage

example: https://fiddle.skia.org/c/@Image_alphaType
*/
func (o Image) AlphaType() AlphaType {
	c_obj := o.sk
	retC := C.misk_Image_alphaType(c_obj)
	return AlphaType(retC)
}

/*
Returns SkColorType if known; otherwise, returns kUnknown_SkColorType.

@return  SkColorType of SkImage

example: https://fiddle.skia.org/c/@Image_colorType
*/
func (o Image) ColorType() ColorType {
	c_obj := o.sk
	retC := C.misk_Image_colorType(c_obj)
	return ColorType(retC)
}

/*
Returns SkColorSpace, the range of colors, associated with SkImage.  The
reference count of SkColorSpace is unchanged. The returned SkColorSpace is
immutable.

SkColorSpace returned was passed to an SkImage constructor,
or was parsed from encoded data. SkColorSpace returned may be ignored when SkImage
is drawn, depending on the capabilities of the SkSurface receiving the drawing.

@return  SkColorSpace in SkImage, or nullptr

example: https://fiddle.skia.org/c/@Image_colorSpace
*/
func (o Image) ColorSpace() ColorSpace {
	c_obj := o.sk
	retC := C.misk_Image_colorSpace(c_obj)
	return ColorSpace{sk: retC}
}

/*
Returns true if SkImage pixels represent transparency only. If true, each pixel
is packed in 8 bits as defined by kAlpha_8_SkColorType.

@return  true if pixels represent a transparency mask

example: https://fiddle.skia.org/c/@Image_isAlphaOnly
*/
func (o Image) IsAlphaOnly() bool {
	c_obj := o.sk
	retC := C.misk_Image_isAlphaOnly(c_obj)
	return bool(retC)
}

/*
Returns true if pixels ignore their alpha value and are treated as fully opaque.

@return  true if SkAlphaType is kOpaque_SkAlphaType
*/
func (o Image) IsOpaque() bool {
	c_obj := o.sk
	retC := C.misk_Image_isOpaque(c_obj)
	return bool(retC)
}

/*
Copies SkRect of pixels from SkImage to dstPixels. Copy starts at offset (srcX, srcY),
and does not exceed SkImage (width(), height()).

dstInfo specifies width, height, SkColorType, SkAlphaType, and SkColorSpace of
destination. dstRowBytes specifies the gap from one destination row to the next.
Returns true if pixels are copied. Returns false if:
- dstInfo.addr() equals nullptr
- dstRowBytes is less than dstInfo.minRowBytes()
- SkPixelRef is nullptr

Pixels are copied only if pixel conversion is possible. If SkImage SkColorType is
kGray_8_SkColorType, or kAlpha_8_SkColorType; dstInfo.colorType() must match.
If SkImage SkColorType is kGray_8_SkColorType, dstInfo.colorSpace() must match.
If SkImage SkAlphaType is kOpaque_SkAlphaType, dstInfo.alphaType() must
match. If SkImage SkColorSpace is nullptr, dstInfo.colorSpace() must match. Returns
false if pixel conversion is not possible.

srcX and srcY may be negative to copy only top or left of source. Returns
false if width() or height() is zero or negative.
Returns false if abs(srcX) >= Image width(), or if abs(srcY) >= Image height().

If cachingHint is kAllow_CachingHint, pixels may be retained locally.
If cachingHint is kDisallow_CachingHint, pixels are not added to the local cache.

@param context      the GrDirectContext in play, if it exists
@param dstInfo      destination width, height, SkColorType, SkAlphaType, SkColorSpace
@param dstPixels    destination pixel storage
@param dstRowBytes  destination row length
@param srcX         column index whose absolute value is less than width()
@param srcY         row index whose absolute value is less than height()
@param cachingHint  whether the pixels should be cached locally
@return             true if pixels are copied to dstPixels
*/
func (o Image) ReadPixels(context GrDirectContext, dstInfo ImageInfo, dstPixels []byte, dstRowBytes uint32, srcX int32, srcY int32, cachingHint ImageCachingHint) bool {
	c_obj := o.sk
	c_context := context.sk
	c_dstInfo := dstInfo.sk
	c_dstPixels := unsafe.Pointer(&dstPixels[0])
	c_dstRowBytes := C.ulong(dstRowBytes)
	c_srcX := C.int(srcX)
	c_srcY := C.int(srcY)
	c_cachingHint := C.uint(cachingHint)
	retC := C.misk_Image_readPixels(c_obj, c_context, c_dstInfo, c_dstPixels, c_dstRowBytes, c_srcX, c_srcY, c_cachingHint)
	return bool(retC)
}

/*
Returns subset of this image.

Returns nullptr if any of the following are true:
- Subset is empty
- Subset is not contained inside the image's bounds
- Pixels in the source image could not be read or copied
- This image is texture-backed and the provided context is null or does not match
the source image's context.

If the source image was texture-backed, the resulting image will be texture-backed also.
Otherwise, the returned image will be raster-backed.

@param direct  the GrDirectContext of the source image (nullptr is ok if the source image
is not texture-backed).
@param subset  bounds of returned SkImage
@return        the subsetted image, or nullptr

example: https://fiddle.skia.org/c/@Image_makeSubset
*/
func (o Image) MakeSubset(direct GrDirectContext, subset IRect) Image {
	c_obj := o.sk
	c_direct := direct.sk
	c_subset := *(*C.sk_SkIRect)(unsafe.Pointer(&subset))
	retC := C.misk_Image_makeSubset(c_obj, c_direct, c_subset)
	return Image{sk: retC}
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
type ImageCachingHint uint32

const (
	// allows internally caching decoded and copied pixels
	ImageCachingHintAllow ImageCachingHint = 0
	// disallows internally caching decoded and copied pixels
	ImageCachingHintDisallow ImageCachingHint = 1
)

/*
Describes pixel dimensions and encoding. SkBitmap, SkImage, PixMap, and SkSurface
can be created from SkImageInfo. SkImageInfo can be retrieved from SkBitmap and
SkPixmap, but not from SkImage and SkSurface. For example, SkImage and SkSurface
implementations may defer pixel depth, so may not completely specify SkImageInfo.

SkImageInfo contains dimensions, the pixel integral width and height. It encodes
how pixel bits describe alpha, transparency; color components red, blue,
and green; and SkColorSpace, the range and linearity of colors.
*/
type ImageInfo struct {
	sk *C.sk_SkImageInfo
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the ImageInfo has not been created.
func (o ImageInfo) IsNil() bool {
	return o.sk == nil
}

/*
Creates an empty SkImageInfo with kUnknown_SkColorType, kUnknown_SkAlphaType,
a width and height of zero, and no SkColorSpace.

@return  empty SkImageInfo
*/
func NewImageInfo() ImageInfo {

	retC := C.misk_new_ImageInfo()
	return ImageInfo{sk: retC}
}

func (o *ImageInfo) Delete() {
	C.misk_delete_SkImageInfo(o.sk)
	o.sk = nil
}

/*
4x4 matrix used by SkCanvas and other parts of Skia.

Skia assumes a right-handed coordinate system:

	+X goes to the right
	+Y goes down
	+Z goes into the screen (away from the viewer)
*/
type M44 struct {
	sk *C.sk_SkM44
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the M44 has not been created.
func (o M44) IsNil() bool {
	return o.sk == nil
}

func NewM44Copy(src M44) M44 {
	c_src := src.sk
	retC := C.misk_new_M44Copy(c_src)
	return M44{sk: retC}
}

func NewM44() M44 {

	retC := C.misk_new_M44()
	return M44{sk: retC}
}

func NewM44AB(a M44, b M44) M44 {
	c_a := a.sk
	c_b := b.sk
	retC := C.misk_new_M44AB(c_a, c_b)
	return M44{sk: retC}
}

/*
The constructor parameters are in row-major order.
*/
func NewM44Scalars(m0 float32, m4 float32, m8 float32, m12 float32, m1 float32, m5 float32, m9 float32, m13 float32, m2 float32, m6 float32, m10 float32, m14 float32, m3 float32, m7 float32, m11 float32, m15 float32) M44 {
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
	retC := C.misk_new_M44Scalars(c_m0, c_m4, c_m8, c_m12, c_m1, c_m5, c_m9, c_m13, c_m2, c_m6, c_m10, c_m14, c_m3, c_m7, c_m11, c_m15)
	return M44{sk: retC}
}

func (o *M44) Delete() {
	C.misk_delete_SkM44(o.sk)
	o.sk = nil
}

type MemoryStream struct {
	sk *C.sk_SkMemoryStream
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the MemoryStream has not been created.
func (o MemoryStream) IsNil() bool {
	return o.sk == nil
}

// AsStream converts the MemoryStream to a Stream.
func (o MemoryStream) AsStream() Stream {
	return Stream{sk: (*C.sk_SkStream)(unsafe.Pointer(o.sk))}
}

func (o *MemoryStream) Delete() {
	C.misk_delete_SkMemoryStream(o.sk)
	o.sk = nil
}

/*
Returns a stream with a shared reference to the input data.
*/
func MemoryStreamMake(data Data) MemoryStream {
	c_data := data.sk
	retC := C.misk_MemoryStream_Make(c_data)
	return MemoryStream{sk: retC}
}

/*
Returns a stream with a bare pointer reference to the input data.
*/
func MemoryStreamMakeDirect(data []byte, length uint32) MemoryStream {
	c_data := unsafe.Pointer(&data[0])
	c_length := C.ulong(length)
	retC := C.misk_MemoryStream_MakeDirect(c_data, c_length)
	return MemoryStream{sk: retC}
}

/*
SkIPoint holds two 32-bit integer coordinates.
*/
type IPoint struct {
	sk *C.sk_SkIPoint
}

// !< x-axis value
func (o IPoint) X() int32 {
	return int32(o.sk.fX)
}

func (o *IPoint) SetX(value int32) {
	o.sk.fX = C.int(value)
}

// !< y-axis value
func (o IPoint) Y() int32 {
	return int32(o.sk.fY)
}

func (o *IPoint) SetY(value int32) {
	o.sk.fY = C.int(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the IPoint has not been created.
func (o IPoint) IsNil() bool {
	return o.sk == nil
}

/*
SkIRect holds four 32-bit integer coordinates describing the upper and
lower bounds of a rectangle. SkIRect may be created from outer bounds or
from position, width, and height. SkIRect describes an area; if its right
is less than or equal to its left, or if its bottom is less than or equal to
its top, it is considered empty.
*/
type IRect C.sk_SkIRect

// !< smaller x-axis bounds
func (o IRect) Left() int32 {
	return int32(o.fLeft)
}

func (o *IRect) SetLeft(value int32) {
	o.fLeft = C.int(value)
}

// !< smaller y-axis bounds
func (o IRect) Top() int32 {
	return int32(o.fTop)
}

func (o *IRect) SetTop(value int32) {
	o.fTop = C.int(value)
}

// !< larger x-axis bounds
func (o IRect) Right() int32 {
	return int32(o.fRight)
}

func (o *IRect) SetRight(value int32) {
	o.fRight = C.int(value)
}

// !< larger y-axis bounds
func (o IRect) Bottom() int32 {
	return int32(o.fBottom)
}

func (o *IRect) SetBottom(value int32) {
	o.fBottom = C.int(value)
}

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
func IRectMakeWH(w int32, h int32) IRect {
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
func IRectMakeLTRB(l int32, t int32, r int32, b int32) IRect {
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
func IRectMakeXYWH(x int32, y int32, w int32, h int32) IRect {
	c_x := C.int(x)
	c_y := C.int(y)
	c_w := C.int(w)
	c_h := C.int(h)
	retC := C.misk_IRect_MakeXYWH(c_x, c_y, c_w, c_h)
	return IRect(retC)
}

/*
Returns constructed SkIRect set to (0, 0, size.width(), size.height()).
Does not validate input; size.width() or size.height() may be negative.

@param size  values for SkIRect width and height
@return      bounds (0, 0, size.width(), size.height())
*/
func IRectMakeSize(size ISize) IRect {
	c_size := *(*C.sk_SkISize)(unsafe.Pointer(&size))
	retC := C.misk_IRect_MakeSize(c_size)
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

/*
Returns left edge of SkIRect, if sorted. Call isEmpty() to see if SkIRect may be invalid,
and sort() to reverse fLeft and fRight if needed.

@return  fLeft
*/
func (o *IRect) X() int32 {
	c_obj := (*C.sk_SkIRect)(o)
	retC := C.misk_IRect_x(c_obj)
	return int32(retC)
}

/*
Returns top edge of SkIRect, if sorted. Call isEmpty() to see if SkIRect may be invalid,
and sort() to reverse fTop and fBottom if needed.

@return  fTop
*/
func (o *IRect) Y() int32 {
	c_obj := (*C.sk_SkIRect)(o)
	retC := C.misk_IRect_y(c_obj)
	return int32(retC)
}

/*
Returns span on the x-axis. This does not check if SkIRect is sorted, or if
result fits in 32-bit signed integer; result may be negative.

@return  fRight minus fLeft
*/
func (o *IRect) Width() int32 {
	c_obj := (*C.sk_SkIRect)(o)
	retC := C.misk_IRect_width(c_obj)
	return int32(retC)
}

/*
Returns span on the y-axis. This does not check if SkIRect is sorted, or if
result fits in 32-bit signed integer; result may be negative.

@return  fBottom minus fTop
*/
func (o *IRect) Height() int32 {
	c_obj := (*C.sk_SkIRect)(o)
	retC := C.misk_IRect_height(c_obj)
	return int32(retC)
}

/*
Returns true if width() or height() are zero or negative.

@return  true if width() or height() are zero or negative
*/
func (o *IRect) IsEmpty() bool {
	c_obj := (*C.sk_SkIRect)(o)
	retC := C.misk_IRect_isEmpty(c_obj)
	return bool(retC)
}

/*
Sets SkIRect to (0, 0, 0, 0).

Many other rectangles are empty; if left is equal to or greater than right,
or if top is equal to or greater than bottom. Setting all members to zero
is a convenience, but does not designate a special empty rectangle.
*/
func (o *IRect) SetEmpty() {
	c_obj := (*C.sk_SkIRect)(o)
	C.misk_IRect_setEmpty(c_obj)
}

/*
Sets SkIRect to (left, top, right, bottom).
left and right are not sorted; left is not necessarily less than right.
top and bottom are not sorted; top is not necessarily less than bottom.

@param left    stored in fLeft
@param top     stored in fTop
@param right   stored in fRight
@param bottom  stored in fBottom
*/
func (o *IRect) SetLTRB(left int32, top int32, right int32, bottom int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_left := C.int(left)
	c_top := C.int(top)
	c_right := C.int(right)
	c_bottom := C.int(bottom)
	C.misk_IRect_setLTRB(c_obj, c_left, c_top, c_right, c_bottom)
}

/*
Sets SkIRect to: (x, y, x + width, y + height).
Does not validate input; width or height may be negative.

@param x       stored in fLeft
@param y       stored in fTop
@param width   added to x and stored in fRight
@param height  added to y and stored in fBottom
*/
func (o *IRect) SetXYWH(x int32, y int32, width int32, height int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_x := C.int(x)
	c_y := C.int(y)
	c_width := C.int(width)
	c_height := C.int(height)
	C.misk_IRect_setXYWH(c_obj, c_x, c_y, c_width, c_height)
}

func (o *IRect) SetWH(width int32, height int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_width := C.int(width)
	c_height := C.int(height)
	C.misk_IRect_setWH(c_obj, c_width, c_height)
}

/*
Offsets SkIRect by adding dx to fLeft, fRight; and by adding dy to fTop, fBottom.

If dx is negative, moves SkIRect returned to the left.
If dx is positive, moves SkIRect returned to the right.
If dy is negative, moves SkIRect returned upward.
If dy is positive, moves SkIRect returned downward.

@param dx  offset added to fLeft and fRight
@param dy  offset added to fTop and fBottom
*/
func (o *IRect) Offset(dx int32, dy int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_dx := C.int(dx)
	c_dy := C.int(dy)
	C.misk_IRect_offset(c_obj, c_dx, c_dy)
}

/*
Offsets SkIRect by adding delta.fX to fLeft, fRight; and by adding delta.fY to
fTop, fBottom.

If delta.fX is negative, moves SkIRect returned to the left.
If delta.fX is positive, moves SkIRect returned to the right.
If delta.fY is negative, moves SkIRect returned upward.
If delta.fY is positive, moves SkIRect returned downward.

@param delta  offset added to SkIRect
*/
func (o *IRect) OffsetPoint(delta IPoint) {
	c_obj := (*C.sk_SkIRect)(o)
	c_delta := delta.sk
	C.misk_IRect_offsetPoint(c_obj, c_delta)
}

/*
Offsets SkIRect so that fLeft equals newX, and fTop equals newY. width and height
are unchanged.

@param newX  stored in fLeft, preserving width()
@param newY  stored in fTop, preserving height()
*/
func (o *IRect) OffsetTo(newX int32, newY int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_newX := C.int(newX)
	c_newY := C.int(newY)
	C.misk_IRect_offsetTo(c_obj, c_newX, c_newY)
}

/*
Insets SkIRect by (dx,dy).

If dx is positive, makes SkIRect narrower.
If dx is negative, makes SkIRect wider.
If dy is positive, makes SkIRect shorter.
If dy is negative, makes SkIRect taller.

@param dx  offset added to fLeft and subtracted from fRight
@param dy  offset added to fTop and subtracted from fBottom
*/
func (o *IRect) Inset(dx int32, dy int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_dx := C.int(dx)
	c_dy := C.int(dy)
	C.misk_IRect_inset(c_obj, c_dx, c_dy)
}

/*
Outsets SkIRect by (dx, dy).

If dx is positive, makes SkIRect wider.
If dx is negative, makes SkIRect narrower.
If dy is positive, makes SkIRect taller.
If dy is negative, makes SkIRect shorter.

@param dx  subtracted to fLeft and added from fRight
@param dy  subtracted to fTop and added from fBottom
*/
func (o *IRect) Outset(dx int32, dy int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_dx := C.int(dx)
	c_dy := C.int(dy)
	C.misk_IRect_outset(c_obj, c_dx, c_dy)
}

/*
Adjusts SkIRect by adding dL to fLeft, dT to fTop, dR to fRight, and dB to fBottom.

If dL is positive, narrows SkIRect on the left. If negative, widens it on the left.
If dT is positive, shrinks SkIRect on the top. If negative, lengthens it on the top.
If dR is positive, narrows SkIRect on the right. If negative, widens it on the right.
If dB is positive, shrinks SkIRect on the bottom. If negative, lengthens it on the bottom.

The resulting SkIRect is not checked for validity. Thus, if the resulting SkIRect left is
greater than right, the SkIRect will be considered empty. Call sort() after this call
if that is not the desired behavior.

@param dL  offset added to fLeft
@param dT  offset added to fTop
@param dR  offset added to fRight
@param dB  offset added to fBottom
*/
func (o *IRect) Adjust(dL int32, dT int32, dR int32, dB int32) {
	c_obj := (*C.sk_SkIRect)(o)
	c_dL := C.int(dL)
	c_dT := C.int(dT)
	c_dR := C.int(dR)
	c_dB := C.int(dB)
	C.misk_IRect_adjust(c_obj, c_dL, c_dT, c_dR, c_dB)
}

/*
Returns true if: fLeft <= x < fRight && fTop <= y < fBottom.
Returns false if SkIRect is empty.

Considers input to describe constructed SkIRect: (x, y, x + 1, y + 1) and
returns true if constructed area is completely enclosed by SkIRect area.

@param x  test SkIPoint x-coordinate
@param y  test SkIPoint y-coordinate
@return   true if (x, y) is inside SkIRect
*/
func (o *IRect) Contains(x int32, y int32) bool {
	c_obj := (*C.sk_SkIRect)(o)
	c_x := C.int(x)
	c_y := C.int(y)
	retC := C.misk_IRect_contains(c_obj, c_x, c_y)
	return bool(retC)
}

/*
Returns true if SkIRect contains r.
Returns false if SkIRect is empty or r is empty.

SkIRect contains r when SkIRect area completely includes r area.

@param r  SkIRect contained
@return   true if all sides of SkIRect are outside r
*/
func (o *IRect) ContainsRect(r IRect) bool {
	c_obj := (*C.sk_SkIRect)(o)
	c_r := *(*C.sk_SkIRect)(unsafe.Pointer(&r))
	retC := C.misk_IRect_containsRect(c_obj, c_r)
	return bool(retC)
}

/*
Returns true if SkIRect contains construction.
Asserts if SkIRect is empty or construction is empty, and if SK_DEBUG is defined.

Return is undefined if SkIRect is empty or construction is empty.

@param r  SkIRect contained
@return   true if all sides of SkIRect are outside r
*/
func (o *IRect) ContainsNoEmptyCheck(r IRect) bool {
	c_obj := (*C.sk_SkIRect)(o)
	c_r := *(*C.sk_SkIRect)(unsafe.Pointer(&r))
	retC := C.misk_IRect_containsNoEmptyCheck(c_obj, c_r)
	return bool(retC)
}

/*
Returns true if SkIRect intersects r, and sets SkIRect to intersection.
Returns false if SkIRect does not intersect r, and leaves SkIRect unchanged.

Returns false if either r or SkIRect is empty, leaving SkIRect unchanged.

@param r  limit of result
@return   true if r and SkIRect have area in common
*/
func (o *IRect) Intersect(r IRect) bool {
	c_obj := (*C.sk_SkIRect)(o)
	c_r := *(*C.sk_SkIRect)(unsafe.Pointer(&r))
	retC := C.misk_IRect_intersect(c_obj, c_r)
	return bool(retC)
}

/*
Sets SkIRect to the union of itself and r.

Has no effect if r is empty. Otherwise, if SkIRect is empty, sets SkIRect to r.

@param r  expansion SkIRect

example: https://fiddle.skia.org/c/@IRect_join_2
*/
func (o *IRect) Join(r IRect) {
	c_obj := (*C.sk_SkIRect)(o)
	c_r := *(*C.sk_SkIRect)(unsafe.Pointer(&r))
	C.misk_IRect_join(c_obj, c_r)
}

/*
Swaps fLeft and fRight if fLeft is greater than fRight; and swaps
fTop and fBottom if fTop is greater than fBottom. Result may be empty,
and width() and height() will be zero or positive.
*/
func (o *IRect) Sort() {
	c_obj := (*C.sk_SkIRect)(o)
	C.misk_IRect_sort(c_obj)
}

type ISize C.sk_SkISize

func (o ISize) Width() int32 {
	return int32(o.fWidth)
}

func (o *ISize) SetWidth(value int32) {
	o.fWidth = C.int(value)
}

func (o ISize) Height() int32 {
	return int32(o.fHeight)
}

func (o *ISize) SetHeight(value int32) {
	o.fHeight = C.int(value)
}

/*
SkMatrix holds a 3x3 matrix for transforming coordinates. This allows mapping
SkPoint and vectors with translation, scaling, skewing, rotation, and
perspective.

SkMatrix elements are in row major order.
SkMatrix constexpr default constructs to identity.

SkMatrix includes a hidden variable that classifies the type of matrix to
improve performance. SkMatrix is not thread safe unless getType() is called first.

example: https://fiddle.skia.org/c/@Matrix_063
*/
type Matrix struct {
	sk *C.sk_SkMatrix
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Matrix has not been created.
func (o Matrix) IsNil() bool {
	return o.sk == nil
}

/*
Creates an identity SkMatrix:

| 1 0 0 |
| 0 1 0 |
| 0 0 1 |
*/
func NewMatrix() Matrix {

	retC := C.misk_new_Matrix()
	return Matrix{sk: retC}
}

func (o *Matrix) Delete() {
	C.misk_delete_SkMatrix(o.sk)
	o.sk = nil
}

/*
Perform a series of path operations, optimized for unioning many paths together.
*/
type OpBuilder struct {
	sk *C.sk_SkOpBuilder
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the OpBuilder has not been created.
func (o OpBuilder) IsNil() bool {
	return o.sk == nil
}

func NewOpBuilder() OpBuilder {

	retC := C.misk_new_OpBuilder()
	return OpBuilder{sk: retC}
}

func (o *OpBuilder) Delete() {
	C.misk_delete_SkOpBuilder(o.sk)
	o.sk = nil
}

/*
Add one or more paths and their operand. The builder is empty before the first
path is added, so the result of a single add is (emptyPath OP path).

@param path The second operand.
@param _operator The operator to apply to the existing and supplied paths.
*/
func (o OpBuilder) Add(path Path, _operator PathOp) {
	c_obj := o.sk
	c_path := path.sk
	c__operator := C.uint(_operator)
	C.misk_OpBuilder_add(c_obj, c_path, c__operator)
}

/*
Computes the sum of all paths and operands, and resets the builder to its
initial state.

@param result The product of the operands.
@return True if the operation succeeded.
*/
func (o OpBuilder) Resolve(result Path) bool {
	c_obj := o.sk
	c_result := result.sk
	retC := C.misk_OpBuilder_resolve(c_obj, c_result)
	return bool(retC)
}

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
func (o *Paint) Delete() {
	C.misk_delete_SkPaint(o.sk)
	o.sk = nil
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

func (o Paint) GetAlpha() byte {
	c_obj := o.sk
	retC := C.misk_Paint_getAlpha(c_obj)
	return byte(retC)
}

func (o Paint) SetAlpha(a uint32) {
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
func (o Paint) SetARGB(a uint32, r uint32, g uint32, b uint32) {
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
Retrieves alpha and RGB, unpremultiplied, packed into 32 bits.
Use helpers SkColorGetA(), SkColorGetR(), SkColorGetG(), and SkColorGetB() to extract
a color component.

@return  unpremultiplied ARGB
*/
func (o Paint) GetColor() Color {
	c_obj := o.sk
	retC := C.misk_Paint_getColor(c_obj)
	return Color(retC)
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
	c_join := C.uchar(join)
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
	c_style := C.uchar(style)
	C.misk_Paint_setStyle(c_obj, c_style)
}

/*
Cap draws at the beginning and end of an open path contour.
*/
type PaintCap uint32

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
type PaintJoin byte

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
type PaintStyle byte

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
func (o *Path) Delete() {
	C.misk_delete_SkPath(o.sk)
	o.sk = nil
}

/*
Returns true if SkPath contain equal verbs and equal weights.
If SkPath contain one or more conics, the weights must match.

conicTo() may add different verbs depending on conic weight, so it is not
trivial to interpolate a pair of SkPath containing conics with different
conic weight values.

@param compare  SkPath to compare
@return         true if SkPath verb array and weights are equivalent

example: https://fiddle.skia.org/c/@Path_isInterpolatable
*/
func (o Path) IsInterpolatable(compare Path) bool {
	c_obj := o.sk
	c_compare := compare.sk
	retC := C.misk_Path_isInterpolatable(c_obj, c_compare)
	return bool(retC)
}

/*
Interpolates between SkPath with SkPoint array of equal size.
Copy verb array and weights to out, and set out SkPoint array to a weighted
average of this SkPoint array and ending SkPoint array, using the formula:
(Path Point * weight) + ending Point * (1 - weight).

weight is most useful when between zero (ending SkPoint array) and
one (this Point_Array); will work with values outside of this
range.

interpolate() returns false and leaves out unchanged if SkPoint array is not
the same size as ending SkPoint array. Call isInterpolatable() to check SkPath
compatibility prior to calling interpolate().

@param ending  SkPoint array averaged with this SkPoint array
@param weight  contribution of this SkPoint array, and
one minus contribution of ending SkPoint array
@param out     SkPath replaced by interpolated averages
@return        true if SkPath contain same number of SkPoint

example: https://fiddle.skia.org/c/@Path_interpolate
*/
func (o Path) Interpolate(ending Path, weight float32, out Path) bool {
	c_obj := o.sk
	c_ending := ending.sk
	c_weight := C.float(weight)
	c_out := out.sk
	retC := C.misk_Path_interpolate(c_obj, c_ending, c_weight, c_out)
	return bool(retC)
}

/*
Returns SkPathFillType, the rule used to fill SkPath.

@return  current SkPathFillType setting
*/
func (o Path) GetFillType() PathFillType {
	c_obj := o.sk
	retC := C.misk_Path_getFillType(c_obj)
	return PathFillType(retC)
}

/*
Sets FillType, the rule used to fill SkPath. While there is no check
that ft is legal, values outside of FillType are not supported.
*/
func (o Path) SetFillType(ft PathFillType) {
	c_obj := o.sk
	c_ft := C.int(ft)
	C.misk_Path_setFillType(c_obj, c_ft)
}

/*
Returns if FillType describes area outside SkPath geometry. The inverse fill area
extends indefinitely.

@return  true if FillType is kInverseWinding or kInverseEvenOdd
*/
func (o Path) IsInverseFillType() bool {
	c_obj := o.sk
	retC := C.misk_Path_isInverseFillType(c_obj)
	return bool(retC)
}

/*
Replaces FillType with its inverse. The inverse of FillType describes the area
unmodified by the original FillType.
*/
func (o Path) ToggleInverseFillType() {
	c_obj := o.sk
	C.misk_Path_toggleInverseFillType(c_obj)
}

/*
Returns true if the path is convex. If necessary, it will first compute the convexity.
*/
func (o Path) IsConvex() bool {
	c_obj := o.sk
	retC := C.misk_Path_isConvex(c_obj)
	return bool(retC)
}

/*
Returns true if this path is recognized as an oval or circle.

bounds receives bounds of oval.

bounds is unmodified if oval is not found.

@param bounds  storage for bounding SkRect of oval; may be nullptr
@return        true if SkPath is recognized as an oval or circle

example: https://fiddle.skia.org/c/@Path_isOval
*/
func (o Path) IsOval(bounds Rect) bool {
	c_obj := o.sk
	c_bounds := *(*C.sk_SkRect)(unsafe.Pointer(&bounds))
	retC := C.misk_Path_isOval(c_obj, c_bounds)
	return bool(retC)
}

/*
Returns true if path is representable as SkRRect.
Returns false if path is representable as oval, circle, or SkRect.

rrect receives bounds of SkRRect.

rrect is unmodified if SkRRect is not found.

@param rrect  storage for bounding SkRect of SkRRect; may be nullptr
@return       true if SkPath contains only SkRRect

example: https://fiddle.skia.org/c/@Path_isRRect
*/
func (o Path) IsRRect(rrect RRect) bool {
	c_obj := o.sk
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	retC := C.misk_Path_isRRect(c_obj, c_rrect)
	return bool(retC)
}

/*
Sets SkPath to its initial state.
Removes verb array, SkPoint array, and weights, and sets FillType to kWinding.
Internal storage associated with SkPath is released.

@return  reference to SkPath

example: https://fiddle.skia.org/c/@Path_reset
*/
func (o Path) Reset() Path {
	c_obj := o.sk
	retC := C.misk_Path_reset(c_obj)
	return Path{sk: &retC}
}

/*
Sets SkPath to its initial state, preserving internal storage.
Removes verb array, SkPoint array, and weights, and sets FillType to kWinding.
Internal storage associated with SkPath is retained.

Use rewind() instead of reset() if SkPath storage will be reused and performance
is critical.

@return  reference to SkPath

example: https://fiddle.skia.org/c/@Path_rewind
*/
func (o Path) Rewind() Path {
	c_obj := o.sk
	retC := C.misk_Path_rewind(c_obj)
	return Path{sk: &retC}
}

/*
Returns if SkPath is empty.
Empty SkPath may have FillType but has no SkPoint, SkPath::Verb, or conic weight.
SkPath() constructs empty SkPath; reset() and rewind() make SkPath empty.

@return  true if the path contains no SkPath::Verb array
*/
func (o Path) IsEmpty() bool {
	c_obj := o.sk
	retC := C.misk_Path_isEmpty(c_obj)
	return bool(retC)
}

/*
Returns if contour is closed.
Contour is closed if SkPath SkPath::Verb array was last modified by close(). When stroked,
closed contour draws SkPaint::Join instead of SkPaint::Cap at first and last SkPoint.

@return  true if the last contour ends with a kClose_Verb

example: https://fiddle.skia.org/c/@Path_isLastContourClosed
*/
func (o Path) IsLastContourClosed() bool {
	c_obj := o.sk
	retC := C.misk_Path_isLastContourClosed(c_obj)
	return bool(retC)
}

/*
Returns true for finite SkPoint array values between negative SK_ScalarMax and
positive SK_ScalarMax. Returns false for any SkPoint array value of
SK_ScalarInfinity, SK_ScalarNegativeInfinity, or SK_ScalarNaN.

@return  true if all SkPoint values are finite
*/
func (o Path) IsFinite() bool {
	c_obj := o.sk
	retC := C.misk_Path_isFinite(c_obj)
	return bool(retC)
}

/*
Returns true if the path is volatile; it will not be altered or discarded
by the caller after it is drawn. SkPath by default have volatile set false, allowing
SkSurface to attach a cache of data which speeds repeated drawing. If true, SkSurface
may not speed repeated drawing.

@return  true if caller will alter SkPath after drawing
*/
func (o Path) IsVolatile() bool {
	c_obj := o.sk
	retC := C.misk_Path_isVolatile(c_obj)
	return bool(retC)
}

/*
Specifies whether SkPath is volatile; whether it will be altered or discarded
by the caller after it is drawn. SkPath by default have volatile set false, allowing
Skia to attach a cache of data which speeds repeated drawing.

Mark temporary paths, discarded or modified after use, as volatile
to inform Skia that the path need not be cached.

Mark animating SkPath volatile to improve performance.
Mark unchanging SkPath non-volatile to improve repeated rendering.

raster surface SkPath draws are affected by volatile for some shadows.
GPU surface SkPath draws are affected by volatile for some shadows and concave geometries.

@param isVolatile  true if caller will alter SkPath after drawing
@return            reference to SkPath
*/
func (o Path) SetIsVolatile(isVolatile bool) Path {
	c_obj := o.sk
	c_isVolatile := C.bool(isVolatile)
	retC := C.misk_Path_setIsVolatile(c_obj, c_isVolatile)
	return Path{sk: &retC}
}

/*
Returns the number of points in SkPath.
SkPoint count is initially zero.

@return  SkPath SkPoint array length

example: https://fiddle.skia.org/c/@Path_countPoints
*/
func (o Path) CountPoints() int32 {
	c_obj := o.sk
	retC := C.misk_Path_countPoints(c_obj)
	return int32(retC)
}

/*
Returns SkPoint at index in SkPoint array. Valid range for index is
0 to countPoints() - 1.
Returns (0, 0) if index is out of range.

@param index  SkPoint array element selector
@return       SkPoint array value or (0, 0)

example: https://fiddle.skia.org/c/@Path_getPoint
*/
func (o Path) GetPoint(index int32) Point {
	c_obj := o.sk
	c_index := C.int(index)
	retC := C.misk_Path_getPoint(c_obj, c_index)
	return Point{sk: &retC}
}

/*
Returns number of points in SkPath. Up to max points are copied.
points may be nullptr; then, max must be zero.
If max is greater than number of points, excess points storage is unaltered.

@param points  storage for SkPath SkPoint array. May be nullptr
@param max     maximum to copy; must be greater than or equal to zero
@return        SkPath SkPoint array length

example: https://fiddle.skia.org/c/@Path_getPoints
*/
func (o Path) GetPoints(points []Point, max int32) int32 {
	c_obj := o.sk
	c_points := (*C.sk_SkPoint)(unsafe.Pointer(&points[0]))
	c_max := C.int(max)
	retC := C.misk_Path_getPoints(c_obj, c_points, c_max)
	return int32(retC)
}

/*
Returns the number of verbs: kMove_Verb, kLine_Verb, kQuad_Verb, kConic_Verb,
kCubic_Verb, and kClose_Verb; added to SkPath.

@return  length of verb array

example: https://fiddle.skia.org/c/@Path_countVerbs
*/
func (o Path) CountVerbs() int32 {
	c_obj := o.sk
	retC := C.misk_Path_countVerbs(c_obj)
	return int32(retC)
}

/*
Returns the number of verbs in the path. Up to max verbs are copied. The
verbs are copied as one byte per verb.

@param verbs  storage for verbs, may be nullptr
@param max    maximum number to copy into verbs
@return       the actual number of verbs in the path

example: https://fiddle.skia.org/c/@Path_getVerbs
*/
func (o Path) GetVerbs(verbs string, max int32) int32 {
	c_obj := o.sk
	c_verbs := (*C.uchar)(unsafe.Pointer(C.CString(verbs)))
	defer C.free(unsafe.Pointer(c_verbs))
	c_max := C.int(max)
	retC := C.misk_Path_getVerbs(c_obj, c_verbs, c_max)
	return int32(retC)
}

/*
Returns the approximate byte size of the SkPath in memory.

@return  approximate size
*/
func (o Path) ApproximateBytesUsed() uint32 {
	c_obj := o.sk
	retC := C.misk_Path_approximateBytesUsed(c_obj)
	return uint32(retC)
}

/*
Exchanges the verb array, SkPoint array, weights, and SkPath::FillType with other.
Cached state is also exchanged. swap() internally exchanges pointers, so
it is lightweight and does not allocate memory.

swap() usage has largely been replaced by operator=(const SkPath& path).
SkPath do not copy their content on assignment until they are written to,
making assignment as efficient as swap().

@param other  SkPath exchanged by value

example: https://fiddle.skia.org/c/@Path_swap
*/
func (o Path) Swap(other Path) {
	c_obj := o.sk
	c_other := other.sk
	C.misk_Path_swap(c_obj, c_other)
}

/*
Returns minimum and maximum axes values of SkPoint array.
Returns (0, 0, 0, 0) if SkPath contains no points. Returned bounds width and height may
be larger or smaller than area affected when SkPath is drawn.

SkRect returned includes all SkPoint added to SkPath, including SkPoint associated with
kMove_Verb that define empty contours.

@return  bounds of all SkPoint in SkPoint array
*/
func (o Path) GetBounds() Rect {
	c_obj := o.sk
	retC := C.misk_Path_getBounds(c_obj)
	return Rect(retC)
}

/*
Updates internal bounds so that subsequent calls to getBounds() are instantaneous.
Unaltered copies of SkPath may also access cached bounds through getBounds().

For now, identical to calling getBounds() and ignoring the returned value.

Call to prepare SkPath subsequently drawn from multiple threads,
to avoid a race condition where each draw separately computes the bounds.
*/
func (o Path) UpdateBoundsCache() {
	c_obj := o.sk
	C.misk_Path_updateBoundsCache(c_obj)
}

/*
Returns minimum and maximum axes values of the lines and curves in SkPath.
Returns (0, 0, 0, 0) if SkPath contains no points.
Returned bounds width and height may be larger or smaller than area affected
when SkPath is drawn.

Includes SkPoint associated with kMove_Verb that define empty
contours.

Behaves identically to getBounds() when SkPath contains
only lines. If SkPath contains curves, computed bounds includes
the maximum extent of the quad, conic, or cubic; is slower than getBounds();
and unlike getBounds(), does not cache the result.

@return  tight bounds of curves in SkPath

example: https://fiddle.skia.org/c/@Path_computeTightBounds
*/
func (o Path) ComputeTightBounds() Rect {
	c_obj := o.sk
	retC := C.misk_Path_computeTightBounds(c_obj)
	return Rect(retC)
}

/*
Returns true if rect is contained by SkPath.
May return false when rect is contained by SkPath.

For now, only returns true if SkPath has one contour and is convex.
rect may share points and edges with SkPath and be contained.
Returns true if rect is empty, that is, it has zero width or height; and
the SkPoint or line described by rect is contained by SkPath.

@param rect  SkRect, line, or SkPoint checked for containment
@return      true if rect is contained

example: https://fiddle.skia.org/c/@Path_conservativelyContainsRect
*/
func (o Path) ConservativelyContainsRect(rect Rect) bool {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	retC := C.misk_Path_conservativelyContainsRect(c_obj, c_rect)
	return bool(retC)
}

/*
Grows SkPath verb array, SkPoint array, and conics to contain additional space.
May improve performance and use less memory by
reducing the number and size of allocations when creating SkPath.

@param extraPtCount  number of additional SkPoint to allocate
@param extraVerbCount  number of additional verbs
@param extraConicCount  number of additional conics

example: https://fiddle.skia.org/c/@Path_incReserve
*/
func (o Path) IncReserve(extraPtCount int32, extraVerbCount int32, extraConicCount int32) {
	c_obj := o.sk
	c_extraPtCount := C.int(extraPtCount)
	c_extraVerbCount := C.int(extraVerbCount)
	c_extraConicCount := C.int(extraConicCount)
	C.misk_Path_incReserve(c_obj, c_extraPtCount, c_extraVerbCount, c_extraConicCount)
}

/*
Adds beginning of contour at SkPoint (x, y).

@param x  x-axis value of contour start
@param y  y-axis value of contour start
@return   reference to SkPath

example: https://fiddle.skia.org/c/@Path_moveTo
*/
func (o Path) MoveToPoint(x float32, y float32) Path {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	retC := C.misk_Path_moveToPoint(c_obj, c_x, c_y)
	return Path{sk: &retC}
}

/*
Adds beginning of contour at SkPoint p.

@param p  contour start
@return   reference to SkPath
*/
func (o Path) MoveTo(p Point) Path {
	c_obj := o.sk
	c_p := p.sk
	retC := C.misk_Path_moveTo(c_obj, c_p)
	return Path{sk: &retC}
}

/*
Adds beginning of contour relative to last point.
If SkPath is empty, starts contour at (dx, dy).
Otherwise, start contour at last point offset by (dx, dy).
Function name stands for "relative move to".

@param dx  offset from last point to contour start on x-axis
@param dy  offset from last point to contour start on y-axis
@return    reference to SkPath

example: https://fiddle.skia.org/c/@Path_rMoveTo
*/
func (o Path) RMoveTo(dx float32, dy float32) Path {
	c_obj := o.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	retC := C.misk_Path_rMoveTo(c_obj, c_dx, c_dy)
	return Path{sk: &retC}
}

/*
Adds line from last point to (x, y). If SkPath is empty, or last SkPath::Verb is
kClose_Verb, last point is set to (0, 0) before adding line.

lineTo() appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed.
lineTo() then appends kLine_Verb to verb array and (x, y) to SkPoint array.

@param x  end of added line on x-axis
@param y  end of added line on y-axis
@return   reference to SkPath

example: https://fiddle.skia.org/c/@Path_lineTo
*/
func (o Path) LineTo(x float32, y float32) Path {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	retC := C.misk_Path_lineTo(c_obj, c_x, c_y)
	return Path{sk: &retC}
}

/*
Adds line from last point to SkPoint p. If SkPath is empty, or last SkPath::Verb is
kClose_Verb, last point is set to (0, 0) before adding line.

lineTo() first appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed.
lineTo() then appends kLine_Verb to verb array and SkPoint p to SkPoint array.

@param p  end SkPoint of added line
@return   reference to SkPath
*/
func (o Path) LineToPoint(p Point) Path {
	c_obj := o.sk
	c_p := p.sk
	retC := C.misk_Path_lineToPoint(c_obj, c_p)
	return Path{sk: &retC}
}

/*
Adds line from last point to vector (dx, dy). If SkPath is empty, or last SkPath::Verb is
kClose_Verb, last point is set to (0, 0) before adding line.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed;
then appends kLine_Verb to verb array and line end to SkPoint array.
Line end is last point plus vector (dx, dy).
Function name stands for "relative line to".

@param dx  offset from last point to line end on x-axis
@param dy  offset from last point to line end on y-axis
@return    reference to SkPath

example: https://fiddle.skia.org/c/@Path_rLineTo
example: https://fiddle.skia.org/c/@Quad_a
example: https://fiddle.skia.org/c/@Quad_b
*/
func (o Path) RLineTo(dx float32, dy float32) Path {
	c_obj := o.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	retC := C.misk_Path_rLineTo(c_obj, c_dx, c_dy)
	return Path{sk: &retC}
}

/*
Adds quad from last point towards (x1, y1), to (x2, y2).
If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to (0, 0)
before adding quad.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed;
then appends kQuad_Verb to verb array; and (x1, y1), (x2, y2)
to SkPoint array.

@param x1  control SkPoint of quad on x-axis
@param y1  control SkPoint of quad on y-axis
@param x2  end SkPoint of quad on x-axis
@param y2  end SkPoint of quad on y-axis
@return    reference to SkPath

example: https://fiddle.skia.org/c/@Path_quadTo
*/
func (o Path) QuadTo(x1 float32, y1 float32, x2 float32, y2 float32) Path {
	c_obj := o.sk
	c_x1 := C.float(x1)
	c_y1 := C.float(y1)
	c_x2 := C.float(x2)
	c_y2 := C.float(y2)
	retC := C.misk_Path_quadTo(c_obj, c_x1, c_y1, c_x2, c_y2)
	return Path{sk: &retC}
}

/*
Adds quad from last point towards SkPoint p1, to SkPoint p2.
If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to (0, 0)
before adding quad.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed;
then appends kQuad_Verb to verb array; and SkPoint p1, p2
to SkPoint array.

@param p1  control SkPoint of added quad
@param p2  end SkPoint of added quad
@return    reference to SkPath
*/
func (o Path) QuadToPoint(p1 Point, p2 Point) Path {
	c_obj := o.sk
	c_p1 := p1.sk
	c_p2 := p2.sk
	retC := C.misk_Path_quadToPoint(c_obj, c_p1, c_p2)
	return Path{sk: &retC}
}

/*
Adds quad from last point towards vector (dx1, dy1), to vector (dx2, dy2).
If SkPath is empty, or last SkPath::Verb
is kClose_Verb, last point is set to (0, 0) before adding quad.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array,
if needed; then appends kQuad_Verb to verb array; and appends quad
control and quad end to SkPoint array.
Quad control is last point plus vector (dx1, dy1).
Quad end is last point plus vector (dx2, dy2).
Function name stands for "relative quad to".

@param dx1  offset from last point to quad control on x-axis
@param dy1  offset from last point to quad control on y-axis
@param dx2  offset from last point to quad end on x-axis
@param dy2  offset from last point to quad end on y-axis
@return     reference to SkPath

example: https://fiddle.skia.org/c/@Conic_Weight_a
example: https://fiddle.skia.org/c/@Conic_Weight_b
example: https://fiddle.skia.org/c/@Conic_Weight_c
example: https://fiddle.skia.org/c/@Path_rQuadTo
*/
func (o Path) RQuadTo(dx1 float32, dy1 float32, dx2 float32, dy2 float32) Path {
	c_obj := o.sk
	c_dx1 := C.float(dx1)
	c_dy1 := C.float(dy1)
	c_dx2 := C.float(dx2)
	c_dy2 := C.float(dy2)
	retC := C.misk_Path_rQuadTo(c_obj, c_dx1, c_dy1, c_dx2, c_dy2)
	return Path{sk: &retC}
}

/*
Adds conic from last point towards (x1, y1), to (x2, y2), weighted by w.
If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to (0, 0)
before adding conic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed.

If w is finite and not one, appends kConic_Verb to verb array;
and (x1, y1), (x2, y2) to SkPoint array; and w to conic weights.

If w is one, appends kQuad_Verb to verb array, and
(x1, y1), (x2, y2) to SkPoint array.

If w is not finite, appends kLine_Verb twice to verb array, and
(x1, y1), (x2, y2) to SkPoint array.

@param x1  control SkPoint of conic on x-axis
@param y1  control SkPoint of conic on y-axis
@param x2  end SkPoint of conic on x-axis
@param y2  end SkPoint of conic on y-axis
@param w   weight of added conic
@return    reference to SkPath
*/
func (o Path) ConicTo(x1 float32, y1 float32, x2 float32, y2 float32, w float32) Path {
	c_obj := o.sk
	c_x1 := C.float(x1)
	c_y1 := C.float(y1)
	c_x2 := C.float(x2)
	c_y2 := C.float(y2)
	c_w := C.float(w)
	retC := C.misk_Path_conicTo(c_obj, c_x1, c_y1, c_x2, c_y2, c_w)
	return Path{sk: &retC}
}

/*
Adds conic from last point towards SkPoint p1, to SkPoint p2, weighted by w.
If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to (0, 0)
before adding conic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed.

If w is finite and not one, appends kConic_Verb to verb array;
and SkPoint p1, p2 to SkPoint array; and w to conic weights.

If w is one, appends kQuad_Verb to verb array, and SkPoint p1, p2
to SkPoint array.

If w is not finite, appends kLine_Verb twice to verb array, and
SkPoint p1, p2 to SkPoint array.

@param p1  control SkPoint of added conic
@param p2  end SkPoint of added conic
@param w   weight of added conic
@return    reference to SkPath
*/
func (o Path) ConicToPoints(p1 Point, p2 Point, w float32) Path {
	c_obj := o.sk
	c_p1 := p1.sk
	c_p2 := p2.sk
	c_w := C.float(w)
	retC := C.misk_Path_conicToPoints(c_obj, c_p1, c_p2, c_w)
	return Path{sk: &retC}
}

/*
Adds conic from last point towards vector (dx1, dy1), to vector (dx2, dy2),
weighted by w. If SkPath is empty, or last SkPath::Verb
is kClose_Verb, last point is set to (0, 0) before adding conic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array,
if needed.

If w is finite and not one, next appends kConic_Verb to verb array,
and w is recorded as conic weight; otherwise, if w is one, appends
kQuad_Verb to verb array; or if w is not finite, appends kLine_Verb
twice to verb array.

In all cases appends SkPoint control and end to SkPoint array.
control is last point plus vector (dx1, dy1).
end is last point plus vector (dx2, dy2).

Function name stands for "relative conic to".

@param dx1  offset from last point to conic control on x-axis
@param dy1  offset from last point to conic control on y-axis
@param dx2  offset from last point to conic end on x-axis
@param dy2  offset from last point to conic end on y-axis
@param w    weight of added conic
@return     reference to SkPath
*/
func (o Path) RConicTo(dx1 float32, dy1 float32, dx2 float32, dy2 float32, w float32) Path {
	c_obj := o.sk
	c_dx1 := C.float(dx1)
	c_dy1 := C.float(dy1)
	c_dx2 := C.float(dx2)
	c_dy2 := C.float(dy2)
	c_w := C.float(w)
	retC := C.misk_Path_rConicTo(c_obj, c_dx1, c_dy1, c_dx2, c_dy2, c_w)
	return Path{sk: &retC}
}

/*
Adds cubic from last point towards (x1, y1), then towards (x2, y2), ending at
(x3, y3). If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to
(0, 0) before adding cubic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed;
then appends kCubic_Verb to verb array; and (x1, y1), (x2, y2), (x3, y3)
to SkPoint array.

@param x1  first control SkPoint of cubic on x-axis
@param y1  first control SkPoint of cubic on y-axis
@param x2  second control SkPoint of cubic on x-axis
@param y2  second control SkPoint of cubic on y-axis
@param x3  end SkPoint of cubic on x-axis
@param y3  end SkPoint of cubic on y-axis
@return    reference to SkPath
*/
func (o Path) CubicTo(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) Path {
	c_obj := o.sk
	c_x1 := C.float(x1)
	c_y1 := C.float(y1)
	c_x2 := C.float(x2)
	c_y2 := C.float(y2)
	c_x3 := C.float(x3)
	c_y3 := C.float(y3)
	retC := C.misk_Path_cubicTo(c_obj, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3)
	return Path{sk: &retC}
}

/*
Adds cubic from last point towards SkPoint p1, then towards SkPoint p2, ending at
SkPoint p3. If SkPath is empty, or last SkPath::Verb is kClose_Verb, last point is set to
(0, 0) before adding cubic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array, if needed;
then appends kCubic_Verb to verb array; and SkPoint p1, p2, p3
to SkPoint array.

@param p1  first control SkPoint of cubic
@param p2  second control SkPoint of cubic
@param p3  end SkPoint of cubic
@return    reference to SkPath
*/
func (o Path) CubicToPoints(p1 Point, p2 Point, p3 Point) Path {
	c_obj := o.sk
	c_p1 := p1.sk
	c_p2 := p2.sk
	c_p3 := p3.sk
	retC := C.misk_Path_cubicToPoints(c_obj, c_p1, c_p2, c_p3)
	return Path{sk: &retC}
}

/*
Adds cubic from last point towards vector (dx1, dy1), then towards
vector (dx2, dy2), to vector (dx3, dy3).
If SkPath is empty, or last SkPath::Verb
is kClose_Verb, last point is set to (0, 0) before adding cubic.

Appends kMove_Verb to verb array and (0, 0) to SkPoint array,
if needed; then appends kCubic_Verb to verb array; and appends cubic
control and cubic end to SkPoint array.
Cubic control is last point plus vector (dx1, dy1).
Cubic end is last point plus vector (dx2, dy2).
Function name stands for "relative cubic to".

@param dx1  offset from last point to first cubic control on x-axis
@param dy1  offset from last point to first cubic control on y-axis
@param dx2  offset from last point to second cubic control on x-axis
@param dy2  offset from last point to second cubic control on y-axis
@param dx3  offset from last point to cubic end on x-axis
@param dy3  offset from last point to cubic end on y-axis
@return    reference to SkPath
*/
func (o Path) RCubicTo(dx1 float32, dy1 float32, dx2 float32, dy2 float32, dx3 float32, dy3 float32) Path {
	c_obj := o.sk
	c_dx1 := C.float(dx1)
	c_dy1 := C.float(dy1)
	c_dx2 := C.float(dx2)
	c_dy2 := C.float(dy2)
	c_dx3 := C.float(dx3)
	c_dy3 := C.float(dy3)
	retC := C.misk_Path_rCubicTo(c_obj, c_dx1, c_dy1, c_dx2, c_dy2, c_dx3, c_dy3)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath. Arc added is part of ellipse
bounded by oval, from startAngle through sweepAngle. Both startAngle and
sweepAngle are measured in degrees, where zero degrees is aligned with the
positive x-axis, and positive sweeps extends arc clockwise.

arcTo() adds line connecting SkPath last SkPoint to initial arc SkPoint if forceMoveTo
is false and SkPath is not empty. Otherwise, added contour begins with first point
of arc. Angles greater than -360 and less than 360 are treated modulo 360.

@param oval         bounds of ellipse containing arc
@param startAngle   starting angle of arc in degrees
@param sweepAngle   sweep, in degrees. Positive is clockwise; treated modulo 360
@param forceMoveTo  true to start a new contour with arc
@return             reference to SkPath

example: https://fiddle.skia.org/c/@Path_arcTo
*/
func (o Path) ArcTo1(oval Rect, startAngle float32, sweepAngle float32, forceMoveTo bool) Path {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_startAngle := C.float(startAngle)
	c_sweepAngle := C.float(sweepAngle)
	c_forceMoveTo := C.bool(forceMoveTo)
	retC := C.misk_Path_arcTo1(c_obj, c_oval, c_startAngle, c_sweepAngle, c_forceMoveTo)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath, after appending line if needed. Arc is implemented by conic
weighted to describe part of circle. Arc is contained by tangent from
last SkPath point to (x1, y1), and tangent from (x1, y1) to (x2, y2). Arc
is part of circle sized to radius, positioned so it touches both tangent lines.

If last Path Point does not start Arc, arcTo appends connecting Line to Path.
The length of Vector from (x1, y1) to (x2, y2) does not affect Arc.

Arc sweep is always less than 180 degrees. If radius is zero, or if
tangents are nearly parallel, arcTo appends Line from last Path Point to (x1, y1).

arcTo appends at most one Line and one conic.
arcTo implements the functionality of PostScript arct and HTML Canvas arcTo.

@param x1      x-axis value common to pair of tangents
@param y1      y-axis value common to pair of tangents
@param x2      x-axis value end of second tangent
@param y2      y-axis value end of second tangent
@param radius  distance from arc to circle center
@return        reference to SkPath

example: https://fiddle.skia.org/c/@Path_arcTo_2_a
example: https://fiddle.skia.org/c/@Path_arcTo_2_b
example: https://fiddle.skia.org/c/@Path_arcTo_2_c
*/
func (o Path) ArcTo2(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) Path {
	c_obj := o.sk
	c_x1 := C.float(x1)
	c_y1 := C.float(y1)
	c_x2 := C.float(x2)
	c_y2 := C.float(y2)
	c_radius := C.float(radius)
	retC := C.misk_Path_arcTo2(c_obj, c_x1, c_y1, c_x2, c_y2, c_radius)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath, after appending line if needed. Arc is implemented by conic
weighted to describe part of circle. Arc is contained by tangent from
last SkPath point to p1, and tangent from p1 to p2. Arc
is part of circle sized to radius, positioned so it touches both tangent lines.

If last SkPath SkPoint does not start arc, arcTo() appends connecting line to SkPath.
The length of vector from p1 to p2 does not affect arc.

Arc sweep is always less than 180 degrees. If radius is zero, or if
tangents are nearly parallel, arcTo() appends line from last SkPath SkPoint to p1.

arcTo() appends at most one line and one conic.
arcTo() implements the functionality of PostScript arct and HTML Canvas arcTo.

@param p1      SkPoint common to pair of tangents
@param p2      end of second tangent
@param radius  distance from arc to circle center
@return        reference to SkPath
*/
func (o Path) ArcTo3(p1 Point, p2 Point, radius float32) Path {
	c_obj := o.sk
	c_p1 := *(*C.sk_SkPoint)(unsafe.Pointer(&p1))
	c_p2 := *(*C.sk_SkPoint)(unsafe.Pointer(&p2))
	c_radius := C.float(radius)
	retC := C.misk_Path_arcTo3(c_obj, c_p1, c_p2, c_radius)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath. Arc is implemented by one or more conics weighted to
describe part of oval with radii (rx, ry) rotated by xAxisRotate degrees. Arc
curves from last SkPath SkPoint to (x, y), choosing one of four possible routes:
clockwise or counterclockwise, and smaller or larger.

Arc sweep is always less than 360 degrees. arcTo() appends line to (x, y) if
either radii are zero, or if last SkPath SkPoint equals (x, y). arcTo() scales radii
(rx, ry) to fit last SkPath SkPoint and (x, y) if both are greater than zero but
too small.

arcTo() appends up to four conic curves.
arcTo() implements the functionality of SVG arc, although SVG sweep-flag value
is opposite the integer value of sweep; SVG sweep-flag uses 1 for clockwise,
while kCW_Direction cast to int is zero.

@param rx           radius on x-axis before x-axis rotation
@param ry           radius on y-axis before x-axis rotation
@param xAxisRotate  x-axis rotation in degrees; positive values are clockwise
@param largeArc     chooses smaller or larger arc
@param sweep        chooses clockwise or counterclockwise arc
@param x            end of arc
@param y            end of arc
@return             reference to SkPath
*/
func (o Path) ArcTo4(rx float32, ry float32, xAxisRotate float32, largeArc PathArcSize, sweep PathDirection, x float32, y float32) Path {
	c_obj := o.sk
	c_rx := C.float(rx)
	c_ry := C.float(ry)
	c_xAxisRotate := C.float(xAxisRotate)
	c_largeArc := C.uint(largeArc)
	c_sweep := C.int(sweep)
	c_x := C.float(x)
	c_y := C.float(y)
	retC := C.misk_Path_arcTo4(c_obj, c_rx, c_ry, c_xAxisRotate, c_largeArc, c_sweep, c_x, c_y)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath. Arc is implemented by one or more conic weighted to describe
part of oval with radii (r.fX, r.fY) rotated by xAxisRotate degrees. Arc curves
from last SkPath SkPoint to (xy.fX, xy.fY), choosing one of four possible routes:
clockwise or counterclockwise,
and smaller or larger.

Arc sweep is always less than 360 degrees. arcTo() appends line to xy if either
radii are zero, or if last SkPath SkPoint equals (xy.fX, xy.fY). arcTo() scales radii r to
fit last SkPath SkPoint and xy if both are greater than zero but too small to describe
an arc.

arcTo() appends up to four conic curves.
arcTo() implements the functionality of SVG arc, although SVG sweep-flag value is
opposite the integer value of sweep; SVG sweep-flag uses 1 for clockwise, while
kCW_Direction cast to int is zero.

@param r            radii on axes before x-axis rotation
@param xAxisRotate  x-axis rotation in degrees; positive values are clockwise
@param largeArc     chooses smaller or larger arc
@param sweep        chooses clockwise or counterclockwise arc
@param xy           end of arc
@return             reference to SkPath
*/
func (o Path) ArcTo5(r Point, xAxisRotate float32, largeArc PathArcSize, sweep PathDirection, xy Point) Path {
	c_obj := o.sk
	c_r := *(*C.sk_SkPoint)(unsafe.Pointer(&r))
	c_xAxisRotate := C.float(xAxisRotate)
	c_largeArc := C.uint(largeArc)
	c_sweep := C.int(sweep)
	c_xy := *(*C.sk_SkPoint)(unsafe.Pointer(&xy))
	retC := C.misk_Path_arcTo5(c_obj, c_r, c_xAxisRotate, c_largeArc, c_sweep, c_xy)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath, relative to last SkPath SkPoint. Arc is implemented by one or
more conic, weighted to describe part of oval with radii (rx, ry) rotated by
xAxisRotate degrees. Arc curves from last SkPath SkPoint to relative end SkPoint:
(dx, dy), choosing one of four possible routes: clockwise or
counterclockwise, and smaller or larger. If SkPath is empty, the start arc SkPoint
is (0, 0).

Arc sweep is always less than 360 degrees. arcTo() appends line to end SkPoint
if either radii are zero, or if last SkPath SkPoint equals end SkPoint.
arcTo() scales radii (rx, ry) to fit last SkPath SkPoint and end SkPoint if both are
greater than zero but too small to describe an arc.

arcTo() appends up to four conic curves.
arcTo() implements the functionality of svg arc, although SVG "sweep-flag" value is
opposite the integer value of sweep; SVG "sweep-flag" uses 1 for clockwise, while
kCW_Direction cast to int is zero.

@param rx           radius before x-axis rotation
@param ry           radius before x-axis rotation
@param xAxisRotate  x-axis rotation in degrees; positive values are clockwise
@param largeArc     chooses smaller or larger arc
@param sweep        chooses clockwise or counterclockwise arc
@param dx           x-axis offset end of arc from last SkPath SkPoint
@param dy           y-axis offset end of arc from last SkPath SkPoint
@return             reference to SkPath
*/
func (o Path) RArcTo(rx float32, ry float32, xAxisRotate float32, largeArc PathArcSize, sweep PathDirection, dx float32, dy float32) Path {
	c_obj := o.sk
	c_rx := C.float(rx)
	c_ry := C.float(ry)
	c_xAxisRotate := C.float(xAxisRotate)
	c_largeArc := C.uint(largeArc)
	c_sweep := C.int(sweep)
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	retC := C.misk_Path_rArcTo(c_obj, c_rx, c_ry, c_xAxisRotate, c_largeArc, c_sweep, c_dx, c_dy)
	return Path{sk: &retC}
}

/*
Appends kClose_Verb to SkPath. A closed contour connects the first and last SkPoint
with line, forming a continuous loop. Open and closed contour draw the same
with SkPaint::kFill_Style. With SkPaint::kStroke_Style, open contour draws
SkPaint::Cap at contour start and end; closed contour draws
SkPaint::Join at contour start and end.

close() has no effect if SkPath is empty or last SkPath SkPath::Verb is kClose_Verb.

@return  reference to SkPath

example: https://fiddle.skia.org/c/@Path_close
*/
func (o Path) Close() Path {
	c_obj := o.sk
	retC := C.misk_Path_close(c_obj)
	return Path{sk: &retC}
}

/*
Returns true if SkPath is equivalent to SkRect when filled.
If false: rect, isClosed, and direction are unchanged.
If true: rect, isClosed, and direction are written to if not nullptr.

rect may be smaller than the SkPath bounds. SkPath bounds may include kMove_Verb points
that do not alter the area drawn by the returned rect.

@param rect       storage for bounds of SkRect; may be nullptr
@param isClosed   storage set to true if SkPath is closed; may be nullptr
@param direction  storage set to SkRect direction; may be nullptr
@return           true if SkPath contains SkRect

example: https://fiddle.skia.org/c/@Path_isRect
*/
func (o Path) IsRect(rect *Rect, isClosed *bool, direction *PathDirection) bool {
	c_obj := o.sk
	c_rect := (*C.sk_SkRect)(unsafe.Pointer(rect))
	c_isClosed := (*C.bool)(isClosed)
	c_direction := (*C.int)(direction)
	retC := C.misk_Path_isRect(c_obj, c_rect, c_isClosed, c_direction)
	return bool(retC)
}

/*
Adds a new contour to the path, defined by the rect, and wound in the
specified direction. The verbs added to the path will be:

kMove, kLine, kLine, kLine, kClose

start specifies which corner to begin the contour:
0: upper-left  corner
1: upper-right corner
2: lower-right corner
3: lower-left  corner

This start point also acts as the implied beginning of the subsequent,
contour, if it does not have an explicit moveTo(). e.g.

path.addRect(...)
// if we don't say moveTo() here, we will use the rect's start point
path.lineTo(...)

@param rect   SkRect to add as a closed contour
@param dir    SkPath::Direction to orient the new contour
@param start  initial corner of SkRect to add
@return       reference to SkPath

example: https://fiddle.skia.org/c/@Path_addRect_2
*/
func (o Path) AddRect1(rect Rect, dir PathDirection, start uint32) Path {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_dir := C.int(dir)
	c_start := C.uint(start)
	retC := C.misk_Path_addRect1(c_obj, c_rect, c_dir, c_start)
	return Path{sk: &retC}
}

func (o Path) AddRect2(rect Rect, dir PathDirection) Path {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_dir := C.int(dir)
	retC := C.misk_Path_addRect2(c_obj, c_rect, c_dir)
	return Path{sk: &retC}
}

func (o Path) AddRect3(left float32, top float32, right float32, bottom float32, dir PathDirection) Path {
	c_obj := o.sk
	c_left := C.float(left)
	c_top := C.float(top)
	c_right := C.float(right)
	c_bottom := C.float(bottom)
	c_dir := C.int(dir)
	retC := C.misk_Path_addRect3(c_obj, c_left, c_top, c_right, c_bottom, c_dir)
	return Path{sk: &retC}
}

/*
Adds oval to path, appending kMove_Verb, four kConic_Verb, and kClose_Verb.
Oval is upright ellipse bounded by SkRect oval with radii equal to half oval width
and half oval height. Oval begins at (oval.fRight, oval.centerY()) and continues
clockwise if dir is kCW_Direction, counterclockwise if dir is kCCW_Direction.

@param oval  bounds of ellipse added
@param dir   SkPath::Direction to wind ellipse
@return      reference to SkPath

example: https://fiddle.skia.org/c/@Path_addOval
*/
func (o Path) AddOval1(oval Rect, dir PathDirection) Path {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_dir := C.int(dir)
	retC := C.misk_Path_addOval1(c_obj, c_oval, c_dir)
	return Path{sk: &retC}
}

/*
Adds oval to SkPath, appending kMove_Verb, four kConic_Verb, and kClose_Verb.
Oval is upright ellipse bounded by SkRect oval with radii equal to half oval width
and half oval height. Oval begins at start and continues
clockwise if dir is kCW_Direction, counterclockwise if dir is kCCW_Direction.

@param oval   bounds of ellipse added
@param dir    SkPath::Direction to wind ellipse
@param start  index of initial point of ellipse
@return       reference to SkPath

example: https://fiddle.skia.org/c/@Path_addOval_2
*/
func (o Path) AddOval2(oval Rect, dir PathDirection, start uint32) Path {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_dir := C.int(dir)
	c_start := C.uint(start)
	retC := C.misk_Path_addOval2(c_obj, c_oval, c_dir, c_start)
	return Path{sk: &retC}
}

/*
Adds circle centered at (x, y) of size radius to SkPath, appending kMove_Verb,
four kConic_Verb, and kClose_Verb. Circle begins at: (x + radius, y), continuing
clockwise if dir is kCW_Direction, and counterclockwise if dir is kCCW_Direction.

Has no effect if radius is zero or negative.

@param x       center of circle
@param y       center of circle
@param radius  distance from center to edge
@param dir     SkPath::Direction to wind circle
@return        reference to SkPath
*/
func (o Path) AddCircle(x float32, y float32, radius float32, dir PathDirection) Path {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	c_radius := C.float(radius)
	c_dir := C.int(dir)
	retC := C.misk_Path_addCircle(c_obj, c_x, c_y, c_radius, c_dir)
	return Path{sk: &retC}
}

/*
Appends arc to SkPath, as the start of new contour. Arc added is part of ellipse
bounded by oval, from startAngle through sweepAngle. Both startAngle and
sweepAngle are measured in degrees, where zero degrees is aligned with the
positive x-axis, and positive sweeps extends arc clockwise.

If sweepAngle <= -360, or sweepAngle >= 360; and startAngle modulo 90 is nearly
zero, append oval instead of arc. Otherwise, sweepAngle values are treated
modulo 360, and arc may or may not draw depending on numeric rounding.

@param oval        bounds of ellipse containing arc
@param startAngle  starting angle of arc in degrees
@param sweepAngle  sweep, in degrees. Positive is clockwise; treated modulo 360
@return            reference to SkPath

example: https://fiddle.skia.org/c/@Path_addArc
*/
func (o Path) AddArc(oval Rect, startAngle float32, sweepAngle float32) Path {
	c_obj := o.sk
	c_oval := *(*C.sk_SkRect)(unsafe.Pointer(&oval))
	c_startAngle := C.float(startAngle)
	c_sweepAngle := C.float(sweepAngle)
	retC := C.misk_Path_addArc(c_obj, c_oval, c_startAngle, c_sweepAngle)
	return Path{sk: &retC}
}

/*
Appends SkRRect to SkPath, creating a new closed contour. SkRRect has bounds
equal to rect; each corner is 90 degrees of an ellipse with radii (rx, ry). If
dir is kCW_Direction, SkRRect starts at top-left of the lower-left corner and
winds clockwise. If dir is kCCW_Direction, SkRRect starts at the bottom-left
of the upper-left corner and winds counterclockwise.

If either rx or ry is too large, rx and ry are scaled uniformly until the
corners fit. If rx or ry is less than or equal to zero, addRoundRect() appends
SkRect rect to SkPath.

After appending, SkPath may be empty, or may contain: SkRect, oval, or SkRRect.

@param rect  bounds of SkRRect
@param rx    x-axis radius of rounded corners on the SkRRect
@param ry    y-axis radius of rounded corners on the SkRRect
@param dir   SkPath::Direction to wind SkRRect
@return      reference to SkPath
*/
func (o Path) AddRoundRect1(rect Rect, rx float32, ry float32, dir PathDirection) Path {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_rx := C.float(rx)
	c_ry := C.float(ry)
	c_dir := C.int(dir)
	retC := C.misk_Path_addRoundRect1(c_obj, c_rect, c_rx, c_ry, c_dir)
	return Path{sk: &retC}
}

/*
Appends SkRRect to SkPath, creating a new closed contour. SkRRect has bounds
equal to rect; each corner is 90 degrees of an ellipse with radii from the
array.

@param rect   bounds of SkRRect
@param radii  array of 8 SkScalar values, a radius pair for each corner
@param dir    SkPath::Direction to wind SkRRect
@return       reference to SkPath
*/
func (o Path) AddRoundRect2(rect Rect, radii []float32, dir PathDirection) Path {
	c_obj := o.sk
	c_rect := *(*C.sk_SkRect)(unsafe.Pointer(&rect))
	c_radii := (*C.float)(unsafe.Pointer(&radii[0]))
	c_dir := C.int(dir)
	retC := C.misk_Path_addRoundRect2(c_obj, c_rect, c_radii, c_dir)
	return Path{sk: &retC}
}

/*
Adds rrect to SkPath, creating a new closed contour. If
dir is kCW_Direction, rrect starts at top-left of the lower-left corner and
winds clockwise. If dir is kCCW_Direction, rrect starts at the bottom-left
of the upper-left corner and winds counterclockwise.

After appending, SkPath may be empty, or may contain: SkRect, oval, or SkRRect.

@param rrect  bounds and radii of rounded rectangle
@param dir    SkPath::Direction to wind SkRRect
@return       reference to SkPath

example: https://fiddle.skia.org/c/@Path_addRRect
*/
func (o Path) AddRRect1(rrect RRect, dir PathDirection) Path {
	c_obj := o.sk
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	c_dir := C.int(dir)
	retC := C.misk_Path_addRRect1(c_obj, c_rrect, c_dir)
	return Path{sk: &retC}
}

/*
Adds rrect to SkPath, creating a new closed contour. If dir is kCW_Direction, rrect
winds clockwise; if dir is kCCW_Direction, rrect winds counterclockwise.
start determines the first point of rrect to add.

@param rrect  bounds and radii of rounded rectangle
@param dir    SkPath::Direction to wind SkRRect
@param start  index of initial point of SkRRect
@return       reference to SkPath

example: https://fiddle.skia.org/c/@Path_addRRect_2
*/
func (o Path) AddRRect2(rrect RRect, dir PathDirection, start uint32) Path {
	c_obj := o.sk
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	c_dir := C.int(dir)
	c_start := C.uint(start)
	retC := C.misk_Path_addRRect2(c_obj, c_rrect, c_dir, c_start)
	return Path{sk: &retC}
}

/*
Adds contour created from line array, adding (count - 1) line segments.
Contour added starts at pts[0], then adds a line for every additional SkPoint
in pts array. If close is true, appends kClose_Verb to SkPath, connecting
pts[count - 1] and pts[0].

If count is zero, append kMove_Verb to path.
Has no effect if count is less than one.

@param pts    array of line sharing end and start SkPoint
@param count  length of SkPoint array
@param close  true to add line connecting contour end and start
@return       reference to SkPath

example: https://fiddle.skia.org/c/@Path_addPoly
*/
func (o Path) AddPoly(pts []Point, count int32, close bool) Path {
	c_obj := o.sk
	c_pts := (*C.sk_SkPoint)(unsafe.Pointer(&pts[0]))
	c_count := C.int(count)
	c_close := C.bool(close)
	retC := C.misk_Path_addPoly(c_obj, c_pts, c_count, c_close)
	return Path{sk: &retC}
}

/*
Appends src to SkPath, offset by (dx, dy).

If mode is kAppend_AddPathMode, src verb array, SkPoint array, and conic weights are
added unaltered. If mode is kExtend_AddPathMode, add line before appending
verbs, SkPoint, and conic weights.

@param src   SkPath verbs, SkPoint, and conic weights to add
@param dx    offset added to src SkPoint array x-axis coordinates
@param dy    offset added to src SkPoint array y-axis coordinates
@param mode  kAppend_AddPathMode or kExtend_AddPathMode
@return      reference to SkPath
*/
func (o Path) AddPath1(src Path, dx float32, dy float32, mode PathAddPathMode) Path {
	c_obj := o.sk
	c_src := src.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	c_mode := C.uint(mode)
	retC := C.misk_Path_addPath1(c_obj, c_src, c_dx, c_dy, c_mode)
	return Path{sk: &retC}
}

/*
Appends src to SkPath.

If mode is kAppend_AddPathMode, src verb array, SkPoint array, and conic weights are
added unaltered. If mode is kExtend_AddPathMode, add line before appending
verbs, SkPoint, and conic weights.

@param src   SkPath verbs, SkPoint, and conic weights to add
@param mode  kAppend_AddPathMode or kExtend_AddPathMode
@return      reference to SkPath
*/
func (o Path) AddPath2(src Path, mode PathAddPathMode) Path {
	c_obj := o.sk
	c_src := src.sk
	c_mode := C.uint(mode)
	retC := C.misk_Path_addPath2(c_obj, c_src, c_mode)
	return Path{sk: &retC}
}

/*
Appends src to SkPath, transformed by matrix. Transformed curves may have different
verbs, SkPoint, and conic weights.

If mode is kAppend_AddPathMode, src verb array, SkPoint array, and conic weights are
added unaltered. If mode is kExtend_AddPathMode, add line before appending
verbs, SkPoint, and conic weights.

@param src     SkPath verbs, SkPoint, and conic weights to add
@param matrix  transform applied to src
@param mode    kAppend_AddPathMode or kExtend_AddPathMode
@return        reference to SkPath
*/
func (o Path) AddPath3(src Path, matrix Matrix, mode PathAddPathMode) Path {
	c_obj := o.sk
	c_src := src.sk
	c_matrix := matrix.sk
	c_mode := C.uint(mode)
	retC := C.misk_Path_addPath3(c_obj, c_src, c_matrix, c_mode)
	return Path{sk: &retC}
}

/*
Appends src to SkPath, from back to front.
Reversed src always appends a new contour to SkPath.

@param src  SkPath verbs, SkPoint, and conic weights to add
@return     reference to SkPath

example: https://fiddle.skia.org/c/@Path_reverseAddPath
*/
func (o Path) ReverseAddPath(src Path) Path {
	c_obj := o.sk
	c_src := src.sk
	retC := C.misk_Path_reverseAddPath(c_obj, c_src)
	return Path{sk: &retC}
}

/*
Offsets SkPoint array by (dx, dy). Offset SkPath replaces dst.
If dst is nullptr, SkPath is replaced by offset data.

@param dx   offset added to SkPoint array x-axis coordinates
@param dy   offset added to SkPoint array y-axis coordinates
@param dst  overwritten, translated copy of SkPath; may be nullptr

example: https://fiddle.skia.org/c/@Path_offset
*/
func (o Path) Offset1(dx float32, dy float32, dst Path) {
	c_obj := o.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	c_dst := dst.sk
	C.misk_Path_offset1(c_obj, c_dx, c_dy, c_dst)
}

/*
Offsets SkPoint array by (dx, dy). SkPath is replaced by offset data.

@param dx  offset added to SkPoint array x-axis coordinates
@param dy  offset added to SkPoint array y-axis coordinates
*/
func (o Path) Offset2(dx float32, dy float32) Path {
	c_obj := o.sk
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	retC := C.misk_Path_offset2(c_obj, c_dx, c_dy)
	return Path{sk: &retC}
}

/*
Transforms verb array, SkPoint array, and weight by matrix.
transform may change verbs and increase their number.
Transformed SkPath replaces dst; if dst is nullptr, original data
is replaced.

@param matrix  SkMatrix to apply to SkPath
@param dst     overwritten, transformed copy of SkPath; may be nullptr
@param pc      whether to apply perspective clipping

example: https://fiddle.skia.org/c/@Path_transform
*/
func (o Path) Transform1(matrix Matrix, dst Path, pc ApplyPerspectiveClip) {
	c_obj := o.sk
	c_matrix := matrix.sk
	c_dst := dst.sk
	c_pc := C.int(pc)
	C.misk_Path_transform1(c_obj, c_matrix, c_dst, c_pc)
}

/*
Transforms verb array, SkPoint array, and weight by matrix.
transform may change verbs and increase their number.
SkPath is replaced by transformed data.

@param matrix  SkMatrix to apply to SkPath
@param pc      whether to apply perspective clipping
*/
func (o Path) Transform2(matrix Matrix, pc ApplyPerspectiveClip) Path {
	c_obj := o.sk
	c_matrix := matrix.sk
	c_pc := C.int(pc)
	retC := C.misk_Path_transform2(c_obj, c_matrix, c_pc)
	return Path{sk: &retC}
}

func (o Path) MakeTransform(m Matrix, pc ApplyPerspectiveClip) Path {
	c_obj := o.sk
	c_m := m.sk
	c_pc := C.int(pc)
	retC := C.misk_Path_makeTransform(c_obj, c_m, c_pc)
	return Path{sk: &retC}
}

func (o Path) MakeScale(sx float32, sy float32) Path {
	c_obj := o.sk
	c_sx := C.float(sx)
	c_sy := C.float(sy)
	retC := C.misk_Path_makeScale(c_obj, c_sx, c_sy)
	return Path{sk: &retC}
}

/*
Returns last point on SkPath in lastPt. Returns false if SkPoint array is empty,
storing (0, 0) if lastPt is not nullptr.

@param lastPt  storage for final SkPoint in SkPoint array; may be nullptr
@return        true if SkPoint array contains one or more SkPoint

example: https://fiddle.skia.org/c/@Path_getLastPt
*/
func (o Path) GetLastPt(lastPt Point) bool {
	c_obj := o.sk
	c_lastPt := lastPt.sk
	retC := C.misk_Path_getLastPt(c_obj, c_lastPt)
	return bool(retC)
}

/*
Sets last point to (x, y). If SkPoint array is empty, append kMove_Verb to
verb array and append (x, y) to SkPoint array.

@param x  set x-axis value of last point
@param y  set y-axis value of last point

example: https://fiddle.skia.org/c/@Path_setLastPt
*/
func (o Path) SetLastPt1(x float32, y float32) {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	C.misk_Path_setLastPt1(c_obj, c_x, c_y)
}

/*
Sets the last point on the path. If SkPoint array is empty, append kMove_Verb to
verb array and append p to SkPoint array.

@param p  set value of last point
*/
func (o Path) SetLastPt2(p Point) {
	c_obj := o.sk
	c_p := p.sk
	C.misk_Path_setLastPt2(c_obj, c_p)
}

/*
Returns a mask, where each set bit corresponds to a SegmentMask constant
if SkPath contains one or more verbs of that type.
Returns zero if SkPath contains no lines, or curves: quads, conics, or cubics.

getSegmentMasks() returns a cached result; it is very fast.

@return  SegmentMask bits or zero
*/
func (o Path) GetSegmentMasks() uint32 {
	c_obj := o.sk
	retC := C.misk_Path_getSegmentMasks(c_obj)
	return uint32(retC)
}

/*
Returns true if the point (x, y) is contained by SkPath, taking into
account FillType.

@param x  x-axis value of containment test
@param y  y-axis value of containment test
@return   true if SkPoint is in SkPath

example: https://fiddle.skia.org/c/@Path_contains
*/
func (o Path) Contains(x float32, y float32) bool {
	c_obj := o.sk
	c_x := C.float(x)
	c_y := C.float(y)
	retC := C.misk_Path_contains(c_obj, c_x, c_y)
	return bool(retC)
}

/*
Writes SkPath to buffer, returning the buffer written to, wrapped in SkData.

serialize() writes SkPath::FillType, verb array, SkPoint array, conic weight, and
additionally writes computed information like SkPath::Convexity and bounds.

serialize() should only be used in concert with readFromMemory().
The format used for SkPath in memory is not guaranteed.

@return  SkPath data wrapped in SkData buffer

example: https://fiddle.skia.org/c/@Path_serialize
*/
func (o Path) Serialize() Data {
	c_obj := o.sk
	retC := C.misk_Path_serialize(c_obj)
	return Data{sk: retC}
}

/*
Returns if SkPath data is consistent. Corrupt SkPath data is detected if
internal values are out of range or internal storage does not match
array dimensions.

@return  true if SkPath data is consistent
*/
func (o Path) IsValid() bool {
	c_obj := o.sk
	retC := C.misk_Path_isValid(c_obj)
	return bool(retC)
}

/*
AddPathMode chooses how addPath() appends. Adding one SkPath to another can extend
the last contour or start a new contour.
*/
type PathAddPathMode uint32

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
type PathArcSize uint32

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
type PathSegmentMask uint32

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
type PathVerb uint32

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
SkPoint holds two 32-bit floating point coordinates.
*/
type Point struct {
	sk *C.sk_SkPoint
}

// !< x-axis value
func (o Point) X() float32 {
	return float32(o.sk.fX)
}

func (o *Point) SetX(value float32) {
	o.sk.fX = C.float(value)
}

// !< y-axis value
func (o Point) Y() float32 {
	return float32(o.sk.fY)
}

func (o *Point) SetY(value float32) {
	o.sk.fY = C.float(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Point has not been created.
func (o Point) IsNil() bool {
	return o.sk == nil
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
	sk *C.sk_SkPixmap
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Pixmap has not been created.
func (o Pixmap) IsNil() bool {
	return o.sk == nil
}

/*
Creates an empty SkPixmap without pixels, with kUnknown_SkColorType, with
kUnknown_SkAlphaType, and with a width and height of zero. Use
reset() to associate pixels, SkColorType, SkAlphaType, width, and height
after SkPixmap has been created.

@return  empty SkPixmap
*/
func NewPixmap() Pixmap {

	retC := C.misk_new_Pixmap()
	return Pixmap{sk: retC}
}

/*
Creates SkPixmap from info width, height, SkAlphaType, and SkColorType.
addr points to pixels, or nullptr. rowBytes should be info.width() times
info.bytesPerPixel(), or larger.

No parameter checking is performed; it is up to the caller to ensure that
addr and rowBytes agree with info.

The memory lifetime of pixels is managed by the caller. When SkPixmap goes
out of scope, addr is unaffected.

SkPixmap may be later modified by reset() to change its size, pixel type, or
storage.

@param info      width, height, SkAlphaType, SkColorType of SkImageInfo
@param addr      pointer to pixels allocated by caller; may be nullptr
@param rowBytes  size of one row of addr; width times pixel size, or larger
@return          initialized SkPixmap
*/
func NewPixmapImageInfo(info ImageInfo, addr []byte, rowBytes uint32) Pixmap {
	c_info := info.sk
	c_addr := unsafe.Pointer(&addr[0])
	c_rowBytes := C.ulong(rowBytes)
	retC := C.misk_new_PixmapImageInfo(c_info, c_addr, c_rowBytes)
	return Pixmap{sk: retC}
}

func (o *Pixmap) Delete() {
	C.misk_delete_SkPixmap(o.sk)
	o.sk = nil
}

/*
SkRect holds four float coordinates describing the upper and
lower bounds of a rectangle. SkRect may be created from outer bounds or
from position, width, and height. SkRect describes an area; if its right
is less than or equal to its left, or if its bottom is less than or equal to
its top, it is considered empty.
*/
type Rect C.sk_SkRect

// !< smaller x-axis bounds
func (o Rect) Left() float32 {
	return float32(o.fLeft)
}

func (o *Rect) SetLeft(value float32) {
	o.fLeft = C.float(value)
}

// !< smaller y-axis bounds
func (o Rect) Top() float32 {
	return float32(o.fTop)
}

func (o *Rect) SetTop(value float32) {
	o.fTop = C.float(value)
}

// !< larger x-axis bounds
func (o Rect) Right() float32 {
	return float32(o.fRight)
}

func (o *Rect) SetRight(value float32) {
	o.fRight = C.float(value)
}

// !< larger y-axis bounds
func (o Rect) Bottom() float32 {
	return float32(o.fBottom)
}

func (o *Rect) SetBottom(value float32) {
	o.fBottom = C.float(value)
}

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

/*
Returns left edge of SkRect, if sorted. Call isSorted() to see if SkRect is valid.
Call sort() to reverse fLeft and fRight if needed.

@return  fLeft
*/
func (o *Rect) X() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_x(c_obj)
	return float32(retC)
}

/*
Returns top edge of SkRect, if sorted. Call isEmpty() to see if SkRect may be invalid,
and sort() to reverse fTop and fBottom if needed.

@return  fTop
*/
func (o *Rect) Y() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_y(c_obj)
	return float32(retC)
}

/*
Returns span on the x-axis. This does not check if SkRect is sorted, or if
result fits in 32-bit float; result may be negative or infinity.

@return  fRight minus fLeft
*/
func (o *Rect) Width() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_width(c_obj)
	return float32(retC)
}

/*
Returns span on the y-axis. This does not check if SkRect is sorted, or if
result fits in 32-bit float; result may be negative or infinity.

@return  fBottom minus fTop
*/
func (o *Rect) Height() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_height(c_obj)
	return float32(retC)
}

/*
Returns average of left edge and right edge. Result does not change if SkRect
is sorted. Result may overflow to infinity if SkRect is far from the origin.

@return  midpoint on x-axis
*/
func (o *Rect) CenterX() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_centerX(c_obj)
	return float32(retC)
}

/*
Returns average of top edge and bottom edge. Result does not change if SkRect
is sorted.

@return  midpoint on y-axis
*/
func (o *Rect) CenterY() float32 {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_centerY(c_obj)
	return float32(retC)
}

/*
Returns true if fLeft is equal to or greater than fRight, or if fTop is equal
to or greater than fBottom. Call sort() to reverse rectangles with negative
width() or height().

@return  true if width() or height() are zero or negative
*/
func (o *Rect) IsEmpty() bool {
	c_obj := (*C.sk_SkRect)(o)
	retC := C.misk_Rect_isEmpty(c_obj)
	return bool(retC)
}

/*
Sets SkRect to (0, 0, 0, 0).

Many other rectangles are empty; if left is equal to or greater than right,
or if top is equal to or greater than bottom. Setting all members to zero
is a convenience, but does not designate a special empty rectangle.
*/
func (o *Rect) SetEmpty() {
	c_obj := (*C.sk_SkRect)(o)
	C.misk_Rect_setEmpty(c_obj)
}

/*
Sets SkRect to (left, top, right, bottom).
left and right are not sorted; left is not necessarily less than right.
top and bottom are not sorted; top is not necessarily less than bottom.

@param left    stored in fLeft
@param top     stored in fTop
@param right   stored in fRight
@param bottom  stored in fBottom
*/
func (o *Rect) SetLTRB(left float32, top float32, right float32, bottom float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_left := C.float(left)
	c_top := C.float(top)
	c_right := C.float(right)
	c_bottom := C.float(bottom)
	C.misk_Rect_setLTRB(c_obj, c_left, c_top, c_right, c_bottom)
}

/*
Sets SkRect to (x, y, x + width, y + height).
Does not validate input; width or height may be negative.

@param x       stored in fLeft
@param y       stored in fTop
@param width   added to x and stored in fRight
@param height  added to y and stored in fBottom
*/
func (o *Rect) SetXYWH(x float32, y float32, width float32, height float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_x := C.float(x)
	c_y := C.float(y)
	c_width := C.float(width)
	c_height := C.float(height)
	C.misk_Rect_setXYWH(c_obj, c_x, c_y, c_width, c_height)
}

/*
Sets SkRect to (0, 0, width, height). Does not validate input;
width or height may be negative.

@param width   stored in fRight
@param height  stored in fBottom
*/
func (o *Rect) SetWH(width float32, height float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_width := C.float(width)
	c_height := C.float(height)
	C.misk_Rect_setWH(c_obj, c_width, c_height)
}

/*
Offsets SkRect by adding dx to fLeft, fRight; and by adding dy to fTop, fBottom.

If dx is negative, moves SkRect to the left.
If dx is positive, moves SkRect to the right.
If dy is negative, moves SkRect upward.
If dy is positive, moves SkRect downward.

@param dx  offset added to fLeft and fRight
@param dy  offset added to fTop and fBottom
*/
func (o *Rect) Offset(dx float32, dy float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	C.misk_Rect_offset(c_obj, c_dx, c_dy)
}

/*
Offsets SkRect so that fLeft equals newX, and fTop equals newY. width and height
are unchanged.

@param newX  stored in fLeft, preserving width()
@param newY  stored in fTop, preserving height()
*/
func (o *Rect) OffsetTo(newX float32, newY float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_newX := C.float(newX)
	c_newY := C.float(newY)
	C.misk_Rect_offsetTo(c_obj, c_newX, c_newY)
}

/*
Insets SkRect by (dx, dy).

If dx is positive, makes SkRect narrower.
If dx is negative, makes SkRect wider.
If dy is positive, makes SkRect shorter.
If dy is negative, makes SkRect taller.

@param dx  added to fLeft and subtracted from fRight
@param dy  added to fTop and subtracted from fBottom
*/
func (o *Rect) Inset(dx float32, dy float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	C.misk_Rect_inset(c_obj, c_dx, c_dy)
}

/*
Outsets SkRect by (dx, dy).

If dx is positive, makes SkRect wider.
If dx is negative, makes SkRect narrower.
If dy is positive, makes SkRect taller.
If dy is negative, makes SkRect shorter.

@param dx  subtracted to fLeft and added from fRight
@param dy  subtracted to fTop and added from fBottom
*/
func (o *Rect) Outset(dx float32, dy float32) {
	c_obj := (*C.sk_SkRect)(o)
	c_dx := C.float(dx)
	c_dy := C.float(dy)
	C.misk_Rect_outset(c_obj, c_dx, c_dy)
}

/*
Returns true if: fLeft <= x < fRight && fTop <= y < fBottom.
Returns false if SkRect is empty.

@param x  test SkPoint x-coordinate
@param y  test SkPoint y-coordinate
@return   true if (x, y) is inside SkRect
*/
func (o *Rect) Contains(x float32, y float32) bool {
	c_obj := (*C.sk_SkRect)(o)
	c_x := C.float(x)
	c_y := C.float(y)
	retC := C.misk_Rect_contains(c_obj, c_x, c_y)
	return bool(retC)
}

/*
Returns true if SkRect contains r.
Returns false if SkRect is empty or r is empty.

SkRect contains r when SkRect area completely includes r area.

@param r  SkRect contained
@return   true if all sides of SkRect are outside r
*/
func (o *Rect) ContainsRect(r Rect) bool {
	c_obj := (*C.sk_SkRect)(o)
	c_r := *(*C.sk_SkRect)(unsafe.Pointer(&r))
	retC := C.misk_Rect_containsRect(c_obj, c_r)
	return bool(retC)
}

/*
Returns true if SkRect intersects r, and sets SkRect to intersection.
Returns false if SkRect does not intersect r, and leaves SkRect unchanged.

Returns false if either r or SkRect is empty, leaving SkRect unchanged.

@param r  limit of result
@return   true if r and SkRect have area in common

example: https://fiddle.skia.org/c/@Rect_intersect
*/
func (o *Rect) Intersect(r Rect) bool {
	c_obj := (*C.sk_SkRect)(o)
	c_r := *(*C.sk_SkRect)(unsafe.Pointer(&r))
	retC := C.misk_Rect_intersect(c_obj, c_r)
	return bool(retC)
}

/*
Sets SkRect to the union of itself and r.

Has no effect if r is empty. Otherwise, if SkRect is empty, sets
SkRect to r.

@param r  expansion SkRect

example: https://fiddle.skia.org/c/@Rect_join_2
*/
func (o *Rect) Join(r Rect) {
	c_obj := (*C.sk_SkRect)(o)
	c_r := *(*C.sk_SkRect)(unsafe.Pointer(&r))
	C.misk_Rect_join(c_obj, c_r)
}

/*
Swaps fLeft and fRight if fLeft is greater than fRight; and swaps
fTop and fBottom if fTop is greater than fBottom. Result may be empty;
and width() and height() will be zero or positive.
*/
func (o *Rect) Sort() {
	c_obj := (*C.sk_SkRect)(o)
	C.misk_Rect_sort(c_obj)
}

/*
SkRRect describes a rounded rectangle with a bounds and a pair of radii for each corner.
The bounds and radii can be set so that SkRRect describes: a rectangle with sharp corners;
a circle; an oval; or a rectangle with one or more rounded corners.

SkRRect allows implementing CSS properties that describe rounded corners.
SkRRect may have up to eight different radii, one for each axis on each of its four
corners.

SkRRect may modify the provided parameters when initializing bounds and radii.
If either axis radii is zero or less: radii are stored as zero; corner is square.
If corner curves overlap, radii are proportionally reduced to fit within bounds.
*/
type RRect C.sk_SkRRect

/*
Initializes bounds at (0, 0), the origin, with zero width and height.
Initializes corner radii to (0, 0), and sets type of kEmpty_Type.

@return  empty SkRRect
*/
func NewRRect() RRect {

	retC := C.misk_new_RRect()
	return *(*RRect)(unsafe.Pointer(&retC))
}

/*
Initializes to copy of rrect bounds and corner radii.

@param rrect  bounds and corner to copy
@return       copy of rrect
*/
func NewRRectCopy(rrect RRect) RRect {
	c_rrect := *(*C.sk_SkRRect)(unsafe.Pointer(&rrect))
	retC := C.misk_new_RRectCopy(c_rrect)
	return *(*RRect)(unsafe.Pointer(&retC))
}

/*
SkRegion describes the set of pixels used to clip SkCanvas. SkRegion is compact,
efficiently storing a single integer rectangle, or a run length encoded array
of rectangles. SkRegion may reduce the current SkCanvas clip, or may be drawn as
one or more integer rectangles. SkRegion iterator returns the scan lines or
rectangles contained by it, optionally intersecting a bounding rectangle.
*/
type Region struct {
	sk *C.sk_SkRegion
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Region has not been created.
func (o Region) IsNil() bool {
	return o.sk == nil
}

/*
Constructs an empty SkRegion. SkRegion is set to empty bounds
at (0, 0) with zero width and height.

@return  empty SkRegion

example: https://fiddle.skia.org/c/@Region_empty_constructor
*/
func NewRegion() Region {

	retC := C.misk_new_Region()
	return Region{sk: retC}
}

/*
Constructs a copy of an existing region.
Copy constructor makes two regions identical by value. Internally, region and
the returned result share pointer values. The underlying SkRect array is
copied when modified.

Creating a SkRegion copy is very efficient and never allocates memory.
SkRegion are always copied by value from the interface; the underlying shared
pointers are not exposed.

@param region  SkRegion to copy by value
@return        copy of SkRegion

example: https://fiddle.skia.org/c/@Region_copy_const_SkRegion
*/
func NewRegionCopy(region Region) Region {
	c_region := region.sk
	retC := C.misk_new_RegionCopy(c_region)
	return Region{sk: retC}
}

/*
Constructs a rectangular SkRegion matching the bounds of rect.

@param rect  bounds of constructed SkRegion
@return      rectangular SkRegion

example: https://fiddle.skia.org/c/@Region_copy_const_SkIRect
*/
func NewRegionCopyRect(rect IRect) Region {
	c_rect := *(*C.sk_SkIRect)(unsafe.Pointer(&rect))
	retC := C.misk_new_RegionCopyRect(c_rect)
	return Region{sk: retC}
}

/*
Releases ownership of any shared data and deletes data if SkRegion is sole owner.

example: https://fiddle.skia.org/c/@Region_destructor
*/
func (o *Region) Delete() {
	C.misk_delete_SkRegion(o.sk)
	o.sk = nil
}

/*
RGBA color value, holding four floating point components. Color components are always in
a known order. kAT determines if the SkRGBA4f's R, G, and B components are premultiplied
by alpha or not.

Skia's public API always uses unpremultiplied colors, which can be stored as
SkRGBA4f<kUnpremul_SkAlphaType>. For convenience, this type can also be referred to
as SkColor4f.
*/
type RGBA4f C.sk_SkRGBA4f

// !< red component
func (o RGBA4f) R() float32 {
	return float32(o.fR)
}

func (o *RGBA4f) SetR(value float32) {
	o.fR = C.float(value)
}

// !< green component
func (o RGBA4f) G() float32 {
	return float32(o.fG)
}

func (o *RGBA4f) SetG(value float32) {
	o.fG = C.float(value)
}

// !< blue component
func (o RGBA4f) B() float32 {
	return float32(o.fB)
}

func (o *RGBA4f) SetB(value float32) {
	o.fB = C.float(value)
}

// !< alpha component
func (o RGBA4f) A() float32 {
	return float32(o.fA)
}

func (o *RGBA4f) SetA(value float32) {
	o.fA = C.float(value)
}

type SamplingOptions struct {
	sk *C.sk_SkSamplingOptions
}

func (o SamplingOptions) MaxAniso() int32 {
	return int32(o.sk.maxAniso)
}

func (o *SamplingOptions) SetmaxAniso(value int32) {
	o.sk.maxAniso = C.int(value)
}

func (o SamplingOptions) UseCubic() bool {
	return bool(o.sk.useCubic)
}

func (o *SamplingOptions) SetuseCubic(value bool) {
	o.sk.useCubic = C.bool(value)
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the SamplingOptions has not been created.
func (o SamplingOptions) IsNil() bool {
	return o.sk == nil
}

func NewSamplingOptions() SamplingOptions {

	retC := C.misk_new_SamplingOptions()
	return SamplingOptions{sk: retC}
}

func NewSamplingOptionsCopy(p0 SamplingOptions) SamplingOptions {
	c_p0 := p0.sk
	retC := C.misk_new_SamplingOptionsCopy(c_p0)
	return SamplingOptions{sk: retC}
}

func (o *SamplingOptions) Delete() {
	C.misk_delete_SkSamplingOptions(o.sk)
	o.sk = nil
}

// /////////////////////////////////////////////////////////////////////////////
type Size C.sk_SkSize

func (o Size) Width() float32 {
	return float32(o.fWidth)
}

func (o *Size) SetWidth(value float32) {
	o.fWidth = C.float(value)
}

func (o Size) Height() float32 {
	return float32(o.fHeight)
}

func (o *Size) SetHeight(value float32) {
	o.fHeight = C.float(value)
}

/*
SkStream -- abstraction for a source of bytes. Subclasses can be backed by
memory, or a file, or something else.
*/
type Stream struct {
	sk *C.sk_SkStream
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Stream has not been created.
func (o Stream) IsNil() bool {
	return o.sk == nil
}

func (o *Stream) Delete() {
	C.misk_delete_SkStream(o.sk)
	o.sk = nil
}

/*
Light weight class for managing strings. Uses reference
counting to make string assignments and copies very fast
with no extra RAM cost. Assumes UTF8 encoding.
*/
type String struct {
	sk *C.sk_SkString
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the String has not been created.
func (o String) IsNil() bool {
	return o.sk == nil
}

func NewString(text string) String {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	retC := C.misk_new_String(c_text)
	return String{sk: retC}
}

func (o *String) Delete() {
	C.misk_delete_SkString(o.sk)
	o.sk = nil
}

func (o String) Data() string {
	c_obj := o.sk
	retC := C.misk_String_data(c_obj)
	return C.GoString(retC)
}

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
	sk *C.sk_SkSurface
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the Surface has not been created.
func (o Surface) IsNil() bool {
	return o.sk == nil
}

func (o *Surface) Unref() {
	C.misk_unref_SkSurface(o.sk)
	o.sk = nil
}

/*
Returns SkCanvas that draws into SkSurface. Subsequent calls return the same SkCanvas.
SkCanvas returned is managed and owned by SkSurface, and is deleted when SkSurface
is deleted.

@return  drawing SkCanvas for SkSurface

example: https://fiddle.skia.org/c/@Surface_getCanvas
*/
func (o Surface) GetCanvas() Canvas {
	c_obj := o.sk
	retC := C.misk_Surface_getCanvas(c_obj)
	return Canvas{sk: retC}
}

/*
Calls makeSurface(ImageInfo) with the same ImageInfo as this surface, but with the

	specified width and height.
*/
func (o Surface) MakeSurface(width int32, height int32) Surface {
	c_obj := o.sk
	c_width := C.int(width)
	c_height := C.int(height)
	retC := C.misk_Surface_makeSurface(c_obj, c_width, c_height)
	return Surface{sk: retC}
}

/*
Returns SkImage capturing SkSurface contents. Subsequent drawing to SkSurface contents
are not captured. SkImage allocation is accounted for if SkSurface was created with
skgpu::Budgeted::kYes.

@return  SkImage initialized with SkSurface contents

example: https://fiddle.skia.org/c/@Surface_makeImageSnapshot
*/
func (o Surface) MakeImageSnapshot() Image {
	c_obj := o.sk
	retC := C.misk_Surface_makeImageSnapshot(c_obj)
	return Image{sk: retC}
}

/*
	Like the no-parameter version, this returns an image of the current surface contents.
	This variant takes a rectangle specifying the subset of the surface that is of interest.
	These bounds will be sanitized before being used.
	- If bounds extends beyond the surface, it will be trimmed to just the intersection of
	  it and the surface.
	- If bounds does not intersect the surface, then this returns nullptr.
	- If bounds == the surface, then this is the same as calling the no-parameter variant.

example: https://fiddle.skia.org/c/@Surface_makeImageSnapshot_2
*/
func (o Surface) MakeImageSnapshotBounds(bounds IRect) Image {
	c_obj := o.sk
	c_bounds := *(*C.sk_SkIRect)(unsafe.Pointer(&bounds))
	retC := C.misk_Surface_makeImageSnapshotBounds(c_obj, c_bounds)
	return Image{sk: retC}
}

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
func NewSurfacePropsPixelGeometry(flags uint32, p1 PixelGeometry) SurfaceProps {
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

func (o *SurfaceProps) Delete() {
	C.misk_delete_SkSurfaceProps(o.sk)
	o.sk = nil
}

type SurfacePropsFlags uint32

const (
	SurfacePropsFlagsDefault_Flag                   SurfacePropsFlags = 0
	SurfacePropsFlagsUseDeviceIndependentFonts_Flag SurfacePropsFlags = 1
	SurfacePropsFlagsDynamicMSAA_Flag               SurfacePropsFlags = 2
	SurfacePropsFlagsAlwaysDither_Flag              SurfacePropsFlags = 4
)

type SVGDOM struct {
	sk *C.sk_SkSVGDOM
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the SVGDOM has not been created.
func (o SVGDOM) IsNil() bool {
	return o.sk == nil
}

func (o *SVGDOM) Unref() {
	C.misk_unref_SkSVGDOM(o.sk)
	o.sk = nil
}
func (o *SVGDOM) Delete() {
	C.misk_delete_SkSVGDOM(o.sk)
	o.sk = nil
}

/*
Returns the root (outermost) SVG element.
*/
func (o SVGDOM) GetRoot() SVGSVG {
	c_obj := o.sk
	retC := C.misk_SVGDOM_getRoot(c_obj)
	return SVGSVG{sk: retC}
}

/*
Specify a "container size" for the SVG dom.

This is used to resolve the initial viewport when the root SVG width/height are specified
in relative units.

If the root dimensions are in absolute units, then the container size has no effect since
the initial viewport is fixed.
*/
func (o SVGDOM) SetContainerSize(p0 Size) {
	c_obj := o.sk
	c_p0 := *(*C.sk_SkSize)(unsafe.Pointer(&p0))
	C.misk_SVGDOM_setContainerSize(c_obj, c_p0)
}

/*
DEPRECATED: use getRoot()->intrinsicSize() to query the root element intrinsic size.

Returns the SVG dom container size.

If the client specified a container size via setContainerSize(), then the same size is
returned.

When unspecified by clients, this returns the intrinsic size of the root element, as defined
by its width/height attributes.  If either width or height is specified in relative units
(e.g. "100%"), then the corresponding intrinsic size dimension is zero.
*/
func (o SVGDOM) ContainerSize() Size {
	c_obj := o.sk
	retC := C.misk_SVGDOM_containerSize(c_obj)
	return Size(retC)
}

func (o SVGDOM) Render(p0 Canvas) {
	c_obj := o.sk
	c_p0 := p0.sk
	C.misk_SVGDOM_render(c_obj, c_p0)
}

func SVGDOMMakeFromStream(str Stream) SVGDOM {
	c_str := str.sk
	retC := C.misk_SVGDOM_MakeFromStream(c_str)
	return SVGDOM{sk: retC}
}

type SVGSVG struct {
	sk *C.sk_SkSVGSVG
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the SVGSVG has not been created.
func (o SVGSVG) IsNil() bool {
	return o.sk == nil
}

func (o SVGSVG) IntrinsicSize(p0 SVGLengthContext) Size {
	c_obj := o.sk
	c_p0 := p0.sk
	retC := C.misk_SVGSVG_intrinsicSize(c_obj, c_p0)
	return Size(retC)
}

type SVGLengthContext struct {
	sk *C.sk_SkSVGLengthContext
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the SVGLengthContext has not been created.
func (o SVGLengthContext) IsNil() bool {
	return o.sk == nil
}

func NewSVGLengthContext(viewport Size, dpi float32) SVGLengthContext {
	c_viewport := *(*C.sk_SkSize)(unsafe.Pointer(&viewport))
	c_dpi := C.float(dpi)
	retC := C.misk_new_SVGLengthContext(c_viewport, c_dpi)
	return SVGLengthContext{sk: retC}
}

func (o *SVGLengthContext) Delete() {
	C.misk_delete_SkSVGLengthContext(o.sk)
	o.sk = nil
}

func (o SVGLengthContext) ViewPort() Size {
	c_obj := o.sk
	retC := C.misk_SVGLengthContext_viewPort(c_obj)
	return Size(retC)
}

func (o SVGLengthContext) SetViewPort(viewport Size) {
	c_obj := o.sk
	c_viewport := *(*C.sk_SkSize)(unsafe.Pointer(&viewport))
	C.misk_SVGLengthContext_setViewPort(c_obj, c_viewport)
}

/*
SkTextBlob combines multiple text runs into an immutable container. Each text
run consists of glyphs, SkPaint, and position. Only parts of SkPaint related to
fonts and text rendering are used by run.
*/
type TextBlob struct {
	sk *C.sk_SkTextBlob
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the TextBlob has not been created.
func (o TextBlob) IsNil() bool {
	return o.sk == nil
}

func (o *TextBlob) Unref() {
	C.misk_unref_SkTextBlob(o.sk)
	o.sk = nil
}

/*
Returns conservative bounding box. Uses SkPaint associated with each glyph to
determine glyph bounds, and unions all bounds. Returned bounds may be
larger than the bounds of all glyphs in runs.

@return  conservative bounding box
*/
func (o TextBlob) Bounds() Rect {
	c_obj := o.sk
	retC := C.misk_TextBlob_bounds(c_obj)
	return Rect(retC)
}

/*
Returns a non-zero value unique among all text blobs.

@return  identifier for SkTextBlob
*/
func (o TextBlob) UniqueID() uint32 {
	c_obj := o.sk
	retC := C.misk_TextBlob_uniqueID(c_obj)
	return uint32(retC)
}

/*
Creates SkTextBlob with a single run. string meaning depends on SkTextEncoding;
by default, string is encoded as UTF-8.

font contains attributes used to define the run text.

When encoding is SkTextEncoding::kUTF8, SkTextEncoding::kUTF16, or
SkTextEncoding::kUTF32, this function uses the default
character-to-glyph mapping from the SkTypeface in font.  It does not
perform typeface fallback for characters not found in the SkTypeface.
It does not perform kerning or other complex shaping; glyphs are
positioned based on their default advances.

@param string   character code points or glyphs drawn
@param font     text size, typeface, text scale, and so on, used to draw
@param encoding text encoding used in the text array
@return         SkTextBlob constructed from one run
*/
func TextBlobMakeFromString(string string, font Font, encoding TextEncoding) TextBlob {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))
	c_font := font.sk
	c_encoding := C.int(encoding)
	retC := C.misk_TextBlob_MakeFromString(c_string, c_font, c_encoding)
	return TextBlob{sk: retC}
}

/*
Returns a textblob built from a single run of text with x-positions and a single y value.
This is equivalent to using SkTextBlobBuilder and calling allocRunPosH().
Returns nullptr if byteLength is zero.

@param text        character code points or glyphs drawn (based on encoding)
@param byteLength  byte length of text array
@param xpos    array of x-positions, must contain values for all of the character points.
@param constY  shared y-position for each character point, to be paired with each xpos.
@param font    SkFont used for this run
@param encoding specifies the encoding of the text array.
@return        new textblob or nullptr
*/
func TextBlobMakeFromPosTextH(text []byte, byteLength uint32, xpos []float32, constY float32, font Font, encoding TextEncoding) TextBlob {
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_xpos := (*C.float)(unsafe.Pointer(&xpos[0]))
	c_constY := C.float(constY)
	c_font := font.sk
	c_encoding := C.int(encoding)
	retC := C.misk_TextBlob_MakeFromPosTextH(c_text, c_byteLength, c_xpos, c_constY, c_font, c_encoding)
	return TextBlob{sk: retC}
}

/*
Returns a textblob built from a single run of text with positions.
This is equivalent to using SkTextBlobBuilder and calling allocRunPos().
Returns nullptr if byteLength is zero.

@param text        character code points or glyphs drawn (based on encoding)
@param byteLength  byte length of text array
@param pos     array of positions, must contain values for all of the character points.
@param font    SkFont used for this run
@param encoding specifies the encoding of the text array.
@return        new textblob or nullptr
*/
func TextBlobMakeFromPosText(text []byte, byteLength uint32, pos []Point, font Font, encoding TextEncoding) TextBlob {
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_pos := (*C.sk_SkPoint)(unsafe.Pointer(&pos[0]))
	c_font := font.sk
	c_encoding := C.int(encoding)
	retC := C.misk_TextBlob_MakeFromPosText(c_text, c_byteLength, c_pos, c_font, c_encoding)
	return TextBlob{sk: retC}
}

/*
Helper class for constructing SkTextBlob.
*/
type TextBlobBuilder struct {
	sk *C.sk_SkTextBlobBuilder
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the TextBlobBuilder has not been created.
func (o TextBlobBuilder) IsNil() bool {
	return o.sk == nil
}

/*
Constructs empty SkTextBlobBuilder. By default, SkTextBlobBuilder has no runs.

@return  empty SkTextBlobBuilder

example: https://fiddle.skia.org/c/@TextBlobBuilder_empty_constructor
*/
func NewTextBlobBuilder() TextBlobBuilder {

	retC := C.misk_new_TextBlobBuilder()
	return TextBlobBuilder{sk: retC}
}

/*
Deletes data allocated internally by SkTextBlobBuilder.
*/
func (o *TextBlobBuilder) Delete() {
	C.misk_delete_SkTextBlobBuilder(o.sk)
	o.sk = nil
}

/*
Returns SkTextBlob built from runs of glyphs added by builder. Returned
SkTextBlob is immutable; it may be copied, but its contents may not be altered.
Returns nullptr if no runs of glyphs were added by builder.

Resets SkTextBlobBuilder to its initial empty state, allowing it to be
reused to build a new set of runs.

@return  SkTextBlob or nullptr

example: https://fiddle.skia.org/c/@TextBlobBuilder_make
*/
func (o TextBlobBuilder) Make() TextBlob {
	c_obj := o.sk
	retC := C.misk_TextBlobBuilder_make(c_obj)
	return TextBlob{sk: retC}
}

/*
Returns run with storage for glyphs. Caller must write count glyphs to
RunBuffer::glyphs before next call to SkTextBlobBuilder.

RunBuffer::pos, RunBuffer::utf8text, and RunBuffer::clusters should be ignored.

Glyphs share metrics in font.

Glyphs are positioned on a baseline at (x, y), using font metrics to
determine their relative placement.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from (x, y) and RunBuffer::glyphs metrics.

@param font    SkFont used for this run
@param count   number of glyphs
@param x       horizontal offset within the blob
@param y       vertical offset within the blob
@param bounds  optional run bounding box
@return writable glyph buffer
*/
func (o TextBlobBuilder) AllocRun(font Font, count int32, x float32, y float32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_x := C.float(x)
	c_y := C.float(y)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRun(c_obj, c_font, c_count, c_x, c_y, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
Returns run with storage for glyphs and positions along baseline. Caller must
write count glyphs to RunBuffer::glyphs and count scalars to RunBuffer::pos
before next call to SkTextBlobBuilder.

RunBuffer::utf8text and RunBuffer::clusters should be ignored.

Glyphs share metrics in font.

Glyphs are positioned on a baseline at y, using x-axis positions written by
caller to RunBuffer::pos.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from y, RunBuffer::pos, and RunBuffer::glyphs metrics.

@param font    SkFont used for this run
@param count   number of glyphs
@param y       vertical offset within the blob
@param bounds  optional run bounding box
@return writable glyph buffer and x-axis position buffer
*/
func (o TextBlobBuilder) AllocRunPosH(font Font, count int32, y float32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_y := C.float(y)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRunPosH(c_obj, c_font, c_count, c_y, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
Returns run with storage for glyphs and SkPoint positions. Caller must
write count glyphs to RunBuffer::glyphs and count SkPoint to RunBuffer::pos
before next call to SkTextBlobBuilder.

RunBuffer::utf8text and RunBuffer::clusters should be ignored.

Glyphs share metrics in font.

Glyphs are positioned using SkPoint written by caller to RunBuffer::pos, using
two scalar values for each SkPoint.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from RunBuffer::pos, and RunBuffer::glyphs metrics.

@param font    SkFont used for this run
@param count   number of glyphs
@param bounds  optional run bounding box
@return writable glyph buffer and SkPoint buffer
*/
func (o TextBlobBuilder) AllocRunPos(font Font, count int32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRunPos(c_obj, c_font, c_count, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
Returns run with storage for glyphs, text, and clusters. Caller must
write count glyphs to RunBuffer::glyphs, textByteCount UTF-8 code units
into RunBuffer::utf8text, and count monotonic indexes into utf8text
into RunBuffer::clusters before next call to SkTextBlobBuilder.

RunBuffer::pos should be ignored.

Glyphs share metrics in font.

Glyphs are positioned on a baseline at (x, y), using font metrics to
determine their relative placement.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from (x, y) and RunBuffer::glyphs metrics.

@param font          SkFont used for this run
@param count         number of glyphs
@param x             horizontal offset within the blob
@param y             vertical offset within the blob
@param textByteCount number of UTF-8 code units
@param bounds        optional run bounding box
@return writable glyph buffer, text buffer, and cluster buffer
*/
func (o TextBlobBuilder) AllocRunText(font Font, count int32, x float32, y float32, textByteCount int32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_x := C.float(x)
	c_y := C.float(y)
	c_textByteCount := C.int(textByteCount)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRunText(c_obj, c_font, c_count, c_x, c_y, c_textByteCount, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
Returns run with storage for glyphs, positions along baseline, text,
and clusters. Caller must write count glyphs to RunBuffer::glyphs,
count scalars to RunBuffer::pos, textByteCount UTF-8 code units into
RunBuffer::utf8text, and count monotonic indexes into utf8text into
RunBuffer::clusters before next call to SkTextBlobBuilder.

Glyphs share metrics in font.

Glyphs are positioned on a baseline at y, using x-axis positions written by
caller to RunBuffer::pos.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from y, RunBuffer::pos, and RunBuffer::glyphs metrics.

@param font          SkFont used for this run
@param count         number of glyphs
@param y             vertical offset within the blob
@param textByteCount number of UTF-8 code units
@param bounds        optional run bounding box
@return writable glyph buffer, x-axis position buffer, text buffer, and cluster buffer
*/
func (o TextBlobBuilder) AllocRunTextPosH(font Font, count int32, y float32, textByteCount int32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_y := C.float(y)
	c_textByteCount := C.int(textByteCount)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRunTextPosH(c_obj, c_font, c_count, c_y, c_textByteCount, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
Returns run with storage for glyphs, SkPoint positions, text, and
clusters. Caller must write count glyphs to RunBuffer::glyphs, count
SkPoint to RunBuffer::pos, textByteCount UTF-8 code units into
RunBuffer::utf8text, and count monotonic indexes into utf8text into
RunBuffer::clusters before next call to SkTextBlobBuilder.

Glyphs share metrics in font.

Glyphs are positioned using SkPoint written by caller to RunBuffer::pos, using
two scalar values for each SkPoint.

bounds defines an optional bounding box, used to suppress drawing when SkTextBlob
bounds does not intersect SkSurface bounds. If bounds is nullptr, SkTextBlob bounds
is computed from RunBuffer::pos, and RunBuffer::glyphs metrics.

@param font          SkFont used for this run
@param count         number of glyphs
@param textByteCount number of UTF-8 code units
@param bounds        optional run bounding box
@return writable glyph buffer, SkPoint buffer, text buffer, and cluster buffer
*/
func (o TextBlobBuilder) AllocRunTextPos(font Font, count int32, textByteCount int32, bounds *Rect) TextBlobBuilderRunBuffer {
	c_obj := o.sk
	c_font := font.sk
	c_count := C.int(count)
	c_textByteCount := C.int(textByteCount)
	c_bounds := (*C.sk_SkRect)(unsafe.Pointer(bounds))
	retC := C.misk_TextBlobBuilder_allocRunTextPos(c_obj, c_font, c_count, c_textByteCount, c_bounds)
	return TextBlobBuilderRunBuffer{sk: &retC}
}

/*
RunBuffer supplies storage for glyphs and positions within a run.

A run is a sequence of glyphs sharing font metrics and positioning.
Each run may position its glyphs in one of three ways:
by specifying where the first glyph is drawn, and allowing font metrics to
determine the advance to subsequent glyphs; by specifying a baseline, and
the position on that baseline for each glyph in run; or by providing SkPoint
array, one per glyph.
*/
type TextBlobBuilderRunBuffer struct {
	sk *C.sk_SkTextBlobBuilderRunBuffer
}

// IsNil returns true if the raw skia object pointer is nil.
// If it is nil is may indicate that the TextBlobBuilderRunBuffer has not been created.
func (o TextBlobBuilderRunBuffer) IsNil() bool {
	return o.sk == nil
}

func (o TextBlobBuilderRunBuffer) Points() Point {
	c_obj := o.sk
	retC := C.misk_TextBlobBuilderRunBuffer_points(c_obj)
	return Point{sk: retC}
}

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

func (o *Typeface) Unref() {
	C.misk_unref_SkTypeface(o.sk)
	o.sk = nil
}

/*
Returns the typeface's intrinsic style attributes.
*/
func (o Typeface) FontStyle() FontStyle {
	c_obj := o.sk
	retC := C.misk_Typeface_fontStyle(c_obj)
	return FontStyle{sk: &retC}
}

/*
Returns true if style() has the kBold bit set.
*/
func (o Typeface) IsBold() bool {
	c_obj := o.sk
	retC := C.misk_Typeface_isBold(c_obj)
	return bool(retC)
}

/*
Returns true if style() has the kItalic bit set.
*/
func (o Typeface) IsItalic() bool {
	c_obj := o.sk
	retC := C.misk_Typeface_isItalic(c_obj)
	return bool(retC)
}

/*
Returns true if the typeface claims to be fixed-pitch.

	This is a style bit, advance widths may vary even if this returns true.
*/
func (o Typeface) IsFixedPitch() bool {
	c_obj := o.sk
	retC := C.misk_Typeface_isFixedPitch(c_obj)
	return bool(retC)
}

/*
Return a 32bit value for this typeface, unique for the underlying font
data. Will never return 0.
*/
func (o Typeface) UniqueID() uint32 {
	c_obj := o.sk
	retC := C.misk_Typeface_uniqueID(c_obj)
	return uint32(retC)
}

/*
Return a new typeface based on this typeface but parameterized as specified in the
SkFontArguments. If the SkFontArguments does not supply an argument for a parameter
in the font then the value from this typeface will be used as the value for that
argument. If the cloned typeface would be exaclty the same as this typeface then
this typeface may be ref'ed and returned. May return nullptr on failure.
*/
func (o Typeface) MakeClone(p0 FontArguments) Typeface {
	c_obj := o.sk
	c_p0 := p0.sk
	retC := C.misk_Typeface_makeClone(c_obj, c_p0)
	return Typeface{sk: retC}
}

/*
Given an array of UTF32 character codes, return their corresponding glyph IDs.

@param chars pointer to the array of UTF32 chars
@param number of chars and glyphs
@param glyphs returns the corresponding glyph IDs for each character.
*/
func (o Typeface) UnicharsToGlyphs(uni []int32, count int32, glyphs []uint16) {
	c_obj := o.sk
	c_uni := (*C.int)(unsafe.Pointer(&uni[0]))
	c_count := C.int(count)
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	C.misk_Typeface_unicharsToGlyphs(c_obj, c_uni, c_count, c_glyphs)
}

func (o Typeface) TextToGlyphs(text []byte, byteLength uint32, encoding TextEncoding, glyphs []uint16, maxGlyphCount int32) int32 {
	c_obj := o.sk
	c_text := unsafe.Pointer(&text[0])
	c_byteLength := C.ulong(byteLength)
	c_encoding := C.int(encoding)
	c_glyphs := (*C.ushort)(unsafe.Pointer(&glyphs[0]))
	c_maxGlyphCount := C.int(maxGlyphCount)
	retC := C.misk_Typeface_textToGlyphs(c_obj, c_text, c_byteLength, c_encoding, c_glyphs, c_maxGlyphCount)
	return int32(retC)
}

/*
Return the glyphID that corresponds to the specified unicode code-point
(in UTF32 encoding). If the unichar is not supported, returns 0.

This is a short-cut for calling unicharsToGlyphs().
*/
func (o Typeface) UnicharToGlyph(unichar int32) uint16 {
	c_obj := o.sk
	c_unichar := C.int(unichar)
	retC := C.misk_Typeface_unicharToGlyph(c_obj, c_unichar)
	return uint16(retC)
}

/*
Return the number of glyphs in the typeface.
*/
func (o Typeface) CountGlyphs() int32 {
	c_obj := o.sk
	retC := C.misk_Typeface_countGlyphs(c_obj)
	return int32(retC)
}

/*
Return the number of tables in the font.
*/
func (o Typeface) CountTables() int32 {
	c_obj := o.sk
	retC := C.misk_Typeface_countTables(c_obj)
	return int32(retC)
}

/*
Return the units-per-em value for this typeface, or zero if there is an
error.
*/
func (o Typeface) GetUnitsPerEm() int32 {
	c_obj := o.sk
	retC := C.misk_Typeface_getUnitsPerEm(c_obj)
	return int32(retC)
}

/*
Return the family name for this typeface. It will always be returned
encoded as UTF8, but the language of the name is whatever the host
platform chooses.
*/
func (o Typeface) GetFamilyName(name String) {
	c_obj := o.sk
	c_name := name.sk
	C.misk_Typeface_getFamilyName(c_obj, c_name)
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

/*
GPU SkImage and SkSurfaces can be stored such that (0, 0) in texture space may correspond to
either the top-left or bottom-left content pixel.
*/
type GrSurfaceOrigin int32

const (
	GrSurfaceOriginTopLeft    GrSurfaceOrigin = 0
	GrSurfaceOriginBottomLeft GrSurfaceOrigin = 1
)

type GrSyncCpu bool

const (
	GrSyncCpuNo  GrSyncCpu = false
	GrSyncCpuYes GrSyncCpu = true
)

/*
Enum used as return value when flush with semaphores so the client knows whether the valid
semaphores will be submitted on the next GrContext::submit call.
*/
type GrSemaphoresSubmitted bool

const (
	GrSemaphoresSubmittedNo  GrSemaphoresSubmitted = false
	GrSemaphoresSubmittedYes GrSemaphoresSubmitted = true
)

/*
\enum SkAlphaType
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
type AlphaType int32

const (
	// uninitialized
	AlphaTypeUnknown AlphaType = 0
	// pixel is opaque
	AlphaTypeOpaque AlphaType = 1
	// pixel components are premultiplied by alpha
	AlphaTypePremul AlphaType = 2
	// pixel components are independent of alpha
	AlphaTypeUnpremul AlphaType = 3
	// last valid value
	AlphaTypeLastEnum AlphaType = 3
)

/*
When we transform points through a matrix containing perspective (the bottom row is something
other than 0,0,1), the bruteforce math can produce confusing results (since we might divide
by 0, or a negative w value). By default, methods that map rects and paths will apply
perspective clipping, but this can be changed by specifying kYes to those methods.
*/
type ApplyPerspectiveClip int32

const (
	// Don't pre-clip the geometry before applying the (perspective) matrix
	ApplyPerspectiveClipNo ApplyPerspectiveClip = 0
	// Do pre-clip the geometry before applying the (perspective) matrix
	ApplyPerspectiveClipYes ApplyPerspectiveClip = 1
)

type BlendMode int32

const ()

type ClipOp int32

const (
	ClipOpDifference    ClipOp = 0
	ClipOpIntersect     ClipOp = 1
	ClipOpMax_EnumValue ClipOp = 1
)

/*
\enum SkColorType
Describes how pixel bits encode color. A pixel may be an alpha mask, a grayscale, RGB, or ARGB.

kN32_SkColorType selects the native 32-bit ARGB format for the current configuration. This can
lead to inconsistent results across platforms, so use with caution.
*/
type ColorType int32

const (
	// uninitialized
	ColorTypeUnknown ColorType = 0
	// pixel with alpha in 8-bit byte
	ColorTypeAlpha_8 ColorType = 1
	// pixel with 5 bits red, 6 bits green, 5 bits blue, in 16-bit word
	ColorTypeRGB_565 ColorType = 2
	// pixel with 4 bits for alpha, red, green, blue; in 16-bit word
	ColorTypeARGB_4444 ColorType = 3
	// pixel with 8 bits for red, green, blue, alpha; in 32-bit word
	ColorTypeRGBA_8888 ColorType = 4
	// pixel with 8 bits each for red, green, blue; in 32-bit word
	ColorTypeRGB_888x ColorType = 5
	// pixel with 8 bits for blue, green, red, alpha; in 32-bit word
	ColorTypeBGRA_8888 ColorType = 6
	// 10 bits for red, green, blue; 2 bits for alpha; in 32-bit word
	ColorTypeRGBA_1010102 ColorType = 7
	// 10 bits for blue, green, red; 2 bits for alpha; in 32-bit word
	ColorTypeBGRA_1010102 ColorType = 8
	// pixel with 10 bits each for red, green, blue; in 32-bit word
	ColorTypeRGB_101010x ColorType = 9
	// pixel with 10 bits each for blue, green, red; in 32-bit word
	ColorTypeBGR_101010x ColorType = 10
	// pixel with 10 bits each for blue, green, red; in 32-bit word, extended range
	ColorTypeBGR_101010x_XR ColorType = 11
	// pixel with 10 bits each for blue, green, red, alpha; in 64-bit word, extended range
	ColorTypeBGRA_10101010_XR ColorType = 12
	// pixel with 10 used bits (most significant) followed by 6 unused
	ColorTypeRGBA_10x6 ColorType = 13
	// pixel with grayscale level in 8-bit byte
	ColorTypeGray_8 ColorType = 14
	// pixel with half floats in [0,1] for red, green, blue, alpha;
	ColorTypeRGBA_F16Norm ColorType = 15
	// pixel with half floats for red, green, blue, alpha;
	ColorTypeRGBA_F16 ColorType = 16
	// pixel using C float for red, green, blue, alpha; in 128-bit word
	ColorTypeRGBA_F32 ColorType = 17
	// pixel with a uint8_t for red and green
	ColorTypeR8G8_unorm ColorType = 18
	// pixel with a half float for alpha
	ColorTypeA16_float ColorType = 19
	// pixel with a half float for red and green
	ColorTypeR16G16_float ColorType = 20
	// pixel with a little endian uint16_t for alpha
	ColorTypeA16_unorm ColorType = 21
	// pixel with a little endian uint16_t for red and green
	ColorTypeR16G16_unorm ColorType = 22
	// pixel with a little endian uint16_t for red, green, blue
	ColorTypeR16G16B16A16_unorm ColorType = 23
	ColorTypeSRGBA_8888         ColorType = 24
	ColorTypeR8_unorm           ColorType = 25
	// last valid value
	ColorTypeLastEnum ColorType = 25
	// native 32-bit RGBA encoding
	ColorTypeN32 ColorType = 4
)

type FontHinting int32

const (
	// glyph outlines unchanged
	FontHintingNone FontHinting = 0
	// minimal modification to improve constrast
	FontHintingSlight FontHinting = 1
	// glyph outlines modified to improve constrast
	FontHintingNormal FontHinting = 2
	// modifies glyph outlines for maximum constrast
	FontHintingFull FontHinting = 3
)

type FilterMode int32

const (
	FilterModeNearest FilterMode = 0
	FilterModeLinear  FilterMode = 1
	FilterModeLast    FilterMode = 1
)

type PathDirection int32

const (
	/*
	   clockwise direction for adding closed contours
	*/
	PathDirectionCW PathDirection = 0
	/*
	   counter-clockwise direction for adding closed contours
	*/
	PathDirectionCCW PathDirection = 1
)

type PathFillType int32

const (
	/*
	   Specifies that "inside" is computed by a non-zero sum of signed edge crossings
	*/
	PathFillTypeWinding PathFillType = 0
	/*
	   Specifies that "inside" is computed by an odd number of edge crossings
	*/
	PathFillTypeEvenOdd PathFillType = 1
	/*
	   Same as Winding, but draws outside of the path, rather than inside
	*/
	PathFillTypeInverseWinding PathFillType = 2
	/*
	   Same as EvenOdd, but draws outside of the path, rather than inside
	*/
	PathFillTypeInverseEvenOdd PathFillType = 3
)

/*
The logical operations that can be performed when combining two paths.
*/
type PathOp uint32

const (
	// subtract the op path from the first path
	PathOpDifference PathOp = 0
	// intersect the two paths
	PathOpIntersect PathOp = 1
	// union (inclusive-or) the two paths
	PathOpUnion PathOp = 2
	// exclusive-or the two paths
	PathOpXOR PathOp = 3
	// subtract the first path from the op path
	PathOpReverseDifference PathOp = 4
)

/*
Description of how the LCD strips are arranged for each pixel. If this is unknown, or the
pixels are meant to be "portable" and/or transformed before showing (e.g. rotated, scaled)
then use kUnknown_SkPixelGeometry.
*/
type PixelGeometry uint32

const (
	PixelGeometryUnknown PixelGeometry = 0
	PixelGeometryRGB_H   PixelGeometry = 1
	PixelGeometryBGR_H   PixelGeometry = 2
	PixelGeometryRGB_V   PixelGeometry = 3
	PixelGeometryBGR_V   PixelGeometry = 4
)

type TextEncoding int32

const (
	// uses bytes to represent UTF-8 or ASCII
	TextEncodingUTF8 TextEncoding = 0
	// uses two byte words to represent most of Unicode
	TextEncodingUTF16 TextEncoding = 1
	// uses four byte words to represent all of Unicode
	TextEncodingUTF32 TextEncoding = 2
	// uses two byte words to represent glyph indices
	TextEncodingGlyphID TextEncoding = 3
)

/*
Is the data protected on the GPU or not.
*/
type SkgpuProtected bool

const (
	SkgpuProtectedNo  SkgpuProtected = false
	SkgpuProtectedYes SkgpuProtected = true
)

func GrBackendRenderTargetsMakeGL(width int32, height int32, sampleCnt int32, stencilBits int32, glInfo GrGLFramebufferInfo) GrBackendRenderTarget {
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
func SkColorSetARGB(a uint32, r uint32, g uint32, b uint32) Color {
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
func SkColorSetA(c Color, a uint32) Color {
	c_c := C.uint(c)
	c_a := C.uint(a)
	retC := C.misk_SkColorSetA(c_c, c_a)
	return Color(retC)
}

/*
	Return a SkImage using the encoded data, but attempts to defer decoding until the
	image is actually used/drawn. This deferral allows the system to cache the result, either on the
	CPU or on the GPU, depending on where the image is drawn. If memory is low, the cache may
	be purged, causing the next draw of the image to have to re-decode.

	If alphaType is nullopt, the image's alpha type will be chosen automatically based on the
	image format. Transparent images will default to kPremul_SkAlphaType. If alphaType contains
	kPremul_SkAlphaType or kUnpremul_SkAlphaType, that alpha type will be used. Forcing opaque
	(passing kOpaque_SkAlphaType) is not allowed, and will return nullptr.

	If the encoded format is not supported, nullptr is returned.

	If possible, clients should use SkCodecs::DeferredImage instead.

	@param encoded  the encoded data
	@return         created SkImage, or nullptr

example: https://fiddle.skia.org/c/@Image_DeferredFromEncodedData
*/
func SkImagesDeferredFromEncodedData(encoded Data, alphaType *AlphaType) Image {
	c_encoded := encoded.sk
	c_alphaType := (*C.int)(alphaType)
	retC := C.misk_SkImagesDeferredFromEncodedData(c_encoded, c_alphaType)
	return Image{sk: retC}
}

/*
Creates CPU-backed SkImage from pixel data described by info.
The pixels data will *not* be copied.

SkImage is returned if SkImageInfo is valid. Valid SkImageInfo parameters include:
dimensions are greater than zero;
each dimension fits in 29 bits;
SkColorType and SkAlphaType are valid, and SkColorType is not kUnknown_SkColorType;
rowBytes are large enough to hold one row of pixels;
pixels is not nullptr, and contains enough data for SkImage.

@param info      contains width, height, SkAlphaType, SkColorType, SkColorSpace
@param pixels    address or pixel storage
@param rowBytes  size of pixel row or larger
@return          SkImage sharing pixels, or nullptr
*/
func SkImagesRasterFromData(info ImageInfo, pixels Data, rowBytes uint32) Image {
	c_info := info.sk
	c_pixels := pixels.sk
	c_rowBytes := C.ulong(rowBytes)
	retC := C.misk_SkImagesRasterFromData(c_info, c_pixels, c_rowBytes)
	return Image{sk: retC}
}

/*
Returns a SkPMColor value from unpremultiplied 8-bit component values.

@param a  amount of alpha, from fully transparent (0) to fully opaque (255)
@param r  amount of red, from no red (0) to full red (255)
@param g  amount of green, from no green (0) to full green (255)
@param b  amount of blue, from no blue (0) to full blue (255)
@return   premultiplied color
*/
func SkPreMultiplyARGB(a uint32, r uint32, g uint32, b uint32) PMColor {
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
Wraps a GPU-backed buffer into SkSurface. Caller must ensure backendRenderTarget
is valid for the lifetime of returned SkSurface.

SkSurface is returned if all parameters are valid. backendRenderTarget is valid if
its pixel configuration agrees with colorSpace and context; for instance, if
backendRenderTarget has an sRGB configuration, then context must support sRGB,
and colorSpace must be present. Further, backendRenderTarget width and height must
not exceed context capabilities, and the context must be able to support
back-end render targets.

Upon success releaseProc is called when it is safe to delete the render target in the
backend API (accounting only for use of the render target by this surface). If SkSurface
creation fails releaseProc is called before this function returns.

@param context                  GPU context
@param backendRenderTarget      GPU intermediate memory buffer
@param colorSpace               range of colors
@param surfaceProps             LCD striping orientation and setting for device independent
fonts; may be nullptr
@param releaseProc              function called when backendRenderTarget can be released
@param releaseContext           state passed to releaseProc
@return                         SkSurface if all parameters are valid; otherwise, nullptr
*/
func SkSurfacesWrapBackendRenderTarget(context GrRecordingContext, backendRenderTarget GrBackendRenderTarget, origin GrSurfaceOrigin, colorType ColorType, colorSpace ColorSpace, surfaceProps SurfaceProps) Surface {
	c_context := context.sk
	c_backendRenderTarget := backendRenderTarget.sk
	c_origin := C.int(origin)
	c_colorType := C.int(colorType)
	c_colorSpace := colorSpace.sk
	c_surfaceProps := surfaceProps.sk
	retC := C.misk_SkSurfacesWrapBackendRenderTarget(c_context, c_backendRenderTarget, c_origin, c_colorType, c_colorSpace, c_surfaceProps)
	return Surface{sk: retC}
}

/*
Set this path to the result of applying the Op to this path and the
specified path: this = (this op operand).
The resulting path will be constructed from non-overlapping contours.
The curve order is reduced where possible so that cubics may be turned
into quadratics, and quadratics maybe turned into lines.

Returns true if operation was able to produce a result;
otherwise, result is unmodified.

@param one The first operand (for difference, the minuend)
@param two The second operand (for difference, the subtrahend)
@param op The operator to apply.
@param result The product of the operands. The result may be one of the
inputs.
@return True if the operation succeeded.
*/
func Op(one Path, two Path, op PathOp, result Path) bool {
	c_one := one.sk
	c_two := two.sk
	c_op := C.uint(op)
	c_result := result.sk
	retC := C.misk_Op(c_one, c_two, c_op, c_result)
	return bool(retC)
}

/*
Set this path to a set of non-overlapping contours that describe the
same area as the original path.
The curve order is reduced where possible so that cubics may
be turned into quadratics, and quadratics maybe turned into lines.

Returns true if operation was able to produce a result;
otherwise, result is unmodified.

@param path The path to simplify.
@param result The simplified path. The result may be the input.
@return True if simplification succeeded.
*/
func Simplify(path Path, result Path) bool {
	c_path := path.sk
	c_result := result.sk
	retC := C.misk_Simplify(c_path, c_result)
	return bool(retC)
}

/*
Set the resulting rectangle to the tight bounds of the path.

@param path The path measured.
@param result The tight bounds of the path.
@return True if the bounds could be computed.
*/
func TightBounds(path Path, result Rect) bool {
	c_path := path.sk
	c_result := *(*C.sk_SkRect)(unsafe.Pointer(&result))
	retC := C.misk_TightBounds(c_path, c_result)
	return bool(retC)
}

/*
Set the result with fill type winding to area equivalent to path.
Returns true if successful. Does not detect if path contains contours which
contain self-crossings or cross other contours; in these cases, may return
true even though result does not fill same area as path.

Returns true if operation was able to produce a result;
otherwise, result is unmodified. The result may be the input.

@param path The path typically with fill type set to even odd.
@param result The equivalent path with fill type set to winding.
@return True if winding path was set.
*/
func AsWinding(path Path, result Path) bool {
	c_path := path.sk
	c_result := result.sk
	retC := C.misk_AsWinding(c_path, c_result)
	return bool(retC)
}

/*
8-bit type for an alpha value. 255 is 100% opaque, zero is 100% transparent.
*/
type Alpha C.uchar

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

/*
Represents fully opaque SkAlpha value. SkAlpha ranges from zero,
fully transparent; to 255, fully opaque.
*/
var AlphaOPAQUE = (Alpha)(C.sk_SK_AlphaOPAQUE)

/*
Represents fully transparent SkAlpha value. SkAlpha ranges from zero,
fully transparent; to 255, fully opaque.
*/
var AlphaTRANSPARENT = (Alpha)(C.sk_SK_AlphaTRANSPARENT)

/*
Represents fully opaque black.
*/
var ColorBLACK = (Color)(C.sk_SK_ColorBLACK)

/*
Represents fully opaque blue.
*/
var ColorBLUE = (Color)(C.sk_SK_ColorBLUE)

/*
Represents fully opaque cyan. HTML aqua is equivalent.
*/
var ColorCYAN = (Color)(C.sk_SK_ColorCYAN)

/*
Represents fully opaque dark gray.
Note that SVG dark gray is equivalent to 0xFFA9A9A9.
*/
var ColorDKGRAY = (Color)(C.sk_SK_ColorDKGRAY)

/*
Represents fully opaque gray.
Note that HTML gray is equivalent to 0xFF808080.
*/
var ColorGRAY = (Color)(C.sk_SK_ColorGRAY)

/*
Represents fully opaque green. HTML lime is equivalent.
Note that HTML green is equivalent to 0xFF008000.
*/
var ColorGREEN = (Color)(C.sk_SK_ColorGREEN)

/*
Represents fully opaque light gray. HTML silver is equivalent to 0xFFC0C0C0.
Note that SVG light gray is equivalent to 0xFFD3D3D3.
*/
var ColorLTGRAY = (Color)(C.sk_SK_ColorLTGRAY)

/*
Represents fully opaque magenta. HTML fuchsia is equivalent.
*/
var ColorMAGENTA = (Color)(C.sk_SK_ColorMAGENTA)

/*
Represents fully opaque red.
*/
var ColorRED = (Color)(C.sk_SK_ColorRED)

/*
Represents fully transparent SkColor. May be used to initialize a destination
containing a mask or a non-rectangular image.
*/
var ColorTRANSPARENT = (Color)(C.sk_SK_ColorTRANSPARENT)

/*
Represents fully opaque white.
*/
var ColorWHITE = (Color)(C.sk_SK_ColorWHITE)

/*
Represents fully opaque yellow.
*/
var ColorYELLOW = (Color)(C.sk_SK_ColorYELLOW)

/*
The generation IDs in Skia reserve 0 has an invalid marker.
*/
var InvalidGenID = (uint32)(C.sk_SK_InvalidGenID)

/*
The unique IDs in Skia reserve 0 has an invalid marker.
*/
var InvalidUniqueID = (uint32)(C.sk_SK_InvalidUniqueID)

/*
Maximum representable milliseconds; 24d 20h 31m 23.647s.
*/
var MSecMax = (uint32)(C.sk_SK_MSecMax)

/*
This value translates to reseting all the context state for any backend.
*/
var All_GrBackendState = (uint32)(C.sk_kAll_GrBackendState)
var GrGLStandardCnt = (int32)(C.sk_kGrGLStandardCnt)
var ColorTypeCnt = (int32)(C.sk_kSkColorTypeCnt)
var FilterModeCount = (int32)(C.sk_kSkFilterModeCount)
var MipmapModeCount = (int32)(C.sk_kSkMipmapModeCount)
var StrAppendS32_MaxSize = (int32)(C.sk_kSkStrAppendS32_MaxSize)
var StrAppendS64_MaxSize = (int32)(C.sk_kSkStrAppendS64_MaxSize)

/*
Floats have at most 8 significant digits, so we limit our %g to that.
However, the total string could be 15 characters: -1.2345678e-005

In theory we should only expect up to 2 digits for the exponent, but on
some platforms we have seen 3 (as in the example above).
*/
var StrAppendScalar_MaxSize = (int32)(C.sk_kSkStrAppendScalar_MaxSize)
var StrAppendU32_MaxSize = (int32)(C.sk_kSkStrAppendU32_MaxSize)
var StrAppendU64_MaxSize = (int32)(C.sk_kSkStrAppendU64_MaxSize)

// FontMgrDefault returns a FontMgr, that can be used to get Typefaces.
func FontMgrDefault() FontMgr {
	return FontMgr{
		sk: C.sk_fontmgr_ref_default(),
	}
}
