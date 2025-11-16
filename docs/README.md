# GoFlow Documentation

Welcome to the complete GoFlow documentation! This guide will help you navigate all available resources.

## 📚 Documentation Index

### Getting Started
- **[Getting Started Guide](./GETTING_STARTED.md)** - Complete beginner tutorial from installation to first app
- **[CLI Documentation](./CLI.md)** - Complete CLI reference for `goflow` command
- **[Project Structure](./PROJECT_STRUCTURE.md)** - Understanding GoFlow project layout

### Core Concepts
- **[Architecture](./ARCHITECTURE.md)** - Framework architecture (widgets, elements, render objects)
- **[Rendering Architecture](./RENDERING.md)** - How rendering works (native backends vs WGPU)
- **[Flutter Inspiration](./FLUTTER_INSPIRATION.md)** - Comparison with Flutter and migration guide
- **[Design Systems](./DESIGN_SYSTEMS.md)** - Material, Cupertino, and Adaptive widgets

### Widget Documentation
- **[Widget Reference](./WIDGETS_REFERENCE.md)** - Complete catalog of all 27+ widgets
- **[Layout Completion Summary](./LAYOUT_COMPLETION_SUMMARY.md)** - Layout system implementation details
- **[Widget Examples](./widgets/)** - Detailed widget documentation by category

### Developer Experience
- **[Hot Reload](../pkg/hotreload/README.md)** - Hot reload system with state preservation
- **[Testing Framework](../pkg/testing/README.md)** - Widget testing, golden tests, and integration testing
- **[Navigation](../pkg/navigation/README.md)** - GetX-style navigation and routing
- **[Input System](./INPUT_SYSTEM.md)** - Mouse, keyboard, and gesture handling

### Advanced Topics
- **[Platform Integration](./PLATFORM_INTEGRATION.md)** - Native platform projects roadmap (Xcode, CMake, VS)
- **[Development Status](./DEV_STATUS.md)** - Current development branch status
- **[Progress Tracking](./PROGRESS.md)** - Detailed progress and roadmap
- **[Workflow Guide](./WORKFLOW.md)** - Complete technical workflow

### Examples
- **[Playground Example](../examples/playground/)** - Comprehensive framework feature test
- **[More Examples](../examples/)** - Additional example applications

## 🚀 Quick Navigation

### I want to...

#### ...get started with GoFlow
→ Start with [Getting Started Guide](./GETTING_STARTED.md)

#### ...understand how GoFlow works internally
→ Read [Architecture](./ARCHITECTURE.md)

