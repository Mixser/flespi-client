package flespi_cdn

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewCDN(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/storage/cdns" {
			t.Errorf("Expected path /storage/cdns, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-cdn",
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	cdn, err := NewCDN(client, "test-cdn")

	if err != nil {
		t.Errorf("NewCDN() error = %v", err)
	}

	if cdn.Id != 123 {
		t.Errorf("Expected ID 123, got %d", cdn.Id)
	}

	if cdn.Name != "test-cdn" {
		t.Errorf("Expected name 'test-cdn', got '%s'", cdn.Name)
	}
}

func TestGetCDN(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-cdn",
				"configuration": {}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	cdn, err := GetCDN(client, 789)
	if err != nil {
		t.Errorf("GetCDN() error = %v", err)
	}

	if cdn.Id != 789 {
		t.Errorf("Expected ID 789, got %d", cdn.Id)
	}
}

func TestListCDNs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "cdn1",
					"configuration": {}
				},
				{
					"id": 2,
					"name": "cdn2",
					"configuration": {}
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	cdns, err := ListCDNs(client)
	if err != nil {
		t.Errorf("ListCDNs() error = %v", err)
	}

	if len(cdns) != 2 {
		t.Errorf("Expected 2 CDNs, got %d", len(cdns))
	}
}

func TestDeleteCDN(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteCDNById(client, 789)
	if err != nil {
		t.Errorf("DeleteCDNById() error = %v", err)
	}
}
