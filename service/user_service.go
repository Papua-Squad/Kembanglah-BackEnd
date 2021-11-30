package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type UserService interface {
	Update(ctx context.Context, request web.UserUpdateRequest) web.User
	Delete(ctx context.Context, sellerId uint)
	FindByID(ctx context.Context, sellerId uint) web.Seller
	FindByUsername(ctx context.Context, username string) web.Seller
	FindAll(ctx context.Context) []web.Seller
}
