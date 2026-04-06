package flespi_stream

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewStream(c flespiapi.APIRequester, name string, protocolId int64, options ...CreateStreamOption) (*Stream, error) {
	stream := Stream{
		Name:          name,
		ProtocolId:    protocolId,
		Configuration: make(map[string]string),
	}

	for _, opt := range options {
		opt(&stream)
	}

	response := streamsResponse{}

	var headers map[string]string
	if stream.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", stream.AccountId),
		}
	}

	accountId := stream.AccountId
	stream.AccountId = 0
	defer func() { stream.AccountId = accountId }()

	if err := c.RequestAPIWithHeaders("POST", "gw/streams", headers, []Stream{stream}, &response); err != nil {
		return nil, err
	}

	return &response.Streams[0], nil
}

func GetStream(c flespiapi.APIRequester, streamId int64) (*Stream, error) {
	response := streamsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/streams/%d?fields=id,name,protocol_id,enabled,queue_ttl,validate_message,configuration,metadata,cid", streamId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Streams[0], nil
}

func ListStreams(c flespiapi.APIRequester) ([]Stream, error) {
	response := streamsResponse{}

	err := c.RequestAPI("GET", "gw/streams/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Streams, nil
}

func UpdateStream(c flespiapi.APIRequester, stream Stream) (*Stream, error) {
	if stream.Id == 0 {
		return nil, fmt.Errorf("ID must be provided")
	}

	streamId := stream.Id
	accountId := stream.AccountId

	stream.Id = 0
	stream.AccountId = 0

	defer func() {
		stream.Id = streamId
		stream.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	response := streamsResponse{}

	if err := c.RequestAPIWithHeaders("PUT", fmt.Sprintf("gw/streams/%d", streamId), headers, stream, &response); err != nil {
		return nil, err
	}

	return &response.Streams[0], nil
}

func DeleteStreamById(c flespiapi.APIRequester, streamId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/streams/%d", streamId), nil, nil)
}

func DeleteStream(c flespiapi.APIRequester, stream Stream) error {
	if stream.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteStreamById(c, stream.Id)
}
