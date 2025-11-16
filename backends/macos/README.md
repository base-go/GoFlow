# macOS Rendering Backend

Native rendering backend for macOS using Core Graphics and Cocoa.

## Overview

This backend implements the GoFlow Canvas interface using:
- **Core Graphics (Quartz 2D)** - 2D rendering API
- **Core Text** - Text layout and rendering
- **Cocoa (AppKit)** - Window management and event handling

## Features

✅ **Implemented:**
- Rectangle drawing (filled and stroked)
- Circle drawing (filled and stroked)
- Line drawing
- Text rendering with font support
- Canvas transformations (translate, scale, rotate)
- Canvas state save/restore
- Clipping rectangles
- Window creation and management
- Basic event loop

⏳ **TODO:**
- Mouse event handling
- Keyboard event handling
- More drawing primitives (paths, images)
- Advanced text layout

## Architecture

```
┌─────────────────────────────────────┐
│   window_macos.go (Go API)          │
├─────────────────────────────────────┤
│   canvas_macos.go (Canvas impl)     │
├─────────────────────────────────────┤
│   CGo Bridge (C/Go boundary)        │
├─────────────────────────────────────┤
│   window_bridge.m (Objective-C)     │
│   cgo_bridge.m (Objective-C)        │
├─────────────────────────────────────┤
│   Core Graphics / Core Text         │
│   Cocoa (AppKit)                    │
└─────────────────────────────────────┘
```

## Usage

```go
package main

import (
    "github.com/base-go/GoFlow/backends/macos"
    "github.com/base-go/GoFlow/goflow"
)

func main() {
    // Initialize app
    macos.InitApp()

    // Create window
    window := macos.NewWindow(800, 600, "My GoFlow App")
    defer window.Destroy()

    // Set draw callback
    window.SetDrawFunc(func(canvas *macos.CoreGraphicsCanvas) {
        // Clear background
        canvas.Clear(goflow.ColorWhite)

        // Draw a rectangle
        paint := goflow.NewPaint()
        paint.Color = goflow.ColorBlue
        rect := goflow.NewRect(goflow.NewOffset(100, 100), goflow.NewSize(200, 150))
        canvas.DrawRect(rect, paint)

        // Draw text
        style := goflow.NewTextStyle()
        style.FontSize = 24
        style.Color = goflow.ColorBlack
        canvas.DrawText("Hello, GoFlow!", goflow.NewOffset(100, 50), style)
    })

    // Show window
    window.Show()

    // Event loop
    for !window.ShouldClose() {
        window.PollEvents()
        window.SetNeedsDisplay()
    }
}
```

## Building

This backend requires macOS and uses CGo to interface with Objective-C:

```bash
go build -o myapp
./myapp
```

## Files

- `canvas_macos.go` - Canvas implementation
- `window_macos.go` - Window management
- `cgo_bridge.h/m` - Core Graphics bridging
- `window_bridge.h/m` - Cocoa window bridging

## Platform Requirements

- **macOS 10.12+** (Sierra or later recommended)
- **Xcode Command Line Tools** (for compilation)

## Performance

The Core Graphics backend is hardware-accelerated and provides excellent performance for 2D rendering. Text rendering uses the system font renderer for perfect quality.

## See Also

- [Core Graphics Documentation](https://developer.apple.com/documentation/coregraphics)
- [Core Text Documentation](https://developer.apple.com/documentation/coretext)
- [Cocoa Documentation](https://developer.apple.com/documentation/appkit)
