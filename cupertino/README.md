# Cupertino Package

Cupertino (iOS/macOS) design implementation for GoFlow.

## Overview

This package provides Cupertino widgets following Apple's Human Interface Guidelines. Cupertino is the design language used for iOS and macOS applications.

## Widgets

### Button

iOS-style buttons with multiple types:

```go
// Filled button (default)
btn := cupertino.NewButton(&widgets.Text{Content: "Click"}, func() {
    fmt.Println("Clicked!")
})

// Text button (plain)
textBtn := cupertino.NewTextButton("Text Button", onPressed)

// Gray button
grayBtn := cupertino.NewGrayButton("Gray Button", onPressed)

// Custom styling
btn.Color = cupertino.DefaultLightTheme().PrimaryColor
btn.BorderRadius = 10.0
btn.MinSize = 44.0 // iOS minimum touch target
```

### Card

iOS-style cards (similar to grouped list items):

```go
card := cupertino.NewCard(content)
card.BorderRadius = 10.0
card.Padding = goflow.NewEdgeInsets(16, 12, 16, 12)
```

### NavigationBar

iOS-style navigation bar:

```go
navBar := cupertino.NewNavigationBar(&widgets.Text{Content: "Title"})
navBar.Leading = backButton
navBar.Trailing = actionButton
```

## Theme

Cupertino themes follow iOS system colors:

```go
// Light theme
theme := cupertino.DefaultLightTheme()

// Dark theme
darkTheme := cupertino.DefaultDarkTheme()

// Access colors
primaryColor := theme.PrimaryColor // iOS Blue
labelColor := theme.LabelColor
```

### Color System

- **Primary**: iOS system blue (#007AFF)
- **Secondary**: iOS system purple
- **Background**: Primary background (White/Black)
- **System Background**: Secondary background (Light Grey/Dark Grey)
- **Label Colors**: Text colors with opacity variants
- **Destructive**: iOS system red for destructive actions

## Design Principles

1. **Clarity**: Text is legible, icons are precise, adornments are subtle
2. **Deference**: Fluid motion and crisp interface help users understand content
3. **Depth**: Visual layers and realistic motion convey hierarchy

## iOS Specifics

- **Rounded Corners**: 10px default (more rounded than Material's 4px)
- **Touch Targets**: Minimum 44pt for interactive elements
- **Typography**: San Francisco font system
- **Blur Effects**: Translucent backgrounds (future enhancement)

## Examples

See `examples/cupertino-demo/main.go` for a complete example.
