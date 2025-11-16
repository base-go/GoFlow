# Platform Embedding Architecture

How Flutter embeds into native platforms (and how GoFlow can do the same).

## Overview

Flutter's iOS/Android/macOS folders don't contain Dart code because they follow an **embedding architecture** where:

1. **Native Shell** - Platform-specific wrapper app (Swift/Kotlin/etc.)
2. **Framework Bridge** - Links to the framework runtime
3. **App Code** - Lives in `lib/` directory (Dart for Flutter, Go for GoFlow)
4. **Build Process** - Generates and links everything at build time

## Flutter's iOS Structure

```
ios/
├── Runner/                    # Native iOS app shell
│   ├── AppDelegate.swift     # App entry point
│   ├── Info.plist            # iOS app configuration
│   ├── Assets.xcassets       # Native assets
│   └── Base.lproj/           # Storyboards (UI)
├── Flutter/                   # Flutter framework integration
│   ├── Generated.xcconfig    # Auto-generated build config
│   ├── Debug.xcconfig        # Debug build settings
│   ├── Release.xcconfig      # Release build settings
│   └── ephemeral/            # Temporary build artifacts
└── Runner.xcodeproj/         # Xcode project
```

### How It Works

#### 1. AppDelegate.swift (Entry Point)

```swift
import Flutter
import UIKit

@main
@objc class AppDelegate: FlutterAppDelegate {
  override func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?
  ) -> Bool {
    GeneratedPluginRegistrant.register(with: self)
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}
```

**What happens:**
1. App launches with native iOS code
2. Inherits from `FlutterAppDelegate` (from Flutter.framework)
3. Flutter engine starts automatically
4. Loads `lib/main.dart` (specified in Generated.xcconfig)
5. Runs Dart code in Flutter engine

#### 2. Generated.xcconfig (Build Configuration)

```
FLUTTER_ROOT=/Users/user/flutter
FLUTTER_APPLICATION_PATH=/path/to/project
FLUTTER_TARGET=lib/main.dart          # Entry point
FLUTTER_BUILD_DIR=build
```

**Key insight:** The Dart code location is configured in build settings, not hardcoded.

#### 3. Build Process

```bash
# When you run "flutter run" or build in Xcode:

1. flutter build ios
   ↓
2. Compiles lib/main.dart → App.framework
   ↓
3. Copies Flutter.framework (engine)
   ↓
4. Xcode links both frameworks
   ↓
5. Native iOS app loads frameworks at runtime
   ↓
6. App.framework contains all Dart code
```

**No Dart code in ios/ because:**
- Dart code is compiled into `App.framework`
- Framework is linked at build time
- Native shell just loads the framework

## GoFlow Can Use the Same Architecture!

### Proposed GoFlow Structure

```
macos/                           # Native macOS app shell
├── GoFlow/                      # GoFlow runtime integration
│   ├── Generated.xcconfig       # Auto-generated config
│   ├── Debug.xcconfig
│   └── Release.xcconfig
├── Runner/                      # Native app wrapper
│   ├── AppDelegate.swift
│   ├── Info.plist
│   └── Assets.xcassets
└── Runner.xcodeproj/

lib/                             # GoFlow app code (Go!)
├── main.go                      # App entry point
└── widgets/                     # UI code
    └── home.go
```

### GoFlow AppDelegate.swift

```swift
import Cocoa
import GoFlowRuntime  // Our Go runtime as framework

@main
class AppDelegate: NSObject, NSApplicationDelegate {
    var window: NSWindow!
    var goflowEngine: GoFlowEngine!

    func applicationDidFinishLaunching(_ notification: Notification) {
        // Initialize GoFlow runtime
        goflowEngine = GoFlowEngine()

        // Load Go app code
        // The path is configured in Generated.xcconfig
        goflowEngine.loadApp(fromPath: getAppPath())

        // Create window with GoFlow view
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 800, height: 600),
            styleMask: [.titled, .closable, .miniaturizable, .resizable],
            backing: .buffered,
            defer: false
        )
        window.contentView = goflowEngine.view
        window.makeKeyAndOrderFront(nil)

        // Start the Go app
        goflowEngine.start()
    }

    private func getAppPath() -> String {
        // Read from bundle or Generated.xcconfig
        return Bundle.main.path(forResource: "app", ofType: "dylib")!
    }
}
```

