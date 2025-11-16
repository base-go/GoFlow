package navigation

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Snackbar is a lightweight notification widget
type Snackbar struct {
	goflow.BaseWidget
	message   string
	action    *SnackbarAction
	onDismiss func()
}

// SnackbarAction represents an action button in a snackbar
type SnackbarAction struct {
	Label   string
	OnPress func()
}

// NewSnackbar creates a new snackbar
func NewSnackbar(message string, onDismiss func()) *Snackbar {
	return &Snackbar{
		message:   message,
		onDismiss: onDismiss,
	}
}

// WithAction adds an action button to the snackbar
func (s *Snackbar) WithAction(label string, onPress func()) *Snackbar {
	s.action = &SnackbarAction{
		Label:   label,
		OnPress: onPress,
	}
	return s
}

// Build builds the snackbar
func (s *Snackbar) Build(context goflow.BuildContext) goflow.Widget {
	children := []goflow.Widget{}

	// Add message text
	messageText := widgets.NewText(s.message)
	messageContainer := widgets.NewContainer()
	messageContainer.Padding = goflow.EdgeInsetsAll(16.0)
	messageContainer.Child = messageText

	children = append(children, messageContainer)

	// Add action button if present
	if s.action != nil {
		actionText := widgets.NewText(s.action.Label)
		actionContainer := widgets.NewContainer()
		actionContainer.Padding = goflow.EdgeInsetsSymmetric(8.0, 16.0)
		actionContainer.Child = actionText
		// TODO: Add GestureDetector when available
		// actionContainer = widgets.NewGestureDetector(actionContainer, ...)

		children = append(children, actionContainer)
	}

	// Create row with message and action
	row := widgets.NewRow(children...)

	// Create snackbar container
	container := widgets.NewContainer()
	container.Color = &goflow.Color{R: 50, G: 50, B: 50, A: 1.0} // Dark gray
	container.Child = row

	// Position at bottom with some margin
	positioned := widgets.NewPositioned(container,
		floatPtr(16.0),  // left
		nil,             // top
		floatPtr(16.0),  // right
		floatPtr(16.0),  // bottom
		nil,             // width
		nil,             // height
	)

	// Wrap in stack to position correctly
	return widgets.NewStack(positioned)
}

// CreateElement creates the element
func (s *Snackbar) CreateElement() goflow.Element {
	return goflow.NewGenericElement(s)
}

// ShowSnackbar is a helper function to show a snackbar
func ShowSnackbar(message string, duration ...int) {
	Get.Snackbar(message, duration...)
}

// ShowSnackbarWithAction shows a snackbar with an action button
func ShowSnackbarWithAction(message string, actionLabel string, onAction func(), duration ...int) {
	snackbar := NewSnackbar(message, func() {
		Get.CloseSnackbar()
	}).WithAction(actionLabel, onAction)

	Get.overlayStack.Update(func(stack []goflow.Widget) []goflow.Widget {
		return append(stack, snackbar)
	})

	// Auto-dismiss after duration
	if len(duration) > 0 {
		// TODO: Implement proper timer-based auto-dismiss
		_ = duration[0]
	}
}
