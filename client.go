// Package flespi provides a Go client for the Flespi telematic platform API.
//
// The client supports all major Flespi resources including webhooks, streams,
// channels, devices, calculators, and more. It provides context support for
// request cancellation, structured error handling, and configurable timeouts.
//
// Example usage:
//
//	client, err := flespi.NewClient("https://flespi.io", "your-token")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	stream, err := flespi_stream.NewStream(client, "my-stream", 1)
//	if err != nil {
//	    log.Fatal(err)
//	}
package flespi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a Flespi API client
type Client struct {
	Host       string
	Token      string
	HTTPClient *http.Client
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}

// NewClient creates a new Flespi API client with the specified host and token.
//
// The host parameter should be the base URL of the Flespi API (e.g., "https://flespi.io").
// The token parameter should be a valid Flespi authentication token.
//
// Optional configuration can be provided using ClientOption functions:
//   - WithTimeout(duration): Set a custom HTTP client timeout
//   - WithHTTPClient(client): Use a custom HTTP client
//
// Example:
//
//	client, err := flespi.NewClient("https://flespi.io", "your-token",
//	    flespi.WithTimeout(30 * time.Second),
//	)
func NewClient(host string, token string, options ...ClientOption) (*Client, error) {
	c := &Client{
		Host:       host,
		Token:      token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	for _, opt := range options {
		opt(c)
	}

	return c, nil
}

func (c *Client) doRequest(req *http.Request, method, endpoint string) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("FlespiToken %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Check for successful status codes (2xx)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, parseAPIError(res.StatusCode, method, endpoint, body)
	}

	return body, nil
}

// RequestAPI makes an API request without context (uses context.Background())
func (c *Client) RequestAPI(method string, endpoint string, payload interface{}, response interface{}) error {
	return c.RequestAPIWithContext(context.Background(), method, endpoint, payload, response)
}

// RequestAPIWithContext makes an API request with context support
func (c *Client) RequestAPIWithContext(ctx context.Context, method string, endpoint string, payload interface{}, response interface{}) error {
	var body io.Reader

	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.Host, endpoint), body)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req, method, endpoint)
	if err != nil {
		return err
	}

	if response != nil {
		err = json.Unmarshal(resp, response)
		if err != nil {
			return err
		}
	}

	return nil
}
