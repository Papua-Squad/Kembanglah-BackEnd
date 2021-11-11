package service

import (
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	userRequest := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: string(hash),
	}
	user, err := service.UserRepository.Create(ctx, userRequest)
	helper.PanicIfError(err)

	return web.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, userId uint) web.User {
	user := domain.User{}
	user.ID = userId

	user, err := service.UserRepository.FindByID(ctx, user)
	helper.PanicIfError(err)

	return web.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}
