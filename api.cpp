// This is a generated file. DO NOT EDIT.

#include "skia.h"

extern "C" {
#include "api.h"

void *skia_new_SkMatrix() { return reinterpret_cast<void *>(new SkMatrix()); }

void *skia_new_SkRefCntBase() {
  return reinterpret_cast<void *>(new SkRefCntBase());
}

void skia_delete_SkRefCntBase(void *obj) {
  delete reinterpret_cast<SkRefCntBase *>(obj);
}

void *skia_new_SkColorInfo() {
  return reinterpret_cast<void *>(new SkColorInfo());
}

void *skia_new_SkColorInfo3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkColorInfo(reinterpret_cast<SkColorInfo &>(c_p0)));
}

void skia_delete_SkColorInfo(void *obj) {
  delete reinterpret_cast<SkColorInfo *>(obj);
}

void *skia_new_SkPixmap() { return reinterpret_cast<void *>(new SkPixmap()); }

void *skia_new_SkPixmap2(void *c_info, void c_addr, ulong c_rowBytes) {
  return reinterpret_cast<void *>(new SkPixmap(
      reinterpret_cast<SkImageInfo &>(c_info), c_addr, c_rowBytes));
}

void *skia_new_SkYUVAInfo() {
  return reinterpret_cast<void *>(new SkYUVAInfo());
}

void *skia_new_SkYUVAInfo2(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkYUVAInfo(reinterpret_cast<SkYUVAInfo &>(c_p0)));
}

void *skia_new_SkYUVAPixmapInfo() {
  return reinterpret_cast<void *>(new SkYUVAPixmapInfo());
}

void *skia_new_SkYUVAPixmapInfo2(void *c_p0, int c_p1, ulong c_rowBytes) {
  return reinterpret_cast<void *>(new SkYUVAPixmapInfo(
      reinterpret_cast<SkYUVAInfo &>(c_p0), c_p1, c_rowBytes));
}

void *skia_new_SkYUVAPixmapInfo3(void *c_p0, int c_p1, ulong c_rowBytes) {
  return reinterpret_cast<void *>(
      new SkYUVAPixmapInfo(reinterpret_cast<SkYUVAInfo &>(c_p0),
                           (SkYUVAPixmapInfo::DataType)c_p1, c_rowBytes));
}

void *skia_new_SkYUVAPixmapInfo4(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkYUVAPixmapInfo(reinterpret_cast<SkYUVAPixmapInfo &>(c_p0)));
}

void *skia_new_SkYUVAPixmaps() {
  return reinterpret_cast<void *>(new SkYUVAPixmaps());
}

void *skia_new_SkYUVAPixmaps3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkYUVAPixmaps(reinterpret_cast<SkYUVAPixmaps &>(c_p0)));
}

void skia_delete_SkYUVAPixmaps(void *obj) {
  delete reinterpret_cast<SkYUVAPixmaps *>(obj);
}

void *skia_new_SkOnce() { return reinterpret_cast<void *>(new SkOnce()); }

void *skia_new_SkNoncopyable() {
  return reinterpret_cast<void *>(new SkNoncopyable());
}

void skia_delete_SkCodec(void *obj) { delete reinterpret_cast<SkCodec *>(obj); }

void *skia_new_SkM44(void *c_src) {
  return reinterpret_cast<void *>(new SkM44(reinterpret_cast<SkM44 &>(c_src)));
}

void *skia_new_SkM442() { return reinterpret_cast<void *>(new SkM44()); }

void *skia_new_SkM443(void *c_a, void *c_b) {
  return reinterpret_cast<void *>(new SkM44(reinterpret_cast<SkM44 &>(c_a),
                                            reinterpret_cast<SkM44 &>(c_b)));
}

void *skia_new_SkM444(int c_p0) {
  return reinterpret_cast<void *>(
      new SkM44((SkM44::Uninitialized_Constructor)c_p0));
}

