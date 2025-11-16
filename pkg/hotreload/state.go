package hotreload

import (
	"encoding/json"
	"fmt"
	"sync"
)

// StateStore preserves state across hot reloads
type StateStore struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

// NewStateStore creates a new state store
func NewStateStore() *StateStore {
	return &StateStore{
		store: make(map[string]interface{}),
	}
}

// Set stores a value with the given key
func (s *StateStore) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

// Get retrieves a value by key
func (s *StateStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.store[key]
	return val, ok
}

// GetOrDefault retrieves a value by key, or returns the default if not found
func (s *StateStore) GetOrDefault(key string, defaultValue interface{}) interface{} {
	if val, ok := s.Get(key); ok {
		return val
	}
	return defaultValue
}

// Delete removes a value by key
func (s *StateStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

// Clear removes all stored values
func (s *StateStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store = make(map[string]interface{})
}

// Keys returns all stored keys
func (s *StateStore) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.store))
	for k := range s.store {
		keys = append(keys, k)
	}
	return keys
}

// Snapshot returns a JSON snapshot of the current state
func (s *StateStore) Snapshot() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return json.Marshal(s.store)
}

// Restore restores state from a JSON snapshot
func (s *StateStore) Restore(snapshot []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store := make(map[string]interface{})
	if err := json.Unmarshal(snapshot, &store); err != nil {
		return fmt.Errorf("failed to restore state: %w", err)
	}

	s.store = store
	return nil
}

// Size returns the number of stored values
func (s *StateStore) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.store)
}