### How GoFlow Would Build

```bash
# goflow build macos

1. Compile lib/main.go → libapp.dylib (or .a)
   go build -buildmode=c-shared -o libapp.dylib lib/main.go

2. Copy GoFlow runtime → GoFlowRuntime.framework

3. Generate Generated.xcconfig with paths
   GOFLOW_APP_PATH=libapp.dylib
   GOFLOW_RUNTIME=GoFlowRuntime.framework

4. Xcode builds native shell

5. Link frameworks
   - GoFlowRuntime.framework (our engine)
   - libapp.dylib (user's Go code)

6. Create .app bundle
   MyApp.app/
   ├── Contents/
   │   ├── MacOS/
   │   │   └── Runner          # Native executable
   │   ├── Frameworks/
   │   │   ├── GoFlowRuntime.framework
   │   │   └── libapp.dylib
   │   └── Resources/
```

## Key Architecture Decisions

### 1. Embedding Model

**Flutter Approach:**
- Dart code → compiled AOT → App.framework
- Flutter engine → Flutter.framework
- Native shell loads both at runtime

**GoFlow Approach (Recommended):**
- Go code → compiled shared library → libapp.dylib
- GoFlow runtime → GoFlowRuntime.framework
- Native shell loads both at runtime

### 2. Build Modes

| Mode | Go Build | Linking | Use Case |
|------|----------|---------|----------|
| Debug | `-buildmode=c-shared` | Dynamic | Hot reload |
| Release | `-buildmode=c-shared` + optimize | Dynamic | Distribution |
| Static | `-buildmode=c-archive` | Static | Single binary |

### 3. Communication Bridge

