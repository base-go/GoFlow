package widgets

import (
	"strings"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// TextEditingController manages text editing state
type TextEditingController struct {
	// Text content
	text string

	// Cursor position (insertion point)
	cursorPosition int

	// Selection start/end (-1 if no selection)
	selectionStart int
	selectionEnd   int

	// Change listeners
	listeners []func()
}

// NewTextEditingController creates a new text editing controller
func NewTextEditingController(initialText string) *TextEditingController {
	return &TextEditingController{
		text:           initialText,
		cursorPosition: len(initialText),
		selectionStart: -1,
		selectionEnd:   -1,
		listeners:      make([]func(), 0),
	}
}

// Text returns the current text
func (t *TextEditingController) Text() string {
	return t.text
}

// SetText updates the text
func (t *TextEditingController) SetText(text string) {
	t.text = text
	t.cursorPosition = len(text)
	t.selectionStart = -1
	t.selectionEnd = -1
	t.notifyListeners()
}

// CursorPosition returns the current cursor position
func (t *TextEditingController) CursorPosition() int {
	return t.cursorPosition
}

// SetCursorPosition updates the cursor position
func (t *TextEditingController) SetCursorPosition(pos int) {
	if pos < 0 {
		pos = 0
	}
	if pos > len(t.text) {
		pos = len(t.text)
	}
	t.cursorPosition = pos
	t.selectionStart = -1
	t.selectionEnd = -1
	t.notifyListeners()
}

// Selection returns the current selection
func (t *TextEditingController) Selection() (start, end int, hasSelection bool) {
	if t.selectionStart >= 0 && t.selectionEnd >= 0 {
		return t.selectionStart, t.selectionEnd, true
	}
	return 0, 0, false
}

// SetSelection updates the selection
func (t *TextEditingController) SetSelection(start, end int) {
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if end > len(t.text) {
		end = len(t.text)
	}
	t.selectionStart = start
	t.selectionEnd = end
	t.cursorPosition = end
	t.notifyListeners()
}

// ClearSelection clears the current selection
func (t *TextEditingController) ClearSelection() {
	t.selectionStart = -1
	t.selectionEnd = -1
	t.notifyListeners()
}

// InsertText inserts text at the cursor position
func (t *TextEditingController) InsertText(text string) {
	// Delete selection if any
	if t.selectionStart >= 0 && t.selectionEnd >= 0 {
		t.text = t.text[:t.selectionStart] + t.text[t.selectionEnd:]
		t.cursorPosition = t.selectionStart
		t.selectionStart = -1
		t.selectionEnd = -1
	}

	// Insert text
	before := t.text[:t.cursorPosition]
	after := t.text[t.cursorPosition:]
	t.text = before + text + after
	t.cursorPosition += len(text)

	t.notifyListeners()
}

// DeleteBackward deletes the character before the cursor
func (t *TextEditingController) DeleteBackward() {
	// Delete selection if any
	if t.selectionStart >= 0 && t.selectionEnd >= 0 {
		t.text = t.text[:t.selectionStart] + t.text[t.selectionEnd:]
		t.cursorPosition = t.selectionStart
		t.selectionStart = -1
		t.selectionEnd = -1
		t.notifyListeners()
		return
	}

	// Delete single character
	if t.cursorPosition > 0 {
		t.text = t.text[:t.cursorPosition-1] + t.text[t.cursorPosition:]
		t.cursorPosition--
		t.notifyListeners()
	}
}

// DeleteForward deletes the character after the cursor
func (t *TextEditingController) DeleteForward() {
	// Delete selection if any
	if t.selectionStart >= 0 && t.selectionEnd >= 0 {
		t.text = t.text[:t.selectionStart] + t.text[t.selectionEnd:]
		t.cursorPosition = t.selectionStart
		t.selectionStart = -1
		t.selectionEnd = -1
		t.notifyListeners()
		return
	}

	// Delete single character
	if t.cursorPosition < len(t.text) {
		t.text = t.text[:t.cursorPosition] + t.text[t.cursorPosition+1:]
		t.notifyListeners()
	}
}

// MoveCursorLeft moves the cursor left by one character
func (t *TextEditingController) MoveCursorLeft(shift bool) {
	if shift {
		// Extend selection
		if t.selectionStart < 0 {
			t.selectionStart = t.cursorPosition
			t.selectionEnd = t.cursorPosition
		}
		if t.cursorPosition > 0 {
			t.cursorPosition--
			if t.cursorPosition < t.selectionStart {
				t.selectionStart = t.cursorPosition
			} else {
				t.selectionEnd = t.cursorPosition
			}
		}
	} else {
		t.ClearSelection()
		if t.cursorPosition > 0 {
			t.cursorPosition--
		}
	}
	t.notifyListeners()
}

// MoveCursorRight moves the cursor right by one character
func (t *TextEditingController) MoveCursorRight(shift bool) {
	if shift {
		// Extend selection
		if t.selectionStart < 0 {
			t.selectionStart = t.cursorPosition
			t.selectionEnd = t.cursorPosition
		}
		if t.cursorPosition < len(t.text) {
			t.cursorPosition++
			if t.cursorPosition > t.selectionEnd {
				t.selectionEnd = t.cursorPosition
			} else {
				t.selectionStart = t.cursorPosition
			}
		}
	} else {
		t.ClearSelection()
		if t.cursorPosition < len(t.text) {
			t.cursorPosition++
		}
	}
	t.notifyListeners()
}

// MoveCursorToStart moves the cursor to the start
func (t *TextEditingController) MoveCursorToStart(shift bool) {
	if shift {
		if t.selectionStart < 0 {
			t.selectionStart = 0
			t.selectionEnd = t.cursorPosition
		} else {
			t.selectionStart = 0
		}
	} else {
		t.ClearSelection()
	}
	t.cursorPosition = 0
	t.notifyListeners()
}

// MoveCursorToEnd moves the cursor to the end
func (t *TextEditingController) MoveCursorToEnd(shift bool) {
	if shift {
		if t.selectionStart < 0 {
			t.selectionStart = t.cursorPosition
			t.selectionEnd = len(t.text)
		} else {
			t.selectionEnd = len(t.text)
		}
	} else {
		t.ClearSelection()
	}
	t.cursorPosition = len(t.text)
	t.notifyListeners()
}

// AddListener adds a change listener
func (t *TextEditingController) AddListener(listener func()) {
	t.listeners = append(t.listeners, listener)
}

// RemoveListener removes a change listener
func (t *TextEditingController) RemoveListener(listener func()) {
	for i, l := range t.listeners {
		if &l == &listener {
			t.listeners = append(t.listeners[:i], t.listeners[i+1:]...)
			return
		}
	}
}

func (t *TextEditingController) notifyListeners() {
	for _, listener := range t.listeners {
		listener()
	}
}

// Dispose cleans up the controller
func (t *TextEditingController) Dispose() {
	t.listeners = nil
}

// TextField is a text input field widget
type TextField struct {
	goflow.BaseWidget

	// Controller for text editing
	Controller *TextEditingController

	// Focus node
	FocusNode *FocusNode

	// Decoration
	Decoration *InputDecoration

	// Style
	Style *goflow.TextStyle

	// Placeholder text
	Placeholder string

	// Placeholder style
	PlaceholderStyle *goflow.TextStyle

	// Obscure text (for passwords)
	ObscureText bool

	// Max lines (1 for single-line, 0 for unlimited)
	MaxLines int

	// Min lines
	MinLines int

	// Callbacks
	OnChanged  func(text string)
	OnSubmitted func(text string)

	// Read-only
	ReadOnly bool

	// Auto-focus
	AutoFocus bool
}

// InputDecoration defines the visual decoration of a TextField
type InputDecoration struct {
	// Border
	Border *InputBorder

	// Enabled border (when not focused)
	EnabledBorder *InputBorder

	// Focused border
	FocusedBorder *InputBorder

	// Error border
	ErrorBorder *InputBorder

	// Fill color
	FillColor *goflow.Color

	// Filled
	Filled bool

	// Content padding
	ContentPadding *goflow.EdgeInsets

	// Hint text
	HintText string

	// Hint style
	HintStyle *goflow.TextStyle

	// Label text
	LabelText string

	// Label style
	LabelStyle *goflow.TextStyle

	// Prefix icon
	PrefixIcon goflow.Widget

	// Suffix icon
	SuffixIcon goflow.Widget
}

// InputBorder defines the border of a TextField
type InputBorder struct {
	// Border side
	BorderSide *BorderSide

	// Border radius
	BorderRadius float64
}

// BorderSide defines a border side
type BorderSide struct {
	Color *goflow.Color
	Width float64
}

// NewTextField creates a new TextField
func NewTextField(controller *TextEditingController) *TextField {
	if controller == nil {
		controller = NewTextEditingController("")
	}

	return &TextField{
		Controller:  controller,
		FocusNode:   NewFocusNode(),
		MaxLines:    1,
		MinLines:    1,
		ObscureText: false,
		ReadOnly:    false,
		AutoFocus:   false,
	}
}

// Build creates the widget tree
func (tf *TextField) Build(context goflow.BuildContext) goflow.Widget {
	// This is a simplified build - in a real implementation,
	// this would create a complex widget tree with borders,
	// padding, text rendering, cursor, selection, etc.

	// For now, just wrap in a Focus widget
	return NewFocus(
		tf.buildContent(context),
		tf.FocusNode,
	)
}

func (tf *TextField) buildContent(context goflow.BuildContext) goflow.Widget {
	// Create editable text with cursor
	editableText := &EditableText{
		Controller:       tf.Controller,
		FocusNode:        tf.FocusNode,
		Style:            tf.Style,
		Placeholder:      tf.Placeholder,
		PlaceholderStyle: tf.PlaceholderStyle,
		ObscureText:      tf.ObscureText,
		OnChanged:        tf.OnChanged,
		OnSubmitted:      tf.OnSubmitted,
	}

	// Wrap in container with decoration
	if tf.Decoration != nil {
		return &Container{
			Child:   editableText,
			Padding: tf.Decoration.ContentPadding,
		}
	}

	return editableText
}

// EditableText is the render object widget for editable text with cursor
type EditableText struct {
	goflow.BaseWidget
	Controller       *TextEditingController
	FocusNode        *FocusNode
	Style            *goflow.TextStyle
	Placeholder      string
	PlaceholderStyle *goflow.TextStyle
	ObscureText      bool
	OnChanged        func(string)
	OnSubmitted      func(string)
}

// CreateElement creates a render object element
func (e *EditableText) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      e,
		renderObject: &RenderEditableText{
			BaseRenderBox:    goflow.NewBaseRenderBox(),
			controller:       e.Controller,
			focusNode:        e.FocusNode,
			style:            e.Style,
			placeholder:      e.Placeholder,
			placeholderStyle: e.PlaceholderStyle,
			obscureText:      e.ObscureText,
			cursorVisible:    true,
		},
	}
}

