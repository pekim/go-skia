package main

import (
	"github.com/pekim/go-skia"
)

// A simple program for verifying that the skia package builds
// and can be used.

func main() {
	p := skia.NewPaint()
	p.Delete()
}
