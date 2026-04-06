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

	// AccountId is the subaccount that owns this stream (returned as "cid" in API responses).
	// On creation it is passed via the x-flespi-cid header, not the request body.
	AccountId int64 `json:"cid,omitempty"`
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

func WithAccountId(accountId int64) CreateStreamOption {
	return func(stream *Stream) {
		stream.AccountId = accountId
	}
}

type streamsResponse struct {
	Streams []Stream `json:"result"`
}
