package api

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) historyHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "History API is running"}
	json.NewEncoder(w).Encode(resp)
}
