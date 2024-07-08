# Project-specific variables
BINARY_NAME=api-culinary-review
MAIN_PACKAGE=cmd/api/main.go

# Define phony targets
.PHONY: all build run clean test help

# Build the Go application
build:
	@echo "Building the application..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

# Run the application
run:
	@echo "Running the application..."
	go run $(MAIN_PACKAGE)

# Tidy a Module
tidy:
	@echo "Tidying the application..."
	go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f bin/$(BINARY_NAME)

#  Init Swagger
swagger:
	@echo "Generating swagger docs..."
	swag init -g cmd/api/main.go

# 
# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Display help
help:
	@echo "Makefile commands:"
	@echo "  make          - Build the application"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make test     - Run tests"
	@echo "  make help     - Display this help message"
