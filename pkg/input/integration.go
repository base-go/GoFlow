package input

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// GestureEventAdapter adapts low-level input events to high-level gesture callbacks
type GestureEventAdapter struct {
	gestureDetector *widgets.GestureDetector
	touchRecognizer *MultiTouchRecognizer
	lastTapTime     int64
}

// NewGestureEventAdapter creates a new gesture event adapter
func NewGestureEventAdapter(detector *widgets.GestureDetector) *GestureEventAdapter {
	adapter := &GestureEventAdapter{
		gestureDetector: detector,
		touchRecognizer: NewMultiTouchRecognizer(nil),
	}

	// Setup gesture callbacks
	adapter.setupGestureCallbacks()

	return adapter
}

// setupGestureCallbacks configures the touch recognizer callbacks
func (a *GestureEventAdapter) setupGestureCallbacks() {
	// Tap gesture
	a.touchRecognizer.SetOnTap(func(gesture *MultiTouchGesture) {
		if a.gestureDetector.OnTap != nil {
			a.gestureDetector.OnTap()
		}
	})

	// Double tap gesture
	a.touchRecognizer.SetOnDoubleTap(func(gesture *MultiTouchGesture) {
		if a.gestureDetector.OnDoubleTap != nil {
			a.gestureDetector.OnDoubleTap()
		}
	})

	// Long press gesture
	a.touchRecognizer.SetOnLongPress(func(gesture *MultiTouchGesture) {
		if a.gestureDetector.OnLongPress != nil {
			a.gestureDetector.OnLongPress()
		}
	})

	// Pan gesture
	a.touchRecognizer.SetOnGestureStart(func(gesture *MultiTouchGesture) {
		if gesture.Type == TouchGesturePan && a.gestureDetector.OnPanStart != nil {
			details := widgets.DragStartDetails{
				GlobalPosition: *gesture.FocalPoint,
				LocalPosition:  *gesture.FocalPoint,
			}
			a.gestureDetector.OnPanStart(details)
		}

		if gesture.TouchCount >= 2 && a.gestureDetector.OnScaleStart != nil {
			details := widgets.ScaleStartDetails{
				FocalPoint: *gesture.FocalPoint,
			}
			a.gestureDetector.OnScaleStart(details)
		}
	})

	a.touchRecognizer.SetOnGestureUpdate(func(gesture *MultiTouchGesture) {
		// Pan update
		if gesture.Type == TouchGesturePan && a.gestureDetector.OnPanUpdate != nil {
			details := widgets.DragUpdateDetails{
				GlobalPosition: *gesture.FocalPoint,
				LocalPosition:  *gesture.FocalPoint,
				Delta:          *gesture.Translation,
			}
			a.gestureDetector.OnPanUpdate(details)
		}

		// Scale update (pinch or rotate)
		if gesture.TouchCount >= 2 && a.gestureDetector.OnScaleUpdate != nil {
			details := widgets.ScaleUpdateDetails{
				FocalPoint: *gesture.FocalPoint,
				Scale:      gesture.Scale,
				Rotation:   gesture.Rotation,
			}
			a.gestureDetector.OnScaleUpdate(details)
		}
	})

	a.touchRecognizer.SetOnGestureEnd(func(gesture *MultiTouchGesture) {
		// Pan end
		if gesture.Type == TouchGesturePan && a.gestureDetector.OnPanEnd != nil {
			velocity := goflow.Offset{X: 0, Y: 0}
			pv := 0.0
			if gesture.Velocity != nil && gesture.Duration.Seconds() > 0 {
				// Calculate velocity magnitude
				vx := gesture.Velocity.X
				vy := gesture.Velocity.Y
				pv = vx*vx + vy*vy // Simplified - sqrt not needed for comparison
				velocity = *gesture.Velocity
			}

			details := widgets.DragEndDetails{
				Velocity:        velocity,
				PrimaryVelocity: pv,
			}
			a.gestureDetector.OnPanEnd(details)
		}

		// Scale end
		if gesture.TouchCount >= 2 && a.gestureDetector.OnScaleEnd != nil {
			details := widgets.ScaleEndDetails{
				Velocity: 0.0, // Could calculate from gesture velocity
			}
			a.gestureDetector.OnScaleEnd(details)
		}
	})
}

