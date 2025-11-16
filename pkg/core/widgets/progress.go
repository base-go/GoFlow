package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// ProgressIndicatorType defines the type of progress indicator
type ProgressIndicatorType int

const (
	CircularProgressIndicator ProgressIndicatorType = iota
	LinearProgressIndicator
)

// ProgressIndicator displays a progress indicator (circular or linear)
type ProgressIndicator struct {
	goflow.BaseWidget
	Type  ProgressIndicatorType
	Value *float64 // nil for indeterminate, 0.0-1.0 for determinate
	Color *goflow.Color
	Size  float64 // For circular: diameter, for linear: height
}

// NewCircularProgressIndicator creates a new circular progress indicator
func NewCircularProgressIndicator() *ProgressIndicator {
	return &ProgressIndicator{
		Type:  CircularProgressIndicator,
		Value: nil, // indeterminate
		Size:  40.0,
	}
}

// NewLinearProgressIndicator creates a new linear progress indicator
func NewLinearProgressIndicator() *ProgressIndicator {
	return &ProgressIndicator{
		Type:  LinearProgressIndicator,
		Value: nil, // indeterminate
		Size:  4.0,
	}
}

// CreateElement creates a render object element
func (p *ProgressIndicator) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      p,
		renderObject: &RenderProgressIndicator{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			progressType:  p.Type,
			value:         p.Value,
			color:         p.Color,
			size:          p.Size,
		},
	}
}

// Build returns nil (this is a render object widget)
func (p *ProgressIndicator) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderProgressIndicator is the render object for progress indicator
type RenderProgressIndicator struct {
	*goflow.BaseRenderBox
	progressType ProgressIndicatorType
	value        *float64
	color        *goflow.Color
	size         float64
}

// PerformLayout performs layout
func (r *RenderProgressIndicator) PerformLayout() {
	constraints := r.GetConstraints()

	var size *goflow.Size
	if r.progressType == CircularProgressIndicator {
		// Circular is always square
		size = goflow.NewSize(r.size, r.size)
	} else {
		// Linear takes full width
		size = goflow.NewSize(constraints.MaxWidth, r.size)
	}

	r.SetSize(constraints.Constrain(size))
}

// Layout performs the layout
func (r *RenderProgressIndicator) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the progress indicator
func (r *RenderProgressIndicator) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	color := r.color
	if color == nil {
		color = goflow.NewColor(33, 150, 243, 255) // Default blue
	}

	if r.progressType == CircularProgressIndicator {
		r.paintCircular(canvas, offset, color)
	} else {
		r.paintLinear(canvas, offset, color)
	}
}

func (r *RenderProgressIndicator) paintCircular(canvas goflow.Canvas, offset *goflow.Offset, color *goflow.Color) {
	size := r.GetSize()
	centerX := offset.X + size.Width/2
	centerY := offset.Y + size.Height/2
	radius := size.Width / 2 * 0.8

	paint := goflow.NewPaint()
	paint.Color = color

	// TODO: Implement circular progress when DrawArc is available
	// For now, draw a simple circle placeholder
	canvas.DrawCircle(goflow.NewOffset(centerX, centerY), radius, paint)
}

func (r *RenderProgressIndicator) paintLinear(canvas goflow.Canvas, offset *goflow.Offset, color *goflow.Color) {
	size := r.GetSize()

	paint := goflow.NewPaint()

	// Draw background track
	bgColor := goflow.NewColor(224, 224, 224, 255) // Light gray
	paint.Color = bgColor
	canvas.DrawRect(goflow.NewRect(offset, size), paint)

	// Draw progress
	paint.Color = color
	progressWidth := size.Width
	if r.value != nil {
		progressWidth = size.Width * (*r.value)
	} else {
		// Indeterminate - show partial width (in real app would animate)
		progressWidth = size.Width * 0.3
	}

	progressRect := goflow.NewRect(offset, goflow.NewSize(progressWidth, size.Height))
	canvas.DrawRect(progressRect, paint)
}
