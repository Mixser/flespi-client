package main

import (
	"fmt"
	"log"

	flespi "github.com/mixser/flespi-client"
	flespi_webhook "github.com/mixser/flespi-client/resources/platform/webhook"
)

func main() {
	// Create a new Flespi client
	client, err := flespi.NewClient("https://flespi.io", "your-flespi-token")
	if err != nil {
		log.Fatal(err)
	}

	// Create a single webhook
	hook, err := client.Webhooks.NewSingle("example-webhook",
		flespi_webhook.SWWithConfiguration(flespi_webhook.CustomServerConfiguration{
			Type:   "custom-server",
			Uri:    "https://example.com/webhook",
			Method: "POST",
			Body:   `{"event": "{{event}}", "timestamp": "{{timestamp}}"}`,
			Headers: []flespi_webhook.Header{
				{Name: "X-Custom-Header", Value: "example-value"},
				{Name: "Content-Type", Value: "application/json"},
			},
		}),
		flespi_webhook.SWWithTrigger(flespi_webhook.Trigger{
			Topic: "platform/messages",
			Filter: &flespi_webhook.TriggerFilter{
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
	chainedHook, err := client.Webhooks.NewChained("chained-webhook",
		flespi_webhook.CWWithConfiguration(flespi_webhook.CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com/first",
			Method:  "POST",
			Body:    `{"step": 1}`,
			Headers: []flespi_webhook.Header{},
		}),
		flespi_webhook.CWWithConfiguration(flespi_webhook.CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com/second",
			Method:  "POST",
			Body:    `{"step": 2}`,
			Headers: []flespi_webhook.Header{},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created chained webhook: %s (ID: %d)\n", chainedHook.Name, chainedHook.Id)

	// List all webhooks
	all, err := client.Webhooks.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total webhooks: %d\n", len(all))

	// Delete webhooks
	err = client.Webhooks.Delete(hook)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Webhooks.Delete(chainedHook)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Webhooks deleted successfully")
}
