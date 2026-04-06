package flespi_device

type Device struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`

	Enabled bool `json:"enabled"`

	Configuration map[string]string `json:"configuration"`

	DeviceTypeId int64 `json:"device_type_id"`

	MessagesTTL    int64 `json:"messages_ttl,omitempty"`
	MessagesRotate int64 `json:"messages_rotate,omitempty"`

	MediaTTL    int64 `json:"media_ttl,omitempty"`
	MediaRotate int64 `json:"media_rotate,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	// AccountId is the subaccount that owns this device (returned as "cid" in API responses).
	// On creation it is passed via the x-flespi-cid header, not the request body.
	AccountId int64 `json:"cid,omitempty"`
}

type CreateDeviceOption func(*Device)

func WithConfiguration(configuration map[string]string) CreateDeviceOption {
	return func(device *Device) {
		if configuration != nil {
			device.Configuration = configuration
		}
	}
}

func WithMessage(ttl int64, rotate int64) CreateDeviceOption {
	return func(device *Device) {
		device.MessagesTTL = ttl
		device.MessagesRotate = rotate
	}
}

func WithMedia(ttl int64, rotate int64) CreateDeviceOption {
	return func(device *Device) {
		device.MediaTTL = ttl
		device.MediaRotate = rotate
	}
}

func WithAccountId(accountId int64) CreateDeviceOption {
	return func(device *Device) {
		device.AccountId = accountId
	}
}

type devicesResponse struct {
	Devices []Device `json:"result"`
}
