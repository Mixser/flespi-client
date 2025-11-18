package flespi_webhook

import (
	"net/http"
	"net/http/httptest"
	"testing"

	flespi "github.com/mixser/flespi-client"
)

func TestNewSingleWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/webhooks" {
			t.Errorf("Expected path /platform/webhooks, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-webhook",
				"triggers": [],
				"configuration": {
					"type": "custom-server",
					"uri": "https://example.com",
					"method": "POST",
					"body": "{}",
					"headers": []
				}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	webhook, err := NewSingleWebhook(client, "test-webhook",
		SWWithConfiguration(CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com",
			Method:  "POST",
			Body:    "{}",
			Headers: []Header{},
		}),
	)

	if err != nil {
		t.Errorf("NewSingleWebhook() error = %v", err)
	}

	if webhook.Id != 123 {
		t.Errorf("Expected ID 123, got %d", webhook.Id)
	}

	if webhook.Name != "test-webhook" {
		t.Errorf("Expected name 'test-webhook', got '%s'", webhook.Name)
	}
}

func TestNewChainedWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 456,
				"name": "chained-webhook",
				"triggers": [],
				"configuration": [{
					"type": "custom-server",
					"uri": "https://example.com",
					"method": "POST",
					"body": "{}",
					"headers": []
				}]
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	webhook, err := NewChainedWebhook(client, "chained-webhook",
		CWWithConfiguration(CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com",
			Method:  "POST",
			Body:    "{}",
			Headers: []Header{},
		}),
	)

	if err != nil {
		t.Errorf("NewChainedWebhook() error = %v", err)
	}

	if webhook.Id != 456 {
		t.Errorf("Expected ID 456, got %d", webhook.Id)
	}
}

func TestGetWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/webhooks/123" {
			t.Errorf("Expected path /platform/webhooks/123, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "test-webhook",
				"triggers": [],
				"configuration": {
					"type": "custom-server",
					"uri": "https://example.com",
					"method": "POST",
					"body": "{}",
					"headers": []
				}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	webhook, err := GetWebhook(client, 123)
	if err != nil {
		t.Errorf("GetWebhook() error = %v", err)
	}

	if webhook.GetId() != 123 {
		t.Errorf("Expected ID 123, got %d", webhook.GetId())
	}
}

func TestListWebhooks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/webhooks/all" {
			t.Errorf("Expected path /platform/webhooks/all, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [
				{
					"id": 123,
					"name": "webhook1",
					"triggers": [],
					"configuration": {
						"type": "custom-server",
						"uri": "https://example.com",
						"method": "POST",
						"body": "{}",
						"headers": []
					}
				},
				{
					"id": 456,
					"name": "webhook2",
					"triggers": [],
					"configuration": [{
						"type": "custom-server",
						"uri": "https://example.com",
						"method": "POST",
						"body": "{}",
						"headers": []
					}]
				}
			]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	webhooks, err := ListWebhooks(client)
	if err != nil {
		t.Errorf("ListWebhooks() error = %v", err)
	}

	if len(webhooks) != 2 {
		t.Errorf("Expected 2 webhooks, got %d", len(webhooks))
	}
}

func TestDeleteWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/webhooks/123" {
			t.Errorf("Expected path /platform/webhooks/123, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	err := DeleteWebhookById(client, 123)
	if err != nil {
		t.Errorf("DeleteWebhookById() error = %v", err)
	}
}

func TestUpdateWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/platform/webhooks/123" {
			t.Errorf("Expected path /platform/webhooks/123, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"result": [{
				"id": 123,
				"name": "updated-webhook",
				"triggers": [],
				"configuration": {
					"type": "custom-server",
					"uri": "https://example.com",
					"method": "POST",
					"body": "{}",
					"headers": []
				}
			}]
		}`))
	}))
	defer server.Close()

	client, _ := flespi.NewClient(server.URL, "test-token")

	webhook := &SingleWebhook{
		Id:   123,
		Name: "updated-webhook",
		Configuration: CustomServerConfiguration{
			Type:    "custom-server",
			Uri:     "https://example.com",
			Method:  "POST",
			Body:    "{}",
			Headers: []Header{},
		},
	}

	updated, err := UpdateWebhook(client, webhook)
	if err != nil {
		t.Errorf("UpdateWebhook() error = %v", err)
	}

	if updated.GetId() != 123 {
		t.Errorf("Expected ID 123, got %d", updated.GetId())
	}
}
