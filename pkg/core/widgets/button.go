package widgets

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// ButtonStyle defines the visual style of a button
type ButtonStyle struct {
	// Colors
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color
	OverlayColor    *goflow.Color // Hover/press overlay

	// Shadow (for elevated buttons)
	Elevation       float64
	ShadowColor     *goflow.Color

	// Border
	Side            *BorderSide

	// Shape
	BorderRadius    float64

	// Padding
	Padding         *goflow.EdgeInsets

	// Minimum size
	MinimumSize     *goflow.Size

	// Text style
	TextStyle       *goflow.TextStyle

	// Animation duration
	AnimationDuration int // milliseconds
}

// DefaultElevatedButtonStyle returns the default style for elevated buttons
func DefaultElevatedButtonStyle() *ButtonStyle {
	return &ButtonStyle{
		BackgroundColor: goflow.NewColor(98, 0, 238, 255),   // Primary blue
		ForegroundColor: goflow.NewColor(255, 255, 255, 255), // White text
		OverlayColor:    goflow.NewColor(255, 255, 255, 30),  // White overlay
		Elevation:       2.0,
		ShadowColor:     goflow.NewColor(0, 0, 0, 100),
		BorderRadius:    4.0,
		Padding: &goflow.EdgeInsets{
			Left:   16,
			Right:  16,
			Top:    8,
			Bottom: 8,
		},
		MinimumSize: &goflow.Size{
			Width:  64,
			Height: 36,
		},
		AnimationDuration: 200,
	}
}

// DefaultTextButtonStyle returns the default style for text buttons
func DefaultTextButtonStyle() *ButtonStyle {
	return &ButtonStyle{
		BackgroundColor: nil, // Transparent
		ForegroundColor: goflow.NewColor(98, 0, 238, 255),    // Primary blue
		OverlayColor:    goflow.NewColor(98, 0, 238, 30),     // Blue overlay
		Elevation:       0,
		BorderRadius:    4.0,
		Padding: &goflow.EdgeInsets{
			Left:   8,
			Right:  8,
			Top:    8,
			Bottom: 8,
		},
		MinimumSize: &goflow.Size{
			Width:  64,
			Height: 36,
		},
		AnimationDuration: 200,
	}
}

// DefaultOutlinedButtonStyle returns the default style for outlined buttons
func DefaultOutlinedButtonStyle() *ButtonStyle {
	return &ButtonStyle{
		BackgroundColor: nil, // Transparent
		ForegroundColor: goflow.NewColor(98, 0, 238, 255), // Primary blue
		OverlayColor:    goflow.NewColor(98, 0, 238, 30),  // Blue overlay
		Side: &BorderSide{
			Color: goflow.NewColor(200, 200, 200, 255),
			Width: 1.0,
		},
		Elevation:    0,
		BorderRadius: 4.0,
		Padding: &goflow.EdgeInsets{
			Left:   16,
			Right:  16,
			Top:    8,
			Bottom: 8,
		},
		MinimumSize: &goflow.Size{
			Width:  64,
			Height: 36,
		},
		AnimationDuration: 200,
	}
}

// ButtonState represents the current state of a button
type ButtonState int

const (
	ButtonStateNormal ButtonState = iota
	ButtonStateHovered
	ButtonStatePressed
	ButtonStateDisabled
	ButtonStateFocused
)

// ElevatedButton is a Material Design elevated button
type ElevatedButton struct {
	goflow.BaseWidget

	// Content (usually Text widget)
	Child goflow.Widget

	// Callback
	OnPressed func()

	// Style
	Style *ButtonStyle

	// Focus node
	FocusNode *FocusNode

	// Auto-focus
	AutoFocus bool

	// Enabled
	Enabled bool
}

// NewElevatedButton creates a new elevated button
func NewElevatedButton(child goflow.Widget, onPressed func()) *ElevatedButton {
	return &ElevatedButton{
		Child:     child,
		OnPressed: onPressed,
		Style:     DefaultElevatedButtonStyle(),
		Enabled:   true,
		AutoFocus: false,
	}
}

// Build creates the widget tree
func (b *ElevatedButton) Build(context goflow.BuildContext) goflow.Widget {
	if !b.Enabled {
		return b.buildDisabled(context)
	}

	// Wrap in gesture detector for tap handling
	detector := &GestureDetector{
		Child:   b.buildContent(context),
		OnTap:   b.OnPressed,
		Behavior: HitTestBehaviorOpaque,
	}

	// Wrap in focus if focus node provided
	if b.FocusNode != nil {
		return NewFocus(detector, b.FocusNode)
	}

	return detector
}

