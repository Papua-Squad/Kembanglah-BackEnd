package controller

import (
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u *UserControllerImpl) Login(ctx echo.Context) error {
	panic("implement me")
}

func (u *UserControllerImpl) Register(ctx echo.Context) error {
	user := new(web.UserCreateRequest)
	err := ctx.Bind(user)
	helper.PanicIfError(err)

	userResponse := u.UserService.Register(ctx.Request().Context(), *user)

	return ctx.JSON(http.StatusCreated, userResponse)
}
