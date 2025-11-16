package goflow

// BaseWidget provides default implementation for Widget interface
type BaseWidget struct {
	Key Key
}

// GetKey returns the key
func (w *BaseWidget) GetKey() Key {
	return w.Key
}

// CreateElement creates a SimpleElement by default
// Widgets that need special elements (like render object widgets) should override this
func (w *BaseWidget) CreateElement() Element {
	// We can't call Build() on w directly since w is *BaseWidget
	// So we create a simple element that will work with embedding widgets
	return &GenericElement{
		BaseElement: BaseElement{widget: nil}, // Will be set during mount
	}
}

// Build provides a default build that returns nil (leaf widget)
func (w *BaseWidget) Build(context BuildContext) Widget {
	// Default implementation returns nil - override in concrete types
	return nil
}

// SimpleElement is a basic element that calls Build
type SimpleElement struct {
	BaseElement
	widget  Widget
	child   Element
	context BuildContext
}

// NewSimpleElement creates a new simple element
func NewSimpleElement(widget Widget) *SimpleElement {
	return &SimpleElement{
		BaseElement: BaseElement{widget: widget},
		widget:      widget,
	}
}

// GetWidgetType returns the widget type
func (e *SimpleElement) GetWidgetType() string {
	return "Widget"
}

// GetElementType returns the element type
func (e *SimpleElement) GetElementType() string {
	return "SimpleElement"
}

// Mount mounts the element
func (e *SimpleElement) Mount(parent Element, slot interface{}) {
	e.parent = parent
	e.context = NewBuildContext(e)
	e.Rebuild()
}

// Update updates with a new widget
func (e *SimpleElement) Update(newWidget Widget) {
	e.widget = newWidget
	e.BaseElement.widget = newWidget
	e.Rebuild()
}

// Unmount unmounts the element
func (e *SimpleElement) Unmount() {
	if e.child != nil {
		e.child.Unmount()
		e.child = nil
	}
}

// Rebuild rebuilds the element
func (e *SimpleElement) Rebuild() {
	e.dirty = false

	// Build the widget
	built := e.widget.Build(e.context)

	if built == nil {
		// No child
		if e.child != nil {
			e.child.Unmount()
			e.child = nil
		}
		return
	}

	if e.child == nil {
		// First build - create new element
		e.child = built.CreateElement()
		e.child.Mount(e, nil)
	} else {
		// Update existing child
		e.child.Update(built)
	}
}

// VisitChildren visits child elements
func (e *SimpleElement) VisitChildren(visitor func(Element)) {
	if e.child != nil {
		visitor(e.child)
	}
}