void *skia_new_SkM445(int c_p0) {
  return reinterpret_cast<void *>(new SkM44((SkM44::NaN_Constructor)c_p0));
}

void *skia_new_SkM446(float c_m0, float c_m4, float c_m8, float c_m12,
                      float c_m1, float c_m5, float c_m9, float c_m13,
                      float c_m2, float c_m6, float c_m10, float c_m14,
                      float c_m3, float c_m7, float c_m11, float c_m15) {
  return reinterpret_cast<void *>(new SkM44(c_m0, c_m4, c_m8, c_m12, c_m1, c_m5,
                                            c_m9, c_m13, c_m2, c_m6, c_m10,
                                            c_m14, c_m3, c_m7, c_m11, c_m15));
}

void *skia_new_SkM447(void *c_src) {
  return reinterpret_cast<void *>(
      new SkM44(reinterpret_cast<SkMatrix &>(c_src)));
}

void *skia_new_SkPaint() { return reinterpret_cast<void *>(new SkPaint()); }

void *skia_new_SkPaint3(void *c_paint) {
  return reinterpret_cast<void *>(
      new SkPaint(reinterpret_cast<SkPaint &>(c_paint)));
}

void skia_delete_SkPaint(void *obj) { delete reinterpret_cast<SkPaint *>(obj); }

void skia_delete_SkRasterHandleAllocator(void *obj) {
  delete reinterpret_cast<SkRasterHandleAllocator *>(obj);
}

void *skia_new_SkString() { return reinterpret_cast<void *>(new SkString()); }

void *skia_new_SkString2(ulong c_len) {
  return reinterpret_cast<void *>(new SkString(c_len));
}

void *skia_new_SkString3(schar c_text) {
  return reinterpret_cast<void *>(new SkString(c_text));
}

void *skia_new_SkString4(schar c_text, ulong c_len) {
  return reinterpret_cast<void *>(new SkString(c_text, c_len));
}

void *skia_new_SkString5(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkString(reinterpret_cast<SkString &>(c_p0)));
}

void *skia_new_SkString7(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkString(reinterpret_cast<std::string &>(c_p0)));
}

void skia_delete_SkString(void *obj) {
  delete reinterpret_cast<SkString *>(obj);
}

void *skia_new_SkSurfaceProps() {
  return reinterpret_cast<void *>(new SkSurfaceProps());
}

void *skia_new_SkSurfaceProps2(uint c_flags, int c_p1) {
  return reinterpret_cast<void *>(
      new SkSurfaceProps(c_flags, (SkPixelGeometry)c_p1));
}

void *skia_new_SkSurfaceProps3(uint c_flags, int c_p1, float c_textContrast,
                               float c_textGamma) {
  return reinterpret_cast<void *>(new SkSurfaceProps(
      c_flags, (SkPixelGeometry)c_p1, c_textContrast, c_textGamma));
}

void *skia_new_SkSurfaceProps4(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSurfaceProps(reinterpret_cast<SkSurfaceProps &>(c_p0)));
}

void *skia_new_SkDeque(ulong c_elemSize, int c_allocCount) {
  return reinterpret_cast<void *>(new SkDeque(c_elemSize, c_allocCount));
}

void *skia_new_SkDeque2(ulong c_elemSize, void c_storage, ulong c_storageSize,
                        int c_allocCount) {
  return reinterpret_cast<void *>(
      new SkDeque(c_elemSize, c_storage, c_storageSize, c_allocCount));
}

void skia_delete_SkDeque(void *obj) { delete reinterpret_cast<SkDeque *>(obj); }

void *skia_new_SkContainerAllocator(ulong c_sizeOfT, int c_maxCapacity) {
  return reinterpret_cast<void *>(
      new SkContainerAllocator(c_sizeOfT, c_maxCapacity));
}

void *skia_new_SkCanvas() { return reinterpret_cast<void *>(new SkCanvas()); }

