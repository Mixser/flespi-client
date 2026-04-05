package flespi_channel

import "github.com/mixser/flespi-client/internal/flespiapi"

// ChannelClient provides receiver-based methods for managing Flespi channels.
// Access it via Client.Channels after creating a flespi.Client.
type ChannelClient struct {
	c flespiapi.APIRequester
}

// NewChannelClient creates a ChannelClient wrapping the given flespiapi.APIRequester.
func NewChannelClient(c flespiapi.APIRequester) *ChannelClient {
	return &ChannelClient{c: c}
}

func (cc *ChannelClient) CreateWithProtocolName(name string, protocolName string, options ...CreateChannelOption) (*Channel, error) {
	return NewChannelWithProtocolName(cc.c, name, protocolName, options...)
}

func (cc *ChannelClient) CreateWithProtocolId(name string, protocolId int64, options ...CreateChannelOption) (*Channel, error) {
	return NewChannelWithProtocolId(cc.c, name, protocolId, options...)
}

func (cc *ChannelClient) List() ([]Channel, error) {
	return ListChannels(cc.c)
}

func (cc *ChannelClient) Get(channelId int64) (*Channel, error) {
	return GetChannel(cc.c, channelId)
}

func (cc *ChannelClient) Update(channel Channel) (*Channel, error) {
	return UpdateChannel(cc.c, channel)
}

func (cc *ChannelClient) Delete(channel Channel) error {
	return DeleteChannel(cc.c, channel)
}

func (cc *ChannelClient) DeleteById(channelId int64) error {
	return DeleteChannelById(cc.c, channelId)
}
