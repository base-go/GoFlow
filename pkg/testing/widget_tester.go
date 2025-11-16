package testing

import (
	"fmt"
	"image"
	"testing"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/input"
)

// WidgetTester provides utilities for testing widgets
type WidgetTester struct {
	t              *testing.T
	element        goflow.Element
	binding        *TestBinding
	pumpDuration   time.Duration
	autoSizeCanvas bool
}

// TestBinding provides a test-specific platform binding
type TestBinding struct {
	canvas      *MockCanvas
	size        *goflow.Size
	events      []input.InputEvent
	frameCount  int
	didSchedule bool
}

// MockCanvas is a canvas implementation for testing
type MockCanvas struct {
	size        *goflow.Size
	pixels      []byte
	drawCalls   []DrawCall
	image       *image.RGBA
	backgroundColor *goflow.Color
}

// DrawCall records a drawing operation
type DrawCall struct {
	Type       string
	Parameters map[string]interface{}
}

// NewWidgetTester creates a new widget tester
func NewWidgetTester(t *testing.T) *WidgetTester {
	return &WidgetTester{
		t:              t,
		binding:        NewTestBinding(),
		pumpDuration:   time.Millisecond * 16, // ~60fps
		autoSizeCanvas: true,
	}
}

// NewTestBinding creates a new test binding
func NewTestBinding() *TestBinding {
	return &TestBinding{
		canvas: NewMockCanvas(800, 600),
		size:   &goflow.Size{Width: 800, Height: 600},
		events: make([]input.InputEvent, 0),
	}
}

// NewMockCanvas creates a new mock canvas
func NewMockCanvas(width, height float64) *MockCanvas {
	w := int(width)
	h := int(height)
	return &MockCanvas{
		size:      &goflow.Size{Width: width, Height: height},
		pixels:    make([]byte, w*h*4), // RGBA
		drawCalls: make([]DrawCall, 0),
		image:     image.NewRGBA(image.Rect(0, 0, w, h)),
	}
}

// PumpWidget builds and renders a widget
func (wt *WidgetTester) PumpWidget(widget goflow.Widget) error {
	// Create element from widget
	wt.element = widget.CreateElement()

	// Mount the element
	wt.element.Mount(nil, nil)

	// Rebuild
	wt.element.Rebuild()

	// Pump frames to allow animations and effects to settle
	return wt.Pump(wt.pumpDuration)
}

// Pump advances time and triggers frame callbacks
func (wt *WidgetTester) Pump(duration time.Duration) error {
	// Simulate frame scheduling
	wt.binding.didSchedule = true
	wt.binding.frameCount++

	// Allow time to pass
	time.Sleep(duration)

	// Trigger rebuild if needed
	if wt.element != nil {
		wt.element.Rebuild()
	}

	return nil
}

// PumpAndSettle pumps frames until there are no more scheduled frames
func (wt *WidgetTester) PumpAndSettle(timeout time.Duration) error {
	start := time.Now()
	maxIterations := 100
	iterations := 0

	for {
		if time.Since(start) > timeout {
			return fmt.Errorf("pumpAndSettle timed out after %v", timeout)
		}

		if iterations >= maxIterations {
			return fmt.Errorf("pumpAndSettle exceeded max iterations (%d)", maxIterations)
		}

		wt.binding.didSchedule = false
		if err := wt.Pump(wt.pumpDuration); err != nil {
			return err
		}

		// If nothing was scheduled, we're settled
		if !wt.binding.didSchedule {
			break
		}

		iterations++
	}

	return nil
}

// GetElement returns the root element
func (wt *WidgetTester) GetElement() goflow.Element {
	return wt.element
}

// GetCanvas returns the mock canvas
func (wt *WidgetTester) GetCanvas() *MockCanvas {
	return wt.binding.canvas
}

// GetSize returns the current size
func (wt *WidgetTester) GetSize() *goflow.Size {
	return wt.binding.size
}

// SetSize sets the canvas size
func (wt *WidgetTester) SetSize(width, height float64) {
	wt.binding.size = &goflow.Size{Width: width, Height: height}
	wt.binding.canvas = NewMockCanvas(width, height)
}

// Tap simulates a tap at the given position
func (wt *WidgetTester) Tap(x, y float64) error {
	return wt.TapAt(&goflow.Offset{X: x, Y: y})
}

