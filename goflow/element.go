package goflow

// Element represents a node in the element tree
// Elements are mutable and manage the widget lifecycle
type Element interface {
	// GetWidget returns the current widget for this element
	GetWidget() Widget

	// GetWidgetType returns the type name of the widget
	GetWidgetType() string

	// GetElementType returns the type name of the element
	GetElementType() string

	// GetParent returns the parent element
	GetParent() Element

	// SetParent sets the parent element
	SetParent(parent Element)

	// GetRenderObject returns the render object, if this is a RenderObjectElement
	GetRenderObject() RenderObject

	// Mount mounts this element into the tree at the given parent
	Mount(parent Element, slot interface{})

	// Update updates this element with a new widget
	Update(newWidget Widget)

	// Unmount removes this element from the tree
	Unmount()

	// MarkNeedsBuild marks this element as needing to rebuild
	MarkNeedsBuild()

	// Rebuild rebuilds this element
	Rebuild()

	// VisitChildren calls the visitor function for each child element
	VisitChildren(visitor func(Element))
}

// BaseElement provides common functionality for all elements
type BaseElement struct {
	widget       Widget
	parent       Element
	dirty        bool
	renderObject RenderObject
}

// GetWidget returns the current widget
func (e *BaseElement) GetWidget() Widget {
	return e.widget
}

// GetWidgetType returns the widget type name
func (e *BaseElement) GetWidgetType() string {
	// This should be overridden by concrete types
	return "Widget"
}

// GetElementType returns the element type name
func (e *BaseElement) GetElementType() string {
	return "Element"
}

// GetParent returns the parent element
func (e *BaseElement) GetParent() Element {
	return e.parent
}

// SetParent sets the parent element
func (e *BaseElement) SetParent(parent Element) {
	e.parent = parent
}

// GetRenderObject returns the render object
func (e *BaseElement) GetRenderObject() RenderObject {
	return e.renderObject
}

// MarkNeedsBuild marks the element as dirty
func (e *BaseElement) MarkNeedsBuild() {
	e.dirty = true
	// In a full implementation, this would schedule a rebuild
}

// VisitChildren is a no-op for base element
func (e *BaseElement) VisitChildren(visitor func(Element)) {
	// Override in subclasses
}
