package goflow

// Color represents an RGBA color
type Color struct {
	R, G, B, A float64
}

// NewColor creates a new color from RGBA values (0-255)
func NewColor(r, g, b, a uint8) *Color {
	return &Color{
		R: float64(r) / 255.0,
		G: float64(g) / 255.0,
		B: float64(b) / 255.0,
		A: float64(a) / 255.0,
	}
}

// Common colors
var (
	ColorBlack       = NewColor(0, 0, 0, 255)
	ColorWhite       = NewColor(255, 255, 255, 255)
	ColorRed         = NewColor(255, 0, 0, 255)
	ColorGreen       = NewColor(0, 255, 0, 255)
	ColorBlue        = NewColor(0, 0, 255, 255)
	ColorYellow      = NewColor(255, 255, 0, 255)
	ColorCyan        = NewColor(0, 255, 255, 255)
	ColorMagenta     = NewColor(255, 0, 255, 255)
	ColorTransparent = NewColor(0, 0, 0, 0)
	ColorGray        = NewColor(128, 128, 128, 255)
)

// Paint describes the style for drawing
type Paint struct {
	Color       *Color
	StrokeWidth float64
	Style       PaintStyle
}

// PaintStyle determines if paint is fill or stroke
type PaintStyle int

const (
	PaintStyleFill PaintStyle = iota
	PaintStyleStroke
)

// NewPaint creates a new paint with default fill style
func NewPaint() *Paint {
	return &Paint{
		Color:       ColorBlack,
		StrokeWidth: 1.0,
		Style:       PaintStyleFill,
	}
}

// TextStyle describes how text should be rendered
type TextStyle struct {
	Color      *Color
	FontSize   float64
	FontFamily string
	FontWeight FontWeight
}

// FontWeight represents font weight
type FontWeight int

const (
	FontWeightNormal FontWeight = 400
	FontWeightBold   FontWeight = 700
)

// NewTextStyle creates a new text style with defaults
func NewTextStyle() *TextStyle {
	return &TextStyle{
		Color:      ColorBlack,
		FontSize:   14.0,
		FontFamily: "sans-serif",
		FontWeight: FontWeightNormal,
	}
}

// Canvas provides drawing operations
// This is an abstraction over the actual rendering backend
type Canvas interface {
	// DrawRect draws a rectangle
	DrawRect(rect *Rect, paint *Paint)

	// DrawCircle draws a circle
	DrawCircle(center *Offset, radius float64, paint *Paint)

	// DrawText draws text at the given position
	DrawText(text string, offset *Offset, style *TextStyle)

	// DrawLine draws a line between two points
	DrawLine(p1, p2 *Offset, paint *Paint)

	// Save saves the current canvas state
	Save()

	// Restore restores the canvas state
	Restore()

	// Translate moves the origin
	Translate(dx, dy float64)

	// Scale scales the canvas
	Scale(sx, sy float64)

	// Rotate rotates the canvas
	Rotate(radians float64)

	// ClipRect sets a clipping rectangle
	ClipRect(rect *Rect)

	// Clear clears the canvas with a color
	Clear(color *Color)
}

// MockCanvas is a simple canvas implementation for testing
type MockCanvas struct {
	operations []string
	width      float64
	height     float64
}

// NewMockCanvas creates a new mock canvas
func NewMockCanvas(width, height float64) *MockCanvas {
	return &MockCanvas{
		operations: make([]string, 0),
		width:      width,
		height:     height,
	}
}

func (c *MockCanvas) DrawRect(rect *Rect, paint *Paint) {
	c.operations = append(c.operations, "DrawRect")
}

func (c *MockCanvas) DrawCircle(center *Offset, radius float64, paint *Paint) {
	c.operations = append(c.operations, "DrawCircle")
}

func (c *MockCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
	c.operations = append(c.operations, "DrawText")
}

func (c *MockCanvas) DrawLine(p1, p2 *Offset, paint *Paint) {
	c.operations = append(c.operations, "DrawLine")
}

func (c *MockCanvas) Save() {
	c.operations = append(c.operations, "Save")
}

func (c *MockCanvas) Restore() {
	c.operations = append(c.operations, "Restore")
}

func (c *MockCanvas) Translate(dx, dy float64) {
	c.operations = append(c.operations, "Translate")
}

func (c *MockCanvas) Scale(sx, sy float64) {
	c.operations = append(c.operations, "Scale")
}

func (c *MockCanvas) Rotate(radians float64) {
	c.operations = append(c.operations, "Rotate")
}

func (c *MockCanvas) ClipRect(rect *Rect) {
	c.operations = append(c.operations, "ClipRect")
}

func (c *MockCanvas) Clear(color *Color) {
	c.operations = append(c.operations, "Clear")
}

// GetOperations returns the recorded operations
func (c *MockCanvas) GetOperations() []string {
	return c.operations
}
