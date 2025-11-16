package material

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// Checkbox is a Material Design checkbox
type Checkbox struct {
	goflow.BaseWidget

	// Current value
	Value bool

	// Callback when value changes
	OnChanged func(bool)

	// Colors
	ActiveColor   *goflow.Color
	CheckColor    *goflow.Color

	// Tristate support (can be null/indeterminate)
	Tristate bool
}

// NewCheckbox creates a new Material checkbox
func NewCheckbox(value bool, onChanged func(bool)) *Checkbox {
	return &Checkbox{
		Value:     value,
		OnChanged: onChanged,
		Tristate:  false,
	}
}

// Build creates the widget tree for the checkbox
func (c *Checkbox) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	activeColor := c.ActiveColor
	if activeColor == nil {
		activeColor = theme.PrimaryColor
	}

	checkColor := c.CheckColor
	if checkColor == nil {
		checkColor = goflow.NewColor(255, 255, 255, 255) // White
	}

	size := 18.0
	bgColor := goflow.NewColor(0, 0, 0, 0) // Transparent
	if c.Value {
		bgColor = activeColor
	}

	return &widgets.Container{
		Width:  &size,
		Height: &size,
		Color:  bgColor,
		Child: &widgets.Center{
			Child: &widgets.Text{
				Data: func() string {
					if c.Value {
						return "✓"
					}
					return ""
				}(),
			},
		},
	}
}

// Radio is a Material Design radio button
type Radio struct {
	goflow.BaseWidget

	// Current value of this radio
	Value interface{}

	// Group value (radio is selected if Value == GroupValue)
	GroupValue interface{}

	// Callback when selected
	OnChanged func(interface{})

	// Active color
	ActiveColor *goflow.Color
}

// NewRadio creates a new Material radio button
func NewRadio(value, groupValue interface{}, onChanged func(interface{})) *Radio {
	return &Radio{
		Value:      value,
		GroupValue: groupValue,
		OnChanged:  onChanged,
	}
}

// Build creates the widget tree for the radio button
func (r *Radio) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	activeColor := r.ActiveColor
	if activeColor == nil {
		activeColor = theme.PrimaryColor
	}

	size := 18.0
	isSelected := r.Value == r.GroupValue

	bgColor := goflow.NewColor(0, 0, 0, 0) // Transparent
	if isSelected {
		bgColor = activeColor
	}

	return &widgets.Container{
		Width:  &size,
		Height: &size,
		Color:  bgColor,
		Child: &widgets.Center{
			Child: &widgets.Text{
				Data: func() string {
					if isSelected {
						return "●"
					}
					return "○"
				}(),
			},
		},
	}
}

// Switch is a Material Design switch
type Switch struct {
	goflow.BaseWidget

	// Current value
	Value bool

	// Callback when value changes
	OnChanged func(bool)

	// Colors
	ActiveColor   *goflow.Color
	ActiveTrackColor *goflow.Color
	InactiveThumbColor *goflow.Color
	InactiveTrackColor *goflow.Color
}

// NewSwitch creates a new Material switch
func NewSwitch(value bool, onChanged func(bool)) *Switch {
	return &Switch{
		Value:     value,
		OnChanged: onChanged,
	}
}

// Build creates the widget tree for the switch
func (s *Switch) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	thumbColor := s.InactiveThumbColor
	trackColor := s.InactiveTrackColor

	if s.Value {
		thumbColor = s.ActiveColor
		trackColor = s.ActiveTrackColor
	}

	if thumbColor == nil {
		if s.Value {
			thumbColor = theme.PrimaryColor
		} else {
			thumbColor = goflow.NewColor(245, 245, 245, 255)
		}
	}

	if trackColor == nil {
		if s.Value {
			trackColor = theme.PrimaryLightColor
		} else {
			trackColor = goflow.NewColor(189, 189, 189, 255)
		}
	}

	trackWidth := 36.0
	trackHeight := 14.0
	thumbSize := 20.0

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

// Slider is a Material Design slider
type Slider struct {
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

	// Colors
	ActiveColor   *goflow.Color
	InactiveColor *goflow.Color

	// Label (shown when dragging)
	Label string
}

// NewSlider creates a new Material slider
func NewSlider(value, min, max float64, onChanged func(float64)) *Slider {
	return &Slider{
		Value:     value,
		Min:       min,
		Max:       max,
		OnChanged: onChanged,
	}
}

// Build creates the widget tree for the slider
func (s *Slider) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	activeColor := s.ActiveColor
	if activeColor == nil {
		activeColor = theme.PrimaryColor
	}

	inactiveColor := s.InactiveColor
	if inactiveColor == nil {
		inactiveColor = theme.DividerColor
	}

	// Calculate slider position
	percentage := (s.Value - s.Min) / (s.Max - s.Min)

	trackHeight := 4.0
	thumbSize := 20.0

	return &widgets.Container{
		Height: &trackHeight,
		Child: &widgets.Stack{
			Children: []goflow.Widget{
				// Inactive track
				&widgets.Container{
					Color: inactiveColor,
				},
				// Active track (positioned based on value)
				&widgets.Align{
					Alignment: widgets.AlignmentCenterLeft,
					Child: &widgets.Container{
						Color: activeColor,
					},
				},
				// Thumb (positioned based on value)
				&widgets.Align{
					Alignment: widgets.AlignmentValue{
						X: -1.0 + (percentage * 2.0),
						Y: 0.0,
					},
					Child: &widgets.Container{
						Width:  &thumbSize,
						Height: &thumbSize,
						Color:  activeColor,
					},
				},
			},
			Fit: widgets.StackFitPassthrough,
		},
	}
}
