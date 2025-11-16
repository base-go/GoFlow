package signals

import "sync"

// Signal is a reactive container for a value of type T
type Signal[T any] struct {
	mu          sync.RWMutex
	value       T
	subscribers map[Subscriber]struct{}
}

// New creates a new Signal with the given initial value
func New[T any](initial T) *Signal[T] {
	return &Signal[T]{
		value:       initial,
		subscribers: make(map[Subscriber]struct{}),
	}
}

// Get returns the current value and tracks the dependency
func (s *Signal[T]) Get() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Track dependency if there's an active observer
	if observer := getCurrentObserver(); observer != nil {
		s.subscribe(observer)
	}

	return s.value
}

// Peek returns the current value without tracking dependencies
func (s *Signal[T]) Peek() T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

// Set updates the value and notifies subscribers
func (s *Signal[T]) Set(value T) {
	s.mu.Lock()
	s.value = value
	subscribers := make([]Subscriber, 0, len(s.subscribers))
	for sub := range s.subscribers {
		subscribers = append(subscribers, sub)
	}
	s.mu.Unlock()

	// Notify all subscribers
	for _, sub := range subscribers {
		notifySubscriber(sub)
	}
}

// Update applies a function to update the value
func (s *Signal[T]) Update(fn func(T) T) {
	s.mu.Lock()
	s.value = fn(s.value)
	subscribers := make([]Subscriber, 0, len(s.subscribers))
	for sub := range s.subscribers {
		subscribers = append(subscribers, sub)
	}
	s.mu.Unlock()

	// Notify all subscribers
	for _, sub := range subscribers {
		notifySubscriber(sub)
	}
}

// subscribe adds a subscriber (called while holding read lock on signal)
func (s *Signal[T]) subscribe(sub Subscriber) {
	s.mu.RUnlock()
	s.mu.Lock()
	s.subscribers[sub] = struct{}{}
	s.mu.Unlock()
	s.mu.RLock()
}

// unsubscribe removes a subscriber
func (s *Signal[T]) unsubscribe(sub Subscriber) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.subscribers, sub)
}
