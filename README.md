# Go Testing

## NOTE

---

- Remember that giving 100% coverage doesn't mean that it is a good metric to measure our testing quality (the result from testing can still be wrong)

- There are 3 type of testing

1. Unit testing (white box testing)
2. Integration testing (black box testing)
3. Functional testing (black box testing)

Integration testing is when you test more than one component and how they function together. For instance, how another system interacts with your system, or the database interacts with your data abstraction layer. Usually, this requires a fully installed system, although in its purest forms it does not.

Functional testing is when you test the system against the functional requirements of the product. Product/Project management usually writes these up and QA formalizes the process of what a user should see and experience, and what the end result of those processes should be. Depending on the product, this can be automated or not.

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
