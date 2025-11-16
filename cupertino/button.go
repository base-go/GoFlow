package cupertino

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// ButtonType defines the visual style of a Cupertino button
type ButtonType int

const (
	// ButtonTypeFilled is a filled button
	ButtonTypeFilled ButtonType = iota
	// ButtonTypeGray is a gray button
	ButtonTypeGray
	// ButtonTypePlain is a plain text button
	ButtonTypePlain
)

// Button is an iOS-style button widget
type Button struct {
	goflow.BaseWidget

	// Required
	OnPressed func()
	Child     goflow.Widget

	// Optional styling
	Type            ButtonType
	Color           *goflow.Color
	DisabledColor   *goflow.Color
	BorderRadius    float64
	Padding         *goflow.EdgeInsets
	MinSize         float64
	Enabled         bool
}

// NewButton creates a new Cupertino button
func NewButton(child goflow.Widget, onPressed func()) *Button {
	return &Button{
		Child:        child,
		OnPressed:    onPressed,
		Type:         ButtonTypeFilled,
		BorderRadius: 8.0, // iOS uses more rounded corners
		Padding:      goflow.NewEdgeInsets(16, 8, 16, 8),
		MinSize:      44.0, // iOS minimum touch target
		Enabled:      true,
	}
}

// NewTextButton creates a plain text Cupertino button
func NewTextButton(text string, onPressed func()) *Button {
	btn := NewButton(&widgets.Text{
		Data: text,
	}, onPressed)
	btn.Type = ButtonTypePlain
	return btn
}

// NewGrayButton creates a gray Cupertino button
func NewGrayButton(text string, onPressed func()) *Button {
	btn := NewButton(&widgets.Text{
		Data: text,
	}, onPressed)
	btn.Type = ButtonTypeGray
	return btn
}

// Build creates the widget tree for the button
func (b *Button) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	// Determine colors based on type and state
	var bgColor *goflow.Color
	var fgColor *goflow.Color

	if !b.Enabled {
		bgColor = b.DisabledColor
		if bgColor == nil {
			bgColor = theme.QuaternaryLabelColor
		}
		fgColor = theme.QuaternaryLabelColor
	} else {
		switch b.Type {
		case ButtonTypeFilled:
			bgColor = b.Color
			if bgColor == nil {
				bgColor = theme.PrimaryColor
			}
			fgColor = goflow.NewColor(255, 255, 255, 255) // White text
		case ButtonTypeGray:
			bgColor = theme.SystemBackgroundColor
			fgColor = theme.LabelColor
		case ButtonTypePlain:
			bgColor = goflow.NewColor(0, 0, 0, 0) // Transparent
			fgColor = b.Color
			if fgColor == nil {
				fgColor = theme.PrimaryColor
			}
		}
	}

	// Build the button content
	content := &widgets.Container{
		Color:   bgColor,
		Padding: b.Padding,
		Child:   b.Child,
	}

	// Wrap with minimum size constraints
	minSize := b.MinSize

	return &widgets.Container{
		Width:  &minSize,
		Height: &minSize,
		Child:  content,
	}
}
