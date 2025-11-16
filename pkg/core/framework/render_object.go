package goflow

// RenderObject is responsible for layout, painting, and hit testing
// This is the base interface for all render objects
type RenderObject interface {
	// Layout computes the size of this render object given the constraints
	Layout(constraints *Constraints)

	// GetSize returns the size computed during layout
	GetSize() *Size

	// SetSize sets the size (used during layout)
	SetSize(size *Size)

	// GetOffset returns the offset of this render object relative to its parent
	GetOffset() *Offset

	// SetOffset sets the offset
	SetOffset(offset *Offset)

	// GetConstraints returns the constraints used in the last layout
	GetConstraints() *Constraints

	// Paint paints this render object to the canvas
	Paint(canvas Canvas, offset *Offset)

	// HitTest checks if a point hits this render object
	HitTest(position *Offset) bool

	// MarkNeedsLayout marks this render object as needing layout
	MarkNeedsLayout()

	// MarkNeedsPaint marks this render object as needing paint
	MarkNeedsPaint()

	// GetParent returns the parent render object
	GetParent() RenderObject

	// SetParent sets the parent render object
	SetParent(parent RenderObject)

	// VisitChildren visits all child render objects
	VisitChildren(visitor func(RenderObject))
}

// BaseRenderObject provides common functionality for render objects
type BaseRenderObject struct {
	size        *Size
	offset      *Offset
	constraints *Constraints
	parent      RenderObject
	needsLayout bool
	needsPaint  bool
}

// NewBaseRenderObject creates a new base render object
func NewBaseRenderObject() *BaseRenderObject {
	return &BaseRenderObject{
		size:        ZeroSize(),
		offset:      ZeroOffset(),
		needsLayout: true,
		needsPaint:  true,
	}
}

func (r *BaseRenderObject) GetSize() *Size {
	return r.size
}

func (r *BaseRenderObject) SetSize(size *Size) {
	r.size = size
}

func (r *BaseRenderObject) GetOffset() *Offset {
	return r.offset
}

func (r *BaseRenderObject) SetOffset(offset *Offset) {
	r.offset = offset
}

func (r *BaseRenderObject) GetConstraints() *Constraints {
	return r.constraints
}

func (r *BaseRenderObject) MarkNeedsLayout() {
	r.needsLayout = true
	r.needsPaint = true
}

func (r *BaseRenderObject) MarkNeedsPaint() {
	r.needsPaint = true
}

func (r *BaseRenderObject) GetParent() RenderObject {
	return r.parent
}

func (r *BaseRenderObject) SetParent(parent RenderObject) {
	r.parent = parent
}

func (r *BaseRenderObject) VisitChildren(visitor func(RenderObject)) {
	// Override in subclasses with children
}

// RenderBox is a render object in the box protocol
// It has a size and can be laid out in 2D
type RenderBox interface {
	RenderObject

	// ComputeSize computes the size given constraints
	ComputeSize(constraints *Constraints) *Size

	// PerformLayout performs the layout for this render object
	PerformLayout()
}

// BaseRenderBox provides base implementation for RenderBox
type BaseRenderBox struct {
	*BaseRenderObject
}

// NewBaseRenderBox creates a new base render box
func NewBaseRenderBox() *BaseRenderBox {
	return &BaseRenderBox{
		BaseRenderObject: NewBaseRenderObject(),
	}
}

// Layout performs layout
func (r *BaseRenderBox) Layout(constraints *Constraints) {
	r.constraints = constraints
	r.needsLayout = false
	// Subclasses should override PerformLayout
}

// HitTest default implementation
func (r *BaseRenderBox) HitTest(position *Offset) bool {
	return position.X >= 0 && position.X <= r.size.Width &&
		position.Y >= 0 && position.Y <= r.size.Height
}
