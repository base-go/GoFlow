package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Divider displays a horizontal or vertical dividing line
type Divider struct {
	goflow.BaseWidget
	Height    *float64
	Thickness float64
	Color     *goflow.Color
	Indent    float64
	EndIndent float64
}

// NewDivider creates a new horizontal divider
func NewDivider() *Divider {
	return &Divider{
		Thickness: 1.0,
		Indent:    0.0,
		EndIndent: 0.0,
	}
}

// CreateElement creates a render object element
func (d *Divider) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      d,
		renderObject: &RenderDivider{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			height:        d.Height,
			thickness:     d.Thickness,
			color:         d.Color,
			indent:        d.Indent,
			endIndent:     d.EndIndent,
		},
	}
}

// Build returns nil (this is a render object widget)
func (d *Divider) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderDivider is the render object for divider
type RenderDivider struct {
	*goflow.BaseRenderBox
	height    *float64
	thickness float64
	color     *goflow.Color
	indent    float64
	endIndent float64
}

// PerformLayout performs layout
func (r *RenderDivider) PerformLayout() {
	constraints := r.GetConstraints()

	height := r.thickness
	if r.height != nil {
		height = *r.height
	}

	size := goflow.NewSize(constraints.MaxWidth, height)
	r.SetSize(constraints.Constrain(size))
}

// Layout performs the layout
func (r *RenderDivider) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the divider
func (r *RenderDivider) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	size := r.GetSize()

	color := r.color
	if color == nil {
		color = goflow.NewColor(224, 224, 224, 255) // Default gray
	}

	paint := goflow.NewPaint()
	paint.Color = color

	// Draw divider line with indents
	x := offset.X + r.indent
	width := size.Width - r.indent - r.endIndent
	y := offset.Y + (size.Height-r.thickness)/2

	dividerRect := goflow.NewRect(
		goflow.NewOffset(x, y),
		goflow.NewSize(width, r.thickness),
	)
	canvas.DrawRect(dividerRect, paint)
}

// VerticalDivider displays a vertical dividing line
type VerticalDivider struct {
	goflow.BaseWidget
	Width     *float64
	Thickness float64
	Color     *goflow.Color
	Indent    float64
	EndIndent float64
}

// NewVerticalDivider creates a new vertical divider
func NewVerticalDivider() *VerticalDivider {
	return &VerticalDivider{
		Thickness: 1.0,
		Indent:    0.0,
		EndIndent: 0.0,
	}
}

// CreateElement creates a render object element
func (v *VerticalDivider) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      v,
		renderObject: &RenderVerticalDivider{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			width:         v.Width,
			thickness:     v.Thickness,
			color:         v.Color,
			indent:        v.Indent,
			endIndent:     v.EndIndent,
		},
	}
}

// Build returns nil (this is a render object widget)
func (v *VerticalDivider) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderVerticalDivider is the render object for vertical divider
type RenderVerticalDivider struct {
	*goflow.BaseRenderBox
	width     *float64
	thickness float64
	color     *goflow.Color
	indent    float64
	endIndent float64
}

// PerformLayout performs layout
func (r *RenderVerticalDivider) PerformLayout() {
	constraints := r.GetConstraints()

	width := r.thickness
	if r.width != nil {
		width = *r.width
	}

	size := goflow.NewSize(width, constraints.MaxHeight)
	r.SetSize(constraints.Constrain(size))
}

// Layout performs the layout
func (r *RenderVerticalDivider) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the vertical divider
func (r *RenderVerticalDivider) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	size := r.GetSize()

	color := r.color
	if color == nil {
		color = goflow.NewColor(224, 224, 224, 255) // Default gray
	}

	paint := goflow.NewPaint()
	paint.Color = color

	// Draw divider line with indents
	y := offset.Y + r.indent
	height := size.Height - r.indent - r.endIndent
	x := offset.X + (size.Width-r.thickness)/2

	dividerRect := goflow.NewRect(
		goflow.NewOffset(x, y),
		goflow.NewSize(r.thickness, height),
	)
	canvas.DrawRect(dividerRect, paint)
}
