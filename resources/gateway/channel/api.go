package flespi_channel

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewChannelWithProtocolName(c flespiapi.APIRequester, name string, protocolName string, options ...CreateChannelOption) (*Channel, error) {
	channel := Channel{
		Name:          name,
		ProtocolName:  protocolName,
		Configuration: make(map[string]interface{}),
		Metadata:      make(map[string]string),
	}

	for _, opt := range options {
		opt(&channel)
	}

	return newChannel(c, channel)
}

func NewChannelWithProtocolId(c flespiapi.APIRequester, name string, protocolId int64, options ...CreateChannelOption) (*Channel, error) {
	channel := Channel{
		Name:          name,
		ProtocolId:    protocolId,
		Configuration: make(map[string]interface{}),
		Metadata:      make(map[string]string),
	}

	for _, opt := range options {
		opt(&channel)
	}

	return newChannel(c, channel)
}

func newChannel(c flespiapi.APIRequester, channel Channel) (*Channel, error) {
	response := channelsResponse{}

	var headers map[string]string
	if channel.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", channel.AccountId),
		}
	}

	// AccountId (cid) is conveyed via header on creation; zero it out so it is not sent in the request body.
	accountId := channel.AccountId
	channel.AccountId = 0
	defer func() {
		channel.AccountId = accountId
	}()

	if err := c.RequestAPIWithHeaders("POST", "gw/channels", headers, []Channel{channel}, &response); err != nil {
		return nil, err
	}

	return &response.Channels[0], nil
}

func ListChannels(c flespiapi.APIRequester) ([]Channel, error) {
	response := channelsResponse{}

	err := c.RequestAPI("GET", "gw/channels/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Channels, nil
}

func GetChannel(c flespiapi.APIRequester, channelId int64) (*Channel, error) {
	response := channelsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/channels/%d?fields=id,name,protocol_id,protocol_name,messages_ttl,enabled,configuration,metadata,cid", channelId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Channels[0], nil
}

func UpdateChannel(c flespiapi.APIRequester, channel Channel) (*Channel, error) {
	response := channelsResponse{}

	channelId := channel.Id
	accountId := channel.AccountId
	protocolName := channel.ProtocolName

	channel.Id = 0
	// ProtocolName is read-only for updates, clear it before sending
	channel.ProtocolName = ""
	channel.AccountId = 0

	defer func() {
		channel.Id = channelId
		channel.AccountId = accountId
		channel.ProtocolName = protocolName
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	err := c.RequestAPIWithHeaders("PUT", fmt.Sprintf("gw/channels/%d", channelId), headers, channel, &response)

	if err != nil {
		return nil, err
	}

	channel.Id = channelId

	return &response.Channels[0], nil
}

func DeleteChannel(c flespiapi.APIRequester, channel Channel) error {
	if channel.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteChannelById(c, channel.Id)
}

func DeleteChannelById(c flespiapi.APIRequester, channelId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/channels/%d", channelId), nil, nil)
}
