# GoFlow REST API Library

A comprehensive REST API client library for GoFlow that integrates seamlessly with the reactive Signals system.

## Features

- **Signal-Based Resources**: Reactive API state management that automatically updates widgets
- **Middleware Support**: Extensible middleware for auth, logging, retry logic, and more
- **Type-Safe**: Generic-based API that ensures compile-time type safety
- **Easy to Use**: Simple, declarative API for making HTTP requests
- **Widget-Friendly**: Designed specifically for GoFlow's reactive widget system

## Quick Start

### 1. Create an API Client

```go
import (
    "github.com/base-go/GoFlow/pkg/api"
    "github.com/base-go/GoFlow/pkg/api/middleware"
)

// Create client with middleware
apiClient := api.NewClient(api.ClientConfig{
    BaseURL: "https://api.example.com",
}).Use(middleware.Logging(nil)).
   Use(middleware.BearerAuth("your-token"))
```

### 2. Define Your Models

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### 3. Create a Service

```go
type UserService struct {
    client *api.Client
}

func NewUserService(client *api.Client) *UserService {
    return &UserService{client: client}
}

func (s *UserService) GetUser(ctx context.Context, id int) *api.Resource[User] {
    path := fmt.Sprintf("/users/%d", id)
    return api.Get[User](ctx, s.client, path)
}

func (s *UserService) GetUsers(ctx context.Context) *api.ResourceList[User] {
    return api.FetchList[User](ctx, s.client, "GET", "/users", nil)
}
```

### 4. Use in Widgets

```go
type UserProfile struct {
    goflow.BaseWidget
    userResource *api.Resource[User]
    initialized  *signals.Signal[bool]
}

func (w *UserProfile) Build(ctx goflow.BuildContext) goflow.Widget {
    // Initialize on first build
    if w.initialized == nil {
        w.initialized = signals.NewSignal(false)
        apiClient := api.GetClient(ctx)
        userService := NewUserService(apiClient)
        w.userResource = userService.GetUser(context.Background(), 1)
        w.initialized.Set(true)
    }

    // Get reactive state
    state := w.userResource.Get()

    // Handle loading
    if state.Loading {
        return &widgets.Text{Text: "Loading..."}
    }

    // Handle error
    if state.Error != nil {
        return &widgets.Text{Text: fmt.Sprintf("Error: %v", state.Error)}
    }

    // Display data
    return &widgets.Text{Text: state.Data.Name}
}
```

### 5. Provide API Client to Widget Tree

```go
func main() {
    apiClient := api.NewClient(api.ClientConfig{
        BaseURL: "https://api.example.com",
    })

    app := &api.APIProvider{
        Client: apiClient,
        Child:  &MyApp{},
    }

    goflow.RunApp(app)
}
```

## Core Concepts

### Resources

Resources are signal-based wrappers around API data that automatically trigger widget rebuilds when data changes.

#### Resource[T]

For single objects:

```go
userResource := api.Get[User](ctx, client, "/users/1")

// In widget
state := userResource.Get()
if state.Loading { /* show loading */ }
if state.Error != nil { /* show error */ }
// Use state.Data
```

#### ResourceList[T]

For arrays:

```go
usersResource := api.FetchList[User](ctx, client, "GET", "/users", nil)

// In widget
state := usersResource.Get()
users := state.Data // []User
```

### Resource State

Every resource has three states:

```go
type ResourceState[T any] struct {
    Data    *T      // The data (nil if not loaded)
    Loading bool    // True while loading
    Error   error   // Error if request failed
}
```

### Client

The HTTP client with middleware support:

```go
client := api.NewClient(api.ClientConfig{
    BaseURL: "https://api.example.com",
    Timeout: 30 * time.Second,
    Headers: map[string]string{
        "X-Custom-Header": "value",
    },
})
```

### Middleware

Add middleware to intercept and modify requests:

```go
import "github.com/base-go/GoFlow/pkg/api/middleware"

client.Use(middleware.Logging(nil)).
       Use(middleware.BearerAuth("token")).
       Use(middleware.Retry(middleware.RetryConfig{
           MaxRetries: 3,
           RetryDelay: 1 * time.Second,
       }))
```

#### Available Middleware

- **BearerAuth**: Add Bearer token authentication
- **BasicAuth**: Add basic authentication
- **APIKey**: Add API key header
- **Logging**: Log requests and responses
- **Retry**: Retry failed requests with exponential backoff

#### Custom Middleware

