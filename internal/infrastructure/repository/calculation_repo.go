package repository

import (
	"context"

	domain "go-prisma-calculator/internal/domain/models"
	"go-prisma-calculator/internal/domain/ports/out"

	// Import the generated Prisma client with an alias 'db' for clarity.
	db "go-prisma-calculator/internal/infrastructure/repository/prisma"
)

// PrismaRepository is the Prisma implementation of our repository port.
type PrismaRepository struct {
	client *db.PrismaClient
}

// NewPrismaRepository is the constructor that fx uses to create an instance.
// It receives the Prisma client as a dependency.
func NewPrismaRepository(client *db.PrismaClient) out.CalculationRepositoryPort {
	return &PrismaRepository{
		client: client,
	}
}

// Save implements the port's contract. It translates the domain model
// into a Prisma model and saves it to the database.
func (r *PrismaRepository) Save(ctx context.Context, calc domain.Calculation) error {
	// Use the Prisma client's fluent API to create a new record.
	_, err := r.client.Calculation.CreateOne(
		db.Calculation.Operation.Set(calc.Operation),
		db.Calculation.A.Set(calc.A),
		db.Calculation.B.Set(calc.B),
		db.Calculation.Result.Set(calc.Result),
	).Exec(ctx)

	return err
}
