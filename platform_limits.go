package flespi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Limit struct {
	Id               int64  `json:"id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	BlockingDuration int    `json:"blocking_duration"`
	ApiCall          int64  `json:"api_calls"`
	ApiTraffic       int64  `json:"api_traffic"`

	ChannelsCount      int64 `json:"channels_count"`
	ChannelMessages    int64 `json:"channel_messages"`
	ChannelStorage     int64 `json:"channel_storage"`
	ChannelTraffic     int64 `json:"channel_traffic"`
	ChannelConnections int64 `json:"channel_connections"`

	ContainersCount  int64 `json:"containers_count"`
	ContainerStorage int64 `json:"container_storage"`

	CdnsCount  int64 `json:"cdns_count"`
	CdnStorage int64 `json:"cdn_storage"`
	CdnTraffic int64 `json:"cdn_traffic"`

	DevicesCount       int64 `json:"devices_count"`
	DeviceStorage      int64 `json:"device_storage"`
	DeviceMediaTraffic int64 `json:"device_media_traffic"`
	DeviceMediaStorage int64 `json:"device_media_storage"`

	StreamsCount  int64 `json:"streams_count"`
	StreamStorage int64 `json:"stream_storage"`
	StreamTraffic int64 `json:"stream_traffic"`

	ModemsCount int64 `json:"modems_count"`

	MqttSessions        int64 `json:"mqtt_sessions"`
	MqttMessages        int64 `json:"mqtt_messages"`
	MqttSessionStorage  int64 `json:"mqtt_session_storage"`
	MqttRetainedStorage int64 `json:"mqtt_retained_storage"`
	MqttSubscriptions   int64 `json:"mqtt_subscriptions"`

	SmsCount int64 `json:"sms_count"`

	TokensCount int64 `json:"tokens_count"`

	SubaccountsCount int64 `json:"subaccounts_count"`

	LimitsCount int64 `json:"limits_count"`

	RealmsCount int64 `json:"realms_count"`

	CalcsCount   int64 `json:"calcs_count"`
	CalcsStorage int64 `json:"calcs_storage"`

	PluginsCount           int64 `json:"plugins_count"`
	PluginTraffic          int64 `json:"plugin_traffic"`
	PluginBufferedMessages int64 `json:"plugin_buffered_messages"`

	GroupsCount int64 `json:"groups_count"`

	WebhooksCount  int64 `json:"webhooks_count"`
	WebhookStorage int64 `json:"webhook_storage"`
	WebhookTraffic int64 `json:"webhook_traffic"`

	GrantsCount int64 `json:"grants_count"`

	IdentityProvidersCount int64 `json:"identity_providers_count"`

	Metadata map[string]string `json:"metadata"`
}

type limitsListResponse struct {
	Limits []Limit `json:"result"`
}

func (c *Client) GetLimits() ([]Limit, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/platform/limits/all", c.Host), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	limitsListResonse := limitsListResponse{}
	err = json.Unmarshal(body, &limitsListResonse)

	if err != nil {
		return nil, err
	}

	return limitsListResonse.Limits, nil

}

func (c *Client) GetLimit(limitId int64) (*Limit, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/platform/limits/%d", c.Host, limitId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	limitsListResonse := limitsListResponse{}
	err = json.Unmarshal(body, &limitsListResonse)

	if err != nil {
		return nil, err
	}

	return &limitsListResonse.Limits[0], nil
}

func (c *Client) NewLimit(limit Limit) (*Limit, error) {
	httpReqBody, err := json.Marshal([]Limit{limit})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/platform/limits", c.Host), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	limitResponse := newLimitResponse{}

	err = json.Unmarshal(resp, &limitResponse)

	if err != nil {
		return nil, err
	}

	return &limitResponse.Limit[0], nil
}

func (c *Client) UpdateLimit(limitId int64, limit Limit) (*Limit, error) {
	limit.Id = 0

	httpReqBody, err := json.Marshal(limit)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/platform/limits/%d", c.Host, limitId), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	limitResponse := newLimitResponse{}

	err = json.Unmarshal(resp, &limitResponse)

	if err != nil {
		return nil, err
	}

	limit.Id = limitId

	return &limit, nil
}

type newLimitResponse struct {
	Limit []Limit `json:"result"`
}


func (c *Client) DeleteLimit(limitId int64) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/platform/limits/%d", c.Host, limitId), nil)

	if err != nil {
		return nil
	}

	_, err = c.doRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}