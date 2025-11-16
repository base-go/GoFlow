package main

import (
	"fmt"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/input"
)

func main() {
	fmt.Println("GoFlow Input System Demo")
	fmt.Println("========================")

	// Create input router
	router := input.NewInputRouter()

	// Demo 1: Keyboard Input
	demonstrateKeyboard(router)

	// Demo 2: Mouse Input
	demonstrateMouse(router)

	// Demo 3: Touch Input
	demonstrateTouch(router)

	// Demo 4: Multitouch Gestures
	demonstrateMultitouch(router)

	// Demo 5: Event Dispatcher
	demonstrateEventDispatcher(router)
}

func demonstrateKeyboard(router *input.InputRouter) {
	fmt.Println("\n1. Keyboard Input Demo")
	fmt.Println("----------------------")

	// Setup keyboard bindings
	kb := router.GetKeyBindingManager()

	// Ctrl+S to save
	kb.AddBinding(input.KeyS, input.ModifierCtrl, func() {
		fmt.Println("✓ Save triggered (Ctrl+S)")
	})

	// Ctrl+C to copy
	kb.AddBinding(input.KeyC, input.ModifierCtrl, func() {
		fmt.Println("✓ Copy triggered (Ctrl+C)")
	})

	// Ctrl+V to paste
	kb.AddBinding(input.KeyV, input.ModifierCtrl, func() {
		fmt.Println("✓ Paste triggered (Ctrl+V)")
	})

	// Simulate keyboard events
	events := []*input.KeyboardEvent{
		input.NewKeyboardEvent(input.EventTypeKeyDown, input.KeyControl, "ControlLeft", "", input.ModifierCtrl, false),
		input.NewKeyboardEvent(input.EventTypeKeyDown, input.KeyS, "KeyS", "s", input.ModifierCtrl, false),
		input.NewKeyboardEvent(input.EventTypeKeyUp, input.KeyS, "KeyS", "s", input.ModifierCtrl, false),
		input.NewKeyboardEvent(input.EventTypeKeyUp, input.KeyControl, "ControlLeft", "", input.ModifierNone, false),
	}

	for _, event := range events {
		router.RouteEvent(event)
	}

	// Check keyboard state
	keyState := router.GetKeyboardState()
	fmt.Printf("Ctrl pressed: %v\n", keyState.IsKeyPressed(input.KeyControl))
}

func demonstrateMouse(router *input.InputRouter) {
	fmt.Println("\n2. Mouse Input Demo")
	fmt.Println("-------------------")

	// Create mouse region
	bounds := goflow.NewRectFromLTWH(10, 10, 200, 100)
	region := input.NewMouseRegion(bounds)
	region.Cursor = input.CursorPointer

	clickCount := 0
	region.OnClick = func(event *input.MouseEvent) {
		clickCount++
		fmt.Printf("✓ Region clicked! (count: %d) at position (%.0f, %.0f)\n",
			clickCount, event.Position.X, event.Position.Y)
	}

	region.OnEnter = func(event *input.MouseEvent) {
		fmt.Println("✓ Mouse entered region")
	}

	region.OnLeave = func(event *input.MouseEvent) {
		fmt.Println("✓ Mouse left region")
	}

	// Add region to manager
	regionMgr := router.GetMouseRegionManager()
	regionMgr.AddRegion(region)

	// Simulate mouse events
	events := []*input.MouseEvent{
		input.NewMouseEvent(input.EventTypeMouseMove, goflow.NewOffset(50, 50), input.MouseButtonNone, input.MouseButtonsNone, 0),
		input.NewMouseEvent(input.EventTypeMouseDown, goflow.NewOffset(50, 50), input.MouseButtonLeft, input.MouseButtonsLeft, 1),
		input.NewMouseEvent(input.EventTypeMouseUp, goflow.NewOffset(50, 50), input.MouseButtonLeft, input.MouseButtonsNone, 1),
		input.NewMouseEvent(input.EventTypeMouseMove, goflow.NewOffset(250, 50), input.MouseButtonNone, input.MouseButtonsNone, 0),
	}

	for _, event := range events {
		router.RouteEvent(event)
	}

	// Mouse wheel
	wheelEvent := input.NewMouseWheelEvent(goflow.NewOffset(50, 50), 0, -120, input.WheelDeltaModePixel)
	fmt.Println("✓ Mouse wheel scrolled:", wheelEvent.DeltaY)
}

func demonstrateTouch(router *input.InputRouter) {
	fmt.Println("\n3. Touch Input Demo")
	fmt.Println("-------------------")

	// Create touch state
	touchState := router.GetTouchState()

	// Simulate touch sequence
	touch1 := touchState.AddTouch(goflow.NewOffset(100, 100))
	fmt.Printf("✓ Touch started: ID=%d, Position=(%.0f, %.0f)\n",
		touch1.ID, touch1.Position.X, touch1.Position.Y)

	// Update touch
	time.Sleep(10 * time.Millisecond)
	touchState.UpdateTouch(touch1.ID, goflow.NewOffset(150, 150), 1.0)
	touch1 = touchState.GetTouch(touch1.ID)
	fmt.Printf("✓ Touch moved: ID=%d, Position=(%.0f, %.0f), Delta=(%.0f, %.0f)\n",
		touch1.ID, touch1.Position.X, touch1.Position.Y,
		touch1.GetDelta().X, touch1.GetDelta().Y)

	// End touch
	touchState.RemoveTouch(touch1.ID)
	fmt.Printf("✓ Touch ended: count=%d\n", touchState.GetTouchCount())
}