func (b *ElevatedButton) buildContent(context goflow.BuildContext) goflow.Widget {
	// Apply text style to child if it's Text
	child := b.Child
	if text, ok := child.(*Text); ok && b.Style.TextStyle != nil {
		styledText := &Text{
			Data:  text.Data,
			Style: b.Style.TextStyle,
		}
		if styledText.Style == nil {
			styledText.Style = goflow.NewTextStyle()
		}
		if b.Style.ForegroundColor != nil {
			styledText.Style.Color = b.Style.ForegroundColor
		}
		child = styledText
	}

	// Wrap in padding
	paddedChild := &Padding{
		Padding: b.Style.Padding,
		Child:   child,
	}

	// Wrap in container with decoration
	width := b.Style.MinimumSize.Width
	height := b.Style.MinimumSize.Height
	return &Container{
		Child:  paddedChild,
		Color:  b.Style.BackgroundColor,
		Width:  &width,
		Height: &height,
	}
}

func (b *ElevatedButton) buildDisabled(context goflow.BuildContext) goflow.Widget {
	// Create disabled style
	disabledStyle := &ButtonStyle{
		BackgroundColor: goflow.NewColor(200, 200, 200, 255),
		ForegroundColor: goflow.NewColor(150, 150, 150, 255),
		BorderRadius:    b.Style.BorderRadius,
		Padding:         b.Style.Padding,
		MinimumSize:     b.Style.MinimumSize,
	}

	child := b.Child
	if text, ok := child.(*Text); ok {
		styledText := &Text{
			Data:  text.Data,
			Style: goflow.NewTextStyle(),
		}
		styledText.Style.Color = disabledStyle.ForegroundColor
		child = styledText
	}

	width := disabledStyle.MinimumSize.Width
	height := disabledStyle.MinimumSize.Height
	return &Container{
		Child:  &Padding{Padding: disabledStyle.Padding, Child: child},
		Color:  disabledStyle.BackgroundColor,
		Width:  &width,
		Height: &height,
	}
}

// TextButton is a Material Design text button (flat)
type TextButton struct {
	goflow.BaseWidget

	// Content (usually Text widget)
	Child goflow.Widget

	// Callback
	OnPressed func()

	// Style
	Style *ButtonStyle

	// Focus node
	FocusNode *FocusNode

	// Auto-focus
	AutoFocus bool

	// Enabled
	Enabled bool
}

// NewTextButton creates a new text button
func NewTextButton(child goflow.Widget, onPressed func()) *TextButton {
	return &TextButton{
		Child:     child,
		OnPressed: onPressed,
		Style:     DefaultTextButtonStyle(),
		Enabled:   true,
		AutoFocus: false,
	}
}

// Build creates the widget tree
func (b *TextButton) Build(context goflow.BuildContext) goflow.Widget {
	if !b.Enabled {
		return b.buildDisabled(context)
	}

	// Wrap in gesture detector for tap handling
	detector := &GestureDetector{
		Child:    b.buildContent(context),
		OnTap:    b.OnPressed,
		Behavior: HitTestBehaviorOpaque,
	}

	// Wrap in focus if focus node provided
	if b.FocusNode != nil {
		return NewFocus(detector, b.FocusNode)
	}

	return detector
}

func (b *TextButton) buildContent(context goflow.BuildContext) goflow.Widget {
	// Apply text style to child if it's Text
	child := b.Child
	if text, ok := child.(*Text); ok {
		styledText := &Text{
			Data:  text.Data,
			Style: goflow.NewTextStyle(),
		}
		if b.Style.ForegroundColor != nil {
			styledText.Style.Color = b.Style.ForegroundColor
		}
		child = styledText
	}

	// Wrap in padding
	paddedChild := &Padding{
		Padding: b.Style.Padding,
		Child:   child,
	}

	// Wrap in container (transparent background)
	width := b.Style.MinimumSize.Width
	height := b.Style.MinimumSize.Height
	return &Container{
		Child:  paddedChild,
		Width:  &width,
		Height: &height,
	}
}

func (b *TextButton) buildDisabled(context goflow.BuildContext) goflow.Widget {
	child := b.Child
	if text, ok := child.(*Text); ok {
		styledText := &Text{
			Data:  text.Data,
			Style: goflow.NewTextStyle(),
		}
		styledText.Style.Color = goflow.NewColor(150, 150, 150, 255)
		child = styledText
	}

	width := b.Style.MinimumSize.Width
	height := b.Style.MinimumSize.Height
	return &Container{
		Child:  &Padding{Padding: b.Style.Padding, Child: child},
		Width:  &width,
		Height: &height,
	}
}

