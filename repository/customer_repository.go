package repository

import (
	"context"
	"kembanglah/model/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, customer domain.Customer) error
	FindByID(ctx context.Context, customerID uint) (domain.Customer, error)
	FindByName(ctx context.Context, name string) ([]domain.Customer, error)
	FindAll(ctx context.Context) ([]domain.Customer, error)
}