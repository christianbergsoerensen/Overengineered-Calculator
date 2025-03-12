package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
	"github.com/stretchr/testify/assert"
)

// We do not care about the calculator/storing logic here
type StubCalculator struct{}

func (s *StubCalculator) Calculate(expression string, a, b float64) (float64, error) {
	return 4.0, nil
}

type StubStorage struct{}

func (s *StubStorage) SaveCalculation(calc storage.CalculationResult) error {
	return nil
}

func (s *StubStorage) GetHistory() ([]storage.CalculationResult, error) {
	return []storage.CalculationResult{}, nil
}

type FailingStorage struct{}

func (s *FailingStorage) SaveCalculation(calc storage.CalculationResult) error {
	return errors.New("database error")
}

func (s *FailingStorage) GetHistory() ([]storage.CalculationResult, error) {
	return []storage.CalculationResult{}, nil
}

func TestCalculateHandler(t *testing.T) {
	handler := handlerCalculate(&StubCalculator{}, &StubStorage{})

	t.Run("Valid input returns correct result", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(`{"add": "2+2"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"result":4}`, rec.Body.String())
	})

	t.Run("Invalid input returns 400 error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(`{"add": 2+2`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("Server error returns 500 error", func(t *testing.T) {
		handler := handlerCalculate(&StubCalculator{}, &FailingStorage{})

		req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(`{"add": "2+2"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, 500, rec.Code)
	})
}
