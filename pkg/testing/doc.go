/*
Package testing provides comprehensive testing utilities for GoFlow widgets and applications.

# Overview

The testing package includes four main components:

1. Widget Testing - Test individual widgets in isolation
2. Golden Tests - Screenshot comparison testing
3. Integration Testing - End-to-end workflow testing
4. Gesture Simulation - Simulate user interactions

# Widget Testing

WidgetTester provides utilities for testing widgets in a controlled environment:

	func TestMyWidget(t *testing.T) {
		wt := testing.NewWidgetTester(t)

		// Build the widget
		widget := NewMyWidget()
		if err := wt.PumpWidget(widget); err != nil {
			t.Fatal(err)
		}

		// Verify the widget tree
		found := wt.VerifyElement(func(e goflow.Element) bool {
			return e.GetWidgetType() == "MyWidget"
		})
		if !found {
			t.Error("MyWidget not found in tree")
		}

		// Simulate interactions
		if err := wt.Tap(100, 100); err != nil {
			t.Fatal(err)
		}

		// Wait for animations to settle
		if err := wt.PumpAndSettle(time.Second * 5); err != nil {
			t.Fatal(err)
		}
	}

# Golden Tests

GoldenTester provides screenshot comparison testing:

	func TestButtonAppearance(t *testing.T) {
		wt := testing.NewWidgetTester(t)
		gt := testing.NewGoldenTester(t, nil)

		// Build widget
		button := NewButton("Click Me")
		wt.PumpWidget(button)

		// Compare with golden
		img := wt.GetCanvas().GetImage()
		if !gt.MatchesGolden("button_default", img) {
			t.Error("Button appearance doesn't match golden")
		}
	}

Update golden files by setting the UPDATE_GOLDENS environment variable:

	UPDATE_GOLDENS=1 go test

# Integration Testing

IntegrationTester provides end-to-end workflow testing with scenarios:

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
					Screenshot: "login_username_entered",
				},
				{
					Name: "Enter password",
					Action: func(it *testing.IntegrationTester) error {
						return it.EnterText("password123")
					},
				},
				{
					Name: "Click login button",
					Action: func(it *testing.IntegrationTester) error {
						return it.Tap(200, 300)
					},
					Verify: func(it *testing.IntegrationTester) error {
						return it.WaitForElement(func(e goflow.Element) bool {
							return e.GetWidgetType() == "HomeScreen"
						}, time.Second * 5)
					},
					Screenshot: "login_success",
				},
			},
		}

		if err := it.RunScenario(scenario); err != nil {
			t.Fatal(err)
		}
	}

# Gesture Simulation

GestureSimulator provides realistic gesture simulation:

	func TestSwipeGesture(t *testing.T) {
		wt := testing.NewWidgetTester(t)
		gs := testing.NewGestureSimulator(wt)

		// Build widget
		wt.PumpWidget(NewScrollableList())

		// Simulate swipe
		config := &testing.GestureConfig{
			Duration: time.Millisecond * 300,
			Steps:    10,
		}
		if err := gs.Swipe(400, 300, testing.SwipeUp, 200, config); err != nil {
			t.Fatal(err)
		}

		// Verify scroll occurred
		wt.PumpAndSettle(time.Second)
	}

Available gestures:
- Tap, DoubleTap, LongPress
- Drag, Swipe, Fling
- Pinch (zoom in/out)
- Scroll, Hover

# Best Practices

1. Use PumpAndSettle() to wait for animations and async operations
2. Take screenshots at key points for visual regression testing
3. Break complex tests into scenarios with clear steps
4. Use WaitFor() instead of arbitrary sleep delays
5. Clean up resources in scenario Teardown functions
6. Name golden files descriptively (e.g., "button_hover_state", "dialog_error")

# Testing Patterns

Widget Unit Test:

	func TestWidget(t *testing.T) {
		wt := testing.NewWidgetTester(t)
		wt.PumpWidget(NewMyWidget())
		// Assertions...
	}

Visual Regression Test:

	func TestVisual(t *testing.T) {
		wt := testing.NewWidgetTester(t)
		gt := testing.NewGoldenTester(t, nil)
		wt.PumpWidget(NewMyWidget())
		gt.MatchesGolden("my_widget", wt.GetCanvas().GetImage())
	}

Integration Test:

	func TestIntegration(t *testing.T) {
		it := testing.NewIntegrationTester(t, nil)
		it.AddScenario(&testing.Scenario{...})
		it.RunAllScenarios()
	}

Gesture Test:

	func TestGesture(t *testing.T) {
		wt := testing.NewWidgetTester(t)
		gs := testing.NewGestureSimulator(wt)
		wt.PumpWidget(NewMyWidget())
		gs.Tap(100, 100)
		// Verify interaction...
	}
*/
package testing
