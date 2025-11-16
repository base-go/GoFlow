package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/ui/adaptive"
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// AdaptiveDemoApp demonstrates platform-adaptive widgets
type AdaptiveDemoApp struct {
	goflow.BaseWidget
}

func (a *AdaptiveDemoApp) Build(context goflow.BuildContext) goflow.Widget {
	platform := goflow.GetPlatform()

	// Create adaptive widgets that will look native on each platform
	return &widgets.Column{
		Children: []goflow.Widget{
			// Adaptive AppBar - will be Material AppBar or Cupertino NavigationBar
			adaptive.NewAppBar(&widgets.Text{
				Data: "Adaptive Demo",
			}),

			// Platform info card
			adaptive.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: fmt.Sprintf("Current Platform: %s", platform.String()),
					},
					&widgets.Text{
						Data: fmt.Sprintf("Design System: %s", platform.DefaultTheme()),
					},
				},
			}),

			// Buttons demo card
			adaptive.NewCard(&widgets.Column{
				Children: []goflow.Widget{
					&widgets.Text{
						Data: "Adaptive Buttons",
					},

					// Filled button
					adaptive.NewButton("Filled Button", func() {
						fmt.Println("Filled button clicked!")
					}),

					// Text button
					adaptive.NewTextButton("Text Button", func() {
						fmt.Println("Text button clicked!")
					}),

					// Outlined button (Material only, will be gray on iOS)
					&adaptive.Button{
						Text:  "Outlined Button",
						Style: adaptive.ButtonStyleOutlined,
						OnPressed: func() {
							fmt.Println("Outlined button clicked!")
						},
					},
				},
			}),

			// Info card
			adaptive.NewCard(&widgets.Text{
				Data: "This UI adapts automatically to your platform. " +
					"On iOS/macOS it uses Cupertino widgets, " +
					"on other platforms it uses Material Design.",
			}),
		},
	}
}

func main() {
	fmt.Println("GoFlow Adaptive Widget Demo")
	fmt.Printf("Platform: %s\n", goflow.GetPlatform().String())
	fmt.Printf("Design System: %s\n", goflow.GetPlatform().DefaultTheme())
	fmt.Println()
	fmt.Println("This example demonstrates platform-adaptive widgets.")
	fmt.Println("The same code automatically switches between Material and Cupertino styles.")
	fmt.Println()

	// Uncomment to test different platforms:
	// goflow.SetPlatform(goflow.PlatformIOS)
	// goflow.SetPlatform(goflow.PlatformAndroid)

	app := &AdaptiveDemoApp{}
	_ = app // In a real app, this would be passed to the GoFlow runtime
}
