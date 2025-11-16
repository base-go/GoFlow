package cupertino

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// NavigationBar is an iOS-style navigation bar
type NavigationBar struct {
	goflow.BaseWidget

	// Middle widget (usually Text with the title)
	Middle goflow.Widget

	// Leading widget (usually a back button or icon)
	Leading goflow.Widget

	// Trailing widget (usually action buttons)
	Trailing goflow.Widget

	// Optional styling
	BackgroundColor *goflow.Color
	BorderColor     *goflow.Color
	Padding         *goflow.EdgeInsets
	Height          float64
}

// NewNavigationBar creates a new Cupertino navigation bar
func NewNavigationBar(middle goflow.Widget) *NavigationBar {
	return &NavigationBar{
		Middle:  middle,
		Padding: goflow.NewEdgeInsets(16, 0, 16, 0),
		Height:  44.0, // Standard iOS navigation bar height
	}
}

// Build creates the widget tree for the navigation bar
func (n *NavigationBar) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := n.BackgroundColor
	if bgColor == nil {
		bgColor = theme.SystemBackgroundColor
	}

	// Build the navigation bar content
	var children []goflow.Widget

	// Add leading widget if present
	if n.Leading != nil {
		children = append(children, n.Leading)
	}

	// Add middle (title) widget - with flex to center it
	if n.Middle != nil {
		children = append(children, &widgets.Center{
			Child: n.Middle,
		})
	}

	// Add trailing widget if present
	if n.Trailing != nil {
		children = append(children, n.Trailing)
	}

	// Create a row with the children
	var content goflow.Widget
	if len(children) > 0 {
		content = &widgets.Row{
			Children:           children,
			MainAxisAlignment:  widgets.MainAxisSpaceBetween,
			CrossAxisAlignment: widgets.CrossAxisCenter,
		}
	}

	height := n.Height

	// Wrap in a container with the navigation bar styling
	return &widgets.Container{
		Color:   bgColor,
		Height:  &height,
		Padding: n.Padding,
		Child:   content,
	}
}
