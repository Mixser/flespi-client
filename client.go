package flespi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Host       string
	Token      string
	HTTPClient *http.Client
}

func NewClient(host string, token string) (*Client, error) {
	c := Client{
		Host:       host,
		Token:      token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", fmt.Sprintf("FlespiToken %s", token))

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}
