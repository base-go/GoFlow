package input

import (
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// MouseEvent represents mouse input events
type MouseEvent struct {
	ButtonEvent
	Delta *goflow.Offset // Movement delta for move events
}

// NewMouseEvent creates a new mouse event
func NewMouseEvent(eventType EventType, position *goflow.Offset, button MouseButton, buttons MouseButtons, clickCount int) *MouseEvent {
	return &MouseEvent{
		ButtonEvent: ButtonEvent{
			PositionEvent: PositionEvent{
				BaseEvent: BaseEvent{
					EventType:  eventType,
					Time:       time.Now(),
					DeviceType: DeviceTypeMouse,
				},
				Position: position,
				Global:   position, // Can be set separately if needed
			},
			Button:     button,
			Buttons:    buttons,
			ClickCount: clickCount,
		},
		Delta: goflow.ZeroOffset(),
	}
}

// MouseWheelEvent represents mouse wheel scrolling
type MouseWheelEvent struct {
	PositionEvent
	DeltaX         float64          // Horizontal scroll delta
	DeltaY         float64          // Vertical scroll delta
	DeltaMode      WheelDeltaMode   // Delta units (pixels, lines, pages)
	Modifiers      KeyModifiers     // Active keyboard modifiers
}

// NewMouseWheelEvent creates a new mouse wheel event
func NewMouseWheelEvent(position *goflow.Offset, deltaX, deltaY float64, deltaMode WheelDeltaMode) *MouseWheelEvent {
	return &MouseWheelEvent{
		PositionEvent: PositionEvent{
			BaseEvent: BaseEvent{
				EventType:  EventTypeMouseWheel,
				Time:       time.Now(),
				DeviceType: DeviceTypeMouse,
			},
			Position: position,
			Global:   position,
		},
		DeltaX:    deltaX,
		DeltaY:    deltaY,
		DeltaMode: deltaMode,
	}
}

// WheelDeltaMode represents mouse wheel delta units
type WheelDeltaMode int

const (
	WheelDeltaModePixel WheelDeltaMode = iota // Delta is in pixels
	WheelDeltaModeLine                        // Delta is in lines
	WheelDeltaModePage                        // Delta is in pages
)

// MouseState tracks the current state of the mouse
type MouseState struct {
	Position      *goflow.Offset // Current mouse position
	Buttons       MouseButtons   // Currently pressed buttons
	LastPosition  *goflow.Offset // Previous position (for delta calculation)
	IsInWindow    bool           // Whether mouse is in window bounds
	HoveredWidget interface{}    // Currently hovered widget (if any)
}

// NewMouseState creates a new mouse state tracker
func NewMouseState() *MouseState {
	return &MouseState{
		Position:     goflow.ZeroOffset(),
		LastPosition: goflow.ZeroOffset(),
		Buttons:      MouseButtonsNone,
		IsInWindow:   false,
	}
}

// IsButtonPressed checks if a specific button is pressed
func (s *MouseState) IsButtonPressed(button MouseButton) bool {
	return s.Buttons.HasButton(button)
}

// HandleEvent updates the mouse state based on an event
func (s *MouseState) HandleEvent(event *MouseEvent) {
	switch event.Type() {
	case EventTypeMouseDown:
		s.Buttons |= s.buttonToBitmask(event.Button)
		s.Position = event.Position
	case EventTypeMouseUp:
		s.Buttons &^= s.buttonToBitmask(event.Button)
		s.Position = event.Position
	case EventTypeMouseMove:
		s.LastPosition = s.Position
		s.Position = event.Position
		event.Delta = s.Position.Subtract(s.LastPosition)
	case EventTypeMouseEnter:
		s.IsInWindow = true
		s.Position = event.Position
	case EventTypeMouseLeave:
		s.IsInWindow = false
		s.Position = event.Position
	}
}

// buttonToBitmask converts a MouseButton to its bitmask
func (s *MouseState) buttonToBitmask(button MouseButton) MouseButtons {
	switch button {
	case MouseButtonLeft:
		return MouseButtonsLeft
	case MouseButtonMiddle:
		return MouseButtonsMiddle
	case MouseButtonRight:
		return MouseButtonsRight
	case MouseButtonBack:
		return MouseButtonsBack
	case MouseButtonForward:
		return MouseButtonsForward
	default:
		return MouseButtonsNone
	}
}

// GetDelta returns the movement delta since last position
func (s *MouseState) GetDelta() *goflow.Offset {
	return s.Position.Subtract(s.LastPosition)
}

// Reset clears the mouse state
func (s *MouseState) Reset() {
	s.Position = goflow.ZeroOffset()
	s.LastPosition = goflow.ZeroOffset()
	s.Buttons = MouseButtonsNone
	s.IsInWindow = false
	s.HoveredWidget = nil
}

// MouseRegion represents a region that can receive mouse events
type MouseRegion struct {
	Bounds    *goflow.Rect
	OnEnter   func(*MouseEvent)
	OnLeave   func(*MouseEvent)
	OnMove    func(*MouseEvent)
	OnDown    func(*MouseEvent)
	OnUp      func(*MouseEvent)
	OnClick   func(*MouseEvent)
	OnWheel   func(*MouseWheelEvent)
	Cursor    CursorType
	IsEnabled bool
	IsHovered bool
}

// NewMouseRegion creates a new mouse region
func NewMouseRegion(bounds *goflow.Rect) *MouseRegion {
	return &MouseRegion{
		Bounds:    bounds,
		IsEnabled: true,
		IsHovered: false,
		Cursor:    CursorDefault,
	}
}

// Contains checks if a position is within this region
func (r *MouseRegion) Contains(position *goflow.Offset) bool {
	return r.Bounds.Contains(position)
}

// HandleMouseEvent processes a mouse event for this region
func (r *MouseRegion) HandleMouseEvent(event *MouseEvent) bool {
	if !r.IsEnabled {
		return false
	}

	wasHovered := r.IsHovered
	r.IsHovered = r.Contains(event.Position)

	// Handle enter/leave
	if r.IsHovered && !wasHovered && r.OnEnter != nil {
		r.OnEnter(event)
		return true
	}
	if !r.IsHovered && wasHovered && r.OnLeave != nil {
		r.OnLeave(event)
		return true
	}

	// Only handle other events if hovered
	if !r.IsHovered {
		return false
	}

	switch event.Type() {
	case EventTypeMouseMove:
		if r.OnMove != nil {
			r.OnMove(event)
			return true
		}
	case EventTypeMouseDown:
		if r.OnDown != nil {
			r.OnDown(event)
			return true
		}
	case EventTypeMouseUp:
		if r.OnUp != nil {
			r.OnUp(event)
		}
		// Check for click (down and up in same region)
		if r.OnClick != nil {
			r.OnClick(event)
		}
		return true
	}

	return false
}

// HandleWheelEvent processes a wheel event for this region
func (r *MouseRegion) HandleWheelEvent(event *MouseWheelEvent) bool {
	if !r.IsEnabled || !r.Contains(event.Position) {
		return false
	}

	if r.OnWheel != nil {
		r.OnWheel(event)
		return true
	}

	return false
}

// CursorType represents different cursor types
type CursorType int

const (
	CursorDefault CursorType = iota
	CursorPointer
	CursorText
	CursorCrosshair
	CursorMove
	CursorNotAllowed
	CursorGrab
	CursorGrabbing
	CursorResizeNS
	CursorResizeEW
	CursorResizeNESW
	CursorResizeNWSE
	CursorZoomIn
	CursorZoomOut
	CursorHelp
	CursorWait
	CursorProgress
)

// MouseRegionManager manages multiple mouse regions
type MouseRegionManager struct {
	regions []*MouseRegion
}

// NewMouseRegionManager creates a new mouse region manager
func NewMouseRegionManager() *MouseRegionManager {
	return &MouseRegionManager{
		regions: make([]*MouseRegion, 0),
	}
}

// AddRegion adds a mouse region
func (m *MouseRegionManager) AddRegion(region *MouseRegion) {
	m.regions = append(m.regions, region)
}

// RemoveRegion removes a mouse region
func (m *MouseRegionManager) RemoveRegion(region *MouseRegion) {
	for i, r := range m.regions {
		if r == region {
			m.regions = append(m.regions[:i], m.regions[i+1:]...)
			return
		}
	}
}

// HandleMouseEvent dispatches a mouse event to appropriate regions
func (m *MouseRegionManager) HandleMouseEvent(event *MouseEvent) bool {
	// Process in reverse order (top to bottom)
	for i := len(m.regions) - 1; i >= 0; i-- {
		if m.regions[i].HandleMouseEvent(event) {
			return true
		}
	}
	return false
}

// HandleWheelEvent dispatches a wheel event to appropriate regions
func (m *MouseRegionManager) HandleWheelEvent(event *MouseWheelEvent) bool {
	// Process in reverse order (top to bottom)
	for i := len(m.regions) - 1; i >= 0; i-- {
		if m.regions[i].HandleWheelEvent(event) {
			return true
		}
	}
	return false
}

// GetCursorAt returns the cursor type for a position
func (m *MouseRegionManager) GetCursorAt(position *goflow.Offset) CursorType {
	// Process in reverse order (top to bottom)
	for i := len(m.regions) - 1; i >= 0; i-- {
		if m.regions[i].IsEnabled && m.regions[i].Contains(position) {
			return m.regions[i].Cursor
		}
	}
	return CursorDefault
}

// Clear removes all regions
func (m *MouseRegionManager) Clear() {
	m.regions = make([]*MouseRegion, 0)
}

// ScrollEvent is an alias for MouseWheelEvent
type ScrollEvent struct {
	PositionEvent
	ScrollDelta *goflow.Offset
}

// NewScrollEvent creates a new scroll event
func NewScrollEvent(position *goflow.Offset, delta *goflow.Offset) *ScrollEvent {
	return &ScrollEvent{
		PositionEvent: PositionEvent{
			BaseEvent: BaseEvent{
				EventType:  EventTypeMouseWheel,
				Time:       time.Now(),
				DeviceType: DeviceTypeMouse,
			},
			Position: position,
			Global:   position,
		},
		ScrollDelta: delta,
	}
}
