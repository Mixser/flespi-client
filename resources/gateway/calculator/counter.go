package flespi_calculator

import (
	"encoding/json"
	"fmt"
)

type Counter interface{}

type CounterExpression struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Expression      string `json:"expression"`
	Method          string `json:"method,omitempty"`
	ValidateMessage string `json:"validate_message,omitempty"`
}

func NewCounterExpression(name string, expression string, options ...CreateCounterExpressionOption) *CounterExpression {
	counterExpression := CounterExpression{
		Name:       name,
		Expression: expression,
		Type:       "expression",
	}

	for _, opt := range options {
		opt(&counterExpression)
	}

	return &counterExpression
}

type CreateCounterExpressionOption func(expression *CounterExpression)

func CEWithMethod(method string) CreateCounterExpressionOption {
	return func(expression *CounterExpression) {
		expression.Method = method
	}
}

type CounterDataset struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Fields []CounterDatasetField `json:"fields"`

	AllowUnknown bool `json:"allow_unknown,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

type CounterDatasetField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewCounterDataset(name string, fields []CounterDatasetField, options ...CreateCounterDatasetOption) *CounterDataset {
	counterDataset := CounterDataset{
		Name:   name,
		Type:   "dataset",
		Fields: fields,
	}

	for _, opt := range options {
		opt(&counterDataset)
	}

	return &counterDataset
}

type CreateCounterDatasetOption func(dataset *CounterDataset)

func CDWithAllowUnknown(allowUnknown bool) CreateCounterDatasetOption {
	return func(dataset *CounterDataset) {
		dataset.AllowUnknown = allowUnknown
	}
}

func CDWithValidateMessage(validateMessage string) CreateCounterDatasetOption {
	return func(dataset *CounterDataset) {
		dataset.ValidateMessage = validateMessage
	}
}

type CounterRoute struct {
	Name string `json:"name"`
	Type string `json:"type"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func NewCounterRoute(name string, options ...CreateCounterRouteOption) *CounterRoute {
	counterRoute := CounterRoute{
		Name: name,
		Type: "route",
	}

	for _, opt := range options {
		opt(&counterRoute)
	}

	return &counterRoute
}

type CreateCounterRouteOption func(route *CounterRoute)

func CRWithValidateMessage(validateMessage string) CreateCounterRouteOption {
	return func(route *CounterRoute) {
		route.ValidateMessage = validateMessage
	}
}

type CounterDatetime struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Format string `json:"format,omitempty"`

	Interval string `json:"interval,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func NewCounterDatetime(name string, options ...CreateCounterDatetimeOption) *CounterDatetime {
	counterDatetime := CounterDatetime{
		Name: name,
		Type: "datetime",
	}

	for _, opt := range options {
		opt(&counterDatetime)
	}

	return &counterDatetime
}

type CreateCounterDatetimeOption func(datetime *CounterDatetime)

func CDatetimeWithFormat(format string) CreateCounterDatetimeOption {
	return func(datetime *CounterDatetime) {
		datetime.Format = format
	}
}

func CDatetimeWithInterval(interval string) CreateCounterDatetimeOption {
	return func(datetime *CounterDatetime) {
		datetime.Interval = interval
	}
}

func CDatetimeWithValidateMessage(validateMessage string) CreateCounterDatetimeOption {
	return func(datetime *CounterDatetime) {
		datetime.ValidateMessage = validateMessage
	}
}

type CounterParameter struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Parameter string `json:"parameter"`
	Method    string `json:"method,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func NewCounterParameter(name string, parameter string, options ...CreateCounterParameterOption) *CounterParameter {
	counterParameter := CounterParameter{
		Name:      name,
		Parameter: parameter,
		Type:      "parameter",
	}

	for _, opt := range options {
		opt(&counterParameter)
	}

	return &counterParameter
}

type CreateCounterParameterOption func(parameter *CounterParameter)

func CPWithMethod(method string) CreateCounterParameterOption {
	return func(expression *CounterParameter) {
		expression.Method = method
	}
}

func CPWithValidateMessage(validateMessage string) CreateCounterParameterOption {
	return func(expression *CounterParameter) {
		expression.ValidateMessage = validateMessage
	}
}

type CounterMessage struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Method string `json:"method,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Extremum CounterMessageExtremum `json:"extremum,omitempty"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

type CounterMessageExtremum struct {
	Type       string `json:"type"`
	Expression string `json:"Expression"`
}

