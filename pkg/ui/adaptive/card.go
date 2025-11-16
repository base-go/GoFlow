package adaptive

import (
	"github.com/base-go/GoFlow/pkg/ui/cupertino"
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/ui/material"
)

// Card is an adaptive card that automatically uses Material or Cupertino style
// based on the platform
type Card struct {
	goflow.BaseWidget

	// Required
	Child goflow.Widget

	// Optional styling (will be adapted for each platform)
	Color   *goflow.Color
	Margin  *goflow.EdgeInsets
	Padding *goflow.EdgeInsets
	Width   *float64
	Height  *float64
}

// NewCard creates a new adaptive card
func NewCard(child goflow.Widget) *Card {
	return &Card{
		Child: child,
	}
}

// Build creates the widget tree for the card
func (c *Card) Build(context goflow.BuildContext) goflow.Widget {
	platform := goflow.GetPlatform()

	// Switch based on platform
	if platform.IsApple() {
		// Use Cupertino style for iOS/macOS
		card := cupertino.NewCard(c.Child)
		card.Color = c.Color
		card.Margin = c.Margin
		card.Padding = c.Padding
		card.Width = c.Width
		card.Height = c.Height
		return card
	} else {
		// Use Material style for Android/Linux/Windows/Web
		card := material.NewCard(c.Child)
		card.Color = c.Color
		card.Margin = c.Margin
		card.Padding = c.Padding
		card.Width = c.Width
		card.Height = c.Height
		return card
	}
}
