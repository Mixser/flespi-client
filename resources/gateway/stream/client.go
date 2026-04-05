package flespi_stream

import "github.com/mixser/flespi-client/internal/flespiapi"

// StreamClient provides receiver-based methods for managing Flespi streams.
// Access it via Client.Streams after creating a flespi.Client.
type StreamClient struct {
	c flespiapi.Doer
}

// NewStreamClient creates a StreamClient wrapping the given flespiapi.Doer.
func NewStreamClient(c flespiapi.Doer) *StreamClient {
	return &StreamClient{c: c}
}

func (sc *StreamClient) Create(name string, protocolId int64, options ...CreateStreamOption) (*Stream, error) {
	return NewStream(sc.c, name, protocolId, options...)
}

func (sc *StreamClient) List() ([]Stream, error) {
	return ListStreams(sc.c)
}

func (sc *StreamClient) Get(streamId int64) (*Stream, error) {
	return GetStream(sc.c, streamId)
}

func (sc *StreamClient) Update(stream Stream) (*Stream, error) {
	return UpdateStream(sc.c, stream)
}

func (sc *StreamClient) Delete(stream Stream) error {
	return DeleteStream(sc.c, stream)
}

func (sc *StreamClient) DeleteById(streamId int64) error {
	return DeleteStreamById(sc.c, streamId)
}
