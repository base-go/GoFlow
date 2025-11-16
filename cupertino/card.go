package cupertino

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// Card is an iOS-style card widget (similar to a grouped list item)
type Card struct {
	goflow.BaseWidget

	// Required
	Child goflow.Widget

	// Optional styling
	Color        *goflow.Color
	BorderColor  *goflow.Color
	BorderRadius float64
	Margin       *goflow.EdgeInsets
	Padding      *goflow.EdgeInsets
	Width        *float64
	Height       *float64
}

// NewCard creates a new Cupertino card
func NewCard(child goflow.Widget) *Card {
	return &Card{
		Child:        child,
		BorderRadius: 10.0, // iOS uses more rounded corners
		Margin:       goflow.NewEdgeInsets(16, 8, 16, 8),
		Padding:      goflow.NewEdgeInsets(16, 12, 16, 12),
	}
}

// Build creates the widget tree for the card
func (c *Card) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	cardColor := c.Color
	if cardColor == nil {
		cardColor = theme.SecondaryBackgroundColor
	}

	// Build the card with padding, color, and margins
	return &widgets.Container{
		Margin:  c.Margin,
		Padding: c.Padding,
		Color:   cardColor,
		Width:   c.Width,
		Height:  c.Height,
		Child:   c.Child,
	}
}
