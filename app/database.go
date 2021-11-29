package app

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kembanglah/config"
	"kembanglah/helper"
)

func NewDatabase(config *config.Config) *gorm.DB {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Host,
		config.Database.Postgres.User,
		config.Database.Postgres.Password,
		config.Database.Postgres.DB,
		config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}
