package goflow

// Widget is the base interface for all widgets in GoFlow
// Widgets describe the configuration for an Element
type Widget interface {
	// Build describes the UI represented by this widget
	Build(context BuildContext) Widget

	// CreateElement creates the element for this widget
	CreateElement() Element

	// GetKey returns the key for this widget, if any
	GetKey() Key
}

// BuildContext provides access to contextual information during build
type BuildContext interface {
	// GetElement returns the element associated with this context
	GetElement() Element

	// FindAncestorWidgetOfType finds the nearest ancestor widget of the given type
	FindAncestorWidgetOfType(widgetType string) Widget

	// FindAncestorElementOfType finds the nearest ancestor element of the given type
	FindAncestorElementOfType(elementType string) Element

	// GetSize returns the size of the render object, if available
	GetSize() *Size

	// MarkNeedsBuild marks this element as needing to rebuild
	MarkNeedsBuild()
}

// buildContext is the default implementation of BuildContext
type buildContext struct {
	element Element
}

// NewBuildContext creates a new build context
func NewBuildContext(element Element) BuildContext {
	return &buildContext{element: element}
}

func (c *buildContext) GetElement() Element {
	return c.element
}

func (c *buildContext) FindAncestorWidgetOfType(widgetType string) Widget {
	current := c.element.GetParent()
	for current != nil {
		if current.GetWidgetType() == widgetType {
			return current.GetWidget()
		}
		current = current.GetParent()
	}
	return nil
}

func (c *buildContext) FindAncestorElementOfType(elementType string) Element {
	current := c.element.GetParent()
	for current != nil {
		if current.GetElementType() == elementType {
			return current
		}
		current = current.GetParent()
	}
	return nil
}

func (c *buildContext) GetSize() *Size {
	ro := c.element.GetRenderObject()
	if ro != nil {
		return ro.GetSize()
	}
	return nil
}

func (c *buildContext) MarkNeedsBuild() {
	c.element.MarkNeedsBuild()
}
