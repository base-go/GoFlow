// +build darwin

package macos

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework QuartzCore
#include "cgo_bridge.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/base-go/GoFlow/pkg/core/framework"
)

// CoreGraphicsCanvas implements the Canvas interface using macOS Core Graphics
type CoreGraphicsCanvas struct {
	ctx    unsafe.Pointer
	width  int
	height int
}

// NewCoreGraphicsCanvas creates a new Core Graphics canvas
func NewCoreGraphicsCanvas(width, height int) *CoreGraphicsCanvas {
	ctx := C.createGraphicsContext(C.int(width), C.int(height))
	if ctx == nil {
		panic("Failed to create Core Graphics context")
	}

	return &CoreGraphicsCanvas{
		ctx:    ctx,
		width:  width,
		height: height,
	}
}

// Destroy frees the graphics context
func (c *CoreGraphicsCanvas) Destroy() {
	if c.ctx != nil {
		C.destroyGraphicsContext(c.ctx)
		c.ctx = nil
	}
}

// Clear clears the canvas with a color
func (c *CoreGraphicsCanvas) Clear(color *goflow.Color) {
	C.clearCanvas(
		c.ctx,
		C.double(color.R),
		C.double(color.G),
		C.double(color.B),
		C.double(color.A),
	)
}

// DrawRect draws a rectangle
func (c *CoreGraphicsCanvas) DrawRect(rect *goflow.Rect, paint *goflow.Paint) {
	filled := C.int(1)
	strokeWidth := C.double(1.0)

	if paint.Style == goflow.PaintStyleStroke {
		filled = C.int(0)
		strokeWidth = C.double(paint.StrokeWidth)
	}

	C.drawRect(
		c.ctx,
		C.double(rect.Left()),
		C.double(rect.Top()),
		C.double(rect.Size.Width),
		C.double(rect.Size.Height),
		C.double(paint.Color.R),
		C.double(paint.Color.G),
		C.double(paint.Color.B),
		C.double(paint.Color.A),
		filled,
		strokeWidth,
	)
}

// DrawCircle draws a circle
func (c *CoreGraphicsCanvas) DrawCircle(center *goflow.Offset, radius float64, paint *goflow.Paint) {
	filled := C.int(1)
	strokeWidth := C.double(1.0)

	if paint.Style == goflow.PaintStyleStroke {
		filled = C.int(0)
		strokeWidth = C.double(paint.StrokeWidth)
	}

	C.drawCircle(
		c.ctx,
		C.double(center.X),
		C.double(center.Y),
		C.double(radius),
		C.double(paint.Color.R),
		C.double(paint.Color.G),
		C.double(paint.Color.B),
		C.double(paint.Color.A),
		filled,
		strokeWidth,
	)
}

// DrawLine draws a line between two points
func (c *CoreGraphicsCanvas) DrawLine(p1, p2 *goflow.Offset, paint *goflow.Paint) {
	C.drawLine(
		c.ctx,
		C.double(p1.X),
		C.double(p1.Y),
		C.double(p2.X),
		C.double(p2.Y),
		C.double(paint.Color.R),
		C.double(paint.Color.G),
		C.double(paint.Color.B),
		C.double(paint.Color.A),
		C.double(paint.StrokeWidth),
	)
}

// DrawText draws text at the given position
func (c *CoreGraphicsCanvas) DrawText(text string, offset *goflow.Offset, style *goflow.TextStyle) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	cFont := C.CString(style.FontFamily)
	defer C.free(unsafe.Pointer(cFont))

	C.drawText(
		c.ctx,
		cText,
		C.double(offset.X),
		C.double(offset.Y),
		cFont,
		C.double(style.FontSize),
		C.double(style.Color.R),
		C.double(style.Color.G),
		C.double(style.Color.B),
		C.double(style.Color.A),
		C.int(style.FontWeight),
	)
}

// Save saves the current canvas state
func (c *CoreGraphicsCanvas) Save() {
	C.saveCanvas(c.ctx)
}

// Restore restores the canvas state
func (c *CoreGraphicsCanvas) Restore() {
	C.restoreCanvas(c.ctx)
}

// Translate moves the origin
func (c *CoreGraphicsCanvas) Translate(dx, dy float64) {
	C.translateCanvas(c.ctx, C.double(dx), C.double(dy))
}

// Scale scales the canvas
func (c *CoreGraphicsCanvas) Scale(sx, sy float64) {
	C.scaleCanvas(c.ctx, C.double(sx), C.double(sy))
}

// Rotate rotates the canvas
func (c *CoreGraphicsCanvas) Rotate(radians float64) {
	C.rotateCanvas(c.ctx, C.double(radians))
}

// ClipRect sets a clipping rectangle
func (c *CoreGraphicsCanvas) ClipRect(rect *goflow.Rect) {
	C.clipRect(
		c.ctx,
		C.double(rect.Left()),
		C.double(rect.Top()),
		C.double(rect.Size.Width),
		C.double(rect.Size.Height),
	)
}

// GetPixelData returns the raw pixel data (for debugging or saving)
func (c *CoreGraphicsCanvas) GetPixelData() []byte {
	bufferSize := c.width * c.height * 4
	buffer := make([]byte, bufferSize)

	C.getPixelData(c.ctx, (*C.uchar)(unsafe.Pointer(&buffer[0])), C.int(bufferSize))

	return buffer
}

// GetDimensions returns the canvas dimensions
func (c *CoreGraphicsCanvas) GetDimensions() (int, int) {
	var width, height C.int
	C.getCanvasDimensions(c.ctx, &width, &height)
	return int(width), int(height)
}
