package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// BottomSheetWrapper wraps a bottom sheet with barrier and dismiss handling
type BottomSheetWrapper struct {
	goflow.BaseWidget
	child         goflow.Widget
	isDismissible bool
	onDismiss     func()
}

// NewBottomSheetWrapper creates a new bottom sheet wrapper
func NewBottomSheetWrapper(child goflow.Widget, isDismissible bool, onDismiss func()) *BottomSheetWrapper {
	return &BottomSheetWrapper{
		child:         child,
		isDismissible: isDismissible,
		onDismiss:     onDismiss,
	}
}

// Build builds the bottom sheet with barrier
func (b *BottomSheetWrapper) Build(context goflow.BuildContext) goflow.Widget {
	// Create a stack with:
	// 1. Semi-transparent barrier
	// 2. Bottom-aligned sheet content

	barrier := widgets.NewContainer()
	barrier.Color = &goflow.Color{R: 0, G: 0, B: 0, A: 0.5} // Semi-transparent black

	// If dismissible, add gesture detector to barrier
	if b.isDismissible {
		// TODO: Add GestureDetector when available
		// barrier = widgets.NewGestureDetector(barrier, widgets.GestureCallbacks{
		// 	OnTap: b.onDismiss,
		// })
	}

	// Create bottom sheet container
	sheetContainer := widgets.NewContainer()
	sheetContainer.Color = &goflow.Color{R: 255, G: 255, B: 255, A: 1.0} // White background
	sheetContainer.Padding = goflow.EdgeInsetsAll(16.0)
	sheetContainer.Child = b.child

	// Align to bottom using Positioned widget
	sheet := widgets.NewPositioned(sheetContainer,
		nil,                      // left
		nil,                      // top
		nil,                      // right
		floatPtr(0.0),            // bottom
		nil,                      // width
		nil,                      // height
	)

	// Stack barrier and bottom sheet
	return widgets.NewStack(barrier, sheet)
}

// CreateElement creates the element
func (b *BottomSheetWrapper) CreateElement() goflow.Element {
	return goflow.NewGenericElement(b)
}

// ModalBottomSheet is a modal bottom sheet widget
type ModalBottomSheet struct {
	goflow.BaseWidget
	child         goflow.Widget
	title         string
	showHandle    bool
}

// NewModalBottomSheet creates a new modal bottom sheet
func NewModalBottomSheet(child goflow.Widget) *ModalBottomSheet {
	return &ModalBottomSheet{
		child:      child,
		showHandle: true,
	}
}

// WithTitle sets the title
func (m *ModalBottomSheet) WithTitle(title string) *ModalBottomSheet {
	m.title = title
	return m
}

// WithHandle controls whether to show the drag handle
func (m *ModalBottomSheet) WithHandle(show bool) *ModalBottomSheet {
	m.showHandle = show
	return m
}

// Build builds the modal bottom sheet
func (m *ModalBottomSheet) Build(context goflow.BuildContext) goflow.Widget {
	children := []goflow.Widget{}

	// Add drag handle if enabled
	if m.showHandle {
		handle := widgets.NewContainer()
		handle.Width = floatPtr(40.0)
		handle.Height = floatPtr(4.0)
		handle.Color = &goflow.Color{R: 200, G: 200, B: 200, A: 1.0} // Light gray

		handleContainer := widgets.NewContainer()
		handleContainer.Padding = goflow.EdgeInsetsSymmetric(8.0, 0)
		handleContainer.Child = widgets.NewCenter(handle)

		children = append(children, handleContainer)
	}

	// Add title if present
	if m.title != "" {
		titleText := widgets.NewText(m.title)
		titleContainer := widgets.NewContainer()
		titleContainer.Padding = goflow.EdgeInsetsAll(16.0)
		titleContainer.Child = titleText

		children = append(children, titleContainer)
	}

	// Add content
	contentContainer := widgets.NewContainer()
	contentContainer.Padding = goflow.EdgeInsetsAll(16.0)
	contentContainer.Child = m.child
	children = append(children, contentContainer)

	// Create column with all children
	column := widgets.NewColumn(children...)

	// Wrap in container with rounded top corners
	container := widgets.NewContainer()
	container.Color = &goflow.Color{R: 255, G: 255, B: 255, A: 1.0}
	container.Child = column

	return container
}

// CreateElement creates the element
func (m *ModalBottomSheet) CreateElement() goflow.Element {
	return goflow.NewGenericElement(m)
}

// ShowBottomSheet is a helper function to show a bottom sheet
func ShowBottomSheet(sheet goflow.Widget, isDismissible ...bool) {
	Get.BottomSheet(sheet, isDismissible...)
}

// ShowModalBottomSheet is a helper to show a modal bottom sheet
func ShowModalBottomSheet(content goflow.Widget, title ...string) {
	sheet := NewModalBottomSheet(content)

	if len(title) > 0 && title[0] != "" {
		sheet.WithTitle(title[0])
	}

	Get.BottomSheet(sheet, true)
}
