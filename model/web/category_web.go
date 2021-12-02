package web

import "kembanglah/model/domain"

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	Type string `json:"type" form:"type" validate:"required"`
}

type CategoryResponse struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Products []domain.Product `json:"product"`
}

type CategoryUpdateRequest struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Type string `json:"type" form:"type"`
}
