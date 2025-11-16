package widgets

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/input"
)

// Focus wraps a child widget and adds focus capabilities
// Similar to Flutter's Focus widget
type Focus struct {
	// Child widget
	Child goflow.Widget

	// Focus node (if nil, one will be created)
	FocusNode *input.FocusNode

	// Whether to autofocus when mounted
	Autofocus bool

	// Whether this node can request focus
	CanRequestFocus bool

	// Skip this node during focus traversal
	SkipTraversal bool

	// Callback when focus changes
	OnFocusChange func(bool)

	// Callback to handle key events
	OnKey func(*input.KeyboardEvent) bool

	// Debug label
	DebugLabel string

	// Descendants are focusable
	DescendantsAreFocusable bool
}

// NewFocus creates a new Focus widget
func NewFocus(child goflow.Widget, options ...FocusOption) *Focus {
	f := &Focus{
		Child:                   child,
		CanRequestFocus:         true,
		DescendantsAreFocusable: true,
	}

	for _, opt := range options {
		opt(f)
	}

	return f
}

// FocusOption configures a Focus widget
type FocusOption func(*Focus)

// WithFocusNode sets an existing focus node
func WithFocusNode(node *input.FocusNode) FocusOption {
	return func(f *Focus) {
		f.FocusNode = node
	}
}

// WithAutofocus enables autofocus
func WithAutofocus(autofocus bool) FocusOption {
	return func(f *Focus) {
		f.Autofocus = autofocus
	}
}

// WithFocusCanRequestFocus sets whether the node can request focus
func WithFocusCanRequestFocus(canRequest bool) FocusOption {
	return func(f *Focus) {
		f.CanRequestFocus = canRequest
	}
}

// WithFocusSkipTraversal sets whether to skip during traversal
func WithFocusSkipTraversal(skip bool) FocusOption {
	return func(f *Focus) {
		f.SkipTraversal = skip
	}
}

// WithFocusOnChange sets the focus change callback
func WithFocusOnChange(callback func(bool)) FocusOption {
	return func(f *Focus) {
		f.OnFocusChange = callback
	}
}

// WithFocusOnKey sets the key event callback
func WithFocusOnKey(callback func(*input.KeyboardEvent) bool) FocusOption {
	return func(f *Focus) {
		f.OnKey = callback
	}
}

// WithFocusDebugLabel sets a debug label
func WithFocusDebugLabel(label string) FocusOption {
	return func(f *Focus) {
		f.DebugLabel = label
	}
}

// Build implements the Widget interface
func (f *Focus) Build(ctx goflow.BuildContext) goflow.Widget {
	return f.Child
}

// CreateElement implements the Widget interface
func (f *Focus) CreateElement() goflow.Element {
	return &FocusElement{
		widget: f,
	}
}

// FocusElement is the element for the Focus widget
type FocusElement struct {
	goflow.BaseElement
	widget    *Focus
	focusNode *input.FocusNode
}

// Mount implements the Element interface
func (e *FocusElement) Mount(parent goflow.Element, newSlot interface{}) {
	e.BaseElement.Mount(parent, newSlot)

	// Create or use existing focus node
	if e.widget.FocusNode != nil {
		e.focusNode = e.widget.FocusNode
	} else {
		options := []input.FocusNodeOption{
			input.WithCanRequestFocus(e.widget.CanRequestFocus),
			input.WithSkipTraversal(e.widget.SkipTraversal),
		}

		if e.widget.DebugLabel != "" {
			options = append(options, input.WithDebugLabel(e.widget.DebugLabel))
		}

		if e.widget.OnFocusChange != nil {
			options = append(options, input.WithOnFocusChange(e.widget.OnFocusChange))
		}

		if e.widget.OnKey != nil {
			options = append(options, input.WithOnKey(e.widget.OnKey))
		}

		e.focusNode = input.NewFocusNode(options...)
	}

	// Register with focus manager
	input.GetFocusManager().RegisterNode(e.focusNode)

	// Autofocus if requested
	if e.widget.Autofocus {
		e.focusNode.RequestFocus()
	}
}

// Unmount implements the Element interface
func (e *FocusElement) Unmount() {
	// Unregister from focus manager
	if e.focusNode != nil {
		input.GetFocusManager().UnregisterNode(e.focusNode)

		// Only dispose if we created it
		if e.widget.FocusNode == nil {
			e.focusNode.Dispose()
		}
	}

	e.BaseElement.Unmount()
}

