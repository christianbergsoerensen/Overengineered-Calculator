package storage

import "time"

type CalculationResult struct {
	ID        int       `json:"id"`
	Operation string    `json:"operation"`
	A         float64   `json:"a"`
	B         float64   `json:"b"`
	Result    float64   `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

type StorageInterface interface {
	SaveCalculation(result CalculationResult) error
	GetHistory() ([]CalculationResult, error)
}
