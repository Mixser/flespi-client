package flespi_calculator

import (
	"encoding/json"
	"fmt"
)

type Calculator struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	MessagesSource MessagesSource `json:"messages_source,omitempty"`

	UpdatePeriod   int64 `json:"update_period,omitempty"`
	UpdateDelay    int64 `json:"update_delay,omitempty"`
	UpdateOnchange bool  `json:"update_onchange,omitempty"`

	IntervalsTTL    int64 `json:"intervals_ttl,omitempty"`
	IntervalsRotate int64 `json:"intervals_rotate,omitempty"`

	Selectors []Selector `json:"selectors"`
	Counters  []Counter  `json:"counters"`

	ValidateInterval string `json:"validate_interval,omitempty"`
	ValidateMessage  string `json:"validate_message,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

func (c *Calculator) UnmarshalJSON(data []byte) error {
	var raw struct {
		Id int64 `json:"id,omitempty"`

		Name string `json:"name"`

		MessagesSource json.RawMessage `json:"messages_source"`

		UpdatePeriod   int64 `json:"update_period"`
		UpdateDelay    int64 `json:"update_delay"`
		UpdateOnchange bool  `json:"update_onchange"`

		IntervalsTTL    int64 `json:"intervals_ttl"`
		IntervalsRotate int64 `json:"intervals_rotate"`

		Selectors []json.RawMessage `json:"selectors"`
		Counters  []json.RawMessage `json:"counters"`

		ValidateInterval string `json:"validate_interval,omitempty"`
		ValidateMessage  string `json:"validate_message,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		Metadata map[string]string `json:"metadata,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.Id = raw.Id
	c.Name = raw.Name
	c.UpdatePeriod = raw.UpdatePeriod
	c.UpdateDelay = raw.UpdateDelay
	c.UpdateOnchange = raw.UpdateOnchange
	c.IntervalsTTL = raw.IntervalsTTL
	c.IntervalsRotate = raw.IntervalsRotate
	c.ValidateInterval = raw.ValidateInterval
	c.ValidateMessage = raw.ValidateMessage
	c.Timezone = raw.Timezone

	var err error
	messagesSource, err := unmarshalMessageSource(raw.MessagesSource)

	if err != nil {
		return err
	}

	c.MessagesSource = messagesSource

	selectors, err := unmarshalSelectors(raw.Selectors)

	if err != nil {
		return err
	}

	c.Selectors = selectors

	counters, err := unmarshalCounters(raw.Counters)

	if err != nil {
		return err
	}

	c.Counters = counters

	c.Metadata = raw.Metadata

	return nil
}

type MessagesSource interface {
	GetSource() string
}

type DeviceSource struct {
	Source string `json:"source"`
}

func (ds *DeviceSource) GetSource() string {
	return "deivce"
}

type CalculatorSource struct {
	Source       string `json:"source"`
	CalculatorId int64  `json:"calculator_id"`
}

func (cs *CalculatorSource) GetSource() string {
	return "calculator"
}

type CreateCalculatorOption func(*Calculator)

func unmarshalMessageSource(raw json.RawMessage) (MessagesSource, error) {
	var source struct {
		Source       string `json:"source"`
		CalculatorId int64  `json:"calculator_id,omitempty"`
	}

	if err := json.Unmarshal(raw, &source); err != nil {
		return nil, err
	}

	var result MessagesSource
	var err error

	switch {
	case source.Source == "device":
		result = &DeviceSource{
			Source: "device",
		}
	case source.Source == "calculator":
		result = &CalculatorSource{
			Source: "calculator",
		}
	default:
		err = fmt.Errorf("unknown source: %s", source.Source)
	}

	return result, err
}

func WithDeviceMessageSource() CreateCalculatorOption {
	return func(calculator *Calculator) {
		calculator.MessagesSource = &DeviceSource{
			Source: "device",
		}
	}
}

func WithCalculatorMessageSource(calculatorId int64) CreateCalculatorOption {
	return func(calculator *Calculator) {
		calculator.MessagesSource = &CalculatorSource{
			Source:       "calculator",
			CalculatorId: calculatorId,
		}
	}
}

func WithSelector(selector Selector) CreateCalculatorOption {
	return func(calculator *Calculator) {
		calculator.Selectors = append(calculator.Selectors, selector)
	}
}

func WithCounter(counter Counter) CreateCalculatorOption {
	return func(calculator *Calculator) {
		calculator.Counters = append(calculator.Counters, counter)
	}
}

type calculatorsResponse struct {
	Calculators []Calculator `json:"result"`
}
