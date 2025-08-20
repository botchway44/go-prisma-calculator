# Get the directory of the downloaded googleapis protos from the git submodule.
GOOGLEAPIS_DIR := googleapis

.PHONY: all gen db-push run docker-up docker-down clean

# The default command, e.g., 'make' or 'make all'
all: gen

# Generate all necessary Go code (Prisma client and Protobuf stubs)
gen:
	@echo "==> Generating Prisma Go client..."
	@go run github.com/steebchen/prisma-client-go generate
	@echo "==> Generating protobuf files from submodule: ${GOOGLEAPIS_DIR}"
	@protoc \
	  -I proto \
	  -I ${GOOGLEAPIS_DIR} \
	  --go_out=. --go_opt=module=go-prisma-calculator \
	  --go-grpc_out=. --go-grpc_opt=module=go-prisma-calculator \
	  --openapiv2_out=./docs \
	  proto/*.proto
	@echo "==> Code generation complete."

# Push the Prisma schema to the database
db-push:
	@echo "==> Pushing Prisma schema to the database..."
	@go run github.com/steebchen/prisma-client-go db push

# Run the main application
run:
	@echo "==> Starting application..."
	@go run cmd/server/main.go

# Start the Docker containers (database)
docker-up:
	@echo "==> Starting Docker containers..."
	@docker-compose up -d

# Stop the Docker containers
docker-down:
	@echo "==> Stopping Docker containers..."
	@docker-compose down

# Clean up generated files
clean:
	@echo "==> Cleaning up generated files..."
	@rm -rf ./generated
	@rm -rf ./docs
	@rm -f ./app.log
	@echo "==> Cleanup complete."