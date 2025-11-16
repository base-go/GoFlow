package api

import (
	"context"
	"fmt"
)

// Fetcher provides helper methods to fetch data into Resources
type Fetcher struct {
	client *Client
}

// NewFetcher creates a new Fetcher
func NewFetcher(client *Client) *Fetcher {
	return &Fetcher{client: client}
}

// Fetch makes a request and loads the result into a Resource
func Fetch[T any](ctx context.Context, client *Client, method, path string, body interface{}) *Resource[T] {
	resource := NewLoadingResource[T]()

	go func() {
		// Create cancellable context
		ctx, cancel := context.WithCancel(ctx)
		resource.setCancel(cancel)

		var resp *Response
		var err error

		switch method {
		case "GET":
			resp, err = client.Get(ctx, path)
		case "POST":
			resp, err = client.Post(ctx, path, body)
		case "PUT":
			resp, err = client.Put(ctx, path, body)
		case "PATCH":
			resp, err = client.Patch(ctx, path, body)
		case "DELETE":
			resp, err = client.Delete(ctx, path)
		default:
			err = fmt.Errorf("unsupported HTTP method: %s", method)
		}

		if err != nil {
			resource.SetError(err)
			return
		}

		if !resp.IsSuccess() {
			resource.SetError(resp.Error())
			return
		}

		var data T
		if err := resp.JSON(&data); err != nil {
			resource.SetError(err)
			return
		}

		resource.SetData(data)
	}()

	return resource
}

// FetchList makes a request and loads the result into a ResourceList
func FetchList[T any](ctx context.Context, client *Client, method, path string, body interface{}) *ResourceList[T] {
	resource := NewResourceList[T]()
	resource.SetLoading()

	go func() {
		var resp *Response
		var err error

		switch method {
		case "GET":
			resp, err = client.Get(ctx, path)
		case "POST":
			resp, err = client.Post(ctx, path, body)
		default:
			err = fmt.Errorf("unsupported HTTP method: %s", method)
		}

		if err != nil {
			resource.SetError(err)
			return
		}

		if !resp.IsSuccess() {
			resource.SetError(resp.Error())
			return
		}

		var data []T
		if err := resp.JSON(&data); err != nil {
			resource.SetError(err)
			return
		}

		resource.SetData(data)
	}()

	return resource
}

// Get is a helper for GET requests
func Get[T any](ctx context.Context, client *Client, path string) *Resource[T] {
	return Fetch[T](ctx, client, "GET", path, nil)
}

// Post is a helper for POST requests
func Post[T any](ctx context.Context, client *Client, path string, body interface{}) *Resource[T] {
	return Fetch[T](ctx, client, "POST", path, body)
}

// Put is a helper for PUT requests
func Put[T any](ctx context.Context, client *Client, path string, body interface{}) *Resource[T] {
	return Fetch[T](ctx, client, "PUT", path, body)
}

// Delete is a helper for DELETE requests
func Delete[T any](ctx context.Context, client *Client, path string) *Resource[T] {
	return Fetch[T](ctx, client, "DELETE", path, nil)
}
