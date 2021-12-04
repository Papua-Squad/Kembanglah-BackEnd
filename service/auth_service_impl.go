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

func (service *AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) (response web.LoginResponse, err error) {
	userResponse, err := service.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		return response, err
	}

	if err := helper.CheckPasswordHash(request.Password, userResponse.Password); err != nil {
		return response, err
	}

	token := helper.GenerateToken(userResponse)
	return web.LoginResponse{
		AuthToken: token,
	}, err
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.RegisterRequest) (response web.RegisterResponse, err error) {
	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return response, err
	}

	userRequest := domain.User{
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Role:     request.Role,
	}

	userResponse, err := service.UserRepository.Create(ctx, userRequest)
	if err != nil {
		return response, err
	}

	return web.RegisterResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Email:    userResponse.Email,
		Username: userResponse.Username,
	}, err
}
