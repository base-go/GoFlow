package goflow

import (
	"runtime"
)

// TargetPlatform represents the platform the app is running on
type TargetPlatform int

const (
	// PlatformAndroid represents Android platform
	PlatformAndroid TargetPlatform = iota
	// PlatformIOS represents iOS platform
	PlatformIOS
	// PlatformMacOS represents macOS platform
	PlatformMacOS
	// PlatformLinux represents Linux platform
	PlatformLinux
	// PlatformWindows represents Windows platform
	PlatformWindows
	// PlatformWeb represents Web platform
	PlatformWeb
	// PlatformFuchsia represents Fuchsia platform
	PlatformFuchsia
)

// String returns the string representation of the platform
func (p TargetPlatform) String() string {
	switch p {
	case PlatformAndroid:
		return "android"
	case PlatformIOS:
		return "ios"
	case PlatformMacOS:
		return "macos"
	case PlatformLinux:
		return "linux"
	case PlatformWindows:
		return "windows"
	case PlatformWeb:
		return "web"
	case PlatformFuchsia:
		return "fuchsia"
	default:
		return "unknown"
	}
}

// IsMobile returns true if the platform is a mobile platform
func (p TargetPlatform) IsMobile() bool {
	return p == PlatformAndroid || p == PlatformIOS
}

// IsDesktop returns true if the platform is a desktop platform
func (p TargetPlatform) IsDesktop() bool {
	return p == PlatformMacOS || p == PlatformLinux || p == PlatformWindows
}

// IsApple returns true if the platform is an Apple platform
func (p TargetPlatform) IsApple() bool {
	return p == PlatformIOS || p == PlatformMacOS
}

// DefaultTheme returns the default theme for the platform
// Returns "cupertino" for Apple platforms, "material" for others
func (p TargetPlatform) DefaultTheme() string {
	if p.IsApple() {
		return "cupertino"
	}
	return "material"
}

// DetectPlatform detects the current platform based on GOOS and GOARCH
func DetectPlatform() TargetPlatform {
	switch runtime.GOOS {
	case "android":
		return PlatformAndroid
	case "ios":
		return PlatformIOS
	case "darwin":
		// Distinguish between macOS and iOS based on architecture
		if runtime.GOARCH == "arm64" || runtime.GOARCH == "arm" {
			// This could be iOS or macOS on Apple Silicon
			// In a real implementation, we'd check more context
			// For now, default to macOS for darwin
			return PlatformMacOS
		}
		return PlatformMacOS
	case "linux":
		return PlatformLinux
	case "windows":
		return PlatformWindows
	case "js":
		return PlatformWeb
	default:
		// Default to Linux for unknown platforms
		return PlatformLinux
	}
}

// Global platform variable
var currentPlatform = DetectPlatform()

// GetPlatform returns the current platform
func GetPlatform() TargetPlatform {
	return currentPlatform
}

// SetPlatform sets the current platform (for testing or overrides)
func SetPlatform(platform TargetPlatform) {
	currentPlatform = platform
}
