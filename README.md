# Go Testing

## NOTE

---

- Remember that giving 100% coverage doesn't mean that it is a good metric to measure our testing quality (the result from testing can still be wrong)

## Testing Command

---

### Run every test starting from current directory

- `go test ./...`

### Run tests in specific folder

- `go test ./folder/...`

### Run tests with coverage

- `go test -cover`

### Run tests with coverage and show the result in HTML

- `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

### Run tests with benchmark

- `go test -bench=.`
