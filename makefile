APP_NAME=blog-crud-api
CMD_DIR=cmd/$(APP_NAME)
PORT?=3001
CACHE_CAPACITY?=30

.PHONY: run build clean test mockgen deps

# Run the application
run:
	go run $(CMD_DIR)/main.go --port=$(PORT) --cache-capacity=$(CACHE_CAPACITY)

# Build the application
build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)/main.go

# Run tests
test:
	go test -cover ./... -v

# Install dependencies
deps:
	go mod tidy
	go mod download


# Clean up binaries and cache
clean:
	rm -rf bin/ $(APP_NAME)
	go clean

# Generate mocks using mockgen
mockgen:
	mockgen -source=internal/business/blogbus/blogbus.go -destination=internal/mock/business/blogbus/blogbus.go -package=mockblogbus

# Generate docs
swag:
	swag init -g internal/api/http/blogapp/route.go -o ./docs
	swag fmt

build-docker:
	docker build -t anazibinurasheed/$(APP_NAME):latest .