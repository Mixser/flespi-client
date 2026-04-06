package flespi_calculator

import (
	"fmt"

	"github.com/mixser/flespi-client/internal/flespiapi"
)

func NewCalculator(client flespiapi.APIRequester, name string, options ...CreateCalculatorOption) (*Calculator, error) {
	calc := Calculator{
		Name: name,
	}

	for _, opt := range options {
		opt(&calc)
	}

	var headers map[string]string
	if calc.AccountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", calc.AccountId),
		}
	}

	accountId := calc.AccountId
	calc.AccountId = 0
	defer func() { calc.AccountId = accountId }()

	response := calculatorsResponse{}

	if err := client.RequestAPIWithHeaders("POST", "gw/calcs", headers, []Calculator{calc}, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func ListCalculators(client flespiapi.APIRequester) ([]Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", "gw/calcs/all", nil, &response); err != nil {
		return nil, err
	}

	return response.Calculators, nil
}

func GetCalculator(client flespiapi.APIRequester, calculatorId int64) (*Calculator, error) {
	response := calculatorsResponse{}

	if err := client.RequestAPI("GET", fmt.Sprintf("gw/calcs/%d?fields=id,name,messages_source,update_period,update_delay,update_onchange,intervals_ttl,intervals_rotate,selectors,counters,validate_interval,validate_message,timezone,metadata,cid", calculatorId), nil, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func UpdateCalculator(client flespiapi.APIRequester, calc Calculator) (*Calculator, error) {
	response := calculatorsResponse{}

	calculatorId := calc.Id
	accountId := calc.AccountId

	calc.Id = 0
	calc.AccountId = 0

	defer func() {
		calc.Id = calculatorId
		calc.AccountId = accountId
	}()

	var headers map[string]string
	if accountId != 0 {
		headers = map[string]string{
			"x-flespi-cid": fmt.Sprintf("%d", accountId),
		}
	}

	if err := client.RequestAPIWithHeaders("PUT", fmt.Sprintf("gw/calcs/%d", calculatorId), headers, calc, &response); err != nil {
		return nil, err
	}

	return &response.Calculators[0], nil
}

func DeleteCalculator(client flespiapi.APIRequester, calc Calculator) error {
	if calc.Id == 0 {
		return fmt.Errorf("calculator id is not set")
	}

	return DeleteCalculatorById(client, calc.Id)
}

func DeleteCalculatorById(client flespiapi.APIRequester, calculatorId int64) error {
	return client.RequestAPI("DELETE", fmt.Sprintf("gw/calcs/%d", calculatorId), nil, nil)
}
