# Flutter-Inspired Design

GoFlow's project structure and CLI are inspired by Flutter, providing a familiar experience for developers who know Flutter while staying true to Go idioms.

## Command Comparison

| Flutter | GoFlow | Description |
|---------|--------|-------------|
| `flutter create` | `goflow new` | Create new project |
| `flutter pub get` | `go mod tidy` | Get dependencies |
| `flutter run` | `cd macos && go run main.go` | Run on platform |
| `flutter build` | `cd macos && go build` | Build for platform |

## Project Structure Comparison

### Flutter Project
```
my_app/
├── lib/
│   └── main.dart
├── macos/
│   ├── Runner.xcodeproj
│   └── Runner/
├── linux/
│   ├── CMakeLists.txt
│   └── runner/
├── windows/
│   ├── CMakeLists.txt
│   └── runner/
├── android/
├── ios/
├── web/
├── test/
├── assets/
├── pubspec.yaml
└── analysis_options.yaml
```

### GoFlow Project
```
my_app/
├── lib/
│   └── main.go              # Go instead of Dart
├── macos/
│   ├── main.go              # Go runner (will become Xcode project)
│   ├── runner/
│   └── .gitignore
├── linux/
│   ├── main.go              # Go runner (will become CMake)
│   ├── runner/
│   └── .gitignore
├── windows/
│   ├── main.go              # Go runner (will become CMake)
│   ├── runner/
│   └── .gitignore
├── test/
├── assets/
│   ├── fonts/
│   ├── images/
│   └── icons/
├── goflow.yaml              # Like pubspec.yaml
├── go.mod                   # Go dependencies
└── analysis_options.yaml    # Future linting config
```

## Configuration Files

### pubspec.yaml vs goflow.yaml

**Flutter's pubspec.yaml:**
```yaml
name: my_app
description: A Flutter application
version: 1.0.0+1

environment:
  sdk: ^3.10.0

dependencies:
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.8

flutter:
  uses-material-design: true
  assets:
    - images/
  fonts:
    - family: Roboto
      fonts:
        - asset: fonts/Roboto-Regular.ttf
```

**GoFlow's goflow.yaml:**
```yaml
name: my_app
description: A GoFlow application
version: 0.1.0+1

# Go module path
module: com.example/my_app

# Minimum Go version
environment:
  go: "1.21"

# Dependencies
dependencies:
  goflow:
    git: https://github.com/base-go/GoFlow
    version: v0.1.0

# Platform configuration
platforms:
  - macos
  - windows
  - linux

# Assets
assets:
  - assets/images/
  - assets/fonts/

# Fonts
fonts:
  - family: Roboto
    fonts:
      - asset: assets/fonts/Roboto-Regular.ttf
      - asset: assets/fonts/Roboto-Bold.ttf
        weight: 700
```

## Code Comparison

### Flutter (Dart)
```dart
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text('$_counter'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        child: Icon(Icons.add),
      ),
    );
  }
}
```

### GoFlow (Go)
```go
package lib

import (
	"fmt"
	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

type MyApp struct {
	goflow.BaseWidget
	counter *signals.Signal[int]
}

func NewMyApp() *MyApp {
	return &MyApp{
		counter: signals.New(0),
	}
}

func (app *MyApp) Build(ctx goflow.BuildContext) goflow.Widget {
	count := app.counter.Get()

	return widgets.NewCenter(
		widgets.NewText(fmt.Sprintf("%d", count)),
	)
}

func (app *MyApp) Increment() {
	app.counter.Update(func(v int) int { return v + 1 })
}

func Run() {
	app := NewMyApp()
	goflow.RunApp(app)
}
```

## Key Differences

### Similarities with Flutter
1. **Project Structure**: Same directory layout with lib/, platform dirs, assets/
2. **Configuration File**: goflow.yaml mirrors pubspec.yaml
3. **Platform Separation**: Platform-specific code in separate directories
4. **CLI Commands**: Similar command structure (new, run, build)
5. **Widget System**: Declarative UI with widget trees
6. **Build Method**: Widgets have Build() methods

### Differences from Flutter
1. **Language**: Go instead of Dart
2. **State Management**: Signals instead of StatefulWidget/setState
3. **Single Widget Type**: No StatelessWidget vs StatefulWidget distinction
4. **Dependencies**: go.mod for packages, not pub
5. **Type System**: Go's static typing with generics for signals
6. **No Hot Reload**: (yet - planned for future)
7. **Desktop Focus**: No iOS/Android (desktop-first framework)

## Platform Integration

### Current State (v0.1.0)
Platform directories contain simple Go runners that call into lib/ code:

```go
// macos/main.go
package main

import "your.module/lib"

func main() {
    lib.Run()
}
```

### Future State (Planned)
Platform directories will contain native projects:

**macOS**: Xcode project with Swift/Objective-C calling into Go
**Windows**: Visual Studio project with C++ calling into Go DLL
**Linux**: CMake project with C++ calling into Go shared library

This matches Flutter's model where each platform has native launcher code.

## Benefits of Flutter-Inspired Design

1. **Familiarity**: Developers coming from Flutter feel at home
2. **Best Practices**: Proven project structure from successful framework
3. **Clear Separation**: Platform code vs shared app code
4. **Scalability**: Structure works for small and large projects
5. **Asset Management**: Centralized assets/ directory
6. **Configuration**: Single source of truth (goflow.yaml)
7. **Consistency**: Similar patterns across all platforms

## Migration Path (Flutter → GoFlow)

For Flutter developers moving to GoFlow:

1. **Project Setup**: `goflow new` instead of `flutter create`
2. **Dependencies**: Use go.mod instead of pubspec.yaml for packages
3. **State**: Replace StatefulWidget with Signals
4. **Widgets**: Similar widget names (Text, Container, Column, etc.)
5. **Build**: `cd platform && go build` instead of `flutter build`
6. **Run**: `cd platform && go run main.go` instead of `flutter run`

## See Also

- [Project Structure Guide](./PROJECT_STRUCTURE.md)
- [Platform Integration Roadmap](./PLATFORM_INTEGRATION.md)
- [CLI Documentation](./CLI.md)
- [Architecture Guide](../ARCHITECTURE.md)
