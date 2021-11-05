package controller

import (
	"github.com/labstack/echo/v4"
	"kembanglah/model/web"
	"net/http"
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
