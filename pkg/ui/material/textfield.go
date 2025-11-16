package material

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// TextField is a Material Design text input field
type TextField struct {
	goflow.BaseWidget

	// Controller for the text value
	Controller *TextEditingController

	// Decoration
	Decoration *InputDecoration

	// Keyboard type
	KeyboardType KeyboardType

	// Text style
	Style *goflow.TextStyle

	// Max lines (1 for single line, >1 for multiline)
	MaxLines int

	// Obscure text (for passwords)
	ObscureText bool

	// Enabled state
	Enabled bool

	// Callbacks
	OnChanged func(string)
	OnSubmitted func(string)
}

// TextEditingController manages the text being edited
type TextEditingController struct {
	Text *signals.Signal[string]
}

// NewTextEditingController creates a new text editing controller
func NewTextEditingController(initialValue string) *TextEditingController {
	return &TextEditingController{
		Text: signals.New(initialValue),
	}
}

// InputDecoration defines the decoration for a text field
type InputDecoration struct {
	HintText      string
	LabelText     string
	HelperText    string
	ErrorText     string
	PrefixIcon    goflow.Widget
	SuffixIcon    goflow.Widget
	Border        InputBorder
	FillColor     *goflow.Color
	Filled        bool
}

// InputBorder defines the border style
type InputBorder int

const (
	InputBorderOutline InputBorder = iota
	InputBorderUnderline
	InputBorderNone
)

// KeyboardType defines the type of keyboard to show
type KeyboardType int

const (
	KeyboardTypeText KeyboardType = iota
	KeyboardTypeNumber
	KeyboardTypePhone
	KeyboardTypeEmail
	KeyboardTypeURL
	KeyboardTypeMultiline
)

// NewTextField creates a new Material text field
func NewTextField() *TextField {
	return &TextField{
		Controller:  NewTextEditingController(""),
		MaxLines:    1,
		ObscureText: false,
		Enabled:     true,
		Decoration: &InputDecoration{
			Border: InputBorderUnderline,
		},
	}
}

// Build creates the widget tree for the text field
func (t *TextField) Build(context goflow.BuildContext) goflow.Widget {
	theme := DefaultLightTheme()

	// Background color
	bgColor := goflow.NewColor(0, 0, 0, 0) // Transparent
	if t.Decoration != nil && t.Decoration.Filled {
		bgColor = t.Decoration.FillColor
		if bgColor == nil {
			bgColor = goflow.NewColor(245, 245, 245, 255) // Light grey
		}
	}

	// Build the text display
	displayText := t.Controller.Text.Get()
	if t.ObscureText && displayText != "" {
		// Replace with dots for password
		displayText = "•••••••"
	}

	var children []goflow.Widget

	// Add prefix icon if present
	if t.Decoration != nil && t.Decoration.PrefixIcon != nil {
		children = append(children, t.Decoration.PrefixIcon)
	}

	// Add text content
	textWidget := &widgets.Text{
		Data:  displayText,
		Style: t.Style,
	}

	if t.Style == nil {
		textWidget.Style = goflow.NewTextStyle()
		textWidget.Style.Color = theme.TextPrimaryColor
	}

	children = append(children, &widgets.Expanded{
		Child: textWidget,
	})

	// Add suffix icon if present
	if t.Decoration != nil && t.Decoration.SuffixIcon != nil {
		children = append(children, t.Decoration.SuffixIcon)
	}

	content := &widgets.Row{
		Children:           children,
		MainAxisAlignment:  widgets.MainAxisStart,
		CrossAxisAlignment: widgets.CrossAxisCenter,
	}

	// Wrap in container with padding and decoration
	minHeight := 48.0
	return &widgets.Container{
		Color:   bgColor,
		Padding: goflow.NewEdgeInsets(12, 8, 12, 8),
		Height:  &minHeight,
		Child:   content,
	}
}
