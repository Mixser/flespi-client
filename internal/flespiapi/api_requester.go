// Package flespiapi defines shared interfaces used across the flespi-client module.
package flespiapi

import "context"

// APIRequester is the interface resource packages use to make API requests.
// *flespi.Client implements this interface.
type APIRequester interface {
	RequestAPI(method, endpoint string, payload, response interface{}) error
	RequestAPIWithContext(ctx context.Context, method, endpoint string, payload, response interface{}) error
	RequestAPIWithHeaders(method, endpoint string, headers map[string]string, payload, response interface{}) error
	RequestAPIWithContextAndHeaders(ctx context.Context, method, endpoint string, headers map[string]string, payload, response interface{}) error
}
