package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/base-go/GoFlow/pkg/hotreload"
)

// AppState represents the application state
type AppState struct {
	Counter    int
	Message    string
	StartTime  time.Time
	ReloadCount int
}

var (
	state *AppState
	store *hotreload.StateStore
)

func main() {
	fmt.Println("GoFlow Hot Reload Demo")
	fmt.Println("======================")
	fmt.Println()

	// Initialize state
	state = &AppState{
		Counter:   0,
		Message:   "Hello, GoFlow!",
		StartTime: time.Now(),
	}

	// Create hot reloader
	reloader, err := hotreload.NewHotReloader(hotreload.Config{
		WatchPaths: []string{"./examples/hotreload-demo"},
		ReloadFunc: handleReload,
		EnableStatePreservation: true,
	})
	if err != nil {
		log.Fatalf("Failed to create hot reloader: %v", err)
	}

	// Get state store
	store = reloader.GetStateStore()

	// Try to restore previous state
	restoreState()

	// Start hot reloader
	if err := reloader.Start(); err != nil {
		log.Fatalf("Failed to start hot reloader: %v", err)
	}
	defer reloader.Stop()

	fmt.Println("Hot reloader started. Watching for changes...")
	fmt.Println("Try modifying this file to see hot reload in action!")
	fmt.Println()

	// Run the app simulation
	runApp()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nShutting down...")

	// Save state before exit
	saveState()
}

func runApp() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			state.Counter++
			printState()
		}
	}()
}

func printState() {
	elapsed := time.Since(state.StartTime)
	fmt.Printf("\n[App State]\n")
	fmt.Printf("  Uptime: %v\n", elapsed.Round(time.Second))
	fmt.Printf("  Counter: %d\n", state.Counter)
	fmt.Printf("  Message: %s\n", state.Message)
	fmt.Printf("  Reload Count: %d\n", state.ReloadCount)
	fmt.Printf("  [Try changing the Message variable in the code!]\n")
}

func handleReload(ctx context.Context) error {
	fmt.Println("\n🔥 Hot reload triggered!")

	// Save current state
	saveState()

	// Simulate rebuild
	fmt.Println("  ✓ Rebuilding widget tree...")
	time.Sleep(100 * time.Millisecond)

	// Restore state
	restoreState()

	// Increment reload counter
	state.ReloadCount++

	fmt.Println("  ✓ State restored")
	fmt.Println("  ✓ Hot reload completed!\n")

	// Print current state
	printState()

	return nil
}

func saveState() {
	store.Set("counter", state.Counter)
	store.Set("message", state.Message)
	store.Set("startTime", state.StartTime)
	store.Set("reloadCount", state.ReloadCount)
}

func restoreState() {
	if val, ok := store.Get("counter"); ok {
		state.Counter = val.(int)
	}
	if val, ok := store.Get("message"); ok {
		state.Message = val.(string)
	}
	if val, ok := store.Get("startTime"); ok {
		state.StartTime = val.(time.Time)
	}
	if val, ok := store.Get("reloadCount"); ok {
		state.ReloadCount = val.(int)
	}
}
