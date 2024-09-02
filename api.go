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

// import (
// 	"unsafe"
// )

type Paint struct {
}

/*
Cap enum specifies the settings for the paint's strokecap.

This is the treatment that is applied to the beginning and end of each non-closed contour (e.g. lines).

If the cap is round or square, the caps are drawn when the contour has a zero length. Zero length contours can be created by following moveTo with a lineTo at the same point, or a moveTo followed by a close.

A dash with an on interval of zero also creates a zero length contour.

The zero length contour draws the square cap without rotation, since the no direction can be inferred.
*/
type PaintCap int

const (
	// begin/end contours with no extension
	PaintCapButt PaintCap = 0
	// begin/end contours with a semi-circle extension
	PaintCapRound PaintCap = 1
	// begin/end contours with a half square extension
	PaintCapSquare PaintCap = 2
)

type PaintJoin int

const (
	// connect path segments with a sharp join
	PaintJoinMiter PaintJoin = 0
	// connect path segments with a round join
	PaintJoinRound PaintJoin = 1
	// connect path segments with a flat bevel join
	PaintJoinBevel PaintJoin = 2
)

type PaintStyle int

const (
	// fill the geometry
	PaintStyleFill PaintStyle = 0
	// stroke the geometry
	PaintStyleStroke PaintStyle = 1
	// fill and stroke the geometry
	PaintStyleStrokeAndFill PaintStyle = 2
)
