package controller

import "github.com/labstack/echo/v4"

type CustomerController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Update(ctx echo.Context) error
}
