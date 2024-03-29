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
	return user, repository.server.DB.WithContext(ctx).Create(&user).Error
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	return user, repository.server.DB.WithContext(ctx).Model(&user).Updates(&user).Error
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, user domain.User) error {
	return repository.server.DB.WithContext(ctx).Delete(&user).Error
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, userID uint) (user domain.User, err error) {
	return user, repository.server.DB.WithContext(ctx).First(&user, userID).Error
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	return user, repository.server.DB.WithContext(ctx).Where("username = ?", username).Find(&user).Error
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context) (users []domain.User, err error) {
	return users, repository.server.DB.WithContext(ctx).Find(&users).Error
}
