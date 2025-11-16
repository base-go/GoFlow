// Package hotreload provides hot reloading capabilities for GoFlow applications.
//
// Hot reloading allows developers to see changes in their application without
// restarting the entire process, preserving application state and improving
// development velocity.
//
// # Features
//
// - File watching: Automatically detects changes to .go files
// - State preservation: Preserves application state across reloads
// - Widget tree diffing: Identifies changes in the widget tree
// - Partial rebuilds: Rebuilds only the changed parts of the widget tree
//
// # Usage
//
// To use hot reloading in your GoFlow application:
//
//	import "github.com/base-go/GoFlow/pkg/hotreload"
//
//	// Create a hot reloader
//	reloader, err := hotreload.NewHotReloader(hotreload.Config{
//		WatchPaths: []string{"./lib"},
//		ReloadFunc: func(ctx context.Context) error {
//			// Rebuild your app
//			app.Rebuild()
//			return nil
//		},
//		EnableStatePreservation: true,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Start watching
//	if err := reloader.Start(); err != nil {
//		log.Fatal(err)
//	}
//	defer reloader.Stop()
//
// # State Preservation
//
// Use the StateStore to preserve state across reloads:
//
//	stateStore := reloader.GetStateStore()
//
//	// Save state before reload
//	stateStore.Set("counter", myCounter)
//
//	// Restore state after reload
//	if val, ok := stateStore.Get("counter"); ok {
//		myCounter = val.(int)
//	}
//
// # Widget Tree Diffing
//
// The differ automatically computes differences between widget trees:
//
//	differ := hotreload.NewDiffer()
//	diffs := differ.Diff(oldTree, newTree, []string{"root"})
//
//	for _, diff := range diffs {
//		log.Printf("Change detected: %s at %v", diff.Type, diff.Path)
//	}
//
// # Partial Rebuilds
//
// The PartialRebuildManager determines which parts of the tree need rebuilding:
//
//	manager := hotreload.NewPartialRebuildManager(stateStore)
//	strategy := manager.DetermineStrategy(diffs)
//
//	switch strategy {
//	case hotreload.RebuildNone:
//		// No rebuild needed
//	case hotreload.RebuildPartial:
//		// Rebuild only affected paths
//		paths := manager.GetAffectedPaths(diffs)
//		for _, path := range paths {
//			rebuildWidget(path)
//		}
//	case hotreload.RebuildFull:
//		// Full rebuild
//		app.Rebuild()
//	}
package hotreload
