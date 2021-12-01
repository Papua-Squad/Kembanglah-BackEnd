package controller

import (
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(ctx echo.Context) error {
	panic("implement me")
}

func (controller *AuthControllerImpl) Register(ctx echo.Context) error {

	user := new(web.RegisterRequest)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse := controller.AuthService.Register(ctx.Request().Context(), *user)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "success create user",
		Data:    userResponse,
	})
}
