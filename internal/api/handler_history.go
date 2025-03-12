package api

import (
	"encoding/json"
	"net/http"

	"github.com/christianbergsoerensen/Overengineered-Calculator/internal/storage"
)

func handlerHistory(store storage.StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerHistoryHelper(w, r, store)
	}
}

func handlerHistoryHelper(w http.ResponseWriter, r *http.Request, store storage.StorageInterface) {
	calcs, err := store.GetHistory()
	if err != nil {
		http.Error(w, "could not fetch history", 500)
		return
	}

	json.NewEncoder(w).Encode(calcs)

}
