package screens

import (
	gf "github.com/base-go/GoFlow"
)

// HomePage is the main landing page with navigation to all widget categories
type HomePage struct {
	gf.BaseWidget
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (h *HomePage) Build(context gf.BuildContext) gf.Widget {
	// Header
	header := gf.Text{Data: "GoFlow Kitchen Sink"}
	subtitle := gf.Text{Data: "Comprehensive Widget Showcase"}

	// Category cards
	categories := []CategoryCard{
		{
			Title:       "Input Widgets",
			Description: "Checkbox, Radio, Switch, Slider, RangeSlider",
			Icon:        "input",
			Color:       gf.NewColor(100, 150, 255, 255),
			Route:       "/input-widgets",
		},
		{
			Title:       "Form Validation",
			Description: "TextFormField, Dropdowns, Pickers, Validation",
			Icon:        "form",
			Color:       gf.NewColor(255, 150, 100, 255),
			Route:       "/form-validation",
		},
		{
			Title:       "Data Display",
			Description: "DataTable, Cards, ExpansionPanel, Lists",
			Icon:        "table",
			Color:       gf.NewColor(150, 200, 100, 255),
			Route:       "/data-display",
		},
		{
			Title:       "Layout Widgets",
			Description: "Row, Column, Stack, Wrap, Table, Flexible",
			Icon:        "layout",
			Color:       gf.NewColor(200, 100, 255, 255),
			Route:       "/layout",
		},
		{
			Title:       "Navigation Patterns",
			Description: "Dialogs, BottomSheets, Snackbars, Routes",
			Icon:        "navigation",
			Color:       gf.NewColor(100, 200, 200, 255),
			Route:       "/navigation",
		},
		{
			Title:       "Animations",
			Description: "AnimatedContainer, Hero, Transitions",
			Icon:        "animation",
			Color:       gf.NewColor(255, 200, 100, 255),
			Route:       "/animations",
		},
		{
			Title:       "Scrolling",
			Description: "ListView, PageView, ScrollController",
			Icon:        "scroll",
			Color:       gf.NewColor(150, 150, 255, 255),
			Route:       "/scrolling",
		},
		{
			Title:       "Display Widgets",
			Description: "Text, Icon, Image, Avatar, Badge, Chip",
			Icon:        "display",
			Color:       gf.NewColor(255, 150, 150, 255),
			Route:       "/display",
		},
	}

	// Build category grid
	var categoryWidgets []gf.Widget
	for _, cat := range categories {
		categoryWidgets = append(categoryWidgets, h.buildCategoryCard(cat))
	}

	// Create grid layout (2 columns)
	var rows []gf.Widget
	for i := 0; i < len(categoryWidgets); i += 2 {
		rowChildren := []gf.Widget{categoryWidgets[i]}
		if i+1 < len(categoryWidgets) {
			rowChildren = append(rowChildren, h.spacer(20), categoryWidgets[i+1])
		}
		row := gf.Row{Children: rowChildren}
		rows = append(rows, &row, h.spacer(20))
	}

	// Build main content
	allChildren := []gf.Widget{
		&header,
		h.spacer(10),
		&subtitle,
		h.spacer(40),
	}
	allChildren = append(allChildren, rows...)

	content := gf.Column{
		Children: allChildren,
	}

	container := gf.Container{
		Padding: gf.NewEdgeInsetsAll(20.0),
		Child:   &content,
	}

	return &container
}

func (h *HomePage) buildCategoryCard(cat CategoryCard) gf.Widget {
	title := gf.Text{Data: cat.Title}
	desc := gf.Text{Data: cat.Description}

	cardContent := gf.Column{
		Children: []gf.Widget{
			&title,
			h.spacer(10),
			&desc,
		},
	}

	width := 300.0
	height := 150.0

	card := gf.Container{
		Padding: gf.NewEdgeInsetsAll(20.0),
		Color:   cat.Color,
		Width:   &width,
		Height:  &height,
		Child:   &cardContent,
	}

	// Wrap in GestureDetector for navigation
	detector := gf.GestureDetector{
		Child: &card,
		OnTap: func() {
			gf.Get.ToNamed(cat.Route, gf.TransitionSlideRight)
		},
	}

	return &detector
}

func (h *HomePage) spacer(height float64) gf.Widget {
	container := gf.Container{
		Height: &height,
	}
	return &container
}


// CategoryCard represents a widget category
type CategoryCard struct {
	Title       string
	Description string
	Icon        string
	Color       *gf.Color
	Route       string
}
