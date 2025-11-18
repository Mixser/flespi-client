package main

import (
	"context"
	"fmt"
	"log"
	"time"

	flespi "github.com/mixser/flespi-client"
	webhook "github.com/mixser/flespi-client/resources/platform/webhook"
)

func main() {
	// Create a new Flespi client
	client, err := flespi.NewClient("https://flespi.io", "your-flespi-token")
	if err != nil {
		log.Fatal(err)
	}

	// Example 1: Handle not found errors
	hook, err := webhook.GetWebhook(client, 999999)
	if err != nil {
		if flespi.IsNotFoundError(err) {
			fmt.Println("Webhook not found - this is expected")
		} else {
			log.Printf("Unexpected error: %v\n", err)
		}
	} else {
		fmt.Printf("Found webhook: %s\n", hook.GetId())
	}

	// Example 2: Handle detailed API errors
	_, err = webhook.NewSingleWebhook(client, "", webhook.SWWithConfiguration(webhook.CustomServerConfiguration{
		Type:    "custom-server",
		Uri:     "invalid-uri",
		Method:  "POST",
		Body:    "{}",
		Headers: []webhook.Header{},
	}))
	if err != nil {
		if apiErr, ok := err.(*flespi.APIError); ok {
			fmt.Printf("API Error: Status %d\n", apiErr.StatusCode)
			fmt.Printf("Method: %s %s\n", apiErr.Method, apiErr.Endpoint)
			if len(apiErr.Errors) > 0 {
				fmt.Printf("Reason: %s\n", apiErr.Errors[0].Reason)
			}
		}
	}

	// Example 3: Use context for timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var response interface{}
	err = client.RequestAPIWithContext(ctx, "GET", "platform/webhooks/all", nil, &response)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else if flespi.IsUnauthorizedError(err) {
			fmt.Println("Invalid authentication token")
		} else if flespi.IsRateLimitError(err) {
			fmt.Println("Rate limit exceeded - please wait before retrying")
		} else {
			log.Printf("Request failed: %v\n", err)
		}
	}

	// Example 4: Handle context cancellation
	ctx2, cancel2 := context.WithCancel(context.Background())

	// Cancel immediately for demonstration
	cancel2()

	err = client.RequestAPIWithContext(ctx2, "GET", "platform/webhooks/all", nil, &response)
	if err != nil {
		if ctx2.Err() == context.Canceled {
			fmt.Println("Request was cancelled")
		}
	}
}
