package widgets

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// Checkbox widget
type Checkbox struct {
	goflow.BaseWidget
	Value           bool
	OnChanged       func(bool)
	ActiveColor     *goflow.Color
	CheckColor      *goflow.Color
	FocusColor      *goflow.Color
	HoverColor      *goflow.Color
	TriState        bool // Allows null state
	Enabled         bool
	Shape           CheckboxShape
	Side            *BorderSide
	VisualDensity   float64
}

// CheckboxShape defines the shape of a checkbox
type CheckboxShape int

const (
	CheckboxShapeRectangle CheckboxShape = iota
	CheckboxShapeCircle
	CheckboxShapeRounded
)

// NewCheckbox creates a new checkbox
func NewCheckbox(value bool, onChanged func(bool)) *Checkbox {
	return &Checkbox{
		Value:         value,
		OnChanged:     onChanged,
		ActiveColor:   goflow.NewColor(33, 150, 243, 255), // Blue
		CheckColor:    goflow.NewColor(255, 255, 255, 255), // White
		Enabled:       true,
		Shape:         CheckboxShapeRounded,
		VisualDensity: 0.0,
	}
}

// Build creates the widget tree
func (c *Checkbox) Build(context goflow.BuildContext) goflow.Widget {
	size := 18.0 + c.VisualDensity*4.0

	// Background color
	bgColor := goflow.NewColor(200, 200, 200, 255)
	if c.Value {
		bgColor = c.ActiveColor
	}
	if !c.Enabled {
		bgColor = goflow.NewColor(220, 220, 220, 255)
	}

	return &GestureDetector{
		OnTap: func() {
			if c.Enabled && c.OnChanged != nil {
				c.OnChanged(!c.Value)
			}
		},
		Child: &Container{
			Width:  &size,
			Height: &size,
			Color:  bgColor,
			Child: c.buildCheckMark(),
		},
	}
}

func (c *Checkbox) buildCheckMark() goflow.Widget {
	if !c.Value {
		return nil
	}

	// Return a checkmark icon
	return &Icon{
		Icon:  "check",
		Size:  14.0,
		Color: c.CheckColor,
	}
}

// Radio widget
type Radio struct {
	goflow.BaseWidget
	Value         interface{}
	GroupValue    interface{}
	OnChanged     func(interface{})
	ActiveColor   *goflow.Color
	FocusColor    *goflow.Color
	HoverColor    *goflow.Color
	Enabled       bool
	ToggleAble    bool
	VisualDensity float64
}

// NewRadio creates a new radio button
func NewRadio(value interface{}, groupValue interface{}, onChanged func(interface{})) *Radio {
	return &Radio{
		Value:         value,
		GroupValue:    groupValue,
		OnChanged:     onChanged,
		ActiveColor:   goflow.NewColor(33, 150, 243, 255), // Blue
		Enabled:       true,
		ToggleAble:    false,
		VisualDensity: 0.0,
	}
}

// Build creates the widget tree
func (r *Radio) Build(context goflow.BuildContext) goflow.Widget {
	size := 20.0 + r.VisualDensity*4.0
	innerSize := size * 0.5

	isSelected := r.Value == r.GroupValue

	// Outer circle color
	outerColor := goflow.NewColor(200, 200, 200, 255)
	if isSelected {
		outerColor = r.ActiveColor
	}
	if !r.Enabled {
		outerColor = goflow.NewColor(220, 220, 220, 255)
	}

	return &GestureDetector{
		OnTap: func() {
			if r.Enabled && r.OnChanged != nil {
				r.OnChanged(r.Value)
			}
		},
		Child: &Container{
			Width:  &size,
			Height: &size,
			Color:  outerColor,
			Child:  r.buildInnerCircle(innerSize, isSelected),
		},
	}
}

func (r *Radio) buildInnerCircle(size float64, isSelected bool) goflow.Widget {
	if !isSelected {
		return nil
	}

	return &Center{
		Child: &Container{
			Width:  &size,
			Height: &size,
			Color:  r.ActiveColor,
		},
	}
}

