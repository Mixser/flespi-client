package flespi_webhook

import "github.com/mixser/flespi-client/internal/flespiapi"

// WebhookClient provides receiver-based methods for managing Flespi webhooks.
// Access it via Client.Webhooks after creating a flespi.Client.
type WebhookClient struct {
	c flespiapi.Doer
}

// NewWebhookClient creates a WebhookClient wrapping the given flespiapi.Doer.
func NewWebhookClient(c flespiapi.Doer) *WebhookClient {
	return &WebhookClient{c: c}
}

func (wc *WebhookClient) NewSingle(name string, options ...CreateSingleWebhookOption) (*SingleWebhook, error) {
	return NewSingleWebhook(wc.c, name, options...)
}

func (wc *WebhookClient) NewChained(name string, options ...CreateChainedWebhookOption) (*ChainedWebhook, error) {
	return NewChainedWebhook(wc.c, name, options...)
}

func (wc *WebhookClient) List() ([]Webhook, error) {
	return ListWebhooks(wc.c)
}

func (wc *WebhookClient) Get(webhookId int64) (Webhook, error) {
	return GetWebhook(wc.c, webhookId)
}

func (wc *WebhookClient) Update(webhook Webhook) (Webhook, error) {
	return UpdateWebhook(wc.c, webhook)
}

func (wc *WebhookClient) Delete(webhook Webhook) error {
	return DeleteWebhook(wc.c, webhook)
}

func (wc *WebhookClient) DeleteById(webhookId int64) error {
	return DeleteWebhookById(wc.c, webhookId)
}
