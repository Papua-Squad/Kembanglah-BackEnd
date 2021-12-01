package web

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Weight      int    `json:"weight" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
