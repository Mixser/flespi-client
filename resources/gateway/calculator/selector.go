package flespi_calculator

import (
	"encoding/json"
	"fmt"

	flespi_geofence "github.com/mixser/flespi-client/resources/gateway/geofence"
)

type Selector interface {
	GetSelectorType() string
}

type SelectorExpression struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Expression string `json:"expression"`

	Invert bool `json:"invert"`

	MaxActive int64 `json:"max_active,omitempty"`
	MinActive int64 `json:"min_active,omitempty"`

	MaxInactive int64 `json:"max_inactive,omitempty"`

	MinDuration int64 `json:"min_duration,omitempty"`

	MaxMessagesTimeDiff int64 `json:"max_messages_time_diff,omitempty"`

	MergeMessageAffter bool `json:"merge_message_after,omitempty"`
	MergeMessageBefore bool `json:"merge_message_before,omitempty"`

	Method string `json:"method,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func (se *SelectorExpression) GetSelectorType() string {
	return "expression"
}

func NewSelectorExpression(name string, expression string, options ...CreateSelectorExpressionOption) *SelectorExpression {
	selectorExpression := SelectorExpression{
		Name:       name,
		Expression: expression,
		Type:       "expression",
	}

	for _, opt := range options {
		opt(&selectorExpression)
	}

	return &selectorExpression
}

type CreateSelectorExpressionOption func(expression *SelectorExpression)

type SelectorDatetime struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Split string `json:"split,omitempty"`

	MaxMessagesTimeDiff int64 `json:"max_messages_time_diff,omitempty"`

	MergeMessageAffter bool `json:"merge_message_after,omitempty"`
	MergeMessageBefore bool `json:"merge_message_before,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func (sd *SelectorDatetime) GetSelectorType() string {
	return "datetime"
}

func NewSelectorDateOrTime(name string, options ...CreateSelectorDateOrTimeOption) *SelectorDatetime {
	selectorDateOrTime := SelectorDatetime{
		Name: name,
		Type: "datetime",
	}

	for _, opt := range options {
		opt(&selectorDateOrTime)
	}

	return &selectorDateOrTime
}

type CreateSelectorDateOrTimeOption func(dateOrTime *SelectorDatetime)

type SelectorGeofence struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Geofences []flespi_geofence.GeofenceGeometry `json:"geofences"`

	MaxActive int64 `json:"max_active,omitempty"`
	MinActive int64 `json:"min_active,omitempty"`

	MaxInactive int64 `json:"max_inactive,omitempty"`

	MinDuration int64 `json:"min_duration,omitempty"`

	MaxMessagesTimeDiff int64 `json:"max_messages_time_diff,omitempty"`

	MergeMessageAffter bool `json:"merge_message_after,omitempty"`
	MergeMessageBefore bool `json:"merge_message_before,omitempty"`

	MergeUnknown bool `json:"merge_unknown,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func (sg *SelectorGeofence) GetSelectorType() string {
	return "geofence"
}

func (sg *SelectorGeofence) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name string `json:"name"`
		Type string `json:"type"`

		Geofences []json.RawMessage `json:"geofences"`

		MaxActive int64 `json:"max_active,omitempty"`
		MinActive int64 `json:"min_active,omitempty"`

		MaxInactive int64 `json:"max_inactive,omitempty"`

		MinDuration int64 `json:"min_duration,omitempty"`

		MaxMessagesTimeDiff int64 `json:"max_messages_time_diff,omitempty"`

		MergeMessageAffter bool `json:"merge_message_after,omitempty"`
		MergeMessageBefore bool `json:"merge_message_before,omitempty"`

		MergeUnknown bool `json:"merge_unknown,omitempty"`

		ValidateMessage string `json:"validate_message,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sg.Name = raw.Name
	sg.Type = raw.Type

	sg.MaxActive = raw.MaxActive
	sg.MinActive = raw.MinActive
	sg.MaxInactive = raw.MaxInactive
	sg.MinDuration = raw.MinDuration
	sg.MaxMessagesTimeDiff = raw.MaxMessagesTimeDiff
	sg.MergeMessageAffter = raw.MergeMessageAffter
	sg.MergeMessageBefore = raw.MergeMessageBefore
	sg.MergeUnknown = raw.MergeUnknown
	sg.ValidateMessage = raw.ValidateMessage

	for _, rawGeofence := range raw.Geofences {
		geometry, err := flespi_geofence.UnmarshalGeometry(rawGeofence)

		if err != nil {
			return err
		}

		sg.Geofences = append(sg.Geofences, geometry)
	}

	return nil
}