// Build returns nil (this is a render object widget)
func (e *EditableText) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderEditableText is the render object for editable text
type RenderEditableText struct {
	*goflow.BaseRenderBox
	controller       *TextEditingController
	focusNode        *FocusNode
	style            *goflow.TextStyle
	placeholder      string
	placeholderStyle *goflow.TextStyle
	obscureText      bool
	cursorVisible    bool
	cursorBlinkTimer float64
}

// PerformLayout performs layout
func (r *RenderEditableText) PerformLayout() {
	constraints := r.GetConstraints()

	// Get text to display
	displayText := r.controller.Text()
	style := r.style
	if style == nil {
		style = goflow.NewTextStyle()
	}

	if displayText == "" && r.placeholder != "" {
		displayText = r.placeholder
		if r.placeholderStyle != nil {
			style = r.placeholderStyle
		}
	}

	// Obscure text if needed
	if r.obscureText && displayText != "" && displayText != r.placeholder {
		displayText = strings.Repeat("•", len(displayText))
	}

	// Calculate text size (simplified)
	width := float64(len(displayText)) * style.FontSize * 0.6
	height := style.FontSize * 1.4 // Extra height for cursor

	// Add padding for cursor
	width += 2.0

	size := constraints.Constrain(goflow.NewSize(width, height))
	r.SetSize(size)
}

