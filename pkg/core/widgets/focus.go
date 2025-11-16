package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// FocusNode represents a node in the focus tree
// It manages focus state for a single widget
type FocusNode struct {
	// Parent scope
	parent *FocusScopeNode

	// Callbacks
	OnFocusChange func(hasFocus bool)

	// State
	hasFocus       bool
	canRequestFocus bool
	skipTraversal  bool

	// Listeners
	listeners []func(bool)
}

// NewFocusNode creates a new FocusNode
func NewFocusNode() *FocusNode {
	return &FocusNode{
		canRequestFocus: true,
		skipTraversal:   false,
		listeners:       make([]func(bool), 0),
	}
}

// HasFocus returns whether this node currently has focus
func (f *FocusNode) HasFocus() bool {
	return f.hasFocus
}

// RequestFocus requests focus for this node
func (f *FocusNode) RequestFocus() {
	if !f.canRequestFocus {
		return
	}

	if f.parent != nil {
		f.parent.RequestFocusForNode(f)
	} else {
		f.setFocus(true)
	}
}

// Unfocus removes focus from this node
func (f *FocusNode) Unfocus() {
	if f.hasFocus {
		f.setFocus(false)
	}
}

// AddListener adds a focus change listener
func (f *FocusNode) AddListener(listener func(bool)) {
	f.listeners = append(f.listeners, listener)
}

// RemoveListener removes a focus change listener
func (f *FocusNode) RemoveListener(listener func(bool)) {
	for i, l := range f.listeners {
		// Compare function pointers (this is a simplification)
		if &l == &listener {
			f.listeners = append(f.listeners[:i], f.listeners[i+1:]...)
			return
		}
	}
}

// setFocus updates focus state and notifies listeners
func (f *FocusNode) setFocus(value bool) {
	if f.hasFocus == value {
		return
	}

	f.hasFocus = value

	// Notify callback
	if f.OnFocusChange != nil {
		f.OnFocusChange(value)
	}

	// Notify listeners
	for _, listener := range f.listeners {
		listener(value)
	}
}

// Dispose cleans up the focus node
func (f *FocusNode) Dispose() {
	f.Unfocus()
	f.listeners = nil
	f.parent = nil
}

// FocusScopeNode manages a scope of focusable widgets
// It handles focus traversal (tab navigation)
type FocusScopeNode struct {
	*FocusNode

	// Children nodes
	children []*FocusNode

	// Currently focused child
	focusedChild *FocusNode

	// Auto-focus first child
	autoFocus bool
}

// NewFocusScopeNode creates a new FocusScopeNode
func NewFocusScopeNode() *FocusScopeNode {
	return &FocusScopeNode{
		FocusNode: NewFocusNode(),
		children:  make([]*FocusNode, 0),
	}
}

// RegisterChild registers a focus node with this scope
func (f *FocusScopeNode) RegisterChild(child *FocusNode) {
	child.parent = f
	f.children = append(f.children, child)

	// Auto-focus first child if enabled
	if f.autoFocus && len(f.children) == 1 {
		child.RequestFocus()
	}
}

// UnregisterChild removes a focus node from this scope
func (f *FocusScopeNode) UnregisterChild(child *FocusNode) {
	for i, c := range f.children {
		if c == child {
			f.children = append(f.children[:i], f.children[i+1:]...)
			child.parent = nil

			// Clear focused child if it was this node
			if f.focusedChild == child {
				f.focusedChild = nil
			}
			return
		}
	}
}

// RequestFocusForNode gives focus to a specific child node
func (f *FocusScopeNode) RequestFocusForNode(node *FocusNode) {
	// Unfocus current child
	if f.focusedChild != nil && f.focusedChild != node {
		f.focusedChild.setFocus(false)
	}

	// Focus new child
	f.focusedChild = node
	node.setFocus(true)
}

