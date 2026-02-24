BINARY_NAME=img2webp

.PHONY: all build clean run check test

all: check build zip

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME).exe ./cmd/img2webp

zip:
	@echo "Zipping $(BINARY_NAME)..."
	powershell -Command "Compress-Archive -Path bin/$(BINARY_NAME).exe -DestinationPath bin/$(BINARY_NAME).zip -Force"

run:
	@echo "Running $(BINARY_NAME)..."
	go run ./cmd/img2webp

clean:
	@echo "Cleaning up..."
	go clean
	rm -f bin/$(BINARY_NAME).exe

check:
	@echo "Formatting and returning go dependencies..."
	go mod tidy
	go fmt ./...
	go vet ./...

test:
	@echo "Running tests..."
	go test ./...

help:
	@echo "Available commands:"
	@echo "  all: Build and zip the executable"
	@echo "  build: Build the executable"
	@echo "  zip: Zip the executable"
	@echo "  run: Run the executable"
	@echo "  clean: Clean the build"
	@echo "  check: Format and vet the code"
	@echo "  test: Run tests"
