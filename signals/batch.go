package signals

// Batch groups multiple signal updates into a single notification cycle
// This prevents unnecessary recomputations when multiple signals change together
func Batch(fn func()) {
	startBatch()
	defer endBatch()
	fn()
}

// Untracked runs a function without tracking any signal dependencies
// Useful for reading signals inside effects without creating subscriptions
func Untracked[T any](fn func() T) T {
	prev := startTracking(nil)
	defer stopTracking(prev)
	return fn()
}
