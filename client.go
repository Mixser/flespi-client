package flespi

import (
	"bytes"
	"encoding/json"
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

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

func (c *Client) RequestAPI(method string, endpoint string, payload interface{}, response interface{}) error {
	var body io.Reader

	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.Host, endpoint), body)

	if err != nil {
		return nil
	}

	resp, err := c.doRequest(req)

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
