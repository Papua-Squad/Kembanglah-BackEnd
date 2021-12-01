package service

import (
	"context"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductRequest) web.ProductResponse {
	productRequest := domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Weight:      request.Weight,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
	}

	productResponse, err := service.ProductRepository.Create(ctx, productRequest)
	helper.PanicIfError(err)

	return web.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Stock:       productResponse.Stock,
		Weight:      productResponse.Stock,
		Description: productResponse.Description,
		ImageUrl:    productResponse.ImageUrl,
	}
}
