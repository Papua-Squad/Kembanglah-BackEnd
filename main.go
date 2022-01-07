package main

import (
	"kembanglah/app"
	"kembanglah/config"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/router"
)

func main() {
	newConfig, err := config.NewConfig()
	helper.PanicIfError(err)

	server := app.NewServer(newConfig, false)

	// Migrate database
	err = server.DB.Migrator().AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Category{},
		&domain.Address{},
		&domain.Transaction{},
	)
	helper.PanicIfError(err)

	router.NewRouter(server)
	err = server.Start(newConfig.Server.Port)
	helper.PanicIfError(err)

}
