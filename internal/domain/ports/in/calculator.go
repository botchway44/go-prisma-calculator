package in

import (
	"context"
	domain "go-prisma-calculator/internal/domain/models"
)

// CalculatorPort is the driving port for our application.
type CalculatorPort interface {
	Add(ctx context.Context, a, b int32) (*domain.Calculation, error)
	Divide(ctx context.Context, a, b int32) (*domain.Calculation, error)
}
