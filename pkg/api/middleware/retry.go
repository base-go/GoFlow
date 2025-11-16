package middleware

import (
	"net/http"
	"time"

	"github.com/base-go/GoFlow/pkg/api"
)

// RetryConfig configures retry behavior
type RetryConfig struct {
	MaxRetries int
	RetryDelay time.Duration
	RetryOn    func(*http.Response, error) bool
}

// DefaultRetryOn returns true for retryable errors
func DefaultRetryOn(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	// Retry on 5xx server errors and 429 rate limit
	return resp.StatusCode >= 500 || resp.StatusCode == 429
}

// Retry adds retry logic to requests
func Retry(config RetryConfig) api.Middleware {
	if config.MaxRetries <= 0 {
		config.MaxRetries = 3
	}
	if config.RetryDelay == 0 {
		config.RetryDelay = 1 * time.Second
	}
	if config.RetryOn == nil {
		config.RetryOn = DefaultRetryOn
	}

	return func(next api.RequestHandler) api.RequestHandler {
		return func(req *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error

			for attempt := 0; attempt <= config.MaxRetries; attempt++ {
				if attempt > 0 {
					// Exponential backoff
					delay := config.RetryDelay * time.Duration(1<<uint(attempt-1))
					time.Sleep(delay)
				}

				resp, err = next(req)

				// Check if we should retry
				if !config.RetryOn(resp, err) {
					break
				}

				// Don't retry on last attempt
				if attempt == config.MaxRetries {
					break
				}
			}

			return resp, err
		}
	}
}
