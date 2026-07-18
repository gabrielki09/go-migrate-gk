package esterepository

import (
	"context"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

type EsteRepository struct {
	db *pgxpool.Pool
}

func NewEsteRepository(db *pgxpool.Pool) *EsteRepository {
	return &EsteRepository{
		db: db,
	}
}

func (f *EsteRepository) GetAll(ctx context.Context) ([]any, error) {
	return nil, nil
}

func (f *EsteRepository) Create(ctx context.Context, payload any) (any, error) {
	return nil, nil
}

func (f *EsteRepository) Update(ctx context.Context, payload any, financialAccountId int) (any, error) {
	return nil, nil
}

func (f *EsteRepository) FindByID(ctx context.Context, financialAccountId int) (any, error) {
	return nil, nil
}

func (f *EsteRepository) Delete(ctx context.Context, financialAccountId int) error {
	return nil
}

func (f *EsteRepository) Active(ctx context.Context, financialAccountId int) error {
	return nil
}


	