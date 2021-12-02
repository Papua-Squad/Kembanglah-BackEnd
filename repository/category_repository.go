package repository

import (
	"context"
	"kembanglah/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, product domain.Category) error
	FindByID(ctx context.Context, categoryID uint) (user domain.Category, err error)
}
