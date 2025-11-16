# Layout Widgets

Layout widgets control how child widgets are positioned and sized.

## Column & Row

Arrange widgets vertically (Column) or horizontally (Row).

### Column

```go
&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Text{Data: "First"},
        &widgets.Text{Data: "Second"},
        &widgets.Text{Data: "Third"},
    },
    MainAxisAlign: widgets.MainAxisCenter,
    CrossAxisAlign: widgets.CrossAxisCenter,
}
```

### Row

```go
&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Text{Data: "Left"},
        &widgets.Text{Data: "Center"},
        &widgets.Text{Data: "Right"},
    },
    MainAxisAlignment: widgets.MainAxisSpaceBetween,
    CrossAxisAlignment: widgets.CrossAxisCenter,
}
```

### Properties

**MainAxisAlignment** (vertical for Column, horizontal for Row):
- `MainAxisStart` - Start of axis
- `MainAxisEnd` - End of axis
- `MainAxisCenter` - Center of axis
- `MainAxisSpaceBetween` - Space evenly between children
- `MainAxisSpaceAround` - Space around children
- `MainAxisSpaceEvenly` - Space evenly including edges

**CrossAxisAlignment** (horizontal for Column, vertical for Row):
- `CrossAxisStart` - Start of cross axis
- `CrossAxisEnd` - End of cross axis
- `CrossAxisCenter` - Center of cross axis
- `CrossAxisStretch` - Stretch to fill

### Example: Responsive Layout

```go
&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Expanded{
            Flex: 2,
            Child: &widgets.Container{
                Color: goflow.NewColor(255, 200, 200, 255),
                Child: &widgets.Text{Data: "2/3 width"},
            },
        },
        &widgets.Expanded{
            Flex: 1,
            Child: &widgets.Container{
                Color: goflow.NewColor(200, 255, 200, 255),
                Child: &widgets.Text{Data: "1/3 width"},
            },
        },
    },
}
```

## Stack

Layer widgets on top of each other.

```go
&widgets.Stack{
    Children: []goflow.Widget{
        // Background
        &widgets.Container{
            Width: floatPtr(200),
            Height: floatPtr(200),
            Color: goflow.NewColor(200, 200, 255, 255),
        },
        // Foreground (centered)
        &widgets.Align{
            Alignment: widgets.AlignmentCenter,
            Child: &widgets.Text{Data: "Centered"},
        },
        // Top-right corner
        &widgets.Align{
            Alignment: widgets.AlignmentTopRight,
            Child: widgets.NewIcon(widgets.IconClose),
        },
    },
    Fit: widgets.StackFitExpand,
}
```

### Properties

**Alignment**: Where non-positioned children are placed
- `StackAlignmentTopLeft`
- `StackAlignmentTopCenter`
- `StackAlignmentTopRight`
- `StackAlignmentCenterLeft`
- `StackAlignmentCenter`
- `StackAlignmentCenterRight`
- `StackAlignmentBottomLeft`
- `StackAlignmentBottomCenter`
- `StackAlignmentBottomRight`

**Fit**: How non-positioned children are sized
- `StackFitLoose` - Natural size
- `StackFitExpand` - Fill the stack
- `StackFitPassthrough` - Pass parent constraints

## Positioned

Position a child within a Stack precisely.

```go
&widgets.Stack{
    Children: []goflow.Widget{
        background,
        &widgets.Positioned{
            Left: floatPtr(10),
            Top: floatPtr(20),
            Child: &widgets.Text{Data: "10px from left, 20px from top"},
        },
        &widgets.Positioned{
            Right: floatPtr(10),
            Bottom: floatPtr(10),
            Width: floatPtr(100),
            Height: floatPtr(50),
            Child: badge,
        },
    },
}
```

### Properties

- `Left` - Distance from left edge
- `Top` - Distance from top edge
- `Right` - Distance from right edge
- `Bottom` - Distance from bottom edge
- `Width` - Width of child (optional)
- `Height` - Height of child (optional)

## Align

Align a child within available space.

```go
&widgets.Align{
    Alignment: widgets.AlignmentBottomRight,
    Child: &widgets.Text{Data: "Bottom Right"},
}
```

### Alignment Values

Predefined alignments:
- `AlignmentTopLeft`, `AlignmentTopCenter`, `AlignmentTopRight`
- `AlignmentCenterLeft`, `AlignmentCenter`, `AlignmentCenterRight`
- `AlignmentBottomLeft`, `AlignmentBottomCenter`, `AlignmentBottomRight`

Custom alignment:
```go
widgets.AlignmentValue{
    X: 0.5,  // -1.0 (left) to 1.0 (right)
    Y: -0.5, // -1.0 (top) to 1.0 (bottom)
}
```

### Properties

- `WidthFactor` - Multiply child width to get our width
- `HeightFactor` - Multiply child height to get our height

```go
// Container sized to 2x child size
&widgets.Align{
    WidthFactor: floatPtr(2.0),
    HeightFactor: floatPtr(2.0),
    Alignment: widgets.AlignmentCenter,
    Child: child,
}
```

## Container

Convenience widget for decoration, padding, margin, and sizing.

