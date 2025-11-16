# Getting Started with GoFlow

A complete guide to building your first GoFlow application - from installation to running a fully functional desktop app.

## What is GoFlow?

GoFlow is a **Flutter-inspired GUI framework for Go** with built-in reactive state management through signals. It allows you to build cross-platform desktop applications using Go with a declarative UI approach.

**Key Features:**
- 🎨 **Declarative UI**: Build UIs like Flutter or React
- ⚡ **Reactive State**: Signals-based state management (no setState!)
- 🖥️ **Desktop Focus**: macOS, Windows, Linux support
- 🦾 **Type-Safe**: Leverage Go's static typing with generics
- 🏗️ **Flutter-Inspired**: Familiar patterns for Flutter developers

## Prerequisites

- **Go 1.21 or later** installed ([download here](https://go.dev/dl/))
- Basic Go knowledge (packages, structs, interfaces)
- Terminal/command line familiarity

## Installation

### Step 1: Install the GoFlow CLI

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

This installs the `goflow` command-line tool to `$GOPATH/bin`.

### Step 2: Verify Installation

```bash
goflow version
```

**Expected output:**
```
GoFlow CLI v0.1.0
```

**Troubleshooting:** If you get "command not found", ensure `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add this to your `~/.bashrc`, `~/.zshrc`, or equivalent to make it permanent.

## Creating Your First App

### Step 1: Create a New Project

```bash
goflow create hello-world
```

This creates a complete project structure:

```
hello-world/
├── lib/                  # Your app code (shared across platforms)
│   └── main.go
├── macos/                # macOS platform runner
│   └── main.go
├── linux/                # Linux platform runner
│   └── main.go
├── windows/              # Windows platform runner
│   └── main.go
├── assets/               # Fonts, images, icons
├── test/                 # Tests
├── goflow.yaml           # Project configuration
├── go.mod                # Go dependencies
└── README.md
```

### Step 2: Navigate to Your Project

```bash
cd hello-world
```

### Step 3: Set Up Dependencies

Since GoFlow is currently in development, you'll need to point to your local GoFlow installation:

```bash
# Replace with your actual GoFlow path
echo "
replace github.com/base-go/GoFlow => /path/to/your/GoFlow
" >> go.mod

go mod tidy
```

**Future:** Once published, you'll just run `go mod tidy` without the replace directive.

### Step 4: Run Your App

```bash
# Run on your current platform
cd macos    # or linux, or windows
go run main.go
```

**Expected output:**
```
╔══════════════════════════════════╗
║    Welcome to Hello World!       ║
╚══════════════════════════════════╝

Starting GoFlow app...
Building widget tree...
✅ App rendered successfully!

Counter: 0
```

🎉 **Congratulations!** You've just run your first GoFlow app!

## Understanding the Generated Code

### The App Structure (`lib/main.go`)

Let's examine the generated app:

```go
package lib

import (
    "fmt"
    "github.com/base-go/GoFlow/goflow"
    "github.com/base-go/GoFlow/signals"
    "github.com/base-go/GoFlow/widgets"
)

// HelloWorldApp is your root widget
type HelloWorldApp struct {
    goflow.BaseWidget
    counter *signals.Signal[int]  // Reactive state
}

// NewHelloWorldApp creates a new app instance
func NewHelloWorldApp() *HelloWorldApp {
    return &HelloWorldApp{
        counter: signals.New(0),  // Initialize with 0
    }
}

// Build describes what the UI should look like
func (app *HelloWorldApp) Build(ctx goflow.BuildContext) goflow.Widget {
    // Reading the signal creates a reactive dependency
    count := app.counter.Get()

    return widgets.NewCenter(
        &widgets.Column{
            Children: []goflow.Widget{
                widgets.NewTextWithStyle(
                    "Welcome to Hello World!",
                    &goflow.TextStyle{
                        Color:      goflow.ColorBlue,
                        FontSize:   32,
                        FontWeight: goflow.FontWeightBold,
                    },
                ),
                spacer(20),
                widgets.NewText(
                    fmt.Sprintf("Counter: %d", count),
                ),
            },
        },
    )
}

// Increment updates the counter
func (app *HelloWorldApp) Increment() {
    app.counter.Update(func(v int) int {
        return v + 1
    })
}

// Run is called by platform runners
func Run() {
    app := NewHelloWorldApp()
    goflow.RunApp(app)
}

func spacer(height float64) goflow.Widget {
    h := height
    return &widgets.Container{Height: &h}
}
```

### Key Concepts

#### 1. **Widgets** - UI Building Blocks

Widgets are immutable descriptions of your UI:

```go
widgets.NewText("Hello")           // Simple text
widgets.NewCenter(child)            // Center a widget
&widgets.Column{Children: [...]}   // Vertical layout
&widgets.Container{...}             // Box with styling
```

#### 2. **Signals** - Reactive State

Instead of `setState()`, GoFlow uses signals:

```go
// Create a signal
counter := signals.New(0)

// Read the value (creates reactive dependency)
count := counter.Get()

// Update the value
counter.Set(5)
counter.Update(func(v int) int { return v + 1 })
```

**When you call `.Get()` inside `Build()`, GoFlow automatically rebuilds your UI when the signal changes!**

#### 3. **Build Method** - Declare Your UI

The `Build()` method returns a widget tree describing your UI:

```go
func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
    // Read state
    count := app.counter.Get()

    // Return widget tree
    return widgets.NewCenter(
        widgets.NewText(fmt.Sprintf("Count: %d", count))
    )
}
```

## Making Your First Modification

Let's add a decrement button to your app!

### Step 1: Add a Decrement Method

Edit `lib/main.go` and add this method after `Increment()`:

```go
func (app *HelloWorldApp) Decrement() {
    app.counter.Update(func(v int) int {
        return v - 1
    })
}
```

### Step 2: Update the Build Method

Modify the `Build()` method to show both operations:

```go
func (app *HelloWorldApp) Build(ctx goflow.BuildContext) goflow.Widget {
    count := app.counter.Get()

    return widgets.NewCenter(
        &widgets.Column{
            Children: []goflow.Widget{
                widgets.NewTextWithStyle(
                    "Counter App",
                    &goflow.TextStyle{
                        Color:      goflow.ColorBlue,
                        FontSize:   32,
                        FontWeight: goflow.FontWeightBold,
                    },
                ),
                spacer(20),
                widgets.NewText(
                    fmt.Sprintf("Counter: %d", count),
                ),
                spacer(10),
                widgets.NewText("Click Increment or Decrement!"),
            },
        },
    )
}
```

### Step 3: Test Your Changes

In `lib/main.go`, temporarily modify the `Run()` function to test:

```go
func Run() {
    app := NewHelloWorldApp()

    // Set up an effect to watch changes
    dispose := signals.NewEffect(func() {
        count := app.counter.Get()
        fmt.Printf("Counter changed: %d\n", count)
    })
    defer dispose()

    // Build the app
    goflow.RunApp(app)

    // Test the counter
    fmt.Println("\nTesting increment...")
    app.Increment()
    app.Increment()
    app.Increment()

    fmt.Println("\nTesting decrement...")
    app.Decrement()

    fmt.Printf("\nFinal count: %d\n", app.counter.Get())
}
```

Run it again:

```bash
cd macos
go run main.go
```

**Expected output:**
```
Counter changed: 0

Testing increment...
Counter changed: 1
Counter changed: 2
Counter changed: 3

Testing decrement...
Counter changed: 2

Final count: 2
```

✅ **Your reactive state management is working!**

## Understanding the Architecture

### The Three Trees

GoFlow maintains three parallel trees (just like Flutter):

```
┌─────────────┐
│   Widget    │  Immutable UI description
│   Tree      │  (What to show)
└──────┬──────┘
       │ Build()
       ↓
┌─────────────┐
│   Element   │  Mutable lifecycle manager
│   Tree      │  (When to rebuild)
└──────┬──────┘
       │ CreateRenderObject()
       ↓
┌─────────────┐
│  Render     │  Layout, Paint, Hit Test
│  Tree       │  (How to draw)
└─────────────┘
```

**You work with widgets** - GoFlow handles the rest!

### How Rendering Works (Currently)

**Current Status (v0.1.0):** GoFlow has the rendering **architecture** but no actual rendering backend yet.

```
Your App → Widget Tree → Element Tree → RenderObject Tree → Canvas Interface
                                                                      ↓
                                                          ⚠️ MockCanvas (testing only)
```

**Future (v0.2.0+):** Native rendering backends:
- **macOS**: Core Graphics + Core Text
- **Windows**: Direct2D + DirectWrite
- **Linux**: Cairo + Pango

See [RENDERING.md](./RENDERING.md) for detailed rendering architecture.

## Advanced: Working with Computed Signals

Computed signals automatically derive values from other signals:

```go
type TemperatureApp struct {
    goflow.BaseWidget
    celsius    *signals.Signal[float64]
    fahrenheit *signals.Computed[float64]
}

func NewTemperatureApp() *TemperatureApp {
    celsius := signals.New(20.0)

    return &TemperatureApp{
        celsius: celsius,
        fahrenheit: signals.NewComputed(func() float64 {
            return celsius.Get()*9/5 + 32  // Auto-updates!
        }),
    }
}

func (app *TemperatureApp) Build(ctx goflow.BuildContext) goflow.Widget {
    c := app.celsius.Get()
    f := app.fahrenheit.Get()

    return widgets.NewCenter(
        &widgets.Column{
            Children: []goflow.Widget{
                widgets.NewText(fmt.Sprintf("%.1f°C", c)),
                widgets.NewText(fmt.Sprintf("%.1f°F", f)),
            },
        },
    )
}
```

When `celsius` changes, `fahrenheit` automatically recomputes!

## Advanced: Signal Collections

### SignalSlice - Reactive Arrays

```go
todos := signals.NewSlice([]string{
    "Learn GoFlow",
    "Build an app",
})

// Reactive operations
todos.Append("Deploy app")
todos.RemoveAt(0)

// Use in Build()
func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
    todoList := app.todos.Get()  // Reactive!

    children := make([]goflow.Widget, len(todoList))
    for i, todo := range todoList {
        children[i] = widgets.NewText(todo)
    }

    return &widgets.Column{Children: children}
}
```

### SignalMap - Reactive Dictionaries

```go
settings := signals.NewMap(map[string]bool{
    "darkMode": false,
    "notifications": true,
})

