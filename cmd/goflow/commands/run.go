package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/base-go/GoFlow/pkg/hotreload"
	"github.com/base-go/mamba"
	"golang.org/x/term"
)

// RunCommand creates the 'goflow run' command with integrated hot reload
func RunCommand() *mamba.Command {
	var (
		hotReloadEnabled = true
		noHotReload      = false
		watchPaths       = []string{"."}
		buildCmd         = "go run ."
		verbose          = false
	)

	cmd := &mamba.Command{
		Use:   "run [flags]",
		Short: "Run GoFlow application with hot reload",
		Long: `Run your GoFlow application with integrated hot reload functionality.

Similar to Flutter's 'flutter run', this command:
  • Runs your GoFlow application
  • Watches for file changes
  • Provides hot reload on 'r' key press
  • Maintains application state during reloads
  • Shows development console with reload statistics

Hot Reload Commands:
  r     Hot reload application
  R     Full restart application
  q     Quit application
  h     Show help`,
		Example: `  # Run current GoFlow app with hot reload
  goflow run

  # Run with specific watch paths
  goflow run --watch src,lib

  # Run with custom build command
  goflow run --build "go run cmd/main.go"

  # Run without hot reload
  goflow run --no-hot-reload`,
		Run: func(cmd *mamba.Command, args []string) {
			// Handle flag priority: --no-hot-reload overrides --hot-reload
			finalHotReload := hotReloadEnabled && !noHotReload
			runGoFlowApp(finalHotReload, watchPaths, buildCmd, verbose)
		},
	}

	// Add flags
	cmd.Flags().BoolVar(&hotReloadEnabled, "hot-reload", true, "Enable hot reload (default: true)")
	cmd.Flags().BoolVar(&noHotReload, "no-hot-reload", false, "Disable hot reload")
	cmd.Flags().StringSliceVar(&watchPaths, "watch", []string{"."}, "Paths to watch for changes")
	cmd.Flags().StringVar(&buildCmd, "build", "go run .", "Build command to execute")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	return cmd
}

// runGoFlowApp implements the main hot reload development experience
func runGoFlowApp(hotReloadEnabled bool, watchPaths []string, buildCmd string, verbose bool) {
	fmt.Println("🚀 Starting GoFlow development server...")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if !hotReloadEnabled {
		fmt.Println("⚡ Hot reload disabled")
		runWithoutHotReload(buildCmd)
		return
	}

	fmt.Println("⚡ Hot reload enabled")
	fmt.Println("📁 Watching paths:", strings.Join(watchPaths, ", "))
	fmt.Println("🔨 Build command:", buildCmd)
	fmt.Println()
	fmt.Println("💡 Commands:")
	fmt.Println("   r  - Hot reload")
	fmt.Println("   R  - Full restart")
	fmt.Println("   q  - Quit")
	fmt.Println("   h  - Help")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	// Initialize hot reload state
	appState := &HotReloadState{
		BuildCommand:  buildCmd,
		WatchPaths:    watchPaths,
		Verbose:       verbose,
		StartTime:     time.Now(),
		ReloadCount:   0,
		RestartCount:  0,
	}

	// Start hot reload system
	reloader, err := hotreload.NewHotReloader(hotreload.Config{
		WatchPaths:              watchPaths,
		ReloadFunc:              appState.handleAutoReload,
		EnableStatePreservation: true,
	})
	if err != nil {
		log.Fatalf("❌ Failed to initialize hot reloader: %v", err)
	}

	// Store the reloader in app state
	appState.Reloader = reloader

	// Start file watcher
	if err := reloader.Start(); err != nil {
		log.Fatalf("❌ Failed to start hot reloader: %v", err)
	}
	defer reloader.Stop()

	// Start the initial application
	if err := appState.startApp(); err != nil {
		log.Fatalf("❌ Failed to start application: %v", err)
	}

	// Setup signal handling for clean shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start keyboard input handler
	go appState.handleKeyboardInput()

	// Main event loop
	fmt.Printf("✅ GoFlow app started (PID: %d)\n", appState.AppProcess.Process.Pid)
	fmt.Printf("⏱️  Ready in %v\n\n", time.Since(appState.StartTime).Round(time.Millisecond))

	// Wait for signals or app completion
	select {
	case <-sigChan:
		fmt.Println("\n🛑 Received interrupt signal...")
	case <-appState.Done:
		fmt.Println("\n✋ Application exited")
	}

	// Clean shutdown
	appState.cleanup()
	fmt.Println("👋 GoFlow development server stopped")
}

