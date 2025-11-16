# Adaptive Package

Platform-adaptive widgets for GoFlow.

## Overview

The adaptive package provides widgets that automatically adapt to the platform's native design language. Write your UI once, and it will automatically use:

- **Cupertino** (iOS-style) on iOS and macOS
- **Material Design** on Android, Linux, Windows, and Web

## Why Adaptive Widgets?

- **Native Feel**: Your app looks and feels native on each platform
- **Write Once**: Single codebase automatically adapts
- **Best Practices**: Follows platform conventions automatically
- **Consistent API**: Same API regardless of underlying design system

## Widgets

### Button

Adaptive button that becomes Material or Cupertino based on platform:

```go
// Basic button
btn := adaptive.NewButton("Click Me", func() {
    fmt.Println("Clicked!")
})

// Text button
textBtn := adaptive.NewTextButton("Text Button", onPressed)

// Custom styling
btn := &adaptive.Button{
    Text: "Custom",
    OnPressed: onPressed,
    Style: adaptive.ButtonStyleFilled,
    Color: customColor,
    Enabled: true,
}
```

### Card

Adaptive card widget:

```go
card := adaptive.NewCard(content)
card.Padding = goflow.NewEdgeInsets(16, 16, 16, 16)
card.Color = customColor
```

### AppBar

Adaptive app bar (Material AppBar or Cupertino NavigationBar):

```go
appBar := adaptive.NewAppBar(&widgets.Text{Content: "Title"})
appBar.Leading = backButton
appBar.Trailing = []goflow.Widget{actionButton}
```

## Platform Detection

Widgets automatically detect the platform:

```go
platform := goflow.GetPlatform()

// Check platform type
if platform.IsApple() {
    // Will use Cupertino widgets
}

if platform.IsMobile() {
    // iOS or Android
}

// Get default theme
theme := platform.DefaultTheme() // "material" or "cupertino"
```

## Button Styles

Adaptive buttons support multiple styles:

- `ButtonStylePlatformDefault`: Platform's default style
- `ButtonStyleFilled`: Filled with solid color
- `ButtonStyleOutlined`: Outlined (Material only, becomes gray on iOS)
- `ButtonStyleText`: Text-only button

```go
btn := &adaptive.Button{
    Text: "Outlined",
    Style: adaptive.ButtonStyleOutlined,
    OnPressed: onPressed,
}
```

## Testing Different Platforms

Override the platform for testing:

```go
// Test iOS appearance
goflow.SetPlatform(goflow.PlatformIOS)

// Test Android appearance
goflow.SetPlatform(goflow.PlatformAndroid)

// Test macOS appearance
goflow.SetPlatform(goflow.PlatformMacOS)
```

## How It Works

Adaptive widgets use the Build pattern to conditionally render platform-specific widgets:

```go
func (b *Button) Build(context goflow.BuildContext) goflow.Widget {
    platform := goflow.GetPlatform()

    if platform.IsApple() {
        return cupertino.NewButton(b.Child, b.OnPressed)
    } else {
        return material.NewButton(b.Child, b.OnPressed)
    }
}
```

## Platform Mapping

| Platform | OS | Design System |
|----------|-------|---------------|
| iOS | ios | Cupertino |
| macOS | darwin | Cupertino |
| Android | android | Material |
| Linux | linux | Material |
| Windows | windows | Material |
| Web | js | Material |

## Best Practices

1. **Use adaptive widgets by default** for cross-platform apps
2. **Test on multiple platforms** during development
3. **Respect platform conventions** (e.g., back button placement)
4. **Customize carefully** - maintain platform feel while adding brand colors
5. **Provide platform-specific overrides** only when necessary

## Examples

See `examples/adaptive-demo/main.go` for a complete example that demonstrates adaptive widgets in action.

## Limitations

Some platform-specific features may not have direct equivalents:

- Material's outlined button becomes gray button on iOS
- Cupertino's navigation bar supports only one trailing widget
- Elevation effects are Material-only

These differences are handled gracefully to maintain a native feel on each platform.
