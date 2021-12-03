package repository

import (
	"context"
	"kembanglah/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, product domain.Category) error
	FindByID(ctx context.Context, categoryID uint) (category domain.Category, err error)
	FindAll(ctx context.Context) (categories []domain.Category, err error)
}
