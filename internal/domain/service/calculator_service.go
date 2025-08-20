package service

import (
	"context"
	"errors"

	"go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/out"
)

type CalculatorService struct {
	repo out.CalculationRepositoryPort
}

// NewCalculatorService creates the service with its dependencies.
func NewCalculatorService(repo out.CalculationRepositoryPort) *CalculatorService {
	return &CalculatorService{repo: repo}
}

func (s *CalculatorService) Add(ctx context.Context, a, b int32) (*domain.Calculation, error) {
	result := a + b
	calculation := domain.Calculation{
		Operation: "add",
		A:         int(a),
		B:         int(b),
		Result:    int(result),
	}

	// Use the repository port to save the data.
	if err := s.repo.Save(ctx, calculation); err != nil {
		return nil, err
	}

	return &calculation, nil
}

func (s *CalculatorService) Divide(ctx context.Context, dividend, divisor int32) (*domain.Calculation, error) {
	if divisor == 0 {
		return nil, errors.New("cannot divide by zero")
	}

	result := dividend / divisor
	calculation := domain.Calculation{
		Operation: "divide",
		A:         int(dividend),
		B:         int(divisor),
		Result:    int(result),
	}

	if err := s.repo.Save(ctx, calculation); err != nil {
		return nil, err
	}

	return &calculation, nil
}