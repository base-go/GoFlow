package templates

import "fmt"

// CreateMainTemplate creates the main.go template
func CreateMainTemplate(projectName string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to %s!")
	fmt.Println("GoFlow app is running...")

	// TODO: Initialize GoFlow app
	// app := goflow.NewApp()
	// app.Run(HomePage())
}
`, projectName)
}

// CreateReadmeTemplate creates the README.md template
func CreateReadmeTemplate(projectName string) string {
	return fmt.Sprintf(`# %s

A GoFlow application.

## Getting Started

### Run the app

`+"```bash"+`
goflow run
`+"```"+`

### Build for production

`+"```bash"+`
goflow build
`+"```"+`

### Build tests

`+"```bash"+`
goflow test
`+"```"+`

## Project Structure

- `+"`lib/`"+` - Application code
- `+"`lib/screens/`"+` - Screen widgets
- `+"`lib/widgets/`"+` - Reusable widgets
- `+"`assets/`"+` - Images, fonts, etc.
- `+"`test/`"+` - Tests
- `+"`macos/`"+` - macOS native shell
- `+"`windows/`"+` - Windows native shell
- `+"`linux/`"+` - Linux native shell

## Learn More

- [GoFlow Documentation](https://github.com/base-go/GoFlow)
- [CLI Reference](https://github.com/base-go/GoFlow/blob/dev/CLI.md)
`, projectName)
}

// CreateGoModTemplate creates the go.mod template
func CreateGoModTemplate(modulePath string) string {
	return fmt.Sprintf(`module %s

go 1.25.0

require github.com/base-go/GoFlow v0.2.0
`, modulePath)
}
