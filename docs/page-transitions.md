# Page Transitions Guide

GoFlow provides a comprehensive page transition library inspired by Flutter's page route animations. This guide covers everything you need to know about creating smooth, beautiful transitions between pages in your application.

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Transition Types](#transition-types)
- [Basic Usage](#basic-usage)
- [Advanced Usage](#advanced-usage)
- [Custom Transitions](#custom-transitions)
- [Animation Curves](#animation-curves)
- [Best Practices](#best-practices)
- [API Reference](#api-reference)

## Overview

Page transitions are animations that play when navigating between different pages in your application. They provide visual continuity and help users understand the navigation flow.

### Features

- **16+ built-in transitions** - From simple fades to complex parallax effects
- **Platform-specific styles** - iOS (Cupertino) and Material Design transitions
- **Fully customizable** - Control duration, curves, and create custom transitions
- **Seamless integration** - Works with GoFlow's navigation system
- **Performant** - Built on GoFlow's efficient animation framework

## Quick Start

### Basic Transition

```go
import (
    "github.com/base-go/GoFlow/pkg/navigation"
    goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// Navigate with a fade transition
navigation.Get.To(NewDetailsPage(), navigation.TransitionFade)
```

### Platform-Specific Transition

```go
// iOS-style transition
navigation.Get.To(NewPage(), navigation.TransitionCupertino)

// Material Design transition
navigation.Get.To(NewPage(), navigation.TransitionMaterial)
```

## Transition Types

### Basic Transitions

#### Fade
```go
navigation.TransitionFade
```
Smoothly fades between pages. Great for subtle transitions.

**Use cases:** Settings pages, overlays, non-directional navigation

#### Slide Right (iOS Style)
```go
navigation.TransitionSlideRight  // or TransitionCupertino
```
Slides new page in from the right, old page slides out to the left.

**Use cases:** iOS apps, forward navigation, drill-down patterns

#### Slide Left
```go
navigation.TransitionSlideLeft
```
Slides new page in from the left.

**Use cases:** Back navigation, bidirectional navigation

#### Slide Up (Material Style)
```go
navigation.TransitionSlideUp  // or TransitionMaterial
```
Slides new page in from the bottom.

**Use cases:** Android apps, modal-like pages, bottom sheets to full page

#### Slide Down
```go
navigation.TransitionSlideDown
```
Slides new page in from the top.

**Use cases:** Notifications to detail, special emphasis

#### Zoom
```go
navigation.TransitionZoom
```
Combines scale and fade for a zoom effect.

**Use cases:** Image galleries, expanding content, focus transitions

#### None
```go
navigation.TransitionNone
```
Instant transition with no animation.

**Use cases:** Tab switching, instant updates, accessibility settings

### Advanced Transition Types

When using the `PageTransition` widget directly, you have access to additional transition types:

```go
// Combined transitions
navigation.PageTransitionTypeRightToLeftWithFade
navigation.PageTransitionTypeLeftToRightWithFade

// Joined transitions (old page slides out while new slides in)
navigation.PageTransitionTypeRightToLeftJoined
navigation.PageTransitionTypeLeftToRightJoined

// Special effects
navigation.PageTransitionTypeRippleEffect
navigation.PageTransitionTypeSlideParallax
navigation.PageTransitionTypeRotate
navigation.PageTransitionTypeScale
```

## Basic Usage

### With Navigation Methods

```go
// Using Get.To
navigation.Get.To(NewMyPage(), navigation.TransitionSlideRight)

// Using Get.ToNamed
navigation.Get.ToNamed("/settings", navigation.TransitionFade)

// Using Get.Off (replace current page)
navigation.Get.Off(NewHomePage(), navigation.TransitionSlideLeft)

// Using Get.OffAll (clear stack)
navigation.Get.OffAll(NewLoginPage(), navigation.TransitionFade)
```

### In Route Definitions

```go
routes := map[string]navigation.RouteBuilder{
    "/home": func() goflow.Widget { return NewHomePage() },
    "/details": func() goflow.Widget { return NewDetailsPage() },
}

// Routes can specify default transitions when building the route
route := navigation.NewPageRoute(NewMyPage(), navigation.TransitionFade)
```

## Advanced Usage

### Using PageTransition Widget

The `PageTransition` widget provides more control over transitions:

```go
import "time"

transition := navigation.NewPageTransition(
    navigation.PageTransitionTypeZoom,
    myPage,
).WithDuration(500 * time.Millisecond).
  WithCurve(animation.EaseInOut)
```

### Configuration Options

```go
// Set custom duration
transition.WithDuration(400 * time.Millisecond)

// Set animation curve
transition.WithCurve(animation.BounceOut)

// Set alignment for scale/zoom transitions
transition.WithAlignment(goflow.Alignment{X: 0, Y: -1}) // Top center

// Set custom builder
transition.WithMatchingBuilder(myCustomBuilder)
```

### PageTransitionsBuilder

For maximum control, create a custom `PageTransitionsBuilder`:

```go
func MyCustomTransition(
    context goflow.BuildContext,
    animation *animation.AnimationController,
    secondaryAnimation *animation.AnimationController,
    child goflow.Widget,
) goflow.Widget {
    curvedAnimation := animation.NewCurvedAnimation(animation, animation.EaseInOut)

    // Create custom animation
    return widgets.NewFadeTransition(curvedAnimation, child)
}
```

## Custom Transitions

### Directional Slide Transitions

Create a slide transition in any direction:

```go
customSlide := navigation.CreateDirectionalSlideTransition(
    navigation.TransitionDirectionUp,
    600.0, // distance in pixels
)

// Use it with CustomPageTransition
custom := navigation.NewCustomPageTransition(
    myPage,
    customSlide,
    300 * time.Millisecond,
    animation.EaseOut,
)
```

### Combined Scale and Rotate

```go
scaleRotate := navigation.CreateScaleRotateTransition(
    0.5,  // initial scale (0.5 = 50%)
    0.25, // rotation in turns (0.25 = 90 degrees)
)
```

### Fully Custom Transition

Implement your own `PageTransitionsBuilder`:

```go
func MyComplexTransition(
    context goflow.BuildContext,
    anim *animation.AnimationController,
    secondaryAnim *animation.AnimationController,
    child goflow.Widget,
) goflow.Widget {
    progress := anim.Value()

    // Example: Slide + Fade + Scale
    offset := (1.0 - progress) * 400
    opacity := progress
    scale := 0.8 + (progress * 0.2)

    scaled := &widgets.Transform{
        Transform: &widgets.ScaleTransform{
            ScaleX: scale,
            ScaleY: scale,
        },
        Child: child,
    }

    translated := &widgets.Transform{
        Transform: &widgets.TranslationTransform{
            X: offset,
            Y: 0,
        },
        Child: scaled,
    }

    return &widgets.Opacity{
        Opacity: opacity,
        Child: translated,
    }
}
```

## Animation Curves

Control the easing of your transitions:

```go
import "github.com/base-go/GoFlow/pkg/core/animation"

// Available curves:
animation.Linear       // No easing
animation.EaseIn       // Slow start
animation.EaseOut      // Slow end
animation.EaseInOut    // Slow start and end
animation.BounceIn     // Bounce at start
animation.BounceOut    // Bounce at end
animation.ElasticIn    // Elastic at start
animation.ElasticOut   // Elastic at end
```

### Using Curves

```go
transition := navigation.NewPageTransition(
    navigation.PageTransitionTypeFade,
    myPage,
).WithCurve(animation.BounceOut)
```

## Best Practices

### 1. Choose Platform-Appropriate Transitions

```go
// For iOS
navigation.Get.To(page, navigation.TransitionCupertino)

// For Android
navigation.Get.To(page, navigation.TransitionMaterial)
```

### 2. Keep Duration Reasonable

```go
// Good: 200-400ms for most transitions
transition.WithDuration(300 * time.Millisecond)

// Avoid: Too fast (jarring) or too slow (sluggish)
transition.WithDuration(50 * time.Millisecond)   // Too fast
transition.WithDuration(1000 * time.Millisecond) // Too slow
```

### 3. Match Transition Direction to Navigation

```go
// Forward navigation: slide from right
navigation.Get.To(nextPage, navigation.TransitionSlideRight)

// Back navigation: use Back() which reverses the transition
navigation.Get.Back()
```

### 4. Use Semantic Transitions

```go
// Modal or bottom sheet style
navigation.Get.To(modalPage, navigation.TransitionSlideUp)

// Expanding detail view
navigation.Get.To(detailPage, navigation.TransitionZoom)

// Settings or preferences
navigation.Get.To(settingsPage, navigation.TransitionFade)
```

### 5. Consider Performance

```go
// For simple transitions, use built-in types
navigation.TransitionFade // More performant

// For complex transitions, optimize carefully
// Avoid multiple nested transforms and animations
```

### 6. Accessibility

```go
// Provide option to reduce or disable transitions
if userPrefersReducedMotion {
    navigation.Get.To(page, navigation.TransitionNone)
} else {
    navigation.Get.To(page, navigation.TransitionFade)
}
```

## API Reference

### Transition Constants

```go
const (
    TransitionFade
    TransitionSlideRight
    TransitionSlideLeft
    TransitionSlideUp
    TransitionSlideDown
    TransitionZoom
    TransitionNone
    TransitionCupertino  // Alias for SlideRight
    TransitionMaterial   // Alias for SlideUp
)
```

### PageTransition Widget

```go
type PageTransition struct {
    Type               PageTransitionType
    Child              goflow.Widget
    Duration           time.Duration
    ReverseDuration    time.Duration
    Curve              animation.Curve
    MatchingBuilder    PageTransitionsBuilder
    Alignment          goflow.Alignment
    IsIOS              bool
}

// Constructor
func NewPageTransition(transitionType PageTransitionType, child goflow.Widget) *PageTransition

// Fluent API
func (p *PageTransition) WithDuration(duration time.Duration) *PageTransition
func (p *PageTransition) WithCurve(curve animation.Curve) *PageTransition
func (p *PageTransition) WithAlignment(alignment goflow.Alignment) *PageTransition
func (p *PageTransition) WithMatchingBuilder(builder PageTransitionsBuilder) *PageTransition
```

### PageTransitionsBuilder

```go
type PageTransitionsBuilder func(
    context goflow.BuildContext,
    animation *animation.AnimationController,
    secondaryAnimation *animation.AnimationController,
    child goflow.Widget,
) goflow.Widget
```

### Built-in Builders

```go
func FadeTransitionBuilder(...) goflow.Widget
func RightToLeftTransitionBuilder(...) goflow.Widget
func LeftToRightTransitionBuilder(...) goflow.Widget
func TopToBottomTransitionBuilder(...) goflow.Widget
func BottomToTopTransitionBuilder(...) goflow.Widget
func ScaleTransitionBuilder(...) goflow.Widget
func RotateTransitionBuilder(...) goflow.Widget
func ZoomTransitionBuilder(...) goflow.Widget
func RippleEffectTransitionBuilder(...) goflow.Widget
func SlideParallaxTransitionBuilder(...) goflow.Widget
```

### Helper Functions

```go
// Create directional slide
func CreateDirectionalSlideTransition(
    direction TransitionDirection,
    distance float64,
) PageTransitionsBuilder

// Create scale + rotate
func CreateScaleRotateTransition(
    initialScale float64,
    rotationTurns float64,
) PageTransitionsBuilder

// Map transition types
func MapPageTransitionTypeToTransition(transitionType PageTransitionType) Transition
```

### TransitionBuilder (Simple)

```go
type TransitionBuilder func(child goflow.Widget, animation float64) goflow.Widget

// Get builder for transition type
func GetTransitionBuilder(transition Transition) TransitionBuilder
```

## Examples

See the complete example in `examples/page-transitions-demo/` for a full demonstration of all transition types.

## Related Documentation

- [Navigation System](navigation.md)
- [Animation Framework](animation.md)
- [Widget Library](widgets.md)

## Comparison with Flutter

GoFlow's page transitions are designed to be familiar to Flutter developers:

| Flutter | GoFlow |
|---------|--------|
| `MaterialPageRoute` | `navigation.TransitionMaterial` |
| `CupertinoPageRoute` | `navigation.TransitionCupertino` |
| `PageRouteBuilder` | `PageTransition` widget |
| `transitionDuration` | `Duration` field |
| `transitionsBuilder` | `PageTransitionsBuilder` |
| `FadeTransition` | `TransitionFade` |
| `SlideTransition` | `TransitionSlide*` |
| `ScaleTransition` | `TransitionZoom` |

## FAQ

**Q: Can I have different transitions for push and pop?**
A: Currently, the reverse transition mirrors the forward transition. Full support for custom reverse transitions is planned.

**Q: How do I disable transitions globally?**
A: You can create a configuration option that always uses `TransitionNone`.

**Q: Can I animate multiple properties at once?**
A: Yes! Create a custom `PageTransitionsBuilder` that combines multiple transforms and opacity.

**Q: What's the difference between `PageTransition` and `Transition`?**
A: `Transition` is a simple enum for common transitions. `PageTransition` is a widget that provides full control over the transition.

**Q: How do I make transitions feel more natural?**
A: Use appropriate curves! `EaseInOut` works well for most transitions. Experiment with `EaseOut` for more natural deceleration.

## Contributing

Want to add a new transition type or improve existing ones? Contributions are welcome! Check out the contributing guide.

## License

GoFlow is licensed under the MIT License. See LICENSE file for details.
