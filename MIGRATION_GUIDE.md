# GoFlow Single Import Architecture - Migration Guide

## Overview

GoFlow has transitioned to a **single import path architecture** to significantly improve developer experience. This guide will help you migrate your existing code to use the new unified import.

## Why This Change?

### Before (Multiple Imports)
```go
import (
    "github.com/base-go/GoFlow/pkg/core/framework"
    "github.com/base-go/GoFlow/pkg/core/widgets"
    "github.com/base-go/GoFlow/pkg/core/signals"
    "github.com/base-go/GoFlow/pkg/ui/material"
)

func main() {
    count := signals.New(0)

    framework.RunApp(widgets.NewContainer(
        widgets.NewText("Hello"),
    ))
}
```

### After (Single Import)
```go
import gf "github.com/base-go/GoFlow"

func main() {
    count := gf.CreateSignal(0)

    gf.RunApp(gf.Container{
        Child: gf.Text{Content: "Hello"},
    })
}
```

## Benefits

✅ **Simpler imports** - One line instead of 5-10
✅ **Better discoverability** - IDE autocomplete shows everything
✅ **Cleaner code** - Single namespace for all GoFlow types
✅ **Easier refactoring** - Internal reorganization won't break your code
✅ **Familiar pattern** - Follows Flutter, React, and other modern UI frameworks

## Migration Steps

### Step 1: Update Your Imports

**Old:**
```go
import (
    "github.com/base-go/GoFlow/pkg/core/framework"
    "github.com/base-go/GoFlow/pkg/core/widgets"
    "github.com/base-go/GoFlow/pkg/core/signals"
    "github.com/base-go/GoFlow/pkg/ui/material"
    "github.com/base-go/GoFlow/pkg/navigation"
)
```

**New:**
```go
import gf "github.com/base-go/GoFlow"
```

### Step 2: Update Type References

Replace package-qualified types with the `gf` alias:

| Old | New |
|-----|-----|
| `framework.Widget` | `gf.Widget` |
| `framework.RunApp(w)` | `gf.RunApp(w)` |
| `widgets.NewText(s)` | `gf.Text{Content: s}` |
| `signals.New(v)` | `gf.CreateSignal(v)` |
| `signals.NewEffect(fn)` | `gf.CreateEffect(fn)` |
| `material.NewButton(...)` | `gf.MaterialButton{...}` |
| `navigation.Get.To(w)` | `gf.Get.To(w)` |

### Step 3: Update Signal Functions

The signal creation functions have been renamed for clarity:

| Old | New |
|-----|-----|
| `signals.New(v)` | `gf.CreateSignal(v)` |
| `signals.NewComputed(fn)` | `gf.CreateComputed(fn)` |
| `signals.NewEffect(fn)` | `gf.CreateEffect(fn)` |
| `signals.NewSlice(v)` | `gf.CreateSignalSlice(v)` |
| `signals.NewMap(v)` | `gf.CreateSignalMap(v)` |
| `signals.Batch(fn)` | `gf.Batch(fn)` |
| `signals.Untracked(fn)` | `gf.Untracked(fn)` |

### Step 4: Update Material/Cupertino Widgets

Material and Cupertino widgets now have prefixes to avoid naming conflicts:

| Old | New |
|-----|-----|
| `material.AppBar` | `gf.MaterialAppBar` |
| `material.Scaffold` | `gf.MaterialScaffold` |
| `material.Button` | `gf.MaterialButton` |
| `cupertino.Button` | `gf.CupertinoButton` |
| `cupertino.NavigationBar` | `gf.CupertinoNavigationBar` |

### Step 5: Update Constants

All constants are now accessible through the `gf` namespace:

| Old | New |
|-----|-----|
| `framework.ColorBlue` | `gf.ColorBlue` |
| `framework.PlatformLinux` | `gf.PlatformLinux` |
| `widgets.MainAxisCenter` | `gf.MainAxisCenter` |
| `widgets.CrossAxisStretch` | `gf.CrossAxisStretch` |

## Complete Example Migration

### Before

```go
package main

import (
    "fmt"
    "github.com/base-go/GoFlow/pkg/core/framework"
    "github.com/base-go/GoFlow/pkg/core/widgets"
    "github.com/base-go/GoFlow/pkg/core/signals"
    "github.com/base-go/GoFlow/pkg/ui/material"
)

func main() {
    count := signals.New(0)

    doubled := signals.NewComputed(func() int {
        return count.Get() * 2
    })

    dispose := signals.NewEffect(func() {
        fmt.Printf("Count: %d\n", count.Get())
    })
    defer dispose()

    app := material.NewScaffold(
        material.NewAppBar("Counter App"),
        widgets.NewColumn(
            widgets.MainAxisCenter,
            widgets.CrossAxisCenter,
            []framework.Widget{
                widgets.NewText(fmt.Sprintf("Count: %d", count.Peek())),
                widgets.NewText(fmt.Sprintf("Doubled: %d", doubled.Get())),
                material.NewButton("Increment", func() {
                    count.Update(func(v int) int { return v + 1 })
                }),
            },
        ),
    )

    framework.RunApp(app)
}
```