// HandleTouchEvent processes a touch event
func (a *GestureEventAdapter) HandleTouchEvent(event *TouchEvent) {
	a.touchRecognizer.ProcessTouchEvent(event)
}

// HandleMouseEvent processes a mouse event as a touch event
func (a *GestureEventAdapter) HandleMouseEvent(event *MouseEvent) {
	// Convert mouse events to touch events for gesture detection
	switch event.Type() {
	case EventTypeMouseDown:
		if event.Button == MouseButtonLeft {
			touch := NewTouch(0, event.Position)
			touchEvent := NewTouchEvent(
				EventTypeTouchStart,
				[]*Touch{touch},
				[]*Touch{touch},
				[]*Touch{touch},
			)
			a.HandleTouchEvent(touchEvent)
		}

	case EventTypeMouseMove:
		if event.Buttons.HasButton(MouseButtonLeft) {
			touch := NewTouch(0, event.Position)
			touchEvent := NewTouchEvent(
				EventTypeTouchMove,
				[]*Touch{touch},
				[]*Touch{touch},
				[]*Touch{touch},
			)
			a.HandleTouchEvent(touchEvent)
		}

	case EventTypeMouseUp:
		if event.Button == MouseButtonLeft {
			touch := NewTouch(0, event.Position)
			touchEvent := NewTouchEvent(
				EventTypeTouchEnd,
				[]*Touch{},
				[]*Touch{touch},
				[]*Touch{},
			)
			a.HandleTouchEvent(touchEvent)
		}
	}

	// Also handle direct click events
	if event.Type() == EventTypeMouseUp && event.Button == MouseButtonLeft {
		if event.ClickCount == 1 && a.gestureDetector.OnTap != nil {
			a.gestureDetector.OnTap()
		}
		if event.ClickCount == 2 && a.gestureDetector.OnDoubleTap != nil {
			a.gestureDetector.OnDoubleTap()
		}
	}
}

// WidgetInputMixin provides input handling capabilities to widgets
type WidgetInputMixin struct {
	// Keyboard handlers
	OnKeyDown   func(*KeyboardEvent) EventPropagation
	OnKeyUp     func(*KeyboardEvent) EventPropagation
	OnKeyRepeat func(*KeyboardEvent) EventPropagation

	// Mouse handlers
	OnMouseDown  func(*MouseEvent) EventPropagation
	OnMouseUp    func(*MouseEvent) EventPropagation
	OnMouseMove  func(*MouseEvent) EventPropagation
	OnMouseEnter func(*MouseEvent) EventPropagation
	OnMouseLeave func(*MouseEvent) EventPropagation
	OnMouseWheel func(*MouseWheelEvent) EventPropagation

	// Touch handlers
	OnTouchStart  func(*TouchEvent) EventPropagation
	OnTouchMove   func(*TouchEvent) EventPropagation
	OnTouchEnd    func(*TouchEvent) EventPropagation
	OnTouchCancel func(*TouchEvent) EventPropagation

	// Pointer handlers (unified)
	OnPointerDown   func(*PointerEvent) EventPropagation
	OnPointerUp     func(*PointerEvent) EventPropagation
	OnPointerMove   func(*PointerEvent) EventPropagation
	OnPointerEnter  func(*PointerEvent) EventPropagation
	OnPointerLeave  func(*PointerEvent) EventPropagation
	OnPointerCancel func(*PointerEvent) EventPropagation

	// Focus handlers
	OnFocus func() EventPropagation
	OnBlur  func() EventPropagation

	// State
	IsFocused     bool
	IsHovered     bool
	AcceptsInput  bool
	TabIndex      int
	InputBounds   *goflow.Rect
}

// NewWidgetInputMixin creates a new widget input mixin
func NewWidgetInputMixin() *WidgetInputMixin {
	return &WidgetInputMixin{
		AcceptsInput: true,
		TabIndex:     0,
	}
}

// HandleInputEvent routes an input event to the appropriate handler
func (m *WidgetInputMixin) HandleInputEvent(event InputEvent) EventPropagation {
	if !m.AcceptsInput {
		return PropagationContinue
	}

	switch e := event.(type) {
	case *KeyboardEvent:
		return m.handleKeyboardEvent(e)
	case *MouseEvent:
		return m.handleMouseEvent(e)
	case *MouseWheelEvent:
		return m.handleMouseWheelEvent(e)
	case *TouchEvent:
		return m.handleTouchEvent(e)
	case *PointerEvent:
		return m.handlePointerEvent(e)
	}

	return PropagationContinue
}

