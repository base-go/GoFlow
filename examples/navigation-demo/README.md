# Navigation Demo

This example demonstrates GoFlow's GetX-style navigation system.

## Features Demonstrated

### 1. Basic Navigation
- Navigate to new pages with `Get.To()`
- Navigate back with `Get.Back()`
- Check if can pop with `Get.CanPop()`

### 2. Named Routes
- Register routes with names
- Navigate using route names with `Get.ToNamed()`
- Clean route management

### 3. Advanced Navigation
- **Off**: Navigate and remove previous route (`Get.Off()`)
- **OffAll**: Navigate and clear entire stack (`Get.OffAll()`)
- **Until**: Navigate back until condition is met (`Get.Until()`)

### 4. Transitions
- Fade transition
- Slide transitions (Right, Left, Up, Down)
- Zoom transition
- Platform-specific transitions (Cupertino, Material)

### 5. Dialogs
- Show custom dialogs with `Get.Dialog()`
- Alert dialogs with `ShowAlertDialog()`
- Dismissible barriers
- Dialog actions

### 6. Bottom Sheets
- Show bottom sheets with `Get.BottomSheet()`
- Modal bottom sheets with `ShowModalBottomSheet()`
- Drag handles
- Dismissible sheets

### 7. Snackbars
- Show temporary notifications with `ShowSnackbar()`
- Custom duration
- Action buttons
- Auto-dismiss

## Running the Example

```bash
cd examples/navigation-demo
go run main.go
```

## Code Structure

### Home Page
```go
// Navigate to new page
navigation.Get.To(NewSecondPage(), navigation.TransitionSlideRight)

// Navigate to named route
navigation.Get.ToNamed("/second", navigation.TransitionFade)

// Show dialog
navigation.ShowAlertDialog(title, content, actions)

// Show bottom sheet
navigation.ShowModalBottomSheet(content, "Title")

// Show snackbar
navigation.ShowSnackbar("Message!", 3000)
```

### Second Page
```go
// Navigate forward
navigation.Get.To(NewThirdPage())

// Replace current route
navigation.Get.Off(NewThirdPage())

// Navigate to named route
navigation.Get.ToNamed("/settings")

// Navigate back
navigation.Get.Back()
```

### Third Page
```go
// Back to home, clearing stack
navigation.Get.OffAllNamed("/")

// Check if can pop
canPop := navigation.Get.CanPop()

// Navigate back until condition
navigation.Get.Until(func(route *Route) bool {
    return route.Name == "/home"
})
```

## Navigation API

### Basic Navigation
- `Get.To(page, transition?)` - Navigate to page
- `Get.ToNamed(routeName, transition?)` - Navigate to named route
- `Get.Back()` - Navigate back

### Stack Management
- `Get.Off(page)` - Navigate and remove previous
- `Get.OffNamed(routeName)` - Navigate to named route and remove previous
- `Get.OffAll(page)` - Navigate and clear stack
- `Get.OffAllNamed(routeName)` - Navigate to named route and clear stack
- `Get.Until(predicate)` - Navigate back until condition
- `Get.CanPop()` - Check if can navigate back

### Overlays
- `Get.Dialog(widget, dismissible?)` - Show dialog
- `Get.CloseDialog()` - Close current dialog
- `Get.BottomSheet(widget, dismissible?)` - Show bottom sheet
- `Get.CloseBottomSheet()` - Close current bottom sheet
- `Get.Snackbar(message, duration?)` - Show snackbar

### Helpers
- `ShowDialog(widget)` - Helper to show dialog
- `ShowAlertDialog(title, content, actions)` - Show alert dialog
- `ShowBottomSheet(widget)` - Helper to show bottom sheet
- `ShowModalBottomSheet(content, title)` - Show modal bottom sheet
- `ShowSnackbar(message, duration)` - Show snackbar
- `ShowSnackbarWithAction(message, label, onAction)` - Show snackbar with action

### Routes Registration
```go
routes := map[string]navigation.RouteBuilder{
    "/": func() goflow.Widget { return NewHomePage() },
    "/details": func() goflow.Widget { return NewDetailsPage() },
}

navigation.Get.RegisterRoutes(routes)
```

## Transitions

Available transition types:
- `TransitionFade` - Fade in/out
- `TransitionSlideRight` - Slide from right (iOS-style)
- `TransitionSlideLeft` - Slide from left
- `TransitionSlideUp` - Slide from bottom (Material-style)
- `TransitionSlideDown` - Slide from top
- `TransitionZoom` - Zoom in/out
- `TransitionNone` - No transition
- `TransitionCupertino` - Alias for SlideRight
- `TransitionMaterial` - Alias for SlideUp

## Best Practices

1. **Use Named Routes** for major app sections
2. **Use Get.To()** for simple forward navigation
3. **Use Get.OffAll()** when navigating to a completely new flow
4. **Always provide dismiss** for dialogs and bottom sheets
5. **Use appropriate transitions** for your platform
6. **Keep snackbar messages short** and actionable

## GetX Compatibility

This implementation is inspired by GetX and provides similar APIs:

| GetX | GoFlow |
|------|--------|
| `Get.to()` | `Get.To()` |
| `Get.back()` | `Get.Back()` |
| `Get.off()` | `Get.Off()` |
| `Get.offAll()` | `Get.OffAll()` |
| `Get.toNamed()` | `Get.ToNamed()` |
| `Get.dialog()` | `Get.Dialog()` |
| `Get.bottomSheet()` | `Get.BottomSheet()` |
| `Get.snackbar()` | `Get.Snackbar()` |

## Notes

- Animations are placeholders and will be implemented when the animation system is ready
- Gesture detection is commented out and will work when GestureDetector is available
- The navigation system uses GoFlow's signals for reactive updates
