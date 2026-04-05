package flespi_device

import (
	"fmt"
	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewDevice(c flespiapi.Doer, name string, enabled bool, deviceTypeId int64, options ...CreateDeviceOption) (*Device, error) {
	device := Device{
		Name:          name,
		Enabled:       enabled,
		DeviceTypeId:  deviceTypeId,
		Configuration: make(map[string]string),
	}

	for _, opt := range options {
		opt(&device)
	}

	response := devicesResponse{}

	err := c.RequestAPI("POST", "gw/devices", []Device{device}, &response)

	if err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func ListDevices(c flespiapi.Doer) ([]Device, error) {
	response := devicesResponse{}

	err := c.RequestAPI("GET", "gw/devices/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Devices, nil
}

func GetDevice(c flespiapi.Doer, deviceId int64) (*Device, error) {
	response := devicesResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/devices/%d", deviceId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func UpdateDevice(c flespiapi.Doer, device Device) (*Device, error) {
	response := devicesResponse{}

	deviceId := device.Id
	device.Id = 0

	err := c.RequestAPI("PUT", fmt.Sprintf("gw/devices/%d", deviceId), device, &response)

	if err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func DeleteDevice(c flespiapi.Doer, device Device) error {
	return DeleteDeviceById(c, device.Id)
}

func DeleteDeviceById(c flespiapi.Doer, deviceId int64) error {
	err := c.RequestAPI("DELETE", fmt.Sprintf("gw/devices/%d", deviceId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
