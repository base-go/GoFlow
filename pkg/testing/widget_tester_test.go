package testing

import (
	"testing"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// TestWidgetTesterCreation tests that WidgetTester can be created
func TestWidgetTesterCreation(t *testing.T) {
	wt := NewWidgetTester(t)

	if wt == nil {
		t.Fatal("WidgetTester should not be nil")
	}

	if wt.t != t {
		t.Error("WidgetTester should store testing.T reference")
	}

	if wt.binding == nil {
		t.Error("WidgetTester should have a binding")
	}

	if wt.pumpDuration == 0 {
		t.Error("WidgetTester should have a default pump duration")
	}
}

// TestMockCanvas tests MockCanvas functionality
func TestMockCanvas(t *testing.T) {
	canvas := NewMockCanvas(800, 600)

	if canvas == nil {
		t.Fatal("MockCanvas should not be nil")
	}

	if canvas.size.Width != 800 || canvas.size.Height != 600 {
		t.Errorf("Canvas size incorrect: got %.0fx%.0f, want 800x600",
			canvas.size.Width, canvas.size.Height)
	}

	if canvas.image == nil {
		t.Error("Canvas image should not be nil")
	}

	// Test recording draw calls
	canvas.RecordDrawCall("fillRect", map[string]interface{}{
		"x":      10.0,
		"y":      20.0,
		"width":  100.0,
		"height": 50.0,
	})

	calls := canvas.GetDrawCalls()
	if len(calls) != 1 {
		t.Errorf("Expected 1 draw call, got %d", len(calls))
	}

	if calls[0].Type != "fillRect" {
		t.Errorf("Expected draw call type 'fillRect', got '%s'", calls[0].Type)
	}

	// Test clearing draw calls
	canvas.ClearDrawCalls()
	if len(canvas.GetDrawCalls()) != 0 {
		t.Error("Draw calls should be cleared")
	}
}

// TestPump tests the Pump function
func TestPump(t *testing.T) {
	wt := NewWidgetTester(t)

	start := time.Now()
	duration := time.Millisecond * 100

	if err := wt.Pump(duration); err != nil {
		t.Fatalf("Pump should not error: %v", err)
	}

	elapsed := time.Since(start)
	if elapsed < duration {
		t.Errorf("Pump should wait at least %v, but only waited %v", duration, elapsed)
	}
}

// TestSetSize tests changing canvas size
func TestSetSize(t *testing.T) {
	wt := NewWidgetTester(t)

	wt.SetSize(1920, 1080)

	size := wt.GetSize()
	if size.Width != 1920 || size.Height != 1080 {
		t.Errorf("Size not updated correctly: got %.0fx%.0f, want 1920x1080",
			size.Width, size.Height)
	}

	canvas := wt.GetCanvas()
	if canvas.size.Width != 1920 || canvas.size.Height != 1080 {
		t.Errorf("Canvas size not updated correctly: got %.0fx%.0f, want 1920x1080",
			canvas.size.Width, canvas.size.Height)
	}
}

// TestEventRecording tests event recording
func TestEventRecording(t *testing.T) {
	wt := NewWidgetTester(t)

	// Initially should have no events
	if len(wt.GetEvents()) != 0 {
		t.Error("Should start with no events")
	}

	// Simulate a tap (which creates events)
	// Note: This will fail if no element is mounted, but tests the mechanism
	wt.Tap(100, 100)

	events := wt.GetEvents()
	if len(events) == 0 {
		t.Error("Tap should create events")
	}

	// Clear events
	wt.ClearEvents()
	if len(wt.GetEvents()) != 0 {
		t.Error("Events should be cleared")
	}
}

// TestFindElement tests element finding (without actual elements)
func TestFindElement(t *testing.T) {
	wt := NewWidgetTester(t)

	// Should return nil when no element is set
	found := wt.FindElement(func(e goflow.Element) bool {
		return true
	})

	if found != nil {
		t.Error("Should not find element when none is mounted")
	}

	// VerifyElement should return false
	if wt.VerifyElement(func(e goflow.Element) bool {
		return true
	}) {
		t.Error("VerifyElement should return false when no element is mounted")
	}
}
