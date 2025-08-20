package providers

import (
	"go-prisma-calculator/internal/application/usecase"
	"go-prisma-calculator/internal/domain/ports/in"
	"go-prisma-calculator/internal/domain/ports/out"
	"go-prisma-calculator/internal/domain/service"
	grpc_adapter "go-prisma-calculator/internal/infrastructure/adapter/grpc"
	rest_adapter "go-prisma-calculator/internal/infrastructure/adapter/rest"
	"go-prisma-calculator/internal/infrastructure/config"
	"go-prisma-calculator/internal/infrastructure/repository"
	db "go-prisma-calculator/internal/infrastructure/repository/prisma"

	"go.uber.org/fx"
)

// Module bundles all of our application's components for fx.
var Module = fx.Options(
	// 1. Provide the application configuration.
	fx.Provide(config.NewConfig),

	// 2. Provide the Prisma database client. It depends on Config.
	fx.Provide(func(c *config.Config) *db.PrismaClient {
		// In a real app, you would use c.DatabaseURL
		client := db.NewClient()
		if err := client.Connect(); err != nil {
			panic(err) // In a real app, handle this more gracefully
		}
		return client
	}),

	// 3. Provide the Repository, mapping the implementation to the port.
	fx.Provide(
		fx.Annotate(
			repository.NewPrismaRepository,
			fx.As(new(out.CalculationRepositoryPort)),
		),
	),

	// 4. Provide the Domain Service, which depends on the repository port.
	fx.Provide(service.NewCalculatorService),

	// 5. Provide the Application Usecase, mapping the implementation to the port.
	fx.Provide(
		fx.Annotate(
			usecase.NewCalculatorUseCase,
			fx.As(new(in.CalculatorPort)),
		),
	),

	// 6. Provide the API adapters, which depend on the usecase port.
	fx.Provide(grpc_adapter.NewAdapter),
	fx.Provide(rest_adapter.NewAdapter),
)
