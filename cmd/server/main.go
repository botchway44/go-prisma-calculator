package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"go-prisma-calculator/internal/infrastructure/providers"
	// Ensure this import path matches your new generated folder structure
	pb "go-prisma-calculator/generated/proto" 

	grpc_adapter "go-prisma-calculator/internal/infrastructure/adapter/grpc"
	rest_adapter "go-prisma-calculator/internal/infrastructure/adapter/rest"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	app := fx.New(
		providers.Module,
		// fx.Invoke tells fx to run a function. We use it to start our servers.
		fx.Invoke(runServers),
	)
	app.Run()
}

// runServers is the function that depends on our adapters to start the servers.
// fx will automatically provide the adapters from the dependency graph.
func runServers(
	lifecycle fx.Lifecycle,
	grpcAdapter *grpc_adapter.Adapter,
	restAdapter *rest_adapter.Adapter,
) {
	// We use the fx Lifecycle to gracefully start and stop our servers
	// when the application starts and stops.
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start gRPC server in a separate goroutine
			go func() {
				lis, err := net.Listen("tcp", ":50051")
				if err != nil {
					log.Fatalf("gRPC failed to listen: %v", err)
				}
				grpcServer := grpc.NewServer()
				pb.RegisterCalculatorServiceServer(grpcServer, grpcAdapter)
				fmt.Println("gRPC server listening on :50051")
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatalf("gRPC server failed to serve: %v", err)
				}
			}()

			// Start Gin REST server in a separate goroutine
			go func() {
				router := gin.Default()
				router.POST("/add", restAdapter.AddHandler)
				// Add a route for the divide handler here if you created one
				fmt.Println("REST (Gin) server listening on :8080")
				if err := http.ListenAndServe(":8080", router); err != nil {
					log.Fatalf("REST server failed to serve: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping servers.")
			// In a real app, you would gracefully shut down servers here.
			return nil
		},
	})
}