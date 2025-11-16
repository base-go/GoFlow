package adaptive

import (
	"github.com/base-go/GoFlow/pkg/ui/cupertino"
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/ui/material"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// Button is an adaptive button that automatically uses Material or Cupertino style
// based on the platform
type Button struct {
	goflow.BaseWidget

	// Required
	OnPressed func()
	Text      string
	Child     goflow.Widget

	// Optional styling (will be applied to both variants)
	Color         *goflow.Color
	DisabledColor *goflow.Color
	Enabled       bool

	// Style preference (if not set, uses platform default)
	Style ButtonStyle
}

// ButtonStyle defines the button style preference
type ButtonStyle int

const (
	// ButtonStylePlatformDefault uses platform-specific default
	ButtonStylePlatformDefault ButtonStyle = iota
	// ButtonStyleFilled is a filled button
	ButtonStyleFilled
	// ButtonStyleOutlined is an outlined button (Material only)
	ButtonStyleOutlined
	// ButtonStyleText is a text-only button
	ButtonStyleText
)

// NewButton creates a new adaptive button
func NewButton(text string, onPressed func()) *Button {
	return &Button{
		Text:      text,
		OnPressed: onPressed,
		Enabled:   true,
		Style:     ButtonStylePlatformDefault,
	}
}

// NewTextButton creates a text-only adaptive button
func NewTextButton(text string, onPressed func()) *Button {
	return &Button{
		Text:      text,
		OnPressed: onPressed,
		Enabled:   true,
		Style:     ButtonStyleText,
	}
}

// Build creates the widget tree for the button
func (b *Button) Build(context goflow.BuildContext) goflow.Widget {
	platform := goflow.GetPlatform()

	// Determine which child to use
	child := b.Child
	if child == nil && b.Text != "" {
		child = &widgets.Text{
			Data: b.Text,
		}
	}

	// Switch based on platform
	if platform.IsApple() {
		// Use Cupertino style for iOS/macOS
		btn := cupertino.NewButton(child, b.OnPressed)
		btn.Enabled = b.Enabled
		btn.Color = b.Color
		btn.DisabledColor = b.DisabledColor

		// Map button style
		switch b.Style {
		case ButtonStyleFilled, ButtonStylePlatformDefault:
			btn.Type = cupertino.ButtonTypeFilled
		case ButtonStyleText:
			btn.Type = cupertino.ButtonTypePlain
		case ButtonStyleOutlined:
			// Cupertino doesn't have outlined, use gray
			btn.Type = cupertino.ButtonTypeGray
		}

		return btn
	} else {
		// Use Material style for Android/Linux/Windows/Web
		btn := material.NewButton(child, b.OnPressed)
		btn.Enabled = b.Enabled
		btn.BackgroundColor = b.Color
		btn.DisabledColor = b.DisabledColor

		// Map button style
		switch b.Style {
		case ButtonStyleFilled, ButtonStylePlatformDefault:
			btn.Style = material.ButtonStyleFilled
		case ButtonStyleText:
			btn.Style = material.ButtonStyleText
		case ButtonStyleOutlined:
			btn.Style = material.ButtonStyleOutlined
		}

		return btn
	}
}
