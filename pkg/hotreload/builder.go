package hotreload

import (
	"fmt"
	"log"
)

// PartialRebuildManager manages partial rebuilds of the widget tree
type PartialRebuildManager struct {
	stateStore *StateStore
	differ     *Differ
}

// NewPartialRebuildManager creates a new partial rebuild manager
func NewPartialRebuildManager(stateStore *StateStore) *PartialRebuildManager {
	return &PartialRebuildManager{
		stateStore: stateStore,
		differ:     NewDiffer(),
	}
}

// MarkDirty marks a widget as needing rebuild
func (m *PartialRebuildManager) MarkDirty(widgetPath string) {
	m.stateStore.Set(fmt.Sprintf("dirty:%s", widgetPath), true)
	log.Printf("Marked widget as dirty: %s", widgetPath)
}

// IsDirty checks if a widget is marked as dirty
func (m *PartialRebuildManager) IsDirty(widgetPath string) bool {
	val, ok := m.stateStore.Get(fmt.Sprintf("dirty:%s", widgetPath))
	if !ok {
		return false
	}
	dirty, ok := val.(bool)
	return ok && dirty
}

// ClearDirty clears the dirty flag for a widget
func (m *PartialRebuildManager) ClearDirty(widgetPath string) {
	m.stateStore.Delete(fmt.Sprintf("dirty:%s", widgetPath))
}

// ComputeDiff computes the difference between old and new widget trees
func (m *PartialRebuildManager) ComputeDiff(oldTree, newTree interface{}) []WidgetDiff {
	return m.differ.Diff(oldTree, newTree, []string{"root"})
}

// RebuildStrategy determines the rebuild strategy based on diffs
type RebuildStrategy int

const (
	// RebuildFull rebuilds the entire tree
	RebuildFull RebuildStrategy = iota

	// RebuildPartial rebuilds only changed subtrees
	RebuildPartial

	// RebuildNone no rebuild needed
	RebuildNone
)

// DetermineStrategy determines the appropriate rebuild strategy
func (m *PartialRebuildManager) DetermineStrategy(diffs []WidgetDiff) RebuildStrategy {
	if len(diffs) == 0 {
		return RebuildNone
	}

	// If there are too many diffs, do a full rebuild
	if len(diffs) > 100 {
		return RebuildFull
	}

	// Check if any diffs are at the root level
	for _, diff := range diffs {
		if len(diff.Path) <= 1 {
			return RebuildFull
		}
	}

	// Otherwise, do a partial rebuild
	return RebuildPartial
}

// GetAffectedPaths returns the paths of widgets that need rebuilding
func (m *PartialRebuildManager) GetAffectedPaths(diffs []WidgetDiff) []string {
	pathSet := make(map[string]bool)

	for _, diff := range diffs {
		// Get the parent path (rebuild from parent level)
		if len(diff.Path) > 0 {
			parentPath := ""
			if len(diff.Path) > 1 {
				parentPath = diff.Path[0]
				for i := 1; i < len(diff.Path)-1; i++ {
					parentPath += "." + diff.Path[i]
				}
			} else {
				parentPath = diff.Path[0]
			}
			pathSet[parentPath] = true
		}
	}

	paths := make([]string, 0, len(pathSet))
	for path := range pathSet {
		paths = append(paths, path)
	}

	return paths
}
