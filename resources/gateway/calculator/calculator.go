package flespi_calculator

import (
	"fmt"

	"github.com/mixser/flespi-client"
)

// TODO: add NewCalculator call

func ListCalculators(client *flespi.Client) ([]Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", "gw/calcs/all", nil, &response); err != nil {
		return nil, err
	}

	return response.Calculators, nil
}

func GetCalculator(client *flespi.Client, calculatorId int64) (*Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", fmt.Sprintf("gw/calcs/%d", calculatorId), nil, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func UpdateCalculator(client *flespi.Client, calc Calculator) (*Calculator, error) {
	response := calculatorsResponse{}

	calculatorId := calc.Id
	calc.Id = 0

	if err := client.RequestAPI("PUT", fmt.Sprintf("gw/calcs/%d", calculatorId), calc, &response); err != nil {
		return nil, err
	}

	calc.Id = calculatorId

	return &response.Calculators[0], nil
}

func DeleteCalculator(client *flespi.Client, calc Calculator) error {
	if calc.Id == 0 {
		return fmt.Errorf("calculator id is not set")
	}

	return DeleteCalculatorById(client, calc.Id)
}

func DeleteCalculatorById(client *flespi.Client, calculatorId int64) error {
	return client.RequestAPI("DELETE", fmt.Sprintf("gw/calcs/%d", calculatorId), nil, nil)
}
