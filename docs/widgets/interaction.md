# Interaction Widgets

Widgets that detect and respond to user gestures.

## GestureDetector

Detects various gestures (tap, drag, scale) on its child widget.

```go
&widgets.GestureDetector{
    Child: myWidget,
    OnTap: func() {
        fmt.Println("Tapped!")
    },
}
```

### Properties

**Child:**
- `Child` - The widget to detect gestures on

**Tap Gestures:**
- `OnTap` - Single tap callback
- `OnDoubleTap` - Double tap callback
- `OnLongPress` - Long press callback

**Drag Gestures:**
- `OnPanStart` - Drag started
- `OnPanUpdate` - Drag position updated
- `OnPanEnd` - Drag ended

**Scale Gestures:**
- `OnScaleStart` - Scale/rotate started
- `OnScaleUpdate` - Scale/rotate updated
- `OnScaleEnd` - Scale/rotate ended

**Other:**
- `Behavior` - Hit test behavior

### Basic Tap Detection

```go
detector := widgets.NewGestureDetector(&widgets.Container{
    Width: floatPtr(100.0),
    Height: floatPtr(100.0),
    Color: goflow.NewColor(33, 150, 243, 255),
})

detector.OnTap = func() {
    fmt.Println("Container tapped!")
}
```

### Multiple Gesture Types

```go
&widgets.GestureDetector{
    Child: myWidget,
    OnTap: func() {
        handleTap()
    },
    OnDoubleTap: func() {
        handleDoubleTap()
    },
    OnLongPress: func() {
        handleLongPress()
    },
}
```

### Drag Detection

```go
&widgets.GestureDetector{
    Child: draggableBox,
    OnPanStart: func(details widgets.DragStartDetails) {
        fmt.Printf("Drag started at: %v\n", details.LocalPosition)
    },
    OnPanUpdate: func(details widgets.DragUpdateDetails) {
        fmt.Printf("Dragging: delta=%v\n", details.Delta)
    },
    OnPanEnd: func(details widgets.DragEndDetails) {
        fmt.Printf("Drag ended with velocity: %v\n", details.Velocity)
    },
}
```

### Example: Interactive Box

```go
position := signals.New(goflow.NewOffset(0, 0))

detector := widgets.NewGestureDetector(&widgets.Container{
    Width: floatPtr(100.0),
    Height: floatPtr(100.0),
    Color: goflow.NewColor(33, 150, 243, 255),
})

detector.OnPanUpdate = func(details widgets.DragUpdateDetails) {
    currentPos := position.Get()
    newPos := goflow.NewOffset(
        currentPos.X + details.Delta.X,
        currentPos.Y + details.Delta.Y,
    )
    position.Set(newPos)
}
```

### Example: Like Button

```go
isLiked := signals.New(false)

icon := widgets.NewIcon(func() widgets.IconData {
    if isLiked.Get() {
        return widgets.IconFavorite
    }
    return widgets.IconFavoriteBorder
}())

detector := widgets.NewGestureDetector(icon)
detector.OnTap = func() {
    isLiked.Set(!isLiked.Get())
}

detector.OnDoubleTap = func() {
    // Double tap for super like
    isLiked.Set(true)
    showAnimation()
}
```

### Example: Pinch to Zoom

```go
scale := signals.New(1.0)

detector := widgets.NewGestureDetector(image)

detector.OnScaleStart = func(details widgets.ScaleStartDetails) {
    fmt.Println("Scaling started")
}

detector.OnScaleUpdate = func(details widgets.ScaleUpdateDetails) {
    scale.Set(details.Scale)
}

detector.OnScaleEnd = func(details widgets.ScaleEndDetails) {
    fmt.Println("Scaling ended")
}
```

### Hit Test Behavior

```go
// Defer to child - only responds if child would
detector.Behavior = widgets.HitTestBehaviorDeferToChild

// Opaque - always responds, even if child wouldn't
detector.Behavior = widgets.HitTestBehaviorOpaque

// Translucent - responds, but lets events pass through
detector.Behavior = widgets.HitTestBehaviorTranslucent
```

