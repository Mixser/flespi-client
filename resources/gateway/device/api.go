package flespi_device

import (
	"fmt"
	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewDevice(c flespiapi.APIRequester, name string, enabled bool, deviceTypeId int64, options ...CreateDeviceOption) (*Device, error) {
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

	var headers map[string]string
	if device.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", device.AccountId),
		}
	}

	accountId := device.AccountId
	device.AccountId = 0
	defer func() { device.AccountId = accountId }()

	if err := c.RequestAPIWithHeaders("POST", "gw/devices", headers, []Device{device}, &response); err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func ListDevices(c flespiapi.APIRequester) ([]Device, error) {
	response := devicesResponse{}

	err := c.RequestAPI("GET", "gw/devices/all", nil, &response)

	if err != nil {
		return nil, err
	}

	return response.Devices, nil
}

func GetDevice(c flespiapi.APIRequester, deviceId int64) (*Device, error) {
	response := devicesResponse{}

	err := c.RequestAPI("GET", fmt.Sprintf("gw/devices/%d?fields=id,name,enabled,device_type_id,messages_ttl,messages_rotate,media_ttl,media_rotate,configuration,metadata,cid", deviceId), nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func UpdateDevice(c flespiapi.APIRequester, device Device) (*Device, error) {
	response := devicesResponse{}

	deviceId := device.Id
	accountId := device.AccountId

	device.Id = 0
	device.AccountId = 0

	defer func() {
		device.Id = deviceId
		device.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	if err := c.RequestAPIWithHeaders("PUT", fmt.Sprintf("gw/devices/%d", deviceId), headers, device, &response); err != nil {
		return nil, err
	}

	return &response.Devices[0], nil
}

func DeleteDevice(c flespiapi.APIRequester, device Device) error {
	return DeleteDeviceById(c, device.Id)
}

func DeleteDeviceById(c flespiapi.APIRequester, deviceId int64) error {
	err := c.RequestAPI("DELETE", fmt.Sprintf("gw/devices/%d", deviceId), nil, nil)

	if err != nil {
		return err
	}

	return nil
}
