# Material Design Package

Material Design implementation for GoFlow.

## Overview

This package provides Material Design widgets following Google's Material Design specifications. Material Design is the design system used for Android applications and is also popular for web and cross-platform apps.

## Widgets

### Button

Material Design buttons with multiple styles:

```go
// Filled button (default)
btn := material.NewButton(&widgets.Text{Content: "Click"}, func() {
    fmt.Println("Clicked!")
})

// Text button
textBtn := material.NewTextButton("Text Button", onPressed)

// Outlined button
outlinedBtn := material.NewOutlinedButton("Outlined", onPressed)

// Custom styling
btn.BackgroundColor = material.DefaultLightTheme().AccentColor
btn.Elevation = 4.0
btn.BorderRadius = 8.0
```

### Card

Material Design cards for grouping content:

```go
card := material.NewCard(content)
card.Elevation = 2.0
card.BorderRadius = 4.0
card.Padding = goflow.NewEdgeInsets(16, 16, 16, 16)
```

### AppBar

Material Design app bar (toolbar):

```go
appBar := material.NewAppBar(&widgets.Text{Content: "Title"})
appBar.Leading = menuButton
appBar.Actions = []goflow.Widget{searchButton, moreButton}
appBar.Elevation = 4.0
```

## Theme

Material Design themes define the color palette:

```go
// Light theme
theme := material.DefaultLightTheme()

// Dark theme
darkTheme := material.DefaultDarkTheme()

// Access colors
primaryColor := theme.PrimaryColor
backgroundColor := theme.BackgroundColor
```

### Color System

- **Primary**: Main brand color (Blue 500)
- **Primary Dark**: Darker variant (Blue 700)
- **Primary Light**: Lighter variant (Blue 300)
- **Accent**: Secondary brand color (Pink A200)
- **Background**: App background (Grey 50)
- **Surface**: Card/surface color (White)
- **Error**: Error state color (Red 500)

## Design Principles

1. **Material is the metaphor**: Surfaces and shadows provide visual cues
2. **Bold, graphic, intentional**: Typography and color are bold and intentional
3. **Motion provides meaning**: Transitions and animations guide user attention

## Examples

See `examples/material-demo/main.go` for a complete example.