## InkWell

Material Design ink splash effect with gesture detection.

```go
&widgets.InkWell{
    Child: myWidget,
    OnTap: func() {
        handleTap()
    },
}
```

### Properties

- `Child` - The widget to wrap
- `OnTap` - Tap callback
- `OnDoubleTap` - Double tap callback
- `OnLongPress` - Long press callback
- `SplashColor` - Color of splash effect
- `HighlightColor` - Color when pressed
- `BorderRadius` - Radius for splash shape

### Basic InkWell

```go
inkwell := widgets.NewInkWell(
    &widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
        Child: &widgets.Text{Data: "Click Me"},
    },
    func() {
        fmt.Println("InkWell tapped!")
    },
)
```

### Styled InkWell

```go
inkwell := widgets.NewInkWell(content, onTap)
inkwell.SplashColor = goflow.NewColor(33, 150, 243, 100) // Blue with alpha
inkwell.HighlightColor = goflow.NewColor(33, 150, 243, 50)
inkwell.BorderRadius = 8.0
```

### Example: List Item

```go
func buildListItem(title string, onTap func()) goflow.Widget {
    return widgets.NewInkWell(
        &widgets.Container{
            Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
            Child: &widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconPerson),
                    &widgets.Expanded{
                        Child: &widgets.Text{Data: title},
                    },
                    widgets.NewIcon(widgets.IconArrowForward),
                },
            },
        },
        onTap,
    )
}
```

### Example: Icon Button with Splash

```go
func buildIconButton(icon widgets.IconData, onTap func()) goflow.Widget {
    inkwell := widgets.NewInkWell(
        &widgets.Container{
            Width: floatPtr(48.0),
            Height: floatPtr(48.0),
            Child: &widgets.Center{
                Child: widgets.NewIcon(icon),
            },
        },
        onTap,
    )
    inkwell.BorderRadius = 24.0 // Circular splash

    return inkwell
}
```

### Example: Card with Ripple

```go
func buildClickableCard(content goflow.Widget, onTap func()) goflow.Widget {
    inkwell := widgets.NewInkWell(content, onTap)
    inkwell.BorderRadius = 8.0

    return material.NewCard(inkwell)
}
```

## Draggable

Makes a widget draggable with drag-and-drop support.

```go
&widgets.Draggable{
    Child: myWidget,
    Feedback: dragPreview,
    Data: "my-data",
}
```

### Properties

- `Child` - Widget shown when not dragging
- `Feedback` - Widget shown while dragging (drag preview)
- `ChildWhenDragging` - Widget shown in place during drag
- `Data` - Data passed to DragTarget
- `OnDragStarted` - Called when drag starts
- `OnDragEnd` - Called when drag ends
- `OnDragCompleted` - Called when successfully dropped
- `OnDraggableCanceled` - Called when drag canceled
- `Axis` - Lock dragging to axis
- `MaxSimultaneousDrags` - Limit concurrent drags

### Basic Draggable

```go
draggable := widgets.NewDraggable(
    &widgets.Container{
        Width: floatPtr(100.0),
        Height: floatPtr(100.0),
        Color: goflow.NewColor(33, 150, 243, 255),
        Child: &widgets.Text{Data: "Drag me"},
    },
    "item-1",
)
```

### Draggable with Feedback

```go
draggable := widgets.NewDraggable(
    buildNormalWidget(),
    "my-data",
)

// Custom feedback widget (shown while dragging)
draggable.Feedback = &widgets.Container{
    Width: floatPtr(100.0),
    Height: floatPtr(100.0),
    Color: goflow.NewColor(33, 150, 243, 150), // Semi-transparent
    Child: &widgets.Text{Data: "Dragging..."},
}

// Widget shown in original position during drag
draggable.ChildWhenDragging = &widgets.Container{
    Width: floatPtr(100.0),
    Height: floatPtr(100.0),
    Color: goflow.NewColor(200, 200, 200, 255), // Gray placeholder
}
```

