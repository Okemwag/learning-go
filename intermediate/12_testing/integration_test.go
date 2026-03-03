//go:build integration

package testingdemo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchStatusCodeIntegration(t *testing.T) {
	// Using integration tests and build tags:
	// this test only runs when requested explicitly, for example:
	// go test -tags=integration ./intermediate/12_testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	client := http.Client{}
	got, err := FetchStatusCode(&client, server.URL)
	if err != nil {
		t.Fatalf("FetchStatusCode() error: %v", err)
	}
	if got != http.StatusAccepted {
		t.Fatalf("FetchStatusCode() = %d, want %d", got, http.StatusAccepted)
	}
}
