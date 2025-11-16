package input

import "time"

// KeyboardEvent represents keyboard input events
type KeyboardEvent struct {
	BaseEvent
	Key       Key           // Physical key
	Code      string        // Key code (e.g., "KeyA", "Digit1")
	Character string        // Character produced (respects shift, etc.)
	Modifiers KeyModifiers  // Active modifiers
	IsRepeat  bool          // True if this is a key repeat event
}

// NewKeyboardEvent creates a new keyboard event
func NewKeyboardEvent(eventType EventType, key Key, code string, character string, modifiers KeyModifiers, isRepeat bool) *KeyboardEvent {
	return &KeyboardEvent{
		BaseEvent: BaseEvent{
			EventType:  eventType,
			Time:       time.Now(),
			DeviceType: DeviceTypeKeyboard,
		},
		Key:       key,
		Code:      code,
		Character: character,
		Modifiers: modifiers,
		IsRepeat:  isRepeat,
	}
}

// Key represents physical keyboard keys
type Key int

const (
	KeyUnknown Key = iota

	// Alphanumeric keys
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	// Number keys
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9

	// Function keys
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12

	// Modifier keys
	KeyShift
	KeyShiftLeft
	KeyShiftRight
	KeyControl
	KeyControlLeft
	KeyControlRight
	KeyAlt
	KeyAltLeft
	KeyAltRight
	KeyMeta
	KeyMetaLeft
	KeyMetaRight

	// Special keys
	KeyEscape
	KeyEnter
	KeyTab
	KeyBackspace
	KeyDelete
	KeyInsert
	KeySpace
	KeyCapsLock
	KeyNumLock
	KeyScrollLock

	// Navigation keys
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyHome
	KeyEnd
	KeyPageUp
	KeyPageDown

	// Editing keys
	KeyCut
	KeyCopy
	KeyPaste
	KeyUndo
	KeyRedo

	// Numpad keys
	KeyNumpad0
	KeyNumpad1
	KeyNumpad2
	KeyNumpad3
	KeyNumpad4
	KeyNumpad5
	KeyNumpad6
	KeyNumpad7
	KeyNumpad8
	KeyNumpad9
	KeyNumpadAdd
	KeyNumpadSubtract
	KeyNumpadMultiply
	KeyNumpadDivide
	KeyNumpadDecimal
	KeyNumpadEnter

	// Punctuation and symbols
	KeyMinus
	KeyEqual
	KeyBracketLeft
	KeyBracketRight
	KeyBackslash
	KeySemicolon
	KeyQuote
	KeyComma
	KeyPeriod
	KeySlash
	KeyBackquote

	// Media keys
	KeyVolumeUp
	KeyVolumeDown
	KeyVolumeMute
	KeyMediaPlayPause
	KeyMediaStop
	KeyMediaNext
	KeyMediaPrevious

	// Browser keys
	KeyBrowserBack
	KeyBrowserForward
	KeyBrowserRefresh
	KeyBrowserHome
)

// KeyModifiers represents keyboard modifier keys state
type KeyModifiers int

const (
	ModifierNone  KeyModifiers = 0
	ModifierShift KeyModifiers = 1 << 0
	ModifierCtrl  KeyModifiers = 1 << 1
	ModifierAlt   KeyModifiers = 1 << 2
	ModifierMeta  KeyModifiers = 1 << 3 // Command on Mac, Windows key on Windows
)

// HasModifier checks if a modifier is active
func (m KeyModifiers) HasModifier(modifier KeyModifiers) bool {
	return m&modifier != 0
}

// HasShift checks if Shift is pressed
func (m KeyModifiers) HasShift() bool {
	return m.HasModifier(ModifierShift)
}

// HasCtrl checks if Control is pressed
func (m KeyModifiers) HasCtrl() bool {
	return m.HasModifier(ModifierCtrl)
}

// HasAlt checks if Alt is pressed
func (m KeyModifiers) HasAlt() bool {
	return m.HasModifier(ModifierAlt)
}

