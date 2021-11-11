package domain

import (
	"gorm.io/gorm"
)

const (
	UserTableName = "user"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200);not_null"`
	Email    string `gorm:"type:varchar(200);not_null"`
	Username string `gorm:"type:varchar(200);not_null"`
	Password string `gorm:"type:varchar(200);not_null"`
}

func (model *User) TableName() string {
	return UserTableName
}
