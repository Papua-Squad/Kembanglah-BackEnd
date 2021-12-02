package service

import (
	"context"
	"kembanglah/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productID uint)
	FindByID(ctx context.Context, productID uint) web.ProductResponse
	FindBySeller(ctx context.Context, sellerID uint) web.ProductResponse
	FindByCategory(ctx context.Context, categoryID uint) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
