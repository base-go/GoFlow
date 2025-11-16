package api_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/base-go/GoFlow/pkg/api"
	"github.com/base-go/GoFlow/pkg/api/middleware"
)

func TestClientGet(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/test" {
			t.Errorf("Expected path /test, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "success"})
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
	})

	resp, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	if !resp.IsSuccess() {
		t.Errorf("Expected success status, got %d", resp.StatusCode)
	}

	var result map[string]string
	if err := resp.JSON(&result); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if result["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", result["message"])
	}
}

func TestClientPost(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		if body["name"] != "test" {
			t.Errorf("Expected name 'test', got '%s'", body["name"])
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": "123", "name": body["name"]})
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
	})

	resp, err := client.Post(context.Background(), "/users", map[string]string{"name": "test"})
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}

func TestClientWithHeaders(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") != "custom-value" {
			t.Errorf("Expected custom header, got '%s'", r.Header.Get("X-Custom-Header"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
		Headers: map[string]string{
			"X-Custom-Header": "custom-value",
		},
	})

	_, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
}

func TestClientTimeout(t *testing.T) {
	// Create slow server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
		Timeout: 50 * time.Millisecond,
	})

	_, err := client.Get(context.Background(), "/test")
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestMiddlewareAuth(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-token" {
			t.Errorf("Expected Bearer token, got '%s'", auth)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
	}).Use(middleware.BearerAuth("test-token"))

	_, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
}

func TestMiddlewareChain(t *testing.T) {
	callOrder := []string{}

	middleware1 := func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			callOrder = append(callOrder, "middleware1-before")
			resp, err := next(req)
			callOrder = append(callOrder, "middleware1-after")
			return resp, err
		}
	}

	middleware2 := func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			callOrder = append(callOrder, "middleware2-before")
			resp, err := next(req)
			callOrder = append(callOrder, "middleware2-after")
			return resp, err
		}
	}

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callOrder = append(callOrder, "handler")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.NewClient(api.ClientConfig{
		BaseURL: server.URL,
	}).Use(middleware1).Use(middleware2)

	_, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	expected := []string{
		"middleware1-before",
		"middleware2-before",
		"handler",
		"middleware2-after",
		"middleware1-after",
	}

	if len(callOrder) != len(expected) {
		t.Fatalf("Expected %d calls, got %d", len(expected), len(callOrder))
	}

	for i, call := range callOrder {
		if call != expected[i] {
			t.Errorf("Call %d: expected '%s', got '%s'", i, expected[i], call)
		}
	}
}
