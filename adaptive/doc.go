// Package adaptive provides platform-adaptive widgets that automatically
// switch between Material Design (Android/Linux/Windows/Web) and Cupertino
// (iOS/macOS) styles based on the runtime platform.
//
// # Overview
//
// The adaptive package allows you to write UI code once and have it
// automatically adapt to the platform's native design language:
//
//   - On iOS and macOS: Uses Cupertino (iOS-style) widgets
//   - On Android, Linux, Windows, and Web: Uses Material Design widgets
//
// # Usage
//
// Simply use the adaptive widgets instead of platform-specific ones:
//
//	button := adaptive.NewButton("Click Me", func() {
//	    fmt.Println("Button clicked!")
//	})
//
// The button will automatically render as:
//   - A Material Design button on Android/Linux/Windows/Web
//   - A Cupertino button on iOS/macOS
//
// # Available Widgets
//
//   - Button: Adaptive button widget
//   - Card: Adaptive card widget
//   - AppBar: Adaptive app bar/navigation bar
//
// # Platform Override
//
// You can override the platform detection for testing:
//
//	goflow.SetPlatform(goflow.PlatformIOS)
//
// This will force all adaptive widgets to use iOS/Cupertino style.
package adaptive
