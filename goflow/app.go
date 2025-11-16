package goflow

import (
	"fmt"
	"time"
)

// App represents the root of a GoFlow application
type App struct {
	rootElement Element
	rootWidget  Widget
	canvas      Canvas
	constraints *Constraints
	renderRoot  RenderObject
}

// NewApp creates a new application
func NewApp(rootWidget Widget, width, height float64) *App {
	return &App{
		rootWidget:  rootWidget,
		canvas:      NewMockCanvas(width, height),
		constraints: TightConstraints(width, height),
	}
}

// RunApp runs a GoFlow application
// This is the main entry point for GoFlow apps
func RunApp(rootWidget Widget) {
	fmt.Println("GoFlow Application Starting...")
	fmt.Println("================================")

	// Create app with default size
	app := NewApp(rootWidget, 800, 600)

	// Build the initial widget tree
	app.Build()

	// Perform initial layout
	app.Layout()

	// Render
	app.Render()

	fmt.Println("\n================================")
	fmt.Println("GoFlow Application Running")
	fmt.Println("(In a real implementation, this would start the event loop)")
	fmt.Println("(For now, this is a demonstration of the framework)")

	// In a real implementation, this would start the platform event loop
	// For now, we'll just show that it works
	time.Sleep(100 * time.Millisecond)
}

// Build builds the widget tree
func (a *App) Build() {
	fmt.Println("\n[Build Phase]")
	fmt.Println("Building widget tree...")

	// Create the root element
	a.rootElement = a.rootWidget.CreateElement()
	a.rootElement.Mount(nil, nil)

	fmt.Println("Widget tree built successfully")
}

// Layout performs layout on the render tree
func (a *App) Layout() {
	fmt.Println("\n[Layout Phase]")
	fmt.Println("Performing layout...")

	// Get the render object from the root element
	a.renderRoot = a.findRenderObject(a.rootElement)

	if a.renderRoot != nil {
		a.renderRoot.Layout(a.constraints)
		size := a.renderRoot.GetSize()
		fmt.Printf("Root size: %.0f×%.0f\n", size.Width, size.Height)
	} else {
		fmt.Println("No render object found")
	}
}

// Render renders the widget tree to the canvas
func (a *App) Render() {
	fmt.Println("\n[Render Phase]")
	fmt.Println("Painting to canvas...")

	if a.renderRoot != nil {
		a.renderRoot.Paint(a.canvas, ZeroOffset())
		if mockCanvas, ok := a.canvas.(*MockCanvas); ok {
			ops := mockCanvas.GetOperations()
			fmt.Printf("Canvas operations: %d\n", len(ops))
			for i, op := range ops {
				if i < 10 { // Show first 10 operations
					fmt.Printf("  - %s\n", op)
				}
			}
			if len(ops) > 10 {
				fmt.Printf("  ... and %d more\n", len(ops)-10)
			}
		}
	}
}

// findRenderObject recursively finds the first render object in the tree
func (a *App) findRenderObject(element Element) RenderObject {
	if element == nil {
		return nil
	}

	// Check if this element has a render object
	ro := element.GetRenderObject()
	if ro != nil {
		return ro
	}

	// Check children
	var result RenderObject
	element.VisitChildren(func(child Element) {
		if result == nil {
			result = a.findRenderObject(child)
		}
	})

	return result
}

// SetCanvas sets a custom canvas (useful for testing or different backends)
func (a *App) SetCanvas(canvas Canvas) {
	a.canvas = canvas
}
