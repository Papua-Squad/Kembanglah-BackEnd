package service

import (
	"context"
	"kembanglah/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductRequest) (response web.ProductResponse, err error)
	Update(ctx context.Context, request web.ProductUpdateRequest) (response web.ProductResponse, err error)
	Delete(ctx context.Context, productID uint) error
	FindByID(ctx context.Context, productID uint) (response web.ProductResponse, err error)
	FindBySeller(ctx context.Context, sellerID uint) (responses []web.ProductResponse, err error)
	FindByCategory(ctx context.Context, categoryID uint) (responses []web.ProductResponse, err error)
	FindAll(ctx context.Context) (responses []web.ProductResponse, err error)
}
