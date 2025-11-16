# GoFlow

A Flutter-inspired GUI framework for Go with reactive state management, inspired by signals pattern from Preact/Solid.js.

## Quick Start

```bash
# Install the CLI
go install github.com/base-go/GoFlow/cmd/goflow@latest

# Create a new project (Flutter-style)
goflow new myapp

# Run it
cd myapp/macos
go run main.go
```

### Project Structure

GoFlow follows Flutter's project structure philosophy:

```
myapp/
├── lib/                    # Shared application code (like Flutter's lib/)
│   └── main.go
├── macos/                  # macOS platform runner
├── linux/                  # Linux platform runner
├── windows/                # Windows platform runner
├── assets/                 # Images, fonts, icons
├── test/                   # Tests
├── goflow.yaml             # Project configuration (like pubspec.yaml)
├── go.mod                  # Go dependencies
└── README.md
```

## Features

- **Signal**: Reactive value containers that notify listeners on change
- **Computed**: Automatically derived values with dependency tracking
- **Effect**: Side effects that run when dependencies change
- **Batch**: Group multiple updates to prevent unnecessary recomputation
- **Untracked**: Read signal values without creating subscriptions
- **Collections**: Reactive slices and maps with built-in helpers (Filter, Map, etc.)

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

## Collections

GoFlow includes reactive collections for common data structures:

### SignalSlice

```go
// Create a reactive slice
todos := signals.NewSlice([]string{"Learn Go", "Build app"})

// Reactive operations
todos.Append("Deploy to production")
todos.Prepend("Plan project")
todos.RemoveAt(1)

// Reactive computed operations
activeTodos := todos.Filter(func(todo string) bool {
    return !strings.HasPrefix(todo, "Done:")
})

// Get length reactively
count := todos.Len()
fmt.Printf("Total todos: %d\n", count.Get())
```

### SignalMap

```go
// Create a reactive map
users := signals.NewMap(map[string]int{
    "alice": 25,
    "bob": 30,
})

// Reactive operations
users.SetKey("charlie", 35)
users.DeleteKey("bob")

// Reactive queries
hasAlice := users.Has("alice")
aliceAge := users.GetKey("alice")
userCount := users.Size()
```

## Advanced Usage

### Batch Updates

Prevent unnecessary re-computations by batching multiple updates:

```go
signals.Batch(func() {
    firstName.Set("John")
    lastName.Set("Doe")
    age.Set(30)
})
// Effect runs only once after all updates
```

### Untracked Reads

Read signal values without creating dependencies:

```go
dispose := signals.NewEffect(func() {
    current := counter.Get() // Tracked
    previous := signals.Untracked(func() int {
        return prevCounter.Get() // Not tracked
    })
    fmt.Printf("Changed from %d to %d\n", previous, current)
})
```

## Examples

See the [examples](examples/) directory for more complete examples:

- [Basic](examples/basic/) - Simple signal usage
- [Counter App](examples/counter-app/) - Interactive counter with multiple computed values
- [Shopping Cart](examples/shopping-cart/) - Shopping cart with reactive total
- [Form Validation](examples/form-validation/) - Real-time form validation
- [Todo List](examples/todo-list/) - Todo list using SignalSlice

## Performance

GoFlow is designed for performance with:

- **Zero allocations** for basic signal operations
- **Lazy evaluation** for computed values
- **Efficient dependency tracking** using graph algorithms
- **Thread-safe** operations with minimal locking overhead

Run benchmarks:
```bash
go test ./signals -bench=. -benchmem
```

## License

MIT
