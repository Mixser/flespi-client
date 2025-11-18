# Flespi Go Client

A Go client library for the [Flespi](https://flespi.io) telematic platform API.

## Features

- Full CRUD operations for Flespi resources
- Context support for request cancellation and timeouts
- Structured error handling with detailed API errors
- Configurable HTTP client and timeouts
- Comprehensive test coverage
- Zero external dependencies

## Installation

```bash
go get github.com/mixser/flespi-client
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    flespi "github.com/mixser/flespi-client"
    "github.com/mixser/flespi-client/resources/gateway/stream"
)

func main() {
    // Create a new client
    client, err := flespi.NewClient("https://flespi.io", "your-flespi-token")
    if err != nil {
        log.Fatal(err)
    }

    // Create a stream
    stream, err := flespi_stream.NewStream(client, "my-stream", 1,
        flespi_stream.WithStatus(true),
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created stream: %s (ID: %d)\n", stream.Name, stream.Id)
}
```

## Usage

### Client Configuration

Create a client with custom configuration:

```go
import (
    "time"

    flespi "github.com/mixser/flespi-client"
)

// Basic client
client, _ := flespi.NewClient("https://flespi.io", "your-token")

// With custom timeout
client, _ := flespi.NewClient("https://flespi.io", "your-token",
    flespi.WithTimeout(30 * time.Second),
)

// With custom HTTP client
httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: customTransport,
}
client, _ := flespi.NewClient("https://flespi.io", "your-token",
    flespi.WithHTTPClient(httpClient),
)
```

### Context Support

Use context for request cancellation and timeouts:

```go
import (
    "context"
    "time"
)

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := client.RequestAPIWithContext(ctx, "GET", "platform/webhooks/all", nil, &response)
```

### Error Handling

The client provides structured error handling:

```go
import flespi "github.com/mixser/flespi-client"

webhook, err := flespi_webhook.GetWebhook(client, 123)
if err != nil {
    // Check for specific error types
    if flespi.IsNotFoundError(err) {
        fmt.Println("Webhook not found")
    } else if flespi.IsUnauthorizedError(err) {
        fmt.Println("Invalid token")
    } else if flespi.IsRateLimitError(err) {
        fmt.Println("Rate limit exceeded")
    } else {
        // Get detailed error information
        if apiErr, ok := err.(*flespi.APIError); ok {
            fmt.Printf("API Error: %d - %s\n", apiErr.StatusCode, apiErr.Error())
        }
    }
}
```

### Working with Webhooks

```go
import (
    flespi "github.com/mixser/flespi-client"
    webhook "github.com/mixser/flespi-client/resources/platform/webhook"
)

// Create a single webhook
hook, err := webhook.NewSingleWebhook(client, "my-webhook",
    webhook.SWWithConfiguration(webhook.CustomServerConfiguration{
        Type:   "custom-server",
        Uri:    "https://my-server.com/webhook",
        Method: "POST",
        Body:   `{"event": "{{event}}"}`,
        Headers: []webhook.Header{
            {Name: "X-Custom-Header", Value: "value"},
        },
    }),
    webhook.SWWithTrigger(webhook.Trigger{
        Topic: "platform/messages",
    }),
)

// List all webhooks
webhooks, err := webhook.ListWebhooks(client)

// Get a specific webhook
hook, err := webhook.GetWebhook(client, 123)

// Update webhook
hook.Name = "updated-name"
updated, err := webhook.UpdateWebhook(client, hook)

// Delete webhook
err = webhook.DeleteWebhookById(client, 123)
```

### Working with Streams

```go
import stream "github.com/mixser/flespi-client/resources/gateway/stream"

// Create a stream
s, err := stream.NewStream(client, "my-stream", 1,
    stream.WithStatus(true),
    stream.WithQueueTTL(86400),
    stream.WithConfigurationItem("key", "value"),
)

// List all streams
streams, err := stream.ListStreams(client)

// Get a specific stream
s, err := stream.GetStream(client, 789)

// Update stream
s.Enabled = false
updated, err := stream.UpdateStream(client, *s)

// Delete stream
err = stream.DeleteStreamById(client, 789)
```

### Working with Channels

```go
import channel "github.com/mixser/flespi-client/resources/gateway/channel"

// Create a channel with protocol name
ch, err := channel.NewChannelWithProtocolName(client, "my-channel", "teltonika",
    channel.WithStatus(true),
    channel.WithMessagesTTL(2592000),
)

// Create a channel with protocol ID
ch, err := channel.NewChannelWithProtocolId(client, "my-channel", 1,
    channel.WithConfigurationItem("port", 5000),
)

// List all channels
channels, err := channel.ListChannels(client)

// Get a specific channel
ch, err := channel.GetChannel(client, 456)

// Update channel
ch.Enabled = false
updated, err := channel.UpdateChannel(client, *ch)

// Delete channel
err = channel.DeleteChannelById(client, 456)
```

### Working with Subaccounts

```go
import subaccount "github.com/mixser/flespi-client/resources/platform/subaccount"

// Create a subaccount
acc, err := subaccount.NewSubaccount(client, "my-subaccount",
    subaccount.WithMetadata(map[string]string{
        "department": "engineering",
    }),
)

// List all subaccounts
accounts, err := subaccount.ListSubaccounts(client)

// Get a specific subaccount
acc, err := subaccount.GetSubaccount(client, 999)

// Update subaccount
acc.Name = "updated-subaccount"
updated, err := subaccount.UpdateSubaccount(client, *acc)

// Delete subaccount
err = subaccount.DeleteSubaccountById(client, 999)
```

## Supported Resources

### Platform
- **Webhooks**: Single and chained webhook configurations
- **Subaccounts**: Manage subaccounts
- **Limits**: Account limits

### Gateway
- **Channels**: Device communication channels
- **Streams**: Message streams
- **Devices**: Connected devices
- **Calculators**: Data calculators with counters and selectors
- **Geofences**: Geographic boundaries
- **Tokens**: Access tokens

### Storage
- **Containers**: Data containers
- **CDN**: Content delivery network resources

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Building

```bash
go build ./...
```

## API Reference

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/mixser/flespi-client).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is provided as-is for use with the Flespi platform.

## Links

- [Flespi Documentation](https://flespi.io/docs)
- [Flespi API Reference](https://flespi.io/docs/#/platform)
- [Flespi Website](https://flespi.io)
