# GoFlow Hot Reload

Hot reloading implementation for GoFlow applications, enabling rapid development with automatic code reloading and state preservation.

## Features

✅ **File Watcher**: Automatically detects changes to `.go` files
✅ **State Preservation**: Preserves application state across reloads
✅ **Widget Tree Diffing**: Identifies changes in the widget tree
✅ **Partial Rebuilds**: Rebuilds only the changed parts of the widget tree

## Architecture

The hot reload system consists of four main components:

### 1. HotReloader (`reload.go`)

The main orchestrator that:
- Watches files for changes using `fsnotify`
- Triggers reloads when changes are detected
- Manages the reload lifecycle
- Provides debouncing to avoid excessive reloads

### 2. StateStore (`state.go`)

Thread-safe state storage that:
- Preserves application state across reloads
- Supports snapshotting and restoration
- Uses `sync.RWMutex` for concurrent access

### 3. Differ (`differ.go`)

Widget tree diffing engine that:
- Compares old and new widget trees
- Identifies added, removed, modified, and reordered widgets
- Recursively diffs structs, slices, maps, and primitives

### 4. PartialRebuildManager (`builder.go`)

Rebuild strategy manager that:
- Determines rebuild strategy (full, partial, or none)
- Identifies affected widget paths
- Manages dirty flags for widgets

## Usage

### Basic Example

```go
import "github.com/base-go/GoFlow/pkg/hotreload"

// Create hot reloader
reloader, err := hotreload.NewHotReloader(hotreload.Config{
    WatchPaths: []string{"./lib"},
    ReloadFunc: func(ctx context.Context) error {
        app.Rebuild()
        return nil
    },
    EnableStatePreservation: true,
})
if err != nil {
    log.Fatal(err)
}

// Start watching
if err := reloader.Start(); err != nil {
    log.Fatal(err)
}
defer reloader.Stop()
```

### State Preservation

```go
// Get state store
store := reloader.GetStateStore()

// Save state before reload
store.Set("counter", myCounter)
store.Set("username", myUsername)

// Restore state after reload
if val, ok := store.Get("counter"); ok {
    myCounter = val.(int)
}
if val, ok := store.Get("username"); ok {
    myUsername = val.(string)
}
```

### Widget Tree Diffing

```go
// Create differ
differ := hotreload.NewDiffer()

// Compare widget trees
diffs := differ.Diff(oldTree, newTree, []string{"root"})

// Process diffs
for _, diff := range diffs {
    log.Printf("Change: %s at %v", diff.Type, diff.Path)
}
```

### Partial Rebuilds

```go
// Create rebuild manager
manager := hotreload.NewPartialRebuildManager(store)

// Compute diffs
diffs := manager.ComputeDiff(oldWidget, newWidget)

// Determine strategy
strategy := manager.DetermineStrategy(diffs)

switch strategy {
case hotreload.RebuildNone:
    // No rebuild needed
case hotreload.RebuildPartial:
    // Rebuild only affected paths
    paths := manager.GetAffectedPaths(diffs)
    for _, path := range paths {
        rebuildWidget(path)
    }
case hotreload.RebuildFull:
    // Full rebuild
    app.Rebuild()
}
```

## Examples

### Simple Hot Reload Demo

See `examples/hotreload-demo/main.go` for a basic example showing:
- File watching
- State preservation
- Automatic reload on file changes

Run with:
```bash
cd examples/hotreload-demo
go run main.go
```

### Widget Tree Hot Reload Demo

See `examples/hotreload-widget-demo/main.go` for an advanced example showing:
- Widget tree diffing
- Partial rebuild strategies
- Integration with GoFlow framework

Run with:
```bash
cd examples/hotreload-widget-demo
go run main.go
```

## Configuration

### Config Options

```go
type Config struct {
    // WatchPaths are the directories or files to watch
    WatchPaths []string

    // ReloadFunc is called when changes are detected
    ReloadFunc ReloadFunc

    // EnableStatePreservation enables state preservation across reloads
    EnableStatePreservation bool

    // Verbose enables verbose logging
    Verbose bool
}
```

### Ignored Files

The hot reloader automatically ignores:
- Non-`.go` files
- `_test.go` files
- Hidden files (starting with `.`)
- Temporary files

## How It Works

1. **File Watching**: The `fsnotify` library watches for file system events
2. **Event Filtering**: Only relevant events (Write, Create, Remove, Rename) on `.go` files are processed
3. **Debouncing**: Multiple rapid changes are debounced to avoid excessive reloads
4. **State Preservation**: Current state is saved to the StateStore
5. **Reload Trigger**: The reload function is called
6. **Tree Diffing**: Old and new widget trees are compared
7. **Rebuild Strategy**: Determines whether to do a full or partial rebuild
8. **State Restoration**: Preserved state is restored
9. **Render**: Updated UI is rendered

## Diff Types

- **DiffTypeNone**: No change
- **DiffTypeAdded**: Widget or value added
- **DiffTypeRemoved**: Widget or value removed
- **DiffTypeModified**: Widget or value modified
- **DiffTypeReordered**: Children reordered or count changed

## Rebuild Strategies

- **RebuildNone**: No rebuild needed (no changes detected)
- **RebuildPartial**: Rebuild only affected subtrees (optimized)
- **RebuildFull**: Rebuild entire tree (fallback for major changes)

## Best Practices

1. **Use State Preservation**: Save critical state to the StateStore
2. **Minimize State**: Only preserve state that needs to survive reloads
3. **Test Reload Logic**: Ensure your app handles reloads gracefully
4. **Watch Specific Paths**: Only watch directories that contain your app code
5. **Handle Errors**: Implement proper error handling in your reload function

## Performance Considerations

- File watching has minimal overhead using native OS APIs
- State preservation uses efficient `sync.RWMutex` for concurrent access
- Widget tree diffing uses reflection but is optimized for shallow comparisons
- Debouncing prevents excessive reloads during rapid file changes

## Limitations

- Currently only watches `.go` files
- Requires recompilation (not true "hot code swapping")
- Large widget trees may take longer to diff
- State preservation requires serializable types

## Future Enhancements

- [ ] Support for watching asset files (images, fonts, etc.)
- [ ] Incremental compilation
- [ ] Better error recovery
- [ ] Hot reload metrics and profiling
- [ ] Custom file filters
- [ ] Integration with build tools

## Contributing

Contributions are welcome! Please see the main GoFlow contributing guidelines.

## License

Part of the GoFlow framework. See main LICENSE file.
