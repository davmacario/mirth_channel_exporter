BINARY_NAME := mirth-channel-exporter
BUILD_DIR := build
BINARY_PATH := ./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: all build clean deps deps-update run fmt

build:
	@echo "Building $(BINARY_PATH)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_PATH) ./main.go

clean:
	@echo "Cleaning..."
	go clean
	@rm -rf ./$(BUILD_DIR)

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

deps-update:
	@echo "Updating dependencies..."
	go mod tidy
	go get -u ./...

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_PATH) $(input_path) $(output_path)

fmt:
	@echo "Formatting code..."
	go fmt ./...

container:
	@echo "Building container image..."
	docker build -t $(BINARY_NAME):latest .
