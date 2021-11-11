package router

import (
	"kembanglah/app"
	"kembanglah/controller"
	"kembanglah/helper"
	"kembanglah/repository"
	"kembanglah/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *app.Server) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Validator = &helper.CustomValidator{Validator: validator.New()}

	homeController := controller.NewHomeController()
	server.Echo.GET("/", homeController.Home)

	//User
	userRepository := repository.NewUserRepository(server)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	server.Echo.POST("/user", userController.Register)

}
