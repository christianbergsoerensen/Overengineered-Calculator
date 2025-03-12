package api

import (
	"encoding/json"
	"net/http"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
)

func handlerHistory(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerHistoryHelper(w, r, store)
	}
}

func handlerHistoryHelper(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	calcs, err := store.GetCalculations()
	if err != nil {
		http.Error(w, "could not fetch history", 500)
		return
	}

	json.NewEncoder(w).Encode(calcs)

}
