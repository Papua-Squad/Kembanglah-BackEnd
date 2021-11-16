package domain

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Addresses []Address
}
