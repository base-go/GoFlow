package signals

import "sync"

// Effect represents a side effect that runs when dependencies change
type Effect struct {
	mu           sync.Mutex
	fn           func()
	dependencies map[interface{}]struct{}
	disposed     bool
}

// NewEffect creates a new effect and runs it immediately
// Returns a dispose function to stop the effect
func NewEffect(fn func()) func() {
	e := &Effect{
		fn:           fn,
		dependencies: make(map[interface{}]struct{}),
	}

	// Run the effect for the first time
	e.run()

	// Return dispose function
	return func() {
		e.dispose()
	}
}

// run executes the effect function and tracks dependencies
func (e *Effect) run() {
	e.mu.Lock()
	if e.disposed {
		e.mu.Unlock()
		return
	}

	// Unsubscribe from old dependencies
	for dep := range e.dependencies {
		if sig, ok := dep.(interface{ unsubscribe(Subscriber) }); ok {
			sig.unsubscribe(e)
		}
	}
	e.dependencies = make(map[interface{}]struct{})

	// Set ourselves as the active observer
	prev := startTracking(e)
	e.mu.Unlock()

	// Run the effect (this will track dependencies)
	e.fn()

	stopTracking(prev)
}

// notify is called when a dependency changes
func (e *Effect) notify() {
	e.run()
}

// dispose stops the effect and unsubscribes from all dependencies
func (e *Effect) dispose() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.disposed {
		return
	}

	e.disposed = true

	// Unsubscribe from all dependencies
	for dep := range e.dependencies {
		if sig, ok := dep.(interface{ unsubscribe(Subscriber) }); ok {
			sig.unsubscribe(e)
		}
	}
	e.dependencies = make(map[interface{}]struct{})
}
