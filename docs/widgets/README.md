# GoFlow Widgets Documentation

Complete reference for all GoFlow widgets with examples and best practices.

## Widget Categories

### Layout Widgets
- [Column & Row](layout.md#column--row) - Vertical and horizontal layouts
- [Stack](layout.md#stack) - Layered widgets
- [Positioned](layout.md#positioned) - Position children within Stack
- [Align](layout.md#align) - Align child within parent
- [Container](layout.md#container) - Decoration, padding, margin, sizing
- [Center](layout.md#center) - Center child widget
- [Padding](layout.md#padding) - Add padding around child
- [SizedBox](layout.md#sizedbox) - Fixed size container
- [Expanded & Flexible](layout.md#expanded--flexible) - Flex children in Row/Column
- [Spacer](layout.md#spacer) - Empty space in flex layouts

### Form Widgets
- [TextField](forms.md#textfield) - Material and Cupertino text input
- [Checkbox](forms.md#checkbox) - Material checkbox
- [Radio](forms.md#radio) - Material radio button
- [Switch](forms.md#switch) - Material and Cupertino toggle switch
- [Slider](forms.md#slider) - Material and Cupertino value slider

### Button Widgets
- [Button](buttons.md#button) - Material and Cupertino primary buttons
- [TextButton](buttons.md#textbutton) - Text-only button (Material)
- [OutlinedButton](buttons.md#outlinedbutton) - Outlined button (Material)
- [IconButton](buttons.md#iconbutton) - Button with icon
- [FloatingActionButton](buttons.md#floatingactionbutton) - Material FAB

### Display Widgets
- [Text](display.md#text) - Display text
- [Icon](display.md#icon) - Display icons
- [Image](display.md#image) - Display images

### Scrolling Widgets
- [ListView](scrolling.md#listview) - Scrollable list
- [ListView.builder](scrolling.md#listviewbuilder) - Lazy-loaded list
- [GridView](scrolling.md#gridview) - Scrollable grid
- [SingleChildScrollView](scrolling.md#singlechildscrollview) - Scrollable single child

### Interaction Widgets
- [GestureDetector](interaction.md#gesturedetector) - Detect gestures
- [InkWell](interaction.md#inkwell) - Material ink splash effect
- [Draggable](interaction.md#draggable) - Make widget draggable
- [DragTarget](interaction.md#dragtarget) - Accept draggable widgets

### App Structure
- [Scaffold](app-structure.md#scaffold) - Material app structure
- [CupertinoPageScaffold](app-structure.md#cupertinopagescaffold) - iOS app structure
- [CupertinoTabScaffold](app-structure.md#cupertinotabscaffold) - iOS tabbed interface
- [AppBar](app-structure.md#appbar) - Material top app bar
- [CupertinoNavigationBar](app-structure.md#cupertinonavigationbar) - iOS navigation bar
- [Drawer](app-structure.md#drawer) - Side navigation drawer
- [BottomNavigationBar](app-structure.md#bottomnavigationbar) - Bottom navigation

### Material-Specific Widgets
- [Card](material.md#card) - Material card
- [ListTile](material.md#listtile) - List item with leading/trailing
- [Dialog](material.md#dialog) - Material dialog
- [AlertDialog](material.md#alertdialog) - Alert dialog with actions
- [DrawerHeader](material.md#drawerheader) - Drawer header

### Cupertino-Specific Widgets
- [CupertinoTextField](cupertino.md#cupertinotextfield) - iOS text field
- [CupertinoSwitch](cupertino.md#cupertinoswitch) - iOS switch
- [CupertinoSlider](cupertino.md#cupertinoslider) - iOS slider

## Quick Reference

### Creating Widgets

```go
import (
    "github.com/base-go/GoFlow/widgets"
    "github.com/base-go/GoFlow/material"
    "github.com/base-go/GoFlow/cupertino"
    "github.com/base-go/GoFlow/adaptive"
)

// Layout
column := &widgets.Column{Children: []goflow.Widget{...}}
row := &widgets.Row{Children: []goflow.Widget{...}}
container := &widgets.Container{Child: child, Padding: padding}

// Material
button := material.NewButton(child, onPressed)
card := material.NewCard(child)

// Cupertino
button := cupertino.NewButton(child, onPressed)
card := cupertino.NewCard(child)

// Adaptive (recommended for cross-platform)
button := adaptive.NewButton("Text", onPressed)
card := adaptive.NewCard(child)
```

### Common Patterns

**Center a widget:**
```go
&widgets.Center{
    Child: myWidget,
}
```

**Add padding:**
```go
&widgets.Container{
    Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
    Child: myWidget,
}
```

**Create a list:**
```go
&widgets.ListView{
    Children: []goflow.Widget{
        item1,
        item2,
        item3,
    },
}
```

**Responsive layout:**
```go
&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Expanded{Child: widget1},
        &widgets.Expanded{Child: widget2},
    },
}
```

## Widget Lifecycle

All widgets in GoFlow follow this lifecycle:

1. **Construction** - Widget is created with `New...()` or struct literal
2. **Build** - `Build(context)` is called to create the widget tree
3. **Mount** - Element is created and mounted to parent
4. **Update** - When properties change, widget is rebuilt
5. **Unmount** - Widget is removed from tree

## Best Practices

### 1. Use Adaptive Widgets for Cross-Platform Apps

```go
// ✅ Good - works on all platforms
adaptive.NewButton("Click Me", onPressed)

// ❌ Less ideal - platform-specific
material.NewButton(child, onPressed)
```

### 2. Prefer Stateless Widgets with Signals

```go
// ✅ Good - reactive with signals
counter := signals.New(0)
&widgets.Text{
    Data: fmt.Sprintf("Count: %d", counter.Get()),
}

// ❌ Less ideal - requires manual updates
```

### 3. Use const for Repeated Values

```go
const (
    padding = 16.0
    spacing = 8.0
)

&widgets.Container{
    Padding: goflow.NewEdgeInsets(padding, padding, padding, padding),
    Child: child,
}
```

### 4. Extract Complex Widgets into Functions

```go
func buildHeader(title string) goflow.Widget {
    return &widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
        Child: &widgets.Text{Data: title},
    }
}
```

### 5. Use ListView.builder for Large Lists

```go
// ✅ Good - efficient for large lists
&widgets.ListViewBuilder{
    ItemCount: 1000,
    ItemBuilder: func(ctx goflow.BuildContext, i int) goflow.Widget {
        return &widgets.Text{Data: fmt.Sprintf("Item %d", i)}
    },
}

// ❌ Avoid - creates all items upfront
&widgets.ListView{
    Children: make1000Items(),
}
```

## Widget Composition

GoFlow widgets are highly composable. Build complex UIs by combining simple widgets:

```go
func buildProfileCard(name, email string) goflow.Widget {
    return material.NewCard(&widgets.Column{
        Children: []goflow.Widget{
            &widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconPerson),
                    &widgets.Expanded{
                        Child: &widgets.Text{Data: name},
                    },
                },
            },
            &widgets.Text{Data: email},
        },
    })
}
```

## Platform-Specific Behavior

Some widgets behave differently on different platforms:

| Widget | Android/Linux/Web | iOS/macOS |
|--------|------------------|-----------|
| Button | Material Design | Cupertino |
| TextField | Underline | Rounded rectangle |
| Switch | Material toggle | iOS toggle |
| AppBar | Material AppBar | NavigationBar |

Use `adaptive` package to handle these differences automatically.

## Next Steps

- Explore widget categories in detail
- Check out [examples](../../examples/) for real-world usage
- Read [DESIGN_SYSTEMS.md](../../DESIGN_SYSTEMS.md) for design system guide
