package flespi_webhook

import "encoding/json"

func unmarshalWebhookResponse(response webhookResponse) ([]Webhook, error) {
	var result []Webhook

	for _, rawValue := range response.RawValue {
		webhook, err := unmarshalWebhook(rawValue)

		if err != nil {
			return nil, err
		}

		result = append(result, webhook)
	}

	return result, nil
}

func unmarshalWebhook(rawValue json.RawMessage) (Webhook, error) {
	var err error = nil
	var singleWebhook SingleWebhook

	if err = json.Unmarshal(rawValue, &singleWebhook); err == nil {
		return &singleWebhook, nil
	}

	var chainedWebhook ChainedWebhook
	if err = json.Unmarshal(rawValue, &chainedWebhook); err == nil {
		return &chainedWebhook, nil
	}

	return nil, err
}

func unmarshalConfiguration(rawValue json.RawMessage) (Configuration, error) {
	var configurationType struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(rawValue, &configurationType); err != nil {
		return nil, err
	}

	var configuration Configuration

	switch configurationType.Type {
	case "custom-server":
		var customServerConfiguration CustomServerConfiguration
		if err := json.Unmarshal(rawValue, &customServerConfiguration); err != nil {
			return nil, err
		}
		configuration = customServerConfiguration
	case "flespi-platform":
		var flespiConfiguration FlespiConfiguration
		if err := json.Unmarshal(rawValue, &flespiConfiguration); err != nil {
			return nil, err
		}
		configuration = flespiConfiguration
	}

	return configuration, nil
}
