package service

import (
	"context"
	"gorm.io/gorm"
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

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductRequest) (response web.ProductResponse, err error) {
	productRequest := domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Weight:      request.Weight,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
		CategoryID:  request.CategoryID,
		SellerID:    request.SellerID,
	}

	productResponse, err := service.ProductRepository.Create(ctx, productRequest)
	if err != nil {
		return response, err
	}

	return web.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Price:       productResponse.Price,
		Stock:       productResponse.Stock,
		Weight:      productResponse.Stock,
		Description: productResponse.Description,
		CategoryID:  productResponse.CategoryID,
		SellerID:    productResponse.SellerID,
		ImageUrl:    productResponse.ImageUrl,
	}, nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (response web.ProductResponse, err error) {
	var product domain.Product
	product.ID = request.ID

	_, err = service.ProductRepository.FindByID(ctx, product.ID)
	if err != nil {
		return response, err
	}

	productResponse, err := service.ProductRepository.Update(ctx, domain.Product{
		Model: gorm.Model{
			ID: request.ID,
		},
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Weight:      request.Weight,
		Description: request.Description,
		ImageUrl:    "",
		CategoryID:  request.CategoryID,
		SellerID:    request.SellerID,
	})
	if err != nil {
		return response, err
	}

	return web.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Price:       productResponse.Price,
		Stock:       productResponse.Stock,
		Weight:      productResponse.Stock,
		Description: productResponse.Description,
		CategoryID:  productResponse.CategoryID,
		SellerID:    productResponse.SellerID,
		ImageUrl:    productResponse.ImageUrl,
	}, nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productID uint) error {
	var product domain.Product
	product.ID = productID

	return service.ProductRepository.Delete(ctx, product)
}

func (service *ProductServiceImpl) FindByID(ctx context.Context, productID uint) (response web.ProductResponse, err error) {
	productResponse, err := service.ProductRepository.FindByID(ctx, productID)
	if err != nil {
		return response, err
	}

	return web.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Price:       productResponse.Price,
		Stock:       productResponse.Stock,
		Weight:      productResponse.Weight,
		Description: productResponse.Description,
		CategoryID:  productResponse.CategoryID,
		SellerID:    productResponse.SellerID,
		ImageUrl:    productResponse.ImageUrl,
	}, nil
}

func (service *ProductServiceImpl) FindBySeller(ctx context.Context, sellerID uint) (responses []web.ProductResponse, err error) {
	productResponse, err := service.ProductRepository.FindBySeller(ctx, sellerID)

	if err != nil || len(productResponse) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	for _, product := range productResponse {
		responses = append(responses, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Stock,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})
	}
	return responses, nil
}

func (service *ProductServiceImpl) FindByCategory(ctx context.Context, categoryID uint) (responses []web.ProductResponse, err error) {
	productResponse, err := service.ProductRepository.FindByCategory(ctx, categoryID)
	if err != nil || len(productResponse) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	for _, product := range productResponse {
		responses = append(responses, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Stock,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})
	}
	return responses, nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) (responses []web.ProductResponse, err error) {
	productResponse, err := service.ProductRepository.FindAll(ctx)
	if err != nil {
		return responses, err
	}

	for _, product := range productResponse {
		responses = append(responses, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Weight:      product.Stock,
			Description: product.Description,
			CategoryID:  product.CategoryID,
			SellerID:    product.SellerID,
			ImageUrl:    product.ImageUrl,
		})
	}
	return responses, nil
}
