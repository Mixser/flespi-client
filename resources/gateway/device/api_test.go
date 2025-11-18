package flespi_device

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewDevice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/gw/devices" {
			t.Errorf("Expected path /gw/devices, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-device",
				"device_type_id": 5,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	device, err := NewDevice(client, "test-device", true, 5)

	if err != nil {
		t.Errorf("NewDevice() error = %v", err)
	}

	if device.Id != 123 {
		t.Errorf("Expected ID 123, got %d", device.Id)
	}

	if device.Name != "test-device" {
		t.Errorf("Expected name 'test-device', got '%s'", device.Name)
	}
}

func TestGetDevice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-device",
				"device_type_id": 1,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	device, err := GetDevice(client, 789)
	if err != nil {
		t.Errorf("GetDevice() error = %v", err)
	}

	if device.Id != 789 {
		t.Errorf("Expected ID 789, got %d", device.Id)
	}
}

func TestListDevices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "device1",
					"device_type_id": 1,
					"configuration": {}
				},
				{
					"id": 2,
					"name": "device2",
					"device_type_id": 2,
					"configuration": {}
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	devices, err := ListDevices(client)
	if err != nil {
		t.Errorf("ListDevices() error = %v", err)
	}

	if len(devices) != 2 {
		t.Errorf("Expected 2 devices, got %d", len(devices))
	}
}

func TestUpdateDevice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "updated-device",
				"device_type_id": 1,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	device := Device{
		Id:            789,
		Name:          "updated-device",
		DeviceTypeId:  1,
		Configuration: make(map[string]string),
	}

	updated, err := UpdateDevice(client, device)
	if err != nil {
		t.Errorf("UpdateDevice() error = %v", err)
	}

	if updated.Name != "updated-device" {
		t.Errorf("Expected name 'updated-device', got '%s'", updated.Name)
	}
}

func TestDeleteDevice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteDeviceById(client, 789)
	if err != nil {
		t.Errorf("DeleteDeviceById() error = %v", err)
	}
}
