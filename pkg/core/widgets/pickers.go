package widgets

import (
	"fmt"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// DatePicker dialog for selecting dates
type DatePicker struct {
	goflow.BaseWidget
	InitialDate      *Date
	FirstDate        *Date
	LastDate         *Date
	CurrentDate      *Date
	SelectableDayPredicate func(*Date) bool
	HelpText         string
	CancelText       string
	ConfirmText      string
	Locale           string
	UseRootNavigator bool
}

// ShowDatePicker shows a date picker dialog
func ShowDatePicker(
	context goflow.BuildContext,
	initialDate *Date,
	firstDate *Date,
	lastDate *Date,
	onDateSelected func(*Date),
) {
	// Create date picker dialog
	picker := &DatePicker{
		InitialDate: initialDate,
		FirstDate:   firstDate,
		LastDate:    lastDate,
		CurrentDate: initialDate,
		HelpText:    "Select Date",
		CancelText:  "Cancel",
		ConfirmText: "OK",
	}

	// Show as dialog (in a real implementation, this would use the Navigator)
	// For now, we'll just call the callback
	if onDateSelected != nil {
		onDateSelected(initialDate)
	}
}

// Build creates the date picker widget tree
func (dp *DatePicker) Build(context goflow.BuildContext) goflow.Widget {
	return &Container{
		Width: func() *float64 {
			w := 320.0
			return &w
		}(),
		Height: func() *float64 {
			h := 400.0
			return &h
		}(),
		Color: goflow.NewColor(255, 255, 255, 255),
		Child: &Column{
			Children: []goflow.Widget{
				// Header
				&Container{
					Color: goflow.NewColor(33, 150, 243, 255),
					Child: &Text{
						Data:  dp.HelpText,
						Style: goflow.NewTextStyle().WithColor(goflow.NewColor(255, 255, 255, 255)),
					},
					Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
				},
				// Calendar view
				dp.buildCalendar(),
				// Actions
				&Row{
					MainAxisAlignment: MainAxisAlignmentEnd,
					Children: []goflow.Widget{
						&Text{
							Data:  dp.CancelText,
							Style: goflow.NewTextStyle(),
						},
						&SizedBox{Width: 8},
						&Text{
							Data:  dp.ConfirmText,
							Style: goflow.NewTextStyle(),
						},
					},
				},
			},
		},
	}
}

func (dp *DatePicker) buildCalendar() goflow.Widget {
	// Build a simple calendar grid
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	dayWidgets := make([]goflow.Widget, len(days))
	for i, day := range days {
		dayWidgets[i] = &Text{
			Data:  day,
			Style: goflow.NewTextStyle(),
		}
	}

	return &Container{
		Child: &Column{
			Children: []goflow.Widget{
				// Day headers
				&Row{
					Children: dayWidgets,
				},
				// Calendar dates (simplified)
				&Text{
					Data:  "Calendar Grid",
					Style: goflow.NewTextStyle(),
				},
			},
		},
		Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
	}
}

// TimePicker dialog for selecting time
type TimePicker struct {
	goflow.BaseWidget
	InitialTime      *TimeOfDay
	HelpText         string
	CancelText       string
	ConfirmText      string
	HourLabelText    string
	MinuteLabelText  string
	Use24HourFormat  bool
	Orientation      Orientation
}

// Orientation defines the orientation of a widget
type Orientation int

const (
	OrientationPortrait Orientation = iota
	OrientationLandscape
)

// ShowTimePicker shows a time picker dialog
func ShowTimePicker(
	context goflow.BuildContext,
	initialTime *TimeOfDay,
	onTimeSelected func(*TimeOfDay),
) {
	// Create time picker dialog
	picker := &TimePicker{
		InitialTime:     initialTime,
		HelpText:        "Select Time",
		CancelText:      "Cancel",
		ConfirmText:     "OK",
		HourLabelText:   "Hour",
		MinuteLabelText: "Minute",
		Use24HourFormat: true,
	}

	// Show as dialog (in a real implementation, this would use the Navigator)
	// For now, we'll just call the callback
	if onTimeSelected != nil {
		onTimeSelected(initialTime)
	}
}

// Build creates the time picker widget tree
func (tp *TimePicker) Build(context goflow.BuildContext) goflow.Widget {
	hourText := fmt.Sprintf("%02d", tp.InitialTime.Hour)
	minuteText := fmt.Sprintf("%02d", tp.InitialTime.Minute)

	return &Container{
		Width: func() *float64 {
			w := 320.0
			return &w
		}(),
		Height: func() *float64 {
			h := 400.0
			return &h
		}(),
		Color: goflow.NewColor(255, 255, 255, 255),
		Child: &Column{
			Children: []goflow.Widget{
				// Header
				&Container{
					Color: goflow.NewColor(33, 150, 243, 255),
					Child: &Text{
						Data:  tp.HelpText,
						Style: goflow.NewTextStyle().WithColor(goflow.NewColor(255, 255, 255, 255)),
					},
					Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
				},
				// Time display
				&Row{
					MainAxisAlignment: MainAxisAlignmentCenter,
					Children: []goflow.Widget{
						&Text{
							Data:  hourText,
							Style: goflow.NewTextStyle().WithFontSize(48),
						},
						&Text{
							Data:  ":",
							Style: goflow.NewTextStyle().WithFontSize(48),
						},
						&Text{
							Data:  minuteText,
							Style: goflow.NewTextStyle().WithFontSize(48),
						},
					},
				},
				// Actions
				&Row{
					MainAxisAlignment: MainAxisAlignmentEnd,
					Children: []goflow.Widget{
						&Text{
							Data:  tp.CancelText,
							Style: goflow.NewTextStyle(),
						},
						&SizedBox{Width: 8},
						&Text{
							Data:  tp.ConfirmText,
							Style: goflow.NewTextStyle(),
						},
					},
				},
			},
		},
	}
}

// ColorPicker widget for selecting colors
type ColorPicker struct {
	goflow.BaseWidget
	PickerColor       *goflow.Color
	OnColorChanged    func(*goflow.Color)
	PaletteType       PaletteType
	EnableAlpha       bool
	EnableLabel       bool
	ColorPickerWidth  float64
	ColorPickerHeight float64
}

// PaletteType defines the type of color palette
type PaletteType int

const (
	PaletteTypePrimary PaletteType = iota
	PaletteTypeAccent
	PaletteTypeBoth
)

// NewColorPicker creates a new color picker
func NewColorPicker(initialColor *goflow.Color, onColorChanged func(*goflow.Color)) *ColorPicker {
	return &ColorPicker{
		PickerColor:       initialColor,
		OnColorChanged:    onColorChanged,
		PaletteType:       PaletteTypePrimary,
		EnableAlpha:       true,
		EnableLabel:       true,
		ColorPickerWidth:  300,
		ColorPickerHeight: 300,
	}
}

// ShowColorPicker shows a color picker dialog
func ShowColorPicker(
	context goflow.BuildContext,
	initialColor *goflow.Color,
	onColorSelected func(*goflow.Color),
) {
	picker := NewColorPicker(initialColor, onColorSelected)
	// Show as dialog (in a real implementation, this would use the Navigator)
	if onColorSelected != nil {
		onColorSelected(initialColor)
	}
}

// Build creates the color picker widget tree
func (cp *ColorPicker) Build(context goflow.BuildContext) goflow.Widget {
	return &Container{
		Width:  &cp.ColorPickerWidth,
		Height: &cp.ColorPickerHeight,
		Color:  goflow.NewColor(255, 255, 255, 255),
		Child: &Column{
			Children: []goflow.Widget{
				// Color preview
				&Container{
					Width: func() *float64 {
						w := 100.0
						return &w
					}(),
					Height: func() *float64 {
						h := 100.0
						return &h
					}(),
					Color: cp.PickerColor,
				},
				// Color palette grid
				cp.buildColorPalette(),
				// Alpha slider (if enabled)
				cp.buildAlphaSlider(),
			},
		},
	}
}

func (cp *ColorPicker) buildColorPalette() goflow.Widget {
	// Build a grid of color swatches
	colors := []goflow.Color{
		*goflow.NewColor(244, 67, 54, 255),   // Red
		*goflow.NewColor(233, 30, 99, 255),   // Pink
		*goflow.NewColor(156, 39, 176, 255),  // Purple
		*goflow.NewColor(103, 58, 183, 255),  // Deep Purple
		*goflow.NewColor(63, 81, 181, 255),   // Indigo
		*goflow.NewColor(33, 150, 243, 255),  // Blue
		*goflow.NewColor(3, 169, 244, 255),   // Light Blue
		*goflow.NewColor(0, 188, 212, 255),   // Cyan
		*goflow.NewColor(0, 150, 136, 255),   // Teal
		*goflow.NewColor(76, 175, 80, 255),   // Green
		*goflow.NewColor(139, 195, 74, 255),  // Light Green
		*goflow.NewColor(205, 220, 57, 255),  // Lime
		*goflow.NewColor(255, 235, 59, 255),  // Yellow
		*goflow.NewColor(255, 193, 7, 255),   // Amber
		*goflow.NewColor(255, 152, 0, 255),   // Orange
		*goflow.NewColor(255, 87, 34, 255),   // Deep Orange
		*goflow.NewColor(121, 85, 72, 255),   // Brown
		*goflow.NewColor(158, 158, 158, 255), // Grey
		*goflow.NewColor(96, 125, 139, 255),  // Blue Grey
	}

	colorWidgets := make([]goflow.Widget, len(colors))
	for i, color := range colors {
		colorCopy := color
		colorWidgets[i] = &GestureDetector{
			OnTap: func() {
				if cp.OnColorChanged != nil {
					cp.OnColorChanged(&colorCopy)
				}
			},
			Child: &Container{
				Width: func() *float64 {
					w := 30.0
					return &w
				}(),
				Height: func() *float64 {
					h := 30.0
					return &h
				}(),
				Color: &colorCopy,
			},
		}
	}

	return &Wrap{
		Spacing:    4,
		RunSpacing: 4,
		Children:   colorWidgets,
	}
}

func (cp *ColorPicker) buildAlphaSlider() goflow.Widget {
	if !cp.EnableAlpha {
		return &SizedBox{Width: 0, Height: 0}
	}

	return &Slider{
		Value: float64(cp.PickerColor.A) / 255.0,
		Min:   0.0,
		Max:   1.0,
		OnChanged: func(value float64) {
			if cp.OnColorChanged != nil {
				newColor := &goflow.Color{
					R: cp.PickerColor.R,
					G: cp.PickerColor.G,
					B: cp.PickerColor.B,
					A: uint8(value * 255),
				}
				cp.OnColorChanged(newColor)
			}
		},
	}
}

// FilePicker allows selecting files from the file system
type FilePicker struct {
	goflow.BaseWidget
	AllowedExtensions []string
	AllowMultiple     bool
	InitialDirectory  string
	DialogTitle       string
}

// NewFilePicker creates a new file picker
func NewFilePicker() *FilePicker {
	return &FilePicker{
		AllowedExtensions: []string{},
		AllowMultiple:     false,
		DialogTitle:       "Select File",
	}
}

// PickFiles shows a file picker dialog
func (fp *FilePicker) PickFiles(onFilesSelected func([]string)) {
	// In a real implementation, this would open a native file picker dialog
	// For now, we'll just simulate it
	if onFilesSelected != nil {
		// Simulate file selection
		onFilesSelected([]string{"/path/to/file.txt"})
	}
}

// PickSingleFile shows a file picker dialog for a single file
func (fp *FilePicker) PickSingleFile(onFileSelected func(string)) {
	fp.PickFiles(func(files []string) {
		if len(files) > 0 && onFileSelected != nil {
			onFileSelected(files[0])
		}
	})
}

// Build creates the file picker widget tree
func (fp *FilePicker) Build(context goflow.BuildContext) goflow.Widget {
	return &Container{
		Child: &Text{
			Data:  "File Picker",
			Style: goflow.NewTextStyle(),
		},
		Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
	}
}

// DirectoryPicker allows selecting directories from the file system
type DirectoryPicker struct {
	goflow.BaseWidget
	InitialDirectory string
	DialogTitle      string
}

// NewDirectoryPicker creates a new directory picker
func NewDirectoryPicker() *DirectoryPicker {
	return &DirectoryPicker{
		DialogTitle: "Select Directory",
	}
}

// PickDirectory shows a directory picker dialog
func (dp *DirectoryPicker) PickDirectory(onDirectorySelected func(string)) {
	// In a real implementation, this would open a native directory picker dialog
	// For now, we'll just simulate it
	if onDirectorySelected != nil {
		// Simulate directory selection
		onDirectorySelected("/path/to/directory")
	}
}

// Build creates the directory picker widget tree
func (dp *DirectoryPicker) Build(context goflow.BuildContext) goflow.Widget {
	return &Container{
		Child: &Text{
			Data:  "Directory Picker",
			Style: goflow.NewTextStyle(),
		},
		Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
	}
}
