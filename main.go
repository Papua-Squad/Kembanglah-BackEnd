package main

import (
	"kembanglah/app"
	"kembanglah/config"
	"kembanglah/helper"
	"kembanglah/router"
)

func main() {
	newConfig, err := config.NewConfig()
	helper.PanicIfError(err)

	server := app.NewServer(newConfig)
	router.NewRouter(server)
	err = server.Start(newConfig.Server.Port)
	helper.PanicIfError(err)

}