### Draggable with Callbacks

```go
draggable := widgets.NewDraggable(child, data)

draggable.OnDragStarted = func() {
    fmt.Println("Started dragging")
}

draggable.OnDragCompleted = func() {
    fmt.Println("Successfully dropped")
}

draggable.OnDraggableCanceled = func() {
    fmt.Println("Drag canceled")
}

draggable.OnDragEnd = func(details widgets.DraggableDetails) {
    fmt.Printf("Drag ended at: %v with velocity: %v\n",
        details.Offset, details.Velocity)
}
```

### Example: Draggable List Item

```go
type Item struct {
    ID   string
    Name string
}

func buildDraggableItem(item Item) goflow.Widget {
    widget := material.NewCard(&widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
        Child: &widgets.Text{Data: item.Name},
    })

    draggable := widgets.NewDraggable(widget, item.ID)

    draggable.Feedback = &widgets.Container{
        Width: floatPtr(200.0),
        Color: goflow.NewColor(33, 150, 243, 200),
        Child: &widgets.Text{Data: item.Name},
    }

    draggable.ChildWhenDragging = &widgets.Container{
        Height: floatPtr(50.0),
        Color: goflow.NewColor(224, 224, 224, 255),
    }

    return draggable
}
```

## DragTarget

Widget that accepts draggable widgets.

```go
&widgets.DragTarget{
    Builder: func(ctx goflow.BuildContext, candidateData, rejectedData []interface{}) goflow.Widget {
        return buildTargetWidget(candidateData)
    },
    OnWillAccept: func(data interface{}) bool {
        return data != nil
    },
    OnAccept: func(data interface{}) {
        handleDrop(data)
    },
}
```

### Properties

- `Builder` - Function to build widget based on drag state
- `OnWillAccept` - Predicate to determine if data is acceptable
- `OnAccept` - Called when acceptable data is dropped
- `OnLeave` - Called when draggable leaves the target

### Basic DragTarget

```go
target := widgets.NewDragTarget(
    func(ctx goflow.BuildContext, candidateData, rejectedData []interface{}) goflow.Widget {
        // Build different UI based on drag state
        if len(candidateData) > 0 {
            return buildHighlightedTarget()
        }
        return buildNormalTarget()
    },
)

target.OnAccept = func(data interface{}) {
    fmt.Printf("Accepted: %v\n", data)
}
```

### DragTarget with Validation

```go
target := widgets.NewDragTarget(builder)

target.OnWillAccept = func(data interface{}) bool {
    // Only accept specific type of data
    itemID, ok := data.(string)
    return ok && itemID != ""
}

target.OnAccept = func(data interface{}) {
    if itemID, ok := data.(string); ok {
        addItemToTarget(itemID)
    }
}

target.OnLeave = func(data interface{}) {
    clearHighlight()
}
```

### Example: Trash Can

```go
func buildTrashCan(onDelete func(string)) goflow.Widget {
    target := widgets.NewDragTarget(
        func(ctx goflow.BuildContext, candidateData, rejectedData []interface{}) goflow.Widget {
            isHovering := len(candidateData) > 0

            color := goflow.NewColor(244, 67, 54, 255) // Red
            if isHovering {
                color = goflow.NewColor(198, 40, 40, 255) // Darker red
            }

            return &widgets.Container{
                Width: floatPtr(100.0),
                Height: floatPtr(100.0),
                Color: color,
                Child: &widgets.Center{
                    Child: widgets.NewIcon(widgets.IconDelete),
                },
            }
        },
    )

    target.OnWillAccept = func(data interface{}) bool {
        _, ok := data.(string)
        return ok
    }

    target.OnAccept = func(data interface{}) {
        if id, ok := data.(string); ok {
            onDelete(id)
        }
    }

    return target
}
```

