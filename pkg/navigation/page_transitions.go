package navigation

import (
	"math"
	"time"

	"github.com/base-go/GoFlow/pkg/core/animation"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// PageTransitionsBuilder is a function that builds a page transition
type PageTransitionsBuilder func(
	context goflow.BuildContext,
	animation *animation.AnimationController,
	secondaryAnimation *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget

// PageTransition wraps a child widget with a page transition animation
type PageTransition struct {
	goflow.BaseWidget
	Type            PageTransitionType
	Child           goflow.Widget
	Duration        time.Duration
	ReverseDuration time.Duration
	Curve           animation.Curve
	MatchingBuilder PageTransitionsBuilder
	Alignment       animation.Alignment
	IsIOS           bool
}

// PageTransitionType defines the type of page transition
type PageTransitionType int

const (
	// PageTransitionTypeFade fades in the new page
	PageTransitionTypeFade PageTransitionType = iota

	// PageTransitionTypeRightToLeft slides from right to left
	PageTransitionTypeRightToLeft

	// PageTransitionTypeLeftToRight slides from left to right
	PageTransitionTypeLeftToRight

	// PageTransitionTypeTopToBottom slides from top to bottom
	PageTransitionTypeTopToBottom

	// PageTransitionTypeBottomToTop slides from bottom to top
	PageTransitionTypeBottomToTop

	// PageTransitionTypeScale scales the page
	PageTransitionTypeScale

	// PageTransitionTypeRotate rotates the page
	PageTransitionTypeRotate

	// PageTransitionTypeSize changes the size
	PageTransitionTypeSize

	// PageTransitionTypeRightToLeftWithFade slides right to left with fade
	PageTransitionTypeRightToLeftWithFade

	// PageTransitionTypeLeftToRightWithFade slides left to right with fade
	PageTransitionTypeLeftToRightWithFade

	// PageTransitionTypeRightToLeftJoined slides with joined animation
	PageTransitionTypeRightToLeftJoined

	// PageTransitionTypeLeftToRightJoined slides with joined animation
	PageTransitionTypeLeftToRightJoined

	// PageTransitionTypeZoom zooms in the page
	PageTransitionTypeZoom

	// PageTransitionTypeFadeIn fades in only (no fade out)
	PageTransitionTypeFadeIn

	// PageTransitionTypeRippleEffect creates a ripple effect
	PageTransitionTypeRippleEffect

	// PageTransitionTypeSlideParallax creates a parallax slide effect
	PageTransitionTypeSlideParallax
)

// NewPageTransition creates a new page transition
func NewPageTransition(transitionType PageTransitionType, child goflow.Widget) *PageTransition {
	return &PageTransition{
		Type:            transitionType,
		Child:           child,
		Duration:        300 * time.Millisecond,
		ReverseDuration: 300 * time.Millisecond,
		Curve:           animation.EaseInOut,
		Alignment:       animation.Alignment{X: 0, Y: 0},
		IsIOS:           false,
	}
}

// WithDuration sets the transition duration
func (p *PageTransition) WithDuration(duration time.Duration) *PageTransition {
	p.Duration = duration
	return p
}

// WithCurve sets the animation curve
func (p *PageTransition) WithCurve(curve animation.Curve) *PageTransition {
	p.Curve = curve
	return p
}

// WithAlignment sets the alignment for scale/zoom transitions
func (p *PageTransition) WithAlignment(alignment animation.Alignment) *PageTransition {
	p.Alignment = alignment
	return p
}

// WithMatchingBuilder sets a custom builder for matching animations
func (p *PageTransition) WithMatchingBuilder(builder PageTransitionsBuilder) *PageTransition {
	p.MatchingBuilder = builder
	return p
}

// Build creates the widget tree
func (p *PageTransition) Build(context goflow.BuildContext) goflow.Widget {
	// For now, return the child directly
	// In a full implementation, this would use AnimationController and listeners
	return p.Child
}

// GetPageTransitionBuilder returns a transition builder for the given type
func GetPageTransitionBuilder(transitionType PageTransitionType) PageTransitionsBuilder {
	switch transitionType {
	case PageTransitionTypeFade:
		return FadeTransitionBuilder
	case PageTransitionTypeRightToLeft:
		return RightToLeftTransitionBuilder
	case PageTransitionTypeLeftToRight:
		return LeftToRightTransitionBuilder
	case PageTransitionTypeTopToBottom:
		return TopToBottomTransitionBuilder
	case PageTransitionTypeBottomToTop:
		return BottomToTopTransitionBuilder
	case PageTransitionTypeScale:
		return ScaleTransitionBuilder
	case PageTransitionTypeRotate:
		return RotateTransitionBuilder
	case PageTransitionTypeSize:
		return SizeTransitionBuilder
	case PageTransitionTypeRightToLeftWithFade:
		return RightToLeftWithFadeTransitionBuilder
	case PageTransitionTypeLeftToRightWithFade:
		return LeftToRightWithFadeTransitionBuilder
	case PageTransitionTypeRightToLeftJoined:
		return RightToLeftJoinedTransitionBuilder
	case PageTransitionTypeLeftToRightJoined:
		return LeftToRightJoinedTransitionBuilder
	case PageTransitionTypeZoom:
		return ZoomTransitionBuilder
	case PageTransitionTypeFadeIn:
		return FadeInTransitionBuilder
	case PageTransitionTypeRippleEffect:
		return RippleEffectTransitionBuilder
	case PageTransitionTypeSlideParallax:
		return SlideParallaxTransitionBuilder
	default:
		return FadeTransitionBuilder
	}
}

// FadeTransitionBuilder creates a fade transition
func FadeTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	return widgets.NewFadeTransition(curvedAnimation, child)
}

// RightToLeftTransitionBuilder creates a right to left slide transition
func RightToLeftTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	// Create a tween that goes from screen width to 0
	// We'll use a simple approach with transform
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (1.0 - curvedAnimation.Value()) * 400, // Assuming screen width, adjust as needed
			Y: 0,
		},
		Child: child,
	}
}

