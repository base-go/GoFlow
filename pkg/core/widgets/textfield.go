package widgets

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/input"
	"strings"
)

// TextEditingController manages the text being edited
type TextEditingController struct {
	text      string
	selection TextSelection
	listeners []func()
}

// NewTextEditingController creates a new text editing controller
func NewTextEditingController(initialText string) *TextEditingController {
	return &TextEditingController{
		text: initialText,
		selection: TextSelection{
			BaseOffset:   0,
			ExtentOffset: 0,
		},
		listeners: make([]func(), 0),
	}
}

// Text returns the current text
func (c *TextEditingController) Text() string {
	return c.text
}

// SetText sets the text and notifies listeners
func (c *TextEditingController) SetText(text string) {
	if c.text != text {
		c.text = text
		c.notifyListeners()
	}
}

// Selection returns the current selection
func (c *TextEditingController) Selection() TextSelection {
	return c.selection
}

// SetSelection sets the selection and notifies listeners
func (c *TextEditingController) SetSelection(selection TextSelection) {
	c.selection = selection
	c.notifyListeners()
}

// Clear clears the text
func (c *TextEditingController) Clear() {
	c.SetText("")
	c.SetSelection(TextSelection{BaseOffset: 0, ExtentOffset: 0})
}

// AddListener adds a change listener
func (c *TextEditingController) AddListener(listener func()) {
	c.listeners = append(c.listeners, listener)
}

// RemoveListener removes a change listener
func (c *TextEditingController) RemoveListener(listener func()) {
	// Note: In production, you'd use a unique ID for each listener
	// This is simplified for now
}

// notifyListeners notifies all listeners of a change
func (c *TextEditingController) notifyListeners() {
	for _, listener := range c.listeners {
		listener()
	}
}

// TextSelection represents a text selection
type TextSelection struct {
	BaseOffset   int // Start of selection
	ExtentOffset int // End of selection (cursor position)
}

// IsCollapsed returns true if the selection is collapsed (just a cursor)
func (s TextSelection) IsCollapsed() bool {
	return s.BaseOffset == s.ExtentOffset
}

// Start returns the start of the selection
func (s TextSelection) Start() int {
	if s.BaseOffset < s.ExtentOffset {
		return s.BaseOffset
	}
	return s.ExtentOffset
}

// End returns the end of the selection
func (s TextSelection) End() int {
	if s.BaseOffset > s.ExtentOffset {
		return s.BaseOffset
	}
	return s.ExtentOffset
}

// TextField is a single-line text input widget
type TextField struct {
	// Controller for managing text
	Controller *TextEditingController

	// Placeholder text when empty
	Placeholder string

	// Text style
	Style *goflow.TextStyle

	// Whether to obscure text (for passwords)
	ObscureText bool

	// Whether the field is enabled
	Enabled bool

	// Maximum length (0 = unlimited)
	MaxLength int

	// Callback when text changes
	OnChanged func(string)

	// Callback when editing is complete (Enter pressed)
	OnSubmitted func(string)

	// Focus node
	FocusNode *input.FocusNode

	// Autofocus
	Autofocus bool

	// Decoration (border, padding, etc.)
	Decoration *InputDecoration
}

// InputDecoration defines the decoration for an input field
type InputDecoration struct {
	// Border color
	BorderColor *goflow.Color

	// Border width
	BorderWidth float64

	// Border radius
	BorderRadius float64

	// Background color
	FillColor *goflow.Color

	// Padding
	ContentPadding *goflow.EdgeInsets

	// Hint text
	HintText string

	// Hint style
	HintStyle *goflow.TextStyle

	// Prefix icon
	PrefixIcon goflow.Widget

	// Suffix icon
	SuffixIcon goflow.Widget
}

// NewTextField creates a new TextField widget
func NewTextField(options ...TextFieldOption) *TextField {
	tf := &TextField{
		Controller: NewTextEditingController(""),
		Enabled:    true,
		Style:      goflow.NewTextStyle(),
		Decoration: &InputDecoration{
			BorderColor:    goflow.NewColor(200, 200, 200, 255),
			BorderWidth:    1.0,
			BorderRadius:   4.0,
			ContentPadding: goflow.NewEdgeInsets(8, 12, 8, 12),
		},
	}

	for _, opt := range options {
		opt(tf)
	}

	return tf
}

