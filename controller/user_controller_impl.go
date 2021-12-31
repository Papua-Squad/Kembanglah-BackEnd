package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"kembanglah/helper"
	"kembanglah/model/web"
	"kembanglah/service"
	"net/http"
	"strconv"
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

	userID, _ := strconv.Atoi(ctx.Param("userID"))
	user.ID = uint(userID)

	if err := helper.BindAndValidate(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse, err := controller.UserService.Update(ctx.Request().Context(), *user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update User",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) UpdatePassword(ctx echo.Context) error {
	user := new(web.UserUpdatePasswordRequest)

	userID, _ := strconv.Atoi(ctx.Param("userID"))
	user.ID = uint(userID)

	if err := helper.BindAndValidate(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse, err := controller.UserService.UpdatePassword(ctx.Request().Context(), *user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update Password User",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) UpdateImage(ctx echo.Context) error {
	user := new(web.UserUpdateImageRequest)
	file, err := ctx.FormFile("image")

	filename := "user_" + random.String(8, random.Alphabetic) + ".jpg"

	err = helper.SaveFile(filename, file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	user.ID = uint(userID)
	user.ImageUrl = "http://159.223.82.24:3000/files/" + filename

	if err := helper.BindAndValidate(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userResponse, err := controller.UserService.UpdateImage(ctx.Request().Context(), *user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success Update Image User",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	if err := controller.UserService.Delete(ctx.Request().Context(), uint(userID)); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Success delete user",
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
		Message: "Get User By ID Success",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) FindByUsername(ctx echo.Context) error {
	username := ctx.Param("username")

	userResponse, err := controller.UserService.FindByUsername(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get User By Username Success",
		Data:    userResponse,
	})
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) error {
	userResponse, err := controller.UserService.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, web.Response{
		Code:    http.StatusOK,
		Message: "Get All Users Success",
		Data:    userResponse,
	})
}
