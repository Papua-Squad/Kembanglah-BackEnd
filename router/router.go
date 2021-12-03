package router

import (
	"kembanglah/app"
	"kembanglah/controller"
	"kembanglah/helper"
	"kembanglah/repository"
	"kembanglah/service"
	"os"

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
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	categoryRepository := repository.NewCategoryRepository(server)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	server.Echo.GET("/", homeController.Home)

	// Authentication
	server.Echo.POST("/register", authController.Register)
	server.Echo.POST("/login", authController.Login)

	// Create Restricted Group, which mean user must log in before using an endpoint
	restricted := server.Echo.Group("/api")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &helper.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	// Files Endpoint
	restricted.GET("/files/*", homeController.Files)

	// Category Endpoint
	categoryEndpoint := restricted.Group("/category")
	categoryEndpoint.POST("/", categoryController.Create)
	categoryEndpoint.PUT("/:categoryID", categoryController.Update)
	categoryEndpoint.DELETE("/:categoryID", categoryController.Delete)
	categoryEndpoint.GET("/:categoryID", categoryController.FindByID)
	categoryEndpoint.GET("/", categoryController.FindAll)

	// Product Endpoint
	productEndpoint := restricted.Group("/product")
	productEndpoint.POST("/", productController.Create)
	productEndpoint.PUT("/:productID", productController.Update)
	productEndpoint.DELETE("/:productID", productController.Delete)
	productEndpoint.GET("/:productID", productController.FindByID)
	productEndpoint.GET("/seller/:sellerID", productController.FindBySeller)
	productEndpoint.GET("/category/:categoryID", productController.FindByCategory)
	productEndpoint.GET("/", productController.FindAll)

}
