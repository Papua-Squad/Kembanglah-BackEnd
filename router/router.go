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

	//seller
	sellerRepository := repository.NewSellerRepository(server)
	sellerService := service.NewSellerService(sellerRepository)
	sellerController := controller.NewSellerController(sellerService)

	server.Echo.POST("/api/seller", sellerController.Register)
	server.Echo.PUT("/api/seller/:sellerId", sellerController.Update)
	server.Echo.GET("/api/seller/:sellerId", sellerController.FindByID)
	server.Echo.GET("/api/seller/:username", sellerController.FindByUsername)
	server.Echo.GET("/api/seller/findAll", sellerController.FindAll)
	server.Echo.DELETE("/api/seller/:sellerId", sellerController.Delete)

	//customer
	customerRepository := repository.NewCustomerRepository(server)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	server.Echo.POST("/api/customer", customerController.Register)
	server.Echo.GET("/api/customer/:customerId", customerController.FindByID)
	server.Echo.PUT("/api/customer/:customerId", customerController.Update)
}