settings.SetKey("darkMode", true)  // Triggers rebuild!
settings.DeleteKey("notifications")

// Check values
if settings.GetKey("darkMode").Get() {
    // Dark mode is enabled
}
```

## Project Templates

GoFlow provides three templates:

### 1. Default Template (Recommended)

```bash
goflow create myapp
```

Includes:
- Counter example
- Signal-based state
- Multiple widgets
- Good starting point

### 2. Minimal Template

```bash
goflow create myapp --template=minimal
```

Simplest possible app:
- Just "Hello World"
- No state management
- Perfect for learning basics

### 3. Material Template

```bash
goflow create myapp --template=material
```

Polished UI example:
- Styled containers
- Card-like layout
- Material Design inspiration
- Production-ready starting point

## Platform-Specific Development

### Create Platform-Specific Project

```bash
# macOS only
goflow create myapp --platforms=macos

# macOS + Linux
goflow create myapp --platforms=macos,linux

# All platforms
goflow create myapp --platforms=macos,windows,linux
```

### Run on Different Platforms

```bash
# macOS
cd myapp/macos
go run main.go

# Linux
cd myapp/linux
go run main.go

# Windows
cd myapp/windows
go run main.go
```

### Platform Runners

Platform runners (`macos/main.go`, etc.) are simple:

```go
package main

