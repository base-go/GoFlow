package main

import (
	"fmt"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
	"github.com/base-go/GoFlow/pkg/navigation"
)

func main() {
	// Register named routes
	routes := map[string]navigation.RouteBuilder{
		"/":         func() goflow.Widget { return NewHomePage() },
		"/second":   func() goflow.Widget { return NewSecondPage() },
		"/third":    func() goflow.Widget { return NewThirdPage() },
		"/settings": func() goflow.Widget { return NewSettingsPage() },
	}

	// Create app with navigation
	app := navigation.GetMaterialApp(
		NewHomePage(),
		routes,
		"Navigation Demo",
	)

	// Run the app
	goflow.RunApp(app)
}

// HomePage - demonstrates basic navigation
type HomePage struct {
	goflow.BaseWidget
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (h *HomePage) Build(context goflow.BuildContext) goflow.Widget {
	// Create navigation buttons
	buttons := []goflow.Widget{
		createButton("Go to Second Page", func() {
			navigation.Get.To(NewSecondPage(), navigation.TransitionSlideRight)
		}),
		createButton("Go to Second (Named)", func() {
			navigation.Get.ToNamed("/second", navigation.TransitionFade)
		}),
		createButton("Show Dialog", func() {
			showExampleDialog()
		}),
		createButton("Show Bottom Sheet", func() {
			showExampleBottomSheet()
		}),
		createButton("Show Snackbar", func() {
			navigation.ShowSnackbar("Hello from Snackbar!", 3000)
		}),
		createButton("Show Alert Dialog", func() {
			navigation.ShowAlertDialog(
				"Confirm Action",
				"Are you sure you want to proceed?",
				[]navigation.DialogAction{
					{
						Label: "Cancel",
						OnPress: func() {
							navigation.Get.CloseDialog()
							navigation.ShowSnackbar("Cancelled")
						},
					},
					{
						Label: "OK",
						OnPress: func() {
							navigation.Get.CloseDialog()
							navigation.ShowSnackbar("Confirmed!", 2000)
						},
						Primary: true,
					},
				},
			)
		}),
	}

	// Create page layout
	title := widgets.NewText("Home Page")
	subtitle := widgets.NewText("GetX-Style Navigation Demo")

	children := []goflow.Widget{
		title,
		subtitle,
		widgets.NewContainer(), // Spacer
	}
	children = append(children, buttons...)

	column := widgets.NewColumn(children...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (h *HomePage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(h)
}

// SecondPage - demonstrates navigation stack
type SecondPage struct {
	goflow.BaseWidget
}

func NewSecondPage() *SecondPage {
	return &SecondPage{}
}

func (s *SecondPage) Build(context goflow.BuildContext) goflow.Widget {
	buttons := []goflow.Widget{
		createButton("Go to Third Page", func() {
			navigation.Get.To(NewThirdPage(), navigation.TransitionSlideLeft)
		}),
		createButton("Replace with Third (Off)", func() {
			navigation.Get.Off(NewThirdPage())
		}),
		createButton("Go to Settings", func() {
			navigation.Get.ToNamed("/settings", navigation.TransitionZoom)
		}),
		createButton("Back to Home", func() {
			navigation.Get.Back()
		}),
	}

	title := widgets.NewText("Second Page")
	subtitle := widgets.NewText("Navigation Stack Demo")

	children := []goflow.Widget{title, subtitle, widgets.NewContainer()}
	children = append(children, buttons...)

	column := widgets.NewColumn(children...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (s *SecondPage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(s)
}

// ThirdPage - demonstrates advanced navigation
type ThirdPage struct {
	goflow.BaseWidget
}

func NewThirdPage() *ThirdPage {
	return &ThirdPage{}
}

func (t *ThirdPage) Build(context goflow.BuildContext) goflow.Widget {
	buttons := []goflow.Widget{
		createButton("Back", func() {
			navigation.Get.Back()
		}),
		createButton("Back to Home (OffAll)", func() {
			navigation.Get.OffAllNamed("/")
		}),
		createButton("Show Modal Bottom Sheet", func() {
			content := createBottomSheetContent()
			navigation.ShowModalBottomSheet(content, "Options")
		}),
		createButton("Can Pop? Check", func() {
			canPop := navigation.Get.CanPop()
			msg := fmt.Sprintf("Can pop: %v", canPop)
			navigation.ShowSnackbar(msg)
		}),
	}

	title := widgets.NewText("Third Page")
	subtitle := widgets.NewText("Advanced Navigation")

	children := []goflow.Widget{title, subtitle, widgets.NewContainer()}
	children = append(children, buttons...)

	column := widgets.NewColumn(children...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (t *ThirdPage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(t)
}

// SettingsPage - demonstrates named routes
type SettingsPage struct {
	goflow.BaseWidget
}

func NewSettingsPage() *SettingsPage {
	return &SettingsPage{}
}

func (s *SettingsPage) Build(context goflow.BuildContext) goflow.Widget {
	buttons := []goflow.Widget{
		createButton("Back", func() {
			navigation.Get.Back()
		}),
		createButton("Home", func() {
			navigation.Get.OffAllNamed("/", navigation.TransitionFade)
		}),
	}

	title := widgets.NewText("Settings Page")
	subtitle := widgets.NewText("Named Route Example")

	children := []goflow.Widget{title, subtitle, widgets.NewContainer()}
	children = append(children, buttons...)

	column := widgets.NewColumn(children...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (s *SettingsPage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(s)
}

// Helper functions

func createButton(label string, onPress func()) goflow.Widget {
	text := widgets.NewText(label)

	button := widgets.NewContainer()
	button.Color = &goflow.Color{R: 100, G: 150, B: 255, A: 1.0}
	button.Padding = goflow.EdgeInsetsSymmetric(12.0, 20.0)
	button.Child = text

	// TODO: Add GestureDetector when available
	// For now, the button is just visual
	// In a real app, wrap with:
	// return widgets.NewGestureDetector(button, widgets.GestureCallbacks{OnTap: onPress})

	// For demonstration, we'll call onPress immediately to show functionality
	// In a real app, this would be triggered by user interaction
	fmt.Printf("Button created: %s (would call onPress on tap)\n", label)

	return button
}

func showExampleDialog() {
	content := widgets.NewText("This is a custom dialog!")

	dialog := widgets.NewContainer()
	dialog.Width = floatPtr(300.0)
	dialog.Height = floatPtr(200.0)
	dialog.Color = &goflow.Color{R: 255, G: 255, B: 255, A: 1.0}
	dialog.Padding = goflow.EdgeInsetsAll(20.0)
	dialog.Child = content

	navigation.Get.Dialog(dialog, true)
}

func showExampleBottomSheet() {
	items := []goflow.Widget{
		widgets.NewText("Option 1"),
		widgets.NewText("Option 2"),
		widgets.NewText("Option 3"),
	}

	column := widgets.NewColumn(items...)

	navigation.Get.BottomSheet(column, true)
}

func createBottomSheetContent() goflow.Widget {
	options := []goflow.Widget{
		createButton("Share", func() {
			navigation.Get.CloseBottomSheet()
			navigation.ShowSnackbar("Shared!")
		}),
		createButton("Edit", func() {
			navigation.Get.CloseBottomSheet()
			navigation.ShowSnackbar("Edit mode")
		}),
		createButton("Delete", func() {
			navigation.Get.CloseBottomSheet()
			navigation.ShowSnackbar("Deleted")
		}),
		createButton("Cancel", func() {
			navigation.Get.CloseBottomSheet()
		}),
	}

	return widgets.NewColumn(options...)
}

func floatPtr(f float64) *float64 {
	return &f
}
