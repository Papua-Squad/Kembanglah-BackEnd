package domain

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ProductID      pq.Int64Array `gorm:"type:integer[]"`
	SellerID       uint
	CustomerID     uint
	Address        string
	PaymentMethode string
	StatusCode     uint
	Status         string
}
