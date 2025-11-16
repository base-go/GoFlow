package goflow

// Key is used to identify widgets in the tree
// Keys help the framework determine which widgets to reuse vs recreate
type Key interface {
	// String returns a string representation of the key
	String() string
}

// ValueKey is a key that uses a value for equality
type ValueKey struct {
	value string
}

// NewValueKey creates a new ValueKey
func NewValueKey(value string) Key {
	return &ValueKey{value: value}
}

// String returns the string representation
func (k *ValueKey) String() string {
	return k.value
}

// ObjectKey is a key that uses object identity
type ObjectKey struct {
	value interface{}
}

// NewObjectKey creates a new ObjectKey
func NewObjectKey(value interface{}) Key {
	return &ObjectKey{value: value}
}

// String returns the string representation
func (k *ObjectKey) String() string {
	return "ObjectKey"
}
