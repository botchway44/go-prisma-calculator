package out

import (
	"context"
	"go-prisma-calculator/internal/domain/models"
)

// CalculationRepositoryPort is the driven port for database operations.
// Our core logic will depend on this, not a concrete database implementation.
type CalculationRepositoryPort interface {
	Save(ctx context.Context, calc domain.Calculation) error
}