package testing

import (
	"fmt"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/input"
)

// GestureSimulator provides gesture simulation for testing
type GestureSimulator struct {
	tester *WidgetTester
}

// NewGestureSimulator creates a new gesture simulator
func NewGestureSimulator(tester *WidgetTester) *GestureSimulator {
	return &GestureSimulator{
		tester: tester,
	}
}

// GestureConfig configures gesture parameters
type GestureConfig struct {
	Duration time.Duration // Total duration of the gesture
	Steps    int           // Number of steps to break the gesture into
	Velocity float64       // Velocity for fling gestures (pixels/second)
}

// DefaultGestureConfig returns default gesture configuration
func DefaultGestureConfig() *GestureConfig {
	return &GestureConfig{
		Duration: time.Millisecond * 300,
		Steps:    10,
		Velocity: 1000.0,
	}
}

// Tap simulates a single tap
func (gs *GestureSimulator) Tap(x, y float64) error {
	return gs.tester.TapAt(&goflow.Offset{X: x, Y: y})
}

// DoubleTap simulates a double tap
func (gs *GestureSimulator) DoubleTap(x, y float64) error {
	if err := gs.Tap(x, y); err != nil {
		return err
	}

	// Wait a brief moment
	if err := gs.tester.Pump(time.Millisecond * 50); err != nil {
		return err
	}

	// Second tap
	return gs.Tap(x, y)
}

// LongPress simulates a long press gesture
func (gs *GestureSimulator) LongPress(x, y float64, duration time.Duration) error {
	offset := &goflow.Offset{X: x, Y: y}

	// Send pointer down
	downEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerDown,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: offset,
		Global:   offset,
	}
	gs.tester.binding.events = append(gs.tester.binding.events, downEvent)

	// Pump and wait
	if err := gs.tester.Pump(duration); err != nil {
		return err
	}

	// Send pointer up
	upEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerUp,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: offset,
		Global:   offset,
	}
	gs.tester.binding.events = append(gs.tester.binding.events, upEvent)

	return gs.tester.Pump(time.Millisecond * 50)
}

// Drag simulates a drag gesture
func (gs *GestureSimulator) Drag(startX, startY, endX, endY float64, config *GestureConfig) error {
	if config == nil {
		config = DefaultGestureConfig()
	}

	return gs.tester.Drag(startX, startY, endX, endY, config.Steps)
}

// Fling simulates a fling gesture with velocity
func (gs *GestureSimulator) Fling(startX, startY float64, velocityX, velocityY float64, config *GestureConfig) error {
	if config == nil {
		config = DefaultGestureConfig()
	}

	// Calculate end position based on velocity and duration
	durationSeconds := config.Duration.Seconds()
	endX := startX + velocityX*durationSeconds
	endY := startY + velocityY*durationSeconds

	return gs.Drag(startX, startY, endX, endY, config)
}

