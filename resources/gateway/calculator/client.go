package flespi_calculator

import "github.com/mixser/flespi-client/internal/flespiapi"

// CalculatorClient provides receiver-based methods for managing Flespi calculators.
// Access it via Client.Calculators after creating a flespi.Client.
type CalculatorClient struct {
	c flespiapi.Doer
}

// NewCalculatorClient creates a CalculatorClient wrapping the given flespiapi.Doer.
func NewCalculatorClient(c flespiapi.Doer) *CalculatorClient {
	return &CalculatorClient{c: c}
}

func (cc *CalculatorClient) Create(name string, options ...CreateCalculatorOption) (*Calculator, error) {
	return NewCalculator(cc.c, name, options...)
}

func (cc *CalculatorClient) List() ([]Calculator, error) {
	return ListCalculators(cc.c)
}

func (cc *CalculatorClient) Get(calculatorId int64) (*Calculator, error) {
	return GetCalculator(cc.c, calculatorId)
}

func (cc *CalculatorClient) Update(calc Calculator) (*Calculator, error) {
	return UpdateCalculator(cc.c, calc)
}

func (cc *CalculatorClient) Delete(calc Calculator) error {
	return DeleteCalculator(cc.c, calc)
}

func (cc *CalculatorClient) DeleteById(calculatorId int64) error {
	return DeleteCalculatorById(cc.c, calculatorId)
}
