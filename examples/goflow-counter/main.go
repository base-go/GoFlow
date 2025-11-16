package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// CounterApp is the root widget for our counter application
type CounterApp struct {
	goflow.BaseWidget
}

// Build creates the widget tree
func (a *CounterApp) Build(context goflow.BuildContext) goflow.Widget {
	return &CounterPage{}
}

// CounterPage is the main page with a counter
type CounterPage struct {
	goflow.BaseWidget
	counter *signals.Signal[int]
}

// NewCounterPage creates a new counter page
func NewCounterPage() *CounterPage {
	return &CounterPage{
		counter: signals.New(0),
	}
}

// Build creates the UI for the counter page
func (p *CounterPage) Build(context goflow.BuildContext) goflow.Widget {
	// Get the current counter value
	count := p.counter.Get()

	// Build UI
	return widgets.NewCenter(
		&widgets.Column{
			Children: []goflow.Widget{
				widgets.NewTextWithStyle(
					"GoFlow Counter App",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				&Container{
					Height: float64Ptr(20),
				},
				widgets.NewText(fmt.Sprintf("Count: %d", count)),
				&Container{
					Height: float64Ptr(10),
				},
				widgets.NewText("(Click + to increment)"),
			},
			MainAxisAlign:  widgets.MainAxisCenter,
			CrossAxisAlign: widgets.CrossAxisCenter,
		},
	)
}

// Increment increments the counter
func (p *CounterPage) Increment() {
	p.counter.Update(func(v int) int {
		return v + 1
	})
}

// Container is a simple container widget
type Container struct {
	goflow.BaseWidget
	Width  *float64
	Height *float64
}

// Build returns nil (just for spacing)
func (c *Container) Build(context goflow.BuildContext) goflow.Widget {
	// In a full implementation, this would create a sized box
	return nil
}

func float64Ptr(f float64) *float64 {
	return &f
}

func main() {
	fmt.Println("=== GoFlow Counter Example ===\n")

	// Create the counter page
	counterPage := NewCounterPage()

	// Set up a signal effect to rebuild when counter changes
	dispose := signals.NewEffect(func() {
		// Access the counter to create a dependency
		_ = counterPage.counter.Get()

		// In a real implementation, this would mark the widget as needing rebuild
		fmt.Println("Counter changed - UI would rebuild here")
	})
	defer dispose()

	// Create and run the app
	app := &CounterApp{}
	goflow.RunApp(app)

	// Simulate some interactions
	fmt.Println("\n=== Simulating User Interactions ===")
	fmt.Println("\nUser clicks '+' button...")
	counterPage.Increment()

	fmt.Println("\nUser clicks '+' button again...")
	counterPage.Increment()

	fmt.Println("\nUser clicks '+' button again...")
	counterPage.Increment()

	fmt.Printf("\nFinal counter value: %d\n", counterPage.counter.Get())
}
