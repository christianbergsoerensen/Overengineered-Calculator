package api

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) calculateHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Calculator API is running"}
	json.NewEncoder(w).Encode(resp)
}
