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

	server := app.NewServer(newConfig)

	// Migrate database
	err = server.DB.Migrator().AutoMigrate(
		&domain.Customer{},
		&domain.Seller{},
		&domain.Product{},
		&domain.Category{},
		&domain.Order{},
		&domain.Transaction{},
		&domain.Address{},
	)
	helper.PanicIfError(err)

	router.NewRouter(server)
	err = server.Start(newConfig.Server.Port)
	helper.PanicIfError(err)

}
