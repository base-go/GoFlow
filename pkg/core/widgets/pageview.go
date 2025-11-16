package widgets

import (
	"math"

	"github.com/base-go/GoFlow/pkg/core/framework"
)

// PageController manages page scrolling
type PageController struct {
	initialPage      int
	keepPage         bool
	viewportFraction float64
	currentPage      float64
	listeners        []func()
}

// NewPageController creates a new page controller
func NewPageController() *PageController {
	return &PageController{
		initialPage:      0,
		keepPage:         true,
		viewportFraction: 1.0,
		currentPage:      0,
		listeners:        make([]func(), 0),
	}
}

// NewPageControllerWithPage creates a page controller with initial page
func NewPageControllerWithPage(initialPage int) *PageController {
	return &PageController{
		initialPage:      initialPage,
		keepPage:         true,
		viewportFraction: 1.0,
		currentPage:      float64(initialPage),
		listeners:        make([]func(), 0),
	}
}

// Page returns the current page index
func (p *PageController) Page() int {
	return int(math.Round(p.currentPage))
}

// PageFloat returns the current page as a float (for partial pages)
func (p *PageController) PageFloat() float64 {
	return p.currentPage
}

// JumpToPage instantly jumps to the given page
func (p *PageController) JumpToPage(page int) {
	p.setPage(float64(page))
}

// AnimateToPage animates to the given page (simplified - instant for now)
func (p *PageController) AnimateToPage(page int, duration float64) {
	// In a real implementation, this would animate over time
	p.setPage(float64(page))
}

// NextPage goes to the next page
func (p *PageController) NextPage(duration float64) {
	p.AnimateToPage(p.Page()+1, duration)
}

// PreviousPage goes to the previous page
func (p *PageController) PreviousPage(duration float64) {
	p.AnimateToPage(p.Page()-1, duration)
}

func (p *PageController) setPage(page float64) {
	if p.currentPage != page {
		p.currentPage = page
		p.notifyListeners()
	}
}

// AddListener adds a change listener
func (p *PageController) AddListener(listener func()) {
	p.listeners = append(p.listeners, listener)
}

// RemoveListener removes a change listener
func (p *PageController) RemoveListener(listener func()) {
	for i, l := range p.listeners {
		if &l == &listener {
			p.listeners = append(p.listeners[:i], p.listeners[i+1:]...)
			return
		}
	}
}

func (p *PageController) notifyListeners() {
	for _, listener := range p.listeners {
		listener()
	}
}

// Dispose cleans up the controller
func (p *PageController) Dispose() {
	p.listeners = nil
}

// PageView displays pages that can be swiped
type PageView struct {
	goflow.BaseWidget
	Children         []goflow.Widget
	Controller       *PageController
	ScrollDirection  Axis
	PageSnapping     bool
	OnPageChanged    func(int)
	Reverse          bool
	AllowImplicitScrolling bool
}

// NewPageView creates a new page view
func NewPageView(children []goflow.Widget) *PageView {
	return &PageView{
		Children:        children,
		Controller:      NewPageController(),
		ScrollDirection: AxisHorizontal,
		PageSnapping:    true,
		AllowImplicitScrolling: false,
	}
}

