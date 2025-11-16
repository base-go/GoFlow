package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/animation"
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// AnimatedBuilder builds a widget based on animation value
type AnimatedBuilder struct {
	goflow.BaseWidget
	Animation animation.Animation
	Builder   func(float64) goflow.Widget
}

// NewAnimatedBuilder creates a new animated builder
func NewAnimatedBuilder(anim animation.Animation, builder func(float64) goflow.Widget) *AnimatedBuilder {
	return &AnimatedBuilder{
		Animation: anim,
		Builder:   builder,
	}
}

// Build creates the widget tree
func (a *AnimatedBuilder) Build(context goflow.BuildContext) goflow.Widget {
	return a.Builder(a.Animation.Value())
}

// AnimatedContainer animates its container properties
type AnimatedContainer struct {
	goflow.BaseWidget
	Child            goflow.Widget
	Width            *float64
	Height           *float64
	Color            *goflow.Color
	Padding          *goflow.EdgeInsets
	Margin           *goflow.EdgeInsets
	Duration         float64 // Duration in seconds
	Curve            animation.Curve
}

// NewAnimatedContainer creates a new animated container
func NewAnimatedContainer(duration float64) *AnimatedContainer {
	return &AnimatedContainer{
		Duration: duration,
		Curve:    animation.Linear,
	}
}

// Build creates the widget tree
func (a *AnimatedContainer) Build(context goflow.BuildContext) goflow.Widget {
	// In a real implementation, this would animate between old and new values
	// For now, just return a regular container
	return &Container{
		Child:   a.Child,
		Width:   a.Width,
		Height:  a.Height,
		Color:   a.Color,
		Padding: a.Padding,
		Margin:  a.Margin,
	}
}

// AnimatedOpacity animates the opacity of its child
type AnimatedOpacity struct {
	goflow.BaseWidget
	Child    goflow.Widget
	Opacity  float64
	Duration float64
	Curve    animation.Curve
}

// NewAnimatedOpacity creates a new animated opacity
func NewAnimatedOpacity(opacity float64, duration float64) *AnimatedOpacity {
	return &AnimatedOpacity{
		Opacity:  opacity,
		Duration: duration,
		Curve:    animation.Linear,
	}
}

// Build creates the widget tree
func (a *AnimatedOpacity) Build(context goflow.BuildContext) goflow.Widget {
	return &Opacity{
		Opacity: a.Opacity,
		Child:   a.Child,
	}
}

// Opacity applies opacity to its child
type Opacity struct {
	goflow.BaseWidget
	Opacity float64
	Child   goflow.Widget
}

// NewOpacity creates a new opacity widget
func NewOpacity(opacity float64, child goflow.Widget) *Opacity {
	return &Opacity{
		Opacity: opacity,
		Child:   child,
	}
}

// CreateElement creates a render object element
func (o *Opacity) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      o,
		renderObject: &RenderOpacity{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			opacity:       o.Opacity,
		},
	}
}

// Build returns the child
func (o *Opacity) Build(context goflow.BuildContext) goflow.Widget {
	return o.Child
}

// RenderOpacity is the render object for opacity
type RenderOpacity struct {
	*goflow.BaseRenderBox
	opacity float64
	child   goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderOpacity) PerformLayout() {
	if r.child != nil {
		r.child.Layout(r.GetConstraints())
		r.SetSize(r.child.GetSize())
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderOpacity) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints with opacity
func (r *RenderOpacity) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child == nil || r.opacity <= 0 {
		return
	}

	// TODO: Implement opacity when SetOpacity is available
	// For now, just paint the child normally
	r.child.Paint(canvas, offset)
}

// AnimatedSize animates the size of its child
type AnimatedSize struct {
	goflow.BaseWidget
	Child    goflow.Widget
	Duration float64
	Curve    animation.Curve
}

// NewAnimatedSize creates a new animated size
func NewAnimatedSize(duration float64) *AnimatedSize {
	return &AnimatedSize{
		Duration: duration,
		Curve:    animation.Linear,
	}
}

// Build creates the widget tree
func (a *AnimatedSize) Build(context goflow.BuildContext) goflow.Widget {
	// In a real implementation, this would animate size changes
	return a.Child
}

// AnimatedPositioned animates the position of a Positioned widget
type AnimatedPositioned struct {
	goflow.BaseWidget
	Child    goflow.Widget
	Left     *float64
	Top      *float64
	Right    *float64
	Bottom   *float64
	Width    *float64
	Height   *float64
	Duration float64
	Curve    animation.Curve
}

// NewAnimatedPositioned creates a new animated positioned
func NewAnimatedPositioned(duration float64) *AnimatedPositioned {
	return &AnimatedPositioned{
		Duration: duration,
		Curve:    animation.Linear,
	}
}

// Build creates the widget tree
func (a *AnimatedPositioned) Build(context goflow.BuildContext) goflow.Widget {
	// In a real implementation, this would animate position changes
	return &Positioned{
		Child:  a.Child,
		Left:   a.Left,
		Top:    a.Top,
		Right:  a.Right,
		Bottom: a.Bottom,
		Width:  a.Width,
		Height: a.Height,
	}
}

