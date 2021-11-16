package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	City        string
	ZipCode     uint
	Country     string
	CustomerID  uint
}
