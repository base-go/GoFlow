package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response wraps an HTTP response with helper methods
type Response struct {
	*http.Response
	body []byte
}

// NewResponse creates a new Response from an http.Response
func NewResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	return &Response{
		Response: resp,
		body:     body,
	}, nil
}

// Body returns the raw response body
func (r *Response) Body() []byte {
	return r.body
}

// JSON decodes the response body as JSON into the provided interface
func (r *Response) JSON(v interface{}) error {
	if err := json.Unmarshal(r.body, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}
	return nil
}

// String returns the response body as a string
func (r *Response) String() string {
	return string(r.body)
}

// IsSuccess returns true if the status code is 2xx
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// IsError returns true if the status code is 4xx or 5xx
func (r *Response) IsError() bool {
	return r.StatusCode >= 400
}

// Error returns an error representation of the response
func (r *Response) Error() error {
	if r.IsSuccess() {
		return nil
	}
	return fmt.Errorf("HTTP %d: %s", r.StatusCode, string(r.body))
}
