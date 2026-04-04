# 02-cli

CLI tool. Build a Go command-line tool that:
- Takes an optional `--url` flag (default `http://localhost:8080`)
- Fetches `GET /items` from the API
- Prints each item as: `[id] name — $price`
- Includes tests using httptest
