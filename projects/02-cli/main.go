package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

func run(baseURL string) error {
	resp, err := http.Get(baseURL + "/items")
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var items []shared.Item
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	for _, item := range items {
		fmt.Printf("[%d] %s — $%.2f\n", item.ID, item.Name, float64(item.Price)/100)
	}
	return nil
}

func main() {
	url := flag.String("url", "http://localhost:8080", "base URL of the API server")
	flag.Parse()

	if err := run(*url); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