func (m *WidgetInputMixin) handleKeyboardEvent(event *KeyboardEvent) EventPropagation {
	if !m.IsFocused {
		return PropagationContinue
	}

	switch event.Type() {
	case EventTypeKeyDown:
		if m.OnKeyDown != nil {
			return m.OnKeyDown(event)
		}
	case EventTypeKeyUp:
		if m.OnKeyUp != nil {
			return m.OnKeyUp(event)
		}
	case EventTypeKeyRepeat:
		if m.OnKeyRepeat != nil {
			return m.OnKeyRepeat(event)
		}
	}

	return PropagationContinue
}

func (m *WidgetInputMixin) handleMouseEvent(event *MouseEvent) EventPropagation {
	// Check if event is within bounds
	if m.InputBounds != nil && !m.InputBounds.Contains(event.Position) {
		if m.IsHovered && m.OnMouseLeave != nil {
			m.IsHovered = false
			return m.OnMouseLeave(event)
		}
		return PropagationContinue
	}

	// Handle enter
	if !m.IsHovered && m.OnMouseEnter != nil {
		m.IsHovered = true
		result := m.OnMouseEnter(event)
		if result != PropagationContinue {
			return result
		}
	}

	switch event.Type() {
	case EventTypeMouseDown:
		if m.OnMouseDown != nil {
			return m.OnMouseDown(event)
		}
	case EventTypeMouseUp:
		if m.OnMouseUp != nil {
			return m.OnMouseUp(event)
		}
	case EventTypeMouseMove:
		if m.OnMouseMove != nil {
			return m.OnMouseMove(event)
		}
	}

	return PropagationContinue
}

func (m *WidgetInputMixin) handleMouseWheelEvent(event *MouseWheelEvent) EventPropagation {
	if m.InputBounds != nil && !m.InputBounds.Contains(event.Position) {
		return PropagationContinue
	}

	if m.OnMouseWheel != nil {
		return m.OnMouseWheel(event)
	}

	return PropagationContinue
}

func (m *WidgetInputMixin) handleTouchEvent(event *TouchEvent) EventPropagation {
	switch event.Type() {
	case EventTypeTouchStart:
		if m.OnTouchStart != nil {
			return m.OnTouchStart(event)
		}
	case EventTypeTouchMove:
		if m.OnTouchMove != nil {
			return m.OnTouchMove(event)
		}
	case EventTypeTouchEnd:
		if m.OnTouchEnd != nil {
			return m.OnTouchEnd(event)
		}
	case EventTypeTouchCancel:
		if m.OnTouchCancel != nil {
			return m.OnTouchCancel(event)
		}
	}

	return PropagationContinue
}

func (m *WidgetInputMixin) handlePointerEvent(event *PointerEvent) EventPropagation {
	switch event.Type() {
	case EventTypePointerDown:
		if m.OnPointerDown != nil {
			return m.OnPointerDown(event)
		}
	case EventTypePointerUp:
		if m.OnPointerUp != nil {
			return m.OnPointerUp(event)
		}
	case EventTypePointerMove:
		if m.OnPointerMove != nil {
			return m.OnPointerMove(event)
		}
	case EventTypePointerEnter:
		if m.OnPointerEnter != nil {
			return m.OnPointerEnter(event)
		}
	case EventTypePointerLeave:
		if m.OnPointerLeave != nil {
			return m.OnPointerLeave(event)
		}
	case EventTypePointerCancel:
		if m.OnPointerCancel != nil {
			return m.OnPointerCancel(event)
		}
	}

	return PropagationContinue
}

// Focus focuses the widget
func (m *WidgetInputMixin) Focus() EventPropagation {
	if m.IsFocused {
		return PropagationContinue
	}

	m.IsFocused = true
	if m.OnFocus != nil {
		return m.OnFocus()
	}

	return PropagationContinue
}

// Blur removes focus from the widget
func (m *WidgetInputMixin) Blur() EventPropagation {
	if !m.IsFocused {
		return PropagationContinue
	}

	m.IsFocused = false
	if m.OnBlur != nil {
		return m.OnBlur()
	}

	return PropagationContinue
}
