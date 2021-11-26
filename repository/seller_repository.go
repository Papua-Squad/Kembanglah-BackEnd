package repository

import (
	"context"
	"kembanglah/model/domain"
)

type SellerRepository interface {
	Create(ctx context.Context, seller domain.Seller) (domain.Seller, error)
	Update(ctx context.Context, seller domain.Seller) (domain.Seller, error)
	Delete(ctx context.Context, seller domain.Seller) error
	FindByID(ctx context.Context, sellerID uint) (domain.Seller, error)
	FindByName(ctx context.Context, sellerName string) ([]domain.Seller, error)
	FindAll(ctx context.Context) ([]domain.Seller, error)
}