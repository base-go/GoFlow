package main

import (
	"fmt"
	"os"

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
	cmd.PrintWarning("Not yet implemented - coming in Phase 2")
	return nil
}

func runBuildCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🔨 Building GoFlow App")
	cmd.PrintWarning("Not yet implemented - coming in Phase 2")
	return nil
}

func runTestCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🧪 Running Tests")
	cmd.PrintWarning("Not yet implemented - coming in Phase 2")
	return nil
}

func runDoctorCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("🏥 GoFlow Doctor")
	cmd.PrintWarning("Not yet implemented - coming in Phase 2")
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
