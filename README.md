# inception-monorepo

A test monorepo for containerized Gas Town. Two projects built simultaneously by different polecats.

## Structure

```
projects/
  01-api/    REST API service (GET /items, GET /health)
  02-cli/    CLI tool that fetches from the API
shared/
  types.go   Shared Item type used by both projects
```
