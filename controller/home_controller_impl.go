package controller

import (
	"kembanglah/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeControllerImpl struct {
}

func NewHomeController() HomeController {
	return &HomeControllerImpl{}
}

func (controller *HomeControllerImpl) Home(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Server Running",
		Data:    nil,
	})
}

func (controller *HomeControllerImpl) Files(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "/files/*")
}