### Example: Reorderable List

```go
items := signals.New([]string{"Item 1", "Item 2", "Item 3"})

func buildReorderableList() goflow.Widget {
    children := make([]goflow.Widget, 0)

    for i, item := range items.Get() {
        index := i

        // Draggable item
        draggable := widgets.NewDraggable(
            buildListItem(item),
            index,
        )

        // Drop target
        target := widgets.NewDragTarget(
            func(ctx goflow.BuildContext, candidateData, rejectedData []interface{}) goflow.Widget {
                return draggable
            },
        )

        target.OnAccept = func(data interface{}) {
            if sourceIndex, ok := data.(int); ok {
                reorderItems(sourceIndex, index)
            }
        }

        children = append(children, target)
    }

    return &widgets.ListView{
        Children: children,
    }
}

func reorderItems(from, to int) {
    currentItems := items.Get()
    item := currentItems[from]

    // Remove from old position
    newItems := append(currentItems[:from], currentItems[from+1:]...)

    // Insert at new position
    newItems = append(newItems[:to], append([]string{item}, newItems[to:]...)...)

    items.Set(newItems)
}
```

## Best Practices

### 1. Use InkWell for Material Interactions

```go
// ✅ Good - Material Design ripple
widgets.NewInkWell(content, onTap)

// ❌ Less ideal - no visual feedback
widgets.NewGestureDetector(content).OnTap = onTap
```

### 2. Provide Visual Feedback

```go
// ✅ Good - shows state changes
isPressed := signals.New(false)

detector := widgets.NewGestureDetector(buildButton())
detector.OnPanStart = func(details widgets.DragStartDetails) {
    isPressed.Set(true)
}
detector.OnPanEnd = func(details widgets.DragEndDetails) {
    isPressed.Set(false)
}
```

### 3. Use Appropriate Hit Test Behavior

```go
// For clickable areas larger than child
detector.Behavior = widgets.HitTestBehaviorOpaque

// For stacked interactive elements
detector.Behavior = widgets.HitTestBehaviorTranslucent
```

### 4. Provide Good Drag Feedback

```go
draggable := widgets.NewDraggable(child, data)

// ✅ Good - clear, semi-transparent feedback
draggable.Feedback = &widgets.Opacity{
    Opacity: 0.7,
    Child: child,
}

// ✅ Also good - placeholder during drag
draggable.ChildWhenDragging = buildPlaceholder()
```

### 5. Validate Drop Targets

```go
target := widgets.NewDragTarget(builder)

// ✅ Good - validates data type and content
target.OnWillAccept = func(data interface{}) bool {
    item, ok := data.(Item)
    return ok && item.IsValid()
}
```

### 6. Handle Gesture Conflicts

```go
// Don't use conflicting gestures on same widget
detector := widgets.NewGestureDetector(child)

// ✅ Good - complementary gestures
detector.OnTap = handleTap
detector.OnLongPress = handleLongPress

// ❌ Avoid - may conflict
detector.OnTap = handleTap
detector.OnDoubleTap = handleDoubleTap // Delays tap detection
```

### 7. Limit Draggable Count

```go
maxDrags := 1
draggable := widgets.NewDraggable(child, data)
draggable.MaxSimultaneousDrags = &maxDrags
```

## Helper Functions

```go
// Float pointer helper
func floatPtr(f float64) *float64 {
    return &f
}

// Build tappable container
func buildTappableBox(color *goflow.Color, size float64, onTap func()) goflow.Widget {
    return widgets.NewInkWell(
        &widgets.Container{
            Width: floatPtr(size),
            Height: floatPtr(size),
            Color: color,
        },
        onTap,
    )
}

// Build drag handle
func buildDragHandle() goflow.Widget {
    return &widgets.Container{
        Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
        Child: widgets.NewIcon(widgets.IconMenu),
    }
}
```
