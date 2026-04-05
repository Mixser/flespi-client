package flespi_device

import "github.com/mixser/flespi-client/internal/flespiapi"

// DeviceClient provides receiver-based methods for managing Flespi devices.
// Access it via Client.Devices after creating a flespi.Client.
type DeviceClient struct {
	c flespiapi.Doer
}

// NewDeviceClient creates a DeviceClient wrapping the given flespiapi.Doer.
func NewDeviceClient(c flespiapi.Doer) *DeviceClient {
	return &DeviceClient{c: c}
}

func (dc *DeviceClient) Create(name string, enabled bool, deviceTypeId int64, options ...CreateDeviceOption) (*Device, error) {
	return NewDevice(dc.c, name, enabled, deviceTypeId, options...)
}

func (dc *DeviceClient) List() ([]Device, error) {
	return ListDevices(dc.c)
}

func (dc *DeviceClient) Get(deviceId int64) (*Device, error) {
	return GetDevice(dc.c, deviceId)
}

func (dc *DeviceClient) Update(device Device) (*Device, error) {
	return UpdateDevice(dc.c, device)
}

func (dc *DeviceClient) Delete(device Device) error {
	return DeleteDevice(dc.c, device)
}

func (dc *DeviceClient) DeleteById(deviceId int64) error {
	return DeleteDeviceById(dc.c, deviceId)
}
