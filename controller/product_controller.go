package controller

import "github.com/labstack/echo/v4"

type ProductController interface {
	Create(ctx echo.Context) error
}