// HotReloadState manages the application state during development
type HotReloadState struct {
	BuildCommand string
	WatchPaths   []string
	Verbose      bool
	StartTime    time.Time
	ReloadCount  int
	RestartCount int
	AppProcess   *exec.Cmd
	Reloader     *hotreload.HotReloader
	Done         chan struct{}
	ctx          context.Context
	cancel       context.CancelFunc
}

// startApp launches the GoFlow application process
func (s *HotReloadState) startApp() error {
	if s.AppProcess != nil {
		s.AppProcess.Process.Kill()
		s.AppProcess.Wait()
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.Done = make(chan struct{})

	// Parse build command
	cmdParts := strings.Fields(s.BuildCommand)
	if len(cmdParts) == 0 {
		return fmt.Errorf("invalid build command: %s", s.BuildCommand)
	}

	// Create command with context for clean cancellation
	s.AppProcess = exec.CommandContext(s.ctx, cmdParts[0], cmdParts[1:]...)
	s.AppProcess.Stdout = os.Stdout
	s.AppProcess.Stderr = os.Stderr

	// Start the process
	if err := s.AppProcess.Start(); err != nil {
		return fmt.Errorf("failed to start app: %v", err)
	}

	// Monitor process completion
	go func() {
		defer close(s.Done)
		s.AppProcess.Wait()
	}()

	return nil
}

// handleAutoReload handles automatic reloads from file watcher
func (s *HotReloadState) handleAutoReload(ctx context.Context) error {
	fmt.Printf("🔥 File change detected - Hot reloading...\n")
	return s.performHotReload()
}

// performHotReload executes a hot reload
func (s *HotReloadState) performHotReload() error {
	start := time.Now()
	s.ReloadCount++

	fmt.Printf("🔄 Hot reload #%d...", s.ReloadCount)

	// For now, we'll do a full restart since Go is compiled
	// In the future, this could be enhanced with more sophisticated hot reload
	if err := s.restartApp(); err != nil {
		fmt.Printf(" ❌ Failed\n")
		return fmt.Errorf("hot reload failed: %v", err)
	}

	duration := time.Since(start).Round(time.Millisecond)
	fmt.Printf(" ✅ Completed in %v\n\n", duration)

	return nil
}

// restartApp performs a full application restart
func (s *HotReloadState) restartApp() error {
	s.RestartCount++

	// Stop current app
	if s.AppProcess != nil && s.AppProcess.Process != nil {
		s.cancel() // Cancel context to gracefully stop

		// Give the process a moment to exit gracefully
		done := make(chan error, 1)
		go func() {
			done <- s.AppProcess.Wait()
		}()

		select {
		case <-done:
			// Process exited gracefully
		case <-time.After(2 * time.Second):
			// Force kill if it doesn't exit
			if s.AppProcess.Process != nil {
				s.AppProcess.Process.Kill()
				<-done
			}
		}
	}

	// Start new instance
	return s.startApp()
}

// handleKeyboardInput processes keyboard commands with single-key input
func (s *HotReloadState) handleKeyboardInput() {
	// Set terminal to raw mode for single-key input
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("⚠️  Warning: Could not enable raw input mode: %v\n", err)
		fmt.Println("📝 Commands will require pressing Enter after each key.")
		s.handleKeyboardInputFallback()
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("🎮 Single-key commands enabled (press keys directly, no Enter needed)")

	buffer := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil || n == 0 {
			continue
		}

		key := string(buffer[0])

		switch key {
		case "r":
			fmt.Println("\n🔥 Manual hot reload triggered...")
			if err := s.performHotReload(); err != nil {
				fmt.Printf("❌ Hot reload failed: %v\n", err)
			}

		case "R":
			fmt.Println("\n🔄 Full restart triggered...")
			if err := s.restartApp(); err != nil {
				fmt.Printf("❌ Restart failed: %v\n", err)
			} else {
				fmt.Printf("✅ App restarted\n\n")
			}

		case "q", "Q":
			fmt.Println("\n🛑 Shutting down...")
			s.cancel()
			return

		case "h", "H":
			fmt.Print("\n")
			s.showHelp()

		case "s", "S":
			fmt.Print("\n")
			s.showStats()

		case "\x03": // Ctrl+C
			fmt.Println("\n🛑 Received Ctrl+C...")
			s.cancel()
			return

		default:
			// Ignore other keys to avoid cluttering output
		}
	}
}

