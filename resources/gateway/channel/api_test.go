package flespi_channel

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewChannelWithProtocolName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/gw/channels" {
			t.Errorf("Expected path /gw/channels, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-channel",
				"protocol_name": "http",
				"enabled": true,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	channel, err := NewChannelWithProtocolName(client, "test-channel", "http")

	if err != nil {
		t.Errorf("NewChannelWithProtocolName() error = %v", err)
	}

	if channel.Id != 123 {
		t.Errorf("Expected ID 123, got %d", channel.Id)
	}

	if channel.Name != "test-channel" {
		t.Errorf("Expected name 'test-channel', got '%s'", channel.Name)
	}
}

func TestNewChannelWithProtocolId(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 456,
				"name": "test-channel",
				"protocol_id": 5,
				"enabled": true,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	channel, err := NewChannelWithProtocolId(client, "test-channel", 5)

	if err != nil {
		t.Errorf("NewChannelWithProtocolId() error = %v", err)
	}

	if channel.Id != 456 {
		t.Errorf("Expected ID 456, got %d", channel.Id)
	}
}

func TestGetChannel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-channel",
				"protocol_id": 1,
				"enabled": true,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	channel, err := GetChannel(client, 789)
	if err != nil {
		t.Errorf("GetChannel() error = %v", err)
	}

	if channel.Id != 789 {
		t.Errorf("Expected ID 789, got %d", channel.Id)
	}
}

func TestListChannels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "channel1",
					"protocol_id": 1,
					"enabled": true,
					"configuration": {}
				},
				{
					"id": 2,
					"name": "channel2",
					"protocol_id": 2,
					"enabled": false,
					"configuration": {}
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	channels, err := ListChannels(client)
	if err != nil {
		t.Errorf("ListChannels() error = %v", err)
	}

	if len(channels) != 2 {
		t.Errorf("Expected 2 channels, got %d", len(channels))
	}
}

func TestUpdateChannel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "updated-channel",
				"protocol_id": 1,
				"enabled": false,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	channel := Channel{
		Id:            789,
		Name:          "updated-channel",
		ProtocolId:    1,
		Enabled:       false,
		Configuration: make(map[string]interface{}),
	}

	updated, err := UpdateChannel(client, channel)
	if err != nil {
		t.Errorf("UpdateChannel() error = %v", err)
	}

	if updated.Name != "updated-channel" {
		t.Errorf("Expected name 'updated-channel', got '%s'", updated.Name)
	}
}

func TestDeleteChannel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteChannelById(client, 789)
	if err != nil {
		t.Errorf("DeleteChannelById() error = %v", err)
	}
}
