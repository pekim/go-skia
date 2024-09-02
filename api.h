// This is a generated file. DO NOT EDIT.

#include <sys/types.h>
typedef signed char schar;
typedef unsigned char uchar;

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

void *skia_new_SkMatrix();
void *skia_new_SkRefCntBase();
void skia_delete_SkRefCntBase(void *obj);
void *skia_new_SkColorInfo();
void *skia_new_SkColorInfo3(void *c_p0);
void skia_delete_SkColorInfo(void *obj);
void *skia_new_SkPixmap();
void *skia_new_SkPixmap2(void *c_info, void c_addr, ulong c_rowBytes);
void *skia_new_SkYUVAInfo();
void *skia_new_SkYUVAInfo2(void *c_p0);
void *skia_new_SkYUVAPixmapInfo();
void *skia_new_SkYUVAPixmapInfo2(void *c_p0, int c_p1, ulong c_rowBytes);
void *skia_new_SkYUVAPixmapInfo3(void *c_p0, int c_p1, ulong c_rowBytes);
void *skia_new_SkYUVAPixmapInfo4(void *c_p0);
void *skia_new_SkYUVAPixmaps();
void *skia_new_SkYUVAPixmaps3(void *c_p0);
void skia_delete_SkYUVAPixmaps(void *obj);
void *skia_new_SkOnce();
void *skia_new_SkNoncopyable();
void skia_delete_SkCodec(void *obj);
void *skia_new_SkM44(void *c_src);
void *skia_new_SkM442();
void *skia_new_SkM443(void *c_a, void *c_b);
void *skia_new_SkM444(int c_p0);
void *skia_new_SkM445(int c_p0);
void *skia_new_SkM446(float c_m0, float c_m4, float c_m8, float c_m12,
                      float c_m1, float c_m5, float c_m9, float c_m13,
                      float c_m2, float c_m6, float c_m10, float c_m14,
                      float c_m3, float c_m7, float c_m11, float c_m15);
void *skia_new_SkM447(void *c_src);
void *skia_new_SkPaint();
void *skia_new_SkPaint3(void *c_paint);
void skia_delete_SkPaint(void *obj);
void skia_delete_SkRasterHandleAllocator(void *obj);
void *skia_new_SkString();
void *skia_new_SkString2(ulong c_len);
void *skia_new_SkString3(schar c_text);
void *skia_new_SkString4(schar c_text, ulong c_len);
void *skia_new_SkString5(void *c_p0);
void skia_delete_SkString(void *obj);
void *skia_new_SkSurfaceProps();
void *skia_new_SkSurfaceProps2(uint c_flags, int c_p1);
void *skia_new_SkSurfaceProps3(uint c_flags, int c_p1, float c_textContrast,
                               float c_textGamma);
void *skia_new_SkSurfaceProps4(void *c_p0);
void *skia_new_SkDeque(ulong c_elemSize, int c_allocCount);
void *skia_new_SkDeque2(ulong c_elemSize, void c_storage, ulong c_storageSize,
                        int c_allocCount);
void skia_delete_SkDeque(void *obj);
void *skia_new_SkContainerAllocator(ulong c_sizeOfT, int c_maxCapacity);
void *skia_new_SkCanvas();
void *skia_new_SkCanvas4(void *c_bitmap);
void *skia_new_SkCanvas5(void *c_bitmap, void *c_props);
void skia_delete_SkCanvas(void *obj);
void skia_delete_SkAutoCanvasRestore(void *obj);
void skia_delete_SkStream(void *obj);
void skia_delete_SkWStream(void *obj);
void *skia_new_SkFILEStream(schar c_path);
void skia_delete_SkFILEStream(void *obj);
void *skia_new_SkMemoryStream();
void *skia_new_SkMemoryStream2(ulong c_length);
void *skia_new_SkMemoryStream3(void c_data, ulong c_length, c_copyData);
void *skia_new_SkFILEWStream(schar c_path);
void skia_delete_SkFILEWStream(void *obj);
void *skia_new_SkDynamicMemoryWStream();
void skia_delete_SkDynamicMemoryWStream(void *obj);
void *skia_new_SkFontStyle(int c_weight, int c_width, int c_slant);
void *skia_new_SkFontStyle2();
void *skia_new_SkWeakRefCnt();
void skia_delete_SkWeakRefCnt(void *obj);
void *skia_new_SkFont();
void *skia_new_SkTextBlobBuilder();
void skia_delete_SkTextBlobBuilder(void *obj);
void *skia_new_SkPath();
void *skia_new_SkPath2(void *c_path);
void skia_delete_SkPath(void *obj);
void skia_delete_SkPicture(void *obj);
void skia_delete_SkEncoder(void *obj);
void *skia_new_GrGLExtensions();
void *skia_new_GrGLExtensions2(void *c_p0);
void *skia_new_GrBackendFormat();
void *skia_new_GrBackendFormat2(void *c_p0);
void skia_delete_GrBackendFormat(void *obj);
void *skia_new_GrBackendTexture();
void *skia_new_GrBackendTexture3(void *c_that);
void skia_delete_GrBackendTexture(void *obj);
void *skia_new_GrBackendRenderTarget();
void *skia_new_GrBackendRenderTarget2(int c_width, int c_height,
                                      int c_sampleCnt, int c_stencilBits,
                                      void *c_mockInfo);
