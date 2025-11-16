package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Wrap displays children in a flow layout, wrapping to the next line when needed
type Wrap struct {
	goflow.BaseWidget
	Children       []goflow.Widget
	Direction      WrapDirection
	Alignment      WrapAlignment
	Spacing        float64
	RunSpacing     float64
	CrossAlignment WrapCrossAlignment
}

// WrapDirection defines the direction of the wrap
type WrapDirection int

const (
	WrapDirectionHorizontal WrapDirection = iota
	WrapDirectionVertical
)

// WrapAlignment defines how runs are aligned
type WrapAlignment int

const (
	WrapAlignmentStart WrapAlignment = iota
	WrapAlignmentEnd
	WrapAlignmentCenter
	WrapAlignmentSpaceBetween
	WrapAlignmentSpaceAround
	WrapAlignmentSpaceEvenly
)

// WrapCrossAlignment defines how children align within a run
type WrapCrossAlignment int

const (
	WrapCrossAlignmentStart WrapCrossAlignment = iota
	WrapCrossAlignmentEnd
	WrapCrossAlignmentCenter
)

// NewWrap creates a new Wrap widget
func NewWrap(children []goflow.Widget) *Wrap {
	return &Wrap{
		Children:       children,
		Direction:      WrapDirectionHorizontal,
		Alignment:      WrapAlignmentStart,
		Spacing:        0,
		RunSpacing:     0,
		CrossAlignment: WrapCrossAlignmentStart,
	}
}

