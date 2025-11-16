package signals

import "sync"

// Subscriber represents anything that can receive notifications
type Subscriber interface {
	notify()
}

// context manages the reactive dependency tracking
type context struct {
	mu           sync.RWMutex
	observer     Subscriber // Currently executing observer (effect or computed)
	batching     int        // Batch nesting level
	pendingNotify map[Subscriber]struct{} // Notifications pending during batch
}

var globalContext = &context{
	pendingNotify: make(map[Subscriber]struct{}),
}

// startTracking sets the current observer for dependency tracking
func startTracking(observer Subscriber) Subscriber {
	globalContext.mu.Lock()
	defer globalContext.mu.Unlock()
	prev := globalContext.observer
	globalContext.observer = observer
	return prev
}

// stopTracking restores the previous observer
func stopTracking(prev Subscriber) {
	globalContext.mu.Lock()
	defer globalContext.mu.Unlock()
	globalContext.observer = prev
}

// getCurrentObserver returns the current observer, if any
func getCurrentObserver() Subscriber {
	globalContext.mu.RLock()
	defer globalContext.mu.RUnlock()
	return globalContext.observer
}

// notifySubscriber notifies a subscriber, respecting batching
func notifySubscriber(sub Subscriber) {
	globalContext.mu.Lock()
	defer globalContext.mu.Unlock()

	if globalContext.batching > 0 {
		// Queue notification for after batch completes
		globalContext.pendingNotify[sub] = struct{}{}
	} else {
		// Notify immediately
		globalContext.mu.Unlock()
		sub.notify()
		globalContext.mu.Lock()
	}
}

// startBatch begins a batch operation
func startBatch() {
	globalContext.mu.Lock()
	defer globalContext.mu.Unlock()
	globalContext.batching++
}

// endBatch completes a batch and flushes pending notifications
func endBatch() {
	globalContext.mu.Lock()
	globalContext.batching--

	if globalContext.batching == 0 {
		// Flush all pending notifications
		pending := globalContext.pendingNotify
		globalContext.pendingNotify = make(map[Subscriber]struct{})
		globalContext.mu.Unlock()

		// Notify outside of lock
		for sub := range pending {
			sub.notify()
		}
	} else {
		globalContext.mu.Unlock()
	}
}
