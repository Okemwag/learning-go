package testingdemo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchStatusCodeWithHTTPTest(t *testing.T) {
	// Using httptest:
	// ResponseRecorder avoids binding a real port in unit tests.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	req := httptest.NewRequest(http.MethodGet, "http://example.test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("rec.Code = %d, want %d", rec.Code, http.StatusCreated)
	}
}