// CreateElement creates a render object element for Wrap
func (w *Wrap) CreateElement() goflow.Element {
	return &MultiChildRenderObjectElement{
		widget: w,
		renderObject: &RenderWrap{
			BaseRenderBox:  goflow.NewBaseRenderBox(),
			direction:      w.Direction,
			alignment:      w.Alignment,
			spacing:        w.Spacing,
			runSpacing:     w.RunSpacing,
			crossAlignment: w.CrossAlignment,
			children:       make([]goflow.RenderObject, 0),
		},
		children: make([]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (w *Wrap) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// GetChildren returns the children
func (w *Wrap) GetChildren() []goflow.Widget {
	return w.Children
}

// RenderWrap is the render object for Wrap
type RenderWrap struct {
	*goflow.BaseRenderBox
	direction      WrapDirection
	alignment      WrapAlignment
	spacing        float64
	runSpacing     float64
	crossAlignment WrapCrossAlignment
	children       []goflow.RenderObject
}

// PerformLayout performs layout for wrap
func (r *RenderWrap) PerformLayout() {
	constraints := r.GetConstraints()

	if r.direction == WrapDirectionHorizontal {
		r.layoutHorizontal(constraints)
	} else {
		r.layoutVertical(constraints)
	}
}

func (r *RenderWrap) layoutHorizontal(constraints *goflow.Constraints) {
	// Layout children in horizontal runs
	runs := make([][]goflow.RenderObject, 0)
	currentRun := make([]goflow.RenderObject, 0)
	currentRunWidth := 0.0

	maxWidth := constraints.MaxWidth

	// First pass: layout all children and organize into runs
	for _, child := range r.children {
		childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
		child.Layout(childConstraints)
		childSize := child.GetSize()

		// Check if we need to wrap
		if len(currentRun) > 0 && currentRunWidth+r.spacing+childSize.Width > maxWidth {
			// Start new run
			runs = append(runs, currentRun)
			currentRun = make([]goflow.RenderObject, 0)
			currentRunWidth = 0
		}

		currentRun = append(currentRun, child)
		currentRunWidth += childSize.Width
		if len(currentRun) > 1 {
			currentRunWidth += r.spacing
		}
	}

	// Add last run
	if len(currentRun) > 0 {
		runs = append(runs, currentRun)
	}

	// Calculate total height
	totalHeight := 0.0
	for i, run := range runs {
		if i > 0 {
			totalHeight += r.runSpacing
		}
		maxRunHeight := 0.0
		for _, child := range run {
			if child.GetSize().Height > maxRunHeight {
				maxRunHeight = child.GetSize().Height
			}
		}
		totalHeight += maxRunHeight
	}

	r.SetSize(goflow.NewSize(maxWidth, totalHeight))

	// Position children
	currentY := 0.0
	for _, run := range runs {
		// Calculate run height
		maxRunHeight := 0.0
		for _, child := range run {
			if child.GetSize().Height > maxRunHeight {
				maxRunHeight = child.GetSize().Height
			}
		}

		// Position children in run
		currentX := 0.0
		for _, child := range run {
			childSize := child.GetSize()

			// Calculate Y position based on cross alignment
			y := currentY
			switch r.crossAlignment {
			case WrapCrossAlignmentStart:
				y = currentY
			case WrapCrossAlignmentEnd:
				y = currentY + maxRunHeight - childSize.Height
			case WrapCrossAlignmentCenter:
				y = currentY + (maxRunHeight-childSize.Height)/2
			}

			child.SetOffset(goflow.NewOffset(currentX, y))
			currentX += childSize.Width + r.spacing
		}

		currentY += maxRunHeight + r.runSpacing
	}
}

func (r *RenderWrap) layoutVertical(constraints *goflow.Constraints) {
	// Layout children in vertical runs (columns)
	runs := make([][]goflow.RenderObject, 0)
	currentRun := make([]goflow.RenderObject, 0)
	currentRunHeight := 0.0

	maxHeight := constraints.MaxHeight

	// First pass: layout all children and organize into runs
	for _, child := range r.children {
		childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
		child.Layout(childConstraints)
		childSize := child.GetSize()

		// Check if we need to wrap
		if len(currentRun) > 0 && currentRunHeight+r.spacing+childSize.Height > maxHeight {
			// Start new run
			runs = append(runs, currentRun)
			currentRun = make([]goflow.RenderObject, 0)
			currentRunHeight = 0
		}

		currentRun = append(currentRun, child)
		currentRunHeight += childSize.Height
		if len(currentRun) > 1 {
			currentRunHeight += r.spacing
		}
	}

	// Add last run
	if len(currentRun) > 0 {
		runs = append(runs, currentRun)
	}

	// Calculate total width
	totalWidth := 0.0
	for i, run := range runs {
		if i > 0 {
			totalWidth += r.runSpacing
		}
		maxRunWidth := 0.0
		for _, child := range run {
			if child.GetSize().Width > maxRunWidth {
				maxRunWidth = child.GetSize().Width
			}
		}
		totalWidth += maxRunWidth
	}

	r.SetSize(goflow.NewSize(totalWidth, maxHeight))

	// Position children
	currentX := 0.0
	for _, run := range runs {
		// Calculate run width
		maxRunWidth := 0.0
		for _, child := range run {
			if child.GetSize().Width > maxRunWidth {
				maxRunWidth = child.GetSize().Width
			}
		}

		// Position children in run
		currentY := 0.0
		for _, child := range run {
			childSize := child.GetSize()

			// Calculate X position based on cross alignment
			x := currentX
			switch r.crossAlignment {
			case WrapCrossAlignmentStart:
				x = currentX
			case WrapCrossAlignmentEnd:
				x = currentX + maxRunWidth - childSize.Width
			case WrapCrossAlignmentCenter:
				x = currentX + (maxRunWidth-childSize.Width)/2
			}

			child.SetOffset(goflow.NewOffset(x, currentY))
			currentY += childSize.Height + r.spacing
		}

		currentX += maxRunWidth + r.runSpacing
	}
}

// Layout performs the layout
func (r *RenderWrap) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the wrap
func (r *RenderWrap) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	for _, child := range r.children {
		childOffset := offset.Add(child.GetOffset())
		child.Paint(canvas, childOffset)
	}
}

// AddChild adds a child render object
func (r *RenderWrap) AddChild(child goflow.RenderObject) {
	r.children = append(r.children, child)
	child.SetParent(r)
}
