package material

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// Scaffold implements the basic Material Design visual layout structure
type Scaffold struct {
	goflow.BaseWidget

	// AppBar at the top
	AppBar goflow.Widget

	// Main content body
	Body goflow.Widget

	// FloatingActionButton
	FloatingActionButton goflow.Widget

	// FAB position
	FloatingActionButtonLocation FABLocation

	// Drawer (side menu)
	Drawer goflow.Widget

	// BottomNavigationBar
	BottomNavigationBar goflow.Widget

	// BottomSheet
	BottomSheet goflow.Widget

	// Background color
	BackgroundColor *goflow.Color

	// Whether to resize when keyboard appears
	ResizeToAvoidBottomInset bool
}

// FABLocation defines where the floating action button is positioned
type FABLocation int

const (
	FABLocationEndFloat FABLocation = iota // Bottom right
	FABLocationCenterFloat                 // Bottom center
	FABLocationStartFloat                  // Bottom left
	FABLocationEndDocked                   // Docked to bottom app bar, right
	FABLocationCenterDocked                // Docked to bottom app bar, center
)

// NewScaffold creates a new Material Scaffold
func NewScaffold() *Scaffold {
	return &Scaffold{
		FloatingActionButtonLocation: FABLocationEndFloat,
		ResizeToAvoidBottomInset:     true,
	}
}

// Build creates the widget tree for the scaffold
func (s *Scaffold) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := s.BackgroundColor
	if bgColor == nil {
		bgColor = theme.BackgroundColor
	}

	var children []goflow.Widget

	// Add app bar if present
	if s.AppBar != nil {
		children = append(children, s.AppBar)
	}

	// Add body if present
	if s.Body != nil {
		children = append(children, &widgets.Expanded{
			Child: &widgets.Container{
				Color: bgColor,
				Child: s.Body,
			},
		})
	}

	// Add bottom navigation bar if present
	if s.BottomNavigationBar != nil {
		children = append(children, s.BottomNavigationBar)
	}

	// Create the main column layout
	mainContent := &widgets.Column{
		Children:       children,
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
		MainAxisSize:   widgets.MainAxisSizeMax,
	}

	// If we have FAB or drawer, wrap in a Stack
	if s.FloatingActionButton != nil || s.Drawer != nil {
		stackChildren := []goflow.Widget{mainContent}

		// Add FAB
		if s.FloatingActionButton != nil {
			// Position FAB based on location
			stackChildren = append(stackChildren, s.FloatingActionButton)
		}

		// Note: Drawer would be shown/hidden based on state (simplified here)

		return &widgets.Stack{
			Children:  stackChildren,
			Alignment: widgets.StackAlignmentTopLeft,
			Fit:       widgets.StackFitExpand,
		}
	}

	return mainContent
}
