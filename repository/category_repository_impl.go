package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type CategoryRepositoryImpl struct {
	Server *app.Server
}

func NewCategoryRepository(server *app.Server) CategoryRepositoryImpl {
	return CategoryRepositoryImpl{
		Server: server,
	}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	return category, repository.Server.DB.WithContext(ctx).Create(&category).Error
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	return category, repository.Server.DB.WithContext(ctx).Model(&category).Updates(&category).Error
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, category domain.Category) error {
	return repository.Server.DB.WithContext(ctx).Delete(&category).Error
}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, categoryID uint) (category domain.Category, err error) {
	return category, repository.Server.DB.WithContext(ctx).Preload("Products").First(&category, categoryID).Error
}
