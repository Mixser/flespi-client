package flespi_device

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mixser/flespi-client/internal/testhelper"
)

func TestDeviceClient_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": [{"id": 1, "name": "test", "device_type_id": 5, "configuration": {}}]}`))
	}))
	defer server.Close()

	c := testhelper.New(server.URL)
	dc := NewDeviceClient(c)

	device, err := dc.Create("test", true, 5)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if device.Id != 1 {
		t.Errorf("expected ID 1, got %d", device.Id)
	}
}

func TestDeviceClient_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": [{"id": 1, "name": "d1", "device_type_id": 1, "configuration": {}}, {"id": 2, "name": "d2", "device_type_id": 2, "configuration": {}}]}`))
	}))
	defer server.Close()

	c := testhelper.New(server.URL)
	dc := NewDeviceClient(c)

	devices, err := dc.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(devices) != 2 {
		t.Errorf("expected 2 devices, got %d", len(devices))
	}
}

func TestDeviceClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": [{"id": 42, "name": "test", "device_type_id": 1, "configuration": {}}]}`))
	}))
	defer server.Close()

	c := testhelper.New(server.URL)
	dc := NewDeviceClient(c)

	device, err := dc.Get(42)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if device.Id != 42 {
		t.Errorf("expected ID 42, got %d", device.Id)
	}
}

func TestDeviceClient_DeleteById(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := testhelper.New(server.URL)
	dc := NewDeviceClient(c)

	if err := dc.DeleteById(1); err != nil {
		t.Fatalf("DeleteById() error = %v", err)
	}
}
