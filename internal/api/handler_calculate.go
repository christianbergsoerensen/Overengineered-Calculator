package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
)

type Handler struct {
}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handlerCalculate(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Calculator API is running"}
	json.NewEncoder(w).Encode(resp)

	calc := calculator.NewCalculator()
	result, _ := calc.Calculate("add", 5, 3)
	fmt.Println(result)
}
