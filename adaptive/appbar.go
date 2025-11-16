package adaptive

import (
	"github.com/base-go/GoFlow/cupertino"
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/material"
)

// AppBar is an adaptive app bar/navigation bar that automatically uses
// Material or Cupertino style based on the platform
type AppBar struct {
	goflow.BaseWidget

	// Title widget (usually Text)
	Title goflow.Widget

	// Leading widget (usually a back button or menu icon)
	Leading goflow.Widget

	// Trailing/Actions widgets (action buttons)
	Trailing []goflow.Widget

	// Optional styling
	BackgroundColor *goflow.Color
}

// NewAppBar creates a new adaptive app bar
func NewAppBar(title goflow.Widget) *AppBar {
	return &AppBar{
		Title: title,
	}
}

// Build creates the widget tree for the app bar
func (a *AppBar) Build(context goflow.BuildContext) goflow.Widget {
	platform := goflow.GetPlatform()

	// Switch based on platform
	if platform.IsApple() {
		// Use Cupertino NavigationBar for iOS/macOS
		navBar := cupertino.NewNavigationBar(a.Title)
		navBar.Leading = a.Leading
		navBar.BackgroundColor = a.BackgroundColor

		// Cupertino only supports one trailing widget
		if len(a.Trailing) > 0 {
			navBar.Trailing = a.Trailing[0]
		}

		return navBar
	} else {
		// Use Material AppBar for Android/Linux/Windows/Web
		appBar := material.NewAppBar(a.Title)
		appBar.Leading = a.Leading
		appBar.BackgroundColor = a.BackgroundColor
		appBar.Actions = a.Trailing

		return appBar
	}
}
