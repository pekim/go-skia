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

type PaintCap int

const (
	PaintCapButt   PaintCap = 0
	PaintCapRound  PaintCap = 1
	PaintCapSquare PaintCap = 2
)

type PaintJoin int

const (
	PaintJoinMiter PaintJoin = 0
	PaintJoinRound PaintJoin = 1
	PaintJoinBevel PaintJoin = 2
)

type PaintStyle int

const (
	PaintStyleFill          PaintStyle = 0
	PaintStyleStroke        PaintStyle = 1
	PaintStyleStrokeAndFill PaintStyle = 2
)
