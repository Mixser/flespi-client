package flespi_stream

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/gw/streams" {
			t.Errorf("Expected path /gw/streams, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-stream",
				"protocol_id": 1,
				"enabled": true,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	stream, err := NewStream(client, "test-stream", 1,
		WithStatus(true),
	)

	if err != nil {
		t.Errorf("NewStream() error = %v", err)
	}

	if stream.Id != 789 {
		t.Errorf("Expected ID 789, got %d", stream.Id)
	}

	if stream.Name != "test-stream" {
		t.Errorf("Expected name 'test-stream', got '%s'", stream.Name)
	}
}

func TestGetStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/gw/streams/789" {
			t.Errorf("Expected path /gw/streams/789, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-stream",
				"protocol_id": 1,
				"enabled": true,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	stream, err := GetStream(client, 789)
	if err != nil {
		t.Errorf("GetStream() error = %v", err)
	}

	if stream.Id != 789 {
		t.Errorf("Expected ID 789, got %d", stream.Id)
	}
}

func TestListStreams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "stream1",
					"protocol_id": 1,
					"enabled": true,
					"configuration": {}
				},
				{
					"id": 2,
					"name": "stream2",
					"protocol_id": 2,
					"enabled": false,
					"configuration": {}
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	streams, err := ListStreams(client)
	if err != nil {
		t.Errorf("ListStreams() error = %v", err)
	}

	if len(streams) != 2 {
		t.Errorf("Expected 2 streams, got %d", len(streams))
	}
}

func TestUpdateStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "updated-stream",
				"protocol_id": 1,
				"enabled": false,
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	stream := Stream{
		Id:            789,
		Name:          "updated-stream",
		ProtocolId:    1,
		Enabled:       false,
		Configuration: make(map[string]string),
	}

	updated, err := UpdateStream(client, stream)
	if err != nil {
		t.Errorf("UpdateStream() error = %v", err)
	}

	if updated.Name != "updated-stream" {
		t.Errorf("Expected name 'updated-stream', got '%s'", updated.Name)
	}

	if updated.Enabled {
		t.Errorf("Expected enabled to be false, got true")
	}
}

func TestDeleteStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteStreamById(client, 789)
	if err != nil {
		t.Errorf("DeleteStreamById() error = %v", err)
	}
}

func TestUpdateStream_MissingID(t *testing.T) {
	client, _ := flespi.NewClient("https://flespi.io", "test-token")

	stream := Stream{
		Name:          "stream-without-id",
		Configuration: make(map[string]string),
	}

	_, err := UpdateStream(client, stream)
	if err == nil {
		t.Errorf("Expected error for stream without ID, got nil")
	}
}
