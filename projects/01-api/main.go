package main

import (
	"encoding/json"
	"net/http"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

var items = []shared.Item{
	{ID: 1, Name: "Widget", Price: 999},
	{ID: 2, Name: "Gadget", Price: 1499},
	{ID: 3, Name: "Doohickey", Price: 299},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/items", itemsHandler)
	return mux
}

func main() {
	http.ListenAndServe(":8080", newMux())
}
