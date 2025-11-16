# GoFlow Testing Framework

A comprehensive testing framework for GoFlow widgets and applications, providing widget testing, golden tests, integration testing, and gesture simulation.

## Features

### 1. Widget Testing
Test individual widgets in isolation with full control over the rendering environment.

```go
func TestMyWidget(t *testing.T) {
    wt := testing.NewWidgetTester(t)

    widget := NewMyWidget()
    wt.PumpWidget(widget)

    // Verify widget tree
    found := wt.VerifyElement(func(e goflow.Element) bool {
        return e.GetWidgetType() == "MyWidget"
    })
    assert.True(t, found)
}
```

**Key Methods:**
- `PumpWidget(widget)` - Build and render a widget
- `Pump(duration)` - Advance time and trigger frame callbacks
- `PumpAndSettle(timeout)` - Wait for all animations to complete
- `VerifyElement(predicate)` - Check if an element exists
- `FindElement(predicate)` - Find an element in the tree

### 2. Golden Tests (Screenshot Comparison)
Visual regression testing through screenshot comparison.

```go
func TestButtonAppearance(t *testing.T) {
    wt := testing.NewWidgetTester(t)
    gt := testing.NewGoldenTester(t, nil)

    button := NewButton("Submit")
    wt.PumpWidget(button)

    img := wt.GetCanvas().GetImage()
    gt.MatchesGolden("button_default", img)
}
```

**Update Golden Files:**
```bash
UPDATE_GOLDENS=1 go test
```

**Configuration:**
```go
opts := &testing.GoldenTestOptions{
    GoldenDir:     "testdata/golden",
    UpdateGoldens: false,
    Threshold:     0.01,  // 1% difference allowed
    DiffDir:       "testdata/golden/diffs",
}
gt := testing.NewGoldenTester(t, opts)
```

### 3. Integration Testing
End-to-end workflow testing with structured scenarios.

```go
func TestLoginFlow(t *testing.T) {
    it := testing.NewIntegrationTester(t, nil)

    scenario := &testing.Scenario{
        Name: "User Login",
        Setup: func(it *testing.IntegrationTester) error {
            return it.PumpWidget(NewLoginScreen())
        },
        Steps: []testing.Step{
            {
                Name: "Enter username",
                Action: func(it *testing.IntegrationTester) error {
                    return it.EnterText("user@example.com")
                },
                Screenshot: "login_username",
            },
            {
                Name: "Click login",
                Action: func(it *testing.IntegrationTester) error {
                    return it.Tap(200, 300)
                },
                Verify: func(it *testing.IntegrationTester) error {
                    return it.AssertElementExists(func(e goflow.Element) bool {
                        return e.GetWidgetType() == "HomeScreen"
                    })
                },
            },
        },
        Timeout: time.Second * 30,
    }

    it.RunScenario(scenario)
}
```

**Scenario Structure:**
- **Setup** - Initialize the test environment
- **Steps** - Sequential test actions with verification
- **Teardown** - Clean up resources
- **Timeout** - Maximum execution time

### 4. Gesture Simulation
Realistic user interaction simulation.

```go
func TestSwipeGesture(t *testing.T) {
    wt := testing.NewWidgetTester(t)
    gs := testing.NewGestureSimulator(wt)

    wt.PumpWidget(NewScrollableList())

    config := &testing.GestureConfig{
        Duration: time.Millisecond * 300,
        Steps:    10,
    }

    gs.Swipe(400, 300, testing.SwipeUp, 200, config)
    wt.PumpAndSettle(time.Second)
}
```

**Available Gestures:**

| Gesture | Description | Example |
|---------|-------------|---------|
| Tap | Single tap | `gs.Tap(x, y)` |
| DoubleTap | Two quick taps | `gs.DoubleTap(x, y)` |
| LongPress | Press and hold | `gs.LongPress(x, y, duration)` |
| Drag | Drag from point A to B | `gs.Drag(x1, y1, x2, y2, config)` |
| Swipe | Directional swipe | `gs.Swipe(x, y, direction, distance, config)` |
| Fling | Swipe with velocity | `gs.Fling(x, y, vx, vy, config)` |
| Pinch | Two-finger zoom | `gs.Pinch(cx, cy, startDist, endDist, config)` |
| Scroll | Mouse wheel scroll | `gs.Scroll(x, y, dx, dy)` |
| Hover | Mouse hover | `gs.Hover(x, y)` |

## Installation

```bash
go get github.com/base-go/GoFlow/pkg/testing
```

## Quick Start

### Basic Widget Test

```go
package mywidget_test

import (
    "testing"
    "github.com/base-go/GoFlow/pkg/testing"
)

func TestMyWidget(t *testing.T) {
    wt := testing.NewWidgetTester(t)

    // Build your widget
    widget := NewMyWidget()
    if err := wt.PumpWidget(widget); err != nil {
        t.Fatal(err)
    }

    // Test interactions
    wt.Tap(100, 100)
    wt.PumpAndSettle(time.Second)

    // Verify results
    found := wt.VerifyElement(func(e goflow.Element) bool {
        return e.GetWidgetType() == "MyWidget"
    })

    if !found {
        t.Error("Widget not found")
    }
}
```

### Visual Regression Test

```go
func TestButtonVisuals(t *testing.T) {
    wt := testing.NewWidgetTester(t)
    gt := testing.NewGoldenTester(t, nil)

    // Default state
    wt.PumpWidget(NewButton("Submit"))
    gt.MatchesGolden("button_default", wt.GetCanvas().GetImage())

    // Hover state
    wt.GetGestureSimulator().Hover(100, 50)
    wt.Pump(time.Millisecond * 16)
    gt.MatchesGolden("button_hover", wt.GetCanvas().GetImage())
}
```

