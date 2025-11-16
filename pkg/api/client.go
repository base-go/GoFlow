package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Middleware is a function that wraps an HTTP request
type Middleware func(next RequestHandler) RequestHandler

// RequestHandler is a function that handles an HTTP request
type RequestHandler func(req *http.Request) (*http.Response, error)

// Client is a REST API client with middleware support
type Client struct {
	baseURL    string
	httpClient *http.Client
	middleware []Middleware
	headers    map[string]string
}

// ClientConfig configures a new API client
type ClientConfig struct {
	BaseURL string
	Timeout time.Duration
	Headers map[string]string
}

// NewClient creates a new API client
func NewClient(config ClientConfig) *Client {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &Client{
		baseURL: config.BaseURL,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		middleware: []Middleware{},
		headers:    config.Headers,
	}
}

// Use adds middleware to the client
func (c *Client) Use(middleware Middleware) *Client {
	c.middleware = append(c.middleware, middleware)
	return c
}

// SetHeader sets a default header for all requests
func (c *Client) SetHeader(key, value string) *Client {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[key] = value
	return c
}

// Get makes a GET request
func (c *Client) Get(ctx context.Context, path string) (*Response, error) {
	return c.Request(ctx, "GET", path, nil)
}

// Post makes a POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.Request(ctx, "POST", path, body)
}

// Put makes a PUT request
func (c *Client) Put(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.Request(ctx, "PUT", path, body)
}

// Patch makes a PATCH request
func (c *Client) Patch(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.Request(ctx, "PATCH", path, body)
}

// Delete makes a DELETE request
func (c *Client) Delete(ctx context.Context, path string) (*Response, error) {
	return c.Request(ctx, "DELETE", path, nil)
}

// Request makes an HTTP request with middleware chain
func (c *Client) Request(ctx context.Context, method, path string, body interface{}) (*Response, error) {
	url := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Set content type for requests with body
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Build middleware chain
	handler := c.buildHandler()

	// Execute request through middleware chain
	resp, err := handler(req)
	if err != nil {
		return nil, err
	}

	return NewResponse(resp)
}

// buildHandler builds the middleware chain
func (c *Client) buildHandler() RequestHandler {
	// Start with the actual HTTP call
	handler := func(req *http.Request) (*http.Response, error) {
		return c.httpClient.Do(req)
	}

	// Apply middleware in reverse order
	for i := len(c.middleware) - 1; i >= 0; i-- {
		handler = c.middleware[i](handler)
	}

	return handler
}
