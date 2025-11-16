package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// DialogWrapper wraps a dialog with barrier and dismiss handling
type DialogWrapper struct {
	goflow.BaseWidget
	child              goflow.Widget
	barrierDismissible bool
	onDismiss          func()
}

// NewDialogWrapper creates a new dialog wrapper
func NewDialogWrapper(child goflow.Widget, barrierDismissible bool, onDismiss func()) *DialogWrapper {
	return &DialogWrapper{
		child:              child,
		barrierDismissible: barrierDismissible,
		onDismiss:          onDismiss,
	}
}

// Build builds the dialog with barrier
func (d *DialogWrapper) Build(context goflow.BuildContext) goflow.Widget {
	// Create a stack with:
	// 1. Semi-transparent barrier
	// 2. Centered dialog content

	barrier := widgets.NewContainer()
	barrier.Color = &goflow.Color{R: 0, G: 0, B: 0, A: 0.5} // Semi-transparent black

	// If dismissible, add gesture detector to barrier
	if d.barrierDismissible {
		barrier = widgets.NewContainer()
		barrier.Color = &goflow.Color{R: 0, G: 0, B: 0, A: 0.5}
		// TODO: Add GestureDetector when available
		// barrier = widgets.NewGestureDetector(barrier, widgets.GestureCallbacks{
		// 	OnTap: d.onDismiss,
		// })
	}

	// Center the dialog content
	dialogContent := widgets.NewCenter(d.child)

	// Stack barrier and dialog
	return widgets.NewStack(barrier, dialogContent)
}

// CreateElement creates the element
func (d *DialogWrapper) CreateElement() goflow.Element {
	return goflow.NewGenericElement(d)
}

// AlertDialog is a Material Design alert dialog
type AlertDialog struct {
	goflow.BaseWidget
	Title   goflow.Widget
	Content goflow.Widget
	Actions []DialogAction
}

// DialogAction represents a dialog action button
type DialogAction struct {
	Label   string
	OnPress func()
	Primary bool
}

// NewAlertDialog creates a new alert dialog
func NewAlertDialog(title, content goflow.Widget, actions []DialogAction) *AlertDialog {
	return &AlertDialog{
		Title:   title,
		Content: content,
		Actions: actions,
	}
}

// Build builds the alert dialog
func (a *AlertDialog) Build(context goflow.BuildContext) goflow.Widget {
	// Create dialog structure
	children := []goflow.Widget{}

	// Add title if present
	if a.Title != nil {
		titleContainer := widgets.NewContainer()
		titleContainer.Padding = goflow.EdgeInsetsAll(16.0)
		titleContainer.Child = a.Title
		children = append(children, titleContainer)
	}

	// Add content if present
	if a.Content != nil {
		contentContainer := widgets.NewContainer()
		contentContainer.Padding = goflow.EdgeInsetsAll(16.0)
		contentContainer.Child = a.Content
		children = append(children, contentContainer)
	}

	// Add actions
	if len(a.Actions) > 0 {
		actionButtons := []goflow.Widget{}
		for _, action := range a.Actions {
			// TODO: Create actual button widgets when available
			// For now, just create text widgets
			button := widgets.NewText(action.Label)
			actionButtons = append(actionButtons, button)
		}

		actionsRow := widgets.NewRow(actionButtons...)
		actionsContainer := widgets.NewContainer()
		actionsContainer.Padding = goflow.EdgeInsetsAll(8.0)
		actionsContainer.Child = actionsRow

		children = append(children, actionsContainer)
	}

	// Create dialog card
	dialog := widgets.NewContainer()
	dialog.Width = floatPtr(300.0)
	dialog.Color = &goflow.Color{R: 255, G: 255, B: 255, A: 1.0} // White background
	dialog.Padding = goflow.EdgeInsetsAll(0)
	dialog.Child = widgets.NewColumn(children...)

	return dialog
}

// CreateElement creates the element
func (a *AlertDialog) CreateElement() goflow.Element {
	return goflow.NewGenericElement(a)
}

// ShowDialog is a helper function to show a dialog
func ShowDialog(dialog goflow.Widget, barrierDismissible ...bool) {
	Get.Dialog(dialog, barrierDismissible...)
}

// ShowAlertDialog is a helper to show an alert dialog
func ShowAlertDialog(title string, content string, actions []DialogAction) {
	titleWidget := widgets.NewText(title)
	contentWidget := widgets.NewText(content)

	dialog := NewAlertDialog(titleWidget, contentWidget, actions)
	Get.Dialog(dialog, true)
}

// floatPtr returns a pointer to a float64
func floatPtr(f float64) *float64 {
	return &f
}
