package controller

import (
	"github.com/labstack/echo/v4"
)

type ProductController interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	FindBySeller(ctx echo.Context) error
	FindByCategory(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
