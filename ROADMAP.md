# GoFlow Roadmap

Strategic plan for GoFlow development, organized by priority and implementation phases.

## Current Status (v0.2.0)

### ✅ Implemented
- **CLI**: Mamba-based CLI with 8 commands (new, run, build, test, doctor, clean, analyze, format)
- **Core Widgets**: Layout (Row, Column, Stack, Center, Align, Flexible, Wrap, Table)
- **Display Widgets**: Text, Icon, Image, Avatar, Badge, Chip, Divider, Progress
- **Input**: TextField, Focus management, Dropdown
- **Interaction**: GestureDetector, Scrolling (ListView, PageView, ScrollController)
- **Navigation**: GetX-style routing (Get.To, Get.Back, dialogs, sheets)
- **Testing**: WidgetTester, GoldenTester, IntegrationTester, GestureSimulator
- **Hot Reload**: File watching, state preservation, widget tree diffing
- **Design Systems**: Material Design, Cupertino, Adaptive widgets
- **Rendering**: macOS native backend with Core Graphics

### 🚧 Partially Implemented
- **Animation**: AnimatedContainer, AnimatedOpacity (basic animations only)
- **Forms**: TextField only (no Form widget, validation, etc.)
- **Platform Integration**: macOS only

### ❌ Not Implemented
- Windows backend
- Linux backend
- Web backend
- Advanced animations
- Theme system
- Asset management
- Persistence layer
- Network utilities

---

## Phase 1: Core Platform Support (Q1 2025)

Priority: **CRITICAL** - Need multi-platform support

### 1.1 Windows Backend
**Goal**: Run GoFlow apps on Windows

**Tasks:**
- [ ] Window creation with Win32 API
- [ ] Direct2D rendering backend
- [ ] Event handling (mouse, keyboard, touch)
- [ ] Text rendering with DirectWrite
- [ ] File system integration
- [ ] Build system integration (`goflow build windows`)

**Example Output:**
```bash
goflow build windows
# → myapp.exe (Windows executable)
```

### 1.2 Linux Backend
**Goal**: Run GoFlow apps on Linux

**Tasks:**
- [ ] GTK3/GTK4 window management
- [ ] Cairo rendering backend
- [ ] Event handling
- [ ] Text rendering with Pango
- [ ] File system integration
- [ ] Build system integration (`goflow build linux`)

**Example Output:**
```bash
goflow build linux
# → myapp (Linux executable)
```

### 1.3 Platform Embedding (All Platforms)
**Goal**: Implement Flutter-style platform embedding

**Tasks:**
- [ ] Create GoFlowRuntime.framework for macOS
- [ ] Compile Go app to shared library (.dylib/.dll/.so)
- [ ] Generate platform folders (macos/, windows/, linux/)
- [ ] Auto-generate xcconfig/build files
- [ ] Create proper app bundles (.app, .exe, .AppImage)

**Reference:** See `docs/MACOS_EMBEDDING.md`

---

## Phase 2: Essential Widgets & Features (Q2 2025)

Priority: **HIGH** - Needed for real apps

### 2.1 Animation System
**Goal**: Flutter-quality animations

**Widgets:**
- [ ] AnimationController
- [ ] Tween (IntTween, DoubleTween, ColorTween)
- [ ] CurvedAnimation
- [ ] AnimatedBuilder
- [ ] ImplicitlyAnimatedWidget
- [ ] AnimatedCrossFade
- [ ] Hero animations
- [ ] Page transitions

**Example:**
```go
controller := NewAnimationController(duration: time.Millisecond * 300)
animation := NewTween(begin: 0.0, end: 1.0).Animate(controller)

AnimatedBuilder(
    animation: animation,
    builder: func(ctx BuildContext, child Widget) Widget {
        return Opacity(
            opacity: animation.Value(),
            child: child,
        )
    },
    child: Text("Fade In"),
)
```

### 2.2 Form System
**Goal**: Complete form handling with validation

**Widgets:**
- [ ] Form widget
- [ ] FormField
- [ ] TextFormField
- [ ] DropdownButtonFormField
- [ ] CheckboxFormField
- [ ] RadioFormField
- [ ] SwitchFormField
- [ ] SliderFormField
- [ ] DatePickerFormField
- [ ] TimePickerFormField

