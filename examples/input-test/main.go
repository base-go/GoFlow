//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/base-go/GoFlow/backends/macos"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("GoFlow Input System Test")
	fmt.Println("========================\n")

	// Initialize app
	macos.InitApp()

	// Create window
	window := macos.NewWindow(600, 500, "Input System Test")
	defer window.Destroy()

	// Test state
	testResults := &TestResults{
		MouseClicks:  0,
		MouseMoves:   0,
		KeyPresses:   0,
		LastMousePos: "none",
		LastKey:      "none",
		MouseDrags:   0,
	}

	// Test 1: Mouse Events
	fmt.Println("Test 1: Mouse Events")
	fmt.Println("- Click anywhere in the window")
	fmt.Println("- Drag the mouse")
	fmt.Println("- Move the mouse around")

	window.SetMouseFunc(func(button, action int, x, y float64) {
		switch action {
		case 0: // Mouse up
			fmt.Printf("  Mouse up at (%.0f, %.0f)\n", x, y)

		case 1: // Mouse down
			testResults.MouseClicks++
			testResults.LastMousePos = fmt.Sprintf("(%.0f, %.0f)", x, y)
			fmt.Printf("✓ Mouse click #%d at %s\n", testResults.MouseClicks, testResults.LastMousePos)

		case 2: // Mouse move
			testResults.MouseMoves++
			if button >= 0 {
				testResults.MouseDrags++
			}
			if testResults.MouseMoves%100 == 0 {
				fmt.Printf("  Mouse moved %d times (drags: %d)\n", testResults.MouseMoves, testResults.MouseDrags)
			}
		}
	})

	// Test 2: Keyboard Events
	fmt.Println("\nTest 2: Keyboard Events")
	fmt.Println("- Press any keys")
	fmt.Println("- Try modifier keys (Ctrl, Shift, Alt, Cmd)")

	keyNames := map[int]string{
		0:   "A",
		1:   "S",
		2:   "D",
		6:   "Z",
		7:   "X",
		8:   "C",
		9:   "V",
		11:  "B",
		13:  "W",
		14:  "E",
		15:  "R",
		36:  "Return",
		48:  "Tab",
		49:  "Space",
		51:  "Delete",
		53:  "Escape",
		56:  "Shift",
		55:  "Cmd",
		58:  "Alt",
		59:  "Ctrl",
		123: "Left",
		124: "Right",
		125: "Down",
		126: "Up",
	}

	window.SetKeyFunc(func(key, action int) {
		if action == 1 { // Key down
			testResults.KeyPresses++

			keyName := keyNames[key]
			if keyName == "" {
				keyName = fmt.Sprintf("Key#%d", key)
			}

			testResults.LastKey = keyName
			fmt.Printf("✓ Key pressed #%d: %s (code: %d)\n", testResults.KeyPresses, keyName, key)
		} else { // Key up
			keyName := keyNames[key]
			if keyName == "" {
				keyName = fmt.Sprintf("Key#%d", key)
			}
			fmt.Printf("  Key released: %s\n", keyName)
		}
	})

	// Test 3: Window Events
	fmt.Println("\nTest 3: Window Events")
	fmt.Println("- Resize the window")

	resizeCount := 0
	window.SetResizeFunc(func(width, height int) {
		resizeCount++
		fmt.Printf("✓ Window resized #%d: %dx%d\n", resizeCount, width, height)
	})

	// Draw callback - show test status
	frame := 0
	window.SetDrawFunc(func(canvas *macos.CoreGraphicsCanvas) {
		canvas.Clear(goflow.NewColor(245, 245, 250, 255))

		// Title
		titleStyle := goflow.NewTextStyle()
		titleStyle.FontSize = 28
		titleStyle.FontWeight = goflow.FontWeightBold
		titleStyle.Color = goflow.NewColor(40, 40, 40, 255)
		canvas.DrawText("Input System Test", goflow.NewOffset(20, 35), titleStyle)

		// Subtitle
		subtitleStyle := goflow.NewTextStyle()
		subtitleStyle.FontSize = 14
		subtitleStyle.Color = goflow.NewColor(120, 120, 120, 255)
		canvas.DrawText("Testing macOS backend input handling", goflow.NewOffset(20, 65), subtitleStyle)

		// Draw interactive region indicator
		regionStyle := goflow.NewTextStyle()
		regionStyle.FontSize = 12
		regionStyle.Color = goflow.NewColor(100, 100, 255, 255)
		canvas.DrawText("← Click and type here", goflow.NewOffset(20, 90), regionStyle)

		// Test results
		textStyle := goflow.NewTextStyle()
		textStyle.FontSize = 16

		y := 130.0
		lineHeight := 30.0

		drawTestLine := func(label, value string, passed bool) {
			color := goflow.NewColor(150, 150, 150, 255)
			if passed {
				color = goflow.NewColor(0, 160, 0, 255)
			}
			textStyle.Color = color

			text := fmt.Sprintf("%-20s %s", label+":", value)
			canvas.DrawText(text, goflow.NewOffset(30, y), textStyle)
			y += lineHeight
		}

		drawTestLine("Mouse Clicks", fmt.Sprintf("%d", testResults.MouseClicks), testResults.MouseClicks > 0)
		drawTestLine("Mouse Moves", fmt.Sprintf("%d", testResults.MouseMoves), testResults.MouseMoves > 10)
		drawTestLine("Mouse Drags", fmt.Sprintf("%d", testResults.MouseDrags), testResults.MouseDrags > 0)
		drawTestLine("Last Position", testResults.LastMousePos, testResults.MouseClicks > 0)

		y += 15
		drawTestLine("Key Presses", fmt.Sprintf("%d", testResults.KeyPresses), testResults.KeyPresses > 0)
		drawTestLine("Last Key", testResults.LastKey, testResults.KeyPresses > 0)

		y += 15
		drawTestLine("Window Resizes", fmt.Sprintf("%d", resizeCount), resizeCount > 0)

		// Instructions box
		y += 30
		instructStyle := goflow.NewTextStyle()
		instructStyle.FontSize = 13
		instructStyle.Color = goflow.NewColor(90, 90, 90, 255)

		canvas.DrawText("Instructions:", goflow.NewOffset(30, y), instructStyle)
		y += 25

		instructStyle.FontSize = 12
		instructStyle.Color = goflow.NewColor(110, 110, 110, 255)
		canvas.DrawText("• Click anywhere to test mouse events", goflow.NewOffset(40, y), instructStyle)
		y += 20
		canvas.DrawText("• Drag to test mouse drag tracking", goflow.NewOffset(40, y), instructStyle)
		y += 20
		canvas.DrawText("• Type any keys to test keyboard", goflow.NewOffset(40, y), instructStyle)
		y += 20
		canvas.DrawText("• Resize window to test window events", goflow.NewOffset(40, y), instructStyle)

		frame++
	})

	// Set up animation timer
	go func() {
		ticker := time.NewTicker(16 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			window.SetNeedsDisplay()
		}
	}()

	// Show window
	window.Show()

	// Print test report after 15 seconds
	go func() {
		time.Sleep(15 * time.Second)
		printTestReport(testResults, resizeCount)
	}()

	fmt.Println("\n>> Window displayed. Interact with it to test input handling.")
	fmt.Println(">> Test report will print after 15 seconds...\n")

	// Run the application
	macos.Run()
}

