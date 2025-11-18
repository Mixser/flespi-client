package flespi

import (
	"fmt"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// ValidateID checks if an ID is valid (greater than 0)
func ValidateID(id int64, fieldName string) error {
	if id <= 0 {
		return &ValidationError{
			Field:   fieldName,
			Message: "must be greater than 0",
		}
	}
	return nil
}

// ValidateRequired checks if a string field is not empty
func ValidateRequired(value string, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return &ValidationError{
			Field:   fieldName,
			Message: "is required and cannot be empty",
		}
	}
	return nil
}

// ValidateURL checks if a URL string is not empty and has a valid format
func ValidateURL(url string, fieldName string) error {
	if err := ValidateRequired(url, fieldName); err != nil {
		return err
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return &ValidationError{
			Field:   fieldName,
			Message: "must start with http:// or https://",
		}
	}

	return nil
}

// ValidateToken checks if a token is valid
func ValidateToken(token string) error {
	return ValidateRequired(token, "token")
}

// ValidateHost checks if a host URL is valid
func ValidateHost(host string) error {
	return ValidateURL(host, "host")
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
