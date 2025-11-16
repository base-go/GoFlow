package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Text displays a string of text with a single style
type Text struct {
	goflow.BaseWidget
	Data  string
	Style *goflow.TextStyle
}

// NewText creates a new Text widget
func NewText(data string) *Text {
	return &Text{
		Data:  data,
		Style: goflow.NewTextStyle(),
	}
}

// NewTextWithStyle creates a new Text widget with a custom style
func NewTextWithStyle(data string, style *goflow.TextStyle) *Text {
	return &Text{
		Data:  data,
		Style: style,
	}
}

// Build creates a RenderObjectWidget for text
func (t *Text) Build(context goflow.BuildContext) goflow.Widget {
	return &TextRenderWidget{
		Data:  t.Data,
		Style: t.Style,
	}
}

// TextRenderWidget is the render object widget for Text
type TextRenderWidget struct {
	goflow.BaseWidget
	Data  string
	Style *goflow.TextStyle
}

// CreateElement creates a RenderObjectElement
func (w *TextRenderWidget) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      w,
		renderObject: &RenderText{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			text:          w.Data,
			style:         w.Style,
		},
	}
}

// Build returns nil (this is a leaf widget)
func (w *TextRenderWidget) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderText is the render object for text
type RenderText struct {
	*goflow.BaseRenderBox
	text  string
	style *goflow.TextStyle
}

// ComputeSize computes the size for text
func (r *RenderText) ComputeSize(constraints *goflow.Constraints) *goflow.Size {
	// Simplified text sizing - in a real implementation, this would
	// measure the actual text using the font and size
	width := float64(len(r.text)) * r.style.FontSize * 0.6
	height := r.style.FontSize * 1.2

	return constraints.Constrain(goflow.NewSize(width, height))
}

// PerformLayout performs layout
func (r *RenderText) PerformLayout() {
	r.SetSize(r.ComputeSize(r.GetConstraints()))
}

// Layout performs the layout
func (r *RenderText) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the text
func (r *RenderText) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	canvas.DrawText(r.text, offset, r.style)
}

// RenderObjectElement is an element that has a render object
type RenderObjectElement struct {
	goflow.BaseElement
	widget       goflow.Widget
	renderObject goflow.RenderObject
}

// GetWidgetType returns the widget type
func (e *RenderObjectElement) GetWidgetType() string {
	return "RenderObjectWidget"
}

// GetElementType returns the element type
func (e *RenderObjectElement) GetElementType() string {
	return "RenderObjectElement"
}

// GetRenderObject returns the render object
func (e *RenderObjectElement) GetRenderObject() goflow.RenderObject {
	return e.renderObject
}

// Mount mounts the element
func (e *RenderObjectElement) Mount(parent goflow.Element, slot interface{}) {
	e.SetParent(parent)
	e.renderObject.MarkNeedsLayout()
}

// Update updates the element
func (e *RenderObjectElement) Update(newWidget goflow.Widget) {
	e.widget = newWidget
	e.renderObject.MarkNeedsLayout()
	e.renderObject.MarkNeedsPaint()
}

// Unmount unmounts the element
func (e *RenderObjectElement) Unmount() {
	// Cleanup
}

// Rebuild rebuilds the element (no-op for render object elements)
func (e *RenderObjectElement) Rebuild() {
	// Render object elements don't rebuild
}
