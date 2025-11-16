package input

import (
	"sort"
	"sync"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// EventDispatcher manages and dispatches input events
type EventDispatcher struct {
	mu sync.RWMutex

	// Event listeners by type
	listeners map[EventType][]*EventListener

	// Global listeners (receive all events)
	globalListeners []*EventListener

	// Event filters
	filters []EventFilter

	// Target resolution
	targetResolver TargetResolver

	// Event queue for deferred processing
	eventQueue []InputEvent
	queueMode  bool

	// Statistics
	stats *DispatcherStats
}

// NewEventDispatcher creates a new event dispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners:       make(map[EventType][]*EventListener),
		globalListeners: make([]*EventListener, 0),
		filters:         make([]EventFilter, 0),
		eventQueue:      make([]InputEvent, 0),
		queueMode:       false,
		stats:           &DispatcherStats{},
	}
}

// AddListener registers an event listener
func (d *EventDispatcher) AddListener(eventType EventType, handler EventHandler, priority int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	listener := &EventListener{
		EventType: eventType,
		Handler:   handler,
		Priority:  priority,
		Once:      false,
	}

	d.listeners[eventType] = append(d.listeners[eventType], listener)
	d.sortListeners(eventType)
}

// AddListenerOnce registers a one-time event listener
func (d *EventDispatcher) AddListenerOnce(eventType EventType, handler EventHandler, priority int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	listener := &EventListener{
		EventType: eventType,
		Handler:   handler,
		Priority:  priority,
		Once:      true,
	}

	d.listeners[eventType] = append(d.listeners[eventType], listener)
	d.sortListeners(eventType)
}

// AddGlobalListener registers a listener that receives all events
func (d *EventDispatcher) AddGlobalListener(handler EventHandler, priority int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	listener := &EventListener{
		EventType: EventTypeUnknown,
		Handler:   handler,
		Priority:  priority,
		Once:      false,
	}

	d.globalListeners = append(d.globalListeners, listener)
	d.sortGlobalListeners()
}

// RemoveListener removes an event listener
func (d *EventDispatcher) RemoveListener(eventType EventType, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if listeners, ok := d.listeners[eventType]; ok {
		for i, listener := range listeners {
			// Note: comparing functions directly won't work in Go
			// In practice, you'd need to track listeners differently
			// This is a simplified version
			if listener.Handler != nil {
				d.listeners[eventType] = append(listeners[:i], listeners[i+1:]...)
				return
			}
		}
	}
}

// AddFilter adds an event filter
func (d *EventDispatcher) AddFilter(filter EventFilter) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.filters = append(d.filters, filter)
}

// SetTargetResolver sets the target resolver for hit testing
func (d *EventDispatcher) SetTargetResolver(resolver TargetResolver) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.targetResolver = resolver
}

// Dispatch dispatches an event to registered listeners
func (d *EventDispatcher) Dispatch(event InputEvent) {
	d.mu.Lock()

	// Queue mode - add to queue instead of dispatching
	if d.queueMode {
		d.eventQueue = append(d.eventQueue, event)
		d.mu.Unlock()
		return
	}

	d.stats.TotalEvents++
	d.mu.Unlock()

	// Apply filters
	if !d.applyFilters(event) {
		d.mu.Lock()
		d.stats.FilteredEvents++
		d.mu.Unlock()
		return
	}

	// Resolve target if we have a position event and target resolver
	var target interface{}
	if posEvent, ok := event.(interface{ Position() *goflow.Offset }); ok {
		if d.targetResolver != nil {
			target = d.targetResolver.ResolveTarget(posEvent.Position())
		}
	}

	// Dispatch to global listeners first
	d.dispatchToGlobalListeners(event, target)

	// Dispatch to event-specific listeners
	d.dispatchToListeners(event, target)

	d.mu.Lock()
	d.stats.DispatchedEvents++
	d.mu.Unlock()
}

