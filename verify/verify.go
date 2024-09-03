package main

import (
	"fmt"

	"github.com/pekim/go-skia"
)

// A simple program for verifying that the skia package builds
// and can be used.

func main() {
	if skia.PaintJoinBevel != 2 {
		panic(fmt.Sprintf("unexpected value for skia.PaintJoinBevel, %d", skia.PaintJoinBevel))
	}

	paint := skia.NewPaint()
	if paint == nil {
		panic("Paint is nil")
	}
	paint2 := skia.NewPaintCopy(paint)
	if paint2 == nil {
		panic("Paint is nil")
	}
	// p.Delete()

	// path := skia.NewPath()
	// path2 := skia.NewPathCopy(path)
	// fmt.Println(path2)

	// s := skia.NewString()
	// skia.NewString5(s)
	// s.Delete()
	// s2.Delete()

	// fmt.Println(skia.NewM44())
	// fmt.Println(skia.NewM442(
	// 	0, 1, 2, 3,
	// 	0, 1, 2, 3,
	// 	0, 1, 2, 3,
	// 	0, 1, 2, 3,
	// ))
}
