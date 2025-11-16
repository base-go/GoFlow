package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Stack overlays children on top of each other
type Stack struct {
	goflow.BaseWidget
	Children  []goflow.Widget
	Alignment StackAlignment
	Fit       StackFit
}

// StackAlignment defines how children are aligned within the stack
type StackAlignment int

const (
	StackAlignmentTopLeft StackAlignment = iota
	StackAlignmentTopCenter
	StackAlignmentTopRight
	StackAlignmentCenterLeft
	StackAlignmentCenter
	StackAlignmentCenterRight
	StackAlignmentBottomLeft
	StackAlignmentBottomCenter
	StackAlignmentBottomRight
)

// StackFit defines how non-positioned children are sized
type StackFit int

const (
	StackFitLoose StackFit = iota // Children sized to their natural size
	StackFitExpand                // Children expanded to fill the stack
	StackFitPassthrough           // Pass parent constraints to children
)

// NewStack creates a new Stack widget
func NewStack(children []goflow.Widget) *Stack {
	return &Stack{
		Children:  children,
		Alignment: StackAlignmentTopLeft,
		Fit:       StackFitLoose,
	}
}

// CreateElement creates a render object element for Stack
func (s *Stack) CreateElement() goflow.Element {
	return &MultiChildRenderObjectElement{
		widget: s,
		renderObject: &RenderStack{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			alignment:     s.Alignment,
			fit:           s.Fit,
			children:      make([]goflow.RenderObject, 0),
		},
		children: make([]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (s *Stack) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// GetChildren returns the children
func (s *Stack) GetChildren() []goflow.Widget {
	return s.Children
}

// RenderStack is the render object for Stack
type RenderStack struct {
	*goflow.BaseRenderBox
	alignment StackAlignment
	fit       StackFit
	children  []goflow.RenderObject
}

// PerformLayout performs layout for stack
func (r *RenderStack) PerformLayout() {
	constraints := r.GetConstraints()

	// Determine stack size
	var width, height float64

	switch r.fit {
	case StackFitLoose:
		// Size to largest child
		for _, child := range r.children {
			childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
			child.Layout(childConstraints)

			childSize := child.GetSize()
			if childSize.Width > width {
				width = childSize.Width
			}
			if childSize.Height > height {
				height = childSize.Height
			}
		}
	case StackFitExpand:
		// Expand to fill constraints
		width = constraints.MaxWidth
		height = constraints.MaxHeight

		for _, child := range r.children {
			childConstraints := goflow.TightConstraintsForSize(goflow.NewSize(width, height))
			child.Layout(childConstraints)
		}
	case StackFitPassthrough:
		// Pass constraints through
		width = constraints.MaxWidth
		height = constraints.MaxHeight

		for _, child := range r.children {
			child.Layout(constraints)
		}
	}

	r.SetSize(goflow.NewSize(width, height))

	// Position children based on alignment
	stackSize := r.GetSize()
	for _, child := range r.children {
		childSize := child.GetSize()

		var x, y float64

		switch r.alignment {
		case StackAlignmentTopLeft:
			x, y = 0, 0
		case StackAlignmentTopCenter:
			x = (stackSize.Width - childSize.Width) / 2
			y = 0
		case StackAlignmentTopRight:
			x = stackSize.Width - childSize.Width
			y = 0
		case StackAlignmentCenterLeft:
			x = 0
			y = (stackSize.Height - childSize.Height) / 2
		case StackAlignmentCenter:
			x = (stackSize.Width - childSize.Width) / 2
			y = (stackSize.Height - childSize.Height) / 2
		case StackAlignmentCenterRight:
			x = stackSize.Width - childSize.Width
			y = (stackSize.Height - childSize.Height) / 2
		case StackAlignmentBottomLeft:
			x = 0
			y = stackSize.Height - childSize.Height
		case StackAlignmentBottomCenter:
			x = (stackSize.Width - childSize.Width) / 2
			y = stackSize.Height - childSize.Height
		case StackAlignmentBottomRight:
			x = stackSize.Width - childSize.Width
			y = stackSize.Height - childSize.Height
		}

		child.SetOffset(goflow.NewOffset(x, y))
	}
}

// Layout performs the layout
func (r *RenderStack) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the stack
func (r *RenderStack) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	for _, child := range r.children {
		childOffset := offset.Add(child.GetOffset())
		child.Paint(canvas, childOffset)
	}
}

// AddChild adds a child render object
func (r *RenderStack) AddChild(child goflow.RenderObject) {
	r.children = append(r.children, child)
	child.SetParent(r)
}

// Positioned positions a child within a Stack
type Positioned struct {
	goflow.BaseWidget
	Child  goflow.Widget
	Left   *float64
	Top    *float64
	Right  *float64
	Bottom *float64
	Width  *float64
	Height *float64
}

// NewPositioned creates a new Positioned widget
func NewPositioned(child goflow.Widget) *Positioned {
	return &Positioned{
		Child: child,
	}
}

// Build wraps the child with positioning information
func (p *Positioned) Build(context goflow.BuildContext) goflow.Widget {
	// In a full implementation, this would create a special render object
	// that communicates position to the parent Stack
	return p.Child
}
