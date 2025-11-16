package signals

import (
	"testing"
	"time"
)

func TestSignalBasic(t *testing.T) {
	s := New(42)

	if got := s.Get(); got != 42 {
		t.Errorf("Get() = %v, want %v", got, 42)
	}

	s.Set(100)

	if got := s.Get(); got != 100 {
		t.Errorf("After Set(100), Get() = %v, want %v", got, 100)
	}
}

func TestSignalUpdate(t *testing.T) {
	s := New(10)

	s.Update(func(v int) int {
		return v * 2
	})

	if got := s.Get(); got != 20 {
		t.Errorf("After Update, Get() = %v, want %v", got, 20)
	}
}

func TestSignalPeek(t *testing.T) {
	s := New(42)
	callCount := 0

	// Create an effect that shouldn't be triggered by Peek
	dispose := NewEffect(func() {
		s.Peek() // Using Peek instead of Get
		callCount++
	})
	defer dispose()

	initialCount := callCount

	// Set should not trigger the effect since we used Peek
	s.Set(100)

	// Give some time for any potential notification
	time.Sleep(10 * time.Millisecond)

	if callCount != initialCount {
		t.Errorf("Peek should not track dependencies, but effect was called %d times", callCount-initialCount)
	}
}

func TestSignalEffect(t *testing.T) {
	s := New(1)
	callCount := 0
	var lastValue int

	dispose := NewEffect(func() {
		lastValue = s.Get()
		callCount++
	})
	defer dispose()

	if callCount != 1 {
		t.Errorf("Effect should run immediately, callCount = %v, want 1", callCount)
	}

	s.Set(2)
	time.Sleep(10 * time.Millisecond)

	if callCount != 2 {
		t.Errorf("Effect should run after Set, callCount = %v, want 2", callCount)
	}

	if lastValue != 2 {
		t.Errorf("lastValue = %v, want 2", lastValue)
	}
}

func TestComputed(t *testing.T) {
	a := New(2)
	b := New(3)

	sum := NewComputed(func() int {
		return a.Get() + b.Get()
	})

	if got := sum.Get(); got != 5 {
		t.Errorf("sum.Get() = %v, want 5", got)
	}

	a.Set(10)

	if got := sum.Get(); got != 13 {
		t.Errorf("After a.Set(10), sum.Get() = %v, want 13", got)
	}

	b.Set(20)

	if got := sum.Get(); got != 30 {
		t.Errorf("After b.Set(20), sum.Get() = %v, want 30", got)
	}
}

func TestComputedLazy(t *testing.T) {
	s := New(1)
	computeCount := 0

	computed := NewComputed(func() int {
		computeCount++
		return s.Get() * 2
	})

	// Computed shouldn't run until accessed
	if computeCount != 0 {
		t.Errorf("Computed should be lazy, but ran %d times before Get()", computeCount)
	}

	computed.Get()
	if computeCount != 1 {
		t.Errorf("After first Get(), computeCount = %v, want 1", computeCount)
	}

	// Getting again without changes shouldn't recompute
	computed.Get()
	if computeCount != 1 {
		t.Errorf("After second Get() without changes, computeCount = %v, want 1", computeCount)
	}

	// Changing dependency should mark as dirty, but not recompute yet
	s.Set(2)
	if computeCount != 1 {
		t.Errorf("After Set() but before Get(), computeCount = %v, want 1", computeCount)
	}

	// Now it should recompute
	computed.Get()
	if computeCount != 2 {
		t.Errorf("After Get() following dependency change, computeCount = %v, want 2", computeCount)
	}
}

func TestComputedWithEffect(t *testing.T) {
	s := New(5)
	doubled := NewComputed(func() int {
		return s.Get() * 2
	})

	var lastValue int
	callCount := 0

	dispose := NewEffect(func() {
		lastValue = doubled.Get()
		callCount++
	})
	defer dispose()

	if callCount != 1 || lastValue != 10 {
		t.Errorf("Initial: callCount=%v lastValue=%v, want callCount=1 lastValue=10", callCount, lastValue)
	}

	s.Set(10)
	time.Sleep(10 * time.Millisecond)

	if callCount != 2 || lastValue != 20 {
		t.Errorf("After Set(10): callCount=%v lastValue=%v, want callCount=2 lastValue=20", callCount, lastValue)
	}
}

func TestBatch(t *testing.T) {
	a := New(1)
	b := New(2)
	callCount := 0

	sum := NewComputed(func() int {
		return a.Get() + b.Get()
	})

	dispose := NewEffect(func() {
		sum.Get()
		callCount++
	})
	defer dispose()

	initialCount := callCount

	// Without batch, this would trigger effect twice
	Batch(func() {
		a.Set(10)
		b.Set(20)
	})

	time.Sleep(10 * time.Millisecond)

	// Should only trigger once due to batching
	if callCount != initialCount+1 {
		t.Errorf("Batch should combine updates, callCount = %v, want %v", callCount, initialCount+1)
	}

	if got := sum.Get(); got != 30 {
		t.Errorf("sum.Get() = %v, want 30", got)
	}
}

func TestNestedBatch(t *testing.T) {
	s := New(1)
	callCount := 0

	dispose := NewEffect(func() {
		s.Get()
		callCount++
	})
	defer dispose()

	initialCount := callCount

	Batch(func() {
		s.Set(2)
		Batch(func() {
			s.Set(3)
		})
		s.Set(4)
	})

	time.Sleep(10 * time.Millisecond)

	// All updates should be batched into one notification
	if callCount != initialCount+1 {
		t.Errorf("Nested batch should combine all updates, callCount = %v, want %v", callCount, initialCount+1)
	}

	if got := s.Get(); got != 4 {
		t.Errorf("s.Get() = %v, want 4", got)
	}
}

func TestUntracked(t *testing.T) {
	s := New(1)
	callCount := 0

	dispose := NewEffect(func() {
		// Read without tracking
		Untracked(func() int {
			return s.Get()
		})
		callCount++
	})
	defer dispose()

	initialCount := callCount

	// This shouldn't trigger the effect
	s.Set(2)
	time.Sleep(10 * time.Millisecond)

	if callCount != initialCount {
		t.Errorf("Untracked should prevent dependency tracking, callCount = %v, want %v", callCount, initialCount)
	}
}

func TestEffectDispose(t *testing.T) {
	s := New(1)
	callCount := 0

	dispose := NewEffect(func() {
		s.Get()
		callCount++
	})

	initialCount := callCount

	s.Set(2)
	time.Sleep(10 * time.Millisecond)

	if callCount != initialCount+1 {
		t.Errorf("Before dispose, callCount = %v, want %v", callCount, initialCount+1)
	}

	dispose()

	beforeDisposeCount := callCount
	s.Set(3)
	time.Sleep(10 * time.Millisecond)

	if callCount != beforeDisposeCount {
		t.Errorf("After dispose, effect should not run, callCount = %v, want %v", callCount, beforeDisposeCount)
	}
}

func TestComplexDependencyGraph(t *testing.T) {
	// a -> b -> d
	//  \-> c -/
	a := New(1)
	b := NewComputed(func() int { return a.Get() * 2 })
	c := NewComputed(func() int { return a.Get() + 10 })
	d := NewComputed(func() int { return b.Get() + c.Get() })

	if got := d.Get(); got != 13 {
		t.Errorf("Initial d.Get() = %v, want 13", got)
	}

	a.Set(5)

	// b = 10, c = 15, d = 25
	if got := d.Get(); got != 25 {
		t.Errorf("After a.Set(5), d.Get() = %v, want 25", got)
	}
}
