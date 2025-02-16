package flespi_channel

import (
	"fmt"

	"github.com/mixser/flespi-client"
)

func NewChannelWithProtocolName(c *flespi.Client, name string, protocolName string, options ...CreateChannelOption) (*Channel, error) {
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

func NewChannelWithProtocolId(c *flespi.Client, name string, protocolId int64, options ...CreateChannelOption) (*Channel, error) {
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

func newChannel(c *flespi.Client, channel Channel) (*Channel, error) {
	response := channelsResponse{}

	err := c.RequestAPI("POST", "gw/channels", []Channel{channel}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Channels[0], nil
}

func ListChannels(c *flespi.Client) ([]Channel, error) {
	response := channelsResponse{}

	err := c.RequestAPI("GET", "gw/channels/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Channels, nil
}

func GetChannel(c *flespi.Client, channelId int64) (*Channel, error) {
	response := channelsResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/channels/%d?fields=id,name,protocol_id,protocol_name,messages_ttl,enabled,configuration", channelId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Channels[0], nil
}

func UpdateChannel(c *flespi.Client, channel Channel) (*Channel, error) {
	response := channelsResponse{}

	channelId := channel.Id
	channel.Id = 0

	err := c.RequestAPI("PUT", fmt.Sprintf("gw/channels/%d", channelId), channel, &response)

	if err != nil {
		return nil, err
	}

	channel.Id = channelId

	return &response.Channels[0], nil
}

func DeleteChannel(c *flespi.Client, channel Channel) error {
	if channel.Id == 0 {
		return fmt.Errorf("ID must be provided")
	}

	return DeleteChannelById(c, channel.Id)
}

func DeleteChannelById(c *flespi.Client, channelId int64) error {
	return c.RequestAPI("DELETE", fmt.Sprintf("gw/channels/%d", channelId), nil, nil)
}
