package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type AddressRepositoryImpl struct {
	Server *app.Server
}

func NewAddressRepository(server *app.Server) AddressRepository {
	return &AddressRepositoryImpl{
		Server: server,
	}
}

func (repository *AddressRepositoryImpl) Create(ctx context.Context, address domain.Address) (domain.Address, error) {
	return address, repository.Server.DB.WithContext(ctx).Create(&address).Error
}

func (repository *AddressRepositoryImpl) Delete(ctx context.Context, address domain.Address) error {
	return repository.Server.DB.WithContext(ctx).Delete(&address).Error
}

func (repository *AddressRepositoryImpl) FindByUser(ctx context.Context, userID uint) (addresses []domain.Address, err error) {
	return addresses, repository.Server.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&addresses).Error
}
