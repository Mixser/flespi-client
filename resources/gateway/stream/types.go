package flespi_stream

type Stream struct {
	Id int64 `json:"id,omitempty"`

	Name       string `json:"name"`
	ProtocolId int64  `json:"protocol_id"`

	Enabled  bool  `json:"enabled"`
	QueueTTL int64 `json:"queue_ttl,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`

	Configuration map[string]string `json:"configuration"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateStreamOption func(*Stream)

func WithStatus(enabled bool) CreateStreamOption {
	return func(stream *Stream) {
		stream.Enabled = enabled
	}
}

func WithQueueTTL(ttl int64) CreateStreamOption {
	return func(stream *Stream) {
		stream.QueueTTL = ttl
	}
}

func WithValidateMessage(validateMessage string) CreateStreamOption {
	return func(stream *Stream) {
		stream.ValidateMessage = validateMessage
	}
}

func WithConfiguration(configuration map[string]string) CreateStreamOption {
	return func(stream *Stream) {
		if configuration != nil {
			stream.Configuration = configuration
		}
	}
}

func WithConfigurationItem(key, value string) CreateStreamOption {
	return func(stream *Stream) {
		stream.Configuration[key] = value
	}
}

func WithMetadata(metadata map[string]string) CreateStreamOption {
	return func(stream *Stream) {
		stream.Metadata = metadata
	}
}

type streamsResponse struct {
	Streams []Stream `json:"result"`
}