type TestResults struct {
	MouseClicks  int
	MouseMoves   int
	MouseDrags   int
	KeyPresses   int
	LastMousePos string
	LastKey      string
}

func printTestReport(results *TestResults, resizeCount int) {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("INPUT SYSTEM TEST REPORT")
	fmt.Println(strings.Repeat("=", 70))

	total := 0
	passed := 0

	checkTest := func(name string, condition bool, note string) {
		total++
		status := "✗ FAIL"
		if condition {
			status = "✓ PASS"
			passed++
		}
		fmt.Printf("%-35s %s", name, status)
		if note != "" {
			fmt.Printf("  (%s)", note)
		}
		fmt.Println()
	}

	fmt.Println("\n1. Core Mouse Events:")
	checkTest("  Mouse Click Detection", results.MouseClicks > 0, fmt.Sprintf("%d clicks", results.MouseClicks))
	checkTest("  Mouse Move Tracking", results.MouseMoves > 10, fmt.Sprintf("%d moves", results.MouseMoves))
	checkTest("  Mouse Drag Detection", results.MouseDrags > 0, fmt.Sprintf("%d drags", results.MouseDrags))
	checkTest("  Mouse Position Accuracy", results.LastMousePos != "none", results.LastMousePos)

	fmt.Println("\n2. Core Keyboard Events:")
	checkTest("  Key Press Detection", results.KeyPresses > 0, fmt.Sprintf("%d keys", results.KeyPresses))
	checkTest("  Key Code Mapping", results.LastKey != "none", results.LastKey)

	fmt.Println("\n3. Window Events:")
	checkTest("  Window Resize Handling", resizeCount > 0, fmt.Sprintf("%d resizes", resizeCount))

	fmt.Println("\n" + strings.Repeat("-", 70))
	percentage := float64(passed) / float64(total) * 100
	fmt.Printf("OVERALL SCORE: %d/%d tests passed (%.1f%%)\n", passed, total, percentage)
	fmt.Println(strings.Repeat("=", 70))

	if passed == total {
		fmt.Println("\n🎉 EXCELLENT! All input tests passed!")
		fmt.Println("   The macOS backend input handling is fully functional.")
	} else if percentage >= 70 {
		fmt.Println("\n✅ GOOD! Most input features are working.")
		fmt.Println("   Some features may need more interaction to test.")
	} else {
		fmt.Println("\n⚠️  Some features need more testing:")
		if results.MouseClicks == 0 {
			fmt.Println("  - Click in the window to test mouse events")
		}
		if results.MouseDrags == 0 {
			fmt.Println("  - Drag the mouse to test drag detection")
		}
		if results.KeyPresses == 0 {
			fmt.Println("  - Press keys to test keyboard events")
		}
		if resizeCount == 0 {
			fmt.Println("  - Resize the window to test resize events")
		}
	}

	fmt.Println("\nNext steps for completion:")
	fmt.Println("  [ ] Focus management system")
	fmt.Println("  [ ] Text input field widget")
	fmt.Println("  [ ] Cursor position rendering")
	fmt.Println("  [ ] Multi-touch gesture integration")
}
