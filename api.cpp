// This is a generated file. DO NOT EDIT.

#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>

extern "C" {
#include "api.h"

void *misk_new_Paint() { return reinterpret_cast<void *>(new SkPaint()); }

void *misk_new_PaintCopy(void *c_paint) {
  return reinterpret_cast<void *>(
      new SkPaint(*reinterpret_cast<SkPaint *>(c_paint)));
}

void *misk_new_Path() { return reinterpret_cast<void *>(new SkPath()); }

void *misk_new_PathCopy(void *c_path) {
  return reinterpret_cast<void *>(
      new SkPath(*reinterpret_cast<SkPath *>(c_path)));
}
}
