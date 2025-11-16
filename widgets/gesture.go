package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// GestureDetector detects gestures on its child
type GestureDetector struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Tap callbacks
	OnTap          func()
	OnDoubleTap    func()
	OnLongPress    func()

	// Drag callbacks
	OnPanStart     func(details DragStartDetails)
	OnPanUpdate    func(details DragUpdateDetails)
	OnPanEnd       func(details DragEndDetails)

	// Scale callbacks
	OnScaleStart   func(details ScaleStartDetails)
	OnScaleUpdate  func(details ScaleUpdateDetails)
	OnScaleEnd     func(details ScaleEndDetails)

	// Hit test behavior
	Behavior HitTestBehavior
}

// HitTestBehavior defines how hit testing works
type HitTestBehavior int

const (
	HitTestBehaviorDeferToChild HitTestBehavior = iota
	HitTestBehaviorOpaque
	HitTestBehaviorTranslucent
)

// DragStartDetails contains drag start information
type DragStartDetails struct {
	GlobalPosition goflow.Offset
	LocalPosition  goflow.Offset
}

// DragUpdateDetails contains drag update information
type DragUpdateDetails struct {
	GlobalPosition goflow.Offset
	LocalPosition  goflow.Offset
	Delta          goflow.Offset
}

// DragEndDetails contains drag end information
type DragEndDetails struct {
	Velocity       goflow.Offset
	PrimaryVelocity float64
}

// ScaleStartDetails contains scale start information
type ScaleStartDetails struct {
	FocalPoint goflow.Offset
}

// ScaleUpdateDetails contains scale update information
type ScaleUpdateDetails struct {
	FocalPoint   goflow.Offset
	Scale        float64
	Rotation     float64
}

// ScaleEndDetails contains scale end information
type ScaleEndDetails struct {
	Velocity float64
}

// NewGestureDetector creates a new GestureDetector
func NewGestureDetector(child goflow.Widget) *GestureDetector {
	return &GestureDetector{
		Child:    child,
		Behavior: HitTestBehaviorDeferToChild,
	}
}

// Build returns the child (gesture detection happens at element level)
func (g *GestureDetector) Build(context goflow.BuildContext) goflow.Widget {
	return g.Child
}

// InkWell is a Material Design ink splash effect with gesture detection
type InkWell struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Tap callback
	OnTap func()

	// Double tap callback
	OnDoubleTap func()

	// Long press callback
	OnLongPress func()

	// Splash color
	SplashColor *goflow.Color

	// Highlight color (when pressed)
	HighlightColor *goflow.Color

	// Border radius for splash
	BorderRadius float64
}

// NewInkWell creates a new InkWell
func NewInkWell(child goflow.Widget, onTap func()) *InkWell {
	return &InkWell{
		Child:        child,
		OnTap:        onTap,
		BorderRadius: 0.0,
	}
}

// Build creates the widget tree
func (i *InkWell) Build(context goflow.BuildContext) goflow.Widget {
	// In a real implementation, this would create a special render object
	// that draws the ink splash effect
	return &GestureDetector{
		Child:       i.Child,
		OnTap:       i.OnTap,
		OnDoubleTap: i.OnDoubleTap,
		OnLongPress: i.OnLongPress,
	}
}

// Draggable makes a widget draggable
type Draggable struct {
	goflow.BaseWidget

	// Child widget (what's shown when not dragging)
	Child goflow.Widget

	// Feedback widget (what's shown while dragging)
	Feedback goflow.Widget

	// Child when dragging
	ChildWhenDragging goflow.Widget

	// Data to pass to DragTarget
	Data interface{}

	// Callbacks
	OnDragStarted  func()
	OnDragEnd      func(details DraggableDetails)
	OnDragCompleted func()
	OnDraggableCanceled func()

	// Axis lock
	Axis *Axis

	// Max simultaneous drags
	MaxSimultaneousDrags *int
}

// DraggableDetails contains drag completion info
type DraggableDetails struct {
	Velocity goflow.Offset
	Offset   goflow.Offset
}

// NewDraggable creates a new Draggable
func NewDraggable(child goflow.Widget, data interface{}) *Draggable {
	return &Draggable{
		Child:    child,
		Feedback: child,
		Data:     data,
	}
}

// Build creates the widget tree
func (d *Draggable) Build(context goflow.BuildContext) goflow.Widget {
	// Simplified - just show the child
	return d.Child
}

// DragTarget is a widget that can accept draggable widgets
type DragTarget struct {
	goflow.BaseWidget

	// Builder function
	Builder func(context goflow.BuildContext, candidateData []interface{}, rejectedData []interface{}) goflow.Widget

	// Callbacks
	OnWillAccept func(data interface{}) bool
	OnAccept     func(data interface{})
	OnLeave      func(data interface{})
}

// NewDragTarget creates a new DragTarget
func NewDragTarget(builder func(goflow.BuildContext, []interface{}, []interface{}) goflow.Widget) *DragTarget {
	return &DragTarget{
		Builder: builder,
	}
}

// Build creates the widget tree
func (d *DragTarget) Build(context goflow.BuildContext) goflow.Widget {
	// Simplified - build with empty data
	return d.Builder(context, []interface{}{}, []interface{}{})
}
