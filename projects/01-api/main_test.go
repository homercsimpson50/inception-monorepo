package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

func TestHealthEndpoint(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("expected status ok, got %q", body["status"])
	}
}

func TestItemsEndpoint(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var got []shared.Item
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(got) != 3 {
		t.Errorf("expected 3 items, got %d", len(got))
	}
}

func TestItemByIDEndpoint(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items/1", nil)
	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var got shared.Item
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}
	if got.Name != "Widget" {
		t.Errorf("expected name Widget, got %q", got.Name)
	}
}

func TestItemByIDNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items/999", nil)
	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}

func TestItemByIDInvalidID(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items/abc", nil)
	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}
