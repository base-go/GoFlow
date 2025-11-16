package widgets

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// TextFormField is a TextField with form validation
type TextFormField struct {
	goflow.BaseWidget
	Controller       *TextEditingController
	FocusNode        *FocusNode
	Decoration       *InputDecoration
	Style            *goflow.TextStyle
	Validator        func(value string) *string
	OnSaved          func(value string)
	OnChanged        func(value string)
	InitialValue     string
	Enabled          bool
	AutoValidate     bool
	ObscureText      bool
	MaxLines         int
	fieldState       *FormFieldStateImpl
}

// NewTextFormField creates a new TextFormField
func NewTextFormField() *TextFormField {
	return &TextFormField{
		Controller:   NewTextEditingController(""),
		FocusNode:    NewFocusNode(),
		Enabled:      true,
		AutoValidate: false,
		ObscureText:  false,
		MaxLines:     1,
	}
}

// WithValidator sets the validator
func (tff *TextFormField) WithValidator(validator func(string) *string) *TextFormField {
	tff.Validator = validator
	return tff
}

// WithInitialValue sets the initial value
func (tff *TextFormField) WithInitialValue(value string) *TextFormField {
	tff.InitialValue = value
	tff.Controller.SetText(value)
	return tff
}

// WithDecoration sets the decoration
func (tff *TextFormField) WithDecoration(decoration *InputDecoration) *TextFormField {
	tff.Decoration = decoration
	return tff
}

// WithObscureText sets whether to obscure text (for passwords)
func (tff *TextFormField) WithObscureText(obscure bool) *TextFormField {
	tff.ObscureText = obscure
	return tff
}

// Build creates the widget tree
func (tff *TextFormField) Build(context goflow.BuildContext) goflow.Widget {
	// Initialize field state if not already done
	if tff.fieldState == nil {
		tff.fieldState = NewFormFieldState(tff.InitialValue, tff.Validator, tff.OnSaved)
	}

	// Create TextField with error display
	textField := &TextField{
		Controller:  tff.Controller,
		FocusNode:   tff.FocusNode,
		Decoration:  tff.Decoration,
		Style:       tff.Style,
		ObscureText: tff.ObscureText,
		MaxLines:    tff.MaxLines,
		ReadOnly:    !tff.Enabled,
		OnChanged: func(text string) {
			tff.fieldState.SetValue(text)

			// Validate on change if auto-validate is enabled
			if tff.AutoValidate {
				tff.fieldState.Validate()
			}

			// Call user's onChange callback
			if tff.OnChanged != nil {
				tff.OnChanged(text)
			}
		},
	}

	// Add error text to decoration if validation failed
	if errorText := tff.fieldState.GetErrorText(); errorText != nil {
		if tff.Decoration == nil {
			tff.Decoration = &InputDecoration{}
		}
		// In a real implementation, we would display the error text below the field
		// For now, we'll just store it in the decoration
	}

	return textField
}

// DropdownFormField is a dropdown with form validation
type DropdownFormField struct {
	goflow.BaseWidget
	Value        interface{}
	Items        []DropdownMenuItem
	OnChanged    func(value interface{})
	OnSaved      func(value interface{})
	Validator    func(value interface{}) *string
	Decoration   *InputDecoration
	Hint         string
	Enabled      bool
	fieldState   *FormFieldStateImpl
}

// DropdownMenuItem represents an item in a dropdown
type DropdownMenuItem struct {
	Value interface{}
	Child goflow.Widget
}

// NewDropdownFormField creates a new dropdown form field
func NewDropdownFormField(items []DropdownMenuItem) *DropdownFormField {
	return &DropdownFormField{
		Items:   items,
		Enabled: true,
	}
}

// WithValidator sets the validator
func (dff *DropdownFormField) WithValidator(validator func(interface{}) *string) *DropdownFormField {
	dff.Validator = validator
	return dff
}

// WithValue sets the initial value
func (dff *DropdownFormField) WithValue(value interface{}) *DropdownFormField {
	dff.Value = value
	return dff
}

// Build creates the widget tree
func (dff *DropdownFormField) Build(context goflow.BuildContext) goflow.Widget {
	// Create dropdown with validation
	return &Dropdown{
		Value: dff.Value,
		Items: dff.Items,
		OnChanged: func(value interface{}) {
			dff.Value = value
			if dff.OnChanged != nil {
				dff.OnChanged(value)
			}
		},
		Hint:    dff.Hint,
		Enabled: dff.Enabled,
	}
}

// DatePickerFormField is a date picker with form validation
type DatePickerFormField struct {
	goflow.BaseWidget
	Value         *Date
	FirstDate     *Date
	LastDate      *Date
	OnSaved       func(date *Date)
	Validator     func(date *Date) *string
	Decoration    *InputDecoration
	Enabled       bool
	Format        string // Date format string
	fieldState    *FormFieldStateImpl
}

// Date represents a date
type Date struct {
	Year  int
	Month int
	Day   int
}

// NewDatePickerFormField creates a new date picker form field
func NewDatePickerFormField() *DatePickerFormField {
	now := &Date{Year: 2025, Month: 1, Day: 1} // Current date placeholder
	return &DatePickerFormField{
		Value:     now,
		FirstDate: &Date{Year: 1900, Month: 1, Day: 1},
		LastDate:  &Date{Year: 2100, Month: 12, Day: 31},
		Enabled:   true,
		Format:    "YYYY-MM-DD",
	}
}