**Features:**
- [ ] Form validation
- [ ] Auto-validation
- [ ] Custom validators
- [ ] Form state management
- [ ] Error display

**Example:**
```go
formKey := NewFormKey()

Form(
    key: formKey,
    child: Column(
        children: []Widget{
            TextFormField(
                validator: func(value string) *string {
                    if value == "" {
                        err := "Email required"
                        return &err
                    }
                    return nil
                },
            ),
            ElevatedButton(
                onPressed: func() {
                    if formKey.CurrentState().Validate() {
                        // Submit form
                    }
                },
                child: Text("Submit"),
            ),
        },
    ),
)
```

### 2.3 Advanced Input Widgets
**Goal**: Rich input components

**Widgets:**
- [ ] Checkbox
- [ ] Radio
- [ ] Switch
- [ ] Slider
- [ ] RangeSlider
- [ ] DatePicker
- [ ] TimePicker
- [ ] ColorPicker
- [ ] FilePicker (platform channels)
- [ ] SearchBar
- [ ] Autocomplete

### 2.4 Data Display Widgets
**Goal**: Display complex data

**Widgets:**
- [ ] DataTable (sortable, selectable)
- [ ] Card (enhanced)
- [ ] ExpansionPanel
- [ ] ExpansionTile
- [ ] Stepper
- [ ] Timeline
- [ ] TreeView
- [ ] Charts (line, bar, pie)

---

## Phase 3: Developer Experience (Q3 2025)

Priority: **HIGH** - Improve development workflow

### 3.1 Theme System
**Goal**: Comprehensive theming support

**Features:**
- [ ] ThemeData
- [ ] ColorScheme
- [ ] TextTheme
- [ ] Dark mode support
- [ ] Custom themes
- [ ] Theme extensions
- [ ] Runtime theme switching

**Example:**
```go
MaterialApp(
    theme: ThemeData(
        colorScheme: ColorScheme.FromSeed(Colors.Blue),
        textTheme: TextTheme{
            HeadlineLarge: TextStyle{FontSize: 32},
            BodyMedium: TextStyle{FontSize: 16},
        },
    ),
    darkTheme: ThemeData.Dark(),
    themeMode: ThemeMode.System, // Auto-detect
    home: HomePage(),
)
```

### 3.2 Asset Management
**Goal**: Flutter-style asset bundling

**Features:**
- [ ] goflow.yaml configuration (like pubspec.yaml)
- [ ] Asset bundling
- [ ] Image assets with multiple resolutions
- [ ] Font loading
- [ ] JSON/data file loading
- [ ] Asset generation

**goflow.yaml:**
```yaml
name: myapp
version: 1.0.0

dependencies:
  goflow: ^0.2.0

assets:
  - assets/images/
  - assets/fonts/

fonts:
  - family: Roboto
    fonts:
      - asset: assets/fonts/Roboto-Regular.ttf
      - asset: assets/fonts/Roboto-Bold.ttf
        weight: 700
```

**Usage:**
```go
Image.Asset("assets/images/logo.png")
```

### 3.3 Hot Reload Enhancement
**Goal**: Near-instant code reload

**Features:**
- [ ] Full hot reload integration with `goflow run`
- [ ] State preservation across reloads
- [ ] Widget tree diffing and partial rebuilds
- [ ] Error recovery
- [ ] Hot restart (full reload)
- [ ] File watcher improvements

**Current:** Basic hot reload exists, needs CLI integration

### 3.4 DevTools
**Goal**: Developer debugging tools

**Features:**
- [ ] Widget inspector
- [ ] Performance profiler
- [ ] Network inspector
- [ ] State viewer
- [ ] Layout debugger
- [ ] Console logging
- [ ] Timeline view

---

## Phase 4: Advanced Features (Q4 2025)

Priority: **MEDIUM** - Nice to have

### 4.1 Web Backend
**Goal**: Run GoFlow apps in browser

**Technologies:**
- [ ] WebAssembly compilation
- [ ] Canvas rendering
- [ ] DOM integration
- [ ] JavaScript interop
- [ ] PWA support
- [ ] Web-specific widgets

**Build:**
```bash
goflow build web
# → Outputs index.html + wasm files
```

### 4.2 Mobile Support (iOS/Android)
**Goal**: Mobile app deployment

**iOS:**
- [ ] UIKit integration
- [ ] Metal rendering
- [ ] Touch events
- [ ] iOS-specific widgets

