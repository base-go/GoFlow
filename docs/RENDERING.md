# GoFlow Rendering Architecture

**Current Status: ⚠️ Architecture Only - No Actual Rendering Yet**

GoFlow has the rendering **architecture** designed and implemented, but no actual rendering backend. This document explains how rendering will work and the options for implementation.

## Current State (v0.1.0)

### What Exists ✅
- **Canvas Interface**: Abstraction for drawing operations
- **RenderObject System**: Layout and paint pipeline
- **Widget → Element → RenderObject** chain
- **MockCanvas**: Testing implementation only

### What Doesn't Exist ❌
- **No Real Rendering Backend**: No actual pixels on screen
- **No Window Creation**: Platform runners exist but don't create windows
- **No GPU Integration**: No Metal/DirectX/Vulkan/OpenGL
- **No Text Rendering**: No font rasterization
- **No Image Loading**: No asset loading

## Rendering Pipeline Architecture

### The Three-Tree System

```
┌─────────────┐
│   Widget    │  Immutable UI description
│   Tree      │  (What to show)
└──────┬──────┘
       │ Build()
       ↓
┌─────────────┐
│   Element   │  Mutable lifecycle manager
│   Tree      │  (When to rebuild)
└──────┬──────┘
       │ CreateRenderObject()
       ↓
┌─────────────┐
│  Render     │  Layout, Paint, Hit Test
│  Tree       │  (How to draw)
└─────────────┘
```

### Rendering Flow

```go
// 1. Build Phase: Widget tree created
app := &MyApp{}
goflow.RunApp(app)
    → app.Build(ctx) returns widget tree
    → Elements update/create based on widgets

// 2. Layout Phase: RenderObjects compute sizes
renderObject.Layout(constraints)
    → Constraints flow down
    → Sizes flow up
    → Parent positions children

// 3. Paint Phase: RenderObjects draw to Canvas
renderObject.Paint(canvas, offset)
    → Depth-first traversal
    → Calls canvas.DrawRect(), DrawText(), etc.
    → Canvas interface calls actual backend

// 4. Backend: Canvas implementation renders pixels
canvas.DrawRect(rect, paint)
    → NativeCanvas calls platform APIs
    → OR WGPUCanvas calls WebGPU
    → Pixels appear on screen
```

## Canvas Interface

The Canvas is GoFlow's abstraction over any rendering backend:

```go
type Canvas interface {
    DrawRect(rect *Rect, paint *Paint)
    DrawCircle(center *Offset, radius float64, paint *Paint)
    DrawText(text string, offset *Offset, style *TextStyle)
    DrawLine(p1, p2 *Offset, paint *Paint)

    // Transform operations
    Save()
    Restore()
    Translate(dx, dy float64)
    Scale(sx, sy float64)
    Rotate(radians float64)
    ClipRect(rect *Rect)

    Clear(color *Color)
}
```

**Any backend that implements this interface can render GoFlow apps.**

## Rendering Backend Options

### Option 1: **Native Renderers** (Recommended for Desktop)

Use each platform's native 2D graphics APIs directly.

#### macOS: Core Graphics / Quartz 2D
```go
// goflow/backends/macos/core_graphics_canvas.go
type CoreGraphicsCanvas struct {
    context CGContextRef  // Core Graphics context
}

func (c *CoreGraphicsCanvas) DrawRect(rect *Rect, paint *Paint) {
    // Convert GoFlow types to CG types
    cgRect := CGRectMake(rect.Left(), rect.Top(), rect.Size.Width, rect.Size.Height)

    // Set color
    CGContextSetRGBFillColor(c.context, paint.Color.R, paint.Color.G, paint.Color.B, paint.Color.A)

    // Draw
    CGContextFillRect(c.context, cgRect)
}

func (c *CoreGraphicsCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    // Use Core Text for text rendering
    ctFont := CTFontCreateWithName(style.FontFamily, style.FontSize, nil)
    // ... render text using Core Text
}
```

**Advantages:**
- Native look and feel
- Perfect font rendering (system fonts)
- Hardware accelerated
- Best performance on each platform
- Access to platform-specific features

**Disadvantages:**
- More code (one implementation per platform)
- Platform-specific APIs to learn
- Harder to keep consistent across platforms

#### Windows: Direct2D
```go
// goflow/backends/windows/direct2d_canvas.go
type Direct2DCanvas struct {
    renderTarget *ID2D1HwndRenderTarget
    factory      *ID2D1Factory
}

func (c *Direct2DCanvas) DrawRect(rect *Rect, paint *Paint) {
    d2dRect := D2D1_RECT_F{
        left:   float32(rect.Left()),
        top:    float32(rect.Top()),
        right:  float32(rect.Right()),
        bottom: float32(rect.Bottom()),
    }

    brush := c.createBrushFromPaint(paint)
    c.renderTarget.FillRectangle(&d2dRect, brush)
}

func (c *Direct2DCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    // Use DirectWrite for text
    textFormat := c.createTextFormat(style)
    c.renderTarget.DrawText(text, textFormat, ...)
}
```

