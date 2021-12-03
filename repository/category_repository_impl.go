package repository

import (
	"context"
	"kembanglah/app"
	"kembanglah/model/domain"
)

type CategoryRepositoryImpl struct {
	Server *app.Server
}

func NewCategoryRepository(server *app.Server) CategoryRepository {
	return &CategoryRepositoryImpl{
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
	return category, repository.Server.DB.WithContext(ctx).First(&category, categoryID).Error
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) (categories []domain.Category, err error) {
	return categories, repository.Server.DB.WithContext(ctx).Find(&categories).Error
}
