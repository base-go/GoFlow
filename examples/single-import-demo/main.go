package main

import (
	"fmt"

	gf "github.com/base-go/GoFlow"
)

func main() {
	fmt.Println("=== GoFlow Single Import Demo ===")
	fmt.Println("This example demonstrates the new single-import architecture.")
	fmt.Println()

	// ========================================
	// Reactive State with Signals
	// ========================================
	fmt.Println("--- Reactive State (Signals) ---")

	// Create signals using the unified import
	count := gf.CreateSignal(0)
	name := gf.CreateSignal("GoFlow")

	// Computed values
	doubled := gf.CreateComputed(func() int {
		return count.Get() * 2
	})

	greeting := gf.CreateComputed(func() string {
		return fmt.Sprintf("Hello from %s! Count: %d", name.Get(), count.Get())
	})

	// Effects
	dispose := gf.CreateEffect(func() {
		fmt.Printf("  Effect triggered: %s\n", greeting.Get())
	})
	defer dispose()

	// Update state
	count.Set(5)
	name.Set("GoFlow 2.0")

	// Batch updates
	fmt.Println("\n--- Batch Updates ---")
	gf.Batch(func() {
		count.Set(10)
		name.Set("Single Import Architecture")
	})

	fmt.Printf("  Final doubled value: %d\n", doubled.Get())

	// ========================================
	// Colors and Styling
	// ========================================
	fmt.Println("\n--- Colors & Styling ---")

	primaryColor := gf.ColorBlue
	secondaryColor := gf.NewColor(255, 100, 50, 255)

	fmt.Printf("  Primary color: %v\n", primaryColor)
	fmt.Printf("  Secondary color: %v\n", secondaryColor)

	textStyle := gf.NewTextStyle(
		"Arial",     // font
		16.0,        // size
		primaryColor, // color
		false,       // bold
		false,       // italic
	)
	fmt.Printf("  Text style created: font=%s, size=%.1f\n", "Arial", 16.0)

	// ========================================
	// Geometry Types
	// ========================================
	fmt.Println("\n--- Geometry Types ---")

	size := gf.NewSize(800, 600)
	offset := gf.NewOffset(100, 50)
	rect := gf.NewRect(offset.X, offset.Y, size.Width, size.Height)

	fmt.Printf("  Window size: %dx%d\n", int(size.Width), int(size.Height))
	fmt.Printf("  Position offset: (%.0f, %.0f)\n", offset.X, offset.Y)
	fmt.Printf("  Bounding rect: x=%.0f, y=%.0f, w=%.0f, h=%.0f\n",
		rect.Left, rect.Top, rect.Width(), rect.Height())

	// ========================================
	// Platform Detection
	// ========================================
	fmt.Println("\n--- Platform Detection ---")

	platform := gf.DetectPlatform()
	fmt.Printf("  Detected platform: %s\n", platform)

	platformNames := map[string]string{
		gf.PlatformLinux:   "Linux",
		gf.PlatformMacOS:   "macOS",
		gf.PlatformWindows: "Windows",
		gf.PlatformWeb:     "Web",
		gf.PlatformAndroid: "Android",
		gf.PlatformIOS:     "iOS",
	}

	if name, ok := platformNames[platform]; ok {
		fmt.Printf("  Running on: %s\n", name)
	}

	// ========================================
	// Widget Type Demonstration
	// ========================================
	fmt.Println("\n--- Widget Types Available ---")
	fmt.Println("  All widgets accessible via 'gf' namespace:")
	fmt.Println("  - Layout: gf.Column, gf.Row, gf.Stack, gf.Container")
	fmt.Println("  - Display: gf.Text, gf.Icon, gf.Image, gf.Card")
	fmt.Println("  - Input: gf.TextField, gf.Checkbox, gf.Slider, gf.Switch")
	fmt.Println("  - Animation: gf.AnimatedContainer, gf.FadeTransition")
	fmt.Println("  - Material: gf.MaterialButton, gf.MaterialScaffold")
	fmt.Println("  - Cupertino: gf.CupertinoButton, gf.CupertinoNavigationBar")
	fmt.Println("  - Navigation: gf.Get, gf.ShowDialog, gf.ShowSnackbar")

	// ========================================
	// Layout Constants
	// ========================================
	fmt.Println("\n--- Layout Constants ---")
	fmt.Printf("  MainAxis alignments: Start=%d, Center=%d, End=%d\n",
		gf.MainAxisStart, gf.MainAxisCenter, gf.MainAxisEnd)
	fmt.Printf("  CrossAxis alignments: Start=%d, Center=%d, Stretch=%d\n",
		gf.CrossAxisStart, gf.CrossAxisCenter, gf.CrossAxisStretch)

	// ========================================
	// Summary
	// ========================================
	fmt.Println("\n=== Summary ===")
	fmt.Println("✓ Single import path: import gf \"github.com/base-go/GoFlow\"")
	fmt.Println("✓ All framework features accessible through 'gf' namespace")
	fmt.Println("✓ Clean, intuitive API surface")
	fmt.Println("✓ Easy discoverability via IDE autocomplete")
	fmt.Println()
	fmt.Println("Compare this to the old multi-import approach:")
	fmt.Println("  ❌ OLD: import \"github.com/base-go/GoFlow/pkg/core/signals\"")
	fmt.Println("  ❌ OLD: import \"github.com/base-go/GoFlow/pkg/core/widgets\"")
	fmt.Println("  ❌ OLD: import \"github.com/base-go/GoFlow/pkg/core/framework\"")
	fmt.Println()
	fmt.Println("  ✅ NEW: import gf \"github.com/base-go/GoFlow\"")
	fmt.Println()
}
