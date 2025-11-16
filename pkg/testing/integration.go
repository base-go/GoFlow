package testing

import (
	"fmt"
	"testing"
	"time"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// IntegrationTester provides integration testing utilities
type IntegrationTester struct {
	widgetTester *WidgetTester
	gestures     *GestureSimulator
	golden       *GoldenTester
	t            *testing.T
	scenarios    []*Scenario
}

// Scenario represents a test scenario
type Scenario struct {
	Name        string
	Description string
	Setup       func(*IntegrationTester) error
	Steps       []Step
	Teardown    func(*IntegrationTester) error
	Timeout     time.Duration
}

// Step represents a single step in a test scenario
type Step struct {
	Name        string
	Description string
	Action      func(*IntegrationTester) error
	Verify      func(*IntegrationTester) error
	Screenshot  string // If set, take a screenshot with this name
}

// NewIntegrationTester creates a new integration tester
func NewIntegrationTester(t *testing.T, goldenOpts *GoldenTestOptions) *IntegrationTester {
	wt := NewWidgetTester(t)
	return &IntegrationTester{
		widgetTester: wt,
		gestures:     NewGestureSimulator(wt),
		golden:       NewGoldenTester(t, goldenOpts),
		t:            t,
		scenarios:    make([]*Scenario, 0),
	}
}

// GetWidgetTester returns the widget tester
func (it *IntegrationTester) GetWidgetTester() *WidgetTester {
	return it.widgetTester
}

// GetGestureSimulator returns the gesture simulator
func (it *IntegrationTester) GetGestureSimulator() *GestureSimulator {
	return it.gestures
}

// GetGoldenTester returns the golden tester
func (it *IntegrationTester) GetGoldenTester() *GoldenTester {
	return it.golden
}

// PumpWidget builds and renders a widget
func (it *IntegrationTester) PumpWidget(widget goflow.Widget) error {
	return it.widgetTester.PumpWidget(widget)
}

// AddScenario adds a test scenario
func (it *IntegrationTester) AddScenario(scenario *Scenario) {
	it.scenarios = append(it.scenarios, scenario)
}

// RunScenario runs a single test scenario
func (it *IntegrationTester) RunScenario(scenario *Scenario) error {
	it.t.Logf("Running scenario: %s", scenario.Name)

	// Set timeout
	timeout := scenario.Timeout
	if timeout == 0 {
		timeout = time.Second * 30 // Default timeout
	}

	done := make(chan error, 1)

	go func() {
		// Run setup
		if scenario.Setup != nil {
			it.t.Logf("  Running setup...")
			if err := scenario.Setup(it); err != nil {
				done <- fmt.Errorf("setup failed: %w", err)
				return
			}
		}

		// Run steps
		for i, step := range scenario.Steps {
			it.t.Logf("  Step %d: %s", i+1, step.Name)

			// Run action
			if step.Action != nil {
				if err := step.Action(it); err != nil {
					done <- fmt.Errorf("step %d action failed: %w", i+1, err)
					return
				}
			}

			// Take screenshot if requested
			if step.Screenshot != "" {
				img := it.widgetTester.GetCanvas().GetImage()
				if !it.golden.MatchesGolden(step.Screenshot, img) {
					done <- fmt.Errorf("step %d screenshot verification failed", i+1)
					return
				}
			}

			// Run verification
			if step.Verify != nil {
				if err := step.Verify(it); err != nil {
					done <- fmt.Errorf("step %d verification failed: %w", i+1, err)
					return
				}
			}
		}

		// Run teardown
		if scenario.Teardown != nil {
			it.t.Logf("  Running teardown...")
			if err := scenario.Teardown(it); err != nil {
				done <- fmt.Errorf("teardown failed: %w", err)
				return
			}
		}

		done <- nil
	}()

	// Wait for completion or timeout
	select {
	case err := <-done:
		if err != nil {
			it.t.Errorf("Scenario %s failed: %v", scenario.Name, err)
			return err
		}
		it.t.Logf("Scenario %s completed successfully", scenario.Name)
		return nil
	case <-time.After(timeout):
		err := fmt.Errorf("scenario %s timed out after %v", scenario.Name, timeout)
		it.t.Error(err)
		return err
	}
}

// RunAllScenarios runs all registered scenarios
func (it *IntegrationTester) RunAllScenarios() error {
	failures := 0
	for _, scenario := range it.scenarios {
		if err := it.RunScenario(scenario); err != nil {
			failures++
		}
	}

	if failures > 0 {
		return fmt.Errorf("%d scenario(s) failed", failures)
	}

	return nil
}

// VerifyWidgetTree verifies the widget tree structure
func (it *IntegrationTester) VerifyWidgetTree(verifier func(goflow.Element) error) error {
	element := it.widgetTester.GetElement()
	if element == nil {
		return fmt.Errorf("no element to verify")
	}
	return verifier(element)
}

// WaitFor waits for a condition to become true
func (it *IntegrationTester) WaitFor(condition func() bool, timeout time.Duration) error {
	start := time.Now()
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	for {
		if condition() {
			return nil
		}

		if time.Since(start) > timeout {
			return fmt.Errorf("condition not met within timeout %v", timeout)
		}

		select {
		case <-ticker.C:
			it.widgetTester.Pump(time.Millisecond * 16)
		}
	}
}

// WaitForElement waits for an element matching the predicate to appear
func (it *IntegrationTester) WaitForElement(predicate func(goflow.Element) bool, timeout time.Duration) (goflow.Element, error) {
	var element goflow.Element
	err := it.WaitFor(func() bool {
		element = it.widgetTester.FindElement(predicate)
		return element != nil
	}, timeout)

	if err != nil {
		return nil, err
	}

	return element, nil
}

// AssertElementExists asserts that an element exists
func (it *IntegrationTester) AssertElementExists(predicate func(goflow.Element) bool) error {
	if !it.widgetTester.VerifyElement(predicate) {
		return fmt.Errorf("element not found")
	}
	return nil
}

// AssertElementNotExists asserts that an element does not exist
func (it *IntegrationTester) AssertElementNotExists(predicate func(goflow.Element) bool) error {
	if it.widgetTester.VerifyElement(predicate) {
		return fmt.Errorf("element should not exist but was found")
	}
	return nil
}

// Tap performs a tap gesture
func (it *IntegrationTester) Tap(x, y float64) error {
	return it.gestures.Tap(x, y)
}

// Drag performs a drag gesture
func (it *IntegrationTester) Drag(startX, startY, endX, endY float64) error {
	return it.gestures.Drag(startX, startY, endX, endY, nil)
}

// EnterText enters text
func (it *IntegrationTester) EnterText(text string) error {
	return it.widgetTester.EnterText(text)
}

// TakeScreenshot takes a screenshot and compares it with golden
func (it *IntegrationTester) TakeScreenshot(name string) error {
	img := it.widgetTester.GetCanvas().GetImage()
	if !it.golden.MatchesGolden(name, img) {
		return fmt.Errorf("screenshot %s does not match golden", name)
	}
	return nil
}

// Pump advances time and triggers frame callbacks
func (it *IntegrationTester) Pump(duration time.Duration) error {
	return it.widgetTester.Pump(duration)
}

// PumpAndSettle pumps frames until settled
func (it *IntegrationTester) PumpAndSettle() error {
	return it.widgetTester.PumpAndSettle(time.Second * 5)
}
