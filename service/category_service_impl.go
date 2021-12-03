package service

import (
	"context"
	"gorm.io/gorm"
	"kembanglah/model/domain"
	"kembanglah/model/web"
	"kembanglah/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository}
}
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryRequest) (response web.CategoryResponse, err error) {
	categoryRequest := domain.Category{
		Name: request.Name,
		Type: request.Type,
	}

	categoryResponse, err := service.CategoryRepository.Create(ctx, categoryRequest)
	if err != nil {
		return response, err
	}

	return web.CategoryResponse{
		ID:   categoryResponse.ID,
		Name: categoryResponse.Name,
		Type: categoryResponse.Type,
	}, nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (response web.CategoryResponse, err error) {
	var category domain.Category
	category.ID = request.ID

	if _, err := service.CategoryRepository.FindByID(ctx, category.ID); err != nil {
		return response, err
	}

	categoryResponse, err := service.CategoryRepository.Update(ctx, domain.Category{
		Model: gorm.Model{ID: category.ID},
		Name:  request.Name,
		Type:  request.Type,
	})

	return web.CategoryResponse{
		ID:   categoryResponse.ID,
		Name: categoryResponse.Name,
		Type: categoryResponse.Type,
	}, nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID uint) error {
	var category domain.Category
	category.ID = categoryID
	return service.CategoryRepository.Delete(ctx, category)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID uint) (response web.CategoryResponse, err error) {
	categoryResponse, err := service.CategoryRepository.FindByID(ctx, categoryID)
	if err != nil {
		return response, err
	}

	return web.CategoryResponse{
		ID:   categoryResponse.ID,
		Name: categoryResponse.Name,
		Type: categoryResponse.Type,
	}, nil
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) (categories []web.CategoryResponse, err error) {
	categoryResponse, err := service.CategoryRepository.FindAll(ctx)
	if err != nil {
		return categories, err
	}

	for _, category := range categoryResponse {
		categories = append(categories, web.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
			Type: category.Type,
		})
	}

	return categories, nil
}