func NewCounterMessage(name string, options ...CreateCounterMessageOption) *CounterMessage {
	counterMessage := CounterMessage{
		Name: name,
		Type: "message",
	}

	for _, opt := range options {
		opt(&counterMessage)
	}

	return &counterMessage
}

type CreateCounterMessageOption func(message *CounterMessage)

func CMWithMethod(method string) CreateCounterMessageOption {
	return func(message *CounterMessage) {
		message.Method = method
	}
}

func CMWithFields(fields []string) CreateCounterMessageOption {
	return func(message *CounterMessage) {
		message.Fields = fields
	}
}

func CMWithExtremum(extremumType string, expression string) CreateCounterMessageOption {
	return func(message *CounterMessage) {
		message.Extremum = CounterMessageExtremum{
			Type:       extremumType,
			Expression: expression,
		}
	}
}

func CMWithValidateMessage(validateMessage string) CreateCounterMessageOption {
	return func(message *CounterMessage) {
		message.ValidateMessage = validateMessage
	}
}

type CounterInterval struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

func NewCounterInterval(name string, expression string) *CounterInterval {
	return &CounterInterval{
		Name:       name,
		Type:       "interval",
		Expression: expression,
	}
}

type CounterActive struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewCounterActive(name string) *CounterActive {
	return &CounterActive{
		Name: name,
		Type: "active",
	}
}

type CounterGeofence struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewCounterGeofence(name string) *CounterGeofence {
	return &CounterGeofence{
		Name: name,
		Type: "geofence",
	}
}

type CounterVariable struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Expression string `json:"expression"`

	ValidateMessage string `json:"validate_message,omitempty"`
}

func NewCounterVariable(name string, expression string, options ...CreateCounterVariableOption) *CounterVariable {
	counterVariable := CounterVariable{
		Name:       name,
		Expression: expression,
		Type:       "variable",
	}

	for _, opt := range options {
		opt(&counterVariable)
	}

	return &counterVariable
}

type CreateCounterVariableOption func(variable *CounterVariable)

func CVWithValidateMessage(validateMessage string) CreateCounterVariableOption {
	return func(variable *CounterVariable) {
		variable.ValidateMessage = validateMessage
	}
}

type CounterSpecifiedString struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewCounterSpecifiedString(name string, value string) *CounterSpecifiedString {
	return &CounterSpecifiedString{
		Name:  name,
		Value: value,
		Type:  "specified",
	}
}

type CounterSpecifiedNumber struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

func NewCounterSpecifiedNumber(name string, value float64) *CounterSpecifiedNumber {
	return &CounterSpecifiedNumber{
		Name:  name,
		Value: value,
		Type:  "specified",
	}
}

type CounterSpecifiedBoolean struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value bool   `json:"value"`
}

func NewCounterSpecifiedBoolean(name string, value bool) *CounterSpecifiedBoolean {
	return &CounterSpecifiedBoolean{
		Name:  name,
		Value: value,
		Type:  "specified",
	}
}

type CounterCalculator struct {
	Name         string `json:"name"`
	CalculatorId int64  `json:"calc_id"`

	Type string `json:"type"`

	AllowFinishAfter bool `json:"allow_finish_after,omitempty"`
	AllowStartBefore bool `json:"allow_start_before,omitempty"`

	// TODO: support here computed fields
	Fields []string `json:"fields,omitempty"`

	Method string `json:"method,omitempty"`

	ValidateInterval string `json:"validate_interval,omitempty"`
}

func NewCounterCalculator(name string, calcId int64, options ...CreateCounterCalculatorOption) *CounterCalculator {
	counterCalculator := CounterCalculator{
		Name:         name,
		CalculatorId: calcId,
		Type:         "calculator",
	}

	for _, opt := range options {
		opt(&counterCalculator)
	}

	return &counterCalculator
}

type CreateCounterCalculatorOption func(calculator *CounterCalculator)

func CCWithAllowFinishAfter(allowFinishAfter bool) CreateCounterCalculatorOption {
	return func(calculator *CounterCalculator) {
		calculator.AllowFinishAfter = allowFinishAfter
	}
}

func CCWithAllowStartBefore(allowStartBefore bool) CreateCounterCalculatorOption {
	return func(calculator *CounterCalculator) {
		calculator.AllowStartBefore = allowStartBefore
	}
}

func CCWithFields(fields []string) CreateCounterCalculatorOption {
	return func(calculator *CounterCalculator) {
		calculator.Fields = fields
	}
}

