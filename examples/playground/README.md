# GoFlow Playground Example

A comprehensive example application that demonstrates and tests all major features of the GoFlow framework.

## Purpose

This playground serves as both a feature demonstration and a testing ground for the GoFlow framework. It exercises all core capabilities in a single, well-documented example.

## Features Tested

### 1. **Basic Signals** ✅
- Counter (int signal)
- Temperature (float64 signal)
- Username (string signal)
- Demonstrates reactive state updates

### 2. **Computed Signals** ✅
- Counter doubled (derived from counter)
- Temperature conversion (Celsius to Fahrenheit)
- Greeting message (derived from username)
- Todo count (derived from todo list length)
- Shows automatic recomputation when dependencies change

### 3. **Signal Collections** ✅

#### SignalSlice
- Todo list management
- Append, Prepend, RemoveAt operations
- Reactive length computation
- Filter and Map transformations

#### SignalMap
- Settings management (key-value pairs)
- Update operations
- Reactive map updates

### 4. **Widget System** ✅
- **Text Widget**: Simple and styled text rendering
- **Container Widget**: Layout with padding, margin, colors, sizing
- **Column Widget**: Vertical layout with alignment
- **Center Widget**: Centering content

### 5. **Layout System** ✅
- Padding (symmetric and all-sides)
- Sizing (width/height constraints)
- Colors and backgrounds
- Alignment options
- EdgeInsets (margin/padding)

### 6. **Reactive Updates** ✅
- Effects for side-effect tracking
- Automatic rebuilds on signal changes
- Batch updates for performance
- Effect cleanup

### 7. **Type System** ✅
- Generic signals (int, float64, string, slices, maps)
- Type-safe operations
- Computed signal inference

## Running the Example

```bash
cd examples/playground
go run main.go
```

## Expected Output

The example will:
1. Set up reactive effects to monitor changes
2. Build the widget tree
3. Run a series of tests demonstrating each feature
4. Print detailed output showing reactive updates
5. Display final state summary

## Code Structure

```go
type PlaygroundApp struct {
    // Basic signals
    counter     *signals.Signal[int]
    temperature *signals.Signal[float64]
    username    *signals.Signal[string]

    // Computed signals
    counterDoubled  *signals.Computed[int]
    temperatureInF  *signals.Computed[float64]
    greeting        *signals.Computed[string]

    // Collections
    todos       *signals.SignalSlice[string]
    settings    *signals.SignalMap[string, bool]
}
```

## What You'll Learn

### Signal Management
```go
// Create a signal
counter := signals.New(0)

// Update with a function
counter.Update(func(v int) int { return v + 1 })

// Get the current value
count := counter.Get()
```

### Computed Values
```go
// Create a computed signal
doubled := signals.NewComputed(func() int {
    return counter.Get() * 2
})
// Automatically updates when counter changes!
```

### Collections
```go
// Signal slice
todos := signals.NewSlice([]string{"item1", "item2"})
todos.Append("item3")
todos.RemoveAt(0)

// Signal map
settings := signals.NewMap(map[string]bool{"key": true})
settings.Update(func(m map[string]bool) map[string]bool {
    m["key"] = !m["key"]
    return m
})
```

### Effects
```go
// Run side effects when signals change
dispose := signals.NewEffect(func() {
    count := counter.Get()
    fmt.Printf("Counter changed: %d\n", count)
})
defer dispose() // Clean up when done
```

### Building UI
```go
func (app *PlaygroundApp) Build(ctx goflow.BuildContext) goflow.Widget {
    count := app.counter.Get() // Creates reactive dependency

    return widgets.NewCenter(
        &widgets.Column{
            Children: []goflow.Widget{
                widgets.NewText(fmt.Sprintf("Count: %d", count)),
                // More widgets...
            },
        },
    )
}
```

## Testing Checklist

When running this example, verify:

- ✅ Signals update correctly
- ✅ Computed signals recalculate automatically
- ✅ Effects trigger on signal changes
- ✅ Collections (slice/map) work reactively
- ✅ Batch updates optimize performance
- ✅ Widget tree builds without errors
- ✅ Layout system handles padding/sizing
- ✅ Colors and styles render correctly
- ✅ Type safety is maintained throughout

## Next Steps

After understanding the playground:

1. **Modify the example** - Add your own signals and widgets
2. **Create custom widgets** - Build reusable components
3. **Experiment with layout** - Try different widget compositions
4. **Add more features** - Implement forms, lists, grids
5. **Build a real app** - Use the patterns learned here

## Architecture Insights

This example demonstrates the core GoFlow architecture:

- **Widget**: Immutable UI descriptions
- **Element**: Manages widget lifecycle
- **RenderObject**: Handles layout and painting
- **Signals**: Reactive state management
- **Build**: Declarative UI construction

See [ARCHITECTURE.md](../../ARCHITECTURE.md) for detailed framework documentation.

## Related Examples

- `goflow-hello`: Minimal "Hello World" example
- `goflow-counter`: Basic counter with interactions
- `todo-list`: Task management app
- `form-validation`: Form handling patterns

## Contributing

Found an issue or want to add more tests? Please contribute to make this playground even more comprehensive!
