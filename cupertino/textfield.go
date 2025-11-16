package cupertino

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/material"
	"github.com/base-go/GoFlow/widgets"
)

// CupertinoTextField is an iOS-style text input field
type CupertinoTextField struct {
	goflow.BaseWidget

	// Controller for the text value
	Controller *material.TextEditingController

	// Placeholder text
	Placeholder string

	// Prefix widget (usually an icon)
	Prefix goflow.Widget

	// Suffix widget
	Suffix goflow.Widget

	// Text style
	Style *goflow.TextStyle

	// Obscure text (for passwords)
	ObscureText bool

	// Enabled state
	Enabled bool

	// Callbacks
	OnChanged   func(string)
	OnSubmitted func(string)
}

// NewCupertinoTextField creates a new iOS-style text field
func NewCupertinoTextField() *CupertinoTextField {
	return &CupertinoTextField{
		Controller:  material.NewTextEditingController(""),
		ObscureText: false,
		Enabled:     true,
	}
}

// Build creates the widget tree for the text field
func (t *CupertinoTextField) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := theme.SystemBackgroundColor

	// Build the text display
	displayText := t.Controller.Text.Get()
	if displayText == "" && t.Placeholder != "" {
		displayText = t.Placeholder
	} else if t.ObscureText && displayText != "" {
		displayText = "•••••••"
	}

	var children []goflow.Widget

	// Add prefix if present
	if t.Prefix != nil {
		children = append(children, t.Prefix)
	}

	// Add text content
	textWidget := &widgets.Text{
		Data:  displayText,
		Style: t.Style,
	}

	if t.Style == nil {
		textWidget.Style = goflow.NewTextStyle()
		textWidget.Style.Color = theme.LabelColor
	}

	children = append(children, &widgets.Expanded{
		Child: textWidget,
	})

	// Add suffix if present
	if t.Suffix != nil {
		children = append(children, t.Suffix)
	}

	content := &widgets.Row{
		Children:           children,
		MainAxisAlignment:  widgets.MainAxisStart,
		CrossAxisAlignment: widgets.CrossAxisCenter,
	}

	minHeight := 44.0 // iOS standard
	return &widgets.Container{
		Color:   bgColor,
		Padding: goflow.NewEdgeInsets(12, 8, 12, 8),
		Height:  &minHeight,
		Child:   content,
	}
}

// CupertinoSwitch is an iOS-style switch
type CupertinoSwitch struct {
	goflow.BaseWidget

	// Current value
	Value bool

	// Callback when value changes
	OnChanged func(bool)

	// Active color
	ActiveColor *goflow.Color
}

// NewCupertinoSwitch creates a new iOS-style switch
func NewCupertinoSwitch(value bool, onChanged func(bool)) *CupertinoSwitch {
	return &CupertinoSwitch{
		Value:     value,
		OnChanged: onChanged,
	}
}

// Build creates the widget tree for the switch
func (s *CupertinoSwitch) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	thumbColor := goflow.NewColor(255, 255, 255, 255) // White
	trackColor := goflow.NewColor(230, 230, 230, 255)  // Light grey

	if s.Value {
		if s.ActiveColor != nil {
			trackColor = s.ActiveColor
		} else {
			trackColor = theme.PrimaryColor
		}
	}

	trackWidth := 51.0  // iOS standard
	trackHeight := 31.0 // iOS standard
	thumbSize := 27.0

	return &widgets.Container{
		Width:  &trackWidth,
		Height: &trackHeight,
		Color:  trackColor,
		Child: &widgets.Align{
			Alignment: func() widgets.AlignmentValue {
				if s.Value {
					return widgets.AlignmentCenterRight
				}
				return widgets.AlignmentCenterLeft
			}(),
			Child: &widgets.Container{
				Width:  &thumbSize,
				Height: &thumbSize,
				Color:  thumbColor,
			},
		},
	}
}

// CupertinoSlider is an iOS-style slider
type CupertinoSlider struct {
	goflow.BaseWidget

	// Current value
	Value float64

	// Min and max values
	Min float64
	Max float64

	// Divisions (for discrete slider)
	Divisions *int

	// Callback when value changes
	OnChanged func(float64)

	// Active color
	ActiveColor *goflow.Color
}

// NewCupertinoSlider creates a new iOS-style slider
func NewCupertinoSlider(value, min, max float64, onChanged func(float64)) *CupertinoSlider {
	return &CupertinoSlider{
		Value:     value,
		Min:       min,
		Max:       max,
		OnChanged: onChanged,
	}
}

// Build creates the widget tree for the slider
func (s *CupertinoSlider) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	activeColor := s.ActiveColor
	if activeColor == nil {
		activeColor = theme.PrimaryColor
	}

	inactiveColor := goflow.NewColor(230, 230, 230, 255)

	// Calculate slider position
	percentage := (s.Value - s.Min) / (s.Max - s.Min)

	trackHeight := 3.0
	thumbSize := 28.0

	return &widgets.Container{
		Height: &trackHeight,
		Child: &widgets.Stack{
			Children: []goflow.Widget{
				// Inactive track
				&widgets.Container{
					Color: inactiveColor,
				},
				// Active track
				&widgets.Align{
					Alignment: widgets.AlignmentCenterLeft,
					Child: &widgets.Container{
						Color: activeColor,
					},
				},
				// Thumb
				&widgets.Align{
					Alignment: widgets.AlignmentValue{
						X: -1.0 + (percentage * 2.0),
						Y: 0.0,
					},
					Child: &widgets.Container{
						Width:  &thumbSize,
						Height: &thumbSize,
						Color:  goflow.NewColor(255, 255, 255, 255),
					},
				},
			},
			Fit: widgets.StackFitPassthrough,
		},
	}
}
