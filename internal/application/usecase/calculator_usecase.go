package usecase

import (
	"context"
	domain "go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/in"
	"go-prisma-calculator/internal/domain/service"
)

// CalculatorUseCase implements the inbound port (in.CalculatorPort).
type CalculatorUseCase struct {
	// It depends on the pure domain service for the actual business logic.
	calcService *service.CalculatorService
}

// NewCalculatorUseCase is the constructor that fx uses to create an instance.
// It receives the domain service as a dependency.
func NewCalculatorUseCase(calcService *service.CalculatorService) in.CalculatorPort {
	return &CalculatorUseCase{calcService: calcService}
}

// Add orchestrates the 'add' operation by calling the domain service.
func (uc *CalculatorUseCase) Add(ctx context.Context, a, b int32) (*domain.Calculation, error) {
	return uc.calcService.Add(ctx, a, b)
}

// Divide orchestrates the 'divide' operation by calling the domain service.
func (uc *CalculatorUseCase) Divide(ctx context.Context, dividend, divisor int32) (*domain.Calculation, error) {
	return uc.calcService.Divide(ctx, dividend, divisor)
}
