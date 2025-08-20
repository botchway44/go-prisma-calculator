package usecase

import (
	"context"
	domain "go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/in"
	"go-prisma-calculator/internal/domain/service"
)

// CalculatorUseCase implements the inbound port.
type CalculatorUseCase struct {
	calcService *service.CalculatorService
}

// NewCalculatorUseCase creates the use case and depends on the domain service.
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
