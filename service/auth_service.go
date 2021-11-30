package service

import (
	"context"
	"kembanglah/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse
}