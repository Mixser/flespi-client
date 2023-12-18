package flespi

import (
	"encoding/json"
	"net/http"
)

type CDN struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Blocked bool   `json:"blocked"`
	Size    int    `json:"size"`
}

type cdnsListResonse struct {
	CDNS []CDN `json:"result"`
}

func (c *Client) GetCDNS() ([]CDN, error) {
	req, err := http.NewRequest("GET", "https://flespi.io/storage/cdns/all", nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	cdnsListResonse := cdnsListResonse{}
	err = json.Unmarshal(body, &cdnsListResonse)

	if err != nil {
		return nil, err
	}

	return cdnsListResonse.CDNS, nil
}

//
//func (c* Client) GetCDN(id int64) (*CDN, error) {
//	req, err := http.NewRequest("GET", fmt.Sprintf("https://flespi.io/storage/cdns/%d", id), nil)
//
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := c.doRequest(req, nil)
//
//	if err != nil {
//		return nil, err
//	}
//
//
//}