// CreateElement creates a render object element for PageView
func (p *PageView) CreateElement() goflow.Element {
	return &PageViewRenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      p,
		renderObject: &RenderPageView{
			BaseRenderBox:   goflow.NewBaseRenderBox(),
			controller:      p.Controller,
			scrollDirection: p.ScrollDirection,
			pageSnapping:    p.PageSnapping,
			onPageChanged:   p.OnPageChanged,
			children:        make([]goflow.RenderObject, 0),
		},
		children: make([]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (p *PageView) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// PageViewRenderObjectElement is a special element for PageView
type PageViewRenderObjectElement struct {
	goflow.BaseElement
	widget       *PageView
	renderObject *RenderPageView
	children     []goflow.Element
}

// GetWidgetType returns the widget type
func (e *PageViewRenderObjectElement) GetWidgetType() string {
	return "PageView"
}

// GetElementType returns the element type
func (e *PageViewRenderObjectElement) GetElementType() string {
	return "PageViewRenderObjectElement"
}

// GetRenderObject returns the render object
func (e *PageViewRenderObjectElement) GetRenderObject() goflow.RenderObject {
	return e.renderObject
}

// Mount mounts the element
func (e *PageViewRenderObjectElement) Mount(parent goflow.Element, slot interface{}) {
	e.SetParent(parent)

	// Mount children
	for _, childWidget := range e.widget.Children {
		childElement := childWidget.CreateElement()
		e.children = append(e.children, childElement)
		childElement.Mount(e, nil)

		if childRenderObj := childElement.GetRenderObject(); childRenderObj != nil {
			e.renderObject.children = append(e.renderObject.children, childRenderObj)
			childRenderObj.SetParent(e.renderObject)
		}
	}

	e.renderObject.MarkNeedsLayout()
}

// Update updates the element
func (e *PageViewRenderObjectElement) Update(newWidget goflow.Widget) {
	e.widget = newWidget.(*PageView)
	e.renderObject.MarkNeedsLayout()
	e.renderObject.MarkNeedsPaint()
}

// Unmount unmounts the element
func (e *PageViewRenderObjectElement) Unmount() {
	for _, child := range e.children {
		child.Unmount()
	}
	e.children = nil
}

// Rebuild rebuilds (no-op for render object elements)
func (e *PageViewRenderObjectElement) Rebuild() {
	// No-op
}

// VisitChildren visits child elements
func (e *PageViewRenderObjectElement) VisitChildren(visitor func(goflow.Element)) {
	for _, child := range e.children {
		visitor(child)
	}
}

// RenderPageView is the render object for PageView
type RenderPageView struct {
	*goflow.BaseRenderBox
	controller      *PageController
	scrollDirection Axis
	pageSnapping    bool
	onPageChanged   func(int)
	children        []goflow.RenderObject
	lastReportedPage int
}

// PerformLayout performs layout for page view
func (r *RenderPageView) PerformLayout() {
	constraints := r.GetConstraints()

	// Each page takes the full viewport size
	size := constraints.Biggest()
	r.SetSize(size)

	// Layout each child to fill the viewport
	for _, child := range r.children {
		childConstraints := goflow.TightConstraintsForSize(size)
		child.Layout(childConstraints)
	}

	// Position children based on current page
	r.positionChildren()
}

func (r *RenderPageView) positionChildren() {
	if len(r.children) == 0 {
		return
	}

	size := r.GetSize()
	currentPage := r.controller.PageFloat()

	// Position each child
	for i, child := range r.children {
		pageOffset := float64(i) - currentPage

		var x, y float64
		if r.scrollDirection == AxisHorizontal {
			x = pageOffset * size.Width
			y = 0
		} else {
			x = 0
			y = pageOffset * size.Height
		}

		child.SetOffset(goflow.NewOffset(x, y))
	}

	// Report page changes
	currentPageInt := r.controller.Page()
	if currentPageInt != r.lastReportedPage && r.onPageChanged != nil {
		r.lastReportedPage = currentPageInt
		r.onPageChanged(currentPageInt)
	}
}

// Layout performs the layout
func (r *RenderPageView) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the page view
func (r *RenderPageView) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	size := r.GetSize()

	// Save canvas state and clip to viewport
	canvas.Save()
	canvas.ClipRect(goflow.NewRect(offset, size))

	// Paint visible pages (current and adjacent)
	currentPage := r.controller.Page()
	startPage := int(math.Max(0, float64(currentPage-1)))
	endPage := int(math.Min(float64(len(r.children)-1), float64(currentPage+1)))

	for i := int(startPage); i <= int(endPage) && i < len(r.children); i++ {
		child := r.children[i]
		childOffset := offset.Add(child.GetOffset())
		child.Paint(canvas, childOffset)
	}

	// Restore canvas
	canvas.Restore()
}

// PageViewBuilder builds pages on demand
type PageViewBuilder struct {
	goflow.BaseWidget
	ItemBuilder     func(int) goflow.Widget
	ItemCount       int
	Controller      *PageController
	ScrollDirection Axis
	PageSnapping    bool
	OnPageChanged   func(int)
}

// NewPageViewBuilder creates a new page view builder
func NewPageViewBuilder(itemBuilder func(int) goflow.Widget, itemCount int) *PageViewBuilder {
	return &PageViewBuilder{
		ItemBuilder:     itemBuilder,
		ItemCount:       itemCount,
		Controller:      NewPageController(),
		ScrollDirection: AxisHorizontal,
		PageSnapping:    true,
	}
}

// Build creates the widget tree
func (p *PageViewBuilder) Build(context goflow.BuildContext) goflow.Widget {
	// Build visible pages
	currentPage := p.Controller.Page()
	children := make([]goflow.Widget, 0)

	// Build current and adjacent pages
	startPage := int(math.Max(0, float64(currentPage-1)))
	endPage := int(math.Min(float64(p.ItemCount-1), float64(currentPage+1)))
	for i := startPage; i <= endPage; i++ {
		if i >= 0 && i < p.ItemCount {
			children = append(children, p.ItemBuilder(i))
		}
	}

	return &PageView{
		Children:        children,
		Controller:      p.Controller,
		ScrollDirection: p.ScrollDirection,
		PageSnapping:    p.PageSnapping,
		OnPageChanged:   p.OnPageChanged,
	}
}