void *skia_new_SkCanvas4(void *c_bitmap) {
  return reinterpret_cast<void *>(
      new SkCanvas(reinterpret_cast<SkBitmap &>(c_bitmap)));
}

void *skia_new_SkCanvas5(void *c_bitmap, void *c_props) {
  return reinterpret_cast<void *>(
      new SkCanvas(reinterpret_cast<SkBitmap &>(c_bitmap),
                   reinterpret_cast<SkSurfaceProps &>(c_props)));
}

void skia_delete_SkCanvas(void *obj) {
  delete reinterpret_cast<SkCanvas *>(obj);
}

void skia_delete_SkAutoCanvasRestore(void *obj) {
  delete reinterpret_cast<SkAutoCanvasRestore *>(obj);
}

void skia_delete_SkStream(void *obj) {
  delete reinterpret_cast<SkStream *>(obj);
}

void skia_delete_SkWStream(void *obj) {
  delete reinterpret_cast<SkWStream *>(obj);
}

void *skia_new_SkFILEStream(schar c_path) {
  return reinterpret_cast<void *>(new SkFILEStream(c_path));
}

void skia_delete_SkFILEStream(void *obj) {
  delete reinterpret_cast<SkFILEStream *>(obj);
}

void *skia_new_SkMemoryStream() {
  return reinterpret_cast<void *>(new SkMemoryStream());
}

void *skia_new_SkMemoryStream2(ulong c_length) {
  return reinterpret_cast<void *>(new SkMemoryStream(c_length));
}

void *skia_new_SkMemoryStream3(void c_data, ulong c_length, c_copyData) {
  return reinterpret_cast<void *>(
      new SkMemoryStream(c_data, c_length, c_copyData));
}

void *skia_new_SkFILEWStream(schar c_path) {
  return reinterpret_cast<void *>(new SkFILEWStream(c_path));
}

void skia_delete_SkFILEWStream(void *obj) {
  delete reinterpret_cast<SkFILEWStream *>(obj);
}

void *skia_new_SkDynamicMemoryWStream() {
  return reinterpret_cast<void *>(new SkDynamicMemoryWStream());
}

void skia_delete_SkDynamicMemoryWStream(void *obj) {
  delete reinterpret_cast<SkDynamicMemoryWStream *>(obj);
}

void *skia_new_SkFontStyle(int c_weight, int c_width, int c_slant) {
  return reinterpret_cast<void *>(
      new SkFontStyle(c_weight, c_width, (SkFontStyle::Slant)c_slant));
}

void *skia_new_SkFontStyle2() {
  return reinterpret_cast<void *>(new SkFontStyle());
}

void *skia_new_SkWeakRefCnt() {
  return reinterpret_cast<void *>(new SkWeakRefCnt());
}

void skia_delete_SkWeakRefCnt(void *obj) {
  delete reinterpret_cast<SkWeakRefCnt *>(obj);
}

void *skia_new_SkFont() { return reinterpret_cast<void *>(new SkFont()); }

void *skia_new_SkTextBlobBuilder() {
  return reinterpret_cast<void *>(new SkTextBlobBuilder());
}

void skia_delete_SkTextBlobBuilder(void *obj) {
  delete reinterpret_cast<SkTextBlobBuilder *>(obj);
}

void *skia_new_SkPath() { return reinterpret_cast<void *>(new SkPath()); }

void *skia_new_SkPath2(void *c_path) {
  return reinterpret_cast<void *>(
      new SkPath(reinterpret_cast<SkPath &>(c_path)));
}

void skia_delete_SkPath(void *obj) { delete reinterpret_cast<SkPath *>(obj); }

void skia_delete_SkPicture(void *obj) {
  delete reinterpret_cast<SkPicture *>(obj);
}

void skia_delete_SkEncoder(void *obj) {
  delete reinterpret_cast<SkEncoder *>(obj);
}

void *skia_new_GrGLExtensions() {
  return reinterpret_cast<void *>(new GrGLExtensions());
}