#### Linux: Cairo
```go
// goflow/backends/linux/cairo_canvas.go
type CairoCanvas struct {
    context *cairo.Context
}

func (c *CairoCanvas) DrawRect(rect *Rect, paint *Paint) {
    cairo.Rectangle(c.context, rect.Left(), rect.Top(), rect.Size.Width, rect.Size.Height)
    cairo.SetSourceRGBA(c.context, paint.Color.R, paint.Color.G, paint.Color.B, paint.Color.A)
    cairo.Fill(c.context)
}

func (c *CairoCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    // Use Pango for text rendering
    layout := pango.CreateLayout(c.context)
    // ... render with Pango
}
```

### Option 2: **WGPU** (WebGPU - Cross-Platform)

Single implementation using wgpu-native (WebGPU in native code).

```go
// goflow/backends/wgpu/wgpu_canvas.go
type WGPUCanvas struct {
    device      *wgpu.Device
    queue       *wgpu.Queue
    surface     *wgpu.Surface
    renderPass  *wgpu.RenderPass
}

func (c *WGPUCanvas) DrawRect(rect *Rect, paint *Paint) {
    // Create GPU buffers for rectangle vertices
    vertices := createRectVertices(rect)
    buffer := c.device.CreateBuffer(vertices)

    // Create pipeline with shaders
    pipeline := c.createRectPipeline(paint.Color)

    // Submit draw command
    c.renderPass.SetPipeline(pipeline)
    c.renderPass.Draw(buffer)
}
```

**Advantages:**
- Single codebase for all platforms
- Modern GPU API
- Future-proof (WebGPU standard)
- Works on Web (WASM)
- Good performance

**Disadvantages:**
- Text rendering is harder (need font rasterization library)
- Not as "native" feeling
- Larger binary size
- More complex setup

### Option 3: **Hybrid** (Like Flutter)

Native window + WGPU for rendering.

```go
// macOS: NSWindow with Metal surface for WGPU
// Windows: HWND with DirectX surface for WGPU
// Linux: X11/Wayland window with Vulkan surface for WGPU
```

This is what Flutter does: Skia (their 2D engine) uses platform backends.

## Recommended Approach: **Start Native, Add WGPU Later**

### Phase 1: Native Backends (v0.2.0)
Implement native rendering for each platform:
- **macOS**: Core Graphics + Core Text
- **Windows**: Direct2D + DirectWrite
- **Linux**: Cairo + Pango

Benefits:
- Best user experience
- Learn platform APIs
- Fast time-to-pixels
- True native apps

### Phase 2: WGPU Option (v0.3.0)
Add WGPU as an alternative backend:
- Same Canvas interface
- Choice at build time: `--backend=native` or `--backend=wgpu`
- WGPU enables Web platform

## Implementation Plan

### Step 1: macOS Native (Easiest Start)

```go
// goflow/backends/macos/cgo_bridge.go
package macos

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework QuartzCore
#import <Cocoa/Cocoa.h>
#import <QuartzCore/QuartzCore.h>

void* createGraphicsContext(int width, int height);
void drawRect(void* ctx, double x, double y, double w, double h, double r, double g, double b, double a);
void drawText(void* ctx, const char* text, double x, double y, const char* font, double size);
*/
import "C"
import "unsafe"

type CoreGraphicsCanvas struct {
    ctx unsafe.Pointer
}

func NewCoreGraphicsCanvas(width, height int) *CoreGraphicsCanvas {
    return &CoreGraphicsCanvas{
        ctx: C.createGraphicsContext(C.int(width), C.int(height)),
    }
}

func (c *CoreGraphicsCanvas) DrawRect(rect *goflow.Rect, paint *goflow.Paint) {
    C.drawRect(
        c.ctx,
        C.double(rect.Left()),
        C.double(rect.Top()),
        C.double(rect.Size.Width),
        C.double(rect.Size.Height),
        C.double(paint.Color.R),
        C.double(paint.Color.G),
        C.double(paint.Color.B),
        C.double(paint.Color.A),
    )
}
```

```objc
// goflow/backends/macos/cgo_bridge.m
#import <Cocoa/Cocoa.h>

void* createGraphicsContext(int width, int height) {
    NSBitmapImageRep *bitmap = [[NSBitmapImageRep alloc]
        initWithBitmapDataPlanes:NULL
        pixelsWide:width
        pixelsHigh:height
        bitsPerSample:8
        samplesPerPixel:4
        hasAlpha:YES
        isPlanar:NO
        colorSpaceName:NSCalibratedRGBColorSpace
        bytesPerRow:0
        bitsPerPixel:0];

    NSGraphicsContext *context = [NSGraphicsContext graphicsContextWithBitmapImageRep:bitmap];
    return (void*)CFBridgingRetain(context);
}

void drawRect(void* ctx, double x, double y, double w, double h, double r, double g, double b, double a) {
    NSGraphicsContext *context = (__bridge NSGraphicsContext*)ctx;
    [NSGraphicsContext setCurrentContext:context];

    [[NSColor colorWithRed:r green:g blue:b alpha:a] setFill];
    NSRectFill(NSMakeRect(x, y, w, h));
}
```

