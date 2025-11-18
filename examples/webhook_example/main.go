package main

import (
	"fmt"
	"log"

	flespi "github.com/mixser/flespi-client"
	webhook "github.com/mixser/flespi-client/resources/platform/webhook"
)

func main() {
	// Create a new Flespi client
	client, err := flespi.NewClient("https://flespi.io", "your-flespi-token")
	if err != nil {
		log.Fatal(err)
	}

	// Create a single webhook
	hook, err := webhook.NewSingleWebhook(client, "example-webhook",
		webhook.SWWithConfiguration(webhook.CustomServerConfiguration{
			Type:   "custom-server",
			Uri:    "https://example.com/webhook",
			Method: "POST",
			Body:   `{"event": "{{event}}", "timestamp": "{{timestamp}}"}`,
			Headers: []webhook.Header{
				{Name: "X-Custom-Header", Value: "example-value"},
				{Name: "Content-Type", Value: "application/json"},
			},
		}),
		webhook.SWWithTrigger(webhook.Trigger{
			Topic: "platform/messages",
			Filter: &webhook.TriggerFilter{
				CID:     123,
				Payload: "ident != null",
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created webhook: %s (ID: %d)\n", hook.Name, hook.Id)

	// Create a chained webhook
	chainedHook, err := webhook.NewChainedWebhook(client, "chained-webhook",
		webhook.CWWithConfiguration(webhook.CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com/first",
			Method:  "POST",
			Body:    `{"step": 1}`,
			Headers: []webhook.Header{},
		}),
		webhook.CWWithConfiguration(webhook.CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com/second",
			Method:  "POST",
			Body:    `{"step": 2}`,
			Headers: []webhook.Header{},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created chained webhook: %s (ID: %d)\n", chainedHook.Name, chainedHook.Id)

	// List all webhooks
	webhooks, err := webhook.ListWebhooks(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total webhooks: %d\n", len(webhooks))

	// Delete webhooks
	err = webhook.DeleteWebhook(client, hook)
	if err != nil {
		log.Fatal(err)
	}

	err = webhook.DeleteWebhook(client, chainedHook)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Webhooks deleted successfully")
}
