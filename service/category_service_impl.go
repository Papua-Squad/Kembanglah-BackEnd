package service

import (
	"context"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"

	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository}
}
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse {
	categoryRequest := domain.Category{
		Name: request.Name,
		Type: request.Type,
	}

	categoryResponse, err := service.CategoryRepository.Create(ctx, categoryRequest)
	helper.PanicIfError(err)

	return web.CategoryResponse{
		ID:   categoryResponse.ID,
		Name: categoryResponse.Name,
		Type: categoryResponse.Type,
	}
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	var category domain.Category
	category.ID = request.ID

	_, err := service.CategoryRepository.FindByID(ctx, category.ID)
	helper.PanicIfError(err)

	categoryResponse, err := service.CategoryRepository.Update(ctx, domain.Category{
		Model: gorm.Model{
			ID: request.ID,
		},
		Name: request.Name,
		Type: request.Type,
	})

	helper.PanicIfError(err)

	return web.CategoryResponse{
		ID:   categoryResponse.ID,
		Name: categoryResponse.Name,
		Type: categoryResponse.Type,
	}
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID uint) {
	var category domain.Category
	category.ID = categoryID

	err := service.CategoryRepository.Delete(ctx, category)
	helper.PanicIfError(err)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID uint) web.CategoryResponse {
	categoryResponse, err := service.CategoryRepository.FindByID(ctx, categoryID)
	helper.PanicIfError(err)

	return web.CategoryResponse{
		ID:       categoryResponse.ID,
		Name:     categoryResponse.Name,
		Type:     categoryResponse.Type,
		Products: categoryResponse.Products,
	}
}
