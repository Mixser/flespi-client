package flespi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetryConfig_ShouldRetry(t *testing.T) {
	config := DefaultRetryConfig()

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "429 should retry",
			err:      &APIError{StatusCode: 429},
			expected: true,
		},
		{
			name:     "500 should retry",
			err:      &APIError{StatusCode: 500},
			expected: true,
		},
		{
			name:     "502 should retry",
			err:      &APIError{StatusCode: 502},
			expected: true,
		},
		{
			name:     "503 should retry",
			err:      &APIError{StatusCode: 503},
			expected: true,
		},
		{
			name:     "404 should not retry",
			err:      &APIError{StatusCode: 404},
			expected: false,
		},
		{
			name:     "400 should not retry",
			err:      &APIError{StatusCode: 400},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := config.shouldRetry(tt.err)
			if result != tt.expected {
				t.Errorf("shouldRetry() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRetryConfig_CalculateBackoff(t *testing.T) {
	config := DefaultRetryConfig()

	tests := []struct {
		attempt         int
		expectedBackoff time.Duration
	}{
		{0, 1 * time.Second},
		{1, 2 * time.Second},
		{2, 4 * time.Second},
		{3, 8 * time.Second},
		{4, 16 * time.Second},
		{5, 30 * time.Second},  // capped at MaxBackoff
		{10, 30 * time.Second}, // still capped
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.attempt)), func(t *testing.T) {
			backoff := config.calculateBackoff(tt.attempt)
			if backoff != tt.expectedBackoff {
				t.Errorf("calculateBackoff(%d) = %v, want %v", tt.attempt, backoff, tt.expectedBackoff)
			}
		})
	}
}

func TestClient_RetryOnTransientErrors(t *testing.T) {
	attempts := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentAttempt := atomic.AddInt32(&attempts, 1)

		// Fail first two attempts with 503, succeed on third
		if currentAttempt < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"errors": [{"reason": "service unavailable"}]}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": []}`))
	}))
	defer server.Close()

	retryConfig := &RetryConfig{
		MaxRetries:        3,
		InitialBackoff:    10 * time.Millisecond,
		MaxBackoff:        100 * time.Millisecond,
		BackoffMultiplier: 2.0,
		RetryableStatusCodes: map[int]bool{
			http.StatusServiceUnavailable: true,
		},
	}

	client, _ := NewClient(server.URL, "test-token", WithRetryConfig(retryConfig))

	type response struct {
		Result []interface{} `json:"result"`
	}

	var resp response
	err := client.RequestAPI("GET", "test/endpoint", nil, &resp)

	if err != nil {
		t.Errorf("Expected request to succeed after retries, got error: %v", err)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestClient_NoRetryOnNonRetryableError(t *testing.T) {
	attempts := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"errors": [{"reason": "not found"}]}`))
	}))
	defer server.Close()

	retryConfig := DefaultRetryConfig()
	client, _ := NewClient(server.URL, "test-token", WithRetryConfig(retryConfig))

	err := client.RequestAPI("GET", "test/endpoint", nil, nil)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if attempts != 1 {
		t.Errorf("Expected 1 attempt (no retries), got %d", attempts)
	}
}

func TestClient_RetryWithContextCancellation(t *testing.T) {
	attempts := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"errors": [{"reason": "service unavailable"}]}`))
	}))
	defer server.Close()

	retryConfig := &RetryConfig{
		MaxRetries:        5,
		InitialBackoff:    100 * time.Millisecond,
		MaxBackoff:        1 * time.Second,
		BackoffMultiplier: 2.0,
		RetryableStatusCodes: map[int]bool{
			http.StatusServiceUnavailable: true,
		},
	}

	client, _ := NewClient(server.URL, "test-token", WithRetryConfig(retryConfig))

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	err := client.RequestAPIWithContext(ctx, "GET", "test/endpoint", nil, nil)

	if err == nil {
		t.Errorf("Expected error due to context cancellation, got nil")
	}

	// Should have attempted once, then cancelled during backoff
	if attempts < 1 || attempts > 2 {
		t.Errorf("Expected 1-2 attempts before cancellation, got %d", attempts)
	}
}
