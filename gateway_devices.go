package flespi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Device struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`

	DeviceTypeId int64 `json:"device_type_id"`

	MessagesTTL    int64 `json:"messages_ttl"`
	MessagesRotate int64 `json:"messages_rotate"`

	MediaTTL    int64 `json:"media_ttl"`
	MediaRotate int64 `json:"media_rotate"`

	Metadata map[string]string `json:"metadata"`
}

type createDeviceResponse struct {
	Devices []Device `json:"result"`
}

type deviceListResponse struct {
	Devices []Device `json:"result"`
}

func (c *Client) NewDevice(device Device) (*Device, error) {
	httpReqBody, err := json.Marshal([]Device{device})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gw/devices", c.Host), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	newDeviceResponse := createDeviceResponse{}
	err = json.Unmarshal(resp, &newDeviceResponse)

	if err != nil {
		return nil, err
	}

	return &newDeviceResponse.Devices[0], nil
}

func (c *Client) GetDevice(deviceId int64) (*Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gw/devices/%d", c.Host, deviceId), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	deviceResponse := deviceListResponse{}
	err = json.Unmarshal(resp, &deviceResponse)

	if err != nil {
		return nil, err
	}

	return &deviceResponse.Devices[0], nil
}

func (c *Client) GetDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gw/devices/all", c.Host), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	deviceResponse := deviceListResponse{}
	err = json.Unmarshal(resp, &deviceResponse)

	return deviceResponse.Devices, nil
}

func (c *Client) UpdateDevice(deviceId int64, device Device) (*Device, error) {
	device.Id = 0
	httpReqBody, err := json.Marshal(device)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/gw/devices/%d", c.Host, deviceId), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	deviceResponse := createDeviceResponse{}

	err = json.Unmarshal(resp, &deviceResponse)

	if err != nil {
		return nil, err
	}

	return &deviceResponse.Devices[0], nil
}
