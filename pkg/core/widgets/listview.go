package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// ListView is a scrollable list of widgets
type ListView struct {
	goflow.BaseWidget

	// Children widgets
	Children []goflow.Widget

	// Scroll direction
	ScrollDirection Axis

	// Padding
	Padding *goflow.EdgeInsets

	// Whether to shrink wrap (size to content)
	ShrinkWrap bool

	// Physics (scroll behavior)
	Physics ScrollPhysics
}

// Axis defines scroll direction
type Axis int

const (
	AxisVertical Axis = iota
	AxisHorizontal
)

// ScrollPhysics defines scroll behavior
type ScrollPhysics int

const (
	ScrollPhysicsAlwaysScrollable ScrollPhysics = iota
	ScrollPhysicsNeverScrollable
	ScrollPhysicsBouncing // iOS-style bouncing
	ScrollPhysicsClamping // Android-style clamping
)

// NewListView creates a new ListView
func NewListView(children []goflow.Widget) *ListView {
	return &ListView{
		Children:        children,
		ScrollDirection: AxisVertical,
		ShrinkWrap:      false,
		Physics:         ScrollPhysicsAlwaysScrollable,
	}
}

// ListViewBuilder creates a ListView with a builder function
type ListViewBuilder struct {
	goflow.BaseWidget

	// Item count
	ItemCount int

	// Item builder
	ItemBuilder func(context goflow.BuildContext, index int) goflow.Widget

	// Scroll direction
	ScrollDirection Axis

	// Padding
	Padding *goflow.EdgeInsets
}

// NewListViewBuilder creates a new ListView.builder
func NewListViewBuilder(itemCount int, itemBuilder func(goflow.BuildContext, int) goflow.Widget) *ListViewBuilder {
	return &ListViewBuilder{
		ItemCount:       itemCount,
		ItemBuilder:     itemBuilder,
		ScrollDirection: AxisVertical,
	}
}

// Build creates the widget tree for ListView.builder
func (l *ListViewBuilder) Build(context goflow.BuildContext) goflow.Widget {
	// Build all items
	children := make([]goflow.Widget, l.ItemCount)
	for i := 0; i < l.ItemCount; i++ {
		children[i] = l.ItemBuilder(context, i)
	}

	return &ListView{
		Children:        children,
		ScrollDirection: l.ScrollDirection,
		Padding:         l.Padding,
	}
}

// Build creates the widget tree for ListView
func (l *ListView) Build(context goflow.BuildContext) goflow.Widget {
	// Wrap children in padding if specified
	children := l.Children
	if l.Padding != nil {
		paddedChildren := make([]goflow.Widget, len(children))
		for i, child := range children {
			paddedChildren[i] = &Padding{
				Padding: l.Padding,
				Child:   child,
			}
		}
		children = paddedChildren
	}

	// Use Column for vertical, Row for horizontal
	if l.ScrollDirection == AxisVertical {
		return &Column{
			Children:       children,
			MainAxisAlign:  MainAxisStart,
			CrossAxisAlign: CrossAxisStretch,
		}
	} else {
		return &Row{
			Children:           children,
			MainAxisAlignment:  MainAxisStart,
			CrossAxisAlignment: CrossAxisStart,
		}
	}
}

// SingleChildScrollView is a scrollable view with a single child
type SingleChildScrollView struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Scroll direction
	ScrollDirection Axis

	// Padding
	Padding *goflow.EdgeInsets

	// Physics
	Physics ScrollPhysics
}

// NewSingleChildScrollView creates a new SingleChildScrollView
func NewSingleChildScrollView(child goflow.Widget) *SingleChildScrollView {
	return &SingleChildScrollView{
		Child:           child,
		ScrollDirection: AxisVertical,
		Physics:         ScrollPhysicsAlwaysScrollable,
	}
}

// Build creates the widget tree
func (s *SingleChildScrollView) Build(context goflow.BuildContext) goflow.Widget {
	child := s.Child

	// Apply padding if specified
	if s.Padding != nil {
		child = &Padding{
			Padding: s.Padding,
			Child:   child,
		}
	}

	return child
}

// GridView is a scrollable 2D array of widgets
type GridView struct {
	goflow.BaseWidget

	// Children widgets
	Children []goflow.Widget

	// Number of columns (crossAxisCount)
	CrossAxisCount int

	// Child aspect ratio
	ChildAspectRatio float64

	// Spacing
	MainAxisSpacing  float64
	CrossAxisSpacing float64

	// Padding
	Padding *goflow.EdgeInsets
}

// NewGridView creates a new GridView
func NewGridView(children []goflow.Widget, crossAxisCount int) *GridView {
	return &GridView{
		Children:         children,
		CrossAxisCount:   crossAxisCount,
		ChildAspectRatio: 1.0,
		MainAxisSpacing:  0.0,
		CrossAxisSpacing: 0.0,
	}
}

// Build creates the widget tree for GridView
func (g *GridView) Build(context goflow.BuildContext) goflow.Widget {
	// Group children into rows
	rows := make([]goflow.Widget, 0)

	for i := 0; i < len(g.Children); i += g.CrossAxisCount {
		end := i + g.CrossAxisCount
		if end > len(g.Children) {
			end = len(g.Children)
		}

		rowChildren := g.Children[i:end]

		// Wrap each child in Expanded to fill row evenly
		expandedChildren := make([]goflow.Widget, len(rowChildren))
		for j, child := range rowChildren {
			expandedChildren[j] = &Expanded{
				Child: child,
			}
		}

		row := &Row{
			Children:           expandedChildren,
			MainAxisAlignment:  MainAxisStart,
			CrossAxisAlignment: CrossAxisStart,
		}

		rows = append(rows, row)
	}

	// Create column of rows
	return &Column{
		Children:       rows,
		MainAxisAlign:  MainAxisStart,
		CrossAxisAlign: CrossAxisStretch,
	}
}
