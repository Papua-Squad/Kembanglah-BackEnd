package repository

import (
	"context"
	"kembanglah/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
}
