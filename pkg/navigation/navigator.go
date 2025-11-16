package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Navigator manages a stack of routes and provides navigation
type Navigator struct {
	goflow.BaseWidget
	initialRoute   goflow.Widget
	onGenerateRoute func(settings *RouteSettings) *Route
	observers      []NavigatorObserver
}

// NavigatorObserver observes navigation events
type NavigatorObserver interface {
	DidPush(route *Route, previousRoute *Route)
	DidPop(route *Route, previousRoute *Route)
	DidReplace(newRoute *Route, oldRoute *Route)
}

// NewNavigator creates a new Navigator widget
func NewNavigator(initialRoute goflow.Widget) *Navigator {
	return &Navigator{
		initialRoute: initialRoute,
		observers:    []NavigatorObserver{},
	}
}

// WithGenerateRoute sets the route generator
func (n *Navigator) WithGenerateRoute(generator func(*RouteSettings) *Route) *Navigator {
	n.onGenerateRoute = generator
	return n
}

// WithObservers adds navigation observers
func (n *Navigator) WithObservers(observers ...NavigatorObserver) *Navigator {
	n.observers = append(n.observers, observers...)
	return n
}

// Build builds the navigator widget tree
func (n *Navigator) Build(context goflow.BuildContext) goflow.Widget {
	// Get the route stack from Get
	routeStack := Get.GetRouteStack()
	dialogStack := Get.GetDialogStack()
	overlayStack := Get.GetOverlayStack()

	// Initialize with initial route if stack is empty
	if len(routeStack.Get()) == 0 {
		initialRoute := NewPageRoute(n.initialRoute, TransitionFade)
		Get.push(initialRoute)
	}

	// Build the stack of pages
	return &NavigatorStack{
		routeStack:   routeStack,
		dialogStack:  dialogStack,
		overlayStack: overlayStack,
	}
}

// CreateElement creates a StatefulElement for Navigator
func (n *Navigator) CreateElement() goflow.Element {
	return goflow.NewGenericElement(n)
}

// NavigatorStack renders the stack of routes, dialogs, and overlays
type NavigatorStack struct {
	goflow.BaseWidget
	routeStack   *signals.Signal[[]*Route]
	dialogStack  *signals.Signal[[]goflow.Widget]
	overlayStack *signals.Signal[[]goflow.Widget]
}

// Build builds the navigator stack
func (ns *NavigatorStack) Build(context goflow.BuildContext) goflow.Widget {
	routes := ns.routeStack.Get()
	dialogs := ns.dialogStack.Get()
	overlays := ns.overlayStack.Get()

	// Build children stack
	children := []goflow.Widget{}

	// Add all route pages
	for _, route := range routes {
		if route.Page != nil {
			children = append(children, route.Page)
		}
	}

	// Add dialogs
	for _, dialog := range dialogs {
		children = append(children, dialog)
	}

	// Add overlays (bottom sheets, snackbars)
	for _, overlay := range overlays {
		children = append(children, overlay)
	}

	// Use Stack to layer all children
	if len(children) == 0 {
		return widgets.NewContainer()
	}

	return widgets.NewStack(children...)
}

// CreateElement creates the element
func (ns *NavigatorStack) CreateElement() goflow.Element {
	return goflow.NewGenericElement(ns)
}

// GetApp wraps the app with Navigator
func GetApp(home goflow.Widget, routes map[string]RouteBuilder) goflow.Widget {
	// Register named routes
	if routes != nil {
		Get.RegisterRoutes(routes)
	}

	// Create navigator with home page
	return NewNavigator(home)
}

// GetMaterialApp creates a Material Design app with navigation
func GetMaterialApp(home goflow.Widget, routes map[string]RouteBuilder, title string) goflow.Widget {
	// Register named routes
	if routes != nil {
		Get.RegisterRoutes(routes)
	}

	// Create navigator wrapped in material app structure
	navigator := NewNavigator(home)

	// In a full implementation, this would wrap with MaterialApp
	// For now, just return the navigator
	return navigator
}

// GetCupertinoApp creates a Cupertino app with navigation
func GetCupertinoApp(home goflow.Widget, routes map[string]RouteBuilder, title string) goflow.Widget {
	// Register named routes
	if routes != nil {
		Get.RegisterRoutes(routes)
	}

	// Create navigator wrapped in cupertino app structure
	navigator := NewNavigator(home)

	// In a full implementation, this would wrap with CupertinoApp
	// For now, just return the navigator
	return navigator
}
