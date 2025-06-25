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

//go:embed test.png
var testPng []byte

//go:embed transparency.png
var transparencyPng []byte

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

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}
	glfw.SwapInterval(0)

	window.SetKeyCallback(func(_ *glfw.Window, key glfw.Key, _scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press && mods == glfw.ModControl && key == glfw.KeyQ {
			window.SetShouldClose(true)
		}
	})

	gl.Enable(gl.BLEND)
	gl.Enable(gl.FRAMEBUFFER_SRGB)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

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
		gl.Viewport(0, 0, int32(width), int32(height))

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

		surface = skia.SurfacesWrapBackendRenderTarget(context.AsGrRecordingContext(), backend, skia.GrSurfaceOriginBottomLeft,
			skia.ColorTypeBGRA_8888, colorspace, skia.NewSurfacePropsPixelGeometry(0, skia.PixelGeometryRGB_H))
		if surface.IsNil() {
			panic("failed to create surface")
		}

	}
	window.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
		framebufferSize(width, height)
	})
	framebufferSize(window.GetFramebufferSize())

	typeface := skia.FontMgrDefault().MatchFamilyStyle("sans-serif", skia.FontStyleNormal())
	font := skia.NewFontTypefaceSize(typeface, 19)
	var metrics skia.FontMetrics
	lineSpacing := font.GetMetrics(&metrics)

	tigerData := skia.DataMakeWithCopy(tigerSVG, uint32(len(tigerSVG)))
	tigerStream := skia.MemoryStreamMake(tigerData).AsStream()
	svgTiger := skia.NewSVGDOMBuilder().Make(tigerStream)
	if svgTiger.IsNil() {
		panic("failed to create tiger svg")
	}
	tigerStream.Delete()
	tigerData.Unref()
	var size skia.Size
	size.SetWidth(300)
	size.SetHeight(300)
	svgTiger.SetContainerSize(size)

	imageData := skia.DataMakeWithCopy(testPng, uint32(len(testPng)))
	image := skia.ImagesDeferredFromEncodedData(imageData, nil)

	image2Data := skia.DataMakeWithCopy(transparencyPng, uint32(len(transparencyPng)))
	image2 := skia.ImagesDeferredFromEncodedData(image2Data, nil)

	draw := func() {
		// start := time.Now()
		// defer func() {
		// 	fmt.Printf("%0.2fms\n", time.Since(start).Seconds()*1000)
		// }()

		canvas := surface.GetCanvas()
		paint := skia.NewPaint()
		defer paint.Delete()
		paint.SetStyle(skia.PaintStyleFill)

		paint.SetColor(skia.ColorWHITE)
		windowW, windowH := window.GetFramebufferSize()
		rectWindow := skia.IRectMakeXYWH(0, 0, int32(windowW), int32(windowH))
		canvas.DrawIRect(rectWindow, paint)

		paint.SetColor(skia.ColorRED)
		rect := skia.RectMakeXYWH(100, 100, 100, 100)
		canvas.DrawRect(rect, paint)

		paint.SetColor(skia.ColorBLACK)
		canvas.DrawString("Some text", 100, 250, font, paint)
		canvas.DrawString(fmt.Sprintf("Font ascent = %.1f", metrics.Ascent()), 100, 250+lineSpacing, font, paint)

		paint.SetColor(skia.ColorGREEN)
		rect2 := skia.RectMakeXYWH(400, 100, 100, 100)
		canvas.DrawRect(rect2, paint)

		canvas.Save()
		canvas.Translate(180, 150)
		scale := float32(2.5)
		canvas.Scale(scale, scale)
		svgTiger.Render(canvas)
		canvas.Restore()

		canvas.DrawImage(image, 1000, 200)

		paint.SetColor(skia.ColorRED)
		paint.SetAlpha(0x20)
		rect3 := skia.RectMakeXYWH(180, 800, float32(image2.Width()), float32(image2.Height()))
		canvas.DrawRect(rect3, paint)
		canvas.DrawImage(image2, 180, 800) // position so that a bit of the tiger is under the translucent image

		context.FlushAndSubmit(skia.GrSyncCpuYes)
	}

	for !window.ShouldClose() {
		draw()
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
