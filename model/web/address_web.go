package web

type AddressRequest struct {
	FirstName   string `json:"first_name" form:"first_name" validate:"required"`
	LastName    string `json:"last_name" form:"last_name" validate:"required"`
	PhoneNumber string `json:"phone" form:"phone" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	City        string `json:"city" form:"city" validate:"required"`
	ZipCode     uint   `json:"zip_code" form:"zip_code" validate:"required"`
	Country     string `json:"country" form:"country" validate:"required"`
	UserID      uint   `json:"user_id" form:"user_id" validate:"required"`
}

type AddressResponse struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     uint   `json:"zip_code"`
	Country     string `json:"country"`
	UserID      uint   `json:"user_id"`
}
