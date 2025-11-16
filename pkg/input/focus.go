package input

import (
	"sync"
)

// FocusNode represents a node in the focus tree
// Similar to Flutter's FocusNode
type FocusNode struct {
	mu sync.RWMutex

	// Hierarchy
	parent   *FocusNode
	children []*FocusNode

	// State
	hasFocus       bool
	hasPrimaryFocus bool
	canRequestFocus bool
	skipTraversal  bool
	descendantsAreFocusable bool

	// Callbacks
	onFocusChange func(bool)
	onKey         func(*KeyboardEvent) bool

	// Debugging
	debugLabel string
}

// NewFocusNode creates a new focus node
func NewFocusNode(options ...FocusNodeOption) *FocusNode {
	node := &FocusNode{
		canRequestFocus:         true,
		descendantsAreFocusable: true,
		children:                make([]*FocusNode, 0),
	}

	for _, opt := range options {
		opt(node)
	}

	return node
}

// FocusNodeOption configures a focus node
type FocusNodeOption func(*FocusNode)

// WithDebugLabel sets a debug label for the focus node
func WithDebugLabel(label string) FocusNodeOption {
	return func(n *FocusNode) {
		n.debugLabel = label
	}
}

// WithOnFocusChange sets the focus change callback
func WithOnFocusChange(callback func(bool)) FocusNodeOption {
	return func(n *FocusNode) {
		n.onFocusChange = callback
	}
}

// WithOnKey sets the key event callback
func WithOnKey(callback func(*KeyboardEvent) bool) FocusNodeOption {
	return func(n *FocusNode) {
		n.onKey = callback
	}
}

// WithCanRequestFocus sets whether the node can request focus
func WithCanRequestFocus(canRequest bool) FocusNodeOption {
	return func(n *FocusNode) {
		n.canRequestFocus = canRequest
	}
}

// WithSkipTraversal sets whether to skip this node during traversal
func WithSkipTraversal(skip bool) FocusNodeOption {
	return func(n *FocusNode) {
		n.skipTraversal = skip
	}
}

// HasFocus returns true if this node has focus
func (n *FocusNode) HasFocus() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.hasFocus
}

// HasPrimaryFocus returns true if this node has primary focus
func (n *FocusNode) HasPrimaryFocus() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.hasPrimaryFocus
}

// CanRequestFocus returns true if this node can request focus
func (n *FocusNode) CanRequestFocus() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.canRequestFocus
}

// RequestFocus requests focus for this node
func (n *FocusNode) RequestFocus() {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.canRequestFocus {
		return
	}

	// Notify focus manager
	if manager := GetFocusManager(); manager != nil {
		manager.RequestFocus(n)
	}
}

// Unfocus removes focus from this node
func (n *FocusNode) Unfocus() {
	if manager := GetFocusManager(); manager != nil {
		manager.Unfocus(n)
	}
}

// NextFocus moves focus to the next focusable node
func (n *FocusNode) NextFocus() {
	if manager := GetFocusManager(); manager != nil {
		manager.NextFocus()
	}
}

// PreviousFocus moves focus to the previous focusable node
func (n *FocusNode) PreviousFocus() {
	if manager := GetFocusManager(); manager != nil {
		manager.PreviousFocus()
	}
}

// HandleKeyEvent handles a keyboard event
func (n *FocusNode) HandleKeyEvent(event *KeyboardEvent) bool {
	n.mu.RLock()
	callback := n.onKey
	n.mu.RUnlock()

	if callback != nil {
		return callback(event)
	}
	return false
}

// Attach attaches a child node
func (n *FocusNode) Attach(child *FocusNode) {
	n.mu.Lock()
	defer n.mu.Unlock()

	child.parent = n
	n.children = append(n.children, child)
}

// Detach detaches a child node
func (n *FocusNode) Detach(child *FocusNode) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			child.parent = nil
			break
		}
	}
}

// Dispose cleans up the focus node
func (n *FocusNode) Dispose() {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Remove from parent
	if n.parent != nil {
		n.parent.Detach(n)
	}

	// Clear callbacks
	n.onFocusChange = nil
	n.onKey = nil

	// Clear children
	n.children = nil
}

// notifyFocusChange notifies listeners of focus change
func (n *FocusNode) notifyFocusChange(focused bool) {
	n.mu.Lock()
	n.hasFocus = focused
	callback := n.onFocusChange
	n.mu.Unlock()

	if callback != nil {
		callback(focused)
	}
}

// setPrimaryFocus sets whether this node has primary focus
func (n *FocusNode) setPrimaryFocus(primary bool) {
	n.mu.Lock()
	n.hasPrimaryFocus = primary
	n.mu.Unlock()
}

