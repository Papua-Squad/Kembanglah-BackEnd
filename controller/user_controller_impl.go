package controller

import (
	"github.com/labstack/echo/v4"
	"kembanglah/app"
)

type UserControllerImpl struct {
	Server app.Server
}

func NewUserController(server app.Server) *UserControllerImpl {
	return &UserControllerImpl{
		server,
	}
}

func (u *UserControllerImpl) Login(ctx echo.Context) error {
	panic("implement me")
}

func (u *UserControllerImpl) Register(ctx echo.Context) error {
	panic("implement me")
}
