package flespi

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"
)

// RetryConfig defines the retry behavior for API requests
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts (default: 3)
	MaxRetries int

	// InitialBackoff is the initial backoff duration (default: 1 second)
	InitialBackoff time.Duration

	// MaxBackoff is the maximum backoff duration (default: 30 seconds)
	MaxBackoff time.Duration

	// BackoffMultiplier is the multiplier for exponential backoff (default: 2.0)
	BackoffMultiplier float64

	// RetryableStatusCodes are HTTP status codes that should trigger a retry
	// Default: 429 (Too Many Requests), 500, 502, 503, 504
	RetryableStatusCodes map[int]bool
}

// DefaultRetryConfig returns a retry configuration with sensible defaults
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:         3,
		InitialBackoff:     1 * time.Second,
		MaxBackoff:         30 * time.Second,
		BackoffMultiplier:  2.0,
		RetryableStatusCodes: map[int]bool{
			http.StatusTooManyRequests:     true, // 429
			http.StatusInternalServerError: true, // 500
			http.StatusBadGateway:          true, // 502
			http.StatusServiceUnavailable:  true, // 503
			http.StatusGatewayTimeout:      true, // 504
		},
	}
}

// WithRetryConfig sets a custom retry configuration
func WithRetryConfig(config *RetryConfig) ClientOption {
	return func(c *Client) {
		c.RetryConfig = config
	}
}

// shouldRetry determines if an error should trigger a retry
func (rc *RetryConfig) shouldRetry(err error) bool {
	if rc == nil || rc.MaxRetries == 0 {
		return false
	}

	if apiErr, ok := err.(*APIError); ok {
		return rc.RetryableStatusCodes[apiErr.StatusCode]
	}

	return false
}

// calculateBackoff calculates the backoff duration for a given attempt
func (rc *RetryConfig) calculateBackoff(attempt int) time.Duration {
	if rc == nil {
		return 0
	}

	backoff := float64(rc.InitialBackoff) * math.Pow(rc.BackoffMultiplier, float64(attempt))
	if backoff > float64(rc.MaxBackoff) {
		backoff = float64(rc.MaxBackoff)
	}

	return time.Duration(backoff)
}

// doRequestWithRetry performs an HTTP request with retry logic
func (c *Client) doRequestWithRetry(ctx context.Context, req *http.Request, method, endpoint string) ([]byte, error) {
	var lastErr error
	maxRetries := 0

	if c.RetryConfig != nil {
		maxRetries = c.RetryConfig.MaxRetries
	}

	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Clone the request for retry attempts
		reqClone := req.Clone(ctx)

		body, err := c.doRequest(reqClone, method, endpoint)
		if err == nil {
			if c.Logger != nil && attempt > 0 {
				c.Logger.Infof("Request succeeded after %d retries: %s %s", attempt, method, endpoint)
			}
			return body, nil
		}

		lastErr = err

		// Check if we should retry
		if attempt < maxRetries && c.RetryConfig != nil && c.RetryConfig.shouldRetry(err) {
			backoff := c.RetryConfig.calculateBackoff(attempt)

			if c.Logger != nil {
				c.Logger.Warnf("Request failed (attempt %d/%d), retrying in %v: %s %s - %v",
					attempt+1, maxRetries+1, backoff, method, endpoint, err)
			}

			// Wait for backoff duration or until context is cancelled
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("request cancelled during retry backoff: %w", ctx.Err())
			case <-time.After(backoff):
				// Continue to next retry
			}
		} else {
			break
		}
	}

	return nil, lastErr
}
