package service

import (
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
	}
}

func (service *CustomerServiceImpl) Login(ctx context.Context, customerId uint) web.Customer {
	customer := domain.Customer{}
	customer.ID = customerId

	customer, err := service.CustomerRepository.FindByID(ctx, customerId)
	helper.PanicIfError(err)

	return web.Customer{
		ID: customer.ID,
	}
}

func (service *CustomerServiceImpl) Register(ctx context.Context, request web.CustomerCreateRequest) web.Customer {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	customerRequest := domain.Customer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hash),
	}
	customer, err := service.CustomerRepository.Create(ctx, customerRequest)
	helper.PanicIfError(err)

	return web.Customer{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Username:  customer.Username,
		Email:     customer.Email,
		Password:  customer.Password,
	}
}

func (service *CustomerServiceImpl) FindByID(ctx context.Context, customerId uint) web.Customer {
	customer, err := service.CustomerRepository.FindByID(ctx, customerId)
	helper.PanicIfError(err)

	return web.Customer{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Username:  customer.Username,
		Email:     customer.Email,
		Password:  customer.Password,
	}
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.Customer {
	var customer domain.Customer
	customer.ID = request.ID

	_, err := service.CustomerRepository.FindByID(ctx, customer.ID)
	helper.PanicIfError(err)

	user, err := service.CustomerRepository.Update(ctx, domain.Customer{
		Model: gorm.Model{
			ID: request.ID,
		},
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
	})
	helper.PanicIfError(err)

	return web.Customer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}
