package web

type ProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Weight      int    `json:"weight" form:"weight" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	SellerID    uint   `json:"seller_id" form:"seller_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Weight      int    `json:"weight"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	SellerID    uint   `json:"seller_id"`
	ImageUrl    string `json:"image_url"`
}

type ProductUpdateRequest struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Weight      int    `json:"weight" form:"weight" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	SellerID    uint   `json:"seller_id" form:"seller_id" validate:"required"`
}
