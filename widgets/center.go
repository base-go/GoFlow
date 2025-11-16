package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Center centers its child within itself
type Center struct {
	goflow.BaseWidget
	Child goflow.Widget
}

// NewCenter creates a new Center widget
func NewCenter(child goflow.Widget) *Center {
	return &Center{
		Child: child,
	}
}

// CreateElement creates a render object element
func (c *Center) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      c,
		renderObject: &RenderCenter{
			BaseRenderBox: goflow.NewBaseRenderBox(),
		},
	}
}

// Build returns the child
func (c *Center) Build(context goflow.BuildContext) goflow.Widget {
	return c.Child
}

// RenderCenter is the render object for centering
type RenderCenter struct {
	*goflow.BaseRenderBox
	child goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderCenter) PerformLayout() {
	if r.child != nil {
		// Layout child with loose constraints
		childConstraints := r.GetConstraints().Loosen()
		r.child.Layout(childConstraints)

		// Our size is the constraint's biggest
		r.SetSize(r.GetConstraints().Biggest())

		// Center the child
		childSize := r.child.GetSize()
		ourSize := r.GetSize()
		childX := (ourSize.Width - childSize.Width) / 2
		childY := (ourSize.Height - childSize.Height) / 2
		r.child.SetOffset(goflow.NewOffset(childX, childY))
	} else {
		r.SetSize(r.GetConstraints().Biggest())
	}
}

// Layout performs the layout
func (r *RenderCenter) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the center and child
func (r *RenderCenter) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child != nil {
		childOffset := offset.Add(r.child.GetOffset())
		r.child.Paint(canvas, childOffset)
	}
}
