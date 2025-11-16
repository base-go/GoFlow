# Platform Integration Roadmap

This document outlines the platform-specific integration strategy for GoFlow, following Flutter's approach of native platform projects that host the framework code.

## Current State (v0.1.0)

Currently, platform directories contain simple Go runners:
```
macos/
└── main.go  # Basic Go file that calls lib.Run()
```

This works for initial development but doesn't provide:
- Proper native app bundles
- Platform-specific UI integration (menus, dock icons, etc.)
- Code signing and distribution
- Platform features (notifications, file associations, etc.)

## Target Architecture

Following Flutter's model, each platform directory should contain a full native project:

### macOS Platform Structure

```
macos/
├── Runner.xcodeproj/           # Xcode project
├── Runner/
│   ├── AppDelegate.swift       # macOS app entry point
│   ├── MainViewController.swift # Window controller
│   ├── Info.plist              # App metadata
│   ├── Assets.xcassets/        # App icon, etc.
│   └── Runner-Bridging-Header.h
├── RunnerTests/
├── build.sh                    # Script to build Go code
└── embed_go.sh                 # Script to embed Go binary

```

**How it works:**
1. Xcode project creates native macOS app (.app bundle)
2. Build phase compiles Go code from `../lib/` to shared library or binary
3. AppDelegate initializes GLFW/WGPU rendering context
4. MainViewController loads and executes Go code
5. Go code uses GLFW for window management, WGPU for rendering
6. Result: Proper macOS app with dock icon, menus, etc.

### Windows Platform Structure

```
windows/
├── Runner.sln                  # Visual Studio solution
├── Runner/
│   ├── Runner.vcxproj          # Project file
│   ├── main.cpp                # Win32 entry point
│   ├── Resource.rc             # Resources (icon, manifest)
│   └── app.manifest            # Windows manifest
├── build.bat                   # Script to build Go code
└── embed_go.bat                # Script to embed Go DLL
```

**How it works:**
1. Visual Studio project creates Windows .exe
2. Build script compiles Go code to DLL
3. Win32 app initializes window and rendering context
4. Loads and calls into Go DLL
5. Result: Windows executable with proper icon, manifest, etc.

### Linux Platform Structure

```
linux/
├── CMakeLists.txt              # CMake build configuration
├── main.cc                     # C++ entry point
├── Makefile                    # Alternative: simple Makefile
├── runner.desktop              # Linux desktop entry
├── icons/                      # App icons (various sizes)
└── build.sh                    # Build script
```

**How it works:**
1. CMake/Make builds native launcher
2. Compiles Go code to shared library
3. Launcher initializes X11/Wayland window
4. Loads Go shared library
5. Result: Linux app with desktop integration

## Platform Integration Steps

### Phase 1: GLFW + WGPU Integration (Current Priority)

Before native platform projects, we need the rendering layer:

1. **GLFW Bindings**
   - Use `github.com/go-gl/glfw` for cross-platform windowing
   - Create window, handle events
   - Provide OpenGL/Vulkan context

2. **WGPU Bindings**
   - Use WebGPU bindings for Go (or cgo to wgpu-native)
   - Abstract rendering backend
   - Support Metal (macOS), DirectX (Windows), Vulkan (Linux/Windows)

3. **Update Platform Runners**
   - Initialize GLFW window
   - Create WGPU rendering context
   - Set up event loop
   - Call into lib code for rendering

### Phase 2: Native Project Templates

Once GLFW+WGPU work, add native projects:

#### macOS (Xcode)
1. Create Xcode project template
2. Add Swift/Objective-C code to:
   - Create NSWindow
   - Initialize Metal layer for WGPU
   - Handle app lifecycle (menus, dock, etc.)
   - Load and execute Go code
3. Add build scripts to compile Go code
4. Support code signing and notarization

#### Windows (Visual Studio)
1. Create VS project template
2. Add C++ code to:
   - Create Win32 window
   - Initialize DirectX context for WGPU
   - Handle Windows messages
   - Load Go DLL
3. Add build scripts
4. Support manifest and code signing

#### Linux (CMake)
1. Create CMake template
2. Add C++ code to:
   - Create X11/Wayland window
   - Initialize Vulkan context
   - Handle events
   - Load Go shared library
3. Create .desktop file for integration
4. Support AppImage/Flatpak/Snap packaging

