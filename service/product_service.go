package service

import (
	"context"
	"kembanglah/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductRequest) web.ProductResponse
}
