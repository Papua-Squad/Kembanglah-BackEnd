package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type UserRepositoryImpl struct {
	server *app.Server
}

func NewUserRepository(server *app.Server) UserRepository {
	return &UserRepositoryImpl{server: server}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	return user, repository.server.DB.Create(&user).Error
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, user domain.User) (domain.User, error) {
	return user, repository.server.DB.First(&user, user.ID).Error
}
