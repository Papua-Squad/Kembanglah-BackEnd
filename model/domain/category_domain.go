package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string
	Type string
}
