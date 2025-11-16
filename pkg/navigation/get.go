package navigation

import (
	"fmt"
	"sync"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
)

// Get is the global navigation manager (GetX-style)
// Usage: Get.To(page), Get.Back(), Get.Dialog(dialog), etc.
var Get = newGetInstance()

// GetInstance manages global navigation state
type GetInstance struct {
	mu           sync.RWMutex
	navigatorKey *GlobalKey
	routeStack   *signals.Signal[[]*Route]
	dialogStack  *signals.Signal[[]goflow.Widget]
	overlayStack *signals.Signal[[]goflow.Widget]
	namedRoutes  map[string]RouteBuilder
}

// RouteBuilder is a function that builds a widget for a route
type RouteBuilder func() goflow.Widget

func newGetInstance() *GetInstance {
	return &GetInstance{
		navigatorKey: NewGlobalKey("navigator"),
		routeStack:   signals.New([]*Route{}),
		dialogStack:  signals.New([]goflow.Widget{}),
		overlayStack: signals.New([]goflow.Widget{}),
		namedRoutes:  make(map[string]RouteBuilder),
	}
}

// To navigates to a new page
func (g *GetInstance) To(page goflow.Widget, transition ...Transition) {
	trans := TransitionFade
	if len(transition) > 0 {
		trans = transition[0]
	}

	route := NewPageRoute(page, trans)
	g.push(route)
}

// ToNamed navigates to a named route
func (g *GetInstance) ToNamed(routeName string, transition ...Transition) error {
	g.mu.RLock()
	builder, exists := g.namedRoutes[routeName]
	g.mu.RUnlock()

	if !exists {
		return fmt.Errorf("route '%s' not found", routeName)
	}

	page := builder()
	g.To(page, transition...)
	return nil
}

// Off navigates to a new page and removes the previous route
func (g *GetInstance) Off(page goflow.Widget, transition ...Transition) {
	g.Back()
	g.To(page, transition...)
}

// OffNamed navigates to a named route and removes the previous route
func (g *GetInstance) OffNamed(routeName string, transition ...Transition) error {
	g.Back()
	return g.ToNamed(routeName, transition...)
}

// OffAll removes all routes and navigates to a new page
func (g *GetInstance) OffAll(page goflow.Widget, transition ...Transition) {
	g.mu.Lock()
	g.routeStack.Set([]*Route{})
	g.mu.Unlock()
	g.To(page, transition...)
}

// OffAllNamed removes all routes and navigates to a named route
func (g *GetInstance) OffAllNamed(routeName string, transition ...Transition) error {
	g.mu.Lock()
	g.routeStack.Set([]*Route{})
	g.mu.Unlock()
	return g.ToNamed(routeName, transition...)
}

// Back navigates back to the previous page
func (g *GetInstance) Back() {
	g.pop()
}

// Until navigates back until a condition is met
func (g *GetInstance) Until(predicate func(*Route) bool) {
	for {
		stack := g.routeStack.Get()
		if len(stack) == 0 {
			break
		}

		current := stack[len(stack)-1]
		if predicate(current) {
			break
		}

		g.pop()
	}
}

// OffUntil navigates to a new page and removes routes until a condition is met
func (g *GetInstance) OffUntil(page goflow.Widget, predicate func(*Route) bool, transition ...Transition) {
	g.Until(predicate)
	g.To(page, transition...)
}

// Dialog shows a dialog overlay
func (g *GetInstance) Dialog(dialog goflow.Widget, barrierDismissible ...bool) {
	dismissible := true
	if len(barrierDismissible) > 0 {
		dismissible = barrierDismissible[0]
	}

	wrapper := NewDialogWrapper(dialog, dismissible, func() {
		g.CloseDialog()
	})

	g.dialogStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		return append(stack, wrapper)
	})
}

// CloseDialog closes the current dialog
func (g *GetInstance) CloseDialog() {
	g.dialogStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		if len(stack) == 0 {
			return stack
		}
		return stack[:len(stack)-1]
	})
}

// BottomSheet shows a bottom sheet overlay
func (g *GetInstance) BottomSheet(sheet goflow.Widget, isDismissible ...bool) {
	dismissible := true
	if len(isDismissible) > 0 {
		dismissible = isDismissible[0]
	}

	wrapper := NewBottomSheetWrapper(sheet, dismissible, func() {
		g.CloseBottomSheet()
	})

	g.overlayStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		return append(stack, wrapper)
	})
}

// CloseBottomSheet closes the current bottom sheet
func (g *GetInstance) CloseBottomSheet() {
	g.overlayStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		if len(stack) == 0 {
			return stack
		}
		return stack[:len(stack)-1]
	})
}

// Snackbar shows a snackbar (lightweight notification)
func (g *GetInstance) Snackbar(message string, duration ...int) {
	// TODO: Implement snackbar with auto-dismiss
	// For now, show as a temporary overlay
	durationMs := 3000
	if len(duration) > 0 {
		durationMs = duration[0]
	}

	snackbar := NewSnackbar(message, func() {
		g.CloseSnackbar()
	})

	g.overlayStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		return append(stack, snackbar)
	})

	// Auto-dismiss after duration
	go func() {
		// In a real implementation, use proper timer
		// For now, just a placeholder
		_ = durationMs
	}()
}

// CloseSnackbar closes the current snackbar
func (g *GetInstance) CloseSnackbar() {
	g.overlayStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		if len(stack) == 0 {
			return stack
		}
		return stack[:len(stack)-1]
	})
}

// RegisterRoute registers a named route
func (g *GetInstance) RegisterRoute(name string, builder RouteBuilder) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.namedRoutes[name] = builder
}

// RegisterRoutes registers multiple named routes
func (g *GetInstance) RegisterRoutes(routes map[string]RouteBuilder) {
	g.mu.Lock()
	defer g.mu.Unlock()
	for name, builder := range routes {
		g.namedRoutes[name] = builder
	}
}

// GetCurrentRoute returns the current route
func (g *GetInstance) GetCurrentRoute() *Route {
	stack := g.routeStack.Get()
	if len(stack) == 0 {
		return nil
	}
	return stack[len(stack)-1]
}

// CanPop returns whether there are routes to pop
func (g *GetInstance) CanPop() bool {
	return len(g.routeStack.Get()) > 1
}

// push adds a route to the stack
func (g *GetInstance) push(route *Route) {
	g.routeStack.Update(func(stack []*Route) []*Route {
		return append(stack, route)
	})
}

// pop removes the top route from the stack
func (g *GetInstance) pop() {
	g.routeStack.Update(func(stack []*Route) []*Route {
		if len(stack) <= 1 {
			// Don't pop the last route
			return stack
		}
		return stack[:len(stack)-1]
	})
}

// GetRouteStack returns the current route stack signal
func (g *GetInstance) GetRouteStack() *signals.Signal[[]*Route] {
	return g.routeStack
}

// GetDialogStack returns the dialog stack signal
func (g *GetInstance) GetDialogStack() *signals.Signal[[]goflow.Widget] {
	return g.dialogStack
}

// GetOverlayStack returns the overlay stack signal
func (g *GetInstance) GetOverlayStack() *signals.Signal[[]goflow.Widget] {
	return g.overlayStack
}
