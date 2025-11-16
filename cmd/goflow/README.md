# GoFlow CLI

The official command-line tool for creating and managing GoFlow projects.

## Installation

### Via `go install` (Recommended)

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

This installs the `goflow` binary to your `$GOPATH/bin` directory. Make sure `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### From Source

```bash
git clone https://github.com/base-go/GoFlow.git
cd GoFlow/cmd/goflow
go build -ldflags="-s -w" -o goflow
sudo mv goflow /usr/local/bin/  # Optional: install globally
```

### Verify Installation

```bash
goflow version
# Output: GoFlow CLI v0.1.0
```

## Quick Start

```bash
# Create a new project (Flutter-style command)
goflow new myapp

# Navigate to project
cd myapp

# Check the Flutter-inspired structure
ls -la
# goflow.yaml, lib/, macos/, linux/, windows/, assets/, etc.

# Run on macOS
cd macos && go run main.go
```

## Usage

### Create a New Project

```bash
goflow new <project_name> [flags]
# or
goflow create <project_name> [flags]  # alias for 'new'
```

**Flags:**
- `--platforms` - Target platforms (default: `macos,windows,linux`)
  - Options: `macos`, `windows`, `linux`, `web`
- `--template` - Project template (default: `default`)
  - Options: `default`, `material`, `minimal`
- `--org` - Organization for module path (default: `com.example`)

**Examples:**

```bash
# Basic project with all platforms
goflow create myapp

# macOS only with material template
goflow create myapp --platforms=macos --template=material

# Custom organization
goflow create myapp --org=io.github.username

# Minimal app for quick prototyping
goflow create quickstart --template=minimal --platforms=macos
```

### Other Commands

```bash
goflow version    # Show CLI version
goflow help       # Show help information
```

## Project Templates

### default
Counter app demonstrating:
- Basic widgets (Text, Container, Column, Center)
- Signal-based state management
- Widget composition
- Event handling

**Best for:** Learning GoFlow, starting new projects

### material
Material Design-inspired app with:
- Styled containers and colors
- App bar layout
- Card-like components
- Polished UI

**Best for:** Production-ready apps, design-focused projects

### minimal
Bare-bones "Hello World" with:
- Single centered text widget
- No state management
- Minimal code

**Best for:** Quick prototyping, learning basics

## Generated Project Structure

```
myapp/
├── lib/                    # Shared application code
│   ├── main.go             # App entry point (package lib)
│   ├── screens/            # Screen widgets
│   ├── widgets/            # Custom widgets
│   ├── models/             # Data models
│   ├── services/           # Business logic
│   └── state/              # State management
├── macos/                  # macOS platform runner
│   ├── main.go
│   ├── runner/
│   └── assets/
├── windows/                # Windows platform runner
│   └── ...
├── linux/                  # Linux platform runner
│   └── ...
├── assets/                 # Shared assets
│   ├── fonts/
│   ├── images/
│   └── icons/
├── test/                   # Tests
├── go.mod                  # Go module
├── .gitignore
└── README.md
```

## Development Workflow

### Running Your App

```bash
cd myapp

# Run via platform runner (recommended)
cd macos  # or windows, or linux
go run main.go
```

### Building Your App

```bash
# macOS
cd macos && go build -o MyApp

# Windows
cd windows && go build -o MyApp.exe

# Linux
cd linux && go build -o myapp
```

### Adding Dependencies

```bash
# Install dependencies
go mod tidy

# Add a new dependency
go get github.com/some/package
```

## Binary Size

The GoFlow CLI binary is optimized for size:

- **With symbols**: ~3.8MB
- **Optimized** (using `-ldflags="-s -w"`): **~2.6MB**
- **Embedded templates**: ~28KB

The CLI uses embedded templates for reliability and offline support, with minimal size impact.

## Advanced Usage

### Custom Module Paths

```bash
# Use your GitHub username
goflow create myapp --org=github.com/username

# Use company domain
goflow create myapp --org=com.company.team

# Generated module: com.company.team/myapp
```

### Platform-Specific Projects

```bash
# macOS only (for Mac App Store apps)
goflow create macapp --platforms=macos

# Cross-platform desktop
goflow create desktop-app --platforms=macos,windows,linux

# Experimental web support
goflow create webapp --platforms=web
```

### Project Naming

Valid project names:
- Must start with a letter
- Can contain letters, numbers, `_`, `-`
- Converted to Go identifiers (e.g., `my-app` → `MyApp`)

```bash
goflow create my-app         # ✅ Valid (becomes MyApp)
goflow create user_dashboard # ✅ Valid (becomes UserDashboard)
goflow create todo-list-app  # ✅ Valid (becomes TodoListApp)
goflow create 123app         # ❌ Invalid (starts with number)
```

## Troubleshooting

### "command not found: goflow"

Add Go bin to PATH:
```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc
```

### "module not found" errors in generated project

The generated projects require GoFlow. For local development:

```bash
cd myapp
# Add replace directive to use local GoFlow
echo 'replace github.com/base-go/GoFlow => /path/to/GoFlow' >> go.mod
go mod tidy
```

For published GoFlow (when available):
```bash
go mod tidy  # Will fetch from GitHub
```

### Templates not found

Make sure you installed via `go install` or built from source with the `templates/` directory present. The templates are embedded into the binary at build time.

## Comparison with Flutter CLI

| Feature | Flutter | GoFlow |
|---------|---------|--------|
| Create project | `flutter create` | `goflow create` |
| Platforms | iOS, Android, Web, Desktop | Desktop (macOS, Windows, Linux), Web |
| Templates | Multiple | default, material, minimal |
| Binary size | ~100MB+ | ~2.6MB |
| Installation | SDK download | `go install` |
| Offline support | Yes | Yes (embedded templates) |

## Contributing

Found a bug or want to add a feature?

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## See Also

- [GoFlow Main README](../../README.md)
- [Project Structure Guide](../../docs/PROJECT_STRUCTURE.md)
- [CLI Documentation](../../docs/CLI.md)
- [Platform Integration Roadmap](../../docs/PLATFORM_INTEGRATION.md)
- [Architecture Guide](../../ARCHITECTURE.md)
- [Examples](../../examples/)

## License

MIT License - see [LICENSE](../../LICENSE) for details