### Phase 3: CLI Integration

Update `goflow create` to generate native projects:

```go
// In createPlatformRunner()
switch platform {
case "macos":
    createXcodeProject(projectPath, modulePath)
case "windows":
    createVisualStudioProject(projectPath, modulePath)
case "linux":
    createCMakeProject(projectPath, modulePath)
}
```

### Phase 4: Platform Features

Once basic integration works, add platform-specific features:

- **macOS**: Menu bar, dock, notifications, app sandbox
- **Windows**: System tray, toast notifications, UAC
- **Linux**: DBus, notifications, desktop themes

## Implementation Example: macOS AppDelegate

```swift
import Cocoa

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {
    var window: NSWindow!
    var goLibrary: UnsafeMutableRawPointer?

    func applicationDidFinishLaunching(_ aNotification: Notification) {
        // 1. Create window
        window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 800, height: 600),
            styleMask: [.titled, .closable, .miniaturizable, .resizable],
            backing: .buffered,
            defer: false
        )
        window.center()
        window.makeKeyAndOrderFront(nil)

        // 2. Initialize Metal for WGPU
        let metalLayer = CAMetalLayer()
        window.contentView?.layer = metalLayer
        window.contentView?.wantsLayer = true

        // 3. Load Go code
        guard let libPath = Bundle.main.path(forResource: "goflow", ofType: "dylib") else {
            fatalError("GoFlow library not found")
        }

        goLibrary = dlopen(libPath, RTLD_NOW)
        guard goLibrary != nil else {
            fatalError("Failed to load GoFlow library: \(String(cString: dlerror()))")
        }

        // 4. Get and call Run function
        typealias RunFunc = @convention(c) () -> Void
        guard let runSymbol = dlsym(goLibrary, "Run"),
              let run = unsafeBitCast(runSymbol, to: RunFunc?.self) else {
            fatalError("Run function not found")
        }

        // 5. Execute Go application
        run()
    }

    func applicationWillTerminate(_ aNotification: Notification) {
        if let lib = goLibrary {
            dlclose(lib)
        }
    }
}
```

## Build Process

### Current (Simple)
```bash
cd macos
go run main.go
```

### Future (Native)
```bash
# macOS
cd macos
xcodebuild -project Runner.xcodeproj -scheme Runner
open build/Release/MyApp.app

# Windows
cd windows
MSBuild Runner.sln /p:Configuration=Release
start bin/Release/MyApp.exe

# Linux
cd linux
cmake . && make
./build/myapp
```

## Distribution

### macOS
- Create .app bundle
- Code sign with Developer ID
- Notarize with Apple
- Distribute as DMG or via Mac App Store

### Windows
- Create installer (NSIS, WiX, or Inno Setup)
- Code sign with certificate
- Distribute as .exe or via Microsoft Store

### Linux
- Create AppImage (self-contained)
- Create Flatpak (Flathub)
- Create Snap (Snapcraft)
- Distribute via package managers

## Benefits of Native Platform Projects

1. **Proper App Bundles**: Real .app, .exe with icons and metadata
2. **Platform Integration**: Menus, dock, system tray, etc.
3. **Distribution Ready**: Can be signed and distributed properly
4. **Performance**: Native event loop and rendering context
5. **Developer Experience**: Use Xcode/VS for platform-specific debugging
6. **Familiar Model**: Follows Flutter's proven approach

## Next Steps

1. ✅ Create CLI tool with basic structure
2. ✅ Design platform integration architecture
3. ⏳ Implement GLFW + WGPU integration in Go
4. ⏳ Create Xcode project template for macOS
5. ⏳ Create Visual Studio template for Windows
6. ⏳ Create CMake template for Linux
7. ⏳ Update CLI to generate native projects
8. ⏳ Add build scripts and automation
9. ⏳ Document platform-specific features
10. ⏳ Support code signing and distribution

## References

- [Flutter Platform Channels](https://docs.flutter.dev/development/platform-integration/platform-channels)
- [Flutter Desktop Embedding](https://github.com/flutter/flutter/tree/master/examples/desktop)
- [GLFW Documentation](https://www.glfw.org/documentation.html)
- [WGPU Native](https://github.com/gfx-rs/wgpu-native)
- [Go Mobile](https://github.com/golang/mobile) - Reference for Go-native binding patterns
