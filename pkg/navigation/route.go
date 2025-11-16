package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Route represents a navigation route
type Route struct {
	Name       string
	Page       goflow.Widget
	Transition Transition
	Settings   *RouteSettings
}

// RouteSettings contains metadata about a route
type RouteSettings struct {
	Name      string
	Arguments interface{}
}

// NewPageRoute creates a new page route
func NewPageRoute(page goflow.Widget, transition Transition) *Route {
	return &Route{
		Page:       page,
		Transition: transition,
		Settings:   &RouteSettings{},
	}
}

// NewNamedRoute creates a new named route
func NewNamedRoute(name string, page goflow.Widget, transition Transition) *Route {
	return &Route{
		Name:       name,
		Page:       page,
		Transition: transition,
		Settings: &RouteSettings{
			Name: name,
		},
	}
}

// Transition defines the type of page transition
type Transition int

const (
	// TransitionFade fades in the new page
	TransitionFade Transition = iota

	// TransitionSlideRight slides in from the right (iOS style)
	TransitionSlideRight

	// TransitionSlideLeft slides in from the left
	TransitionSlideLeft

	// TransitionSlideUp slides in from the bottom (Android style)
	TransitionSlideUp

	// TransitionSlideDown slides in from the top
	TransitionSlideDown

	// TransitionZoom zooms in the new page
	TransitionZoom

	// TransitionNone no transition
	TransitionNone

	// TransitionCupertino iOS-style slide from right
	TransitionCupertino

	// TransitionMaterial Android-style slide from bottom
	TransitionMaterial
)

// TransitionBuilder builds a transition widget
type TransitionBuilder func(child goflow.Widget, animation float64) goflow.Widget

// GetTransitionBuilder returns a transition builder for the given transition type
func GetTransitionBuilder(transition Transition) TransitionBuilder {
	switch transition {
	case TransitionFade:
		return FadeTransition
	case TransitionSlideRight:
		return SlideRightTransition
	case TransitionSlideLeft:
		return SlideLeftTransition
	case TransitionSlideUp:
		return SlideUpTransition
	case TransitionSlideDown:
		return SlideDownTransition
	case TransitionZoom:
		return ZoomTransition
	case TransitionCupertino:
		return SlideRightTransition // iOS-style
	case TransitionMaterial:
		return SlideUpTransition // Android-style
	case TransitionNone:
		return NoTransition
	default:
		return FadeTransition
	}
}

// FadeTransition creates a fade transition
func FadeTransition(child goflow.Widget, animation float64) goflow.Widget {
	return &widgets.Opacity{
		Opacity: animation,
		Child:   child,
	}
}

// SlideRightTransition creates a slide from right transition (iOS style)
func SlideRightTransition(child goflow.Widget, animation float64) goflow.Widget {
	// Slide from right to left as animation goes from 0 to 1
	offset := (1.0 - animation) * 400.0 // Screen width approximation
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: offset,
			Y: 0,
		},
		Child: child,
	}
}

// SlideLeftTransition creates a slide from left transition
func SlideLeftTransition(child goflow.Widget, animation float64) goflow.Widget {
	// Slide from left to right as animation goes from 0 to 1
	offset := -(1.0 - animation) * 400.0
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: offset,
			Y: 0,
		},
		Child: child,
	}
}

// SlideUpTransition creates a slide from bottom transition (Material style)
func SlideUpTransition(child goflow.Widget, animation float64) goflow.Widget {
	// Slide from bottom to top as animation goes from 0 to 1
	offset := (1.0 - animation) * 600.0 // Screen height approximation
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: 0,
			Y: offset,
		},
		Child: child,
	}
}

// SlideDownTransition creates a slide from top transition
func SlideDownTransition(child goflow.Widget, animation float64) goflow.Widget {
	// Slide from top to bottom as animation goes from 0 to 1
	offset := -(1.0 - animation) * 600.0
	return &widgets.Transform{
		Transform: &widgets.TranslationTransform{
			X: 0,
			Y: offset,
		},
		Child: child,
	}
}

// ZoomTransition creates a zoom transition
func ZoomTransition(child goflow.Widget, animation float64) goflow.Widget {
	// Combine scale and fade for a smooth zoom effect
	scaleChild := &widgets.Transform{
		Transform: &widgets.ScaleTransform{
			ScaleX: animation,
			ScaleY: animation,
		},
		Child: child,
	}
	return &widgets.Opacity{
		Opacity: animation,
		Child:   scaleChild,
	}
}

// NoTransition returns the child without any transition
func NoTransition(child goflow.Widget, animation float64) goflow.Widget {
	return child
}
