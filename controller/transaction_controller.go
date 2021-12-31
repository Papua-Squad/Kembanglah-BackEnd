package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	Save(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindBySeller(ctx echo.Context) error
	FindByCustomer(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
