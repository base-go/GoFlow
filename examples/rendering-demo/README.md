# GoFlow Rendering Demo

A demonstration of the GoFlow rendering engine using the native macOS backend.

## Features Demonstrated

- **Window Creation**: Native macOS window with Cocoa
- **Canvas Drawing**: Core Graphics rendering
- **Shapes**: Rectangles, circles, and lines
- **Text Rendering**: Using Core Text with custom fonts
- **Colors**: Full color support with transparency
- **Stroke vs Fill**: Both filled and stroked shapes
- **Animation**: Simple animation loop

## Building and Running

**macOS only:**

```bash
cd examples/rendering-demo
go run main.go
```

Or build and run:

```bash
go build -o rendering-demo
./rendering-demo
```

## What You'll See

The demo displays:

1. **Title and subtitle** - Text rendering with different sizes and weights
2. **Filled rectangle** - Blue rectangle with solid fill
3. **Stroked rectangle** - Red rectangle with outline only
4. **Filled circle** - Green circle with solid fill
5. **Stroked circle** - Yellow circle with outline only
6. **Colored lines** - Multiple lines in different colors
7. **Animated circle** - Purple circle (animation placeholder)
8. **Color palette** - All basic colors displayed as swatches

## Platform Requirements

- macOS 10.12+ (Sierra or later)
- Xcode Command Line Tools

## Architecture

This demo uses:
- `backends/macos` - Native macOS rendering backend
- Core Graphics (Quartz 2D) - 2D rendering
- Core Text - Text layout and rendering
- Cocoa (AppKit) - Window management

## Controls

- Close the window to exit the application

## Notes

This is a basic rendering demonstration. Future examples will include:
- Mouse interaction
- Keyboard input
- More complex animations
- Widget rendering
- Full GoFlow app integration
