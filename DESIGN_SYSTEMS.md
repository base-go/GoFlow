# GoFlow Design Systems

GoFlow supports multiple design systems that allow you to create native-feeling applications across different platforms.

## Overview

GoFlow provides three approaches to building UI:

1. **Adaptive Widgets** - Automatically switch between design systems based on platform (recommended)
2. **Material Design** - Google's design system (Android, Linux, Windows, Web)
3. **Cupertino** - Apple's design language (iOS, macOS)

## Adaptive Widgets (Recommended)

Adaptive widgets automatically select the appropriate design system based on the runtime platform:

- **iOS/macOS**: Uses Cupertino (iOS-style) widgets
- **Android/Linux/Windows/Web**: Uses Material Design widgets

### Usage

```go
import "github.com/base-go/GoFlow/adaptive"

// Create a button that automatically adapts to the platform
button := adaptive.NewButton("Click Me", func() {
    fmt.Println("Button clicked!")
})

// Create a card
card := adaptive.NewCard(content)

// Create an app bar
appBar := adaptive.NewAppBar(title)
```

### Platform Detection

GoFlow automatically detects the platform at runtime:

```go
import "github.com/base-go/GoFlow/goflow"

platform := goflow.GetPlatform() // Returns: PlatformAndroid, PlatformIOS, etc.
theme := platform.DefaultTheme()  // Returns: "material" or "cupertino"
```

### Override Platform (for testing)

```go
goflow.SetPlatform(goflow.PlatformIOS)     // Force iOS style
goflow.SetPlatform(goflow.PlatformAndroid) // Force Material style
```

## Material Design

Material Design is Google's design system, used primarily on Android but also suitable for web and desktop applications.

### Features

- **Elevation**: Shadow-based depth system
- **Typography**: Roboto font family
- **Color System**: Primary, accent, and semantic colors
- **Angular corners**: 4px default border radius

### Widgets

```go
import "github.com/base-go/GoFlow/material"

// Theme
theme := material.DefaultLightTheme()
darkTheme := material.DefaultDarkTheme()

// Buttons
filledBtn := material.NewButton(child, onPressed)
textBtn := material.NewTextButton("Text", onPressed)
outlinedBtn := material.NewOutlinedButton("Outlined", onPressed)

// Card
card := material.NewCard(content)

// AppBar
appBar := material.NewAppBar(title)
appBar.Leading = backButton
appBar.Actions = []goflow.Widget{actionButton1, actionButton2}
```

### Material Theme

```go
type MaterialTheme struct {
    // Primary colors
    PrimaryColor      *goflow.Color  // Blue 500
    PrimaryDarkColor  *goflow.Color  // Blue 700
    PrimaryLightColor *goflow.Color  // Blue 300
    AccentColor       *goflow.Color  // Pink A200

    // Background colors
    BackgroundColor *goflow.Color    // Grey 50
    SurfaceColor    *goflow.Color    // White
    ErrorColor      *goflow.Color    // Red 500

    // Text colors
    TextPrimaryColor   *goflow.Color  // Almost Black
    TextSecondaryColor *goflow.Color  // Grey 600
    TextDisabledColor  *goflow.Color  // Grey 400
}
```

## Cupertino (iOS/macOS)

Cupertino is Apple's design language, providing an iOS and macOS native look and feel.

### Features

- **Blur effects**: Translucent backgrounds
- **San Francisco font**: System font family
- **Rounded corners**: 10px default border radius
- **System colors**: Dynamic colors that adapt to light/dark mode
- **44pt touch targets**: Minimum size for interactive elements

### Widgets

```go
import "github.com/base-go/GoFlow/cupertino"

// Theme
theme := cupertino.DefaultLightTheme()
darkTheme := cupertino.DefaultDarkTheme()

// Buttons
filledBtn := cupertino.NewButton(child, onPressed)
textBtn := cupertino.NewTextButton("Text", onPressed)
grayBtn := cupertino.NewGrayButton("Gray", onPressed)

// Card
card := cupertino.NewCard(content)

// NavigationBar
navBar := cupertino.NewNavigationBar(title)
navBar.Leading = backButton
navBar.Trailing = actionButton
```

### Cupertino Theme

```go
type CupertinoTheme struct {
    // Primary colors
    PrimaryColor   *goflow.Color  // iOS Blue (#007AFF)
    SecondaryColor *goflow.Color  // iOS Purple

    // Background colors
    BackgroundColor              *goflow.Color  // White
    SystemBackgroundColor        *goflow.Color  // Light Grey
    SecondaryBackgroundColor     *goflow.Color  // White

    // Label colors (with opacity support)
    LabelColor           *goflow.Color  // Black
    SecondaryLabelColor  *goflow.Color  // 60% opacity
    TertiaryLabelColor   *goflow.Color  // 30% opacity

    // Destructive color
    DestructiveColor *goflow.Color  // iOS Red (#FF3B30)
}
```

## Platform Support

| Platform | GOOS | Default Design System |
|----------|------|-----------------------|
| Android | android | Material |
| iOS | ios | Cupertino |
| macOS | darwin | Cupertino |
| Linux | linux | Material |
| Windows | windows | Material |
| Web | js | Material |

## Examples

### Adaptive Example

See `examples/adaptive-demo/main.go` for a complete example of adaptive widgets.

### Material Example

See `examples/material-demo/main.go` for Material Design widgets.

### Cupertino Example

See `examples/cupertino-demo/main.go` for Cupertino widgets.

## Best Practices

1. **Use Adaptive Widgets**: For cross-platform apps, prefer adaptive widgets to automatically match platform conventions.

2. **Respect Platform Conventions**: When using platform-specific widgets, follow the design guidelines:
   - Material: [Material Design Guidelines](https://material.io/design)
   - Cupertino: [Human Interface Guidelines](https://developer.apple.com/design/human-interface-guidelines/)

3. **Test on Multiple Platforms**: Use `goflow.SetPlatform()` to test different design systems during development.

4. **Custom Styling**: You can customize colors and other properties while maintaining the overall design language.

5. **Consistency**: Stick to one design system per platform for consistency. Don't mix Material and Cupertino in the same app unless intentional.

## Future Enhancements

- Fluent Design (Windows 11)
- Custom theme support
- Theme inheritance and customization
- Dark mode automatic switching
- Accessibility features
- Animation and transitions