func demonstrateMultitouch(router *input.InputRouter) {
	fmt.Println("\n4. Multitouch Gestures Demo")
	fmt.Println("---------------------------")

	// Create gesture recognizer
	recognizer := router.GetGestureRecognizer()

	// Setup gesture callbacks
	recognizer.SetOnTap(func(gesture *input.MultiTouchGesture) {
		fmt.Printf("✓ Tap gesture at (%.0f, %.0f)\n",
			gesture.FocalPoint.X, gesture.FocalPoint.Y)
	})

	recognizer.SetOnPinch(func(gesture *input.MultiTouchGesture) {
		fmt.Printf("✓ Pinch gesture: scale=%.2f\n", gesture.Scale)
	})

	recognizer.SetOnRotate(func(gesture *input.MultiTouchGesture) {
		fmt.Printf("✓ Rotate gesture: rotation=%.2f radians\n", gesture.Rotation)
	})

	recognizer.SetOnSwipe(func(gesture *input.MultiTouchGesture) {
		direction := "unknown"
		switch gesture.Type {
		case input.TouchGestureSwipeLeft:
			direction = "left"
		case input.TouchGestureSwipeRight:
			direction = "right"
		case input.TouchGestureSwipeUp:
			direction = "up"
		case input.TouchGestureSwipeDown:
			direction = "down"
		}
		fmt.Printf("✓ Swipe gesture: direction=%s\n", direction)
	})

	// Simulate tap gesture
	touch1 := input.NewTouch(0, goflow.NewOffset(100, 100))
	tapStart := input.NewTouchEvent(
		input.EventTypeTouchStart,
		[]*input.Touch{touch1},
		[]*input.Touch{touch1},
		[]*input.Touch{touch1},
	)
	recognizer.ProcessTouchEvent(tapStart)

	time.Sleep(50 * time.Millisecond)

	tapEnd := input.NewTouchEvent(
		input.EventTypeTouchEnd,
		[]*input.Touch{},
		[]*input.Touch{touch1},
		[]*input.Touch{},
	)
	recognizer.ProcessTouchEvent(tapEnd)

	// Simulate pinch gesture
	recognizer.Reset()
	touch2 := input.NewTouch(0, goflow.NewOffset(100, 100))
	touch3 := input.NewTouch(1, goflow.NewOffset(200, 100))

	pinchStart := input.NewTouchEvent(
		input.EventTypeTouchStart,
		[]*input.Touch{touch2, touch3},
		[]*input.Touch{touch2, touch3},
		[]*input.Touch{touch2, touch3},
	)
	recognizer.ProcessTouchEvent(pinchStart)

	// Move touches closer (pinch in)
	touch2.Update(goflow.NewOffset(125, 100), 1.0)
	touch3.Update(goflow.NewOffset(175, 100), 1.0)

	pinchMove := input.NewTouchEvent(
		input.EventTypeTouchMove,
		[]*input.Touch{touch2, touch3},
		[]*input.Touch{touch2, touch3},
		[]*input.Touch{touch2, touch3},
	)
	recognizer.ProcessTouchEvent(pinchMove)
}

func demonstrateEventDispatcher(router *input.InputRouter) {
	fmt.Println("\n5. Event Dispatcher Demo")
	fmt.Println("------------------------")

	dispatcher := router.GetDispatcher()

	// Add event listeners with different priorities
	dispatcher.AddListener(input.EventTypeKeyDown, func(event input.InputEvent) input.EventPropagation {
		kbEvent := event.(*input.KeyboardEvent)
		fmt.Printf("✓ High priority listener: Key %d pressed\n", kbEvent.Key)
		return input.PropagationContinue
	}, 100)

	dispatcher.AddListener(input.EventTypeKeyDown, func(event input.InputEvent) input.EventPropagation {
		fmt.Println("✓ Low priority listener: Key event received")
		return input.PropagationContinue
	}, 10)

	// Add global listener
	dispatcher.AddGlobalListener(func(event input.InputEvent) input.EventPropagation {
		fmt.Printf("✓ Global listener: Event type %d\n", event.Type())
		return input.PropagationContinue
	}, 50)

	// Add event filter
	filter := input.NewEventTypeFilter(input.EventTypeKeyDown, input.EventTypeKeyUp)
	dispatcher.AddFilter(filter)

	// Dispatch events
	keyEvent := input.NewKeyboardEvent(input.EventTypeKeyDown, input.KeyA, "KeyA", "a", input.ModifierNone, false)
	dispatcher.Dispatch(keyEvent)

	// Mouse event should be filtered
	mouseEvent := input.NewMouseEvent(input.EventTypeMouseMove, goflow.NewOffset(0, 0), input.MouseButtonNone, input.MouseButtonsNone, 0)
	dispatcher.Dispatch(mouseEvent)

	// Check stats
	stats := dispatcher.GetStats()
	fmt.Printf("\n✓ Dispatcher stats: Total=%d, Dispatched=%d, Filtered=%d\n",
		stats.TotalEvents, stats.DispatchedEvents, stats.FilteredEvents)
}
