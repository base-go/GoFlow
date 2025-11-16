package material

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// BottomNavigationBar is a Material Design bottom navigation bar
type BottomNavigationBar struct {
	goflow.BaseWidget

	// Items
	Items []BottomNavigationBarItem

	// Current index
	CurrentIndex int

	// On tap callback
	OnTap func(int)

	// Background color
	BackgroundColor *goflow.Color

	// Selected item color
	SelectedItemColor *goflow.Color

	// Unselected item color
	UnselectedItemColor *goflow.Color

	// Show labels
	ShowSelectedLabels   bool
	ShowUnselectedLabels bool

	// Type
	Type BottomNavigationBarType

	// Elevation
	Elevation float64
}

// BottomNavigationBarType defines the style
type BottomNavigationBarType int

const (
	BottomNavigationBarTypeFixed BottomNavigationBarType = iota
	BottomNavigationBarTypeShifting
)

// BottomNavigationBarItem represents an item in the bottom nav bar
type BottomNavigationBarItem struct {
	Icon             widgets.IconData
	ActiveIcon       *widgets.IconData
	Label            string
	BackgroundColor  *goflow.Color
	Tooltip          string
}

// NewBottomNavigationBar creates a new bottom navigation bar
func NewBottomNavigationBar(items []BottomNavigationBarItem) *BottomNavigationBar {
	return &BottomNavigationBar{
		Items:                items,
		CurrentIndex:         0,
		Type:                 BottomNavigationBarTypeFixed,
		Elevation:            8.0,
		ShowSelectedLabels:   true,
		ShowUnselectedLabels: true,
	}
}

// Build creates the widget tree for the bottom navigation bar
func (b *BottomNavigationBar) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := b.BackgroundColor
	if bgColor == nil {
		bgColor = theme.SurfaceColor
	}

	selectedColor := b.SelectedItemColor
	if selectedColor == nil {
		selectedColor = theme.PrimaryColor
	}

	unselectedColor := b.UnselectedItemColor
	if unselectedColor == nil {
		unselectedColor = theme.TextSecondaryColor
	}

	// Build items
	itemWidgets := make([]goflow.Widget, len(b.Items))
	for i, item := range b.Items {
		isSelected := i == b.CurrentIndex

		iconColor := unselectedColor
		if isSelected {
			iconColor = selectedColor
		}

		icon := item.Icon
		if isSelected && item.ActiveIcon != nil {
			icon = *item.ActiveIcon
		}

		var children []goflow.Widget
		children = append(children, &widgets.Icon{
			Icon:  icon,
			Size:  24.0,
			Color: iconColor,
		})

		// Add label if needed
		showLabel := (isSelected && b.ShowSelectedLabels) || (!isSelected && b.ShowUnselectedLabels)
		if showLabel && item.Label != "" {
			children = append(children, &widgets.Text{
				Data:  item.Label,
				Style: &goflow.TextStyle{Color: iconColor, FontSize: 12.0},
			})
		}

		itemWidgets[i] = &widgets.Expanded{
			Child: &widgets.Container{
				Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
				Child: &widgets.Column{
					Children:       children,
					MainAxisAlign:  widgets.MainAxisCenter,
					CrossAxisAlign: widgets.CrossAxisCenter,
				},
			},
		}
	}

	height := 56.0

	return &widgets.Container{
		Height: &height,
		Color:  bgColor,
		Child: &widgets.Row{
			Children:           itemWidgets,
			MainAxisAlignment:  widgets.MainAxisSpaceAround,
			CrossAxisAlignment: widgets.CrossAxisCenter,
		},
	}
}

// Dialog is a Material Design dialog
type Dialog struct {
	goflow.BaseWidget

	// Child widget (dialog content)
	Child goflow.Widget

	// Background color
	BackgroundColor *goflow.Color

	// Elevation
	Elevation float64

	// Shape (border radius)
	BorderRadius float64

	// Inset padding
	InsetPadding *goflow.EdgeInsets
}

// NewDialog creates a new Material dialog
func NewDialog(child goflow.Widget) *Dialog {
	return &Dialog{
		Child:        child,
		Elevation:    24.0,
		BorderRadius: 4.0,
		InsetPadding: goflow.NewEdgeInsets(40, 24, 40, 24),
	}
}

// Build creates the widget tree for the dialog
func (d *Dialog) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := d.BackgroundColor
	if bgColor == nil {
		bgColor = theme.SurfaceColor
	}

	return &widgets.Container{
		Color:   bgColor,
		Padding: goflow.NewEdgeInsets(24, 20, 24, 24),
		Child:   d.Child,
	}
}

// AlertDialog is a Material Design alert dialog
type AlertDialog struct {
	goflow.BaseWidget

	// Title widget
	Title goflow.Widget

	// Content widget
	Content goflow.Widget

	// Actions (usually buttons)
	Actions []goflow.Widget

	// Background color
	BackgroundColor *goflow.Color
}

// NewAlertDialog creates a new alert dialog
func NewAlertDialog() *AlertDialog {
	return &AlertDialog{}
}

// Build creates the widget tree for the alert dialog
func (a *AlertDialog) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := a.BackgroundColor
	if bgColor == nil {
		bgColor = theme.SurfaceColor
	}

	var children []goflow.Widget

	// Add title if present
	if a.Title != nil {
		children = append(children, &widgets.Container{
			Padding: goflow.NewEdgeInsets(24, 24, 24, 0),
			Child:   a.Title,
		})
	}

	// Add content if present
	if a.Content != nil {
		children = append(children, &widgets.Container{
			Padding: goflow.NewEdgeInsets(24, 20, 24, 24),
			Child:   a.Content,
		})
	}

	// Add actions if present
	if len(a.Actions) > 0 {
		children = append(children, &widgets.Container{
			Padding: goflow.NewEdgeInsets(0, 0, 8, 8),
			Child: &widgets.Row{
				Children:           a.Actions,
				MainAxisAlignment:  widgets.MainAxisEnd,
				CrossAxisAlignment: widgets.CrossAxisCenter,
			},
		})
	}

	return &widgets.Container{
		Color: bgColor,
		Child: &widgets.Column{
			Children:       children,
			MainAxisAlign:  widgets.MainAxisStart,
			CrossAxisAlign: widgets.CrossAxisStretch,
			MainAxisSize:   widgets.MainAxisSizeMin,
		},
	}
}