// TextFieldOption configures a TextField
type TextFieldOption func(*TextField)

// WithController sets the text editing controller
func WithController(controller *TextEditingController) TextFieldOption {
	return func(tf *TextField) {
		tf.Controller = controller
	}
}

// WithPlaceholder sets the placeholder text
func WithPlaceholder(text string) TextFieldOption {
	return func(tf *TextField) {
		tf.Placeholder = text
	}
}

// WithTextStyle sets the text style
func WithTextStyle(style *goflow.TextStyle) TextFieldOption {
	return func(tf *TextField) {
		tf.Style = style
	}
}

// WithObscureText sets whether to obscure text
func WithObscureText(obscure bool) TextFieldOption {
	return func(tf *TextField) {
		tf.ObscureText = obscure
	}
}

// WithMaxLength sets the maximum length
func WithMaxLength(length int) TextFieldOption {
	return func(tf *TextField) {
		tf.MaxLength = length
	}
}

// WithOnChanged sets the change callback
func WithOnChanged(callback func(string)) TextFieldOption {
	return func(tf *TextField) {
		tf.OnChanged = callback
	}
}

// WithOnSubmitted sets the submit callback
func WithOnSubmitted(callback func(string)) TextFieldOption {
	return func(tf *TextField) {
		tf.OnSubmitted = callback
	}
}

// WithTextFieldFocusNode sets the focus node
func WithTextFieldFocusNode(node *input.FocusNode) TextFieldOption {
	return func(tf *TextField) {
		tf.FocusNode = node
	}
}

// WithTextFieldAutofocus sets autofocus
func WithTextFieldAutofocus(autofocus bool) TextFieldOption {
	return func(tf *TextField) {
		tf.Autofocus = autofocus
	}
}

// WithDecoration sets the input decoration
func WithDecoration(decoration *InputDecoration) TextFieldOption {
	return func(tf *TextField) {
		tf.Decoration = decoration
	}
}

// Build implements the Widget interface
func (tf *TextField) Build(ctx goflow.BuildContext) goflow.Widget {
	// Build the text field UI
	content := tf.buildContent()

	// Wrap with Focus widget
	return NewFocus(
		content,
		WithFocusNode(tf.FocusNode),
		WithAutofocus(tf.Autofocus),
		WithFocusOnKey(tf.handleKeyEvent),
	)
}

// buildContent builds the visual content of the text field
func (tf *TextField) buildContent() goflow.Widget {
	// Get the display text
	displayText := tf.Controller.Text()
	if tf.ObscureText && len(displayText) > 0 {
		displayText = strings.Repeat("•", len(displayText))
	}

	// If empty, show placeholder
	if len(displayText) == 0 && tf.Placeholder != "" {
		displayText = tf.Placeholder
	}

	// Create the text widget
	textWidget := &Text{
		Text:  displayText,
		Style: tf.Style,
	}

	// Wrap in container with decoration
	return &Container{
		Padding: tf.Decoration.ContentPadding,
		Decoration: &BoxDecoration{
			Color: tf.Decoration.FillColor,
			Border: &Border{
				Color: tf.Decoration.BorderColor,
				Width: tf.Decoration.BorderWidth,
			},
			BorderRadius: &BorderRadius{
				TopLeft:     tf.Decoration.BorderRadius,
				TopRight:    tf.Decoration.BorderRadius,
				BottomLeft:  tf.Decoration.BorderRadius,
				BottomRight: tf.Decoration.BorderRadius,
			},
		},
		Child: textWidget,
	}
}

