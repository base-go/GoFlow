// +build darwin

package macos

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework QuartzCore
#include "window_bridge.h"
#include <stdlib.h>

// Forward declarations for callbacks
extern void goDrawCallback(WindowHandle window, void* userData);
extern void goResizeCallback(WindowHandle window, int width, int height, void* userData);
*/
import "C"
import (
	"sync"
	"unsafe"
)

// DrawFunc is called when the window needs to be redrawn
type DrawFunc func(canvas *CoreGraphicsCanvas)

// ResizeFunc is called when the window is resized
type ResizeFunc func(width, height int)

// Window represents a macOS window
type Window struct {
	handle     C.WindowHandle
	width      int
	height     int
	title      string
	drawFunc   DrawFunc
	resizeFunc ResizeFunc
	mu         sync.Mutex
}

var (
	windowRegistry   = make(map[C.WindowHandle]*Window)
	windowRegistryMu sync.RWMutex
)

// NewWindow creates a new macOS window
func NewWindow(width, height int, title string) *Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	handle := C.createWindow(C.int(width), C.int(height), cTitle)
	if handle == nil {
		panic("Failed to create window")
	}

	window := &Window{
		handle: handle,
		width:  width,
		height: height,
		title:  title,
	}

	// Register window
	windowRegistryMu.Lock()
	windowRegistry[handle] = window
	windowRegistryMu.Unlock()

	// Set up callbacks
	C.setDrawCallback(handle, C.DrawCallback(C.goDrawCallback), unsafe.Pointer(handle))
	C.setResizeCallback(handle, C.ResizeCallback(C.goResizeCallback), unsafe.Pointer(handle))

	return window
}

// Destroy destroys the window
func (w *Window) Destroy() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		windowRegistryMu.Lock()
		delete(windowRegistry, w.handle)
		windowRegistryMu.Unlock()

		C.destroyWindow(w.handle)
		w.handle = nil
	}
}

// Show shows the window
func (w *Window) Show() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.showWindow(w.handle)
	}
}

// Hide hides the window
func (w *Window) Hide() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.hideWindow(w.handle)
	}
}

// ShouldClose returns true if the window should close
func (w *Window) ShouldClose() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle == nil {
		return true
	}

	return C.windowShouldClose(w.handle) != 0
}

// SetShouldClose sets whether the window should close
func (w *Window) SetShouldClose(shouldClose bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		if shouldClose {
			C.setWindowShouldClose(w.handle, 1)
		} else {
			C.setWindowShouldClose(w.handle, 0)
		}
	}
}

// PollEvents processes pending events
func (w *Window) PollEvents() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.pollEvents(w.handle)
	}
}

// WaitEvents waits for events
func (w *Window) WaitEvents() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.waitEvents(w.handle)
	}
}

// GetCanvas returns a canvas for drawing
func (w *Window) GetCanvas() *CoreGraphicsCanvas {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle == nil {
		return nil
	}

	ctx := C.getWindowGraphicsContext(w.handle)
	if ctx == nil {
		return nil
	}

	// Return a canvas that wraps the window's graphics context
	// Note: We don't own this context, so we shouldn't destroy it
	return &CoreGraphicsCanvas{
		ctx:    ctx,
		width:  w.width,
		height: w.height,
	}
}

// Present updates the window display
func (w *Window) Present() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.presentWindow(w.handle)
	}
}

// SetNeedsDisplay marks the window as needing redraw
func (w *Window) SetNeedsDisplay() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.setWindowNeedsDisplay(w.handle)
	}
}

// GetSize returns the window size
func (w *Window) GetSize() (int, int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle == nil {
		return w.width, w.height
	}

	var width, height C.int
	C.getWindowSize(w.handle, &width, &height)
	w.width = int(width)
	w.height = int(height)
	return w.width, w.height
}

// SetSize sets the window size
func (w *Window) SetSize(width, height int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.setWindowSize(w.handle, C.int(width), C.int(height))
		w.width = width
		w.height = height
	}
}

// GetPosition returns the window position
func (w *Window) GetPosition() (int, int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle == nil {
		return 0, 0
	}

	var x, y C.int
	C.getWindowPosition(w.handle, &x, &y)
	return int(x), int(y)
}

// SetPosition sets the window position
func (w *Window) SetPosition(x, y int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.handle != nil {
		C.setWindowPosition(w.handle, C.int(x), C.int(y))
	}
}

// SetTitle sets the window title
func (w *Window) SetTitle(title string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.title = title
	if w.handle != nil {
		cTitle := C.CString(title)
		defer C.free(unsafe.Pointer(cTitle))
		C.setWindowTitle(w.handle, cTitle)
	}
}

// SetDrawFunc sets the draw callback
func (w *Window) SetDrawFunc(fn DrawFunc) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.drawFunc = fn
}

// SetResizeFunc sets the resize callback
func (w *Window) SetResizeFunc(fn ResizeFunc) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.resizeFunc = fn
}

// InitApp initializes the Cocoa application
// This must be called before creating any windows
func InitApp() {
	C.initApp()
}

// Run starts the Cocoa event loop (blocking)
// This should be called after setting up windows
func Run() {
	C.runApp()
}

//export goDrawCallback
func goDrawCallback(handle C.WindowHandle, userData unsafe.Pointer) {
	windowRegistryMu.RLock()
	window, ok := windowRegistry[handle]
	windowRegistryMu.RUnlock()

	if ok && window.drawFunc != nil {
		canvas := window.GetCanvas()
		if canvas != nil {
			window.drawFunc(canvas)
		}
	}
}

//export goResizeCallback
func goResizeCallback(handle C.WindowHandle, width, height C.int, userData unsafe.Pointer) {
	windowRegistryMu.RLock()
	window, ok := windowRegistry[handle]
	windowRegistryMu.RUnlock()

	if ok {
		window.width = int(width)
		window.height = int(height)

		if window.resizeFunc != nil {
			window.resizeFunc(int(width), int(height))
		}
	}
}
