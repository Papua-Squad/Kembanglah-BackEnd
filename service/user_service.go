package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type UserService interface {
	Update(ctx context.Context, request web.UserUpdateRequest) (response web.UserResponse, err error)
	UpdatePassword(ctx context.Context, request web.UserUpdatePasswordRequest) (response web.UserResponse, err error)
	UpdateImage(ctx context.Context, request web.UserUpdateImageRequest) (response web.UserResponse, err error)
	Delete(ctx context.Context, userID uint) error
	FindByID(ctx context.Context, userID uint) (response web.UserResponse, err error)
	FindByUsername(ctx context.Context, username string) (response web.UserResponse, err error)
	FindAll(ctx context.Context) (responses []web.UserResponse, err error)
}
