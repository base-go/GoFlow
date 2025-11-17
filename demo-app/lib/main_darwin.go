//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"runtime"
	"time"

	"com.example/demo-app/lib/screens"
	gf "github.com/base-go/GoFlow"
	"github.com/base-go/GoFlow/backends/macos"
)

func init() {
	runtime.LockOSThread()
}

// PageType represents different pages in the app
type PageType int

const (
	PageHome PageType = iota
	PageInputWidgets
	PageLayoutDemo
	PageForms
)

// ButtonState represents the different visual states a button can have
type ButtonState int

const (
	ButtonNormal ButtonState = iota
	ButtonHovered
	ButtonPressed
	ButtonReleased
)

// ButtonVariant represents different button styles like Nuxt UI
type ButtonVariant int

const (
	ButtonElevated ButtonVariant = iota // Default: filled background with elevation (Material Elevated)
	ButtonText                          // Text-only button without background (Material TextButton)
	ButtonOutlined                      // Outlined button with border but no fill
)

// AppState holds the current application state
type AppState struct {
	currentPage      PageType
	clickedButton    int
	checkboxState    bool
	sliderValue      float64
	textInput1       string
	textInput2       string
	formName         string
	formEmail        string
	radioSelection   int
	activeTextInput  int
	hoveredTextInput int
	lastCursorBlink  int64 // For blinking cursor animation

	// Enhanced button state tracking
	hoveredButton int                 // ID of currently hovered button
	pressedButton int                 // ID of currently pressed button
	buttonStates  map[int]ButtonState // Track state of each button
}

var appState = &AppState{
	currentPage:      PageHome,
	clickedButton:    -1,
	checkboxState:    false,
	sliderValue:      50.0,
	textInput1:       "",
	textInput2:       "",
	formName:         "",
	formEmail:        "",
	radioSelection:   0,
	activeTextInput:  0,
	hoveredTextInput: 0,
	lastCursorBlink:  0,
	hoveredButton:    0,
	pressedButton:    0,
	buttonStates:     make(map[int]ButtonState),
}

func main() {
	fmt.Println("GoFlow Kitchen Sink Demo - macOS")
	fmt.Println("=================================\n")

	// Register named routes (kept for framework compatibility)
	routes := map[string]gf.RouteBuilder{
		"/":              func() gf.Widget { return screens.NewHomePage() },
		"/input-widgets": func() gf.Widget { return screens.NewInputWidgetsPage() },
	}
	gf.Get.RegisterRoutes(routes)

	// Initialize macOS app
	macos.InitApp()

	// Create window
	window := macos.NewWindow(900, 700, "GoFlow Kitchen Sink")
	defer window.Destroy()

	// Set mouse callback for button interactions and hover effects
	window.SetMouseFunc(func(button, action int, x, y float64) {
		if action == 1 { // Mouse down
			handleMouseDown(x, y)
			window.SetNeedsDisplay() // Trigger redraw
		} else if action == 0 { // Mouse up
			handleMouseUp(x, y)
			window.SetNeedsDisplay() // Trigger redraw
		} else if action == 2 { // Mouse move
			handleMouseMove(x, y)
			window.SetNeedsDisplay() // Trigger redraw for hover effects
		}
	})

	// Set draw callback
	window.SetDrawFunc(func(canvas *macos.CoreGraphicsCanvas) {
		// Clear background
		canvas.Clear(gf.NewColor(245, 245, 245, 255))

		// Render current page
		switch appState.currentPage {
		case PageHome:
			drawHomePage(canvas)
		case PageInputWidgets:
			drawInputWidgetsPage(canvas)
		case PageLayoutDemo:
			drawLayoutDemoPage(canvas)
		case PageForms:
			drawFormsPage(canvas)
		}
	})

	// Set resize callback
	window.SetResizeFunc(func(width, height int) {
		fmt.Printf("Window resized to: %dx%d\n", width, height)
	})

	// Set keyboard callback
	window.SetKeyFunc(func(key, action int) {
		if action == 1 { // Key press
			// ESC key to go back
			if key == 53 {
				if appState.currentPage != PageHome {
					appState.currentPage = PageHome
					window.SetNeedsDisplay()
				}
			}

			// Handle text input for active text fields
			if appState.activeTextInput > 0 {
				handleTextInput(key, window)
			}
		}
	})

	// Show the window
	window.Show()

	fmt.Println("\n✅ Window created successfully!")
	fmt.Println("📱 GoFlow Kitchen Sink - Multi-Page Demo")
	fmt.Println("\nFeatures:")
	fmt.Println("  • 4 interactive pages with routing")
	fmt.Println("  • Click buttons to navigate between pages")
	fmt.Println("  • Press ESC to go back to home")
	fmt.Println("  • Interactive widgets with state management")
	fmt.Println("\nClose the window to exit.\n")

	// Trigger initial draw
	window.SetNeedsDisplay()

	// Run the event loop
	macos.Run()
}

// drawHomePage renders the home page
func drawHomePage(canvas *macos.CoreGraphicsCanvas) {
	// Draw header
	titleStyle := gf.NewTextStyle()
	titleStyle.FontSize = 36
	titleStyle.Color = gf.NewColor(33, 33, 33, 255)
	canvas.DrawText("GoFlow Kitchen Sink", gf.NewOffset(50, 50), titleStyle)

	subtitleStyle := gf.NewTextStyle()
	subtitleStyle.FontSize = 18
	subtitleStyle.Color = gf.NewColor(100, 100, 100, 255)
	canvas.DrawText("Multi-Page Demo with RELOADED Routing", gf.NewOffset(50, 95), subtitleStyle)

	// Draw page selection buttons
	y := 160.0
	buttons := []struct {
		label string
		page  PageType
		color struct{ r, g, b uint8 }
	}{
		{"Input Widgets Page →", PageInputWidgets, struct{ r, g, b uint8 }{66, 133, 244}},
		{"Forms & Text Input →", PageForms, struct{ r, g, b uint8 }{255, 152, 0}},
		{"Layout Demo Page →", PageLayoutDemo, struct{ r, g, b uint8 }{52, 168, 83}},
	}

	for i, btn := range buttons {
		drawButton(canvas, 50, y, 300, 60, btn.label, btn.color.r, btn.color.g, btn.color.b, i+1)
		y += 80
	}

	// Draw instructions
	infoStyle := gf.NewTextStyle()
	infoStyle.FontSize = 14
	infoStyle.Color = gf.NewColor(120, 120, 120, 255)
	canvas.DrawText("Click the buttons above to navigate to different pages", gf.NewOffset(50, 400), infoStyle)
	canvas.DrawText("Press ESC key to return to home from any page", gf.NewOffset(50, 425), infoStyle)
}

