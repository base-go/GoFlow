package main

import (
	"fmt"

	"github.com/base-go/GoFlow/cupertino"
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// CupertinoDemoApp demonstrates Cupertino (iOS/macOS) widgets
type CupertinoDemoApp struct {
	goflow.BaseWidget
}

func (c *CupertinoDemoApp) Build(context goflow.BuildContext) goflow.Widget {
	theme := cupertino.DefaultLightTheme()

	return &widgets.Column{
		Children: []goflow.Widget{
			// Cupertino NavigationBar
			cupertino.NewNavigationBar(&widgets.Text{
				Data: "Cupertino Demo",
			}),

			// Theme info card
			cupertino.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: "Cupertino Components (iOS/macOS)",
					},
					&widgets.Text{
						Data: fmt.Sprintf("Primary Color: #%02x%02x%02x",
							theme.PrimaryColor.R,
							theme.PrimaryColor.G,
							theme.PrimaryColor.B),
					},
				},
			}),

			// Buttons showcase
			cupertino.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: "Cupertino Buttons",
					},

					// Filled button
					cupertino.NewButton(&widgets.Text{
						Data: "Filled Button",
					}, func() {
						fmt.Println("Filled button clicked!")
					}),

					// Text button
					cupertino.NewTextButton("Text Button", func() {
						fmt.Println("Text button clicked!")
					}),

					// Gray button
					cupertino.NewGrayButton("Gray Button", func() {
						fmt.Println("Gray button clicked!")
					}),
				},
			}),

			// Additional cards
			cupertino.NewCard(&widgets.Text{
				Data: "Cupertino is Apple's design language for iOS and macOS applications.",
			}),
		},
	}
}

func main() {
	fmt.Println("GoFlow Cupertino Demo")
	fmt.Println("This example demonstrates Cupertino (iOS/macOS) widgets.")
	fmt.Println()

	app := &CupertinoDemoApp{}
	_ = app // In a real app, this would be passed to the GoFlow runtime
}
