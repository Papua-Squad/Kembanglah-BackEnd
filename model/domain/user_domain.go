package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Email    string
	Username string
	Password string
	Role     string
}
