package flespi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		token   string
		options []ClientOption
		wantErr bool
	}{
		{
			name:    "basic client",
			host:    "https://flespi.io",
			token:   "test-token",
			wantErr: false,
		},
		{
			name:  "client with custom timeout",
			host:  "https://flespi.io",
			token: "test-token",
			options: []ClientOption{
				WithTimeout(30 * time.Second),
			},
			wantErr: false,
		},
		{
			name:  "client with custom HTTP client",
			host:  "https://flespi.io",
			token: "test-token",
			options: []ClientOption{
				WithHTTPClient(&http.Client{Timeout: 5 * time.Second}),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.host, tt.token, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if client.Host != tt.host {
					t.Errorf("NewClient() Host = %v, want %v", client.Host, tt.host)
				}
				if client.Token != tt.token {
					t.Errorf("NewClient() Token = %v, want %v", client.Token, tt.token)
				}
			}
		})
	}
}

func TestClient_RequestAPI_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify authorization header
		auth := r.Header.Get("Authorization")
		if auth != "FlespiToken test-token" {
			t.Errorf("Expected Authorization header 'FlespiToken test-token', got '%s'", auth)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": [{"id": 123, "name": "test"}]}`))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, "test-token")

	type response struct {
		Result []map[string]interface{} `json:"result"`
	}

	var resp response
	err := client.RequestAPI("GET", "test/endpoint", nil, &resp)
	if err != nil {
		t.Errorf("RequestAPI() error = %v", err)
	}

	if len(resp.Result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(resp.Result))
	}
}

func TestClient_RequestAPI_ErrorHandling(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedErrMsg string
	}{
		{
			name:           "404 not found",
			statusCode:     http.StatusNotFound,
			responseBody:   `{"errors": [{"reason": "resource not found"}]}`,
			expectedErrMsg: "resource not found",
		},
		{
			name:           "401 unauthorized",
			statusCode:     http.StatusUnauthorized,
			responseBody:   `{"errors": [{"reason": "invalid token"}]}`,
			expectedErrMsg: "invalid token",
		},
		{
			name:           "429 rate limit",
			statusCode:     http.StatusTooManyRequests,
			responseBody:   `{"errors": [{"reason": "rate limit exceeded"}]}`,
			expectedErrMsg: "rate limit exceeded",
		},
		{
			name:           "500 server error",
			statusCode:     http.StatusInternalServerError,
			responseBody:   `{"errors": [{"reason": "internal server error"}]}`,
			expectedErrMsg: "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client, _ := NewClient(server.URL, "test-token")

			err := client.RequestAPI("GET", "test/endpoint", nil, nil)
			if err == nil {
				t.Errorf("Expected error, got nil")
				return
			}

			apiErr, ok := err.(*APIError)
			if !ok {
				t.Errorf("Expected *APIError, got %T", err)
				return
			}

			if apiErr.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, apiErr.StatusCode)
			}

			if len(apiErr.Errors) > 0 && apiErr.Errors[0].Reason != tt.expectedErrMsg {
				t.Errorf("Expected error message '%s', got '%s'", tt.expectedErrMsg, apiErr.Errors[0].Reason)
			}
		})
	}
}

func TestClient_RequestAPIWithContext(t *testing.T) {
	t.Run("context cancellation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		client, _ := NewClient(server.URL, "test-token")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		err := client.RequestAPIWithContext(ctx, "GET", "test/endpoint", nil, nil)
		if err == nil {
			t.Errorf("Expected timeout error, got nil")
		}
	})
}

func TestErrorHelpers(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		checkFunc  func(error) bool
		expected   bool
	}{
		{
			name:       "IsNotFoundError - true",
			statusCode: 404,
			checkFunc:  IsNotFoundError,
			expected:   true,
		},
		{
			name:       "IsNotFoundError - false",
			statusCode: 500,
			checkFunc:  IsNotFoundError,
			expected:   false,
		},
		{
			name:       "IsUnauthorizedError - true",
			statusCode: 401,
			checkFunc:  IsUnauthorizedError,
			expected:   true,
		},
		{
			name:       "IsRateLimitError - true",
			statusCode: 429,
			checkFunc:  IsRateLimitError,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &APIError{StatusCode: tt.statusCode}
			result := tt.checkFunc(err)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestClient_AcceptsMultipleSuccessStatusCodes(t *testing.T) {
	statusCodes := []int{200, 201, 202, 204}

	for _, code := range statusCodes {
		t.Run(http.StatusText(code), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(code)
				if code != 204 {
					w.Write([]byte(`{"result": []}`))
				}
			}))
			defer server.Close()

			client, _ := NewClient(server.URL, "test-token")

			type response struct {
				Result []interface{} `json:"result"`
			}

			// For 204 No Content, we should not pass a response object
			if code == 204 {
				err := client.RequestAPI("GET", "test/endpoint", nil, nil)
				if err != nil {
					t.Errorf("Expected no error for status %d, got %v", code, err)
				}
			} else {
				var resp response
				err := client.RequestAPI("GET", "test/endpoint", nil, &resp)
				if err != nil {
					t.Errorf("Expected no error for status %d, got %v", code, err)
				}
			}
		})
	}
}
