package repository

import (
	"context"

	domain "go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/out"

	// Only one import is needed, with the alias 'db'.
	db "go-prisma-calculator/internal/infrastructure/repository/prisma"
)

// PrismaRepository is the Prisma implementation of our repository port.
type PrismaRepository struct {
	client *db.PrismaClient
}

// NewPrismaRepository creates a new repository with a Prisma client.
func NewPrismaRepository(client *db.PrismaClient) out.CalculationRepositoryPort {
	return &PrismaRepository{
		client: client,
	}
}

// Save implements the port and saves a calculation to the database.
func (r *PrismaRepository) Save(ctx context.Context, calc domain.Calculation) error {
	_, err := r.client.Calculation.CreateOne(
		db.Calculation.Operation.Set(calc.Operation),
		db.Calculation.A.Set(calc.A),
		db.Calculation.B.Set(calc.B),
		db.Calculation.Result.Set(calc.Result),
	).Exec(ctx)

	return err
}
