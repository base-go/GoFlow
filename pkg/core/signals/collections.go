package signals

// SignalSlice provides reactive slice operations
type SignalSlice[T any] struct {
	*Signal[[]T]
}

// NewSlice creates a new reactive slice
func NewSlice[T any](initial []T) *SignalSlice[T] {
	if initial == nil {
		initial = []T{}
	}
	return &SignalSlice[T]{
		Signal: New(initial),
	}
}

// Len returns the current length of the slice
func (s *SignalSlice[T]) Len() *Computed[int] {
	return NewComputed(func() int {
		return len(s.Get())
	})
}

// Append adds items to the end of the slice
func (s *SignalSlice[T]) Append(items ...T) {
	s.Update(func(slice []T) []T {
		return append(slice, items...)
	})
}

// Prepend adds items to the beginning of the slice
func (s *SignalSlice[T]) Prepend(items ...T) {
	s.Update(func(slice []T) []T {
		newSlice := make([]T, 0, len(slice)+len(items))
		newSlice = append(newSlice, items...)
		newSlice = append(newSlice, slice...)
		return newSlice
	})
}

// RemoveAt removes the item at the given index
func (s *SignalSlice[T]) RemoveAt(index int) {
	s.Update(func(slice []T) []T {
		if index < 0 || index >= len(slice) {
			return slice
		}
		newSlice := make([]T, 0, len(slice)-1)
		newSlice = append(newSlice, slice[:index]...)
		newSlice = append(newSlice, slice[index+1:]...)
		return newSlice
	})
}

// Filter creates a computed signal with filtered items
func (s *SignalSlice[T]) Filter(predicate func(T) bool) *Computed[[]T] {
	return NewComputed(func() []T {
		slice := s.Get()
		filtered := make([]T, 0)
		for _, item := range slice {
			if predicate(item) {
				filtered = append(filtered, item)
			}
		}
		return filtered
	})
}

// Map creates a computed signal with mapped items
func (s *SignalSlice[T]) Map(mapper func(T) T) *Computed[[]T] {
	return NewComputed(func() []T {
		slice := s.Get()
		mapped := make([]T, len(slice))
		for i, item := range slice {
			mapped[i] = mapper(item)
		}
		return mapped
	})
}

// Clear removes all items from the slice
func (s *SignalSlice[T]) Clear() {
	s.Set([]T{})
}

// SignalMap provides reactive map operations
type SignalMap[K comparable, V any] struct {
	*Signal[map[K]V]
}

// NewMap creates a new reactive map
func NewMap[K comparable, V any](initial map[K]V) *SignalMap[K, V] {
	if initial == nil {
		initial = make(map[K]V)
	}
	return &SignalMap[K, V]{
		Signal: New(initial),
	}
}

// SetKey sets a key-value pair in the map
func (m *SignalMap[K, V]) SetKey(key K, value V) {
	m.Update(func(mp map[K]V) map[K]V {
		newMap := make(map[K]V, len(mp)+1)
		for k, v := range mp {
			newMap[k] = v
		}
		newMap[key] = value
		return newMap
	})
}

// DeleteKey removes a key from the map
func (m *SignalMap[K, V]) DeleteKey(key K) {
	m.Update(func(mp map[K]V) map[K]V {
		newMap := make(map[K]V, len(mp))
		for k, v := range mp {
			if k != key {
				newMap[k] = v
			}
		}
		return newMap
	})
}

// Has creates a computed signal that checks if a key exists
func (m *SignalMap[K, V]) Has(key K) *Computed[bool] {
	return NewComputed(func() bool {
		mp := m.Get()
		_, exists := mp[key]
		return exists
	})
}

// GetKey creates a computed signal for a specific key's value
func (m *SignalMap[K, V]) GetKey(key K) *Computed[V] {
	return NewComputed(func() V {
		mp := m.Get()
		return mp[key]
	})
}

// Keys creates a computed signal with all keys
func (m *SignalMap[K, V]) Keys() *Computed[[]K] {
	return NewComputed(func() []K {
		mp := m.Get()
		keys := make([]K, 0, len(mp))
		for k := range mp {
			keys = append(keys, k)
		}
		return keys
	})
}

// Values creates a computed signal with all values
func (m *SignalMap[K, V]) Values() *Computed[[]V] {
	return NewComputed(func() []V {
		mp := m.Get()
		values := make([]V, 0, len(mp))
		for _, v := range mp {
			values = append(values, v)
		}
		return values
	})
}

// Size creates a computed signal with the map size
func (m *SignalMap[K, V]) Size() *Computed[int] {
	return NewComputed(func() int {
		return len(m.Get())
	})
}

// Clear removes all entries from the map
func (m *SignalMap[K, V]) Clear() {
	m.Set(make(map[K]V))
}
