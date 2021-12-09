package web

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ImageUrl string `json:"image_url"`
}

type UserUpdateRequest struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Username string `json:"username" form:"username" validate:"required"`
}

type UserUpdatePasswordRequest struct {
	ID       uint   `json:"id"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UserUpdateImageRequest struct {
	ID       uint   `json:"id"`
	ImageUrl string `json:"image_url"  validate:"required"`
}
