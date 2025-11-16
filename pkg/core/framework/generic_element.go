package goflow

// GenericElement is a generic element that works with any widget
type GenericElement struct {
	BaseElement
	child   Element
	context BuildContext
}

// GetWidgetType returns the widget type
func (e *GenericElement) GetWidgetType() string {
	return "Widget"
}

// GetElementType returns the element type
func (e *GenericElement) GetElementType() string {
	return "GenericElement"
}

// Mount mounts the element
func (e *GenericElement) Mount(parent Element, slot interface{}) {
	e.parent = parent
	e.context = NewBuildContext(e)
	e.Rebuild()
}

// Update updates with a new widget
func (e *GenericElement) Update(newWidget Widget) {
	e.widget = newWidget
	e.BaseElement.widget = newWidget
	e.Rebuild()
}

// Unmount unmounts the element
func (e *GenericElement) Unmount() {
	if e.child != nil {
		e.child.Unmount()
		e.child = nil
	}
}

// Rebuild rebuilds the element
func (e *GenericElement) Rebuild() {
	e.dirty = false

	if e.widget == nil {
		return
	}

	// Build the widget
	built := e.widget.Build(e.context)

	if built == nil {
		// No child - this is a leaf widget
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
func (e *GenericElement) VisitChildren(visitor func(Element)) {
	if e.child != nil {
		visitor(e.child)
	}
}
