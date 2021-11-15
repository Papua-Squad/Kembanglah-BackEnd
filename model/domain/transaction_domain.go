package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	OrderID    uint
	SellerID   uint
	CustomerID uint
}
