package service

import (
	"context"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"gorm.io/gorm"
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

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	var product domain.Product
	product.ID = request.ID

	_, err := service.ProductRepository.FindByID(ctx, product.ID)
	helper.PanicIfError(err)

	productResponse, err := service.ProductRepository.Update(ctx, domain.Product{
		Model: gorm.Model{
			ID: request.ID,
		},
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Weight:      request.Weight,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
	})

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

func (service *ProductServiceImpl) Delete(ctx context.Context, productID uint) {
	var product domain.Product
	product.ID = productID

	err := service.ProductRepository.Delete(ctx, product)
	helper.PanicIfError(err)
}

func (service *ProductServiceImpl) FindByID(ctx context.Context, productID uint) web.ProductResponse {
	productResponse, err := service.ProductRepository.FindByID(ctx, productID)
	helper.PanicIfError(err)

	return web.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Stock:       productResponse.Stock,
		Weight:      productResponse.Stock,
		Description: productResponse.Description,
		ImageUrl:    productResponse.ImageUrl,
		Categories:  productResponse.Categories,
	}
}

func (service *ProductServiceImpl) FindBySeller(ctx context.Context, sellerID uint) web.ProductResponse {
	panic("implement me")

}

func (service *ProductServiceImpl) FindByCategory(ctx context.Context, categoryID uint) web.ProductResponse {
	panic("impl me")
	// productResponse, err := service.ProductRepository.FindByCategory(ctx, categoryID)
	// helper.PanicIfError(err)

	// return web.ProductResponse{}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	var newProduct []web.ProductResponse

	product, err := service.ProductRepository.FindAll(ctx)
	helper.PanicIfError(err)

	for _, productResponse := range product {
		newProduct = append(newProduct, web.ProductResponse{
			ID:          productResponse.ID,
			Name:        productResponse.Name,
			Stock:       productResponse.Stock,
			Weight:      productResponse.Stock,
			Description: productResponse.Description,
			ImageUrl:    productResponse.ImageUrl,
			Categories:  productResponse.Categories,
		})
	}
	return newProduct

}
