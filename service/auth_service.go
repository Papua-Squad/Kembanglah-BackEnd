package service

import (
	"context"
	"kembanglah/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) (response web.LoginResponse, err error)
	Register(ctx context.Context, request web.RegisterRequest) (response web.RegisterResponse, err error)
}
