package signals

import "sync"

// Computed is a derived value that automatically recomputes when dependencies change
type Computed[T any] struct {
	mu           sync.RWMutex
	compute      func() T
	value        T
	dirty        bool
	dependencies map[interface{}]struct{}
	subscribers  map[Subscriber]struct{}
}

// NewComputed creates a new Computed signal
func NewComputed[T any](compute func() T) *Computed[T] {
	c := &Computed[T]{
		compute:      compute,
		dirty:        true,
		dependencies: make(map[interface{}]struct{}),
		subscribers:  make(map[Subscriber]struct{}),
	}
	return c
}

// Get returns the current value, recomputing if necessary
func (c *Computed[T]) Get() T {
	c.mu.RLock()
	if !c.dirty {
		value := c.value
		// Track dependency if there's an active observer
		if observer := getCurrentObserver(); observer != nil {
			c.mu.RUnlock()
			c.subscribe(observer)
			c.mu.RLock()
		}
		c.mu.RUnlock()
		return value
	}
	c.mu.RUnlock()

	// Need to recompute
	c.mu.Lock()
	// Double-check after acquiring write lock
	if !c.dirty {
		value := c.value
		c.mu.Unlock()
		return value
	}

	// Unsubscribe from old dependencies
	for dep := range c.dependencies {
		if sig, ok := dep.(interface{ unsubscribe(Subscriber) }); ok {
			sig.unsubscribe(c)
		}
	}
	c.dependencies = make(map[interface{}]struct{})

	// Set ourselves as the active observer to track new dependencies
	prev := startTracking(c)
	c.mu.Unlock()

	// Compute new value (this will track dependencies via Get() calls)
	newValue := c.compute()

	c.mu.Lock()
	c.value = newValue
	c.dirty = false
	stopTracking(prev)

	// Track dependency if there's an active observer
	if observer := getCurrentObserver(); observer != nil {
		c.subscribers[observer] = struct{}{}
	}

	c.mu.Unlock()
	return newValue
}

// Peek returns the current value without tracking dependencies or recomputing
func (c *Computed[T]) Peek() T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// notify is called when a dependency changes
func (c *Computed[T]) notify() {
	c.mu.Lock()
	wasClean := !c.dirty
	c.dirty = true

	if !wasClean {
		c.mu.Unlock()
		return
	}

	subscribers := make([]Subscriber, 0, len(c.subscribers))
	for sub := range c.subscribers {
		subscribers = append(subscribers, sub)
	}
	c.mu.Unlock()

	// Notify subscribers that we've changed
	for _, sub := range subscribers {
		notifySubscriber(sub)
	}
}

// subscribe adds a subscriber
func (c *Computed[T]) subscribe(sub Subscriber) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.subscribers[sub] = struct{}{}
}

// unsubscribe removes a subscriber
func (c *Computed[T]) unsubscribe(sub Subscriber) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.subscribers, sub)
}
