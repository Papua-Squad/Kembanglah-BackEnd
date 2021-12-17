package service

import (
	"context"
	"kembanglah/model/web"
)

type AddressService interface {
	Create(ctx context.Context, request web.AddressRequest) (response web.AddressResponse, err error)
	Delete(ctx context.Context, addressID uint) error
	FindByUser(ctx context.Context, userID uint) (responses []web.AddressResponse, err error)
}
