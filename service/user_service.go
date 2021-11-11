package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type UserService interface {
	Login(ctx context.Context, userId uint) web.User
	Register(ctx context.Context, request web.UserCreateRequest) web.User
}