### After

```go
package main

import (
    "fmt"
    gf "github.com/base-go/GoFlow"
)

func main() {
    count := gf.CreateSignal(0)

    doubled := gf.CreateComputed(func() int {
        return count.Get() * 2
    })

    dispose := gf.CreateEffect(func() {
        fmt.Printf("Count: %d\n", count.Get())
    })
    defer dispose()

    app := gf.MaterialScaffold{
        AppBar: gf.MaterialAppBar{Title: "Counter App"},
        Body: gf.Column{
            MainAxisAlignment: gf.MainAxisCenter,
            CrossAxisAlignment: gf.CrossAxisCenter,
            Children: []gf.Widget{
                gf.Text{Content: fmt.Sprintf("Count: %d", count.Peek())},
                gf.Text{Content: fmt.Sprintf("Doubled: %d", doubled.Get())},
                gf.MaterialButton{
                    Text: "Increment",
                    OnPressed: func() {
                        count.Update(func(v int) int { return v + 1 })
                    },
                },
            },
        },
    }

    gf.RunApp(app)
}
```

## Backwards Compatibility

The old package structure (`pkg/core/framework`, `pkg/core/widgets`, etc.) will continue to work for a transition period, but is considered **deprecated**. We recommend migrating to the single import as soon as possible.

### Gradual Migration

You can migrate gradually by using both styles in the same codebase:

```go
import (
    gf "github.com/base-go/GoFlow"  // New style
    "github.com/base-go/GoFlow/pkg/core/signals"  // Old style (deprecated)
)

func main() {
    // Mix old and new (works but not recommended)
    count := signals.New(0)  // Old
    app := gf.Container{}    // New
}
```

However, we strongly recommend switching entirely to the new import style for consistency.

## IDE Support

### VS Code / Gopls

The single import works seamlessly with Go's language server. You'll get full autocomplete after typing `gf.`:

```go
import gf "github.com/base-go/GoFlow"

// Type "gf." and get autocomplete for:
// - gf.Widget
// - gf.Container
// - gf.CreateSignal
// - gf.RunApp
// - etc.
```

### Organizing Imports

Use `goimports` or your IDE's organize imports feature to automatically clean up:

```bash
goimports -w .
```

## Common Patterns

### Creating Widgets

**Old:**
```go
container := widgets.NewContainer(
    widgets.NewPadding(
        framework.NewEdgeInsets(10, 10, 10, 10),
        widgets.NewText("Hello"),
    ),
)
```

**New:**
```go
container := gf.Container{
    Padding: gf.NewEdgeInsets(10, 10, 10, 10),
    Child: gf.Text{Content: "Hello"},
}
```

### Using Signals

**Old:**
```go
count := signals.New(0)
doubled := signals.NewComputed(func() int {
    return count.Get() * 2
})
signals.NewEffect(func() {
    fmt.Println(count.Get())
})
```

**New:**
```go
count := gf.CreateSignal(0)
doubled := gf.CreateComputed(func() int {
    return count.Get() * 2
})
gf.CreateEffect(func() {
    fmt.Println(count.Get())
})
```

### Navigation

**Old:**
```go
import "github.com/base-go/GoFlow/pkg/navigation"

navigation.Get.To(myWidget)
navigation.ShowSnackbar("Hello")
```

**New:**
```go
import gf "github.com/base-go/GoFlow"

gf.Get.To(myWidget)
gf.ShowSnackbar("Hello")
```

## Advanced Usage

### Optional Subpackages

Some specialized packages remain separate for advanced use cases:

```go
import (
    gf "github.com/base-go/GoFlow"
    "github.com/base-go/GoFlow/pkg/testing"     // Widget testing
    "github.com/base-go/GoFlow/pkg/hotreload"   // Hot reload
)
```

These are imported separately because they're typically only needed in specific contexts (tests, development tools, etc.).

## FAQs

### Q: Why not keep the granular imports?

**A:** While granular imports can help with compile times in some languages, Go's fast compiler makes this less critical. The improved developer experience of a single import far outweighs any minor compilation benefits.

### Q: What if I prefer the old style?

**A:** The old imports will continue to work during a deprecation period, but we strongly encourage adopting the new style for long-term maintainability and consistency with the community.

### Q: How do I know which import to use for specialized features?

**A:** If it's related to building UIs (widgets, state, navigation, styling), use the single import. Only specialized tools (testing, hot reload, platform-specific code) need separate imports.

### Q: Will this break my existing code?

**A:** No! The old imports will continue to work. However, we recommend migrating to get the best experience.

## Need Help?

- Check the updated examples in `examples/single-import-demo/`
- See the [full API documentation](https://pkg.go.dev/github.com/base-go/GoFlow)
- Open an issue on [GitHub](https://github.com/base-go/GoFlow/issues)

## Summary

The single import architecture makes GoFlow easier to learn, use, and maintain. The migration is straightforward and can be done incrementally. Start by trying it in a new feature or file, then gradually migrate the rest of your codebase.

**Simple rule:** If you're using it to build UI, import it once:

```go
import gf "github.com/base-go/GoFlow"
```

Happy coding! 🎉