import "com.example/myapp/lib"

func main() {
    // Future: Initialize GLFW window
    // Future: Initialize WGPU rendering
    // Future: Set up event loop

    lib.Run()  // Calls into shared lib code
}
```

**All your app logic lives in `lib/`** - platform runners just bootstrap!

## Next Steps

### 1. Explore Examples

Check out the [examples](../examples/) directory:
- **playground**: Comprehensive framework test
- **goflow-hello**: Minimal example
- **goflow-counter**: Interactive counter

### 2. Read the Documentation

- [**Architecture Guide**](../ARCHITECTURE.md) - Deep dive into widgets, elements, render objects
- [**CLI Documentation**](./CLI.md) - Complete CLI reference
- [**Project Structure**](./PROJECT_STRUCTURE.md) - Organize your code
- [**Rendering Architecture**](./RENDERING.md) - How rendering works
- [**Flutter Inspiration**](./FLUTTER_INSPIRATION.md) - Compare with Flutter

### 3. Build Something!

Try building:
- **Todo List App**: Use `SignalSlice` for tasks
- **Settings Screen**: Use `SignalMap` for configuration
- **Form Validation**: Real-time validation with signals
- **Dashboard**: Multiple computed values

### 4. Contribute

GoFlow is in active development! Check the [GitHub repository](https://github.com/base-go/GoFlow) to:
- Report issues
- Request features
- Submit pull requests
- Join discussions

## Common Patterns

### Loading State

```go
type MyApp struct {
    goflow.BaseWidget
    isLoading *signals.Signal[bool]
    data      *signals.Signal[string]
}

