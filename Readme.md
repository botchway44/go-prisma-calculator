

# Calculator Service

[](https://golang.org)
[](https://opensource.org/licenses/MIT)

A sample Go microservice that demonstrates Hexagonal Architecture (Ports & Adapters) for a simple Calculator application. This project uses gRPC and a RESTful API (via Gin) as its entry points, Prisma as its ORM for a PostgreSQL database, and Uber's `fx` for dependency injection.

This project is designed as a learning tool to showcase a clean, testable, and maintainable application structure using modern Go tooling.

-----

## 🏛️ Architecture Overview

This project is built using the **Hexagonal (Ports & Adapters) Architecture**. The core principle is to isolate the application's central business logic from external concerns like databases, APIs, and frameworks.

This is achieved by structuring the application into three distinct layers:

1.  **Domain**: The heart of the application. It contains the pure business logic and rules and defines the contracts (**ports**) for how it interacts with the outside world. It has zero knowledge of the database or gRPC.
2.  **Application**: This layer orchestrates the business logic to fulfill specific use cases. It implements the "inbound" ports defined by the domain.
3.  **Infrastructure**: This is the outermost layer. It contains all the external details: gRPC and REST handlers, the Prisma database repository, configuration management, etc.

**Dependency Rule:** Dependencies only ever point inward: **Infrastructure → Application → Domain**.

-----

## 📂 Folder Structure

The project follows a clean, layered structure to enforce the architectural principles.

```
.
├── cmd/server/             # Main application entry point.
├── docs/                   # Generated OpenAPI (Swagger) documentation.
├── generated/              # Generated Go code from Protobuf files.
├── googleapis/             # Git submodule for Google API protos.
├── internal/
│   ├── domain/             # Core business logic and rules.
│   ├── application/        # Use case implementations.
│   └── infrastructure/     # All external concerns (adapters, config, etc.).
├── proto/                  # Source .proto files.
├── prisma/                 # Prisma schema and migration files.
├── .env                    # Local environment variables (ignored by git).
├── Taskfile.yml            # Modern task runner for automating commands.
├── Dockerfile              # For containerizing the application.
└── docker-compose.yml      # For running the entire stack with Docker.
```

-----

## 🚀 Getting Started

You can run this project in two ways: locally for development or fully containerized with Docker.

### Prerequisites

  * [Go](https://golang.org/doc/install) (version 1.25 or higher)
  * [Docker](https://www.docker.com/products/docker-desktop) and Docker Compose
  * [protoc](https://grpc.io/docs/protoc-installation/) (the Protobuf compiler)
  * [Task](https://taskfile.dev/installation/) (the recommended task runner)

### Method 1: Running Locally (Recommended for Development)

1.  **Clone the Repository with Submodules**
    It's crucial to use the `--recurse-submodules` flag to download the `googleapis` dependency.

    ```bash
    git clone --recurse-submodules https://github.com/botchway44/go-prisma-calculator.git
    cd go-prisma-calculator
    ```

2.  **Set Up Environment Variables**
    Create a `.env` file for your local environment variables.

    ```bash
    cp .env.example .env
    ```

3.  **Start the Database**
    This command uses Docker Compose to start a PostgreSQL container. You can setup a local postgres server and connect to.

    ```bash
    task db:up
    ```

4.  **Install Go Dependencies & Tools**
    This will download all the necessary Go modules and development tools.

    ```bash
    go mod tidy
    go install github.com/air-verse/air@latest
    ```

5.  **Push the Database Schema**
    This command syncs your Prisma schema with the running database.

    ```bash
    task db:push
    ```

6.  **Generate All Code**
    This single command generates both the Prisma client and all Protobuf-related code.

    ```bash
    task gen
    ```

7.  **Run the Server with Live-Reloading**

    ```bash
    task dev
    ```

### Method 2: Running Everything with Docker (Production-like)

This method builds the application into a container and runs both the app and the database with a single command.

1.  **Clone the Repository with Submodules**
    ```bash
    git clone --recurse-submodules https://github.com/botchway44/go-prisma-calculator.git
    cd go-prisma-calculator
    ```
2.  **Set Up Environment Variables**
    ```bash
    cp .env.example .env
    ```
3.  **Run the Entire Stack**
    This command will build the application's Docker image and start all services.
    ```bash
    docker-compose up --build
    ```

-----

## 💻 Common Commands

The `Taskfile.yml` provides several useful commands for the local development workflow:

  * `task dev`: Run the server with live-reloading via Air.
  * `task run`: Run the application once without live-reloading.
  * `task gen`: Generate all Go code (Prisma & Protobuf).
  * `task db:push`: Apply schema changes to the database.
  * `task db:up`: Start the Docker database.
  * `task db:down`: Stop the Docker database.

-----

## 📖 API & Documentation

The application exposes both a gRPC and a RESTful API.

  * **gRPC server** is available on `:50051`
  * **REST (Gin) server** is available on `:8080`

### REST API Docs (Swagger)

Once the application is running, you can access the interactive Swagger UI in your browser at:
**[http://localhost:8080/swagger](https://www.google.com/search?q=http://localhost:8080/swagger)**

### Example REST Request

You can test the `add` endpoint using `curl`:

```bash
curl -X POST http://localhost:8080/add \
-H "Content-Type: application/json" \
-d '{"a": 10, "b": 32}'
```

-----
