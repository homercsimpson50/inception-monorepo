package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

func TestRun(t *testing.T) {
	mockItems := []shared.Item{
		{ID: 1, Name: "Widget", Price: 999},
		{ID: 2, Name: "Gadget", Price: 1499},
		{ID: 3, Name: "Doohickey", Price: 299},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/items" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockItems)
	}))
	defer srv.Close()

	if err := run(srv.URL); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
}

func TestRunServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not json"))
	}))
	defer srv.Close()

	if err := run(srv.URL); err == nil {
		t.Fatal("expected error for invalid JSON response, got nil")
	}
}
