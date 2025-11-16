package cupertino

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// CupertinoPageScaffold implements the basic iOS page structure
type CupertinoPageScaffold struct {
	goflow.BaseWidget

	// NavigationBar at the top
	NavigationBar goflow.Widget

	// Main content
	Child goflow.Widget

	// Background color
	BackgroundColor *goflow.Color

	// Whether to resize when keyboard appears
	ResizeToAvoidBottomInset bool
}

// NewCupertinoPageScaffold creates a new iOS-style scaffold
func NewCupertinoPageScaffold(child goflow.Widget) *CupertinoPageScaffold {
	return &CupertinoPageScaffold{
		Child:                    child,
		ResizeToAvoidBottomInset: true,
	}
}

// Build creates the widget tree for the scaffold
func (c *CupertinoPageScaffold) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := c.BackgroundColor
	if bgColor == nil {
		bgColor = theme.BackgroundColor
	}

	var children []goflow.Widget

	// Add navigation bar if present
	if c.NavigationBar != nil {
		children = append(children, c.NavigationBar)
	}

	// Add main content
	if c.Child != nil {
		children = append(children, &widgets.Expanded{
			Child: &widgets.Container{
				Color: bgColor,
				Child: c.Child,
			},
		})
	}

	return &widgets.Column{
		Children:       children,
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStretch,
		MainAxisSize:   widgets.MainAxisSizeMax,
	}
}

// CupertinoTabScaffold implements an iOS-style tabbed interface
type CupertinoTabScaffold struct {
	goflow.BaseWidget

	// Tab bar at the bottom
	TabBar goflow.Widget

	// Tab views (one for each tab)
	TabBuilder func(context goflow.BuildContext, index int) goflow.Widget

	// Current tab index
	CurrentIndex int

	// Background color
	BackgroundColor *goflow.Color
}

// NewCupertinoTabScaffold creates a new iOS-style tab scaffold
func NewCupertinoTabScaffold(tabBar goflow.Widget, tabBuilder func(goflow.BuildContext, int) goflow.Widget) *CupertinoTabScaffold {
	return &CupertinoTabScaffold{
		TabBar:       tabBar,
		TabBuilder:   tabBuilder,
		CurrentIndex: 0,
	}
}

// Build creates the widget tree for the tab scaffold
func (c *CupertinoTabScaffold) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	bgColor := c.BackgroundColor
	if bgColor == nil {
		bgColor = theme.BackgroundColor
	}

	// Build current tab content
	currentTabContent := c.TabBuilder(context, c.CurrentIndex)

	return &widgets.Column{
		Children: []goflow.Widget{
			// Tab content
			&widgets.Expanded{
				Child: &widgets.Container{
					Color: bgColor,
					Child: currentTabContent,
				},
			},
			// Tab bar
			c.TabBar,
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStretch,
		MainAxisSize:   widgets.MainAxisSizeMax,
	}
}
