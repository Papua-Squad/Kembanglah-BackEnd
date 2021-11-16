package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID  uint
	CustomerID uint
	Address    string
	SubTotal   uint
	Status     string
}
