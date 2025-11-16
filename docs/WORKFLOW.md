# GoFlow Complete Workflow Guide

This document explains the complete workflow of GoFlow from installation to rendering, showing how all the pieces fit together.

## Table of Contents

1. [Overview](#overview)
2. [CLI Tool](#cli-tool)
3. [Framework Packages](#framework-packages)
4. [Project Structure](#project-structure)
5. [Development Workflow](#development-workflow)
6. [Complete Execution Flow](#complete-execution-flow)
7. [Rendering Pipeline](#rendering-pipeline)
8. [State Management Flow](#state-management-flow)

## Overview

GoFlow consists of multiple interconnected components:

```
┌─────────────────────────────────────────────────────────┐
│                    GoFlow Ecosystem                     │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  ┌───────────────┐         ┌──────────────────┐       │
│  │   CLI Tool    │────────▶│  Framework Core  │       │
│  │  (cmd/goflow) │ creates │   (goflow/,      │       │
│  │               │ projects│    signals/,     │       │
│  └───────────────┘   using │    widgets/)     │       │
│         │                   └──────────────────┘       │
│         │                            │                 │
│         ▼                            │                 │
│  ┌───────────────────────────────────┼───────────┐    │
│  │        Generated Project          │           │    │
│  │  ┌──────────┐  ┌───────────────┐ │           │    │
│  │  │   lib/   │  │   Platform    │◀┘           │    │
│  │  │  main.go │◀─│   Runners     │             │    │
│  │  └──────────┘  │ (macos/etc)   │             │    │
│  │                 └───────────────┘             │    │
│  └───────────────────────────────────────────────┘    │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

## CLI Tool

### What is it?

The GoFlow CLI (`cmd/goflow/main.go`) is a **command-line tool** for creating and managing GoFlow projects.

### Installation

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

This compiles the CLI tool and installs the `goflow` binary to `$GOPATH/bin`.

### What it does

1. **Templates Management**: Embeds project templates using Go's `embed.FS`
2. **Project Generation**: Creates directory structure and files
3. **Platform Support**: Generates platform-specific runners
4. **Configuration**: Creates `goflow.yaml`, `go.mod`, and other configs

### CLI Execution Flow

```
User runs: goflow create myapp
         ↓
Parse command-line arguments
         ↓
Select template (default, minimal, material)
         ↓
Create directory structure
         ↓
Generate files from templates:
  - lib/main.go (app code)
  - platform/main.go (runners)
  - goflow.yaml (config)
  - go.mod (dependencies)
         ↓
Initialize git repository
         ↓
Print success message
```

## Framework Packages

### Core Packages

GoFlow framework consists of multiple Go packages:

```
github.com/base-go/GoFlow/
├── goflow/          # Core framework (Widget, Element, RenderObject)
├── signals/         # Reactive state management
└── widgets/         # Built-in widgets (Text, Container, etc.)
```

### How it's used

Your generated project imports these packages:

```go
import (
    "github.com/base-go/GoFlow/pkg/core/framework"
    "github.com/base-go/GoFlow/pkg/core/signals"
    "github.com/base-go/GoFlow/pkg/core/widgets"
)
```

### Dual Nature

GoFlow is **both**:
1. **CLI tool**: Creates projects
2. **Framework library**: Provides APIs for building UIs

Think of it like Flutter:
- `flutter` command creates projects
- Flutter SDK provides widgets, state management, etc.

## Project Structure

### Generated Structure

When you run `goflow create myapp`:

```
myapp/
├── lib/                    # Your application code (package lib)
│   └── main.go            # Defines MyappApp struct and Run() function
│
├── macos/                 # macOS platform (package main)
│   └── main.go            # Imports lib and calls lib.Run()
│
├── linux/                 # Linux platform (package main)
│   └── main.go            # Imports lib and calls lib.Run()
│
├── windows/               # Windows platform (package main)
│   └── main.go            # Imports lib and calls lib.Run()
│
├── go.mod                 # Go module definition
├── goflow.yaml            # GoFlow project config
└── README.md
```

### Package Relationship

```
┌─────────────────────┐
│   macos/main.go     │  package main
│                     │
│  import "com.ex/    │
│    myapp/lib"       │
│                     │
│  func main() {      │
│    lib.Run()        │──┐
│  }                  │  │
└─────────────────────┘  │
                         │
                         ▼
                  ┌─────────────────────┐
                  │   lib/main.go       │  package lib
                  │                     │
                  │  type MyappApp      │
                  │    struct {...}     │
                  │                     │
                  │  func Run() {       │
                  │    app := New...()  │
                  │    goflow.RunApp()  │
                  │  }                  │
                  └─────────────────────┘
```

**Key Point:** Platform runners are **entry points** that call into shared `lib/` code.

## Development Workflow

### Step-by-Step Workflow

#### 1. Install CLI Tool

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

**What happens:**
- Go downloads GoFlow source
- Compiles `cmd/goflow/main.go`
- Installs binary to `$GOPATH/bin/goflow`

#### 2. Create Project

```bash
goflow create myapp
```

**What happens:**
- CLI creates `myapp/` directory
- Generates all files from embedded templates
- Creates `go.mod` with module path `com.example/myapp`
- Initializes git repository

#### 3. Set Up Dependencies

```bash
cd myapp
echo "replace github.com/base-go/GoFlow => /path/to/GoFlow" >> go.mod
go mod tidy
```

**What happens:**
- Points to local GoFlow development copy
- Downloads and caches dependencies
- Creates `go.sum` checksum file

**In production:** Users would skip the replace directive and just run `go mod tidy`.

#### 4. Run Application

```bash
cd macos  # or windows, or linux
go run main.go
```

**What happens:**
- Go compiles `macos/main.go`
- Imports `com.example/myapp/lib`
- Calls `lib.Run()`
- `lib.Run()` calls `goflow.RunApp(app)`
- GoFlow builds widget tree
- Prints output (currently no real rendering)

## Complete Execution Flow

### From Command to Pixels (Current + Future)

```
┌─────────────────────────────────────────────────────────────┐
│ 1. USER ACTION                                              │
└─────────────────────────────────────────────────────────────┘
         │
         │  $ cd myapp/macos && go run main.go
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 2. PLATFORM RUNNER (macos/main.go)                          │
│                                                              │
│    package main                                              │
│                                                              │
│    import "com.example/myapp/lib"                            │
│                                                              │
│    func main() {                                             │
│        // TODO: Initialize GLFW window                       │
│        // TODO: Initialize WGPU rendering                    │
│        // TODO: Set up event loop                            │
│                                                              │
│        lib.Run()  ◄─── Calls into shared lib                │
│    }                                                         │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 3. APPLICATION CODE (lib/main.go)                           │
│                                                              │
│    package lib                                               │
│                                                              │
│    type MyappApp struct {                                    │
│        goflow.BaseWidget                                     │
│        counter *signals.Signal[int]                          │
│    }                                                         │
│                                                              │
│    func (app *MyappApp) Build(ctx BuildContext) Widget {    │
│        count := app.counter.Get()  // Reactive!             │
│        return widgets.NewCenter(...)                         │
│    }                                                         │
│                                                              │
│    func Run() {                                              │
│        app := NewMyappApp()                                  │
│        goflow.RunApp(app)  ◄─── Enters framework            │
│    }                                                         │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 4. GOFLOW FRAMEWORK (goflow.RunApp)                         │
│                                                              │
│    func RunApp(app Widget) {                                │
│        // 1. Create root element                             │
│        element := app.CreateElement()                        │
│                                                              │
│        // 2. Mount element tree                              │
│        element.Mount(nil, nil)                               │
│                                                              │
│        // 3. Build widget tree                               │
│        element.Rebuild()                                     │
│           └─▶ app.Build(ctx) ──┐                            │
│                                 │                            │
│        // 4. Layout render tree │                            │
│        renderObject.Layout(...)─┤                            │
│                                 │                            │
│        // 5. Paint to canvas    │                            │
│        renderObject.Paint(...)──┘                            │
│    }                                                         │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 5. WIDGET TREE (app.Build returns)                          │
│                                                              │
│    Center                                                    │
│      └─ Column                                               │
│           ├─ Text("Welcome")                                 │
│           ├─ Container (spacer)                              │
│           └─ Text("Counter: 5")                              │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 6. ELEMENT TREE (parallel to widgets)                       │
│                                                              │
│    CenterElement                                             │
│      └─ ColumnElement                                        │
│           ├─ TextElement                                     │
│           ├─ ContainerElement                                │
│           └─ TextElement                                     │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 7. RENDER TREE (subset with RenderObjects)                  │
│                                                              │
│    RenderCenter                                              │
│      └─ RenderColumn                                         │
│           ├─ RenderText("Welcome")                           │
│           ├─ RenderPadding (spacer)                          │
│           └─ RenderText("Counter: 5")                        │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 8. LAYOUT PHASE                                              │
│                                                              │
│    RenderCenter.Layout(constraints)                          │
│      ├─ Passes constraints to child                          │
│      ▼                                                       │
│    RenderColumn.Layout(constraints)                          │
│      ├─ Measures children                                    │
│      ├─ Computes total height                                │
│      ├─ Returns size to parent                               │
│      ▼                                                       │
│    Each child computes its size                              │
│      └─ Sizes bubble up                                      │
│      ▲─ Constraints flow down                                │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 9. PAINT PHASE                                               │
│                                                              │
│    RenderCenter.Paint(canvas, offset)                        │
│      └─ Centers child, calls child.Paint()                   │
│           ▼                                                  │
│    RenderColumn.Paint(canvas, offset)                        │
│      └─ Paints each child at computed position               │
│           ├─ RenderText.Paint() → canvas.DrawText()         │
│           ├─ RenderPadding.Paint() → (nothing)              │
│           └─ RenderText.Paint() → canvas.DrawText()         │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 10. CANVAS INTERFACE                                         │
│                                                              │
│    type Canvas interface {                                   │
│        DrawRect(rect, paint)                                 │
│        DrawText(text, offset, style)                         │
│        DrawCircle(center, radius, paint)                     │
│        // ... transform operations                           │
│    }                                                         │
│                                                              │
│    Current: MockCanvas (testing only)                        │
│    Future:  NativeCanvas (Core Graphics, Direct2D, Cairo)   │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 11. RENDERING BACKEND (Future)                               │
│                                                              │
│    macOS:   Core Graphics → GPU → Screen                     │
│    Windows: Direct2D → GPU → Screen                          │
│    Linux:   Cairo → GPU → Screen                             │
│                                                              │
│    OR: WGPU → Metal/DirectX/Vulkan → Screen                  │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 12. PIXELS ON SCREEN! 🎨                                     │
└─────────────────────────────────────────────────────────────┘
```

## Rendering Pipeline

### Current State (v0.1.0)

```
Widget Tree → Element Tree → RenderObject Tree → Canvas Interface
                                                        ↓
                                                  MockCanvas
                                                        ↓
                                              (No actual rendering)
```

### Future State (v0.2.0+)

#### Option 1: Native Backends (Recommended)

```
Widget Tree → Element Tree → RenderObject Tree → Canvas Interface
                                                        ↓
                                          ┌─────────────┴──────────────┐
                                          ▼                            ▼
                                   macOS/Windows/Linux          Native APIs
                                          ↓                            ↓
                                   CoreGraphicsCanvas         Core Graphics
                                   Direct2DCanvas       OR    Direct2D
                                   CairoCanvas                Cairo
                                          ↓                            ↓
                                   Platform GPU APIs              Metal
                                          ↓                       DirectX
                                       Screen                    Vulkan
```

#### Option 2: WGPU Backend

```
Widget Tree → Element Tree → RenderObject Tree → Canvas Interface
                                                        ↓
                                                   WGPUCanvas
                                                        ↓
                                                  wgpu-native
                                                        ↓
                                          ┌─────────────┴──────────────┐
                                          ▼                            ▼
                                    Metal (macOS)           DirectX (Windows)
                                          ↓                            ↓
                                       Screen                       Screen

                                          ▼
                                   Vulkan (Linux)
                                          ↓
                                       Screen
```

See [RENDERING.md](./RENDERING.md) for detailed rendering architecture.

## State Management Flow

### Signal-Based Reactivity

```
┌─────────────────────────────────────────────────────────────┐
│ 1. SIGNAL CREATION                                           │
│                                                              │
│    counter := signals.New(0)                                 │
│                                                              │
│    // Creates reactive container with:                       │
│    // - Value: 0                                             │
│    // - Subscribers: []                                      │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 2. BUILD PHASE (Creates Dependency)                          │
│                                                              │
│    func (app *MyApp) Build(ctx BuildContext) Widget {       │
│        count := app.counter.Get()  ◄─ Reading signal        │
│        //                             creates dependency     │
│        return widgets.NewText(fmt.Sprintf("%d", count))      │
│    }                                                         │
│                                                              │
│    After .Get():                                             │
│    counter.subscribers = [Build function]                    │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 3. STATE UPDATE                                              │
│                                                              │
│    counter.Set(5)                                            │
│        ↓                                                     │
│    counter.value = 5                                         │
│        ↓                                                     │
│    Notify all subscribers                                    │
│        ↓                                                     │
│    For each subscriber:                                      │
│        element.MarkNeedsBuild()  ◄─ Mark for rebuild        │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 4. REBUILD PHASE                                             │
│                                                              │
│    element.Rebuild()                                         │
│        ↓                                                     │
│    newWidget = app.Build(ctx)  ◄─ Calls Build again         │
│        ↓                                                     │
│    count := counter.Get()  // Now returns 5                  │
│        ↓                                                     │
│    element.Update(newWidget)                                 │
│        ↓                                                     │
│    Re-layout and re-paint if needed                          │
└─────────────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│ 5. UI UPDATES                                                │
│                                                              │
│    Screen shows new value: "Counter: 5"                      │
└─────────────────────────────────────────────────────────────┘
```

### Computed Signals

```
counter := signals.New(0)
doubled := signals.NewComputed(func() int {
    return counter.Get() * 2  ◄─ Creates dependency
})

// Dependency graph:
counter ──▶ doubled ──▶ Build()
    │
    └──────────────────▶ Build() (if used directly)

// When counter.Set(5):
counter changes → doubled recomputes → Build() reruns → UI updates
```

### Effects

```
dispose := signals.NewEffect(func() {
    count := counter.Get()  ◄─ Creates dependency
    fmt.Printf("Count: %d\n", count)
})

// When counter.Set(5):
counter changes → effect reruns → prints "Count: 5"

// Cleanup:
defer dispose()  // Removes subscription
```

## How Everything Connects

### Development Dependencies

```
Your Machine:
  │
  ├─ GoFlow Source (/path/to/GoFlow/)
  │    ├─ cmd/goflow/         ◄─ CLI tool source
  │    ├─ goflow/             ◄─ Framework package
  │    ├─ signals/            ◄─ Signals package
  │    └─ widgets/            ◄─ Widgets package
  │
  ├─ $GOPATH/bin/
  │    └─ goflow              ◄─ Compiled CLI binary
  │
  └─ Your Project (myapp/)
       ├─ lib/main.go         ◄─ Imports GoFlow packages
       ├─ macos/main.go       ◄─ Imports lib/
       └─ go.mod              ◄─ Points to GoFlow
            └─ replace directive to /path/to/GoFlow/
```

### Runtime Flow

```
1. User runs: cd myapp/macos && go run main.go
              ↓
2. Go compiles macos/main.go
              ↓
3. Resolves import: com.example/myapp/lib
              ↓
4. Resolves imports in lib/main.go:
   - github.com/base-go/GoFlow/pkg/core/framework
   - github.com/base-go/GoFlow/pkg/core/signals
   - github.com/base-go/GoFlow/pkg/core/widgets
              ↓
5. Uses replace directive in go.mod
   → Loads from /path/to/GoFlow/
              ↓
6. Compiles everything into single binary
              ↓
7. Executes: main.main() → lib.Run() → goflow.RunApp()
              ↓
8. Builds widget tree, creates elements, renders
```

## Summary

### Key Takeaways

1. **CLI Tool** (`goflow`) creates projects with proper structure
2. **Framework Packages** (`goflow/`, `signals/`, `widgets/`) provide APIs
3. **Platform Runners** bootstrap the app and call into `lib/`
4. **Shared Code** in `lib/` is cross-platform
5. **Three Trees** (Widget → Element → RenderObject) power the UI
6. **Signals** provide reactive state management
7. **Canvas Interface** abstracts rendering backend
8. **Native Backends** (future) render to actual pixels

### Complete Flow Summary

```
goflow create myapp
    ↓
Edit lib/main.go (your app)
    ↓
cd platform && go run main.go
    ↓
Platform runner → lib.Run() → goflow.RunApp()
    ↓
Widget Tree → Element Tree → RenderObject Tree
    ↓
Layout → Paint → Canvas
    ↓
(Future) Native Backend → GPU → Screen
```

## Next Steps

- [Getting Started Guide](./GETTING_STARTED.md) - Step-by-step tutorial
- [Architecture](../ARCHITECTURE.md) - Deep dive into three-tree system
- [Rendering Architecture](./RENDERING.md) - Native vs WGPU backends
- [CLI Reference](./CLI.md) - Complete CLI documentation

---

**Now you understand the complete GoFlow workflow!** 🎉
