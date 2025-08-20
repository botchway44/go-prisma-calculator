# --- Stage 1: The Builder ---
# This stage builds our Go application into a single binary.
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker's build cache.
# This step is only re-run when the dependencies change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the application. CGO_ENABLED=0 is important for creating a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./server ./cmd/server/main.go


# --- Stage 2: The Final Image ---
# This stage creates our lightweight final image.
FROM alpine:latest

WORKDIR /app

# Copy the compiled server binary from the 'builder' stage.
COPY --from=builder /app/server .

# Copy the .env file for configuration and the docs for Swagger.
COPY .env .
COPY docs ./docs

# Expose the ports our application uses.
EXPOSE 8080
EXPOSE 50051

# This is the command that will be run when the container starts.
CMD ["./server"]