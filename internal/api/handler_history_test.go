package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// We do not care about the storage, we just want to ensure that the endpoint can be called correctly
func TestHistoryHandler(t *testing.T) {
	storage := &StubStorage{}
	handler := handlerHistory(storage)

	req, err := http.NewRequest(http.MethodGet, "/history", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", rec.Code)
	}
}
