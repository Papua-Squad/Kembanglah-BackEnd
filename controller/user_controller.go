package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}