void *skia_new_GrGLExtensions2(void *c_p0) {
  return reinterpret_cast<void *>(
      new GrGLExtensions(reinterpret_cast<GrGLExtensions &>(c_p0)));
}

void *skia_new_GrBackendFormat() {
  return reinterpret_cast<void *>(new GrBackendFormat());
}

void *skia_new_GrBackendFormat2(void *c_p0) {
  return reinterpret_cast<void *>(
      new GrBackendFormat(reinterpret_cast<GrBackendFormat &>(c_p0)));
}

void skia_delete_GrBackendFormat(void *obj) {
  delete reinterpret_cast<GrBackendFormat *>(obj);
}

void *skia_new_GrBackendTexture() {
  return reinterpret_cast<void *>(new GrBackendTexture());
}

void *skia_new_GrBackendTexture3(void *c_that) {
  return reinterpret_cast<void *>(
      new GrBackendTexture(reinterpret_cast<GrBackendTexture &>(c_that)));
}

void skia_delete_GrBackendTexture(void *obj) {
  delete reinterpret_cast<GrBackendTexture *>(obj);
}

void *skia_new_GrBackendRenderTarget() {
  return reinterpret_cast<void *>(new GrBackendRenderTarget());
}

void *skia_new_GrBackendRenderTarget2(int c_width, int c_height,
                                      int c_sampleCnt, int c_stencilBits,
                                      void *c_mockInfo) {
  return reinterpret_cast<void *>(new GrBackendRenderTarget(
      c_width, c_height, c_sampleCnt, c_stencilBits,
      reinterpret_cast<GrMockRenderTargetInfo &>(c_mockInfo)));
}

void *skia_new_GrBackendRenderTarget3(void *c_that) {
  return reinterpret_cast<void *>(new GrBackendRenderTarget(
      reinterpret_cast<GrBackendRenderTarget &>(c_that)));
}

void skia_delete_GrBackendRenderTarget(void *obj) {
  delete reinterpret_cast<GrBackendRenderTarget *>(obj);
}

void *skia_new_GrDriverBugWorkarounds() {
  return reinterpret_cast<void *>(new GrDriverBugWorkarounds());
}

void *skia_new_GrDriverBugWorkarounds2(void *c_p0) {
  return reinterpret_cast<void *>(new GrDriverBugWorkarounds(
      reinterpret_cast<GrDriverBugWorkarounds &>(c_p0)));
}

void *skia_new_GrDriverBugWorkarounds3(void *c_workarounds) {
  return reinterpret_cast<void *>(new GrDriverBugWorkarounds(
      reinterpret_cast<std::vector<int32_t> &>(c_workarounds)));
}

void skia_delete_GrDriverBugWorkarounds(void *obj) {
  delete reinterpret_cast<GrDriverBugWorkarounds *>(obj);
}

void skia_delete_GrContext_Base(void *obj) {
  delete reinterpret_cast<GrContext_Base *>(obj);
}

void *skia_new_SkSemaphore(int c_count) {
  return reinterpret_cast<void *>(new SkSemaphore(c_count));
}

void skia_delete_SkSemaphore(void *obj) {
  delete reinterpret_cast<SkSemaphore *>(obj);
}

void *skia_new_SkMutex() { return reinterpret_cast<void *>(new SkMutex()); }

void skia_delete_SkMutex(void *obj) { delete reinterpret_cast<SkMutex *>(obj); }

void *skia_new_SkAutoMutexExclusive(void *c_mutex) {
  return reinterpret_cast<void *>(
      new SkAutoMutexExclusive(reinterpret_cast<SkMutex &>(c_mutex)));
}

void *skia_new_SkAutoMutexExclusive2(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkAutoMutexExclusive(reinterpret_cast<SkAutoMutexExclusive &>(c_p0)));
}

void skia_delete_SkAutoMutexExclusive(void *obj) {
  delete reinterpret_cast<SkAutoMutexExclusive *>(obj);
}

