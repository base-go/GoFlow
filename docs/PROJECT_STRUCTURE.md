# GoFlow Project Structure

GoFlow follows a Flutter-inspired project structure with platform-specific directories for desktop and web targets.

## Standard GoFlow Project Layout

```
myapp/
├── lib/                        # Shared application code (cross-platform)
│   ├── main.go                 # Application entry point
│   ├── screens/                # Screen/page widgets
│   │   ├── home_screen.go
│   │   └── settings_screen.go
│   ├── widgets/                # Custom reusable widgets
│   │   └── custom_button.go
│   ├── models/                 # Data models
│   │   └── user.go
│   ├── services/               # Business logic & services
│   │   └── api_service.go
│   └── state/                  # State management
│       └── app_state.go
├── macos/                      # macOS platform-specific code
│   ├── main.go                 # macOS entry point
│   ├── runner/                 # Platform runner
│   │   └── window.go
│   └── assets/                 # macOS-specific assets
├── windows/                    # Windows platform-specific code
│   ├── main.go                 # Windows entry point
│   ├── runner/
│   │   └── window.go
│   └── assets/
├── linux/                      # Linux platform-specific code
│   ├── main.go                 # Linux entry point
│   ├── runner/
│   │   └── window.go
│   └── assets/
├── web/                        # Web/WASM platform (future)
│   ├── main.go
│   └── index.html
├── assets/                     # Shared assets (fonts, images, etc.)
│   ├── fonts/
│   ├── images/
│   └── icons/
├── test/                       # Tests
│   └── widget_test.go
├── go.mod                      # Go module definition
├── go.sum                      # Go dependency checksums
├── .gitignore
└── README.md
```

## Directory Descriptions

### `lib/`
Contains all shared, platform-agnostic application code. This is where you spend most of your development time.

- **main.go**: Defines your root widget and app logic
- **screens/**: Full-screen views/pages
- **widgets/**: Custom reusable UI components
- **models/**: Data structures and business entities
- **services/**: API clients, database access, etc.
- **state/**: Signals and state management

### `macos/`, `windows/`, `linux/`
Platform-specific runners that:
1. Initialize the platform window system (GLFW)
2. Set up the rendering backend (WGPU)
3. Create the event loop
4. Import and run your `lib/main.go` application

Each platform directory has:
- **main.go**: Platform entry point that creates window and starts app
- **runner/**: Platform-specific window management code
- **assets/**: Platform-specific resources (icons, manifests, etc.)

### `web/`
Future web platform support via WebAssembly:
- Compiles Go to WASM
- Runs GoFlow in the browser
- Uses WebGPU for rendering

### `assets/`
Shared application resources:
- **fonts/**: TTF, OTF font files
- **images/**: PNG, JPG, SVG images
- **icons/**: App icons for different platforms

### `test/`
Unit tests, widget tests, and integration tests.

## Building for Different Platforms

### macOS
```bash
cd macos
go build -o MyApp.app/Contents/MacOS/myapp
```

### Windows
```bash
cd windows
go build -o myapp.exe
```

### Linux
```bash
cd linux
go build -o myapp
```

### Web (Future)
```bash
cd web
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Creating a New Project

Use the GoFlow CLI:

```bash
# Create a new project
goflow create myapp

# Create with specific platforms only
goflow create myapp --platforms=macos,linux

# Create with a template
goflow create myapp --template=material
```

## Comparison with Flutter

| Flutter | GoFlow | Purpose |
|---------|--------|---------|
| `lib/` | `lib/` | Shared app code |
| `macos/` | `macos/` | macOS platform code |
| `windows/` | `windows/` | Windows platform code |
| `linux/` | `linux/` | Linux platform code |
| `web/` | `web/` | Web platform code |
| `ios/` | ❌ | Not supported (desktop focus) |
| `android/` | ❌ | Not supported (desktop focus) |
| `pubspec.yaml` | `go.mod` | Dependencies |
| `assets/` | `assets/` | Resources |
| `test/` | `test/` | Tests |

## Architecture Patterns

### Feature-First (Recommended for large apps)
```
lib/
├── main.go
├── auth/
│   ├── login_screen.go
│   ├── signup_screen.go
│   ├── auth_service.go
│   └── auth_state.go
├── dashboard/
│   ├── dashboard_screen.go
│   ├── dashboard_widgets.go
│   └── dashboard_state.go
└── shared/
    ├── widgets/
    └── utils/
```

### Layer-First (Recommended for small apps)
```
lib/
├── main.go
├── screens/
│   ├── login_screen.go
│   ├── dashboard_screen.go
│   └── settings_screen.go
├── widgets/
│   └── custom_widgets.go
├── services/
│   └── api_service.go
└── state/
    └── app_state.go
```

## Best Practices

1. **Keep lib/ platform-agnostic**: No platform-specific imports in lib/
2. **Use signals for state**: Leverage GoFlow's reactive state management
3. **Organize by feature for large apps**: Group related code together
4. **Share assets**: Put common assets in root assets/, platform-specific in platform dirs
5. **Test your widgets**: Write tests in test/ directory
6. **Follow Go conventions**: Use Go naming, file structure, and idioms

## Next Steps

- See [CREATING_APPS.md](./CREATING_APPS.md) for detailed app creation guide
- See [ARCHITECTURE.md](../ARCHITECTURE.md) for framework architecture
- See [examples/](../examples/) for sample applications
