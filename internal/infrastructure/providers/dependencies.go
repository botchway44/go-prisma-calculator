package providers

import (
	"context"
	"log/slog"

	"go-prisma-calculator/internal/application/usecase"
	"go-prisma-calculator/internal/domain/ports/in"
	"go-prisma-calculator/internal/domain/ports/out"
	"go-prisma-calculator/internal/domain/service"
	grpc_adapter "go-prisma-calculator/internal/infrastructure/adapter/grpc"
	rest_adapter "go-prisma-calculator/internal/infrastructure/adapter/rest"
	"go-prisma-calculator/internal/infrastructure/config"
	"go-prisma-calculator/internal/infrastructure/logger"
	"go-prisma-calculator/internal/infrastructure/repository"
	db "go-prisma-calculator/internal/infrastructure/repository/prisma"

	"go.uber.org/fx"
)

// Module bundles all of our application's components for fx.
var Module = fx.Options(
	// 1. Provide the Logger, managing the file lifecycle with fx.
	fx.Provide(func(lifecycle fx.Lifecycle) *slog.Logger {
		l, f := logger.NewLogger()
		if f != nil {
			lifecycle.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					return f.Close()
				},
			})
		}
		return l
	}),

	// 2. Provide the application configuration.
	fx.Provide(config.NewConfig),

	// 3. Provide the Prisma database client, which depends on Config.
	fx.Provide(func(c *config.Config) *db.PrismaClient {
		// In a real app, you would use c.DatabaseURL
		client := db.NewClient()
		if err := client.Connect(); err != nil {
			panic(err) // In a real app, handle this more gracefully
		}
		return client
	}),

	// 4. Provide the Repository, mapping the implementation to the outbound port.
	fx.Provide(
		fx.Annotate(
			repository.NewPrismaRepository,
			fx.As(new(out.CalculationRepositoryPort)),
		),
	),

	// 5. Provide the Domain Service, which depends on the repository port.
	fx.Provide(service.NewCalculatorService),

	// 6. Provide the Application Usecase, mapping the implementation to the inbound port.
	fx.Provide(
		fx.Annotate(
			usecase.NewCalculatorUseCase,
			fx.As(new(in.CalculatorPort)),
		),
	),

	// 7. Provide the API adapters, which depend on the usecase port and the logger.
	fx.Provide(grpc_adapter.NewAdapter),
	fx.Provide(rest_adapter.NewAdapter),
)