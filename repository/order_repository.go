package repository

import (
	"context"
	"kembanglah/model/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, order domain.Order) error
	FindByID(ctx context.Context, orderID uint) (domain.Order, error)
}