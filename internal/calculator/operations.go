package calculator

import (
	"errors"
	"math"
)

type Operation interface {
	Calculate(a, b float64) (float64, error)
}

type AddOperation struct{}

func (op AddOperation) Calculate(a, b float64) (float64, error) {
	result := a + b
	if err := CheckOverflow(result); err != nil {
		return 0, err
	}
	return result, nil
}

type SubtractOperation struct{}

func (so SubtractOperation) Calculate(a, b float64) (float64, error) {
	result := a - b
	if err := CheckOverflow(result); err != nil {
		return 0, err
	}
	return result, nil
}

type MultiplyOperation struct{}

func (mo MultiplyOperation) Calculate(a, b float64) (float64, error) {
	result := a * b
	if err := CheckOverflow(result); err != nil {
		return 0, err
	}
	return result, nil
}

type DivideOperation struct{}

func (do DivideOperation) Calculate(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0 is not allowed")
	}

	result := a / b
	if err := CheckOverflow(result); err != nil {
		return 0, err
	}
	return result, nil
}

func CheckOverflow(result float64) error {
	if result > math.MaxFloat64 || result < -math.MaxFloat64 {
		return errors.New("overflow error")
	}
	return nil
}
