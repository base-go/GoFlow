package material

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Drawer is a Material Design navigation drawer
type Drawer struct {
	goflow.BaseWidget

	// Child widget (drawer content)
	Child goflow.Widget

	// Background color
	BackgroundColor *goflow.Color

	// Elevation
	Elevation float64

	// Width
	Width *float64
}

// NewDrawer creates a new Material drawer
func NewDrawer(child goflow.Widget) *Drawer {
	defaultWidth := 304.0 // Material Design standard
	return &Drawer{
		Child:     child,
		Elevation: 16.0,
		Width:     &defaultWidth,
	}
}

// Build creates the widget tree for the drawer
func (d *Drawer) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := d.BackgroundColor
	if bgColor == nil {
		bgColor = theme.SurfaceColor
	}

	return &widgets.Container{
		Width: d.Width,
		Color: bgColor,
		Child: d.Child,
	}
}

// DrawerHeader is a header for a drawer
type DrawerHeader struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Decoration
	Decoration *BoxDecoration

	// Padding
	Padding *goflow.EdgeInsets

	// Margin
	Margin *goflow.EdgeInsets
}

// BoxDecoration defines decoration for a box
type BoxDecoration struct {
	Color        *goflow.Color
	Gradient     *Gradient
	BorderRadius float64
}

// Gradient defines a gradient
type Gradient struct {
	Colors []goflow.Color
	Stops  []float64
}

// NewDrawerHeader creates a new drawer header
func NewDrawerHeader(child goflow.Widget) *DrawerHeader {
	return &DrawerHeader{
		Child:   child,
		Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
		Margin:  goflow.ZeroEdgeInsets(),
	}
}

// Build creates the widget tree for the drawer header
func (d *DrawerHeader) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	height := 160.0 // Material Design standard
	bgColor := theme.PrimaryColor

	if d.Decoration != nil && d.Decoration.Color != nil {
		bgColor = d.Decoration.Color
	}

	return &widgets.Container{
		Height:  &height,
		Color:   bgColor,
		Padding: d.Padding,
		Margin:  d.Margin,
		Child:   d.Child,
	}
}

// ListTile is a single fixed-height row (typically used in lists)
type ListTile struct {
	goflow.BaseWidget

	// Leading widget (usually an icon or avatar)
	Leading goflow.Widget

	// Title widget
	Title goflow.Widget

	// Subtitle widget
	Subtitle goflow.Widget

	// Trailing widget
	Trailing goflow.Widget

	// Is it three lines?
	IsThreeLine bool

	// Dense (smaller height)
	Dense bool

	// Enabled
	Enabled bool

	// Selected
	Selected bool

	// On tap callback
	OnTap func()

	// On long press callback
	OnLongPress func()

	// Content padding
	ContentPadding *goflow.EdgeInsets
}

// NewListTile creates a new ListTile
func NewListTile(title goflow.Widget) *ListTile {
	return &ListTile{
		Title:          title,
		Enabled:        true,
		Selected:       false,
		ContentPadding: goflow.NewEdgeInsets(16, 8, 16, 8),
	}
}

// Build creates the widget tree for the list tile
func (l *ListTile) Build(context goflow.BuildContext) goflow.Widget {
	var children []goflow.Widget

	// Add leading if present
	if l.Leading != nil {
		children = append(children, l.Leading)
	}

	// Build title/subtitle column
	titleColumn := []goflow.Widget{l.Title}
	if l.Subtitle != nil {
		titleColumn = append(titleColumn, l.Subtitle)
	}

	children = append(children, &widgets.Expanded{
		Child: &widgets.Column{
			Children:       titleColumn,
			MainAxisAlign:  widgets.MainAxisCenter,
			CrossAxisAlign: widgets.CrossAxisStart,
		},
	})

	// Add trailing if present
	if l.Trailing != nil {
		children = append(children, l.Trailing)
	}

	minHeight := 56.0
	if l.Dense {
		minHeight = 48.0
	}
	if l.IsThreeLine {
		minHeight = 88.0
	}

	return &widgets.Container{
		Padding: l.ContentPadding,
		Height:  &minHeight,
		Child: &widgets.Row{
			Children:           children,
			MainAxisAlignment:  widgets.MainAxisStart,
			CrossAxisAlignment: widgets.CrossAxisCenter,
		},
	}
}
