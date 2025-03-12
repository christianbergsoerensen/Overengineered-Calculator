package calculator

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	type testInfo struct {
		name      string
		op        string
		a, b      float64
		want      float64
		expectErr bool
	}

	tests := []testInfo{}
	tests = append(tests, testInfo{"Valid Addition", "add", 8, 5, 13, false})
	tests = append(tests, testInfo{"Valid Subtraction", "subtract", 34, 11, 23, false})
	tests = append(tests, testInfo{"Valid Multiplication", "multiply", 4, 11, 44, false})
	tests = append(tests, testInfo{"Overflow Calculation", "multiply", 1e155, 1e154, 0, true})
	tests = append(tests, testInfo{"Valid Division", "divide", 10, 2, 5, false})
	tests = append(tests, testInfo{"Invalid Division", "divide", 10, 0, 0, true})
	tests = append(tests, testInfo{"Invalid Operation", "log", 8, 2, 0, true})

	calc := NewCalculator()

	for _, testInfo := range tests {
		t.Run(testInfo.name, func(t *testing.T) {
			got, err := calc.Calculate(testInfo.op, testInfo.a, testInfo.b)
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
