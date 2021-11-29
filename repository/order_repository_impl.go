package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type OrderRepositoryImpl struct {
	Server *app.Server
}

func NewOrderRepository(server *app.Server) OrderRepositoryImpl {
	return OrderRepositoryImpl{Server: server}
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	panic("implement me")
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	panic("implement me")
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, order domain.Order) error {
	panic("implement me")
}

func (repository *OrderRepositoryImpl) FindByID(ctx context.Context, orderID uint) (domain.Order, error) {
	panic("implement me")
}
