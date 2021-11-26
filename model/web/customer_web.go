package web

type Customer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CustomerCreateRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

type CustomerUpdateRequest struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}
