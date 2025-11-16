# GoFlow Documentation

Welcome to the complete GoFlow documentation! This guide will help you navigate all available resources.

## 📚 Documentation Index

### Getting Started
- **[Getting Started Guide](./GETTING_STARTED.md)** - Complete beginner tutorial from installation to first app
- **[CLI Documentation](./CLI.md)** - Complete CLI reference for `goflow` command
- **[Project Structure](./PROJECT_STRUCTURE.md)** - Understanding GoFlow project layout

### Core Concepts
- **[Architecture](../ARCHITECTURE.md)** - Framework architecture (widgets, elements, render objects)
- **[Rendering Architecture](./RENDERING.md)** - How rendering works (native backends vs WGPU)
- **[Flutter Inspiration](./FLUTTER_INSPIRATION.md)** - Comparison with Flutter and migration guide

### Advanced Topics
- **[Platform Integration](./PLATFORM_INTEGRATION.md)** - Native platform projects roadmap (Xcode, CMake, VS)

### Examples
- **[Playground Example](../examples/playground/)** - Comprehensive framework feature test
- **[More Examples](../examples/)** - Additional example applications

## 🚀 Quick Navigation

### I want to...

#### ...get started with GoFlow
→ Start with [Getting Started Guide](./GETTING_STARTED.md)

#### ...understand how GoFlow works internally
→ Read [Architecture](../ARCHITECTURE.md)

#### ...create a new project
→ See [CLI Documentation](./CLI.md#goflow-create)

#### ...understand the project structure
→ Read [Project Structure](./PROJECT_STRUCTURE.md)

#### ...know how rendering works
→ Read [Rendering Architecture](./RENDERING.md)

#### ...migrate from Flutter
→ See [Flutter Inspiration](./FLUTTER_INSPIRATION.md)

#### ...contribute to GoFlow
→ Check out [Platform Integration](./PLATFORM_INTEGRATION.md) for roadmap

## 📖 Documentation by Topic

### Installation & Setup
1. [Installing the CLI](./CLI.md#installation)
2. [Creating your first project](./GETTING_STARTED.md#creating-your-first-app)
3. [Setting up dependencies](./GETTING_STARTED.md#step-3-set-up-dependencies)
4. [Running your app](./GETTING_STARTED.md#step-4-run-your-app)

### Core Framework
1. [Widgets](../ARCHITECTURE.md#1-widgets) - UI building blocks
2. [Elements](../ARCHITECTURE.md#2-elements) - Lifecycle management
3. [RenderObjects](../ARCHITECTURE.md#3-renderobjects) - Layout & painting
4. [The Three Trees](../ARCHITECTURE.md#the-three-trees)
5. [Layout System](../ARCHITECTURE.md#layout-system)

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
5. [Troubleshooting](./CLI.md#troubleshooting)

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
1. Read [Architecture](../ARCHITECTURE.md) thoroughly
2. Read [Rendering Architecture](./RENDERING.md)
3. Review [Platform Integration](./PLATFORM_INTEGRATION.md) roadmap
4. Check out the codebase structure
5. Look for contribution opportunities

## 📝 Document Status

| Document | Status | Version |
|----------|--------|---------|
| Getting Started | ✅ Complete | v0.1.0 |
| CLI Documentation | ✅ Complete | v0.1.0 |
| Project Structure | ✅ Complete | v0.1.0 |
| Architecture | ✅ Complete | v0.1.0 |
| Rendering | ✅ Complete | v0.1.0 |
| Flutter Inspiration | ✅ Complete | v0.1.0 |
| Platform Integration | ✅ Roadmap | v0.1.0 |

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
- Hot Reload Guide (when implemented)
- Event Handling Guide (when implemented)
- Animation System Guide (when implemented)
- Testing Guide
- Performance Optimization Guide
- Publishing & Distribution Guide

---

**Happy learning! 🚀**
