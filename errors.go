package flespi

import (
	"encoding/json"
	"fmt"
)

// APIError represents an error returned by the Flespi API
type APIError struct {
	StatusCode int
	Method     string
	Endpoint   string
	Message    string
	RawBody    []byte
	Errors     []ErrorDetail
}

// ErrorDetail represents a single error from the Flespi API response
type ErrorDetail struct {
	Reason string `json:"reason"`
	ID     int64  `json:"id,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("flespi API error: %s %s (status %d): %s",
			e.Method, e.Endpoint, e.StatusCode, e.Errors[0].Reason)
	}
	return fmt.Sprintf("flespi API error: %s %s (status %d): %s",
		e.Method, e.Endpoint, e.StatusCode, e.Message)
}

// errorResponse represents the error structure returned by Flespi API
type errorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

// parseAPIError attempts to parse the error response from Flespi API
func parseAPIError(statusCode int, method, endpoint string, body []byte) error {
	apiErr := &APIError{
		StatusCode: statusCode,
		Method:     method,
		Endpoint:   endpoint,
		RawBody:    body,
	}

	var errResp errorResponse
	if err := json.Unmarshal(body, &errResp); err == nil && len(errResp.Errors) > 0 {
		apiErr.Errors = errResp.Errors
	} else {
		apiErr.Message = string(body)
	}

	return apiErr
}

// IsNotFoundError checks if the error is a 404 Not Found error
func IsNotFoundError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsUnauthorizedError checks if the error is a 401 Unauthorized error
func IsUnauthorizedError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsRateLimitError checks if the error is a 429 Too Many Requests error
func IsRateLimitError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 429
	}
	return false
}
