package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// DropdownButton displays a button that opens a dropdown menu
type DropdownButton struct {
	goflow.BaseWidget
	Value         interface{}
	Items         []DropdownMenuItem
	OnChanged     func(interface{})
	Hint          string
	Elevation     float64
	Style         *DropdownButtonStyle
	IsExpanded    bool
	UnderlineColor *goflow.Color
}

// DropdownMenuItem represents an item in a dropdown
type DropdownMenuItem struct {
	Value interface{}
	Child goflow.Widget
}

// DropdownButtonStyle defines the style for dropdown button
type DropdownButtonStyle struct {
	BackgroundColor *goflow.Color
	TextColor       *goflow.Color
	IconColor       *goflow.Color
	BorderRadius    float64
	Padding         *goflow.EdgeInsets
}

// NewDropdownButton creates a new DropdownButton
func NewDropdownButton(value interface{}, items []DropdownMenuItem, onChanged func(interface{})) *DropdownButton {
	return &DropdownButton{
		Value:     value,
		Items:     items,
		OnChanged: onChanged,
		Elevation: 8.0,
	}
}

// Build creates the widget tree for DropdownButton
func (d *DropdownButton) Build(context goflow.BuildContext) goflow.Widget {
	// Find the selected item
	var selectedChild goflow.Widget
	if d.Value != nil {
		for _, item := range d.Items {
			if item.Value == d.Value {
				selectedChild = item.Child
				break
			}
		}
	}

	// If no selection and hint is provided, show hint
	if selectedChild == nil && d.Hint != "" {
		hintStyle := goflow.NewTextStyle()
		hintStyle.Color = goflow.NewColor(158, 158, 158, 255) // Gray
		selectedChild = NewTextWithStyle(d.Hint, hintStyle)
	}

	// Create icon
	iconColor := goflow.NewColor(0, 0, 0, 255)
	if d.Style != nil && d.Style.IconColor != nil {
		iconColor = d.Style.IconColor
	}

	dropdownIcon := &Icon{
		Icon:  IconArrowDropDown,
		Size:  24,
		Color: iconColor,
	}

	// Create button content
	buttonContent := NewRow([]goflow.Widget{
		&Flexible{
			Child: selectedChild,
			Flex:  1,
		},
		dropdownIcon,
	})
	buttonContent.MainAxisAlignment = MainAxisSpaceBetween
	buttonContent.CrossAxisAlignment = CrossAxisCenter

	// Apply padding
	padding := goflow.NewEdgeInsets(8, 12, 8, 12)
	if d.Style != nil && d.Style.Padding != nil {
		padding = d.Style.Padding
	}

	paddedContent := &Container{
		Child:   buttonContent,
		Padding: padding,
	}

	// Apply background color
	var backgroundColor *goflow.Color
	if d.Style != nil {
		backgroundColor = d.Style.BackgroundColor
	}

	container := &Container{
		Child: paddedContent,
		Color: backgroundColor,
	}

	// Add underline if specified
	if d.UnderlineColor != nil {
		children := []goflow.Widget{
			container,
			NewDivider(),
		}
		column := NewColumn(children)
		column.MainAxisSize = MainAxisSizeMin
		container = &Container{
			Child: column,
		}
	}

	// Wrap in gesture detector to handle taps
	return &GestureDetector{
		OnTap: func() {
			// In a real implementation, this would show the dropdown menu
			// For now, we'll just cycle through values
			if d.OnChanged != nil && len(d.Items) > 0 {
				currentIndex := -1
				for i, item := range d.Items {
					if item.Value == d.Value {
						currentIndex = i
						break
					}
				}
				nextIndex := (currentIndex + 1) % len(d.Items)
				d.OnChanged(d.Items[nextIndex].Value)
			}
		},
		Child: container,
	}
}

// PopupMenuButton displays a button that shows a popup menu
type PopupMenuButton struct {
	goflow.BaseWidget
	ItemBuilder func() []PopupMenuItem
	OnSelected  func(interface{})
	Icon        goflow.Widget
	Child       goflow.Widget
	Offset      *goflow.Offset
	Elevation   float64
}

// PopupMenuItem represents an item in a popup menu
type PopupMenuItem struct {
	Value   interface{}
	Child   goflow.Widget
	Enabled bool
}

// NewPopupMenuButton creates a new PopupMenuButton
func NewPopupMenuButton(itemBuilder func() []PopupMenuItem, onSelected func(interface{})) *PopupMenuButton {
	return &PopupMenuButton{
		ItemBuilder: itemBuilder,
		OnSelected:  onSelected,
		Elevation:   8.0,
	}
}

// Build creates the widget tree for PopupMenuButton
func (p *PopupMenuButton) Build(context goflow.BuildContext) goflow.Widget {
	// Determine button content
	var buttonContent goflow.Widget
	if p.Child != nil {
		buttonContent = p.Child
	} else if p.Icon != nil {
		buttonContent = p.Icon
	} else {
		// Default icon
		buttonContent = &Icon{
			Icon:  IconMoreVert,
			Size:  24,
			Color: goflow.NewColor(0, 0, 0, 255),
		}
	}

	// Wrap in gesture detector
	return &GestureDetector{
		OnTap: func() {
			// In a real implementation, this would show the popup menu
			// For now, we'll just trigger the first item
			if p.OnSelected != nil && p.ItemBuilder != nil {
				items := p.ItemBuilder()
				if len(items) > 0 {
					p.OnSelected(items[0].Value)
				}
			}
		},
		Child: buttonContent,
	}
}

// Autocomplete provides autocomplete functionality for text input
type Autocomplete struct {
	goflow.BaseWidget
	OptionsBuilder  func(string) []AutocompleteOption
	OnSelected      func(interface{})
	InitialValue    string
	FieldViewBuilder func(func(), goflow.Widget, func(string)) goflow.Widget
}

// AutocompleteOption represents an autocomplete option
type AutocompleteOption struct {
	Value interface{}
	Label string
}

// NewAutocomplete creates a new Autocomplete widget
func NewAutocomplete(optionsBuilder func(string) []AutocompleteOption, onSelected func(interface{})) *Autocomplete {
	return &Autocomplete{
		OptionsBuilder: optionsBuilder,
		OnSelected:     onSelected,
	}
}

// Build creates the widget tree for Autocomplete
func (a *Autocomplete) Build(context goflow.BuildContext) goflow.Widget {
	// Create text field
	textField := &TextField{
		OnChanged: func(value string) {
			// In a real implementation, this would show autocomplete options
			if a.OnSelected != nil && a.OptionsBuilder != nil {
				options := a.OptionsBuilder(value)
				if len(options) > 0 {
					a.OnSelected(options[0].Value)
				}
			}
		},
	}

	return textField
}
