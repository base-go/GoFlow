# REST API Library Implementation

**Date:** 2025-11-16
**Branch:** claude/rest-api-libraries-017raMCMgMaSRSLBcArTrSDA

## Overview

This document describes the implementation of the REST API library for GoFlow, which provides a signal-based reactive approach to making HTTP requests that integrates seamlessly with GoFlow's widget system.

## Motivation

GoFlow widgets need a way to fetch data from REST APIs that:
- Integrates with the Signals reactive system
- Automatically triggers widget rebuilds when data changes
- Provides clear loading and error states
- Is type-safe and easy to use
- Follows familiar patterns from Flutter and modern web frameworks

## Architecture

### Core Components

```
pkg/api/
├── client.go          # HTTP client with middleware support
├── response.go        # Response wrapper with helper methods
├── resource.go        # Signal-based Resource[T] and ResourceList[T]
├── fetcher.go         # Helper functions for common requests
├── provider.go        # Context provider for API client
├── middleware/
│   ├── auth.go        # Authentication middleware
│   ├── logging.go     # Request/response logging
│   └── retry.go       # Retry logic with exponential backoff
└── services/
    ├── user_service.go    # Example user service
    └── post_service.go    # Example post service
```

### Design Patterns

#### 1. Signal-Based Resources

Resources wrap API data in Signals that trigger widget rebuilds:

```go
type ResourceState[T any] struct {
    Data    *T
    Loading bool
    Error   error
}

type Resource[T any] struct {
    state *signals.Signal[ResourceState[T]]
    // ... other fields
}
```

This allows widgets to reactively respond to data changes:

```go
func (w *MyWidget) Build(ctx BuildContext) Widget {
    state := w.resource.Get() // Creates reactive dependency

    if state.Loading { return LoadingWidget() }
    if state.Error != nil { return ErrorWidget(state.Error) }
    return DataWidget(state.Data)
}
```

#### 2. Middleware Chain

Middleware wraps HTTP requests in a composable chain:

```go
type Middleware func(next RequestHandler) RequestHandler

client.Use(middleware.Logging(nil)).
       Use(middleware.BearerAuth("token")).
       Use(middleware.Retry(config))
```

Each middleware can:
- Modify the request before sending
- Inspect/modify the response
- Handle errors and retries
- Add headers, logging, timing, etc.

#### 3. Service Layer Pattern

Services encapsulate related API endpoints:

```go
type UserService struct {
    client *api.Client
}

func (s *UserService) GetUser(ctx context.Context, id int) *api.Resource[User] {
    return api.Get[User](ctx, s.client, fmt.Sprintf("/users/%d", id))
}
```

This provides:
- Separation of concerns
- Testability (easy to mock)
- Reusability across widgets
- Type safety with generics

#### 4. Context Provider Pattern

API client is provided through BuildContext:

```go
&api.APIProvider{
    Client: apiClient,
    Child: &MyApp{},
}

// In descendants:
apiClient := api.GetClient(ctx)
```

## Implementation Details

### Client

The `Client` struct (`client.go:22-28`) provides:
- Base URL configuration
- Timeout settings
- Default headers
- Middleware chain
- HTTP verb methods (GET, POST, PUT, PATCH, DELETE)

### Resource

The `Resource[T]` type (`resource.go:14-22`) provides:
- Thread-safe state management
- Reactive signal-based updates
- Cancellation support
- Helper methods for common checks

### Middleware

Three built-in middleware types:

1. **Authentication** (`middleware/auth.go`)
   - Bearer token
   - Basic auth
   - API key headers

2. **Logging** (`middleware/logging.go`)
   - Request/response logging
   - Timing information
   - Custom logger support

3. **Retry** (`middleware/retry.go`)
   - Exponential backoff
   - Configurable retry conditions
   - Max retries setting

### Helper Functions

