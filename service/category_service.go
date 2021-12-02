package service

import (
	"context"
	"kembanglah/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryID uint)
	FindByID(ctx context.Context, categoryID uint) web.CategoryResponse
}