// OutlinedButton is a Material Design outlined button
type OutlinedButton struct {
	goflow.BaseWidget

	// Content (usually Text widget)
	Child goflow.Widget

	// Callback
	OnPressed func()

	// Style
	Style *ButtonStyle

	// Focus node
	FocusNode *FocusNode

	// Auto-focus
	AutoFocus bool

	// Enabled
	Enabled bool
}

// NewOutlinedButton creates a new outlined button
func NewOutlinedButton(child goflow.Widget, onPressed func()) *OutlinedButton {
	return &OutlinedButton{
		Child:     child,
		OnPressed: onPressed,
		Style:     DefaultOutlinedButtonStyle(),
		Enabled:   true,
		AutoFocus: false,
	}
}

// Build creates the widget tree
func (b *OutlinedButton) Build(context goflow.BuildContext) goflow.Widget {
	if !b.Enabled {
		return b.buildDisabled(context)
	}

	// Wrap in gesture detector for tap handling
	detector := &GestureDetector{
		Child:    b.buildContent(context),
		OnTap:    b.OnPressed,
		Behavior: HitTestBehaviorOpaque,
	}

	// Wrap in focus if focus node provided
	if b.FocusNode != nil {
		return NewFocus(detector, b.FocusNode)
	}

	return detector
}

func (b *OutlinedButton) buildContent(context goflow.BuildContext) goflow.Widget {
	// Apply text style to child if it's Text
	child := b.Child
	if text, ok := child.(*Text); ok {
		styledText := &Text{
			Data:  text.Data,
			Style: goflow.NewTextStyle(),
		}
		if b.Style.ForegroundColor != nil {
			styledText.Style.Color = b.Style.ForegroundColor
		}
		child = styledText
	}

	// Wrap in padding
	paddedChild := &Padding{
		Padding: b.Style.Padding,
		Child:   child,
	}

	// Wrap in container with border
	// Note: Border decoration would be added here in full implementation
	width := b.Style.MinimumSize.Width
	height := b.Style.MinimumSize.Height
	return &Container{
		Child:  paddedChild,
		Width:  &width,
		Height: &height,
	}
}

func (b *OutlinedButton) buildDisabled(context goflow.BuildContext) goflow.Widget {
	child := b.Child
	if text, ok := child.(*Text); ok {
		styledText := &Text{
			Data:  text.Data,
			Style: goflow.NewTextStyle(),
		}
		styledText.Style.Color = goflow.NewColor(150, 150, 150, 255)
		child = styledText
	}

	width := b.Style.MinimumSize.Width
	height := b.Style.MinimumSize.Height
	return &Container{
		Child:  &Padding{Padding: b.Style.Padding, Child: child},
		Width:  &width,
		Height: &height,
	}
}

// FloatingActionButton is a Material Design floating action button
type FloatingActionButton struct {
	goflow.BaseWidget

	// Icon or child widget
	Child goflow.Widget

	// Callback
	OnPressed func()

	// Colors
	BackgroundColor *goflow.Color
	ForegroundColor *goflow.Color

	// Elevation
	Elevation float64

	// Tooltip
	Tooltip string

	// Mini variant
	Mini bool

	// Enabled
	Enabled bool
}

// NewFloatingActionButton creates a new FAB
func NewFloatingActionButton(child goflow.Widget, onPressed func()) *FloatingActionButton {
	return &FloatingActionButton{
		Child:           child,
		OnPressed:       onPressed,
		BackgroundColor: goflow.NewColor(98, 0, 238, 255),
		ForegroundColor: goflow.NewColor(255, 255, 255, 255),
		Elevation:       6.0,
		Mini:            false,
		Enabled:         true,
	}
}

// Build creates the widget tree
func (f *FloatingActionButton) Build(context goflow.BuildContext) goflow.Widget {
	size := 56.0
	if f.Mini {
		size = 40.0
	}

	// Center the child
	centeredChild := &Center{
		Child: f.Child,
	}

	// Wrap in container with circular shape
	container := &Container{
		Width:  &size,
		Height: &size,
		Color:  f.BackgroundColor,
		Child:  centeredChild,
	}

	if !f.Enabled {
		return container
	}

	// Wrap in gesture detector
	return &GestureDetector{
		Child:    container,
		OnTap:    f.OnPressed,
		Behavior: HitTestBehaviorOpaque,
	}
}
