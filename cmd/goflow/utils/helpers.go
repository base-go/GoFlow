package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

// RunCommand executes a command and returns the output
func RunCommand(name string, args ...string) (string, error) {
	var out strings.Builder
	cmd := exec.Command(name, args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

// DetectPlatform returns the current platform
func DetectPlatform() string {
	switch runtime.GOOS {
	case "darwin":
		return "macos"
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	default:
		return runtime.GOOS
	}
}

// GetAppName returns the app executable name for the platform
func GetAppName(platform string) string {
	switch platform {
	case "windows":
		return "app.exe"
	case "macos":
		return "app"
	case "linux":
		return "app"
	default:
		return "app"
	}
}

// ParsePlatforms parses the platforms string into a list
func ParsePlatforms(platforms string) []string {
	if platforms == "" {
		// Default to current platform
		return []string{DetectPlatform()}
	}

	if platforms == "all" {
		return []string{"macos", "windows", "linux"}
	}

	// Split by comma
	result := []string{}
	for _, p := range SplitString(platforms, ",") {
		p = TrimString(p)
		if p != "" {
			result = append(result, p)
		}
	}

	return result
}

// SplitString splits a string by separator
func SplitString(s, sep string) []string {
	if s == "" {
		return []string{}
	}

	var result []string
	var current string

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, current)
			current = ""
			i += len(sep) - 1
		} else {
			current += string(s[i])
		}
	}

	result = append(result, current)
	return result
}

// TrimString trims whitespace from a string
func TrimString(s string) string {
	start := 0
	end := len(s)

	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}

	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}

	return s[start:end]
}
