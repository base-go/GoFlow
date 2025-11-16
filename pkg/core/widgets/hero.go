package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/animation"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// Hero widget for hero animations between routes
type Hero struct {
	goflow.BaseWidget
	Tag   string        // Unique tag to match heroes across routes
	Child goflow.Widget // The widget to animate
}

// NewHero creates a new Hero widget
func NewHero(tag string, child goflow.Widget) *Hero {
	return &Hero{
		Tag:   tag,
		Child: child,
	}
}

// Build creates the widget tree
func (h *Hero) Build(context goflow.BuildContext) goflow.Widget {
	// In a real implementation, this would register with the HeroController
	// For now, just return the child
	return h.Child
}

// HeroController manages hero animations
type HeroController struct {
	heroes map[string]*HeroFlight
}

// NewHeroController creates a new hero controller
func NewHeroController() *HeroController {
	return &HeroController{
		heroes: make(map[string]*HeroFlight),
	}
}

// HeroFlight represents an in-flight hero animation
type HeroFlight struct {
	tag             string
	fromWidget      goflow.Widget
	toWidget        goflow.Widget
	fromRect        *goflow.Rect
	toRect          *goflow.Rect
	animation       *animation.AnimationController
	curvedAnimation *animation.CurvedAnimation
}

// StartFlight starts a hero animation
func (hc *HeroController) StartFlight(tag string, fromWidget, toWidget goflow.Widget, fromRect, toRect *goflow.Rect) {
	controller := animation.NewAnimationController(300 * 1000000) // 300ms in nanoseconds
	curvedAnim := animation.NewCurvedAnimation(controller, animation.EaseInOut)

	flight := &HeroFlight{
		tag:             tag,
		fromWidget:      fromWidget,
		toWidget:        toWidget,
		fromRect:        fromRect,
		toRect:          toRect,
		animation:       controller,
		curvedAnimation: curvedAnim,
	}

	hc.heroes[tag] = flight
	controller.Forward()

	// Clean up after animation completes
	controller.AddStatusListener(func(status animation.AnimationStatus) {
		if status == animation.AnimationStatusCompleted {
			delete(hc.heroes, tag)
			controller.Dispose()
		}
	})
}

// PageTransition defines a page transition animation
type PageTransition struct {
	Type     PageTransitionType
	Duration int64 // Duration in nanoseconds
	Curve    animation.Curve
}

// PageTransitionType defines the type of page transition
type PageTransitionType int

const (
	PageTransitionFade PageTransitionType = iota
	PageTransitionSlideRight
	PageTransitionSlideLeft
	PageTransitionSlideUp
	PageTransitionSlideDown
	PageTransitionZoom
	PageTransitionRotation
	PageTransitionScale
	PageTransitionMaterial   // Material Design transition
	PageTransitionCupertino  // iOS-style transition
)

// NewPageTransition creates a new page transition
func NewPageTransition(transitionType PageTransitionType, duration int64) *PageTransition {
	return &PageTransition{
		Type:     transitionType,
		Duration: duration,
		Curve:    animation.EaseInOut,
	}
}

// PageRouteBuilder builds a page route with transitions
type PageRouteBuilder struct {
	builder     func(goflow.BuildContext) goflow.Widget
	transition  *PageTransition
	maintainState bool
	fullscreenDialog bool
}

// NewPageRouteBuilder creates a new page route builder
func NewPageRouteBuilder(builder func(goflow.BuildContext) goflow.Widget) *PageRouteBuilder {
	return &PageRouteBuilder{
		builder:     builder,
		transition:  NewPageTransition(PageTransitionFade, 300*1000000), // Default 300ms fade
		maintainState: true,
		fullscreenDialog: false,
	}
}

// WithTransition sets the transition animation
func (p *PageRouteBuilder) WithTransition(transition *PageTransition) *PageRouteBuilder {
	p.transition = transition
	return p
}

// WithMaintainState sets whether to maintain state
func (p *PageRouteBuilder) WithMaintainState(maintain bool) *PageRouteBuilder {
	p.maintainState = maintain
	return p
}

// WithFullscreenDialog sets whether this is a fullscreen dialog
func (p *PageRouteBuilder) WithFullscreenDialog(fullscreen bool) *PageRouteBuilder {
	p.fullscreenDialog = fullscreen
	return p
}

