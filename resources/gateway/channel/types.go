package flespi_channel

type Channel struct {
	Id            int64                  `json:"id,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	Name string `json:"name"`

	ProtocolId   int64  `json:"protocol_id,omitempty"`
	ProtocolName string `json:"protocol_name,omitempty"`

	Enabled bool `json:"enabled"`

	MessagesTTL int64 `json:"messages_ttl,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateChannelOption func(*Channel)

func WithStatus(enabled bool) CreateChannelOption {
	return func(channel *Channel) {
		channel.Enabled = enabled
	}
}

func WithMessagesTTL(ttl int64) CreateChannelOption {
	return func(channel *Channel) {
		channel.MessagesTTL = ttl
	}
}

func WithConfiguration(configuration map[string]interface{}) CreateChannelOption {
	return func(channel *Channel) {
		if configuration != nil {
			channel.Configuration = configuration
		}
	}
}

func WithConfigurationItem(key string, value interface{}) CreateChannelOption {
	return func(channel *Channel) {
		channel.Configuration[key] = value
	}
}

func WithMetadata(metadata map[string]string) CreateChannelOption {
	return func(channel *Channel) {
		channel.Metadata = metadata
	}
}

func WithMetadataItem(key, value string) CreateChannelOption {
	return func(channel *Channel) {
		channel.Metadata[key] = value
	}
}

type channelsResponse struct {
	Channels []Channel `json:"result"`
}