func CCWithMethod(method string) CreateCounterCalculatorOption {
	return func(calculator *CounterCalculator) {
		calculator.Method = method
	}
}

func CCWithValidateInterval(validateInterval string) CreateCounterCalculatorOption {
	return func(calculator *CounterCalculator) {
		calculator.ValidateInterval = validateInterval
	}
}

type CounterAccumulator struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Counter string `json:"counter"`

	ResetExpression string `json:"reset_expression,omitempty"`
	ResetInterval   string `json:"reset_interval,omitempty"`
}

func NewCounterAccumulator(name string, counter string, options ...CreateCounterAccumulatorOption) *CounterAccumulator {
	counterAccumulator := CounterAccumulator{
		Name:    name,
		Counter: counter,
		Type:    "accumulator",
	}

	for _, opt := range options {
		opt(&counterAccumulator)
	}

	return &counterAccumulator
}

type CreateCounterAccumulatorOption func(accumulator *CounterAccumulator)

func CAWithResetExpression(resetExpression string) CreateCounterAccumulatorOption {
	return func(accumulator *CounterAccumulator) {
		accumulator.ResetExpression = resetExpression
	}
}

func CAWithResetInterval(resetInterval string) CreateCounterAccumulatorOption {
	return func(accumulator *CounterAccumulator) {
		accumulator.ResetInterval = resetInterval
	}
}

func unmarshalCounters(rawCounters []json.RawMessage) ([]Counter, error) {
	var result []Counter

	for _, rawCounter := range rawCounters {
		counter, err := unmarshalCounter(rawCounter)

		if err != nil {
			return nil, err
		}

		result = append(result, counter)
	}

	return result, nil
}

func unmarshalCounter(raw json.RawMessage) (Counter, error) {
	var counterType struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(raw, &counterType); err != nil {
		return nil, err
	}

	var result Counter
	var err error

	switch counterType.Type {
	case "expression":
		var counterExpression CounterExpression

		if err = json.Unmarshal(raw, &counterExpression); err == nil {
			result = &counterExpression
		}
	case "dataset":
		var counterDataset CounterDataset

		if err = json.Unmarshal(raw, &counterDataset); err == nil {
			result = &counterDataset
		}
	case "route":
		var counterRoute CounterRoute

		if err = json.Unmarshal(raw, &counterRoute); err == nil {
			result = &counterRoute
		}
	case "datetime":
		var counterDatetime CounterDatetime

		if err = json.Unmarshal(raw, &counterDatetime); err == nil {
			result = &counterDatetime
		}
	case "parameter":
		var counterParameter CounterParameter

		if err = json.Unmarshal(raw, &counterParameter); err == nil {
			result = &counterParameter
		}
	case "message":
		var counterMessage CounterMessage

		if err = json.Unmarshal(raw, &counterMessage); err == nil {
			result = &counterMessage
		}
	case "interval":
		var counterInterval CounterInterval

		if err = json.Unmarshal(raw, &counterInterval); err == nil {
			result = &counterInterval
		}
	case "active":
		var counterActive CounterActive

		if err = json.Unmarshal(raw, &counterActive); err == nil {
			result = &counterActive
		}
	case "geofence":
		var counterGeofence CounterGeofence

		if err = json.Unmarshal(raw, &counterGeofence); err == nil {
			result = &counterGeofence
		}
	case "variable":
		var counterVariable CounterVariable

		if err = json.Unmarshal(raw, &counterVariable); err == nil {
			result = &counterVariable
		}
	case "specified":
		var specified struct {
			Name  string      `json:"name"`
			Value interface{} `json:"value"`
			Type  string      `json:"type"`
		}

		if err = json.Unmarshal(raw, &specified); err == nil {
			switch value := specified.Value.(type) {
			case string:
				result = NewCounterSpecifiedString(specified.Name, value)
			case float64:
				result = NewCounterSpecifiedNumber(specified.Name, value)
			case bool:
				result = NewCounterSpecifiedBoolean(specified.Name, value)
			}
		}
	case "calculator":
		var counterCalculator CounterCalculator

		if err = json.Unmarshal(raw, &counterCalculator); err == nil {
			result = &counterCalculator
		}
	case "accumulator":
		var counterAccumulator CounterAccumulator

		if err = json.Unmarshal(raw, &counterAccumulator); err == nil {
			result = &counterAccumulator
		}
	default:
		err = fmt.Errorf("unknown counter type: %s", counterType.Type)
	}

	return result, err
}
