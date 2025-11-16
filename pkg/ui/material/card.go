package material

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Card is a Material Design card widget
// A card is a sheet of material that serves as an entry point to more detailed information
type Card struct {
	goflow.BaseWidget

	// Required
	Child goflow.Widget

	// Optional styling
	Elevation      float64
	Color          *goflow.Color
	ShadowColor    *goflow.Color
	BorderRadius   float64
	Margin         *goflow.EdgeInsets
	Padding        *goflow.EdgeInsets
	Width          *float64
	Height         *float64
}

// NewCard creates a new Material card
func NewCard(child goflow.Widget) *Card {
	return &Card{
		Child:        child,
		Elevation:    1.0,
		BorderRadius: 4.0,
		Margin:       goflow.NewEdgeInsets(8, 8, 8, 8),
		Padding:      goflow.NewEdgeInsets(16, 16, 16, 16),
	}
}

// Build creates the widget tree for the card
func (c *Card) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	cardColor := c.Color
	if cardColor == nil {
		cardColor = theme.SurfaceColor
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
