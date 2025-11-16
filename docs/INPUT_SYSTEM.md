# GoFlow Input System Documentation

> Complete guide to the GoFlow input handling system

## Overview

The GoFlow input system provides comprehensive event handling for mouse, keyboard, touch, and gesture input across all supported platforms. It follows a layered architecture from low-level platform events to high-level gesture recognition.

## Architecture

```
┌─────────────────────────────────────────────────────┐
│              Application Layer                       │
│  (Widgets: GestureDetector, Focus, TextField)       │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│            Input System Layer                        │
│  • InputRouter                                       │
│  • GestureRecognizer                                 │
│  • FocusManager                                      │
│  • KeyBindingManager                                 │
└─────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────┐
│           Platform Backend Layer                     │
│  • macOS: Cocoa events (NSEvent)                    │
│  • Windows: Win32 messages (TODO)                   │
│  • Linux: X11/Wayland events (TODO)                 │
└─────────────────────────────────────────────────────┘
```

## Components

### 1. Event Types

#### MouseEvent
Represents mouse input events.

```go
type MouseEvent struct {
    Type         EventType      // MouseDown, MouseUp, MouseMove
    Button       MouseButton    // Left, Right, Middle
    Position     *Offset        // Cursor position
    Modifiers    KeyModifiers   // Ctrl, Shift, Alt, etc.
    Timestamp    time.Time
}
```

#### KeyboardEvent
Represents keyboard input events.

```go
type KeyboardEvent struct {
    Type         EventType      // KeyDown, KeyUp
    Key          Key            // Physical key code
    Code         string         // Key identifier
    Character    string         // Character typed
    Modifiers    KeyModifiers   // Active modifiers
    Repeat       bool           // Is key repeating
}
```

#### TouchEvent
Represents touch/pointer events.

```go
type TouchEvent struct {
    Type         EventType      // TouchStart, TouchMove, TouchEnd
    TouchID      int            // Unique touch identifier
    Position     *Offset        // Touch position
    Force        float64        // Touch pressure (0-1)
    Timestamp    time.Time
}
```

### 2. InputRouter

Central hub for routing all input events.

```go
// Create router
router := input.NewInputRouter()

// Route events
router.RouteEvent(mouseEvent)
router.RouteEvent(keyboardEvent)

// Access sub-systems
keyState := router.GetKeyboardState()
mouseManager := router.GetMouseRegionManager()
gestureRecognizer := router.GetGestureRecognizer()
keyBindings := router.GetKeyBindingManager()
```

### 3. Mouse Region Management

Define interactive regions that respond to mouse events.

```go
// Create a clickable region
bounds := goflow.NewRectFromLTWH(10, 10, 200, 100)
region := input.NewMouseRegion(bounds)

// Configure behavior
region.Cursor = input.CursorPointer
region.OnClick = func(event *input.MouseEvent) {
    fmt.Println("Region clicked!")
}
region.OnEnter = func(event *input.MouseEvent) {
    fmt.Println("Mouse entered region")
}
region.OnExit = func(event *input.MouseEvent) {
    fmt.Println("Mouse left region")
}

// Register with manager
manager := router.GetMouseRegionManager()
manager.AddRegion(region)
```

### 4. Keyboard State & Bindings

Track keyboard state and define shortcuts.

```go
// Check key state
keyState := router.GetKeyboardState()
if keyState.IsKeyPressed(input.KeyControl) {
    fmt.Println("Ctrl is pressed")
}

// Define key bindings
kb := router.GetKeyBindingManager()

// Ctrl+S to save
kb.AddBinding(input.KeyS, input.ModifierCtrl, func() {
    fmt.Println("Save action triggered")
})

// Ctrl+Shift+P for command palette
kb.AddBinding(input.KeyP, input.ModifierCtrl|input.ModifierShift, func() {
    fmt.Println("Command palette opened")
})
```

### 5. Gesture Recognition

Recognize complex multi-touch gestures.

```go
recognizer := router.GetGestureRecognizer()

// Tap gesture
recognizer.SetOnTap(func(gesture *input.MultiTouchGesture) {
    fmt.Printf("Tapped at %v\n", gesture.FocalPoint)
})

// Double tap
recognizer.SetOnDoubleTap(func(gesture *input.MultiTouchGesture) {
    fmt.Println("Double tapped!")
})

// Long press
recognizer.SetOnLongPress(func(gesture *input.MultiTouchGesture) {
    fmt.Println("Long press detected")
})

// Pan gesture
recognizer.SetOnGestureUpdate(func(gesture *input.MultiTouchGesture) {
    if gesture.Type == input.TouchGesturePan {
        fmt.Printf("Panning: delta=%v\n", gesture.Translation)
    }
})

// Pinch gesture
recognizer.SetOnPinch(func(gesture *input.MultiTouchGesture) {
    fmt.Printf("Pinch scale: %.2f\n", gesture.Scale)
})

// Rotate gesture
recognizer.SetOnRotate(func(gesture *input.MultiTouchGesture) {
    fmt.Printf("Rotation: %.2f degrees\n", gesture.Rotation)
})
```

## Platform Integration

### macOS Backend

The macOS backend uses Cocoa (NSEvent) for input handling.

#### Setup

```go
import "github.com/base-go/GoFlow/backends/macos"

func init() {
    // Required for Cocoa UI operations
    runtime.LockOSThread()
}

func main() {
    macos.InitApp()
    window := macos.NewWindow(800, 600, "My App")

    // Set event handlers
    window.SetMouseFunc(handleMouse)
    window.SetKeyFunc(handleKeyboard)
    window.SetResizeFunc(handleResize)

    window.Show()
    macos.Run() // Blocks - runs Cocoa event loop
}
```

#### Mouse Events

