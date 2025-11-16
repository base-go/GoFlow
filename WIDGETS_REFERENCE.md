# GoFlow Widgets Reference

Complete reference documentation for all GoFlow widgets.

## Table of Contents

- [Layout Widgets](#layout-widgets)
  - [Column](#column)
  - [Row](#row)
  - [Stack](#stack)
  - [Positioned](#positioned)
  - [Align](#align)
  - [Center](#center)
  - [Flexible & Expanded](#flexible--expanded)
  - [Spacer](#spacer)
- [Container & Decoration](#container--decoration)
  - [Container](#container)
  - [Padding](#padding)
  - [SizedBox](#sizedbox)
  - [ColoredBox](#coloredbox)
- [Display Widgets](#display-widgets)
  - [Text](#text)
  - [Icon](#icon)
  - [IconButton](#iconbutton)
  - [Image](#image)
- [Scrolling Widgets](#scrolling-widgets)
  - [ListView](#listview)
  - [ListViewBuilder](#listviewbuilder)
  - [SingleChildScrollView](#singlechildscrollview)
  - [GridView](#gridview)
- [Interaction Widgets](#interaction-widgets)
  - [GestureDetector](#gesturedetector)
  - [InkWell](#inkwell)
  - [Draggable & DragTarget](#draggable--dragtarget)

---

## Layout Widgets

### Column

Displays children vertically in a column.

**File:** `pkg/core/widgets/column.go`

**Properties:**
- `Children []goflow.Widget` - Child widgets to display
- `MainAxisAlign MainAxisAlignment` - Vertical alignment (default: MainAxisStart)
- `CrossAxisAlign CrossAxisAlignment` - Horizontal alignment (default: CrossAxisCenter)
- `MainAxisSize MainAxisSize` - How much vertical space to occupy (default: MainAxisSizeMax)

**MainAxisAlignment Values:**
- `MainAxisStart` - Children at the top
- `MainAxisEnd` - Children at the bottom
- `MainAxisCenter` - Children centered vertically
- `MainAxisSpaceBetween` - Space distributed between children
- `MainAxisSpaceAround` - Space around each child
- `MainAxisSpaceEvenly` - Even spacing including edges

**CrossAxisAlignment Values:**
- `CrossAxisStart` - Children aligned to the left
- `CrossAxisEnd` - Children aligned to the right
- `CrossAxisCenter` - Children centered horizontally
- `CrossAxisStretch` - Children stretched to full width

**Example:**
```go
&widgets.Column{
    Children: []goflow.Widget{
        widgets.NewText("Item 1"),
        widgets.NewText("Item 2"),
        widgets.NewText("Item 3"),
    },
    MainAxisAlign:  widgets.MainAxisCenter,
    CrossAxisAlign: widgets.CrossAxisStart,
}
```

---

### Row

Displays children horizontally in a row.

**File:** `pkg/core/widgets/row.go`

**Properties:**
- `Children []goflow.Widget` - Child widgets to display
- `MainAxisAlignment MainAxisAlignment` - Horizontal alignment (default: MainAxisStart)
- `CrossAxisAlignment CrossAxisAlignment` - Vertical alignment (default: CrossAxisCenter)
- `MainAxisSize MainAxisSize` - How much horizontal space to occupy (default: MainAxisSizeMax)

**Alignment values same as Column**

**Example:**
```go
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewIcon(widgets.IconHome),
        widgets.NewText("Home"),
    },
    MainAxisAlignment:  widgets.MainAxisStart,
    CrossAxisAlignment: widgets.CrossAxisCenter,
}
```

---

### Stack

Overlays children on top of each other (z-axis stacking).

**File:** `pkg/core/widgets/stack.go`

**Properties:**
- `Children []goflow.Widget` - Children to stack (later children appear on top)
- `Alignment StackAlignment` - How non-positioned children align (default: StackAlignmentTopLeft)
- `Fit StackFit` - How non-positioned children are sized (default: StackFitLoose)

**StackAlignment Values:**
- `StackAlignmentTopLeft`, `StackAlignmentTopCenter`, `StackAlignmentTopRight`
- `StackAlignmentCenterLeft`, `StackAlignmentCenter`, `StackAlignmentCenterRight`
- `StackAlignmentBottomLeft`, `StackAlignmentBottomCenter`, `StackAlignmentBottomRight`

**StackFit Values:**
- `StackFitLoose` - Children sized to their natural size
- `StackFitExpand` - Children expanded to fill the stack
- `StackFitPassthrough` - Pass parent constraints to children

**Example:**
```go
widgets.NewStack([]goflow.Widget{
    // Background
    &widgets.ColoredBox{
        Color: goflow.ColorBlue,
    },
    // Foreground
    widgets.NewCenter(
        widgets.NewText("Overlaid Text"),
    ),
})
```

---

### Positioned

Positions a child within a Stack at absolute coordinates.

**File:** `pkg/core/widgets/stack.go`

**Properties:**
- `Child goflow.Widget` - The child to position
- `Left *float64` - Distance from left edge (nil = not constrained)
- `Top *float64` - Distance from top edge
- `Right *float64` - Distance from right edge
- `Bottom *float64` - Distance from bottom edge
- `Width *float64` - Explicit width
- `Height *float64` - Explicit height

**Example:**
```go
widgets.NewStack([]goflow.Widget{
    widgets.NewPositioned(
        widgets.NewText("Top Right"),
    ).Top(10).Right(10),
})
```

---

### Align

Aligns a child within itself using precise coordinates.

**File:** `pkg/core/widgets/align.go`

**Properties:**
- `Child goflow.Widget` - The child to align
- `Alignment AlignmentValue` - Alignment coordinates (X: -1.0 to 1.0, Y: -1.0 to 1.0)
- `WidthFactor *float64` - If set, width = child.width * widthFactor
- `HeightFactor *float64` - If set, height = child.height * heightFactor

**Predefined Alignments:**
- `AlignmentTopLeft`, `AlignmentTopCenter`, `AlignmentTopRight`
- `AlignmentCenterLeft`, `AlignmentCenter`, `AlignmentCenterRight`
- `AlignmentBottomLeft`, `AlignmentBottomCenter`, `AlignmentBottomRight`

**Example:**
```go
widgets.NewAlign(
    widgets.AlignmentBottomRight,
    widgets.NewText("Bottom Right Corner"),
)
```

---

### Center

Centers its child both horizontally and vertically.

**File:** `pkg/core/widgets/center.go`

**Properties:**
- `Child goflow.Widget` - The child to center

**Example:**
```go
widgets.NewCenter(
    widgets.NewText("Centered Text"),
)
```

---

### Flexible & Expanded

Allow children in Row/Column to expand and fill available space.

**File:** `pkg/core/widgets/flexible.go`

**Flexible Properties:**
- `Child goflow.Widget` - The child widget
- `Flex int` - Flex factor (how much space relative to siblings, default: 1)
- `Fit FlexFit` - How to fit (FlexFitLoose or FlexFitTight)

**Expanded:**
Shorthand for Flexible with FlexFitTight (child must fill space).

**Example:**
```go
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewText("Fixed"),
        &widgets.Expanded{
            Child: widgets.NewText("Fills remaining space"),
        },
        widgets.NewText("Fixed"),
    },
}
```

---

### Spacer

Creates flexible empty space in Row/Column.

**File:** `pkg/core/widgets/flexible.go`

**Properties:**
- `Flex int` - Flex factor (default: 1)

**Example:**
```go
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewText("Left"),
        widgets.NewSpacer(), // Pushes items apart
        widgets.NewText("Right"),
    },
}
```

---

## Container & Decoration

### Container

Convenience widget combining padding, color, and sizing.

**File:** `pkg/core/widgets/container.go`

**Properties:**
- `Child goflow.Widget` - The child widget
- `Width *float64` - Explicit width
- `Height *float64` - Explicit height
- `Color *goflow.Color` - Background color
- `Padding *goflow.EdgeInsets` - Inner padding
- `Margin *goflow.EdgeInsets` - Outer margin (not yet implemented in render)

**Example:**
```go
&widgets.Container{
    Width:   float64Ptr(200),
    Height:  float64Ptr(100),
    Color:   goflow.ColorBlue,
    Padding: goflow.NewEdgeInsetsAll(16),
    Child:   widgets.NewText("Padded Box"),
}
```

---

### Padding

Adds padding around a child.

**File:** `pkg/core/widgets/container.go`

**Properties:**
- `Padding *goflow.EdgeInsets` - Padding amount
- `Child goflow.Widget` - The child widget

**EdgeInsets Constructors:**
- `NewEdgeInsetsAll(value float64)` - Same padding on all sides
- `NewEdgeInsetsSymmetric(horizontal, vertical float64)`
- `NewEdgeInsets(left, top, right, bottom float64)`

**Example:**
```go
widgets.NewPadding(
    goflow.NewEdgeInsetsAll(20),
    widgets.NewText("Padded Text"),
)
```

---

### SizedBox

Constrains child to a specific size.

**File:** `pkg/core/widgets/container.go`

**Properties:**
- `Width *float64` - Explicit width
- `Height *float64` - Explicit height
- `Child goflow.Widget` - The child widget

**Example:**
```go
&widgets.SizedBox{
    Width:  float64Ptr(100),
    Height: float64Ptr(50),
    Child:  widgets.NewText("Fixed Size"),
}
```

---

### ColoredBox

Paints a colored background behind its child.

**File:** `pkg/core/widgets/container.go`

**Properties:**
- `Color *goflow.Color` - Background color
- `Child goflow.Widget` - The child widget

**Example:**
```go
&widgets.ColoredBox{
    Color: goflow.NewColor(255, 200, 200, 255),
    Child: widgets.NewText("Pink Background"),
}
```

---

## Display Widgets

### Text

Displays a string of text with styling.

**File:** `pkg/core/widgets/text.go`

**Properties:**
- `Data string` - The text to display
- `Style *goflow.TextStyle` - Text styling

**TextStyle Properties:**
- `FontSize float64` - Font size (default: 14)
- `Color *goflow.Color` - Text color (default: black)
- `FontWeight FontWeight` - Font weight (Normal, Bold)

**Example:**
```go
widgets.NewTextWithStyle(
    "Hello World",
    &goflow.TextStyle{
        FontSize:   24,
        Color:      goflow.ColorBlue,
        FontWeight: goflow.FontWeightBold,
    },
)
```

---

### Icon

Displays an icon from Material Icons.

**File:** `pkg/core/widgets/icon.go`

**Properties:**
- `Icon IconData` - The icon to display
- `Size float64` - Icon size (default: 24)
- `Color *goflow.Color` - Icon color
- `SemanticLabel string` - Accessibility label

**Predefined Icons:**
- `IconHome`, `IconMenu`, `IconSearch`, `IconSettings`
- `IconFavorite`, `IconAdd`, `IconRemove`, `IconClose`
- `IconCheck`, `IconArrowBack`, `IconArrowForward`
- `IconEdit`, `IconDelete`, `IconShare`
- `IconPerson`, `IconEmail`, `IconPhone`
- `IconInfo`, `IconWarning`, `IconError`

**Example:**
```go
icon := widgets.NewIcon(widgets.IconHome)
icon.Size = 32
icon.Color = goflow.ColorBlue
```

---

### IconButton

A clickable button with an icon.

**File:** `pkg/core/widgets/icon.go`

**Properties:**
- `Icon IconData` - The icon to display
- `OnPressed func()` - Callback when pressed
- `IconSize float64` - Icon size (default: 24)
- `Color *goflow.Color` - Icon color
- `Tooltip string` - Tooltip text
- `Padding *goflow.EdgeInsets` - Padding (default: 8px all sides)

**Example:**
```go
widgets.NewIconButton(
    widgets.IconAdd,
    func() {
        fmt.Println("Add clicked!")
    },
)
```

---

### Image

Displays an image.

**File:** `pkg/core/widgets/image.go`

**Properties:**
- `Source ImageSource` - Image source (AssetImage, NetworkImage, FileImage)
- `Fit ImageFit` - How to inscribe the image (default: ImageFitContain)
- `Width *float64` - Explicit width
- `Height *float64` - Explicit height
- `SemanticLabel string` - Accessibility label

**ImageFit Values:**
- `ImageFitFill` - Fill box, may distort aspect ratio
- `ImageFitContain` - Fit within box, preserve aspect ratio
- `ImageFitCover` - Cover box, preserve aspect ratio
- `ImageFitFitWidth`, `ImageFitFitHeight` - Fit one dimension
- `ImageFitNone` - No scaling
- `ImageFitScaleDown` - Align within box, scale down if needed

**Example:**
```go
widgets.NewImage(&widgets.AssetImage{
    Path: "assets/logo.png",
})
```

---

## Scrolling Widgets

### ListView

A scrollable list of widgets.

**File:** `pkg/core/widgets/listview.go`

**Properties:**
- `Children []goflow.Widget` - List items
- `ScrollDirection Axis` - Scroll direction (default: AxisVertical)
- `Padding *goflow.EdgeInsets` - Padding around items
- `ShrinkWrap bool` - Size to content (default: false)
- `Physics ScrollPhysics` - Scroll behavior

**Example:**
```go
widgets.NewListView([]goflow.Widget{
    widgets.NewText("Item 1"),
    widgets.NewText("Item 2"),
    widgets.NewText("Item 3"),
})
```

---

### ListViewBuilder

Creates a ListView with a builder function.

**File:** `pkg/core/widgets/listview.go`

**Properties:**
- `ItemCount int` - Number of items
- `ItemBuilder func(BuildContext, int) Widget` - Function to build each item
- `ScrollDirection Axis` - Scroll direction
- `Padding *goflow.EdgeInsets` - Padding

**Example:**
```go
widgets.NewListViewBuilder(
    10,
    func(ctx goflow.BuildContext, index int) goflow.Widget {
        return widgets.NewText(fmt.Sprintf("Item %d", index))
    },
)
```

---

### SingleChildScrollView

Makes a single child scrollable.

**File:** `pkg/core/widgets/listview.go`

**Properties:**
- `Child goflow.Widget` - The child to make scrollable
- `ScrollDirection Axis` - Scroll direction (default: AxisVertical)
- `Padding *goflow.EdgeInsets` - Padding
- `Physics ScrollPhysics` - Scroll behavior

**Example:**
```go
widgets.NewSingleChildScrollView(
    &widgets.Column{
        Children: []goflow.Widget{
            // Many children...
        },
    },
)
```

---

### GridView

A scrollable 2D grid of widgets.

**File:** `pkg/core/widgets/listview.go`

**Properties:**
- `Children []goflow.Widget` - Grid items
- `CrossAxisCount int` - Number of columns
- `ChildAspectRatio float64` - Aspect ratio of each child (default: 1.0)
- `MainAxisSpacing float64` - Vertical spacing (default: 0)
- `CrossAxisSpacing float64` - Horizontal spacing (default: 0)
- `Padding *goflow.EdgeInsets` - Padding

**Example:**
```go
widgets.NewGridView(
    []goflow.Widget{
        widgets.NewText("1"),
        widgets.NewText("2"),
        widgets.NewText("3"),
        widgets.NewText("4"),
    },
    2, // 2 columns
)
```

---

## Interaction Widgets

### GestureDetector

Detects gestures on its child.

**File:** `pkg/core/widgets/gesture.go`

**Properties:**
- `Child goflow.Widget` - The child widget
- `OnTap func()` - Called on tap
- `OnDoubleTap func()` - Called on double tap
- `OnLongPress func()` - Called on long press
- `OnPanStart, OnPanUpdate, OnPanEnd` - Drag/pan callbacks
- `OnScaleStart, OnScaleUpdate, OnScaleEnd` - Scale/pinch callbacks
- `Behavior HitTestBehavior` - Hit test behavior

**Example:**
```go
widgets.NewGestureDetector(
    widgets.NewText("Tap Me"),
).OnTap = func() {
    fmt.Println("Tapped!")
}
```

---

### InkWell

Material Design ink splash effect with gestures.

**File:** `pkg/core/widgets/gesture.go`

**Properties:**
- `Child goflow.Widget` - The child widget
- `OnTap func()` - Tap callback
- `OnDoubleTap func()` - Double tap callback
- `OnLongPress func()` - Long press callback
- `SplashColor *goflow.Color` - Splash effect color
- `HighlightColor *goflow.Color` - Highlight color when pressed
- `BorderRadius float64` - Border radius for splash (default: 0)

**Example:**
```go
widgets.NewInkWell(
    widgets.NewText("Click Me"),
    func() {
        fmt.Println("InkWell tapped!")
    },
)
```

---

### Draggable & DragTarget

Enable drag-and-drop interactions.

**File:** `pkg/core/widgets/gesture.go`

**Draggable Properties:**
- `Child goflow.Widget` - Widget shown when not dragging
- `Feedback goflow.Widget` - Widget shown while dragging
- `ChildWhenDragging goflow.Widget` - Widget shown in place when dragging
- `Data interface{}` - Data to pass to DragTarget
- `OnDragStarted, OnDragEnd, OnDragCompleted, OnDraggableCanceled` - Callbacks
- `Axis *Axis` - Lock dragging to one axis
- `MaxSimultaneousDrags *int` - Maximum simultaneous drags

**DragTarget Properties:**
- `Builder func(BuildContext, []interface{}, []interface{}) Widget` - Builds target
- `OnWillAccept func(interface{}) bool` - Determine if data is acceptable
- `OnAccept func(interface{})` - Called when drag is accepted
- `OnLeave func(interface{})` - Called when drag leaves

**Example:**
```go
// Draggable
widgets.NewDraggable(
    widgets.NewText("Drag me"),
    "my-data",
)

// DragTarget
widgets.NewDragTarget(
    func(ctx goflow.BuildContext, candidates, rejected []interface{}) goflow.Widget {
        if len(candidates) > 0 {
            return widgets.NewText("Drop here!")
        }
        return widgets.NewText("Drag here")
    },
)
```

---

## Layout System Status

| Feature | Status | Notes |
|---------|--------|-------|
| Column | ✅ Complete | All alignment modes working |
| Row | ✅ Complete | All alignment modes working |
| Stack | ✅ Complete | All alignment modes working |
| Positioned | ⚠️ Basic | Exists but needs Stack integration |
| Flexible/Expanded | ⚠️ Partial | Defined but not used by Row/Column yet |
| MainAxis Alignment | ✅ Complete | 6 modes fully implemented |
| CrossAxis Alignment | ✅ Complete | 4 modes fully implemented |
| Intrinsic Dimensions | ❌ Not Implemented | Future feature |
| Baseline Alignment | ❌ Not Implemented | Future feature |

---

## Common Patterns

### Responsive Layout
```go
&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Expanded{
            Flex: 2,
            Child: widgets.NewText("Takes 2/3"),
        },
        &widgets.Expanded{
            Flex: 1,
            Child: widgets.NewText("Takes 1/3"),
        },
    },
}
```

### Card Layout
```go
&widgets.Container{
    Color:   goflow.ColorWhite,
    Padding: goflow.NewEdgeInsetsAll(16),
    Child: &widgets.Column{
        Children: []goflow.Widget{
            widgets.NewTextWithStyle("Title", titleStyle),
            spacer(8),
            widgets.NewText("Content goes here"),
        },
        MainAxisAlign:  widgets.MainAxisStart,
        CrossAxisAlign: widgets.CrossAxisStart,
    },
}
```

### Button
```go
&widgets.InkWell{
    Child: &widgets.Container{
        Padding: goflow.NewEdgeInsetsSymmetric(24, 12),
        Color:   goflow.ColorBlue,
        Child:   widgets.NewText("Click Me"),
    },
    OnTap: func() {
        fmt.Println("Button pressed!")
    },
}
```

### Overlay
```go
widgets.NewStack([]goflow.Widget{
    // Background
    widgets.NewImage(backgroundImage),
    // Overlay
    &widgets.ColoredBox{
        Color: goflow.NewColor(0, 0, 0, 128), // Semi-transparent black
        Child: widgets.NewCenter(
            widgets.NewText("Overlaid Text"),
        ),
    },
})
```

---

## Next Steps

Future widget implementations:
- Material Design components (AppBar, Card, Drawer, etc.)
- Cupertino (iOS-style) widgets
- Form widgets (TextField, Checkbox, Radio, Switch, Slider)
- Animation widgets (AnimatedContainer, FadeTransition, etc.)
- Custom painting widgets (CustomPaint)
- Sliver widgets for advanced scrolling

---

**Total Widgets Documented:** 27 widgets across 13 files

This reference covers all currently implemented widgets in the GoFlow framework.