**CGo Bridge (like Flutter's platform channels):**

```go
// lib/main.go - Go side
package main

import "C"
import (
    "github.com/base-go/GoFlow"
)

//export StartApp
func StartApp() {
    app := goflow.NewApp()
    app.Run()
}

//export HandleEvent
func HandleEvent(eventType C.int, data unsafe.Pointer) {
    // Handle platform events
}

func main() {} // Required for c-shared
```

```swift
// AppDelegate.swift - Swift side
import GoFlowRuntime

let engine = GoFlowEngine()

// Call Go functions
StartApp()  // Exported from Go

// Pass events to Go
engine.handleEvent(type: .mouseClick, data: data)
```

## Advantages of This Approach

### ✅ Clean Separation
- Platform code in platform directories
- App code in lib/
- Clear boundaries

### ✅ Build System Integration
- Uses native build tools (Xcode, Visual Studio)
- IDE support out of the box
- Familiar to platform developers

### ✅ Easy Distribution
- Creates standard .app bundles
- Can submit to App Store
- Platform-native packaging

### ✅ Plugin System
- Plugins can have platform-specific code
- Bridges to native APIs
- Similar to Flutter plugins

### ✅ Hot Reload
- Native shell stays running
- Reload Go shared library
- No full app restart needed

## Implementation Roadmap

### Phase 1: Basic Embedding (Current)
- ✅ Go app as separate binary
- ✅ Native window creation
- ✅ Basic rendering

### Phase 2: Framework Approach
- [ ] Build Go code as shared library
- [ ] Create GoFlowRuntime.framework
- [ ] Template for macos/ directory
- [ ] Auto-generate xcconfig files

### Phase 3: Build Integration
- [ ] `goflow build macos` → creates .app
- [ ] `goflow run macos` → launches Xcode
- [ ] Hot reload with dynamic library

### Phase 4: Advanced Features
- [ ] Plugin system with platform channels
- [ ] Asset bundling
- [ ] Code signing
- [ ] App Store submission

## Example: Complete Build Flow

### goflow new myapp

```bash
myapp/
├── lib/
│   └── main.go              # Go app code
├── macos/
│   ├── Runner/
│   │   └── AppDelegate.swift
│   ├── GoFlow/
│   │   └── Generated.xcconfig
│   └── Runner.xcodeproj/
├── windows/
├── linux/
└── go.mod
```

### goflow build macos

```bash
# 1. Compile Go code
go build -buildmode=c-shared \
  -ldflags="-s -w" \
  -o macos/Runner/libapp.dylib \
  lib/main.go

# 2. Generate xcconfig
cat > macos/GoFlow/Generated.xcconfig <<EOF
GOFLOW_ROOT=/usr/local/goflow
GOFLOW_APP_PATH=libapp.dylib
GOFLOW_RUNTIME=GoFlowRuntime.framework
EOF

# 3. Build with Xcode
xcodebuild -project macos/Runner.xcodeproj \
  -scheme Runner \
  -configuration Release \
  -derivedDataPath build/macos

# 4. Create .app bundle
cp -r build/macos/Build/Products/Release/MyApp.app \
  ./build/MyApp.app
```

### Result

```
build/
└── MyApp.app/
    ├── Contents/
    │   ├── MacOS/
    │   │   └── Runner              # Native executable
    │   ├── Frameworks/
    │   │   ├── GoFlowRuntime.framework
    │   │   └── libapp.dylib        # Your Go code!
    │   ├── Resources/
    │   │   └── Assets.car
    │   └── Info.plist
```

Double-click MyApp.app → Runs your Go GUI app!

## Platform Channels (Like Flutter)

### Go Side

```go
package main

import "C"

//export PlatformChannelInvoke
func PlatformChannelInvoke(channel *C.char, method *C.char, args unsafe.Pointer) {
    channelName := C.GoString(channel)
    methodName := C.GoString(method)

    if channelName == "file_picker" {
        if methodName == "pick" {
            // Handle file picker
            result := pickFile()
            // Send result back to Swift
        }
    }
}
```

### Swift Side

```swift
class GoFlowEngine {
    func invokeMethod(channel: String, method: String, args: Any?) -> Any? {
        // Call Go exported function
        let result = PlatformChannelInvoke(
            channel.cString(using: .utf8),
            method.cString(using: .utf8),
            args
        )
        return result
    }
}

// User code
let result = engine.invokeMethod(
    channel: "file_picker",
    method: "pick",
    args: nil
)
```

## Comparison: Flutter vs GoFlow

| Aspect | Flutter | GoFlow |
|--------|---------|--------|
| App Language | Dart | Go |
| Native Shell | Swift/Kotlin | Swift/Kotlin |
| Compilation | AOT → Framework | Shared Library |
| Runtime | Flutter Engine | GoFlow Runtime |
| Hot Reload | ✅ VM-based | ✅ Library reload |
| Platform API | Platform Channels | CGo + Channels |
| Distribution | .app/.apk | .app/.exe |

## Next Steps for GoFlow

1. **Create GoFlowRuntime.framework**
   - Window management
   - Rendering engine
   - Event handling
   - CGo bridge

2. **Template Generation**
   - `goflow new` creates platform folders
   - Pre-configured Xcode projects
   - Build scripts

3. **Build Integration**
   - Compile Go → shared library
   - Generate xcconfig
   - Invoke Xcode build

4. **Developer Experience**
   - `goflow run macos` → opens in Xcode
   - Hot reload support
   - Debug integration

## Resources

- Flutter Embedding: https://flutter.dev/docs/development/platform-integration/platform-channels
- CGo Shared Libraries: https://golang.org/cmd/cgo/
- Xcode Build Settings: https://developer.apple.com/documentation/xcode

---

**TL;DR:** Flutter's platform folders are just native wrappers that load the framework at runtime. GoFlow can do the same by compiling Go code to a shared library and loading it from a native shell. This gives us platform integration, proper .app bundles, and hot reload capability.
