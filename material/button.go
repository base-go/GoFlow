package material

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// ButtonStyle defines the visual style of a Material button
type ButtonStyle int

const (
	// ButtonStyleFilled is a filled button with elevation
	ButtonStyleFilled ButtonStyle = iota
	// ButtonStyleOutlined is an outlined button
	ButtonStyleOutlined
	// ButtonStyleText is a text-only button
	ButtonStyleText
)

// Button is a Material Design button widget
type Button struct {
	goflow.BaseWidget

	// Required
	OnPressed func()
	Child     goflow.Widget

	// Optional styling
	Style           ButtonStyle
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color
	DisabledColor   *goflow.Color
	Elevation       float64
	BorderRadius    float64
	Padding         *goflow.EdgeInsets
	MinWidth        float64
	MinHeight       float64
	Enabled         bool
}

// NewButton creates a new Material button
func NewButton(child goflow.Widget, onPressed func()) *Button {
	return &Button{
		Child:        child,
		OnPressed:    onPressed,
		Style:        ButtonStyleFilled,
		Elevation:    2.0,
		BorderRadius: 4.0,
		Padding:      goflow.NewEdgeInsets(16, 8, 16, 8), // left, top, right, bottom
		MinWidth:     88.0,
		MinHeight:    36.0,
		Enabled:      true,
	}
}

// NewTextButton creates a text-only Material button
func NewTextButton(text string, onPressed func()) *Button {
	btn := NewButton(&widgets.Text{
		Data: text,
	}, onPressed)
	btn.Style = ButtonStyleText
	btn.Elevation = 0.0
	return btn
}

// NewOutlinedButton creates an outlined Material button
func NewOutlinedButton(text string, onPressed func()) *Button {
	btn := NewButton(&widgets.Text{
		Data: text,
	}, onPressed)
	btn.Style = ButtonStyleOutlined
	btn.Elevation = 0.0
	return btn
}

// Build creates the widget tree for the button
func (b *Button) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	// Determine colors based on style and state
	var bgColor *goflow.Color
	var fgColor *goflow.Color

	if !b.Enabled {
		bgColor = b.DisabledColor
		if bgColor == nil {
			bgColor = theme.TextDisabledColor
		}
		fgColor = theme.TextDisabledColor
	} else {
		switch b.Style {
		case ButtonStyleFilled:
			bgColor = b.BackgroundColor
			if bgColor == nil {
				bgColor = theme.PrimaryColor
			}
			fgColor = b.ForegroundColor
			if fgColor == nil {
				fgColor = goflow.NewColor(255, 255, 255, 255) // White text
			}
		case ButtonStyleOutlined:
			bgColor = goflow.NewColor(0, 0, 0, 0) // Transparent
			fgColor = b.ForegroundColor
			if fgColor == nil {
				fgColor = theme.PrimaryColor
			}
		case ButtonStyleText:
			bgColor = goflow.NewColor(0, 0, 0, 0) // Transparent
			fgColor = b.ForegroundColor
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
	minWidth := b.MinWidth
	minHeight := b.MinHeight

	return &widgets.Container{
		Width:  &minWidth,
		Height: &minHeight,
		Child:  content,
	}
}
