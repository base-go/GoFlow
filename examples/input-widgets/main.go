package main

import (
	"fmt"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
	"github.com/base-go/GoFlow/pkg/input"
)

func main() {
	fmt.Println("GoFlow Input + Widgets Integration Demo")
	fmt.Println("========================================")

	// Demo 1: GestureDetector with Input System
	demonstrateGestureIntegration()

	// Demo 2: Widget Input Mixin
	demonstrateWidgetInputMixin()

	// Demo 3: Interactive Button
	demonstrateInteractiveButton()
}

func demonstrateGestureIntegration() {
	fmt.Println("\n1. GestureDetector Integration")
	fmt.Println("------------------------------")

	// Create a GestureDetector
	child := widgets.NewText("Click Me")
	detector := widgets.NewGestureDetector(child)

	// Setup gesture callbacks
	detector.OnTap = func() {
		fmt.Println("✓ Widget tapped!")
	}

	detector.OnDoubleTap = func() {
		fmt.Println("✓ Widget double-tapped!")
	}

	detector.OnLongPress = func() {
		fmt.Println("✓ Widget long-pressed!")
	}

	detector.OnPanStart = func(details widgets.DragStartDetails) {
		fmt.Printf("✓ Pan started at (%.0f, %.0f)\n",
			details.GlobalPosition.X, details.GlobalPosition.Y)
	}

	detector.OnPanUpdate = func(details widgets.DragUpdateDetails) {
		fmt.Printf("✓ Pan update: delta=(%.0f, %.0f)\n",
			details.Delta.X, details.Delta.Y)
	}

	detector.OnScaleUpdate = func(details widgets.ScaleUpdateDetails) {
		fmt.Printf("✓ Scale update: scale=%.2f, rotation=%.2f\n",
			details.Scale, details.Rotation)
	}

	// Create adapter to connect input events to gestures
	adapter := input.NewGestureEventAdapter(detector)

	// Simulate touch events
	touch1 := input.NewTouch(0, goflow.NewOffset(100, 100))

	// Tap
	tapStart := input.NewTouchEvent(
		input.EventTypeTouchStart,
		[]*input.Touch{touch1},
		[]*input.Touch{touch1},
		[]*input.Touch{touch1},
	)
	adapter.HandleTouchEvent(tapStart)

	tapEnd := input.NewTouchEvent(
		input.EventTypeTouchEnd,
		[]*input.Touch{},
		[]*input.Touch{touch1},
		[]*input.Touch{},
	)
	adapter.HandleTouchEvent(tapEnd)

	// Simulate mouse click
	mouseDown := input.NewMouseEvent(
		input.EventTypeMouseDown,
		goflow.NewOffset(100, 100),
		input.MouseButtonLeft,
		input.MouseButtonsLeft,
		1,
	)
	adapter.HandleMouseEvent(mouseDown)

	mouseUp := input.NewMouseEvent(
		input.EventTypeMouseUp,
		goflow.NewOffset(100, 100),
		input.MouseButtonLeft,
		input.MouseButtonsNone,
		1,
	)
	adapter.HandleMouseEvent(mouseUp)
}

func demonstrateWidgetInputMixin() {
	fmt.Println("\n2. Widget Input Mixin")
	fmt.Println("---------------------")

	// Create a widget with input capabilities
	mixin := input.NewWidgetInputMixin()
	mixin.InputBounds = goflow.NewRectFromLTWH(0, 0, 200, 50)

	// Setup keyboard handlers
	mixin.OnKeyDown = func(event *input.KeyboardEvent) input.EventPropagation {
		fmt.Printf("✓ Key pressed: %d (char: %s)\n", event.Key, event.Character)
		return input.PropagationContinue
	}

	// Setup mouse handlers
	mixin.OnMouseEnter = func(event *input.MouseEvent) input.EventPropagation {
		fmt.Println("✓ Mouse entered widget")
		return input.PropagationContinue
	}

	mixin.OnMouseLeave = func(event *input.MouseEvent) input.EventPropagation {
		fmt.Println("✓ Mouse left widget")
		return input.PropagationContinue
	}

	mixin.OnMouseDown = func(event *input.MouseEvent) input.EventPropagation {
		fmt.Printf("✓ Mouse button %d pressed\n", event.Button)
		return input.PropagationStop
	}

	// Setup touch handlers
	mixin.OnTouchStart = func(event *input.TouchEvent) input.EventPropagation {
		fmt.Printf("✓ Touch started with %d touches\n", len(event.Touches))
		return input.PropagationContinue
	}

	// Focus handlers
	mixin.OnFocus = func() input.EventPropagation {
		fmt.Println("✓ Widget focused")
		return input.PropagationContinue
	}

	mixin.OnBlur = func() input.EventPropagation {
		fmt.Println("✓ Widget blurred")
		return input.PropagationContinue
	}

	// Simulate focus
	mixin.Focus()

	// Simulate keyboard event
	keyEvent := input.NewKeyboardEvent(
		input.EventTypeKeyDown,
		input.KeyA,
		"KeyA",
		"a",
		input.ModifierNone,
		false,
	)
	mixin.HandleInputEvent(keyEvent)

	// Simulate mouse events
	mouseEnter := input.NewMouseEvent(
		input.EventTypeMouseMove,
		goflow.NewOffset(100, 25),
		input.MouseButtonNone,
		input.MouseButtonsNone,
		0,
	)
	mixin.HandleInputEvent(mouseEnter)

	mouseDown := input.NewMouseEvent(
		input.EventTypeMouseDown,
		goflow.NewOffset(100, 25),
		input.MouseButtonLeft,
		input.MouseButtonsLeft,
		1,
	)
	mixin.HandleInputEvent(mouseDown)

	// Blur
	mixin.Blur()
}

