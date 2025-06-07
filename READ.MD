# Intro
- TODO

# Install/ Update
## Update
- `make update`

# Testing
## Run all tests in the current directory
`go test`

## Run tests with verbose output
`go test -v`

## Run a specific test
`go test -run TestXPCommand`

## Run tests in all subdirectories
`go test ./...`

## Run tests with coverage
`go test -cover`

## Run tests with detailed coverage report
`go test -coverprofile=coverage.out`
`go tool cover -html=coverage.out`

# Other things