# GoFlow Input System

Comprehensive input handling for GoFlow applications supporting keyboard, mouse, touch, and multitouch with advanced gesture recognition.

## Features

- **Keyboard Input**: Full keyboard event handling with modifier tracking and key bindings
- **Mouse Input**: Mouse events, button state, cursor management, and region-based hit testing
- **Touch Input**: Single and multi-touch support with individual touch tracking
- **Multitouch Gestures**: Advanced gesture recognition (tap, pinch, rotate, swipe, pan, etc.)
- **Event Dispatcher**: Flexible event routing with priorities and filters
- **Widget Integration**: Seamless integration with GoFlow's widget system
- **Thread-Safe**: Dispatcher supports concurrent access from multiple goroutines

## Quick Start

### Basic Setup

```go
import "github.com/base-go/GoFlow/pkg/input"

// Create input router (combines all input systems)
router := input.NewInputRouter()

// Route events
router.RouteEvent(keyboardEvent)
router.RouteEvent(mouseEvent)
router.RouteEvent(touchEvent)
```

### Keyboard Input

```go
// Track keyboard state
keyState := router.GetKeyboardState()
if keyState.IsKeyPressed(input.KeyA) {
    // Key A is pressed
}

// Add keyboard shortcuts
kb := router.GetKeyBindingManager()
kb.AddBinding(input.KeyS, input.ModifierCtrl, func() {
    // Save action (Ctrl+S)
})
```

### Mouse Input

```go
// Create mouse region
region := input.NewMouseRegion(bounds)
region.Cursor = input.CursorPointer
region.OnClick = func(event *input.MouseEvent) {
    fmt.Println("Clicked at", event.Position)
}

// Add to manager
manager := router.GetMouseRegionManager()
manager.AddRegion(region)
```

### Touch Gestures

```go
// Setup gesture recognition
recognizer := router.GetGestureRecognizer()

recognizer.SetOnTap(func(gesture *input.MultiTouchGesture) {
    fmt.Println("Tapped at", gesture.FocalPoint)
})

recognizer.SetOnPinch(func(gesture *input.MultiTouchGesture) {
    fmt.Println("Pinch scale:", gesture.Scale)
})

recognizer.SetOnSwipe(func(gesture *input.MultiTouchGesture) {
    fmt.Println("Swiped:", gesture.Type)
})
```

## Core Components

### InputRouter

High-level interface combining all input systems:

```go
router := input.NewInputRouter()

// Access components
dispatcher := router.GetDispatcher()
keyState := router.GetKeyboardState()
mouseState := router.GetMouseState()
touchState := router.GetTouchState()
```

### EventDispatcher

Manages event routing with priorities and filters:

```go
dispatcher := input.NewEventDispatcher()

// Add listener
dispatcher.AddListener(
    input.EventTypeKeyDown,
    func(event input.InputEvent) input.EventPropagation {
        // Handle event
        return input.PropagationContinue
    },
    100, // priority (higher = first)
)

// Add filter
filter := input.NewEventTypeFilter(
    input.EventTypeKeyDown,
    input.EventTypeMouseDown,
)
dispatcher.AddFilter(filter)
```

### KeyboardState

Tracks keyboard state and modifiers:

```go
state := input.NewKeyboardState()

// Check key state
if state.IsKeyPressed(input.KeyShift) {
    // Shift is pressed
}

// Check modifiers
mods := state.GetModifiers()
if mods.HasCtrl() {
    // Control is pressed
}
```

### MouseState

Tracks mouse position and buttons:

```go
state := input.NewMouseState()

// Check button state
if state.IsButtonPressed(input.MouseButtonLeft) {
    // Left button is pressed
}

// Get position
position := state.Position
delta := state.GetDelta()
```

### TouchState

Tracks individual touches:

```go
state := input.NewTouchState()

// Add touch
touch := state.AddTouch(position)

// Update touch
state.UpdateTouch(touch.ID, newPosition, force)

// Get all touches
touches := state.GetAllTouches()
```

### MultiTouchRecognizer

Recognizes complex gestures:

```go
config := &input.GestureConfig{
    TapTimeout:       300 * time.Millisecond,
    LongPressTime:    500 * time.Millisecond,
    SwipeMinDistance: 50.0,
    PinchMinScale:    0.1,
    RotationMinAngle: 5.0,
}

recognizer := input.NewMultiTouchRecognizer(config)
```

## Event Types

### Keyboard Events

- `EventTypeKeyDown`: Key pressed
- `EventTypeKeyUp`: Key released
- `EventTypeKeyRepeat`: Key auto-repeat

### Mouse Events

- `EventTypeMouseDown`: Mouse button pressed
- `EventTypeMouseUp`: Mouse button released
- `EventTypeMouseMove`: Mouse moved
- `EventTypeMouseEnter`: Mouse entered region
- `EventTypeMouseLeave`: Mouse left region
- `EventTypeMouseWheel`: Mouse wheel scrolled

### Touch Events

- `EventTypeTouchStart`: Touch started
- `EventTypeTouchMove`: Touch moved
- `EventTypeTouchEnd`: Touch ended
- `EventTypeTouchCancel`: Touch canceled

### Pointer Events

