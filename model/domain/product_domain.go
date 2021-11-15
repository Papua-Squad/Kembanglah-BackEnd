package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Stock       int
	Weight      int
	Description string
	UserID      uint
	Categories  []Category
}
