package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

func mockServer(items []shared.Item) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/items" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}))
}

var testItems = []shared.Item{
	{ID: 1, Name: "Widget", Price: 999},
	{ID: 2, Name: "Gadget", Price: 1499},
	{ID: 3, Name: "Doohickey", Price: 299},
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestRunTableFormat(t *testing.T) {
	srv := mockServer(testItems)
	defer srv.Close()

	output := captureStdout(t, func() {
		if err := run(srv.URL, "table"); err != nil {
			t.Fatalf("run returned error: %v", err)
		}
	})

	for _, item := range testItems {
		if !strings.Contains(output, item.Name) {
			t.Errorf("table output missing item name %q", item.Name)
		}
	}
	if !strings.Contains(output, "[1]") {
		t.Error("table output missing item ID prefix")
	}
	if !strings.Contains(output, "$9.99") {
		t.Error("table output missing formatted price")
	}
}

func TestRunJSONFormat(t *testing.T) {
	srv := mockServer(testItems)
	defer srv.Close()

	output := captureStdout(t, func() {
		if err := run(srv.URL, "json"); err != nil {
			t.Fatalf("run returned error: %v", err)
		}
	})

	var got []shared.Item
	if err := json.Unmarshal([]byte(output), &got); err != nil {
		t.Fatalf("json output not valid JSON: %v\noutput: %s", err, output)
	}
	if len(got) != len(testItems) {
		t.Fatalf("expected %d items, got %d", len(testItems), len(got))
	}
	for i, item := range got {
		if item != testItems[i] {
			t.Errorf("item %d: got %+v, want %+v", i, item, testItems[i])
		}
	}
}

func TestRunDefaultIsTable(t *testing.T) {
	srv := mockServer(testItems)
	defer srv.Close()

	output := captureStdout(t, func() {
		if err := run(srv.URL, "table"); err != nil {
			t.Fatalf("run returned error: %v", err)
		}
	})

	// Default (table) should NOT be valid JSON
	var items []shared.Item
	if err := json.Unmarshal([]byte(output), &items); err == nil {
		t.Error("table output should not be valid JSON")
	}
}

func TestRunInvalidFormat(t *testing.T) {
	srv := mockServer(testItems)
	defer srv.Close()

	err := run(srv.URL, "csv")
	if err == nil {
		t.Fatal("expected error for invalid format, got nil")
	}
	if !strings.Contains(err.Error(), "unsupported format") {
		t.Errorf("error should mention unsupported format, got: %v", err)
	}
}

func TestRunServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not json"))
	}))
	defer srv.Close()

	if err := run(srv.URL, "table"); err == nil {
		t.Fatal("expected error for invalid JSON response, got nil")
	}
}