// FocusManager manages focus for the application
type FocusManager struct {
	mu sync.RWMutex

	rootScope    *FocusScopeNode
	primaryFocus *FocusNode
	focusedNode  *FocusNode
	traversalOrder []*FocusNode
}

var (
	globalFocusManager *FocusManager
	focusManagerOnce   sync.Once
)

// GetFocusManager returns the global focus manager
func GetFocusManager() *FocusManager {
	focusManagerOnce.Do(func() {
		globalFocusManager = NewFocusManager()
	})
	return globalFocusManager
}

// NewFocusManager creates a new focus manager
func NewFocusManager() *FocusManager {
	manager := &FocusManager{
		rootScope:      NewFocusScopeNode(),
		traversalOrder: make([]*FocusNode, 0),
	}
	return manager
}

// GetPrimaryFocus returns the node with primary focus
func (m *FocusManager) GetPrimaryFocus() *FocusNode {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.primaryFocus
}

// RequestFocus requests focus for a node
func (m *FocusManager) RequestFocus(node *FocusNode) {
	if node == nil || !node.CanRequestFocus() {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Unfocus current node
	if m.focusedNode != nil && m.focusedNode != node {
		m.focusedNode.notifyFocusChange(false)
		m.focusedNode.setPrimaryFocus(false)
	}

	// Focus new node
	m.focusedNode = node
	m.primaryFocus = node

	node.notifyFocusChange(true)
	node.setPrimaryFocus(true)
}

// Unfocus removes focus from a node
func (m *FocusManager) Unfocus(node *FocusNode) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.focusedNode == node {
		node.notifyFocusChange(false)
		node.setPrimaryFocus(false)
		m.focusedNode = nil
		m.primaryFocus = nil
	}
}

// NextFocus moves focus to the next node
func (m *FocusManager) NextFocus() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.traversalOrder) == 0 {
		return
	}

	currentIndex := -1
	if m.focusedNode != nil {
		for i, node := range m.traversalOrder {
			if node == m.focusedNode {
				currentIndex = i
				break
			}
		}
	}

	nextIndex := (currentIndex + 1) % len(m.traversalOrder)
	nextNode := m.traversalOrder[nextIndex]

	if nextNode != nil && nextNode.CanRequestFocus() {
		m.RequestFocus(nextNode)
	}
}

// PreviousFocus moves focus to the previous node
func (m *FocusManager) PreviousFocus() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.traversalOrder) == 0 {
		return
	}

	currentIndex := -1
	if m.focusedNode != nil {
		for i, node := range m.traversalOrder {
			if node == m.focusedNode {
				currentIndex = i
				break
			}
		}
	}

	prevIndex := currentIndex - 1
	if prevIndex < 0 {
		prevIndex = len(m.traversalOrder) - 1
	}

	prevNode := m.traversalOrder[prevIndex]

	if prevNode != nil && prevNode.CanRequestFocus() {
		m.RequestFocus(prevNode)
	}
}

// HandleKeyEvent routes keyboard events to the focused node
func (m *FocusManager) HandleKeyEvent(event *KeyboardEvent) bool {
	m.mu.RLock()
	focused := m.focusedNode
	m.mu.RUnlock()

	if focused != nil {
		return focused.HandleKeyEvent(event)
	}

	return false
}

// RegisterNode registers a node in the traversal order
func (m *FocusManager) RegisterNode(node *FocusNode) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if already registered
	for _, n := range m.traversalOrder {
		if n == node {
			return
		}
	}

	m.traversalOrder = append(m.traversalOrder, node)
}

// UnregisterNode removes a node from the traversal order
func (m *FocusManager) UnregisterNode(node *FocusNode) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.traversalOrder {
		if n == node {
			m.traversalOrder = append(m.traversalOrder[:i], m.traversalOrder[i+1:]...)
			break
		}
	}

	// If this was the focused node, clear focus
	if m.focusedNode == node {
		m.focusedNode = nil
		m.primaryFocus = nil
	}
}

// FocusScopeNode represents a scope in the focus tree
type FocusScopeNode struct {
	*FocusNode
	focusedChild *FocusNode
}

// NewFocusScopeNode creates a new focus scope node
func NewFocusScopeNode() *FocusScopeNode {
	return &FocusScopeNode{
		FocusNode: NewFocusNode(),
	}
}

// SetFocusedChild sets the focused child within this scope
func (n *FocusScopeNode) SetFocusedChild(child *FocusNode) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.focusedChild = child
}

// AutoFocus automatically focuses the first focusable descendant
func (n *FocusScopeNode) AutoFocus() {
	n.mu.RLock()
	children := n.children
	n.mu.RUnlock()

	for _, child := range children {
		if child.CanRequestFocus() && !child.skipTraversal {
			child.RequestFocus()
			return
		}
	}
}
