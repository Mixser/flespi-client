package main

import (
	"fmt"
	"log"

	flespi "github.com/mixser/flespi-client"
	flespi_stream "github.com/mixser/flespi-client/resources/gateway/stream"
)

func main() {
	// Create a new Flespi client
	client, err := flespi.NewClient("https://flespi.io", "your-flespi-token")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new stream
	stream, err := client.Streams.Create("example-stream", 1,
		flespi_stream.WithStatus(true),
		flespi_stream.WithQueueTTL(86400),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created stream: %s (ID: %d)\n", stream.Name, stream.Id)

	// List all streams
	all, err := client.Streams.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total streams: %d\n", len(all))

	// Update the stream
	stream.Enabled = false
	updated, err := client.Streams.Update(*stream)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated stream enabled status: %v\n", updated.Enabled)

	// Delete the stream
	err = client.Streams.Delete(*stream)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stream deleted successfully")
}