// Update implements the Element interface
func (e *FocusElement) Update(newWidget goflow.Widget) {
	oldWidget := e.widget
	e.widget = newWidget.(*Focus)

	// If focus node changed, update it
	if oldWidget.FocusNode != e.widget.FocusNode && e.widget.FocusNode != nil {
		// Unregister old node
		if e.focusNode != nil {
			input.GetFocusManager().UnregisterNode(e.focusNode)
			if oldWidget.FocusNode == nil {
				e.focusNode.Dispose()
			}
		}

		// Use new node
		e.focusNode = e.widget.FocusNode
		input.GetFocusManager().RegisterNode(e.focusNode)
	}

	e.BaseElement.Update(newWidget)
}

// FocusScope creates a scope for managing focus within a subtree
// Similar to Flutter's FocusScope widget
type FocusScope struct {
	// Child widget
	Child goflow.Widget

	// Focus scope node (if nil, one will be created)
	Node *input.FocusScopeNode

	// Whether to autofocus the first focusable child
	Autofocus bool

	// Callback when focus changes within this scope
	OnFocusChange func(bool)

	// Debug label
	DebugLabel string
}

// NewFocusScope creates a new FocusScope widget
func NewFocusScope(child goflow.Widget, options ...FocusScopeOption) *FocusScope {
	fs := &FocusScope{
		Child: child,
	}

	for _, opt := range options {
		opt(fs)
	}

	return fs
}

// FocusScopeOption configures a FocusScope widget
type FocusScopeOption func(*FocusScope)

// WithScopeNode sets an existing scope node
func WithScopeNode(node *input.FocusScopeNode) FocusScopeOption {
	return func(fs *FocusScope) {
		fs.Node = node
	}
}

// WithScopeAutofocus enables autofocus for the scope
func WithScopeAutofocus(autofocus bool) FocusScopeOption {
	return func(fs *FocusScope) {
		fs.Autofocus = autofocus
	}
}

// WithScopeOnFocusChange sets the focus change callback
func WithScopeOnFocusChange(callback func(bool)) FocusScopeOption {
	return func(fs *FocusScope) {
		fs.OnFocusChange = callback
	}
}

// WithScopeDebugLabel sets a debug label
func WithScopeDebugLabel(label string) FocusScopeOption {
	return func(fs *FocusScope) {
		fs.DebugLabel = label
	}
}

// Build implements the Widget interface
func (fs *FocusScope) Build(ctx goflow.BuildContext) goflow.Widget {
	return fs.Child
}

// CreateElement implements the Widget interface
func (fs *FocusScope) CreateElement() goflow.Element {
	return &FocusScopeElement{
		widget: fs,
	}
}

// FocusScopeElement is the element for the FocusScope widget
type FocusScopeElement struct {
	goflow.BaseElement
	widget    *FocusScope
	scopeNode *input.FocusScopeNode
}

// Mount implements the Element interface
func (e *FocusScopeElement) Mount(parent goflow.Element, newSlot interface{}) {
	e.BaseElement.Mount(parent, newSlot)

	// Create or use existing scope node
	if e.widget.Node != nil {
		e.scopeNode = e.widget.Node
	} else {
		e.scopeNode = input.NewFocusScopeNode()
	}

	// Autofocus if requested
	if e.widget.Autofocus {
		e.scopeNode.AutoFocus()
	}
}

// Unmount implements the Element interface
func (e *FocusScopeElement) Unmount() {
	// Clean up scope node if we created it
	if e.scopeNode != nil && e.widget.Node == nil {
		e.scopeNode.Dispose()
	}

	e.BaseElement.Unmount()
}

// Update implements the Element interface
func (e *FocusScopeElement) Update(newWidget goflow.Widget) {
	oldWidget := e.widget
	e.widget = newWidget.(*FocusScope)

	// If scope node changed, update it
	if oldWidget.Node != e.widget.Node && e.widget.Node != nil {
		// Clean up old node
		if e.scopeNode != nil && oldWidget.Node == nil {
			e.scopeNode.Dispose()
		}

		// Use new node
		e.scopeNode = e.widget.Node
	}

	e.BaseElement.Update(newWidget)
}
