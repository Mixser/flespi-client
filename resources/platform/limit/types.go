package flespi_limit

type Limit struct {
	Id               int64  `json:"id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	BlockingDuration int    `json:"blocking_duration,omitempty"`
	ApiCall          int64  `json:"api_calls,omitempty"`
	ApiTraffic       int64  `json:"api_traffic,omitempty"`

	ChannelsCount      int64 `json:"channels_count,omitempty"`
	ChannelMessages    int64 `json:"channel_messages,omitempty"`
	ChannelStorage     int64 `json:"channel_storage,omitempty"`
	ChannelTraffic     int64 `json:"channel_traffic,omitempty"`
	ChannelConnections int64 `json:"channel_connections,omitempty"`

	ContainersCount  int64 `json:"containers_count,omitempty"`
	ContainerStorage int64 `json:"container_storage,omitempty"`

	CdnsCount  int64 `json:"cdns_count,omitempty"`
	CdnStorage int64 `json:"cdn_storage,omitempty"`
	CdnTraffic int64 `json:"cdn_traffic,omitempty"`

	DevicesCount       int64 `json:"devices_count,omitempty"`
	DeviceStorage      int64 `json:"device_storage,omitempty"`
	DeviceMediaTraffic int64 `json:"device_media_traffic,omitempty"`
	DeviceMediaStorage int64 `json:"device_media_storage,omitempty"`

	StreamsCount  int64 `json:"streams_count,omitempty"`
	StreamStorage int64 `json:"stream_storage,omitempty"`
	StreamTraffic int64 `json:"stream_traffic,omitempty"`

	ModemsCount int64 `json:"modems_count,omitempty"`

	MqttSessions        int64 `json:"mqtt_sessions,omitempty"`
	MqttMessages        int64 `json:"mqtt_messages,omitempty"`
	MqttSessionStorage  int64 `json:"mqtt_session_storage,omitempty"`
	MqttRetainedStorage int64 `json:"mqtt_retained_storage,omitempty"`
	MqttSubscriptions   int64 `json:"mqtt_subscriptions,omitempty"`

	SmsCount int64 `json:"sms_count,omitempty"`

	TokensCount int64 `json:"tokens_count,omitempty"`

	SubaccountsCount int64 `json:"subaccounts_count,omitempty"`

	LimitsCount int64 `json:"limits_count,omitempty"`

	RealmsCount int64 `json:"realms_count,omitempty"`

	CalcsCount   int64 `json:"calcs_count,omitempty"`
	CalcsStorage int64 `json:"calcs_storage,omitempty"`

	PluginsCount           int64 `json:"plugins_count,omitempty"`
	PluginTraffic          int64 `json:"plugin_traffic,omitempty"`
	PluginBufferedMessages int64 `json:"plugin_buffered_messages,omitempty"`

	GroupsCount int64 `json:"groups_count,omitempty"`

	WebhooksCount  int64 `json:"webhooks_count,omitempty"`
	WebhookStorage int64 `json:"webhook_storage,omitempty"`
	WebhookTraffic int64 `json:"webhook_traffic,omitempty"`

	GrantsCount int64 `json:"grants_count,omitempty"`

	IdentityProvidersCount int64 `json:"identity_providers_count,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateLimitOption func(*Limit)

func WithDescription(description string) CreateLimitOption {
	return func (limit *Limit) {
		limit.Description = description
	}
}

func WithBlockingDurationLimit(duration int) CreateLimitOption {
	return func(limit *Limit) {
		limit.BlockingDuration = duration
	}
}

func WithApiLimit(apiCall int64, apiTraffic int64) CreateLimitOption {
	return func (limit *Limit) {
		limit.ApiCall = apiCall
		limit.ApiTraffic = apiTraffic
	}
}


func WithChannelLimit(count int64, messages int64, storage int64, traffic int64, connections int64) CreateLimitOption {
	return func (limit *Limit) {
		limit.ChannelsCount = count
		limit.ChannelMessages = messages
		limit.ChannelStorage = storage
		limit.ChannelTraffic = traffic
		limit.ChannelConnections = connections
	}
}

func WithContainerLimit(count int64, storage int64) CreateLimitOption {
	return func (limit *Limit) {
		limit.ContainersCount = count
		limit.ContainerStorage = storage
	}
}

func WithCdnLimit(count int64, storage int64, traffic int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.CdnsCount = count
		limit.CdnStorage = storage
		limit.CdnTraffic = traffic
	}
}

func WithDeviceLimit(count int64, storage int64, mediaTraffic int64, mediaStorage int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.DevicesCount = count
		limit.DeviceStorage = storage
		limit.DeviceMediaTraffic = mediaTraffic
		limit.DeviceMediaStorage = mediaStorage
	}
}

func WithStreamLimit(count int64, storage int64, traffic int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.StreamsCount = count
		limit.StreamStorage = storage
		limit.StreamTraffic = traffic
	}
}

func WithModelLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.ModemsCount = count
	}
}

func WithMqttLimit(sessions int64, messages int64, sessionStorage int64, retainedStorage int64, subscriptions int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.MqttSessions = sessions
		limit.MqttMessages = messages
		limit.MqttSessionStorage = sessionStorage
		limit.MqttRetainedStorage = retainedStorage
		limit.MqttSubscriptions = subscriptions
	}
}

func WithSmsLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.SmsCount = count
	}
}

func WithTokenLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.TokensCount = count
	}
}


func WithLimitLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.LimitsCount = count
	}
}

func WithRealmLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.RealmsCount = count
	}
}


func WithCalcLimit(count int64, storage int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.CalcsCount = count
		limit.CalcsStorage = storage
	}
}

func WithPluginLimit(count int64, traffic int64, bufferedMessages int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.PluginsCount = count
		limit.PluginTraffic = traffic
		limit.PluginBufferedMessages = bufferedMessages
	}
}

func WithGroupLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.GroupsCount = count
	}
}

func WithWebhookLimit(count int64, storage int64, traffic int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.WebhooksCount = count
		limit.WebhookStorage = storage
		limit.WebhookTraffic = traffic
	}
}

func WithGrantLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.GrantsCount = count
	}
}

func WithIdentityProviderLimit(count int64) CreateLimitOption {
	return func(limit *Limit) {
		limit.IdentityProvidersCount = count
	}
}

type limitsResponse struct {
	Limits []Limit `json:"result"`
}