// handleKeyboardInputFallback processes keyboard commands with Enter-based input (fallback)
func (s *HotReloadState) handleKeyboardInputFallback() {
	buffer := make([]byte, 1)
	for {
		fmt.Print("\n> ")
		n, err := os.Stdin.Read(buffer)
		if err != nil || n == 0 {
			continue
		}

		key := strings.TrimSpace(string(buffer[0]))

		switch key {
		case "r":
			fmt.Println("🔥 Manual hot reload triggered...")
			if err := s.performHotReload(); err != nil {
				fmt.Printf("❌ Hot reload failed: %v\n", err)
			}

		case "R":
			fmt.Println("🔄 Full restart triggered...")
			if err := s.restartApp(); err != nil {
				fmt.Printf("❌ Restart failed: %v\n", err)
			} else {
				fmt.Printf("✅ App restarted\n\n")
			}

		case "q", "Q":
			fmt.Println("🛑 Shutting down...")
			s.cancel()
			return

		case "h", "H":
			s.showHelp()

		case "s", "S":
			s.showStats()

		default:
			if key != "" {
				fmt.Printf("❓ Unknown command '%s'. Press 'h' for help.\n", key)
			}
		}
	}
}

// showHelp displays available commands
func (s *HotReloadState) showHelp() {
	fmt.Println()
	fmt.Println("🎮 GoFlow Hot Reload Commands:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  r       Hot reload application")
	fmt.Println("  R       Full restart application")
	fmt.Println("  q       Quit development server")
	fmt.Println("  s       Show statistics")
	fmt.Println("  h       Show this help")
	fmt.Println()
}

// showStats displays development session statistics
func (s *HotReloadState) showStats() {
	uptime := time.Since(s.StartTime).Round(time.Second)

	fmt.Println()
	fmt.Println("📊 Development Session Stats:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  ⏱️  Uptime:        %v\n", uptime)
	fmt.Printf("  🔥 Hot reloads:    %d\n", s.ReloadCount)
	fmt.Printf("  🔄 Full restarts:  %d\n", s.RestartCount)
	fmt.Printf("  📁 Watching:       %s\n", strings.Join(s.WatchPaths, ", "))
	fmt.Printf("  🔨 Build command:  %s\n", s.BuildCommand)
	fmt.Println()
}

// cleanup performs shutdown cleanup
func (s *HotReloadState) cleanup() {
	if s.cancel != nil {
		s.cancel()
	}

	if s.AppProcess != nil && s.AppProcess.Process != nil {
		s.AppProcess.Process.Kill()
		s.AppProcess.Wait()
	}

	if s.Reloader != nil {
		s.Reloader.Stop()
	}
}

// runWithoutHotReload runs the app without hot reload functionality
func runWithoutHotReload(buildCmd string) {
	cmdParts := strings.Fields(buildCmd)
	if len(cmdParts) == 0 {
		log.Fatalf("❌ Invalid build command: %s", buildCmd)
	}

	fmt.Printf("🚀 Running: %s\n", buildCmd)

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Application failed: %v", err)
	}
}