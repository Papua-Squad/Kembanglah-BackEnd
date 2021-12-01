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
	user := new(web.LoginRequest)
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

	loginResponse, err := controller.AuthService.Login(ctx.Request().Context(), *user)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, web.Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "success login",
		Data:    loginResponse,
	})
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

	registerResponse, err := controller.AuthService.Register(ctx.Request().Context(), *user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "success register",
		Data:    registerResponse,
	})
}
