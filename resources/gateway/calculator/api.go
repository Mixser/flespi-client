package flespi_calculator

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewCalculator(client flespiapi.Doer, name string, options ...CreateCalculatorOption) (*Calculator, error) {
	calc := Calculator{
		Name: name,
	}

	for _, opt := range options {
		opt(&calc)
	}

	response := calculatorsResponse{}

	if err := client.RequestAPI("POST", "gw/calcs", []Calculator{calc}, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func ListCalculators(client flespiapi.Doer) ([]Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", "gw/calcs/all", nil, &response); err != nil {
		return nil, err
	}

	return response.Calculators, nil
}

func GetCalculator(client flespiapi.Doer, calculatorId int64) (*Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", fmt.Sprintf("gw/calcs/%d", calculatorId), nil, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func UpdateCalculator(client flespiapi.Doer, calc Calculator) (*Calculator, error) {
	response := calculatorsResponse{}

	calculatorId := calc.Id
	calc.Id = 0

	if err := client.RequestAPI("PUT", fmt.Sprintf("gw/calcs/%d", calculatorId), calc, &response); err != nil {
		return nil, err
	}

	calc.Id = calculatorId

	return &response.Calculators[0], nil
}

func DeleteCalculator(client flespiapi.Doer, calc Calculator) error {
	if calc.Id == 0 {
		return fmt.Errorf("calculator id is not set")
	}

	return DeleteCalculatorById(client, calc.Id)
}

func DeleteCalculatorById(client flespiapi.Doer, calculatorId int64) error {
	return client.RequestAPI("DELETE", fmt.Sprintf("gw/calcs/%d", calculatorId), nil, nil)
}
