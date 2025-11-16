# GoFlow CLI Documentation

Complete guide to the GoFlow command-line interface.

## Table of Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Commands](#commands)
  - [new](#goflow-new) - Create a new project
  - [run](#goflow-run) - Run your app
  - [build](#goflow-build) - Build for production
  - [test](#goflow-test) - Run tests
  - [doctor](#goflow-doctor) - Check environment
  - [clean](#goflow-clean) - Clean artifacts
  - [analyze](#goflow-analyze) - Analyze code
  - [format](#goflow-format) - Format code
  - [version](#goflow-version) - Show version
- [Examples](#examples)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)

## Installation

### From Source

```bash
git clone https://github.com/base-go/GoFlow.git
cd GoFlow
go build -o goflow ./cmd/goflow
sudo mv goflow /usr/local/bin/
```

### From Release

```bash
# macOS (Apple Silicon)
curl -L https://github.com/base-go/GoFlow/releases/latest/download/goflow-macos-arm64 -o goflow
chmod +x goflow
sudo mv goflow /usr/local/bin/

# Linux (AMD64)
curl -L https://github.com/base-go/GoFlow/releases/latest/download/goflow-linux-amd64 -o goflow
chmod +x goflow
sudo mv goflow /usr/local/bin/
```

### Verify Installation

```bash
goflow version
# Output: GoFlow CLI v0.2.0
```

## Getting Started

### Create Your First App

```bash
# Create a new project
goflow new myapp

# Navigate to project
cd myapp

# Run the app
goflow run

# Run tests
goflow test

# Build for production
goflow build --release
```

## Commands

### goflow new

Create a new GoFlow project with Flutter-style structure.

**Usage:**
```bash
goflow new <project_name> [flags]
```

**Flags:**
- `-p, --platforms <string>` - Target platforms (default: "macos,windows,linux")
- `-t, --template <string>` - Project template (default, material, minimal) (default: "default")
- `-o, --org <string>` - Organization for module path (default: "com.example")
- `--interactive` - Interactive mode with prompts

**Examples:**
```bash
# Basic project
goflow new myapp

# Multi-platform with custom org
goflow new myapp --platforms=macos,linux --org=com.mycompany

# Material Design template
goflow new myapp --template=material

# Interactive mode
goflow new myapp --interactive
```

**Project Structure:**
```
myapp/
├── main.go              # Application entry point
├── go.mod               # Go module file
├── lib/                 # Application code
│   ├── main.go
│   └── widgets/
├── test/                # Tests
│   └── widget_test.go
├── assets/              # Static assets
├── build/               # Build output
└── README.md
```

---

### goflow run

Run your GoFlow app with hot reload support.

**Usage:**
```bash
goflow run [platform] [flags]
```

**Flags:**
- `--no-hot-reload` - Disable hot reload
- `-r, --release` - Run release build
- `--target <string>` - Target device/emulator
- `-p, --port <int>` - Hot reload server port (default: 8080)

**Examples:**
```bash
# Auto-detect platform
goflow run

# Specific platform
goflow run macos

# Release mode (no hot reload)
goflow run --release

# Disable hot reload
goflow run --no-hot-reload

# Custom port
goflow run --port 3000
```

**Hot Reload:**
- Watches `.go` files for changes
- Automatically rebuilds and restarts
- Preserves application state (when possible)
- Real-time feedback during development

**Note:** Full hot reload integration coming soon. Currently builds and runs the app.

---

### goflow build

Build your GoFlow app for production with optimizations.

**Usage:**
```bash
goflow build [platform] [flags]
```

**Flags:**
- `-r, --release` - Build release version (default: true)
- `--bundle` - Create app bundle/installer
- `-o, --output <string>` - Output directory
- `--strip` - Strip debug symbols (default: true)

**Examples:**
```bash
# Build for current platform
goflow build

# Build for macOS
goflow build macos

# Debug build (with symbols)
goflow build --release=false --strip=false

# Custom output directory
goflow build --output ./dist

# Create app bundle (macOS)
goflow build macos --bundle
```

**Build Modes:**
- **Debug**: Full symbols, no optimization, faster compilation
- **Release**: Optimized, stripped symbols, smaller binary

**Output:**
```
build/
└── app           # Executable binary
```

---

### goflow test

Run widget tests, golden tests, and integration tests.

**Usage:**
```bash
goflow test [packages...] [flags]
```

**Flags:**
- `-c, --coverage` - Generate coverage report
- `--golden` - Update golden files
- `-r, --run <string>` - Run only tests matching pattern
- `-v, --verbose` - Verbose output

**Examples:**
```bash
# Run all tests
goflow test

# Run tests in specific package
goflow test ./lib/...

# Generate coverage report
goflow test --coverage

# Update golden files
goflow test --golden

# Run tests matching pattern
goflow test --run TestButton

# Verbose output with coverage
goflow test -v --coverage
```

**Test Types:**

1. **Widget Tests**: Test individual widgets
   ```go
   func TestMyWidget(t *testing.T) {
       wt := testing.NewWidgetTester(t)
       widget := NewMyWidget()
       wt.PumpWidget(widget)
       // Assert widget behavior
   }
   ```

2. **Golden Tests**: Visual regression testing
   ```go
   func TestButtonAppearance(t *testing.T) {
       wt := testing.NewWidgetTester(t)
       gt := testing.NewGoldenTester(t, nil)
       wt.PumpWidget(NewButton("Click"))
       gt.MatchesGolden("button_default", wt.GetCanvas().GetImage())
   }
   ```

3. **Integration Tests**: End-to-end workflows
   ```go
   func TestLoginFlow(t *testing.T) {
       it := testing.NewIntegrationTester(t, nil)
       scenario := &testing.Scenario{
           Name: "User Login",
           Steps: []testing.Step{...},
       }
       it.RunScenario(scenario)
   }
   ```

**Coverage Output:**
- `coverage.out` - Coverage data
- `coverage.html` - HTML report (open in browser)

---

### goflow doctor

Check your environment and dependencies.

**Usage:**
```bash
goflow doctor
```

**Checks:**
- ✓ Go installation and version
- ✓ GoFlow version
- ✓ Platform SDKs (Xcode, MSVC, GTK)
- ✓ Required Go packages

**Example Output:**
```
🏥 GoFlow Doctor

ℹ Checking Go installation...
✓ Go version: go version go1.25.0 darwin/arm64
ℹ Checking GoFlow version...
✓ GoFlow CLI: v0.2.0
ℹ Checking platform tools...
✓ Xcode installed
ℹ Checking Go module dependencies...
✓ All dependencies available

✓ Everything looks good! You're ready to build GoFlow apps.
```

**Platform Requirements:**

| Platform | Required Tools |
|----------|---------------|
| macOS    | Xcode Command Line Tools |
| Windows  | Visual Studio with C++ tools |
| Linux    | GTK3 development libraries |

**Installing Platform Tools:**

**macOS:**
```bash
xcode-select --install
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install libgtk-3-dev
```

**Windows:**
Download Visual Studio Community with "Desktop development with C++"

---

### goflow clean

Remove build artifacts, temporary files, and caches.

**Usage:**
```bash
goflow clean [flags]
```

**Flags:**
- `--deep` - Deep clean including Go cache
- `-f, --force` - Force clean without confirmation

**Examples:**
```bash
# Clean build artifacts
goflow clean

# Deep clean (includes Go cache)
goflow clean --deep

# Force clean without prompt
goflow clean --force

# Deep clean with force
goflow clean --deep --force
```

**What Gets Cleaned:**

**Standard Clean:**
- `./build/` - Build directory
- `./dist/` - Distribution directory
- `*.test` - Test binaries
- `coverage.out` - Coverage data
- `coverage.html` - Coverage report

**Deep Clean (--deep):**
- All standard clean items
- `./go.work` - Go workspace file
- `./go.work.sum` - Go workspace checksum
- Go build cache
- Go module cache

**Example Output:**
```
🧹 Cleaning Build Artifacts

ℹ Deep clean: enabled
ℹ Removing: ./build
ℹ Removing: coverage.out
ℹ Cleaning Go build cache...
✓ Go cache cleaned
✓ Cleaned 2 items successfully
```

---

### goflow analyze

Analyze and lint your code using Go's static analysis tools.

**Usage:**
```bash
goflow analyze [flags]
```

**Flags:**
- `--fix` - Auto-fix issues (requires golangci-lint)
- `--ignore <strings>` - Patterns to ignore

**Examples:**
```bash
# Run analysis
goflow analyze

# Auto-fix issues
goflow analyze --fix

# Ignore specific patterns
goflow analyze --ignore=vendor,generated
```

**Analysis Tools:**

1. **go vet** (built-in)
   - Always runs
   - Checks for common mistakes
   - Reports suspicious constructs

2. **staticcheck** (optional)
   - Runs if installed
   - Advanced static analysis
   - Install: `go install honnef.co/go/tools/cmd/staticcheck@latest`

3. **golangci-lint** (optional, for --fix)
   - Runs multiple linters
   - Can auto-fix issues
   - Install: https://golangci-lint.run/usage/install/

**Example Output:**
```
🔍 Analyzing Code

ℹ Running go vet...
✓ go vet: no issues
ℹ Checking for staticcheck...
ℹ Running staticcheck...
✓ staticcheck: no issues

✓ Code analysis complete: no issues found
```

---

### goflow format

Format your Go code using gofmt and goimports.

**Usage:**
```bash
goflow format [flags]
```

**Flags:**
- `--check` - Check formatting without changes
- `-w, --write` - Write changes to files (default: true)

**Examples:**
```bash
# Format all .go files
goflow format

# Check formatting without changes
goflow format --check

# Explicit write mode
goflow format --write
```

**Formatting Tools:**

1. **gofmt** (built-in)
   - Always runs
   - Standard Go formatting

2. **goimports** (optional)
   - Runs if installed
   - Organizes imports
   - Install: `go install golang.org/x/tools/cmd/goimports@latest`

**Example Output:**
```
✨ Formatting Code

ℹ Mode: format
ℹ Running gofmt...
✓ Code formatted successfully
ℹ Checking for goimports...
ℹ Running goimports...
✓ Imports organized
```

**Check Mode Output:**
```
✨ Formatting Code

ℹ Mode: check
ℹ Running gofmt...
⚠ Files need formatting:
lib/main.go
lib/widgets/button.go
```

---

### goflow version

Show version information.

**Usage:**
```bash
goflow version
```

**Alias:**
```bash
goflow v
```

**Example Output:**
```
ℹ GoFlow CLI v0.2.0
```

---

## Examples

### Complete Development Workflow

```bash
# 1. Create new project
goflow new todo-app --template=material

# 2. Navigate to project
cd todo-app

# 3. Check environment
goflow doctor

# 4. Run in development mode
goflow run

# 5. Make code changes (hot reload applies automatically)

# 6. Run tests
goflow test --coverage

# 7. Format code
goflow format

# 8. Analyze code
goflow analyze

# 9. Build for production
goflow build --release

# 10. Clean up
goflow clean
```

### CI/CD Pipeline

```bash
# Continuous Integration
goflow doctor              # Check environment
goflow format --check      # Verify formatting
goflow analyze             # Static analysis
goflow test --coverage     # Run tests with coverage
goflow build --release     # Test production build

# Continuous Deployment
goflow build --release --bundle  # Create release bundle
```

### Testing Workflow

```bash
# Run all tests
goflow test

# Test-driven development
goflow test --run TestMyFeature -v

# Update visual tests
goflow test --golden

# Generate coverage report
goflow test --coverage
open coverage.html
```

## Configuration

### Project Configuration

Create `.goflow.yml` in project root (future feature):

```yaml
version: 0.2.0

platforms:
  - macos
  - linux
  - windows

build:
  release:
    strip: true
    optimize: true

test:
  coverage:
    threshold: 80
  golden:
    directory: testdata/golden

format:
  exclude:
    - vendor/
    - generated/
```

### Environment Variables

- `UPDATE_GOLDENS=1` - Update golden files (alternative to --golden)
- `GOFLOW_NO_COLOR=1` - Disable colored output
- `GOFLOW_VERBOSE=1` - Enable verbose output

## Troubleshooting

### Common Issues

#### "No main.go found"
**Solution:**
```bash
# Ensure you're in a GoFlow project directory
ls main.go

# Or specify the correct directory
cd path/to/project
goflow run
```

#### "Build failed"
**Solution:**
```bash
# Check for compile errors
go build ./...

# Check dependencies
go mod tidy
go mod download

# Run doctor
goflow doctor
```

#### "Tests failing"
**Solution:**
```bash
# Run with verbose output
goflow test -v

# Run specific test
goflow test --run TestName -v

# Clean and retry
goflow clean
goflow test
```

#### "Xcode not found" (macOS)
**Solution:**
```bash
xcode-select --install
sudo xcode-select --switch /Applications/Xcode.app
```

#### "GTK not found" (Linux)
**Solution:**
```bash
# Ubuntu/Debian
sudo apt-get install libgtk-3-dev

# Fedora
sudo dnf install gtk3-devel

# Arch
sudo pacman -S gtk3
```

### Getting Help

```bash
# General help
goflow --help

# Command-specific help
goflow new --help
goflow run --help
goflow test --help
```

### Debug Mode

Enable debug output:

```bash
# Verbose command output
goflow run -v

# Go build verbose
go build -v -o myapp ./cmd/app
```

## Terminal Features

GoFlow CLI uses the [Mamba](https://github.com/base-go/Mamba) framework for beautiful terminal UI:

- ✓ Styled output with colors
- ✓ Emoji indicators (✓, ✗, ⚠, ℹ)
- ✓ Progress indicators
- ✓ Formatted help text
- ✓ Responsive layout

### Output Types

- **Success**: ✓ Green text
- **Error**: ✗ Red text
- **Warning**: ⚠ Yellow text
- **Info**: ℹ Blue text
- **Header**: Bold with emoji

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Related Documentation

- [Getting Started Guide](README.md)
- [Widget Documentation](pkg/core/widgets/README.md)
- [Testing Framework](pkg/testing/README.md)
- [Hot Reload System](pkg/hotreload/README.md)
- [Release Process](RELEASES.md)

## Version History

- **v0.2.0** (2025-01-15)
  - Migrated to Mamba framework
  - Added development commands (run, test, doctor, build)
  - Added utility commands (clean, analyze, format)
  - Added version command
  - Beautiful terminal UI

- **v0.1.0** (2024-12-01)
  - Initial CLI release
  - Basic project creation

## Support

- GitHub Issues: https://github.com/base-go/GoFlow/issues
- Discussions: https://github.com/base-go/GoFlow/discussions
- Documentation: https://goflow.dev/docs
