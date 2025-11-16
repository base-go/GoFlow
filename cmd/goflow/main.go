package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

const version = "0.1.0"

type TemplateData struct {
	ProjectName string
	AppName     string
	ModulePath  string
	Platform    string
	Platforms   []string
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "new", "create":
		createCommand()
	case "version", "--version", "-v":
		fmt.Printf("GoFlow CLI v%s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("GoFlow - Flutter-inspired GUI framework for Go")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  goflow new <project_name> [flags]        Create a new GoFlow project")
	fmt.Println("  goflow create <project_name> [flags]     Alias for 'new'")
	fmt.Println("  goflow version                           Show version information")
	fmt.Println("  goflow help                              Show this help message")
	fmt.Println()
	fmt.Println("Flags for 'new'/'create':")
	fmt.Println("  --platforms string    Comma-separated list of platforms (default: macos,windows,linux)")
	fmt.Println("                        Available: macos, windows, linux, web")
	fmt.Println("  --template string     Project template (default: default)")
	fmt.Println("                        Available: default, material, minimal")
	fmt.Println("  --org string          Organization name for the module (default: com.example)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  goflow new myapp")
	fmt.Println("  goflow new myapp --platforms=macos,linux")
	fmt.Println("  goflow new myapp --template=material --org=com.mycompany")
	fmt.Println()
	fmt.Println("Installation:")
	fmt.Println("  go install github.com/base-go/GoFlow/cmd/goflow@latest")
}

func createCommand() {
	createFlags := flag.NewFlagSet("create", flag.ExitOnError)
	platforms := createFlags.String("platforms", "macos,windows,linux", "Comma-separated list of platforms")
	templateName := createFlags.String("template", "default", "Project template")
	org := createFlags.String("org", "com.example", "Organization name")

	if len(os.Args) < 3 {
		fmt.Println("Error: project name required")
		fmt.Println()
		fmt.Println("Usage: goflow new <project_name> [flags]")
		os.Exit(1)
	}

	projectName := os.Args[2]
	createFlags.Parse(os.Args[3:])

	// Validate project name
	if !isValidProjectName(projectName) {
		fmt.Printf("Error: '%s' is not a valid project name\n", projectName)
		fmt.Println("Project name must:")
		fmt.Println("  - Start with a letter")
		fmt.Println("  - Contain only letters, numbers, underscores, and hyphens")
		os.Exit(1)
	}

	// Parse platforms
	platformList := strings.Split(*platforms, ",")
	validPlatforms := validatePlatforms(platformList)

	if len(validPlatforms) == 0 {
		fmt.Println("Error: no valid platforms specified")
		os.Exit(1)
	}

	fmt.Printf("Creating GoFlow project: %s\n", projectName)
	fmt.Printf("Organization: %s\n", *org)
	fmt.Printf("Platforms: %s\n", strings.Join(validPlatforms, ", "))
	fmt.Printf("Template: %s\n", *templateName)
	fmt.Println()

	// Create project
	if err := createProject(projectName, *org, validPlatforms, *templateName); err != nil {
		fmt.Printf("Error creating project: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("✅ Project created successfully!")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Println("  go mod tidy")
	fmt.Println()
	fmt.Println("To run your app (development):")
	firstPlatform := validPlatforms[0]
	fmt.Printf("  cd %s && go run main.go\n", firstPlatform)
	fmt.Println()
	fmt.Println("To build for specific platform:")
	for _, p := range validPlatforms {
		fmt.Printf("  cd %s && go build    # %s\n", p, strings.Title(p))
	}
}

func isValidProjectName(name string) bool {
	if len(name) == 0 {
		return false
	}

	// Must start with a letter
	if !((name[0] >= 'a' && name[0] <= 'z') || (name[0] >= 'A' && name[0] <= 'Z')) {
		return false
	}

	// Must contain only letters, numbers, underscores, hyphens
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-') {
			return false
		}
	}

	return true
}

