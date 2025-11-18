package flespi_subaccount

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewSubaccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/subaccounts" {
			t.Errorf("Expected path /platform/subaccounts, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-subaccount"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	subaccount, err := NewSubaccount(client, "test-subaccount")

	if err != nil {
		t.Errorf("NewSubaccount() error = %v", err)
	}

	if subaccount.Id != 123 {
		t.Errorf("Expected ID 123, got %d", subaccount.Id)
	}

	if subaccount.Name != "test-subaccount" {
		t.Errorf("Expected name 'test-subaccount', got '%s'", subaccount.Name)
	}
}

func TestGetSubaccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"name": "test-subaccount"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	subaccount, err := GetSubaccount(client, 789)
	if err != nil {
		t.Errorf("GetSubaccount() error = %v", err)
	}

	if subaccount.Id != 789 {
		t.Errorf("Expected ID 789, got %d", subaccount.Id)
	}
}

func TestListSubaccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"name": "subaccount1"
				},
				{
					"id": 2,
					"name": "subaccount2"
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	subaccounts, err := ListSubaccounts(client)
	if err != nil {
		t.Errorf("ListSubaccounts() error = %v", err)
	}

	if len(subaccounts) != 2 {
		t.Errorf("Expected 2 subaccounts, got %d", len(subaccounts))
	}
}

func TestDeleteSubaccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteSubaccountById(client, 789)
	if err != nil {
		t.Errorf("DeleteSubaccountById() error = %v", err)
	}
}