#### ...create a new project
→ See [CLI Documentation](./CLI.md#goflow-create)

#### ...understand the project structure
→ Read [Project Structure](./PROJECT_STRUCTURE.md)

#### ...know how rendering works
→ Read [Rendering Architecture](./RENDERING.md)

#### ...see all available widgets
→ Browse [Widget Reference](./WIDGETS_REFERENCE.md)

#### ...implement hot reload in my app
→ Check [Hot Reload Guide](../pkg/hotreload/README.md)

#### ...write tests for my widgets
→ See [Testing Framework](../pkg/testing/README.md)

#### ...add navigation to my app
→ Read [Navigation Guide](../pkg/navigation/README.md)

#### ...migrate from Flutter
→ See [Flutter Inspiration](./FLUTTER_INSPIRATION.md)

#### ...contribute to GoFlow
→ Check [Progress Tracking](./PROGRESS.md) for roadmap

## 📖 Documentation by Topic

### Installation & Setup
1. [Installing the CLI](./CLI.md#installation)
2. [Creating your first project](./GETTING_STARTED.md#creating-your-first-app)
3. [Setting up dependencies](./GETTING_STARTED.md#step-3-set-up-dependencies)
4. [Running your app](./GETTING_STARTED.md#step-4-run-your-app)

### Core Framework
1. [Widgets](./ARCHITECTURE.md#1-widgets) - UI building blocks
2. [Elements](./ARCHITECTURE.md#2-elements) - Lifecycle management
3. [RenderObjects](./ARCHITECTURE.md#3-renderobjects) - Layout & painting
4. [The Three Trees](./ARCHITECTURE.md#the-three-trees)
5. [Layout System](./ARCHITECTURE.md#layout-system)
6. [Widget Reference](./WIDGETS_REFERENCE.md) - All available widgets

### State Management
1. [Signals](../README.md#quick-start) - Reactive values
2. [Computed Signals](./GETTING_STARTED.md#advanced-working-with-computed-signals)
3. [Signal Collections](./GETTING_STARTED.md#advanced-signal-collections)
4. [Effects](../README.md#quick-start) - Side effects
5. [Batch Updates](../README.md#batch-updates)

### Widgets & UI
1. [Built-in Widgets](../ARCHITECTURE.md#core-concepts)
2. [Text Widget](./GETTING_STARTED.md#1-widgets---ui-building-blocks)
3. [Container Widget](./GETTING_STARTED.md#1-widgets---ui-building-blocks)
4. [Layout Widgets](./GETTING_STARTED.md#1-widgets---ui-building-blocks)
5. [Building Custom Widgets](../examples/playground/main.go)

### Rendering & Platform
1. [Rendering Overview](./RENDERING.md#goflow-rendering-architecture)
2. [Canvas Interface](./RENDERING.md#canvas-interface)
3. [Native Backends](./RENDERING.md#option-1-native-renderers-recommended-for-desktop)
4. [WGPU Backend](./RENDERING.md#option-2-wgpu-webgpu---cross-platform)
5. [Platform Integration Roadmap](./PLATFORM_INTEGRATION.md)

### CLI & Tools
1. [CLI Commands](./CLI.md#commands)
2. [goflow create](./CLI.md#goflow-create)
3. [Project Templates](./CLI.md#available-templates)
4. [Platform Selection](./CLI.md#available-platforms)
5. [Hot Reload](../pkg/hotreload/README.md)
6. [Testing Framework](../pkg/testing/README.md)
7. [Troubleshooting](./CLI.md#troubleshooting)

## 🎯 Learning Paths

### Path 1: Complete Beginner
1. Read [Getting Started](./GETTING_STARTED.md)
2. Run the [Playground Example](../examples/playground/)
3. Create your own app using `goflow create`
4. Read [Architecture](../ARCHITECTURE.md) when ready for details

### Path 2: Flutter Developer
1. Read [Flutter Inspiration](./FLUTTER_INSPIRATION.md) for comparison
2. Skim [Getting Started](./GETTING_STARTED.md) focusing on differences
3. Review [Signals](../README.md) (replaces setState)
4. Start building!

### Path 3: Framework Contributor
1. Read [Architecture](./ARCHITECTURE.md) thoroughly
2. Read [Rendering Architecture](./RENDERING.md)
3. Review [Progress Tracking](./PROGRESS.md) for roadmap
4. Check [Development Status](./DEV_STATUS.md) for current state
5. Review [Platform Integration](./PLATFORM_INTEGRATION.md) roadmap
6. Check out the codebase structure
7. Look for contribution opportunities

## 📝 Document Status

| Document | Status | Version |
|----------|--------|---------|
| Getting Started | ✅ Complete | v0.2.0 |
| CLI Documentation | ✅ Complete | v0.2.0 |
| Project Structure | ✅ Complete | v0.2.0 |
| Architecture | ✅ Complete | v0.2.0 |
| Rendering | ✅ Complete | v0.2.0 |
| Flutter Inspiration | ✅ Complete | v0.2.0 |
| Design Systems | ✅ Complete | v0.2.0 |
| Widget Reference | ✅ Complete | v0.2.0 |
| Layout Summary | ✅ Complete | v0.2.0 |
| Hot Reload | ✅ Complete | v0.2.0 |
| Testing Framework | ✅ Complete | v0.2.0 |
| Navigation | ✅ Complete | v0.2.0 |
| Input System | ✅ Complete | v0.2.0 |
| Progress Tracking | ✅ Complete | v0.2.0 |
| Development Status | ✅ Complete | v0.2.0 |
| Workflow Guide | ✅ Complete | v0.2.0 |
| Platform Integration | ✅ Roadmap | v0.2.0 |

## 🔗 External Resources

### Go Resources
- [Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

### Framework Inspiration
- [Flutter Architecture](https://flutter.dev/docs/resources/architectural-overview)
- [Preact Signals](https://preactjs.com/guide/v10/signals/)
- [Solid.js Reactivity](https://www.solidjs.com/tutorial/introduction_signals)

### Rendering & Graphics
- [Core Graphics (macOS)](https://developer.apple.com/documentation/coregraphics)
- [Direct2D (Windows)](https://learn.microsoft.com/en-us/windows/win32/direct2d/direct2d-portal)
- [Cairo (Linux)](https://www.cairographics.org/)
- [WebGPU](https://www.w3.org/TR/webgpu/)

## 💡 Tips for Reading

- **Beginner?** Start with Getting Started and follow along
- **Skimming?** Each doc has a table of contents
- **Looking for something specific?** Use the "I want to..." section above
- **Want details?** Architecture and Rendering docs are comprehensive
- **Need examples?** Check the examples directory

## 🆘 Getting Help

Can't find what you're looking for?

1. Check the [FAQ](./GETTING_STARTED.md#faqs)
2. Search through all docs (Cmd/Ctrl+F across files)
3. Look at [example code](../examples/)
4. Ask in [GitHub Discussions](https://github.com/base-go/GoFlow/discussions)
5. [Report missing documentation](https://github.com/base-go/GoFlow/issues)

## 🎨 Contributing to Docs

Found an error? Have a suggestion? Want to improve the docs?

- Docs are written in Markdown
- Located in `/docs` directory
- Follow existing structure and style
- Add examples where helpful
- Keep explanations clear and concise

## 📅 What's Next?

Upcoming documentation:
- Animation System Guide (when implemented)
- Performance Optimization Guide
- Publishing & Distribution Guide
- Advanced Gesture Handling
- Custom Render Objects Guide
- Plugin Development Guide

---

**Happy learning! 🚀**
