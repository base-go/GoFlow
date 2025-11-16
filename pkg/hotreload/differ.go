package hotreload

import (
	"fmt"
	"reflect"
)

// WidgetDiff represents a difference between two widget trees
type WidgetDiff struct {
	Type     DiffType
	Path     []string // Path to the widget in the tree
	OldValue interface{}
	NewValue interface{}
}

// DiffType represents the type of difference
type DiffType int

const (
	DiffTypeNone DiffType = iota
	DiffTypeAdded
	DiffTypeRemoved
	DiffTypeModified
	DiffTypeReordered
)

func (d DiffType) String() string {
	switch d {
	case DiffTypeNone:
		return "None"
	case DiffTypeAdded:
		return "Added"
	case DiffTypeRemoved:
		return "Removed"
	case DiffTypeModified:
		return "Modified"
	case DiffTypeReordered:
		return "Reordered"
	default:
		return "Unknown"
	}
}

// Differ compares widget trees and identifies differences
type Differ struct {
	diffs []WidgetDiff
}

// NewDiffer creates a new widget tree differ
func NewDiffer() *Differ {
	return &Differ{
		diffs: make([]WidgetDiff, 0),
	}
}

// Diff compares two widget trees and returns the differences
func (d *Differ) Diff(oldTree, newTree interface{}, path []string) []WidgetDiff {
	d.diffs = make([]WidgetDiff, 0)
	d.compareValues(oldTree, newTree, path)
	return d.diffs
}

// compareValues compares two values recursively
func (d *Differ) compareValues(oldVal, newVal interface{}, path []string) {
	// Handle nil cases
	if oldVal == nil && newVal == nil {
		return
	}

	if oldVal == nil {
		d.addDiff(DiffTypeAdded, path, oldVal, newVal)
		return
	}

	if newVal == nil {
		d.addDiff(DiffTypeRemoved, path, oldVal, newVal)
		return
	}

	// Get reflect values
	oldReflect := reflect.ValueOf(oldVal)
	newReflect := reflect.ValueOf(newVal)

	// If types are different, consider it a modification
	if oldReflect.Type() != newReflect.Type() {
		d.addDiff(DiffTypeModified, path, oldVal, newVal)
		return
	}

	// Compare based on kind
	switch oldReflect.Kind() {
	case reflect.Struct:
		d.compareStructs(oldReflect, newReflect, path)
	case reflect.Slice, reflect.Array:
		d.compareSlices(oldReflect, newReflect, path)
	case reflect.Map:
		d.compareMaps(oldReflect, newReflect, path)
	case reflect.Ptr:
		if oldReflect.IsNil() && newReflect.IsNil() {
			return
		}
		if oldReflect.IsNil() || newReflect.IsNil() {
			d.addDiff(DiffTypeModified, path, oldVal, newVal)
			return
		}
		d.compareValues(oldReflect.Elem().Interface(), newReflect.Elem().Interface(), path)
	default:
		// For primitive types, check equality
		if !reflect.DeepEqual(oldVal, newVal) {
			d.addDiff(DiffTypeModified, path, oldVal, newVal)
		}
	}
}

// compareStructs compares two structs field by field
func (d *Differ) compareStructs(oldVal, newVal reflect.Value, path []string) {
	t := oldVal.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		fieldPath := append(path, field.Name)
		oldField := oldVal.Field(i)
		newField := newVal.Field(i)

		d.compareValues(oldField.Interface(), newField.Interface(), fieldPath)
	}
}

// compareSlices compares two slices
func (d *Differ) compareSlices(oldVal, newVal reflect.Value, path []string) {
	oldLen := oldVal.Len()
	newLen := newVal.Len()

	// Check if lengths differ
	if oldLen != newLen {
		d.addDiff(DiffTypeReordered, path, oldLen, newLen)
	}

	// Compare common elements
	minLen := oldLen
	if newLen < minLen {
		minLen = newLen
	}

	for i := 0; i < minLen; i++ {
		elemPath := append(path, fmt.Sprintf("[%d]", i))
		d.compareValues(oldVal.Index(i).Interface(), newVal.Index(i).Interface(), elemPath)
	}

	// Handle added elements
	for i := oldLen; i < newLen; i++ {
		elemPath := append(path, fmt.Sprintf("[%d]", i))
		d.addDiff(DiffTypeAdded, elemPath, nil, newVal.Index(i).Interface())
	}

	// Handle removed elements
	for i := newLen; i < oldLen; i++ {
		elemPath := append(path, fmt.Sprintf("[%d]", i))
		d.addDiff(DiffTypeRemoved, elemPath, oldVal.Index(i).Interface(), nil)
	}
}

// compareMaps compares two maps
func (d *Differ) compareMaps(oldVal, newVal reflect.Value, path []string) {
	// Check for removed keys
	for _, key := range oldVal.MapKeys() {
		keyPath := append(path, fmt.Sprintf("[%v]", key.Interface()))
		newValue := newVal.MapIndex(key)

		if !newValue.IsValid() {
			d.addDiff(DiffTypeRemoved, keyPath, oldVal.MapIndex(key).Interface(), nil)
		} else {
			d.compareValues(oldVal.MapIndex(key).Interface(), newValue.Interface(), keyPath)
		}
	}

	// Check for added keys
	for _, key := range newVal.MapKeys() {
		oldValue := oldVal.MapIndex(key)
		if !oldValue.IsValid() {
			keyPath := append(path, fmt.Sprintf("[%v]", key.Interface()))
			d.addDiff(DiffTypeAdded, keyPath, nil, newVal.MapIndex(key).Interface())
		}
	}
}

// addDiff adds a difference to the list
func (d *Differ) addDiff(diffType DiffType, path []string, oldValue, newValue interface{}) {
	// Create a copy of the path
	pathCopy := make([]string, len(path))
	copy(pathCopy, path)

	d.diffs = append(d.diffs, WidgetDiff{
		Type:     diffType,
		Path:     pathCopy,
		OldValue: oldValue,
		NewValue: newValue,
	})
}

// ApplyDiff applies a set of diffs to rebuild only the changed parts
func ApplyDiff(diffs []WidgetDiff) error {
	// This is where we would apply the diffs to the widget tree
	// For now, just log the diffs
	for _, diff := range diffs {
		fmt.Printf("Diff [%s]: %v -> %v at %v\n",
			diff.Type, diff.OldValue, diff.NewValue, diff.Path)
	}
	return nil
}
