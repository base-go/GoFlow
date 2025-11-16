// Package signals provides a reactive state management system inspired by
// Preact Signals and the Dart signals package.
//
// The signals package implements fine-grained reactivity with automatic
// dependency tracking and lazy evaluation.
//
// Core Concepts:
//
// Signal - A reactive container for values that change over time. When a
// signal's value changes, all dependent computed values and effects are
// automatically notified.
//
//	counter := signals.New(0)
//	counter.Set(1)              // Update the value
//	value := counter.Get()      // Read and track dependency
//	value := counter.Peek()     // Read without tracking
//
// Computed - A derived value that automatically recomputes when its
// dependencies change. Computed values are lazy - they only recalculate
// when accessed.
//
//	doubled := signals.NewComputed(func() int {
//	    return counter.Get() * 2
//	})
//
// Effect - A side effect that runs when its dependencies change. Effects
// run immediately when created and again whenever any tracked signal changes.
//
//	dispose := signals.NewEffect(func() {
//	    fmt.Println("Count:", counter.Get())
//	})
//	defer dispose() // Clean up when done
//
// Batch - Groups multiple signal updates into a single notification cycle,
// preventing unnecessary recomputations.
//
//	signals.Batch(func() {
//	    firstName.Set("John")
//	    lastName.Set("Doe")
//	})
//
// Untracked - Reads signal values without creating dependencies, useful
// for avoiding infinite loops in effects.
//
//	signals.Untracked(func() int {
//	    return counter.Peek()
//	})
//
// Thread Safety:
//
// All signal operations are thread-safe and can be used concurrently from
// multiple goroutines.
package signals
