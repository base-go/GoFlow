# Single Import Demo

This example demonstrates GoFlow's new **single import path architecture**.

## Overview

Instead of importing multiple packages:

```go
import (
    "github.com/base-go/GoFlow/pkg/core/framework"
    "github.com/base-go/GoFlow/pkg/core/widgets"
    "github.com/base-go/GoFlow/pkg/core/signals"
)
```

You now import everything from one place:

```go
import gf "github.com/base-go/GoFlow"
```

## What This Demo Shows

This example demonstrates accessing all GoFlow features through the unified `gf` namespace:

- ✅ **Reactive State (Signals)** - `gf.CreateSignal()`, `gf.CreateComputed()`, `gf.CreateEffect()`
- ✅ **Colors & Styling** - `gf.ColorBlue`, `gf.NewColor()`, `gf.NewTextStyle()`
- ✅ **Geometry Types** - `gf.NewSize()`, `gf.NewOffset()`, `gf.NewRect()`
- ✅ **Platform Detection** - `gf.DetectPlatform()`, platform constants
- ✅ **Widget Types** - All widgets accessible via `gf` namespace
- ✅ **Layout Constants** - Alignment and sizing constants

## Running the Demo

```bash
cd examples/single-import-demo
go run main.go
```

## Expected Output

```
=== GoFlow Single Import Demo ===
This example demonstrates the new single-import architecture.

--- Reactive State (Signals) ---
  Effect triggered: Hello from GoFlow! Count: 5
  Effect triggered: Hello from GoFlow 2.0! Count: 5

--- Batch Updates ---
  Effect triggered: Hello from Single Import Architecture! Count: 10
  Final doubled value: 20

--- Colors & Styling ---
  Primary color: {0 0 255 255}
  Secondary color: {255 100 50 255}
  Text style created: font=Arial, size=16.0

--- Geometry Types ---
  Window size: 800x600
  Position offset: (100, 50)
  Bounding rect: x=100, y=50, w=800, h=600

--- Platform Detection ---
  Detected platform: linux
  Running on: Linux

--- Widget Types Available ---
  All widgets accessible via 'gf' namespace:
  - Layout: gf.Column, gf.Row, gf.Stack, gf.Container
  - Display: gf.Text, gf.Icon, gf.Image, gf.Card
  - Input: gf.TextField, gf.Checkbox, gf.Slider, gf.Switch
  - Animation: gf.AnimatedContainer, gf.FadeTransition
  - Material: gf.MaterialButton, gf.MaterialScaffold
  - Cupertino: gf.CupertinoButton, gf.CupertinoNavigationBar
  - Navigation: gf.Get, gf.ShowDialog, gf.ShowSnackbar

--- Layout Constants ---
  MainAxis alignments: Start=0, Center=2, End=1
  CrossAxis alignments: Start=0, Center=2, Stretch=3

=== Summary ===
✓ Single import path: import gf "github.com/base-go/GoFlow"
✓ All framework features accessible through 'gf' namespace
✓ Clean, intuitive API surface
✓ Easy discoverability via IDE autocomplete

Compare this to the old multi-import approach:
  ❌ OLD: import "github.com/base-go/GoFlow/pkg/core/signals"
  ❌ OLD: import "github.com/base-go/GoFlow/pkg/core/widgets"
  ❌ OLD: import "github.com/base-go/GoFlow/pkg/core/framework"

  ✅ NEW: import gf "github.com/base-go/GoFlow"
```

## Key Takeaways

1. **One Import to Rule Them All** - Everything you need is in `github.com/base-go/GoFlow`

2. **Clean Namespace** - Use `gf.` prefix for all GoFlow types and functions

3. **Better IDE Support** - Autocomplete shows all available types after typing `gf.`

4. **Familiar Pattern** - Similar to Flutter (`package:flutter/material.dart`) and React

5. **Easier to Learn** - New developers don't need to hunt for which package contains what

## Next Steps

- See [MIGRATION_GUIDE.md](../../MIGRATION_GUIDE.md) for migrating existing code
- Check out other examples to see the pattern in action
- Start building with `import gf "github.com/base-go/GoFlow"`!