- `EventTypePointerDown`: Pointer down (unified)
- `EventTypePointerUp`: Pointer up
- `EventTypePointerMove`: Pointer moved
- `EventTypePointerEnter`: Pointer entered
- `EventTypePointerLeave`: Pointer left
- `EventTypePointerCancel`: Pointer canceled

## Recognized Gestures

- `TouchGestureTap`: Single tap
- `TouchGestureDoubleTap`: Double tap
- `TouchGestureLongPress`: Long press
- `TouchGestureSwipeLeft`: Swipe left
- `TouchGestureSwipeRight`: Swipe right
- `TouchGestureSwipeUp`: Swipe up
- `TouchGestureSwipeDown`: Swipe down
- `TouchGesturePan`: Drag/pan
- `TouchGesturePinchIn`: Pinch to zoom out
- `TouchGesturePinchOut`: Pinch to zoom in
- `TouchGestureRotate`: Two-finger rotation

## Widget Integration

### With GestureDetector

```go
child := widgets.NewText("Click Me", nil)
detector := widgets.NewGestureDetector(child)

detector.OnTap = func() {
    fmt.Println("Tapped!")
}

// Create adapter
adapter := input.NewGestureEventAdapter(detector)
adapter.HandleTouchEvent(touchEvent)
adapter.HandleMouseEvent(mouseEvent)
```

### With WidgetInputMixin

```go
mixin := input.NewWidgetInputMixin()
mixin.InputBounds = bounds

mixin.OnKeyDown = func(event *input.KeyboardEvent) input.EventPropagation {
    // Handle key
    return input.PropagationContinue
}

mixin.OnMouseDown = func(event *input.MouseEvent) input.EventPropagation {
    // Handle click
    return input.PropagationStop
}

// Handle events
mixin.HandleInputEvent(event)
```

## Configuration

### Gesture Configuration

```go
config := &input.GestureConfig{
    // Tap settings
    TapTimeout:       300 * time.Millisecond,
    TapMaxMovement:   10.0,
    DoubleTapTimeout: 300 * time.Millisecond,

    // Long press settings
    LongPressTime:     500 * time.Millisecond,
    LongPressDistance: 10.0,

    // Swipe settings
    SwipeMinDistance: 50.0,
    SwipeMaxTime:     300 * time.Millisecond,

    // Pan settings
    PanMinDistance: 5.0,

    // Pinch settings
    PinchMinScale: 0.1,

    // Rotation settings
    RotationMinAngle: 5.0,

    // General
    MaxTouchPoints: 10,
}
```

## Advanced Usage

### Event Propagation

Control how events propagate through handlers:

```go
handler := func(event input.InputEvent) input.EventPropagation {
    // Process event

    // Continue to next handler
    return input.PropagationContinue

    // Stop propagation to remaining handlers
    // return input.PropagationStop

    // Stop all propagation immediately
    // return input.PropagationStopImmediate
}
```

### Event Filtering

Filter events before dispatch:

```go
// Type filter
typeFilter := input.NewEventTypeFilter(
    input.EventTypeKeyDown,
    input.EventTypeMouseDown,
)
dispatcher.AddFilter(typeFilter)

// Custom filter
customFilter := input.NewFunctionFilter(func(event input.InputEvent) bool {
    // Return true to allow, false to filter
    return event.Device() != input.DeviceTypeUnknown
})
dispatcher.AddFilter(customFilter)
```

### Event Queuing

Queue events for batch processing:

```go
dispatcher.EnableQueueMode()

// Events are queued but not dispatched
dispatcher.Dispatch(event1)
dispatcher.Dispatch(event2)

// Process all queued events
dispatcher.ProcessQueue()

dispatcher.DisableQueueMode()
```

### Mouse Regions

Create interactive regions:

```go
region := input.NewMouseRegion(bounds)
region.Cursor = input.CursorPointer
region.IsEnabled = true

region.OnEnter = func(event *input.MouseEvent) {
    // Mouse entered
}

region.OnLeave = func(event *input.MouseEvent) {
    // Mouse left
}

region.OnClick = func(event *input.MouseEvent) {
    // Clicked
}

region.OnWheel = func(event *input.MouseWheelEvent) {
    // Scrolled
}
```

## Examples

See the examples directory for complete demonstrations:

- `examples/input-demo/` - Basic input system usage
- `examples/input-widgets/` - Widget integration

## Performance

- Event dispatch: O(n) where n = number of listeners
- Mouse region hit testing: O(n) where n = number of regions
- Touch state tracking: O(1) for add/update/remove
- Gesture recognition: O(n) where n = number of active touches

## Thread Safety

- **EventDispatcher**: Thread-safe, supports concurrent access
- **KeyboardState**: Not thread-safe, use from UI thread
- **MouseState**: Not thread-safe, use from UI thread
- **TouchState**: Not thread-safe, use from UI thread
- **Recognizers**: Not thread-safe, use from UI thread

## Best Practices

1. Use `InputRouter` for comprehensive input handling
2. Configure gesture parameters for your application
3. Use event priorities to control dispatch order
4. Clean up listeners when no longer needed
5. Use mouse regions for efficient hit testing
6. Process events on the UI thread
7. Use appropriate propagation modes
8. Filter events for better performance

## License

Part of the GoFlow framework.