// Switch widget
type Switch struct {
	goflow.BaseWidget
	Value               bool
	OnChanged           func(bool)
	ActiveColor         *goflow.Color
	ActiveTrackColor    *goflow.Color
	InactiveThumbColor  *goflow.Color
	InactiveTrackColor  *goflow.Color
	FocusColor          *goflow.Color
	HoverColor          *goflow.Color
	Enabled             bool
	AutoFocus           bool
	DragStartBehavior   DragStartBehavior
}

// DragStartBehavior defines when drag starts
type DragStartBehavior int

const (
	DragStartDown DragStartBehavior = iota
	DragStartStart
)

// NewSwitch creates a new switch
func NewSwitch(value bool, onChanged func(bool)) *Switch {
	return &Switch{
		Value:               value,
		OnChanged:           onChanged,
		ActiveColor:         goflow.NewColor(33, 150, 243, 255), // Blue
		ActiveTrackColor:    goflow.NewColor(33, 150, 243, 100), // Light blue
		InactiveThumbColor:  goflow.NewColor(200, 200, 200, 255),
		InactiveTrackColor:  goflow.NewColor(150, 150, 150, 100),
		Enabled:             true,
		AutoFocus:           false,
		DragStartBehavior:   DragStartStart,
	}
}

// Build creates the widget tree
func (s *Switch) Build(context goflow.BuildContext) goflow.Widget {
	trackWidth := 36.0
	trackHeight := 20.0
	thumbSize := 16.0

	// Colors
	trackColor := s.InactiveTrackColor
	thumbColor := s.InactiveThumbColor
	if s.Value {
		trackColor = s.ActiveTrackColor
		thumbColor = s.ActiveColor
	}
	if !s.Enabled {
		trackColor = goflow.NewColor(150, 150, 150, 50)
		thumbColor = goflow.NewColor(200, 200, 200, 255)
	}

	// Thumb position
	thumbLeft := 2.0
	if s.Value {
		thumbLeft = trackWidth - thumbSize - 2.0
	}

	return &GestureDetector{
		OnTap: func() {
			if s.Enabled && s.OnChanged != nil {
				s.OnChanged(!s.Value)
			}
		},
		Child: &Container{
			Width:  &trackWidth,
			Height: &trackHeight,
			Color:  trackColor,
			Child: &Stack{
				Children: []goflow.Widget{
					&Positioned{
						Left: &thumbLeft,
						Top: func() *float64 {
							top := 2.0
							return &top
						}(),
						Child: &Container{
							Width:  &thumbSize,
							Height: &thumbSize,
							Color:  thumbColor,
						},
					},
				},
			},
		},
	}
}

// Slider widget
type Slider struct {
	goflow.BaseWidget
	Value         float64
	OnChanged     func(float64)
	OnChangeStart func(float64)
	OnChangeEnd   func(float64)
	Min           float64
	Max           float64
	Divisions     int
	Label         string
	ActiveColor   *goflow.Color
	InactiveColor *goflow.Color
	ThumbColor    *goflow.Color
	Enabled       bool
	AutoFocus     bool
}

// NewSlider creates a new slider
func NewSlider(value float64, min float64, max float64, onChanged func(float64)) *Slider {
	return &Slider{
		Value:         value,
		Min:           min,
		Max:           max,
		OnChanged:     onChanged,
		ActiveColor:   goflow.NewColor(33, 150, 243, 255), // Blue
		InactiveColor: goflow.NewColor(200, 200, 200, 255),
		ThumbColor:    goflow.NewColor(33, 150, 243, 255),
		Enabled:       true,
		AutoFocus:     false,
	}
}

