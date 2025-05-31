.PHONY: run build test

# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=urlshortener
BINARY_PATH=bin/$(BINARY_NAME)

# Build application
build:
	@echo "Building..."
	@$(GOBUILD) -o $(BINARY_PATH) -v ./cmd/server

# Run application
run:
	@echo "Running..."
	@$(GOCMD) run ./cmd/server

# Run w/ hot reload
dev:
	@air

# Test
test:
	@echo "Testing..."
	@$(GOTEST) -v ./...

# Test w/ coverage
	@echo "Testing with coverage..."
	@$(GOTEST) -v -cover -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf bin/
	@rm -f coverage.out coverage.html

# Run linter (requires golangci-lint)
lint:
	@echo "Linting..."
	@golangci-lint run

# Format code
fmt:
	@echo "Formatting..."
	@$(GOCMD) fmt ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@$(GOMOD) download

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@$(GOMOD) tidy

# Docker commands
docker-build:
	@echo "Building Docker image..."
	@docker build -t urlshortener:latest -f docker/Dockerfile .

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-logs:
	@docker-compose logs -f

# Db commands
migrate-up:
	@echo "Running migrations..."
	@migrate -path ./migrations -database "$(DATABASE_URL)" up

migrate-down:
	@echo "Rolling back migrations..."
	@migrate -path ./migrations -database "$(DATABASE_URL)" down

migrate-create:
	@echo "Creating migration..."
	@migrate create -ext sql -dir ./migrations -seq $(name)
