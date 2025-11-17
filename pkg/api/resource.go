package api

import (
	"context"
	"sync"

	"github.com/base-go/GoFlow/pkg/core/signals"
)

// ResourceState represents the state of an API resource
type ResourceState[T any] struct {
	Data    *T
	Loading bool
	Error   error
}

// Resource is a signal-based reactive API resource
type Resource[T any] struct {
	state  *signals.Signal[ResourceState[T]]
	mutex  sync.RWMutex
	cancel context.CancelFunc
}

// NewResource creates a new Resource
func NewResource[T any]() *Resource[T] {
	return &Resource[T]{
		state: signals.New(ResourceState[T]{
			Loading: false,
		}),
	}
}

// NewLoadingResource creates a new Resource in loading state
func NewLoadingResource[T any]() *Resource[T] {
	return &Resource[T]{
		state: signals.New(ResourceState[T]{
			Loading: true,
		}),
	}
}

// Get returns the current state (creates reactive dependency in widgets)
func (r *Resource[T]) Get() ResourceState[T] {
	return r.state.Get()
}

// GetData returns just the data (or nil)
func (r *Resource[T]) GetData() *T {
	return r.state.Get().Data
}

// IsLoading returns true if the resource is currently loading
func (r *Resource[T]) IsLoading() bool {
	return r.state.Get().Loading
}

// HasError returns true if the resource has an error
func (r *Resource[T]) HasError() bool {
	return r.state.Get().Error != nil
}

// GetError returns the current error (or nil)
func (r *Resource[T]) GetError() error {
	return r.state.Get().Error
}

// SetLoading sets the resource to loading state
func (r *Resource[T]) SetLoading() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	current := r.state.Get()
	r.state.Set(ResourceState[T]{
		Data:    current.Data, // Preserve existing data
		Loading: true,
		Error:   nil,
	})
}

// SetData sets the resource data and clears loading/error states
func (r *Resource[T]) SetData(data T) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.state.Set(ResourceState[T]{
		Data:    &data,
		Loading: false,
		Error:   nil,
	})
}

// SetError sets the resource error and clears loading state
func (r *Resource[T]) SetError(err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	current := r.state.Get()
	r.state.Set(ResourceState[T]{
		Data:    current.Data, // Preserve existing data
		Loading: false,
		Error:   err,
	})
}

// Update updates the resource based on a response
func (r *Resource[T]) Update(data *T, err error) {
	if err != nil {
		r.SetError(err)
	} else if data != nil {
		r.SetData(*data)
	}
}

// Cancel cancels any ongoing request
func (r *Resource[T]) Cancel() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.cancel != nil {
		r.cancel()
		r.cancel = nil
	}
}

// setCancel stores a cancel function for the current request
func (r *Resource[T]) setCancel(cancel context.CancelFunc) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Cancel any previous request
	if r.cancel != nil {
		r.cancel()
	}
	r.cancel = cancel
}

// ResourceList is a signal-based reactive list of resources
type ResourceList[T any] struct {
	state *signals.Signal[ResourceState[[]T]]
	mutex sync.RWMutex
}

// NewResourceList creates a new ResourceList
func NewResourceList[T any]() *ResourceList[T] {
	return &ResourceList[T]{
		state: signals.New(ResourceState[[]T]{
			Loading: false,
		}),
	}
}

// Get returns the current state
func (r *ResourceList[T]) Get() ResourceState[[]T] {
	return r.state.Get()
}

// GetData returns just the data slice (or nil)
func (r *ResourceList[T]) GetData() []T {
	data := r.state.Get().Data
	if data == nil {
		return nil
	}
	return *data
}

// IsLoading returns true if the resource is currently loading
func (r *ResourceList[T]) IsLoading() bool {
	return r.state.Get().Loading
}

// SetLoading sets the resource to loading state
func (r *ResourceList[T]) SetLoading() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	current := r.state.Get()
	r.state.Set(ResourceState[[]T]{
		Data:    current.Data,
		Loading: true,
		Error:   nil,
	})
}

// SetData sets the resource data
func (r *ResourceList[T]) SetData(data []T) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.state.Set(ResourceState[[]T]{
		Data:    &data,
		Loading: false,
		Error:   nil,
	})
}

// SetError sets the resource error
func (r *ResourceList[T]) SetError(err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	current := r.state.Get()
	r.state.Set(ResourceState[[]T]{
		Data:    current.Data,
		Loading: false,
		Error:   err,
	})
}
