package test

import (
	"testing"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
)

func TestOperations(t *testing.T) {
	type testInfo struct {
		name      string
		op        calculator.Operation
		a, b      float64
		want      float64
		expectErr bool
	}

	tests := []testInfo{}
	tests = append(tests, testInfo{"Addition", calculator.AddOperation{}, 8, 5, 13, false})
	tests = append(tests, testInfo{"Subtraction", calculator.SubtractOperation{}, 34, 11, 23, false})
	tests = append(tests, testInfo{"Multiplication", calculator.MultiplyOperation{}, 4, 11, 44, false})
	tests = append(tests, testInfo{"Division", calculator.DivideOperation{}, 10, 2, 5, false})
	tests = append(tests, testInfo{"Division by zero", calculator.DivideOperation{}, 10, 0, 0, true})

	for _, testInfo := range tests {
		t.Run(testInfo.name, func(t *testing.T) {
			got, err := testInfo.op.Calculate(testInfo.a, testInfo.b)
			if (err != nil) && !testInfo.expectErr {
				t.Errorf("recieved error %v, when expecting %v", err, testInfo.expectErr)
			}

			if err == nil && testInfo.expectErr {
				t.Errorf("did not recieve error, when expecting %v", testInfo.expectErr)
			}

			if got != testInfo.want {
				t.Errorf("got %v, expected %v", got, testInfo.want)
			}
		})
	}
}