// WithValidator sets the validator
func (dpff *DatePickerFormField) WithValidator(validator func(*Date) *string) *DatePickerFormField {
	dpff.Validator = validator
	return dpff
}

// Build creates the widget tree
func (dpff *DatePickerFormField) Build(context goflow.BuildContext) goflow.Widget {
	// Create a button that opens a date picker dialog
	dateText := "Select Date"
	if dpff.Value != nil {
		dateText = formatDate(dpff.Value, dpff.Format)
	}

	return &Container{
		Child: &Text{
			Data:  dateText,
			Style: goflow.NewTextStyle(),
		},
		Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
	}
}

// formatDate formats a date according to the format string
func formatDate(date *Date, format string) string {
	// Simple date formatting (can be enhanced)
	return "2025-01-01" // Placeholder
}

// TimePickerFormField is a time picker with form validation
type TimePickerFormField struct {
	goflow.BaseWidget
	Value      *TimeOfDay
	OnSaved    func(time *TimeOfDay)
	Validator  func(time *TimeOfDay) *string
	Decoration *InputDecoration
	Enabled    bool
	Format     string // Time format string
	fieldState *FormFieldStateImpl
}

// TimeOfDay represents a time of day
type TimeOfDay struct {
	Hour   int
	Minute int
}

// NewTimePickerFormField creates a new time picker form field
func NewTimePickerFormField() *TimePickerFormField {
	return &TimePickerFormField{
		Value:   &TimeOfDay{Hour: 12, Minute: 0},
		Enabled: true,
		Format:  "HH:MM",
	}
}

// WithValidator sets the validator
func (tpff *TimePickerFormField) WithValidator(validator func(*TimeOfDay) *string) *TimePickerFormField {
	tpff.Validator = validator
	return tff
}

// Build creates the widget tree
func (tpff *TimePickerFormField) Build(context goflow.BuildContext) goflow.Widget {
	// Create a button that opens a time picker dialog
	timeText := "Select Time"
	if tpff.Value != nil {
		timeText = formatTime(tpff.Value, tpff.Format)
	}

	return &Container{
		Child: &Text{
			Data:  timeText,
			Style: goflow.NewTextStyle(),
		},
		Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
	}
}

// formatTime formats a time according to the format string
func formatTime(time *TimeOfDay, format string) string {
	// Simple time formatting (can be enhanced)
	return "12:00" // Placeholder
}

// CheckboxFormField is a checkbox with form validation
type CheckboxFormField struct {
	goflow.BaseWidget
	Value      bool
	OnSaved    func(value bool)
	Validator  func(value bool) *string
	Title      goflow.Widget
	Enabled    bool
	fieldState *FormFieldStateImpl
}

// NewCheckboxFormField creates a new checkbox form field
func NewCheckboxFormField(title goflow.Widget) *CheckboxFormField {
	return &CheckboxFormField{
		Title:   title,
		Value:   false,
		Enabled: true,
	}
}

// WithValidator sets the validator
func (cff *CheckboxFormField) WithValidator(validator func(bool) *string) *CheckboxFormField {
	cff.Validator = validator
	return cff
}

// Build creates the widget tree
func (cff *CheckboxFormField) Build(context goflow.BuildContext) goflow.Widget {
	return &Row{
		Children: []goflow.Widget{
			&Checkbox{
				Value: cff.Value,
				OnChanged: func(value bool) {
					cff.Value = value
				},
				Enabled: cff.Enabled,
			},
			cff.Title,
		},
	}
}

// RadioFormField is a radio button group with form validation
type RadioFormField struct {
	goflow.BaseWidget
	Value      interface{}
	GroupValue interface{}
	OnSaved    func(value interface{})
	Validator  func(value interface{}) *string
	Title      goflow.Widget
	Enabled    bool
	fieldState *FormFieldStateImpl
}

// NewRadioFormField creates a new radio form field
func NewRadioFormField(value interface{}, groupValue interface{}, title goflow.Widget) *RadioFormField {
	return &RadioFormField{
		Value:      value,
		GroupValue: groupValue,
		Title:      title,
		Enabled:    true,
	}
}

// Build creates the widget tree
func (rff *RadioFormField) Build(context goflow.BuildContext) goflow.Widget {
	return &Row{
		Children: []goflow.Widget{
			&Radio{
				Value:      rff.Value,
				GroupValue: rff.GroupValue,
				OnChanged: func(value interface{}) {
					rff.GroupValue = value
				},
				Enabled: rff.Enabled,
			},
			rff.Title,
		},
	}
}

// SliderFormField is a slider with form validation
type SliderFormField struct {
	goflow.BaseWidget
	Value      float64
	Min        float64
	Max        float64
	Divisions  int
	OnSaved    func(value float64)
	Validator  func(value float64) *string
	Label      string
	Enabled    bool
	fieldState *FormFieldStateImpl
}

// NewSliderFormField creates a new slider form field
func NewSliderFormField(min, max float64) *SliderFormField {
	return &SliderFormField{
		Value:   min,
		Min:     min,
		Max:     max,
		Enabled: true,
	}
}

// WithValidator sets the validator
func (sff *SliderFormField) WithValidator(validator func(float64) *string) *SliderFormField {
	sff.Validator = validator
	return sff
}

// Build creates the widget tree
func (sff *SliderFormField) Build(context goflow.BuildContext) goflow.Widget {
	return &Slider{
		Value: sff.Value,
		Min:   sff.Min,
		Max:   sff.Max,
		OnChanged: func(value float64) {
			sff.Value = value
		},
		Enabled: sff.Enabled,
	}
}
