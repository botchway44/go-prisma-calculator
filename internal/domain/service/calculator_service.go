package service

import (
	"context"
	"errors"

	domain "go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/out"
)

// CalculatorService contains the pure business logic for calculations.
type CalculatorService struct {
	// It depends on the outbound repository port to save data.
	repo out.CalculationRepositoryPort
}

// NewCalculatorService is the constructor that fx uses.
// It receives the repository as a dependency.
func NewCalculatorService(repo out.CalculationRepositoryPort) *CalculatorService {
	return &CalculatorService{repo: repo}
}

// Add performs the addition, creates a domain model, and saves it.
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

// Divide performs the division, creates a domain model, and saves it.
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