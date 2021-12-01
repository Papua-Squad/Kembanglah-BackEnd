package service

import (
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.User {
	var usr domain.User
	usr.ID = request.ID

	_, err := service.UserRepository.FindByID(ctx, usr.ID)
	helper.PanicIfError(err)

	user, err := service.UserRepository.Update(ctx, domain.User{
		Model: gorm.Model{
			ID: request.ID,
		},
		FullName: request.FullName,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})

	helper.PanicIfError(err)

	return web.User{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) {
	var user domain.User
	user.ID = userId

	err := service.UserRepository.Delete(ctx, user)
	helper.PanicIfError(err)
}

func (service *UserServiceImpl) FindByID(ctx context.Context, userId uint) web.User {
	user, err := service.UserRepository.FindByID(ctx, userId)
	helper.PanicIfError(err)

	return web.User{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) web.User {
	user, err := service.UserRepository.FindByUsername(ctx, username)
	helper.PanicIfError(err)

	return web.User{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.User {
	var newUser []web.User

	user, err := service.UserRepository.FindAll(ctx)
	helper.PanicIfError(err)

	for _, data := range user {
		newUser = append(newUser, web.User{
			ID:       data.ID,
			FullName: data.FullName,
			Username: data.Username,
			Email:    data.Email,
			Password: data.Password,
		})
	}
	return newUser

}
