package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/base-go/GoFlow/cmd/goflow/templates"
	"github.com/base-go/GoFlow/cmd/goflow/utils"
	"github.com/base-go/mamba"
)

// NewCommand creates the 'goflow new' command
func NewCommand() *mamba.Command {
	cmd := &mamba.Command{
		Use:   "new [project-name]",
		Short: "Create a new GoFlow project",
		Long:  "Creates a new GoFlow project with the specified name and template.",
		Run: func(c *mamba.Command, args []string) {
			if err := runNewCommand(c, args); err != nil {
				c.PrintError(fmt.Sprintf("Error: %v", err))
			}
		},
		Args: mamba.ExactArgs(1),
	}

	cmd.Flags().StringP("template", "t", "default", "Project template (default, material, cupertino)")
	cmd.Flags().StringP("platforms", "p", "", "Target platforms (macos, windows, linux, all)")
	cmd.Flags().StringP("org", "o", "com.example", "Organization identifier")

	return cmd
}

func runNewCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintHeader("📦 Creating New GoFlow Project")

	projectName := args[0]

	// Get flags
	platforms, _ := cmd.Flags().GetString("platforms")
	template, _ := cmd.Flags().GetString("template")
	org, _ := cmd.Flags().GetString("org")

	cmd.PrintInfo(fmt.Sprintf("Project: %s", projectName))
	cmd.PrintInfo(fmt.Sprintf("Template: %s", template))
	cmd.PrintInfo(fmt.Sprintf("Platforms: %s", platforms))

	// Create project directory
	if _, err := os.Stat(projectName); err == nil {
		cmd.PrintError(fmt.Sprintf("✗ Directory '%s' already exists", projectName))
		return fmt.Errorf("directory already exists")
	}

	cmd.PrintInfo("Creating project structure...")

	// Create base directories
	dirs := []string{
		projectName,
		filepath.Join(projectName, "lib"),
		filepath.Join(projectName, "lib", "screens"),
		filepath.Join(projectName, "lib", "widgets"),
		filepath.Join(projectName, "assets"),
		filepath.Join(projectName, "test"),
	}

	// Add platform directories based on --platforms flag
	platformList := utils.ParsePlatforms(platforms)
	for _, platform := range platformList {
		switch platform {
		case "macos":
			dirs = append(dirs,
				filepath.Join(projectName, "macos"),
				filepath.Join(projectName, "macos", "Runner"),
				filepath.Join(projectName, "macos", "GoFlow"),
				filepath.Join(projectName, "macos", "GoFlow", "ephemeral"),
			)
		case "windows":
			dirs = append(dirs,
				filepath.Join(projectName, "windows"),
				filepath.Join(projectName, "windows", "runner"),
			)
		case "linux":
			dirs = append(dirs,
				filepath.Join(projectName, "linux"),
				filepath.Join(projectName, "linux", "runner"),
			)
		}
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			cmd.PrintError(fmt.Sprintf("✗ Failed to create %s: %v", dir, err))
			return err
		}
	}

	// Create go.mod
	cmd.PrintInfo("Initializing Go module...")
	modulePath := fmt.Sprintf("%s/%s", org, projectName)
	goMod := templates.CreateGoModTemplate(modulePath)

	if err := os.WriteFile(filepath.Join(projectName, "go.mod"), []byte(goMod), 0644); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Failed to create go.mod: %v", err))
		return err
	}

	// Create main.go
	cmd.PrintInfo("Creating application code...")
	mainGo := templates.CreateMainTemplate(projectName)
	if err := os.WriteFile(filepath.Join(projectName, "lib", "main.go"), []byte(mainGo), 0644); err != nil {
		cmd.PrintError(fmt.Sprintf("✗ Failed to create main.go: %v", err))
		return err
	}

	// Create README
	readme := templates.CreateReadmeTemplate(projectName)
	if err := os.WriteFile(filepath.Join(projectName, "README.md"), []byte(readme), 0644); err != nil {
		cmd.PrintWarning(fmt.Sprintf("⚠ Failed to create README.md: %v", err))
	}

	// Create platform-specific files
	cmd.PrintInfo("Creating platform files...")
	for _, platform := range platformList {
		if err := createPlatformFiles(projectName, platform, org); err != nil {
			cmd.PrintWarning(fmt.Sprintf("⚠ Failed to create %s platform files: %v", platform, err))
		} else {
			cmd.PrintSuccess(fmt.Sprintf("✓ Created %s platform files", platform))
		}
	}

	cmd.PrintSuccess(fmt.Sprintf("✓ Created project: %s", projectName))

	// Print next steps
	fmt.Println()
	cmd.PrintInfo("Next steps:")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Println("  go mod tidy")
	fmt.Println("  goflow run")

	return nil
}

// createPlatformFiles creates platform-specific files
func createPlatformFiles(projectName, platform, org string) error {
	switch platform {
	case "macos":
		return createMacOSFiles(projectName, org)
	case "windows":
		return createWindowsFiles(projectName)
	case "linux":
		return createLinuxFiles(projectName)
	default:
		return fmt.Errorf("unsupported platform: %s", platform)
	}
}

// createMacOSFiles creates macOS platform files
func createMacOSFiles(projectName, org string) error {
	basePath := filepath.Join(projectName, "macos")

	// AppDelegate.swift
	appDelegate := templates.CreateMacOSAppDelegateTemplate(projectName)
	if err := os.WriteFile(filepath.Join(basePath, "Runner", "AppDelegate.swift"), []byte(appDelegate), 0644); err != nil {
		return err
	}

	// MainWindow.swift
	mainWindow := templates.CreateMacOSMainWindowTemplate(projectName)
	if err := os.WriteFile(filepath.Join(basePath, "Runner", "MainWindow.swift"), []byte(mainWindow), 0644); err != nil {
		return err
	}

	// Info.plist
	infoPlist := templates.CreateMacOSInfoPlistTemplate(projectName, org)
	if err := os.WriteFile(filepath.Join(basePath, "Runner", "Info.plist"), []byte(infoPlist), 0644); err != nil {
		return err
	}

	// GoFlow-Generated.xcconfig
	xcconfig := templates.CreateMacOSXCConfigTemplate(projectName)
	if err := os.WriteFile(filepath.Join(basePath, "GoFlow", "ephemeral", "GoFlow-Generated.xcconfig"), []byte(xcconfig), 0644); err != nil {
		return err
	}

	// README
	readme := templates.CreateMacOSReadmeTemplate()
	if err := os.WriteFile(filepath.Join(basePath, "README.md"), []byte(readme), 0644); err != nil {
		return err
	}

	return nil
}

// createWindowsFiles creates Windows platform files (placeholder for now)
func createWindowsFiles(projectName string) error {
	basePath := filepath.Join(projectName, "windows")
	readme := templates.CreateWindowsReadmeTemplate()
	return os.WriteFile(filepath.Join(basePath, "README.md"), []byte(readme), 0644)
}

// createLinuxFiles creates Linux platform files (placeholder for now)
func createLinuxFiles(projectName string) error {
	basePath := filepath.Join(projectName, "linux")
	readme := templates.CreateLinuxReadmeTemplate()
	return os.WriteFile(filepath.Join(basePath, "README.md"), []byte(readme), 0644)
}
