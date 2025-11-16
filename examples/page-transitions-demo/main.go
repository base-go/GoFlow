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
		"/":      func() goflow.Widget { return NewTransitionsHomePage() },
		"/demo":  func() goflow.Widget { return NewTransitionDemoPage("Demo Page") },
		"/page2": func() goflow.Widget { return NewTransitionDemoPage("Second Page") },
		"/page3": func() goflow.Widget { return NewTransitionDemoPage("Third Page") },
	}

	// Create app with navigation
	app := navigation.GetMaterialApp(
		NewTransitionsHomePage(),
		routes,
		"Page Transitions Demo - GoFlow",
	)

	// Run the app
	goflow.RunApp(app)
}

// TransitionsHomePage - demonstrates all available page transitions
type TransitionsHomePage struct {
	goflow.BaseWidget
}

func NewTransitionsHomePage() *TransitionsHomePage {
	return &TransitionsHomePage{}
}

func (h *TransitionsHomePage) Build(context goflow.BuildContext) goflow.Widget {
	// Create section for basic transitions
	basicTransitions := []goflow.Widget{
		createSectionTitle("Basic Transitions"),
		createTransitionButton("Fade Transition", navigation.TransitionFade),
		createTransitionButton("Slide Right (iOS)", navigation.TransitionSlideRight),
		createTransitionButton("Slide Left", navigation.TransitionSlideLeft),
		createTransitionButton("Slide Up (Material)", navigation.TransitionSlideUp),
		createTransitionButton("Slide Down", navigation.TransitionSlideDown),
		createTransitionButton("Zoom Transition", navigation.TransitionZoom),
		createTransitionButton("No Transition", navigation.TransitionNone),
	}

	// Create section for platform-specific transitions
	platformTransitions := []goflow.Widget{
		createSectionTitle("Platform-Specific"),
		createTransitionButton("Cupertino (iOS Style)", navigation.TransitionCupertino),
		createTransitionButton("Material (Android Style)", navigation.TransitionMaterial),
	}

	// Create info section
	infoSection := []goflow.Widget{
		widgets.NewContainer(), // Spacer
		createInfoBox(),
	}

	// Combine all sections
	allChildren := []goflow.Widget{
		createPageTitle("Page Transitions Demo"),
		createSubtitle("Tap any button to see the transition effect"),
		widgets.NewContainer(), // Spacer
	}
	allChildren = append(allChildren, basicTransitions...)
	allChildren = append(allChildren, widgets.NewContainer()) // Spacer
	allChildren = append(allChildren, platformTransitions...)
	allChildren = append(allChildren, infoSection...)

	column := widgets.NewColumn(allChildren...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (h *TransitionsHomePage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(h)
}

// TransitionDemoPage - a simple page to navigate to
type TransitionDemoPage struct {
	goflow.BaseWidget
	Title string
}

func NewTransitionDemoPage(title string) *TransitionDemoPage {
	return &TransitionDemoPage{Title: title}
}

func (d *TransitionDemoPage) Build(context goflow.BuildContext) goflow.Widget {
	// Navigation buttons
	buttons := []goflow.Widget{
		createActionButton("Back to Transitions", func() {
			navigation.Get.Back()
		}),
		createActionButton("Go to Page 2 (Fade)", func() {
			navigation.Get.ToNamed("/page2", navigation.TransitionFade)
		}),
		createActionButton("Go to Page 3 (Zoom)", func() {
			navigation.Get.ToNamed("/page3", navigation.TransitionZoom)
		}),
		createActionButton("Replace with Home (Slide)", func() {
			navigation.Get.Off(NewTransitionsHomePage(), navigation.TransitionSlideLeft)
		}),
		createActionButton("Clear Stack & Home", func() {
			navigation.Get.OffAllNamed("/", navigation.TransitionFade)
		}),
	}

	// Create info about current page
	info := createDemoPageInfo()

	// Build page content
	children := []goflow.Widget{
		createPageTitle(d.Title),
		createSubtitle("Try different navigation actions"),
		widgets.NewContainer(), // Spacer
		info,
		widgets.NewContainer(), // Spacer
	}
	children = append(children, buttons...)

	column := widgets.NewColumn(children...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (d *TransitionDemoPage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(d)
}

// Helper functions for creating styled widgets

func createPageTitle(text string) goflow.Widget {
	title := widgets.NewText(text)
	// In a real app, you'd style this with larger font, bold, etc.
	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsSymmetric(0, 10.0)
	container.Child = title
	return container
}

func createSubtitle(text string) goflow.Widget {
	subtitle := widgets.NewText(text)
	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsSymmetric(0, 5.0)
	container.Child = subtitle
	return container
}

func createSectionTitle(text string) goflow.Widget {
	title := widgets.NewText(text)
	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsSymmetric(0, 10.0)
	container.Color = &goflow.Color{R: 230, G: 230, B: 230, A: 1.0}
	container.Child = title
	return container
}

func createTransitionButton(label string, transition navigation.Transition) goflow.Widget {
	text := widgets.NewText(label)

	button := widgets.NewContainer()
	button.Color = &goflow.Color{R: 66, G: 165, B: 245, A: 1.0} // Material Blue
	button.Padding = goflow.EdgeInsetsSymmetric(12.0, 20.0)
	var margin goflow.EdgeInsets
	margin.Bottom = 8.0
	button.Margin = &margin
	button.Child = text

	// Log button creation (in real app, this would have gesture detection)
	fmt.Printf("Created transition button: %s with transition type %d\n", label, transition)

	// Note: In a real app with gesture detection:
	// return widgets.NewGestureDetector(button, widgets.GestureCallbacks{
	//     OnTap: func() {
	//         navigation.Get.To(NewTransitionDemoPage("Demo Page"), transition)
	//     },
	// })

	return button
}

func createActionButton(label string, onPress func()) goflow.Widget {
	text := widgets.NewText(label)

	button := widgets.NewContainer()
	button.Color = &goflow.Color{R: 76, G: 175, B: 80, A: 1.0} // Material Green
	button.Padding = goflow.EdgeInsetsSymmetric(12.0, 20.0)
	var margin goflow.EdgeInsets
	margin.Bottom = 8.0
	button.Margin = &margin
	button.Child = text

	fmt.Printf("Created action button: %s\n", label)

	// Note: In a real app with gesture detection:
	// return widgets.NewGestureDetector(button, widgets.GestureCallbacks{OnTap: onPress})

	return button
}

func createInfoBox() goflow.Widget {
	infoText := widgets.NewText(
		"Page Transitions Library\n\n" +
			"This demo showcases GoFlow's page transition library, " +
			"similar to Flutter's page route animations.\n\n" +
			"Features:\n" +
			"• Fade transitions\n" +
			"• Slide transitions (all directions)\n" +
			"• Zoom/Scale transitions\n" +
			"• Platform-specific styles (iOS/Material)\n" +
			"• Custom transition support\n" +
			"• Configurable duration and curves\n\n" +
			"Coming Soon:\n" +
			"• Gesture-based navigation\n" +
			"• Hero animations\n" +
			"• Shared element transitions",
	)

	infoBox := widgets.NewContainer()
	infoBox.Color = &goflow.Color{R: 255, G: 249, B: 196, A: 1.0} // Light yellow
	infoBox.Padding = goflow.EdgeInsetsAll(15.0)
	infoBox.Child = infoText

	return infoBox
}

func createDemoPageInfo() goflow.Widget {
	text := widgets.NewText(
		"This page demonstrates navigation actions.\n" +
			"Notice how the transition animation changes based on the button you tap.",
	)

	container := widgets.NewContainer()
	container.Color = &goflow.Color{R: 232, G: 245, B: 233, A: 1.0} // Light green
	container.Padding = goflow.EdgeInsetsAll(12.0)
	container.Child = text

	return container
}

// Additional example pages for advanced transitions

// TransitionComparisonPage - shows transitions side by side
type TransitionComparisonPage struct {
	goflow.BaseWidget
}

func NewTransitionComparisonPage() *TransitionComparisonPage {
	return &TransitionComparisonPage{}
}

func (t *TransitionComparisonPage) Build(context goflow.BuildContext) goflow.Widget {
	comparisonInfo := []goflow.Widget{
		createPageTitle("Transition Comparison"),
		createSubtitle("iOS vs Material Design"),
		widgets.NewContainer(), // Spacer

		createComparisonRow("iOS (Cupertino)", navigation.TransitionCupertino),
		createComparisonRow("Material Design", navigation.TransitionMaterial),

		widgets.NewContainer(), // Spacer
		createSubtitle("Custom Transitions"),

		createComparisonRow("Fade In/Out", navigation.TransitionFade),
		createComparisonRow("Zoom In/Out", navigation.TransitionZoom),

		widgets.NewContainer(), // Spacer
		createActionButton("Back", func() {
			navigation.Get.Back()
		}),
	}

	column := widgets.NewColumn(comparisonInfo...)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsAll(20.0)
	container.Child = column

	return widgets.NewCenter(container)
}

func (t *TransitionComparisonPage) CreateElement() goflow.Element {
	return goflow.NewGenericElement(t)
}

func createComparisonRow(label string, transition navigation.Transition) goflow.Widget {
	labelText := widgets.NewText(label)

	tryButton := createTransitionButton("Try It", transition)

	row := widgets.NewRow(labelText, tryButton)

	container := widgets.NewContainer()
	container.Padding = goflow.EdgeInsetsSymmetric(0, 8.0)
	container.Child = row

	return container
}
