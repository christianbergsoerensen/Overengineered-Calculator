package calculator

import "errors"

type Operation interface {
	Calculate(a, b float64) (float64, error)
}

type AddOperation struct{}

func (ao AddOperation) Calculate(a, b float64) (float64, error) {
	return a + b, nil
}

type SubtractOperation struct{}

func (so SubtractOperation) Calculate(a, b float64) (float64, error) {
	return a - b, nil
}

type MultiplyOperation struct{}

func (so MultiplyOperation) Calculate(a, b float64) (float64, error) {
	return a * b, nil
}

type DivideOperation struct{}

func (so DivideOperation) Calculate(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0 is not allowed")
	}

	return a / b, nil
}
