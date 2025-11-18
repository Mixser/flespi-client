package flespi_webhook

import "encoding/json"

const (
	ActionBreak = "break"
	ActionSkip  = "skip"
	ActionRetry = "retry"
)

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Validator struct {
	Expression string `json:"expression"`
	Action     string `json:"action"`
}

type Configuration interface {
	isConfigurationInstance()
}

type CustomServerConfiguration struct {
	Type     string     `json:"type"`
	Uri      string     `json:"uri"`
	Method   string     `json:"method"`
	Body     string     `json:"body"`
	CA       *string    `json:"ca,omitempty"`
	Headers  []Header   `json:"headers"`
	Validate *Validator `json:"validate,omitempty"`
}

type FlespiConfiguration struct {
	Type     string     `json:"type"`
	Uri      string     `json:"uri"`
	Method   string     `json:"method"`
	Body     string     `json:"body"`
	CID      string     `json:"cid"`
	Validate *Validator `json:"validate"`
}

func (c CustomServerConfiguration) isConfigurationInstance() {
	return
}


func (c FlespiConfiguration) isConfigurationInstance() {
	return
}


type TriggerFilter struct {
	CID     int64  `json:"cid"`
	Payload string `json:"payload"`
}

type Trigger struct {
	Topic  string         `json:"topic"`
	Filter *TriggerFilter `json:"filter,omitempty"`
}

type Webhook interface {

	GetId() int64
	isWebhookObject()
}

type SingleWebhook struct {
	Id            int64         `json:"id,omitempty"`
	Name          string        `json:"name"`
	Triggers      []Trigger     `json:"triggers"`
	Configuration Configuration `json:"configuration"`
}

func (sw *SingleWebhook) GetId() int64 {
	return sw.Id
}

func (sw *SingleWebhook) UnmarshalJSON(data []byte) error {
	var raw struct {
		Id       int64     `json:"id,omitempty"`
		Name     string    `json:"name"`
		Triggers []Trigger `json:"triggers"`
		Configuration json.RawMessage `json:"configuration"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sw.Id = raw.Id
	sw.Name = raw.Name
	sw.Triggers = raw.Triggers

	configuration, err := unmarshalConfiguration(raw.Configuration)

	if err != nil {
		return err
	}

	sw.Configuration = configuration

	return nil
}



func (sw *SingleWebhook) isWebhookObject() {
	return
}

type ChainedWebhook struct {
	Id            int64           `json:"id,omitempty"`
	Name          string          `json:"name"`
	Triggers      []Trigger       `json:"triggers"`
	Configuration []Configuration `json:"configuration"`
}

func (cw *ChainedWebhook) GetId() int64{
	return cw.Id
}

func (cw *ChainedWebhook) UnmarshalJSON(data []byte) error {
	var raw struct {
		Id       int64     `json:"id,omitempty"`
		Name     string    `json:"name"`
		Triggers []Trigger `json:"triggers"`
		Configuration []json.RawMessage `json:"configuration"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	cw.Id = raw.Id
	cw.Name = raw.Name
	cw.Triggers = raw.Triggers

	for _, rawCfg := range raw.Configuration {
		cfg, err := unmarshalConfiguration(rawCfg)

		if err != nil {
			return err
		}

		cw.Configuration = append(cw.Configuration, cfg)
	}
	return nil
}


func (cw *ChainedWebhook) isWebhookObject() {
	return
}


type CreateSingleWebhookOption func(*SingleWebhook)

func SWWithTrigger(trigger Trigger) CreateSingleWebhookOption {
	return func(webhook *SingleWebhook) {
		webhook.Triggers = append(webhook.Triggers, trigger)
	}
}

func SWWithTriggers(triggers []Trigger) CreateSingleWebhookOption {
	return func(webhook *SingleWebhook) {
		webhook.Triggers = triggers
	}
}

func SWWithConfiguration(cfg Configuration) CreateSingleWebhookOption {
	return func(webhook *SingleWebhook) {
		webhook.Configuration = cfg
	}
}

func CWWithTrigger(trigger Trigger) CreateChainedWebhookOption {
	return func(webhook *ChainedWebhook) {
		webhook.Triggers = append(webhook.Triggers, trigger)
	}
}

func CWWithTriggers(triggers []Trigger) CreateChainedWebhookOption {
	return func(webhook *ChainedWebhook) {
		webhook.Triggers = triggers
	}
}

func CWWithConfiguration(cfg Configuration) CreateChainedWebhookOption {
	return func(webhook *ChainedWebhook) {
		webhook.Configuration = append(webhook.Configuration, cfg)
	}
}

func CWWithConfigurations(configurations []Configuration) CreateChainedWebhookOption {
	return func(webhook *ChainedWebhook) {
		webhook.Configuration = configurations
	}
}


type CreateChainedWebhookOption func(*ChainedWebhook)


type webhookResponse struct {
	RawValue []json.RawMessage `json:"result"`
}