// Layout performs the layout
func (r *RenderEditableText) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the editable text with cursor
func (r *RenderEditableText) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// Get text to display
	displayText := r.controller.Text()
	style := r.style
	if style == nil {
		style = goflow.NewTextStyle()
	}

	isPlaceholder := false
	if displayText == "" && r.placeholder != "" {
		displayText = r.placeholder
		isPlaceholder = true
		if r.placeholderStyle != nil {
			style = r.placeholderStyle
		}
	}

	// Obscure text if needed
	obscuredText := displayText
	if r.obscureText && displayText != "" && !isPlaceholder {
		obscuredText = strings.Repeat("•", len(displayText))
	}

	// Draw text
	if obscuredText != "" {
		canvas.DrawText(obscuredText, offset, style)
	}

	// Draw cursor if focused and visible
	if r.focusNode != nil && r.focusNode.HasFocus() && r.cursorVisible {
		cursorPosition := r.controller.CursorPosition()

		// Calculate cursor X position (simplified)
		cursorX := offset.X
		if cursorPosition > 0 {
			textBeforeCursor := obscuredText
			if cursorPosition < len(obscuredText) {
				textBeforeCursor = obscuredText[:cursorPosition]
			}
			cursorX += float64(len(textBeforeCursor)) * style.FontSize * 0.6
		}

		// Draw cursor line
		paint := goflow.NewPaint()
		paint.Color = style.Color
		if paint.Color == nil {
			paint.Color = goflow.NewColor(0, 0, 0, 255)
		}
		paint.StrokeWidth = 2.0

		cursorTop := offset.Y
		cursorBottom := offset.Y + style.FontSize * 1.2

		canvas.DrawLine(
			goflow.NewOffset(cursorX, cursorTop),
			goflow.NewOffset(cursorX, cursorBottom),
			paint,
		)
	}

	// Draw selection if any
	start, end, hasSelection := r.controller.Selection()
	if hasSelection && start != end {
		// Calculate selection bounds (simplified)
		selectionStart := offset.X + float64(start)*style.FontSize*0.6
		selectionEnd := offset.X + float64(end)*style.FontSize*0.6
		selectionWidth := selectionEnd - selectionStart

		// Draw selection background
		paint := goflow.NewPaint()
		paint.Color = goflow.NewColor(33, 150, 243, 100) // Semi-transparent blue

		selectionRect := goflow.NewRect(
			goflow.NewOffset(selectionStart, offset.Y),
			goflow.NewSize(selectionWidth, style.FontSize*1.2),
		)
		canvas.DrawRect(selectionRect, paint)
	}
}
