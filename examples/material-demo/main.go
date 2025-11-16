package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/ui/material"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// MaterialDemoApp demonstrates Material Design widgets
type MaterialDemoApp struct {
	goflow.BaseWidget
}

func (m *MaterialDemoApp) Build(context goflow.BuildContext) goflow.Widget {
	theme := material.DefaultLightTheme()

	return &widgets.Column{
		Children: []goflow.Widget{
			// Material AppBar
			material.NewAppBar(&widgets.Text{
				Data: "Material Design Demo",
			}),

			// Theme info card
			material.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: "Material Design Components",
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
			material.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: "Material Buttons",
					},

					// Filled button
					material.NewButton(&widgets.Text{
						Data: "Filled Button",
					}, func() {
						fmt.Println("Filled button clicked!")
					}),

					// Text button
					material.NewTextButton("Text Button", func() {
						fmt.Println("Text button clicked!")
					}),

					// Outlined button
					material.NewOutlinedButton("Outlined Button", func() {
						fmt.Println("Outlined button clicked!")
					}),
				},
			}),

			// Additional cards
			material.NewCard(&widgets.Text{
				Data: "Material Design is Google's design system for Android, Web, and cross-platform apps.",
			}),
		},
	}
}

func main() {
	fmt.Println("GoFlow Material Design Demo")
	fmt.Println("This example demonstrates Material Design widgets.")
	fmt.Println()

	app := &MaterialDemoApp{}
	_ = app // In a real app, this would be passed to the GoFlow runtime
}