// TapAt simulates a tap at the given offset
func (wt *WidgetTester) TapAt(offset *goflow.Offset) error {
	// Send pointer down event
	downEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerDown,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: offset,
		Global:   offset,
	}
	wt.binding.events = append(wt.binding.events, downEvent)

	// Pump to process event
	if err := wt.Pump(time.Millisecond * 50); err != nil {
		return err
	}

	// Send pointer up event
	upEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerUp,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: offset,
		Global:   offset,
	}
	wt.binding.events = append(wt.binding.events, upEvent)

	// Pump to process event
	return wt.Pump(time.Millisecond * 50)
}

// Drag simulates a drag gesture
func (wt *WidgetTester) Drag(startX, startY, endX, endY float64, steps int) error {
	start := &goflow.Offset{X: startX, Y: startY}
	end := &goflow.Offset{X: endX, Y: endY}

	// Send pointer down
	downEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerDown,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: start,
		Global:   start,
	}
	wt.binding.events = append(wt.binding.events, downEvent)

	if err := wt.Pump(time.Millisecond * 50); err != nil {
		return err
	}

	// Send move events
	dx := (end.X - start.X) / float64(steps)
	dy := (end.Y - start.Y) / float64(steps)

	for i := 1; i <= steps; i++ {
		pos := &goflow.Offset{
			X: start.X + dx*float64(i),
			Y: start.Y + dy*float64(i),
		}

		moveEvent := &input.PositionEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypePointerMove,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeTouch,
			},
			Position: pos,
			Global:   pos,
		}
		wt.binding.events = append(wt.binding.events, moveEvent)

		if err := wt.Pump(time.Millisecond * 16); err != nil {
			return err
		}
	}

	// Send pointer up
	upEvent := &input.PositionEvent{
		BaseEvent: input.BaseEvent{
			EventType:  input.EventTypePointerUp,
			Time:       time.Now(),
			DeviceType: input.DeviceTypeTouch,
		},
		Position: end,
		Global:   end,
	}
	wt.binding.events = append(wt.binding.events, upEvent)

	return wt.Pump(time.Millisecond * 50)
}

// EnterText simulates text entry
func (wt *WidgetTester) EnterText(text string) error {
	for _, ch := range text {
		keyEvent := &input.KeyEvent{
			BaseEvent: input.BaseEvent{
				EventType:  input.EventTypeKeyDown,
				Time:       time.Now(),
				DeviceType: input.DeviceTypeKeyboard,
			},
			Key:       input.Key(ch),
			Character: string(ch),
		}
		wt.binding.events = append(wt.binding.events, keyEvent)

		if err := wt.Pump(time.Millisecond * 16); err != nil {
			return err
		}
	}
	return nil
}

// GetEvents returns all recorded events
func (wt *WidgetTester) GetEvents() []input.InputEvent {
	return wt.binding.events
}

// ClearEvents clears all recorded events
func (wt *WidgetTester) ClearEvents() {
	wt.binding.events = make([]input.InputEvent, 0)
}

// VerifyElement verifies that an element exists in the tree
func (wt *WidgetTester) VerifyElement(predicate func(goflow.Element) bool) bool {
	if wt.element == nil {
		return false
	}
	return wt.findElement(wt.element, predicate) != nil
}

// FindElement finds an element matching the predicate
func (wt *WidgetTester) FindElement(predicate func(goflow.Element) bool) goflow.Element {
	if wt.element == nil {
		return nil
	}
	return wt.findElement(wt.element, predicate)
}

func (wt *WidgetTester) findElement(element goflow.Element, predicate func(goflow.Element) bool) goflow.Element {
	if predicate(element) {
		return element
	}

	var found goflow.Element
	element.VisitChildren(func(child goflow.Element) {
		if found == nil {
			found = wt.findElement(child, predicate)
		}
	})

	return found
}

// GetDrawCalls returns all recorded draw calls
func (mc *MockCanvas) GetDrawCalls() []DrawCall {
	return mc.drawCalls
}

// ClearDrawCalls clears all recorded draw calls
func (mc *MockCanvas) ClearDrawCalls() {
	mc.drawCalls = make([]DrawCall, 0)
}

// GetImage returns the rendered image
func (mc *MockCanvas) GetImage() *image.RGBA {
	return mc.image
}

// RecordDrawCall records a drawing operation
func (mc *MockCanvas) RecordDrawCall(callType string, params map[string]interface{}) {
	mc.drawCalls = append(mc.drawCalls, DrawCall{
		Type:       callType,
		Parameters: params,
	})
}
