package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type CustomerRepositoryImpl struct {
	server *app.Server
}

func NewCustomerRepository(server *app.Server) CustomerRepository {
	return &CustomerRepositoryImpl{server: server}
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, user domain.Customer) (domain.Customer, error) {
	return user, repository.server.DB.WithContext(ctx).Create(&user).Error
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	return customer, repository.server.DB.WithContext(ctx).Model(&customer).Updates(&customer).Error
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customer domain.Customer) error {
	return repository.server.DB.WithContext(ctx).Delete(&customer).Error
}

func (repository *CustomerRepositoryImpl) FindByID(ctx context.Context, customerID uint) (customer domain.Customer, err error) {
	return customer, repository.server.DB.WithContext(ctx).First(&customer, customerID).Error
}

func (repository *CustomerRepositoryImpl) FindByUsername(ctx context.Context, username string) (customer domain.Customer, err error) {
	return customer, repository.server.DB.WithContext(ctx).Where("username = ?", username).First(&customer).Error
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) (customer []domain.Customer, err error) {
	return  customer, repository.server.DB.WithContext(ctx).Find(&customer).Error
}
