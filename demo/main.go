package main

import (
	"fmt"
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pekim/go-skia"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "skia demo", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press && mods == glfw.ModControl && key == glfw.KeyQ {
			window.SetShouldClose(true)
		}
	})
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	iface := skia.GrGLMakeNativeInterface()
	context := skia.GrDirectContextsMakeGLInterface(iface)
	var fbo int32
	gl.GetIntegerv(gl.FRAMEBUFFER_BINDING, &fbo)
	var fbInfo skia.GrGLFramebufferInfo
	fbInfo.SetFBOID(uint32(fbo))
	fbInfo.SetFormat(gl.RGBA8)

	backend := skia.GrBackendRenderTargetsMakeGL(800, 600, 1, 8, fbInfo)
	if backend.IsNil() {
		panic("failed to create backend")
	}
	target := skia.NewGrBackendRenderTargetCopy(backend)
	if target.IsNil() {
		panic("failed to create target")
	}
	var colorspace skia.ColorSpace

	surface := skia.SkSurfacesWrapBackendRenderTarget(context.AsGrRecordingContext(), backend, skia.GrSurfaceOriginBottomLeft,
		skia.ColorTypeBGRA_8888, colorspace, skia.NewSurfacePropsPixelGeometry(0, skia.PixelGeometryRGB_H))
	if surface.IsNil() {
		panic("failed to create surface")
	}

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		canvas := surface.GetCanvas()
		paint := skia.NewPaint()
		paint.SetColor(skia.ColorRED)
		paint.SetStyle(skia.PaintStyleFill)

		rect := skia.RectMakeXYWH(100, 100, 100, 100)
		canvas.DrawRect(rect, paint)

		typeface := skia.FontMgrRefDefault().MatchFamilyStyle("sans-serif", skia.FontStyleNormal())
		font := skia.NewFontTypefaceSize(typeface, 22)
		var metrics skia.FontMetrics
		lineSpacing := font.GetMetrics(&metrics)

		fmt.Println(font.UnicharToGlyph('B'))
		chars := []int32{'A', 'B', 'C', 'D', 'Y', 'Z'}
		glyphs := make([]uint16, len(chars))
		font.UnicharsToGlyphs(chars, int32(len(chars)), glyphs)
		fmt.Println(glyphs)

		paint.SetColor(skia.ColorBLACK)
		canvas.DrawString("Some text", 100, 250, font, paint)
		canvas.DrawString(fmt.Sprintf("Font ascent = %.1f", metrics.Ascent()), 100, 250+lineSpacing, font, paint)

		context.FlushAndSubmit(skia.GrSyncCpuYes)

		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
