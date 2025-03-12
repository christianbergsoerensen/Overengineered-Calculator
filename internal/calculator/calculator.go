package calculator

import (
	"errors"
)

type CalculatorInterface interface {
	Calculate(operation string, a, b float64) (float64, error)
}

type Calculator struct {
	Operations map[string]Operation
}

func NewCalculator() *Calculator {

	ops := make(map[string]Operation)
	ops["add"] = AddOperation{}
	ops["subtract"] = SubtractOperation{}
	ops["multiply"] = MultiplyOperation{}
	ops["divide"] = DivideOperation{}

	calc := &Calculator{ops}
	return calc
}

func (c *Calculator) Calculate(operation string, a, b float64) (float64, error) {
	op, ok := c.Operations[operation]
	if !ok {
		return 0, errors.New("operation does not exist")
	}

	result, err := op.Calculate(a, b)
	if err != nil {
		return 0, err
	}

	return result, nil

}
