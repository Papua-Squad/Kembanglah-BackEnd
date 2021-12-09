package service

import (
	"context"
	"gorm.io/gorm"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (response web.UserResponse, err error) {
	var user domain.User
	user.ID = request.ID

	_, err = service.UserRepository.FindByID(ctx, user.ID)
	if err != nil {
		return response, err
	}

	userResponse, err := service.UserRepository.Update(ctx, domain.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
	})
	if err != nil {
		return response, err
	}

	return web.UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Role:     userResponse.Role,
		ImageUrl: userResponse.ImageUrl,
	}, nil
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, request web.UserUpdatePasswordRequest) (response web.UserResponse, err error) {
	var user domain.User
	user.ID = request.ID

	_, err = service.UserRepository.FindByID(ctx, user.ID)
	if err != nil {
		return response, err
	}

	userResponse, err := service.UserRepository.Update(ctx, domain.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		Password: request.Password,
	})
	if err != nil {
		return response, err
	}

	return web.UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Role:     userResponse.Role,
		ImageUrl: userResponse.ImageUrl,
	}, nil
}

func (service *UserServiceImpl) UpdateImage(ctx context.Context, request web.UserUpdateImageRequest) (response web.UserResponse, err error) {
	var user domain.User
	user.ID = request.ID

	_, err = service.UserRepository.FindByID(ctx, user.ID)
	if err != nil {
		return response, err
	}

	userResponse, err := service.UserRepository.Update(ctx, domain.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		return response, err
	}

	return web.UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Role:     userResponse.Role,
		ImageUrl: userResponse.ImageUrl,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userID uint) error {
	var user domain.User
	user.ID = userID

	return service.UserRepository.Delete(ctx, user)
}

func (service *UserServiceImpl) FindByID(ctx context.Context, userID uint) (response web.UserResponse, err error) {
	userResponse, err := service.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return response, err
	}

	return web.UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Role:     userResponse.Role,
		ImageUrl: userResponse.ImageUrl,
	}, nil
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) (response web.UserResponse, err error) {
	userResponse, err := service.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		return response, err
	}

	return web.UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Role:     userResponse.Role,
		ImageUrl: userResponse.ImageUrl,
	}, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) (responses []web.UserResponse, err error) {
	userResponse, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		return responses, err
	}

	for _, user := range userResponse {
		responses = append(responses, web.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			ImageUrl: user.ImageUrl,
		})
	}
	return responses, nil
}
