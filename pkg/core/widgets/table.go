package widgets

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// Table displays children in a table layout
type Table struct {
	goflow.BaseWidget
	Children         []TableRow
	ColumnWidths     map[int]TableColumnWidth
	DefaultRowHeight float64
	Border           *TableBorder
	DefaultVerticalAlignment TableCellVerticalAlignment
}

// TableRow represents a row in a table
type TableRow struct {
	Children   []goflow.Widget
	Decoration *goflow.Color
}

// TableColumnWidth defines how a column should be sized
type TableColumnWidth interface {
	GetWidth(constraints *goflow.Constraints, rows [][]goflow.RenderObject, columnIndex int) float64
}

// FixedColumnWidth uses a fixed width
type FixedColumnWidth struct {
	Width float64
}

func (f *FixedColumnWidth) GetWidth(constraints *goflow.Constraints, rows [][]goflow.RenderObject, columnIndex int) float64 {
	return f.Width
}

// FlexColumnWidth uses a flex factor
type FlexColumnWidth struct {
	Flex float64
}

func (f *FlexColumnWidth) GetWidth(constraints *goflow.Constraints, rows [][]goflow.RenderObject, columnIndex int) float64 {
	// This will be calculated during layout
	return 0
}

// IntrinsicColumnWidth sizes based on content
type IntrinsicColumnWidth struct{}

func (i *IntrinsicColumnWidth) GetWidth(constraints *goflow.Constraints, rows [][]goflow.RenderObject, columnIndex int) float64 {
	maxWidth := 0.0
	for _, row := range rows {
		if columnIndex < len(row) {
			childSize := row[columnIndex].GetSize()
			if childSize.Width > maxWidth {
				maxWidth = childSize.Width
			}
		}
	}
	return maxWidth
}

// TableBorder defines borders for a table
type TableBorder struct {
	Top          *BorderSide
	Right        *BorderSide
	Bottom       *BorderSide
	Left         *BorderSide
	Horizontal   *BorderSide
	Vertical     *BorderSide
	BorderRadius float64
}

// TableCellVerticalAlignment defines vertical alignment in a table cell
type TableCellVerticalAlignment int

const (
	TableCellVerticalAlignmentTop TableCellVerticalAlignment = iota
	TableCellVerticalAlignmentMiddle
	TableCellVerticalAlignmentBottom
)

// NewTable creates a new Table widget
func NewTable(children []TableRow) *Table {
	return &Table{
		Children:                 children,
		ColumnWidths:             make(map[int]TableColumnWidth),
		DefaultRowHeight:         0, // Auto-size
		DefaultVerticalAlignment: TableCellVerticalAlignmentTop,
	}
}

// CreateElement creates a render object element for Table
func (t *Table) CreateElement() goflow.Element {
	return &TableRenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      t,
		renderObject: &RenderTable{
			BaseRenderBox:        goflow.NewBaseRenderBox(),
			rows:                 make([][]goflow.RenderObject, 0),
			columnWidths:         t.ColumnWidths,
			defaultRowHeight:     t.DefaultRowHeight,
			border:               t.Border,
			verticalAlignment:    t.DefaultVerticalAlignment,
		},
		children: make([][]goflow.Element, 0),
	}
}

// Build returns nil (this is a render object widget)
func (t *Table) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// TableRenderObjectElement is a special element for table
type TableRenderObjectElement struct {
	goflow.BaseElement
	widget       *Table
	renderObject *RenderTable
	children     [][]goflow.Element
}

// GetWidgetType returns the widget type
func (e *TableRenderObjectElement) GetWidgetType() string {
	return "Table"
}

// GetElementType returns the element type
func (e *TableRenderObjectElement) GetElementType() string {
	return "TableRenderObjectElement"
}

// GetRenderObject returns the render object
func (e *TableRenderObjectElement) GetRenderObject() goflow.RenderObject {
	return e.renderObject
}

// Mount mounts the element
func (e *TableRenderObjectElement) Mount(parent goflow.Element, slot interface{}) {
	e.SetParent(parent)

	// Mount children
	for _, row := range e.widget.Children {
		rowElements := make([]goflow.Element, 0)
		rowRenderObjects := make([]goflow.RenderObject, 0)

		for _, cellWidget := range row.Children {
			cellElement := cellWidget.CreateElement()
			rowElements = append(rowElements, cellElement)
			cellElement.Mount(e, nil)

			if childRenderObj := cellElement.GetRenderObject(); childRenderObj != nil {
				rowRenderObjects = append(rowRenderObjects, childRenderObj)
			}
		}

		e.children = append(e.children, rowElements)
		e.renderObject.rows = append(e.renderObject.rows, rowRenderObjects)
	}

	e.renderObject.MarkNeedsLayout()
}

// Update updates the element
func (e *TableRenderObjectElement) Update(newWidget goflow.Widget) {
	e.widget = newWidget.(*Table)
	e.renderObject.MarkNeedsLayout()
	e.renderObject.MarkNeedsPaint()
}

// Unmount unmounts the element
func (e *TableRenderObjectElement) Unmount() {
	for _, rowElements := range e.children {
		for _, child := range rowElements {
			child.Unmount()
		}
	}
	e.children = nil
}

// Rebuild rebuilds (no-op for render object elements)
func (e *TableRenderObjectElement) Rebuild() {
	// No-op
}

// VisitChildren visits child elements
func (e *TableRenderObjectElement) VisitChildren(visitor func(goflow.Element)) {
	for _, rowElements := range e.children {
		for _, child := range rowElements {
			visitor(child)
		}
	}
}

