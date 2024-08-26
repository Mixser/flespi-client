package flespi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Subaccount struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	LimitId int64  `json:"limit_id"`

	Metadata map[string]string `json:"metadata"`
}

type subaccountsListResponse struct {
	Subaccounts []Subaccount `json:"result"`
}

func (c *Client) NewSubaccount(subaccount Subaccount) (*Subaccount, error) {
	httpReqBody, err := json.Marshal([]Subaccount{subaccount})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/platform/subaccounts", c.Host), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	subaccountsResponse := subaccountsListResponse{}

	err = json.Unmarshal(resp, &subaccountsResponse)

	if err != nil {
		return nil, err
	}

	return &subaccountsResponse.Subaccounts[0], nil
}

func (c *Client) GetSubaccount(subaccountId int64) (*Subaccount, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/platform/subaccounts/%d", c.Host, subaccountId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	response := subaccountsListResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response.Subaccounts[0], nil
}

func (c *Client) UpdateSubaccount(subaccountId int64, subaccount Subaccount) (*Subaccount, error) {
	subaccount.Id = 0 // Fill with zero to ommit this field

	httpReqBody, err := json.Marshal(subaccount)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/platform/subaccounts/%d", c.Host, subaccountId), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	subaccountResponse := limitsListResponse{}

	err = json.Unmarshal(resp, &subaccountResponse)

	if err != nil {
		return nil, err
	}

	subaccount.Id = subaccountId
	return &subaccount, nil
}

func (c *Client) DeleteSubaccount(subaccountId int64) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/platform/subaccounts/%d", c.Host, subaccountId), nil)

	if err != nil {
		return nil
	}

	_, err = c.doRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}
