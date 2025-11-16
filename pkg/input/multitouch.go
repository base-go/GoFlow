package input

import (
	"math"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// MultiTouchGesture represents complex multi-touch gestures
type MultiTouchGesture struct {
	Type            TouchGesture
	TouchCount      int
	FocalPoint      *goflow.Offset  // Center point of all touches
	Scale           float64         // Pinch scale factor
	Rotation        float64         // Rotation angle in radians
	Velocity        *goflow.Offset  // Velocity vector
	Translation     *goflow.Offset  // Total translation
	StartTime       time.Time
	Duration        time.Duration
}

// NewMultiTouchGesture creates a new multi-touch gesture
func NewMultiTouchGesture(gestureType TouchGesture, touchCount int) *MultiTouchGesture {
	now := time.Now()
	return &MultiTouchGesture{
		Type:        gestureType,
		TouchCount:  touchCount,
		FocalPoint:  goflow.ZeroOffset(),
		Scale:       1.0,
		Rotation:    0.0,
		Velocity:    goflow.ZeroOffset(),
		Translation: goflow.ZeroOffset(),
		StartTime:   now,
		Duration:    0,
	}
}

// MultiTouchRecognizer provides advanced multi-touch gesture recognition
type MultiTouchRecognizer struct {
	// Configuration
	config *GestureConfig

	// Current state
	activeTouches map[int]*Touch
	activeGesture *MultiTouchGesture

	// Initial state (when gesture started)
	initialFocalPoint *goflow.Offset
	initialScale      float64
	initialRotation   float64
	initialTime       time.Time

	// Previous frame state (for velocity calculation)
	previousFocalPoint *goflow.Offset
	previousScale      float64
	previousRotation   float64
	previousTime       time.Time

	// Gesture callbacks
	onTap           func(*MultiTouchGesture)
	onDoubleTap     func(*MultiTouchGesture)
	onLongPress     func(*MultiTouchGesture)
	onPan           func(*MultiTouchGesture)
	onPinch         func(*MultiTouchGesture)
	onRotate        func(*MultiTouchGesture)
	onSwipe         func(*MultiTouchGesture)
	onGestureStart  func(*MultiTouchGesture)
	onGestureUpdate func(*MultiTouchGesture)
	onGestureEnd    func(*MultiTouchGesture)
}

// GestureConfig configures gesture recognition parameters
type GestureConfig struct {
	// Tap configuration
	TapTimeout       time.Duration // Max time for tap
	TapMaxMovement   float64       // Max movement for tap
	DoubleTapTimeout time.Duration // Max time between taps

	// Long press configuration
	LongPressTime     time.Duration // Min time for long press
	LongPressDistance float64       // Max movement for long press

	// Swipe configuration
	SwipeMinDistance float64 // Min distance for swipe
	SwipeMaxTime     time.Duration

	// Pan configuration
	PanMinDistance float64 // Min distance before pan starts

	// Pinch configuration
	PinchMinScale float64 // Min scale change to detect pinch

	// Rotation configuration
	RotationMinAngle float64 // Min angle change (in degrees)

	// General
	MaxTouchPoints int // Maximum simultaneous touch points to track
}

// DefaultGestureConfig returns default gesture configuration
func DefaultGestureConfig() *GestureConfig {
	return &GestureConfig{
		TapTimeout:        300 * time.Millisecond,
		TapMaxMovement:    10.0,
		DoubleTapTimeout:  300 * time.Millisecond,
		LongPressTime:     500 * time.Millisecond,
		LongPressDistance: 10.0,
		SwipeMinDistance:  50.0,
		SwipeMaxTime:      300 * time.Millisecond,
		PanMinDistance:    5.0,
		PinchMinScale:     0.1,
		RotationMinAngle:  5.0,
		MaxTouchPoints:    10,
	}
}

// NewMultiTouchRecognizer creates a new multi-touch recognizer
func NewMultiTouchRecognizer(config *GestureConfig) *MultiTouchRecognizer {
	if config == nil {
		config = DefaultGestureConfig()
	}

	return &MultiTouchRecognizer{
		config:        config,
		activeTouches: make(map[int]*Touch),
	}
}

// SetOnTap sets the tap callback
func (r *MultiTouchRecognizer) SetOnTap(callback func(*MultiTouchGesture)) {
	r.onTap = callback
}

// SetOnDoubleTap sets the double tap callback
func (r *MultiTouchRecognizer) SetOnDoubleTap(callback func(*MultiTouchGesture)) {
	r.onDoubleTap = callback
}

// SetOnLongPress sets the long press callback
func (r *MultiTouchRecognizer) SetOnLongPress(callback func(*MultiTouchGesture)) {
	r.onLongPress = callback
}

// SetOnPan sets the pan callback
func (r *MultiTouchRecognizer) SetOnPan(callback func(*MultiTouchGesture)) {
	r.onPan = callback
}

// SetOnPinch sets the pinch callback
func (r *MultiTouchRecognizer) SetOnPinch(callback func(*MultiTouchGesture)) {
	r.onPinch = callback
}

// SetOnRotate sets the rotate callback
func (r *MultiTouchRecognizer) SetOnRotate(callback func(*MultiTouchGesture)) {
	r.onRotate = callback
}

// SetOnSwipe sets the swipe callback
func (r *MultiTouchRecognizer) SetOnSwipe(callback func(*MultiTouchGesture)) {
	r.onSwipe = callback
}

// SetOnGestureStart sets the gesture start callback
func (r *MultiTouchRecognizer) SetOnGestureStart(callback func(*MultiTouchGesture)) {
	r.onGestureStart = callback
}

// SetOnGestureUpdate sets the gesture update callback
func (r *MultiTouchRecognizer) SetOnGestureUpdate(callback func(*MultiTouchGesture)) {
	r.onGestureUpdate = callback
}

// SetOnGestureEnd sets the gesture end callback
func (r *MultiTouchRecognizer) SetOnGestureEnd(callback func(*MultiTouchGesture)) {
	r.onGestureEnd = callback
}

// ProcessTouchEvent processes a touch event
func (r *MultiTouchRecognizer) ProcessTouchEvent(event *TouchEvent) {
	switch event.Type() {
	case EventTypeTouchStart:
		r.handleTouchStart(event)
	case EventTypeTouchMove:
		r.handleTouchMove(event)
	case EventTypeTouchEnd:
		r.handleTouchEnd(event)
	case EventTypeTouchCancel:
		r.handleTouchCancel()
	}
}

func (r *MultiTouchRecognizer) handleTouchStart(event *TouchEvent) {
	// Add new touches
	for _, touch := range event.ChangedTouches {
		if len(r.activeTouches) < r.config.MaxTouchPoints {
			r.activeTouches[touch.ID] = touch.Clone()
		}
	}

	// Initialize gesture if this is the first touch
	if len(r.activeTouches) == len(event.ChangedTouches) {
		r.initialTime = time.Now()
		r.previousTime = r.initialTime

		r.initialFocalPoint = r.calculateFocalPoint()
		r.previousFocalPoint = r.initialFocalPoint

		if len(r.activeTouches) >= 2 {
			r.initialScale = r.calculateScale()
			r.initialRotation = r.calculateRotation()
			r.previousScale = r.initialScale
			r.previousRotation = r.initialRotation
		} else {
			r.initialScale = 1.0
			r.initialRotation = 0.0
		}

		// Create gesture
		r.activeGesture = NewMultiTouchGesture(TouchGestureNone, len(r.activeTouches))
		r.activeGesture.FocalPoint = r.initialFocalPoint

		if r.onGestureStart != nil {
			r.onGestureStart(r.activeGesture)
		}
	}
}

func (r *MultiTouchRecognizer) handleTouchMove(event *TouchEvent) {
	// Update touches
	for _, touch := range event.ChangedTouches {
		if existing, ok := r.activeTouches[touch.ID]; ok {
			existing.Update(touch.Position, touch.Force)
		}
	}

	if r.activeGesture == nil {
		return
	}

	now := time.Now()

	// Calculate current state
	currentFocalPoint := r.calculateFocalPoint()
	currentScale := 1.0
	currentRotation := 0.0

	if len(r.activeTouches) >= 2 {
		currentScale = r.calculateScale()
		currentRotation = r.calculateRotation()
	}

	// Update gesture
	r.activeGesture.FocalPoint = currentFocalPoint
	r.activeGesture.Translation = currentFocalPoint.Subtract(r.initialFocalPoint)
	r.activeGesture.Duration = now.Sub(r.initialTime)

	// Calculate velocity
	dt := now.Sub(r.previousTime).Seconds()
	if dt > 0 {
		delta := currentFocalPoint.Subtract(r.previousFocalPoint)
		r.activeGesture.Velocity = &goflow.Offset{
			X: delta.X / dt,
			Y: delta.Y / dt,
		}
	}

	// Detect gesture type
	touchCount := len(r.activeTouches)

	if touchCount == 1 {
		// Single touch - pan or swipe
		distance := r.calculateDistance(r.initialFocalPoint, currentFocalPoint)
		if distance > r.config.PanMinDistance {
			r.activeGesture.Type = TouchGesturePan
			if r.onPan != nil {
				r.onPan(r.activeGesture)
			}
		}
	} else if touchCount >= 2 {
		// Multi-touch - pinch and/or rotate
		scaleChange := math.Abs(currentScale/r.initialScale - 1.0)
		rotationChange := math.Abs(currentRotation - r.initialRotation)

		r.activeGesture.Scale = currentScale / r.initialScale
		r.activeGesture.Rotation = currentRotation - r.initialRotation

		if scaleChange > r.config.PinchMinScale {
			if r.activeGesture.Scale > 1.0 {
				r.activeGesture.Type = TouchGesturePinchOut
			} else {
				r.activeGesture.Type = TouchGesturePinchIn
			}
			if r.onPinch != nil {
				r.onPinch(r.activeGesture)
			}
		}

		if rotationChange > r.config.RotationMinAngle*math.Pi/180.0 {
			r.activeGesture.Type = TouchGestureRotate
			if r.onRotate != nil {
				r.onRotate(r.activeGesture)
			}
		}
	}

	// Update previous state
	r.previousFocalPoint = currentFocalPoint
	r.previousScale = currentScale
	r.previousRotation = currentRotation
	r.previousTime = now

	if r.onGestureUpdate != nil {
		r.onGestureUpdate(r.activeGesture)
	}
}

func (r *MultiTouchRecognizer) handleTouchEnd(event *TouchEvent) {
	// Remove ended touches
	for _, touch := range event.ChangedTouches {
		delete(r.activeTouches, touch.ID)
	}

	if r.activeGesture == nil {
		return
	}

	// Update duration
	r.activeGesture.Duration = time.Now().Sub(r.initialTime)

	// Detect final gesture type if not already set
	if r.activeGesture.Type == TouchGestureNone && r.activeGesture.TouchCount == 1 {
		distance := r.calculateDistance(r.initialFocalPoint, r.activeGesture.FocalPoint)

		// Check for long press
		if r.activeGesture.Duration >= r.config.LongPressTime &&
			distance <= r.config.LongPressDistance {
			r.activeGesture.Type = TouchGestureLongPress
			if r.onLongPress != nil {
				r.onLongPress(r.activeGesture)
			}
		} else if distance >= r.config.SwipeMinDistance &&
			r.activeGesture.Duration <= r.config.SwipeMaxTime {
			// Swipe detection
			r.activeGesture.Type = r.detectSwipeDirection()
			if r.onSwipe != nil {
				r.onSwipe(r.activeGesture)
			}
		} else if r.activeGesture.Duration <= r.config.TapTimeout &&
			distance <= r.config.TapMaxMovement {
			// Tap
			r.activeGesture.Type = TouchGestureTap
			if r.onTap != nil {
				r.onTap(r.activeGesture)
			}
		}
	}

	// Call gesture end
	if r.onGestureEnd != nil {
		r.onGestureEnd(r.activeGesture)
	}

	// Clear gesture if no more touches
	if len(r.activeTouches) == 0 {
		r.activeGesture = nil
	}
}

func (r *MultiTouchRecognizer) handleTouchCancel() {
	r.activeTouches = make(map[int]*Touch)
	r.activeGesture = nil
}

// calculateFocalPoint calculates the center point of all active touches
func (r *MultiTouchRecognizer) calculateFocalPoint() *goflow.Offset {
	if len(r.activeTouches) == 0 {
		return goflow.ZeroOffset()
	}

	sumX := 0.0
	sumY := 0.0
	count := 0

	for _, touch := range r.activeTouches {
		sumX += touch.Position.X
		sumY += touch.Position.Y
		count++
	}

	return &goflow.Offset{
		X: sumX / float64(count),
		Y: sumY / float64(count),
	}
}

// calculateScale calculates the average distance between touches
func (r *MultiTouchRecognizer) calculateScale() float64 {
	if len(r.activeTouches) < 2 {
		return 1.0
	}

	// Get all touches as slice
	touches := make([]*Touch, 0, len(r.activeTouches))
	for _, touch := range r.activeTouches {
		touches = append(touches, touch)
	}

	// Calculate average distance from focal point
	focalPoint := r.calculateFocalPoint()
	sumDistance := 0.0

	for _, touch := range touches {
		sumDistance += r.calculateDistance(touch.Position, focalPoint)
	}

	return sumDistance / float64(len(touches))
}

// calculateRotation calculates the rotation angle of touches
func (r *MultiTouchRecognizer) calculateRotation() float64 {
	if len(r.activeTouches) < 2 {
		return 0.0
	}

	// Get first two touches
	touches := make([]*Touch, 0, 2)
	for _, touch := range r.activeTouches {
		touches = append(touches, touch)
		if len(touches) == 2 {
			break
		}
	}

	// Calculate angle between first two touches
	dx := touches[1].Position.X - touches[0].Position.X
	dy := touches[1].Position.Y - touches[0].Position.Y
	return math.Atan2(dy, dx)
}

// calculateDistance calculates distance between two points
func (r *MultiTouchRecognizer) calculateDistance(p1, p2 *goflow.Offset) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// detectSwipeDirection detects swipe direction based on translation
func (r *MultiTouchRecognizer) detectSwipeDirection() TouchGesture {
	if r.activeGesture == nil || r.activeGesture.Translation == nil {
		return TouchGestureNone
	}

	dx := r.activeGesture.Translation.X
	dy := r.activeGesture.Translation.Y

	// Determine primary direction
	if math.Abs(dx) > math.Abs(dy) {
		if dx > 0 {
			return TouchGestureSwipeRight
		}
		return TouchGestureSwipeLeft
	} else {
		if dy > 0 {
			return TouchGestureSwipeDown
		}
		return TouchGestureSwipeUp
	}
}

// GetActiveGesture returns the currently active gesture
func (r *MultiTouchRecognizer) GetActiveGesture() *MultiTouchGesture {
	return r.activeGesture
}

// GetActiveTouchCount returns the number of active touches
func (r *MultiTouchRecognizer) GetActiveTouchCount() int {
	return len(r.activeTouches)
}

// Reset clears all gesture state
func (r *MultiTouchRecognizer) Reset() {
	r.activeTouches = make(map[int]*Touch)
	r.activeGesture = nil
}
