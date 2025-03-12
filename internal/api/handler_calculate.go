package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/calculator"
	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
)

type CalculateRequest struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func handlerCalculate(calc calculator.CalculatorInterface, store storage.StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCalculateHelper(w, r, calc, store)
	}
}

func handleCalculateHelper(w http.ResponseWriter, r *http.Request, calc calculator.CalculatorInterface, store storage.StorageInterface) {
	req := CalculateRequest{}
	//Expecting r.body to have keys operation , a and b, which then can be decoded into a CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	res, err := calc.Calculate(req.Operation, req.A, req.B)
	//Only division by 0 error at the moment
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//By now we know that the calc is successful so we save it to storage

	err = store.SaveCalculation(storage.CalculationResult{
		Operation: req.Operation,
		A:         req.A,
		B:         req.B,
		Result:    res,
		Timestamp: time.Now(),
	})
	if err != nil {
		http.Error(w, "server failed to save the result", 500)
		return
	}

	json.NewEncoder(w).Encode(CalculateResponse{Result: res})

}
