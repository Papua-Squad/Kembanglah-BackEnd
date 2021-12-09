package controller

import (
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"

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

func (controller *UserControllerImpl) Update(ctx echo.Context) error {
	user := new(web.UserUpdateRequest)

	userID, err := strconv.Atoi(ctx.Param("userID"))
	helper.PanicIfError(err)
	user.ID = uint(userID)

	if err := helper.BindAndValidate(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse := controller.UserService.Update(ctx.Request().Context(), *user)

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Succes Update UserResponse",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	helper.PanicIfError(err)

	controller.UserService.Delete(ctx.Request().Context(), uint(userID))
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "OK",
	})
}

func (controller *UserControllerImpl) FindByID(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	userResponse, err := controller.UserService.FindByID(ctx.Request().Context(), uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get UserResponse By ID Success",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) FindByUsername(ctx echo.Context) error {
	username := ctx.Param("username")

	userResponse := controller.UserService.FindByUsername(ctx.Request().Context(), string(username))
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get UserResponse By Username Success",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) error {
	userResponse := controller.UserService.FindAll(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get All UserResponse Success",
		Data:    userResponse,
	})
}
