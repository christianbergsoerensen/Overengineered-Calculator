package api

import (
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
	"github.com/go-chi/chi/v5"
)

func NewRouter(calc calculator.CalculatorInterface, store storage.StorageInterface) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/calculate", handlerCalculate(calc, store))
	r.Get("/history", handlerHistory(store))

	return r
}
