package repository

import (
	"context"
	"kembanglah/model/domain"
)

type AddressRepository interface {
	Create(ctx context.Context, address domain.Address) (domain.Address, error)
	Delete(ctx context.Context, address domain.Address) error
	FindByUser(ctx context.Context, userID uint) (addresses []domain.Address, err error)
}
