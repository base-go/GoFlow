package main

import (
	"fmt"

	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/material"
	"github.com/base-go/GoFlow/widgets"
)

// WidgetsShowcaseApp demonstrates all available widgets
type WidgetsShowcaseApp struct {
	goflow.BaseWidget
}

func (w *WidgetsShowcaseApp) Build(context goflow.BuildContext) goflow.Widget {
	return &material.Scaffold{
		AppBar: material.NewAppBar(&widgets.Text{
			Data: "GoFlow Widgets Showcase",
		}),
		Body: &widgets.SingleChildScrollView{
			Child: &widgets.Column{
				Children: []goflow.Widget{
					// Layout Widgets Section
					buildSectionHeader("Layout Widgets"),

					// Stack example
					material.NewCard(&widgets.Container{
						Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
						Child: &widgets.Column{
							Children: []goflow.Widget{
								&widgets.Text{Data: "Stack Example"},
								&widgets.Stack{
									Children: []goflow.Widget{
										&widgets.Container{
											Width:  floatPtr(200),
											Height: floatPtr(100),
											Color:  goflow.NewColor(200, 200, 255, 255),
										},
										&widgets.Align{
											Alignment: widgets.AlignmentCenter,
											Child: &widgets.Text{
												Data: "Centered Text",
											},
										},
									},
									Fit: widgets.StackFitLoose,
								},
							},
						},
					}),

					// Buttons Section
					buildSectionHeader("Buttons"),

					material.NewCard(&widgets.Container{
						Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
						Child: &widgets.Row{
							Children: []goflow.Widget{
								material.NewButton(&widgets.Text{Data: "Filled"}, func() {
									fmt.Println("Filled button clicked")
								}),
								material.NewTextButton("Text", func() {
									fmt.Println("Text button clicked")
								}),
								material.NewOutlinedButton("Outlined", func() {
									fmt.Println("Outlined button clicked")
								}),
							},
							MainAxisAlignment: widgets.MainAxisSpaceAround,
						},
					}),

					// Form Widgets Section
					buildSectionHeader("Form Widgets"),

					material.NewCard(&widgets.Container{
						Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
						Child: &widgets.Column{
							Children: []goflow.Widget{
								&widgets.Text{Data: "Checkbox & Switch"},
								material.NewCheckbox(true, func(b bool) {
									fmt.Println("Checkbox:", b)
								}),
								material.NewSwitch(true, func(b bool) {
									fmt.Println("Switch:", b)
								}),
							},
						},
					}),

					// Icons Section
					buildSectionHeader("Icons"),

					material.NewCard(&widgets.Container{
						Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
						Child: &widgets.Row{
							Children: []goflow.Widget{
								widgets.NewIcon(widgets.IconHome),
								widgets.NewIcon(widgets.IconFavorite),
								widgets.NewIcon(widgets.IconSearch),
								widgets.NewIcon(widgets.IconSettings),
							},
							MainAxisAlignment: widgets.MainAxisSpaceAround,
						},
					}),

					// List Tiles Section
					buildSectionHeader("List Tiles"),

					material.NewListTile(&widgets.Text{Data: "List Tile 1"}),
					material.NewListTile(&widgets.Text{Data: "List Tile 2"}),
					material.NewListTile(&widgets.Text{Data: "List Tile 3"}),
				},
			},
		},
		FloatingActionButton: material.NewFloatingActionButton(
			widgets.NewIcon(widgets.IconAdd),
			func() {
				fmt.Println("FAB clicked!")
			},
		),
	}
}

func buildSectionHeader(title string) goflow.Widget {
	return &widgets.Container{
		Padding: goflow.NewEdgeInsets(16, 16, 16, 8),
		Child: &widgets.Text{
			Data: title,
			Style: &goflow.TextStyle{
				FontSize: 18.0,
				Color:    goflow.NewColor(33, 150, 243, 255),
			},
		},
	}
}

func floatPtr(f float64) *float64 {
	return &f
}

func main() {
	fmt.Println("GoFlow Widgets Showcase")
	fmt.Println("This example demonstrates all available widgets in GoFlow")
	fmt.Println()

	app := &WidgetsShowcaseApp{}
	_ = app // In a real app, this would be passed to the GoFlow runtime
}
