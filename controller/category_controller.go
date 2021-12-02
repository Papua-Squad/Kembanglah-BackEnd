package controller

import "github.com/labstack/echo/v4"

type CategoryController interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	FindByID(ctx echo.Context) error
}