void skia_delete_GrImageContext(void *obj) {
  delete reinterpret_cast<GrImageContext *>(obj);
}

void skia_delete_GrRecordingContext(void *obj) {
  delete reinterpret_cast<GrRecordingContext *>(obj);
}

void skia_delete_GrDirectContext(void *obj) {
  delete reinterpret_cast<GrDirectContext *>(obj);
}

void *skia_new_SkTDStorage(int c_sizeOfT) {
  return reinterpret_cast<void *>(new SkTDStorage(c_sizeOfT));
}

void *skia_new_SkTDStorage2(void c_src, int c_size, int c_sizeOfT) {
  return reinterpret_cast<void *>(new SkTDStorage(c_src, c_size, c_sizeOfT));
}

void *skia_new_SkTDStorage3(void *c_that) {
  return reinterpret_cast<void *>(
      new SkTDStorage(reinterpret_cast<SkTDStorage &>(c_that)));
}

void skia_delete_SkTDStorage(void *obj) {
  delete reinterpret_cast<SkTDStorage *>(obj);
}

void skia_delete_SkPDFUnion(void *obj) {
  delete reinterpret_cast<SkPDFUnion *>(obj);
}

void skia_delete_SkPDFObject(void *obj) {
  delete reinterpret_cast<SkPDFObject *>(obj);
}

void *skia_new_SkPDFArray() {
  return reinterpret_cast<void *>(new SkPDFArray());
}

void skia_delete_SkPDFArray(void *obj) {
  delete reinterpret_cast<SkPDFArray *>(obj);
}

void *skia_new_SkPDFDict(schar c_type) {
  return reinterpret_cast<void *>(new SkPDFDict(c_type));
}

void skia_delete_SkPDFDict(void *obj) {
  delete reinterpret_cast<SkPDFDict *>(obj);
}

void *skia_new_SkArenaAlloc(schar c_block, ulong c_blockSize,
                            ulong c_firstHeapAllocation) {
  return reinterpret_cast<void *>(
      new SkArenaAlloc(c_block, c_blockSize, c_firstHeapAllocation));
}

void *skia_new_SkArenaAlloc2(ulong c_firstHeapAllocation) {
  return reinterpret_cast<void *>(new SkArenaAlloc(c_firstHeapAllocation));
}

void *skia_new_SkArenaAlloc3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkArenaAlloc(reinterpret_cast<SkArenaAlloc &>(c_p0)));
}

void skia_delete_SkArenaAlloc(void *obj) {
  delete reinterpret_cast<SkArenaAlloc *>(obj);
}

void *skia_new_SkArenaAllocWithReset(schar c_block, ulong c_blockSize,
                                     ulong c_firstHeapAllocation) {
  return reinterpret_cast<void *>(
      new SkArenaAllocWithReset(c_block, c_blockSize, c_firstHeapAllocation));
}

void *skia_new_SkArenaAllocWithReset2(ulong c_firstHeapAllocation) {
  return reinterpret_cast<void *>(
      new SkArenaAllocWithReset(c_firstHeapAllocation));
}

void *skia_new_SkPDFTagTree() {
  return reinterpret_cast<void *>(new SkPDFTagTree());
}

void skia_delete_SkPDFTagTree(void *obj) {
  delete reinterpret_cast<SkPDFTagTree *>(obj);
}

void skia_delete_SkPDFDocument(void *obj) {
  delete reinterpret_cast<SkPDFDocument *>(obj);
}

void skia_delete_SkShaper(void *obj) {
  delete reinterpret_cast<SkShaper *>(obj);
}

void *skia_new_SkSVGLength() {
  return reinterpret_cast<void *>(new SkSVGLength());
}

void *skia_new_SkSVGLength2(float c_v, int c_u) {
  return reinterpret_cast<void *>(new SkSVGLength(c_v, (SkSVGLength::Unit)c_u));
}

