package web

import "kembanglah/model/domain"

type ProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Weight      int    `json:"weight" form:"weight"  validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	ImageUrl    string `json:"image_url" form:"image_url" validate:"required"`
}

type ProductResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	Weight      int               `json:"weight"`
	Description string            `json:"description"`
	ImageUrl    string            `json:"image_url"`
	Categories  []domain.Category `json:"category"`
}

type ProductUpdateRequest struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Weight      int    `json:"weight" form:"weight"`
	Description string `json:"description" form:"description"`
	ImageUrl    string `json:"image_url" form:"image_url"`
}