```go
func CustomMiddleware() api.Middleware {
    return func(next api.RequestHandler) api.RequestHandler {
        return func(req *http.Request) (*http.Response, error) {
            // Modify request
            req.Header.Set("X-Custom", "value")

            // Call next middleware
            resp, err := next(req)

            // Handle response
            return resp, err
        }
    }
}

client.Use(CustomMiddleware())
```

## API Reference

### Client Methods

```go
// HTTP verbs
Get(ctx context.Context, path string) (*Response, error)
Post(ctx context.Context, path string, body interface{}) (*Response, error)
Put(ctx context.Context, path string, body interface{}) (*Response, error)
Patch(ctx context.Context, path string, body interface{}) (*Response, error)
Delete(ctx context.Context, path string) (*Response, error)

// Configuration
Use(middleware Middleware) *Client
SetHeader(key, value string) *Client
```

### Helper Functions

```go
// Fetch single resource
Get[T](ctx context.Context, client *Client, path string) *Resource[T]
Post[T](ctx context.Context, client *Client, path string, body interface{}) *Resource[T]
Put[T](ctx context.Context, client *Client, path string, body interface{}) *Resource[T]
Delete[T](ctx context.Context, client *Client, path string) *Resource[T]

// Fetch list resource
FetchList[T](ctx context.Context, client *Client, method, path string, body interface{}) *ResourceList[T]
```

### Resource Methods

```go
// Get current state (creates reactive dependency)
Get() ResourceState[T]

// Get just the data
GetData() *T

// Check state
IsLoading() bool
HasError() bool
GetError() error

// Update state
SetLoading()
SetData(data T)
SetError(err error)

// Cancel ongoing request
Cancel()
```

## Examples

See the `examples/api_demo` directory for a complete working example that demonstrates:

- Setting up an API client with middleware
- Creating services (UserService, PostService)
- Using resources in widgets
- Handling loading and error states
- Interactive UI with comments

Run the example:

```bash
cd examples/api_demo
go run .
```

## Best Practices

### 1. Initialize Resources Once

```go
type MyWidget struct {
    goflow.BaseWidget
    resource    *api.Resource[Data]
    initialized *signals.Signal[bool]
}

func (w *MyWidget) Build(ctx goflow.BuildContext) goflow.Widget {
    if w.initialized == nil {
        w.initialized = signals.NewSignal(false)
        w.resource = w.loadData(ctx)
        w.initialized.Set(true)
    }
    // Use w.resource...
}
```

### 2. Use Services for Organization

Group related API calls into services:

```go
type UserService struct {
    client *api.Client
}

func (s *UserService) GetUser(ctx context.Context, id int) *api.Resource[User]
func (s *UserService) UpdateUser(ctx context.Context, id int, user User) *api.Resource[User]
func (s *UserService) DeleteUser(ctx context.Context, id int) *api.Resource[interface{}]
```

### 3. Handle All States

Always handle loading, error, and success states:

```go
state := resource.Get()

if state.Loading {
    return &widgets.LoadingSpinner{}
}

if state.Error != nil {
    return &widgets.ErrorDisplay{Error: state.Error}
}

return &widgets.DataDisplay{Data: state.Data}
```

### 4. Use Context Properly

Pass context to cancel requests when widgets unmount:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

resource := api.Get[User](ctx, client, "/users/1")
```

### 5. Provide Client at Root

Use `APIProvider` at the root of your app:

```go
&api.APIProvider{
    Client: apiClient,
    Child:  &MyApp{},
}
```

Then access it in descendants:

```go
apiClient := api.GetClient(ctx)
```

## Architecture

The library follows these design principles:

1. **Reactive First**: Resources are signals that trigger widget rebuilds
2. **Type Safety**: Generics ensure compile-time type checking
3. **Separation of Concerns**: Client → Service → Widget layers
4. **Middleware Pattern**: Composable request/response interceptors
5. **Context Propagation**: API clients passed through BuildContext

```
┌─────────────────────────────────────┐
│  Widget Layer (UI)                  │
│  - Displays data                    │
│  - Handles user interaction         │
├─────────────────────────────────────┤
│  Service Layer                      │
│  - Business logic                   │
│  - API calls                        │
│  - Returns Resources                │
├─────────────────────────────────────┤
│  API Client Layer                   │
│  - HTTP requests                    │
│  - Middleware chain                 │
│  - Response handling                │
├─────────────────────────────────────┤
│  Resource Layer                     │
│  - Signal-based state               │
│  - Reactive updates                 │
└─────────────────────────────────────┘
```

## Contributing

Contributions are welcome! Please follow the GoFlow contribution guidelines.

## License

This library is part of the GoFlow project and follows the same license.