**Android:**
- [ ] Android View system
- [ ] Skia rendering
- [ ] Touch events
- [ ] Android-specific widgets

### 4.3 Persistence Layer
**Goal**: Easy data persistence

**Features:**
- [ ] SharedPreferences (key-value storage)
- [ ] SQLite integration
- [ ] Secure storage
- [ ] File I/O helpers
- [ ] Cache management

**Example:**
```go
prefs := SharedPreferences.GetInstance()
prefs.SetString("username", "john")
username := prefs.GetString("username", defaultValue: "")
```

### 4.4 Network Utilities
**Goal**: HTTP client and REST helpers

**Features:**
- [ ] HTTP client with middleware
- [ ] JSON serialization helpers
- [ ] REST API builder
- [ ] WebSocket support
- [ ] GraphQL client
- [ ] Network monitoring

**Example:**
```go
response, err := http.Get("https://api.example.com/users")
var users []User
json.Unmarshal(response.Body, &users)
```

### 4.5 State Management Extensions
**Goal**: Advanced state patterns

**Patterns:**
- [ ] Provider pattern
- [ ] BLoC pattern
- [ ] MobX-style observables
- [ ] Redux-style store
- [ ] Riverpod-style providers

**Example:**
```go
// Provider pattern
userProvider := Provider.New(func() *User {
    return fetchUser()
})

Consumer(
    provider: userProvider,
    builder: func(ctx BuildContext, user *User) Widget {
        return Text(user.Name)
    },
)
```

---

## Phase 5: Ecosystem & Tooling (2026)

Priority: **LOW** - Long-term vision

### 5.1 Plugin System
**Goal**: Community extensions

**Features:**
- [ ] Plugin API
- [ ] Platform channels (Go ↔ Native)
- [ ] Plugin registry
- [ ] Package manager integration
- [ ] Plugin discovery

**Example:**
```go
// Camera plugin
camera := Camera.Initialize()
image := camera.TakePicture()
```

### 5.2 IDE Extensions
**Goal**: Better IDE support

**Extensions:**
- [ ] VS Code extension (syntax highlighting, snippets)
- [ ] GoLand plugin
- [ ] Widget preview
- [ ] Code generation
- [ ] Refactoring tools

### 5.3 CI/CD Integration
**Goal**: Automated builds and deployment

**Features:**
- [ ] GitHub Actions templates
- [ ] GitLab CI templates
- [ ] Docker images for builds
- [ ] Automated testing in CI
- [ ] Release automation

### 5.4 Documentation Site
**Goal**: Comprehensive online docs

**Features:**
- [ ] Interactive examples
- [ ] API documentation
- [ ] Video tutorials
- [ ] Code snippets
- [ ] Community showcase

---

## Feature Requests & Community Input

### Most Requested Features

1. **Windows/Linux Support** (20+ requests)
   → Planned for Phase 1

2. **Better Animation System** (15+ requests)
   → Planned for Phase 2

3. **Theme Customization** (12+ requests)
   → Planned for Phase 3

4. **Form Validation** (10+ requests)
   → Planned for Phase 2

5. **Web Support** (8+ requests)
   → Planned for Phase 4

### Community Contributions

Want to contribute? Check:
- Issues labeled `good-first-issue`
- Issues labeled `help-wanted`
- [CONTRIBUTING.md](CONTRIBUTING.md)

---

## Versioning Strategy

Following [Semantic Versioning](https://semver.org/):

- **v0.x.x** - Pre-1.0 development (current)
- **v1.0.0** - Stable API, multi-platform support (Phase 1-2 complete)
- **v1.x.x** - Minor features, backwards-compatible
- **v2.0.0** - Breaking changes (if needed)

### Upcoming Releases

- **v0.3.0** (Q1 2025) - Windows/Linux support
- **v0.4.0** (Q2 2025) - Animation & Forms
- **v0.5.0** (Q3 2025) - Theme system & Assets
- **v1.0.0** (Q4 2025) - Stable release

---

## Get Involved

- **GitHub**: https://github.com/base-go/GoFlow
- **Discussions**: https://github.com/base-go/GoFlow/discussions
- **Issues**: https://github.com/base-go/GoFlow/issues
- **Discord**: (coming soon)

## Feedback

Have ideas for new features? Open an issue or discussion!

---

*Last updated: 2025-01-15*
