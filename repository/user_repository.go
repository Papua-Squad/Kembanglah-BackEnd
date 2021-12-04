package repository

import (
	"context"
	"kembanglah/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
	FindByID(ctx context.Context, userId uint) (user domain.User, err error)
	FindByUsername(ctx context.Context, username string) (user domain.User, err error)
	FindAll(ctx context.Context) (sellers []domain.User, err error)
}