// handleKeyEvent handles keyboard events
func (tf *TextField) handleKeyEvent(event *input.KeyboardEvent) bool {
	if !tf.Enabled {
		return false
	}

	// Only handle key down events
	if event.Type != input.EventTypeKeyDown {
		return false
	}

	handled := false

	// Handle printable characters
	if len(event.Character) > 0 {
		tf.insertText(event.Character)
		handled = true
	}

	// Handle special keys
	switch event.Key {
	case input.KeyBackspace:
		tf.handleBackspace()
		handled = true

	case input.KeyDelete:
		tf.handleDelete()
		handled = true

	case input.KeyEnter, input.KeyReturn:
		if tf.OnSubmitted != nil {
			tf.OnSubmitted(tf.Controller.Text())
		}
		handled = true

	case input.KeyLeft:
		tf.moveCursorLeft()
		handled = true

	case input.KeyRight:
		tf.moveCursorRight()
		handled = true

	case input.KeyHome:
		tf.moveCursorToStart()
		handled = true

	case input.KeyEnd:
		tf.moveCursorToEnd()
		handled = true
	}

	return handled
}

// insertText inserts text at the current cursor position
func (tf *TextField) insertText(text string) {
	// Check max length
	if tf.MaxLength > 0 && len(tf.Controller.Text())+len(text) > tf.MaxLength {
		return
	}

	currentText := tf.Controller.Text()
	selection := tf.Controller.Selection()

	// Insert text at cursor
	newText := currentText[:selection.ExtentOffset] + text + currentText[selection.ExtentOffset:]
	newCursor := selection.ExtentOffset + len(text)

	tf.Controller.SetText(newText)
	tf.Controller.SetSelection(TextSelection{
		BaseOffset:   newCursor,
		ExtentOffset: newCursor,
	})

	// Notify callback
	if tf.OnChanged != nil {
		tf.OnChanged(newText)
	}
}

// handleBackspace handles backspace key
func (tf *TextField) handleBackspace() {
	currentText := tf.Controller.Text()
	selection := tf.Controller.Selection()

	if selection.ExtentOffset > 0 {
		newText := currentText[:selection.ExtentOffset-1] + currentText[selection.ExtentOffset:]
		newCursor := selection.ExtentOffset - 1

		tf.Controller.SetText(newText)
		tf.Controller.SetSelection(TextSelection{
			BaseOffset:   newCursor,
			ExtentOffset: newCursor,
		})

		if tf.OnChanged != nil {
			tf.OnChanged(newText)
		}
	}
}

// handleDelete handles delete key
func (tf *TextField) handleDelete() {
	currentText := tf.Controller.Text()
	selection := tf.Controller.Selection()

	if selection.ExtentOffset < len(currentText) {
		newText := currentText[:selection.ExtentOffset] + currentText[selection.ExtentOffset+1:]

		tf.Controller.SetText(newText)

		if tf.OnChanged != nil {
			tf.OnChanged(newText)
		}
	}
}

// moveCursorLeft moves the cursor left
func (tf *TextField) moveCursorLeft() {
	selection := tf.Controller.Selection()
	if selection.ExtentOffset > 0 {
		newCursor := selection.ExtentOffset - 1
		tf.Controller.SetSelection(TextSelection{
			BaseOffset:   newCursor,
			ExtentOffset: newCursor,
		})
	}
}

// moveCursorRight moves the cursor right
func (tf *TextField) moveCursorRight() {
	selection := tf.Controller.Selection()
	if selection.ExtentOffset < len(tf.Controller.Text()) {
		newCursor := selection.ExtentOffset + 1
		tf.Controller.SetSelection(TextSelection{
			BaseOffset:   newCursor,
			ExtentOffset: newCursor,
		})
	}
}

// moveCursorToStart moves the cursor to the start
func (tf *TextField) moveCursorToStart() {
	tf.Controller.SetSelection(TextSelection{
		BaseOffset:   0,
		ExtentOffset: 0,
	})
}

// moveCursorToEnd moves the cursor to the end
func (tf *TextField) moveCursorToEnd() {
	endPos := len(tf.Controller.Text())
	tf.Controller.SetSelection(TextSelection{
		BaseOffset:   endPos,
		ExtentOffset: endPos,
	})
}
