package widgets

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Badge displays a small status descriptor for another widget
type Badge struct {
	goflow.BaseWidget
	Child           goflow.Widget
	Label           string
	Count           *int
	IsVisible       bool
	BackgroundColor *goflow.Color
	TextColor       *goflow.Color
	Offset          *goflow.Offset
	Alignment       BadgeAlignment
}

// BadgeAlignment defines where the badge is positioned
type BadgeAlignment int

const (
	BadgeTopEnd BadgeAlignment = iota
	BadgeTopStart
	BadgeBottomEnd
	BadgeBottomStart
)

// NewBadge creates a new Badge widget
func NewBadge(child goflow.Widget) *Badge {
	return &Badge{
		Child:     child,
		IsVisible: true,
		Alignment: BadgeTopEnd,
	}
}

// NewBadgeWithCount creates a badge with a count
func NewBadgeWithCount(child goflow.Widget, count int) *Badge {
	return &Badge{
		Child:     child,
		Count:     &count,
		IsVisible: true,
		Alignment: BadgeTopEnd,
	}
}

// Build creates the widget tree for Badge
func (b *Badge) Build(context goflow.BuildContext) goflow.Widget {
	if !b.IsVisible {
		return b.Child
	}

	// Determine badge content
	badgeText := b.Label
	if b.Count != nil {
		if *b.Count > 99 {
			badgeText = "99+"
		} else {
			badgeText = fmt.Sprintf("%d", *b.Count)
		}
	}

	// If no content, just return child
	if badgeText == "" {
		return b.Child
	}

	// Create badge widget
	backgroundColor := b.BackgroundColor
	if backgroundColor == nil {
		backgroundColor = goflow.NewColor(244, 67, 54, 255) // Default red
	}

	textColor := b.TextColor
	if textColor == nil {
		textColor = goflow.NewColor(255, 255, 255, 255) // White
	}

	textStyle := goflow.NewTextStyle()
	textStyle.Color = textColor
	textStyle.FontSize = 10

	badgeContent := &Container{
		Child:   NewTextWithStyle(badgeText, textStyle),
		Color:   backgroundColor,
		Padding: goflow.NewEdgeInsets(2, 6, 2, 6),
	}

	// Create positioned badge
	var positioned *Positioned
	switch b.Alignment {
	case BadgeTopEnd:
		right := 0.0
		top := 0.0
		positioned = &Positioned{
			Child:  badgeContent,
			Right:  &right,
			Top:    &top,
		}
	case BadgeTopStart:
		left := 0.0
		top := 0.0
		positioned = &Positioned{
			Child: badgeContent,
			Left:  &left,
			Top:   &top,
		}
	case BadgeBottomEnd:
		right := 0.0
		bottom := 0.0
		positioned = &Positioned{
			Child:  badgeContent,
			Right:  &right,
			Bottom: &bottom,
		}
	case BadgeBottomStart:
		left := 0.0
		bottom := 0.0
		positioned = &Positioned{
			Child:  badgeContent,
			Left:   &left,
			Bottom: &bottom,
		}
	}

	// Create stack with child and positioned badge
	stack := NewStack([]goflow.Widget{
		b.Child,
		positioned,
	})
	stack.Fit = StackFitLoose
	return stack
}
