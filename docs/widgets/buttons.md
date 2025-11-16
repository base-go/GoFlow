# Button Widgets

Buttons trigger actions when pressed.

## Button

Primary button with Material or Cupertino styling.

### Material Button

```go
button := material.NewButton(
    &widgets.Text{Data: "Click Me"},
    func() {
        fmt.Println("Button clicked!")
    },
)
```

### Cupertino Button

```go
button := cupertino.NewButton(
    &widgets.Text{Data: "Click Me"},
    func() {
        fmt.Println("Button clicked!")
    },
)
```

### Adaptive Button (Recommended)

```go
button := adaptive.NewButton("Click Me", func() {
    fmt.Println("Button clicked!")
})
```

### Button Styles

**Material:**
- `ButtonStyleFilled` - Filled with solid color (default)
- `ButtonStyleOutlined` - Outlined with border
- `ButtonStyleText` - Text only, no background

```go
// Filled button
filledBtn := material.NewButton(child, onPressed)
filledBtn.Style = material.ButtonStyleFilled

// Outlined button
outlinedBtn := material.NewOutlinedButton("Outlined", onPressed)

// Text button
textBtn := material.NewTextButton("Text", onPressed)
```

**Cupertino:**
- `ButtonTypeFilled` - Filled with color (default)
- `ButtonTypeGray` - Gray background
- `ButtonTypePlain` - Text only

```go
// Filled button
filledBtn := cupertino.NewButton(child, onPressed)

// Gray button
grayBtn := cupertino.NewGrayButton("Gray", onPressed)

// Text button
textBtn := cupertino.NewTextButton("Text", onPressed)
```

### Properties

**Material:**
- `OnPressed` - Callback when pressed
- `Child` - Button content
- `Style` - Button style
- `BackgroundColor` - Custom background color
- `ForegroundColor` - Text/icon color
- `DisabledColor` - Color when disabled
- `Elevation` - Shadow elevation
- `BorderRadius` - Corner radius
- `Padding` - Internal padding
- `MinWidth` / `MinHeight` - Minimum size
- `Enabled` - Whether button is enabled

**Cupertino:**
- `OnPressed` - Callback when pressed
- `Child` - Button content
- `Type` - Button type
- `Color` - Custom color
- `DisabledColor` - Color when disabled
- `BorderRadius` - Corner radius (default 8.0)
- `Padding` - Internal padding
- `MinSize` - Minimum size (default 44.0)
- `Enabled` - Whether button is enabled

### Example: Custom Styled Button

```go
button := material.NewButton(
    &widgets.Text{Data: "Custom Button"},
    onPressed,
)
button.BackgroundColor = goflow.NewColor(76, 175, 80, 255) // Green
button.ForegroundColor = goflow.NewColor(255, 255, 255, 255) // White
button.Elevation = 4.0
button.BorderRadius = 8.0
```

### Example: Disabled Button

```go
canSubmit := signals.New(false)

button := material.NewButton(
    &widgets.Text{Data: "Submit"},
    func() {
        if !canSubmit.Get() {
            return
        }
        submit()
    },
)
button.Enabled = canSubmit.Get()
```

## TextButton

Text-only button (Material).

```go
button := material.NewTextButton("Cancel", func() {
    fmt.Println("Cancelled")
})
```

Equivalent to:
```go
button := material.NewButton(&widgets.Text{Data: "Cancel"}, onPressed)
button.Style = material.ButtonStyleText
```

### Use Cases

- Secondary actions
- Navigation links
- Dialog actions
- Less prominent actions

### Example: Dialog Actions

```go
&widgets.Row{
    Children: []goflow.Widget{
        material.NewTextButton("Cancel", onCancel),
        material.NewTextButton("OK", onOK),
    },
    MainAxisAlignment: widgets.MainAxisEnd,
}
```

## OutlinedButton

Outlined button with border (Material).

```go
button := material.NewOutlinedButton("Learn More", func() {
    openLearnMore()
})
```

### Use Cases

- Medium emphasis actions
- Alternative to filled buttons
- Actions that need visibility but not dominance

### Example: Button Group

```go
&widgets.Row{
    Children: []goflow.Widget{
        material.NewButton(&widgets.Text{Data: "Primary"}, onPrimary),
        material.NewOutlinedButton("Secondary", onSecondary),
        material.NewTextButton("Tertiary", onTertiary),
    },
    MainAxisAlignment: widgets.MainAxisSpaceEvenly,
}
```

## IconButton

Button with icon.

```go
button := widgets.NewIconButton(
    widgets.IconFavorite,
    func() {
        toggleFavorite()
    },
)
button.Color = goflow.NewColor(244, 67, 54, 255) // Red
button.IconSize = 28.0
```

### Properties

