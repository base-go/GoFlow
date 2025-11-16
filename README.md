# GoFlow

A reactive state management library for Go, inspired by signals pattern from Preact/Dart.

## Features

- **Signal**: Reactive value containers that notify listeners on change
- **Computed**: Automatically derived values with dependency tracking
- **Effect**: Side effects that run when dependencies change
- **Batch**: Group multiple updates to prevent unnecessary recomputation
- **Untracked**: Read signal values without creating subscriptions

## Installation

```bash
go get github.com/base-go/GoFlow
```

## Quick Start

```go
import "github.com/base-go/GoFlow/signals"

// Create signals
counter := signals.New(0)
doubled := signals.NewComputed(func() int {
    return counter.Get() * 2
})

// React to changes
dispose := signals.NewEffect(func() {
    fmt.Printf("Counter: %d, Doubled: %d\n", counter.Get(), doubled.Get())
})
defer dispose()

// Update (triggers effect)
counter.Set(5) // Prints: Counter: 5, Doubled: 10
```

## License

MIT
