package repository

import (
	"context"
	"kembanglah/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, product domain.Product) error
	FindByID(ctx context.Context, productID uint) (product domain.Product, err error)
	FindBySeller(ctx context.Context, sellerID uint) (products []domain.Product, err error)
	FindByCategory(ctx context.Context, categoryID uint) (products []domain.Product, err error)
	FindAll(ctx context.Context) (products []domain.Product, err error)
}
