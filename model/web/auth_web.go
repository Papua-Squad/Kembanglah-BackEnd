package web

type RegisterRequest struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Role     string `json:"role" form:"role" validate:"required"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AuthToken string           `json:"auth_token"`
	Profile   RegisterResponse `json:"profile"`
}
