# GoFlow Framework Architecture

## Overview

GoFlow is a Flutter-inspired GUI framework for Go with built-in reactive state management through signals. Unlike Flutter, GoFlow simplifies state management by using signals instead of separate StatelessWidget and StatefulWidget classes.

## Design Philosophy

1. **Signals-First**: All state management is handled through the signals package
2. **Single Widget Type**: No distinction between stateless and stateful widgets
3. **Declarative UI**: Describe what you want, not how to build it
4. **Composition**: Build complex UIs from simple, reusable widgets
5. **Reactive**: UI automatically updates when signals change

## Architecture Layers

```
┌──────────────────────────────────────┐
│      Application Layer                │
│  (User code: MyApp, CounterApp, etc)  │
├──────────────────────────────────────┤
│         Widget Layer                  │
│  (Text, Container, Column, etc)       │
├──────────────────────────────────────┤
│       Framework Core                  │
│  (Widget, Element, BuildContext)      │
├──────────────────────────────────────┤
│      Rendering Layer                  │
│  (RenderObject, Layout, Paint)        │
├──────────────────────────────────────┤
│      Signals Layer                    │
│  (Reactive State Management)          │
├──────────────────────────────────────┤
│    Platform Layer (Future)            │
│  (WGPU, Window Management)            │
└──────────────────────────────────────┘
```

## Core Concepts

### 1. Widgets

Widgets are immutable descriptions of part of the UI.

```go
type Widget interface {
    Build(context BuildContext) Widget
    CreateElement() Element
    GetKey() Key
}
```

**Key Points:**
- Widgets are configuration objects
- They describe **what** to show, not **how** to show it
- They are cheap to create and recreate
- A widget's Build() method returns another widget (or nil for leaf widgets)

### 2. Elements

Elements are mutable objects that manage the widget tree lifecycle.

```go
type Element interface {
    Mount(parent Element, slot interface{})
    Update(newWidget Widget)
    Unmount()
    Rebuild()
    GetRenderObject() RenderObject
    // ... more methods
}
```

**Key Points:**
- Elements live longer than widgets
- They manage state and the relationship between widgets and render objects
- One element per widget instance
- Elements form a tree parallel to the widget tree

### 3. RenderObjects

RenderObjects handle layout, painting, and hit testing.

```go
type RenderObject interface {
    Layout(constraints *Constraints)
    Paint(canvas Canvas, offset *Offset)
    HitTest(position *Offset) bool
    GetSize() *Size
    // ... more methods
}
```

**Key Points:**
- Do the actual layout calculations
- Perform painting to the canvas
- Handle hit testing for events
- Form a tree (subset of element tree)

### 4. Signals for State

Instead of StatefulWidget, use signals:

```go
type CounterApp struct {
    goflow.BaseWidget
    counter *signals.Signal[int]
}

func NewCounterApp() *CounterApp {
    return &CounterApp{
        counter: signals.New(0),
    }
}

func (a *CounterApp) Build(ctx goflow.BuildContext) goflow.Widget {
    count := a.counter.Get() // Reactive!
    return widgets.NewText(fmt.Sprintf("Count: %d", count))
}
```

## The Three Trees

GoFlow maintains three parallel trees:

### 1. Widget Tree
- **Immutable** configuration
- Created by calling Build() methods
- Cheap to recreate
- Example: `Text("Hello")` → `Container()` → `Column()`

### 2. Element Tree
- **Mutable** objects
- Manages widget lifecycle
- Lives longer than widgets
- One-to-one with widget instances
- Example: `TextElement` → `ContainerElement` → `ColumnElement`

### 3. Render Tree
- **Subset** of element tree
- Only elements with RenderObjects
- Performs layout and painting
- Example: `RenderText` → `RenderPadding` → `RenderColumn`

## Layout System

GoFlow uses Flutter's box protocol for layout:

### Constraints Go Down

Parent passes constraints to children:

```go
constraints := &Constraints{
    MinWidth:  0,
    MaxWidth:  800,
    MinHeight: 0,
    MaxHeight: 600,
}
child.Layout(constraints)
```

### Sizes Go Up

Children return their size to parents:

```go
func (r *RenderText) Layout(constraints *Constraints) {
    size := r.ComputeSize(constraints)
    r.SetSize(size)
}
```

### Parents Set Positions

