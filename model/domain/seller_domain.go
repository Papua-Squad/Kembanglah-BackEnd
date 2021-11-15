package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Products  []Product
}
