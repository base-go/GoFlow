package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Chip displays a compact element representing an attribute, text, entity, or action
type Chip struct {
	goflow.BaseWidget
	Label            string
	Avatar           goflow.Widget
	DeleteIcon       goflow.Widget
	OnDeleted        func()
	BackgroundColor  *goflow.Color
	LabelColor       *goflow.Color
	Padding          *goflow.EdgeInsets
	LabelPadding     *goflow.EdgeInsets
	DeleteIconColor  *goflow.Color
}

// NewChip creates a new Chip widget
func NewChip(label string) *Chip {
	return &Chip{
		Label:   label,
		Padding: goflow.NewEdgeInsets(4, 8, 4, 8),
	}
}

// Build creates the widget tree for Chip
func (c *Chip) Build(context goflow.BuildContext) goflow.Widget {
	children := make([]goflow.Widget, 0)

	// Add avatar if present
	if c.Avatar != nil {
		avatarContainer := &Container{
			Child:  c.Avatar,
			Margin: goflow.NewEdgeInsets(0, 0, 0, 8),
		}
		children = append(children, avatarContainer)
	}

	// Add label
	labelStyle := goflow.NewTextStyle()
	if c.LabelColor != nil {
		labelStyle.Color = c.LabelColor
	}
	label := NewTextWithStyle(c.Label, labelStyle)
	children = append(children, label)

	// Add delete icon if present
	if c.DeleteIcon != nil {
		deleteContainer := &Container{
			Child:  c.DeleteIcon,
			Margin: goflow.NewEdgeInsets(0, 8, 0, 0),
		}
		if c.OnDeleted != nil {
			// Wrap with gesture detector
			deleteContainer = &Container{
				Child: &GestureDetector{
					OnTap: c.OnDeleted,
					Child: deleteContainer,
				},
			}
		}
		children = append(children, deleteContainer)
	}

	// Create row with all elements
	row := NewRow(children)
	row.MainAxisSize = MainAxisSizeMin
	row.CrossAxisAlignment = CrossAxisCenter

	// Wrap in container with background and padding
	backgroundColor := c.BackgroundColor
	if backgroundColor == nil {
		backgroundColor = goflow.NewColor(224, 224, 224, 255) // Default gray
	}

	padding := c.Padding
	if padding == nil {
		padding = goflow.NewEdgeInsets(4, 12, 4, 12)
	}

	container := &Container{
		Child:   row,
		Color:   backgroundColor,
		Padding: padding,
	}

	return container
}

// ActionChip represents a chip that triggers an action
type ActionChip struct {
	goflow.BaseWidget
	Label           string
	Avatar          goflow.Widget
	OnPressed       func()
	BackgroundColor *goflow.Color
	LabelColor      *goflow.Color
	Padding         *goflow.EdgeInsets
}

// NewActionChip creates a new ActionChip widget
func NewActionChip(label string, onPressed func()) *ActionChip {
	return &ActionChip{
		Label:     label,
		OnPressed: onPressed,
		Padding:   goflow.NewEdgeInsets(4, 12, 4, 12),
	}
}

// Build creates the widget tree for ActionChip
func (a *ActionChip) Build(context goflow.BuildContext) goflow.Widget {
	children := make([]goflow.Widget, 0)

	// Add avatar if present
	if a.Avatar != nil {
		avatarContainer := &Container{
			Child:  a.Avatar,
			Margin: goflow.NewEdgeInsets(0, 0, 0, 8),
		}
		children = append(children, avatarContainer)
	}

	// Add label
	labelStyle := goflow.NewTextStyle()
	if a.LabelColor != nil {
		labelStyle.Color = a.LabelColor
	}
	label := NewTextWithStyle(a.Label, labelStyle)
	children = append(children, label)

	// Create row
	row := NewRow(children)
	row.MainAxisSize = MainAxisSizeMin
	row.CrossAxisAlignment = CrossAxisCenter

	// Wrap in gesture detector
	var child goflow.Widget = row
	if a.OnPressed != nil {
		child = &GestureDetector{
			OnTap: a.OnPressed,
			Child: row,
		}
	}

	// Wrap in container with background and padding
	backgroundColor := a.BackgroundColor
	if backgroundColor == nil {
		backgroundColor = goflow.NewColor(224, 224, 224, 255) // Default gray
	}

	padding := a.Padding
	if padding == nil {
		padding = goflow.NewEdgeInsets(4, 12, 4, 12)
	}

	container := &Container{
		Child:   child,
		Color:   backgroundColor,
		Padding: padding,
	}

	return container
}
