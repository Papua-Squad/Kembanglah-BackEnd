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
	UserID      uint
	//User        User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" `
}
