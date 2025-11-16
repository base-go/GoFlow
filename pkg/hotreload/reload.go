package hotreload

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// HotReloader manages hot reloading of the application
type HotReloader struct {
	watcher     *fsnotify.Watcher
	watchPaths  []string
	reloadFunc  ReloadFunc
	mu          sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
	stateStore  *StateStore
	debounce    chan struct{}
	running     bool
}

// ReloadFunc is called when a reload is triggered
type ReloadFunc func(ctx context.Context) error

// Config holds configuration for the hot reloader
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

// NewHotReloader creates a new hot reloader
func NewHotReloader(cfg Config) (*HotReloader, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	hr := &HotReloader{
		watcher:    watcher,
		watchPaths: cfg.WatchPaths,
		reloadFunc: cfg.ReloadFunc,
		ctx:        ctx,
		cancel:     cancel,
		stateStore: NewStateStore(),
		debounce:   make(chan struct{}, 1),
	}

	// Add watch paths
	for _, path := range cfg.WatchPaths {
		if err := hr.addWatchPath(path); err != nil {
			watcher.Close()
			cancel()
			return nil, fmt.Errorf("failed to add watch path %s: %w", path, err)
		}
	}

	return hr, nil
}

// addWatchPath adds a path to watch, recursively if it's a directory
func (h *HotReloader) addWatchPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return h.watcher.Add(path)
	}

	// For directories, walk and add all subdirectories
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// Skip hidden directories
			if len(info.Name()) > 0 && info.Name()[0] == '.' {
				return filepath.SkipDir
			}

			return h.watcher.Add(p)
		}

		return nil
	})
}

// Start starts the hot reloader
func (h *HotReloader) Start() error {
	h.mu.Lock()
	if h.running {
		h.mu.Unlock()
		return fmt.Errorf("hot reloader already running")
	}
	h.running = true
	h.mu.Unlock()

	go h.watch()
	log.Println("Hot reloader started")

	return nil
}

// Stop stops the hot reloader
func (h *HotReloader) Stop() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.running {
		return nil
	}

	h.cancel()
	h.running = false

	if err := h.watcher.Close(); err != nil {
		return fmt.Errorf("failed to close watcher: %w", err)
	}

	log.Println("Hot reloader stopped")
	return nil
}

// watch watches for file changes
func (h *HotReloader) watch() {
	for {
		select {
		case <-h.ctx.Done():
			return

		case event, ok := <-h.watcher.Events:
			if !ok {
				return
			}

			// Filter out events we don't care about
			if h.shouldIgnoreEvent(event) {
				continue
			}

			log.Printf("File change detected: %s (%s)", event.Name, event.Op)

			// Debounce: trigger reload after a short delay
			h.triggerReload()

		case err, ok := <-h.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

// shouldIgnoreEvent checks if an event should be ignored
func (h *HotReloader) shouldIgnoreEvent(event fsnotify.Event) bool {
	// Ignore non-Go files for now
	ext := filepath.Ext(event.Name)
	if ext != ".go" {
		return true
	}

	// Ignore temporary files
	base := filepath.Base(event.Name)
	if len(base) > 0 && base[0] == '.' {
		return true
	}

	// Ignore _test.go files
	if len(base) > 8 && base[len(base)-8:] == "_test.go" {
		return true
	}

	// Only watch Write, Create, Remove, Rename events
	if event.Op&fsnotify.Write == 0 &&
		event.Op&fsnotify.Create == 0 &&
		event.Op&fsnotify.Remove == 0 &&
		event.Op&fsnotify.Rename == 0 {
		return true
	}

	return false
}

// triggerReload triggers a reload with debouncing
func (h *HotReloader) triggerReload() {
	select {
	case h.debounce <- struct{}{}:
		go h.performReload()
	default:
		// Reload already scheduled
	}
}

// performReload performs the actual reload
func (h *HotReloader) performReload() {
	// Wait for debounce
	<-h.debounce

	log.Println("Performing hot reload...")

	if h.reloadFunc != nil {
		if err := h.reloadFunc(h.ctx); err != nil {
			log.Printf("Reload failed: %v", err)
			return
		}
	}

	log.Println("Hot reload completed successfully")
}

// GetStateStore returns the state store for preserving state across reloads
func (h *HotReloader) GetStateStore() *StateStore {
	return h.stateStore
}
