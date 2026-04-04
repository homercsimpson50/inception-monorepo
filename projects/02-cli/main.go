package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/homercsimpson50/inception-monorepo/shared"
)

func run(baseURL, format string) error {
	if format != "table" && format != "json" {
		return fmt.Errorf("unsupported format: %q (must be \"table\" or \"json\")", format)
	}

	resp, err := http.Get(baseURL + "/items")
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var items []shared.Item
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	switch format {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(items); err != nil {
			return fmt.Errorf("json encode failed: %w", err)
		}
	default:
		for _, item := range items {
			fmt.Printf("[%d] %s — $%.2f\n", item.ID, item.Name, float64(item.Price)/100)
		}
	}
	return nil
}

func main() {
	url := flag.String("url", "http://localhost:8080", "base URL of the API server")
	format := flag.String("format", "table", "output format: table or json")
	flag.Parse()

	if err := run(*url, *format); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