### Step 2: Integrate with Platform Runners

```go
// macos/main.go
package main

import (
    "com.example/myapp/lib"
    "github.com/base-go/GoFlow/goflow"
    "github.com/base-go/GoFlow/backends/macos"
)

func main() {
    // Create native window
    window := macos.CreateWindow(800, 600, "My GoFlow App")

    // Create rendering backend
    canvas := macos.NewCoreGraphicsCanvas(800, 600)

    // Run app with native backend
    app := lib.NewMyApp()
    goflow.RunAppWithBackend(app, canvas, window)
}
```

### Step 3: Event Loop Integration

```go
// Platform runner handles events
for {
    event := window.PollEvent()

    switch event.Type {
    case EventResize:
        // Trigger re-layout
        goflow.MarkNeedsLayout()

    case EventMouseClick:
        // Hit test and dispatch
        goflow.DispatchPointerEvent(event.Position)

    case EventPaint:
        // Re-paint
        goflow.MarkNeedsPaint()
        renderObject.Paint(canvas, offset)
        window.Present()
    }
}
```

## Text Rendering

Each native backend uses its platform's text engine:

### macOS: Core Text
```go
func (c *CoreGraphicsCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    // Create attributed string
    ctFont := CTFontCreateWithName(style.FontFamily, style.FontSize, nil)
    attributes := CFDictionaryCreate(...)

    // Create text layout
    line := CTLineCreateWithAttributedString(attrString)

    // Draw
    CGContextSetTextPosition(c.context, offset.X, offset.Y)
    CTLineDraw(line, c.context)
}
```

### Windows: DirectWrite
```go
func (c *Direct2DCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    textFormat := c.factory.CreateTextFormat(
        style.FontFamily,
        nil,
        DWRITE_FONT_WEIGHT_NORMAL,
        DWRITE_FONT_STYLE_NORMAL,
        DWRITE_FONT_STRETCH_NORMAL,
        style.FontSize,
        "en-us",
    )

    c.renderTarget.DrawText(text, textFormat, rect, brush)
}
```

### Linux: Pango
```go
func (c *CairoCanvas) DrawText(text string, offset *Offset, style *TextStyle) {
    layout := pango.CreateLayout(c.context)
    layout.SetText(text)

    font := pango.FontDescriptionFromString(fmt.Sprintf("%s %f", style.FontFamily, style.FontSize))
    layout.SetFontDescription(font)

    cairo.MoveTo(c.context, offset.X, offset.Y)
    pango.ShowLayout(c.context, layout)
}
```

## Performance Considerations

### Native Backends
- **Immediate Mode**: Each frame redraws everything
- **Dirty Tracking**: Only repaint what changed
- **Layer Caching**: Cache complex widgets as bitmaps
- **Hardware Acceleration**: Use GPU when available

### Optimization Strategy
1. **Dirty Rectangles**: Track which areas need repainting
2. **Render Layers**: Cache static content
3. **Async Rendering**: Render in background thread
4. **GPU Compositing**: Use platform compositing APIs

## Next Steps

1. **✅ Document rendering architecture** (this doc)
2. **⏳ Implement macOS Core Graphics backend**
3. **⏳ Implement Windows Direct2D backend**
4. **⏳ Implement Linux Cairo backend**
5. **⏳ Add GLFW window creation**
6. **⏳ Integrate event loop**
7. **⏳ Add text rendering**
8. **⏳ Add image loading**
9. **⏳ WGPU backend (optional)**

## See Also

- [Platform Integration](./PLATFORM_INTEGRATION.md) - Native platform projects
- [Architecture](../ARCHITECTURE.md) - Framework design
- [Flutter Comparison](./FLUTTER_INSPIRATION.md) - How Flutter does it

## References

### Native APIs
- **macOS**: [Core Graphics](https://developer.apple.com/documentation/coregraphics), [Core Text](https://developer.apple.com/documentation/coretext)
- **Windows**: [Direct2D](https://learn.microsoft.com/en-us/windows/win32/direct2d/direct2d-portal), [DirectWrite](https://learn.microsoft.com/en-us/windows/win32/directwrite/direct-write-portal)
- **Linux**: [Cairo](https://www.cairographics.org/), [Pango](https://pango.gnome.org/)

### Cross-Platform Options
- **WGPU**: [wgpu-native](https://github.com/gfx-rs/wgpu-native)
- **Skia**: [Skia Graphics Library](https://skia.org/) (what Flutter uses)
- **GLFW**: [Window creation](https://www.glfw.org/)
