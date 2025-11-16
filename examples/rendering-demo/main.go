//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"math"

	"github.com/base-go/GoFlow/backends/macos"
	"github.com/base-go/GoFlow/pkg/core/framework"
)

func main() {
	fmt.Println("GoFlow Rendering Demo - macOS")
	fmt.Println("==============================")

	// Initialize app
	macos.InitApp()

	// Create window
	window := macos.NewWindow(800, 600, "GoFlow Rendering Demo")
	defer window.Destroy()

	// Animation state
	frame := 0

	// Set draw callback
	window.SetDrawFunc(func(canvas *macos.CoreGraphicsCanvas) {
		// Clear background with light gray
		canvas.Clear(goflow.NewColor(240, 240, 240, 255))

		// Draw title
		titleStyle := goflow.NewTextStyle()
		titleStyle.FontSize = 32
		titleStyle.FontWeight = goflow.FontWeightBold
		titleStyle.Color = goflow.NewColor(50, 50, 50, 255)
		canvas.DrawText("GoFlow Rendering Demo", goflow.NewOffset(50, 50), titleStyle)

		// Draw subtitle
		subtitleStyle := goflow.NewTextStyle()
		subtitleStyle.FontSize = 16
		subtitleStyle.Color = goflow.NewColor(100, 100, 100, 255)
		canvas.DrawText("Native macOS rendering with Core Graphics", goflow.NewOffset(50, 90), subtitleStyle)

		// Draw rectangles
		drawRectangles(canvas)

		// Draw circles
		drawCircles(canvas)

		// Draw lines
		drawLines(canvas)

		// Draw animated circle
		drawAnimatedCircle(canvas, frame)

		// Draw color palette
		drawColorPalette(canvas)

		frame++
	})

	// Set resize callback
	window.SetResizeFunc(func(width, height int) {
		fmt.Printf("Window resized to: %dx%d\n", width, height)
	})

	// Show window
	window.Show()

	fmt.Println("Window created and shown. Close the window to exit.")

	// Event loop
	for !window.ShouldClose() {
		window.PollEvents()
		window.SetNeedsDisplay()
	}

	fmt.Println("Demo finished!")
}

func drawRectangles(canvas *macos.CoreGraphicsCanvas) {
	// Filled rectangle
	paint1 := goflow.NewPaint()
	paint1.Color = goflow.NewColor(66, 133, 244, 255) // Google Blue
	rect1 := goflow.NewRect(goflow.NewOffset(50, 130), goflow.NewSize(150, 100))
	canvas.DrawRect(rect1, paint1)

	// Stroked rectangle
	paint2 := goflow.NewPaint()
	paint2.Style = goflow.PaintStyleStroke
	paint2.Color = goflow.NewColor(234, 67, 53, 255) // Google Red
	paint2.StrokeWidth = 3
	rect2 := goflow.NewRect(goflow.NewOffset(220, 130), goflow.NewSize(150, 100))
	canvas.DrawRect(rect2, paint2)

	// Labels
	labelStyle := goflow.NewTextStyle()
	labelStyle.FontSize = 12
	labelStyle.Color = goflow.NewColor(80, 80, 80, 255)
	canvas.DrawText("Filled Rect", goflow.NewOffset(50, 240), labelStyle)
	canvas.DrawText("Stroked Rect", goflow.NewOffset(220, 240), labelStyle)
}

func drawCircles(canvas *macos.CoreGraphicsCanvas) {
	// Filled circle
	paint1 := goflow.NewPaint()
	paint1.Color = goflow.NewColor(52, 168, 83, 255) // Google Green
	canvas.DrawCircle(goflow.NewOffset(475, 180), 50, paint1)

	// Stroked circle
	paint2 := goflow.NewPaint()
	paint2.Style = goflow.PaintStyleStroke
	paint2.Color = goflow.NewColor(251, 188, 5, 255) // Google Yellow
	paint2.StrokeWidth = 3
	canvas.DrawCircle(goflow.NewOffset(625, 180), 50, paint2)

	// Labels
	labelStyle := goflow.NewTextStyle()
	labelStyle.FontSize = 12
	labelStyle.Color = goflow.NewColor(80, 80, 80, 255)
	canvas.DrawText("Filled Circle", goflow.NewOffset(425, 240), labelStyle)
	canvas.DrawText("Stroked Circle", goflow.NewOffset(565, 240), labelStyle)
}

func drawLines(canvas *macos.CoreGraphicsCanvas) {
	paint := goflow.NewPaint()
	paint.Style = goflow.PaintStyleStroke
	paint.StrokeWidth = 2

	y := 280.0
	colors := []*goflow.Color{
		goflow.ColorRed,
		goflow.ColorGreen,
		goflow.ColorBlue,
		goflow.ColorYellow,
		goflow.ColorCyan,
		goflow.ColorMagenta,
	}

	for i, color := range colors {
		paint.Color = color
		y1 := y + float64(i)*15
		canvas.DrawLine(
			goflow.NewOffset(50, y1),
			goflow.NewOffset(350, y1),
			paint,
		)
	}

	// Label
	labelStyle := goflow.NewTextStyle()
	labelStyle.FontSize = 12
	labelStyle.Color = goflow.NewColor(80, 80, 80, 255)
	canvas.DrawText("Lines with different colors", goflow.NewOffset(50, 395), labelStyle)
}

func drawAnimatedCircle(canvas *macos.CoreGraphicsCanvas, frame int) {
	// Animated bouncing circle
	centerX := 550.0
	centerY := 350.0
	radius := 30.0

	// Simple sine wave animation
	t := float64(frame) * 0.05
	offsetX := math.Sin(t) * 100
	offsetY := math.Cos(t) * 50

	paint := goflow.NewPaint()
	paint.Color = goflow.NewColor(156, 39, 176, 255) // Purple
	canvas.DrawCircle(goflow.NewOffset(centerX+offsetX, centerY+offsetY), radius, paint)

	// Label
	labelStyle := goflow.NewTextStyle()
	labelStyle.FontSize = 12
	labelStyle.Color = goflow.NewColor(80, 80, 80, 255)
	canvas.DrawText("Animated Circle", goflow.NewOffset(490, 420), labelStyle)
}

func drawColorPalette(canvas *macos.CoreGraphicsCanvas) {
	// Draw a color palette
	colors := []*goflow.Color{
		goflow.ColorBlack,
		goflow.ColorWhite,
		goflow.ColorRed,
		goflow.ColorGreen,
		goflow.ColorBlue,
		goflow.ColorYellow,
		goflow.ColorCyan,
		goflow.ColorMagenta,
		goflow.ColorGray,
	}

	paint := goflow.NewPaint()
	x := 50.0
	y := 450.0
	size := 40.0

	for _, color := range colors {
		paint.Color = color
		rect := goflow.NewRect(goflow.NewOffset(x, y), goflow.NewSize(size, size))
		canvas.DrawRect(rect, paint)

		// Draw border
		borderPaint := goflow.NewPaint()
		borderPaint.Style = goflow.PaintStyleStroke
		borderPaint.Color = goflow.ColorGray
		borderPaint.StrokeWidth = 1
		canvas.DrawRect(rect, borderPaint)

		x += size + 5
	}

	// Label
	labelStyle := goflow.NewTextStyle()
	labelStyle.FontSize = 12
	labelStyle.Color = goflow.NewColor(80, 80, 80, 255)
	canvas.DrawText("Color Palette", goflow.NewOffset(50, 505), labelStyle)
}
