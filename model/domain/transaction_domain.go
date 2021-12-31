package domain

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ProductsID      pq.Int64Array `gorm:"type:integer[]"`
	SellerID        uint
	CustomerID      uint
	ShippingAddress string
	PaymentMethode  string
	Status          string
	SubTotal        int
}
