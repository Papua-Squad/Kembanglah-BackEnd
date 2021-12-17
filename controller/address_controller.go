package controller

import "github.com/labstack/echo/v4"

type AddressController interface {
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindByUser(ctx echo.Context) error
}