func validatePlatforms(platforms []string) []string {
	validPlatformNames := map[string]bool{
		"macos":   true,
		"windows": true,
		"linux":   true,
		"web":     true,
	}

	result := make([]string, 0)
	seen := make(map[string]bool)

	for _, p := range platforms {
		p = strings.TrimSpace(strings.ToLower(p))
		if validPlatformNames[p] && !seen[p] {
			result = append(result, p)
			seen[p] = true
		}
	}

	return result
}

func createProject(projectName, org string, platforms []string, templateName string) error {
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	projectPath, err := filepath.Abs(projectName)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Create directory structure
	dirs := []string{
		"lib/screens",
		"lib/widgets",
		"lib/models",
		"lib/services",
		"lib/state",
		"assets/fonts",
		"assets/images",
		"assets/icons",
		"test",
	}

	// Add platform directories
	for _, platform := range platforms {
		dirs = append(dirs, filepath.Join(platform, "runner"))
		dirs = append(dirs, filepath.Join(platform, "assets"))
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(projectPath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Prepare template data
	modulePath := fmt.Sprintf("%s/%s", org, projectName)
	appName := capitalize(projectName)
	data := TemplateData{
		ProjectName: projectName,
		AppName:     appName,
		ModulePath:  modulePath,
		Platforms:   platforms,
	}

	// Create root-level files from templates
	if err := createFromTemplate(filepath.Join(projectPath, "go.mod"), "templates/go.mod.tmpl", data); err != nil {
		return err
	}

	if err := createFromTemplate(filepath.Join(projectPath, ".gitignore"), "templates/.gitignore.tmpl", data); err != nil {
		return err
	}

	if err := createFromTemplate(filepath.Join(projectPath, "README.md"), "templates/README.md.tmpl", data); err != nil {
		return err
	}

	// Create goflow.yaml (similar to Flutter's pubspec.yaml)
	if err := createFromTemplate(filepath.Join(projectPath, "goflow.yaml"), "templates/goflow.yaml.tmpl", data); err != nil {
		return err
	}

	// Create analysis_options.yaml (for future linting)
	if err := createFromTemplate(filepath.Join(projectPath, "analysis_options.yaml"), "templates/analysis_options.yaml.tmpl", data); err != nil {
		return err
	}

	// Create lib/main.go from template
	libTemplatePath := fmt.Sprintf("templates/%s/lib_main.go.tmpl", templateName)
	if err := createFromTemplate(filepath.Join(projectPath, "lib", "main.go"), libTemplatePath, data); err != nil {
		return err
	}

	// Create platform runners
	for _, platform := range platforms {
		data.Platform = platform
		runnerPath := filepath.Join(projectPath, platform, "main.go")
		if err := createFromTemplate(runnerPath, "templates/platform_main.go.tmpl", data); err != nil {
			return err
		}

		// Create platform-specific .gitignore
		var platformGitignoreTmpl string
		switch platform {
		case "macos":
			platformGitignoreTmpl = "templates/macos_.gitignore.tmpl"
		case "linux":
			platformGitignoreTmpl = "templates/linux_.gitignore.tmpl"
		case "windows":
			platformGitignoreTmpl = "templates/windows_.gitignore.tmpl"
		}

		if platformGitignoreTmpl != "" {
			gitignorePath := filepath.Join(projectPath, platform, ".gitignore")
			if err := createFromTemplate(gitignorePath, platformGitignoreTmpl, data); err != nil {
				return err
			}
		}
	}

	return nil
}

func createFromTemplate(outputPath, templatePath string, data TemplateData) error {
	// Read template from embedded FS
	tmplContent, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
	}

	// Parse template with custom functions
	funcMap := template.FuncMap{
		"title": strings.Title,
		"hasPlatform": func(platforms []string, platform string) bool {
			for _, p := range platforms {
				if p == platform {
					return true
				}
			}
			return false
		},
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer outFile.Close()

	// Execute template
	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	return nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	// Convert to title case and remove hyphens/underscores
	// e.g., "material-demo" -> "MaterialDemo", "my_app" -> "MyApp"
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return strings.Join(parts, "")
}
