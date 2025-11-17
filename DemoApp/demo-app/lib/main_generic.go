//go:build !darwin
// +build !darwin

package main

import (
	gf "github.com/base-go/GoFlow"

	"com.example/demo-app/lib/screens"
)

func main() {
	// Register named routes for all pages
	routes := map[string]gf.RouteBuilder{
		"/":                  func() gf.Widget { return screens.NewHomePage() },
		"/input-widgets":     func() gf.Widget { return screens.NewInputWidgetsPage() },
		"/form-validation":   func() gf.Widget { return createPlaceholderPage("Form Validation", "Forms with validation - coming soon") },
		"/data-display":      func() gf.Widget { return createPlaceholderPage("Data Display", "DataTable, Cards, etc. - coming soon") },
		"/layout":            func() gf.Widget { return createPlaceholderPage("Layout Widgets", "Coming soon: Row, Column, Stack, Wrap, etc.") },
		"/navigation":        func() gf.Widget { return createPlaceholderPage("Navigation", "Coming soon: Dialogs, Sheets, Snackbars, etc.") },
		"/animations":        func() gf.Widget { return createPlaceholderPage("Animations", "Coming soon: AnimatedContainer, Hero, etc.") },
		"/scrolling":         func() gf.Widget { return createPlaceholderPage("Scrolling", "Coming soon: ListView, PageView, etc.") },
		"/display":           func() gf.Widget { return createPlaceholderPage("Display", "Coming soon: Text, Icon, Image, etc.") },
	}

	gf.Get.RegisterRoutes(routes)

	// Create the app with home page
	app := gf.GetMaterialApp(
		screens.NewHomePage(),
		routes,
		"GoFlow Kitchen Sink",
	)

	// Run the app
	gf.RunApp(app)
}

// createPlaceholderPage creates a simple placeholder page for unimplemented sections
func createPlaceholderPage(title, message string) gf.Widget {
	return &PlaceholderPage{
		Title:   title,
		Message: message,
	}
}

// PlaceholderPage is a simple page showing a title and message
type PlaceholderPage struct {
	gf.BaseWidget
	Title   string
	Message string
}

func (p *PlaceholderPage) Build(context gf.BuildContext) gf.Widget {
	// Create a simple placeholder layout
	titleWidget := gf.Text{Data: p.Title}
	messageWidget := gf.Text{Data: p.Message}

	backButton := gf.Text{Data: "← Back"}
	backContainer := gf.Container{
		Padding: gf.NewEdgeInsets(8, 16, 8, 16),
		Color:   gf.NewColor(200, 200, 200, 255),
		Child:   &backButton,
	}

	backBtn := gf.GestureDetector{
		Child: &backContainer,
		OnTap: func() {
			gf.Get.Back()
		},
	}

	spacer1 := 20.0
	spacer2 := 40.0

	content := gf.Column{
		Children: []gf.Widget{
			&backBtn,
			&gf.SizedBox{Height: &spacer1},
			&titleWidget,
			&gf.SizedBox{Height: &spacer2},
			&messageWidget,
		},
	}

	container := gf.Container{
		Padding: gf.NewEdgeInsets(20, 20, 20, 20),
		Child:   &content,
	}

	return &container
}