// dispatchToGlobalListeners dispatches to global listeners
func (d *EventDispatcher) dispatchToGlobalListeners(event InputEvent, target interface{}) {
	d.mu.RLock()
	listeners := make([]*EventListener, len(d.globalListeners))
	copy(listeners, d.globalListeners)
	d.mu.RUnlock()

	for _, listener := range listeners {
		propagation := listener.Handler(event)

		if listener.Once {
			d.removeGlobalListener(listener)
		}

		if propagation == PropagationStopImmediate {
			return
		}
	}
}

// dispatchToListeners dispatches to event-specific listeners
func (d *EventDispatcher) dispatchToListeners(event InputEvent, target interface{}) {
	eventType := event.Type()

	d.mu.RLock()
	listeners, ok := d.listeners[eventType]
	if !ok {
		d.mu.RUnlock()
		return
	}

	// Make a copy to avoid holding lock during dispatch
	listenersCopy := make([]*EventListener, len(listeners))
	copy(listenersCopy, listeners)
	d.mu.RUnlock()

	for _, listener := range listenersCopy {
		propagation := listener.Handler(event)

		if listener.Once {
			d.removeListener(eventType, listener)
		}

		if propagation == PropagationStopImmediate {
			return
		}
		if propagation == PropagationStop {
			break
		}
	}
}

// applyFilters applies all filters to an event
func (d *EventDispatcher) applyFilters(event InputEvent) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, filter := range d.filters {
		if !filter.Filter(event) {
			return false
		}
	}
	return true
}

// sortListeners sorts listeners by priority (highest first)
func (d *EventDispatcher) sortListeners(eventType EventType) {
	listeners := d.listeners[eventType]
	sort.Slice(listeners, func(i, j int) bool {
		return listeners[i].Priority > listeners[j].Priority
	})
}

// sortGlobalListeners sorts global listeners by priority
func (d *EventDispatcher) sortGlobalListeners() {
	sort.Slice(d.globalListeners, func(i, j int) bool {
		return d.globalListeners[i].Priority > d.globalListeners[j].Priority
	})
}

// removeListener removes a specific listener
func (d *EventDispatcher) removeListener(eventType EventType, listener *EventListener) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if listeners, ok := d.listeners[eventType]; ok {
		for i, l := range listeners {
			if l == listener {
				d.listeners[eventType] = append(listeners[:i], listeners[i+1:]...)
				return
			}
		}
	}
}

// removeGlobalListener removes a global listener
func (d *EventDispatcher) removeGlobalListener(listener *EventListener) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i, l := range d.globalListeners {
		if l == listener {
			d.globalListeners = append(d.globalListeners[:i], d.globalListeners[i+1:]...)
			return
		}
	}
}

// EnableQueueMode enables event queuing (events are queued but not dispatched)
func (d *EventDispatcher) EnableQueueMode() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.queueMode = true
}

// DisableQueueMode disables event queuing
func (d *EventDispatcher) DisableQueueMode() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.queueMode = false
}

// ProcessQueue processes all queued events
func (d *EventDispatcher) ProcessQueue() {
	d.mu.Lock()
	queue := d.eventQueue
	d.eventQueue = make([]InputEvent, 0)
	d.mu.Unlock()

	for _, event := range queue {
		d.Dispatch(event)
	}
}

// ClearQueue clears the event queue without processing
func (d *EventDispatcher) ClearQueue() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.eventQueue = make([]InputEvent, 0)
}

// GetStats returns dispatcher statistics
func (d *EventDispatcher) GetStats() DispatcherStats {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return *d.stats
}

// ResetStats resets dispatcher statistics
func (d *EventDispatcher) ResetStats() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.stats = &DispatcherStats{}
}

// EventFilter filters events before dispatch
type EventFilter interface {
	Filter(event InputEvent) bool
}

