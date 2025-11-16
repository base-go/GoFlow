# Page Transitions Demo

This example demonstrates GoFlow's comprehensive page transition library, inspired by Flutter's page route animations.

## Overview

The page transitions library provides smooth, customizable animations when navigating between pages in your GoFlow application. It supports various transition types including fades, slides, zooms, and custom transitions.

## Running the Demo

```bash
cd examples/page-transitions-demo
go run main.go
```

## Features Demonstrated

### Basic Transitions

1. **Fade Transition** - Smooth fade in/out effect
2. **Slide Transitions** - Slide from any direction:
   - Right to Left (iOS style)
   - Left to Right
   - Bottom to Top (Material style)
   - Top to Bottom
3. **Zoom Transition** - Scale and fade combined
4. **No Transition** - Instant navigation

### Platform-Specific Transitions

1. **Cupertino (iOS)** - Standard iOS slide from right
2. **Material (Android)** - Standard Material Design slide from bottom

### Advanced Features

- **Custom Transitions** - Create your own transition effects
- **Configurable Duration** - Control animation speed
- **Easing Curves** - Choose from multiple animation curves:
  - Linear
  - EaseIn, EaseOut, EaseInOut
  - BounceIn, BounceOut
  - ElasticIn, ElasticOut

## Usage Examples

### Basic Navigation with Transition

```go
// Navigate to a page with fade transition
navigation.Get.To(NewMyPage(), navigation.TransitionFade)

// Navigate with slide right (iOS style)
navigation.Get.To(NewMyPage(), navigation.TransitionSlideRight)

// Navigate with zoom
navigation.Get.To(NewMyPage(), navigation.TransitionZoom)
```

### Named Routes with Transitions

```go
// Navigate to named route with specific transition
navigation.Get.ToNamed("/settings", navigation.TransitionMaterial)
```

### Stack Operations with Transitions

```go
// Replace current page
navigation.Get.Off(NewHomePage(), navigation.TransitionFade)

// Clear stack and navigate
navigation.Get.OffAll(NewHomePage(), navigation.TransitionSlideLeft)
```

### Custom Page Transitions

```go
// Create a custom transition
customTransition := navigation.NewPageTransition(
    navigation.PageTransitionTypeRightToLeftWithFade,
    myWidget,
).WithDuration(500 * time.Millisecond).
  WithCurve(animation.EaseInOut)
```

### Advanced Custom Transitions

```go
// Create a directional slide transition
customBuilder := navigation.CreateDirectionalSlideTransition(
    navigation.TransitionDirectionUp,
    600.0, // distance
)

// Create a scale + rotate transition
scaleRotateBuilder := navigation.CreateScaleRotateTransition(
    0.5,  // initial scale
    0.25, // rotation (in turns)
)
```

## Available Transition Types

### Simple Transitions
- `TransitionFade` - Fade in/out
- `TransitionSlideRight` - Slide from right
- `TransitionSlideLeft` - Slide from left
- `TransitionSlideUp` - Slide from bottom
- `TransitionSlideDown` - Slide from top
- `TransitionZoom` - Zoom in/out
- `TransitionNone` - No animation

### Platform Transitions
- `TransitionCupertino` - iOS style (slide right)
- `TransitionMaterial` - Material Design style (slide up)

### Advanced Page Transition Types

When using `PageTransition` widget directly:

- `PageTransitionTypeFade` - Simple fade
- `PageTransitionTypeRightToLeft` - Slide right to left
- `PageTransitionTypeLeftToRight` - Slide left to right
- `PageTransitionTypeTopToBottom` - Slide top to bottom
- `PageTransitionTypeBottomToTop` - Slide bottom to top
- `PageTransitionTypeScale` - Scale animation
- `PageTransitionTypeRotate` - Rotation animation
- `PageTransitionTypeSize` - Size change animation
- `PageTransitionTypeRightToLeftWithFade` - Combined slide and fade
- `PageTransitionTypeLeftToRightWithFade` - Combined slide and fade
- `PageTransitionTypeRightToLeftJoined` - Joined page transition
- `PageTransitionTypeLeftToRightJoined` - Joined page transition
- `PageTransitionTypeZoom` - Zoom with fade
- `PageTransitionTypeFadeIn` - Fade in only
- `PageTransitionTypeRippleEffect` - Ripple/circular reveal
- `PageTransitionTypeSlideParallax` - Parallax slide effect

## Architecture

The page transitions library is built on GoFlow's animation framework and consists of:

1. **PageTransition Widget** - Main widget for creating transitions
2. **Transition Builders** - Functions that build animated widgets
3. **Route Integration** - Seamless integration with navigation system
4. **Animation Controller** - Manages animation lifecycle
5. **Curves** - Various easing functions for smooth animations

## Code Structure

```
pkg/navigation/
├── page_transitions.go  # Main transitions library
├── route.go            # Route and transition types
└── navigator.go        # Navigator integration

examples/page-transitions-demo/
└── main.go            # This demo
```

## Key Concepts

### Transition Types

Transitions are defined as constants that specify the animation style:

```go
type Transition int

const (
    TransitionFade
    TransitionSlideRight
    TransitionSlideLeft
    // ... etc
)
```

### Transition Builders

Builders are functions that create animated widgets:

```go
type TransitionBuilder func(child goflow.Widget, animation float64) goflow.Widget
```

### Page Transitions Builder

More advanced builders with full animation controller access:

```go
type PageTransitionsBuilder func(
    context goflow.BuildContext,
    animation *animation.AnimationController,
    secondaryAnimation *animation.AnimationController,
    child goflow.Widget,
) goflow.Widget
```

## Comparison with Flutter

GoFlow's page transitions library closely mirrors Flutter's approach:

| Flutter | GoFlow |
|---------|--------|
| `PageRouteBuilder` | `PageRouteBuilder` |
| `PageTransitionsBuilder` | `PageTransitionsBuilder` |
| `FadeTransition` | `TransitionFade` |
| `SlideTransition` | `TransitionSlideRight/Left/Up/Down` |
| `ScaleTransition` | `TransitionZoom` |
| Custom transitions | `CustomPageTransition` |

## Future Enhancements

- [ ] Gesture-driven transitions
- [ ] Hero animations (shared element transitions)
- [ ] Curved path transitions
- [ ] Physics-based transitions
- [ ] Transition reversal on back navigation
- [ ] Transition observers for analytics
- [ ] Per-route transition configuration
- [ ] Adaptive transitions (platform-aware)

## Learn More

- See [Flutter's Page Route Animation Cookbook](https://docs.flutter.dev/cookbook/animation/page-route-animation) for inspiration
- Check out the main navigation demo: `examples/navigation-demo`
- Read the animation framework docs: `pkg/core/animation/`

## Tips

1. **Choose the right transition** - Use platform-appropriate transitions for better UX
2. **Keep it smooth** - Default 300ms duration works well for most cases
3. **Don't overdo it** - Sometimes `TransitionNone` is the right choice
4. **Test on target platform** - Transitions may feel different on different devices
5. **Consider accessibility** - Some users may prefer reduced motion

## Contributing

Found a bug or want to add a new transition type? Contributions are welcome!
