package middleware

import (
	"net/http"

	"github.com/base-go/GoFlow/pkg/api"
)

// BearerAuth adds a Bearer token to the Authorization header
func BearerAuth(token string) api.Middleware {
	return func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			req.Header.Set("Authorization", "Bearer "+token)
			return next(req)
		}
	}
}

// BasicAuth adds basic authentication to requests
func BasicAuth(username, password string) api.Middleware {
	return func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			req.SetBasicAuth(username, password)
			return next(req)
		}
	}
}

// APIKey adds an API key header to requests
func APIKey(headerName, apiKey string) api.Middleware {
	return func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			req.Header.Set(headerName, apiKey)
			return next(req)
		}
	}
}
