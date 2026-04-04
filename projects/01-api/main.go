package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/items", itemsHandler)
	mux.HandleFunc("GET /items/{id}", itemHandler)
	return mux
}

func main() {
	http.ListenAndServe(":8080", newMux())
}
