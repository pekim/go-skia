package main

import (
	_ "embed"
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pekim/go-skia"
)

//go:embed tiger.svg
var tigerSVG []byte

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

	if err := gl.Init(); err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(0)

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press && mods == glfw.ModControl && key == glfw.KeyQ {
			window.SetShouldClose(true)
		}
	})

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

	var backend skia.GrBackendRenderTarget
	var target skia.GrBackendRenderTarget
	var surface skia.Surface
	framebufferSize := func(width int, height int) {
		if !surface.IsNil() {
			surface.Unref()
		}
		if !target.IsNil() {
			target.Delete()
		}
		if !backend.IsNil() {
			backend.Delete()
		}

		backend = skia.NewGrBackendRenderTargetCopy(skia.GrBackendRenderTargetsMakeGL(int32(width), int32(height), 1, 8, fbInfo))
		if backend.IsNil() {
			panic("failed to create backend")
		}
		target = skia.NewGrBackendRenderTargetCopy(backend)
		if target.IsNil() {
			panic("failed to create target")
		}
		var colorspace skia.ColorSpace

		surface = skia.SkSurfacesWrapBackendRenderTarget(context.AsGrRecordingContext(), backend, skia.GrSurfaceOriginBottomLeft,
			skia.ColorTypeBGRA_8888, colorspace, skia.NewSurfacePropsPixelGeometry(0, skia.PixelGeometryRGB_H))
		if surface.IsNil() {
			panic("failed to create surface")
		}

	}
	window.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
		framebufferSize(width, height)
	})
	framebufferSize(window.GetFramebufferSize())

	typeface := skia.FontMgrRefDefault().MatchFamilyStyle("sans-serif", skia.FontStyleNormal())
	font := skia.NewFontTypefaceSize(typeface, 22)
	var metrics skia.FontMetrics
	lineSpacing := font.GetMetrics(&metrics)

	tigerData := skia.DataMakeWithCopy(tigerSVG, uint32(len(tigerSVG)))
	tigerStream := skia.MemoryStreamMake(tigerData).AsStream()
	svgTiger := skia.SVGDOMMakeFromStream(tigerStream)
	if svgTiger.IsNil() {
		panic("failed to create tiger svg")
	}
	tigerStream.Delete()
	tigerData.Unref()
	var size skia.Size
	size.SetWidth(300)
	size.SetHeight(300)
	svgTiger.SetContainerSize(size)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		canvas := surface.GetCanvas()
		paint := skia.NewPaint()
		paint.SetColor(skia.ColorRED)
		paint.SetStyle(skia.PaintStyleFill)

		rect := skia.RectMakeXYWH(100, 100, 100, 100)
		canvas.DrawRect(rect, paint)

		paint.SetColor(skia.ColorBLACK)
		canvas.DrawString("Some text", 100, 250, font, paint)
		canvas.DrawString(fmt.Sprintf("Font ascent = %.1f", metrics.Ascent()), 100, 250+lineSpacing, font, paint)

		canvas.Save()
		// paint.SetColor(skia.ColorTRANSPARENT)
		// canvas.DrawPaint(paint)
		canvas.Translate(200, 200)
		canvas.Scale(2.5, 2.5)
		svgTiger.Render(canvas)
		canvas.Restore()

		context.FlushAndSubmit(skia.GrSyncCpuYes)

		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