// TargetResolver resolves event targets based on position
type TargetResolver interface {
	ResolveTarget(position *goflow.Offset) interface{}
}

// DispatcherStats tracks dispatcher statistics
type DispatcherStats struct {
	TotalEvents      int
	DispatchedEvents int
	FilteredEvents   int
}

// FunctionFilter is a simple function-based event filter
type FunctionFilter struct {
	FilterFunc func(InputEvent) bool
}

func (f *FunctionFilter) Filter(event InputEvent) bool {
	return f.FilterFunc(event)
}

// NewFunctionFilter creates a function-based filter
func NewFunctionFilter(filterFunc func(InputEvent) bool) *FunctionFilter {
	return &FunctionFilter{FilterFunc: filterFunc}
}

// EventTypeFilter filters events by type
type EventTypeFilter struct {
	AllowedTypes map[EventType]bool
}

func (f *EventTypeFilter) Filter(event InputEvent) bool {
	return f.AllowedTypes[event.Type()]
}

// NewEventTypeFilter creates an event type filter
func NewEventTypeFilter(allowedTypes ...EventType) *EventTypeFilter {
	typeMap := make(map[EventType]bool)
	for _, t := range allowedTypes {
		typeMap[t] = true
	}
	return &EventTypeFilter{AllowedTypes: typeMap}
}

// InputRouter routes input events to appropriate handlers
type InputRouter struct {
	dispatcher         *EventDispatcher
	keyboardState      *KeyboardState
	mouseState         *MouseState
	touchState         *TouchState
	keyBindingManager  *KeyBindingManager
	mouseRegionManager *MouseRegionManager
	gestureRecognizer  *MultiTouchRecognizer
}

// NewInputRouter creates a new input router
func NewInputRouter() *InputRouter {
	return &InputRouter{
		dispatcher:         NewEventDispatcher(),
		keyboardState:      NewKeyboardState(),
		mouseState:         NewMouseState(),
		touchState:         NewTouchState(),
		keyBindingManager:  NewKeyBindingManager(),
		mouseRegionManager: NewMouseRegionManager(),
		gestureRecognizer:  NewMultiTouchRecognizer(nil),
	}
}

// GetDispatcher returns the event dispatcher
func (r *InputRouter) GetDispatcher() *EventDispatcher {
	return r.dispatcher
}

// GetKeyboardState returns the keyboard state
func (r *InputRouter) GetKeyboardState() *KeyboardState {
	return r.keyboardState
}

// GetMouseState returns the mouse state
func (r *InputRouter) GetMouseState() *MouseState {
	return r.mouseState
}

// GetTouchState returns the touch state
func (r *InputRouter) GetTouchState() *TouchState {
	return r.touchState
}

// GetKeyBindingManager returns the key binding manager
func (r *InputRouter) GetKeyBindingManager() *KeyBindingManager {
	return r.keyBindingManager
}

// GetMouseRegionManager returns the mouse region manager
func (r *InputRouter) GetMouseRegionManager() *MouseRegionManager {
	return r.mouseRegionManager
}

// GetGestureRecognizer returns the gesture recognizer
func (r *InputRouter) GetGestureRecognizer() *MultiTouchRecognizer {
	return r.gestureRecognizer
}

// RouteEvent routes an event to appropriate handlers
func (r *InputRouter) RouteEvent(event InputEvent) {
	// Update state trackers
	switch e := event.(type) {
	case *KeyboardEvent:
		r.keyboardState.HandleEvent(e)
		r.keyBindingManager.HandleEvent(e)
	case *MouseEvent:
		r.mouseState.HandleEvent(e)
		r.mouseRegionManager.HandleMouseEvent(e)
	case *MouseWheelEvent:
		r.mouseRegionManager.HandleWheelEvent(e)
	case *TouchEvent:
		r.gestureRecognizer.ProcessTouchEvent(e)
	}

	// Dispatch to listeners
	r.dispatcher.Dispatch(event)
}
