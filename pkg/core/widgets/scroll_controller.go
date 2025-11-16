package widgets

import (
	"math"

	"github.com/base-go/GoFlow/pkg/core/framework"
)

// ScrollController manages scrolling for scrollable widgets
type ScrollController struct {
	// Current scroll offset
	offset float64

	// Min/max scroll extent
	minScrollExtent float64
	maxScrollExtent float64

	// Viewport dimensions
	viewportDimension float64

	// Initial scroll offset
	initialScrollOffset float64

	// Listeners
	listeners []func()
}

// NewScrollController creates a new scroll controller
func NewScrollController() *ScrollController {
	return &ScrollController{
		offset:              0,
		minScrollExtent:     0,
		maxScrollExtent:     0,
		viewportDimension:   0,
		initialScrollOffset: 0,
		listeners:           make([]func(), 0),
	}
}

// NewScrollControllerWithOffset creates a scroll controller with initial offset
func NewScrollControllerWithOffset(initialOffset float64) *ScrollController {
	return &ScrollController{
		offset:              initialOffset,
		minScrollExtent:     0,
		maxScrollExtent:     0,
		viewportDimension:   0,
		initialScrollOffset: initialOffset,
		listeners:           make([]func(), 0),
	}
}

// Offset returns the current scroll offset
func (s *ScrollController) Offset() float64 {
	return s.offset
}

// Position returns scroll position metrics
func (s *ScrollController) Position() *ScrollPosition {
	return &ScrollPosition{
		Pixels:            s.offset,
		MinScrollExtent:   s.minScrollExtent,
		MaxScrollExtent:   s.maxScrollExtent,
		ViewportDimension: s.viewportDimension,
	}
}

// JumpTo instantly jumps to the given offset
func (s *ScrollController) JumpTo(offset float64) {
	s.setOffset(offset)
}

// AnimateTo animates to the given offset (simplified - instant for now)
func (s *ScrollController) AnimateTo(offset float64, duration float64) {
	// In a real implementation, this would animate over time
	s.setOffset(offset)
}

// setOffset sets the scroll offset, clamping to valid range
func (s *ScrollController) setOffset(offset float64) {
	// Clamp to valid range
	offset = math.Max(s.minScrollExtent, math.Min(s.maxScrollExtent, offset))

	if s.offset != offset {
		s.offset = offset
		s.notifyListeners()
	}
}

// UpdateExtents updates the scroll extents
func (s *ScrollController) UpdateExtents(minExtent, maxExtent, viewportDimension float64) {
	s.minScrollExtent = minExtent
	s.maxScrollExtent = maxExtent
	s.viewportDimension = viewportDimension

	// Clamp current offset to new extents
	s.setOffset(s.offset)
}

// AddListener adds a change listener
func (s *ScrollController) AddListener(listener func()) {
	s.listeners = append(s.listeners, listener)
}

// RemoveListener removes a change listener
func (s *ScrollController) RemoveListener(listener func()) {
	for i, l := range s.listeners {
		if &l == &listener {
			s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
			return
		}
	}
}

func (s *ScrollController) notifyListeners() {
	for _, listener := range s.listeners {
		listener()
	}
}

// Dispose cleans up the controller
func (s *ScrollController) Dispose() {
	s.listeners = nil
}

// ScrollPosition represents the scroll position
type ScrollPosition struct {
	Pixels            float64
	MinScrollExtent   float64
	MaxScrollExtent   float64
	ViewportDimension float64
}

// AtEdge returns true if at the start or end edge
func (p *ScrollPosition) AtEdge() bool {
	return p.Pixels <= p.MinScrollExtent || p.Pixels >= p.MaxScrollExtent
}

// OutOfRange returns true if the position is out of range
func (p *ScrollPosition) OutOfRange() bool {
	return p.Pixels < p.MinScrollExtent || p.Pixels > p.MaxScrollExtent
}

// Extents returns the total scrollable extent
func (p *ScrollPosition) Extents() float64 {
	return p.MaxScrollExtent - p.MinScrollExtent
}

// ScrollableState manages the state of a scrollable widget
type ScrollableState struct {
	controller *ScrollController
	axis       Axis
}

// Scrollable is a render object widget that provides scrolling
type Scrollable struct {
	goflow.BaseWidget
	Child      goflow.Widget
	Controller *ScrollController
	Axis       Axis
}

// CreateElement creates a render object element
func (s *Scrollable) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      s,
		renderObject: &RenderScrollable{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			controller:    s.Controller,
			axis:          s.Axis,
		},
	}
}

// Build returns the child
func (s *Scrollable) Build(context goflow.BuildContext) goflow.Widget {
	return s.Child
}

// RenderScrollable is the render object for scrollable
type RenderScrollable struct {
	*goflow.BaseRenderBox
	controller *ScrollController
	axis       Axis
	child      goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderScrollable) PerformLayout() {
	if r.child == nil {
		r.SetSize(r.GetConstraints().Smallest())
		return
	}

	constraints := r.GetConstraints()

	// Create unbounded constraints in scroll direction
	var childConstraints *goflow.Constraints
	if r.axis == AxisVertical {
		childConstraints = goflow.NewConstraints(
			constraints.MinWidth,
			constraints.MaxWidth,
			0,
			math.MaxFloat64,
		)
	} else {
		childConstraints = goflow.NewConstraints(
			0,
			math.MaxFloat64,
			constraints.MinHeight,
			constraints.MaxHeight,
		)
	}

	r.child.Layout(childConstraints)
	childSize := r.child.GetSize()

	// Set our size to constraints
	r.SetSize(constraints.Biggest())

	// Update scroll controller extents
	size := r.GetSize()
	var maxExtent float64
	var viewportDimension float64

	if r.axis == AxisVertical {
		maxExtent = math.Max(0, childSize.Height-size.Height)
		viewportDimension = size.Height
	} else {
		maxExtent = math.Max(0, childSize.Width-size.Width)
		viewportDimension = size.Width
	}

	r.controller.UpdateExtents(0, maxExtent, viewportDimension)

	// Position child based on scroll offset
	offset := -r.controller.Offset()
	if r.axis == AxisVertical {
		r.child.SetOffset(goflow.NewOffset(0, offset))
	} else {
		r.child.SetOffset(goflow.NewOffset(offset, 0))
	}
}

// Layout performs the layout
func (r *RenderScrollable) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the scrollable
func (r *RenderScrollable) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child == nil {
		return
	}

	size := r.GetSize()

	// Save canvas state and clip to viewport
	canvas.Save()
	canvas.ClipRect(goflow.NewRect(offset, size))

	// Paint child
	childOffset := offset.Add(r.child.GetOffset())
	r.child.Paint(canvas, childOffset)

	// Restore canvas
	canvas.Restore()
}
