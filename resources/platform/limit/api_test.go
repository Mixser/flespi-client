package flespi_limit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestGetLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-limit",
				"type": "devices"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	limit, err := GetLimit(client, 789)
	if err != nil {
		t.Errorf("GetLimit() error = %v", err)
	}

	if limit.Id != 789 {
		t.Errorf("Expected ID 789, got %d", limit.Id)
	}
}

func TestListLimits(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "limit1",
					"type": "devices"
				},
				{
					"id": 2,
					"name": "limit2",
					"type": "channels"
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	limits, err := ListLimits(client)
	if err != nil {
		t.Errorf("ListLimits() error = %v", err)
	}

	if len(limits) != 2 {
		t.Errorf("Expected 2 limits, got %d", len(limits))
	}
}