void *skia_new_GrBackendRenderTarget3(void *c_that);
void skia_delete_GrBackendRenderTarget(void *obj);
void *skia_new_GrDriverBugWorkarounds();
void *skia_new_GrDriverBugWorkarounds2(void *c_p0);
void skia_delete_GrDriverBugWorkarounds(void *obj);
void skia_delete_GrContext_Base(void *obj);
void *skia_new_SkSemaphore(int c_count);
void skia_delete_SkSemaphore(void *obj);
void *skia_new_SkMutex();
void skia_delete_SkMutex(void *obj);
void *skia_new_SkAutoMutexExclusive(void *c_mutex);
void *skia_new_SkAutoMutexExclusive2(void *c_p0);
void skia_delete_SkAutoMutexExclusive(void *obj);
void skia_delete_GrImageContext(void *obj);
void skia_delete_GrRecordingContext(void *obj);
void skia_delete_GrDirectContext(void *obj);
void *skia_new_SkTDStorage(int c_sizeOfT);
void *skia_new_SkTDStorage2(void c_src, int c_size, int c_sizeOfT);
void *skia_new_SkTDStorage3(void *c_that);
void skia_delete_SkTDStorage(void *obj);
void skia_delete_SkPDFUnion(void *obj);
void skia_delete_SkPDFObject(void *obj);
void *skia_new_SkPDFArray();
void skia_delete_SkPDFArray(void *obj);
void *skia_new_SkPDFDict(schar c_type);
void skia_delete_SkPDFDict(void *obj);
void *skia_new_SkArenaAlloc(schar c_block, ulong c_blockSize,
                            ulong c_firstHeapAllocation);
void *skia_new_SkArenaAlloc2(ulong c_firstHeapAllocation);
void *skia_new_SkArenaAlloc3(void *c_p0);
void skia_delete_SkArenaAlloc(void *obj);
void *skia_new_SkArenaAllocWithReset(schar c_block, ulong c_blockSize,
                                     ulong c_firstHeapAllocation);
void *skia_new_SkArenaAllocWithReset2(ulong c_firstHeapAllocation);
void *skia_new_SkPDFTagTree();
void skia_delete_SkPDFTagTree(void *obj);
void skia_delete_SkPDFDocument(void *obj);
void skia_delete_SkShaper(void *obj);
void *skia_new_SkSVGLength();
void *skia_new_SkSVGLength2(float c_v, int c_u);
void *skia_new_SkSVGLength3(void *c_p0);
void *skia_new_SkSVGIRI();
void *skia_new_SkSVGIRI2(int c_t, void *c_iri);
void *skia_new_SkSVGColor();
void *skia_new_SkSVGColor2(void *c_c);
void *skia_new_SkSVGColor5(void *c_p0);
void *skia_new_SkSVGPaint();
void *skia_new_SkSVGPaint2(int c_t);
void *skia_new_SkSVGPaint5(void *c_p0);
void *skia_new_SkSVGFuncIRI();
void *skia_new_SkSVGFuncIRI2(int c_t);
void *skia_new_SkSVGLineJoin();
void *skia_new_SkSVGLineJoin2(int c_t);
void *skia_new_SkSVGLineJoin3(void *c_p0);
void *skia_new_SkSVGSpreadMethod();
void *skia_new_SkSVGSpreadMethod2(int c_t);
void *skia_new_SkSVGSpreadMethod3(void *c_p0);
void *skia_new_SkSVGFillRule();
void *skia_new_SkSVGFillRule2(int c_t);
void *skia_new_SkSVGFillRule3(void *c_p0);
void *skia_new_SkSVGVisibility();
void *skia_new_SkSVGVisibility2(int c_t);
void *skia_new_SkSVGVisibility3(void *c_p0);
void *skia_new_SkSVGDashArray();
void *skia_new_SkSVGDashArray2(int c_t);
void *skia_new_SkSVGDashArray4(void *c_p0);
void *skia_new_SkSVGStopColor();
void *skia_new_SkSVGStopColor2(int c_t);
void *skia_new_SkSVGStopColor3(void *c_c);
void *skia_new_SkSVGStopColor4(void *c_p0);
void *skia_new_SkSVGObjectBoundingBoxUnits();
void *skia_new_SkSVGObjectBoundingBoxUnits2(int c_t);
void *skia_new_SkSVGFontFamily();
void *skia_new_SkSVGFontFamily2(schar c_family);
void *skia_new_SkSVGFontStyle();
void *skia_new_SkSVGFontStyle2(int c_t);
void *skia_new_SkSVGFontSize();
void *skia_new_SkSVGFontSize2(void *c_s);
void *skia_new_SkSVGFontWeight();
void *skia_new_SkSVGFontWeight2(int c_t);
void *skia_new_SkSVGTextAnchor();
void *skia_new_SkSVGTextAnchor2(int c_t);
void *skia_new_SkSVGFeInputType();
void *skia_new_SkSVGFeInputType2(int c_t);
void *skia_new_SkSVGFeInputType3(void *c_id);
void *skia_new_SkSVGFeTurbulenceBaseFrequency();
void *skia_new_SkSVGFeTurbulenceBaseFrequency2(float c_freqX, float c_freqY);
void *skia_new_SkSVGAttributeParser(schar c_p0);
void skia_delete_SkSVGNode(void *obj);
void *skia_new_SkSVGLengthContext(void *c_viewport, float c_dpi);
void *skia_new_SkSVGRenderContext2(void *c_p0);
void skia_delete_SkSVGRenderContext(void *obj);

#ifdef __cplusplus
}
#endif // __cplusplus