// RenderTable is the render object for Table
type RenderTable struct {
	*goflow.BaseRenderBox
	rows              [][]goflow.RenderObject
	columnWidths      map[int]TableColumnWidth
	defaultRowHeight  float64
	border            *TableBorder
	verticalAlignment TableCellVerticalAlignment
}

// PerformLayout performs layout for table
func (r *RenderTable) PerformLayout() {
	constraints := r.GetConstraints()

	if len(r.rows) == 0 {
		r.SetSize(constraints.Smallest())
		return
	}

	// Determine number of columns
	numColumns := 0
	for _, row := range r.rows {
		if len(row) > numColumns {
			numColumns = len(row)
		}
	}

	// First pass: layout all cells to get intrinsic sizes
	for _, row := range r.rows {
		for _, cell := range row {
			childConstraints := goflow.LooseConstraints(constraints.MaxWidth, constraints.MaxHeight)
			cell.Layout(childConstraints)
		}
	}

	// Calculate column widths
	columnWidths := make([]float64, numColumns)
	flexColumns := make(map[int]float64)
	totalFixedWidth := 0.0

	for col := 0; col < numColumns; col++ {
		if colWidth, ok := r.columnWidths[col]; ok {
			switch w := colWidth.(type) {
			case *FixedColumnWidth:
				columnWidths[col] = w.Width
				totalFixedWidth += w.Width
			case *FlexColumnWidth:
				flexColumns[col] = w.Flex
			case *IntrinsicColumnWidth:
				width := w.GetWidth(constraints, r.rows, col)
				columnWidths[col] = width
				totalFixedWidth += width
			}
		} else {
			// Default to intrinsic
			width := (&IntrinsicColumnWidth{}).GetWidth(constraints, r.rows, col)
			columnWidths[col] = width
			totalFixedWidth += width
		}
	}

	// Distribute remaining space to flex columns
	if len(flexColumns) > 0 {
		remainingWidth := constraints.MaxWidth - totalFixedWidth
		totalFlex := 0.0
		for _, flex := range flexColumns {
			totalFlex += flex
		}

		if totalFlex > 0 {
			for col, flex := range flexColumns {
				columnWidths[col] = (flex / totalFlex) * remainingWidth
			}
		}
	}

	// Calculate row heights
	rowHeights := make([]float64, len(r.rows))
	for rowIdx, row := range r.rows {
		maxHeight := r.defaultRowHeight
		for _, cell := range row {
			cellHeight := cell.GetSize().Height
			if cellHeight > maxHeight {
				maxHeight = cellHeight
			}
		}
		rowHeights[rowIdx] = maxHeight
	}

	// Calculate total table size
	totalWidth := 0.0
	for _, width := range columnWidths {
		totalWidth += width
	}

	totalHeight := 0.0
	for _, height := range rowHeights {
		totalHeight += height
	}

	r.SetSize(goflow.NewSize(totalWidth, totalHeight))

	// Position cells
	currentY := 0.0
	for rowIdx, row := range r.rows {
		currentX := 0.0
		rowHeight := rowHeights[rowIdx]

		for colIdx, cell := range row {
			if colIdx >= numColumns {
				break
			}

			cellSize := cell.GetSize()
			columnWidth := columnWidths[colIdx]

			// Calculate vertical alignment
			y := currentY
			switch r.verticalAlignment {
			case TableCellVerticalAlignmentTop:
				y = currentY
			case TableCellVerticalAlignmentMiddle:
				y = currentY + (rowHeight-cellSize.Height)/2
			case TableCellVerticalAlignmentBottom:
				y = currentY + rowHeight - cellSize.Height
			}

			// Center horizontally in cell
			x := currentX + (columnWidth-cellSize.Width)/2

			cell.SetOffset(goflow.NewOffset(x, y))
			currentX += columnWidth
		}

		currentY += rowHeight
	}
}

// Layout performs the layout
func (r *RenderTable) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the table
func (r *RenderTable) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// Paint cells
	for _, row := range r.rows {
		for _, cell := range row {
			childOffset := offset.Add(cell.GetOffset())
			cell.Paint(canvas, childOffset)
		}
	}

	// Paint borders if specified
	if r.border != nil {
		r.paintBorders(canvas, offset)
	}
}

func (r *RenderTable) paintBorders(canvas goflow.Canvas, offset *goflow.Offset) {
	size := r.GetSize()
	paint := goflow.NewPaint()

	// Paint outer borders
	if r.border.Top != nil {
		paint.Color = r.border.Top.Color
		paint.StrokeWidth = r.border.Top.Width
		canvas.DrawLine(
			offset,
			goflow.NewOffset(offset.X+size.Width, offset.Y),
			paint,
		)
	}

	if r.border.Bottom != nil {
		paint.Color = r.border.Bottom.Color
		paint.StrokeWidth = r.border.Bottom.Width
		canvas.DrawLine(
			goflow.NewOffset(offset.X, offset.Y+size.Height),
			goflow.NewOffset(offset.X+size.Width, offset.Y+size.Height),
			paint,
		)
	}

	if r.border.Left != nil {
		paint.Color = r.border.Left.Color
		paint.StrokeWidth = r.border.Left.Width
		canvas.DrawLine(
			offset,
			goflow.NewOffset(offset.X, offset.Y+size.Height),
			paint,
		)
	}

	if r.border.Right != nil {
		paint.Color = r.border.Right.Color
		paint.StrokeWidth = r.border.Right.Width
		canvas.DrawLine(
			goflow.NewOffset(offset.X+size.Width, offset.Y),
			goflow.NewOffset(offset.X+size.Width, offset.Y+size.Height),
			paint,
		)
	}
}
