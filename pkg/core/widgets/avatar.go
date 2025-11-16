package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// CircleAvatar displays a circle representing a user
type CircleAvatar struct {
	goflow.BaseWidget
	Child           goflow.Widget
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color
	Radius          float64
	MinRadius       float64
	MaxRadius       float64
	BackgroundImage ImageSource
}

// NewCircleAvatar creates a new CircleAvatar widget
func NewCircleAvatar() *CircleAvatar {
	return &CircleAvatar{
		Radius: 20.0,
	}
}

// Build creates the widget tree for CircleAvatar
func (c *CircleAvatar) Build(context goflow.BuildContext) goflow.Widget {
	// Determine effective radius
	radius := c.Radius
	if c.MinRadius > 0 && radius < c.MinRadius {
		radius = c.MinRadius
	}
	if c.MaxRadius > 0 && radius > c.MaxRadius {
		radius = c.MaxRadius
	}

	// Create circular container
	diameter := radius * 2

	// Build content
	var content goflow.Widget
	if c.BackgroundImage != nil {
		// Use background image
		content = &Image{
			Source: c.BackgroundImage,
			Width:  &diameter,
			Height: &diameter,
			Fit:    ImageFitCover,
		}
		if c.Child != nil {
			// Overlay child on image
			content = NewStack([]goflow.Widget{
				content,
				&Center{Child: c.Child},
			})
		}
	} else if c.Child != nil {
		// Use child as content
		content = &Center{Child: c.Child}
	}

	backgroundColor := c.BackgroundColor
	if backgroundColor == nil {
		backgroundColor = goflow.NewColor(189, 189, 189, 255) // Default gray
	}

	// Create circular clipped container
	return &ClipOval{
		Child: &Container{
			Width:  &diameter,
			Height: &diameter,
			Color:  backgroundColor,
			Child:  content,
		},
	}
}

// ClipOval clips its child to an oval shape
type ClipOval struct {
	goflow.BaseWidget
	Child goflow.Widget
}

// NewClipOval creates a new ClipOval widget
func NewClipOval(child goflow.Widget) *ClipOval {
	return &ClipOval{
		Child: child,
	}
}

// CreateElement creates a render object element
func (c *ClipOval) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      c,
		renderObject: &RenderClipOval{
			BaseRenderBox: goflow.NewBaseRenderBox(),
		},
	}
}

// Build returns the child
func (c *ClipOval) Build(context goflow.BuildContext) goflow.Widget {
	return c.Child
}

// RenderClipOval is the render object for clip oval
type RenderClipOval struct {
	*goflow.BaseRenderBox
	child goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderClipOval) PerformLayout() {
	if r.child != nil {
		r.child.Layout(r.GetConstraints())
		r.SetSize(r.child.GetSize())
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderClipOval) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the clipped oval
func (r *RenderClipOval) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// TODO: Implement circular clipping when ClipOval is available
	// For now, just paint the child without clipping
	if r.child != nil {
		r.child.Paint(canvas, offset)
	}
}

// ClipRRect clips its child to a rounded rectangle
type ClipRRect struct {
	goflow.BaseWidget
	Child        goflow.Widget
	BorderRadius float64
}

// NewClipRRect creates a new ClipRRect widget
func NewClipRRect(child goflow.Widget, borderRadius float64) *ClipRRect {
	return &ClipRRect{
		Child:        child,
		BorderRadius: borderRadius,
	}
}

// CreateElement creates a render object element
func (c *ClipRRect) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      c,
		renderObject: &RenderClipRRect{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			borderRadius:  c.BorderRadius,
		},
	}
}

// Build returns the child
func (c *ClipRRect) Build(context goflow.BuildContext) goflow.Widget {
	return c.Child
}

// RenderClipRRect is the render object for clip rounded rect
type RenderClipRRect struct {
	*goflow.BaseRenderBox
	child        goflow.RenderObject
	borderRadius float64
}

// PerformLayout performs layout
func (r *RenderClipRRect) PerformLayout() {
	if r.child != nil {
		r.child.Layout(r.GetConstraints())
		r.SetSize(r.child.GetSize())
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderClipRRect) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the clipped rounded rect
func (r *RenderClipRRect) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// TODO: Implement rounded rectangle clipping when ClipRRect is available
	// For now, just paint the child without clipping
	if r.child != nil {
		r.child.Paint(canvas, offset)
	}
}
