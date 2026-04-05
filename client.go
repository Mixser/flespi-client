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
//	device, err := client.Devices.Create("my-device", true, 5)
//	streams, err := client.Streams.List()
//	err = client.Webhooks.DeleteById(42)
package flespi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mixser/flespi-client/internal/flespiapi"
	flespi_calculator "github.com/mixser/flespi-client/resources/gateway/calculator"
	flespi_channel "github.com/mixser/flespi-client/resources/gateway/channel"
	flespi_device "github.com/mixser/flespi-client/resources/gateway/device"
	flespi_geofence "github.com/mixser/flespi-client/resources/gateway/geofence"
	flespi_stream "github.com/mixser/flespi-client/resources/gateway/stream"
	flespi_token "github.com/mixser/flespi-client/resources/gateway/token"
	flespi_limit "github.com/mixser/flespi-client/resources/platform/limit"
	flespi_subaccount "github.com/mixser/flespi-client/resources/platform/subaccount"
	flespi_webhook "github.com/mixser/flespi-client/resources/platform/webhook"
	flespi_cdn "github.com/mixser/flespi-client/resources/storage/cdn"
)

// Doer is the interface resource packages use to make API requests.
// *Client implements this interface.
type Doer = flespiapi.Doer

// compile-time check that *Client implements Doer
var _ Doer = (*Client)(nil)

// Client represents a Flespi API client
type Client struct {
	Host        string
	Token       string
	HTTPClient  *http.Client
	RetryConfig *RetryConfig
	Logger      Logger

	// Gateway sub-clients
	Devices     *flespi_device.DeviceClient
	Streams     *flespi_stream.StreamClient
	Channels    *flespi_channel.ChannelClient
	Tokens      *flespi_token.TokenClient
	Calculators *flespi_calculator.CalculatorClient
	Geofences   *flespi_geofence.GeofenceClient

	// Platform sub-clients
	Webhooks    *flespi_webhook.WebhookClient
	Subaccounts *flespi_subaccount.SubaccountClient
	Limits      *flespi_limit.LimitClient

	// Storage sub-clients
	CDNs *flespi_cdn.CDNClient
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

	c.Devices = flespi_device.NewDeviceClient(c)
	c.Streams = flespi_stream.NewStreamClient(c)
	c.Channels = flespi_channel.NewChannelClient(c)
	c.Tokens = flespi_token.NewTokenClient(c)
	c.Calculators = flespi_calculator.NewCalculatorClient(c)
	c.Geofences = flespi_geofence.NewGeofenceClient(c)
	c.Webhooks = flespi_webhook.NewWebhookClient(c)
	c.Subaccounts = flespi_subaccount.NewSubaccountClient(c)
	c.Limits = flespi_limit.NewLimitClient(c)
	c.CDNs = flespi_cdn.NewCDNClient(c)

	return c, nil
}

func (c *Client) doRequest(req *http.Request, method, endpoint string) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("FlespiToken %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req) //nolint:gosec // G704: URL is constructed from c.Host, set at client init — not user input
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
	c.logRequest(method, endpoint, payload)

	var body io.Reader

	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			c.logError("Failed to marshal payload: %v", err)
			return err
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.Host, endpoint), body)
	if err != nil {
		c.logError("Failed to create request: %v", err)
		return err
	}

	var resp []byte
	if c.RetryConfig != nil && c.RetryConfig.MaxRetries > 0 {
		resp, err = c.doRequestWithRetry(ctx, req, method, endpoint)
	} else {
		resp, err = c.doRequest(req, method, endpoint)
	}

	if err != nil {
		c.logResponse(method, endpoint, 0, err)
		return err
	}

	c.logResponse(method, endpoint, 200, nil)

	if response != nil {
		err = json.Unmarshal(resp, response)
		if err != nil {
			c.logError("Failed to unmarshal response: %v", err)
			return err
		}
	}

	return nil
}
