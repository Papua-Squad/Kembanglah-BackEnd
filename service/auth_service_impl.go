package service

import (
	"context"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	userRequest :=
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse {
	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	userRequest := domain.User{
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Role:     request.Role,
	}

	userResponse, err := service.UserRepository.Create(ctx, userRequest)
	helper.PanicIfError(err)

	return web.RegisterResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Email:    userResponse.Email,
		Username: userResponse.Username,
	}
}
