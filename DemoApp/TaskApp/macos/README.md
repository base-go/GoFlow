# macOS Runner

This directory contains the macOS native shell for your GoFlow application.

## Structure

- **Runner/** - Swift code for the native macOS app shell
  - AppDelegate.swift - Application lifecycle management
  - MainWindow.swift - Main window and GoFlow engine integration
  - Info.plist - macOS application metadata

- **GoFlow/** - GoFlow-specific configuration
  - ephemeral/GoFlow-Generated.xcconfig - Auto-generated build configuration

## How It Works

This follows the same architecture as Flutter's platform embedding:

1. Your Go code in `lib/` is compiled to a shared library (`libapp.dylib`)
2. The Swift runner (this folder) creates a native macOS window
3. The runner loads and initializes your Go code
4. GoFlow renders your UI using Core Graphics

## Building

The `goflow build macos` command handles the full build process:

1. Compiles Go code → libapp.dylib
2. Generates xcconfig with paths
3. Builds Swift runner with Xcode
4. Links everything together
5. Creates .app bundle

## Development

For development with hot reload:

```bash
goflow run
```

This will:
- Watch for file changes
- Recompile Go code
- Reload the shared library
- Preserve app state

## Learn More

See `docs/MACOS_EMBEDDING.md` in the GoFlow repository for detailed documentation.
