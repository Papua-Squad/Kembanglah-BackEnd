package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type UserService interface {
	Update(ctx context.Context, request web.UserUpdateRequest) web.User
	Delete(ctx context.Context, userID uint)
	FindByID(ctx context.Context, userID uint) (response web.User, err error)
	FindByUsername(ctx context.Context, username string) web.User
	FindAll(ctx context.Context) []web.User
}