// NextFocus moves focus to the next focusable widget
func (f *FocusScopeNode) NextFocus() bool {
	if len(f.children) == 0 {
		return false
	}

	// Find current focused index
	currentIndex := -1
	if f.focusedChild != nil {
		for i, child := range f.children {
			if child == f.focusedChild {
				currentIndex = i
				break
			}
		}
	}

	// Find next focusable widget
	nextIndex := currentIndex + 1
	for i := 0; i < len(f.children); i++ {
		idx := (nextIndex + i) % len(f.children)
		child := f.children[idx]

		if !child.skipTraversal && child.canRequestFocus {
			child.RequestFocus()
			return true
		}
	}

	return false
}

// PreviousFocus moves focus to the previous focusable widget
func (f *FocusScopeNode) PreviousFocus() bool {
	if len(f.children) == 0 {
		return false
	}

	// Find current focused index
	currentIndex := len(f.children)
	if f.focusedChild != nil {
		for i, child := range f.children {
			if child == f.focusedChild {
				currentIndex = i
				break
			}
		}
	}

	// Find previous focusable widget
	prevIndex := currentIndex - 1
	if prevIndex < 0 {
		prevIndex = len(f.children) - 1
	}

	for i := 0; i < len(f.children); i++ {
		idx := prevIndex - i
		if idx < 0 {
			idx += len(f.children)
		}

		child := f.children[idx]
		if !child.skipTraversal && child.canRequestFocus {
			child.RequestFocus()
			return true
		}
	}

	return false
}

// Unfocus removes focus from all children
func (f *FocusScopeNode) Unfocus() {
	if f.focusedChild != nil {
		f.focusedChild.setFocus(false)
		f.focusedChild = nil
	}
	f.FocusNode.Unfocus()
}

// Dispose cleans up the scope
func (f *FocusScopeNode) Dispose() {
	for _, child := range f.children {
		child.Dispose()
	}
	f.children = nil
	f.focusedChild = nil
	f.FocusNode.Dispose()
}

// Focus widget attaches a FocusNode to a widget subtree
type Focus struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Focus node (optional - will be created if nil)
	FocusNode *FocusNode

	// Auto-focus this widget when mounted
	AutoFocus bool

	// Skip this widget in focus traversal
	SkipTraversal bool

	// Can this widget request focus
	CanRequestFocus bool

	// Focus change callback
	OnFocusChange func(hasFocus bool)
}

// NewFocus creates a new Focus widget
func NewFocus(child goflow.Widget, focusNode *FocusNode) *Focus {
	return &Focus{
		Child:           child,
		FocusNode:       focusNode,
		AutoFocus:       false,
		SkipTraversal:   false,
		CanRequestFocus: true,
	}
}

// Build returns the child (focus handling happens at element level)
func (f *Focus) Build(context goflow.BuildContext) goflow.Widget {
	// Initialize focus node if not provided
	if f.FocusNode == nil {
		f.FocusNode = NewFocusNode()
	}

	// Configure focus node
	f.FocusNode.skipTraversal = f.SkipTraversal
	f.FocusNode.canRequestFocus = f.CanRequestFocus
	f.FocusNode.OnFocusChange = f.OnFocusChange

	// Auto-focus if requested
	if f.AutoFocus {
		f.FocusNode.RequestFocus()
	}

	return f.Child
}

// FocusScope widget creates a scope for focus traversal
type FocusScope struct {
	goflow.BaseWidget

	// Child widget
	Child goflow.Widget

	// Scope node (optional - will be created if nil)
	Node *FocusScopeNode

	// Auto-focus first child
	AutoFocus bool
}

// NewFocusScope creates a new FocusScope widget
func NewFocusScope(child goflow.Widget, node *FocusScopeNode) *FocusScope {
	return &FocusScope{
		Child:     child,
		Node:      node,
		AutoFocus: false,
	}
}

// Build returns the child (scope handling happens at element level)
func (f *FocusScope) Build(context goflow.BuildContext) goflow.Widget {
	// Initialize scope node if not provided
	if f.Node == nil {
		f.Node = NewFocusScopeNode()
	}

	f.Node.autoFocus = f.AutoFocus

	return f.Child
}
