# GoFlow Navigation

GetX-style navigation and routing for GoFlow.

## Quick Start

```go
import "github.com/base-go/GoFlow/pkg/navigation"

// Create app with navigation
app := navigation.GetMaterialApp(
    NewHomePage(),
    routes,
    "My App",
)

// Navigate to a page
navigation.Get.To(NewSecondPage())

// Navigate back
navigation.Get.Back()
```

## API Reference

### Basic Navigation

```go
// Navigate to a new page
Get.To(page goflow.Widget, transition ...Transition)

// Navigate to a named route
Get.ToNamed(routeName string, transition ...Transition)

// Navigate back
Get.Back()

// Check if can navigate back
Get.CanPop() bool
```

### Stack Management

```go
// Navigate and remove previous route
Get.Off(page goflow.Widget, transition ...Transition)
Get.OffNamed(routeName string, transition ...Transition)

// Navigate and clear entire stack
Get.OffAll(page goflow.Widget, transition ...Transition)
Get.OffAllNamed(routeName string, transition ...Transition)

// Navigate back until condition is met
Get.Until(predicate func(*Route) bool)

// Navigate to page and remove routes until condition
Get.OffUntil(page goflow.Widget, predicate func(*Route) bool, transition ...Transition)
```

### Dialogs

```go
// Show a dialog
Get.Dialog(dialog goflow.Widget, barrierDismissible ...bool)

// Close current dialog
Get.CloseDialog()

// Helper: Show alert dialog
ShowAlertDialog(title, content string, actions []DialogAction)

// Example
ShowAlertDialog(
    "Confirm",
    "Are you sure?",
    []DialogAction{
        {Label: "Cancel", OnPress: func() { Get.CloseDialog() }},
        {Label: "OK", OnPress: func() { /* do something */ }, Primary: true},
    },
)
```

### Bottom Sheets

```go
// Show a bottom sheet
Get.BottomSheet(sheet goflow.Widget, isDismissible ...bool)

// Close current bottom sheet
Get.CloseBottomSheet()

// Helper: Show modal bottom sheet
ShowModalBottomSheet(content goflow.Widget, title ...string)

// Example
ShowModalBottomSheet(
    widgets.NewText("Options here"),
    "Choose an option",
)
```

### Snackbars

```go
// Show a snackbar (auto-dismiss after duration)
Get.Snackbar(message string, duration ...int)

// Helper: Show simple snackbar
ShowSnackbar(message string, duration ...int)

// Helper: Show snackbar with action
ShowSnackbarWithAction(message, actionLabel string, onAction func(), duration ...int)

// Examples
ShowSnackbar("Saved!", 3000)
ShowSnackbarWithAction("Deleted", "Undo", func() { /* undo */ })
```

### Named Routes

```go
// Register a single route
Get.RegisterRoute(name string, builder RouteBuilder)

// Register multiple routes
Get.RegisterRoutes(routes map[string]RouteBuilder)

// Example
routes := map[string]navigation.RouteBuilder{
    "/home": func() goflow.Widget { return NewHomePage() },
    "/profile": func() goflow.Widget { return NewProfilePage() },
    "/settings": func() goflow.Widget { return NewSettingsPage() },
}
Get.RegisterRoutes(routes)

// Navigate to named route
Get.ToNamed("/profile")
```

### Transitions

Available transition types:

```go
TransitionFade           // Fade in/out
TransitionSlideRight     // Slide from right (iOS-style)
TransitionSlideLeft      // Slide from left
TransitionSlideUp        // Slide from bottom (Material-style)
TransitionSlideDown      // Slide from top
TransitionZoom           // Zoom in/out
TransitionNone           // No transition
TransitionCupertino      // Alias for SlideRight
TransitionMaterial       // Alias for SlideUp
```

Usage:

```go
Get.To(NewPage(), navigation.TransitionSlideRight)
Get.ToNamed("/settings", navigation.TransitionFade)
```

### App Setup

```go
// Material Design app
app := navigation.GetMaterialApp(
    homePage,
    routes,
    "App Title",
)

// Cupertino (iOS) app
app := navigation.GetCupertinoApp(
    homePage,
    routes,
    "App Title",
)

// Basic app with Navigator
app := navigation.GetApp(homePage, routes)
```

### Navigation Observers

```go
type MyObserver struct{}

func (o *MyObserver) DidPush(route, previous *Route) {
    fmt.Println("Pushed:", route.Name)
}

func (o *MyObserver) DidPop(route, previous *Route) {
    fmt.Println("Popped:", route.Name)
}

func (o *MyObserver) DidReplace(new, old *Route) {
    fmt.Println("Replaced")
}

// Add observer to navigator
navigator := navigation.NewNavigator(home).
    WithObservers(&MyObserver{})
```

### Advanced: Reactive Navigation

```go
// Get the route stack signal
routeStack := Get.GetRouteStack()

// React to route changes
signals.NewEffect(func() {
    routes := routeStack.Get()
    fmt.Printf("Stack size: %d\n", len(routes))
})

// Get current route
currentRoute := Get.GetCurrentRoute()
if currentRoute != nil {
    fmt.Println("Current route:", currentRoute.Name)
}
```

## Examples

See [examples/navigation-demo](../../examples/navigation-demo) for a comprehensive example.

## GetX Compatibility

| GetX | GoFlow |
|------|--------|
| `Get.to()` | `Get.To()` |
| `Get.back()` | `Get.Back()` |
| `Get.off()` | `Get.Off()` |
| `Get.offAll()` | `Get.OffAll()` |
| `Get.toNamed()` | `Get.ToNamed()` |
| `Get.offNamed()` | `Get.OffNamed()` |
| `Get.offAllNamed()` | `Get.OffAllNamed()` |
| `Get.until()` | `Get.Until()` |
| `Get.dialog()` | `Get.Dialog()` |
| `Get.bottomSheet()` | `Get.BottomSheet()` |
| `Get.snackbar()` | `Get.Snackbar()` |

## Features

- ✅ GetX-style global navigation
- ✅ Named routes registry
- ✅ Route stack management
- ✅ Multiple transition types
- ✅ Dialogs with barriers
- ✅ Bottom sheets
- ✅ Snackbars
- ✅ Navigation observers
- ✅ Reactive route stack (using signals)
- ✅ Alert dialogs
- ✅ Modal bottom sheets
- ⏳ Route transitions (placeholders, will work when animation system is ready)
- ⏳ Gesture-based dismissal (will work when GestureDetector is ready)

## Architecture

The navigation system is built on GoFlow's signals for reactivity:

- **Route Stack**: `Signal[[]*Route]` - Reactive list of routes
- **Dialog Stack**: `Signal[[]Widget]` - Reactive list of dialogs
- **Overlay Stack**: `Signal[[]Widget]` - Reactive list of overlays

All navigation operations update these signals, causing the Navigator widget to rebuild automatically.

## Best Practices

1. **Use named routes** for major sections of your app
2. **Use Get.To()** for simple forward navigation
3. **Use Get.Off()** when you want to replace the current route
4. **Use Get.OffAll()** when navigating to a completely new flow (e.g., login → home)
5. **Always provide a dismiss action** for dialogs and bottom sheets
6. **Use appropriate transitions** for your platform:
   - iOS/macOS: `TransitionCupertino` (slide from right)
   - Android/Web: `TransitionMaterial` (slide from bottom)
7. **Keep snackbar messages short** and actionable

## Notes

- Transition animations are placeholders and will be implemented when the animation system is ready
- The navigation system uses GoFlow's signals for reactive state management
- All APIs are thread-safe and can be called from any goroutine
- The global `Get` instance is initialized automatically
