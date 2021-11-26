package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type SellerRepositoryImpl struct {
	server *app.Server
}

func NewSellerRepository(server *app.Server) SellerRepository {
	return &SellerRepositoryImpl{server: server}
}

func (repository *SellerRepositoryImpl) Create(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	return seller, repository.server.DB.Create(&seller).Error
}

func (repository *SellerRepositoryImpl) Update(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	return seller, repository.server.DB.Model(&seller).Updates(&seller).Error
}

func (repository *SellerRepositoryImpl) Delete(ctx context.Context, seller domain.Seller) error {
	if err := repository.
		server.DB.
		WithContext(ctx).
		First(&seller, seller.ID).Error; err != nil {
		return err
	}

	return repository.
		server.DB.
		WithContext(ctx).
		Delete(&seller).Error
}

func (repository *SellerRepositoryImpl) FindByID(ctx context.Context, sellerId uint) (seller domain.Seller, err error) {
	return seller, repository.server.DB.First(&seller, sellerId).Error
}

func (repository *SellerRepositoryImpl) FindByUsername(ctx context.Context, username string) (seller domain.Seller, err error) {
	return seller, repository.server.DB.Take(&seller, username).Error
}

func (repository *SellerRepositoryImpl) FindAll(ctx context.Context) (seller []domain.Seller, err error) {
	return seller, repository.server.DB.Find(&seller).Error
}
