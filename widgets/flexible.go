package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Flexible allows a child of Row, Column, or Flex to expand
type Flexible struct {
	goflow.BaseWidget
	Child goflow.Widget
	Flex  int        // How much space to take relative to siblings
	Fit   FlexFit    // How to fit the child
}

// FlexFit defines how a flexible child fits within its parent
type FlexFit int

const (
	FlexFitLoose FlexFit = iota // Child can be smaller than flex space
	FlexFitTight                // Child must fill flex space
)

// NewFlexible creates a new Flexible widget
func NewFlexible(child goflow.Widget) *Flexible {
	return &Flexible{
		Child: child,
		Flex:  1,
		Fit:   FlexFitLoose,
	}
}

// Build returns the child (parent will handle flex behavior)
func (f *Flexible) Build(context goflow.BuildContext) goflow.Widget {
	return f.Child
}

// Expanded is a shorthand for Flexible with FlexFitTight
type Expanded struct {
	goflow.BaseWidget
	Child goflow.Widget
	Flex  int
}

// NewExpanded creates a new Expanded widget
func NewExpanded(child goflow.Widget) *Expanded {
	return &Expanded{
		Child: child,
		Flex:  1,
	}
}

// Build returns a Flexible with tight fit
func (e *Expanded) Build(context goflow.BuildContext) goflow.Widget {
	return &Flexible{
		Child: e.Child,
		Flex:  e.Flex,
		Fit:   FlexFitTight,
	}
}

// Spacer creates space between widgets in a Row or Column
type Spacer struct {
	goflow.BaseWidget
	Flex int
}

// NewSpacer creates a new Spacer widget
func NewSpacer() *Spacer {
	return &Spacer{
		Flex: 1,
	}
}

// Build returns an Expanded with an empty container
func (s *Spacer) Build(context goflow.BuildContext) goflow.Widget {
	return &Expanded{
		Child: &Container{},
		Flex:  s.Flex,
	}
}
