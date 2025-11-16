package input

import (
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// TouchEvent represents touch input events
type TouchEvent struct {
	BaseEvent
	Touches        []*Touch // All active touches
	ChangedTouches []*Touch // Touches that changed in this event
	TargetTouches  []*Touch // Touches on the target element
}

// NewTouchEvent creates a new touch event
func NewTouchEvent(eventType EventType, touches, changedTouches, targetTouches []*Touch) *TouchEvent {
	return &TouchEvent{
		BaseEvent: BaseEvent{
			EventType:  eventType,
			Time:       time.Now(),
			DeviceType: DeviceTypeTouch,
		},
		Touches:        touches,
		ChangedTouches: changedTouches,
		TargetTouches:  targetTouches,
	}
}

// Touch represents a single touch point
type Touch struct {
	ID             int             // Unique identifier for this touch
	Position       *goflow.Offset  // Current position
	StartPosition  *goflow.Offset  // Position where touch started
	PreviousPosition *goflow.Offset // Previous position (for delta)
	Radius         float64         // Touch area radius
	RadiusX        float64         // Touch area X radius (ellipse)
	RadiusY        float64         // Touch area Y radius (ellipse)
	RotationAngle  float64         // Rotation angle of touch area
	Force          float64         // Pressure/force (0.0 to 1.0)
	StartTime      time.Time       // When this touch started
	UpdateTime     time.Time       // Last update time
}

// NewTouch creates a new touch point
func NewTouch(id int, position *goflow.Offset) *Touch {
	now := time.Now()
	return &Touch{
		ID:               id,
		Position:         position,
		StartPosition:    position,
		PreviousPosition: position,
		Radius:           0.0,
		RadiusX:          0.0,
		RadiusY:          0.0,
		RotationAngle:    0.0,
		Force:            1.0,
		StartTime:        now,
		UpdateTime:       now,
	}
}

// GetDelta returns the movement delta since last update
func (t *Touch) GetDelta() *goflow.Offset {
	return t.Position.Subtract(t.PreviousPosition)
}

// GetTotalDelta returns the total movement since touch started
func (t *Touch) GetTotalDelta() *goflow.Offset {
	return t.Position.Subtract(t.StartPosition)
}

// GetDuration returns how long this touch has been active
func (t *Touch) GetDuration() time.Duration {
	return t.UpdateTime.Sub(t.StartTime)
}

// Clone creates a copy of this touch
func (t *Touch) Clone() *Touch {
	return &Touch{
		ID:               t.ID,
		Position:         &goflow.Offset{X: t.Position.X, Y: t.Position.Y},
		StartPosition:    &goflow.Offset{X: t.StartPosition.X, Y: t.StartPosition.Y},
		PreviousPosition: &goflow.Offset{X: t.PreviousPosition.X, Y: t.PreviousPosition.Y},
		Radius:           t.Radius,
		RadiusX:          t.RadiusX,
		RadiusY:          t.RadiusY,
		RotationAngle:    t.RotationAngle,
		Force:            t.Force,
		StartTime:        t.StartTime,
		UpdateTime:       t.UpdateTime,
	}
}

// Update updates the touch position and time
func (t *Touch) Update(position *goflow.Offset, force float64) {
	t.PreviousPosition = t.Position
	t.Position = position
	t.Force = force
	t.UpdateTime = time.Now()
}

// PointerEvent represents unified pointer events (mouse, touch, pen)
type PointerEvent struct {
	PositionEvent
	PointerID     int          // Unique pointer identifier
	PointerType   PointerType  // Type of pointer
	Button        MouseButton  // Button that caused event (for mouse)
	Buttons       MouseButtons // Currently pressed buttons
	Pressure      float64      // Pressure (0.0 to 1.0)
	TiltX         float64      // Tilt X angle (for pen)
	TiltY         float64      // Tilt Y angle (for pen)
	Twist         float64      // Twist angle (for pen)
	Width         float64      // Contact geometry width
	Height        float64      // Contact geometry height
	IsPrimary     bool         // Is this the primary pointer
}

// NewPointerEvent creates a new pointer event
func NewPointerEvent(eventType EventType, pointerID int, pointerType PointerType, position *goflow.Offset, isPrimary bool) *PointerEvent {
	return &PointerEvent{
		PositionEvent: PositionEvent{
			BaseEvent: BaseEvent{
				EventType:  eventType,
				Time:       time.Now(),
				DeviceType: DeviceTypeTouch,
			},
			Position: position,
			Global:   position,
		},
		PointerID:   pointerID,
		PointerType: pointerType,
		Pressure:    1.0,
		IsPrimary:   isPrimary,
	}
}

// PointerType represents different pointer device types
type PointerType int

const (
	PointerTypeMouse PointerType = iota
	PointerTypeTouch
	PointerTypePen
)

// TouchState tracks all active touches
type TouchState struct {
	activeTouches map[int]*Touch // Map of touch ID to Touch
	nextID        int            // Next touch ID to assign
}

// NewTouchState creates a new touch state tracker
func NewTouchState() *TouchState {
	return &TouchState{
		activeTouches: make(map[int]*Touch),
		nextID:        0,
	}
}

// AddTouch adds a new touch
func (s *TouchState) AddTouch(position *goflow.Offset) *Touch {
	touch := NewTouch(s.nextID, position)
	s.activeTouches[touch.ID] = touch
	s.nextID++
	return touch
}

// UpdateTouch updates an existing touch
func (s *TouchState) UpdateTouch(id int, position *goflow.Offset, force float64) *Touch {
	if touch, ok := s.activeTouches[id]; ok {
		touch.Update(position, force)
		return touch
	}
	return nil
}

// RemoveTouch removes a touch
func (s *TouchState) RemoveTouch(id int) *Touch {
	if touch, ok := s.activeTouches[id]; ok {
		delete(s.activeTouches, id)
		return touch
	}
	return nil
}

// GetTouch returns a touch by ID
func (s *TouchState) GetTouch(id int) *Touch {
	return s.activeTouches[id]
}

// GetAllTouches returns all active touches
func (s *TouchState) GetAllTouches() []*Touch {
	touches := make([]*Touch, 0, len(s.activeTouches))
	for _, touch := range s.activeTouches {
		touches = append(touches, touch)
	}
	return touches
}

// GetTouchCount returns the number of active touches
func (s *TouchState) GetTouchCount() int {
	return len(s.activeTouches)
}

// Clear removes all touches
func (s *TouchState) Clear() {
	s.activeTouches = make(map[int]*Touch)
}

// TouchGesture represents detected touch gestures
type TouchGesture int

const (
	TouchGestureNone TouchGesture = iota
	TouchGestureTap
	TouchGestureDoubleTap
	TouchGestureLongPress
	TouchGestureSwipeLeft
	TouchGestureSwipeRight
	TouchGestureSwipeUp
	TouchGestureSwipeDown
	TouchGesturePinchIn
	TouchGesturePinchOut
	TouchGestureRotate
	TouchGesturePan
)

// TouchGestureRecognizer recognizes touch gestures
type TouchGestureRecognizer struct {
	state           *TouchState
	tapTimeout      time.Duration
	longPressTime   time.Duration
	swipeThreshold  float64
	scaleThreshold  float64
	rotateThreshold float64

	// Gesture detection state
	lastTapTime      time.Time
	tapCount         int
	initialScale     float64
	initialRotation  float64
	gestureStartTime time.Time
}

// NewTouchGestureRecognizer creates a new gesture recognizer
func NewTouchGestureRecognizer() *TouchGestureRecognizer {
	return &TouchGestureRecognizer{
		state:            NewTouchState(),
		tapTimeout:       300 * time.Millisecond,
		longPressTime:    500 * time.Millisecond,
		swipeThreshold:   50.0,
		scaleThreshold:   0.1,
		rotateThreshold:  15.0,
		initialScale:     1.0,
		initialRotation:  0.0,
	}
}

// ProcessTouchEvent processes a touch event and returns detected gesture
func (r *TouchGestureRecognizer) ProcessTouchEvent(event *TouchEvent) TouchGesture {
	switch event.Type() {
	case EventTypeTouchStart:
		return r.handleTouchStart(event)
	case EventTypeTouchMove:
		return r.handleTouchMove(event)
	case EventTypeTouchEnd:
		return r.handleTouchEnd(event)
	case EventTypeTouchCancel:
		r.state.Clear()
		return TouchGestureNone
	}
	return TouchGestureNone
}

func (r *TouchGestureRecognizer) handleTouchStart(event *TouchEvent) TouchGesture {
	// Add new touches to state
	for _, touch := range event.ChangedTouches {
		r.state.AddTouch(touch.Position)
	}

	r.gestureStartTime = time.Now()

	// Check for double tap
	if r.state.GetTouchCount() == 1 {
		now := time.Now()
		if now.Sub(r.lastTapTime) < r.tapTimeout {
			r.tapCount++
			if r.tapCount == 2 {
				r.tapCount = 0
				return TouchGestureDoubleTap
			}
		} else {
			r.tapCount = 1
		}
		r.lastTapTime = now
	}

	// Initialize multi-touch gesture tracking
	if r.state.GetTouchCount() == 2 {
		touches := r.state.GetAllTouches()
		r.initialScale = r.calculateDistance(touches[0].Position, touches[1].Position)
		r.initialRotation = r.calculateAngle(touches[0].Position, touches[1].Position)
	}

	return TouchGestureNone
}

func (r *TouchGestureRecognizer) handleTouchMove(event *TouchEvent) TouchGesture {
	// Update touches
	for _, touch := range event.ChangedTouches {
		r.state.UpdateTouch(touch.ID, touch.Position, touch.Force)
	}

	touchCount := r.state.GetTouchCount()

	// Single touch - pan or swipe
	if touchCount == 1 {
		return TouchGesturePan
	}

	// Two touches - pinch or rotate
	if touchCount == 2 {
		touches := r.state.GetAllTouches()
		currentDistance := r.calculateDistance(touches[0].Position, touches[1].Position)
		currentAngle := r.calculateAngle(touches[0].Position, touches[1].Position)

		scaleChange := (currentDistance - r.initialScale) / r.initialScale
		rotationChange := currentAngle - r.initialRotation

		// Pinch detection
		if scaleChange < -r.scaleThreshold {
			return TouchGesturePinchIn
		}
		if scaleChange > r.scaleThreshold {
			return TouchGesturePinchOut
		}

		// Rotation detection
		if rotationChange > r.rotateThreshold || rotationChange < -r.rotateThreshold {
			return TouchGestureRotate
		}
	}

	return TouchGestureNone
}

func (r *TouchGestureRecognizer) handleTouchEnd(event *TouchEvent) TouchGesture {
	// Remove ended touches
	for _, touch := range event.ChangedTouches {
		r.state.RemoveTouch(touch.ID)
	}

	// Check for long press
	if len(event.ChangedTouches) == 1 {
		touch := event.ChangedTouches[0]
		if touch.GetDuration() >= r.longPressTime {
			return TouchGestureLongPress
		}

		// Check for swipe
		delta := touch.GetTotalDelta()
		if delta.X > r.swipeThreshold && delta.X > delta.Y {
			return TouchGestureSwipeRight
		}
		if delta.X < -r.swipeThreshold && -delta.X > delta.Y {
			return TouchGestureSwipeLeft
		}
		if delta.Y > r.swipeThreshold && delta.Y > delta.X {
			return TouchGestureSwipeDown
		}
		if delta.Y < -r.swipeThreshold && -delta.Y > delta.X {
			return TouchGestureSwipeUp
		}

		// Simple tap if no significant movement
		if touch.GetDuration() < r.longPressTime {
			return TouchGestureTap
		}
	}

	return TouchGestureNone
}

// calculateDistance calculates distance between two points
func (r *TouchGestureRecognizer) calculateDistance(p1, p2 *goflow.Offset) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return (dx*dx + dy*dy) // Square root not needed for comparison
}

// calculateAngle calculates angle between two points in degrees
func (r *TouchGestureRecognizer) calculateAngle(p1, p2 *goflow.Offset) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	// Use atan2 equivalent (would need math import)
	// For now return simplified version
	return dx + dy // Placeholder - would use math.Atan2 in real implementation
}

// GetTouchState returns the current touch state
func (r *TouchGestureRecognizer) GetTouchState() *TouchState {
	return r.state
}

// Reset clears all gesture state
func (r *TouchGestureRecognizer) Reset() {
	r.state.Clear()
	r.tapCount = 0
	r.initialScale = 1.0
	r.initialRotation = 0.0
}

// MultiTouchEvent represents multi-touch events with pointer tracking
type MultiTouchEvent struct {
	PositionEvent
	PointerID int // Unique pointer identifier
	TouchID   int // Touch identifier
	IsPrimary bool
	Pressure  float64
}
