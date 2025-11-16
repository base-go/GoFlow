// Package navigation provides GetX-style navigation and routing for GoFlow
//
// # Navigation Overview
//
// The navigation package provides a simple, reactive navigation system inspired by GetX.
// It includes support for:
//   - Page navigation with transitions
//   - Named routes
//   - Dialogs and bottom sheets
//   - Snackbars
//   - Navigation observers
//
// # Basic Usage
//
// Simple navigation:
//
//	import "github.com/base-go/GoFlow/pkg/navigation"
//
//	// Navigate to a new page
//	navigation.Get.To(NewSecondPage())
//
//	// Navigate back
//	navigation.Get.Back()
//
//	// Navigate with custom transition
//	navigation.Get.To(NewPage(), navigation.TransitionSlideRight)
//
// # Named Routes
//
// Register and use named routes:
//
//	// Register routes
//	navigation.Get.RegisterRoutes(map[string]navigation.RouteBuilder{
//	    "/home": func() goflow.Widget { return NewHomePage() },
//	    "/profile": func() goflow.Widget { return NewProfilePage() },
//	    "/settings": func() goflow.Widget { return NewSettingsPage() },
//	})
//
//	// Navigate to named route
//	navigation.Get.ToNamed("/profile")
//
// # Advanced Navigation
//
// Remove routes from stack:
//
//	// Navigate and remove previous route
//	navigation.Get.Off(NewPage())
//
//	// Navigate and remove all previous routes
//	navigation.Get.OffAll(NewHomePage())
//
//	// Navigate back until condition is met
//	navigation.Get.Until(func(route *navigation.Route) bool {
//	    return route.Name == "/home"
//	})
//
// # Dialogs
//
// Show dialogs:
//
//	// Simple dialog
//	dialog := widgets.NewContainer()
//	dialog.Child = widgets.NewText("Hello Dialog!")
//	navigation.Get.Dialog(dialog)
//
//	// Alert dialog with actions
//	navigation.ShowAlertDialog(
//	    "Confirm",
//	    "Are you sure?",
//	    []navigation.DialogAction{
//	        {Label: "Cancel", OnPress: func() { navigation.Get.CloseDialog() }},
//	        {Label: "OK", OnPress: func() { /* do something */ }, Primary: true},
//	    },
//	)
//
//	// Close dialog
//	navigation.Get.CloseDialog()
//
// # Bottom Sheets
//
// Show bottom sheets:
//
//	// Simple bottom sheet
//	content := widgets.NewText("Bottom Sheet Content")
//	navigation.Get.BottomSheet(content)
//
//	// Modal bottom sheet with title
//	navigation.ShowModalBottomSheet(
//	    widgets.NewText("Sheet content"),
//	    "Sheet Title",
//	)
//
//	// Close bottom sheet
//	navigation.Get.CloseBottomSheet()
//
// # Snackbars
//
// Show snackbars (temporary notifications):
//
//	// Simple snackbar (auto-dismisses after 3 seconds)
//	navigation.ShowSnackbar("Action completed!")
//
//	// Snackbar with custom duration
//	navigation.ShowSnackbar("Message", 5000) // 5 seconds
//
//	// Snackbar with action
//	navigation.ShowSnackbarWithAction(
//	    "Item deleted",
//	    "Undo",
//	    func() { /* undo action */ },
//	)
//
// # GetMaterialApp / GetCupertinoApp
//
// Create an app with navigation:
//
//	func main() {
//	    routes := map[string]navigation.RouteBuilder{
//	        "/": func() goflow.Widget { return NewHomePage() },
//	        "/details": func() goflow.Widget { return NewDetailsPage() },
//	    }
//
//	    app := navigation.GetMaterialApp(
//	        NewHomePage(),
//	        routes,
//	        "My App",
//	    )
//
//	    goflow.RunApp(app)
//	}
//
// # Transitions
//
// Available transition types:
//   - TransitionFade: Fade in/out
//   - TransitionSlideRight: Slide from right (iOS-style)
//   - TransitionSlideLeft: Slide from left
//   - TransitionSlideUp: Slide from bottom (Material-style)
//   - TransitionSlideDown: Slide from top
//   - TransitionZoom: Zoom in/out
//   - TransitionNone: No transition
//   - TransitionCupertino: Alias for SlideRight
//   - TransitionMaterial: Alias for SlideUp
//
// Usage:
//
//	navigation.Get.To(NewPage(), navigation.TransitionSlideRight)
//	navigation.Get.ToNamed("/settings", navigation.TransitionFade)
//
// # Navigation Observers
//
// Observe navigation events:
//
//	type MyObserver struct{}
//
//	func (o *MyObserver) DidPush(route, previous *navigation.Route) {
//	    fmt.Println("Pushed route:", route.Name)
//	}
//
//	func (o *MyObserver) DidPop(route, previous *navigation.Route) {
//	    fmt.Println("Popped route:", route.Name)
//	}
//
//	func (o *MyObserver) DidReplace(new, old *navigation.Route) {
//	    fmt.Println("Replaced route")
//	}
//
//	// Add observer to navigator
//	navigator := navigation.NewNavigator(homePage).
//	    WithObservers(&MyObserver{})
//
// # Reactive Navigation
//
// The navigation system uses signals for reactive updates:
//
//	// Get the current route stack
//	routeStack := navigation.Get.GetRouteStack()
//
//	// Create an effect that runs when routes change
//	signals.NewEffect(func() {
//	    routes := routeStack.Get()
//	    fmt.Printf("Current stack size: %d\n", len(routes))
//	})
//
// # Best Practices
//
// 1. Use named routes for major app sections
// 2. Use Get.To() for simple forward navigation
// 3. Use Get.OffAll() when navigating to a completely new flow
// 4. Always provide a way to dismiss dialogs and bottom sheets
// 5. Use appropriate transitions for your platform (Cupertino for iOS, Material for Android)
// 6. Keep snackbar messages short and actionable
//
// # GetX Compatibility
//
// This package is inspired by GetX and provides similar APIs:
//   - Get.to() → Get.To()
//   - Get.back() → Get.Back()
//   - Get.off() → Get.Off()
//   - Get.offAll() → Get.OffAll()
//   - Get.toNamed() → Get.ToNamed()
//   - Get.dialog() → Get.Dialog()
//   - Get.bottomSheet() → Get.BottomSheet()
//   - Get.snackbar() → Get.Snackbar()
//
// The main difference is Go's naming conventions (PascalCase for exported functions).
package navigation
