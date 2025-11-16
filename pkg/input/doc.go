/*
Package input provides comprehensive input handling for GoFlow applications.

This package implements a complete input system supporting keyboard, mouse, touch,
and multitouch interactions with gesture recognition.

# Overview

The input package consists of several key components:

1. Event System - Core event types and handling
2. Keyboard Input - Key events, state tracking, and bindings
3. Mouse Input - Mouse events, button tracking, and regions
4. Touch Input - Touch events and pointer abstraction
5. Multitouch - Advanced gesture recognition
6. Event Dispatcher - Event routing and filtering

# Event System

The event system provides a unified interface for all input events:

	type InputEvent interface {
		Type() EventType
		Timestamp() time.Time
		Device() DeviceType
	}

Event types include:
  - Keyboard events (KeyDown, KeyUp, KeyRepeat)
  - Mouse events (MouseDown, MouseUp, MouseMove, MouseWheel)
  - Touch events (TouchStart, TouchMove, TouchEnd, TouchCancel)
  - Pointer events (unified mouse/touch/pen events)

# Keyboard Input

Keyboard input supports key events with modifier tracking:

	// Create keyboard state tracker
	keyState := input.NewKeyboardState()

	// Handle keyboard events
	event := input.NewKeyboardEvent(
		input.EventTypeKeyDown,
		input.KeyA,
		"KeyA",
		"a",
		input.ModifierCtrl,
		false,
	)
	keyState.HandleEvent(event)

	// Check key state
	if keyState.IsKeyPressed(input.KeyA) {
		// Key is pressed
	}

Keyboard shortcuts can be managed with KeyBindingManager:

	manager := input.NewKeyBindingManager()
	manager.AddBinding(input.KeyS, input.ModifierCtrl, func() {
		// Save action
	})

# Mouse Input

Mouse input provides button tracking, movement, and regions:

	// Create mouse state tracker
	mouseState := input.NewMouseState()

	// Create mouse region
	region := input.NewMouseRegion(bounds)
	region.OnClick = func(event *input.MouseEvent) {
		// Handle click
	}

	// Manage regions
	manager := input.NewMouseRegionManager()
	manager.AddRegion(region)

Mouse wheel events are also supported:

	wheelEvent := input.NewMouseWheelEvent(
		position,
		deltaX,
		deltaY,
		input.WheelDeltaModePixel,
	)

# Touch Input

Touch input supports single and multi-touch:

	// Create touch state tracker
	touchState := input.NewTouchState()

	// Add touch
	touch := touchState.AddTouch(position)

	// Update touch
	touchState.UpdateTouch(touch.ID, newPosition, force)

	// Remove touch
	touchState.RemoveTouch(touch.ID)

# Multitouch Gestures

The multitouch system recognizes complex gestures:

	// Create gesture recognizer
	recognizer := input.NewMultiTouchRecognizer(nil)

	// Set gesture callbacks
	recognizer.SetOnPinch(func(gesture *input.MultiTouchGesture) {
		scale := gesture.Scale
		// Handle pinch gesture
	})

	recognizer.SetOnRotate(func(gesture *input.MultiTouchGesture) {
		rotation := gesture.Rotation
		// Handle rotation gesture
	})

	recognizer.SetOnSwipe(func(gesture *input.MultiTouchGesture) {
		direction := gesture.Type
		// Handle swipe gesture
	})

Supported gestures:
  - Tap, DoubleTap, LongPress
  - Swipe (Left, Right, Up, Down)
  - Pan (drag)
  - Pinch (In/Out)
  - Rotate

# Event Dispatcher

The event dispatcher manages event routing:

	// Create dispatcher
	dispatcher := input.NewEventDispatcher()

	// Add event listener
	dispatcher.AddListener(
		input.EventTypeKeyDown,
		func(event input.InputEvent) input.EventPropagation {
			// Handle event
			return input.PropagationContinue
		},
		100, // priority
	)

	// Dispatch event
	dispatcher.Dispatch(event)

Event filters can be used to filter events:

	filter := input.NewEventTypeFilter(
		input.EventTypeKeyDown,
		input.EventTypeKeyUp,
	)
	dispatcher.AddFilter(filter)

# Input Router

The InputRouter provides a high-level interface combining all components:

	// Create router
	router := input.NewInputRouter()

	// Route events
	router.RouteEvent(keyboardEvent)
	router.RouteEvent(mouseEvent)
	router.RouteEvent(touchEvent)

	// Access state
	keyState := router.GetKeyboardState()
	mouseState := router.GetMouseState()
	touchState := router.GetTouchState()

# Configuration

Gesture recognition can be configured:

	config := &input.GestureConfig{
		TapTimeout:       300 * time.Millisecond,
		LongPressTime:    500 * time.Millisecond,
		SwipeMinDistance: 50.0,
		PinchMinScale:    0.1,
		RotationMinAngle: 5.0,
	}

	recognizer := input.NewMultiTouchRecognizer(config)

# Example: Complete Input Handling

	// Create input router
	router := input.NewInputRouter()

	// Setup keyboard bindings
	kb := router.GetKeyBindingManager()
	kb.AddBinding(input.KeyS, input.ModifierCtrl, func() {
		println("Save")
	})

	// Setup mouse region
	regionMgr := router.GetMouseRegionManager()
	region := input.NewMouseRegion(bounds)
	region.OnClick = func(e *input.MouseEvent) {
		println("Clicked!")
	}
	regionMgr.AddRegion(region)

	// Setup gesture recognition
	gr := router.GetGestureRecognizer()
	gr.SetOnPinch(func(g *input.MultiTouchGesture) {
		println("Pinch scale:", g.Scale)
	})

	// Route events
	router.RouteEvent(keyEvent)
	router.RouteEvent(mouseEvent)
	router.RouteEvent(touchEvent)

# Integration with GoFlow Widgets

The input system integrates with GoFlow's existing gesture system in
pkg/core/widgets/gesture.go, allowing widgets to receive input events
through a unified interface.

# Thread Safety

The EventDispatcher is thread-safe and can be used from multiple goroutines.
Other components (KeyboardState, MouseState, TouchState) are not thread-safe
and should be accessed from a single goroutine (typically the UI thread).

# Performance

The input system is designed for high performance:
  - Event dispatch is O(n) where n is the number of listeners
  - Mouse region hit testing is O(n) where n is the number of regions
  - Touch state tracking is O(1) for add/update/remove operations
  - Gesture recognition is O(n) where n is the number of active touches

# Best Practices

1. Use the InputRouter for comprehensive input handling
2. Configure gesture parameters based on your application needs
3. Use event priorities to control dispatch order
4. Clean up listeners when they're no longer needed
5. Use mouse regions for efficient hit testing
6. Process touch events on the UI thread

*/
package input
