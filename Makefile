.PHONY: all build test clean generate dev test test-unit test-coverage

# Go related variables
BINARY_NAME=abierr
MAIN_PACKAGE=./example/main.go


# Development server with hot reload
dev:
	go run $(MAIN_PACKAGE)

# Build the application
build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

# Clean build files
clean:
	@echo "Cleaning..."
	rm -rf bin/


# Generate smart contract bindings
gen-contract:
	@echo "Generating smart contract bindings..."
	abigen --abi example/erc20/abi/ERC20Owner.json --pkg contracts --type Erc20Owner --out example/erc20/contracts/erc20_owner_contract.go
	abigen --abi example/erc20/abi/ERC20.json --pkg contracts --type Erc20 --out example/erc20/contracts/erc20_contract.go
	abigen --abi example/optimex/abi/EVMBTC.json --pkg contracts --type EvmBtc --out example/optimex/contracts/evmbtc_contract.go

# Linting and formatting
lint:
	golangci-lint run

fmt:
	go fmt ./...

# Dev environment setup
setup: install-tools
	@echo "Setting up development environment..."
	cp .env.example .env
	go mod download
	go mod tidy

# Run all tests
test:
	go test   ./...

# Run tests with coverage
test-coverage:
	go test  -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Help
help:
	@echo "Available commands:"
	@echo "Development:"
	@echo "  make setup              - Install tools and setup development environment"
	@echo "  make dev               - Run development server with hot reload"
	@echo "  make build             - Build the application"
	@echo "  make fmt               - Format code"
	@echo "  make lint              - Run linter"
	@echo ""
	@echo "Testing:"
	@echo "  make test              - Run all tests with race detection"
	@echo "  make test-coverage     - Run tests with coverage report"
	@echo ""
	@echo "Code Generation:"
	@echo "  make gen-contract      - Generate smart contract bindings"