// Build builds the page widget
func (p *PageRouteBuilder) Build(context goflow.BuildContext) goflow.Widget {
	return p.builder(context)
}

// BuildTransition builds the transition widget
func (p *PageRouteBuilder) BuildTransition(
	context goflow.BuildContext,
	animation animation.Animation,
	secondaryAnimation animation.Animation,
	child goflow.Widget,
) goflow.Widget {
	switch p.transition.Type {
	case PageTransitionFade:
		return NewFadeTransition(animation, child)

	case PageTransitionSlideRight:
		offsetTween := animation.NewTween(-1.0, 0.0)
		return &SlideTransition{
			Position: &tweenAnimation{
				parent: animation,
				tween:  offsetTween,
			},
			Child: child,
		}

	case PageTransitionSlideLeft:
		offsetTween := animation.NewTween(1.0, 0.0)
		return &SlideTransition{
			Position: &tweenAnimation{
				parent: animation,
				tween:  offsetTween,
			},
			Child: child,
		}

	case PageTransitionSlideUp:
		return &SlideTransition{
			Position: &tweenAnimation{
				parent: animation,
				tween:  animation.NewTween(1.0, 0.0),
			},
			Child: child,
		}

	case PageTransitionSlideDown:
		return &SlideTransition{
			Position: &tweenAnimation{
				parent: animation,
				tween:  animation.NewTween(-1.0, 0.0),
			},
			Child: child,
		}

	case PageTransitionZoom:
		return NewScaleTransition(&tweenAnimation{
			parent: animation,
			tween:  animation.NewTween(0.0, 1.0),
		}, child)

	case PageTransitionScale:
		return NewScaleTransition(&tweenAnimation{
			parent: animation,
			tween:  animation.NewTween(0.8, 1.0),
		}, child)

	case PageTransitionRotation:
		return NewRotationTransition(&tweenAnimation{
			parent: animation,
			tween:  animation.NewTween(0.0, 1.0),
		}, child)

	case PageTransitionMaterial:
		// Material Design transition (fade + scale)
		return NewFadeTransition(animation, NewScaleTransition(&tweenAnimation{
			parent: animation,
			tween:  animation.NewTween(0.8, 1.0),
		}, child))

	case PageTransitionCupertino:
		// iOS-style transition (slide from right)
		offsetTween := animation.NewTween(-1.0, 0.0)
		return &SlideTransition{
			Position: &tweenAnimation{
				parent: animation,
				tween:  offsetTween,
			},
			Child: child,
		}

	default:
		return NewFadeTransition(animation, child)
	}
}

// tweenAnimation wraps an animation with a tween
type tweenAnimation struct {
	parent animation.Animation
	tween  *animation.Tween
}

// Value returns the tweened value
func (t *tweenAnimation) Value() float64 {
	return t.tween.Lerp(t.parent.Value())
}

// AddListener adds a listener
func (t *tweenAnimation) AddListener(listener func()) {
	t.parent.AddListener(listener)
}

// RemoveListener removes a listener
func (t *tweenAnimation) RemoveListener(listener func()) {
	t.parent.RemoveListener(listener)
}

// AddStatusListener adds a status listener
func (t *tweenAnimation) AddStatusListener(listener func(animation.AnimationStatus)) {
	t.parent.AddStatusListener(listener)
}

// RemoveStatusListener removes a status listener
func (t *tweenAnimation) RemoveStatusListener(listener func(animation.AnimationStatus)) {
	t.parent.RemoveStatusListener(listener)
}

// MaterialPageRoute is a Material Design page route
func MaterialPageRoute(builder func(goflow.BuildContext) goflow.Widget) *PageRouteBuilder {
	return NewPageRouteBuilder(builder).WithTransition(
		NewPageTransition(PageTransitionMaterial, 300*1000000),
	)
}

// CupertinoPageRoute is an iOS-style page route
func CupertinoPageRoute(builder func(goflow.BuildContext) goflow.Widget) *PageRouteBuilder {
	return NewPageRouteBuilder(builder).WithTransition(
		NewPageTransition(PageTransitionCupertino, 350*1000000),
	)
}

// FadePageRoute is a simple fade transition
func FadePageRoute(builder func(goflow.BuildContext) goflow.Widget) *PageRouteBuilder {
	return NewPageRouteBuilder(builder).WithTransition(
		NewPageTransition(PageTransitionFade, 200*1000000),
	)
}
