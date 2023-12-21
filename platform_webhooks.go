package flespi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ActionBreak = "break"
	ActionSkip  = "skip"
	ActionRetry = "retry"
)

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Validator struct {
	Expression string `json:"expression"`
	Action     string `json:"action"`
}

type Configuration interface {
	isConfigurationInstance()
}

type CustomServerConfiguration struct {
	Type     string     `json:"type"`
	Uri      string     `json:"uri"`
	Method   string     `json:"method"`
	Body     string     `json:"body"`
	CA       *string    `json:"ca,omitempty"`
	Headers  []Header   `json:"headers"`
	Validate *Validator `json:"validate,omitempty"`
}

func (c CustomServerConfiguration) isConfigurationInstance() {
	return
}

type FlespiConfiguration struct {
	Type     string     `json:"type"`
	Uri      string     `json:"uri"`
	Method   string     `json:"method"`
	Body     string     `json:"body"`
	CID      string     `json:"cid"`
	Validate *Validator `json:"validate"`
}

func (c FlespiConfiguration) isConfigurationInstance() {
	return
}

type TriggerFilter struct {
	CID     int64  `json:"cid"`
	Payload string `json:"payload"`
}

type Trigger struct {
	Topic  string         `json:"topic"`
	Filter *TriggerFilter `json:"filter,omitempty"`
}

type Webhook interface {
	isWebhookObject()
}

type SingleWebhook struct {
	Id            int64         `json:"id,omitempty"`
	Name          string        `json:"name"`
	Triggers      []Trigger     `json:"triggers"`
	Configuration Configuration `json:"configuration"`
}

func (sw *SingleWebhook) isWebhookObject() {
	return
}

func (sw *SingleWebhook) UnmarshalJSON(data []byte) error {
	var staticFieldsStruct struct {
		Id       int64     `json:"id,omitempty"`
		Name     string    `json:"name"`
		Triggers []Trigger `json:"triggers"`
	}

	if err := json.Unmarshal(data, &staticFieldsStruct); err != nil {
		return err
	}

	sw.Id = staticFieldsStruct.Id
	sw.Name = staticFieldsStruct.Name
	sw.Triggers = staticFieldsStruct.Triggers

	var raw struct {
		Configuration json.RawMessage `json:"configuration"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	configuration, err := unmarshalConfiguration(raw.Configuration)

	if err != nil {
		return err
	}

	sw.Configuration = configuration

	return nil
}

type ChainedWebhook struct {
	Id            int64           `json:"id,omitempty"`
	Name          string          `json:"name"`
	Triggers      []Trigger       `json:"triggers"`
	Configuration []Configuration `json:"configuration"`
}

func (cw *ChainedWebhook) isWebhookObject() {
	return
}

func (cw *ChainedWebhook) UnmarshalJSON(data []byte) error {
	var staticFieldsStruct struct {
		Id       int64     `json:"id,omitempty"`
		Name     string    `json:"name"`
		Triggers []Trigger `json:"triggers"`
	}

	if err := json.Unmarshal(data, &staticFieldsStruct); err != nil {
		return err
	}

	cw.Id = staticFieldsStruct.Id
	cw.Name = staticFieldsStruct.Name
	cw.Triggers = staticFieldsStruct.Triggers

	var raw struct {
		Configuration []json.RawMessage `json:"configuration"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, rawValue := range raw.Configuration {
		configuration, err := unmarshalConfiguration(rawValue)

		if err != nil {
			return err
		}

		cw.Configuration = append(cw.Configuration, configuration)
	}

	return nil
}

func unmarshalConfiguration(rawValue json.RawMessage) (Configuration, error) {
	var configurationType struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(rawValue, &configurationType); err != nil {
		return nil, err
	}

	var configuration Configuration

	switch configurationType.Type {
	case "custom-server":
		var customServerConfiguration CustomServerConfiguration
		if err := json.Unmarshal(rawValue, &customServerConfiguration); err != nil {
			return nil, err
		}
		configuration = customServerConfiguration
	case "flespi-platform":
		var flespiConfiguration FlespiConfiguration
		if err := json.Unmarshal(rawValue, &flespiConfiguration); err != nil {
			return nil, err
		}
		configuration = flespiConfiguration
	}

	return configuration, nil
}

func (c *Client) NewWebhook(webhook Webhook) (Webhook, error) {
	httpReqBody, err := json.Marshal([]Webhook{webhook})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/platform/webhooks", c.Host), bytes.NewBuffer(httpReqBody))

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(resp)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func (c *Client) GetWebhook(webhookId int64) (Webhook, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/platform/webhooks/%d", c.Host, webhookId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(body)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func (c *Client) UpdateWebhook(webhookId int64, webhook Webhook) (Webhook, error) {
	httpReqBody, err := json.Marshal(webhook)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/platform/webhooks/%d", c.Host, webhookId), bytes.NewBuffer(httpReqBody))

	resp, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(resp)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func (c *Client) DeleteWebhook(webhookId int64) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/platform/webhooks/%d", c.Host, webhookId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}

func unmarshalWebhook(rawValue json.RawMessage) (Webhook, error) {
	var err error = nil
	var singleWebhook SingleWebhook

	if err = json.Unmarshal(rawValue, &singleWebhook); err == nil {
		return &singleWebhook, nil
	}

	var chainedWebhook ChainedWebhook
	if err = json.Unmarshal(rawValue, &chainedWebhook); err == nil {
		return &chainedWebhook, nil
	}

	return nil, err
}

func unmarshalWebhookResponse(data []byte) ([]Webhook, error) {
	var response struct {
		RawValue []json.RawMessage `json:"result"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	var result []Webhook

	for _, rawValue := range response.RawValue {
		webhook, err := unmarshalWebhook(rawValue)

		if err != nil {
			return nil, err
		}

		result = append(result, webhook)
	}

	return result, nil
}
