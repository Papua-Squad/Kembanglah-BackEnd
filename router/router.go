package router

import (
	"kembanglah/app"
	"kembanglah/controller"
	"kembanglah/helper"
	"kembanglah/model/domain"
	"kembanglah/repository"
	"kembanglah/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *app.Server) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Validator = &helper.CustomValidator{Validator: validator.New()}

	// Initialize
	homeController := controller.NewHomeController()

	authRepository := repository.NewUserRepository(server)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	productRepository := repository.NewProductRepository(server)
	productService := service.NewProductService(&productRepository)
	productController := controller.NewProductController(productService)

	server.Echo.GET("/", homeController.Home)

	// Authentication
	server.Echo.POST("/register", authController.Register)
	server.Echo.POST("/login", authController.Login)

	jwtConfig := middleware.JWTConfig{
		Claims:     &domain.JwtCustomClaims{},
		SigningKey: []byte(server.Config.JwtSecret),
	}

	// Create Restricted Group, which mean user must login before using an endpoint
	restricted := server.Echo.Group("/api")
	restricted.Use(middleware.JWTWithConfig(jwtConfig))

	// Endpoint untuk files
	restricted.GET("/files/*", homeController.Files)

	restricted.POST("/product", productController.Create)

}