// LeftToRightTransitionBuilder creates a left to right slide transition
func LeftToRightTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (curvedAnimation.Value() - 1.0) * 400,
			Y: 0,
		},
		Child: child,
	}
}

// TopToBottomTransitionBuilder creates a top to bottom slide transition
func TopToBottomTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: 0,
			Y: (curvedAnimation.Value() - 1.0) * 600,
		},
		Child: child,
	}
}

// BottomToTopTransitionBuilder creates a bottom to top slide transition
func BottomToTopTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: 0,
			Y: (1.0 - curvedAnimation.Value()) * 600,
		},
		Child: child,
	}
}

// ScaleTransitionBuilder creates a scale transition
func ScaleTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	return widgets.NewScaleTransition(curvedAnimation, child)
}

// RotateTransitionBuilder creates a rotation transition
func RotateTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	return widgets.NewRotationTransition(curvedAnimation, child)
}

// SizeTransitionBuilder creates a size transition
func SizeTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	scale := curvedAnimation.Value()

	return &widgets.Transform{
		Transform: &widgets.ScaleTransform{
			ScaleX: 1.0,
			ScaleY: scale,
		},
		Child: child,
	}
}

// RightToLeftWithFadeTransitionBuilder combines slide and fade
func RightToLeftWithFadeTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	slideChild := &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (1.0 - curvedAnimation.Value()) * 400,
			Y: 0,
		},
		Child: child,
	}

	return widgets.NewFadeTransition(curvedAnimation, slideChild)
}

// LeftToRightWithFadeTransitionBuilder combines slide and fade
func LeftToRightWithFadeTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	slideChild := &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (curvedAnimation.Value() - 1.0) * 400,
			Y: 0,
		},
		Child: child,
	}

	return widgets.NewFadeTransition(curvedAnimation, slideChild)
}

// RightToLeftJoinedTransitionBuilder creates a joined slide transition
// The old page slides out while the new page slides in
func RightToLeftJoinedTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	// New page slides in from right
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (1.0 - curvedAnimation.Value()) * 400,
			Y: 0,
		},
		Child: child,
	}
}

// LeftToRightJoinedTransitionBuilder creates a joined slide transition
func LeftToRightJoinedTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (curvedAnimation.Value() - 1.0) * 400,
			Y: 0,
		},
		Child: child,
	}
}

// ZoomTransitionBuilder creates a zoom transition
func ZoomTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	scale := curvedAnimation.Value()

	scaleChild := &widgets.Transform{
		Transform: &widgets.ScaleTransform{
			ScaleX: scale,
			ScaleY: scale,
		},
		Child: child,
	}

	return widgets.NewFadeTransition(curvedAnimation, scaleChild)
}

// FadeInTransitionBuilder creates a fade in transition (no fade out)
func FadeInTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.Linear)
	return widgets.NewFadeTransition(curvedAnimation, child)
}

// RippleEffectTransitionBuilder creates a ripple/circular reveal effect
func RippleEffectTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
	scale := curvedAnimation.Value()

	// Combine scale and fade for a ripple effect
	scaleChild := &widgets.Transform{
		Transform: &widgets.ScaleTransform{
			ScaleX: 0.8 + (scale * 0.2),
			ScaleY: 0.8 + (scale * 0.2),
		},
		Child: child,
	}

	return widgets.NewFadeTransition(curvedAnimation, scaleChild)
}

// SlideParallaxTransitionBuilder creates a parallax slide effect
func SlideParallaxTransitionBuilder(
	context goflow.BuildContext,
	anim *animation.AnimationController,
	secondaryAnim *animation.AnimationController,
	child goflow.Widget,
) goflow.Widget {
	curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)

	// Parallax effect: new page moves faster than old page
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: (1.0 - curvedAnimation.Value()) * 300, // Slightly less distance for parallax
			Y: 0,
		},
		Child: child,
	}
}

// CustomPageTransition allows creating custom transitions
type CustomPageTransition struct {
	goflow.BaseWidget
	Child    goflow.Widget
	Builder  PageTransitionsBuilder
	Duration time.Duration
	Curve    animation.Curve
}