### Integration Test

```go
func TestFullWorkflow(t *testing.T) {
    it := testing.NewIntegrationTester(t, nil)

    it.PumpWidget(NewApp())

    // Navigate through app
    it.Tap(100, 100)  // Click menu
    it.PumpAndSettle()
    it.TakeScreenshot("menu_open")

    it.Tap(200, 150)  // Click item
    it.PumpAndSettle()
    it.TakeScreenshot("item_selected")

    // Verify final state
    it.AssertElementExists(func(e goflow.Element) bool {
        return e.GetWidgetType() == "DetailView"
    })
}
```

## Best Practices

### 1. Use PumpAndSettle for Animations
```go
// ✅ Good
wt.Tap(100, 100)
wt.PumpAndSettle(time.Second * 5)

// ❌ Bad - arbitrary delays
wt.Tap(100, 100)
time.Sleep(time.Second * 2)
```

### 2. Descriptive Golden File Names
```go
// ✅ Good
gt.MatchesGolden("button_primary_hover_state", img)
gt.MatchesGolden("dialog_error_message", img)

// ❌ Bad
gt.MatchesGolden("test1", img)
gt.MatchesGolden("screenshot", img)
```

### 3. Break Complex Tests into Scenarios
```go
// ✅ Good
scenario := &testing.Scenario{
    Name: "Checkout Flow",
    Steps: []testing.Step{
        {Name: "Add to cart", Action: ...},
        {Name: "Enter shipping", Action: ...},
        {Name: "Enter payment", Action: ...},
        {Name: "Confirm order", Action: ...},
    },
}

// ❌ Bad - monolithic test function
func TestCheckout(t *testing.T) {
    // 300 lines of sequential operations...
}
```

### 4. Use WaitFor Instead of Sleep
```go
// ✅ Good
it.WaitFor(func() bool {
    return dataLoaded
}, time.Second * 5)

// ❌ Bad
time.Sleep(time.Second * 3)
```

### 5. Clean Up in Teardown
```go
scenario := &testing.Scenario{
    Setup: func(it *testing.IntegrationTester) error {
        db = openDatabase()
        return it.PumpWidget(NewApp())
    },
    Teardown: func(it *testing.IntegrationTester) error {
        db.Close()
        return nil
    },
}
```

## Testing Patterns

### Widget Unit Test
Focus on a single widget in isolation.

```go
func TestCounter(t *testing.T) {
    wt := testing.NewWidgetTester(t)
    counter := NewCounter(0)
    wt.PumpWidget(counter)

    // Click increment button
    wt.Tap(100, 50)
    wt.Pump(time.Millisecond * 16)

    // Verify count increased
    // (implementation specific)
}
```

### Integration Test Pattern
Test multiple widgets working together.

```go
func TestForm(t *testing.T) {
    it := testing.NewIntegrationTester(t, nil)
    it.PumpWidget(NewForm())

    // Fill out form
    it.Tap(100, 50)  // Focus first field
    it.EnterText("John")

    it.Tap(100, 100)  // Focus second field
    it.EnterText("Doe")

    // Submit
    it.Tap(100, 200)
    it.PumpAndSettle()

    // Verify
    it.AssertElementExists(func(e goflow.Element) bool {
        return e.GetWidgetType() == "SuccessMessage"
    })
}
```

### Visual Regression Pattern
Compare visual output across changes.

```go
func TestTheme(t *testing.T) {
    wt := testing.NewWidgetTester(t)
    gt := testing.NewGoldenTester(t, nil)

    // Test each theme variant
    themes := []string{"light", "dark", "high_contrast"}

    for _, theme := range themes {
        wt.PumpWidget(NewThemedButton(theme))
        img := wt.GetCanvas().GetImage()
        gt.MatchesGolden("button_"+theme, img)
    }
}
```

## Advanced Usage

### Custom Canvas Size
```go
wt := testing.NewWidgetTester(t)
wt.SetSize(1920, 1080)  // Desktop
wt.PumpWidget(widget)
```

### Multiple Scenarios
```go
it := testing.NewIntegrationTester(t, nil)

scenarios := []*testing.Scenario{
    createLoginScenario(),
    createCheckoutScenario(),
    createSettingsScenario(),
}

for _, scenario := range scenarios {
    it.AddScenario(scenario)
}

it.RunAllScenarios()
```

### Event Inspection
```go
wt := testing.NewWidgetTester(t)
wt.PumpWidget(widget)

wt.Tap(100, 100)

events := wt.GetEvents()
for _, event := range events {
    t.Logf("Event: %v", event.Type())
}
```

## Troubleshooting

### Golden Tests Failing
1. Visually inspect diff image in `testdata/golden/diffs/`
2. If change is intentional, update goldens:
   ```bash
   UPDATE_GOLDENS=1 go test
   ```
3. Adjust threshold if minor rendering differences:
   ```go
   opts := &testing.GoldenTestOptions{Threshold: 0.02}
   ```

### PumpAndSettle Timeout
- Increase timeout: `wt.PumpAndSettle(time.Second * 10)`
- Check for infinite animations
- Verify async operations complete

### Element Not Found
- Use `wt.FindElement()` to debug tree structure
- Check element type strings match exactly
- Ensure widget has been built with `PumpWidget()`

## Examples

See `examples_test.go` for comprehensive examples of all testing patterns.

## License

MIT License - see LICENSE file for details
