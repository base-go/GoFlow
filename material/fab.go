package material

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// FloatingActionButton is a circular Material Design button
type FloatingActionButton struct {
	goflow.BaseWidget

	// Required
	OnPressed func()

	// Child widget (usually an icon)
	Child goflow.Widget

	// Optional styling
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color
	Elevation       float64
	Mini            bool // Smaller size
	Tooltip         string
}

// NewFloatingActionButton creates a new FAB
func NewFloatingActionButton(child goflow.Widget, onPressed func()) *FloatingActionButton {
	return &FloatingActionButton{
		Child:     child,
		OnPressed: onPressed,
		Elevation: 6.0,
		Mini:      false,
	}
}

// Build creates the widget tree for the FAB
func (f *FloatingActionButton) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := f.BackgroundColor
	if bgColor == nil {
		bgColor = theme.AccentColor
	}

	fgColor := f.ForegroundColor
	if fgColor == nil {
		fgColor = goflow.NewColor(255, 255, 255, 255) // White
	}

	size := 56.0
	if f.Mini {
		size = 40.0
	}

	return &widgets.Container{
		Width:  &size,
		Height: &size,
		Color:  bgColor,
		Child: &widgets.Center{
			Child: f.Child,
		},
	}
}