// NewCustomPageTransition creates a custom page transition
func NewCustomPageTransition(
	child goflow.Widget,
	builder PageTransitionsBuilder,
	duration time.Duration,
	curve animation.Curve,
) *CustomPageTransition {
	return &CustomPageTransition{
		Child:    child,
		Builder:  builder,
		Duration: duration,
		Curve:    curve,
	}
}

// Build creates the widget tree
func (c *CustomPageTransition) Build(context goflow.BuildContext) goflow.Widget {
	// In a full implementation, this would create and manage the animation controller
	return c.Child
}

// Helper function to create a combined transform and fade transition
func createCombinedTransition(
	curvedAnimation *animation.CurvedAnimation,
	offsetX, offsetY float64,
	child goflow.Widget,
) goflow.Widget {
	slideChild := &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: offsetX,
			Y: offsetY,
		},
		Child: child,
	}

	return widgets.NewFadeTransition(curvedAnimation, slideChild)
}

// PageRouteBuilder is a convenience function to create page routes with transitions
func PageRouteBuilder(
	page goflow.Widget,
	transitionType PageTransitionType,
	duration time.Duration,
	curve animation.Curve,
) *Route {
	transition := MapPageTransitionTypeToTransition(transitionType)
	return NewPageRoute(page, transition)
}

// MapPageTransitionTypeToTransition maps PageTransitionType to Transition
func MapPageTransitionTypeToTransition(transitionType PageTransitionType) Transition {
	switch transitionType {
	case PageTransitionTypeFade:
		return TransitionFade
	case PageTransitionTypeRightToLeft:
		return TransitionSlideRight
	case PageTransitionTypeLeftToRight:
		return TransitionSlideLeft
	case PageTransitionTypeTopToBottom:
		return TransitionSlideDown
	case PageTransitionTypeBottomToTop:
		return TransitionSlideUp
	case PageTransitionTypeScale, PageTransitionTypeZoom:
		return TransitionZoom
	default:
		return TransitionFade
	}
}

// AdvancedTransitionConfig provides advanced configuration for transitions
type AdvancedTransitionConfig struct {
	Duration           time.Duration
	ReverseDuration    time.Duration
	Curve              animation.Curve
	ReverseCurve       animation.Curve
	Fullscreen         bool
	Opaque             bool
	Barrrier           bool
	BarrierColor       *goflow.Color
	BarrierDismissible bool
	MaintainState      bool
	TransitionBuilder  PageTransitionsBuilder
}

// NewAdvancedTransitionConfig creates a default advanced transition config
func NewAdvancedTransitionConfig() *AdvancedTransitionConfig {
	return &AdvancedTransitionConfig{
		Duration:           300 * time.Millisecond,
		ReverseDuration:    300 * time.Millisecond,
		Curve:              animation.EaseInOut,
		ReverseCurve:       animation.EaseInOut,
		Fullscreen:         false,
		Opaque:             true,
		Barrrier:           false,
		BarrierDismissible: true,
		MaintainState:      true,
	}
}

// TransitionDirection specifies the direction of the transition
type TransitionDirection int

const (
	TransitionDirectionUp TransitionDirection = iota
	TransitionDirectionDown
	TransitionDirectionLeft
	TransitionDirectionRight
)

// CreateDirectionalSlideTransition creates a slide transition in the specified direction
func CreateDirectionalSlideTransition(
	direction TransitionDirection,
	distance float64,
) PageTransitionsBuilder {
	return func(
		context goflow.BuildContext,
		anim *animation.AnimationController,
		secondaryAnim *animation.AnimationController,
		child goflow.Widget,
	) goflow.Widget {
		curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
		progress := curvedAnimation.Value()

		var offsetX, offsetY float64
		switch direction {
		case TransitionDirectionRight:
			offsetX = (1.0 - progress) * distance
		case TransitionDirectionLeft:
			offsetX = -(1.0 - progress) * distance
		case TransitionDirectionDown:
			offsetY = (1.0 - progress) * distance
		case TransitionDirectionUp:
			offsetY = -(1.0 - progress) * distance
		}

		return &widgets.Transform{
			Transform: &widgets.TranslationTransform{
				X: offsetX,
				Y: offsetY,
			},
			Child: child,
		}
	}
}

// CreateScaleRotateTransition creates a combined scale and rotation transition
func CreateScaleRotateTransition(
	initialScale float64,
	rotationTurns float64,
) PageTransitionsBuilder {
	return func(
		context goflow.BuildContext,
		anim *animation.AnimationController,
		secondaryAnim *animation.AnimationController,
		child goflow.Widget,
	) goflow.Widget {
		curvedAnimation := animation.NewCurvedAnimation(anim, animation.EaseInOut)
		progress := curvedAnimation.Value()

		scale := initialScale + (1.0-initialScale)*progress
		rotation := rotationTurns * (1.0 - progress)

		rotateChild := &widgets.Transform{
			Transform: &widgets.RotationTransform{
				Angle: rotation * 2 * math.Pi,
			},
			Child: child,
		}

		return &widgets.Transform{
			Transform: &widgets.ScaleTransform{
				ScaleX: scale,
				ScaleY: scale,
			},
			Child: rotateChild,
		}
	}
}
