package repository

import "context"

type UserRepository interface {
	Create(ctx context.Context)
	FindByID(ctx context.Context)
}
