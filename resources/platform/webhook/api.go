package flespi_webhook

import (
	"fmt"
	"github.com/mixser/flespi-client"
)

func NewSignleWebhook(c *flespi.Client, name string, options ...CreateSingleWebhookOption) (*SingleWebhook, error) {
	webhook := SingleWebhook{Name: name}

	for _, opt := range options {
		opt(&webhook)
	}

	result, err := newWebhook(c, &webhook)

	if err != nil {
		return nil, err
	}

	return result.(*SingleWebhook), nil
}

func NewChainedWebhook(c *flespi.Client, name string, options ...CreateChaniedWebhookOption) (*ChainedWebhook, error) {
	webhook := ChainedWebhook{Name: name}

	for _, opt := range options {
		opt(&webhook)
	}

	result, err := newWebhook(c, &webhook)

	if err != nil {
		return nil, err
	}

	return result.(*ChainedWebhook), nil
}


func newWebhook(c *flespi.Client, webhook Webhook) (Webhook, error) {
	response := webhookResponse{}

	err := c.RequestAPI("POST", "platform/webhooks", []Webhook{webhook}, &response)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(response)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func GetWebhook(c *flespi.Client, webhookId int64) (Webhook, error) {
	response := webhookResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("platform/webhooks/%d", webhookId), nil, &response)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(response)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func ListWebhooks(c *flespi.Client) ([]Webhook, error) {
	response := webhookResponse{}

	err := c.RequestAPI("GET", "platform/webhooks/all", nil, &response)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(response)

	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func UpdateWebhook(c *flespi.Client, webhook Webhook) (Webhook, error) {
	response := webhookResponse{}

	err := c.RequestAPI("PUT", fmt.Sprintf("platform/webhooks/%d", webhook.GetId()), webhook, &response)

	if err != nil {
		return nil, err
	}

	webhooks, err := unmarshalWebhookResponse(response)

	if err != nil {
		return nil, err
	}

	return webhooks[0], nil
}

func DeleteWebhook(c *flespi.Client, webhook Webhook) error {
	return DeleteWebhookById(c, webhook.GetId())
}

func DeleteWebhookById(c *flespi.Client, webhookId int64) error {
	err := c.RequestAPI("DELETE", fmt.Sprintf("platform/webhooks/%d", webhookId), nil, nil)

	if err != nil {
		return err
	}
	return nil
}