void *skia_new_SkSVGLength3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGLength(reinterpret_cast<SkSVGLength &>(c_p0)));
}

void *skia_new_SkSVGIRI() { return reinterpret_cast<void *>(new SkSVGIRI()); }

void *skia_new_SkSVGIRI2(int c_t, void *c_iri) {
  return reinterpret_cast<void *>(new SkSVGIRI(
      (SkSVGIRI::Type)c_t, reinterpret_cast<SkSVGStringType &>(c_iri)));
}

void *skia_new_SkSVGColor() {
  return reinterpret_cast<void *>(new SkSVGColor());
}

void *skia_new_SkSVGColor2(void *c_c) {
  return reinterpret_cast<void *>(
      new SkSVGColor(reinterpret_cast<SkSVGColorType &>(c_c)));
}

void *skia_new_SkSVGColor5(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGColor(reinterpret_cast<SkSVGColor &>(c_p0)));
}

void *skia_new_SkSVGPaint() {
  return reinterpret_cast<void *>(new SkSVGPaint());
}

void *skia_new_SkSVGPaint2(int c_t) {
  return reinterpret_cast<void *>(new SkSVGPaint((SkSVGPaint::Type)c_t));
}

void *skia_new_SkSVGPaint5(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGPaint(reinterpret_cast<SkSVGPaint &>(c_p0)));
}

void *skia_new_SkSVGFuncIRI() {
  return reinterpret_cast<void *>(new SkSVGFuncIRI());
}

void *skia_new_SkSVGFuncIRI2(int c_t) {
  return reinterpret_cast<void *>(new SkSVGFuncIRI((SkSVGFuncIRI::Type)c_t));
}

void *skia_new_SkSVGLineJoin() {
  return reinterpret_cast<void *>(new SkSVGLineJoin());
}

void *skia_new_SkSVGLineJoin2(int c_t) {
  return reinterpret_cast<void *>(new SkSVGLineJoin((SkSVGLineJoin::Type)c_t));
}

void *skia_new_SkSVGLineJoin3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGLineJoin(reinterpret_cast<SkSVGLineJoin &>(c_p0)));
}

void *skia_new_SkSVGSpreadMethod() {
  return reinterpret_cast<void *>(new SkSVGSpreadMethod());
}

void *skia_new_SkSVGSpreadMethod2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGSpreadMethod((SkSVGSpreadMethod::Type)c_t));
}

void *skia_new_SkSVGSpreadMethod3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGSpreadMethod(reinterpret_cast<SkSVGSpreadMethod &>(c_p0)));
}

void *skia_new_SkSVGFillRule() {
  return reinterpret_cast<void *>(new SkSVGFillRule());
}

void *skia_new_SkSVGFillRule2(int c_t) {
  return reinterpret_cast<void *>(new SkSVGFillRule((SkSVGFillRule::Type)c_t));
}

void *skia_new_SkSVGFillRule3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGFillRule(reinterpret_cast<SkSVGFillRule &>(c_p0)));
}

void *skia_new_SkSVGVisibility() {
  return reinterpret_cast<void *>(new SkSVGVisibility());
}

void *skia_new_SkSVGVisibility2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGVisibility((SkSVGVisibility::Type)c_t));
}

void *skia_new_SkSVGVisibility3(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGVisibility(reinterpret_cast<SkSVGVisibility &>(c_p0)));
}

void *skia_new_SkSVGDashArray() {
  return reinterpret_cast<void *>(new SkSVGDashArray());
}

void *skia_new_SkSVGDashArray2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGDashArray((SkSVGDashArray::Type)c_t));
}

void *skia_new_SkSVGDashArray4(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGDashArray(reinterpret_cast<SkSVGDashArray &>(c_p0)));
}

void *skia_new_SkSVGStopColor() {
  return reinterpret_cast<void *>(new SkSVGStopColor());
}

void *skia_new_SkSVGStopColor2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGStopColor((SkSVGStopColor::Type)c_t));
}

