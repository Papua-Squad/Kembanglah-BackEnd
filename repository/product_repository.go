package repository

import (
	"context"
	"kembanglah/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, product domain.Product) error
	FindByID(ctx context.Context, productID uint) (domain.Product, error)
	FindBySeller(ctx context.Context, sellerID uint) ([]domain.Product, error)
	FindByCategory(ctx context.Context, categoryID uint) ([]domain.Product, error)
}
