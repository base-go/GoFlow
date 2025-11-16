package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/base-go/GoFlow/pkg/api"
)

// Logger is a function that logs messages
type Logger func(format string, args ...interface{})

// Logging adds request/response logging
func Logging(logger Logger) api.Middleware {
	if logger == nil {
		logger = func(format string, args ...interface{}) {
			fmt.Printf(format+"\n", args...)
		}
	}

	return func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			start := time.Now()

			logger("[API] → %s %s", req.Method, req.URL.String())

			resp, err := next(req)

			duration := time.Since(start)

			if err != nil {
				logger("[API] ✗ %s %s - Error: %v (took %v)",
					req.Method, req.URL.String(), err, duration)
			} else {
				logger("[API] ← %s %s - %d %s (took %v)",
					req.Method, req.URL.String(), resp.StatusCode, resp.Status, duration)
			}

			return resp, err
		}
	}
}