void *skia_new_SkSVGStopColor3(void *c_c) {
  return reinterpret_cast<void *>(
      new SkSVGStopColor(reinterpret_cast<SkSVGColorType &>(c_c)));
}

void *skia_new_SkSVGStopColor4(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGStopColor(reinterpret_cast<SkSVGStopColor &>(c_p0)));
}

void *skia_new_SkSVGObjectBoundingBoxUnits() {
  return reinterpret_cast<void *>(new SkSVGObjectBoundingBoxUnits());
}

void *skia_new_SkSVGObjectBoundingBoxUnits2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGObjectBoundingBoxUnits((SkSVGObjectBoundingBoxUnits::Type)c_t));
}

void *skia_new_SkSVGFontFamily() {
  return reinterpret_cast<void *>(new SkSVGFontFamily());
}

void *skia_new_SkSVGFontFamily2(schar c_family) {
  return reinterpret_cast<void *>(new SkSVGFontFamily(c_family));
}

void *skia_new_SkSVGFontStyle() {
  return reinterpret_cast<void *>(new SkSVGFontStyle());
}

void *skia_new_SkSVGFontStyle2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGFontStyle((SkSVGFontStyle::Type)c_t));
}

void *skia_new_SkSVGFontSize() {
  return reinterpret_cast<void *>(new SkSVGFontSize());
}

void *skia_new_SkSVGFontSize2(void *c_s) {
  return reinterpret_cast<void *>(
      new SkSVGFontSize(reinterpret_cast<SkSVGLength &>(c_s)));
}

void *skia_new_SkSVGFontWeight() {
  return reinterpret_cast<void *>(new SkSVGFontWeight());
}

void *skia_new_SkSVGFontWeight2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGFontWeight((SkSVGFontWeight::Type)c_t));
}

void *skia_new_SkSVGTextAnchor() {
  return reinterpret_cast<void *>(new SkSVGTextAnchor());
}

void *skia_new_SkSVGTextAnchor2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGTextAnchor((SkSVGTextAnchor::Type)c_t));
}

void *skia_new_SkSVGFeInputType() {
  return reinterpret_cast<void *>(new SkSVGFeInputType());
}

void *skia_new_SkSVGFeInputType2(int c_t) {
  return reinterpret_cast<void *>(
      new SkSVGFeInputType((SkSVGFeInputType::Type)c_t));
}

void *skia_new_SkSVGFeInputType3(void *c_id) {
  return reinterpret_cast<void *>(
      new SkSVGFeInputType(reinterpret_cast<SkSVGStringType &>(c_id)));
}

void *skia_new_SkSVGFeTurbulenceBaseFrequency() {
  return reinterpret_cast<void *>(new SkSVGFeTurbulenceBaseFrequency());
}

void *skia_new_SkSVGFeTurbulenceBaseFrequency2(float c_freqX, float c_freqY) {
  return reinterpret_cast<void *>(
      new SkSVGFeTurbulenceBaseFrequency(c_freqX, c_freqY));
}

void *skia_new_SkSVGAttributeParser(schar c_p0) {
  return reinterpret_cast<void *>(new SkSVGAttributeParser(c_p0));
}

void skia_delete_SkSVGNode(void *obj) {
  delete reinterpret_cast<SkSVGNode *>(obj);
}

void *skia_new_SkSVGLengthContext(void *c_viewport, float c_dpi) {
  return reinterpret_cast<void *>(
      new SkSVGLengthContext(reinterpret_cast<SkSize &>(c_viewport), c_dpi));
}

void *skia_new_SkSVGRenderContext2(void *c_p0) {
  return reinterpret_cast<void *>(
      new SkSVGRenderContext(reinterpret_cast<SkSVGRenderContext &>(c_p0)));
}

void skia_delete_SkSVGRenderContext(void *obj) {
  delete reinterpret_cast<SkSVGRenderContext *>(obj);
}
}
