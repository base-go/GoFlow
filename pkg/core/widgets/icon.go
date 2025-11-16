package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Icon displays an icon
type Icon struct {
	goflow.BaseWidget

	// Icon data (name or code point)
	Icon IconData

	// Size
	Size float64

	// Color
	Color *goflow.Color

	// Semantic label for accessibility
	SemanticLabel string
}

// IconData represents an icon
type IconData struct {
	Name      string
	CodePoint rune
}

// Common Material Icons
var (
	IconHome       = IconData{Name: "home", CodePoint: '\ue88a'}
	IconMenu       = IconData{Name: "menu", CodePoint: '\ue5d2'}
	IconSearch     = IconData{Name: "search", CodePoint: '\ue8b6'}
	IconSettings   = IconData{Name: "settings", CodePoint: '\ue8b8'}
	IconFavorite   = IconData{Name: "favorite", CodePoint: '\ue87d'}
	IconAdd        = IconData{Name: "add", CodePoint: '\ue145'}
	IconRemove     = IconData{Name: "remove", CodePoint: '\ue15b'}
	IconClose      = IconData{Name: "close", CodePoint: '\ue5cd'}
	IconCheck      = IconData{Name: "check", CodePoint: '\ue5ca'}
	IconArrowBack  = IconData{Name: "arrow_back", CodePoint: '\ue5c4'}
	IconArrowForward = IconData{Name: "arrow_forward", CodePoint: '\ue5c8'}
	IconEdit       = IconData{Name: "edit", CodePoint: '\ue3c9'}
	IconDelete     = IconData{Name: "delete", CodePoint: '\ue872'}
	IconShare      = IconData{Name: "share", CodePoint: '\ue80d'}
	IconPerson     = IconData{Name: "person", CodePoint: '\ue7fd'}
	IconEmail      = IconData{Name: "email", CodePoint: '\ue0be'}
	IconPhone      = IconData{Name: "phone", CodePoint: '\ue0cd'}
	IconInfo       = IconData{Name: "info", CodePoint: '\ue88e'}
	IconWarning    = IconData{Name: "warning", CodePoint: '\ue002'}
	IconError      = IconData{Name: "error", CodePoint: '\ue000'}
	IconArrowDropDown = IconData{Name: "arrow_drop_down", CodePoint: '\ue5c5'}
	IconMoreVert   = IconData{Name: "more_vert", CodePoint: '\ue5d4'}
)

// NewIcon creates a new Icon widget
func NewIcon(icon IconData) *Icon {
	return &Icon{
		Icon: icon,
		Size: 24.0,
	}
}

// CreateElement creates a render object element
func (i *Icon) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      i,
		renderObject: &RenderIcon{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			icon:          i.Icon,
			size:          i.Size,
			color:         i.Color,
		},
	}
}

// Build returns nil (this is a render object widget)
func (i *Icon) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderIcon is the render object for Icon
type RenderIcon struct {
	*goflow.BaseRenderBox
	icon  IconData
	size  float64
	color *goflow.Color
}

// PerformLayout performs layout
func (r *RenderIcon) PerformLayout() {
	r.SetSize(r.GetConstraints().Constrain(goflow.NewSize(r.size, r.size)))
}

// Layout performs the layout
func (r *RenderIcon) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the icon
func (r *RenderIcon) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// In a real implementation, this would draw the actual icon glyph
	// For now, we'll use the icon name as text
	style := goflow.NewTextStyle()
	style.FontSize = r.size
	if r.color != nil {
		style.Color = r.color
	} else {
		style.Color = goflow.NewColor(0, 0, 0, 255)
	}

	canvas.DrawText(r.icon.Name, offset, style)
}

// IconButton is a button with an icon
type IconButton struct {
	goflow.BaseWidget

	// Icon to display
	Icon IconData

	// Callback when pressed
	OnPressed func()

	// Icon size
	IconSize float64

	// Icon color
	Color *goflow.Color

	// Tooltip
	Tooltip string

	// Padding
	Padding *goflow.EdgeInsets
}

// NewIconButton creates a new icon button
func NewIconButton(icon IconData, onPressed func()) *IconButton {
	return &IconButton{
		Icon:      icon,
		OnPressed: onPressed,
		IconSize:  24.0,
		Padding:   goflow.NewEdgeInsets(8, 8, 8, 8),
	}
}

// Build creates the widget tree for the icon button
func (i *IconButton) Build(context goflow.BuildContext) goflow.Widget {
	return &Container{
		Padding: i.Padding,
		Child: &Icon{
			Icon:  i.Icon,
			Size:  i.IconSize,
			Color: i.Color,
		},
	}
}
