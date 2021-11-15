package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint
	UserID    uint
	SubTotal  uint
	Status    string
}
