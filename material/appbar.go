package material

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// AppBar is a Material Design app bar (toolbar)
type AppBar struct {
	goflow.BaseWidget

	// Title widget (usually Text)
	Title goflow.Widget

	// Leading widget (usually an icon button)
	Leading goflow.Widget

	// Actions (usually icon buttons)
	Actions []goflow.Widget

	// Optional styling
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color
	Elevation       float64
	Height          float64
}

// NewAppBar creates a new Material app bar
func NewAppBar(title goflow.Widget) *AppBar {
	return &AppBar{
		Title:     title,
		Elevation: 4.0,
		Height:    56.0, // Standard Material app bar height
	}
}

// Build creates the widget tree for the app bar
func (a *AppBar) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := a.BackgroundColor
	if bgColor == nil {
		bgColor = theme.PrimaryColor
	}

	// Build the app bar content
	var children []goflow.Widget

	// Add leading widget if present
	if a.Leading != nil {
		children = append(children, a.Leading)
	}

	// Add title (with flex to take remaining space)
	if a.Title != nil {
		children = append(children, a.Title)
	}

	// Add actions if present
	if len(a.Actions) > 0 {
		for _, action := range a.Actions {
			children = append(children, action)
		}
	}

	// Create a row with the children
	var content goflow.Widget
	if len(children) > 0 {
		content = &widgets.Row{
			Children:           children,
			MainAxisAlignment:  widgets.MainAxisStart,
			CrossAxisAlignment: widgets.CrossAxisCenter,
		}
	}

	height := a.Height

	// Wrap in a container with the app bar styling
	return &widgets.Container{
		Color:   bgColor,
		Height:  &height,
		Padding: goflow.NewEdgeInsets(16, 0, 16, 0), // Horizontal padding
		Child:   content,
	}
}