func demonstrateInteractiveButton() {
	fmt.Println("\n3. Interactive Button Widget")
	fmt.Println("-----------------------------")

	// Create an interactive button using both systems
	button := &InteractiveButton{
		Label:  "Submit",
		Bounds: goflow.NewRectFromLTWH(10, 10, 100, 40),
	}

	// Setup input handling
	router := input.NewInputRouter()

	// Add mouse region for the button
	mouseRegion := input.NewMouseRegion(button.Bounds)
	mouseRegion.Cursor = input.CursorPointer
	mouseRegion.OnClick = func(event *input.MouseEvent) {
		button.HandleClick()
	}
	mouseRegion.OnEnter = func(event *input.MouseEvent) {
		button.IsHovered = true
		fmt.Println("✓ Button hovered")
	}
	mouseRegion.OnLeave = func(event *input.MouseEvent) {
		button.IsHovered = false
		fmt.Println("✓ Button unhovered")
	}

	router.GetMouseRegionManager().AddRegion(mouseRegion)

	// Add keyboard binding for Enter key
	router.GetKeyBindingManager().AddBinding(
		input.KeyEnter,
		input.ModifierNone,
		func() {
			button.HandleClick()
		},
	)

	// Setup gesture recognition for touch
	detector := widgets.NewGestureDetector(widgets.NewText(button.Label))
	detector.OnTap = func() {
		button.HandleClick()
	}

	adapter := input.NewGestureEventAdapter(detector)

	// Simulate mouse click
	mouseClick := input.NewMouseEvent(
		input.EventTypeMouseDown,
		goflow.NewOffset(50, 25),
		input.MouseButtonLeft,
		input.MouseButtonsLeft,
		1,
	)
	router.RouteEvent(mouseClick)

	mouseRelease := input.NewMouseEvent(
		input.EventTypeMouseUp,
		goflow.NewOffset(50, 25),
		input.MouseButtonLeft,
		input.MouseButtonsNone,
		1,
	)
	router.RouteEvent(mouseRelease)

	// Simulate keyboard Enter
	keyDown := input.NewKeyboardEvent(
		input.EventTypeKeyDown,
		input.KeyEnter,
		"Enter",
		"\n",
		input.ModifierNone,
		false,
	)
	router.RouteEvent(keyDown)

	// Simulate touch tap
	touch := input.NewTouch(0, goflow.NewOffset(50, 25))
	touchStart := input.NewTouchEvent(
		input.EventTypeTouchStart,
		[]*input.Touch{touch},
		[]*input.Touch{touch},
		[]*input.Touch{touch},
	)
	adapter.HandleTouchEvent(touchStart)

	touchEnd := input.NewTouchEvent(
		input.EventTypeTouchEnd,
		[]*input.Touch{},
		[]*input.Touch{touch},
		[]*input.Touch{},
	)
	adapter.HandleTouchEvent(touchEnd)

	fmt.Printf("\nButton click count: %d\n", button.ClickCount)
}

// InteractiveButton is a custom button widget with input handling
type InteractiveButton struct {
	Label      string
	Bounds     *goflow.Rect
	OnClick    func()
	IsHovered  bool
	IsPressed  bool
	ClickCount int
}

func (b *InteractiveButton) HandleClick() {
	b.ClickCount++
	fmt.Printf("✓ Button '%s' clicked! (count: %d)\n", b.Label, b.ClickCount)

	if b.OnClick != nil {
		b.OnClick()
	}
}
