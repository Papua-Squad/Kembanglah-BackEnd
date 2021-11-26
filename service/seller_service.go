package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type SellerService interface {
	// Login(ctx context.Context, userId uint) web.User
	Register(ctx context.Context, request web.SellerCreateRequest) web.Seller
	Update(ctx context.Context, request web.SellerUpdateRequest) web.Seller
	Delete(ctx context.Context, sellerId uint)
	FindByID(ctx context.Context, sellerId uint) web.Seller
	FindByUsername(ctx context.Context, username string) web.Seller
	FindAll(ctx context.Context) []web.Seller
}