- `Icon` - Icon to display
- `OnPressed` - Callback when pressed
- `IconSize` - Size of icon (default 24.0)
- `Color` - Icon color
- `Tooltip` - Tooltip text
- `Padding` - Padding around icon

### Example: App Bar Actions

```go
appBar := material.NewAppBar(&widgets.Text{Data: "Title"})
appBar.Actions = []goflow.Widget{
    widgets.NewIconButton(widgets.IconSearch, onSearch),
    widgets.NewIconButton(widgets.IconSettings, onSettings),
}
```

### Example: Like Button

```go
isLiked := signals.New(false)

likeButton := widgets.NewIconButton(
    func() widgets.IconData {
        if isLiked.Get() {
            return widgets.IconFavorite
        }
        return widgets.IconFavoriteBorder
    }(),
    func() {
        isLiked.Set(!isLiked.Get())
    },
)
likeButton.Color = func() *goflow.Color {
    if isLiked.Get() {
        return goflow.NewColor(244, 67, 54, 255) // Red
    }
    return nil
}()
```

## FloatingActionButton

Material Design FAB for primary actions.

```go
fab := material.NewFloatingActionButton(
    widgets.NewIcon(widgets.IconAdd),
    func() {
        createNew()
    },
)
```

### Properties

- `OnPressed` - Callback when pressed
- `Child` - FAB content (usually icon)
- `BackgroundColor` - Background color
- `ForegroundColor` - Icon color
- `Elevation` - Shadow elevation (default 6.0)
- `Mini` - Use smaller size
- `Tooltip` - Tooltip text

### Example: Mini FAB

```go
fab := material.NewFloatingActionButton(
    widgets.NewIcon(widgets.IconEdit),
    onEdit,
)
fab.Mini = true
fab.Elevation = 4.0
```

### Example: Custom FAB

```go
fab := material.NewFloatingActionButton(
    widgets.NewIcon(widgets.IconSend),
    sendMessage,
)
fab.BackgroundColor = goflow.NewColor(33, 150, 243, 255) // Blue
fab.ForegroundColor = goflow.NewColor(255, 255, 255, 255) // White
```

### FAB in Scaffold

```go
&material.Scaffold{
    AppBar: appBar,
    Body: body,
    FloatingActionButton: material.NewFloatingActionButton(
        widgets.NewIcon(widgets.IconAdd),
        onAdd,
    ),
}
```

## Button Best Practices

### 1. Use Appropriate Button Types

```go
// Primary action - filled button
material.NewButton(&widgets.Text{Data: "Save"}, onSave)

// Secondary action - outlined button
material.NewOutlinedButton("Preview", onPreview)

// Tertiary action - text button
material.NewTextButton("Cancel", onCancel)
```

### 2. Provide Feedback

```go
isLoading := signals.New(false)

button := material.NewButton(
    func() goflow.Widget {
        if isLoading.Get() {
            return &widgets.Row{
                Children: []goflow.Widget{
                    &widgets.ProgressIndicator{},
                    &widgets.Text{Data: "Loading..."},
                },
            }
        }
        return &widgets.Text{Data: "Submit"}
    }(),
    func() {
        isLoading.Set(true)
        // ... async operation
        isLoading.Set(false)
    },
)
button.Enabled = !isLoading.Get()
```

### 3. Use Icons for Clarity

```go
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewIcon(widgets.IconSave),
        &widgets.Text{Data: "Save"},
    },
}
```

### 4. Maintain Touch Targets

```go
// Material minimum: 88x36 (default)
// Cupertino minimum: 44x44 (default)

button := material.NewButton(child, onPressed)
// Already has appropriate minimum size

// For custom buttons, ensure minimum size
customButton := &widgets.Container{
    MinWidth: 88.0,
    MinHeight: 36.0,
    Child: buttonContent,
}
```

### 5. Group Related Actions

```go
&widgets.Row{
    Children: []goflow.Widget{
        material.NewOutlinedButton("Back", onBack),
        widgets.NewSpacer(),
        material.NewButton(&widgets.Text{Data: "Next"}, onNext),
    },
}
```

### 6. Disable When Appropriate

```go
formValid := signals.NewComputed(func() bool {
    return name.Get() != "" && email.Get() != ""
})

submitButton := material.NewButton(
    &widgets.Text{Data: "Submit"},
    onSubmit,
)
submitButton.Enabled = formValid.Get()
```

### 7. Use Consistent Styling

```go
// Define button styles
var (
    primaryButtonColor = goflow.NewColor(33, 150, 243, 255)
    secondaryButtonColor = goflow.NewColor(156, 39, 176, 255)
)

func createPrimaryButton(text string, onPressed func()) goflow.Widget {
    btn := material.NewButton(&widgets.Text{Data: text}, onPressed)
    btn.BackgroundColor = primaryButtonColor
    return btn
}
```
