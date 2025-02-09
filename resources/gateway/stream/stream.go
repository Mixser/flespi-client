package flespi_stream

import (
	"fmt"

	"github.com/mixser/flespi-client"
)

func NewStream(c *flespi.Client, name string, protocolId int64, options ...CreateStreamOption) (*Stream, error) {
	stream := Stream{
		Name:          name,
		ProtocolId:    protocolId,
		Configuration: make(map[string]string),
	}

	for _, opt := range options {
		opt(&stream)
	}

	response := streamsResponse{}

	err := c.RequestAPI("POST", "gw/streams", []Stream{stream}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Streams[0], nil
}

func GetStream(c *flespi.Client, streamId int64) (*Stream, error) {
	response := streamsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/streams/%d", streamId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Streams[0], nil
}

func ListStreams(c *flespi.Client) ([]Stream, error) {
	response := streamsResponse{}

	err := c.RequestAPI("GET", "gw/streams/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Streams, nil
}

func UpdateStream(c *flespi.Client, stream Stream) (*Stream, error) {
	if stream.Id == 0 {
		return nil, fmt.Errorf("ID must be provided")
	}

	streamId := stream.Id
	stream.Id = 0

	response := streamsResponse{}

	err := c.RequestAPI("PUT", fmt.Sprintf("gw/streams/%d", streamId), stream, &response)

	if err != nil {
		return nil, err
	}

	stream.Id = streamId
	return &response.Streams[0], nil
}

func DeleteStreamById(c *flespi.Client, streamId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/streams/%d", streamId), nil, nil)
}

func DeleteStream(c *flespi.Client, stream Stream) error {
	if stream.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteStreamById(c, stream.Id)
}