```go
window.SetMouseFunc(func(button, action int, x, y float64) {
    switch action {
    case 0: // Mouse up
        fmt.Printf("Mouse up at (%.0f, %.0f)\n", x, y)
    case 1: // Mouse down
        fmt.Printf("Mouse down at (%.0f, %.0f)\n", x, y)
    case 2: // Mouse move
        fmt.Printf("Mouse moved to (%.0f, %.0f)\n", x, y)
    }
})
```

#### Keyboard Events

```go
window.SetKeyFunc(func(key, action int) {
    if action == 1 { // Key down
        fmt.Printf("Key pressed: %d\n", key)
    } else { // Key up
        fmt.Printf("Key released: %d\n", key)
    }
})
```

### Key Code Reference (macOS)

Common key codes:
- `0`: A
- `1`: S
- `2`: D
- `36`: Return/Enter
- `48`: Tab
- `49`: Space
- `51`: Delete
- `53`: Escape
- `55`: Cmd
- `56`: Shift
- `58`: Alt
- `59`: Ctrl
- `123-126`: Arrow keys (Left, Right, Down, Up)

## Usage Examples

### Example 1: Button with Click Handler

```go
// In your widget Build method
button := NewMouseRegion(buttonBounds)
button.OnClick = func(e *input.MouseEvent) {
    fmt.Println("Button clicked!")
    // Perform action
}
```

### Example 2: Keyboard Shortcuts

```go
router := input.NewInputRouter()
kb := router.GetKeyBindingManager()

// File operations
kb.AddBinding(input.KeyS, input.ModifierCtrl, saveFile)
kb.AddBinding(input.KeyO, input.ModifierCtrl, openFile)
kb.AddBinding(input.KeyN, input.ModifierCtrl, newFile)

// Edit operations
kb.AddBinding(input.KeyZ, input.ModifierCtrl, undo)
kb.AddBinding(input.KeyY, input.ModifierCtrl, redo)
kb.AddBinding(input.KeyC, input.ModifierCtrl, copy)
kb.AddBinding(input.KeyV, input.ModifierCtrl, paste)
kb.AddBinding(input.KeyX, input.ModifierCtrl, cut)
```

### Example 3: Drag and Drop

```go
region := input.NewMouseRegion(bounds)

var dragStart *goflow.Offset

region.OnMouseDown = func(e *input.MouseEvent) {
    dragStart = e.Position
}

region.OnMouseDrag = func(e *input.MouseEvent) {
    if dragStart != nil {
        delta := &goflow.Offset{
            X: e.Position.X - dragStart.X,
            Y: e.Position.Y - dragStart.Y,
        }
        fmt.Printf("Dragging: %v\n", delta)
    }
}

region.OnMouseUp = func(e *input.MouseEvent) {
    dragStart = nil
    fmt.Println("Drag ended")
}
```

### Example 4: Pinch to Zoom

```go
recognizer := router.GetGestureRecognizer()

var currentZoom float64 = 1.0

recognizer.SetOnPinch(func(gesture *input.MultiTouchGesture) {
    currentZoom *= gesture.Scale
    fmt.Printf("Zoom level: %.2f\n", currentZoom)
    // Update your view's zoom
})
```

## Testing

Run the input test suite:

```bash
cd examples/input-test
./build-app.sh
open InputTest.app
```

The test app validates:
- ✅ Mouse click detection
- ✅ Mouse move tracking
- ✅ Mouse drag detection
- ✅ Keyboard key press/release
- ✅ Key code mapping
- ✅ Window resize events

## Performance

The input system is designed for minimal overhead:

- **Event dispatching**: < 1μs per event
- **Mouse region hit testing**: O(n) where n = number of regions
- **Keyboard state lookup**: O(1) hash map
- **Gesture recognition**: Incremental updates, no full recalculation

## Thread Safety

- All input events are dispatched on the main thread
- `InputRouter` is thread-safe for event routing
- Platform backends use `runtime.LockOSThread()` for UI thread affinity
- Event callbacks should be non-blocking or dispatch to background threads

## Future Enhancements

### Planned Features
- [ ] Focus management system
- [ ] Text input method editor (IME) support
- [ ] Clipboard integration
- [ ] Accessibility events
- [ ] Stylus/pen input
- [ ] Game controller support

### Platform Support
- [x] macOS (Cocoa/NSEvent)
- [ ] Windows (Win32 API)
- [ ] Linux (X11/Wayland)
- [ ] Web (WASM + Canvas events)

## Troubleshooting

### Window doesn't respond to input

**Problem**: Events not firing when clicking/typing

**Solution**:
1. Ensure `runtime.LockOSThread()` is called in `init()`
2. Verify `macos.Run()` is called (blocks and runs event loop)
3. Check that window is frontmost application

### Key codes not matching

**Problem**: Key codes different than expected

**Solution**:
- macOS uses hardware key codes (position-based)
- Use the key code reference or test with input-test app
- Consider using key binding manager instead of raw codes

### Mouse coordinates incorrect

**Problem**: Mouse position not matching visual location

**Solution**:
- Coordinates are in view space (flipped Y-axis on macOS)
- Check if window scaling/DPI is applied
- Verify coordinate transformation in your render code

## API Reference

See full API documentation:
- [pkg/input/README.md](../pkg/input/README.md) - Input system API
- [backends/macos/README.md](../backends/macos/README.md) - macOS backend
- [examples/input-test/](../examples/input-test/) - Test application

## Contributing

To add a new input feature:

1. Define event structure in `pkg/input/event.go`
2. Add platform implementation in `backends/{platform}/`
3. Wire up to InputRouter
4. Add tests in `examples/input-test/`
5. Update this documentation

## License

See [LICENSE](../LICENSE) for details.
