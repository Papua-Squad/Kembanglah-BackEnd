package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type ProductRepositoryImpl struct {
	Server *app.Server
}

func NewProductRepository(server *app.Server) ProductRepository {
	return &ProductRepositoryImpl{
		Server: server,
	}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	return product, repository.Server.DB.WithContext(ctx).Create(&product).Error
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	return product, repository.Server.DB.WithContext(ctx).Model(&product).Updates(&product).Error
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	return repository.Server.DB.WithContext(ctx).Delete(&product).Error
}

func (repository *ProductRepositoryImpl) FindByID(ctx context.Context, productID uint) (product domain.Product, err error) {
	return product, repository.Server.DB.WithContext(ctx).First(&product, productID).Error
}

func (repository *ProductRepositoryImpl) FindBySeller(ctx context.Context, sellerID uint) (products []domain.Product, err error) {
	return products, repository.Server.DB.WithContext(ctx).Where("seller_id = ?", sellerID).Find(&products).Error
}

func (repository *ProductRepositoryImpl) FindByCategory(ctx context.Context, categoryID uint) (products []domain.Product, err error) {
	return products, repository.Server.DB.WithContext(ctx).Where("category_id = ?", categoryID).Find(&products).Error

}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) (products []domain.Product, err error) {
	return products, repository.Server.DB.WithContext(ctx).Find(&products).Error
}