// Pinch simulates a pinch gesture (two-finger zoom)
func (gs *GestureSimulator) Pinch(centerX, centerY, startDistance, endDistance float64, config *GestureConfig) error {
	if config == nil {
		config = DefaultGestureConfig()
	}

	// Calculate finger positions
	startFinger1 := &goflow.Offset{X: centerX - startDistance/2, Y: centerY}
	startFinger2 := &goflow.Offset{X: centerX + startDistance/2, Y: centerY}
	endFinger1 := &goflow.Offset{X: centerX - endDistance/2, Y: centerY}
	endFinger2 := &goflow.Offset{X: centerX + endDistance/2, Y: centerY}

	// Send pointer down for both fingers
	down1 := &input.MultiTouchEvent{
		PositionEvent: input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypePointerDown,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeTouch,
			},
			Position: startFinger1,
			Global:   startFinger1,
		},
		PointerID: 0,
		TouchID:   0,
	}

	down2 := &input.MultiTouchEvent{
		PositionEvent: input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypePointerDown,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeTouch,
			},
			Position: startFinger2,
			Global:   startFinger2,
		},
		PointerID: 1,
		TouchID:   1,
	}

	gs.tester.binding.events = append(gs.tester.binding.events, down1, down2)

	if err := gs.tester.Pump(time.Millisecond * 50); err != nil {
		return err
	}

	// Animate movement
	dx1 := (endFinger1.X - startFinger1.X) / float64(config.Steps)
	dy1 := (endFinger1.Y - startFinger1.Y) / float64(config.Steps)
	dx2 := (endFinger2.X - startFinger2.X) / float64(config.Steps)
	dy2 := (endFinger2.Y - startFinger2.Y) / float64(config.Steps)

	for i := 1; i <= config.Steps; i++ {
		pos1 := &goflow.Offset{
			X: startFinger1.X + dx1*float64(i),
			Y: startFinger1.Y + dy1*float64(i),
		}
		pos2 := &goflow.Offset{
			X: startFinger2.X + dx2*float64(i),
			Y: startFinger2.Y + dy2*float64(i),
		}

		move1 := &input.MultiTouchEvent{
			PositionEvent: input.PositionEvent{
				BaseEvent: input.BaseEvent{
					EventType:  input.EventTypePointerMove,
					Time:       time.Now(),
					DeviceType: input.DeviceTypeTouch,
				},
				Position: pos1,
				Global:   pos1,
			},
			PointerID: 0,
			TouchID:   0,
		}

		move2 := &input.MultiTouchEvent{
			PositionEvent: input.PositionEvent{
				BaseEvent: input.BaseEvent{
					EventType:  input.EventTypePointerMove,
					Time:       time.Now(),
					DeviceType: input.DeviceTypeTouch,
				},
				Position: pos2,
				Global:   pos2,
			},
			PointerID: 1,
			TouchID:   1,
		}

		gs.tester.binding.events = append(gs.tester.binding.events, move1, move2)

		if err := gs.tester.Pump(config.Duration / time.Duration(config.Steps)); err != nil {
			return err
		}
	}

	// Send pointer up for both fingers
	up1 := &input.MultiTouchEvent{
		PositionEvent: input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypePointerUp,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeTouch,
			},
			Position: endFinger1,
			Global:   endFinger1,
		},
		PointerID: 0,
		TouchID:   0,
	}

	up2 := &input.MultiTouchEvent{
		PositionEvent: input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypePointerUp,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeTouch,
			},
			Position: endFinger2,
			Global:   endFinger2,
		},
		PointerID: 1,
		TouchID:   1,
	}

	gs.tester.binding.events = append(gs.tester.binding.events, up1, up2)

	return gs.tester.Pump(time.Millisecond * 50)
}

// Swipe simulates a swipe gesture in a direction
func (gs *GestureSimulator) Swipe(startX, startY float64, direction SwipeDirection, distance float64, config *GestureConfig) error {
	if config == nil {
		config = DefaultGestureConfig()
	}

	var endX, endY float64
	switch direction {
	case SwipeUp:
		endX = startX
		endY = startY - distance
	case SwipeDown:
		endX = startX
		endY = startY + distance
	case SwipeLeft:
		endX = startX - distance
		endY = startY
	case SwipeRight:
		endX = startX + distance
		endY = startY
	default:
		return fmt.Errorf("invalid swipe direction: %d", direction)
	}

	return gs.Drag(startX, startY, endX, endY, config)
}

// SwipeDirection represents the direction of a swipe
type SwipeDirection int

const (
	SwipeUp SwipeDirection = iota
	SwipeDown
	SwipeLeft
	SwipeRight
)

// Rotate simulates a rotation gesture (two-finger rotation)
func (gs *GestureSimulator) Rotate(centerX, centerY, radius, startAngle, endAngle float64, config *GestureConfig) error {
	if config == nil {
		config = DefaultGestureConfig()
	}

	// Not fully implemented - would require complex multi-touch simulation
	// This is a placeholder for the API
	return fmt.Errorf("rotate gesture not yet fully implemented")
}

// Scroll simulates a scroll gesture
func (gs *GestureSimulator) Scroll(x, y, deltaX, deltaY float64) error {
	// Create scroll/wheel event
	scrollEvent := &input.ScrollEvent{
		PositionEvent: input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypeMouseWheel,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeMouse,
			},
			Position: &goflow.Offset{X: x, Y: y},
			Global:   &goflow.Offset{X: x, Y: y},
		},
		ScrollDelta: &goflow.Offset{X: deltaX, Y: deltaY},
	}

	gs.tester.binding.events = append(gs.tester.binding.events, scrollEvent)

	return gs.tester.Pump(time.Millisecond * 50)
}

// Hover simulates a hover gesture (mouse movement without clicking)
func (gs *GestureSimulator) Hover(x, y float64) error {
	hoverEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypeMouseMove,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeMouse,
		},
		Position: &goflow.Offset{X: x, Y: y},
		Global:   &goflow.Offset{X: x, Y: y},
	}

	gs.tester.binding.events = append(gs.tester.binding.events, hoverEvent)

	return gs.tester.Pump(time.Millisecond * 50)
}
