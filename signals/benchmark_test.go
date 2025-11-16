package signals

import (
	"testing"
)

func BenchmarkSignalGet(b *testing.B) {
	s := New(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Get()
	}
}

func BenchmarkSignalSet(b *testing.B) {
	s := New(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set(i)
	}
}

func BenchmarkSignalUpdate(b *testing.B) {
	s := New(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Update(func(v int) int {
			return v + 1
		})
	}
}

func BenchmarkComputedSimple(b *testing.B) {
	s := New(1)
	c := NewComputed(func() int {
		return s.Get() * 2
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Get()
	}
}

func BenchmarkComputedChain(b *testing.B) {
	s := New(1)
	c1 := NewComputed(func() int { return s.Get() * 2 })
	c2 := NewComputed(func() int { return c1.Get() + 10 })
	c3 := NewComputed(func() int { return c2.Get() * 3 })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c3.Get()
	}
}

func BenchmarkEffect(b *testing.B) {
	s := New(0)
	var count int
	dispose := NewEffect(func() {
		count = s.Get()
	})
	defer dispose()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set(i)
	}
	_ = count
}

func BenchmarkBatchUpdates(b *testing.B) {
	s1 := New(0)
	s2 := New(0)
	s3 := New(0)
	var count int
	c := NewComputed(func() int {
		return s1.Get() + s2.Get() + s3.Get()
	})
	dispose := NewEffect(func() {
		count = c.Get()
	})
	defer dispose()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Batch(func() {
			s1.Set(i)
			s2.Set(i)
			s3.Set(i)
		})
	}
	_ = count
}

func BenchmarkSignalSliceAppend(b *testing.B) {
	s := NewSlice([]int{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Append(i)
	}
}

func BenchmarkSignalSliceFilter(b *testing.B) {
	s := NewSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	filtered := s.Filter(func(n int) bool {
		return n%2 == 0
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filtered.Get()
	}
}

func BenchmarkSignalMapSetKey(b *testing.B) {
	m := NewMap(map[string]int{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.SetKey("key", i)
	}
}

func BenchmarkSignalMapGetKey(b *testing.B) {
	m := NewMap(map[string]int{"key": 42})
	value := m.GetKey("key")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = value.Get()
	}
}

func BenchmarkComplexReactiveGraph(b *testing.B) {
	// Simulate a complex reactive graph
	s1 := New(1)
	s2 := New(2)
	s3 := New(3)

	c1 := NewComputed(func() int { return s1.Get() + s2.Get() })
	c2 := NewComputed(func() int { return s2.Get() + s3.Get() })
	c3 := NewComputed(func() int { return c1.Get() + c2.Get() })
	c4 := NewComputed(func() int { return c3.Get() * 2 })

	var result int
	dispose := NewEffect(func() {
		result = c4.Get()
	})
	defer dispose()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Set(i)
	}
	_ = result
}

func BenchmarkManyDependents(b *testing.B) {
	s := New(0)

	// Create 100 computed values depending on the same signal
	computeds := make([]*Computed[int], 100)
	for i := 0; i < 100; i++ {
		computeds[i] = NewComputed(func() int {
			return s.Get() * 2
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Set(i)
		// Access all computeds
		for _, c := range computeds {
			_ = c.Get()
		}
	}
}
