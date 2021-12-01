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

	//user
	authRepository := repository.NewUserRepository(server)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	server.Echo.POST("/api/user", authController.Register)

	//product
	productRepository := repository.NewProductRepository(server)
	productService := service.NewProductService(&productRepository)
	productController := controller.NewProductController(productService)

	server.Echo.POST("/api/product", productController.Create)

}
