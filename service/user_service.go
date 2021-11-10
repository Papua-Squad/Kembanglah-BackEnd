package service

import "golang.org/x/net/context"

type UserService interface {
	Login(ctx context.Context)
	Register(ctx context.Context)
}
