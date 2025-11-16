package cupertino

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// CupertinoTheme defines iOS/macOS design theme colors
type CupertinoTheme struct {
	// Primary colors (iOS blue)
	PrimaryColor      *goflow.Color
	SecondaryColor    *goflow.Color

	// Background colors
	BackgroundColor        *goflow.Color
	SystemBackgroundColor  *goflow.Color
	SecondaryBackgroundColor *goflow.Color

	// Label colors
	LabelColor           *goflow.Color
	SecondaryLabelColor  *goflow.Color
	TertiaryLabelColor   *goflow.Color
	QuaternaryLabelColor *goflow.Color

	// Separator color
	SeparatorColor *goflow.Color

	// Destructive color
	DestructiveColor *goflow.Color
}

// DefaultLightTheme returns the default iOS/macOS light theme
func DefaultLightTheme() *CupertinoTheme {
	return &CupertinoTheme{
		PrimaryColor:   goflow.NewColor(0, 122, 255, 255),   // iOS Blue
		SecondaryColor: goflow.NewColor(88, 86, 214, 255),   // iOS Purple

		BackgroundColor:              goflow.NewColor(255, 255, 255, 255), // White
		SystemBackgroundColor:        goflow.NewColor(242, 242, 247, 255), // Light Grey
		SecondaryBackgroundColor:     goflow.NewColor(255, 255, 255, 255), // White

		LabelColor:           goflow.NewColor(0, 0, 0, 255),       // Black
		SecondaryLabelColor:  goflow.NewColor(60, 60, 67, 153),    // 60% opacity
		TertiaryLabelColor:   goflow.NewColor(60, 60, 67, 76),     // 30% opacity
		QuaternaryLabelColor: goflow.NewColor(60, 60, 67, 46),     // 18% opacity

		SeparatorColor: goflow.NewColor(60, 60, 67, 73), // 29% opacity

		DestructiveColor: goflow.NewColor(255, 59, 48, 255), // iOS Red
	}
}

// DefaultDarkTheme returns the default iOS/macOS dark theme
func DefaultDarkTheme() *CupertinoTheme {
	return &CupertinoTheme{
		PrimaryColor:   goflow.NewColor(10, 132, 255, 255),  // iOS Blue (dark mode)
		SecondaryColor: goflow.NewColor(94, 92, 230, 255),   // iOS Purple (dark mode)

		BackgroundColor:              goflow.NewColor(0, 0, 0, 255),       // Black
		SystemBackgroundColor:        goflow.NewColor(28, 28, 30, 255),    // Dark Grey
		SecondaryBackgroundColor:     goflow.NewColor(44, 44, 46, 255),    // Slightly lighter

		LabelColor:           goflow.NewColor(255, 255, 255, 255),    // White
		SecondaryLabelColor:  goflow.NewColor(235, 235, 245, 153),    // 60% opacity
		TertiaryLabelColor:   goflow.NewColor(235, 235, 245, 76),     // 30% opacity
		QuaternaryLabelColor: goflow.NewColor(235, 235, 245, 46),     // 18% opacity

		SeparatorColor: goflow.NewColor(84, 84, 88, 163), // 65% opacity

		DestructiveColor: goflow.NewColor(255, 69, 58, 255), // iOS Red (dark mode)
	}
}