func NewSelectorGeofence(name string, options ...CreateSelectorGeofenceOption) *SelectorGeofence {
	selectorGeofence := SelectorGeofence{
		Name: name,
		Type: "geofence",
	}

	for _, opt := range options {
		opt(&selectorGeofence)
	}

	return &selectorGeofence
}

type CreateSelectorGeofenceOption func(selector *SelectorGeofence)


func WithGeometry(geometry flespi_geofence.GeofenceGeometry) CreateSelectorGeofenceOption {
	return func(geofence *SelectorGeofence) {
		geofence.Geofences = append(geofence.Geofences, geometry)
	}
}

type SelectorCalculator struct {
	CalculatorId int64  `json:"calculator_id"`
	Type         string `json:"type"`

	Invert bool `json:"invert"`

	MaxActive int64 `json:"max_active,omitempty"`
	MinActive int64 `json:"min_active,omitempty"`

	MaxInactive int64 `json:"max_inactive,omitempty"`

	MinDuration int64 `json:"min_duration,omitempty"`

	MaxMessagesTimeDiff int64 `json:"max_messages_time_diff,omitempty"`

	ValidateInterval string `json:"validate_interval,omitempty"`
}

func (sc *SelectorCalculator) GetSelectorType() string {
	return "calculator"
}

func NewSelectorCalculator(calculatorId int64, options ...CreateSelectorCalculatorOption) *SelectorCalculator {
	selectorCalculator := SelectorCalculator{
		CalculatorId: calculatorId,
		Type:         "calculator",
	}

	for _, opt := range options {
		opt(&selectorCalculator)
	}

	return &selectorCalculator
}

type CreateSelectorCalculatorOption func(selector *SelectorCalculator)

type SelectorInactive struct {
	Name string `json:"name"`
	Type string `json:"type"`

	DelayThreshold int64 `json:"delay_threshold"`
}

func (si *SelectorInactive) GetSelectorType() string {
	return "inactive"
}

func NewSelectorInactive(name string, delayThreshold int64) *SelectorInactive {
	selectorInactive := SelectorInactive{
		Name:           name,
		Type:           "inactive",
		DelayThreshold: delayThreshold,
	}

	return &selectorInactive
}

func unmarshalSelectors(rawSelectors []json.RawMessage) ([]Selector, error) {
	var selectors []Selector = []Selector{}

	for _, rawSelector := range rawSelectors {
		selector, err := unmarshallSelector(rawSelector)
		if err != nil {
			return nil, err
		}
		selectors = append(selectors, selector)
	}

	return selectors, nil
}

func unmarshallSelector(raw json.RawMessage) (Selector, error) {
	var rawSelector struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(raw, &rawSelector); err != nil {
		return nil, err
	}

	var result Selector = nil
	var err error = nil

	switch {
	case rawSelector.Type == "expression":
		selector := &SelectorExpression{}
		err = json.Unmarshal(raw, &selector)
		result = selector
	case rawSelector.Type == "datetime":
		selector := &SelectorDatetime{}
		err = json.Unmarshal(raw, &selector)
		result = selector
	case rawSelector.Type == "geofence":
		selector := &SelectorGeofence{}
		err = json.Unmarshal(raw, &selector)
		result = selector
	case rawSelector.Type == "calculator":
		selector := &SelectorCalculator{}
		err = json.Unmarshal(raw, &selector)
		result = selector
	case rawSelector.Type == "inactive":
		selector := &SelectorInactive{}
		err = json.Unmarshal(raw, &selector)
		result = selector
	default:
		err = fmt.Errorf("unknown selector type: %s", rawSelector.Type)
	}

	return result, err
}