After layout, parents position children:

```go
child.SetOffset(goflow.NewOffset(x, y))
```

## Rendering Pipeline

### 1. Build Phase

```
Widget.Build() → New Widget Tree → Element.Update()
```

- User code builds widget tree
- Framework updates element tree
- Minimal rebuilds (dirty tracking)

### 2. Layout Phase

```
RenderObject.Layout() → Constraint System → Sizes Computed
```

- Constraints flow down
- Sizes flow up
- Parent positions children

### 3. Paint Phase

```
RenderObject.Paint() → Canvas Commands → Visual Output
```

- Depth-first traversal
- Parent paints before children
- Respects clipping/transforms

## Example Flow

```go
// 1. User creates widget
app := &MyApp{}

// 2. RunApp builds element tree
goflow.RunApp(app)
    → app.CreateElement() → GenericElement
    → element.Mount(nil, nil)
    → element.Rebuild()
    → app.Build(ctx) → child widget
    → childElement.Mount(element, nil)
    → ... recursive build

// 3. Layout
    → findRenderObject(rootElement)
    → renderObject.Layout(constraints)
    → ... recursive layout

// 4. Paint
    → renderObject.Paint(canvas, offset)
    → ... recursive paint
```

## Signals Integration

### Reactive Rebuilds

```go
type TodoApp struct {
    goflow.BaseWidget
    todos *signals.SignalSlice[Todo]
}

func (a *TodoApp) Build(ctx goflow.BuildContext) goflow.Widget {
    // Access signal - creates dependency
    todoList := a.todos.Get()

    return widgets.NewColumn(
        // Build UI from signal data
        buildTodoWidgets(todoList),
    )
}
```

### Effect-Based Rebuilds (Future)

```go
// In a future version, elements could use effects:
dispose := signals.NewEffect(func() {
    element.MarkNeedsBuild()
})
```

## File Organization

```
GoFlow/
├── signals/              # Reactive state management
│   ├── signal.go
│   ├── computed.go
│   ├── effect.go
│   ├── batch.go
│   └── collections.go
├── goflow/               # Core framework
│   ├── widget.go         # Widget interface
│   ├── element.go        # Element system
│   ├── render_object.go  # RenderObject system
│   ├── constraints.go    # Layout constraints
│   ├── geometry.go       # Size, Offset, Rect
│   ├── canvas.go         # Painting abstraction
│   ├── key.go            # Widget keys
│   └── app.go            # RunApp entry point
├── widgets/              # Built-in widgets
│   ├── text.go
│   ├── container.go
│   ├── column.go
│   └── center.go
└── examples/             # Example apps
    ├── goflow-hello/
    └── goflow-counter/
```

## Comparison with Flutter

| Feature | Flutter | GoFlow |
|---------|---------|--------|
| State Management | StatelessWidget + StatefulWidget | Single Widget + Signals |
| Language | Dart | Go |
| Build Method | `Widget build(BuildContext)` | `Widget Build(BuildContext)` |
| State Updates | `setState(() => ...)` | `signal.Set(value)` |
| Reactive | Through setState | Through signals |
| Hot Reload | Yes | Planned |
| Platform | Mobile, Web, Desktop | Desktop (planned) |

## Key Differences

1. **No StatefulWidget**: Use signals instead
2. **Explicit reactivity**: Signal.Get() creates dependencies
3. **Go idioms**: Interfaces, composition, explicit errors
4. **Simpler**: One way to manage state (signals)
5. **Type-safe**: Go's static typing + generics for signals

## Future Work

1. **Platform Integration**
   - WGPU rendering backend
   - Window management (GLFW)
   - Event handling (mouse, keyboard)

2. **More Widgets**
   - Row (horizontal flex)
   - Stack (overlay)
   - Gesture detection
   - Input widgets

3. **Advanced Features**
   - Hot reload
   - Dev tools
   - Performance profiling
   - Animation system

4. **Optimizations**
   - Render object reuse
   - Layer caching
   - Incremental layout
   - Paint culling

## Learning Resources

- Flutter Architecture: https://flutter.dev/docs/resources/architectural-overview
- Signals Pattern: Preact Signals, Solid.js
- Box Layout: https://flutter.dev/docs/development/ui/layout

## Contributing

See CONTRIBUTING.md for development guidelines.

## License

MIT License - see LICENSE file
