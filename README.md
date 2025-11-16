# GoFlow

A reactive GUI framework for Go with signals-based state management and platform-adaptive design systems.

## Features

### Core Framework
- **Signal**: Reactive value containers that notify listeners on change
- **Computed**: Automatically derived values with dependency tracking
- **Effect**: Side effects that run when dependencies change
- **Batch**: Group multiple updates to prevent unnecessary recomputation
- **Untracked**: Read signal values without creating subscriptions
- **Collections**: Reactive slices and maps with built-in helpers (Filter, Map, etc.)

### GUI & Design Systems
- **Adaptive Widgets**: Automatically switch between Material and Cupertino based on platform
- **Material Design**: Google's design system (Android, Linux, Windows, Web)
- **Cupertino**: Apple's design language (iOS, macOS)
- **Platform Detection**: Automatic platform-specific styling
- **Widget System**: Declarative UI with reactive updates

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

## Design Systems

GoFlow provides three approaches to building cross-platform UIs:

### Adaptive Widgets (Recommended)
Write once, automatically adapts to platform:

```go
import "github.com/base-go/GoFlow/adaptive"

// Button automatically becomes Material or Cupertino
button := adaptive.NewButton("Click Me", func() {
    fmt.Println("Clicked!")
})

// Card adapts to platform style
card := adaptive.NewCard(content)

// AppBar becomes Material AppBar or Cupertino NavigationBar
appBar := adaptive.NewAppBar(title)
```

### Platform-Specific Widgets

Use Material Design or Cupertino explicitly:

```go
// Material Design (Android, Web, Desktop)
import "github.com/base-go/GoFlow/material"
btn := material.NewButton(child, onPressed)

// Cupertino (iOS, macOS)
import "github.com/base-go/GoFlow/cupertino"
btn := cupertino.NewButton(child, onPressed)
```

See [DESIGN_SYSTEMS.md](DESIGN_SYSTEMS.md) for complete documentation.

## Examples

See the [examples](examples/) directory for more complete examples:

### Signals Examples
- [Basic](examples/basic/) - Simple signal usage
- [Counter App](examples/counter-app/) - Interactive counter with multiple computed values
- [Shopping Cart](examples/shopping-cart/) - Shopping cart with reactive total
- [Form Validation](examples/form-validation/) - Real-time form validation
- [Todo List](examples/todo-list/) - Todo list using SignalSlice

### Design System Examples
- [Adaptive Demo](examples/adaptive-demo/) - Platform-adaptive widgets
- [Material Demo](examples/material-demo/) - Material Design widgets
- [Cupertino Demo](examples/cupertino-demo/) - iOS/macOS widgets

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
