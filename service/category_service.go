package service

import (
	"context"
	"kembanglah/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequest) (response web.CategoryResponse, err error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (response web.CategoryResponse, err error)
	Delete(ctx context.Context, categoryID uint) error
	FindByID(ctx context.Context, categoryID uint) (response web.CategoryResponse, err error)
	FindAll(ctx context.Context) (categories []web.CategoryResponse, err error)
}
