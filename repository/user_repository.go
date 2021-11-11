package repository

import (
	"context"
	"kembanglah/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, user domain.User) (domain.User, error)
}
