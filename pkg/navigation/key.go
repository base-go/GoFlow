package navigation

import (
	"sync"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// GlobalKey is a key that can be used to access widgets globally
type GlobalKey struct {
	name    string
	element goflow.Element
	mu      sync.RWMutex
}

// NewGlobalKey creates a new global key
func NewGlobalKey(name string) *GlobalKey {
	return &GlobalKey{
		name: name,
	}
}

// GetCurrentContext returns the build context for this key
func (k *GlobalKey) GetCurrentContext() goflow.BuildContext {
	k.mu.RLock()
	defer k.mu.RUnlock()

	if k.element != nil {
		return goflow.NewBuildContext(k.element)
	}
	return nil
}

// GetCurrentWidget returns the widget associated with this key
func (k *GlobalKey) GetCurrentWidget() goflow.Widget {
	k.mu.RLock()
	defer k.mu.RUnlock()

	if k.element != nil {
		return k.element.GetWidget()
	}
	return nil
}

// GetCurrentElement returns the element associated with this key
func (k *GlobalKey) GetCurrentElement() goflow.Element {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.element
}

// AssociateElement associates an element with this key
func (k *GlobalKey) AssociateElement(element goflow.Element) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.element = element
}

// ClearElement clears the associated element
func (k *GlobalKey) ClearElement() {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.element = nil
}

// String returns the string representation of the key
func (k *GlobalKey) String() string {
	return k.name
}
