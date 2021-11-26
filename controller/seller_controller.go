package controller

import "github.com/labstack/echo/v4"

type SellerController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	FindByUsername(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
