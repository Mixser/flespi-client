package flespi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CDN struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	Blocked bool   `json:"blocked,omitempty"`
	Size    int64  `json:"size,omitempty"`
}

type cdnsListResponse struct {
	CDNS []CDN `json:"result"`
}

type newCDNResposne struct {
	CDNS []CDN `json:"result"`
}

func (c *Client) GetCDNS() ([]CDN, error) {
	req, err := http.NewRequest("GET", "https://flespi.io/storage/cdns/all", nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	cdnsListResonse := cdnsListResponse{}
	err = json.Unmarshal(body, &cdnsListResonse)

	if err != nil {
		return nil, err
	}

	return cdnsListResonse.CDNS, nil
}

func (c *Client) GetCDN(id int64) (*CDN, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://flespi.io/storage/cdns/%d", id), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	reposne := cdnsListResponse{}
	err = json.Unmarshal(body, &reposne)

	if err != nil {
		return nil, err
	}

	return &reposne.CDNS[0], nil
}

func (c *Client) NewCDN(cdn CDN) (*CDN, error) {
	httpReqBody, err := json.Marshal([]CDN{cdn})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/storage/cdns", c.Host), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	cdnResponse := newCDNResposne{}

	err = json.Unmarshal(resp, &cdnResponse)

	if err != nil {
		return nil, err
	}

	return &cdnResponse.CDNS[0], nil
}

func (c *Client) UpdateCDN(cdnId int64, cdn CDN) (*CDN, error) {
	cdn.Id = 0
	defer func() {
		cdn.Id = cdnId
	}()

	httpReqBody, err := json.Marshal(cdn)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/storage/cdns/%d", c.Host, cdnId), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	cdnResponse := newCDNResposne{}

	err = json.Unmarshal(resp, &cdnResponse)

	if err != nil {
		return nil, err
	}

	return &cdn, nil
}

func (c *Client) DeleteCDN(cdnId int64) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/storage/cdns/%d", c.Host, cdnId), nil)

	if err != nil {
		return nil
	}

	_, err = c.doRequest(req)

	if err != nil {
		return err
	}

	return nil
}
