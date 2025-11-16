package widgets

import (
	"sort"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// DataTable displays data in a table format
type DataTable struct {
	goflow.BaseWidget
	Columns           []DataColumn
	Rows              []DataRow
	SortColumnIndex   *int
	SortAscending     bool
	OnSelectAll       func(bool)
	Headings          *TableHeadings
	DataRowHeight     float64
	HeadingRowHeight  float64
	HorizontalMargin  float64
	ColumnSpacing     float64
	ShowCheckboxColumn bool
	DividerThickness  float64
}

// DataColumn represents a column in a DataTable
type DataColumn struct {
	Label     goflow.Widget
	Tooltip   string
	Numeric   bool
	OnSort    func(int, bool)
}

// DataRow represents a row in a DataTable
type DataRow struct {
	Cells    []DataCell
	Selected bool
	OnSelectChanged func(bool)
	Color    *goflow.Color
}

// DataCell represents a cell in a DataTable
type DataCell struct {
	Child      goflow.Widget
	Placeholder bool
}

// TableHeadings represents custom table headings
type TableHeadings struct {
	HeadingRowColor *goflow.Color
	HeadingTextStyle *goflow.TextStyle
}

// NewDataTable creates a new data table
func NewDataTable(columns []DataColumn, rows []DataRow) *DataTable {
	return &DataTable{
		Columns:           columns,
		Rows:              rows,
		SortAscending:     true,
		DataRowHeight:     48.0,
		HeadingRowHeight:  56.0,
		HorizontalMargin:  24.0,
		ColumnSpacing:     56.0,
		ShowCheckboxColumn: false,
		DividerThickness:  1.0,
	}
}

// WithSorting enables sorting on a column
func (dt *DataTable) WithSorting(columnIndex int, ascending bool) *DataTable {
	dt.SortColumnIndex = &columnIndex
	dt.SortAscending = ascending
	return dt
}

// Build creates the widget tree
func (dt *DataTable) Build(context goflow.BuildContext) goflow.Widget {
	children := []goflow.Widget{
		// Header row
		dt.buildHeaderRow(),
	}

	// Data rows
	for i := range dt.Rows {
		children = append(children, dt.buildDataRow(i))
	}

	return &Container{
		Child: &Column{
			Children: children,
		},
	}
}

func (dt *DataTable) buildHeaderRow() goflow.Widget {
	cells := make([]goflow.Widget, len(dt.Columns))

	for i, column := range dt.Columns {
		isCurrent := dt.SortColumnIndex != nil && *dt.SortColumnIndex == i

		cell := &Container{
			Child: &Row{
				Children: []goflow.Widget{
					column.Label,
					dt.buildSortIndicator(isCurrent),
				},
			},
			Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
		}

		if column.OnSort != nil {
			columnIndex := i
			cell = &GestureDetector{
				OnTap: func() {
					newAscending := true
					if isCurrent {
						newAscending = !dt.SortAscending
					}
					dt.SortColumnIndex = &columnIndex
					dt.SortAscending = newAscending
					column.OnSort(columnIndex, newAscending)
				},
				Child: cell,
			}.(*GestureDetector)
		}

		cells[i] = cell
	}

	height := dt.HeadingRowHeight
	return &Container{
		Height: &height,
		Color:  goflow.NewColor(245, 245, 245, 255),
		Child: &Row{
			Children: cells,
		},
	}
}

func (dt *DataTable) buildSortIndicator(isCurrent bool) goflow.Widget {
	if !isCurrent {
		return &SizedBox{Width: 0, Height: 0}
	}

	iconName := "arrow_upward"
	if !dt.SortAscending {
		iconName = "arrow_downward"
	}

	return &Icon{
		Icon: iconName,
		Size: 16,
	}
}

func (dt *DataTable) buildDataRow(rowIndex int) goflow.Widget {
	row := dt.Rows[rowIndex]
	cells := make([]goflow.Widget, len(row.Cells))

	for i, cell := range row.Cells {
		cells[i] = &Container{
			Child:   cell.Child,
			Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
		}
	}

	height := dt.DataRowHeight
	bgColor := row.Color
	if bgColor == nil {
		bgColor = goflow.NewColor(255, 255, 255, 255)
	}

	return &Container{
		Height: &height,
		Color:  bgColor,
		Child: &Row{
			Children: cells,
		},
	}
}

// Card widget with elevation and rounded corners
type Card struct {
	goflow.BaseWidget
	Child            goflow.Widget
	Color            *goflow.Color
	ShadowColor      *goflow.Color
	Elevation        float64
	Shape            CardShape
	BorderOnForeground bool
	Margin           *goflow.EdgeInsets
	ClipBehavior     Clip
	Semantics        string
}

// CardShape defines the shape of a card
type CardShape struct {
	BorderRadius float64
	Border       *BorderSide
}

// Clip defines how to clip content
type Clip int

const (
	ClipNone Clip = iota
	ClipHardEdge
	ClipAntiAlias
	ClipAntiAliasWithSaveLayer
)

// NewCard creates a new card
func NewCard(child goflow.Widget) *Card {
	return &Card{
		Child:       child,
		Color:       goflow.NewColor(255, 255, 255, 255),
		ShadowColor: goflow.NewColor(0, 0, 0, 40),
		Elevation:   1.0,
		Shape: CardShape{
			BorderRadius: 4.0,
		},
		BorderOnForeground: true,
		ClipBehavior:      ClipNone,
	}
}

// WithElevation sets the elevation
func (c *Card) WithElevation(elevation float64) *Card {
	c.Elevation = elevation
	return c
}

// WithMargin sets the margin
func (c *Card) WithMargin(margin *goflow.EdgeInsets) *Card {
	c.Margin = margin
	return c
}

// Build creates the widget tree
func (c *Card) Build(context goflow.BuildContext) goflow.Widget {
	// Create container with shadow effect (simulated with elevation)
	return &Container{
		Color:  c.Color,
		Child:  c.Child,
		Margin: c.Margin,
		// In a real implementation, elevation would create a shadow
		// For now, we'll just use the container
	}
}

// ExpansionPanel is a panel that can be expanded or collapsed
type ExpansionPanel struct {
	goflow.BaseWidget
	HeaderBuilder func(goflow.BuildContext, bool) goflow.Widget
	Body          goflow.Widget
	IsExpanded    bool
	CanTapOnHeader bool
}

// NewExpansionPanel creates a new expansion panel
func NewExpansionPanel(
	headerBuilder func(goflow.BuildContext, bool) goflow.Widget,
	body goflow.Widget,
) *ExpansionPanel {
	return &ExpansionPanel{
		HeaderBuilder: headerBuilder,
		Body:          body,
		IsExpanded:    false,
		CanTapOnHeader: true,
	}
}

// Build creates the widget tree
func (ep *ExpansionPanel) Build(context goflow.BuildContext) goflow.Widget {
	children := []goflow.Widget{
		// Header
		ep.HeaderBuilder(context, ep.IsExpanded),
	}

	// Add body if expanded
	if ep.IsExpanded {
		children = append(children, ep.Body)
	}

	return &Column{
		Children: children,
	}
}

// ExpansionPanelList is a list of expansion panels
type ExpansionPanelList struct {
	goflow.BaseWidget
	Children              []ExpansionPanel
	ExpansionCallback     func(int, bool)
	AnimationDuration     int64
	ElevationCallback     func(int) float64
	ExpandedHeaderPadding *goflow.EdgeInsets
}

// NewExpansionPanelList creates a new expansion panel list
func NewExpansionPanelList(panels []ExpansionPanel) *ExpansionPanelList {
	return &ExpansionPanelList{
		Children:          panels,
		AnimationDuration: 200 * 1000000, // 200ms
	}
}

// Build creates the widget tree
func (epl *ExpansionPanelList) Build(context goflow.BuildContext) goflow.Widget {
	children := make([]goflow.Widget, len(epl.Children))

	for i := range epl.Children {
		children[i] = &epl.Children[i]
	}

	return &Column{
		Children: children,
	}
}

// ExpansionTile is a single-line ListTile with a trailing button that expands or collapses
type ExpansionTile struct {
	goflow.BaseWidget
	Title             goflow.Widget
	Subtitle          goflow.Widget
	Leading           goflow.Widget
	Trailing          goflow.Widget
	Children          []goflow.Widget
	InitiallyExpanded bool
	OnExpansionChanged func(bool)
	TilePadding       *goflow.EdgeInsets
	ExpandedCrossAxisAlignment CrossAxisAlignment
	ExpandedAlignment Alignment
	ChildrenPadding   *goflow.EdgeInsets
	BackgroundColor   *goflow.Color
	CollapsedBackgroundColor *goflow.Color
	TextColor         *goflow.Color
	CollapsedTextColor *goflow.Color
	IconColor         *goflow.Color
	CollapsedIconColor *goflow.Color
	isExpanded        bool
}

// NewExpansionTile creates a new expansion tile
func NewExpansionTile(title goflow.Widget, children []goflow.Widget) *ExpansionTile {
	return &ExpansionTile{
		Title:             title,
		Children:          children,
		InitiallyExpanded: false,
		isExpanded:        false,
		ExpandedCrossAxisAlignment: CrossAxisAlignmentCenter,
	}
}

// Build creates the widget tree
func (et *ExpansionTile) Build(context goflow.BuildContext) goflow.Widget {
	headerChildren := []goflow.Widget{}

	if et.Leading != nil {
		headerChildren = append(headerChildren, et.Leading)
	}

	headerChildren = append(headerChildren, &Column{
		CrossAxisAlignment: CrossAxisAlignmentStart,
		Children: []goflow.Widget{
			et.Title,
			et.Subtitle,
		},
	})

	// Trailing icon (expand/collapse indicator)
	iconName := "expand_more"
	if et.isExpanded {
		iconName = "expand_less"
	}

	headerChildren = append(headerChildren, &Icon{
		Icon: iconName,
		Size: 24,
	})

	header := &GestureDetector{
		OnTap: func() {
			et.isExpanded = !et.isExpanded
			if et.OnExpansionChanged != nil {
				et.OnExpansionChanged(et.isExpanded)
			}
		},
		Child: &Container{
			Color: et.BackgroundColor,
			Child: &Row{
				Children: headerChildren,
			},
			Padding: et.TilePadding,
		},
	}

	widgets := []goflow.Widget{header}

	// Add children if expanded
	if et.isExpanded {
		childrenContainer := &Container{
			Child: &Column{
				Children: et.Children,
			},
			Padding: et.ChildrenPadding,
		}
		widgets = append(widgets, childrenContainer)
	}

	return &Column{
		Children: widgets,
	}
}

// TreeView displays hierarchical data
type TreeView struct {
	goflow.BaseWidget
	Nodes             []TreeNode
	OnNodeTap         func(string)
	OnNodeDoubleTap   func(string)
	OnNodeExpanded    func(string, bool)
	IndentPerLevel    float64
	ShowRootNode      bool
	AllowSelection    bool
	ExpanderBuilder   func(bool) goflow.Widget
}

// TreeNode represents a node in a tree
type TreeNode struct {
	ID          string
	Label       goflow.Widget
	Icon        goflow.Widget
	Children    []TreeNode
	IsExpanded  bool
	IsSelected  bool
	Data        interface{}
}

// NewTreeView creates a new tree view
func NewTreeView(nodes []TreeNode) *TreeView {
	return &TreeView{
		Nodes:          nodes,
		IndentPerLevel: 24.0,
		ShowRootNode:   true,
		AllowSelection: true,
	}
}

// Build creates the widget tree
func (tv *TreeView) Build(context goflow.BuildContext) goflow.Widget {
	return &Column{
		Children: tv.buildNodes(tv.Nodes, 0),
	}
}

func (tv *TreeView) buildNodes(nodes []TreeNode, level int) []goflow.Widget {
	widgets := make([]goflow.Widget, 0)

	for _, node := range nodes {
		// Build node
		widgets = append(widgets, tv.buildNode(node, level))

		// Build children if expanded
		if node.IsExpanded && len(node.Children) > 0 {
			childWidgets := tv.buildNodes(node.Children, level+1)
			widgets = append(widgets, childWidgets...)
		}
	}

	return widgets
}

func (tv *TreeView) buildNode(node TreeNode, level int) goflow.Widget {
	indent := tv.IndentPerLevel * float64(level)

	// Expander icon
	expanderIcon := tv.buildExpander(node)

	children := []goflow.Widget{
		&SizedBox{Width: indent},
		expanderIcon,
	}

	if node.Icon != nil {
		children = append(children, node.Icon)
	}

	children = append(children, node.Label)

	return &GestureDetector{
		OnTap: func() {
			if tv.OnNodeTap != nil {
				tv.OnNodeTap(node.ID)
			}
		},
		OnDoubleTap: func() {
			if tv.OnNodeDoubleTap != nil {
				tv.OnNodeDoubleTap(node.ID)
			}
		},
		Child: &Container{
			Child: &Row{
				Children: children,
			},
			Padding: goflow.NewEdgeInsets(4, 4, 4, 4),
		},
	}
}

func (tv *TreeView) buildExpander(node TreeNode) goflow.Widget {
	if len(node.Children) == 0 {
		return &SizedBox{Width: 16, Height: 16}
	}

	if tv.ExpanderBuilder != nil {
		return tv.ExpanderBuilder(node.IsExpanded)
	}

	iconName := "chevron_right"
	if node.IsExpanded {
		iconName = "expand_more"
	}

	return &Icon{
		Icon: iconName,
		Size: 16,
	}
}

// DataTableSource provides data for paginated data tables
type DataTableSource interface {
	GetRowCount() int
	GetRow(index int) DataRow
	OnRowsPerPageChanged(rowsPerPage int)
	OnPageChanged(page int)
}

// PaginatedDataTable is a data table with pagination
type PaginatedDataTable struct {
	goflow.BaseWidget
	Source           DataTableSource
	Header           goflow.Widget
	Columns          []DataColumn
	RowsPerPage      int
	AvailableRowsPerPage []int
	OnRowsPerPageChanged func(int)
	OnPageChanged    func(int)
	CurrentPage      int
	SortColumnIndex  *int
	SortAscending    bool
}

// NewPaginatedDataTable creates a new paginated data table
func NewPaginatedDataTable(source DataTableSource, columns []DataColumn) *PaginatedDataTable {
	return &PaginatedDataTable{
		Source:               source,
		Columns:              columns,
		RowsPerPage:          10,
		AvailableRowsPerPage: []int{5, 10, 20, 50},
		CurrentPage:          0,
		SortAscending:        true,
	}
}

// Build creates the widget tree
func (pdt *PaginatedDataTable) Build(context goflow.BuildContext) goflow.Widget {
	// Calculate rows for current page
	startRow := pdt.CurrentPage * pdt.RowsPerPage
	endRow := startRow + pdt.RowsPerPage
	totalRows := pdt.Source.GetRowCount()

	if endRow > totalRows {
		endRow = totalRows
	}

	rows := make([]DataRow, 0)
	for i := startRow; i < endRow; i++ {
		rows = append(rows, pdt.Source.GetRow(i))
	}

	dataTable := NewDataTable(pdt.Columns, rows)
	if pdt.SortColumnIndex != nil {
		dataTable.WithSorting(*pdt.SortColumnIndex, pdt.SortAscending)
	}

	return &Column{
		Children: []goflow.Widget{
			pdt.Header,
			dataTable,
			pdt.buildPagination(totalRows, startRow, endRow),
		},
	}
}

func (pdt *PaginatedDataTable) buildPagination(totalRows, startRow, endRow int) goflow.Widget {
	return &Row{
		MainAxisAlignment: MainAxisAlignmentEnd,
		Children: []goflow.Widget{
			&Text{
				Data:  fmt.Sprintf("%d-%d of %d", startRow+1, endRow, totalRows),
				Style: goflow.NewTextStyle(),
			},
		},
	}
}

// SortableDataTable adds sorting functionality to DataTable
type SortableDataTable struct {
	*DataTable
	sortComparator func(DataRow, DataRow) bool
}

// NewSortableDataTable creates a new sortable data table
func NewSortableDataTable(columns []DataColumn, rows []DataRow) *SortableDataTable {
	return &SortableDataTable{
		DataTable: NewDataTable(columns, rows),
	}
}

// Sort sorts the table by the given column
func (sdt *SortableDataTable) Sort(columnIndex int, ascending bool) {
	if sdt.sortComparator == nil {
		return
	}

	sort.SliceStable(sdt.Rows, func(i, j int) bool {
		result := sdt.sortComparator(sdt.Rows[i], sdt.Rows[j])
		if !ascending {
			result = !result
		}
		return result
	})
}