// FadeTransition fades its child in or out
type FadeTransition struct {
	goflow.BaseWidget
	Opacity animation.Animation
	Child   goflow.Widget
}

// NewFadeTransition creates a new fade transition
func NewFadeTransition(opacity animation.Animation, child goflow.Widget) *FadeTransition {
	return &FadeTransition{
		Opacity: opacity,
		Child:   child,
	}
}

// Build creates the widget tree
func (f *FadeTransition) Build(context goflow.BuildContext) goflow.Widget {
	return &Opacity{
		Opacity: f.Opacity.Value(),
		Child:   f.Child,
	}
}

// ScaleTransition scales its child
type ScaleTransition struct {
	goflow.BaseWidget
	Scale animation.Animation
	Child goflow.Widget
}

// NewScaleTransition creates a new scale transition
func NewScaleTransition(scale animation.Animation, child goflow.Widget) *ScaleTransition {
	return &ScaleTransition{
		Scale: scale,
		Child: child,
	}
}

// Build creates the widget tree
func (s *ScaleTransition) Build(context goflow.BuildContext) goflow.Widget {
	scale := s.Scale.Value()
	return &Transform{
		Transform: &ScaleTransform{
			ScaleX: scale,
			ScaleY: scale,
		},
		Child: s.Child,
	}
}

// Transform applies a transformation to its child
type Transform struct {
	goflow.BaseWidget
	Transform TransformType
	Child     goflow.Widget
}

// TransformType defines a transformation
type TransformType interface {
	Apply(canvas goflow.Canvas)
}

// ScaleTransform scales the child
type ScaleTransform struct {
	ScaleX float64
	ScaleY float64
}

func (s *ScaleTransform) Apply(canvas goflow.Canvas) {
	canvas.Scale(s.ScaleX, s.ScaleY)
}

// RotationTransform rotates the child
type RotationTransform struct {
	Angle float64 // In radians
}

func (r *RotationTransform) Apply(canvas goflow.Canvas) {
	canvas.Rotate(r.Angle)
}

// TranslationTransform translates the child
type TranslationTransform struct {
	X, Y float64
}

func (t *TranslationTransform) Apply(canvas goflow.Canvas) {
	canvas.Translate(t.X, t.Y)
}

// NewTransform creates a new transform widget
func NewTransform(transform TransformType, child goflow.Widget) *Transform {
	return &Transform{
		Transform: transform,
		Child:     child,
	}
}

// CreateElement creates a render object element
func (t *Transform) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      t,
		renderObject: &RenderTransform{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			transform:     t.Transform,
		},
	}
}

// Build returns the child
func (t *Transform) Build(context goflow.BuildContext) goflow.Widget {
	return t.Child
}

// RenderTransform is the render object for transform
type RenderTransform struct {
	*goflow.BaseRenderBox
	transform TransformType
	child     goflow.RenderObject
}

// PerformLayout performs layout
func (r *RenderTransform) PerformLayout() {
	if r.child != nil {
		r.child.Layout(r.GetConstraints())
		r.SetSize(r.child.GetSize())
	} else {
		r.SetSize(r.GetConstraints().Smallest())
	}
}

// Layout performs the layout
func (r *RenderTransform) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints with transform
func (r *RenderTransform) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	if r.child == nil {
		return
	}

	canvas.Save()
	canvas.Translate(offset.X, offset.Y)
	r.transform.Apply(canvas)
	canvas.Translate(-offset.X, -offset.Y)
	r.child.Paint(canvas, offset)
	canvas.Restore()
}

// RotationTransition rotates its child
type RotationTransition struct {
	goflow.BaseWidget
	Turns animation.Animation // Rotation in turns (1.0 = 360 degrees)
	Child goflow.Widget
}

// NewRotationTransition creates a new rotation transition
func NewRotationTransition(turns animation.Animation, child goflow.Widget) *RotationTransition {
	return &RotationTransition{
		Turns: turns,
		Child: child,
	}
}

// Build creates the widget tree
func (r *RotationTransition) Build(context goflow.BuildContext) goflow.Widget {
	return &Transform{
		Transform: &RotationTransform{
			Angle: r.Turns.Value() * 2 * 3.14159, // Convert turns to radians
		},
		Child: r.Child,
	}
}

// SlideTransition slides its child
type SlideTransition struct {
	goflow.BaseWidget
	Position animation.Animation
	Child    goflow.Widget
}

// NewSlideTransition creates a new slide transition
func NewSlideTransition(position animation.Animation, child goflow.Widget) *SlideTransition {
	return &SlideTransition{
		Position: position,
		Child:    child,
	}
}

// Build creates the widget tree
func (s *SlideTransition) Build(context goflow.BuildContext) goflow.Widget {
	offset := s.Position.Value()
	return &Transform{
		Transform: &TranslationTransform{
			X: offset,
			Y: 0,
		},
		Child: s.Child,
	}
}