```go
&widgets.Container{
    Width: floatPtr(200),
    Height: floatPtr(100),
    Color: goflow.NewColor(100, 149, 237, 255),
    Padding: goflow.NewEdgeInsets(16, 12, 16, 12),
    Margin: goflow.NewEdgeInsets(8, 8, 8, 8),
    Child: &widgets.Text{Data: "Styled Container"},
}
```

### Properties

- `Width` - Fixed width
- `Height` - Fixed height
- `Color` - Background color
- `Padding` - Inner spacing
- `Margin` - Outer spacing
- `Child` - Child widget

### Example: Card-like Container

```go
&widgets.Container{
    Color: goflow.NewColor(255, 255, 255, 255),
    Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
    Margin: goflow.NewEdgeInsets(8, 8, 8, 8),
    Child: &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: "Title"},
            &widgets.Text{Data: "Subtitle"},
        },
    },
}
```

## Center

Center a child widget.

```go
&widgets.Center{
    Child: &widgets.Text{Data: "Centered Text"},
}
```

Equivalent to:
```go
&widgets.Align{
    Alignment: widgets.AlignmentCenter,
    Child: child,
}
```

## Padding

Add padding around a child.

```go
&widgets.Padding{
    Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
    Child: myWidget,
}
```

Or use Container:
```go
&widgets.Container{
    Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
    Child: myWidget,
}
```

### EdgeInsets

Create padding values:

```go
// Uniform padding
goflow.NewEdgeInsets(16, 16, 16, 16) // left, top, right, bottom

// Symmetric padding
goflow.NewEdgeInsets(16, 8, 16, 8) // horizontal: 16, vertical: 8

// Zero padding
goflow.ZeroEdgeInsets()
```

## SizedBox

Fixed size container (part of Container functionality).

```go
width := 100.0
height := 50.0
&widgets.Container{
    Width: &width,
    Height: &height,
    Child: myWidget,
}
```

### Common Use Cases

**Fixed height spacer:**
```go
height := 20.0
&widgets.Container{
    Height: &height,
}
```

**Constrain child size:**
```go
maxWidth := 300.0
&widgets.Container{
    Width: &maxWidth,
    Child: &widgets.Text{Data: "Constrained width"},
}
```

## Expanded & Flexible

Make children expand to fill available space in Row/Column.

### Expanded

Fill remaining space (tight fit):

```go
&widgets.Row{
    Children: []goflow.Widget{
        fixedWidget,
        &widgets.Expanded{
            Child: flexibleWidget, // Takes remaining space
        },
        anotherFixedWidget,
    },
}
```

### Flexible

Allow flexibility (loose fit):

```go
&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Flexible{
            Flex: 2,
            Fit: widgets.FlexFitLoose,
            Child: widget1, // Can be smaller than allocated space
        },
        &widgets.Flexible{
            Flex: 1,
            Fit: widgets.FlexFitTight,
            Child: widget2, // Must fill allocated space
        },
    },
}
```

### Properties

- `Flex` - Flex factor (default 1)
- `Fit` - How to fit child (`FlexFitLoose` or `FlexFitTight`)

### Example: Three Column Layout

```go
&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Expanded{
            Flex: 1,
            Child: leftColumn,
        },
        &widgets.Expanded{
            Flex: 2,
            Child: mainColumn,
        },
        &widgets.Expanded{
            Flex: 1,
            Child: rightColumn,
        },
    },
}
```

## Spacer

Empty space in flex layouts (Row/Column).

```go
&widgets.Row{
    Children: []goflow.Widget{
        leftWidget,
        widgets.NewSpacer(), // Pushes widgets apart
        rightWidget,
    },
}
```

Equivalent to:
```go
&widgets.Expanded{
    Child: &widgets.Container{}, // Empty container
}
```

### Example: Space Between Buttons

```go
&widgets.Row{
    Children: []goflow.Widget{
        material.NewButton(child1, onPressed1),
        widgets.NewSpacer(),
        material.NewButton(child2, onPressed2),
        widgets.NewSpacer(),
        material.NewButton(child3, onPressed3),
    },
}
```

## Layout Tips

### Centering

**Center horizontally in Column:**
```go
&widgets.Column{
    CrossAxisAlign: widgets.CrossAxisCenter,
    Children: children,
}
```

**Center vertically in Row:**
```go
&widgets.Row{
    CrossAxisAlignment: widgets.CrossAxisCenter,
    Children: children,
}
```

**Center both axes:**
```go
&widgets.Center{
    Child: myWidget,
}
```

### Spacing

**Space between items:**
```go
&widgets.Column{
    Children: []goflow.Widget{
        item1,
        &widgets.Container{Height: floatPtr(16)}, // 16px spacer
        item2,
    },
}
```

**Or use MainAxisAlignment:**
```go
&widgets.Column{
    MainAxisAlign: widgets.MainAxisSpaceBetween,
    Children: children,
}
```

### Overflow

If content overflows, wrap in `SingleChildScrollView`:
```go
&widgets.SingleChildScrollView{
    Child: &widgets.Column{
        Children: manyChildren,
    },
}
```
