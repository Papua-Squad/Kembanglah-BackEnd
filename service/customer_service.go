package service

import (
	"kembanglah/model/web"

	"golang.org/x/net/context"
)

type CustomerService interface {
	Login(ctx context.Context, customerId uint) web.Customer
	Register(ctx context.Context, request web.CustomerCreateRequest) web.Customer
	FindByID(ctx context.Context, customerId uint) web.Customer
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.Customer
}
