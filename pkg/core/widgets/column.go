package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Column displays children vertically
type Column struct {
	goflow.BaseWidget
	Children         []goflow.Widget
	MainAxisAlign    MainAxisAlignment
	CrossAxisAlign   CrossAxisAlignment
	MainAxisSize     MainAxisSize
}

// MainAxisAlignment defines how children are aligned along the main axis
type MainAxisAlignment int

const (
	MainAxisStart MainAxisAlignment = iota
	MainAxisEnd
	MainAxisCenter
	MainAxisSpaceBetween
	MainAxisSpaceAround
	MainAxisSpaceEvenly
)

// CrossAxisAlignment defines how children are aligned along the cross axis
type CrossAxisAlignment int

const (
	CrossAxisStart CrossAxisAlignment = iota
	CrossAxisEnd
	CrossAxisCenter
	CrossAxisStretch
)

// MainAxisSize defines how much space the flex should occupy
type MainAxisSize int

const (
	MainAxisSizeMin MainAxisSize = iota
	MainAxisSizeMax
)

// NewColumn creates a new Column widget
func NewColumn(children []goflow.Widget) *Column {
	return &Column{
		Children:      children,
		MainAxisAlign: MainAxisStart,
		CrossAxisAlign: CrossAxisCenter,
		MainAxisSize:  MainAxisSizeMax,
	}
}

// CreateElement creates a render object element for Column
func (c *Column) CreateElement() goflow.Element {
	return &MultiChildRenderObjectElement{
		widget: c,
		renderObject: &RenderColumn{
			BaseRenderBox:  goflow.NewBaseRenderBox(),
			mainAxisAlign:  c.MainAxisAlign,
			crossAxisAlign: c.CrossAxisAlign,
			mainAxisSize:   c.MainAxisSize,
			children:       make([]goflow.RenderObject, 0),
		},
		children: make([]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (c *Column) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// GetChildren returns the children
func (c *Column) GetChildren() []goflow.Widget {
	return c.Children
}

// RenderColumn is the render object for Column
type RenderColumn struct {
	*goflow.BaseRenderBox
	mainAxisAlign  MainAxisAlignment
	crossAxisAlign CrossAxisAlignment
	mainAxisSize   MainAxisSize
	children       []goflow.RenderObject
}

// PerformLayout performs layout for column
func (r *RenderColumn) PerformLayout() {
	constraints := r.GetConstraints()

	// Layout children and calculate total height
	totalHeight := 0.0
	maxWidth := 0.0

	for _, child := range r.children {
		childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
		child.Layout(childConstraints)

		childSize := child.GetSize()
		totalHeight += childSize.Height
		if childSize.Width > maxWidth {
			maxWidth = childSize.Width
		}
	}

	// Determine our size
	width := constraints.MaxWidth
	if r.mainAxisSize == MainAxisSizeMin {
		width = maxWidth
	}
	width = constraints.ConstrainWidth(width)

	height := totalHeight
	if r.mainAxisSize == MainAxisSizeMax {
		height = constraints.MaxHeight
	}
	height = constraints.ConstrainHeight(height)

	r.SetSize(goflow.NewSize(width, height))

	// Position children
	availableSpace := height - totalHeight
	spacing := 0.0
	currentY := 0.0

	switch r.mainAxisAlign {
	case MainAxisStart:
		currentY = 0
	case MainAxisEnd:
		currentY = availableSpace
	case MainAxisCenter:
		currentY = availableSpace / 2
	case MainAxisSpaceBetween:
		if len(r.children) > 1 {
			spacing = availableSpace / float64(len(r.children)-1)
		}
	case MainAxisSpaceAround:
		spacing = availableSpace / float64(len(r.children))
		currentY = spacing / 2
	case MainAxisSpaceEvenly:
		spacing = availableSpace / float64(len(r.children)+1)
		currentY = spacing
	}

	for _, child := range r.children {
		childSize := child.GetSize()

		// Calculate X position based on cross axis alignment
		x := 0.0
		switch r.crossAxisAlign {
		case CrossAxisStart:
			x = 0
		case CrossAxisEnd:
			x = width - childSize.Width
		case CrossAxisCenter:
			x = (width - childSize.Width) / 2
		case CrossAxisStretch:
			x = 0
			// In a full implementation, we'd resize the child
		}

		child.SetOffset(goflow.NewOffset(x, currentY))
		currentY += childSize.Height + spacing
	}
}

// Layout performs the layout
func (r *RenderColumn) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the column
func (r *RenderColumn) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	for _, child := range r.children {
		childOffset := offset.Add(child.GetOffset())
		child.Paint(canvas, childOffset)
	}
}

// AddChild adds a child render object
func (r *RenderColumn) AddChild(child goflow.RenderObject) {
	r.children = append(r.children, child)
	child.SetParent(r)
}

// MultiChildRenderObjectElement is an element with multiple children
type MultiChildRenderObjectElement struct {
	goflow.BaseElement
	widget       interface{}
	renderObject goflow.RenderObject
	children     []goflow.Element
}

// GetWidgetType returns the widget type
func (e *MultiChildRenderObjectElement) GetWidgetType() string {
	return "MultiChildRenderObjectWidget"
}

// GetElementType returns the element type
func (e *MultiChildRenderObjectElement) GetElementType() string {
	return "MultiChildRenderObjectElement"
}

// GetRenderObject returns the render object
func (e *MultiChildRenderObjectElement) GetRenderObject() goflow.RenderObject {
	return e.renderObject
}

// MultiChildWidget interface for widgets with multiple children
type MultiChildWidget interface {
	GetChildren() []goflow.Widget
}

// MultiChildRenderObject interface for render objects with multiple children
type MultiChildRenderObject interface {
	AddChild(child goflow.RenderObject)
}

// Mount mounts the element
func (e *MultiChildRenderObjectElement) Mount(parent goflow.Element, slot interface{}) {
	e.SetParent(parent)

	// Mount children - works for any multi-child widget
	if multiChildWidget, ok := e.widget.(MultiChildWidget); ok {
		if multiChildRender, ok := e.renderObject.(MultiChildRenderObject); ok {
			for _, childWidget := range multiChildWidget.GetChildren() {
				childElement := childWidget.CreateElement()
				e.children = append(e.children, childElement)
				childElement.Mount(e, nil)

				// Add child's render object
				if childRenderObj := childElement.GetRenderObject(); childRenderObj != nil {
					multiChildRender.AddChild(childRenderObj)
				}
			}
		}
	}

	e.renderObject.MarkNeedsLayout()
}

// Update updates the element
func (e *MultiChildRenderObjectElement) Update(newWidget goflow.Widget) {
	e.widget = newWidget
	e.renderObject.MarkNeedsLayout()
	e.renderObject.MarkNeedsPaint()
}

// Unmount unmounts the element
func (e *MultiChildRenderObjectElement) Unmount() {
	for _, child := range e.children {
		child.Unmount()
	}
	e.children = nil
}

// Rebuild rebuilds (no-op for render object elements)
func (e *MultiChildRenderObjectElement) Rebuild() {
	// No-op
}

// VisitChildren visits child elements
func (e *MultiChildRenderObjectElement) VisitChildren(visitor func(goflow.Element)) {
	for _, child := range e.children {
		visitor(child)
	}
}
