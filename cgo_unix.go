//go:build unix

package skia

// #cgo CXXFLAGS: -D SKIA_UNIX
// #cgo pkg-config: fontconfig
import "C"
