# Makefile for Palindromic Fuel Calculator

.PHONY: build test fmt vet clean run web help

# Build the binary
build:
	go build -o palindromic-fuel main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-cover:
	go test -cover ./...

# Run benchmarks
bench:
	go test -bench=. -benchmem

# Format code
fmt:
	gofmt -w main.go main_test.go

# Check formatting
fmt-check:
	gofmt -d main.go main_test.go

# Vet code
vet:
	go vet main.go

# Clean build artifacts
clean:
	rm -f palindromic-fuel *.csv

# Run the program (example)
run:
	./palindromic-fuel -price=128.9 -max=100

# Start web server
web:
	./palindromic-fuel -web

# Install dependencies (if any)
deps:
	go mod tidy

# Full check: fmt, vet, test
check: fmt vet test

# Help
help:
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  test-cover - Run tests with coverage"
	@echo "  bench      - Run benchmarks"
	@echo "  fmt        - Format code"
	@echo "  fmt-check  - Check code formatting"
	@echo "  vet        - Vet code"
	@echo "  clean      - Clean build artifacts"
	@echo "  run        - Run example command"
	@echo "  web        - Start web server"
	@echo "  deps       - Tidy dependencies"
	@echo "  check      - Run fmt, vet, and test"
	@echo "  help       - Show this help"