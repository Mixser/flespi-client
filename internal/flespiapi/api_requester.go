// Package flespiapi defines shared interfaces used across the flespi-client module.
package flespiapi

import "context"

// Doer is the interface resource packages use to make API requests.
// *flespi.Client implements this interface.
type Doer interface {
	RequestAPI(method, endpoint string, payload, response interface{}) error
	RequestAPIWithContext(ctx context.Context, method, endpoint string, payload, response interface{}) error
}
