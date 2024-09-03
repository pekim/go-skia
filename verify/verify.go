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

	// p := skia.NewPaint()
	// p.Delete()

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