func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
    loading := app.isLoading.Get()

    if loading {
        return widgets.NewText("Loading...")
    }

    data := app.data.Get()
    return widgets.NewText(data)
}
```

### Conditional Rendering

```go
func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
    isLoggedIn := app.isLoggedIn.Get()

    if isLoggedIn {
        return app.buildDashboard()
    } else {
        return app.buildLoginScreen()
    }
}
```

### List Rendering

```go
func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
    items := app.items.Get()

    children := make([]goflow.Widget, len(items))
    for i, item := range items {
        children[i] = widgets.NewText(item)
    }

    return &widgets.Column{Children: children}
}
```

## Troubleshooting

### "Cannot find package"

Add a replace directive:

```bash
echo "
replace github.com/base-go/GoFlow => /path/to/GoFlow
" >> go.mod
go mod tidy
```

### "goflow: command not found"

Ensure `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### "package lib is not in GOROOT"

Make sure `lib/main.go` has `package lib` (not `package main`)!

### Platform runner won't compile

Check that:
1. `lib/main.go` has `package lib`
2. `Run()` function is exported (capital R)
3. No syntax errors in lib code

## FAQs

### Q: Does GoFlow have hot reload like Flutter?

**A:** Not yet, but it's planned for a future version!

### Q: Can I build for iOS/Android?

**A:** No, GoFlow is desktop-focused (macOS, Windows, Linux). For mobile, use Flutter!

### Q: How does GoFlow compare to Flutter?

See the [Flutter Inspiration](./FLUTTER_INSPIRATION.md) guide for detailed comparison.

### Q: What rendering backend does GoFlow use?

Currently none - just a mock canvas for testing. Future versions will use **native backends** (Core Graphics, Direct2D, Cairo). See [RENDERING.md](./RENDERING.md).

### Q: Can I use GoFlow in production?

**Not yet!** GoFlow is in early development (v0.1.0). Current status:
- ✅ Widget system works
- ✅ Signals work
- ✅ Layout system works
- ❌ No actual rendering yet
- ❌ No event handling yet
- ❌ No window creation yet

## Summary

You've learned:
- ✅ How to install the GoFlow CLI
- ✅ How to create a new project
- ✅ The basic project structure
- ✅ How to use signals for reactive state
- ✅ How to build UIs with widgets
- ✅ How to run your app on different platforms
- ✅ The three-tree architecture
- ✅ Advanced patterns (computed signals, collections)

**Now go build something awesome!** 🚀

## Get Help

- 📖 [Full Documentation](../README.md)
- 💬 [GitHub Discussions](https://github.com/base-go/GoFlow/discussions)
- 🐛 [Report Issues](https://github.com/base-go/GoFlow/issues)
- 📧 Community Support (coming soon)

---

**Happy GoFlowing!** 🎨✨
