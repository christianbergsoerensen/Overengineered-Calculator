package api

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	handler := newHandler()

	r.Post("/calculate", handler.handlerCalculate)
	r.Get("/history", handler.handlerHistory)

	return r
}
