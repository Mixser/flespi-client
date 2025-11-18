package flespi_token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/tokens" {
			t.Errorf("Expected path /platform/tokens, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"info": "test-token"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	token, err := NewToken(client, "test-token")

	if err != nil {
		t.Errorf("NewToken() error = %v", err)
	}

	if token.Id != 123 {
		t.Errorf("Expected ID 123, got %d", token.Id)
	}

	if token.Info != "test-token" {
		t.Errorf("Expected info 'test-token', got '%s'", token.Info)
	}
}

func TestGetToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 789,
				"info": "test-token"
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	token, err := GetToken(client, 789)
	if err != nil {
		t.Errorf("GetToken() error = %v", err)
	}

	if token.Id != 789 {
		t.Errorf("Expected ID 789, got %d", token.Id)
	}
}

func TestListTokens(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 1,
					"info": "token1"
				},
				{
					"id": 2,
					"info": "token2"
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	tokens, err := ListTokens(client)
	if err != nil {
		t.Errorf("ListTokens() error = %v", err)
	}

	if len(tokens) != 2 {
		t.Errorf("Expected 2 tokens, got %d", len(tokens))
	}
}

func TestDeleteToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteTokenById(client, 789)
	if err != nil {
		t.Errorf("DeleteTokenById() error = %v", err)
	}
}
