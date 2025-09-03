APP_NAME=jwt_and_rest_api

build:
	@echo "Building the project..."
	go build -o $(APP_NAME) ./cmd/main.go

run:
run:
	@echo "Running the application..."
	./jwt_and_rest_api

clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(APP_NAME)

documentation:
	@echo "Generating Docs..."
	swag init -g cmd/main.go -o docs


help:
	@echo "Available commands:"
	@echo "  build        Build the Go application"
	@echo "  run          Run the Go application"
	@echo "  docs          Generate swagger docs"
	@echo "  clean        Clean build artifacts"