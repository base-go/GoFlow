package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Container is a convenience widget that combines common painting, positioning, and sizing widgets
type Container struct {
	goflow.BaseWidget
	Child       goflow.Widget
	Width       *float64
	Height      *float64
	Color       *goflow.Color
	Padding     *goflow.EdgeInsets
	Margin      *goflow.EdgeInsets
}

// NewContainer creates a new Container
func NewContainer() *Container {
	return &Container{
		Padding:   goflow.ZeroEdgeInsets(),
		Margin:    goflow.ZeroEdgeInsets(),
	}
}

// Build creates the widget tree for Container
func (c *Container) Build(context goflow.BuildContext) goflow.Widget {
	child := c.Child

	// Wrap with padding if specified
	if c.Padding != nil && (c.Padding.Left > 0 || c.Padding.Top > 0 || c.Padding.Right > 0 || c.Padding.Bottom > 0) {
		child = &Padding{
			Padding: c.Padding,
			Child:   child,
		}
	}

	// Wrap with color box if specified
	if c.Color != nil {
		child = &ColoredBox{
			Color: c.Color,
			Child: child,
		}
	}

	// Wrap with sized box if size specified
	if c.Width != nil || c.Height != nil {
		child = &SizedBox{
			Width:  c.Width,
			Height: c.Height,
			Child:  child,
		}
	}

	return child
}

// Padding adds padding around a child widget
type Padding struct {
	goflow.BaseWidget
	Padding *goflow.EdgeInsets
	Child   goflow.Widget
}

// NewPadding creates a new Padding widget
func NewPadding(padding *goflow.EdgeInsets, child goflow.Widget) *Padding {
	return &Padding{
		Padding: padding,
		Child:   child,
	}
}

// CreateElement creates a render object element
func (p *Padding) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      p,
		renderObject: &RenderPadding{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			padding:       p.Padding,
			child:         nil,
		},
	}
}

// Build returns the child
func (p *Padding) Build(context goflow.BuildContext) goflow.Widget {
	return p.Child
}

// RenderPadding is the render object for padding
type RenderPadding struct {
	*goflow.BaseRenderBox
	padding *goflow.EdgeInsets
	child   goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderPadding) PerformLayout() {
	if r.child != nil {
		// Deflate constraints by padding
		childConstraints := r.GetConstraints().Deflate(r.padding)
		r.child.Layout(childConstraints)

		// Set child offset
		r.child.SetOffset(goflow.NewOffset(r.padding.Left, r.padding.Top))

		// Our size is child size plus padding
		childSize := r.child.GetSize()
		r.SetSize(goflow.NewSize(
			childSize.Width+r.padding.Horizontal(),
			childSize.Height+r.padding.Vertical(),
		))
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderPadding) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the padding and child
func (r *RenderPadding) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child != nil {
		childOffset := offset.Add(r.child.GetOffset())
		r.child.Paint(canvas, childOffset)
	}
}

// ColoredBox draws a colored box
type ColoredBox struct {
	goflow.BaseWidget
	Color *goflow.Color
	Child goflow.Widget
}

// CreateElement creates a render object element
func (c *ColoredBox) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      c,
		renderObject: &RenderColoredBox{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			color:         c.Color,
		},
	}
}

// Build returns the child
func (c *ColoredBox) Build(context goflow.BuildContext) goflow.Widget {
	return c.Child
}

// RenderColoredBox is the render object for colored box
type RenderColoredBox struct {
	*goflow.BaseRenderBox
	color *goflow.Color
	child goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderColoredBox) PerformLayout() {
	if r.child != nil {
		r.child.Layout(r.GetConstraints())
		r.SetSize(r.child.GetSize())
	} else {
		r.SetSize(r.GetConstraints().Biggest())
	}
}

// Layout performs the layout
func (r *RenderColoredBox) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the colored box
func (r *RenderColoredBox) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// Draw the background color
	paint := goflow.NewPaint()
	paint.Color = r.color
	rect := goflow.NewRect(offset, r.GetSize())
	canvas.DrawRect(rect, paint)

	// Paint child
	if r.child != nil {
		r.child.Paint(canvas, offset)
	}
}

// SizedBox constrains its child to a specific size
type SizedBox struct {
	goflow.BaseWidget
	Width  *float64
	Height *float64
	Child  goflow.Widget
}

// CreateElement creates a render object element
func (s *SizedBox) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      s,
		renderObject: &RenderSizedBox{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			width:         s.Width,
			height:        s.Height,
		},
	}
}

// Build returns the child
func (s *SizedBox) Build(context goflow.BuildContext) goflow.Widget {
	return s.Child
}

// RenderSizedBox is the render object for sized box
type RenderSizedBox struct {
	*goflow.BaseRenderBox
	width  *float64
	height *float64
	child  goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderSizedBox) PerformLayout() {
	size := r.GetConstraints().Biggest()

	if r.width != nil {
		size.Width = *r.width
	}
	if r.height != nil {
		size.Height = *r.height
	}

	r.SetSize(r.GetConstraints().Constrain(size))

	if r.child != nil {
		childConstraints := goflow.TightConstraintsForSize(r.GetSize())
		r.child.Layout(childConstraints)
	}
}

// Layout performs the layout
func (r *RenderSizedBox) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the sized box
func (r *RenderSizedBox) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child != nil {
		r.child.Paint(canvas, offset)
	}
}
