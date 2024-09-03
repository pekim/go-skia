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

type class unsafe.Pointer

/*
SkPaint controls options applied when drawing. SkPaint collects all
options outside of the SkCanvas clip and SkCanvas matrix.

Various options apply to strokes and fills, and images.

SkPaint collects effects and filters that describe single-pass and multiple-pass
algorithms that alter the drawing geometry, color, and transparency. For instance,
SkPaint does not directly implement dashing or blur, but contains the objects that do so.
*/
type Paint class

/*
	Cap draws at the beginning and end of an open path contour.
*/
type PaintCap int64

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

/*
	Set Style to fill, stroke, or both fill and stroke geometry.
	The stroke and fill
	share all paint attributes; for instance, they are drawn with the same color.

	Use kStrokeAndFill_Style to avoid hitting the same pixels twice with a stroke draw and
	a fill draw.
*/
type PaintStyle int64

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
type Path class

/*
	AddPathMode chooses how addPath() appends. Adding one SkPath to another can extend
	the last contour or start a new contour.
*/
type PathAddPathMode int64

/*
	Four oval parts with radii (rx, ry) start at last SkPath SkPoint and ends at (x, y).
	ArcSize and Direction select one of the four oval parts.
*/
type PathArcSize int64

/*
	SegmentMask constants correspond to each drawing Verb type in SkPath; for
	instance, if SkPath only contains lines, only the kLine_SegmentMask bit is set.
*/
type PathSegmentMask int64

/*
	Verb instructs SkPath how to interpret one or more SkPoint and optional conic weight;
	manage contour, and terminate SkPath.
*/
type PathVerb int64
