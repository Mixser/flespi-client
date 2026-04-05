// Package testhelper provides utilities for testing resource packages
// without importing the root flespi package (which would create a cycle).
package testhelper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TestClient is a minimal Doer implementation for use in tests.
type TestClient struct {
	baseURL string
	token   string
	http    *http.Client
}

// New creates a TestClient pointed at the given base URL.
func New(baseURL string) *TestClient {
	return &TestClient{
		baseURL: baseURL,
		token:   "test-token",
		http:    &http.Client{},
	}
}

func (c *TestClient) RequestAPI(method, endpoint string, payload, response interface{}) error {
	return c.RequestAPIWithContext(context.Background(), method, endpoint, payload, response)
}

func (c *TestClient) RequestAPIWithContext(_ context.Context, method, endpoint string, payload, response interface{}) error {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.baseURL, endpoint), body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("FlespiToken %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, respBody)
	}

	if response != nil {
		return json.Unmarshal(respBody, response)
	}
	return nil
}
