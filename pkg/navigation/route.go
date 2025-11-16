package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
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
	// TODO: Implement proper opacity animation when animation system is ready
	// For now, just return the child
	return child
}

// SlideRightTransition creates a slide from right transition
func SlideRightTransition(child goflow.Widget, animation float64) goflow.Widget {
	// TODO: Implement proper slide animation when animation system is ready
	return child
}

// SlideLeftTransition creates a slide from left transition
func SlideLeftTransition(child goflow.Widget, animation float64) goflow.Widget {
	// TODO: Implement proper slide animation
	return child
}

// SlideUpTransition creates a slide from bottom transition
func SlideUpTransition(child goflow.Widget, animation float64) goflow.Widget {
	// TODO: Implement proper slide animation
	return child
}

// SlideDownTransition creates a slide from top transition
func SlideDownTransition(child goflow.Widget, animation float64) goflow.Widget {
	// TODO: Implement proper slide animation
	return child
}

// ZoomTransition creates a zoom transition
func ZoomTransition(child goflow.Widget, animation float64) goflow.Widget {
	// TODO: Implement proper zoom animation
	return child
}

// NoTransition returns the child without any transition
func NoTransition(child goflow.Widget, animation float64) goflow.Widget {
	return child
}
