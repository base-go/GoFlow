package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Row displays children horizontally
type Row struct {
	goflow.BaseWidget
	Children           []goflow.Widget
	MainAxisAlignment  MainAxisAlignment
	CrossAxisAlignment CrossAxisAlignment
	MainAxisSize       MainAxisSize
}

// NewRow creates a new Row widget
func NewRow(children []goflow.Widget) *Row {
	return &Row{
		Children:           children,
		MainAxisAlignment:  MainAxisStart,
		CrossAxisAlignment: CrossAxisCenter,
		MainAxisSize:       MainAxisSizeMax,
	}
}

// CreateElement creates a render object element for Row
func (r *Row) CreateElement() goflow.Element {
	return &MultiChildRenderObjectElement{
		widget: r,
		renderObject: &RenderRow{
			BaseRenderBox:  goflow.NewBaseRenderBox(),
			mainAxisAlign:  r.MainAxisAlignment,
			crossAxisAlign: r.CrossAxisAlignment,
			mainAxisSize:   r.MainAxisSize,
			children:       make([]goflow.RenderObject, 0),
		},
		children: make([]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (r *Row) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// GetChildren returns the children
func (r *Row) GetChildren() []goflow.Widget {
	return r.Children
}

// RenderRow is the render object for Row
type RenderRow struct {
	*goflow.BaseRenderBox
	mainAxisAlign  MainAxisAlignment
	crossAxisAlign CrossAxisAlignment
	mainAxisSize   MainAxisSize
	children       []goflow.RenderObject
}

// PerformLayout performs layout for row
func (r *RenderRow) PerformLayout() {
	constraints := r.GetConstraints()

	// Layout children and calculate total width
	totalWidth := 0.0
	maxHeight := 0.0

	for _, child := range r.children {
		childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
		child.Layout(childConstraints)

		childSize := child.GetSize()
		totalWidth += childSize.Width
		if childSize.Height > maxHeight {
			maxHeight = childSize.Height
		}
	}

	// Determine our size
	width := totalWidth
	if r.mainAxisSize == MainAxisSizeMax {
		width = constraints.MaxWidth
	}
	width = constraints.ConstrainWidth(width)

	height := constraints.MaxHeight
	if r.mainAxisSize == MainAxisSizeMin {
		height = maxHeight
	}
	height = constraints.ConstrainHeight(height)

	r.SetSize(goflow.NewSize(width, height))

	// Position children
	availableSpace := width - totalWidth
	spacing := 0.0
	currentX := 0.0

	switch r.mainAxisAlign {
	case MainAxisStart:
		currentX = 0
	case MainAxisEnd:
		currentX = availableSpace
	case MainAxisCenter:
		currentX = availableSpace / 2
	case MainAxisSpaceBetween:
		if len(r.children) > 1 {
			spacing = availableSpace / float64(len(r.children)-1)
		}
	case MainAxisSpaceAround:
		spacing = availableSpace / float64(len(r.children))
		currentX = spacing / 2
	case MainAxisSpaceEvenly:
		spacing = availableSpace / float64(len(r.children)+1)
		currentX = spacing
	}

	for _, child := range r.children {
		childSize := child.GetSize()

		// Calculate Y position based on cross axis alignment
		y := 0.0
		switch r.crossAxisAlign {
		case CrossAxisStart:
			y = 0
		case CrossAxisEnd:
			y = height - childSize.Height
		case CrossAxisCenter:
			y = (height - childSize.Height) / 2
		case CrossAxisStretch:
			y = 0
			// In a full implementation, we'd resize the child
		}

		child.SetOffset(goflow.NewOffset(currentX, y))
		currentX += childSize.Width + spacing
	}
}

// Layout performs the layout
func (r *RenderRow) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the row
func (r *RenderRow) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	for _, child := range r.children {
		childOffset := offset.Add(child.GetOffset())
		child.Paint(canvas, childOffset)
	}
}

// AddChild adds a child render object
func (r *RenderRow) AddChild(child goflow.RenderObject) {
	r.children = append(r.children, child)
	child.SetParent(r)
}
