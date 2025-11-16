package goflow

import "math"

// Constraints define the size constraints for layout
// Following Flutter's box constraints model
type Constraints struct {
	MinWidth  float64
	MaxWidth  float64
	MinHeight float64
	MaxHeight float64
}

// NewConstraints creates new constraints
func NewConstraints(minWidth, maxWidth, minHeight, maxHeight float64) *Constraints {
	return &Constraints{
		MinWidth:  minWidth,
		MaxWidth:  maxWidth,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
	}
}

// TightConstraints creates constraints with exact dimensions
func TightConstraints(width, height float64) *Constraints {
	return &Constraints{
		MinWidth:  width,
		MaxWidth:  width,
		MinHeight: height,
		MaxHeight: height,
	}
}

// TightConstraintsForSize creates tight constraints from a size
func TightConstraintsForSize(size *Size) *Constraints {
	return TightConstraints(size.Width, size.Height)
}

// LooseConstraints creates loose constraints (min = 0, max = given values)
func LooseConstraints(maxWidth, maxHeight float64) *Constraints {
	return &Constraints{
		MinWidth:  0,
		MaxWidth:  maxWidth,
		MinHeight: 0,
		MaxHeight: maxHeight,
	}
}

// UnboundedConstraints creates unbounded constraints
func UnboundedConstraints() *Constraints {
	return &Constraints{
		MinWidth:  0,
		MaxWidth:  math.Inf(1),
		MinHeight: 0,
		MaxHeight: math.Inf(1),
	}
}

// IsTight returns true if the constraints can only be satisfied by a single size
func (c *Constraints) IsTight() bool {
	return c.MinWidth >= c.MaxWidth && c.MinHeight >= c.MaxHeight
}

// IsNormalized returns true if the constraints are valid
func (c *Constraints) IsNormalized() bool {
	return c.MinWidth >= 0 && c.MinWidth <= c.MaxWidth &&
		c.MinHeight >= 0 && c.MinHeight <= c.MaxHeight
}

// IsSatisfiedBy checks if a size satisfies these constraints
func (c *Constraints) IsSatisfiedBy(size *Size) bool {
	return size.Width >= c.MinWidth && size.Width <= c.MaxWidth &&
		size.Height >= c.MinHeight && size.Height <= c.MaxHeight
}

// Constrain returns the size that satisfies the constraints
func (c *Constraints) Constrain(size *Size) *Size {
	return &Size{
		Width:  clamp(size.Width, c.MinWidth, c.MaxWidth),
		Height: clamp(size.Height, c.MinHeight, c.MaxHeight),
	}
}

// ConstrainWidth constrains the width
func (c *Constraints) ConstrainWidth(width float64) float64 {
	return clamp(width, c.MinWidth, c.MaxWidth)
}

// ConstrainHeight constrains the height
func (c *Constraints) ConstrainHeight(height float64) float64 {
	return clamp(height, c.MinHeight, c.MaxHeight)
}

// Tighten creates tighter constraints
func (c *Constraints) Tighten(width, height *float64) *Constraints {
	minWidth := c.MinWidth
	maxWidth := c.MaxWidth
	minHeight := c.MinHeight
	maxHeight := c.MaxHeight

	if width != nil {
		minWidth = clamp(*width, c.MinWidth, c.MaxWidth)
		maxWidth = minWidth
	}
	if height != nil {
		minHeight = clamp(*height, c.MinHeight, c.MaxHeight)
		maxHeight = minHeight
	}

	return &Constraints{
		MinWidth:  minWidth,
		MaxWidth:  maxWidth,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
	}
}

// Loosen creates looser constraints (sets minimums to 0)
func (c *Constraints) Loosen() *Constraints {
	return &Constraints{
		MinWidth:  0,
		MaxWidth:  c.MaxWidth,
		MinHeight: 0,
		MaxHeight: c.MaxHeight,
	}
}

// Deflate reduces the constraints by the given edge insets
func (c *Constraints) Deflate(insets *EdgeInsets) *Constraints {
	horizontal := insets.Horizontal()
	vertical := insets.Vertical()

	return &Constraints{
		MinWidth:  math.Max(0, c.MinWidth-horizontal),
		MaxWidth:  math.Max(0, c.MaxWidth-horizontal),
		MinHeight: math.Max(0, c.MinHeight-vertical),
		MaxHeight: math.Max(0, c.MaxHeight-vertical),
	}
}

// Biggest returns the biggest size that satisfies the constraints
func (c *Constraints) Biggest() *Size {
	return &Size{
		Width:  c.MaxWidth,
		Height: c.MaxHeight,
	}
}

// Smallest returns the smallest size that satisfies the constraints
func (c *Constraints) Smallest() *Size {
	return &Size{
		Width:  c.MinWidth,
		Height: c.MinHeight,
	}
}

// clamp restricts a value to be within min and max
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