// HasMeta checks if Meta/Command is pressed
func (m KeyModifiers) HasMeta() bool {
	return m.HasModifier(ModifierMeta)
}

// KeyboardState tracks the current state of all keyboard keys
type KeyboardState struct {
	keys      map[Key]bool
	modifiers KeyModifiers
}

// NewKeyboardState creates a new keyboard state tracker
func NewKeyboardState() *KeyboardState {
	return &KeyboardState{
		keys: make(map[Key]bool),
	}
}

// IsKeyPressed checks if a key is currently pressed
func (s *KeyboardState) IsKeyPressed(key Key) bool {
	return s.keys[key]
}

// GetModifiers returns the current modifier state
func (s *KeyboardState) GetModifiers() KeyModifiers {
	return s.modifiers
}

// HandleEvent updates the keyboard state based on an event
func (s *KeyboardState) HandleEvent(event *KeyboardEvent) {
	switch event.Type() {
	case EventTypeKeyDown:
		s.keys[event.Key] = true
		s.updateModifiers(event.Key, true)
	case EventTypeKeyUp:
		s.keys[event.Key] = false
		s.updateModifiers(event.Key, false)
	}
}

// updateModifiers updates modifier state when modifier keys are pressed/released
func (s *KeyboardState) updateModifiers(key Key, pressed bool) {
	var modifier KeyModifiers

	switch key {
	case KeyShift, KeyShiftLeft, KeyShiftRight:
		modifier = ModifierShift
	case KeyControl, KeyControlLeft, KeyControlRight:
		modifier = ModifierCtrl
	case KeyAlt, KeyAltLeft, KeyAltRight:
		modifier = ModifierAlt
	case KeyMeta, KeyMetaLeft, KeyMetaRight:
		modifier = ModifierMeta
	default:
		return
	}

	if pressed {
		s.modifiers |= modifier
	} else {
		s.modifiers &^= modifier
	}
}

// Reset clears all keyboard state
func (s *KeyboardState) Reset() {
	s.keys = make(map[Key]bool)
	s.modifiers = ModifierNone
}

// KeyBinding represents a keyboard shortcut
type KeyBinding struct {
	Key       Key
	Modifiers KeyModifiers
	Action    func()
}

// Matches checks if the binding matches a keyboard event
func (kb *KeyBinding) Matches(event *KeyboardEvent) bool {
	return event.Key == kb.Key && event.Modifiers == kb.Modifiers
}

// KeyBindingManager manages keyboard shortcuts
type KeyBindingManager struct {
	bindings []*KeyBinding
}

// NewKeyBindingManager creates a new key binding manager
func NewKeyBindingManager() *KeyBindingManager {
	return &KeyBindingManager{
		bindings: make([]*KeyBinding, 0),
	}
}

// AddBinding adds a keyboard shortcut
func (m *KeyBindingManager) AddBinding(key Key, modifiers KeyModifiers, action func()) {
	m.bindings = append(m.bindings, &KeyBinding{
		Key:       key,
		Modifiers: modifiers,
		Action:    action,
	})
}

// HandleEvent processes a keyboard event and triggers matching bindings
func (m *KeyBindingManager) HandleEvent(event *KeyboardEvent) bool {
	if event.Type() != EventTypeKeyDown {
		return false
	}

	for _, binding := range m.bindings {
		if binding.Matches(event) {
			binding.Action()
			return true
		}
	}

	return false
}

// RemoveBinding removes a keyboard shortcut
func (m *KeyBindingManager) RemoveBinding(key Key, modifiers KeyModifiers) {
	for i, binding := range m.bindings {
		if binding.Key == key && binding.Modifiers == modifiers {
			m.bindings = append(m.bindings[:i], m.bindings[i+1:]...)
			return
		}
	}
}

// Clear removes all bindings
func (m *KeyBindingManager) Clear() {
	m.bindings = make([]*KeyBinding, 0)
}

// KeyEvent is an alias for KeyboardEvent for convenience
type KeyEvent = KeyboardEvent
