package repository

import (
	"context"
	"kembanglah/model/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, user domain.Customer) (domain.Customer, error)
	Update(ctx context.Context, user domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, customer domain.Customer) error
	FindByID(ctx context.Context, customerId uint) (domain.Customer, error)
	FindByUsername(ctx context.Context, username string) ([]domain.Customer, error)
	FindAll(ctx context.Context) ([]domain.Customer, error)
}
