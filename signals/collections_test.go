package signals

import (
	"testing"
	"time"
)

func TestSignalSlice(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	if len(s.Get()) != 3 {
		t.Errorf("Initial length = %v, want 3", len(s.Get()))
	}

	s.Append(4, 5)
	if len(s.Get()) != 5 {
		t.Errorf("After Append, length = %v, want 5", len(s.Get()))
	}

	s.Prepend(0)
	slice := s.Get()
	if len(slice) != 6 || slice[0] != 0 {
		t.Errorf("After Prepend, slice = %v, want [0 1 2 3 4 5]", slice)
	}
}

func TestSignalSliceRemoveAt(t *testing.T) {
	s := NewSlice([]string{"a", "b", "c", "d"})

	s.RemoveAt(2) // Remove "c"
	slice := s.Get()
	if len(slice) != 3 || slice[2] != "d" {
		t.Errorf("After RemoveAt(2), slice = %v, want [a b d]", slice)
	}

	// Invalid index should not panic
	s.RemoveAt(10)
	if len(s.Get()) != 3 {
		t.Errorf("After invalid RemoveAt, length should still be 3")
	}
}

func TestSignalSliceFilter(t *testing.T) {
	s := NewSlice([]int{1, 2, 3, 4, 5, 6})

	even := s.Filter(func(n int) bool {
		return n%2 == 0
	})

	filtered := even.Get()
	if len(filtered) != 3 {
		t.Errorf("Filtered length = %v, want 3", len(filtered))
	}

	// Should react to changes
	s.Append(8)
	filtered = even.Get()
	if len(filtered) != 4 {
		t.Errorf("After append, filtered length = %v, want 4", len(filtered))
	}
}

func TestSignalSliceMap(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	doubled := s.Map(func(n int) int {
		return n * 2
	})

	mapped := doubled.Get()
	if len(mapped) != 3 || mapped[0] != 2 || mapped[2] != 6 {
		t.Errorf("Mapped = %v, want [2 4 6]", mapped)
	}
}

func TestSignalSliceLen(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})
	length := s.Len()

	if length.Get() != 3 {
		t.Errorf("Len() = %v, want 3", length.Get())
	}

	callCount := 0
	dispose := NewEffect(func() {
		length.Get()
		callCount++
	})
	defer dispose()

	initialCount := callCount
	s.Append(4)
	time.Sleep(10 * time.Millisecond)

	if callCount != initialCount+1 {
		t.Errorf("Effect should trigger on length change, callCount = %v", callCount)
	}
}

func TestSignalSliceClear(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})
	s.Clear()

	if len(s.Get()) != 0 {
		t.Errorf("After Clear, length = %v, want 0", len(s.Get()))
	}
}

func TestSignalMap(t *testing.T) {
	m := NewMap(map[string]int{
		"a": 1,
		"b": 2,
	})

	if len(m.Get()) != 2 {
		t.Errorf("Initial size = %v, want 2", len(m.Get()))
	}

	m.SetKey("c", 3)
	if len(m.Get()) != 3 {
		t.Errorf("After SetKey, size = %v, want 3", len(m.Get()))
	}

	m.DeleteKey("b")
	mp := m.Get()
	if len(mp) != 2 {
		t.Errorf("After DeleteKey, size = %v, want 2", len(mp))
	}
	if _, exists := mp["b"]; exists {
		t.Error("Key 'b' should be deleted")
	}
}

func TestSignalMapHas(t *testing.T) {
	m := NewMap(map[string]int{"a": 1})

	hasA := m.Has("a")
	hasB := m.Has("b")

	if !hasA.Get() {
		t.Error("Has('a') should be true")
	}
	if hasB.Get() {
		t.Error("Has('b') should be false")
	}

	m.SetKey("b", 2)
	if !hasB.Get() {
		t.Error("After SetKey('b'), Has('b') should be true")
	}
}

func TestSignalMapGetKey(t *testing.T) {
	m := NewMap(map[string]int{"a": 1})

	valueA := m.GetKey("a")

	if valueA.Get() != 1 {
		t.Errorf("GetKey('a') = %v, want 1", valueA.Get())
	}

	m.SetKey("a", 10)
	if valueA.Get() != 10 {
		t.Errorf("After update, GetKey('a') = %v, want 10", valueA.Get())
	}
}

func TestSignalMapKeysValues(t *testing.T) {
	m := NewMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})

	keys := m.Keys()
	values := m.Values()

	if len(keys.Get()) != 3 {
		t.Errorf("Keys length = %v, want 3", len(keys.Get()))
	}
	if len(values.Get()) != 3 {
		t.Errorf("Values length = %v, want 3", len(values.Get()))
	}
}

func TestSignalMapSize(t *testing.T) {
	m := NewMap(map[string]int{"a": 1})
	size := m.Size()

	if size.Get() != 1 {
		t.Errorf("Size() = %v, want 1", size.Get())
	}

	m.SetKey("b", 2)
	if size.Get() != 2 {
		t.Errorf("After SetKey, Size() = %v, want 2", size.Get())
	}
}

func TestSignalMapClear(t *testing.T) {
	m := NewMap(map[string]int{"a": 1, "b": 2})
	m.Clear()

	if len(m.Get()) != 0 {
		t.Errorf("After Clear, size = %v, want 0", len(m.Get()))
	}
}
