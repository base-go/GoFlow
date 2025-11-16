# GoFlow CLI Documentation

The GoFlow CLI is a command-line tool for creating and managing GoFlow projects, inspired by Flutter's CLI.

## Installation

### Via `go install` (Recommended)

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

This installs the `goflow` binary to your `$GOPATH/bin`. Ensure it's in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Build from Source

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

## Commands

### `goflow create`

Creates a new GoFlow project with the specified configuration.

#### Syntax

```bash
goflow create <project_name> [flags]
```

#### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--platforms` | string | `macos,windows,linux` | Comma-separated list of target platforms |
| `--template` | string | `default` | Project template to use |
| `--org` | string | `com.example` | Organization name for Go module path |

#### Available Platforms

- `macos` - macOS desktop (GLFW + WGPU)
- `windows` - Windows desktop (GLFW + WGPU)
- `linux` - Linux desktop (GLFW + WGPU)
- `web` - WebAssembly (future)

####Available Templates

**default** - Standard counter app template
- Demonstrates basic widgets (Text, Container, Column, Center)
- Shows signal-based state management
- Includes increment function
- Good starting point for most apps

**minimal** - Minimal "Hello World" template
- Simplest possible GoFlow app
- Just displays centered text
- Perfect for learning the basics
- Quick prototyping

**material** - Material Design-inspired template
- Card-like layout with header
- Colored containers and styling
- More polished UI example
- Best for production-like apps

#### Examples

**Basic Project Creation**
```bash
# Create with all default platforms
goflow create myapp

# Created directory structure:
# myapp/
# ├── lib/
# ├── macos/
# ├── windows/
# ├── linux/
# ├── assets/
# └── go.mod
```

**Platform-Specific Creation**
```bash
# Create for macOS only
goflow create myapp --platforms=macos

# Create for macOS and Linux
goflow create myapp --platforms=macos,linux

# Create with web support (future)
goflow create myapp --platforms=macos,web
```

**Template Selection**
```bash
# Create with minimal template
goflow create quickstart --template=minimal

# Create with material design template
goflow create myapp --template=material

# Create with default template (explicit)
goflow create myapp --template=default
```

**Custom Organization**
```bash
# Use custom organization
goflow create myapp --org=com.mycompany

# Creates module: com.mycompany/myapp
# Instead of: com.example/myapp
```

**Combined Flags**
```bash
goflow create myapp \
  --platforms=macos,linux \
  --template=material \
  --org=io.github.username
```

### `goflow version`

Display the GoFlow CLI version.

```bash
goflow version
# or
goflow --version
# or
goflow -v
```

### `goflow help`

Display help information about the CLI and its commands.

```bash
goflow help
# or
goflow --help
# or
goflow -h
```

## Project Structure

After creating a project with `goflow create myapp`, you'll have:

```
myapp/
├── lib/                        # Shared application code
│   ├── main.go                 # App entry point (package lib)
│   ├── screens/                # Screen widgets
│   ├── widgets/                # Custom widgets
│   ├── models/                 # Data models
│   ├── services/               # Business logic
│   └── state/                  # State management
├── macos/                      # macOS platform code
│   ├── main.go                 # Platform runner
│   ├── runner/                 # Platform-specific code
│   └── assets/                 # Platform assets
├── windows/                    # Windows platform code
│   └── ...
├── linux/                      # Linux platform code
│   └── ...
├── assets/                     # Shared assets
│   ├── fonts/
│   ├── images/
│   └── icons/
├── test/                       # Tests
├── go.mod                      # Go module definition
├── .gitignore
└── README.md
```

## Working with Created Projects

### Running Your App

```bash
cd myapp

# Run via platform runner (recommended)
cd macos  # or windows, or linux
go run main.go

# The platform runner will call lib.Run()
```

### Building Your App

```bash
cd myapp

# macOS
cd macos
go build -o MyApp
./MyApp

# Windows
cd windows
go build -o myapp.exe
myapp.exe

# Linux
cd linux
go build -o myapp
./myapp
```

### Project Dependencies

After creating a project, set up dependencies:

```bash
cd myapp

# If using local GoFlow (development)
echo "\nreplace github.com/base-go/GoFlow => /path/to/GoFlow" >> go.mod
go mod tidy

# If using published GoFlow (future)
go mod tidy
```

## Template Details

### Default Template

Generated `lib/main.go`:
- `DemoApp` struct with counter signal
- `Build()` method creating UI
- `Increment()` method for interactions
- `Run()` function for platform runners
- Spacer helper function

### Minimal Template

Generated `lib/main.go`:
- Simple `App` struct
- Basic `Build()` method
- `Run()` function
- No state management
- Ideal for learning

### Material Template

Generated `lib/main.go`:
- `MaterialApp` with styled UI
- App bar-style header
- Card-like content container
- Counter with Material styling
- `Run()` function

## Project Naming

### Valid Project Names

- Must start with a letter
- Can contain letters, numbers, underscores, hyphens
- Will be converted to valid Go identifiers:
  - `my-app` → `MyApp` (type name)
  - `todo_list` → `TodoList` (type name)
  - `user-auth-service` → `UserAuthService` (type name)

### Examples

```bash
goflow create myapp          # ✅ Valid
goflow create my-app         # ✅ Valid (becomes MyApp)
goflow create my_cool_app    # ✅ Valid (becomes MyCoolApp)
goflow create user-dashboard # ✅ Valid (becomes UserDashboard)
goflow create 123app         # ❌ Invalid (starts with number)
goflow create my app         # ❌ Invalid (contains space)
```

## Troubleshooting

### "command not found: goflow"

The CLI is not in your PATH. Either:
1. Run from the `cmd/goflow` directory: `./goflow create myapp`
2. Add to PATH: `export PATH=$PATH:/path/to/GoFlow/cmd/goflow`
3. Install globally: `sudo mv goflow /usr/local/bin/`

### "module not found" errors

Add a replace directive to use local GoFlow:
```bash
cd myapp
echo "\nreplace github.com/base-go/GoFlow => /path/to/GoFlow" >> go.mod
go mod tidy
```

### Platform runner won't compile

Make sure your `lib/main.go` has:
- `package lib` (not `package main`)
- Exported `Run()` function
- Valid Go code (no syntax errors)

## Next Steps

After creating a project:

1. **Explore the structure** - Understand lib/, platform dirs, and assets/
2. **Modify lib/main.go** - Customize your app's UI and logic
3. **Run your app** - Use platform runners to test
4. **Add features** - Create new widgets, screens, and state
5. **Build for distribution** - Compile platform-specific binaries

## See Also

- [Project Structure Guide](./PROJECT_STRUCTURE.md)
- [Architecture Documentation](../ARCHITECTURE.md)
- [Examples](../examples/)
- [Main README](../README.md)
