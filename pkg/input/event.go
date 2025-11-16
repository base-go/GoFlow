package input

import (
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// EventType represents the type of input event
type EventType int

const (
	EventTypeUnknown EventType = iota

	// Keyboard events
	EventTypeKeyDown
	EventTypeKeyUp
	EventTypeKeyRepeat

	// Mouse events
	EventTypeMouseDown
	EventTypeMouseUp
	EventTypeMouseMove
	EventTypeMouseEnter
	EventTypeMouseLeave
	EventTypeMouseWheel

	// Touch events
	EventTypeTouchStart
	EventTypeTouchMove
	EventTypeTouchEnd
	EventTypeTouchCancel

	// Pointer events (unified mouse/touch/pen)
	EventTypePointerDown
	EventTypePointerUp
	EventTypePointerMove
	EventTypePointerEnter
	EventTypePointerLeave
	EventTypePointerCancel
	EventTypePointerWheel
)

// InputEvent is the base interface for all input events
type InputEvent interface {
	Type() EventType
	Timestamp() time.Time
	Device() DeviceType
}

// DeviceType represents the input device type
type DeviceType int

const (
	DeviceTypeUnknown DeviceType = iota
	DeviceTypeKeyboard
	DeviceTypeMouse
	DeviceTypeTouch
	DeviceTypePen
	DeviceTypeTrackpad
)

// BaseEvent provides common event functionality
type BaseEvent struct {
	EventType  EventType
	Time       time.Time
	DeviceType DeviceType
}

func (e *BaseEvent) Type() EventType {
	return e.EventType
}

func (e *BaseEvent) Timestamp() time.Time {
	return e.Time
}

func (e *BaseEvent) Device() DeviceType {
	return e.DeviceType
}

// PositionEvent represents events with a position
type PositionEvent struct {
	BaseEvent
	Position *goflow.Offset // Position relative to widget/window
	Global   *goflow.Offset // Global screen position
}

// ButtonEvent represents events with button state
type ButtonEvent struct {
	PositionEvent
	Button       MouseButton
	Buttons      MouseButtons // Bitmask of currently pressed buttons
	ClickCount   int          // Number of clicks (1=single, 2=double, etc.)
}

// MouseButton represents mouse buttons
type MouseButton int

const (
	MouseButtonNone MouseButton = iota
	MouseButtonLeft
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonBack
	MouseButtonForward
)

// MouseButtons is a bitmask of mouse buttons
type MouseButtons int

const (
	MouseButtonsNone    MouseButtons = 0
	MouseButtonsLeft    MouseButtons = 1 << 0
	MouseButtonsMiddle  MouseButtons = 1 << 1
	MouseButtonsRight   MouseButtons = 1 << 2
	MouseButtonsBack    MouseButtons = 1 << 3
	MouseButtonsForward MouseButtons = 1 << 4
)

// HasButton checks if a button is pressed
func (b MouseButtons) HasButton(button MouseButton) bool {
	switch button {
	case MouseButtonLeft:
		return b&MouseButtonsLeft != 0
	case MouseButtonMiddle:
		return b&MouseButtonsMiddle != 0
	case MouseButtonRight:
		return b&MouseButtonsRight != 0
	case MouseButtonBack:
		return b&MouseButtonsBack != 0
	case MouseButtonForward:
		return b&MouseButtonsForward != 0
	default:
		return false
	}
}

// EventPropagation controls event flow
type EventPropagation int

const (
	PropagationContinue EventPropagation = iota // Continue propagation
	PropagationStop                             // Stop propagation to other handlers
	PropagationStopImmediate                    // Stop all propagation immediately
)

// EventHandler is a function that handles input events
type EventHandler func(event InputEvent) EventPropagation

// EventListener represents an event listener registration
type EventListener struct {
	EventType EventType
	Handler   EventHandler
	Priority  int  // Higher priority handlers are called first
	Once      bool // If true, handler is removed after first invocation
}
