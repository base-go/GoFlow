# GoFlow

A Flutter-inspired GUI framework for Go with reactive state management and platform-adaptive design systems.

## 🚀 Quick Start

```bash
# Install the CLI
go install github.com/base-go/GoFlow/cmd/goflow@latest

# Create a new project (Flutter-style)
goflow create myapp

# Run it
cd myapp/macos
go run main.go
```

**New to GoFlow?** → Start with the [**Getting Started Guide**](./docs/GETTING_STARTED.md)

## 📚 Documentation

- **[Getting Started](./docs/GETTING_STARTED.md)** - Complete beginner tutorial
- **[Architecture](./ARCHITECTURE.md)** - Framework architecture deep dive
- **[CLI Reference](./docs/CLI.md)** - Complete CLI documentation
- **[Project Structure](./docs/PROJECT_STRUCTURE.md)** - Project organization guide
- **[Rendering Architecture](./docs/RENDERING.md)** - How rendering works
- **[Flutter Comparison](./docs/FLUTTER_INSPIRATION.md)** - For Flutter developers
- **[Workflow Guide](./docs/WORKFLOW.md)** - Complete technical flow
- **[All Documentation](./docs/)** - Browse all docs

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

## Available Widgets

GoFlow includes a comprehensive set of widgets inspired by Flutter:

### Layout Widgets
- **Column/Row**: Vertical/horizontal layout
- **Stack**: Layered widgets
- **Positioned**: Position children within Stack
- **Align**: Align child within parent
- **Container**: Padding, margin, sizing, colors
- **Center**: Center child widget
- **Padding**: Add padding around child
- **SizedBox**: Fixed size container
- **Expanded/Flexible**: Flex children in Row/Column
- **Spacer**: Empty space in flex layouts

### Form Widgets
- **TextField** (Material/Cupertino): Text input
- **Checkbox**: Material checkbox
- **Radio**: Material radio button
- **Switch** (Material/Cupertino): Toggle switch
- **Slider** (Material/Cupertino): Value slider

### Button Widgets
- **Button** (Material/Cupertino): Primary buttons
- **TextButton**: Text-only button (Material)
- **OutlinedButton**: Outlined button (Material)
- **IconButton**: Button with icon
- **FloatingActionButton**: Material FAB

### Display Widgets
- **Text**: Display text
- **Icon**: Display icons
- **Image**: Display images

### Scrolling Widgets
- **ListView**: Scrollable list
- **ListView.builder**: Lazy-loaded list
- **GridView**: Scrollable grid
- **SingleChildScrollView**: Scrollable single child

### Interaction Widgets
- **GestureDetector**: Detect gestures
- **InkWell**: Material ink splash effect
- **Draggable**: Make widget draggable
- **DragTarget**: Accept draggable widgets

### App Structure
- **Scaffold** (Material): Basic app structure
- **AppBar** (Material): Top app bar
- **Drawer**: Side navigation drawer
- **BottomNavigationBar**: Bottom navigation
- **CupertinoPageScaffold**: iOS app structure
- **CupertinoNavigationBar**: iOS navigation bar
- **CupertinoTabScaffold**: iOS tabbed interface

### Material-Specific
- **Card**: Material card
- **ListTile**: List item with leading/trailing
- **Dialog**: Material dialog
- **AlertDialog**: Alert dialog with actions
- **DrawerHeader**: Drawer header

### Cupertino-Specific
- **CupertinoTextField**: iOS text field
- **CupertinoSwitch**: iOS switch
- **CupertinoSlider**: iOS slider

## 📖 Examples

See the [examples](examples/) directory for complete examples:

### Signals Examples
- **[Playground](examples/playground/)** - Comprehensive framework feature test
- [Basic](examples/basic/) - Simple signal usage
- [Counter App](examples/counter-app/) - Interactive counter with multiple computed values
- [Shopping Cart](examples/shopping-cart/) - Shopping cart with reactive total
- [Form Validation](examples/form-validation/) - Real-time form validation
- [Todo List](examples/todo-list/) - Todo list using SignalSlice

### Design System Examples
- [Adaptive Demo](examples/adaptive-demo/) - Platform-adaptive widgets
- [Material Demo](examples/material-demo/) - Material Design widgets
- [Cupertino Demo](examples/cupertino-demo/) - iOS/macOS widgets
- [Widgets Showcase](examples/widgets-showcase/) - Comprehensive widget demonstration

## ⚡ Performance

GoFlow is designed for performance with:

- **Zero allocations** for basic signal operations
- **Lazy evaluation** for computed values
- **Efficient dependency tracking** using graph algorithms
- **Thread-safe** operations with minimal locking overhead

Run benchmarks:
```bash
go test ./signals -bench=. -benchmem
```

## 🎨 Framework Status (v0.1.0)

### ✅ What Works
- Widget system (Text, Container, Column, Center)
- Reactive signals (Signal, Computed, Effect)
- Signal collections (SignalSlice, SignalMap)
- Layout system (constraints, sizing)
- Three-tree architecture (Widget → Element → RenderObject)
- CLI tool for project creation
- Flutter-inspired project structure

### ⏳ In Progress
- Native rendering backends (Core Graphics, Direct2D, Cairo)
- GLFW window integration
- Event handling (mouse, keyboard)
- More built-in widgets
- Design system implementations (Material, Cupertino)

### 🔮 Planned
- Hot reload
- Animation system
- Advanced layout widgets (Row, Stack, Grid)
- Input widgets (TextField, Button)
- Platform features (menus, notifications)
- Complete Material Design implementation
- Complete Cupertino implementation

See [RENDERING.md](./docs/RENDERING.md) and [PLATFORM_INTEGRATION.md](./docs/PLATFORM_INTEGRATION.md) for details.

## 🤝 Contributing

GoFlow is in active development! We welcome contributions:

- 🐛 Report bugs
- 💡 Suggest features
- 📖 Improve documentation
- 🔧 Submit pull requests

See our [GitHub repository](https://github.com/base-go/GoFlow) for more information.

## 📚 Learn More

- [Architecture Guide](./ARCHITECTURE.md) - Deep dive into the framework
- [Getting Started](./docs/GETTING_STARTED.md) - Step-by-step tutorial
- [Workflow Guide](./docs/WORKFLOW.md) - Complete technical flow
- [Examples](./examples/) - Sample applications
- [API Documentation](https://pkg.go.dev/github.com/base-go/GoFlow) - Go package docs

## License

MIT
