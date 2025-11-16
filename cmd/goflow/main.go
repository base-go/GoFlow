package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/base-go/mamba"
)

const cliVersion = "0.2.0"

var rootCmd = &mamba.Command{
	Use:   "goflow",
	Short: "GoFlow - Flutter-inspired GUI framework for Go",
	Long: `GoFlow is a modern GUI framework for Go that brings Flutter's
developer experience to desktop and web applications.

Features:
  • Reactive UI with Signals
  • Hot reload for rapid development
  • Comprehensive widget library
  • Navigation & routing
  • Testing framework
  • Cross-platform support`,
	Version: cliVersion,
}

func main() {
	// Add commands
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(analyzeCmd)
	rootCmd.AddCommand(formatCmd)

	// Execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// newCmd creates a new GoFlow project
var newCmd = &mamba.Command{
	Use:     "new <project_name>",
	Aliases: []string{"create"},
	Short:   "Create a new GoFlow project",
	Long: `Create a new GoFlow project with Flutter-style structure.

Examples:
  goflow new myapp
  goflow new myapp --platforms=macos,linux
  goflow new myapp --template=material
  goflow new myapp --org=com.mycompany`,
	Args: mamba.ExactArgs(1),
	RunE: runNewCommand,
}

// runCmd runs the GoFlow app
var runCmd = &mamba.Command{
	Use:   "run [platform]",
	Short: "Run the GoFlow app with hot reload",
	Long: `Run your GoFlow app on the specified platform with hot reload enabled.

If no platform is specified, auto-detects the current platform.

Examples:
  goflow run              # Auto-detect platform
  goflow run macos        # Run on macOS
  goflow run --no-hot-reload  # Disable hot reload
  goflow run --release    # Run release build`,
	Args: mamba.MaximumNArgs(1),
	RunE: runRunCommand,
}

// buildCmd builds the GoFlow app
var buildCmd = &mamba.Command{
	Use:   "build [platform]",
	Short: "Build the GoFlow app for production",
	Long: `Build your GoFlow app for the specified platform with optimizations.

Examples:
  goflow build macos      # Build for macOS
  goflow build --release  # Optimized release build
  goflow build --bundle   # Create app bundle/installer`,
	Args: mamba.MaximumNArgs(1),
	RunE: runBuildCommand,
}

// testCmd runs tests
var testCmd = &mamba.Command{
	Use:   "test [packages...]",
	Short: "Run tests",
	Long: `Run widget tests, golden tests, and integration tests.

Examples:
  goflow test             # Run all tests
  goflow test ./lib/...   # Run specific tests
  goflow test --coverage  # With coverage
  goflow test --golden    # Update golden files`,
	RunE: runTestCommand,
}

// doctorCmd checks the environment
var doctorCmd = &mamba.Command{
	Use:   "doctor",
	Short: "Check your environment and dependencies",
	Long: `Verify that all required tools and dependencies are installed.

Checks:
  • Go installation and version
  • GoFlow version
  • Platform SDKs (Xcode, MSVC, GTK)
  • Required Go packages`,
	RunE: runDoctorCommand,
}

// cleanCmd cleans build artifacts
var cleanCmd = &mamba.Command{
	Use:   "clean",
	Short: "Clean build artifacts and caches",
	Long: `Remove build artifacts, temporary files, and caches.

Examples:
  goflow clean        # Clean build directory
  goflow clean --deep # Clean cache too`,
	RunE: runCleanCommand,
}

// analyzeCmd analyzes code
var analyzeCmd = &mamba.Command{
	Use:   "analyze",
	Short: "Analyze and lint your code",
	Long: `Run static analysis and linting on your GoFlow project.

Examples:
  goflow analyze          # Analyze code
  goflow analyze --fix    # Auto-fix issues`,
	RunE: runAnalyzeCommand,
}

// formatCmd formats code
var formatCmd = &mamba.Command{
	Use:   "format",
	Short: "Format your code",
	Long: `Format Go code using gofmt.

Examples:
  goflow format           # Format all .go files
  goflow format --check   # Check formatting without changes`,
	RunE: runFormatCommand,
}

func init() {
	// Flags for 'new' command
	newCmd.Flags().StringP("platforms", "p", "macos,windows,linux", "Target platforms (macos,windows,linux,web)")
	newCmd.Flags().StringP("template", "t", "default", "Project template (default,material,minimal)")
	newCmd.Flags().StringP("org", "o", "com.example", "Organization for module path")
	newCmd.Flags().Bool("interactive", false, "Interactive mode with prompts")

	// Flags for 'run' command
	runCmd.Flags().Bool("no-hot-reload", false, "Disable hot reload")
	runCmd.Flags().BoolP("release", "r", false, "Run release build")
	runCmd.Flags().String("target", "", "Target device/emulator")
	runCmd.Flags().IntP("port", "p", 8080, "Hot reload server port")

	// Flags for 'build' command
	buildCmd.Flags().BoolP("release", "r", true, "Build release version")
	buildCmd.Flags().Bool("bundle", false, "Create app bundle/installer")
	buildCmd.Flags().StringP("output", "o", "", "Output directory")
	buildCmd.Flags().Bool("strip", true, "Strip debug symbols")

	// Flags for 'test' command
	testCmd.Flags().BoolP("coverage", "c", false, "Generate coverage report")
	testCmd.Flags().Bool("golden", false, "Update golden files")
	testCmd.Flags().StringP("run", "r", "", "Run only tests matching pattern")
	testCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	// Flags for 'clean' command
	cleanCmd.Flags().Bool("deep", false, "Deep clean including cache")
	cleanCmd.Flags().BoolP("force", "f", false, "Force clean without confirmation")

	// Flags for 'analyze' command
	analyzeCmd.Flags().Bool("fix", false, "Auto-fix issues")
	analyzeCmd.Flags().StringSlice("ignore", []string{}, "Patterns to ignore")

	// Flags for 'format' command
	formatCmd.Flags().Bool("check", false, "Check formatting without changes")
	formatCmd.Flags().BoolP("write", "w", true, "Write changes to files")
}

// Command implementations (stubs for now, will implement in next phases)

func runNewCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("📦 Creating New GoFlow Project")
	// Will implement with current logic + interactive mode
	cmd.PrintWarning("Not yet fully implemented - migrating from old CLI...")
	return nil
}

func runRunCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🚀 Running GoFlow App")

	// Determine platform
	platform := "macos"
	if len(args) > 0 {
		platform = args[0]
	} else {
		platform = detectPlatform()
	}

	cmd.PrintInfo(fmt.Sprintf("Building for platform: %s", platform))

	// Check for main.go
	mainFile := "main.go"
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		mainFile = "cmd/app/main.go"
		if _, err := os.Stat(mainFile); os.IsNotExist(err) {
			cmd.PrintError("✗ No main.go found. Run this command in a GoFlow project directory.")
			return fmt.Errorf("main.go not found")
		}
	}

	// Get flags
	noHotReload, _ := cmd.Flags().GetBool("no-hot-reload")
	release, _ := cmd.Flags().GetBool("release")

	// Build the app
	buildMode := "debug"
	if release {
		buildMode = "release"
	}

	cmd.PrintInfo(fmt.Sprintf("Building in %s mode...", buildMode))

	buildArgs := []string{"build"}
	if release {
		buildArgs = append(buildArgs, "-ldflags", "-s -w") // Strip debug symbols
	}
	buildArgs = append(buildArgs, "-o", "./build/app", mainFile)

	if output, err := runCommand("go", buildArgs...); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Build failed:\n%s", output))
		return fmt.Errorf("build failed")
	}

	cmd.PrintSuccess("✓ Build completed successfully")

	if !noHotReload && !release {
		cmd.PrintInfo("Hot reload: enabled (watching for changes...)")
		cmd.PrintWarning("⚠ Full hot reload integration coming soon")
	}

	// Run the app
	cmd.PrintInfo("Starting application...")
	runArgs := []string{"./build/app"}

	if output, err := runCommand(runArgs[0]); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ App failed:\n%s", output))
		return fmt.Errorf("app failed")
	}

	return nil
}

func runBuildCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🔨 Building GoFlow App")

	// Determine platform
	platform := "macos"
	if len(args) > 0 {
		platform = args[0]
	} else {
		platform = detectPlatform()
	}

	// Get flags
	release, _ := cmd.Flags().GetBool("release")
	bundle, _ := cmd.Flags().GetBool("bundle")
	outputDir, _ := cmd.Flags().GetString("output")
	strip, _ := cmd.Flags().GetBool("strip")

	if outputDir == "" {
		outputDir = "./build"
	}

	cmd.PrintInfo(fmt.Sprintf("Building for platform: %s", platform))
	cmd.PrintInfo(fmt.Sprintf("Output directory: %s", outputDir))

	// Check for main.go
	mainFile := "main.go"
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		mainFile = "cmd/app/main.go"
		if _, err := os.Stat(mainFile); os.IsNotExist(err) {
			cmd.PrintError("✗ No main.go found")
			return fmt.Errorf("main.go not found")
		}
	}

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Failed to create output directory: %v", err))
		return err
	}

	// Build arguments
	buildArgs := []string{"build"}

	if release && strip {
		buildArgs = append(buildArgs, "-ldflags", "-s -w")
		cmd.PrintInfo("Build mode: Release (optimized, stripped)")
	} else {
		cmd.PrintInfo("Build mode: Debug")
	}

	outputFile := filepath.Join(outputDir, getAppName(platform))
	buildArgs = append(buildArgs, "-o", outputFile, mainFile)

	cmd.PrintInfo("Building...")

	if output, err := runCommand("go", buildArgs...); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Build failed:\n%s", output))
		return fmt.Errorf("build failed")
	}

	cmd.PrintSuccess(fmt.Sprintf("✓ Build completed: %s", outputFile))

	if bundle {
		cmd.PrintWarning("⚠ App bundling not yet implemented")
	}

	return nil
}

func runTestCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🧪 Running Tests")

	// Get flags
	coverage, _ := cmd.Flags().GetBool("coverage")
	golden, _ := cmd.Flags().GetBool("golden")
	runPattern, _ := cmd.Flags().GetString("run")
	verbose, _ := cmd.Flags().GetBool("verbose")

	// Build test arguments
	testArgs := []string{"test"}

	if verbose {
		testArgs = append(testArgs, "-v")
	}

	if coverage {
		testArgs = append(testArgs, "-coverprofile=coverage.out")
		cmd.PrintInfo("Coverage: enabled")
	}

	if runPattern != "" {
		testArgs = append(testArgs, "-run", runPattern)
		cmd.PrintInfo(fmt.Sprintf("Running tests matching: %s", runPattern))
	}

	if golden {
		cmd.PrintInfo("Golden files: will be updated")
		os.Setenv("UPDATE_GOLDENS", "1")
	}

	// Add package arguments
	if len(args) > 0 {
		testArgs = append(testArgs, args...)
	} else {
		testArgs = append(testArgs, "./...")
	}

	cmd.PrintInfo("Running tests...")

	if output, err := runCommand("go", testArgs...); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Tests failed:\n%s", output))
		return fmt.Errorf("tests failed")
	} else {
		cmd.PrintSuccess("✓ All tests passed")
		if verbose {
			fmt.Println(output)
		}
	}

	if coverage {
		cmd.PrintInfo("Generating coverage report...")
		if _, err := runCommand("go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html"); err != nil {
			cmd.PrintWarning(fmt.Sprintf("⚠ Failed to generate HTML coverage: %v", err))
		} else {
			cmd.PrintSuccess("✓ Coverage report: coverage.html")
		}
	}

	return nil
}

func runDoctorCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🏥 GoFlow Doctor")

	allOK := true

	// Check Go installation
	cmd.PrintInfo("Checking Go installation...")
	goVersion, err := checkGoVersion()
	if err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Go not found: %v", err))
		allOK = false
	} else {
		cmd.PrintSuccess(fmt.Sprintf("✓ Go version: %s", goVersion))
	}

	// Check GoFlow version
	cmd.PrintInfo("Checking GoFlow version...")
	cmd.PrintSuccess(fmt.Sprintf("✓ GoFlow CLI: v%s", cliVersion))

	// Check platform-specific tools
	cmd.PrintInfo("Checking platform tools...")
	if err := checkPlatformTools(cmd); err != nil {
		allOK = false
	}

	// Check required packages
	cmd.PrintInfo("Checking Go module dependencies...")
	if err := checkModuleDeps(cmd); err != nil {
		cmd.PrintWarning(fmt.Sprintf("⚠ Module check: %v", err))
	} else {
		cmd.PrintSuccess("✓ All dependencies available")
	}

	fmt.Println()
	if allOK {
		cmd.PrintSuccess("✓ Everything looks good! You're ready to build GoFlow apps.")
	} else {
		cmd.PrintWarning("⚠ Some issues detected. Please install missing dependencies.")
		return fmt.Errorf("environment check failed")
	}

	return nil
}

func runCleanCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🧹 Cleaning Build Artifacts")
	cmd.PrintWarning("Not yet implemented - coming in Phase 3")
	return nil
}

func runAnalyzeCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🔍 Analyzing Code")
	cmd.PrintWarning("Not yet implemented - coming in Phase 3")
	return nil
}

func runFormatCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("✨ Formatting Code")
	cmd.PrintWarning("Not yet implemented - coming in Phase 3")
	return nil
}

// Helper functions for doctor command

func checkGoVersion() (string, error) {
	output, err := runCommand("go", "version")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

func checkPlatformTools(cmd *mamba.Command) error {
	hasIssue := false

	switch getOS() {
	case "darwin":
		// Check for Xcode
		if _, err := runCommand("xcodebuild", "-version"); err != nil {
			cmd.PrintWarning("⚠ Xcode not found (required for macOS apps)")
			hasIssue = true
		} else {
			cmd.PrintSuccess("✓ Xcode installed")
		}
	case "windows":
		// Check for MSVC
		cmd.PrintInfo("ℹ Windows: Ensure Visual Studio with C++ tools is installed")
	case "linux":
		// Check for GTK
		if _, err := runCommand("pkg-config", "--modversion", "gtk+-3.0"); err != nil {
			cmd.PrintWarning("⚠ GTK3 not found (required for Linux apps)")
			hasIssue = true
		} else {
			cmd.PrintSuccess("✓ GTK3 installed")
		}
	}

	if hasIssue {
		return fmt.Errorf("platform tools missing")
	}
	return nil
}

func checkModuleDeps(cmd *mamba.Command) error {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return fmt.Errorf("not in a Go module directory")
	}

	output, err := runCommand("go", "list", "-m", "all")
	if err != nil {
		return err
	}

	// Check for key dependencies
	required := []string{
		"github.com/base-go/GoFlow",
		"github.com/fsnotify/fsnotify",
	}

	for _, dep := range required {
		if !strings.Contains(output, dep) {
			return fmt.Errorf("missing required dependency: %s", dep)
		}
	}

	return nil
}

func getOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "darwin"
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	default:
		return "unknown"
	}
}

func runCommand(name string, args ...string) (string, error) {
	var out strings.Builder
	cmd := exec.Command(name, args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func detectPlatform() string {
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

func getAppName(platform string) string {
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
