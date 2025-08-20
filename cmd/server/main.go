package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	// Import your providers and adapters
	grpc_adapter "go-prisma-calculator/internal/infrastructure/adapter/grpc"
	rest_adapter "go-prisma-calculator/internal/infrastructure/adapter/rest"
	"go-prisma-calculator/internal/infrastructure/providers"

	// Import your generated protobuf package
	pb "go-prisma-calculator/generated/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	// Create a new fx application, providing all the dependency modules.
	app := fx.New(
		providers.Module,
		// fx.Invoke tells fx to run a function. We use it to start our servers.
		fx.Invoke(runServers),
	)

	// Run the application. fx will manage the entire lifecycle.
	app.Run()
}

// runServers is the function that depends on our adapters and logger to start the servers.
// fx will automatically provide these dependencies from the graph.
func runServers(
	lifecycle fx.Lifecycle,
	logger *slog.Logger,
	grpcAdapter *grpc_adapter.Adapter,
	restAdapter *rest_adapter.Adapter,
) {
	// We use the fx Lifecycle to gracefully start and stop our servers.
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start gRPC server in a separate goroutine
			go func() {
				lis, err := net.Listen("tcp", ":50051")
				if err != nil {
					logger.Error("gRPC failed to listen", slog.String("error", err.Error()))
					return
				}
				grpcServer := grpc.NewServer()
				pb.RegisterCalculatorServiceServer(grpcServer, grpcAdapter)
				logger.Info("gRPC server listening on :50051")
				if err := grpcServer.Serve(lis); err != nil {
					logger.Error("gRPC server failed to serve", slog.String("error", err.Error()))
				}
			}()

			// Start Gin REST server in a separate goroutine
			go func() {
				router := gin.Default()
				router.POST("/add", restAdapter.AddHandler)
				router.POST("/divide", restAdapter.DivideHandler)
				
				// Routes for Swagger/OpenAPI documentation
				router.StaticFile("/swagger.json", "./docs/calculator.swagger.json")
				router.GET("/swagger", func(c *gin.Context) {
					c.Header("Content-Type", "text/html; charset=utf-8")
					c.String(http.StatusOK, swaggerHTML)
				})

				logger.Info("REST (Gin) server listening on :8080")
				logger.Info("Find Swagger UI at http://localhost:8080/swagger")

				if err := http.ListenAndServe(":8080", router); err != nil {
					logger.Error("REST server failed to serve", slog.String("error", err.Error()))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping servers.")
			// In a real app, you would gracefully shut down servers here.
			return nil
		},
	})
}

// swaggerHTML contains the simple HTML page for rendering the Swagger UI.
const swaggerHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
    <style>
        html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
        body { margin: 0; background: #fafafa; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
        window.onload = function() {
            SwaggerUIBundle({
                url: "/swagger.json", // Points to the JSON file we are serving
                dom_id: '#swagger-ui',
            });
        };
    </script>
</body>
</html>
`