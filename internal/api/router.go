package api

import (
	"fmt"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
	"github.com/go-chi/chi/v5"
)

func NewRouter(calc *calculator.Calculator, store storage.Storage) *chi.Mux {
	r := chi.NewRouter()
	fmt.Println(calc)
	r.Post("/calculate", handlerCalculate(calc, store))
	r.Get("/history", handlerHistory(store))

	return r
}
