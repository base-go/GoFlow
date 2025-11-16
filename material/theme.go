package material

import (
	"github.com/base-go/GoFlow/goflow"
)

// MaterialTheme defines Material Design theme colors and typography
type MaterialTheme struct {
	// Primary colors
	PrimaryColor      *goflow.Color
	PrimaryDarkColor  *goflow.Color
	PrimaryLightColor *goflow.Color
	AccentColor       *goflow.Color

	// Background colors
	BackgroundColor *goflow.Color
	SurfaceColor    *goflow.Color
	ErrorColor      *goflow.Color

	// Text colors
	TextPrimaryColor   *goflow.Color
	TextSecondaryColor *goflow.Color
	TextDisabledColor  *goflow.Color

	// Divider color
	DividerColor *goflow.Color

	// Elevation shadows
	Elevation1 float64
	Elevation2 float64
	Elevation3 float64
	Elevation4 float64
}

// DefaultLightTheme returns the default Material Design light theme
func DefaultLightTheme() *MaterialTheme {
	return &MaterialTheme{
		PrimaryColor:      goflow.NewColor(33, 150, 243, 255),  // Blue 500
		PrimaryDarkColor:  goflow.NewColor(25, 118, 210, 255),  // Blue 700
		PrimaryLightColor: goflow.NewColor(100, 181, 246, 255), // Blue 300
		AccentColor:       goflow.NewColor(255, 64, 129, 255),  // Pink A200

		BackgroundColor: goflow.NewColor(250, 250, 250, 255), // Grey 50
		SurfaceColor:    goflow.NewColor(255, 255, 255, 255), // White
		ErrorColor:      goflow.NewColor(244, 67, 54, 255),   // Red 500

		TextPrimaryColor:   goflow.NewColor(33, 33, 33, 255),     // Almost Black
		TextSecondaryColor: goflow.NewColor(117, 117, 117, 255),  // Grey 600
		TextDisabledColor:  goflow.NewColor(189, 189, 189, 255),  // Grey 400

		DividerColor: goflow.NewColor(224, 224, 224, 255), // Grey 200

		Elevation1: 2.0,
		Elevation2: 4.0,
		Elevation3: 8.0,
		Elevation4: 16.0,
	}
}

// DefaultDarkTheme returns the default Material Design dark theme
func DefaultDarkTheme() *MaterialTheme {
	return &MaterialTheme{
		PrimaryColor:      goflow.NewColor(144, 202, 249, 255), // Blue 200
		PrimaryDarkColor:  goflow.NewColor(100, 181, 246, 255), // Blue 300
		PrimaryLightColor: goflow.NewColor(187, 222, 251, 255), // Blue 100
		AccentColor:       goflow.NewColor(255, 64, 129, 255),  // Pink A200

		BackgroundColor: goflow.NewColor(18, 18, 18, 255),  // Dark Grey
		SurfaceColor:    goflow.NewColor(33, 33, 33, 255),  // Slightly lighter
		ErrorColor:      goflow.NewColor(239, 83, 80, 255), // Red 400

		TextPrimaryColor:   goflow.NewColor(255, 255, 255, 255),  // White
		TextSecondaryColor: goflow.NewColor(189, 189, 189, 255),  // Grey 400
		TextDisabledColor:  goflow.NewColor(117, 117, 117, 255),  // Grey 600

		DividerColor: goflow.NewColor(66, 66, 66, 255), // Grey 800

		Elevation1: 2.0,
		Elevation2: 4.0,
		Elevation3: 8.0,
		Elevation4: 16.0,
	}
}
