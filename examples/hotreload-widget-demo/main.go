package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
	"github.com/base-go/GoFlow/pkg/hotreload"
)

var (
	app      *goflow.App
	reloader *hotreload.HotReloader
	manager  *hotreload.PartialRebuildManager
)

func main() {
	fmt.Println("GoFlow Hot Reload Widget Demo")
	fmt.Println("==============================")
	fmt.Println()

	// Create the initial app
	rootWidget := buildUI()
	app = goflow.NewApp(rootWidget, 800, 600)

	// Build initial widget tree
	app.Build()
	app.Layout()
	app.Render()

	// Create hot reloader
	var err error
	reloader, err = hotreload.NewHotReloader(hotreload.Config{
		WatchPaths: []string{"./examples/hotreload-widget-demo"},
		ReloadFunc: handleHotReload,
		EnableStatePreservation: true,
	})
	if err != nil {
		log.Fatalf("Failed to create hot reloader: %v", err)
	}

	// Create partial rebuild manager
	manager = hotreload.NewPartialRebuildManager(reloader.GetStateStore())

	// Start hot reloader
	if err := reloader.Start(); err != nil {
		log.Fatalf("Failed to start hot reloader: %v", err)
	}
	defer reloader.Stop()

	fmt.Println("\nHot reloader started!")
	fmt.Println("Try modifying the buildUI() function below to see hot reload in action.")
	fmt.Println("Press Ctrl+C to exit.\n")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down...")
}

// buildUI builds the user interface
// Try modifying this function and save the file to trigger hot reload!
func buildUI() goflow.Widget {
	// Create a container with padding
	container := widgets.NewContainer()
	container.Child = widgets.NewText("Nested Widget")
	container.Padding = &goflow.EdgeInsets{
		Top:    10,
		Bottom: 10,
		Left:   20,
		Right:  20,
	}

	// Create column with center alignment
	column := widgets.NewColumn(
		[]goflow.Widget{
			widgets.NewText("Hot Reload Demo"),
			widgets.NewText("Modify this text and save!"),
			container,
		},
	)
	column.MainAxisAlign = widgets.MainAxisCenter
	column.CrossAxisAlign = widgets.CrossAxisCenter

	return widgets.NewCenter(column)
}

func handleHotReload(ctx context.Context) error {
	fmt.Println("\n🔥 Hot reload triggered!")

	// Build new widget tree
	fmt.Println("  ✓ Building new widget tree...")
	newRootWidget := buildUI()

	// Compare old and new trees
	fmt.Println("  ✓ Computing widget tree diff...")
	oldWidget := app.GetRootWidget()
	diffs := manager.ComputeDiff(oldWidget, newRootWidget)

	if len(diffs) == 0 {
		fmt.Println("  ℹ No changes detected")
		return nil
	}

	fmt.Printf("  ℹ Found %d changes:\n", len(diffs))
	for i, diff := range diffs {
		if i < 5 { // Show first 5 diffs
			fmt.Printf("    - %s at %v\n", diff.Type, diff.Path)
		}
	}
	if len(diffs) > 5 {
		fmt.Printf("    ... and %d more\n", len(diffs)-5)
	}

	// Determine rebuild strategy
	strategy := manager.DetermineStrategy(diffs)
	fmt.Printf("  ✓ Rebuild strategy: %v\n", getStrategyName(strategy))

	switch strategy {
	case hotreload.RebuildNone:
		// No rebuild needed
		fmt.Println("  ✓ No rebuild required")

	case hotreload.RebuildPartial:
		// Partial rebuild
		paths := manager.GetAffectedPaths(diffs)
		fmt.Printf("  ✓ Rebuilding %d affected paths\n", len(paths))

		// In a real implementation, we would rebuild only these paths
		// For now, we'll do a full rebuild for simplicity
		rebuildApp(newRootWidget)

	case hotreload.RebuildFull:
		// Full rebuild
		fmt.Println("  ✓ Performing full rebuild")
		rebuildApp(newRootWidget)
	}

	fmt.Println("  ✓ Hot reload completed!\n")
	return nil
}

func rebuildApp(newRootWidget goflow.Widget) {
	// Update the app with new widget
	app.UpdateRootWidget(newRootWidget)

	// Rebuild, layout, and render
	app.Build()
	app.Layout()
	app.Render()
}

func getStrategyName(strategy hotreload.RebuildStrategy) string {
	switch strategy {
	case hotreload.RebuildNone:
		return "None"
	case hotreload.RebuildPartial:
		return "Partial"
	case hotreload.RebuildFull:
		return "Full"
	default:
		return "Unknown"
	}
}
