package goflow

import "math"

// Size represents a 2D size with width and height
type Size struct {
	Width  float64
	Height float64
}

// NewSize creates a new Size
func NewSize(width, height float64) *Size {
	return &Size{Width: width, Height: height}
}

// Zero returns a zero-sized Size
func ZeroSize() *Size {
	return &Size{Width: 0, Height: 0}
}

// Infinite returns an infinite Size
func InfiniteSize() *Size {
	return &Size{Width: math.Inf(1), Height: math.Inf(1)}
}

// IsEmpty returns true if either dimension is zero
func (s *Size) IsEmpty() bool {
	return s.Width == 0 || s.Height == 0
}

// IsFinite returns true if both dimensions are finite
func (s *Size) IsFinite() bool {
	return !math.IsInf(s.Width, 0) && !math.IsInf(s.Height, 0)
}

// Offset represents a 2D offset with x and y coordinates
type Offset struct {
	X float64
	Y float64
}

// NewOffset creates a new Offset
func NewOffset(x, y float64) *Offset {
	return &Offset{X: x, Y: y}
}

// ZeroOffset returns a zero offset
func ZeroOffset() *Offset {
	return &Offset{X: 0, Y: 0}
}

// Add adds two offsets
func (o *Offset) Add(other *Offset) *Offset {
	return &Offset{X: o.X + other.X, Y: o.Y + other.Y}
}

// Subtract subtracts an offset from this offset
func (o *Offset) Subtract(other *Offset) *Offset {
	return &Offset{X: o.X - other.X, Y: o.Y - other.Y}
}

// Rect represents a rectangle with offset and size
type Rect struct {
	Offset *Offset
	Size   *Size
}

// NewRect creates a new Rect
func NewRect(offset *Offset, size *Size) *Rect {
	return &Rect{Offset: offset, Size: size}
}

// NewRectFromLTWH creates a Rect from left, top, width, height
func NewRectFromLTWH(left, top, width, height float64) *Rect {
	return &Rect{
		Offset: &Offset{X: left, Y: top},
		Size:   &Size{Width: width, Height: height},
	}
}

// Left returns the left edge
func (r *Rect) Left() float64 {
	return r.Offset.X
}

// Top returns the top edge
func (r *Rect) Top() float64 {
	return r.Offset.Y
}

// Right returns the right edge
func (r *Rect) Right() float64 {
	return r.Offset.X + r.Size.Width
}

// Bottom returns the bottom edge
func (r *Rect) Bottom() float64 {
	return r.Offset.Y + r.Size.Height
}

// Contains checks if a point is inside the rectangle
func (r *Rect) Contains(point *Offset) bool {
	return point.X >= r.Left() && point.X <= r.Right() &&
		point.Y >= r.Top() && point.Y <= r.Bottom()
}

// EdgeInsets represents padding or margin
type EdgeInsets struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

// NewEdgeInsets creates edge insets with individual values
func NewEdgeInsets(left, top, right, bottom float64) *EdgeInsets {
	return &EdgeInsets{Left: left, Top: top, Right: right, Bottom: bottom}
}

// NewEdgeInsetsAll creates edge insets with the same value for all sides
func NewEdgeInsetsAll(value float64) *EdgeInsets {
	return &EdgeInsets{Left: value, Top: value, Right: value, Bottom: value}
}

// NewEdgeInsetsSymmetric creates edge insets with horizontal and vertical values
func NewEdgeInsetsSymmetric(horizontal, vertical float64) *EdgeInsets {
	return &EdgeInsets{Left: horizontal, Top: vertical, Right: horizontal, Bottom: vertical}
}

// ZeroEdgeInsets returns zero edge insets
func ZeroEdgeInsets() *EdgeInsets {
	return &EdgeInsets{Left: 0, Top: 0, Right: 0, Bottom: 0}
}

// Horizontal returns the total horizontal insets
func (e *EdgeInsets) Horizontal() float64 {
	return e.Left + e.Right
}

// Vertical returns the total vertical insets
func (e *EdgeInsets) Vertical() float64 {
	return e.Top + e.Bottom
}