// Build creates the widget tree
func (s *Slider) Build(context goflow.BuildContext) goflow.Widget {
	trackHeight := 4.0
	thumbSize := 20.0
	width := 200.0

	// Calculate thumb position based on value
	percentage := (s.Value - s.Min) / (s.Max - s.Min)
	thumbLeft := (width - thumbSize) * percentage

	activeWidth := width * percentage
	inactiveWidth := width * (1 - percentage)

	return &GestureDetector{
		OnPanUpdate: func(details *PanDetails) {
			if !s.Enabled {
				return
			}

			// Calculate new value based on pan position
			newPercentage := details.LocalPosition.X / width
			if newPercentage < 0 {
				newPercentage = 0
			}
			if newPercentage > 1 {
				newPercentage = 1
			}

			newValue := s.Min + (s.Max-s.Min)*newPercentage

			// Apply divisions if set
			if s.Divisions > 0 {
				step := (s.Max - s.Min) / float64(s.Divisions)
				newValue = float64(int(newValue/step+0.5)) * step
			}

			if s.OnChanged != nil {
				s.OnChanged(newValue)
			}
		},
		OnPanStart: func(details *PanDetails) {
			if s.OnChangeStart != nil {
				s.OnChangeStart(s.Value)
			}
		},
		OnPanEnd: func(details *PanDetails) {
			if s.OnChangeEnd != nil {
				s.OnChangeEnd(s.Value)
			}
		},
		Child: &Container{
			Width:  &width,
			Height: &thumbSize,
			Child: &Stack{
				Children: []goflow.Widget{
					// Inactive track
					&Positioned{
						Left: &activeWidth,
						Top: func() *float64 {
							top := (thumbSize - trackHeight) / 2
							return &top
						}(),
						Child: &Container{
							Width:  &inactiveWidth,
							Height: &trackHeight,
							Color:  s.InactiveColor,
						},
					},
					// Active track
					&Positioned{
						Left: func() *float64 {
							left := 0.0
							return &left
						}(),
						Top: func() *float64 {
							top := (thumbSize - trackHeight) / 2
							return &top
						}(),
						Child: &Container{
							Width:  &activeWidth,
							Height: &trackHeight,
							Color:  s.ActiveColor,
						},
					},
					// Thumb
					&Positioned{
						Left: &thumbLeft,
						Top: func() *float64 {
							top := 0.0
							return &top
						}(),
						Child: &Container{
							Width:  &thumbSize,
							Height: &thumbSize,
							Color:  s.ThumbColor,
						},
					},
				},
			},
		},
	}
}

// PanDetails contains details about a pan gesture
type PanDetails struct {
	GlobalPosition *goflow.Offset
	LocalPosition  *goflow.Offset
	Delta          *goflow.Offset
}

// RangeSlider widget for selecting a range
type RangeSlider struct {
	goflow.BaseWidget
	Values        RangeValues
	OnChanged     func(RangeValues)
	OnChangeStart func(RangeValues)
	OnChangeEnd   func(RangeValues)
	Min           float64
	Max           float64
	Divisions     int
	Labels        RangeLabels
	ActiveColor   *goflow.Color
	InactiveColor *goflow.Color
	Enabled       bool
}

// RangeValues represents a range of values
type RangeValues struct {
	Start float64
	End   float64
}

// RangeLabels represents labels for range values
type RangeLabels struct {
	Start string
	End   string
}

// NewRangeSlider creates a new range slider
func NewRangeSlider(start float64, end float64, min float64, max float64, onChanged func(RangeValues)) *RangeSlider {
	return &RangeSlider{
		Values: RangeValues{
			Start: start,
			End:   end,
		},
		Min:           min,
		Max:           max,
		OnChanged:     onChanged,
		ActiveColor:   goflow.NewColor(33, 150, 243, 255),
		InactiveColor: goflow.NewColor(200, 200, 200, 255),
		Enabled:       true,
	}
}

// Build creates the widget tree
func (rs *RangeSlider) Build(context goflow.BuildContext) goflow.Widget {
	// Similar to Slider but with two thumbs
	return &Container{
		Width: func() *float64 {
			w := 200.0
			return &w
		}(),
		Height: func() *float64 {
			h := 20.0
			return &h
		}(),
		Color: rs.InactiveColor,
	}
}
