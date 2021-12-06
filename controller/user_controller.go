package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	FindByUsername(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