// drawInputWidgetsPage renders the input widgets demo page
func drawInputWidgetsPage(canvas *macos.CoreGraphicsCanvas) {
	// Draw header
	titleStyle := gf.NewTextStyle()
	titleStyle.FontSize = 28
	titleStyle.Color = gf.NewColor(33, 33, 33, 255)
	canvas.DrawText("Input Widgets Demo", gf.NewOffset(50, 50), titleStyle)

	// Back button
	drawButton(canvas, 50, 95, 150, 35, "← Back", 100, 100, 100, 10)

	labelStyle := gf.NewTextStyle()
	labelStyle.FontSize = 14
	labelStyle.Color = gf.NewColor(100, 100, 100, 255)

	sectionStyle := gf.NewTextStyle()
	sectionStyle.FontSize = 18
	sectionStyle.Color = gf.NewColor(50, 50, 50, 255)

	y := 160.0

	// BUTTONS SECTION
	canvas.DrawText("Buttons", gf.NewOffset(50, y), sectionStyle)
	canvas.DrawText("Click to change state", gf.NewOffset(50, y+20), labelStyle)
	y += 45

	// Show different button variants
	buttonConfigs := []struct {
		label   string
		variant ButtonVariant
		color   struct{ r, g, b uint8 }
	}{
		{"Elevated", ButtonElevated, struct{ r, g, b uint8 }{66, 133, 244}},
		{"Text", ButtonText, struct{ r, g, b uint8 }{66, 133, 244}},
		{"Outlined", ButtonOutlined, struct{ r, g, b uint8 }{66, 133, 244}},
	}

	for i, config := range buttonConfigs {
		color := config.color
		if appState.clickedButton == i {
			color = struct{ r, g, b uint8 }{25, 103, 210}
		}
		drawButtonWithVariant(canvas, 50, y, 150, 45, config.label, color.r, color.g, color.b, 20+i, config.variant)
		y += 55
	}

	// CHECKBOX SECTION
	y = 160.0
	x := 450.0
	canvas.DrawText("Checkbox", gf.NewOffset(x, y), sectionStyle)
	canvas.DrawText("Toggle on/off", gf.NewOffset(x, y+20), labelStyle)
	y += 45

	drawCheckbox(canvas, x, y, appState.checkboxState, 30)
	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 15
	textStyle.Color = gf.NewColor(33, 33, 33, 255)
	canvas.DrawText("Enable notifications", gf.NewOffset(x+35, y+15), textStyle)

	// RADIO BUTTONS SECTION
	y += 60
	canvas.DrawText("Radio Buttons", gf.NewOffset(x, y), sectionStyle)
	canvas.DrawText("Select one option", gf.NewOffset(x, y+20), labelStyle)
	y += 45

	radioOptions := []string{"Option A", "Option B", "Option C"}
	for i, option := range radioOptions {
		drawRadioButton(canvas, x, y, appState.radioSelection == i, 40+i)
		canvas.DrawText(option, gf.NewOffset(x+35, y+15), textStyle)
		y += 35
	}

	// SLIDER SECTION
	y += 20
	canvas.DrawText("Slider", gf.NewOffset(x, y), sectionStyle)
	canvas.DrawText("Adjust value (0-100)", gf.NewOffset(x, y+20), labelStyle)
	y += 45

	drawSlider(canvas, x, y, 250, appState.sliderValue, 31)
	canvas.DrawText(fmt.Sprintf("Value: %.0f%%", appState.sliderValue), gf.NewOffset(x, y+35), textStyle)

	// Status display
	statusY := 580.0
	statusStyle := gf.NewTextStyle()
	statusStyle.FontSize = 13
	statusStyle.Color = gf.NewColor(120, 120, 120, 255)
	canvas.DrawText("Widget State:", gf.NewOffset(50, statusY), sectionStyle)
	statusY += 25
	buttonLabels := []string{"Elevated", "Text", "Outlined"}
	canvas.DrawText(fmt.Sprintf("• Last button clicked: %s", buttonLabels[max(0, appState.clickedButton)]), gf.NewOffset(50, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Checkbox: %v", appState.checkboxState), gf.NewOffset(50, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Radio: %s", radioOptions[appState.radioSelection]), gf.NewOffset(50, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Slider: %.0f%%", appState.sliderValue), gf.NewOffset(50, statusY), statusStyle)
}

// drawLayoutDemoPage renders the layout demo page
func drawLayoutDemoPage(canvas *macos.CoreGraphicsCanvas) {
	// Draw header
	titleStyle := gf.NewTextStyle()
	titleStyle.FontSize = 32
	titleStyle.Color = gf.NewColor(33, 33, 33, 255)
	canvas.DrawText("Layout Demo", gf.NewOffset(50, 50), titleStyle)

	// Back button
	drawButton(canvas, 50, 100, 150, 40, "← Back to Home", 100, 100, 100, 10)

	// Row layout example
	labelStyle := gf.NewTextStyle()
	labelStyle.FontSize = 18
	labelStyle.Color = gf.NewColor(50, 50, 50, 255)
	canvas.DrawText("Row Layout:", gf.NewOffset(50, 170), labelStyle)

	colors := []struct{ r, g, b uint8 }{
		{244, 67, 54},  // Red
		{33, 150, 243}, // Blue
		{76, 175, 80},  // Green
	}
	x := 50.0
	for _, color := range colors {
		paint := gf.NewPaint()
		paint.Color = gf.NewColor(color.r, color.g, color.b, 255)
		rect := gf.NewRect(gf.NewOffset(x, 200), gf.NewSize(100, 100))
		canvas.DrawRect(rect, paint)
		x += 120
	}

	// Column layout example
	canvas.DrawText("Column Layout:", gf.NewOffset(50, 330), labelStyle)

	y := 360.0
	for _, color := range colors {
		paint := gf.NewPaint()
		paint.Color = gf.NewColor(color.r, color.g, color.b, 255)
		rect := gf.NewRect(gf.NewOffset(50, y), gf.NewSize(100, 50))
		canvas.DrawRect(rect, paint)
		y += 60
	}

	// Stack layout example
	canvas.DrawText("Stack Layout:", gf.NewOffset(450, 170), labelStyle)

	// Draw stacked rectangles (largest to smallest)
	stackColors := []struct {
		r, g, b uint8
		size    float64
	}{
		{156, 39, 176, 120}, // Purple - largest
		{103, 58, 183, 90},  // Deep Purple - medium
		{63, 81, 181, 60},   // Indigo - smallest
	}

	for _, item := range stackColors {
		paint := gf.NewPaint()
		paint.Color = gf.NewColor(item.r, item.g, item.b, 255)
		offset := (120 - item.size) / 2
		rect := gf.NewRect(gf.NewOffset(450+offset, 200+offset), gf.NewSize(item.size, item.size))
		canvas.DrawRect(rect, paint)
	}
}

// drawButton draws a button with the specified variant (defaults to Elevated)
func drawButton(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, label string, r, g, b uint8, id int) {
	drawButtonWithVariant(canvas, x, y, width, height, label, r, g, b, id, ButtonElevated)
}

// drawButtonWithVariant draws a button with the specified variant type
func drawButtonWithVariant(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, label string, r, g, b uint8, id int, variant ButtonVariant) {
	switch variant {
	case ButtonText:
		drawTextButton(canvas, x, y, width, height, label, r, g, b, id)
	case ButtonOutlined:
		drawOutlinedButton(canvas, x, y, width, height, label, r, g, b, id)
	default: // ButtonElevated
		drawElevatedButton(canvas, x, y, width, height, label, r, g, b, id)
	}
}

// drawElevatedButton implements Flutter-style Elevated button with proper 9-scale appearance
func drawElevatedButton(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, label string, r, g, b uint8, id int) {
	// Get current button state
	state := appState.buttonStates[id]
	if appState.hoveredButton == id {
		state = ButtonHovered
	}
	if appState.pressedButton == id {
		state = ButtonPressed
	}

	// Flutter Material Design elevation and colors
	var elevation float64 = 3 // Increased base elevation
	var finalR, finalG, finalB uint8
	var shadowAlpha uint8 = 100

	switch state {
	case ButtonHovered:
		// Material Design hover: slight overlay + increased elevation
		finalR = uint8(min(255, int(r)+15))
		finalG = uint8(min(255, int(g)+15))
		finalB = uint8(min(255, int(b)+15))
		elevation = 6 // Higher elevation on hover for more dramatic effect
		shadowAlpha = 140
	case ButtonPressed:
		// Material Design pressed: darker overlay + reduced elevation
		finalR = uint8(max(0, int(r)-20))
		finalG = uint8(max(0, int(g)-20))
		finalB = uint8(max(0, int(b)-20))
		elevation = 2 // Lower elevation when pressed
		shadowAlpha = 80
	default: // ButtonNormal
		finalR, finalG, finalB = r, g, b
	}

	// Draw multi-layer shadow for Material Design elevation effect
	drawMaterialShadow(canvas, x, y, width, height, elevation, shadowAlpha)

	// Draw rounded rectangle button (9-scale approach)
	cornerRadius := 8.0
	drawRoundedRect(canvas, x, y, width, height, cornerRadius, finalR, finalG, finalB, 255)

	// Draw subtle border for definition
	borderR := uint8(max(0, int(finalR)-30))
	borderG := uint8(max(0, int(finalG)-30))
	borderB := uint8(max(0, int(finalB)-30))
	drawRoundedRectBorder(canvas, x, y, width, height, cornerRadius, borderR, borderG, borderB, 80, 1.0)

	// Calculate text dimensions for perfect centering
	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 16
	textStyle.Color = gf.NewColor(255, 255, 255, 255)

	// Estimate text dimensions (more accurate measurement)
	textWidth := estimateTextWidth(label, textStyle.FontSize)
	textHeight := textStyle.FontSize

	// Center the text perfectly
	textX := x + (width-textWidth)/2
	textY := y + (height+textHeight)/2 - 2 // Slight adjustment for baseline

	// Apply pressed offset for tactile feedback
	if state == ButtonPressed {
		textX += 0.5
		textY += 0.5
	}

	canvas.DrawText(label, gf.NewOffset(textX, textY), textStyle)
}

// drawMaterialShadow creates layered shadows for Material Design elevation effect
func drawMaterialShadow(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, elevation float64, alpha uint8) {
	if elevation <= 0 {
		return // No shadow for zero elevation
	}

	// Material Design shadow layers based on elevation
	// Reference: https://material.io/design/environment/elevation.html
	shadowLayers := []struct {
		offsetX, offsetY float64
		blur             float64
		spread           float64
		alpha            uint8
		description      string
	}{
		// Umbra shadow (key light source)
		{0, elevation * 0.3, elevation * 1.0, elevation * 0.2, uint8(float64(alpha) * 0.2), "umbra"},
		// Penumbra shadow (ambient light)
		{0, elevation * 0.6, elevation * 1.8, elevation * 0.4, uint8(float64(alpha) * 0.14), "penumbra"},
		// Ambient shadow (general diffusion)
		{0, elevation * 1.0, elevation * 3.0, elevation * 0.1, uint8(float64(alpha) * 0.12), "ambient"},
	}

	cornerRadius := 8.0 // Same as button corner radius

	for _, shadow := range shadowLayers {
		// Draw multiple rectangles to simulate blur effect
		blurSteps := int(max(1, int(shadow.blur)))
		stepAlpha := shadow.alpha / uint8(blurSteps)

		for step := 0; step < blurSteps; step++ {
			blurOffset := float64(step) * 0.5
			shadowExpansion := shadow.spread + blurOffset

			shadowPaint := gf.NewPaint()
			shadowPaint.Color = gf.NewColor(0, 0, 0, stepAlpha)

			// Calculate shadow position and size
			shadowX := x + shadow.offsetX - shadowExpansion
			shadowY := y + shadow.offsetY - shadowExpansion
			shadowW := width + (shadowExpansion * 2)
			shadowH := height + (shadowExpansion * 2)

			// Draw rounded shadow using same approach as button
			drawRoundedRect(canvas, shadowX, shadowY, shadowW, shadowH, cornerRadius+shadowExpansion, 0, 0, 0, stepAlpha)
		}
	}
}

// drawRoundedRect draws a 9-scale rounded rectangle using corner composition
func drawRoundedRect(canvas *macos.CoreGraphicsCanvas, x, y, width, height, radius float64, r, g, b, a uint8) {
	paint := gf.NewPaint()
	paint.Color = gf.NewColor(r, g, b, a)

	// 9-scale approach: draw center rectangle and corner pieces

	// Center rectangle (main body)
	centerRect := gf.NewRect(
		gf.NewOffset(x+radius, y+radius),
		gf.NewSize(width-radius*2, height-radius*2),
	)
	canvas.DrawRect(centerRect, paint)

	// Top and bottom strips
	topRect := gf.NewRect(gf.NewOffset(x+radius, y), gf.NewSize(width-radius*2, radius))
	bottomRect := gf.NewRect(gf.NewOffset(x+radius, y+height-radius), gf.NewSize(width-radius*2, radius))
	canvas.DrawRect(topRect, paint)
	canvas.DrawRect(bottomRect, paint)

	// Left and right strips
	leftRect := gf.NewRect(gf.NewOffset(x, y+radius), gf.NewSize(radius, height-radius*2))
	rightRect := gf.NewRect(gf.NewOffset(x+width-radius, y+radius), gf.NewSize(radius, height-radius*2))
	canvas.DrawRect(leftRect, paint)
	canvas.DrawRect(rightRect, paint)

	// Corner circles for rounded effect
	canvas.DrawCircle(gf.NewOffset(x+radius, y+radius), radius, paint)              // Top-left
	canvas.DrawCircle(gf.NewOffset(x+width-radius, y+radius), radius, paint)        // Top-right
	canvas.DrawCircle(gf.NewOffset(x+radius, y+height-radius), radius, paint)       // Bottom-left
	canvas.DrawCircle(gf.NewOffset(x+width-radius, y+height-radius), radius, paint) // Bottom-right
}

// drawRoundedRectBorder draws a border around a rounded rectangle
func drawRoundedRectBorder(canvas *macos.CoreGraphicsCanvas, x, y, width, height, radius float64, r, g, b, a uint8, thickness float64) {
	// For now, draw a simple border approximation
	// In a full implementation, this would draw proper rounded borders
	borderPaint := gf.NewPaint()
	borderPaint.Color = gf.NewColor(r, g, b, a)

	// Top border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x+radius, y), gf.NewSize(width-radius*2, thickness)), borderPaint)
	// Bottom border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x+radius, y+height-thickness), gf.NewSize(width-radius*2, thickness)), borderPaint)
	// Left border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x, y+radius), gf.NewSize(thickness, height-radius*2)), borderPaint)
	// Right border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x+width-thickness, y+radius), gf.NewSize(thickness, height-radius*2)), borderPaint)
}

// estimateTextWidth provides a rough estimate of text width for centering
func estimateTextWidth(text string, fontSize float64) float64 {
	// Rough character width estimation (can be made more accurate with proper font metrics)
	avgCharWidth := fontSize * 0.55 // Approximation for typical fonts
	return float64(len(text)) * avgCharWidth
}

// drawTextButton implements Material Design TextButton (no background, text only)
func drawTextButton(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, label string, r, g, b uint8, id int) {
	// Get current button state
	state := appState.buttonStates[id]
	if appState.hoveredButton == id {
		state = ButtonHovered
	}
	if appState.pressedButton == id {
		state = ButtonPressed
	}

	// Text button only shows background on hover/press
	var bgAlpha uint8 = 0
	var textColorMultiplier float64 = 1.0

	switch state {
	case ButtonHovered:
		bgAlpha = 25              // Very subtle background on hover
		textColorMultiplier = 1.1 // Slightly brighter text
	case ButtonPressed:
		bgAlpha = 40              // Slightly more visible background when pressed
		textColorMultiplier = 0.9 // Slightly darker text when pressed
	}

	// Draw subtle background only on interaction
	if bgAlpha > 0 {
		cornerRadius := 8.0
		bgR := uint8(min(255, int(float64(r)*textColorMultiplier)))
		bgG := uint8(min(255, int(float64(g)*textColorMultiplier)))
		bgB := uint8(min(255, int(float64(b)*textColorMultiplier)))
		drawRoundedRect(canvas, x, y, width, height, cornerRadius, bgR, bgG, bgB, bgAlpha)
	}

	// Calculate text color based on background color (contrasting text)
	textR := uint8(min(255, int(float64(r)*textColorMultiplier)))
	textG := uint8(min(255, int(float64(g)*textColorMultiplier)))
	textB := uint8(min(255, int(float64(b)*textColorMultiplier)))

	// Calculate text dimensions for perfect centering
	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 16
	textStyle.Color = gf.NewColor(textR, textG, textB, 255)

	// Estimate text dimensions (more accurate measurement)
	textWidth := estimateTextWidth(label, textStyle.FontSize)
	textHeight := textStyle.FontSize

	// Center the text perfectly
	textX := x + (width-textWidth)/2
	textY := y + (height+textHeight)/2 - 2 // Slight adjustment for baseline

	// Apply pressed offset for tactile feedback
	if state == ButtonPressed {
		textX += 0.5
		textY += 0.5
	}

	canvas.DrawText(label, gf.NewOffset(textX, textY), textStyle)
}

// drawOutlinedButton implements Material Design OutlinedButton (border with no fill)
func drawOutlinedButton(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, label string, r, g, b uint8, id int) {
	// Get current button state
	state := appState.buttonStates[id]
	if appState.hoveredButton == id {
		state = ButtonHovered
	}
	if appState.pressedButton == id {
		state = ButtonPressed
	}

	// Outlined button colors
	var borderAlpha uint8 = 180
	var bgAlpha uint8 = 0
	var textColorMultiplier float64 = 1.0
	var borderWidth float64 = 1.5

	switch state {
	case ButtonHovered:
		bgAlpha = 15      // Very subtle background fill on hover
		borderAlpha = 220 // Stronger border on hover
		textColorMultiplier = 1.1
		borderWidth = 2.0
	case ButtonPressed:
		bgAlpha = 30      // Slightly more background when pressed
		borderAlpha = 255 // Full opacity border when pressed
		textColorMultiplier = 0.9
		borderWidth = 1.0 // Thinner border when pressed for pressed effect
	}

	cornerRadius := 8.0

	// Draw background only on interaction
	if bgAlpha > 0 {
		bgR := uint8(min(255, int(float64(r)*textColorMultiplier)))
		bgG := uint8(min(255, int(float64(g)*textColorMultiplier)))
		bgB := uint8(min(255, int(float64(b)*textColorMultiplier)))
		drawRoundedRect(canvas, x, y, width, height, cornerRadius, bgR, bgG, bgB, bgAlpha)
	}

	// Draw border
	borderR := uint8(min(255, int(float64(r)*textColorMultiplier)))
	borderG := uint8(min(255, int(float64(g)*textColorMultiplier)))
	borderB := uint8(min(255, int(float64(b)*textColorMultiplier)))
	drawRoundedRectBorder(canvas, x, y, width, height, cornerRadius, borderR, borderG, borderB, borderAlpha, borderWidth)

	// Calculate text color
	textR := uint8(min(255, int(float64(r)*textColorMultiplier)))
	textG := uint8(min(255, int(float64(g)*textColorMultiplier)))
	textB := uint8(min(255, int(float64(b)*textColorMultiplier)))

	// Calculate text dimensions for perfect centering
	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 16
	textStyle.Color = gf.NewColor(textR, textG, textB, 255)

	// Estimate text dimensions (more accurate measurement)
	textWidth := estimateTextWidth(label, textStyle.FontSize)
	textHeight := textStyle.FontSize

	// Center the text perfectly
	textX := x + (width-textWidth)/2
	textY := y + (height+textHeight)/2 - 2 // Slight adjustment for baseline

	// Apply pressed offset for tactile feedback
	if state == ButtonPressed {
		textX += 0.5
		textY += 0.5
	}

	canvas.DrawText(label, gf.NewOffset(textX, textY), textStyle)
}

// drawCheckbox draws a checkbox
func drawCheckbox(canvas *macos.CoreGraphicsCanvas, x, y float64, checked bool, id int) {
	// Draw box background
	paint := gf.NewPaint()
	paint.Color = gf.NewColor(230, 230, 230, 255)
	rect := gf.NewRect(gf.NewOffset(x, y), gf.NewSize(24, 24))
	canvas.DrawRect(rect, paint)

	// Draw checkmark if checked
	if checked {
		fillPaint := gf.NewPaint()
		fillPaint.Color = gf.NewColor(66, 133, 244, 255)
		innerRect := gf.NewRect(gf.NewOffset(x+4, y+4), gf.NewSize(16, 16))
		canvas.DrawRect(innerRect, fillPaint)
	}
}

// drawSlider draws a slider control
func drawSlider(canvas *macos.CoreGraphicsCanvas, x, y, width, value float64, id int) {
	// Draw track
	trackPaint := gf.NewPaint()
	trackPaint.Color = gf.NewColor(200, 200, 200, 255)
	trackRect := gf.NewRect(gf.NewOffset(x, y+8), gf.NewSize(width, 4))
	canvas.DrawRect(trackRect, trackPaint)

	// Draw filled portion
	fillWidth := (value / 100.0) * width
	fillPaint := gf.NewPaint()
	fillPaint.Color = gf.NewColor(66, 133, 244, 255)
	fillRect := gf.NewRect(gf.NewOffset(x, y+8), gf.NewSize(fillWidth, 4))
	canvas.DrawRect(fillRect, fillPaint)

	// Draw thumb
	thumbX := x + fillWidth - 10
	thumbPaint := gf.NewPaint()
	thumbPaint.Color = gf.NewColor(66, 133, 244, 255)
	canvas.DrawCircle(gf.NewOffset(thumbX+10, y+10), 10, thumbPaint)
}

// drawRadioButton draws a radio button
func drawRadioButton(canvas *macos.CoreGraphicsCanvas, x, y float64, selected bool, id int) {
	// Draw outer circle
	outerPaint := gf.NewPaint()
	outerPaint.Color = gf.NewColor(200, 200, 200, 255)
	canvas.DrawCircle(gf.NewOffset(x+12, y+12), 12, outerPaint)

	// Draw inner circle if selected
	if selected {
		innerPaint := gf.NewPaint()
		innerPaint.Color = gf.NewColor(66, 133, 244, 255)
		canvas.DrawCircle(gf.NewOffset(x+12, y+12), 7, innerPaint)
	} else {
		// Draw white center for unselected
		centerPaint := gf.NewPaint()
		centerPaint.Color = gf.NewColor(255, 255, 255, 255)
		canvas.DrawCircle(gf.NewOffset(x+12, y+12), 9, centerPaint)
	}
}

// drawFormsPage renders the forms and text input demo page
func drawFormsPage(canvas *macos.CoreGraphicsCanvas) {
	// Draw header
	titleStyle := gf.NewTextStyle()
	titleStyle.FontSize = 28
	titleStyle.Color = gf.NewColor(33, 33, 33, 255)
	canvas.DrawText("Forms & Text Input Demo", gf.NewOffset(50, 50), titleStyle)

	// Back button
	drawButton(canvas, 50, 95, 150, 35, "← Back", 100, 100, 100, 10)

	labelStyle := gf.NewTextStyle()
	labelStyle.FontSize = 14
	labelStyle.Color = gf.NewColor(100, 100, 100, 255)

	sectionStyle := gf.NewTextStyle()
	sectionStyle.FontSize = 18
	sectionStyle.Color = gf.NewColor(50, 50, 50, 255)

	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 15
	textStyle.Color = gf.NewColor(33, 33, 33, 255)

	y := 160.0

	// TEXT INPUT SECTION
	canvas.DrawText("Text Input Fields", gf.NewOffset(50, y), sectionStyle)
	canvas.DrawText("Static text input demonstration", gf.NewOffset(50, y+20), labelStyle)
	y += 45

	// Name field
	canvas.DrawText("Name:", gf.NewOffset(50, y), textStyle)
	y += 25
	drawTextInput(canvas, 50, y, 300, 35, appState.textInput1, 50)
	y += 50

	// Email field (from form)
	canvas.DrawText("Message:", gf.NewOffset(50, y), textStyle)
	y += 25
	drawTextInput(canvas, 50, y, 300, 35, appState.textInput2, 51)
	y += 60

	// FORM SECTION
	canvas.DrawText("Form Example", gf.NewOffset(50, y), sectionStyle)
	canvas.DrawText("Sample user registration form", gf.NewOffset(50, y+20), labelStyle)
	y += 45

	// Form Name
	canvas.DrawText("Full Name:", gf.NewOffset(50, y), textStyle)
	y += 25
	drawTextInput(canvas, 50, y, 300, 35, appState.formName, 52)
	y += 50

	// Form Email
	canvas.DrawText("Email Address:", gf.NewOffset(50, y), textStyle)
	y += 25
	drawTextInput(canvas, 50, y, 300, 35, appState.formEmail, 53)
	y += 50

	// Submit button
	drawButton(canvas, 50, y, 150, 40, "Submit Form", 76, 175, 80, 60)

	// Form data display
	statusY := 580.0
	canvas.DrawText("Current Form Values:", gf.NewOffset(450, statusY), sectionStyle)
	statusY += 25

	statusStyle := gf.NewTextStyle()
	statusStyle.FontSize = 13
	statusStyle.Color = gf.NewColor(120, 120, 120, 255)

	canvas.DrawText(fmt.Sprintf("• Text Input 1: %s", appState.textInput1), gf.NewOffset(450, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Text Input 2: %s", appState.textInput2), gf.NewOffset(450, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Form Name: %s", appState.formName), gf.NewOffset(450, statusY), statusStyle)
	statusY += 20
	canvas.DrawText(fmt.Sprintf("• Form Email: %s", appState.formEmail), gf.NewOffset(450, statusY), statusStyle)
}

// drawTextInput draws an interactive text input field with visual feedback
func drawTextInput(canvas *macos.CoreGraphicsCanvas, x, y, width, height float64, value string, id int) {
	isActive := appState.activeTextInput == id
	isHovered := appState.hoveredTextInput == id

	// Draw background with different color based on state
	bgPaint := gf.NewPaint()
	if isActive {
		bgPaint.Color = gf.NewColor(248, 250, 255, 255) // Light blue background when active
	} else if isHovered {
		bgPaint.Color = gf.NewColor(250, 250, 250, 255) // Light gray background when hovered
	} else {
		bgPaint.Color = gf.NewColor(255, 255, 255, 255) // White background when inactive
	}
	bgRect := gf.NewRect(gf.NewOffset(x, y), gf.NewSize(width, height))
	canvas.DrawRect(bgRect, bgPaint)

	// Draw border with different thickness and color based on state
	borderPaint := gf.NewPaint()
	borderThickness := 1.0
	if isActive {
		borderPaint.Color = gf.NewColor(66, 133, 244, 255) // Blue border when active
		borderThickness = 2.0
	} else if isHovered {
		borderPaint.Color = gf.NewColor(150, 150, 150, 255) // Darker gray border when hovered
		borderThickness = 1.5
	} else {
		borderPaint.Color = gf.NewColor(200, 200, 200, 255) // Gray border when inactive
	}

	// Top border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x, y), gf.NewSize(width, borderThickness)), borderPaint)
	// Bottom border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x, y+height-borderThickness), gf.NewSize(width, borderThickness)), borderPaint)
	// Left border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x, y), gf.NewSize(borderThickness, height)), borderPaint)
	// Right border
	canvas.DrawRect(gf.NewRect(gf.NewOffset(x+width-borderThickness, y), gf.NewSize(borderThickness, height)), borderPaint)

	// Draw placeholder text or actual value
	textStyle := gf.NewTextStyle()
	textStyle.FontSize = 14

	var displayText string
	if value == "" {
		// Show placeholder text
		displayText = getPlaceholderText(id)
		textStyle.Color = gf.NewColor(150, 150, 150, 255) // Gray placeholder
	} else {
		displayText = value
		textStyle.Color = gf.NewColor(33, 33, 33, 255) // Black text
	}

	// Draw the text with proper vertical centering
	textX := x + 12
	textY := y + height/2 - 2 // Better vertical centering
	canvas.DrawText(displayText, gf.NewOffset(textX, textY), textStyle)

	// Draw blinking cursor if this field is active
	if isActive {
		drawTextCursor(canvas, textX, textY, value, textStyle)
	}
}

// getPlaceholderText returns appropriate placeholder text for each field
func getPlaceholderText(id int) string {
	switch id {
	case 50: // Name field
		return "Enter your name..."
	case 51: // Message field
		return "Enter your message..."
	case 52: // Form Name field
		return "Enter full name..."
	case 53: // Form Email field
		return "Enter email address..."
	default:
		return "Type here..."
	}
}

// drawTextCursor draws a blinking cursor at the end of the text
func drawTextCursor(canvas *macos.CoreGraphicsCanvas, textX, textY float64, text string, textStyle *gf.TextStyle) {
	// Calculate cursor position at the end of text
	// For simplicity, estimate character width (this could be more precise)
	charWidth := 8.0
	cursorX := textX + float64(len(text))*charWidth

	// Implement blinking animation (blink every 500ms)
	currentTime := time.Now().UnixMilli()
	shouldShowCursor := (currentTime/500)%2 == 0

	if shouldShowCursor {
		// Draw cursor line with better positioning
		cursorPaint := gf.NewPaint()
		cursorPaint.Color = gf.NewColor(33, 33, 33, 255) // Black cursor
		cursorRect := gf.NewRect(gf.NewOffset(cursorX, textY-10), gf.NewSize(1, 14))
		canvas.DrawRect(cursorRect, cursorPaint)
	}
}

// getButtonIdAtPosition returns the button ID at the given position, or 0 if none
func getButtonIdAtPosition(x, y float64) int {
	switch appState.currentPage {
	case PageHome:
		// Check home page buttons (3 buttons at y=160, 240, 320)
		if isInRect(x, y, 50, 160, 300, 60) {
			return 1 // Input Widgets
		} else if isInRect(x, y, 50, 240, 300, 60) {
			return 2 // Forms
		} else if isInRect(x, y, 50, 320, 300, 60) {
			return 3 // Layout Demo
		}

	case PageInputWidgets:
		// Back button
		if isInRect(x, y, 50, 95, 150, 35) {
			return 10
		}
		// Check button clicks (3 buttons starting at y=205)
		for i := 0; i < 3; i++ {
			btnY := 205.0 + float64(i)*55
			if isInRect(x, y, 50, btnY, 150, 45) {
				return 20 + i
			}
		}

	case PageLayoutDemo:
		// Back button
		if isInRect(x, y, 50, 100, 150, 40) {
			return 10
		}

	case PageForms:
		// Back button
		if isInRect(x, y, 50, 95, 150, 35) {
			return 10
		}
		// Submit button (y=550)
		if isInRect(x, y, 50, 550, 150, 40) {
			return 60
		}
	}
	return 0
}

// handleMouseDown processes mouse down events
func handleMouseDown(x, y float64) {
	fmt.Printf("Mouse down at (%.0f, %.0f) - Page: %d\n", x, y, appState.currentPage)

	// Check if clicking on a button
	buttonId := getButtonIdAtPosition(x, y)
	if buttonId > 0 {
		appState.pressedButton = buttonId
		fmt.Printf("Button %d pressed\n", buttonId)
		return
	}

	// Original click handling logic for non-button elements
	handleClickLogic(x, y)
}

// handleMouseUp processes mouse up events
func handleMouseUp(x, y float64) {
	fmt.Printf("Mouse up at (%.0f, %.0f) - Page: %d\n", x, y, appState.currentPage)

	// If there was a pressed button, handle the click action
	if appState.pressedButton > 0 {
		buttonId := getButtonIdAtPosition(x, y)
		if buttonId == appState.pressedButton {
			// Valid button click (mouse up on same button as mouse down)
			handleButtonClick(buttonId)
		}
		appState.pressedButton = 0 // Clear pressed state
	}
}

// handleButtonClick processes button click actions
func handleButtonClick(buttonId int) {
	fmt.Printf("Button %d clicked!\n", buttonId)

	switch buttonId {
	case 1: // Home -> Input Widgets
		appState.currentPage = PageInputWidgets
		fmt.Println("→ Navigating to Input Widgets page")
	case 2: // Home -> Forms
		appState.currentPage = PageForms
		fmt.Println("→ Navigating to Forms page")
	case 3: // Home -> Layout Demo
		appState.currentPage = PageLayoutDemo
		fmt.Println("→ Navigating to Layout Demo page")
	case 10: // Back button
		appState.currentPage = PageHome
		fmt.Println("← Navigating back to Home")
	case 20, 21, 22: // Input widget buttons
		appState.clickedButton = buttonId - 20
		fmt.Printf("Widget button %d clicked\n", appState.clickedButton+1)
	case 60: // Submit button
		submitForm()
	}
}

// handleClickLogic processes non-button click logic (widgets, text inputs, etc.)
func handleClickLogic(x, y float64) {
	fmt.Printf("Click at (%.0f, %.0f) - Page: %d\n", x, y, appState.currentPage)

	switch appState.currentPage {
	case PageHome:
		// Check home page buttons (3 buttons at y=160, 240, 320)
		if isInRect(x, y, 50, 160, 300, 60) {
			appState.currentPage = PageInputWidgets
			fmt.Println("→ Navigating to Input Widgets page")
		} else if isInRect(x, y, 50, 240, 300, 60) {
			appState.currentPage = PageForms
			fmt.Println("→ Navigating to Forms page")
		} else if isInRect(x, y, 50, 320, 300, 60) {
			appState.currentPage = PageLayoutDemo
			fmt.Println("→ Navigating to Layout Demo page")
		}

	case PageInputWidgets:
		// Back button
		if isInRect(x, y, 50, 95, 150, 35) {
			appState.currentPage = PageHome
			fmt.Println("← Navigating back to Home")
			return
		}
		// Check button clicks (3 buttons starting at y=205)
		for i := 0; i < 3; i++ {
			btnY := 205.0 + float64(i)*55
			if isInRect(x, y, 50, btnY, 150, 45) {
				appState.clickedButton = i
				fmt.Printf("Button %d clicked\n", i+1)
			}
		}
		// Check checkbox (at x=450, y=205)
		if isInRect(x, y, 450, 205, 24, 24) {
			appState.checkboxState = !appState.checkboxState
			fmt.Printf("Checkbox toggled: %v\n", appState.checkboxState)
		}
		// Check radio buttons (3 radio buttons starting at y=310)
		for i := 0; i < 3; i++ {
			radioY := 310.0 + float64(i)*35
			if isInRect(x, y, 450, radioY, 24, 24) {
				appState.radioSelection = i
				fmt.Printf("Radio button %d selected\n", i)
			}
		}
		// Check slider (at y=485 approximately)
		if isInRect(x, y, 450, 475, 250, 30) {
			relX := x - 450
			appState.sliderValue = (relX / 250.0) * 100
			if appState.sliderValue < 0 {
				appState.sliderValue = 0
			}
			if appState.sliderValue > 100 {
				appState.sliderValue = 100
			}
			fmt.Printf("Slider moved to: %.0f%%\n", appState.sliderValue)
		}

	case PageLayoutDemo:
		// Back button
		if isInRect(x, y, 50, 100, 150, 40) {
			appState.currentPage = PageHome
			fmt.Println("← Navigating back to Home")
		}

	case PageForms:
		// Back button
		if isInRect(x, y, 50, 95, 150, 35) {
			appState.currentPage = PageHome
			fmt.Println("← Navigating back to Home")
			return
		}

		// Text Input 1 - Name (y=220)
		if isInRect(x, y, 50, 220, 300, 35) {
			fmt.Println("📝 Clicked on Name input")
			appState.activeTextInput = 1
		}

		// Text Input 2 - Message (y=295)
		if isInRect(x, y, 50, 295, 300, 35) {
			fmt.Println("📝 Clicked on Message input")
			appState.activeTextInput = 2
		}

		// Form Name input (y=425)
		if isInRect(x, y, 50, 425, 300, 35) {
			fmt.Println("📝 Clicked on Form Name input")
			appState.activeTextInput = 3
		}

		// Form Email input (y=500)
		if isInRect(x, y, 50, 500, 300, 35) {
			fmt.Println("📝 Clicked on Form Email input")
			appState.activeTextInput = 4
		}

		// Submit button (y=550)
		if isInRect(x, y, 50, 550, 150, 40) {
			submitForm()
		}

		// Clear active text input if clicking outside any text field
		if !isClickingAnyTextField(x, y) {
			appState.activeTextInput = 0
		}
	}
}

// handleMouseMove processes mouse movement for hover effects
func handleMouseMove(x, y float64) {
	// Handle button hover states
	oldHoveredButton := appState.hoveredButton
	appState.hoveredButton = getButtonIdAtPosition(x, y)

	if oldHoveredButton != appState.hoveredButton {
		fmt.Printf("Button hover changed: %d -> %d\n", oldHoveredButton, appState.hoveredButton)
	}

	// Handle text input hover states
	if appState.currentPage == PageForms {
		// Check if hovering over any text field
		oldHovered := appState.hoveredTextInput
		appState.hoveredTextInput = 0

		if isInRect(x, y, 50, 220, 300, 35) { // Name input
			appState.hoveredTextInput = 1
		} else if isInRect(x, y, 50, 295, 300, 35) { // Message input
			appState.hoveredTextInput = 2
		} else if isInRect(x, y, 50, 425, 300, 35) { // Form Name input
			appState.hoveredTextInput = 3
		} else if isInRect(x, y, 50, 500, 300, 35) { // Form Email input
			appState.hoveredTextInput = 4
		}

		// Only log if hover state changed
		if oldHovered != appState.hoveredTextInput {
			fmt.Printf("Text input hover changed: %d -> %d\n", oldHovered, appState.hoveredTextInput)
		}
	}
}

// handleTextInput processes keyboard input for text fields
func handleTextInput(key int, window *macos.Window) {
	if appState.activeTextInput == 0 {
		return
	}

	// Handle backspace (key 51)
	if key == 51 {
		switch appState.activeTextInput {
		case 1:
			if len(appState.textInput1) > 0 {
				appState.textInput1 = appState.textInput1[:len(appState.textInput1)-1]
			}
		case 2:
			if len(appState.textInput2) > 0 {
				appState.textInput2 = appState.textInput2[:len(appState.textInput2)-1]
			}
		case 3:
			if len(appState.formName) > 0 {
				appState.formName = appState.formName[:len(appState.formName)-1]
			}
		case 4:
			if len(appState.formEmail) > 0 {
				appState.formEmail = appState.formEmail[:len(appState.formEmail)-1]
			}
		}
		window.SetNeedsDisplay()
		return
	}

	// Convert key code to character
	char := keyCodeToChar(key)
	if char == "" {
		return
	}

	// Add character to the appropriate text field
	switch appState.activeTextInput {
	case 1:
		appState.textInput1 += char
		fmt.Printf("Text Input 1: '%s'\n", appState.textInput1)
	case 2:
		appState.textInput2 += char
		fmt.Printf("Text Input 2: '%s'\n", appState.textInput2)
	case 3:
		appState.formName += char
		fmt.Printf("Form Name: '%s'\n", appState.formName)
	case 4:
		appState.formEmail += char
		fmt.Printf("Form Email: '%s'\n", appState.formEmail)
	}

	window.SetNeedsDisplay()
}

// keyCodeToChar converts macOS key codes to characters
func keyCodeToChar(key int) string {
	// Map of common key codes to characters (macOS US keyboard layout)
	keyMap := map[int]string{
		// Letters
		0: "a", 1: "s", 2: "d", 3: "f", 4: "h", 5: "g", 6: "z", 7: "x", 8: "c", 9: "v",
		11: "b", 12: "q", 13: "w", 14: "e", 15: "r", 16: "y", 17: "t", 31: "o", 32: "u",
		34: "i", 35: "p", 37: "l", 38: "j", 40: "k", 45: "n", 46: "m",
		// Numbers
		18: "1", 19: "2", 20: "3", 21: "4", 23: "5", 22: "6", 26: "7", 28: "8", 25: "9", 29: "0",
		// Space
		49: " ",
		// Common symbols
		39: "'", 41: ";", 43: ",", 47: ".", 44: "/", 50: "`",
		24: "=", 27: "-", 33: "[", 30: "]", 42: "\\",
	}

	char, exists := keyMap[key]
	if exists {
		return char
	}
	return ""
}

// submitForm handles form submission with validation
func submitForm() {
	fmt.Println("\n🚀 Form Submission Attempted!")
	fmt.Println("=====================================")

	// Validation
	hasErrors := false

	if appState.formName == "" {
		fmt.Println("❌ Validation Error: Full Name is required")
		hasErrors = true
	}

	if appState.formEmail == "" {
		fmt.Println("❌ Validation Error: Email Address is required")
		hasErrors = true
	} else if !isValidEmail(appState.formEmail) {
		fmt.Println("❌ Validation Error: Email format is invalid")
		hasErrors = true
	}

	if hasErrors {
		fmt.Println("\n💡 Please fix the errors above and try again")
		return
	}

	// Success
	fmt.Println("✅ Form submitted successfully!")
	fmt.Printf("📝 Name: %s\n", appState.formName)
	fmt.Printf("📧 Email: %s\n", appState.formEmail)
	fmt.Printf("💭 Text Input 1: %s\n", appState.textInput1)
	fmt.Printf("💬 Text Input 2: %s\n", appState.textInput2)
	fmt.Println("=====================================")
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	// Simple email validation - contains @ and .
	hasAt := false
	hasDot := false
	for _, char := range email {
		if char == '@' {
			hasAt = true
		} else if char == '.' {
			hasDot = true
		}
	}
	return hasAt && hasDot && len(email) > 5
}

// isClickingAnyTextField checks if the click is within any text field
func isClickingAnyTextField(x, y float64) bool {
	return isInRect(x, y, 50, 220, 300, 35) || // Name input
		isInRect(x, y, 50, 295, 300, 35) || // Message input
		isInRect(x, y, 50, 425, 300, 35) || // Form Name input
		isInRect(x, y, 50, 500, 300, 35) // Form Email input
}

// isInRect checks if a point is inside a rectangle
func isInRect(x, y, rectX, rectY, width, height float64) bool {
	return x >= rectX && x <= rectX+width && y >= rectY && y <= rectY+height
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// createPlaceholderPage creates a simple placeholder page (kept for compatibility)
func createPlaceholderPage(title, message string) gf.Widget {
	return &PlaceholderPage{
		Title:   title,
		Message: message,
	}
}

type PlaceholderPage struct {
	gf.BaseWidget
	Title   string
	Message string
}

func (p *PlaceholderPage) Build(context gf.BuildContext) gf.Widget {
	titleWidget := gf.Text{Data: p.Title}
	messageWidget := gf.Text{Data: p.Message}

	backButton := gf.Text{Data: "← Back"}
	backContainer := gf.Container{
		Padding: gf.NewEdgeInsets(8, 16, 8, 16),
		Color:   gf.NewColor(200, 200, 200, 255),
		Child:   &backButton,
	}

	backBtn := gf.GestureDetector{
		Child: &backContainer,
		OnTap: func() {
			gf.Get.Back()
		},
	}

	spacer1 := 20.0
	spacer2 := 40.0

	content := gf.Column{
		Children: []gf.Widget{
			&backBtn,
			&gf.SizedBox{Height: &spacer1},
			&titleWidget,
			&gf.SizedBox{Height: &spacer2},
			&messageWidget,
		},
	}

	container := gf.Container{
		Padding: gf.NewEdgeInsets(20, 20, 20, 20),
		Child:   &content,
	}

	return &container
}
