package controller

import "github.com/labstack/echo/v4"

type HomeController interface {
	Home(ctx echo.Context) error
}
