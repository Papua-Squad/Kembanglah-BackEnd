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
	return user, repository.server.DB.Create(&user).Error
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	return customer, repository.server.DB.Model(&customer).Updates(&customer).Error
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, seller domain.Customer) error {
	panic("implement me")
}

func (repository *CustomerRepositoryImpl) FindByID(ctx context.Context, customerId uint) (customer domain.Customer, err error) {
	return customer, repository.server.DB.First(&customer, customerId).Error
}

func (repository *CustomerRepositoryImpl) FindByUsername(ctx context.Context, username string) (seller []domain.Customer, err error) {
	panic("implement me")
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) (seller []domain.Customer, err error) {
	panic("implement me")
}