`fetcher.go` provides convenience functions:
- `Get[T]()` - GET request returning Resource[T]
- `Post[T]()` - POST request returning Resource[T]
- `FetchList[T]()` - Request returning ResourceList[T]

## Usage Example

### Setup

```go
// Create API client
apiClient := api.NewClient(api.ClientConfig{
    BaseURL: "https://api.example.com",
}).Use(middleware.Logging(nil))

// Wrap app with provider
app := &api.APIProvider{
    Client: apiClient,
    Child: &HomePage{},
}
```

### In Widgets

```go
type UserProfile struct {
    goflow.BaseWidget
    userResource *api.Resource[User]
    initialized  *signals.Signal[bool]
}

func (w *UserProfile) Build(ctx goflow.BuildContext) goflow.Widget {
    // Initialize once
    if w.initialized == nil {
        w.initialized = signals.NewSignal(false)
        apiClient := api.GetClient(ctx)
        userService := NewUserService(apiClient)
        w.userResource = userService.GetUser(context.Background(), 1)
        w.initialized.Set(true)
    }

    // Get reactive state
    state := w.userResource.Get()

    // Handle states
    if state.Loading {
        return &widgets.Text{Text: "Loading..."}
    }
    if state.Error != nil {
        return &widgets.Text{Text: fmt.Sprintf("Error: %v", state.Error)}
    }

    // Display data
    return &widgets.Text{Text: state.Data.Name}
}
```

## Examples

### Simple User List (`examples/api_demo/main.go`)

Demonstrates:
- Fetching a list of users
- Displaying loading state
- Error handling
- Rendering list items

### Interactive Posts (`examples/api_demo/posts_widget.go`)

Demonstrates:
- User posts by ID
- Expandable comments
- Multiple resources in one widget
- Interactive state with signals

## Testing

Comprehensive tests in:
- `client_test.go` - HTTP client tests, middleware chain
- `resource_test.go` - Resource state management

To run tests:
```bash
go test ./pkg/api/... -v
```

## Benefits

### For Widget Developers

✅ **Simple API**: Familiar GET/POST/PUT/DELETE methods
✅ **Reactive**: Automatic widget rebuilds
✅ **Type Safe**: Compile-time type checking
✅ **Error Handling**: Built-in error state
✅ **Loading States**: Automatic loading indicators

### For Application Architects

✅ **Testable**: Easy to mock services
✅ **Extensible**: Middleware pattern
✅ **Organized**: Service layer pattern
✅ **Reusable**: Share services across widgets
✅ **Maintainable**: Clear separation of concerns

## Future Enhancements

Potential improvements:

1. **Caching Layer**
   - Response caching
   - Cache invalidation
   - TTL support

2. **Optimistic Updates**
   - Update UI before server confirms
   - Rollback on error

3. **Pagination Support**
   - Built-in pagination helpers
   - Infinite scroll support

4. **GraphQL Support**
   - GraphQL client
   - Query builder

5. **WebSocket Support**
   - Real-time updates
   - SSE (Server-Sent Events)

6. **Request Debouncing**
   - Prevent duplicate requests
   - Cancel in-flight requests

7. **Upload/Download Progress**
   - Track progress
   - Show progress bars

## Integration with GoFlow Roadmap

This implementation advances **Phase 4.4** (Q4 2025) of the GoFlow roadmap:
- ✅ HTTP client with middleware
- ✅ JSON serialization helpers
- ✅ REST API builder pattern
- ⏳ WebSocket support (future)
- ⏳ GraphQL client (future)

## Documentation

See `pkg/api/README.md` for:
- Complete API reference
- Best practices
- Advanced examples
- Migration guide

## Conclusion

The REST API library provides a robust, reactive foundation for making HTTP requests in GoFlow applications. It integrates seamlessly with the Signals system and provides patterns familiar to developers from Flutter and modern web frameworks.

The library is production-ready for basic use cases and provides a solid foundation for future enhancements.
