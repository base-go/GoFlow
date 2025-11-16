package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Align aligns its child within itself
type Align struct {
	goflow.BaseWidget
	Child         goflow.Widget
	Alignment     AlignmentValue
	WidthFactor   *float64 // If set, width = child.width * widthFactor
	HeightFactor  *float64 // If set, height = child.height * heightFactor
}

// AlignmentValue defines how a child is aligned
type AlignmentValue struct {
	X float64 // -1.0 (left) to 1.0 (right)
	Y float64 // -1.0 (top) to 1.0 (bottom)
}

// Common alignment values
var (
	AlignmentTopLeft      = AlignmentValue{X: -1.0, Y: -1.0}
	AlignmentTopCenter    = AlignmentValue{X: 0.0, Y: -1.0}
	AlignmentTopRight     = AlignmentValue{X: 1.0, Y: -1.0}
	AlignmentCenterLeft   = AlignmentValue{X: -1.0, Y: 0.0}
	AlignmentCenter       = AlignmentValue{X: 0.0, Y: 0.0}
	AlignmentCenterRight  = AlignmentValue{X: 1.0, Y: 0.0}
	AlignmentBottomLeft   = AlignmentValue{X: -1.0, Y: 1.0}
	AlignmentBottomCenter = AlignmentValue{X: 0.0, Y: 1.0}
	AlignmentBottomRight  = AlignmentValue{X: 1.0, Y: 1.0}
)

// NewAlign creates a new Align widget
func NewAlign(alignment AlignmentValue, child goflow.Widget) *Align {
	return &Align{
		Child:     child,
		Alignment: alignment,
	}
}

// CreateElement creates a render object element
func (a *Align) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      a,
		renderObject: &RenderAlign{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			alignment:     a.Alignment,
			widthFactor:   a.WidthFactor,
			heightFactor:  a.HeightFactor,
		},
	}
}

// Build returns the child
func (a *Align) Build(context goflow.BuildContext) goflow.Widget {
	return a.Child
}

// RenderAlign is the render object for Align
type RenderAlign struct {
	*goflow.BaseRenderBox
	alignment    AlignmentValue
	widthFactor  *float64
	heightFactor *float64
	child        goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderAlign) PerformLayout() {
	if r.child != nil {
		// Layout child with loose constraints
		r.child.Layout(goflow.LooseConstraints(
			r.GetConstraints().MaxWidth,
			r.GetConstraints().MaxHeight,
		))

		childSize := r.child.GetSize()

		// Determine our size
		width := childSize.Width
		height := childSize.Height

		if r.widthFactor != nil {
			width = childSize.Width * (*r.widthFactor)
		} else {
			width = r.GetConstraints().MaxWidth
		}

		if r.heightFactor != nil {
			height = childSize.Height * (*r.heightFactor)
		} else {
			height = r.GetConstraints().MaxHeight
		}

		r.SetSize(r.GetConstraints().Constrain(goflow.NewSize(width, height)))

		// Position child based on alignment
		ourSize := r.GetSize()
		x := (ourSize.Width - childSize.Width) * (r.alignment.X + 1.0) / 2.0
		y := (ourSize.Height - childSize.Height) * (r.alignment.Y + 1.0) / 2.0

		r.child.SetOffset(goflow.NewOffset(x, y))
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderAlign) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the align and child
func (r *RenderAlign) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child != nil {
		childOffset := offset.Add(r.child.GetOffset())
		r.child.Paint(canvas, childOffset)
	}
}
